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
	"flag"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	host  = flag.String("host", "localhost", "Remote host to connect to.")
	port  = flag.Uint("port", 8080, "Port of remote host.")
	id    = uuid.New().String()
	reqid = flag.String("requestid", id, "The page to request from server, default is a random uuid.")
	level = flag.Bool("verbose", false, "Display detailed information.")
)

func main() {
	flag.Parse()
	request := "http://" + *host + ":" + strconv.FormatUint(uint64(*port), 10) + "/" + *reqid
	if *level {
		log.Println("Requesting:", request)
	}
	if resp, err := http.Get(request); err != nil {
		log.Println("Failed to make request:", err)
	} else {
		if resp.StatusCode != 200 {
			log.Println("Status code != 200:", resp.StatusCode)
			return
		}
		buff := make([]byte, 1024)
		_, err := resp.Body.Read(buff)
		if err != nil && err != io.EOF {
			log.Println("Failed to read response body:", err)
			return
		}
		response := string(buff)
		if *level {
			log.Println("Response:", string(buff))
		}
		words := strings.SplitN(response, " ", 8)
		if len(words) != 8 {
			log.Println("Invalid respose payload")
			return
		}
		r := strings.SplitN(words[7], "\n", 2)[0]
		if r != *reqid {
			log.Println("Request id does not match:", *reqid, "(expected)", r, "(received)")
			return
		}
		log.Println("Test OK!")
	}
}
