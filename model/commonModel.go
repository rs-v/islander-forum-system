package model

import (
	"log"

	"github.com/forum_server/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func getDsn() string {
	conf := config.GetConfig()
	dsn := conf.UserName + ":" + conf.PassWord + "@tcp(" + conf.Ip + ")/" + conf.Database
	return dsn
}

func newDB() *gorm.DB {
	dsn := getDsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// func get(query string, res ...interface{}) {
// 	db := newDB()
// 	defer db.Close()
// 	rows, err := db.Query(query)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		rows.Scan(res...)
// 	}
// }

// func insert(query string, data ...interface{}) int {
// 	db := newDB()
// 	defer db.Close()
// 	res, err := db.Exec(query, data...)
// 	if err != nil {
// 		panic(err)
// 	}
// 	lastId, err := res.LastInsertId()
// 	return int(lastId)
// }
