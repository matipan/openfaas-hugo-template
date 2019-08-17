FROM openfaas/of-watchdog:0.5.5 AS watchdog
FROM alpine:3.10 AS build

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
RUN chmod +x /usr/bin/fwatchdog

RUN apk update && apk add py-pygments ca-certificates bash git && rm -rf /var/cache/apk/*

ENV HUGO_VERSION 0.55.0
ENV HUGO_BINARY hugo_${HUGO_VERSION}_Linux-64bit
ENV HUGO_ENV=production

RUN mkdir /usr/local/hugo
ADD https://github.com/spf13/hugo/releases/download/v${HUGO_VERSION}/${HUGO_BINARY}.tar.gz /usr/local/hugo/
RUN tar xzf /usr/local/hugo/${HUGO_BINARY}.tar.gz -C /usr/local/hugo/ \
	&& ln -s /usr/local/hugo/hugo /usr/local/bin/hugo \
	&& rm /usr/local/hugo/${HUGO_BINARY}.tar.gz

WORKDIR /home/app
COPY ./function .

RUN ["hugo", "--gc", "--minify"]

FROM alpine:3.10 AS runtime

WORKDIR /home/app

COPY --from=build /home/app/public public
COPY --from=build /usr/bin/fwatchdog .

ENV mode="static"
ENV static_path="/home/app/public"

HEALTHCHECK --interval=3s CMD [ -e /tmp/.lock ] || exit 1

CMD ["./fwatchdog"]
