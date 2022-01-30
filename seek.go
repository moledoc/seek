// Program that seeks patterns from files, filenames and directories.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/moledoc/walks"
)

// todos is a variable, where we store pre-defined search pattern of TODOs, NOTEs etc
// The multiplicity of the last letter in todos keyword shows the urgency of the todo.
// NB! the todos are not sorted by the urgency, at least not currently.
// Idea gotten from https://github.com/tsoding/snitch (https://www.youtube.com/watch?v=oyoTEzgFqq8&ab_channel=TsodingDaily)
// and Tsoding took the idea from https://github.com/rolandwalker/fixmee
var todos string = "TOD(O*):|NOT(E*):|HAC(K*):|DEBU(G*):|FIXM(E*):|REVIE(W*):|BU(G*):|TES(T*):|TESTM(E*):|MAYB(E*):"

// typeStr is a variable to hold the type of our search (dir, file, pat)
var typeStr string

// format is a variable that holds the format of the found result
var format string

// fileAction is a function where we perform pattern searcing from a file or a filename.
func fileAction(path string) {
	if (typeStr == "" || typeStr == "file") && walks.Search.MatchString(path) {
		fmt.Println(path)
	}
	if typeStr != "pat" && typeStr != "" {
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
		if walks.Search.MatchString(line) {
			formattedLine := fmt.Sprintf(format, path+":"+fmt.Sprint(lineNr+1)+":", strings.TrimLeft(line, " |\t"))
			fmt.Print(formattedLine)
		}
	}
}

// dirAction is a function where we perform pattern searcing from a directory name.
func dirAction(path string) {
	if (typeStr == "" || typeStr == "dir") && walks.Search.MatchString(path) {
		fmt.Println(path)
	}
}

// help is a function to print program help.
func help() {
	fmt.Println("USAGE\n\tseek [FLAGS] PATTERNS")
	fmt.Println("DESCRIPTION\n\tThis is a tool to search patterns from files, filenames and directory names, written in Go.")
	fmt.Println("FLAGS")
	flag.PrintDefaults()
	os.Exit(0)
}

func main() {
	typeFlag := flag.String("type", "", "The search type: do we search the pattern from 'dir'=directory name; 'file'=filename; 'pat'=pattern inside a file. By default search from everywhere.")
	ignore := flag.String("ignore", "\\.git", "REGEXP_PATTERN that we want to ignore.")
	indent := flag.Int("indent", 30, "The size of indentation between filepath and found pattern.")
	depth := flag.Int("depth", -1, "The depth of directory structure recursion, -1 is exhaustive recursion.")
	from := flag.String("from", ".", "Specify a file or a directory from which we seek the pattern.")
	helpBool := flag.Bool("help", false, "Print this help.")
	findTodos := flag.Bool("todo", false, fmt.Sprintf("In addtion to REGEXP_PATTERN, search for %s", todos))
	flag.Parse()
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	if *helpBool {
		help()
	}
	if !(*typeFlag == "dir" || *typeFlag == "file" || *typeFlag == "pat" || *typeFlag == "") {
		log.Fatalf("Unknown -type value. Expected 'dir', 'file' or 'pat', got '%s'\n", *typeFlag)
	}

	format = "%-" + fmt.Sprint(*indent) + "s%s\n"
	typeStr = *typeFlag

	var search string
	if *findTodos {
		search += todos
	}
	for i := flag.NFlag() + 1; i < len(os.Args); i++ {
		arg := os.Args[i]
		arg = strings.Replace(arg, "\\", "\\\\", -1)
		if search == "" {
			search += arg
			continue
		}
		search += "|" + arg
	}
	walks.Search = regexp.MustCompile(search)
	walks.Ignore = regexp.MustCompile(*ignore)

	root, err := os.Stat(*from)
	if err != nil {
		log.Fatal(err)
	}

	switch mode := root.Mode(); {
	default:
		log.Fatalf("Unreachable: %s is not a file nor a directory\n", *from)
	case mode.IsRegular():
		fileAction(*from)
	case mode.IsDir():
		walks.Walk(*from, fileAction, dirAction, *depth)
		// wait until waitgroups are Done
		walks.WaitGroup.Wait()
	}

}
