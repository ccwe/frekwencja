PROJECT := frekwencja

# BUILD_DIR := $(shell pwd)/bin
BUILD_DIR := ./bin
SRC_DIR := .

run:
	go run .

.PHONY: all
all: windows macos linux android ios

.PHONY: desktop
desktop: windows macos linux

.PHONY: windows
windows: $(BUILD_DIR)/$(PROJECT)-windows.exe

.PHONY: macos
macos: $(BUILD_DIR)/$(PROJECT)-macos

.PHONY: linux
linux: $(BUILD_DIR)/$(PROJECT)-linux

.PHONY: android
android: $(BUILD_DIR)/$(PROJECT)-android

.PHONY: ios
ios: $(BUILD_DIR)/$(PROJECT)-ios

$(BUILD_DIR)/$(PROJECT)-windows.exe: $(SRC_DIR)/*.go
	cd $(SRC_DIR) && GOOS=windows go build -o $(BUILD_DIR)/$(PROJECT)-windows.exe .

$(BUILD_DIR)/$(PROJECT)-macos: $(SRC_DIR)/*.go
	cd $(SRC_DIR) && GOOS=darwin go build -o $(BUILD_DIR)/$(PROJECT)-macos .

$(BUILD_DIR)/$(PROJECT)-linux: $(SRC_DIR)/*.go
	cd $(SRC_DIR) && GOOS=linux go build -o $(BUILD_DIR)/$(PROJECT)-linux .

$(BUILD_DIR)/$(PROJECT)-android: $(SRC_DIR)/*.go
	cd $(SRC_DIR) && GOOS=linux GOARCH=arm GOARM=7 CGO_ENABLED=1 go build -o $(BUILD_DIR)/$(PROJECT)-android .

$(BUILD_DIR)/$(PROJECT)-ios: $(SRC_DIR)/*.go
	cd $(SRC_DIR) && GOOS=ios CGO_ENABLED=1 go build -o $(BUILD_DIR)/$(PROJECT)-ios .

.PHONY: clean
clean:
	-rm -rf $(BUILD_DIR)
