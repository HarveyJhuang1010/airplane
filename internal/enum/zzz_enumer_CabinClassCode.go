// Code generated by "enumer -type=CabinClassCode -trimprefix=CabinClassCode -yaml -json -text -sql -transform=snake --output=zzz_enumer_CabinClassCode.go"; DO NOT EDIT.

package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _CabinClassCodeName = "economy_standardeconomy_flexbusiness_basicbusiness_standard"

var _CabinClassCodeIndex = [...]uint8{0, 16, 28, 42, 59}

const _CabinClassCodeLowerName = "economy_standardeconomy_flexbusiness_basicbusiness_standard"

func (i CabinClassCode) String() string {
	if i < 0 || i >= CabinClassCode(len(_CabinClassCodeIndex)-1) {
		return fmt.Sprintf("CabinClassCode(%d)", i)
	}
	return _CabinClassCodeName[_CabinClassCodeIndex[i]:_CabinClassCodeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _CabinClassCodeNoOp() {
	var x [1]struct{}
	_ = x[CabinClassCodeEconomyStandard-(0)]
	_ = x[CabinClassCodeEconomyFlex-(1)]
	_ = x[CabinClassCodeBusinessBasic-(2)]
	_ = x[CabinClassCodeBusinessStandard-(3)]
}

var _CabinClassCodeValues = []CabinClassCode{CabinClassCodeEconomyStandard, CabinClassCodeEconomyFlex, CabinClassCodeBusinessBasic, CabinClassCodeBusinessStandard}

var _CabinClassCodeNameToValueMap = map[string]CabinClassCode{
	_CabinClassCodeName[0:16]:       CabinClassCodeEconomyStandard,
	_CabinClassCodeLowerName[0:16]:  CabinClassCodeEconomyStandard,
	_CabinClassCodeName[16:28]:      CabinClassCodeEconomyFlex,
	_CabinClassCodeLowerName[16:28]: CabinClassCodeEconomyFlex,
	_CabinClassCodeName[28:42]:      CabinClassCodeBusinessBasic,
	_CabinClassCodeLowerName[28:42]: CabinClassCodeBusinessBasic,
	_CabinClassCodeName[42:59]:      CabinClassCodeBusinessStandard,
	_CabinClassCodeLowerName[42:59]: CabinClassCodeBusinessStandard,
}

var _CabinClassCodeNames = []string{
	_CabinClassCodeName[0:16],
	_CabinClassCodeName[16:28],
	_CabinClassCodeName[28:42],
	_CabinClassCodeName[42:59],
}

// CabinClassCodeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func CabinClassCodeString(s string) (CabinClassCode, error) {
	if val, ok := _CabinClassCodeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _CabinClassCodeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to CabinClassCode values", s)
}

// CabinClassCodeValues returns all values of the enum
func CabinClassCodeValues() []CabinClassCode {
	return _CabinClassCodeValues
}

// CabinClassCodeStrings returns a slice of all String values of the enum
func CabinClassCodeStrings() []string {
	strs := make([]string, len(_CabinClassCodeNames))
	copy(strs, _CabinClassCodeNames)
	return strs
}

// IsACabinClassCode returns "true" if the value is listed in the enum definition. "false" otherwise
func (i CabinClassCode) IsACabinClassCode() bool {
	for _, v := range _CabinClassCodeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for CabinClassCode
func (i CabinClassCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for CabinClassCode
func (i *CabinClassCode) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("CabinClassCode should be a string, got %s", data)
	}

	var err error
	*i, err = CabinClassCodeString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for CabinClassCode
func (i CabinClassCode) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for CabinClassCode
func (i *CabinClassCode) UnmarshalText(text []byte) error {
	var err error
	*i, err = CabinClassCodeString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for CabinClassCode
func (i CabinClassCode) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for CabinClassCode
func (i *CabinClassCode) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = CabinClassCodeString(s)
	return err
}

func (i CabinClassCode) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *CabinClassCode) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	var str string
	switch v := value.(type) {
	case []byte:
		str = string(v)
	case string:
		str = v
	case fmt.Stringer:
		str = v.String()
	default:
		return fmt.Errorf("invalid value of CabinClassCode: %[1]T(%[1]v)", value)
	}

	val, err := CabinClassCodeString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
