# SimpleHTTP

**Very simple** client server written in Go for testing connections
from the command line. Ideal for placing into Docker images and
running in GNS3 for testing a firewalls `HTTP` configuration.

The server writes out a single line of text for each request that
details the client and page requested.

The client just performs a `HTTP GET` to the server address with a
unique request code that the server will echo back. Further, the
client will check the unique code echo'ed back is the same as it
expected.

The two `Dockerfile`s provided have been tested with GNS3 and built
over a recent `alpine` image. The set of network tools include are
similar to the `gns3/ipterm` image. To change the server port in GNS3,
set the environment variable `PORT` on the server container, there is
no need to expose ports.

To build the Docker images, the following commands provide a guide:
```bash
docker build -t simple-server:latest server
docker build -t simple-client:latest client
```

The Dockerfiles will install temporarily install `golang` to compile
the client/server. It is suggested that quick tests are perfomed
outside Docker to avoid this.

Copyright (c) 2018 Karim Kanso. All Rights Reserved.
