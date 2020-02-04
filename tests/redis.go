package tests

import (
	"log"
	"os"
	"strconv"

	"github.com/go-redis/redis/v7"
	"github.com/joho/godotenv"
)

// SetupRedis creates a new Redis connection
func SetupRedis() (*redis.Client, error) {
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

	conn := redis.NewClient(&redis.Options{
		Addr: RedisURL,
		DB:   dbN,
	})

	err = conn.Ping().Err()
	if err != nil {
		log.Fatalln("Couldn't establish a connection: ", err)
	}

	TidyRedis(conn)

	return conn, nil
}

// TidRedis clears the Redis db
func TidyRedis(r *redis.Client) {
	err := r.FlushDB().Err()
	if err != nil {
		log.Fatalln("Couldn't FLUSHDB on Redis database: ", err)
	}
}
