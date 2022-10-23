package mandatory

import (
	"fmt"
	"github.com/paulusrobin/gogen-golib/encoding/json"
	"strings"
)

type (
	// DeviceType enum of device type.
	DeviceType int

	// DeviceInfo object device.
	DeviceInfo struct {
		id   DeviceType
		code string
		name string
	}
)

// ID getter function to get id.
func (d DeviceInfo) ID() DeviceType {
	return d.id
}

// Name getter function to get name.
func (d DeviceInfo) Name() string {
	return d.name
}

// Code getter function to get code.
func (d DeviceInfo) Code() string {
	return d.code
}

// Valid getter function to get code.
func (d DeviceInfo) Valid() bool {
	idx := int(d.id) - 1
	if idx < 0 || idx >= len(deviceInfos) {
		return false
	}
	return true
}

const (
	Android DeviceType = iota + 1
	Ios
	Web
	MobileWeb
)

var deviceInfos = []DeviceInfo{
	{Android, "ANDROID", "Android"},
	{Ios, "IOS", "Ios"},
	{Web, "WEB", "Website"},
	{MobileWeb, "MWEB", "Mobile Website"},
}

// Info function to get DeviceInfo.
func (s DeviceType) Info() DeviceInfo {
	if int(s) < 0 || int(s) >= len(deviceInfos)-1 {
		return DeviceInfo{
			id:   0,
			code: "Unknown",
			name: "Unknown",
		}
	}
	return deviceInfos[s-1]
}

// MarshalJSON implements the marshaller interface.
func (s DeviceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Info().Name())
}

// UnmarshalJSON implements the un marshaller interface.
func (s *DeviceType) UnmarshalJSON(b []byte) error {
	var strAction = ""
	err := json.Unmarshal(b, &strAction)
	if err != nil {
		return err
	}
	*s, err = DeviceFromString(strAction)
	return err
}

// ErrInvalidDeviceType error invalid device type.
var ErrInvalidDeviceType = fmt.Errorf("invalid device")

// DeviceFromString function to get DeviceType from string.
func DeviceFromString(str string) (DeviceType, error) {
	lowerStr := strings.ToLower(str)
	for i, j := 0, len(deviceInfos)-1; i <= j; i, j = i+1, j-1 {
		if strings.ToLower(deviceInfos[i].name) == lowerStr {
			return DeviceType(i + 1), nil
		}
		if strings.ToLower(deviceInfos[j].name) == lowerStr {
			return DeviceType(j + 1), nil
		}
	}
	return -1, ErrInvalidDeviceType
}

// DeviceFromStringCode function to get DeviceType from string code.
func DeviceFromStringCode(str string) (DeviceType, error) {
	lowerStr := strings.ToLower(str)
	for i, j := 0, len(deviceInfos)-1; i <= j; i, j = i+1, j-1 {
		if strings.ToLower(deviceInfos[i].code) == lowerStr {
			return DeviceType(i + 1), nil
		}
		if strings.ToLower(deviceInfos[j].code) == lowerStr {
			return DeviceType(j + 1), nil
		}
	}
	return -1, ErrInvalidDeviceType
}
