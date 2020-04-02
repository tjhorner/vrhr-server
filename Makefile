.PHONY: dist dist-win dist-macos dist-linux ensure-dist-dir build build-frontend install uninstall

GOBUILD=GO111MODULE=on packr2 build -ldflags="-s -w"
INSTALLPATH=/usr/local/bin

ensure-dist-dir:
	@- mkdir -p dist

build-frontend:
	cd frontend && npm run build

dist-win: ensure-dist-dir
	# Build for Windows x64
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o dist/vrhr-windows-amd64.exe *.go
	packr2 clean

dist-macos: ensure-dist-dir
	# Build for macOS x64
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o dist/vrhr-darwin-amd64 *.go
	packr2 clean

dist-linux: ensure-dist-dir
	# Build for Linux x64
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o dist/vrhr-linux-amd64 *.go
	packr2 clean

dist: build-frontend dist-win dist-macos dist-linux clean

clean:
	packr2 clean

build: build-frontend
	@- mkdir -p bin
	$(GOBUILD) -o bin/vrhr *.go
	make clean
	@- chmod +x bin/vrhr

install: build
	mv bin/vrhr $(INSTALLPATH)/vrhr
	@- rm -rf bin
	@echo "vrhr was installed to $(INSTALLPATH)/vrhr. Run make uninstall to get rid of it, or just remove the binary yourself."

uninstall:
	rm $(INSTALLPATH)/vrhr

run:
	@- go run *.go