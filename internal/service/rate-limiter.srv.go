package service

import (
	"sync"
	"time"

	"github.com/Cyber-cicco/jardin-pc/internal/dao"
	"github.com/Cyber-cicco/jardin-pc/internal/dto"
)

const (
    NO_BLOCK = iota
    TEMP_BLOCK
    PERMA_BLOCK
)

type AttemptMap struct {
    Map map[string]map[string][]*dto.RequestMachineInfos
    mu sync.RWMutex
}

var Attempts *AttemptMap

func NewAttemptMap() {
    attemptMap := &AttemptMap{
        Map : make(map[string]map[string][]*dto.RequestMachineInfos),
    }
    attemptMap.StartCleanupTask()
    Attempts = attemptMap
}

func (a *AttemptMap) AddEntry(login *dto.LoginDto, infos *dto.RequestMachineInfos) {
    infos.Date = time.Now()
    a.mu.Lock()
    defer a.mu.Unlock()
    //Check if the Ip has already made a failed attempt
    entry, ok := a.Map[infos.IpAdress]

    //If it ip doesn't already exist, append it
    if !ok {
        a.Map[infos.IpAdress] = make(map[string][]*dto.RequestMachineInfos)
        a.Map[infos.IpAdress][login.Email] = []*dto.RequestMachineInfos{infos} 
        return
    }
    // If it exists, check if it has already made an attempt for this user
    userEntry, ok := entry[login.Email]

    // If there is not yet an entry, add one
    if !ok {
        entry[login.Email] = []*dto.RequestMachineInfos{infos}
        return
    }

    // Else append it
    userEntry = append(userEntry, infos)
    entry[login.Email] = userEntry
}

func (a *AttemptMap) CheckIfBlocked(login *dto.LoginDto, infos *dto.RequestMachineInfos) int {

    if dao.CheckIpIsBlocked(infos.IpAdress) {
        return PERMA_BLOCK
    }

    a.mu.RLock()
    defer a.mu.RUnlock()

    block_status := a.checkIfBlockedForUser(login, infos)

    if block_status != NO_BLOCK {
        
        return block_status
    }

    return NO_BLOCK
}

func (a *AttemptMap) checkIfBlockedForUser(login *dto.LoginDto, currMachine *dto.RequestMachineInfos) int {
    entry, ok := a.Map[currMachine.IpAdress]

    // if no entry in map, return false
    if !ok {
        return NO_BLOCK
    }

    userEntry, ok := entry[login.Email]

    // if no entry for this user, return false
    if !ok {
        return NO_BLOCK
    }

    // If there is an entry, check its length.
    // If less than 5, return false
    if len(userEntry) < 5 {
        return NO_BLOCK
    }

    weight := 100

    if len(entry) > 5 {
        weight += 100
    }

    applyFilters := func(pastMachine *dto.RequestMachineInfos) {

        if pastMachine.OS == currMachine.OS {
            weight += 20
        }

        // Discrimination pure et simple
        if pastMachine.IsLinux() {
            weight += 5
        }

        if pastMachine.IsUnknown() {
            weight += 20
        }

        if pastMachine.Device == currMachine.Device {
            weight += 20
        }
    }

    for i, pastMachine := range userEntry {

        weight += 60
        applyFilters(pastMachine)

        if weight >= 500 {
            blocked := pastMachine.Date.After(time.Now().Add(-time.Minute * 5))

            if blocked {
                for _, otherMachine := range userEntry[i:] {
                    applyFilters(otherMachine)

                    if weight > 5000 {

                        if otherMachine.Date.After(time.Now().Add(-time.Hour)) {
                            dao.BlockIp(currMachine.IpAdress)
                            return PERMA_BLOCK
                        }
                    }
                }
                return TEMP_BLOCK
            }
        }

    }
    return NO_BLOCK
    
}

func (a *AttemptMap) StartCleanupTask() {
    ticker := time.NewTicker(10 * time.Minute)  // Run every 10 minutes
    
    go func() {
        for range ticker.C {
            a.cleanupOldAttempts()
        }
    }()
}

func (a *AttemptMap) cleanupOldAttempts() {
    a.mu.Lock()
    defer a.mu.Unlock()
    
    oneHourAgo := time.Now().Add(-1 * time.Hour)
    
    for ip, userMap := range a.Map {
        for email, attempts := range userMap {
            // Keep only recent attempts
            validAttempts := attempts[:0]
            for _, attempt := range attempts {
                if attempt.Date.After(oneHourAgo) {
                    validAttempts = append(validAttempts, attempt)
                }
            }
            if len(validAttempts) == 0 {
                delete(userMap, email)
            } else {
                userMap[email] = validAttempts
            }
        }
        if len(userMap) == 0 {
            delete(a.Map, ip)
        }
    }
}


