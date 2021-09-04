# go-enum

[![Actions Status](https://github.com/abice/go-enum/actions/workflows/build_and_test.yml/badge.svg)](https://github.com/abice/go-enum/actions/workflows/build_and_test.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/abice/go-enum)](https://goreportcard.com/report/github.com/abice/go-enum)
[![Coverage Status](https://coveralls.io/repos/github/abice/go-enum/badge.svg)](https://coveralls.io/github/abice/go-enum)
[![GoDoc](https://godoc.org/github.com/abice/go-enum?status.svg)](https://godoc.org/github.com/abice/go-enum)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)

An enum generator for go

## How it works

The goal of go-enum is to create an easy to use enum generator that will take a decorated type declaration like `type EnumName int` and create the associated constant values and funcs that will make life a little easier for adding new values.
It's not perfect, but I think it's useful.

I took the output of the [Stringer](https://godoc.org/golang.org/x/tools/cmd/stringer) command as the `String()` method, and added a way to parse a string value.

## Command options

``` shell
go-enum --help
NAME:
   go-enum - An enum generator for go

USAGE:
   go-enum [global options] [arguments...]

GLOBAL OPTIONS:
   --file value, -f value      The file(s) to generate enums.  Use more than one flag for more files.
   --noprefix                  Prevents the constants generated from having the Enum as a prefix. (default: false)
   --lower                     Adds lowercase variants of the enum strings for lookup. (default: false)
   --nocase                    Adds case insensitive parsing to the enumeration (forces lower flag). (default: false)
   --marshal                   Adds text (and inherently json) marshalling functions. (default: false)
   --sql                       Adds SQL database scan and value functions. (default: false)
   --flag                      Adds golang flag functions. (default: false)
   --prefix value              Replaces the prefix with a user one.
   --names                     Generates a 'Names() []string' function, and adds the possible enum values in the error response during parsing (default: false)
   --nocamel                   Removes the snake_case to CamelCase name changing (default: false)
   --ptr                       Adds a pointer method to get a pointer from const values (default: false)
   --sqlnullint                Adds a Null{{ENUM}} type for marshalling a nullable int value to sql (default: false)
   --sqlnullstr                Adds a Null{{ENUM}} type for marshalling a nullable string value to sql.  If sqlnullint is specified too, it will be Null{{ENUM}}Str (default: false)
   --template value, -t value  Additional template file(s) to generate enums.  Use more than one flag for more files. Templates will be executed in alphabetical order.
   --help, -h                  show help (default: false)
```

### Syntax

The parser looks for comments on your type defs and parse the enum declarations from it.
The parser will look for `ENUM(` and continue to look for comma separated values until it finds a `)`.  You can put values on the same line, or on multiple lines.\
If you need to have a specific value jump in the enum, you can now specify that by adding `=numericValue` to the enum declaration.  Keep in mind, this resets the data for all following values.  So if you specify `50` in the middle of an enum, each value after that will be `51, 52, 53...`

#### Comments

You can use comments inside enum that start with `//`\
The comment must be at the end of the same line as the comment value, only then it will be added as a comment to the generated constant.

```go
// Commented is an enumeration of commented values
/*
ENUM(
value1 // Commented value 1
value2
value3 // Commented value 3
)
*/
type Commented int
```

The generated comments in code will look something like:

```go
...
const (
    // CommentedValue1 is a Commented of type Value1
    // Commented value 1
    CommentedValue1 Commented = iota
    // CommentedValue2 is a Commented of type Value2
    CommentedValue2
    // CommentedValue3 is a Commented of type Value3
    // Commented value 3
    CommentedValue3
)
...
```

#### Example

There are a few examples in the `example` [directory](example).
I've included one here for easy access, but can't guarantee it's up to date.

``` go
// Color is an enumeration of colors that are allowed.
/* ENUM(
Black, White, Red
Green = 33 // Green starts with 33
*/
// Blue
// grey=
// yellow
// blue-green
// red-orange
// yellow_green
// red-orange-blue
// )
type Color int32
```

The generated code will look something like:

``` go
// Code generated by go-enum
// DO NOT EDIT!

package example

import (
    "fmt"
    "strings"
)

const (
    // ColorBlack is a Color of type Black
    ColorBlack Color = iota
    // ColorWhite is a Color of type White
    ColorWhite
    // ColorRed is a Color of type Red
    ColorRed
    // ColorGreen is a Color of type Green
    // Green starts with 33
    ColorGreen Color = iota + 30
    // ColorBlue is a Color of type Blue
    ColorBlue
    // ColorGrey is a Color of type Grey
    ColorGrey
    // ColorYellow is a Color of type Yellow
    ColorYellow
    // ColorBlueGreen is a Color of type Blue-Green
    ColorBlueGreen
    // ColorRedOrange is a Color of type Red-Orange
    ColorRedOrange
    // ColorYellowGreen is a Color of type Yellow_green
    ColorYellowGreen
    // ColorRedOrangeBlue is a Color of type Red-Orange-Blue
    ColorRedOrangeBlue
)

const _ColorName = "BlackWhiteRedGreenBluegreyyellowblue-greenred-orangeyellow_greenred-orange-blue"

var _ColorMap = map[Color]string{
    0:  _ColorName[0:5],
    1:  _ColorName[5:10],
    2:  _ColorName[10:13],
    33: _ColorName[13:18],
    34: _ColorName[18:22],
    35: _ColorName[22:26],
    36: _ColorName[26:32],
    37: _ColorName[32:42],
    38: _ColorName[42:52],
    39: _ColorName[52:64],
    40: _ColorName[64:79],
}

// String implements the Stringer interface.
func (x Color) String() string {
    if str, ok := _ColorMap[x]; ok {
        return str
    }
    return fmt.Sprintf("Color(%d)", x)
}

var _ColorValue = map[string]Color{
    _ColorName[0:5]:                    0,
    strings.ToLower(_ColorName[0:5]):   0,
    _ColorName[5:10]:                   1,
    strings.ToLower(_ColorName[5:10]):  1,
    _ColorName[10:13]:                  2,
    strings.ToLower(_ColorName[10:13]): 2,
    _ColorName[13:18]:                  33,
    strings.ToLower(_ColorName[13:18]): 33,
    _ColorName[18:22]:                  34,
    strings.ToLower(_ColorName[18:22]): 34,
    _ColorName[22:26]:                  35,
    strings.ToLower(_ColorName[22:26]): 35,
    _ColorName[26:32]:                  36,
    strings.ToLower(_ColorName[26:32]): 36,
    _ColorName[32:42]:                  37,
    strings.ToLower(_ColorName[32:42]): 37,
    _ColorName[42:52]:                  38,
    strings.ToLower(_ColorName[42:52]): 38,
    _ColorName[52:64]:                  39,
    strings.ToLower(_ColorName[52:64]): 39,
    _ColorName[64:79]:                  40,
    strings.ToLower(_ColorName[64:79]): 40,
}

// ParseColor attempts to convert a string to a Color
func ParseColor(name string) (Color, error) {
    if x, ok := _ColorValue[name]; ok {
        return x, nil
    }
    return Color(0), fmt.Errorf("%s is not a valid Color", name)
}

// MarshalText implements the text marshaller method
func (x Color) MarshalText() ([]byte, error) {
    return []byte(x.String()), nil
}

// UnmarshalText implements the text unmarshaller method
func (x *Color) UnmarshalText(text []byte) error {
    name := string(text)
    tmp, err := ParseColor(name)
    if err != nil {
        return err
    }
    *x = tmp
    return nil
}

```

## Adding it to your project

1. `go get github.com/abice/go-enum`
1. Add a go:generate line to your file like so... `//go:generate go-enum -f=$GOFILE --marshal`
1. Run go generate like so `go generate ./...`
1. Enjoy your newly created Enumeration
