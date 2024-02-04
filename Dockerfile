FROM golang:latest

WORKDIR /src

COPY . .

RUN chmod +x build.sh
RUN ./build.sh release
COPY out/pearbutter.toml out

ENTRYPOINT ["./out/pearbutter"]