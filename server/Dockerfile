FROM alpine:3.8

RUN apk --no-cache add \
        curl \
        tcpdump \
        netcat-openbsd \
        openssh-client \
        net-tools \
        bind-tools \
        wpa_supplicant \
        busybox-extras \
        nmap \
        nmap-nselibs \
        nmap-scripts \
        nmap-ncat \
        nmap-nping \
        ip6tables \
        iproute2

RUN apk --no-cache add \
        ncurses \
        libcurl \
        libpcap \
        libnet \
        pcre && \
        mkdir /ettercap && \
        cd /ettercap && \
        curl -o ettercap.tar.gz -L https://github.com/Ettercap/ettercap/releases/download/v0.8.2/ettercap-0.8.2.tar.gz && \
        tar xf ettercap.tar.gz && \
        cd ettercap-0.8.2/ && \
        apk --no-cache add --virtual build-dependencies \
        cmake \
        make \
        gcc \
        musl-dev \
        ncurses-dev \
        curl-dev \
        libpcap-dev \
        libnet-dev \
        pcre-dev \
        flex \
        bison \
        linux-headers && \
        mkdir build && \
        cd build && \
        cmake -DENABLE_GTK=OFF -DCMAKE_INSTALL_PREFIX=/usr \
        -DINSTALL_DESKTOP=OFF ../ && \
        make && \
        make install && \
        cd / && \
        rm -rf /ettercap && \
        apk del build-dependencies

COPY server/Main.go /root/go/src/server/Main.go

RUN apk --no-cache add --virtual build-dependencies \
        go git musl-dev && \
        \
        go get server && \
        cp /root/go/bin/server /usr/bin/ && \
        rm -rf /root/go && \
        apk del build-dependencies

ENV PORT 8080

VOLUME ["/root"]

CMD ["sh", "-c", "/usr/bin/server --port $PORT"]
