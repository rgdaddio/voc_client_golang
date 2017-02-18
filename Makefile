CC      = gccgo
CFLAGS  = -g
RM      = rm -f
EXPORT  := export GOPATH=/usr/local/go/
APATH   := $(shell pwd)
ROOT_PATH := export GOROOT=$(APATH)/go
DWN_EXPORT := export GOPATH=$(APATH)/libs
BIN_PATH   := export PATH=$(GOROOT)/bin
BUILDER := $(CURDIR)/go/bin/go
PKG_DIR :=
RM_DIR  :=  /tmp

default: all

all: voc_client_go_app

voc_client_go_app: test_client.go
	rm -f go1.5.linux-amd64.tar*
	wget https://storage.googleapis.com/golang/go1.5.linux-amd64.tar.gz
	tar -xvzf go1.5.linux-amd64.tar.gz
	@echo $(DWN_EXPORT)
	@echo $(ROOT_PATH)
	$(ROOT_PATH)
	$(DWN_EXPORT)
	$(BIN_PATH)
	$(BUILDER) get golang.org/x/net/context
	$(BUILDER) get -u github.com/mattn/go-sqlite3
	rm -fr /tmp/go-build*
	rm -fr /home/rdaddio/output/*
	$(EXPORT)
	$(BUILDER)  build -work -x -compiler gccgo -gccgoflags -pthread github.com/mattn/go-sqlite3/.
	cp -fr /tmp/go-*/github.com/ /home/rdaddio/output/
	cp -fr /tmp/go-*/golang.org/ /home/rdaddio/output/
	gccgo -v -g -o voc_client_go_app test_client.go client_utilities.go -I /home/rdaddio/output /home/rdaddio/output/github.com/mattn/libgo-sqlite3.a /home/rdaddio/output/golang.org/x/net/libcontext.a -pthread -ldl  

clean veryclean:
	$(RM) voc_client_go_app

