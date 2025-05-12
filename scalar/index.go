package scalar

const templateHTML = `
<!doctype html>
<html>
  <head>
    <title>{{.Title}}</title>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width, initial-scale=1" />

    {{- if .CustomStyle }}
    <style>
      :root {
        {{ .CustomStyle }}
      }
    </style>
    {{ end }}
  </head>

  <body>
    <div id="app"></div>
    <script>
      function initScalar() {
        if (typeof Scalar !== 'undefined') {
          Scalar.createApiReference('#app', {
            content: {{.FileContentString}},
            {{- if .ProxyUrl}}
            proxyUrl: "{{.ProxyUrl}}",
            {{ end }}
          });
        }else{
	      console.error("Failed to load Scalar API Reference");
	      document.querySelector('#app').innerHTML = "<p>Something went wrong. Please report this bug at <a href='https://github.com/yokeTH/gofiber-scalar/issues'>https://github.com/yokeTH/gofiber-scalar/issues</a></p>";
        }
      }
    </script>
    <script src="{{ .Extra.FallbackUrl }}" onload="initScalar()" ></script>
  </body>
</html>`
