# seek

This is a tool to search patterns from files, filenames and directory names, written in Go.

## Quick start

* Use build binaries from bin/ (seek.exe for windows, seek for linux).
* Build it yourself:
  * **Dependency**: need to have Go installed
  ```sh
  go build -o seek seek.go
  ```
  * **[optional]** install the binary to local user with
  ```sh
  go install .
  ```

To build binary (64-bit) from linux to windows, run

```console
GOOS=windows GOARCH=amd64 go build -o seek.exe seek.go
```

## Usage

seek \[FLAGS\] PATTERNS

## FLAGS

**NB!** When using up to one flag, then using `=` between flag and value can be omitted.

**NB!** Program does run when omitting `=` between flag and value for multiple flags, but it does not work as excpected.
For specifics of why, visit [Go flag package](https://pkg.go.dev/flag).

**Note**: single dash (`-`) can be used instead of double dash (`--`).

**--depth**=int
```
The depth of directory structure recursion, -1 is exhaustive recursion. (default -1)
```
**--from**=string
```
Specify a file or a directory from which we seek the pattern. (default ".")
```
**--help**
```
Prints help.
```
**--ignore**=string
```
REGEXP_PATTERN that we want to ignore. (default "\\.git")
```
**--indent**=int
```
The size of indentation between filepath and found pattern. (default 30)
```
**--type**=string
```
The search type: do we search the pattern from 'dir'=directory name; 'file'=filename; 'pat'=pattern inside a file. By default search from everywhere.
```
**--todo**
```
In addtion to REGEXP_PATTERN, search for TODO:|NOTE:|HACK:|DEBUG:|FIXME:|REVIEW:|BUG:|TEST:|TESTME:|MAYBE:
```

## Notes

* in vim/ directory I have provided the vimscript function, key remapping and commands related to seek that I use in my `.vimrc`.

## Author

Meelis Utt
