// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"petstore/pkg/models/shared"
)

type GetPetByIDSecurity struct {
	APIKey       *string `security:"scheme,type=apiKey,subtype=header,name=api_key"`
	PetstoreAuth *string `security:"scheme,type=oauth2,name=Authorization"`
}

type GetPetByIDRequest struct {
	// ID of pet to return
	PetID int64 `pathParam:"style=simple,explode=false,name=petId"`
}

type GetPetByIDResponse struct {
	Body        []byte
	ContentType string
	// successful operation
	Pet         *shared.Pet
	StatusCode  int
	RawResponse *http.Response
}
