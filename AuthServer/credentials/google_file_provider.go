package credentials

import (
	"github.com/meik99/CoffeeToGO/AuthServer/auth_json"
	"io"
	"log"
	"os"
)

type googleFileProvider struct {
	path                string
	credentialsSourceFn func(*googleFileProvider) io.Reader
}

func NewGoogleFileProvider(path string) Provider {
	return &googleFileProvider{
		path:                path,
		credentialsSourceFn: readerFromFile,
	}
}

func readerFromFile(gfp *googleFileProvider) io.Reader {
	credentialsJson, err := os.Open(gfp.path)
	if err != nil {
		log.Fatalf(err.Error(), err)
	}
	return credentialsJson
}

func (gfp *googleFileProvider) GetCredentials() (OAuthCredentials, error) {
	var gCredentials googleCredentials
	credentialsJson := gfp.credentialsSourceFn(gfp)

	err := auth_json.ParseJSONToInterface(credentialsJson, &gCredentials)
	if err != nil {
		return OAuthCredentials{}, err
	}
	return gCredentials.Web, nil
}
