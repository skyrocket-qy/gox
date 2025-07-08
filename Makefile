bk:
	golines . -w -m 100
	gofumpt -w .
	git add .
	git commit -m "backup"
	git push

PATCHVER=$(shell git pull --tags | ./semver get patch $(shell git describe --tags $(shell git rev-list --tags --max-count=1)))
NEWVER=$(shell echo $(PATCHVER) + 1 | bc)

nt: 
	@git tag "v1.0.$(NEWVER)"
	@git push
	@git push --tags

push: 
	@git pull --tags
	@git add .
	@git commit -m "update protos"
	@make nt
