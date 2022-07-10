package mandatory

type Request struct {
	traceID         string
	ipAddress       []string
	valid           bool
	bringDeviceType bool
	language        string

	authorization Authorization
	deviceType    DeviceType
	device        Device
	userAgent     UserAgent
	os            OS
	user          User
}

func (m Request) TraceID() string {
	return m.traceID
}

func (m Request) IpAddresses() []string {
	return m.ipAddress
}

func (m Request) Language() string {
	if m.language == "" {
		return "en"
	}
	return m.language
}

func (m Request) Authorization() Authorization {
	return m.authorization
}

func (m Request) DeviceType() DeviceType {
	return m.deviceType
}

func (m Request) Device() Device {
	return m.device
}

func (m Request) UserAgent() UserAgent {
	return m.userAgent
}

func (m Request) OS() OS {
	return m.os
}

func (m Request) User() User {
	return m.user
}

/*
===========================
	Utilities Function
===========================
*/
func (m Request) Valid() bool {
	return m.valid
}

func (m Request) IsUserLogin() bool {
	return m.user.IsLogin()
}

func (m Request) IsMobileApp() bool {
	return Android == m.deviceType || Ios == m.deviceType
}

func (m Request) IsWebApp() bool {
	return Web == m.deviceType || MobileWeb == m.deviceType
}
