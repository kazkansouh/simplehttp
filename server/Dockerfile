FROM alpine:3.8

RUN apk --no-cache add \
        curl \
        tcpdump \
        netcat-openbsd \
        dhclient \
        openssh-client \
        net-tools \
        iproute2

COPY Main.go /root/go/src/server/Main.go

RUN apk --no-cache add --virtual build-dependencies \
        go git musl-dev && \
        \
        go get server && \
        cp /root/go/bin/server /usr/bin/ && \
        rm -rf /root/go && \
        apk del build-dependencies

ENV PORT 8080

VOLUME ["/root"]

CMD ["sh", "-c", "/bin/server --port $PORT"]
