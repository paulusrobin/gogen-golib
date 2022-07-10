package mandatory

import (
	"bytes"
	"github.com/ua-parser/uap-go/uaparser"
	"strings"
)

type Builder struct {
	Request
	uaParser *uaparser.Parser
}

func NewMandatoryRequestBuilder() (Builder, error) {
	parser, err := uaparser.NewFromBytes(bytes.NewBufferString(regexes).Bytes())
	if err != nil {
		return Builder{}, err
	}

	return Builder{
		Request:  Request{},
		uaParser: parser,
	}, nil
}

func (m Builder) Build() Request {
	m.valid = true
	return m.Request
}

func (m Builder) WithTraceID(traceID string) Builder {
	m.traceID = traceID
	return m
}

func (m Builder) WithIpAddresses(ipAddress []string) Builder {
	m.ipAddress = ipAddress
	return m
}

func (m Builder) WithAuthorization(authorization string) Builder {
	m.authorization.authorization = authorization
	m.authorization.token = strings.ReplaceAll(authorization, "Bearer ", "")
	return m
}

func (m Builder) WithApiKey(apiKey string) Builder {
	m.authorization.apiKey = apiKey
	return m
}

func (m Builder) WithServiceSecret(ID, secret string) Builder {
	m.authorization.serviceID = ID
	m.authorization.serviceSecret = secret
	return m
}

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

func (m Builder) WithApplication(deviceID, appsVersion string) Builder {
	m.device.deviceID = deviceID
	m.device.appVersion = appsVersion
	if m.userAgent.value != "" {
		return m.WithUserAgent(m.userAgent.value)
	}
	return m
}

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

func (m Builder) WithUser(ID uint64, email string) Builder {
	m.user.login = true
	m.user.id = ID
	m.user.email = email
	return m
}

func (m Builder) WithUserPhone(ID uint64, email string, phone string) Builder {
	m.user.login = true
	m.user.id = ID
	m.user.email = email
	m.user.phone = phone
	return m
}

func (m Builder) WithPhone(ID uint64, phone string) Builder {
	m.user.login = true
	m.user.id = ID
	m.user.phone = phone
	return m
}

func (m Builder) WithLanguage(language string) Builder {
	m.language = language
	return m
}
