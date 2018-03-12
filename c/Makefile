
exe := ../bin/sandbox-launcher
objects := sandbox-launcher.o

all: $(exe)
clean:
	rm -f $(exe) *.o *.d

%.o: %.c
	$(CC) $(CFLAGS) -MMD -c -o $@ $<
$(exe): $(objects)
	$(CC) $(LDFLAGS) -o $@ $(objects)

-include *.d

.SUFFIXES:
.PHONY: all clean
