bk:
	golines . -w -m 100
	gofumpt -w .
	git add .
	git commit -m "backup"
	git push

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