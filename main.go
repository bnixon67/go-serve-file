/*
Copyright 2023 Bill Nixon

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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
