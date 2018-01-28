package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/danott/envflag"
)

var description = "Super Simple File Web Server"
var version = "0.1"

var (
	root string
	bind string
)

func init() {
	flag.StringVar(&root, "root", ".", "web root to serv files from")
	flag.StringVar(&bind, "bind", "127.0.0.1:8080", "address and port to bind")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "%s v%s\n\nUsage:\n", description, version)
		flag.PrintDefaults()
	}
}

func main() {
	envflag.Parse()

	log.Printf("Server binding to %s with root:%s\n", bind, root)
	log.Fatal(http.ListenAndServe(bind, http.FileServer(http.Dir(root))))
}
