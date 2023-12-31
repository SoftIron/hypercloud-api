// Code generated by "enumer -type InstanceState -linecomment -text"; DO NOT EDIT.

package hc

import (
	"fmt"
	"strings"
)

const (
	_InstanceStateName_0      = "allnot_doneinitpendingholdactivestoppedsuspendeddone"
	_InstanceStateLowerName_0 = "allnot_doneinitpendingholdactivestoppedsuspendeddone"
	_InstanceStateName_1      = "offundeployedcloningcloning_failed"
	_InstanceStateLowerName_1 = "offundeployedcloningcloning_failed"
)

var (
	_InstanceStateIndex_0 = [...]uint8{0, 3, 11, 15, 22, 26, 32, 39, 48, 52}
	_InstanceStateIndex_1 = [...]uint8{0, 3, 13, 20, 34}
)

func (i InstanceState) String() string {
	switch {
	case -2 <= i && i <= 6:
		i -= -2
		return _InstanceStateName_0[_InstanceStateIndex_0[i]:_InstanceStateIndex_0[i+1]]
	case 8 <= i && i <= 11:
		i -= 8
		return _InstanceStateName_1[_InstanceStateIndex_1[i]:_InstanceStateIndex_1[i+1]]
	default:
		return fmt.Sprintf("InstanceState(%d)", i)
	}
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _InstanceStateNoOp() {
	var x [1]struct{}
	_ = x[AnyState-(-2)]
	_ = x[NotDoneState-(-1)]
	_ = x[InitState-(0)]
	_ = x[PendingState-(1)]
	_ = x[HoldState-(2)]
	_ = x[ActiveState-(3)]
	_ = x[StoppedState-(4)]
	_ = x[SuspendedState-(5)]
	_ = x[DoneState-(6)]
	_ = x[OffState-(8)]
	_ = x[UndeployedState-(9)]
	_ = x[CloningState-(10)]
	_ = x[CloningFailedState-(11)]
}

var _InstanceStateValues = []InstanceState{AnyState, NotDoneState, InitState, PendingState, HoldState, ActiveState, StoppedState, SuspendedState, DoneState, OffState, UndeployedState, CloningState, CloningFailedState}

var _InstanceStateNameToValueMap = map[string]InstanceState{
	_InstanceStateName_0[0:3]:        AnyState,
	_InstanceStateLowerName_0[0:3]:   AnyState,
	_InstanceStateName_0[3:11]:       NotDoneState,
	_InstanceStateLowerName_0[3:11]:  NotDoneState,
	_InstanceStateName_0[11:15]:      InitState,
	_InstanceStateLowerName_0[11:15]: InitState,
	_InstanceStateName_0[15:22]:      PendingState,
	_InstanceStateLowerName_0[15:22]: PendingState,
	_InstanceStateName_0[22:26]:      HoldState,
	_InstanceStateLowerName_0[22:26]: HoldState,
	_InstanceStateName_0[26:32]:      ActiveState,
	_InstanceStateLowerName_0[26:32]: ActiveState,
	_InstanceStateName_0[32:39]:      StoppedState,
	_InstanceStateLowerName_0[32:39]: StoppedState,
	_InstanceStateName_0[39:48]:      SuspendedState,
	_InstanceStateLowerName_0[39:48]: SuspendedState,
	_InstanceStateName_0[48:52]:      DoneState,
	_InstanceStateLowerName_0[48:52]: DoneState,
	_InstanceStateName_1[0:3]:        OffState,
	_InstanceStateLowerName_1[0:3]:   OffState,
	_InstanceStateName_1[3:13]:       UndeployedState,
	_InstanceStateLowerName_1[3:13]:  UndeployedState,
	_InstanceStateName_1[13:20]:      CloningState,
	_InstanceStateLowerName_1[13:20]: CloningState,
	_InstanceStateName_1[20:34]:      CloningFailedState,
	_InstanceStateLowerName_1[20:34]: CloningFailedState,
}

var _InstanceStateNames = []string{
	_InstanceStateName_0[0:3],
	_InstanceStateName_0[3:11],
	_InstanceStateName_0[11:15],
	_InstanceStateName_0[15:22],
	_InstanceStateName_0[22:26],
	_InstanceStateName_0[26:32],
	_InstanceStateName_0[32:39],
	_InstanceStateName_0[39:48],
	_InstanceStateName_0[48:52],
	_InstanceStateName_1[0:3],
	_InstanceStateName_1[3:13],
	_InstanceStateName_1[13:20],
	_InstanceStateName_1[20:34],
}

// InstanceStateString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func InstanceStateString(s string) (InstanceState, error) {
	if val, ok := _InstanceStateNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _InstanceStateNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to InstanceState values", s)
}

// InstanceStateValues returns all values of the enum
func InstanceStateValues() []InstanceState {
	return _InstanceStateValues
}

// InstanceStateStrings returns a slice of all String values of the enum
func InstanceStateStrings() []string {
	strs := make([]string, len(_InstanceStateNames))
	copy(strs, _InstanceStateNames)
	return strs
}

// IsAInstanceState returns "true" if the value is listed in the enum definition. "false" otherwise
func (i InstanceState) IsAInstanceState() bool {
	for _, v := range _InstanceStateValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for InstanceState
func (i InstanceState) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for InstanceState
func (i *InstanceState) UnmarshalText(text []byte) error {
	var err error
	*i, err = InstanceStateString(string(text))
	return err
}
