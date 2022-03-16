package model

import (
	"context"
	"encoding/json"
	"log"

	"github.com/forum_server/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ctx = context.Background()

func getDsn() string {
	conf := config.GetConfig()
	dsn := conf.UserName + ":" + conf.PassWord + "@tcp(" + conf.Ip + ")/" + conf.Database + "?charset=utf8&timeout=30s"
	return dsn
}

func newRdb() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.GetConfig().RedisIp,
		Password: "",
		DB:       0,
	})
	return rdb
}

func newDB() *gorm.DB {
	dsn := getDsn()
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func AddZsetBuff(key string, score int, data interface{}) {
	value, _ := json.Marshal(data)
	rdb := newRdb()
	err := rdb.ZAdd(ctx, key, &redis.Z{
		Score:  float64(score),
		Member: value,
	}).Err()
	if err != nil {
		log.Println(err)
	}
}

func GetZsetArr(key string, first int64, end int64) []string {
	rdb := newRdb()
	res, err := rdb.ZRevRange(ctx, "post", first, end).Result()
	if err != nil {
		log.Println(err)
	}
	return res
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
