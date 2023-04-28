# Pet

## Overview

Everything about your Pets

Find out more
<http://swagger.io>
### Available Operations

* [AddPetForm](#addpetform) - Add a new pet to the store
* [AddPetJSON](#addpetjson) - Add a new pet to the store
* [AddPetRaw](#addpetraw) - Add a new pet to the store
* [DeletePet](#deletepet) - Deletes a pet
* [FindPetsByStatus](#findpetsbystatus) - Finds Pets by status
* [FindPetsByTags](#findpetsbytags) - Finds Pets by tags
* [GetPetByID](#getpetbyid) - Find pet by ID
* [UpdatePetWithForm](#updatepetwithform) - Updates a pet in the store with form data
* [UpdatePetForm](#updatepetform) - Update an existing pet
* [UpdatePetJSON](#updatepetjson) - Update an existing pet
* [UpdatePetRaw](#updatepetraw) - Update an existing pet
* [UploadFile](#uploadfile) - uploads an image

## AddPetForm

Add a new pet to the store

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/shared"
	"petstore/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()    
    req := shared.Pet{
        Category: &shared.Category{
            ID: sdk.Int64(1),
            Name: sdk.String("Dogs"),
        },
        ID: sdk.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "ipsam",
        },
        Status: shared.PetStatusEnumSold.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: sdk.Int64(778157),
                Name: sdk.String("Teri Strosin"),
            },
            shared.Tag{
                ID: sdk.Int64(799159),
                Name: sdk.String("Erik Lebsack"),
            },
            shared.Tag{
                ID: sdk.Int64(118274),
                Name: sdk.String("Luke McCullough"),
            },
            shared.Tag{
                ID: sdk.Int64(944669),
                Name: sdk.String("Everett Breitenberg"),
            },
        },
    }

    res, err := s.Pet.AddPetForm(ctx, req, operations.AddPetFormSecurity{
        PetstoreAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

## AddPetJSON

Add a new pet to the store

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/shared"
	"petstore/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()    
    req := shared.Pet{
        Category: &shared.Category{
            ID: sdk.Int64(1),
            Name: sdk.String("Dogs"),
        },
        ID: sdk.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "qui",
            "impedit",
        },
        Status: shared.PetStatusEnumSold.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: sdk.Int64(216550),
                Name: sdk.String("Brandon Auer"),
            },
            shared.Tag{
                ID: sdk.Int64(149675),
                Name: sdk.String("Curtis Morissette"),
            },
        },
    }

    res, err := s.Pet.AddPetJSON(ctx, req, operations.AddPetJSONSecurity{
        PetstoreAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

## AddPetRaw

Add a new pet to the store

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/shared"
	"petstore/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()    
    req := []byte("saepe")

    res, err := s.Pet.AddPetRaw(ctx, req, operations.AddPetRawSecurity{
        PetstoreAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

## DeletePet

Deletes a pet

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
    req := operations.DeletePetRequest{
        APIKey: sdk.String("fuga"),
        PetID: 449950,
    }

    res, err := s.Pet.DeletePet(ctx, req, operations.DeletePetSecurity{
        PetstoreAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## FindPetsByStatus

Multiple status values can be provided with comma separated strings

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
    req := operations.FindPetsByStatusRequest{
        Status: operations.FindPetsByStatusStatusEnumPending.ToPointer(),
    }

    res, err := s.Pet.FindPetsByStatus(ctx, req, operations.FindPetsByStatusSecurity{
        PetstoreAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pets != nil {
        // handle response
    }
}
```

## FindPetsByTags

Multiple tags can be provided with comma separated strings. Use tag1, tag2, tag3 for testing.

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
    req := operations.FindPetsByTagsRequest{
        Tags: []string{
            "iure",
            "saepe",
            "quidem",
        },
    }

    res, err := s.Pet.FindPetsByTags(ctx, req, operations.FindPetsByTagsSecurity{
        PetstoreAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pets != nil {
        // handle response
    }
}
```

## GetPetByID

Returns a single pet

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
    req := operations.GetPetByIDRequest{
        PetID: 99280,
    }

    res, err := s.Pet.GetPetByID(ctx, req, operations.GetPetByIDSecurity{
        APIKey: sdk.String("YOUR_API_KEY_HERE"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

## UpdatePetWithForm

Updates a pet in the store with form data

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
    req := operations.UpdatePetWithFormRequest{
        Name: sdk.String("Lela Orn"),
        PetID: 170909,
        Status: sdk.String("dolorem"),
    }

    res, err := s.Pet.UpdatePetWithForm(ctx, req, operations.UpdatePetWithFormSecurity{
        PetstoreAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## UpdatePetForm

Update an existing pet by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/shared"
	"petstore/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()    
    req := shared.Pet{
        Category: &shared.Category{
            ID: sdk.Int64(1),
            Name: sdk.String("Dogs"),
        },
        ID: sdk.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "explicabo",
            "nobis",
        },
        Status: shared.PetStatusEnumAvailable.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: sdk.Int64(363711),
                Name: sdk.String("Velma Batz"),
            },
            shared.Tag{
                ID: sdk.Int64(988374),
                Name: sdk.String("Juan O'Hara"),
            },
            shared.Tag{
                ID: sdk.Int64(161309),
                Name: sdk.String("Shaun McCullough"),
            },
        },
    }

    res, err := s.Pet.UpdatePetForm(ctx, req, operations.UpdatePetFormSecurity{
        PetstoreAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

## UpdatePetJSON

Update an existing pet by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/shared"
	"petstore/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()    
    req := shared.Pet{
        Category: &shared.Category{
            ID: sdk.Int64(1),
            Name: sdk.String("Dogs"),
        },
        ID: sdk.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "molestiae",
            "velit",
        },
        Status: shared.PetStatusEnumPending.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: sdk.Int64(338007),
                Name: sdk.String("Kayla O'Kon"),
            },
        },
    }

    res, err := s.Pet.UpdatePetJSON(ctx, req, operations.UpdatePetJSONSecurity{
        PetstoreAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

## UpdatePetRaw

Update an existing pet by Id

### Example Usage

```go
package main

import(
	"context"
	"log"
	"petstore"
	"petstore/pkg/models/shared"
	"petstore/pkg/models/operations"
)

func main() {
    s := sdk.New()

    ctx := context.Background()    
    req := []byte("quo")

    res, err := s.Pet.UpdatePetRaw(ctx, req, operations.UpdatePetRawSecurity{
        PetstoreAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Pet != nil {
        // handle response
    }
}
```

## UploadFile

uploads an image

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
    req := operations.UploadFileRequest{
        RequestBody: []byte("sequi"),
        AdditionalMetadata: sdk.String("tenetur"),
        PetID: 368725,
    }

    res, err := s.Pet.UploadFile(ctx, req, operations.UploadFileSecurity{
        PetstoreAuth: "Bearer YOUR_ACCESS_TOKEN_HERE",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.APIResponse != nil {
        // handle response
    }
}
```
