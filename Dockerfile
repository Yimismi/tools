FROM golang:1.13
WORKDIR /go/src/tools

# Copy the directory contents into the container at /var/local
COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct

RUN go mod download

RUN sh ./build.sh

RUN sh ./install.sh

EXPOSE 8080

# Define environment variable
ENV NAME TOOLS
CMD ["sh", "/var/local/tools/bin/restart.sh"]