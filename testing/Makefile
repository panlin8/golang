all:release

SOURCE= hello.go
TARGET= test

#gdb -> source /usr/local/src/go/src/runtime/runtime-gdb.py   load go support
debug:
	go build -gcflags "-N -l" -o $(TARGET) $(SOURCE)

release:
	go build -o $(TARGET) $(SOURCE)

test:
	go test -v

bench:
	go test -v -run="none" -bench="BenchmarkSprintf"

testall:
	go test -bench=".*"

# generate cpuprofile file
cprof:
	go test -run=none -bench=".*" -cpuprofile=cprof

# generate hotspot map use pprof tool
pprof:
	go tool pprof --text testing.test cprof

asm:
	go build -gcflags -S $(SOURCE)

escape:
	go build -gcflags=-m -o $(TARGET) $(SOURCE)

clean:
	rm -rf $(TARGET) testing.test cprof mprof
