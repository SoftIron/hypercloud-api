// Code generated by "enumer -type NetworkRecovery -linecomment -text"; DO NOT EDIT.

package cloud

import (
	"fmt"
	"strings"
)

const _NetworkRecoveryName = "failuresuccessretry_vnetdelete_vnet"

var _NetworkRecoveryIndex = [...]uint8{0, 7, 14, 24, 35}

const _NetworkRecoveryLowerName = "failuresuccessretry_vnetdelete_vnet"

func (i NetworkRecovery) String() string {
	if i < 0 || i >= NetworkRecovery(len(_NetworkRecoveryIndex)-1) {
		return fmt.Sprintf("NetworkRecovery(%d)", i)
	}
	return _NetworkRecoveryName[_NetworkRecoveryIndex[i]:_NetworkRecoveryIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _NetworkRecoveryNoOp() {
	var x [1]struct{}
	_ = x[FailureNetworkRecovery-(0)]
	_ = x[SuccessNetworkRecovery-(1)]
	_ = x[RetryNetworkRecovery-(2)]
	_ = x[DeleteNetworkRecovery-(3)]
}

var _NetworkRecoveryValues = []NetworkRecovery{FailureNetworkRecovery, SuccessNetworkRecovery, RetryNetworkRecovery, DeleteNetworkRecovery}

var _NetworkRecoveryNameToValueMap = map[string]NetworkRecovery{
	_NetworkRecoveryName[0:7]:        FailureNetworkRecovery,
	_NetworkRecoveryLowerName[0:7]:   FailureNetworkRecovery,
	_NetworkRecoveryName[7:14]:       SuccessNetworkRecovery,
	_NetworkRecoveryLowerName[7:14]:  SuccessNetworkRecovery,
	_NetworkRecoveryName[14:24]:      RetryNetworkRecovery,
	_NetworkRecoveryLowerName[14:24]: RetryNetworkRecovery,
	_NetworkRecoveryName[24:35]:      DeleteNetworkRecovery,
	_NetworkRecoveryLowerName[24:35]: DeleteNetworkRecovery,
}

var _NetworkRecoveryNames = []string{
	_NetworkRecoveryName[0:7],
	_NetworkRecoveryName[7:14],
	_NetworkRecoveryName[14:24],
	_NetworkRecoveryName[24:35],
}

// NetworkRecoveryString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func NetworkRecoveryString(s string) (NetworkRecovery, error) {
	if val, ok := _NetworkRecoveryNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _NetworkRecoveryNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to NetworkRecovery values", s)
}

// NetworkRecoveryValues returns all values of the enum
func NetworkRecoveryValues() []NetworkRecovery {
	return _NetworkRecoveryValues
}

// NetworkRecoveryStrings returns a slice of all String values of the enum
func NetworkRecoveryStrings() []string {
	strs := make([]string, len(_NetworkRecoveryNames))
	copy(strs, _NetworkRecoveryNames)
	return strs
}

// IsANetworkRecovery returns "true" if the value is listed in the enum definition. "false" otherwise
func (i NetworkRecovery) IsANetworkRecovery() bool {
	for _, v := range _NetworkRecoveryValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for NetworkRecovery
func (i NetworkRecovery) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for NetworkRecovery
func (i *NetworkRecovery) UnmarshalText(text []byte) error {
	var err error
	*i, err = NetworkRecoveryString(string(text))
	return err
}