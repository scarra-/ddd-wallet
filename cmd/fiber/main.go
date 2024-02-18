package main

import (
	"fmt"
	"os"
	"syscall"

	"github.com/aadejanovs/wallet/internal/infrastructure"
	"github.com/vrecan/death"
)

func main() {
	server := infrastructure.Setup()

	go func() {
		err := server.Listen("0.0.0.0:80")
		if err != nil {
			panic(err)
		}
	}()

	d := death.NewDeath(syscall.SIGINT, syscall.SIGTERM)

	d.WaitForDeathWithFunc(func() {
		fmt.Println("server_shutting_down")

		err := server.Shutdown()
		if err != nil {
			fmt.Printf("server_shutdown_error: %s\n", err)
		}
	})

	d.WaitForDeath()
	os.Exit(0)
}
