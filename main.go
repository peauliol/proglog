package main

import (
	"fmt"
	"log"
	"github.com/peauliol/proglog/internal/server"
)

func main() {
	fmt.Println("Starting service at port 8080")
	srv := server.NewHTTPServer(":8080")
	log.Fatal(srv.ListenAndServe())
}

