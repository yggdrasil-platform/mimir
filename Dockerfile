############################
# Image: Build executable binary.
############################
FROM golang:1.14.2-alpine AS builder

# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Create appuser.
ENV USER=appuser
ENV UID=10001
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"

# Copy the source files.
WORKDIR $GOPATH/src/server
COPY . .

# Download and verify dependencies.
RUN go mod download
RUN go mod verify

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags='-w -s -extldflags "-static"' -a \
      -o /go/bin/server ./cmd/server

############################
# Image: Execute the binary.
############################
FROM scratch

ARG env
ARG port
ARG service_name
ARG version

ENV ENV=$env
ENV PORT=$port
ENV SERVICE_NAME=$service_name
ENV VERSION=$version

# Import the user and group files from the builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable.
COPY --from=builder /go/bin/server /usr/app/server

# Use an unprivileged user.
USER appuser:appuser

EXPOSE ${PORT}

# Execute the binary.
ENTRYPOINT ["/usr/app/server"]
