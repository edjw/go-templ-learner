## Go with Templ for templating and Tailwind and HTMX and deploy on Render.com

Absolutely obscene build command for Render.com to build Go with Templ, Tailwind and HTMX.

```bash
export GOPATH=/opt/render/project/go && export PATH=$PATH:$GOPATH/bin && go install github.com/a-h/templ/cmd/templ@v0.2.408 && $GOPATH/bin/templ generate && npm install && npm run build_tailwind && go build -tags netgo -ldflags '-s -w' -o app
```

1. Install Templ after some funky path stuff.
2. Generate the Go files from the Templ templates
3. Install node modules (needed for the Tailwind typography module)
4. Run Tailwind to generate the CSS
5. Build the Go app using Render's build command