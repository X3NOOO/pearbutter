FROM golang:latest as builder
LABEL stage=builder

WORKDIR /src

COPY . .

RUN mkdir -p out

RUN go mod download
RUN go build -o out/pearbutter github.com/X3NOOO/pearbutter

FROM scratch
WORKDIR /pearbutter
COPY --from=builder /src/out/* .

ENTRYPOINT ["/pearbutter/pearbutter", "--config", "/pearbutter/pearbutter.toml"]


