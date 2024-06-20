# Adding new API endpoints

Table of contents:
1. [Defining a new route](#defining-a-new-route)
2. [Defining a handler for your route](#defining-a-handler-for-your-route)

## Defining a new route

All routes are defined in `cmd/api/routes.go`, so simply add your new route there. For example:

```go
// ... other routes ...
router.Get("/my-endpoint", api.handleGetRecord)
// ... other routes ...
```

## Defining a handler for your route

To create a new set of handlers, add a new Go file to `cmd/api` with the filename `handle-<task>`. 

Your handlers should have the following signature

```go
func (api *Config) handleSomeTask(w http.ResponseWriter, r *http.Request) {
  // define your handler logic here...
}
```

### Receiving JSON payload

If your handler accepts a JSON payload, then define the structure of your payload as a type in `cmd/api/types.go`, and each field should have their respective JSON annotation. For example:
```go
type RequestPayload struct {
  Id string `json:"id"`
  ColumnOne string `json:"column_1"`
  ColumnTwo string `json:"column_2"`
}
```

Then, your handler can use the `readJSON` helper function defined at `cmd/api/json.go` to read the incoming JSON payload
```go
func (api *Config) handleSomeTask(w http.ResponseWriter, r *http.Request) {
  var payload RequestPayload

  if err := api.readJSON(w, r, &payload); err != nil {
    api.errorJSON(w, err, http.StatusBadRequest)
    return
  }

  // you can now use payload.Id, payload.ColumnOne, and payload.ColumnTwo
  // define the rest of your handler logic here...
}
```

### Returning JSON payload

If your handler function needs to return a JSON payload, you can do so by using the `writeJSON` helper function defined in `cmd/api/json.go`

```go
func (api *Config) handleSomeTask(w http.ResponseWriter, r *http.Request) {
  // your handler logic...
  // ...
  // your handler logic...
  api.writeJSON(w, <Status Code>, payload)
}
```

### Loading a URL parameter
If your route has a dynamic part, for example
```go
router.Get("/my-endpoint/{id}")
```

You can get the parameter like so
```go
func (api *Config) handleSomeTask(w http.ResponseWriter, r *http.Request) {
  id := chi.URLParam(r, "id")

  // define the rest of your handler logic here...
}
```

### Adding payload validators

Should you require to do error-handling for your JSON payload, please add them to `/cmd/api/validate.go` and use it in your handler.
