package model

import (
	"context"
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/forum_server/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var ctx = context.Background()

// 时间从设置里拿
var buffTime = time.Second * time.Duration(config.GetConfig().BuffTime)

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

func addZsetBuff(key string, score int, data interface{}) {
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

func getZsetArr(key string, first int64, end int64) []string {
	rdb := newRdb()
	res, err := rdb.ZRevRange(ctx, key, first, end).Result()
	if err != nil {
		log.Println(err)
	}
	rdb.Expire(ctx, key, buffTime)
	return res
}

func getZsetCount(key string) int {
	rdb := newRdb()
	num, err := rdb.ZCard(ctx, key).Result()
	if err != nil {
		log.Println(err)
	}
	return int(num)
}

func delKey(key string) int {
	rdb := newRdb()
	num, err := rdb.Del(ctx, key).Result()
	if err != nil {
		log.Println(err)
	}
	return int(num)
}

func checkKey(key string) bool {
	rdb := newRdb()
	num, err := rdb.Exists(ctx, key).Result()
	if err != nil {
		log.Println(err)
	}
	if num > 0 {
		return true
	} else {
		return false
	}
}

func tranPost(buffArr []string) []ForumPost {
	res := make([]ForumPost, len(buffArr))
	for i := 0; i < len(buffArr); i++ {
		json.Unmarshal([]byte(buffArr[i]), &res[i])
	}
	return res
}

func setCount(key string, count int) {
	rdb := newRdb()
	rdb.Set(ctx, key, count, buffTime)
}

func getCount(key string) int {
	rdb := newRdb()
	buff, err := rdb.Get(ctx, key).Result()
	if err != nil {
		log.Println(err)
	}
	count, err := strconv.Atoi(buff)
	if err != nil {
		log.Println(err)
	}
	return count
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
