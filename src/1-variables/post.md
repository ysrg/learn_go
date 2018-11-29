# Variables

Declaring variables that are set to their zero values can be done like this:

```go
var a int
var b string
var c float64
var d bool
```

Go has a concept of ‘zero value’. What this means is that every value or variable we create must be initialized, and if we don’t do the initialization ourselves, it gets initialized with ‘zero value’.  The entire allocation of memory for the specific variable type (4,8bytes, etc) , every bit of that will be reset to zero —  `int` is assigned zero, `string` - an empty string, `float` - float and `bool`— false (1 byte). You can see that we don’t necessarily have to specify the bit size of the `int`  type  (8,32,64 bits), Go is smart enough to look at our hardware architecture and check the pointer (address) size and infer it automatically.Eg — on 64-bit machines — it will be 64bit `int`s.

Another way to declare variables is the shorthand `:=` operator. Besides declaring, it also initializes it.

```go
aa := 19
bb := "hi"
cc := 3.1415
dd := true
```
