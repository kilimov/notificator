# Build Path
BUILD_PATH=./cmd/

# This how we want to name the binary output
BINARY=./bin/notificator

# These are the values we want to pass for VERSION and BUILD
VERSION=`git describe --abbrev=6 --always --tag`

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS=-ldflags "-X main.version=${VERSION}"
GOFLAGS=-a -tags notificator -installsuffix notificator -mod=vendor

run:
	go run $(BUILD_PATH)/main.go

bin:
	echo "  >  Building binary \"notificator\" $(VERSION)..."
	go build $(GOFLAGS) $(LDFLAGS) -o $(BINARY) $(BUILD_PATH)

bin-linux:
	echo "  >  Building linux-amd64 binary \"notificator\" $(VERSION)..."
	GOOS=linux GOARCH=amd64 go build $(GOFLAGS) $(LDFLAGS) -o $(BINARY)-linux $(BUILD_PATH)

bin-windows:
	echo "  >  Building windows-amd64 binary \"notificator\" $(VERSION)..."
	GOOS=windows GOARCH=amd64 go build $(GOFLAGS) $(LDFLAGS) -o $(BINARY).exe $(BUILD_PATH)

bin-cross-platform: bin-linux bin-windows

build:
	$(MAKE) bin

build-cross-platform:
	$(MAKE) bin-cross-platform

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	if [ -f ${BINARY} ] ; then rm ${BINARY}-linux ; fi
	if [ -f ${BINARY} ] ; then rm ${BINARY}.exe ; fi
	if [ -f coverage.html ] ; then rm coverage.html ; fi
	if [ -d .cover ] ; then rm -rf .cover ; fi
	docker-compose down --rmi all -v 2>/dev/null || true
	docker-compose stop >/dev/null
	docker-compose rm >/dev/null

rebuild:
	docker-compose build notificator
	docker-compose build unit

unit:
	docker-compose run --rm unit

coverage:
	docker-compose run --rm unit && [ -f ./coverage.html ] && xdg-open coverage.html

swagger:
	swag init -g ./app/apiserver/start.go -o ./api -d ./internal

.DEFAULT_GOAL := run

.PHONY: all bin bin-linux bin-windows build clean coverage rebuild run unit swagger
