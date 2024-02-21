FROM golang:1.22 AS builder
COPY . .
RUN go build -tags netgo,osusergo -ldflags '-extldflags "-static"' -o /usr/local/bin/fly-autoscaler-metric ./cmd/fly-autoscaler-metric
RUN go build -tags netgo,osusergo -ldflags '-extldflags "-static"' -o /usr/local/bin/fly-autoscaler-worker ./cmd/fly-autoscaler-worker

FROM alpine AS metric
COPY --from=builder /usr/local/bin/fly-autoscaler-metric /usr/local/bin/fly-autoscaler-metric
CMD fly-autoscaler-metric

FROM alpine AS worker
COPY --from=builder /usr/local/bin/fly-autoscaler-worker /usr/local/bin/fly-autoscaler-worker
CMD fly-autoscaler-worker

