// Code generated by "stringer -type=VPType"; DO NOT EDIT.

package gno

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[VPUverse-0]
	_ = x[VPBlock-1]
	_ = x[VPField-2]
	_ = x[VPValMethod-3]
	_ = x[VPPtrMethod-4]
	_ = x[VPInterface-5]
	_ = x[VPSubrefField-6]
	_ = x[VPDerefField-18]
	_ = x[VPDerefValMethod-19]
	_ = x[VPDerefPtrMethod-20]
	_ = x[VPDerefInterface-21]
	_ = x[VPNative-32]
}

const (
	_VPType_name_0 = "VPUverseVPBlockVPFieldVPValMethodVPPtrMethodVPInterfaceVPSubrefField"
	_VPType_name_1 = "VPDerefFieldVPDerefValMethodVPDerefPtrMethodVPDerefInterface"
	_VPType_name_2 = "VPNative"
)

var (
	_VPType_index_0 = [...]uint8{0, 8, 15, 22, 33, 44, 55, 68}
	_VPType_index_1 = [...]uint8{0, 12, 28, 44, 60}
)

func (i VPType) String() string {
	switch {
	case 0 <= i && i <= 6:
		return _VPType_name_0[_VPType_index_0[i]:_VPType_index_0[i+1]]
	case 18 <= i && i <= 21:
		i -= 18
		return _VPType_name_1[_VPType_index_1[i]:_VPType_index_1[i+1]]
	case i == 32:
		return _VPType_name_2
	default:
		return "VPType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
