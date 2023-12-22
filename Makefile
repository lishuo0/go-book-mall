DEFAULT := ./cmd
OUTPUT := ./bin/mall

default: build

#编译标识注意 -N -l ，禁止编译优化
build:
		go build  -gcflags "-N -l"  -o $(OUTPUT) $(DEFAULT)