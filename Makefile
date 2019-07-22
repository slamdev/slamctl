APPS := $(shell find cmd -mindepth 1 -maxdepth 1 -type d | xargs -n1 basename) web
BUILD_TARGETS := $(foreach m,$(APPS),build/$(m))

build: $(BUILD_TARGETS)

build/%:
	@echo "Building $*"
	go build -o bin/$* ./cmd/$*

build/daemon: build/web pack-assets

build/web:
	@echo "Building web"
	cd web && npm run build

pack-assets: build/web
ifeq ($(shell which packr2),)
	go install github.com/gobuffalo/packr/v2/packr2
endif
	cd internal && packr2

clean:
	cd internal && packr2 clean
	rm -rf bin/ web/dist
