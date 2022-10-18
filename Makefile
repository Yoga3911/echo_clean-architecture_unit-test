tag:
	git tag -a v$(t) -m "$(m)"

add: tag
	git add .

commit: add
	git commit -m "$(m)"

push: commit
	git push origin v$(t)

release: push

.PHONY: tag, add, commit, push, release 