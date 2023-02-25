module github.com/mchmarny/klogo

go 1.13

require (
	cloud.google.com/go v0.46.3
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/json-iterator/go v1.1.7 // indirect
	github.com/mattn/go-isatty v0.0.9 // indirect
	github.com/mchmarny/gcputil v0.2.1
	github.com/stretchr/testify v1.4.0
	github.com/ugorji/go/codec v0.0.0-20190128213124-ee1426cffec0 // indirect
	golang.org/x/sys v0.1.0 // indirect
	gopkg.in/yaml.v2 v2.2.4 // indirect
)

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
