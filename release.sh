rsrc -manifest ./qr-serve.manifest
go build -ldflags="-H windowsgui"
mkdir -p dist
mv qr-serve.exe dist/
