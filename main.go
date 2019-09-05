package main

import (
	"fmt"
	"os"

	"github.com/flavio/guestbook-go/handlers"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/urfave/cli"
)

func main() {
	var redisHost, redisSocket, redisPassword string
	var port, redisPort int

	app := cli.NewApp()
	app.Name = "guestbook-go"
	app.Usage = "simple web application guestbook"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "redis-unix-socket",
			Usage:       "Path to the unix socket `FILE` used by the Redis server",
			EnvVar:      "GUESTBOOK_REDIS_SOCKET",
			Destination: &redisSocket,
		},
		cli.StringFlag{
			Name:        "redis-host",
			Value:       "localhost",
			Usage:       "Redis `HOST`",
			EnvVar:      "GUESTBOOK_REDIS_HOST",
			Destination: &redisHost,
		},
		cli.IntFlag{
			Name:        "redis-port",
			Value:       6379,
			Usage:       "`PORT` used by the Redis server",
			EnvVar:      "GUESTBOOK_REDIS_PORT",
			Destination: &redisPort,
		},
		cli.StringFlag{
			Name:        "redis-password",
			Usage:       "`PASSWORD` used by the Redis server",
			EnvVar:      "GUESTBOOK_REDIS_PASSWORD",
			Destination: &redisPassword,
		},
		cli.IntFlag{
			Name:        "port, p",
			Value:       4000,
			Usage:       "Listen to `PORT`",
			EnvVar:      "GUESTBOOK_PORT",
			Destination: &port,
		},
	}

	app.Action = func(c *cli.Context) error {
		if redisSocket != "" && (redisHost != "" || redisPort != 0) {
			fmt.Println("Using redis socket")
		}

		fmt.Println("Connecting to redis server")
		db := setupRedisClient(redisHost, redisPort, redisSocket, redisPassword)
		pong, err := db.Ping().Result()
		fmt.Println(pong, err)

		e := echo.New()
		e.File("/", "public/index.html")
		e.GET("/messages", handlers.GetMessages(db))
		e.PUT("/messages", handlers.PutMessage(db))
		e.DELETE("/messages/:index", handlers.DeleteMessage(db))

		if err := e.Start(fmt.Sprintf(":%d", port)); err != nil {
			return cli.NewExitError(err, 1)
		}

		return nil
	}

	app.Run(os.Args)
}

func setupRedisClient(host string, port int, socket string, password string) *redis.Client {
	options := redis.Options{
		Password: password,
		DB:       0, // use default DB
	}

	if socket != "" {
		options.Network = "unix"
		options.Addr = socket
	} else {
		options.Addr = fmt.Sprintf("%s:%d", host, port)
	}

	return redis.NewClient(&options)
}
