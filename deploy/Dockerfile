FROM debian:bookworm-20240423
# debian 12.5

LABEL org.opencontainers.image.authors="Goudan Woo"

ARG TARGETOS
ARG TARGETARCH

COPY --chmod=0755 build/cross/${TARGETOS}/${TARGETARCH}/* /usr/bin/

ENTRYPOINT [ "poo-minter" ]
