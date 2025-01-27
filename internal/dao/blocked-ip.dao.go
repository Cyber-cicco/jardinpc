package dao

import (

	"github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/model"
	. "github.com/Cyber-cicco/jardin-pc/.gen/jardinpc/table"
	. "github.com/go-jet/jet/v2/mysql"
)

func CheckIpIsBlocked(ip string) bool {
    var res model.BannedIP
    stmt := SELECT(BannedIP.IP).FROM(BannedIP).WHERE(BannedIP.IP.EQ(String(ip)))
    return stmt.Query(db, &res) == nil
}

func BlockIp(ip string) {
    stmt := RawStatement(`
        INSERT INTO banned_ip(ip) VALUES(#ip)
        `, RawArgs{"#ip" : ip})
    stmt.Exec(db)
}

