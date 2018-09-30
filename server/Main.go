/*
 * Copyright (c) 2018 Karim Kanso. All Rights Reserved.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package main

import (
	"bytes"
	"flag"
	"log"
	"net/http"
	"strconv"
	"strings"
)

const version string = "0.1"

var port = flag.Uint("port", 8080, "Port to run web server on.")

func baseHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received for", r.URL.Path)
	page := strings.TrimPrefix(r.URL.Path, "/")
	var buff bytes.Buffer
	f := func(s string) {
		buff.Write([]byte(s))
	}
	f("Hello ")
	f(r.RemoteAddr)
	f(" you requested ")
	f(r.Host)
	f(" for page ")
	f(page)
	f("\n")
	w.Header().Add("Content-Type", "text/plain")
	w.Write(buff.Bytes())
}

func main() {
	flag.Parse()
	log.Println("Starting simple server! version", version)
	http.HandleFunc("/", baseHandler)
	port := strconv.FormatUint(uint64(*port), 10)
	log.Println("Listening on port", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		panic(err)
	}
}
