# tmpl

Code generation from pure Go templates.

## Example

```go
var template = `package template

func Foo(x int) int {
	return x + ConstAdditive
}
`

func main() {
	t, err := tmpl.ParseFile("", []byte(template))
	if err != nil {
		log.Fatal(err)
	}

	err = t.Apply(
		tmpl.SetPackageName("arith"),
		tmpl.Rename("Foo", "Add"),
		tmpl.DefineIntDecimal("ConstAdditive", 42),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = t.Format(os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
```

Output:

```go
package arith

func Add(x int) int {
	return x + 42
}
```
