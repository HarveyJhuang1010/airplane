// Code generated by "enumer -type=UserLoginSessionStep -trimprefix=UserLoginSessionStep -yaml -json -text -transform=snake --output=zzz_enumer_userLoginSessionStep.go"; DO NOT EDIT.

package enum

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _UserLoginSessionStepName = "identifyotp_resendotp_verifiedfinish"

var _UserLoginSessionStepIndex = [...]uint8{0, 8, 18, 30, 36}

const _UserLoginSessionStepLowerName = "identifyotp_resendotp_verifiedfinish"

func (i UserLoginSessionStep) String() string {
	i -= 1
	if i < 0 || i >= UserLoginSessionStep(len(_UserLoginSessionStepIndex)-1) {
		return fmt.Sprintf("UserLoginSessionStep(%d)", i+1)
	}
	return _UserLoginSessionStepName[_UserLoginSessionStepIndex[i]:_UserLoginSessionStepIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _UserLoginSessionStepNoOp() {
	var x [1]struct{}
	_ = x[UserLoginSessionStepIdentify-(1)]
	_ = x[UserLoginSessionStepOtpResend-(2)]
	_ = x[UserLoginSessionStepOtpVerified-(3)]
	_ = x[UserLoginSessionStepFinish-(4)]
}

var _UserLoginSessionStepValues = []UserLoginSessionStep{UserLoginSessionStepIdentify, UserLoginSessionStepOtpResend, UserLoginSessionStepOtpVerified, UserLoginSessionStepFinish}

var _UserLoginSessionStepNameToValueMap = map[string]UserLoginSessionStep{
	_UserLoginSessionStepName[0:8]:        UserLoginSessionStepIdentify,
	_UserLoginSessionStepLowerName[0:8]:   UserLoginSessionStepIdentify,
	_UserLoginSessionStepName[8:18]:       UserLoginSessionStepOtpResend,
	_UserLoginSessionStepLowerName[8:18]:  UserLoginSessionStepOtpResend,
	_UserLoginSessionStepName[18:30]:      UserLoginSessionStepOtpVerified,
	_UserLoginSessionStepLowerName[18:30]: UserLoginSessionStepOtpVerified,
	_UserLoginSessionStepName[30:36]:      UserLoginSessionStepFinish,
	_UserLoginSessionStepLowerName[30:36]: UserLoginSessionStepFinish,
}

var _UserLoginSessionStepNames = []string{
	_UserLoginSessionStepName[0:8],
	_UserLoginSessionStepName[8:18],
	_UserLoginSessionStepName[18:30],
	_UserLoginSessionStepName[30:36],
}

// UserLoginSessionStepString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func UserLoginSessionStepString(s string) (UserLoginSessionStep, error) {
	if val, ok := _UserLoginSessionStepNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _UserLoginSessionStepNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to UserLoginSessionStep values", s)
}

// UserLoginSessionStepValues returns all values of the enum
func UserLoginSessionStepValues() []UserLoginSessionStep {
	return _UserLoginSessionStepValues
}

// UserLoginSessionStepStrings returns a slice of all String values of the enum
func UserLoginSessionStepStrings() []string {
	strs := make([]string, len(_UserLoginSessionStepNames))
	copy(strs, _UserLoginSessionStepNames)
	return strs
}

// IsAUserLoginSessionStep returns "true" if the value is listed in the enum definition. "false" otherwise
func (i UserLoginSessionStep) IsAUserLoginSessionStep() bool {
	for _, v := range _UserLoginSessionStepValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for UserLoginSessionStep
func (i UserLoginSessionStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for UserLoginSessionStep
func (i *UserLoginSessionStep) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("UserLoginSessionStep should be a string, got %s", data)
	}

	var err error
	*i, err = UserLoginSessionStepString(s)
	return err
}

// MarshalText implements the encoding.TextMarshaler interface for UserLoginSessionStep
func (i UserLoginSessionStep) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for UserLoginSessionStep
func (i *UserLoginSessionStep) UnmarshalText(text []byte) error {
	var err error
	*i, err = UserLoginSessionStepString(string(text))
	return err
}

// MarshalYAML implements a YAML Marshaler for UserLoginSessionStep
func (i UserLoginSessionStep) MarshalYAML() (interface{}, error) {
	return i.String(), nil
}

// UnmarshalYAML implements a YAML Unmarshaler for UserLoginSessionStep
func (i *UserLoginSessionStep) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	if err := unmarshal(&s); err != nil {
		return err
	}

	var err error
	*i, err = UserLoginSessionStepString(s)
	return err
}
