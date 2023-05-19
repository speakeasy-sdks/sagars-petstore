<!-- Start SDK Example Usage -->
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
    res, err := s.Pet.AddPetForm(ctx, shared.Pet{
        Category: &shared.Category{
            ID: sdk.Int64(1),
            Name: sdk.String("Dogs"),
        },
        ID: sdk.Int64(10),
        Name: "doggie",
        PhotoUrls: []string{
            "provident",
            "distinctio",
            "quibusdam",
        },
        Status: shared.PetStatusPending.ToPointer(),
        Tags: []shared.Tag{
            shared.Tag{
                ID: sdk.Int64(544883),
                Name: sdk.String("Ben Mueller"),
            },
            shared.Tag{
                ID: sdk.Int64(437587),
                Name: sdk.String("Raquel Bednar"),
            },
            shared.Tag{
                ID: sdk.Int64(383441),
                Name: sdk.String("Alexandra Schulist"),
            },
            shared.Tag{
                ID: sdk.Int64(568045),
                Name: sdk.String("Mrs. Sophie Smith MD"),
            },
        },
    }, operations.AddPetFormSecurity{
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
<!-- End SDK Example Usage -->