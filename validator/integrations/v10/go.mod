module github.com/paulusrobin/gogen-golib/validator/integrations/v10

go 1.18

require (
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.11.0
	github.com/paulusrobin/gogen-golib/validator/interface v0.0.0-20220709001715-f945a568eddb
	github.com/rs/zerolog v1.27.0
)

require (
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	golang.org/x/crypto v0.0.0-20211215153901-e495a2d5b3d3 // indirect
	golang.org/x/sys v0.0.0-20210927094055-39ccf1dd6fa6 // indirect
	golang.org/x/text v0.3.7 // indirect
)

replace github.com/paulusrobin/gogen-golib/validator/interface => ./../../interface
