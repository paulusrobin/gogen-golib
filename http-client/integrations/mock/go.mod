module github.com/paulusrobin/gogen-golib/http-client/mock

go 1.18

require github.com/paulusrobin/gogen-golib/http-client latest

replace (
	github.com/paulusrobin/gogen-golib/http-client latest => ./../../interface
)
