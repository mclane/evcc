#
# STEP 1 build executable binary
#

FROM golang:1.14-buster as builder

WORKDIR /build

# cache modules
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN make clean install assets build

#
# STEP 2 build a small image including module support
#

FROM debian:buster-slim

WORKDIR /evcc

# Import from builder
# COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=builder /build/evcc /usr/local/bin/evcc

COPY bin/* /evcc/

# UI and /api
EXPOSE 7070/tcp
# KEBA charger
EXPOSE 7090/udp
# SMA Energy Manager
EXPOSE 9522/udp

ENTRYPOINT [ "/evcc/entrypoint.sh" ]
CMD [ "evcc" ]
