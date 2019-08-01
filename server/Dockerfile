FROM golang:1.10.8 AS build

WORKDIR $GOPATH/src/github.com/matipan/server

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build --ldflags "-s -w" -a -installsuffix cgo -o server

FROM alpine:3.10 AS runtime

# Add non root user and certs
RUN apk --no-cache add ca-certificates \
    && addgroup -S server && adduser -S -g server server \
    && mkdir -p /home/server \
    && chown server /home/server

WORKDIR /home/server

COPY --from=build /go/src/github.com/matipan/server/server .

ENV PUBLISH="public"
ENV PORT=8080

ENTRYPOINT ./server --publish=$PUBLISH --port=$PORT