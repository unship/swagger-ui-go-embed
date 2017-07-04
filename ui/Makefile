GOBINDATA_TARGET = embedded.go

.DEFAULT_GOAL := all
.PHONY: all
all: generate

.PHONY: generate
generate: $(GOBINDATA_TARGET)

.PHONY: local
local:
	rm -rf swagger-ui-local
	cp -r swagger-ui swagger-ui-local
	rm swagger-ui-local/.git
	cd swagger-ui-local && \
    		rm -rf dist && \
    		npm install && \
    		webpack -p

.PHONY: $(GOBINDATA_TARGET)
$(GOBINDATA_TARGET):
	go-bindata -nometadata -pkg ui -o $@ -prefix swagger-ui-local/dist swagger-ui-local/dist/...

.PHONY: rm
rm:
	rm $(GOBINDATA_TARGET)