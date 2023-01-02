
exe := ../_build/tempest-sandbox-launcher
objects := sandbox-launcher.o

all: $(exe)
clean:
	rm -f $(exe) *.o *.d

%.o: %.c
	$(CC) $(CFLAGS) -std=c11 -Wall -Wextra -MMD -c -o $@ $<
$(exe): $(objects)
	$(CC) $(LDFLAGS) -o $@ $(objects)

-include *.d

.SUFFIXES:
.PHONY: all clean
