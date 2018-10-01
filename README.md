# SimpleHTTP

**Very simple** client server written in Go for testing connections
from the command line. Ideal for placing into Docker images and
running in GNS3 for testing a firewalls `HTTP` configuration.

The server writes out a single line of text for each request that
details the client and page requested.

The client just performs a `HTTP GET` to the server address with a
unique request code that the server will echo back.

The two `Dockerfile`s provided have been tested with GNS3 and built
over the `gns3/ipterm:latest` image. To change the server port in
GNS3, set the environment variable `PORT` on the server container,
there is no need to expose ports.

To build the Docker images, the following commands provide a guide:
```bash
cd server
go build -o output.bin Main.go
docker build -t simple-server:latest .

cd ../client
go build -o output.bin Main.go
docker build -t simple-client:latest .
```

The included `Makefile` attempts to run the above commands to build
the docker images along with some bookkeeping of old docker images.

Copyright (c) 2018 Karim Kanso. All Rights Reserved.
