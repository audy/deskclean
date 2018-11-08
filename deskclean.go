package main

import (
	"flag"
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

func getFileTypes() map[string]*regexp.Regexp {
	// directory name -> regular expression that captures it
	return map[string]*regexp.Regexp{
		"textfiles": regexp.MustCompile(`.(rtf|rtfd|md|txt|docx?|rtf|html?|pdf)$`),
		"data":      regexp.MustCompile(`.(ab1|csv|sam|fasta|fastq|fa|fna|faa|gbk?|gbf|gff|numbers|aln|zip|tar.gz|xlsx?|sqlite|json?)$`),
		"scripts":   regexp.MustCompile(`.(rmd|go|sql|pl|py|sh|rb|js|ts|coffee|c|r|ipynb)$`),
		"images":    regexp.MustCompile(`.(svg|jpe?g|png|gif|gifv|bmp|mp4|mov|m4v|ai)$`),
	}
}

func parseArgs() map[string]string {

	// default home directory: ~/Desktop
	home := getHomeDirectory()
	desktop := path.Join(home, "Desktop")

	path := flag.String("path", desktop, "path to clean")

	flag.Parse()

	return map[string]string{
		"path": *path,
	}
}

func main() {

	// parse arguments
	args := parseArgs()

	filetypes := getFileTypes()

	files, err := ioutil.ReadDir(args["path"])

	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		origin := path.Join(args["path"], f.Name())

		matched := false

		if !isDirectory(origin) {
			for dir, re := range filetypes {
				matches := re.MatchString(origin)
				if matches {
					matched = true
					destinationDirectory := path.Join(args["path"], dir)

					// make destination directory
					// will not fail if directory already exists
					os.Mkdir(destinationDirectory, os.FileMode(0700))

					destination := path.Join(destinationDirectory, f.Name())
					fmt.Printf("-> %s -> %s\n", origin, destination)
					err := os.Rename(origin, destination)
					if err != nil {
						log.Fatal(fmt.Sprintf("xx %s -> %s\n", origin, destination))
					}
				}
			}
			if !matched {
				fmt.Printf("?? %s\n", origin)
			}
		}

	}
}
