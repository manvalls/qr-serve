rsrc -manifest ./qr-serve.manifest -ico ./qr-serve.ico
go build -ldflags="-H windowsgui"
mkdir -p dist
mv qr-serve.exe dist/
