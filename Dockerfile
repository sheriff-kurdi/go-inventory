FROM golang:1.18-alpine AS builder

# Move to working directory (/build).
WORKDIR /build

# Copy the build into the container.
COPY build/. .

# Export necessary port.
EXPOSE 3000

# Command to run when starting the container.
ENTRYPOINT ["./inventory"]

