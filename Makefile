# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:trivialVersions=true"
GV ?= "device:v1beta1"


# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

OUTPUT_DIR=bin
GOFLAGS=-mod=vendor
define ALL_HELP_INFO
# Build code.
#
# Args:
#   WHAT: Directory names to build.  If any of these directories has a 'main'
#     package, the build will produce executable files under $(OUT_DIR).
#     If not specified, "everything" will be built.
#   GOFLAGS: Extra flags to pass to 'go' when building.
#   GOLDFLAGS: Extra linking flags passed to 'go' when building.
#   GOGCFLAGS: Additional go compile flags passed to 'go' when building.
endef

# 执行代码风格和静态检查
check:
	gofmt -w ./pkg/ ./cmd/
#    go vet ./pkg/... ./cmd/...

# 编译二进制文件
build: check
	hack/gobuild.sh

# 生成CRD文件，使用 sigs.k8s.io/controller-tools v0.2.9 的工具
#
# git clone https://github.com/kubernetes-sigs/controller-tools.git
# git checkout -b v0.2.9 v0.2.9
# go build -o /usr/local/bin/controller-gen cmd/controller-gen/main.go
#
manifests:
	controller-gen \
	object:headerFile=./hack/boilerplate.go.txt \
	paths=./pkg/apis/... \
	rbac:roleName=controller-perms \
	${CRD_OPTIONS} \
	output:crd:artifacts:config=config/crds

# 生成 deepcopy
# --input-dirs 	   :指定需要生成deepcopy的package，如果有多个用逗号分割
# --output-base    :表示生成目录，生成好后将对应的 zz_generated.deepcopy.go copy 到对应的package 下
deepcopy:
	rm -rf ./generate_output
	deepcopy-gen \
	--input-dirs github.com/QQGoblin/device-watcher/pkg/apis/device/v1beta1 \
	-O zz_generated.deepcopy \
	--output-base ./generate_output \
	-h "./hack/boilerplate.go.txt"

# 生成 clientset 接口
#
# --input-base		: CRD 定义的基础目录
# --input    		: 需要生成的CRD包地址，多个时用逗号隔开
# --output-package  : 输出的包路径
# --output-base     ：输出文件夹
generate-client:
	client-gen --clientset-name versioned \
	--input-base github.com/QQGoblin/device-watcher/pkg/apis \
	--input device/v1beta1 \
	--output-package github.com/QQGoblin/device-watcher/pkg/client/clientset \
	--output-base ./generate_output \
    -h "./hack/boilerplate.go.txt"

# 生成 lister 接口
#
# --input-dirs		: 指定需要生成deepcopy的package，如果有多个用逗号分割
# --output-package  : 输出的包路径
# --output-base     ：输出文件夹
generate-lister:
	lister-gen \
	--input-dirs github.com/QQGoblin/device-watcher/pkg/apis/device/v1beta1 \
	--output-package github.com/QQGoblin/device-watcher/pkg/client/listers \
	--output-base ./generate_output \
    -h "./hack/boilerplate.go.txt"

# 生成 informer 接口
#
# --input-dirs		             : 指定需要生成deepcopy的package，如果有多个用逗号分割
# --versioned-clientset-package  : 引用的client包路径
# --listers-package              : 引用的listers包路径
# --output-package               : 输出的包路径
# --output-base                  : 输出文件夹
generate-informer:
	informer-gen \
	--input-dirs github.com/QQGoblin/device-watcher/pkg/apis/device/v1beta1 \
	--versioned-clientset-package github.com/QQGoblin/device-watcher/pkg/client/clientset/versioned \
	--listers-package github.com/QQGoblin/device-watcher/pkg/client/listers \
	--output-package github.com/QQGoblin/device-watcher/pkg/client/informers \
	--output-base ./generate_output \
    -h "./hack/boilerplate.go.txt"

generate-clean:
	rm -rf ./generate_output

generate-all: generate-clean generate-client generate-lister generate-informer