# xbiggest

>Find the x biggest files in your directories.

- [Installation](#setup--installation)
- [Usage](#usage)
  + [Query](#query-syntax)
    * [Attribute](#attribute)
  + [Examples](#examples-3)
- [Contribute](#contribute)
- [Credits](#credits)
- [License](#license)

## Setup / installation

Requires Go to be [installed](https://golang.org/doc/install) and [configured](https://golang.org/doc/install#testing).

Install with `go get`:

```console
$ go get -u -v github.com/AnujanM/xbiggest/...
$ which xbiggest
$GOPATH/bin/xbiggest
```

Or, build directly via source:

```console
$ git clone https://github.com/AnujanM/xbiggest.git $GOPATH/src/github.com/AnujanM/xbiggest
$ cd $GOPATH/src/github.com/AnujanM/xbiggest
$ make
$ ./xbiggest
```

## Usage

Pass your query to xbiggest via command line argument. View the usage dialogue with the `-help` flag.

```sh
$ xbiggest Directories [number of files]
$ xbiggest "~, ~Documents"
$ xbiggest "~, ~Documents" 25
```

### Query syntax

Every query needs at least 1 directory to search through, the number of files returned can be set on the third argument or the default is 10

#### Source

Each source should be a relative or absolute path to some directory on your machine. Use ~ for your Home Directory.

##### Examples

```sh
$ xbiggest "." 
```

```sh
$ xbiggest "~/Desktop"
```

```sh
$ xbiggest "~/Desktop, $GOPATH"
```

## Contribute

This project is completely open source, feel free to [open an issue](https://github.com/AnujanM/xbiggest/issues) or [submit a pull request](https://github.com/AnujanM/xbiggest/pulls).

Before submitting code, please ensure your changes comply with [Golint](https://github.com/golang/lint). Use `make lint` to test this.

## Credits

Inspired by [kshvmdn](https://github.com/kshvmdn) ([fsql](https://github.com/kshvmdn/fsql))

## License

xbiggest source code is available under the [MIT license](./LICENSE).