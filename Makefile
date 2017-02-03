CC      = gccgo
CFLAGS  = -g
RM      = rm -f

default: all

all: voc_client_go_app
voc_client_go_app: test_client.go
	$(CC) $(CFLAGS) -o voc_client_go_app test_client.go

clean veryclean:
	$(RM) voc_client_go_app

