# petstore

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/sagars-petstore
```
<!-- End SDK Installation -->

## SDK Example Usage
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
        Status: shared.PetStatusEnumPending.ToPointer(),
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Pet](docs/pet/README.md)

* [AddPetForm](docs/pet/README.md#addpetform) - Add a new pet to the store
* [AddPetJSON](docs/pet/README.md#addpetjson) - Add a new pet to the store
* [AddPetRaw](docs/pet/README.md#addpetraw) - Add a new pet to the store
* [DeletePet](docs/pet/README.md#deletepet) - Deletes a pet
* [FindPetsByStatus](docs/pet/README.md#findpetsbystatus) - Finds Pets by status
* [FindPetsByTags](docs/pet/README.md#findpetsbytags) - Finds Pets by tags
* [GetPetByID](docs/pet/README.md#getpetbyid) - Find pet by ID
* [UpdatePetWithForm](docs/pet/README.md#updatepetwithform) - Updates a pet in the store with form data
* [UpdatePetForm](docs/pet/README.md#updatepetform) - Update an existing pet
* [UpdatePetJSON](docs/pet/README.md#updatepetjson) - Update an existing pet
* [UpdatePetRaw](docs/pet/README.md#updatepetraw) - Update an existing pet
* [UploadFile](docs/pet/README.md#uploadfile) - uploads an image

### [Store](docs/store/README.md)

* [DeleteOrder](docs/store/README.md#deleteorder) - Delete purchase order by ID
* [GetInventory](docs/store/README.md#getinventory) - Returns pet inventories by status
* [GetOrderByID](docs/store/README.md#getorderbyid) - Find purchase order by ID
* [PlaceOrderForm](docs/store/README.md#placeorderform) - Place an order for a pet
* [PlaceOrderJSON](docs/store/README.md#placeorderjson) - Place an order for a pet
* [PlaceOrderRaw](docs/store/README.md#placeorderraw) - Place an order for a pet

### [User](docs/user/README.md)

* [CreateUserForm](docs/user/README.md#createuserform) - Create user
* [CreateUserJSON](docs/user/README.md#createuserjson) - Create user
* [CreateUserRaw](docs/user/README.md#createuserraw) - Create user
* [CreateUsersWithListInput](docs/user/README.md#createuserswithlistinput) - Creates list of users with given input array
* [DeleteUser](docs/user/README.md#deleteuser) - Delete user
* [GetUserByName](docs/user/README.md#getuserbyname) - Get user by user name
* [LoginUser](docs/user/README.md#loginuser) - Logs user into the system
* [LogoutUser](docs/user/README.md#logoutuser) - Logs out current logged in user session
* [UpdateUserForm](docs/user/README.md#updateuserform) - Update user
* [UpdateUserJSON](docs/user/README.md#updateuserjson) - Update user
* [UpdateUserRaw](docs/user/README.md#updateuserraw) - Update user
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
