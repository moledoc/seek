# seek

This is a tool to search patterns from directories and files, written in Go

## Usage

seek [SUBCOMMAND | FLAGS] PATTERN(S)

## SUBCOMMAND

**help**

```
Prints help
```

## FLAGS

**type**={dir | file | pat}

```
Specifies if we seek for a directory, file or a pattern inside a file.
The default behavior is to search all of the previously mentioned.
```

**ignore**=REGEXP_PATTERN

```
Pattern of directory and filenames to ignore
```

**indent**=int

```
The size of indentation between filepath and ado. (default 60)
```

**-depth**=int

```
The depth of directory structure recursion, -1 is exhaustive recursion. (default -1)
```

## Examples

TODO:

## Author

Meelis Utt