// Code generated by "enumer -type=BookingStatus -trimprefix=BookingStatus -yaml -json -text -sql -transform=snake --output=zzz_enumer_BookingStatus.go"; DO NOT EDIT.

package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strings"
)

const _BookingStatusName = "pendingconfirmedcancelledexpiredoverbookedconfirming"

var _BookingStatusIndex = [...]uint8{0, 7, 16, 25, 32, 42, 52}

const _BookingStatusLowerName = "pendingconfirmedcancelledexpiredoverbookedconfirming"

func (i BookingStatus) String() string {
	if i < 0 || i >= BookingStatus(len(_BookingStatusIndex)-1) {
		return fmt.Sprintf("BookingStatus(%d)", i)
	}
	return _BookingStatusName[_BookingStatusIndex[i]:_BookingStatusIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _BookingStatusNoOp() {
	var x [1]struct{}
	_ = x[BookingStatusPending-(0)]
	_ = x[BookingStatusConfirmed-(1)]
	_ = x[BookingStatusCancelled-(2)]
	_ = x[BookingStatusExpired-(3)]
	_ = x[BookingStatusOverbooked-(4)]
	_ = x[BookingStatusConfirming-(5)]
}

var _BookingStatusValues = []BookingStatus{BookingStatusPending, BookingStatusConfirmed, BookingStatusCancelled, BookingStatusExpired, BookingStatusOverbooked, BookingStatusConfirming}

var _BookingStatusNameToValueMap = map[string]BookingStatus{
	_BookingStatusName[0:7]:        BookingStatusPending,
	_BookingStatusLowerName[0:7]:   BookingStatusPending,
	_BookingStatusName[7:16]:       BookingStatusConfirmed,
	_BookingStatusLowerName[7:16]:  BookingStatusConfirmed,
	_BookingStatusName[16:25]:      BookingStatusCancelled,
	_BookingStatusLowerName[16:25]: BookingStatusCancelled,
	_BookingStatusName[25:32]:      BookingStatusExpired,
	_BookingStatusLowerName[25:32]: BookingStatusExpired,
	_BookingStatusName[32:42]:      BookingStatusOverbooked,
	_BookingStatusLowerName[32:42]: BookingStatusOverbooked,
	_BookingStatusName[42:52]:      BookingStatusConfirming,
	_BookingStatusLowerName[42:52]: BookingStatusConfirming,
}

var _BookingStatusNames = []string{
	_BookingStatusName[0:7],
	_BookingStatusName[7:16],
	_BookingStatusName[16:25],
	_BookingStatusName[25:32],
	_BookingStatusName[32:42],
	_BookingStatusName[42:52],
}

// BookingStatusString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func BookingStatusString(s string) (BookingStatus, error) {
	if val, ok := _BookingStatusNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _BookingStatusNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to BookingStatus values", s)
}

// BookingStatusValues returns all values of the enum
func BookingStatusValues() []BookingStatus {
	return _BookingStatusValues
}

// BookingStatusStrings returns a slice of all String values of the enum
func BookingStatusStrings() []string {
	strs := make([]string, len(_BookingStatusNames))
	copy(strs, _BookingStatusNames)
	return strs
}

// IsABookingStatus returns "true" if the value is listed in the enum definition. "false" otherwise
func (i BookingStatus) IsABookingStatus() bool {
	for _, v := range _BookingStatusValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for BookingStatus
func (i BookingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for BookingStatus
func (i *BookingStatus) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("BookingStatus should be a string, got %s", data)
	}

	var err error
	*i, err = BookingStatusString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for BookingStatus
func (i BookingStatus) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for BookingStatus
func (i *BookingStatus) UnmarshalText(text []byte) error {
	var err error
	*i, err = BookingStatusString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for BookingStatus
func (i BookingStatus) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for BookingStatus
func (i *BookingStatus) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = BookingStatusString(s)
	return err
}

func (i BookingStatus) Value() (driver.Value, error) {
	return i.String(), nil
}

func (i *BookingStatus) Scan(value interface{}) error {
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
		return fmt.Errorf("invalid value of BookingStatus: %[1]T(%[1]v)", value)
	}

	val, err := BookingStatusString(str)
	if err != nil {
		return err
	}

	*i = val
	return nil
}
