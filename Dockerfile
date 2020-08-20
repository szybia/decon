FROM golang:1.14.7-alpine AS builder

WORKDIR /app

COPY . .

#	Build for linux without CGO and strip binary
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags="-s -w" -o decon ./cmd/rest


#	Compress binary even further
FROM gruebel/upx:latest as upx

WORKDIR /app

COPY --from=builder /app/decon decon_original

RUN upx --best --lzma -o decon decon_original


FROM scratch

WORKDIR /app

COPY --from=upx /app/decon .

ENV GIN_MODE release

CMD [ "./decon" ]