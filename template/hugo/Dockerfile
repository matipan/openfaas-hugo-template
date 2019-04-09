FROM matipan/gocv-alpine:4.0.1-buildstage as build-stage

# Add the watchdog
RUN apk upgrade --no-cache && apk --no-cache add curl \ 
    && echo "Pulling watchdog binary from Github." \
    && curl -sSL https://github.com/openfaas/faas/releases/download/0.8.10/fwatchdog > /usr/bin/fwatchdog \
    && chmod +x /usr/bin/fwatchdog \
    && apk del curl --no-cache

WORKDIR /go/src/handler
COPY . .

RUN go build -o handler .

FROM matipan/gocv-alpine:4.0.1-runtime

RUN echo -e '@edgunity http://nl.alpinelinux.org/alpine/edge/community\n\
@edge http://nl.alpinelinux.org/alpine/edge/main\n\
@testing http://nl.alpinelinux.org/alpine/edge/testing\n\
@community http://dl-cdn.alpinelinux.org/alpine/edge/community'\
  >> /etc/apk/repositories

RUN apk --no-cache add ca-certificates
# libstdc++ libjpeg libtbb@testing libpng jasper-libs tiff openblas libwebp

# Add non root user
RUN addgroup -S app && adduser -S -g app app
RUN mkdir -p /home/app
RUN chown app /home/app

WORKDIR /home/app
COPY --from=build-stage /go/src/handler/cascades   ./cascades
COPY --from=build-stage /go/src/handler/models   ./models
COPY --from=build-stage /go/src/handler/handler    .
COPY --from=build-stage /usr/bin/fwatchdog         .

RUN chown -R app ./cascades

USER app

ENV fprocess="./handler"

CMD ["./fwatchdog"]
