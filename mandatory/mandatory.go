package mandatory

// Request object of mandatory request.
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

// TraceID getter function to get trace id.
func (m Request) TraceID() string {
	return m.traceID
}

// IpAddresses getter function to get ip addresses.
func (m Request) IpAddresses() []string {
	return m.ipAddress
}

// Language getter function to get language.
func (m Request) Language() string {
	if m.language == "" {
		return "en"
	}
	return m.language
}

// Authorization getter function to get authorization.
func (m Request) Authorization() Authorization {
	return m.authorization
}

// DeviceType getter function to get device type.
func (m Request) DeviceType() DeviceType {
	return m.deviceType
}

// Device getter function to get device.
func (m Request) Device() Device {
	return m.device
}

// UserAgent getter function to get user agent.
func (m Request) UserAgent() UserAgent {
	return m.userAgent
}

// OS getter function to get os.
func (m Request) OS() OS {
	return m.os
}

// User getter function to get user.
func (m Request) User() User {
	return m.user
}

/*
===========================
	Utilities Function
===========================
*/

// Valid getter function to check Request object is valid.
func (m Request) Valid() bool {
	return m.valid
}

// IsUserLogin getter function to check Request object is Login User.
func (m Request) IsUserLogin() bool {
	return m.user.IsLogin()
}

// IsMobileApp getter function to check Request object is MobileApp.
func (m Request) IsMobileApp() bool {
	return Android == m.deviceType || Ios == m.deviceType
}

// IsWebApp getter function to check Request object is WebApp.
func (m Request) IsWebApp() bool {
	return Web == m.deviceType || MobileWeb == m.deviceType
}
