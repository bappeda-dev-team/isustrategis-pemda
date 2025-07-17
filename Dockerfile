ARG GO_VERSION=1.23.0
FROM golang:${GO_VERSION}-alpine AS build

WORKDIR /app

# Install deps for building and flyway
RUN apk add --no-cache git curl unzip openjdk17

# Copy source code
COPY . .

# Build Go binary
RUN go build -o api main.go

# Install Flyway CLI (versi bisa kamu ganti)
ENV FLYWAY_VERSION=10.12.0
RUN curl -L "https://repo1.maven.org/maven2/org/flywaydb/flyway-commandline/${FLYWAY_VERSION}/flyway-commandline-${FLYWAY_VERSION}-linux-x64.tar.gz" \
    | tar zx -C /opt && \
    ln -s /opt/flyway-${FLYWAY_VERSION}/flyway /usr/local/bin/flyway

# ===== Final stage =====
FROM alpine:3.20

WORKDIR /app

# Install minimal runtime dependencies
RUN apk add --no-cache openjdk17 libc6-compat

# Copy binary and Flyway from build stage
COPY --from=build /app/api /app/api
COPY --from=build /usr/local/bin/flyway /usr/local/bin/flyway
COPY --from=build /opt /opt

# Copy migration scripts
COPY db/migrations /app/migrations

# Set env vars (override via docker run or compose)
ENV FLYWAY_LOCATIONS=filesystem:/app/migrations

# Set default DB connection (can override in docker-compose or env)
ENV DB_HOST=localhost
ENV DB_PORT=5432
ENV DB_NAME=yourdb
ENV DB_USER=youruser
ENV DB_PASSWORD=yourpass
ENV DB_JDBC_URL=jdbc:postgresql://${DB_HOST}:${DB_PORT}/${DB_NAME}

# Run migration then app
ENTRYPOINT sh -c 'flyway -url=$DB_JDBC_URL -user=$DB_USER -password=$DB_PASSWORD migrate && exec /app/api'
