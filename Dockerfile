FROM golang:alpine AS builder

WORKDIR /build
COPY . .
RUN apk add --update make
RUN make

FROM alpine:latest
WORKDIR /app
COPY --from=builder /build/bin/xcpng-csi /app

CMD ["./xcpng-csi"]
