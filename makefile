CC=g++
CFLAGS=-c -Wall
LDFLAGS=
CPP_SOURCES=setup.cpp classes/AgpmlToHtml.cpp classes/Chooser.cpp
C_SOURCES=$(wildcard .c) $(wildcard classes/.c) $(wildcard **/*.c)
OBJECTS=$(addprefix objects/,$(CPP_SOURCES:.cpp=.o)) $(addprefix objects/,$(C_SOURCES:.c=.o))
EXECUTABLE=agpmlc

all: objects $(EXECUTABLE)

objects:
	mkdir -p objects
	mkdir -p objects/classes

$(EXECUTABLE): $(OBJECTS)
	$(CC) $(LDFLAGS) $^ -o $@

objects/classes/%.o: classes/%.cpp
	$(CC) $(CFLAGS) $< -o objects/classes/$(notdir $@)

objects/%.o: %.cpp
	$(CC) $(CFLAGS) $< -o objects/$(notdir $@)

objects/%.o: %.c
	$(CC) $(CFLAGS) $< -o objects/$(notdir $@)

clean:
	rm -rf objects/ $(EXECUTABLE)