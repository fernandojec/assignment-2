FROM golang:alpine as builder

WORKDIR /app

COPY . .

RUN go build -o app-web cmd/web/main.go

FROM progrium/busybox
RUN  opkg-install ca-certificates

WORKDIR /app

COPY --from=builder /app/app-web .
COPY --from=builder /app/views /app/views

EXPOSE 443

CMD [ "./app-web" ]