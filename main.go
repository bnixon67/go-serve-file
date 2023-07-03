package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func serveFileHandler(file string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, file)
	}
}

func main() {
	addr := flag.String("addr", ":8080", "address (host:port) to listen on")
	flag.Parse()

	if len(flag.Args()) != 1 {
		fmt.Printf("usage: %s filename\n", os.Args[0])
		os.Exit(1)
	}

	file := flag.Arg(0)

	http.HandleFunc("/file", serveFileHandler(file))

	fmt.Printf("file %s available at %s/file\n", file, *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
		os.Exit(1)
	}
}
