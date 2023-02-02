
build_dir := ../_build
exe := $(build_dir)/tempest-sandbox-launcher
objects := sandbox-launcher.o

all: $(exe)
clean:
	rm -f $(exe) *.o *.d

%.o: %.c
	$(CC) $(CFLAGS) -I $(build_dir) -std=c11 -Wall -Wextra -MMD -c -o $@ $<
$(exe): $(objects)
	$(CC) $(LDFLAGS) -o $@ $(objects)

sandbox-launcher.o: $(build_dir)/bpf_filter.h

$(build_dir)/gen-clean-h: gen-clean-h.o
	$(CC) $(LDFLAGS) -o $@ $<
$(build_dir)/constants.h: $(build_dir)/gen-clean-h
	$(build_dir)/gen-clean-h > $@
$(build_dir)/filter_preproc.s: $(build_dir)/constants.h filter.s
	cpp -I $(build_dir) filter.s -o $@
$(build_dir)/bpf_filter.h: $(build_dir)/filter_preproc.s
	bpf_asm -c < $< > $@

-include *.d

.SUFFIXES:
.PHONY: all clean
