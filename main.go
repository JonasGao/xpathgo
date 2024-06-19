package main

import (
	"fmt"
	"github.com/antchfx/xmlquery"
	"os"
)

func printAndExit(obj interface{}) {
	_, err := fmt.Fprintf(os.Stderr, "Error: %v", obj)
	if err != nil {
		return
	}
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: xpathgo <xpath> <xml file path>")
		return
	}
	path := os.Args[2]
	xpath := os.Args[1]
	defer func() {
		if r := recover(); r != nil {
			printAndExit(r)
		}
	}()
	var err error
	var file *os.File
	var doc *xmlquery.Node
	if _, err = os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			printAndExit("File not exists: " + path)
		}
		printAndExit("Can not open file: " + path)
	}
	file, err = os.Open(path)
	if err != nil {
		panic(err)
	}
	doc, err = xmlquery.Parse(file)
	if err != nil {
		panic(err)
	}
	list := xmlquery.Find(doc, xpath)
	for _, node := range list {
		fmt.Println(node.OutputXML(true))
	}
}
