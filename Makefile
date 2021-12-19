
BIN		:= 	""

# check GOBIN
ifneq ($(GOBIN),)
	BIN=$(GOBIN)
else
	# check GOPATH
	ifneq ($(GOPATH),)
		BIN=$(GOPATH)/bin
	endif
endif

clean:
	@echo ${BIN}
	@echo "clean finished."

# .....