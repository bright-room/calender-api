FROM alpine/curl:latest AS builder

ENV MIGRATE_VERSION=v4.18.3
ARG TARGETARCH

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/${MIGRATE_VERSION}/migrate.linux-${TARGETARCH}.tar.gz | tar xvz

FROM alpine:latest AS generate-template

COPY --from=builder /migrate /usr/local/bin/migrate
COPY generate-template.sh /entrypoint.sh

RUN chmod +x /usr/local/bin/migrate && \
    chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]

FROM alpine:latest AS migration

COPY --from=builder /migrate /usr/local/bin/migrate
COPY migration.sh /entrypoint.sh

RUN chmod +x /usr/local/bin/migrate && \
    chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]
