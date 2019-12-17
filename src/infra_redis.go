package main

import (
	"github.com/gomodule/redigo/redis"
	"log"
)

var (
	pool *redis.Pool
)

func StartRedisPool() {

	address := "localhost:6379"

	pool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, e := redis.Dial("tcp", address)
			if e != nil {
				log.Fatal(Fail("REDIS  : Service pool connection failed: ", e))
			}
			return conn, e
		},
	}
	log.Printf(Detail("REDIS  : Service connection started : Host=%s"), address)

}

func StopRedisPool() {

	e := pool.Close()
	if e != nil {
		log.Println(Fail("REDIS  : Service connection failed to close:", e))
	} else {
		log.Println(Detail("REDIS  : Service connection stopped."))
	}

}