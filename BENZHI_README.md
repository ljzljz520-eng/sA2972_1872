# sA2972_1872

Go project for module `parking-offline`.

## Standard commands

```bash
go build ./...
go test -count=1 ./...
```

## Run

```bash
go run ./cmd/parking-capi
go run ./cmd/parking-offline
```

## Docker validation

```bash
chmod +x build_benzhi_docker.sh
./build_benzhi_docker.sh my-go-task linux/arm64
./build_benzhi_docker.sh my-go-task linux/amd64
docker run -it my-go-task:latest
```

## Known initial failures

See `BUG_REPRO.md` for the exact command and output captured during packaging.
