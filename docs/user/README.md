# User

## Overview

Operations about user

### Available Operations

* [CreateUserForm](#createuserform) - Create user
* [CreateUserJSON](#createuserjson) - Create user
* [CreateUserRaw](#createuserraw) - Create user
* [CreateUsersWithListInput](#createuserswithlistinput) - Creates list of users with given input array
* [DeleteUser](#deleteuser) - Delete user
* [GetUserByName](#getuserbyname) - Get user by user name
* [LoginUser](#loginuser) - Logs user into the system
* [LogoutUser](#logoutuser) - Logs out current logged in user session
* [UpdateUserForm](#updateuserform) - Update user
* [UpdateUserJSON](#updateuserjson) - Update user
* [UpdateUserRaw](#updateuserraw) - Update user

## CreateUserForm

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/shared"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.User.CreateUserForm(ctx, shared.User{
        Email: sdk.String("john@email.com"),
        FirstName: sdk.String("John"),
        ID: sdk.Int64(10),
        LastName: sdk.String("James"),
        Password: sdk.String("12345"),
        Phone: sdk.String("12345"),
        UserStatus: sdk.Int(1),
        Username: sdk.String("theUser"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## CreateUserJSON

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/shared"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.User.CreateUserJSON(ctx, shared.User{
        Email: sdk.String("john@email.com"),
        FirstName: sdk.String("John"),
        ID: sdk.Int64(10),
        LastName: sdk.String("James"),
        Password: sdk.String("12345"),
        Phone: sdk.String("12345"),
        UserStatus: sdk.Int(1),
        Username: sdk.String("theUser"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## CreateUserRaw

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/shared"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.User.CreateUserRaw(ctx, []byte("quasi"))
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## CreateUsersWithListInput

Creates list of users with given input array

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/shared"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.User.CreateUsersWithListInput(ctx, []shared.User{
        shared.User{
            Email: sdk.String("john@email.com"),
            FirstName: sdk.String("John"),
            ID: sdk.Int64(10),
            LastName: sdk.String("James"),
            Password: sdk.String("12345"),
            Phone: sdk.String("12345"),
            UserStatus: sdk.Int(1),
            Username: sdk.String("theUser"),
        },
        shared.User{
            Email: sdk.String("john@email.com"),
            FirstName: sdk.String("John"),
            ID: sdk.Int64(10),
            LastName: sdk.String("James"),
            Password: sdk.String("12345"),
            Phone: sdk.String("12345"),
            UserStatus: sdk.Int(1),
            Username: sdk.String("theUser"),
        },
        shared.User{
            Email: sdk.String("john@email.com"),
            FirstName: sdk.String("John"),
            ID: sdk.Int64(10),
            LastName: sdk.String("James"),
            Password: sdk.String("12345"),
            Phone: sdk.String("12345"),
            UserStatus: sdk.Int(1),
            Username: sdk.String("theUser"),
        },
        shared.User{
            Email: sdk.String("john@email.com"),
            FirstName: sdk.String("John"),
            ID: sdk.Int64(10),
            LastName: sdk.String("James"),
            Password: sdk.String("12345"),
            Phone: sdk.String("12345"),
            UserStatus: sdk.Int(1),
            Username: sdk.String("theUser"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

## DeleteUser

This can only be done by the logged in user.

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
    res, err := s.User.DeleteUser(ctx, operations.DeleteUserRequest{
        Username: "Weston_Thiel",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## GetUserByName

Get user by user name

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
    res, err := s.User.GetUserByName(ctx, operations.GetUserByNameRequest{
        Username: "Whitney.Bednar",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.User != nil {
        // handle response
    }
}
```

## LoginUser

Logs user into the system

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
    res, err := s.User.LoginUser(ctx, operations.LoginUserRequest{
        Password: sdk.String("cum"),
        Username: sdk.String("Aiyana.Batz"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.LoginUser200ApplicationJSONString != nil {
        // handle response
    }
}
```

## LogoutUser

Logs out current logged in user session

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
    res, err := s.User.LogoutUser(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## UpdateUserForm

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/operations"
	"petstore/pkg/models/shared"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.User.UpdateUserForm(ctx, operations.UpdateUserFormRequest{
        User: &shared.User{
            Email: sdk.String("john@email.com"),
            FirstName: sdk.String("John"),
            ID: sdk.Int64(10),
            LastName: sdk.String("James"),
            Password: sdk.String("12345"),
            Phone: sdk.String("12345"),
            UserStatus: sdk.Int(1),
            Username: sdk.String("theUser"),
        },
        Username: "Wilfrid.Carter",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## UpdateUserJSON

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/operations"
	"petstore/pkg/models/shared"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.User.UpdateUserJSON(ctx, operations.UpdateUserJSONRequest{
        User: &shared.User{
            Email: sdk.String("john@email.com"),
            FirstName: sdk.String("John"),
            ID: sdk.Int64(10),
            LastName: sdk.String("James"),
            Password: sdk.String("12345"),
            Phone: sdk.String("12345"),
            UserStatus: sdk.Int(1),
            Username: sdk.String("theUser"),
        },
        Username: "Jayden.Carter88",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## UpdateUserRaw

This can only be done by the logged in user.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/operations"
	"petstore/pkg/models/shared"
)

func main() {
    s := sdk.New()

    ctx := context.Background()
    res, err := s.User.UpdateUserRaw(ctx, operations.UpdateUserRawRequest{
        RequestBody: []byte("commodi"),
        Username: "Terrill69",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```
