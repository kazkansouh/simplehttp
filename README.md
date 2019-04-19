# SimpleHTTP

**Very simple** client server written in Go for testing connections
from the command line. Ideal for placing into Docker images and
running in [GNS3](www.gns3.com) for testing a firewalls `HTTP`
configuration.

The server writes out a single line of text for each request that
details the client and page requested.

The client just performs a `HTTP GET` to the server address with a
unique request code that the server will echo back. Further, the
client will check the unique code echo'ed back is the same as it
expected.

To install directly, follow standard `Go` procedure:
```bash
go get github.com/kazkansouh/simplehttp/server
go get github.com/kazkansouh/simplehttp/client
```

## Usage

```bash
user@hostname:~$ client -help
Usage of ./bin/client:
  -host string
        Remote host to connect to. (default "localhost")
  -port uint
        Port of remote host. (default 8080)
  -requestid string
        The page to request from server, default is a random uuid. (default "hostname-118e24a3-48c0-4b96-b266-01ab13c50c18")
  -verbose
        Display detailed information.

```

```bash
user@hostname:~$ server -help
Usage of ./bin/server:
  -port uint
        Port to run web server on. (default 8080)
```

While the server is running, it lists out each connection it receives,
notice it lists out the requested URL (when used with the client it
also lists the hostname of the client):

```bash
username@hostname:~$ server
2019/04/19 08:44:34 Starting simple server! version 0.1
2019/04/19 08:44:34 Listening on port 8080
2019/04/19 08:44:42 127.0.0.1:50720 requested /hosta-69f68c8e-a1ce-4b88-a0b2-3a149c8e60b5
2019/04/19 08:44:43 192.168.0.2:50732 requested /hostb-0cf6aed3-7b82-4a56-876a-e3608197b138
2019/04/19 08:44:44 127.0.0.1:50738 requested /hosta-b07b3587-a539-498f-8155-95fac18dce70
```

## GNS3 Docker image

The suggested method to deploy within GNS3 is to use the Docker images
available on DockerHub
([client](https://hub.docker.com/r/karimkanso/simple-client),
[server](https://hub.docker.com/r/karimkanso/simple-server)) and the
provided appliance files. In the
[appliance](https://github.com/kazkansouh/simplehttp/tree/master/appliances)
directory there are two `gns3a` files to simplify the import
process. Simply download the files and open with GNS3.  To change the
server port in GNS3, set the environment variable `PORT` on the server
container, there is no need to expose ports.

The Docker images are intended to provide a similar level of
functionality to the standard
[`gns3/ipterm`](https://github.com/GNS3/gns3-registry/tree/master/docker/ipterm)
image that is available with GNS3.

### Docker image building

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

### Advanced usage

This section summaries some non-canonical configurations supported by
the Docker image when used in GNS3.

#### 802.1X wpa_supplicant

To utilise 802.1X authentication within GNS3, configure the interface
`ifupdown` with a `pre-up` directive as follows:

```
auto eth0
iface eth0 inet static
    address 192.168.0.2
    netmask 255.255.255.0
    gateway 192.168.0.1
    pre-up wpa_supplicant -Dwired -ieth0 -c/root/wpa.conf &
```

Here the file `/root/wpa.conf` contains the credentials for the
supplicant and should be present in the Docker image. As the image
marks `/root` as persistent, the file can be manually created as
follows:

```bash
$ cat - > /root/wpa.conf << EOF
ctrl_interface=/var/run/wpa_supplicant
ctrl_interface_group=0
eapol_version=2
ap_scan=0
network={
        key_mgmt=IEEE8021X
        eap=PEAP
        identity="user1"
        password="cisco"
        phase2="autheap=MSCHAPV2"
        eapol_flags=0
}
EOF
```

Then stop/start the image (or use `ifdown`/`ifup`) to get the
supplicant working as desired.

More information about this configuration, including a Docker RADIUS
image and using Cisco L2 devices can be found
[here](https://hub.docker.com/r/karimkanso/gns3-freeradius).

#### ettercap

The image contains version 0.8.2 of
[ettercap](https://github.com/Ettercap/ettercap/tree/v0.8.2) built
from source (as its not available currently in Alpine packages) for
basic penetration testing.

#### ISC dhclient

The `:dhclient` tag (both in GitHub and DockerHub) provides the [ISC
dhclient](https://www.isc.org/downloads/dhcp/) installed into the
image. This provides a DHCP client that is vastly more configurable
than the client provided by `busybox` (that GNS3 injects into the
image to configure the network).

For example, it is not possible to use the `busybox` DHCP client to
configure the DHCP client identifier that is sent to the DHCP
server. To do this with `:dhclient` variant, configure the `interface`
file with the following to set the client identifier to `GNS3`:

```
auto eth0
iface eth0 inet dhcp
	pre-up echo "send dhcp-client-identifier 0:67:6e:73:33;"  > /etc/dhcp/dhclient.conf
```

However, this only works in GNS3 2.2+ that has the initialisation path
updated to include `/sbin` within the
[`init.sh`](https://github.com/GNS3/gns3-server/pull/1421/commits/14fb64b9411a411e582ebf558ceaec32e31ab404#diff-4506833306c4c2bd672c76e499b11245)
script so that it can see the `dhclient` installed in the image.

## Other bits

Copyright (c) 2019 Karim Kanso. All Rights Reserved.
