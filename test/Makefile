TARNAME=gh_2.23.0_linux_amd64.tar.gz
FOLDERNAME=gh_2.23.0_linux_amd64
VAGRANT=/mnt/c/HashiCorp/Vagrant/bin/vagrant.exe


# test
BNAME=ya

## ci-doc-run
CINAME=ya-ci-dock
# GOBURL=github.com/cbot918/ya
ci-doc-run: $(BNAME)
	docker run -dit --name $(CINAME)  cbot918/ubugo
	docker cp $(BNAME) $(CINAME):.
	rm $(BNAME)
	docker exec $(CINAME) bash -c "./$(BNAME) install docker"
	docker exec -it $(CINAME) bash

ci-doc-clean:
	docker stop $(CINAME)
	docker rm $(CINAME)

## ci-vm-run
ci-vm-run: $(BNAME)
	$(VAGRANT) scp $(BNAME) ci-vm:/home/vagrant
	$(VAGRANT) ssh -c "sudo chmod +x /home/vagrant/$(BNAME) && sudo mv /home/vagrant/$(BNAME) /usr/local/bin"
	rm $(BNAME)

ci-vm-init:
	$(VAGRANT) up
	$(VAGRANT) snapshot push

ci-vm-reset:
	$(VAGRANT) snapshot pop
	$(VAGRANT) snapshot push

ci-vm-destroy:
	$(VAGRANT) destroy ci-vm -f


build: main.go
	go build -o ya . 

ya: main.go
	go build -o ya . 





.PHONY: clean test test-install ci
.SILENT: clean test test-install ci clean ci-vm-run build