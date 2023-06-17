# Small -> Image Size: 30 ~ 40 MB
# Use Alpine Linux Environment (Use Golang Environment in Build Stage)
# Build Stage
FROM golang:1.18 AS builder

WORKDIR /src

COPY . .

RUN CGO_ENABLED=0 go build -o a_ramen .

# Base Stage
FROM alpine:3.13

RUN apk add tzdata
RUN cp /usr/share/zoneinfo/Asia/Bangkok /etc/localtime
RUN echo "Asia/Bangkok" >  /etc/timezone

WORKDIR /usr/src/app

COPY --from=builder /src/a_ramen /usr/src/app/a_ramen

ADD .env .
ADD static/ static/


EXPOSE 8080
CMD ["./a_ramen"]

