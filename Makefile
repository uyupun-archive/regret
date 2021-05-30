setup:
	go get github.com/cespare/reflex

run:
	reflex -r '\.go|config.yml\z' -s -- sh -c 'go run cmd/main.go cmd/router.go'
