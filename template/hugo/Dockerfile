FROM alpine:3.10 AS build

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

FROM matipan/static-server:0.0.1 AS runtime

WORKDIR /home/server

COPY --from=build /home/app/public .

ENV PUBLISH="public"
ENV PORT=8080

# The entrypoint is inheritated from the static server image