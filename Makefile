CC      = gccgo
CFLAGS  = -g
RM      = rm -f
EXPORT  := export GOPATH=/usr/local/go/
BUILDER := /home/rdaddio/lgo/go/bin/go
PKG_DIR :=
RM_DIR  :=  /tmp
default: all

all: voc_client_go_app
voc_client_go_app: test_client.go
	rm -fr /tmp/go-build*
	rm -fr /home/rdaddio/output/*
	$(EXPORT)
	$(BUILDER)  build -work -x -compiler gccgo -gccgoflags -pthread github.com/mattn/go-sqlite3/.
	cp -fr /tmp/go-*/github.com/ /home/rdaddio/output/
	cp -fr /tmp/go-*/golang.org/ /home/rdaddio/output/
	gccgo -v -g -o voc_client_go_app test_client.go client_utilities.go -I /home/rdaddio/output /home/rdaddio/output/github.com/mattn/libgo-sqlite3.a /home/rdaddio/output/golang.org/x/net/libcontext.a -pthread -ldl  
#$(CC) $(CFLAGS) -o voc_client_go_app test_client.go client_utilities.go

clean veryclean:
	$(RM) voc_client_go_app

