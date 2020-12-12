# GSKEMA 

gskema is a json schema validator

## Getting Started

Install package

```bash
go get github.com/ahsayde/gskema
```

## Usage

string values

```go
schema := gskema.String
schema.MinLength(3).MaxLength(5)

val, err := schema.Validate("go")           // invalid
val, err = schema.Validate("golang")        // valid
val, err = schema.Validate("go lang")       // invalid
```

## To Do
- [ ] Support multiple schemas (anyOf, oneOf, allOff).
- [ ] Add validator for string formats like email, ip, mac, etc..
- [ ] Add Enum support.
- [ ] Support required fields in struct.
- [ ] Support default values for struct fields.
- [ ] Add more examples.