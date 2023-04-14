# Paths to tools needed in dependencies
GO := $(shell which go)


# Paths to locations, etc
BUILD_DIR := build
CMD_DIR := $(wildcard cmd/*)
BUILD_FLAGS = -ldflags "-s -w" 


# Determine file extension based on OS
ifeq ($(OS),Windows_NT)
    EXT := .exe
else
    EXT :=
endif

# Targets
all: clean cmd

cmd: $(wildcard cmd/*)


$(CMD_DIR): dependencies mkdir
	@echo Build cmd $(notdir $@)
	@${GO} build ${BUILD_FLAGS} -o ${BUILD_DIR}/$(notdir $@)${EXT} ./$@

FORCE:


dependencies:
	@test -f "${GO}" && test -x "${GO}"  || (echo "Missing go binary" && exit 1)

mkdir:
	@echo Mkdir ${BUILD_DIR}
	@install -d ${BUILD_DIR}

clean:
	@echo Clean
	@rm -fr $(BUILD_DIR)
	@${GO} mod tidy
	@${GO} clean