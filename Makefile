GOCMD = go
GOBUILD = $(GOCMD) build
GOINSTALL = $(GOCMD) install
BINARY_NAME = goadmin
CLI = adm

all: serve

serve:
	$(GOCMD) run .

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./build/$(BINARY_NAME) -v ./

generate:
	$(GOINSTALL) github.com/GoAdminGroup/go-admin/adm
	$(CLI) generate -c adm_config.ini