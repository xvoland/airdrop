# for arm64
#CGO_ENABLED=1 GOARCH=arm64 go build -o airdrop_arm64 \
#  -ldflags="-extldflags '-L. -lairdrop -framework Cocoa -framework CoreFoundation -framework Security'" \
#  main.go

# for x86_64
#CGO_ENABLED=1 GOARCH=amd64 go build -o airdrop_x86_64 \
#  -ldflags="-extldflags '-L. -lairdrop -framework Cocoa -framework CoreFoundation -framework Security'" \
#  main.go

# Combining binaries
# lipo -create -output airdrop airdrop_arm64 airdrop_x86_64


# swiftc -c AirDropBridge.swift -o AirDropBridge.o
swiftc -c -parse-as-library AirDropBridge.swift -o AirDropBridge.o

ar rcs libairdrop.a AirDropBridge.o

# runtime lib libairdrop.a

# [?] sign app airdrop
# codesign --force --sign - airdrop


### for Makefile
#
#ARCH=arm64
#OBJ=AirDropBridge.o
#LIB=libairdrop.a
#BIN=airdrop

#all: $(BIN)

#$(OBJ): AirDropBridge.swift airdrop.h
#	swiftc -c -target $(ARCH)-apple-macos11 -import-objc-header airdrop.h -parse-as-library AirDropBridge.swift -o $(OBJ)

#$(LIB): $(OBJ)
#	ar rcs $(LIB) $(OBJ)

#$(BIN): main.go $(LIB)
#	CGO_ENABLED=1 GOARCH=$(ARCH) go build -o $(BIN) \
#		-ldflags="-extldflags '-L. -lairdrop -framework Cocoa -framework CoreFoundation -framework Security'" \
#		main.go

#clean:
#	rm -f $(OBJ) $(LIB) $(BIN)
