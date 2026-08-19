FROM --platform=$BUILDPLATFORM golang:1.26-alpine AS build
WORKDIR /src
COPY . .
ARG TARGETOS TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -o /field-archive ./cmd/field-archive
FROM alpine:3.22
RUN mkdir -p /data
COPY --from=build /field-archive /field-archive
EXPOSE 8080
ENTRYPOINT ["/field-archive"]
