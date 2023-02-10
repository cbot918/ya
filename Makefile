install:
	./install.sh

test:
	./test/ete.sh

test-clean:
	docker container stop ete
	docker container rm ete

.PHONY: install test test-clean
.SILENT: install