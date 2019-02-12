package model

import (
	"constant"
	"math/rand"
	"time"

	"gopkg.in/mgo.v2/bson"
)

var defaultSelector = bson.M{}

func getredisDefaultExpire() int64 {
	rand.Seed(time.Now().UnixNano())
	return constant.RedisDefaultExpire + rand.Int63n(constant.RedisDefaultRandExpire)
}
