FROM alpine:3.7

RUN apk update && apk add py-pygments ca-certificates bash git && rm -rf /var/cache/apk/*

ENV HUGO_VERSION 0.55.0
ENV HUGO_BINARY hugo_${HUGO_VERSION}_Linux-64bit
ENV HUGO_ENV=production

RUN mkdir /usr/local/hugo
ADD https://github.com/spf13/hugo/releases/download/v${HUGO_VERSION}/${HUGO_BINARY}.tar.gz /usr/local/hugo/
RUN tar xzf /usr/local/hugo/${HUGO_BINARY}.tar.gz -C /usr/local/hugo/ \
	&& ln -s /usr/local/hugo/hugo /usr/local/bin/hugo \
	&& rm /usr/local/hugo/${HUGO_BINARY}.tar.gz

RUN addgroup -S web && adduser -S -g web web
RUN mkdir -p /home/web
RUN chown web /home/web

WORKDIR /home/web
COPY ./function .

EXPOSE 8080

ENTRYPOINT ["hugo", "server", "--minify", "-p", "8080", "--bind", "0.0.0.0", "--log"]