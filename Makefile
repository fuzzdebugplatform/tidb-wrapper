

all: resource bin

resource:
	go generate ./resources
bin:
	go build


.PHONY: all resource bin