export GO111MODULE := on
export PATH := bin:$(PATH)


VERSION       	       ?= $(shell git describe --tags --always --dirty --match="v*" 2> /dev/null || cat $(CURDIR)/.version 2> /dev/null || echo v1.0.0)
SCRIPTS_DIR            ?= ./scripts
K8S_DIR                ?= ./.k8s
K8S_BUILD_DIR          ?= ./.build_k8s
K8S_FILES              := $(shell find $(K8S_DIR) -name '*.yaml' | sed 's:$(K8S_DIR)/::g')
DOCKER_REGISTRY_DOMAIN ?= docker.io
DOCKER_REGISTRY_PATH   ?= zikes
DOCKER_IMAGE           ?= $(DOCKER_REGISTRY_PATH)/$(PACKAGE):$(VERSION)
DOCKER_IMAGE_DOMAIN    ?= $(DOCKER_REGISTRY_DOMAIN)/$(DOCKER_IMAGE)

GOCMD				   = go
GOLINT				   :=$(shell which golint)
GOIMPORT			   :=$(shell which goimports)
GOFMT				   :=$(shell which gofmt)
GOBUILD				   :=$(GOCMD) build
GOCLEAN				   :=$(GOCMD) clean
GOTEST				   :=$(GOCMD) test
GOMOD				   :=$(GOCMD) mod
GOLIST				   :=$(GOCMD) list
GOVET				   :=$(GOCMD) vet

GOFILES				   :=$(shell find . -name "*.go" -type f)
PACKAGES			   :=$(shell $(GOLIST) ./...)

GO      = go

MAKE_ENV += PACKAGE VERSION DOCKER_IMAGE DOCKER_IMAGE_DOMAIN

SHELL_EXPORT := $(foreach v,$(MAKE_ENV),$(v)='$($(v))' )


.PHONY: docs
docs:
	bash $(SCRIPTS_DIR)/0_create_docs.sh

.PHONY: build
build: docs
	bash $(SCRIPTS_DIR)/1_build_srv_images.sh

.PHONY: push-docker
push: build
	bash $(SCRIPTS_DIR)/2_push_images.sh

.PHONY: delete-images
delete-images:
	bash $(SCRIPTS_DIR)/7_delete_images.sh

.PHONY: create-gke
create-gke:
	bash $(SCRIPTS_DIR)/3_create_gke_cluster.sh

.PHONY: install-istio
install-istio:
	bash $(SCRIPTS_DIR)/4_install_istio.sh

# Builds the Kubernetes build directory if it does not exist
# The @ symbol prevents Make from echoing the results of the
# command.
$(K8S_BUILD_DIR):
	@mkdir -p $(K8S_BUILD_DIR)

.PHONY: build-k8s
build-k8s: $(K8S_BUILD_DIR)
	@for file in $(K8S_FILES); do \
		mkdir -p `dirname "$(K8S_BUILD_DIR)/$$file"` ; \
		$(SHELL_EXPORT) envsubst <$(K8S_DIR)/$$file >$(K8S_BUILD_DIR)/$$file ;\
	done

.PHONY: deploy
deploy: build-k8s push-docker
#	todo: this
	kubectl apply -f $(K8S_BUILD_DIR)


# GoLang Scripts

.PHONY: install
install:
	$(GOMOD) download

.PHONY: lint
lint:
	for PKG in $(PACKAGES); do golint -set_exit_status $$PKG || exit 1; done;

.PHONY: vet
vet:
	$(GOVET) $(PACKAGES)

.PHONY: fmt
fmt:
	$(GOFMT) -s -w $(GOFILES)

.PHONY: fmt-check
fmt-check:
	@diff=$$($(GOFMT) -s -d $(GOFILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

.PHONY: view-covered
view-covered:
	$(GOTEST) -coverprofile=cover.out ./...
	$(GOCMD) tool cover -html=cover.out