package mandatory

// UserAgent object of mandatory UserAgent.
type UserAgent struct {
	value  string
	family string
	major  string
	minor  string
	patch  string
}

// Value getter function to get mandatory UserAgent Value.
func (u UserAgent) Value() string {
	return u.value
}

// Family getter function to get mandatory UserAgent Family.
func (u UserAgent) Family() string {
	return u.family
}

// Major getter function to get mandatory UserAgent Major.
func (u UserAgent) Major() string {
	return u.major
}

// Minor getter function to get mandatory UserAgent Minor.
func (u UserAgent) Minor() string {
	return u.minor
}

// Patch getter function to get mandatory UserAgent Patch.
func (u UserAgent) Patch() string {
	return u.patch
}
