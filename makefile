bk:
	git add .
	git commit -m "backup"
	git push
	./semver

cover:
	go test ./... -coverprofile=coverage.out
	go tool cover -func=coverage.out

latest-tag:
	git describe --tags --abbrev=0

pull-all:
	for d in ./*/ ; do (cd "$d" && git pull); done