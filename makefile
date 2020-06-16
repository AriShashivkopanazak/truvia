# test makefile

truvia:
	go build src/*.go
	mv tasks build/truvia
	mv api build/truvia
