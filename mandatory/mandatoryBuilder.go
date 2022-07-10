package mandatory

import (
	"bytes"
	"github.com/ua-parser/uap-go/uaparser"
	"strings"
	"sync"
)

// Builder builder object for Request.
type Builder struct {
	Request
	uaParser *uaparser.Parser
}

var (
	parser *uaparser.Parser
	once   sync.Once
)

// NewMandatoryRequestBuilder initialize builder object for Request.
func NewMandatoryRequestBuilder() (Builder, error) {
	var err error
	once.Do(func() {
		parser, err = uaparser.NewFromBytes(bytes.NewBufferString(regexes).Bytes())
	})
	if err != nil {
		return Builder{}, err
	}

	return Builder{
		Request:  Request{},
		uaParser: parser,
	}, nil
}

// Build initialize builder object for Request.
func (m Builder) Build() Request {
	m.valid = true
	return m.Request
}

// WithTraceID getter function to set TraceID.
func (m Builder) WithTraceID(traceID string) Builder {
	m.traceID = traceID
	return m
}

// WithIpAddresses getter function to set IpAddresses.
func (m Builder) WithIpAddresses(ipAddress []string) Builder {
	m.ipAddress = ipAddress
	return m
}

// WithAuthorization getter function to set Authorization.
func (m Builder) WithAuthorization(authorization string) Builder {
	m.authorization.authorization = authorization
	m.authorization.token = strings.ReplaceAll(authorization, "Bearer ", "")
	return m
}

// WithApiKey getter function to set ApiKey.
func (m Builder) WithApiKey(apiKey string) Builder {
	m.authorization.apiKey = apiKey
	return m
}

// WithServiceSecret getter function to set ServiceSecret.
func (m Builder) WithServiceSecret(ID, secret string) Builder {
	m.authorization.serviceID = ID
	m.authorization.serviceSecret = secret
	return m
}

// WithUserAgent getter function to set UserAgent.
func (m Builder) WithUserAgent(userAgent string) Builder {
	client := m.uaParser.Parse(userAgent)
	m.userAgent.value = userAgent
	m.userAgent.family = client.UserAgent.Family
	m.userAgent.major = client.UserAgent.Major
	m.userAgent.minor = client.UserAgent.Minor
	m.userAgent.patch = client.UserAgent.Patch
	m.os.family = client.Os.Family
	m.os.major = client.Os.Major
	m.os.minor = client.Os.Minor
	m.os.patch = client.Os.Patch
	m.os.patchMinor = client.Os.PatchMinor
	m.device.family = client.Device.Family
	m.device.brand = client.Device.Brand
	m.device.model = client.Device.Model

	if m.bringDeviceType {
		return m
	}

	switch strings.ToLower(m.os.family) {
	case "android":
		m.deviceType = Android
		break
	case "ios":
		m.deviceType = Ios
		break
	}

	if m.device.deviceID == "" {
		if strings.Contains(strings.ToLower(m.userAgent.family), "mobile") {
			m.deviceType = MobileWeb
		} else {
			m.deviceType = Web
		}
	}
	return m
}

// WithApplication getter function to set Application.
func (m Builder) WithApplication(deviceID, appsVersion string) Builder {
	m.device.deviceID = deviceID
	m.device.appVersion = appsVersion
	if m.userAgent.value != "" {
		return m.WithUserAgent(m.userAgent.value)
	}
	return m
}

// WithDeviceType getter function to set DeviceType.
func (m Builder) WithDeviceType(deviceType string) Builder {
	if "" == deviceType {
		return m
	}

	var err error
	m.deviceType, err = DeviceFromStringCode(deviceType)
	if err != nil {
		return m
	}

	m.bringDeviceType = true
	return m
}

// WithUser getter function to set User.
func (m Builder) WithUser(ID uint64, email string) Builder {
	m.user.login = true
	m.user.id = ID
	m.user.email = email
	return m
}

// WithUserPhone getter function to set UserPhone.
func (m Builder) WithUserPhone(ID uint64, email string, phone string) Builder {
	m.user.login = true
	m.user.id = ID
	m.user.email = email
	m.user.phone = phone
	return m
}

// WithPhone getter function to set Phone.
func (m Builder) WithPhone(ID uint64, phone string) Builder {
	m.user.login = true
	m.user.id = ID
	m.user.phone = phone
	return m
}

// WithLanguage getter function to set Language.
func (m Builder) WithLanguage(language string) Builder {
	m.language = language
	return m
}
