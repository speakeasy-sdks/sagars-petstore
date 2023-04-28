// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"petstore/pkg/models/shared"
)

type GetUserByNameRequest struct {
	// The name that needs to be fetched. Use user1 for testing.
	Username string `pathParam:"style=simple,explode=false,name=username"`
}

type GetUserByNameResponse struct {
	Body        []byte
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// successful operation
	User *shared.User
}
