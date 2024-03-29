package tmpl

import (
	"bytes"
	"go/format"
	"testing"
)

var cases = []struct {
	Name       string
	Source     string
	Transforms []Transform
	Expect     string
}{
	{
		Name:       "none",
		Source:     "package none\nfunc Hello(x int) int { return x+1 }\n",
		Transforms: nil,
		Expect:     "package none\nfunc Hello(x int) int { return x+1 }\n",
	},
	{
		Name:   "set_package_name",
		Source: "package foo\nfunc Hello(x int) int { return x+1 }\n",
		Transforms: []Transform{
			SetPackageName("bar"),
		},
		Expect: "package bar\nfunc Hello(x int) int { return x+1 }\n",
	},
	{
		Name: "rename_function",
		Source: `package pkg
func Foo(x int) int { return x+1 }
func Bar(x int) int { return 2*Foo(x) }
`,
		Transforms: []Transform{
			Rename("Foo", "Baz"),
		},
		Expect: `package pkg
func Baz(x int) int { return x+1 }
func Bar(x int) int { return 2*Baz(x) }
`,
	},
	{
		Name: "define_string",
		Source: `package pkg
func Greet() string { return ConstGreeting }
`,
		Transforms: []Transform{
			DefineString("ConstGreeting", "Hello, World!"),
		},
		Expect: `package pkg
func Greet() string { return "Hello, World!" }
`,
	},
	{
		Name: "define_int",
		Source: `package pkg
func Foo(x int) int { return (x+ConstAdditive) / ConstDivisor }
`,
		Transforms: []Transform{
			DefineIntDecimal("ConstAdditive", 42),
			DefineIntHex("ConstDivisor", 0x37),
		},
		Expect: `package pkg
func Foo(x int) int { return (x+42)/0x37 }
`,
	},
	{
		// Regression test for a bug in earlier versions where variable replacement
		// would cause field references to be formatted incorrectly.
		Name: "rename_variable_formatting",
		Source: `package pkg
var x Thing
func init() {
	x.Field = 12
}
`,
		Transforms: []Transform{
			Rename("x", "baz"),
		},
		Expect: `package pkg
var baz Thing
func init() {
	baz.Field = 12
}
`,
	},
}

func TestTransforms(t *testing.T) {
	for _, c := range cases {
		c := c // scopelint
		t.Run(c.Name, func(t *testing.T) {
			tpl, err := ParseFile("", []byte(c.Source))
			if err != nil {
				t.Fatal(err)
			}

			err = tpl.Apply(c.Transforms...)
			if err != nil {
				t.Fatal(err)
			}

			got, err := tpl.Bytes()
			if err != nil {
				t.Fatal(err)
			}

			expect, err := format.Source([]byte(c.Expect))
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(expect, got) {
				t.Logf("got:\n%s", got)
				t.Logf("expect:\n%s", expect)
				t.Fatal("mismatch")
			}
		})
	}
}
