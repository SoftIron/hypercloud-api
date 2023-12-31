// Code generated by "enumer -type Period -transform lower -text"; DO NOT EDIT.

package snapshot

import (
	"fmt"
	"strings"
)

const _PeriodName = "minutelyhourlydailyweeklymonthlyyearly"

var _PeriodIndex = [...]uint8{0, 8, 14, 19, 25, 32, 38}

const _PeriodLowerName = "minutelyhourlydailyweeklymonthlyyearly"

func (i Period) String() string {
	i -= 1
	if i < 0 || i >= Period(len(_PeriodIndex)-1) {
		return fmt.Sprintf("Period(%d)", i+1)
	}
	return _PeriodName[_PeriodIndex[i]:_PeriodIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _PeriodNoOp() {
	var x [1]struct{}
	_ = x[Minutely-(1)]
	_ = x[Hourly-(2)]
	_ = x[Daily-(3)]
	_ = x[Weekly-(4)]
	_ = x[Monthly-(5)]
	_ = x[Yearly-(6)]
}

var _PeriodValues = []Period{Minutely, Hourly, Daily, Weekly, Monthly, Yearly}

var _PeriodNameToValueMap = map[string]Period{
	_PeriodName[0:8]:        Minutely,
	_PeriodLowerName[0:8]:   Minutely,
	_PeriodName[8:14]:       Hourly,
	_PeriodLowerName[8:14]:  Hourly,
	_PeriodName[14:19]:      Daily,
	_PeriodLowerName[14:19]: Daily,
	_PeriodName[19:25]:      Weekly,
	_PeriodLowerName[19:25]: Weekly,
	_PeriodName[25:32]:      Monthly,
	_PeriodLowerName[25:32]: Monthly,
	_PeriodName[32:38]:      Yearly,
	_PeriodLowerName[32:38]: Yearly,
}

var _PeriodNames = []string{
	_PeriodName[0:8],
	_PeriodName[8:14],
	_PeriodName[14:19],
	_PeriodName[19:25],
	_PeriodName[25:32],
	_PeriodName[32:38],
}

// PeriodString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func PeriodString(s string) (Period, error) {
	if val, ok := _PeriodNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _PeriodNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Period values", s)
}

// PeriodValues returns all values of the enum
func PeriodValues() []Period {
	return _PeriodValues
}

// PeriodStrings returns a slice of all String values of the enum
func PeriodStrings() []string {
	strs := make([]string, len(_PeriodNames))
	copy(strs, _PeriodNames)
	return strs
}

// IsAPeriod returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Period) IsAPeriod() bool {
	for _, v := range _PeriodValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for Period
func (i Period) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for Period
func (i *Period) UnmarshalText(text []byte) error {
	var err error
	*i, err = PeriodString(string(text))
	return err
}
