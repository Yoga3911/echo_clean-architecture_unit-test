tag:
	git tag -a v$(tag) -m "$(msg)"

add: tag
	git add .

commit: add
	git commit -m "$(msg)"

push: commit
	git push origin v$(tag)
	git push

release: push

.PHONY: tag, add, commit, push, release 