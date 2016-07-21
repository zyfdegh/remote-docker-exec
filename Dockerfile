# Dockerfile for backend web-console

FROM alpine
Maintainer Zhang Yifa <yzhang3@linkernetworks.com>

WORKDIR /linker

# fix library dependencies
# otherwise golang binary may encounter 'not found' error
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2

COPY cert.pem /linker/cert.pem
COPY key.pem /linker/key.pem
COPY ca.pem /linker/ca.pem

COPY remote-docker-exec /linker/remote-docker-exec
COPY gotty /linker/gotty
RUN chmod +x /linker/remote-docker-exec && chmod +x /linker/gotty

EXPOSE 8080

CMD ["/linker/gotty","-w","bash"]
