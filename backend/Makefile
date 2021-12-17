PROJECT := frekwencja

BUILD_DIR := $(shell pwd)/bin
SRC_DIR := .

$(BUILD_DIR)/$(PROJECT): $(SRC_DIR)/*.go
	cd $(SRC_DIR) && GOOS=linux go build -o $(BUILD_DIR)/ .

.PHONY: zip
zip: $(BUILD_DIR)/$(PROJECT)
	-rm -f ./package.zip
	zip -j package $(BUILD_DIR)/*

.PHONY: clean
clean:
	-rm -rf $(BUILD_DIR)
