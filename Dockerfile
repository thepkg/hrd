FROM cn007b/go:1.19-alpine AS build

WORKDIR  /go/src/github.com/cn007b/hrd
COPY . .
RUN make build

FROM cn007b/alpine:3.16
RUN apk add --no-cache ca-certificates
COPY --from=build /go/src/github.com/cn007b/hrd/hrd /usr/local/bin/

CMD ["/usr/local/bin/hrd"]
