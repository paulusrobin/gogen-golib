module github.com/paulusrobin/gogen-golib/http-client/mock

go 1.18

require github.com/paulusrobin/gogen-golib/http-client/interface eb97a40a8f327e107d464202d244676d4c389124

replace (
	github.com/paulusrobin/gogen-golib/http-client/interface => ./../../interface
)
