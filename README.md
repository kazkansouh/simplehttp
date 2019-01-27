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

In the [appliance](appliance) directory there are two `gns3a` files to
simplify the import process.  To change the server port in GNS3, set
the environment variable `PORT` on the server container, there is no
need to expose ports.

Two `Dockerfile`s provided have been tested with GNS3 and built over a
recent `alpine` image. The set of network tools included are similar
to the `gns3/ipterm` image, this was needed as the `gns3/ipterm`
(which the first version of this code used) is based off Debian
Jessie, hence many of the software packages in the repo are also
becoming dated.

To build the Docker images, the following commands provide a guide:
```bash
docker build -t simple-server:latest . -f server/Dockerfile
docker build -t simple-client:latest . -f client/Dockerfile
```

The Dockerfiles will install temporarily install `golang` to compile
the client/server. It is suggested that quick tests are perfomed
outside Docker to avoid this.

Copyright (c) 2019 Karim Kanso. All Rights Reserved.
