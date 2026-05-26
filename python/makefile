FILE_MAIN := main.py
FILE_GIT_KEEP := .gitkeep
FILE_GIT_IGNORE := .gitignore

DIR_BUILD := ./build


.PHONY: build run init


build:
	@:


run:
	@python3 $(FILE_MAIN)


init:
	mkdir -p $(DIR_BUILD)
	touch $(FILE_MAIN)
	touch $(FILE_GIT_IGNORE)
	touch $(DIR_BUILD)/$(FILE_GIT_KEEP)
	git init
