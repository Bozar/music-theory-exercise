FILE_MAIN := main.py
FILE_GIT_KEEP := .gitkeep
FILE_GIT_IGNORE := .gitignore

DIR_BIN := ./bin


.PHONY: build
build:
	@:


.PHONY: run
run:
	@python3 $(FILE_MAIN)


.PHONY: git
git:
	@git log --oneline -n 5
	@printf "\n"
	@git status


.PHONY: tig
tig:
	@tig


.PHONY: init
init:
	mkdir -p $(DIR_BIN)
	touch $(FILE_MAIN)
	touch $(FILE_GIT_IGNORE)
	touch $(DIR_BIN)/$(FILE_GIT_KEEP)
	git init
