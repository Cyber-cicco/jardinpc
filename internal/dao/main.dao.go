package dao

import (
	"database/sql"
	"fmt"
	"sync"
    _ "github.com/go-sql-driver/mysql" 

	"github.com/Cyber-cicco/jardin-pc/internal/config"
)

var once sync.Once
var db *sql.DB

func InitDB() {

	once.Do(func() {
		_db, err := sql.Open(
			"mysql",
			fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", config.Conf.DBUser, config.Conf.DBPassword, config.Conf.DBAdress, config.Conf.DBName),
		)
		if err != nil {
			panic(err)
		}
		db = _db
	})
}

func BeginTransaction() *sql.Tx {
    tx, _ := db.Begin()
    return tx
}
