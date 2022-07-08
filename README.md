# dirShuffle
Shuffle files by directory

main takes a directory to search, followed by a list of allowed file extensions, and prints a list of filtered files in the directory, shuffled by directory.
The original goal was to shuffle music by album, but there may be other uses.

sample usage:
```
go build .
# print a list of go files in the parent directory (and all subdirectories), shuffled by directory
./dirShuffle $(dirname $(pwd)) ".go"
```
