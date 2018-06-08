// Code generated by go-enum
// DO NOT EDIT!

package example

import (
	"fmt"
	"strings"
)

const (
	// MakeToyota is a Make of type Toyota
	MakeToyota Make = iota
	// Skipped value
	_
	// MakeChevy is a Make of type Chevy
	MakeChevy
	// Skipped value
	_
	// MakeFord is a Make of type Ford
	MakeFord
	// Skipped value
	_
	// MakeTesla is a Make of type Tesla
	MakeTesla
	// Skipped value
	_
	// MakeHyundai is a Make of type Hyundai
	MakeHyundai
	// Skipped value
	_
	// MakeNissan is a Make of type Nissan
	MakeNissan
	// Skipped value
	_
	// MakeJaguar is a Make of type Jaguar
	MakeJaguar
	// Skipped value
	_
	// MakeAudi is a Make of type Audi
	MakeAudi
	// Skipped value
	_
	// MakeBMW is a Make of type BMW
	MakeBMW
	// Skipped value
	_
	// MakeMercedesBenz is a Make of type Mercedes-Benz
	MakeMercedesBenz
	// Skipped value
	_
	// MakeVolkswagon is a Make of type Volkswagon
	MakeVolkswagon
)

const _MakeName = "ToyotaChevyFordTeslaHyundaiNissanJaguarAudiBMWMercedes-BenzVolkswagon"

var _MakeNames = []string{
	_MakeName[0:6],
	_MakeName[6:11],
	_MakeName[11:15],
	_MakeName[15:20],
	_MakeName[20:27],
	_MakeName[27:33],
	_MakeName[33:39],
	_MakeName[39:43],
	_MakeName[43:46],
	_MakeName[46:59],
	_MakeName[59:69],
}

func MakeNames() []string {
	tmp := make([]string, len(_MakeNames))
	copy(tmp, _MakeNames)
	return tmp
}

var _MakeMap = map[Make]string{
	0:  _MakeName[0:6],
	2:  _MakeName[6:11],
	4:  _MakeName[11:15],
	6:  _MakeName[15:20],
	8:  _MakeName[20:27],
	10: _MakeName[27:33],
	12: _MakeName[33:39],
	14: _MakeName[39:43],
	16: _MakeName[43:46],
	18: _MakeName[46:59],
	20: _MakeName[59:69],
}

func (i Make) String() string {
	if str, ok := _MakeMap[i]; ok {
		return str
	}
	return fmt.Sprintf("Make(%d)", i)
}

var _MakeValue = map[string]Make{
	_MakeName[0:6]:                    0,
	strings.ToLower(_MakeName[0:6]):   0,
	_MakeName[6:11]:                   2,
	strings.ToLower(_MakeName[6:11]):  2,
	_MakeName[11:15]:                  4,
	strings.ToLower(_MakeName[11:15]): 4,
	_MakeName[15:20]:                  6,
	strings.ToLower(_MakeName[15:20]): 6,
	_MakeName[20:27]:                  8,
	strings.ToLower(_MakeName[20:27]): 8,
	_MakeName[27:33]:                  10,
	strings.ToLower(_MakeName[27:33]): 10,
	_MakeName[33:39]:                  12,
	strings.ToLower(_MakeName[33:39]): 12,
	_MakeName[39:43]:                  14,
	strings.ToLower(_MakeName[39:43]): 14,
	_MakeName[43:46]:                  16,
	strings.ToLower(_MakeName[43:46]): 16,
	_MakeName[46:59]:                  18,
	strings.ToLower(_MakeName[46:59]): 18,
	_MakeName[59:69]:                  20,
	strings.ToLower(_MakeName[59:69]): 20,
}

// ParseMake attempts to convert a string to a Make
func ParseMake(name string) (Make, error) {
	if x, ok := _MakeValue[name]; ok {
		return Make(x), nil
	}
	return Make(0), fmt.Errorf("%s is not a valid Make, try [%s]", name, strings.Join(_MakeNames, ", "))
}

func (x *Make) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *Make) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseMake(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

// Set implements the Golang flag.Value interface func
func (x *Make) Set(val string) error {
	v, err := ParseMake(val)
	*x = v
	return err
}

// Get implements the Golang flag.Getter interface func
func (x *Make) Get() interface{} {
	return *x
}

// Type implements the github.com/spf13/pFlag Value interface
func (x *Make) Type() string {
	return "Make"
}

const (
	// NoZerosStart is a NoZeros of type Start
	NoZerosStart NoZeros = iota + 20
	// NoZerosMiddle is a NoZeros of type Middle
	NoZerosMiddle
	// NoZerosEnd is a NoZeros of type End
	NoZerosEnd
	// NoZerosPs is a NoZeros of type Ps
	NoZerosPs
	// NoZerosPps is a NoZeros of type Pps
	NoZerosPps
	// NoZerosPpps is a NoZeros of type Ppps
	NoZerosPpps
)

const _NoZerosName = "startmiddleendpsppsppps"

var _NoZerosNames = []string{
	_NoZerosName[0:5],
	_NoZerosName[5:11],
	_NoZerosName[11:14],
	_NoZerosName[14:16],
	_NoZerosName[16:19],
	_NoZerosName[19:23],
}

func NoZerosNames() []string {
	tmp := make([]string, len(_NoZerosNames))
	copy(tmp, _NoZerosNames)
	return tmp
}

var _NoZerosMap = map[NoZeros]string{
	20: _NoZerosName[0:5],
	21: _NoZerosName[5:11],
	22: _NoZerosName[11:14],
	23: _NoZerosName[14:16],
	24: _NoZerosName[16:19],
	25: _NoZerosName[19:23],
}

func (i NoZeros) String() string {
	if str, ok := _NoZerosMap[i]; ok {
		return str
	}
	return fmt.Sprintf("NoZeros(%d)", i)
}

var _NoZerosValue = map[string]NoZeros{
	_NoZerosName[0:5]:                    20,
	strings.ToLower(_NoZerosName[0:5]):   20,
	_NoZerosName[5:11]:                   21,
	strings.ToLower(_NoZerosName[5:11]):  21,
	_NoZerosName[11:14]:                  22,
	strings.ToLower(_NoZerosName[11:14]): 22,
	_NoZerosName[14:16]:                  23,
	strings.ToLower(_NoZerosName[14:16]): 23,
	_NoZerosName[16:19]:                  24,
	strings.ToLower(_NoZerosName[16:19]): 24,
	_NoZerosName[19:23]:                  25,
	strings.ToLower(_NoZerosName[19:23]): 25,
}

// ParseNoZeros attempts to convert a string to a NoZeros
func ParseNoZeros(name string) (NoZeros, error) {
	if x, ok := _NoZerosValue[name]; ok {
		return NoZeros(x), nil
	}
	return NoZeros(0), fmt.Errorf("%s is not a valid NoZeros, try [%s]", name, strings.Join(_NoZerosNames, ", "))
}

func (x *NoZeros) MarshalText() ([]byte, error) {
	return []byte(x.String()), nil
}

func (x *NoZeros) UnmarshalText(text []byte) error {
	name := string(text)
	tmp, err := ParseNoZeros(name)
	if err != nil {
		return err
	}
	*x = tmp
	return nil
}

// Set implements the Golang flag.Value interface func
func (x *NoZeros) Set(val string) error {
	v, err := ParseNoZeros(val)
	*x = v
	return err
}

// Get implements the Golang flag.Getter interface func
func (x *NoZeros) Get() interface{} {
	return *x
}

// Type implements the github.com/spf13/pFlag Value interface
func (x *NoZeros) Type() string {
	return "NoZeros"
}
