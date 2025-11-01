package scalar

import (
	_ "embed"
	"fmt"
	"path"
	"strings"
	"text/template"

	"github.com/gofiber/fiber/v3"
	"github.com/swaggo/swag/v2"
)

//go:embed scalar.min.js
var embeddedJS []byte

func New(config ...Config) fiber.Handler {
	// Set default config
	cfg := configDefault

	// Override config if provided
	if len(config) > 0 {
		cfg = config[0]

		// Set default values
		if len(cfg.BasePath) == 0 {
			cfg.BasePath = configDefault.BasePath
		}
		if len(cfg.Path) == 0 {
			cfg.Path = configDefault.Path
		}
		if len(cfg.Title) == 0 {
			cfg.Title = configDefault.Title
		}
		if len(cfg.RawSpecUrl) == 0 {
			cfg.RawSpecUrl = configDefault.RawSpecUrl
		}
		if cfg.ForceOffline == nil {
			cfg.ForceOffline = configDefault.ForceOffline
		}
		if cfg.FallbackCacheAge == 0 {
			cfg.FallbackCacheAge = configDefault.FallbackCacheAge
		}
		if cfg.Theme == "" {
			cfg.Theme = ThemeNone
		}
	}

	rawSpec := cfg.FileContentString
	if len(rawSpec) == 0 {
		doc, err := swag.ReadDoc()
		if err != nil {
			panic(err)
		}
		rawSpec = doc
	}

	cfg.FileContentString = string(rawSpec)

	html, err := template.New("index.html").Parse(templateHTML)
	if err != nil {
		panic(fmt.Errorf("failed to parse html template:%v", err))
	}

	var forceOfflineResolved bool
	if cfg.ForceOffline != nil {
		forceOfflineResolved = *cfg.ForceOffline
	} else if configDefault.ForceOffline != nil {
		forceOfflineResolved = *configDefault.ForceOffline
	} else {
		forceOfflineResolved = false
	}

	htmlData := struct {
		Config
		ForceOffline bool
		Extra        map[string]any
	}{
		Config:       cfg,
		ForceOffline: forceOfflineResolved,
		Extra:        map[string]any{},
	}

	return func(ctx fiber.Ctx) error {
		if cfg.Next != nil && cfg.Next(ctx) {
			return ctx.Next()
		}

		resolvedBasePath := cfg.BasePath
		if xf := ctx.Get("X-Forwarded-Prefix"); xf != "" {
			resolvedBasePath = xf
		} else if xf2 := ctx.Get("X-Forwarded-Path"); xf2 != "" {
			resolvedBasePath = xf2
		}
		scalarUIPath := cfg.Path
		specURL := path.Join(scalarUIPath, cfg.RawSpecUrl)
		jsFallbackPath := path.Join(resolvedBasePath, scalarUIPath, "/js/api-reference.min.js")

		// fallback js
		if ctx.Path() == jsFallbackPath {
			ctx.Set("Cache-Control", fmt.Sprintf("public, max-age=%d", cfg.FallbackCacheAge))
			return ctx.Send(embeddedJS)
		}

		if cfg.CacheAge > 0 {
			ctx.Set("Cache-Control", fmt.Sprintf("public, max-age=%d", cfg.CacheAge))
		} else {
			ctx.Set("Cache-Control", "no-store")
		}

		if ctx.Path() == specURL {
			ctx.Set("Content-Type", "application/json")
			return ctx.SendString(rawSpec)
		}

		if !strings.HasPrefix(ctx.Path(), scalarUIPath) && ctx.Path() != specURL && ctx.Path() != jsFallbackPath {
			return ctx.Next()
		}

		htmlData.Extra["FallbackUrl"] = jsFallbackPath
		ctx.Type("html")
		return html.Execute(ctx, htmlData)
	}
}
