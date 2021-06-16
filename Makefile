all: yamlcfg

yamlcfg:
	go build -o bin/yamlcfg ./cmd/yamlcfg

.PHONY: yamlcfg