// Code generated by "enumer -type ImageType -transform upper -text"; DO NOT EDIT.

package cloud

import (
	"fmt"
	"strings"
)

const _ImageTypeName = "OSCDROMDATABLOCKKERNELRAMDISKCONTEXT"

var _ImageTypeIndex = [...]uint8{0, 2, 7, 16, 22, 29, 36}

const _ImageTypeLowerName = "oscdromdatablockkernelramdiskcontext"

func (i ImageType) String() string {
	i -= 1
	if i < 0 || i >= ImageType(len(_ImageTypeIndex)-1) {
		return fmt.Sprintf("ImageType(%d)", i+1)
	}
	return _ImageTypeName[_ImageTypeIndex[i]:_ImageTypeIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ImageTypeNoOp() {
	var x [1]struct{}
	_ = x[OS-(1)]
	_ = x[CDROM-(2)]
	_ = x[DataBlock-(3)]
	_ = x[Kernel-(4)]
	_ = x[RAMDisk-(5)]
	_ = x[Context-(6)]
}

var _ImageTypeValues = []ImageType{OS, CDROM, DataBlock, Kernel, RAMDisk, Context}

var _ImageTypeNameToValueMap = map[string]ImageType{
	_ImageTypeName[0:2]:        OS,
	_ImageTypeLowerName[0:2]:   OS,
	_ImageTypeName[2:7]:        CDROM,
	_ImageTypeLowerName[2:7]:   CDROM,
	_ImageTypeName[7:16]:       DataBlock,
	_ImageTypeLowerName[7:16]:  DataBlock,
	_ImageTypeName[16:22]:      Kernel,
	_ImageTypeLowerName[16:22]: Kernel,
	_ImageTypeName[22:29]:      RAMDisk,
	_ImageTypeLowerName[22:29]: RAMDisk,
	_ImageTypeName[29:36]:      Context,
	_ImageTypeLowerName[29:36]: Context,
}

var _ImageTypeNames = []string{
	_ImageTypeName[0:2],
	_ImageTypeName[2:7],
	_ImageTypeName[7:16],
	_ImageTypeName[16:22],
	_ImageTypeName[22:29],
	_ImageTypeName[29:36],
}

// ImageTypeString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ImageTypeString(s string) (ImageType, error) {
	if val, ok := _ImageTypeNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ImageTypeNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ImageType values", s)
}

// ImageTypeValues returns all values of the enum
func ImageTypeValues() []ImageType {
	return _ImageTypeValues
}

// ImageTypeStrings returns a slice of all String values of the enum
func ImageTypeStrings() []string {
	strs := make([]string, len(_ImageTypeNames))
	copy(strs, _ImageTypeNames)
	return strs
}

// IsAImageType returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ImageType) IsAImageType() bool {
	for _, v := range _ImageTypeValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalText implements the encoding.TextMarshaler interface for ImageType
func (i ImageType) MarshalText() ([]byte, error) {
	return []byte(i.String()), nil
}

// UnmarshalText implements the encoding.TextUnmarshaler interface for ImageType
func (i *ImageType) UnmarshalText(text []byte) error {
	var err error
	*i, err = ImageTypeString(string(text))
	return err
}
