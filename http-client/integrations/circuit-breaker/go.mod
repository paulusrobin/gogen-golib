module github.com/paulusrobin/gogen-golib/http-client/integrations/circuit-breaker

go 1.18

require (
	github.com/paulusrobin/gogen-golib/http-client/interface v0.0.0-20220708235924-9c2ce8534e56
	github.com/sony/gobreaker v0.5.0
)

require (
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/paulusrobin/gogen-golib/encoding/json v1.0.0 // indirect
	github.com/rs/zerolog v1.27.0 // indirect
	golang.org/x/sys v0.0.0-20210927094055-39ccf1dd6fa6 // indirect
)

replace github.com/paulusrobin/gogen-golib/http-client/interface => ./../../interface
