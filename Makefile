GOCMD = go
GOBUILD = $(GOCMD) build
GOINSTALL = $(GOCMD) install
GOTEST = $(GOCMD) test
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

test: black-box-test user-acceptance-test

black-box-test:
	cp admin.db admin_test.db
	$(GOTEST) -v -test.run=TestExampleBlackBox
	rm admin_test.db

user-acceptance-test:
	cp admin.db admin_test.db
	$(GOTEST) -v -test.run=TestExampleUserAcceptance
	rm admin_test.db
