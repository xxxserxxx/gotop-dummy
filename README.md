go build -buildmode=plugin -o dummy.so .

export CC=aarch64-linux-gnu-gcc
GOARCH=arm64 CGO_ENABLED=1 go build --buildmode=plugin -o dummy-arm.so .
