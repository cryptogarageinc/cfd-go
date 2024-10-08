.PHONY: all
all: generate format

# for docker
get-cache:
	go install golang.org/x/tools/cmd/goimports@v0.17.0
	go install go.uber.org/mock/mockgen@v0.4.0
	go mod download

update:
	go mod tidy

generate:
	go generate ./apis/... ./service/...

format:
	go run golang.org/x/tools/cmd/goimports@v0.17.0 -w .
	go vet . ./types/... ./errors ./utils ./config ./apis/... ./service/... ./tests
	go run github.com/google/yamlfmt/cmd/yamlfmt@v0.10.0
	go mod tidy

lint:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2 run

lint-fix:
	go run github.com/golangci/golangci-lint/cmd/golangci-lint@v1.55.2 run --fix

build-lib:
	echo "build for Linux/MacOS"
	./tools/simple_build.sh

build-lib-win:
	echo "build for Windows"
	.\tools\build_mingw.bat

build-all: cleanup build-lib

build-all-win: cleanup-win build-lib-win

cleanup:
	echo "cleanup for Linux/MacOS"
	./tools/cmake_cleanup.sh

cleanup-win:
	echo "cleanup for Windows"
	.\tools\cmake_cleanup.bat

test:
	echo "test for Linux/MacOS"
	go version
	go mod download
	./go_test.sh

test-win:
	echo "test for Windows"
	go version
	go mod download
	go_test.bat

gen-swig:
	./tools/gen_swig.sh
	make format
