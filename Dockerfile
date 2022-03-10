# FROM archlinux

# WORKDIR /app

# COPY ./app/bitme ./

# ENV PORT 8000
# EXPOSE 8000

# CMD ["./bitme"]

FROM golang:latest AS build

WORKDIR /bitme

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN make build

FROM scratch

WORKDIR /bitme

COPY --from=build /bitme/app/bitme /bitme/bitme
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Europe/Moscow

ENV PORT 8000
EXPOSE 8000

CMD ["/bitme/bitme"]
