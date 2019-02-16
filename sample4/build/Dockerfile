FROM alpine:3.6 as app
RUN apk update \
    && apk upgrade \
    && apk add --no-cache \
        ca-certificates \
    && update-ca-certificates 2>/dev/null || true
ADD restapp /root/delivery/restapplication/restapp
#COPY config/devconfig.ini /root/config/devconfig.ini
#COPY config/prodconfig.ini /root/config/prodconfig.ini
EXPOSE 8011
WORKDIR /root/delivery/restapplication
ENV PATH="/root:${PATH}"
ENTRYPOINT [ "./restapp" ]