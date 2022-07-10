package mandatory

// Device object of mandatory request.
type Device struct {
	appVersion string
	deviceID   string
	family     string
	brand      string
	model      string
}

// AppVersion getter function to get mandatory Device AppVersion.
func (d Device) AppVersion() string {
	return d.appVersion
}

// DeviceID getter function to get mandatory Device DeviceID.
func (d Device) DeviceID() string {
	return d.deviceID
}

// Family getter function to get mandatory Device Family.
func (d Device) Family() string {
	return d.family
}

// Brand getter function to get mandatory Device Brand.
func (d Device) Brand() string {
	return d.brand
}

// Model getter function to get mandatory Device Model.
func (d Device) Model() string {
	return d.model
}
