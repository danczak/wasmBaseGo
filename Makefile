all:
	@$(MAKE) -C browser -f browser.mk
	@$(MAKE) -C server -f server.mk

clean:
	@$(MAKE) -C browser -f browser.mk clean
	@$(MAKE) -C server -f server.mk clean