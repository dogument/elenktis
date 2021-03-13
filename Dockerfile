FROM golang:1.16.2
WORKDIR /app
COPY . .
RUN make

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app/bin/elenktis .
EXPOSE 8080
CMD ["./elenktis"]