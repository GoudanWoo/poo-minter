FROM alpine:3.19.0

LABEL org.opencontainers.image.authors="Goudan Woo"

ARG TARGETOS
ARG TARGETARCH

COPY --chmod=0755 build/cross/${TARGETOS}/${TARGETARCH}/* /usr/bin/

ENTRYPOINT [ "poo-minter" ]
