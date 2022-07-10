package mandatory

type OS struct {
	name    string
	version string

	family     string
	major      string
	minor      string
	patch      string
	patchMinor string
}

func (o OS) Name() string {
	return o.name
}

func (o OS) Version() string {
	return o.version
}

func (o OS) Family() string {
	return o.family
}

func (o OS) Major() string {
	return o.major
}

func (o OS) Minor() string {
	return o.minor
}

func (o OS) Patch() string {
	return o.patch
}

func (o OS) PatchMinor() string {
	return o.patchMinor
}
