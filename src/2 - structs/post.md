# Struct types

Go allows you to declare your own user defined types. In other languages those are called classes, in go - structs.

```go
type example struct {
	flag bool
	counter int16
	pi float32
}
```

Now, let us suppose we use the `example` struct later to initialize a new variable `e1`  of type example set to its zero value:

```go
func main() {
	var e1 example
}
```

The question now is — how much memory is being allocated for `e1` ? A struct is a contiguous block of memory and based on its fields we have a `bool` (1byte),  a `int16` (2bytes)  and a `float32` (4bytes) => so we have 7 bytes total for our struct, right? No, actually we have 8 bytes, because of paddings and alignments. In our case we will have a one padding byte between the `bool` and `int16`. What Go says about alignment is that any `n` type value must fall on a `n` B boundary.  In our example our `int16` must fall on a 2B boundary, but as `bool` is 1B - hence the padding; also, if we have a 8B value - it has to fall on a 8B boundary - in our case we’d have 7B padding.  So,  all 8B will have zero values (`bool` - `false`, `int` and `float` - `0`.  A good rule of thumb is to lay our struct fields from highest to smallest:

```go
type example struct {
	counter1 int64
	counter2 int64
	counter3 int64
	pi float32
	flag1 bool
	flag2 bool
}
```

 This will push any padding down to the bottom. The rule is — the largest field represents the padding for the entire struct . This is more an optimization for performance, not correctness.
Now lets use the method where we declare a variable of type `example` and initialize it at the same time using a *struct literal*:

```go
e2 := example{
	flat: true,
	counter: 10,
	pi: 3.14,
}
```

Also, we can declare a variable of an anonymous type and init it using a struct literal (anonymous structs) :

```go
e := struct {
	flag bool
	counter int16
	pi float32
}{
	flag: true,
	counter: 10,
	pi: 3.14,
}
```

To see if two structs are identical you’ll need explicit conversion. Anonymous types, on the other hand you, can be compared more freely.

```go
var greet struct {
	flag true
	pi float32
}

var cat struct {
	flag true
	pi float32
}
var b greet
var c cat
b = c //compiler error cant use c as type greet in assignment

//we have to specify the conversion explicitly
b = greet(c)

//for anonynous structs this works without explicit conversion

b = e //works as expected as e is not based on an unamed type byt is an anonymous struct type
```
This will be very useful when passing around anonymous functions (eg working with http requests, functions are first-class in Go).



