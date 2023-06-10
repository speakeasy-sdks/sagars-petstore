# Order


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `Complete`                                         | **bool*                                            | :heavy_minus_sign:                                 | N/A                                                |                                                    |
| `ID`                                               | **int64*                                           | :heavy_minus_sign:                                 | N/A                                                | 10                                                 |
| `PetID`                                            | **int64*                                           | :heavy_minus_sign:                                 | N/A                                                | 198772                                             |
| `Quantity`                                         | **int*                                             | :heavy_minus_sign:                                 | N/A                                                | 7                                                  |
| `ShipDate`                                         | [*time.Time](https://pkg.go.dev/time#Time)         | :heavy_minus_sign:                                 | N/A                                                |                                                    |
| `Status`                                           | [*OrderStatus](../../models/shared/orderstatus.md) | :heavy_minus_sign:                                 | Order Status                                       | approved                                           |