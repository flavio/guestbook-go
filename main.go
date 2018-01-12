package main

import (
	"fmt"
	"os"

	"github.com/flavio/guestbook-go/handlers"

	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	var redisHost, redisSocket string
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

		db := setupRedisClient(redisHost, redisPort, redisSocket)

		e := echo.New()
		e.File("/", "public/index.html")
		e.GET("/messages", handlers.GetMessages(db))
		e.PUT("/messages", handlers.PutMessage(db))
		e.DELETE("/messages/:index", handlers.DeleteMessage(db))

		err := e.Start(fmt.Sprintf(":%d", port))
		if err != nil {
			return cli.NewExitError(err, 1)
		}

		return nil
	}

	app.Run(os.Args)
}

func setupRedisClient(host string, port int, socket string) *redis.Client {
	options := redis.Options{
		Password: "", // no password set
		DB:       0,  // use default DB
	}

	if socket != "" {
		options.Network = "unix"
		options.Addr = socket
	} else {
		options.Addr = fmt.Sprintf("%s:%d", host, port)
	}

	return redis.NewClient(&options)
}
