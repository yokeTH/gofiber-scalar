package scalar

import (
	"html/template"

	"github.com/gofiber/fiber/v2"
)

// Config defines the config for middleware.
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
	// Optional. Default: 1 min
	CacheAge int

	// Scalar theme
	//
	// Optional. Default: ThemeNone
	Theme Theme

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
	ForceOffline *bool

	// Fallback scalar cache
	//
	// Optional. Default: 86400 (1 Days)
	FallbackCacheAge int
}

var configDefault = Config{
	Next:             nil,
	BasePath:         "/",
	Path:             "/docs",
	Title:            "Fiber API documentation",
	CacheAge:         60,
	Theme:            ThemeNone,
	RawSpecUrl:       "doc.json",
	ForceOffline:     ForceOfflineTrue,
	FallbackCacheAge: 86400,
}

func ptr[T any](v T) *T {
	return &v
}

var (
	ForceOfflineTrue  = ptr(true)
	ForceOfflineFalse = ptr(false)
)
