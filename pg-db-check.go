package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"time"

	"github.com/go-pg/pg/v10"
)

func main() {
	args := os.Args[1:]
	argsNum := len(args)

	tries := 0
	for {
    conn, _ := net.DialTimeout("tcp", args[0], 3 * time.Second)
    if conn != nil {
			conn.Close()
			break
    }
		if tries == 10 {
			os.Exit(1)
		}
		time.Sleep(1 * time.Second)
		tries += 1
	}

	if argsNum < 4 {
    fmt.Println("Too few arguments")
		os.Exit(2)
	}

	db := pg.Connect(&pg.Options{
		Addr: args[0],
		User: args[1],
		Password: args[2],
		Database: args[3],
	})
	defer db.Close()

	ctx := context.Background()

	if argsNum >= 5 {
		_, err := db.ExecContext(ctx, args[4])
		if err != nil {
			os.Exit(3)
		}
	}

	if err := db.Ping(ctx); err != nil {
		os.Exit(3)
	}
}
