FROM golang:1.20-alpine

LABEL maintainer="araitkozha"
WORKDIR /
COPY . .
RUN go build -o ascii-art-web
EXPOSE 8080
CMD ["./ascii-art-web"]