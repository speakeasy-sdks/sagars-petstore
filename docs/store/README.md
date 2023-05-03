# Store

## Overview

Access to Petstore orders

Find out more about our store
<http://swagger.io>
### Available Operations

* [DeleteOrder](#deleteorder) - Delete purchase order by ID
* [GetInventory](#getinventory) - Returns pet inventories by status
* [GetOrderByID](#getorderbyid) - Find purchase order by ID
* [PlaceOrderForm](#placeorderform) - Place an order for a pet
* [PlaceOrderJSON](#placeorderjson) - Place an order for a pet
* [PlaceOrderRaw](#placeorderraw) - Place an order for a pet

## DeleteOrder

For valid response try integer IDs with value < 1000. Anything above 1000 or nonintegers will generate API errors

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.Store.DeleteOrder(ctx, operations.DeleteOrderRequest{
        OrderID: 662527,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## GetInventory

Returns a map of status codes to quantities

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.Store.GetInventory(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetInventory200ApplicationJSONObject != nil {
        // handle response
    }
}
```

## GetOrderByID

For valid response try integer IDs with value <= 5 or > 10. Other values will generate exceptions.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.Store.GetOrderByID(ctx, operations.GetOrderByIDRequest{
        OrderID: 820994,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Order != nil {
        // handle response
    }
}
```

## PlaceOrderForm

Place a new order in the store

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/shared"
	"petstore/pkg/types"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.Store.PlaceOrderForm(ctx, shared.Order{
        Complete: sdk.Bool(false),
        ID: sdk.Int64(10),
        PetID: sdk.Int64(198772),
        Quantity: sdk.Int(7),
        ShipDate: types.MustTimeFromString("2022-11-26T13:23:33.410Z"),
        Status: shared.OrderStatusEnumApproved.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Order != nil {
        // handle response
    }
}
```

## PlaceOrderJSON

Place a new order in the store

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/shared"
	"petstore/pkg/types"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.Store.PlaceOrderJSON(ctx, shared.Order{
        Complete: sdk.Bool(false),
        ID: sdk.Int64(10),
        PetID: sdk.Int64(198772),
        Quantity: sdk.Int(7),
        ShipDate: types.MustTimeFromString("2021-04-29T07:12:18.684Z"),
        Status: shared.OrderStatusEnumApproved.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Order != nil {
        // handle response
    }
}
```

## PlaceOrderRaw

Place a new order in the store

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/shared"
	"petstore/pkg/types"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.Store.PlaceOrderRaw(ctx, []byte("laborum"))
    if err != nil {
        log.Fatal(err)
    }

    if res.Order != nil {
        // handle response
    }
}
```
