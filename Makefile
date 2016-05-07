
sandstorm_repo ?=

capnp_dir := $(sandstorm_repo)/src/sandstorm
capnp_in  := $(wildcard $(capnp_dir)/*.capnp)
capnp_out := $(capnp_in:$(capnp_dir)/%.capnp=capnp/capnp_gen/%.capnp)
go_out    := $(capnp_out:capnp/capnp_gen/%.capnp=capnp/%/capnp.go)

all: $(go_out)
	echo $(go_out)


capnp/capnp_gen/%.capnp: $(capnp_dir)/%.capnp
	@echo GEN $@
	@[ -d capnp/capnp_gen ] || mkdir -p capnp/capnp_gen
	@cp $< $@
	@printf 'using Go = import "/zombiezen.com/go/capnproto2/go.capnp";\n' >> $@
	@printf '$$Go.package("$*");\n' >> $@
	@printf '$$Go.import("zenhack.net/go/sandstorm/capnp/$*");\n' >> $@

# TODO: Target name here is wrong; need to figure out how to specify this
# correctly.
capnp/%/capnp.go: capnp/capnp_gen/%.capnp $(capnp_out)
	@echo COMPILE $<
	@[ -d capnp/$* ] || mkdir capnp/$*
	@capnp compile -I $(GOPATH)/src -ogo:capnp/$* $<

.PHONY: all
