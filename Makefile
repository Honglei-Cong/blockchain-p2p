

BUILD_TAGS?=blockchain-p2p
BUILD_FLAGS = -ldflags "-X github.com/9thchain/blockchain-p2p/version.GitCommit=`git rev-parse --short=8 HEAD`"

build:
		go build $(BUILD_FLAGS) -tags '$(BUILD_TAGS)' -o build/blockchain-p2p ./cmd/
