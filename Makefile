# gobox

NAME ?= gobox
PREFIX ?= /usr/local/bin
CGO_ENABLED?=0
export CGO_ENABLED

all:	${NAME}


build:
	@echo 'Building ${NAME} version ${RELEASE}'
	go get -d -x -v .
	go build -o ${NAME} -x --ldflags "-s -extldflags='-static'"
	@echo 'Successfully built ${NAME}'

${NAME}: build

install: 
	@echo 'PREFIX=${PREFIX}'
	@mkdir -p ${PREFIX}
	@mv ${NAME} ${PREFIX}/${NAME}
	@echo 'Successfully installed ${NAME} to ${PREFIX}'

