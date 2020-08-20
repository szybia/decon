FROM golang:1.14.7-alpine AS builder

WORKDIR /app

RUN apk add --no-cache make

COPY . .

RUN make build


#	Compress binary even further
FROM gruebel/upx:latest as upx

WORKDIR /app

COPY --from=builder /app/bin/rest rest_original

RUN upx --best --lzma -o rest rest_original


FROM scratch

WORKDIR /app

EXPOSE 80 443

COPY --from=upx /app/rest .

ENV GIN_MODE release
ENV DECON_ENDPOINT :80
ENV DECON_REDIS_ADDR redis:6397

CMD [ "./rest" ]