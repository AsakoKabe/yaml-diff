package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/sters/yaml-diff/yamldiff"
)

func main() {
	file1 := flag.String("file1", "", "")
	file2 := flag.String("file2", "", "")
	flag.Parse()

	yamls1 := yamldiff.Load(load(*file1))
	yamls2 := yamldiff.Load(load(*file2))

	for _, diff := range yamldiff.Do(yamls1, yamls2) {
		fmt.Println(diff.Diff)
	}

	fmt.Print()
}

func load(f string) string {
	file, err := os.Open(f)
	defer file.Close()
	if err != nil {
		log.Fatalf("%+v, %s", err, f)
	}

	b, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("%+v, %s", err, f)
	}

	return string(b)
}