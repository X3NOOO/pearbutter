FROM golang:latest

WORKDIR /src

COPY . .

RUN mkdir -p out

RUN go mod download
RUN go build -o out/pearbutter github.com/X3NOOO/pearbutter

WORKDIR /pearbutter
RUN mv /src/out/* .
RUN rm -rf /src

ENTRYPOINT ["/pearbutter/pearbutter", "--config", "/pearbutter/pearbutter.toml"]


