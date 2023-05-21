.PHONY: all

libtest: test.cpp
	g++ -o libtest.so test.cpp \
	-std=c++17 -O3 -Wall -Wextra -fPIC -shared

test_cpp: test_cpp.go libtest
	go build test_cpp.go
	./test_cpp

all:
	go run hello-world.go
	go run values.go
	go run variables.go
	go run constant.go
	go run for.go
	go run if-else.go
	go run switch.go
	go run arrays.go
	go run slice.go
	go run map.go
	go run range.go
	go run functions.go
	go run pointer.go
	go run strings-and-runes.go
	go run structs.go
	go run interfaces.go
	go run struct-embedding.go
	# go run generics.go
	go run errors.go
	go run goroutines.go
	go run channels.go
	go run channel-synchronization.go
	go run channel-directions.go
	go run select.go
	go run timeouts.go
	go run non-blocking-channel-operations.go 
	go run closing-channels.go