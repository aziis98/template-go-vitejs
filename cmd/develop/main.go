package main

import (
	"bufio"
	"io"
	"log"
	"os/exec"

	"aziis98.com/template-go-vitejs/backend/database"
	"aziis98.com/template-go-vitejs/backend/server"
	"golang.org/x/sync/errgroup"
)

func main() {
	srv := server.New(server.Dependencies{
		DB: &database.Mem{},
	})

	g := &errgroup.Group{}

	g.Go(func() error {
		log.Printf(`[develop] listening on port :4000...`)
		return srv.Listen(":4000")
	})
	g.Go(func() error {
		log.Printf(`[develop] starting vitejs development server...`)

		r, w := io.Pipe()
		cmd := exec.Command("npm", "run", "dev")
		cmd.Stdout = w
		stdout := bufio.NewScanner(r)

		go func() {
			for stdout.Scan() {
				log.Printf(`[develop] [vitejs] %s`, stdout.Text())
			}
		}()

		return cmd.Run()
	})

	log.Fatal(g.Wait())
}
