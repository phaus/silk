package main

import (
	"flag"
	"fmt"
	"testing"

	"./runner"
)

var (
	showVersion = flag.Bool("version", false, "show version and exit")
	url         = flag.String("silk.url", "", "(required) target url")
	user        = flag.String("silk.user", "", "(optional) username for basic authentication")
	pass        = flag.String("silk.pass", "", "(optional) password for basic authentication")
	help        = flag.Bool("help", false, "show help")
	paths       []string
)

func main() {
	flag.Parse()
	if *showVersion {
		printversion()
		return
	}
	if *help {
		printhelp()
		return
	}
	if *url == "" {
		fmt.Println("silk.url argument is required")
		return
	}
	paths = flag.Args()
	testing.Main(func(pat, str string) (bool, error) { return true, nil },
		[]testing.InternalTest{{Name: "silk", F: testFunc}},
		nil,
		nil)
}

func testFunc(t *testing.T) {
	r := runner.New(t, *url, *user, *pass)
	fmt.Println("silk: running", len(paths), "file(s)...")
	r.RunGlob(paths, nil)
}

func printhelp() {
	printversion()
	fmt.Println("usage: silk [file] [file2 [file3 [...]]")
	fmt.Println("  e.g: silk ./test/*.silk.md")
	flag.PrintDefaults()
}

func printversion() {
	fmt.Println("silk", version)
}
