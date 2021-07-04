package internal

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

func RedisConnection() redis.Conn {
	const IPPort = "172.28.0.2:6379"
	// const IPPort = "127.0.0.1:6379"
	err := *new(error)
	rc, err := redis.Dial("tcp", IPPort)
	if err != nil {
		fmt.Println("db conn error")
		panic(err)
	}
	// conn = rc
	fmt.Println("db conn success")
	return rc
}

func RedisSet(key string, value string, rc redis.Conn) {
	rc.Do("SET", key, value)
}

func RedisSetTimeout(key string, value string, timeout int, rc redis.Conn) {
	rc.Do("SETEX", key, timeout, value)
}

func RedisCheckKey(key string, rc redis.Conn) bool {
	res, err := redis.Bool(rc.Do("EXISTS", key))
	if err != nil {
		fmt.Println(err)
		return false
	}
	return res
}

func RedisGet(key string, rc redis.Conn) string {
	s, err := redis.String(rc.Do("GET", key))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return s
}

func RedisDelete(key string, rc redis.Conn) {
	rc.Do("DEL", key)
}
