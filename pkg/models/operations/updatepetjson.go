// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"petstore/pkg/models/shared"
)

type UpdatePetJSONSecurity struct {
	PetstoreAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

type UpdatePetJSONResponse struct {
	Body        []byte
	ContentType string
	// Successful operation
	Pet         *shared.Pet
	StatusCode  int
	RawResponse *http.Response
}
