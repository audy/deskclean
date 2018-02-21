package main

import (
  "fmt"
  "io/ioutil"
  "log"
  "path"
  "os/user"
  "os"
)

func getHomeDirectory() string {
  usr, _ := user.Current()
  return usr.HomeDir
}

// check if path is directory
func isDirectory(filepath string) bool {
  file, err := os.Stat(filepath)

  if err != nil {
    log.Fatal(err)
  }

  return file.Mode().IsDir()
}

func main() {

  home := getHomeDirectory()

  fmt.Println(home)

  files, err := ioutil.ReadDir(path.Join(home, "Desktop"))

  if err != nil {
    log.Fatal(err)
  }

  for _, f := range files {
    filepath := path.Join(home, "Desktop", f.Name())
    fmt.Println(filepath)

    if isDirectory(filepath) {
      fmt.Println(filepath)
    }

  }
}
