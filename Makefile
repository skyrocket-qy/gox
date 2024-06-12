backup:
	golines . -w -m 93
	gofumpt -w .
	git add .
	git commit -m "backup"
	git push