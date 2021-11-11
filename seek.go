// Program that seeks directories, files and keywords/patterns from a file.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"gitlab.com/utt_meelis/walks"
)

// TODO: documentation

// TODO: switch from Search to walks.Search
// Search is a variable to hold the regexp search pattern
var Search *regexp.Regexp = regexp.MustCompile("")

// typeStr is a variable to hold the type of our search (dir, file, kw)
var typeStr string

//
var format string

func fileAction(path string) {
	if (typeStr == "" || typeStr == "file") && Search.MatchString(path) {
		fmt.Println(path)
	}
	if typeStr == "file" || typeStr == "dir" {
		return
	}
	// Will ignore file contents, that does not have extensions (eg binaries) or have .exe extension.
	if filepath.Ext(path) == "" || filepath.Ext(path) == ".exe" {
		return
	}
	contents, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	for lineNr, line := range strings.Split(string(contents), "\n") {
		// if walks.Search.MatchString(line) {
		if Search.MatchString(line) {
			formattedLine := fmt.Sprintf(format, path+":"+fmt.Sprint(lineNr+1)+":", strings.TrimLeft(line, " |\t"))
			fmt.Print(formattedLine)
		}
	}
}
func dirAction(path string) {
	// if walks.Search.MatchString(path) {
	if (typeStr == "" || typeStr == "dir") && Search.MatchString(path) {
		fmt.Println(path)
	}
}

func help() {
	fmt.Println("TODO:")
}

func main() {

	typeFlag := flag.String("type", "", "Type (directory, file, kw) that we are searching for.")
	ignore := flag.String("ignore", "\\.git", "REGEXP_PATTERN that we want to ignore.")
	indent := flag.Int("indent", 60, "The size of indentation between filepath found keyword.")
	depth := flag.Int("depth", -1, "The depth of directory structure recursion, -1 is exhaustive recursion.")
	flag.Parse()

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if strings.Contains(os.Args[1], "help") {
		help()
		os.Exit(0)
	}

	format = "%-" + fmt.Sprint(*indent) + "s%s\n"
	typeStr = *typeFlag

	root, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	var search string
	for i := flag.NFlag() + 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		if arg == "." || arg == ".." || strings.Contains(arg, ".") {
			arg = strings.Replace(arg, ".", "\\.", -1)
		}
		if i == flag.NFlag()+1 {
			search += arg
			continue
		}
		search += "|" + arg
	}

	// walks.Search = regexp.MustCompile(search)
	Search = regexp.MustCompile(search)
	walks.Ignore = regexp.MustCompile(*ignore)
	walks.Walk(root, fileAction, dirAction, *depth)
	// wait until waitgroups are Done
	walks.WaitGroup.Wait()
}
