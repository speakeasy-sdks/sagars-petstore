// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"petstore/pkg/models/shared"
)

type UploadFileSecurity struct {
	PetstoreAuth string `security:"scheme,type=oauth2,name=Authorization"`
}

type UploadFileRequest struct {
	RequestBody []byte `request:"mediaType=application/octet-stream"`
	// Additional Metadata
	AdditionalMetadata *string `queryParam:"style=form,explode=true,name=additionalMetadata"`
	// ID of pet to update
	PetID int64 `pathParam:"style=simple,explode=false,name=petId"`
}

type UploadFileResponse struct {
	// successful operation
	APIResponse *shared.APIResponse
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}
