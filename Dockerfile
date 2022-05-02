FROM golang:buster as builder
WORKDIR /go/src/noise
COPY . /go/src/noise
RUN CGO_ENABLED=0 GOOS=linux go build -o noise *.go

FROM busybox:latest
COPY --from=builder /go/src/noise/noise /bin/noise
RUN chmod +x /bin/noise
ENV GOGC=off
ENTRYPOINT [ "/bin/noise" ]
