package main

import (
	"log"
	"os"

	"github.com/gomodule/redigo/redis"
)

var (
	pool *redis.Pool
)


func initPool(address string) error {

	var err error

	pool = &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000,
		Dial: func() (redis.Conn, error) {
			conn, e := redis.Dial("tcp", address)
			if e != nil {
				log.Printf("ERROR: fail init redis pool: %s", err.Error())
				os.Exit(1)
			}
			err = e
			return conn, e
		},
	}

	return err
}

func OpenRedisConnection() {

	address := "localhost:6379"
	e := initPool(address)
	if e != nil {
		log.Fatal(Fail("REDIS  : service connection failed to start."))
	}
	log.Printf(Detail("REDIS  : Service connection started : Host=%s"), address)

}

func CloseRedisConnection() {

	e := pool.Close()
	if e != nil {
		log.Println(Fail("REDIS  : Service connection failed to close:", e))
	} else {
		log.Println(Detail("REDIS  : Service connection stopped."))
	}

}