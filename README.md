# mdtable2csv - Convert Markdown Table To Csv File

This application convert markdown table to csv file.

[![GitHub license](https://img.shields.io/github/license/BatuhanKucukali/mdtable2csv)](https://github.com/BatuhanKucukali/istekbin-api/blob/master/LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/BatuhanKucukali/mdtable2csv)](https://github.com/BatuhanKucukali/istekbin-api/issues)
[![Go Report Card](https://goreportcard.com/badge/github.com/BatuhanKucukali/mdtable2csv)](https://goreportcard.com/report/github.com/BatuhanKucukali/istekbin-api)

## Install via Homebrew

```bash
brew tap BatuhanKucukali/mdtable2csv
brew install mdtable2csv
```

## Usage

mdtable2csv convert {filePath} {delimiter(default ',')}

Examples

```bash
mdtable2csv convert example/markdown.md
```

```bash
mdtable2csv convert example/markdown.md ;
```

## Result

```
| Tables        | Are           | Cool  |
| ------------- |:-------------:| -----:|
| col 3 is      | right-aligned | $1600 |
| col 2 is      | centered      |   $12 |
| zebra stripes | are neat      |    $1 |
```

```
Tables,Are,Cool
col 3 is,right-aligned,$1600
col 2 is,centered,$12
zebra stripes,are neat,$1
```

## Run this project

1 . Clone project on your machine
```bash
git clone git@github.com:BatuhanKucukali/mdtable2csv.git
```
2 . Change directory
```bash
cd mdtable2csv
```
3 . Build
```bash
go build .
```
4 . Run
```bash
./mdtable2csv convert example/markdown.md
```

## Run Test

```bash
go test ./cmd/
```

## TODO

- [X] Documentation
- [X] Brew repository
- [X] Test
- [ ] Argument helper
- [ ] Argument for trim option
- [ ] Argument for new file name

## Getting help ##

If you're having trouble getting this project running, feel free to [open an issue](https://github.com/BatuhanKucukali/mdtable2csv/issues/new)
