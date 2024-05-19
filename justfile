fmt:
    go fmt ./...

gen:
    cd tool/generator && go run .
    just fmt
