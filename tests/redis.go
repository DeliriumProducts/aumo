package tests

import (
	"log"
	"os"
	"strconv"

	"github.com/gomodule/redigo/redis"
	"github.com/joho/godotenv"
)

func SetupRedis() (redis.Conn, error) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Println(".env file not found, reading directly from env variables")
	}

	RedisURL := os.Getenv("REDIS_URL_TEST")
	RedisDatabase := os.Getenv("REDIS_DATABASE_TEST")
	dbN, err := strconv.Atoi(RedisDatabase)
	if err != nil {
		log.Fatalln("Couldn't parse REDIS_DATABASE_TEST as an int: ", err)
	}

	conn, err := redis.DialURL(RedisURL, redis.DialDatabase(dbN))
	if err != nil {
		log.Fatalln("Couldn't establish a connection: ", err)
	}

	defer conn.Close()

	TidyRedis(conn)

	return conn, nil
}

func TidyRedis(r redis.Conn) {
	_, err := r.Do("FLUSHDB")
	if err != nil {
		log.Fatalln("Couldn't FLUSHDB on Redis database")
	}
}
