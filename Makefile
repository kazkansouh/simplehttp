default: simple-server simple-client

.PHONY: default clean

simple-%: %/output.bin %/Dockerfile
	if test ! -z "$$(docker images -q $@:latest)" ; then \
		docker tag $@:latest $@:old ; \
	fi
	docker build -t $@:latest $(subst simple-,,$@)/
	if test ! -z "$$(docker images -q $@:old)" ; then \
		docker rmi $@:old ; \
	fi

%/output.bin: %/Main.go
	go build -o $@ $<

clean:
	rm -f client/output.bin
	rm -f server/output.bin
	-docker rmi simple-client:latest
	-docker rmi simple-server:latest
