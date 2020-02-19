FROM golang:buster AS builder

WORKDIR /build
COPY . .
RUN apt update && apt install make
RUN make

FROM debian:buster
WORKDIR /app
COPY --from=builder /build/bin/xcpng-csi /app

CMD ["./xcpng-csi"]
