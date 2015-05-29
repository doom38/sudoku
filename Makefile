
all: clean compile

clean:
	rm -rfv bin pkg

compile:
	GOPATH=`pwd` go install ./src/sudoku/...