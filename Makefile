CXX = "g++"
CXXFLAGS = "-Wall"

all: bin/arrays_and_strings

bin/%: src/%.cc
	$(CXX) $< $(CXXFLAGS) -o $@

