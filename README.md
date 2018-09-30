# SimpleHTTP

**Very simple** client server written in Go for testing connections
from the command line. Ideal for placing into Docker images and
running in GNS3.

The server writes out a single line of text for each request that
details the client and page requested.

The client just performs a http get to the server address with a
unique request code that the server will echo back.

The two Dockerfiles provided have been tested with GNS3 and built over
the `gns3/ipterm:latest` image. To change the server port in GNS3, set
the environment variable `PORT` on the server container, there is no
need to expose ports.

To build the Docker images, the following commands provide a guide:
```bash
cd server
go build -o server Main.go
docker build -t simple-server:latest .

cd ../client
go build -o client Main.go
docker build -t simple-client:latest .
```

Copyright (c) 2018 Karim Kanso. All Rights Reserved.
