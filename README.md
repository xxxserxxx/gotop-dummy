This repository is archived
===========================

Plugin development for gotop has been halted due to no good path forward. As
such, this demo package is deprecated.


go build -buildmode=plugin -o dummy.so .

export CC=aarch64-linux-gnu-gcc
GOARCH=arm64 CGO_ENABLED=1 go build --buildmode=plugin -o dummy-arm.so .
