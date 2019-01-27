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
	e "encoding/json"
	"flag"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func new_id() string {
	host, err := os.Hostname()
	if err == nil {
		return host + "-" + uuid.New().String()
	}
	return uuid.New().String()
}

var (
	host  = flag.String("host", "localhost", "Remote host to connect to.")
	port  = flag.Uint("port", 8080, "Port of remote host.")
	reqid = flag.String("requestid", new_id(), "The page to request from server, default is a random uuid.")
	level = flag.Bool("verbose", false, "Display detailed information.")
)

type response struct {
	Client string `json:"client"`
	Server string `json:"server"`
	Page   string `json:"page"`
}

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
		n, err := resp.Body.Read(buff)
		if err != nil && err != io.EOF {
			log.Println("Failed to read response body:", err)
			return
		}
		var r response
		if err := e.Unmarshal(buff[:n], &r); err != nil {
			log.Println("Failed to parse result from server:", err)
			log.Println("Returned data:", string(buff))
			return
		}

		if *level {
			log.Printf("Response parsed ok: %+v\n", r)
		}

		if r.Page != *reqid {
			log.Println("Request id does not match:", *reqid, "(expected)", r.Page, "(received)")
			return
		}
		log.Println("Test OK!")
	}
}
