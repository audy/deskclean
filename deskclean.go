package main

import (
	"flag"
	"fmt"
	"github.com/pelletier/go-toml/v2"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path"
    "path/filepath"
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

type FileTypeConfig map[string]map[string]string

func readConfig() ([]byte, error) {
    homedir, err := os.UserHomeDir()

    if err != nil {
        return nil, fmt.Errorf("Error getting home directory: %v", err)
    }

    filename := filepath.Join(homedir, ".config", "deskclean", "config.toml")

	data , err := os.ReadFile(filename)
	if err != nil {
        if os.IsNotExist(err) {
            log.Printf("Using default config")
            defaultConfig := `
            [textfiles]
            pattern = ".*\\.(rtf|rtfd|md|txt|docx?|rtf|html?|pdf|log)$"
            
            [data]
            pattern = ".*\\.(ab1|csv|sam|fasta|fastq|fa|fna|faa|gbk?|gbf|gff|numbers|aln|zip|tar.gz|xlsx?|sqlite|json?)$"
            
            [scripts]
            pattern = ".*\\.(rmd|go|sql|pl|py|sh|rb|js|ts|coffee|c|r|R|ipynb)$"
            
            [images]
            pattern = ".*\\.(svg|jpe?g|png|gif|gifv|bmp|mp4|mov|m4v|ai|webp)$"
            `

            return []byte(defaultConfig), nil
        }

        // if there is another error (permission denied, etc), return the error
		return nil, fmt.Errorf("Error reading file %s: %v", filename, err)
	}

    log.Printf("Reading config from %s", filename)
    return data, nil
}

func getFileTypes() (map[string]*regexp.Regexp, error) {
	var config FileTypeConfig

    data, err := readConfig()

    if err != nil {
        log.Fatalf("Failed to read TOML file: %v", err)
    }

	if err := toml.Unmarshal([]byte(data), &config); err != nil {
		return nil, fmt.Errorf("Error parsing TOML file: %v", err)
	}

	fileTypes := make(map[string]*regexp.Regexp)

	for key, value := range config {
		if pattern, ok := value["pattern"]; ok {
			compiledRegex, err := regexp.Compile(pattern)
			if err != nil {
				return nil, fmt.Errorf("Error compiling regex for key %s: %v", key, err)
			}
			fileTypes[key] = compiledRegex
		} else {
			log.Printf("No pattern found for key %s", key)
		}
	}

	return fileTypes, nil
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

	filetypes, err := getFileTypes()

	if err != nil {
		log.Fatalf("Failed to get file types: %v", err)
	}

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
