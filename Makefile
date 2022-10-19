tag: commit
	git tag -a v$(t) -m "$(m)"

add: 
	git add .

commit: add
	git commit -m "$(m)"

push: tag
	git push origin v$(t)

release: push

.PHONY: tag, add, commit, push, release 