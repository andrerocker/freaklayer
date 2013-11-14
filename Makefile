bootstrap:
	GOPATH=~/.bricklayer go get "github.com/vaughan0/go-ini"

run:
	GOPATH=~/.bricklayer go run src/freaklayer.go --config etc/bricklayer/bricklayer.conf
