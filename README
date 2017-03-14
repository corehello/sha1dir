# A tool to sha1 all the files in a directory


## build

```
go build .
```

## Usage:
```
Usage of sha1dir:
  -filter string
        the blacklist of the directory, support regexp
  -output string
        the output file name or path (default "sha1result")
  -root string
        the root directory which to be walked (default ".")
```

### Example Usage:
```
./sha1dir -root .
./sha1dir -root . -filter .git
./sha1dir -root . -filter first,second
./sha1dir -root . -filter first,second -output /tmp/sha1
```

## import

```
go get github.com/corehello/sha1dir/sha1dir

import (
  ...
  "github.com/corehello/sha1dir/sha1dir"
  ...
)

```
