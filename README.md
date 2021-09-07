# bringo!

Simple unofficial [bring!](https://web.getbring.com/) API-wrapper written in Go

## Usage

```go
bring, err := bringo.NewWithLogin("my@email.com", "my password")
```

### Get All Lists (Meta)

Returns the name, UUID and theme of all lists. However, **this request does not include any items in the list.**
To also get the items in a list, use `GetList(listUUID string)` or `GetLists()`.

```go
var (
    meta []*bring.BringListMeta
    err error
)
meta, err = bring.GetListMetas()
for _, m := range meta {
    // ...
}
```

### Get Single List By ID

Returns the UUID, the status, items to buy and previous items.

```go
var (
    list *bring.BringList
    err error
)
list, err = bring.GetList("ffffffff-ffff-ffff-ffff-ffffffffffff")
// ...
```

### Get All Lists (Full)

combines `GetLists()` and `GetList(listUUID string)`

```go
var (
    lists []*bring.BringListExpensive
    err error
)
lists, err = bring.GetLists()
for _, l := range lists {
    // ...
}
```

### Create/Save a new item
SaveItem(listUUID, itemName, specification string
```go
var (
	err error
	listUUID = "ffffffff-ffff-ffff-ffff-ffffffffffff"
	itemName = "Butter"
	specification = ""
)
err = bring.SaveItem(listUUID, itemName, specification string)
```