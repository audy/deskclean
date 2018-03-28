deskclean:
	go fmt
	go build

test:
	mkdir test-desktop
	touch test-desktop/textfile.md
	touch test-desktop/image.png
	touch test-desktop/script.py
	touch test-desktop/data.fasta
	./deskclean -path test-desktop
	rm -rf test-desktop/

install: deskclean
	cp deskclean /usr/local/bin/deskclean
