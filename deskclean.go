package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
	"regexp"
)

func getHomeDirectory() string {
	usr, _ := user.Current()
	return usr.HomeDir
}

// check if path is directory
func isDirectory(origin string) bool {
	file, err := os.Stat(origin)

	if err != nil {
		log.Fatal(err)
	}

	return file.Mode().IsDir()
}

func main() {

	// directory name -> regular expression that captures it
	filetypes := map[string]*regexp.Regexp{
		"textfiles": regexp.MustCompile(`.(rtf|rtfd|md|txt|docx?|rtf|html?|pdf)$`),
		"data":      regexp.MustCompile(`.(ab1|csv|sam|fasta|fastq|fa|fna|faa|gbk?|gbf|gff|aln|zip|tar.gz|xlsx?|sqlite|json?)$`),
		"scripts":   regexp.MustCompile(`.(rmd|go|sql|pl|py|sh|rb|js|ts|coffee|c|r|ipynb)$`),
		"images":    regexp.MustCompile(`.(svg|jpe?g|png|gif|gifv|bmp|mp4|mov|m4v|ai)$`),
	}

	home := getHomeDirectory()

	files, err := ioutil.ReadDir(path.Join(home, "Desktop"))

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		origin := path.Join(home, "Desktop", f.Name())

		matched := false

		if !isDirectory(origin) {
			for dir, re := range filetypes {
				matches := re.MatchString(origin)
				if matches {
					matched = true
					destination := path.Join(home, "Desktop", dir, f.Name())
					fmt.Printf("-> %s -> %s\n", origin, destination)
					err := os.Rename(origin, destination)
					if err != nil {
						log.Fatal(fmt.Sprintf("xx %s -> %s", origin, destination))
					}
				}
			}
			if !matched {
				fmt.Printf("?? %s\n", origin)
			}
		}

	}
}
