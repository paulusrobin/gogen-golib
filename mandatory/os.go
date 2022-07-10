package mandatory

// OS object of mandatory OS.
type OS struct {
	name    string
	version string

	family     string
	major      string
	minor      string
	patch      string
	patchMinor string
}

// Name getter function to get OS name.
func (o OS) Name() string {
	return o.name
}

// Version getter function to get OS version.
func (o OS) Version() string {
	return o.version
}

// Family getter function to get OS family.
func (o OS) Family() string {
	return o.family
}

// Major getter function to get OS major version.
func (o OS) Major() string {
	return o.major
}

// Minor getter function to get OS minor version.
func (o OS) Minor() string {
	return o.minor
}

// Patch getter function to get OS patch version.
func (o OS) Patch() string {
	return o.patch
}

// PatchMinor getter function to get OS patch minor version.
func (o OS) PatchMinor() string {
	return o.patchMinor
}
