// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"petstore/pkg/models/shared"
)

type GetOrderByIDRequest struct {
	// ID of order that needs to be fetched
	OrderID int64 `pathParam:"style=simple,explode=false,name=orderId"`
}

type GetOrderByIDResponse struct {
	Body        []byte
	ContentType string
	// successful operation
	Order       *shared.Order
	StatusCode  int
	RawResponse *http.Response
}
