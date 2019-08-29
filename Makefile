CXX = g++
CXXFLAGS = -Wall -Werror -g -Og --std=c++17

all: bin/arrays_and_strings

bin/%: src/%.cc
	$(CXX) $< $(CXXFLAGS) -o $@

