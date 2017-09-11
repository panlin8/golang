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

clean:
	    rm -rf $(TARGET)
