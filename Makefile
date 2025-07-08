bk:
	golines . -w -m 100
	gofumpt -w .
	git add .
	git commit -m "backup"
	git push
	./semver