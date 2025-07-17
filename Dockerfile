ARG GO_VERSION=1.23.0
FROM golang:${GO_VERSION}-alpine AS build

WORKDIR /app

# Install build & Flyway deps
RUN apk add --no-cache git curl unzip openjdk17

# Copy source
COPY . .

# Build Go binary
RUN go build -o api main.go

# Install Flyway
ENV FLYWAY_VERSION=11.8.2
RUN curl -L "https://repo1.maven.org/maven2/org/flywaydb/flyway-commandline/${FLYWAY_VERSION}/flyway-commandline-${FLYWAY_VERSION}-linux-x64.tar.gz" \
  | tar zx -C /opt && \
  ln -s /opt/flyway-${FLYWAY_VERSION}/flyway /usr/local/bin/flyway

# ===== Final image =====
FROM alpine:3.20

WORKDIR /app

# Install runtime deps
RUN apk add --no-cache openjdk17 libc6-compat

# Copy binary & Flyway
COPY --from=build /app/api /app/api
COPY --from=build /usr/local/bin/flyway /usr/local/bin/flyway
COPY --from=build /opt /opt

# Copy migration scripts
COPY db/migrations /app/migrations

# Set default ENV (optional, override via `-e`)
ENV DB_HOST=localhost
ENV DB_PORT=3306
ENV DB_NAME=yourdb
ENV DB_USER=youruser
ENV DB_PASSWORD=yourpass
ENV FLYWAY_LOCATIONS=filesystem:/app/migrations

# ENTRYPOINT: hitung URL saat runtime
ENTRYPOINT sh -c 'flyway -url=jdbc:mysql://$DB_HOST:$DB_PORT/$DB_NAME -user=$DB_USER -password=$DB_PASSWORD migrate && exec /app/api'
