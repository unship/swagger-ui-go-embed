GOBINDATA_TARGET = embedded.go

.DEFAULT_GOAL := all
.PHONY: all
all: generate

.PHONY: generate
generate: $(GOBINDATA_TARGET)

local:
	rm -rf swagger-ui-local
	cp -r swagger-ui swagger-ui-local
	rm swagger-ui-local/.git

.PHONY: $(GOBINDATA_TARGET)
$(GOBINDATA_TARGET):
#	cd swagger-ui && \
#		rm -rf dist && \
#		npm install && \
#		webpack -p
	go-bindata -nometadata -pkg ui -o $@ -prefix swagger-ui-local/dist swagger-ui-local/dist/...

rm:
	rm $(GOBINDATA_TARGET)