// Code generated by "enumer -type=JournalManagerAction -trimprefix=JournalManagerAction -yaml -json -text -sql -transform=snake --output=zzz_enumer_journalManagerAction.go"; DO NOT EDIT.

package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _JournalManagerActionName = "unknown"

var _JournalManagerActionIndex = [...]uint8{0, 7}

const _JournalManagerActionLowerName = "unknown"

func (i JournalManagerAction) String() string {
	if i < 0 || i >= JournalManagerAction(len(_JournalManagerActionIndex)-1) {
		return fmt.Sprintf("JournalManagerAction(%d)", i)
	}
	return _JournalManagerActionName[_JournalManagerActionIndex[i]:_JournalManagerActionIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _JournalManagerActionNoOp() {
	var x [1]struct{}
	_ = x[JournalManagerActionUnknown-(0)]
}

var _JournalManagerActionValues = []JournalManagerAction{JournalManagerActionUnknown}

var _JournalManagerActionNameToValueMap = map[string]JournalManagerAction{
	_JournalManagerActionName[0:7]:      JournalManagerActionUnknown,
	_JournalManagerActionLowerName[0:7]: JournalManagerActionUnknown,
}

var _JournalManagerActionNames = []string{
	_JournalManagerActionName[0:7],
}

// JournalManagerActionString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func JournalManagerActionString(s string) (JournalManagerAction, error) {
	if val, ok := _JournalManagerActionNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _JournalManagerActionNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to JournalManagerAction values", s)
}

// JournalManagerActionValues returns all values of the enum
func JournalManagerActionValues() []JournalManagerAction {
	return _JournalManagerActionValues
}

// JournalManagerActionStrings returns a slice of all String values of the enum
func JournalManagerActionStrings() []string {
	strs := make([]string, len(_JournalManagerActionNames))
	copy(strs, _JournalManagerActionNames)
	return strs
}

// IsAJournalManagerAction returns "true" if the value is listed in the enum definition. "false" otherwise
func (i JournalManagerAction) IsAJournalManagerAction() bool {
	for _, v := range _JournalManagerActionValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for JournalManagerAction
func (i JournalManagerAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for JournalManagerAction
func (i *JournalManagerAction) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("JournalManagerAction should be a string, got %s", data)
	}

	var err error
	*i, err = JournalManagerActionString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for JournalManagerAction
func (i JournalManagerAction) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for JournalManagerAction
func (i *JournalManagerAction) UnmarshalText(text []byte) error {
	var err error
	*i, err = JournalManagerActionString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for JournalManagerAction
func (i JournalManagerAction) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for JournalManagerAction
func (i *JournalManagerAction) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = JournalManagerActionString(s)
	return err
}

func (i JournalManagerAction) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *JournalManagerAction) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of JournalManagerAction: %[1]T(%[1]v)", value)
	}

	val, err := JournalManagerActionString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
