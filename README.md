# Gofiber Scalar

Scalar middleware for [Fiber](https://github.com/gofiber/fiber). The middleware handles Scalar UI.

**Note: Requires Go 1.23.0 and above**

### Table of Contents
- [Signatures](#signatures)
- [Installation](#installation)
- [Examples](#examples)
- [Config](#config)
- [Default Config](#default-config)

### Signatures
```go
func New(config ...scalar.Config) fiber.Handler
```

### Installation
Scalar is tested on the latest [Go versions](https://golang.org/dl/) with support for modules. So make sure to initialize one first if you didn't do that yet:
```bash
go mod init github.com/<user>/<repo>
```
And then install the Scalar middleware:
```bash
go get -u github.com/yokeTH/gofiber-scalar/scalar
```

### Examples
Using swaggo to generate documents default output path is `(root)/docs`:
```bash
swag init -v3.1
```

Import the middleware package and generated docs
```go
import (
  "YOUR_MODULE/docs"

  "github.com/gofiber/fiber/v2"
  "github.com/yokeTH/gofiber-scalar/scalar"
  "github.com/swaggo/swag" // for register docs DO NOT USE v2
)
```

After Imported:
Register Swag the Docs
```go
swag.Register(docs.SwaggerInfo.InstanceName(), docs.SwaggerInfo)
```

Using the default config:
```go
app.Use(scalar.New())
```
Now you can access scalar API documentation UI at `{HOSTNAME}/docs` and JSON documentation at `{HOSTNAME}/docs/doc.json`. Additionally, you can modify the path by configuring the middleware to suit your application's requirements.

Using as the handler: for an example `localhost:8080/yourpath`

```go
app.Get("/yourpath/*", scalar.New(scalar.Config{
	Path:     "/yourpath",
}))
```

Using a custom config:
```go
cfg := scalar.Config{
    BasePath: "/",
    Path:     "swagger", // replace original swagger path
    Title:    "Your app API Docs",
}

app.Use(scalar.New(cfg))
```

Use program data for Swagger content:
```go
cfg := scalar.Config{
    BasePath:          "/",
    FileContentString: jsonString,
    Path:              "scalar",
    Title:             "Scalar API Docs",
}

app.Use(scalar.New(cfg))
```

### Config
```go
type Config struct {
	// Next defines a function to skip this middleware when returned true.
	//
	// Optional. Default: nil
	Next func(c *fiber.Ctx) bool

	// BasePath for the UI path
	//
	// Optional. Default: /
	BasePath string

	// FileContent for the content of the swagger.json or swagger.yaml file.
	//
	// Optional. Default: nil
	FileContentString string

	// Path combines with BasePath for the full UI path
	//
	// Optional. Default: docs
	Path string

	// Title for the documentation site
	//
	// Optional. Default: Fiber API documentation
	Title string

	// CacheAge defines the max-age for the Cache-Control header in seconds.
	//
	// Optional. Default: 1 min (no cache)
	CacheAge int

	// Custom Scalar Style
	// Ref: https://github.com/scalar/scalar/blob/main/packages/themes/src/variables.css
	// Optional. Default: ""
	CustomStyle template.CSS

	// Proxy to avoid CORS issues
	// Optional.
	ProxyUrl string

	// Raw Space Url
	// Optional. Default: doc.json
	RawSpecUrl string

	// ForceOffline
	// Optional: Default: true
	ForceOffline bool

	// Fallback scalar cache
	//
	// Optional. Default: 86400 (1 Days)
	FallbackCacheAge int
}
```

### Default Config
```go
var configDefault = Config{
	Next:             nil,
	BasePath:         "/",
	Path:             "docs",
	Title:            "Fiber API documentation",
	CacheAge:         60,
	RawSpecUrl:       "doc.json",
	ForceOffline:     true,
	FallbackCacheAge: 86400,
}
```
