![actions](https://github.com/ahsayde/gskma/workflows/test/badge.svg)
[![codecov](https://codecov.io/gh/ahsayde/gskma/branch/main/graph/badge.svg?token=CIBAB97XQS)](https://codecov.io/gh/ahsayde/gskma)

# GSKEMA 

gskema is a json schema validator

## Getting Started

Install package

```bash
go get github.com/ahsayde/gskema
```

## Usage

```go
schema := gskema.String
schema.MinLength(3).MaxLength(5)

val, err := schema.Validate("go")           // invalid
val, err = schema.Validate("golang")        // valid
val, err = schema.Validate("go lang")       // invalid
```

```go
schema := gskema.TypeOf([]string{})
schema.MaxItems(3).MinItems(1)

val, err = schema.Validate([]string{"a"})
```

schema from struct

```go
type Person struct {
    Name        string      `json:"name,minlen=2,maxlen=10"`
    Age         string      `json:"age,min=1,max=150"`
    Address     []string    `json:"age,minitems=1,maxitems=2"`
}

schema := gskema.TypeOf(Person{})
val, err = schema.Validate(Test{Name: "ahmed"})
```



## To Do
- [ ] Write more docs
- [ ] Support multiple schemas (anyOf, oneOf, allOff).
- [ ] Add validator for string formats like email, ip, mac, etc..
- [ ] Add Enum support.
- [ ] Support required fields in struct.
- [ ] Support default values for struct fields.
- [ ] Add more examples.