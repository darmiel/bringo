# bringo!
Simple unofficial [bring!](https://web.getbring.com/) API-wrapper written in Go

## Usage

```go
bring, err := bringo.NewWithLogin("my@email.com", "my password")
```

### Get All Lists (Meta)

**Type:**

```go
UUID  string `json:"listUuid"`
Name  string `json:"name"`
Theme string `json:"theme"`
```

**Request:**

```go
meta, err := bring.GetListMetas()
for _, m := range meta {
// ...
}
```

### Get All Lists (Full)
(combines `GetLists()` and `GetList(listUUID string)`)
**Type:**
```go
UUID     string       `json:"listUuid"`
Name     string       `json:"name"`
Theme    string       `json:"theme"`
UUID     string       `json:"uuid"`
Status   string       `json:"status"`
Purchase []*BringItem `json:"purchase"`
Recently []*BringItem `json:"recently"
```
**Request:**
```go
lists, err := bring.GetLists()
for _, l := range lists {
	// ...
}
```

### Get Single List By ID
**Type:**
```go
UUID     string       `json:"uuid"`
Status   string       `json:"status"`
Purchase []*BringItem `json:"purchase"`
Recently []*BringItem `json:"recently"`
```
**Request:**
```go
list, err := bring.GetList("ffffffff-ffff-ffff-ffff-ffffffffffff")
// ...
```