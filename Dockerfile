FROM alpine:latest
COPY ./build/hproxy /bin/hproxy
CMD ["/bin/hproxy"]
