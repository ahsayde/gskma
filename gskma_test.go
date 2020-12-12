package gskma

import (
	"reflect"
	"testing"
)

type Test struct {
	A string `json:"a,required,maxlen=10,minlen=1"`
	B int32  `json:"b,max=3,min=1"`
	C int64
	D float32
	E float64
	F bool
	G map[string]int
	H []string
}

type OptionTestCase struct {
	schema Schema
	value  interface{}
	err    bool
}

func TestBasicTypes(t *testing.T) {

	type testCase struct {
		schema   Schema
		expected interface{}
		value    interface{}
		err      bool
	}

	cases := []testCase{
		{
			schema:   Int32(),
			value:    int(1),
			expected: int32(1),
			err:      false,
		},
		{
			schema:   Int32(),
			value:    int8(1),
			expected: int32(1),
			err:      false,
		},
		{
			schema:   Int32(),
			value:    int16(1),
			expected: int32(1),
			err:      false,
		},
		{
			schema:   Int32(),
			value:    int32(1),
			expected: int32(1),
			err:      false,
		},
		{
			schema:   Int32(),
			value:    int64(1),
			expected: int32(1),
			err:      false,
		},
		{
			schema:   Int32(),
			value:    "1",
			expected: int32(1),
			err:      false,
		},
		{
			schema:   Int32(),
			value:    "1.5",
			expected: nil,
			err:      true,
		},
		{
			schema:   Int32(),
			value:    "string",
			expected: nil,
			err:      true,
		},
		{
			schema:   Int32(),
			value:    []string{},
			expected: nil,
			err:      true,
		},
		{
			schema:   Int64(),
			value:    int(1),
			expected: int64(1),
			err:      false,
		},
		{
			schema:   Int64(),
			value:    int8(1),
			expected: int64(1),
			err:      false,
		},
		{
			schema:   Int64(),
			value:    int16(1),
			expected: int64(1),
			err:      false,
		},
		{
			schema:   Int64(),
			value:    int32(1),
			expected: int64(1),
			err:      false,
		},
		{
			schema:   Int64(),
			value:    int64(1),
			expected: int64(1),
			err:      false,
		},
		{
			schema:   Int64(),
			value:    "1",
			expected: int64(1),
			err:      false,
		},
		{
			schema:   Int64(),
			value:    "1.5",
			expected: nil,
			err:      true,
		},
		{
			schema:   Int64(),
			value:    "string",
			expected: nil,
			err:      true,
		},
		{
			schema:   Int64(),
			value:    []string{},
			expected: nil,
			err:      true,
		},
		{
			schema:   Float32(),
			value:    float32(1.5),
			expected: float32(1.5),
			err:      false,
		},
		{
			schema:   Float32(),
			value:    float64(1.5),
			expected: float32(1.5),
			err:      false,
		},
		{
			schema:   Float32(),
			value:    "1.5",
			expected: float32(1.5),
			err:      false,
		},
		{
			schema:   Float32(),
			value:    "1",
			expected: float32(1),
			err:      false,
		},
		{
			schema:   Float32(),
			value:    "string",
			expected: nil,
			err:      true,
		},
		{
			schema:   Float32(),
			value:    []string{},
			expected: nil,
			err:      true,
		},
		{
			schema:   Float64(),
			value:    float64(1.5),
			expected: float64(1.5),
			err:      false,
		},
		{
			schema:   Float64(),
			value:    float64(1.5),
			expected: float64(1.5),
			err:      false,
		},
		{
			schema:   Float64(),
			value:    "1.5",
			expected: float64(1.5),
			err:      false,
		},
		{
			schema:   Float64(),
			value:    "1",
			expected: float64(1),
			err:      false,
		},
		{
			schema:   Float64(),
			value:    "string",
			expected: nil,
			err:      true,
		},
		{
			schema:   Float64(),
			value:    []string{},
			expected: nil,
			err:      true,
		},
		{
			schema:   Boolean(),
			value:    true,
			expected: true,
			err:      false,
		},
		{
			schema:   Boolean(),
			value:    false,
			expected: false,
			err:      false,
		},
		{
			schema:   Boolean(),
			value:    "true",
			expected: true,
			err:      false,
		},
		{
			schema:   Boolean(),
			value:    "false",
			expected: false,
			err:      false,
		},
		{
			schema:   TypeOf(map[string]string{}),
			value:    map[string]string{"k": "v"},
			expected: map[string]string{"k": "v"},
			err:      false,
		},
		{
			schema:   TypeOf(map[string]string{}),
			value:    map[string]int{"k": 1},
			expected: nil,
			err:      true,
		},
		{
			schema:   TypeOf([]string{}),
			value:    []string{"a", "b"},
			expected: []string{"a", "b"},
			err:      false,
		},
		{
			schema:   TypeOf([]string{}),
			value:    []int{1, 2},
			expected: nil,
			err:      true,
		},
		{
			schema:   TypeOf(Test{}),
			value:    Test{A: "string", B: 1, C: 1, D: 1.5, E: 1.5},
			expected: Test{A: "string", B: 1, C: 1, D: 1.5, E: 1.5},
			err:      false,
		},
		{
			schema:   TypeOf(Test{}),
			value:    map[string]string{},
			expected: nil,
			err:      true,
		},
	}

	for i, c := range cases {
		s := c.schema
		v, err := s.Validate(c.value)

		if c.err != (err != nil) {
			t.Errorf("Test Case #%d: Unexpected error: %v %v", i, err, c.schema)
		}
		if !reflect.DeepEqual(v, c.expected) {
			t.Errorf("aaaaaaaaaaa")
		}
	}
}

func TestDefault(t *testing.T) {
	s := String()
	s.Default("test")

	v, _ := s.Validate(nil)
	if v != "test" {
		t.Errorf("default is not working for string")
	}

	s = Int32()
	s.Default(1)

	v, _ = s.Validate(nil)
	if v.(int32) != 1 {
		t.Errorf("default is not working for int32")
	}

	s = Int64()
	s.Default(1)

	v, _ = s.Validate(nil)
	if v.(int64) != 1 {
		t.Errorf("default is not working for int64")
	}

	s = Float32()
	s.Default(1.5)

	v, _ = s.Validate(nil)
	if v.(float32) != 1.5 {
		t.Errorf("default is not working for float32")
	}

	s = Float64()
	s.Default(1.5)

	v, _ = s.Validate(nil)
	if v.(float64) != 1.5 {
		t.Errorf("default is not working for float64")
	}

	s = Boolean()
	s.Default(true)

	v, _ = s.Validate(nil)
	if v.(bool) != true {
		t.Errorf("default is not working for float64")
	}

	var d interface{}

	s = TypeOf(map[string]string{})
	d = map[string]string{"a": "a"}
	s.Default(d)

	v, _ = s.Validate(nil)
	if !reflect.DeepEqual(v, d) {
		t.Errorf("default is not working for maps")
	}

	s = TypeOf([]string{})
	d = []string{"a"}
	s.Default(d)

	v, _ = s.Validate(nil)
	if !reflect.DeepEqual(v, d) {
		t.Errorf("default is not working for arrays")
	}

	s = TypeOf(Test{})
	d = Test{A: "string", B: 1, C: 1, D: 1.5, E: 1.5}
	s.Default(d)

	v, _ = s.Validate(nil)
	if !reflect.DeepEqual(v, d) {
		t.Errorf("default is not working for structs")
	}

	s = TypeOf([]Test{})
	d = []Test{{A: "string", B: 1, C: 1, D: 1.5, E: 1.5}}
	s.Default(d)

	v, _ = s.Validate(nil)
	if !reflect.DeepEqual(v, d) {
		t.Errorf("default is not working for array of struct")
	}
}

func TestMinimum(t *testing.T) {
	cases := []OptionTestCase{
		{
			schema: Int32(),
			value:  2,
			err:    false,
		},
		{
			schema: Int32(),
			value:  1,
			err:    false,
		},
		{
			schema: Int32(),
			value:  0,
			err:    true,
		},
		{
			schema: Int64(),
			value:  2,
			err:    false,
		},
		{
			schema: Int64(),
			value:  1,
			err:    false,
		},
		{
			schema: Int64(),
			value:  0,
			err:    true,
		},
		{
			schema: Float32(),
			value:  2.0,
			err:    false,
		},
		{
			schema: Float32(),
			value:  1.0,
			err:    false,
		},
		{
			schema: Float32(),
			value:  .95,
			err:    true,
		},
		{
			schema: Float64(),
			value:  2.0,
			err:    false,
		},
		{
			schema: Float64(),
			value:  1.0,
			err:    false,
		},
		{
			schema: Float64(),
			value:  0.95,
			err:    true,
		},
	}

	for _, c := range cases {
		c.schema.Minimum(1)
		_, err := c.schema.Validate(c.value)

		if c.err != (err != nil) {
			t.Error("Minimum is not working for", c.schema.data.rkind)
		}
	}
}

func TestMaximum(t *testing.T) {
	cases := []OptionTestCase{
		{
			schema: Int32(),
			value:  0,
			err:    false,
		},
		{
			schema: Int32(),
			value:  1,
			err:    false,
		},
		{
			schema: Int32(),
			value:  2,
			err:    true,
		},
		{
			schema: Int64(),
			value:  0,
			err:    false,
		},
		{
			schema: Int64(),
			value:  1,
			err:    false,
		},
		{
			schema: Int64(),
			value:  2,
			err:    true,
		},
		{
			schema: Float32(),
			value:  0.0,
			err:    false,
		},
		{
			schema: Float32(),
			value:  1.0,
			err:    false,
		},
		{
			schema: Float32(),
			value:  1.1,
			err:    true,
		},
		{
			schema: Float64(),
			value:  0.0,
			err:    false,
		},
		{
			schema: Float64(),
			value:  1.0,
			err:    false,
		},
		{
			schema: Float64(),
			value:  1.1,
			err:    true,
		},
	}

	for _, c := range cases {
		c.schema.Maximum(1)
		_, err := c.schema.Validate(c.value)
		if c.err != (err != nil) {
			t.Error("Maximum is not working for", c.schema.data.rkind)
		}
	}
}

func TestExclusiveMaximum(t *testing.T) {
	cases := []OptionTestCase{
		{
			schema: Int32(),
			value:  0,
			err:    false,
		},
		{
			schema: Int32(),
			value:  1,
			err:    true,
		},
		{
			schema: Int32(),
			value:  2,
			err:    true,
		},
		{
			schema: Int64(),
			value:  0,
			err:    false,
		},
		{
			schema: Int64(),
			value:  1,
			err:    true,
		},
		{
			schema: Int64(),
			value:  2,
			err:    true,
		},
		{
			schema: Float32(),
			value:  0.0,
			err:    false,
		},
		{
			schema: Float32(),
			value:  1.0,
			err:    true,
		},
		{
			schema: Float32(),
			value:  1.1,
			err:    true,
		},
		{
			schema: Float64(),
			value:  0.0,
			err:    false,
		},
		{
			schema: Float64(),
			value:  1.0,
			err:    true,
		},
		{
			schema: Float64(),
			value:  1.1,
			err:    true,
		},
	}

	for _, c := range cases {
		c.schema.ExclusiveMaximum(1)
		_, err := c.schema.Validate(c.value)
		if c.err != (err != nil) {
			t.Error("ExclusiveMaximum is not working for", c.schema.data.rkind)
		}
	}
}

func TestExclusiveMinimum(t *testing.T) {
	cases := []OptionTestCase{
		{
			schema: Int32(),
			value:  2,
			err:    false,
		},
		{
			schema: Int32(),
			value:  1,
			err:    true,
		},
		{
			schema: Int32(),
			value:  0,
			err:    true,
		},
		{
			schema: Int64(),
			value:  2,
			err:    false,
		},
		{
			schema: Int64(),
			value:  1,
			err:    true,
		},
		{
			schema: Int64(),
			value:  0,
			err:    true,
		},
		{
			schema: Float32(),
			value:  2.0,
			err:    false,
		},
		{
			schema: Float32(),
			value:  1.0,
			err:    true,
		},
		{
			schema: Float32(),
			value:  .95,
			err:    true,
		},
		{
			schema: Float64(),
			value:  2.0,
			err:    false,
		},
		{
			schema: Float64(),
			value:  1.0,
			err:    true,
		},
		{
			schema: Float64(),
			value:  0.95,
			err:    true,
		},
	}
	for _, c := range cases {
		c.schema.ExclusiveMinimum(1)
		_, err := c.schema.Validate(c.value)
		if c.err != (err != nil) {
			t.Error("ExclusiveMinimum is not working for", c.schema.data.rkind)
		}
	}
}

func TestMultipleOf(t *testing.T) {
	cases := []OptionTestCase{
		{
			schema: Int32(),
			value:  2,
			err:    false,
		},
		{
			schema: Int32(),
			value:  3,
			err:    true,
		},
		{
			schema: Int32(),
			value:  4,
			err:    false,
		},
		{
			schema: Int64(),
			value:  2,
			err:    false,
		},
		{
			schema: Int64(),
			value:  3,
			err:    true,
		},
		{
			schema: Int64(),
			value:  4,
			err:    false,
		},
	}
	for _, c := range cases {
		c.schema.MultipleOf(2)
		_, err := c.schema.Validate(c.value)
		if c.err != (err != nil) {
			t.Error("TestMultipleOf is not working for type", c.schema.data.rkind)
		}
	}
}

func TestMaxLength(t *testing.T) {
	cases := []OptionTestCase{
		{
			schema: String(),
			value:  "",
			err:    false,
		},
		{
			schema: String(),
			value:  "s",
			err:    false,
		},
		{
			schema: String(),
			value:  "ss",
			err:    true,
		},
	}
	for _, c := range cases {
		c.schema.MaxLength(1)
		_, err := c.schema.Validate(c.value)
		if c.err != (err != nil) {
			t.Error("MaxLength is not working for", c.schema.data.rkind)
		}
	}
}

func TestMinLength(t *testing.T) {
	cases := []OptionTestCase{
		{
			schema: String(),
			value:  "",
			err:    true,
		},
		{
			schema: String(),
			value:  "s",
			err:    false,
		},
		{
			schema: String(),
			value:  "ss",
			err:    false,
		},
	}
	for _, c := range cases {
		c.schema.MinLength(1)
		_, err := c.schema.Validate(c.value)
		if c.err != (err != nil) {
			t.Error("MinLength is not working for", c.schema.data.rkind)
		}
	}
}

func TestMaxItems(t *testing.T) {
	cases := []OptionTestCase{
		{
			schema: TypeOf([]string{}),
			value:  []string{},
			err:    false,
		},
		{
			schema: TypeOf([]string{}),
			value:  []string{"a"},
			err:    false,
		},
		{
			schema: TypeOf([]string{}),
			value:  []string{"a", "b"},
			err:    true,
		},
	}
	for _, c := range cases {
		c.schema.MaxItems(1)
		_, err := c.schema.Validate(c.value)
		if c.err != (err != nil) {
			t.Error("MaxItems is not working")
		}
	}
}

func TestMinItems(t *testing.T) {
	cases := []OptionTestCase{
		{
			schema: TypeOf([]string{}),
			value:  []string{},
			err:    true,
		},
		{
			schema: TypeOf([]string{}),
			value:  []string{"a"},
			err:    false,
		},
		{
			schema: TypeOf([]string{}),
			value:  []string{"a", "b"},
			err:    false,
		},
	}
	for _, c := range cases {
		c.schema.MinItems(1)
		_, err := c.schema.Validate(c.value)
		if c.err != (err != nil) {
			t.Error("MinItems is not working")
		}
	}
}

func TestMaxProperties(t *testing.T) {
	cases := []OptionTestCase{
		{
			schema: TypeOf(map[string]string{}),
			value:  map[string]string{},
			err:    false,
		},
		{
			schema: TypeOf(map[string]string{}),
			value:  map[string]string{"a": "a"},
			err:    false,
		},
		{
			schema: TypeOf(map[string]string{}),
			value:  map[string]string{"a": "a", "b": "b"},
			err:    true,
		},
	}
	for _, c := range cases {
		c.schema.MaxProperties(1)
		_, err := c.schema.Validate(c.value)
		if c.err != (err != nil) {
			t.Error("MaxProperties is not working")
		}
	}
}

func TestMinProperties(t *testing.T) {
	cases := []OptionTestCase{
		{
			schema: TypeOf(map[string]string{}),
			value:  map[string]string{},
			err:    true,
		},
		{
			schema: TypeOf(map[string]string{}),
			value:  map[string]string{"a": "a"},
			err:    false,
		},
		{
			schema: TypeOf(map[string]string{}),
			value:  map[string]string{"a": "a", "b": "b"},
			err:    false,
		},
	}
	for _, c := range cases {
		c.schema.MinProperties(1)
		_, err := c.schema.Validate(c.value)
		if c.err != (err != nil) {
			t.Error("MinProperties is not working")
		}
	}
}

func TestSchemaFromStructTag(t *testing.T) {
	type test struct {
		A string            `json:"a,maxlen=10,minlen=1"`
		B int               `json:"b,max=3,min=1,multof=2"`
		C int               `json:"c,exclmax=3,exclmin=1"`
		D []string          `json:"d,maxItems=5,minItems=2"`
		E map[string]string `json:"e,maxprops=6,minprops=3"`
	}

	s := TypeOf(test{})

	if *s.data.Properties["a"].MaxLength != 10 {
		t.Error("maxlen tag is not working")
	}

	if *s.data.Properties["a"].MinLength != 1 {
		t.Error("minlen tag is not working")
	}

	if *s.data.Properties["b"].Maximum != 3 {
		t.Error("max tag is not working")
	}

	if *s.data.Properties["b"].Minimum != 1 {
		t.Error("min tag is not working")
	}

	if *s.data.Properties["b"].MultipleOf != 2 {
		t.Error("multof tag is not working")
	}

	if *s.data.Properties["c"].ExclusiveMaximum != 3 {
		t.Error("exclmax tag is not working")
	}

	if *s.data.Properties["c"].ExclusiveMinimum != 1 {
		t.Error("exclmax tag is not working")
	}

	if *s.data.Properties["d"].MaxItems != 5 {
		t.Error("maxItems tag is not working")
	}

	if *s.data.Properties["d"].MinItems != 2 {
		t.Error("minItems tag is not working")
	}

	if *s.data.Properties["e"].MaxProperties != 6 {
		t.Error("maxprops tag is not working")
	}

	if *s.data.Properties["e"].MinProperties != 3 {
		t.Error("minprops tag is not working")
	}
}
