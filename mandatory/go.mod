module github.com/paulusrobin/gogen-golib/mandatory

go 1.18

require (
	github.com/paulusrobin/gogen-golib/encoding/json v1.0.0
	github.com/ua-parser/uap-go v0.0.0-20211112212520-00c877edfe0f
)

require (
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	gopkg.in/yaml.v2 v2.2.1 // indirect
)

replace github.com/paulusrobin/gogen-golib/encoding/json => ./../encoding/json
