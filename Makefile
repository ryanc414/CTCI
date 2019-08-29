CXX = g++
CXXFLAGS = -Wall -Werror -g -Og

all: bin/arrays_and_strings

bin/%: src/%.cc
	$(CXX) $< $(CXXFLAGS) -o $@

