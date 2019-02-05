FROM alpine:3.8

RUN apk --no-cache add \
        curl \
        tcpdump \
        netcat-openbsd \
        openssh-client \
        net-tools \
        bind-tools \
        wpa_supplicant \
        iproute2

COPY client/Main.go /root/go/src/client/Main.go

RUN apk --no-cache add --virtual build-dependencies \
        go git musl-dev && \
        \
        go get client && \
        cp /root/go/bin/client /usr/bin/ && \
        rm -rf /root/go && \
        apk del build-dependencies

VOLUME ["/root"]
