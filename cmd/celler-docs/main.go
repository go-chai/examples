package main

import (
	"fmt"

	"github.com/go-chai/chai"
	"github.com/go-chai/chai/openapi2"
	"github.com/go-chai/examples/pkg/router"
	"github.com/go-openapi/spec"
)

func main() {
	r := router.GetRoutes()

	// This must be used only during development to generate the swagger spec
	docs, err := openapi2.Docs(r)
	if err != nil {
		panic(fmt.Sprintf("openapi2.Docs() failed: %+v", err))
	}

	addCustomDocs(docs)

	LogYAML(docs)

	// This must be used only during development to store the swagger spec
	err = openapi2.WriteDocs(docs, &openapi2.GenConfig{
		OutputDir: "cmd/celler/docs",
	})
	if err != nil {
		panic(fmt.Sprintf("gen.New().Generate() failed: %+v", err))
	}
}

func addCustomDocs(docs *spec.Swagger) {
	docs.Swagger = "2.0"
	docs.Host = "localhost:8080"
	docs.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Description:    "This is a sample celler server.",
			Title:          "Swagger Example API",
			TermsOfService: "http://swagger.io/terms/",
			Contact: &spec.ContactInfo{
				ContactInfoProps: spec.ContactInfoProps{
					Name:  "API Support",
					URL:   "http://www.swagger.io/support",
					Email: "support@swagger.io",
				},
			},
			License: &spec.License{
				LicenseProps: spec.LicenseProps{
					Name: "Apache 2.0",
					URL:  "http://www.apache.org/licenses/LICENSE-2.0.html",
				},
			},
			Version: "1.0",
		},
	}
	docs.SecurityDefinitions = map[string]*spec.SecurityScheme{
		"BasicAuth": {
			SecuritySchemeProps: spec.SecuritySchemeProps{
				Type: "basic",
			},
		},
		"ApiKeyAuth": {
			SecuritySchemeProps: spec.SecuritySchemeProps{
				Type: "apiKey",
				In:   "header",
				Name: "Authorization",
			},
		},
		"OAuth2Implicit": {
			SecuritySchemeProps: spec.SecuritySchemeProps{
				Description:      "Use with the OAuth2 Implicit Grant to retrieve a token",
				Type:             "oauth2",
				Flow:             "implicit",
				AuthorizationURL: "https://example.com/oauth/authorize",
				TokenURL:         "",
				Scopes: map[string]string{
					"admin": "Grants read and write access to administrative information",
					"write": "Grants write access",
				},
			},
		},
		"OAuth2Application": {
			SecuritySchemeProps: spec.SecuritySchemeProps{
				Description: "Use with the OAuth2 Implicit Grant to retrieve a token",
				Type:        "oauth2",
				Flow:        "application",
				TokenURL:    "https://example.com/oauth/token",
				Scopes: map[string]string{
					"admin": "Grants read and write access to administrative information",
					"write": "Grants write access",
				},
			},
		},

		"OAuth2Password": {
			SecuritySchemeProps: spec.SecuritySchemeProps{
				Type:     "oauth2",
				Flow:     "password",
				TokenURL: "https://example.com/oauth/token",
				Scopes: map[string]string{
					"admin": "Grants read and write access to administrative information",
					"write": "Grants write access",
					"read":  "Grants read access",
				},
			},
		},
		"OAuth2AccessToken": {
			SecuritySchemeProps: spec.SecuritySchemeProps{
				Type:             "oauth2",
				Flow:             "accessCode",
				AuthorizationURL: "https://example.com/oauth/authorize",
				TokenURL:         "https://example.com/oauth/token",
				Scopes: map[string]string{
					"admin": "Grants read and write access to administrative information",
				},
			},
		},
	}
}

func LogYAML(v interface{}) {
	bytes, err := chai.MarshalYAML(v)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bytes))

	return
}
