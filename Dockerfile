FROM golang:latest AS build

WORKDIR /bitme

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN make build

FROM scratch

WORKDIR /bitme

COPY --from=build /bitme/app/bitme bitme/bitme
COPY --from=build /bitme/cmd/config/config.env bitme/config/config.env
COPY --from=build /bitme/cmd/error.log bitme/error.log
COPY --from=build /bitme/cmd/adminurl.db bitme/adminurl.db
COPY --from=build /bitme/cmd/data.db bitme/data.db
COPY --from=build /bitme/cmd/ip.db bitme/ip.db
COPY --from=build /bitme/cmd/shorturl.db bitme/shorturl.db
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Europe/Moscow

EXPOSE 8000

CMD ["./bitme"]
