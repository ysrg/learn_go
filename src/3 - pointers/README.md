# 3 - Pointers

Everything in Go is about **pass by value**. Pointers serve a very important use case - and that is *sharing* a value across a program boundary, the most common one will be sharing across function calls or go routines. So, to share a value, we need a pointer.

```go
func main() {
  //declare a variable of type int with a value of 10
  count := 10

  //display the 'value of' and 'address of' count
  Println("Before:", count, &count)

  //pass the 'value of' the variable count
  increment(count)
  Println("After", count, &count)
}

func increment(inc int) {
  inc++
  Println("Inc: ", inc, &inc)
}

```

When this program starts up, the Go runtime is going to create what we call a goroutine. Think of a goroutine as a:

* separate path of execution that contains the instructions we write that need to get executed by the machine
* thread (for now) -- the Go scheduler at some point has to place that goroutine on an OS thread so it can execute.

So, our previous program contains one goroutine (the main goroutine) that is going to be created by the runtime and therefore executed. Every goroutine is given a block of memory that we call a stack. In Go the stack memory starts at 2K. To put this in perspective, every OS thread generally is given 1MB of memory. Every time a function is called, a piece of the stack is used to help the function run. At some point the runtime is going to have the `main` goroutine execute the `main` function and when it does it will take a piece of memory from the stack -- called *stack frame*. The stack frame is known at compile time, this means that no value can be placed on a stack unless the compiler knows the size of it ahead of time. If we dont know the size of something at compile time it has to be on this other piece of memory that we call the *heap*.
In our case, we know the size of `count` variable -- its 4B on a 32b machine, and we put the value 10 in those 4B of memory on the stack frame for `main`. Knowing that, we can go back to our example and see the result:

```go
Before: 10 0x10347fa2
Inc:    11 0x10347e39
After:  10 0x10347fa2
```

Let's pass the address of `count` to `increment` function, not the value like earlier. This is still passed by value (the address is a value), not by reference:

```go
increment(&count)
```

To store the address in the function, we need a variable that can store the address, and this is where pointer variables come in. The `inc` argument is not of type `int` anymore but of type `*int`. For every type that is declared, you get for free a pointer type, that has its name starting with a `*`. This means that the pointer type declares a variable whose sole purpose is to store the address value. When passing address values like `&count` across program boundaries it is pointer variable's (`*int`) job to store that value.

```go
import fmt
func increment(inc *int) {
  *inc++
  fmt.Println("Inc: ", *inc, &inc, inc) // inc is the value that the pointer points to
}
```

`*` is just the value the pointer points to. So now we can increment the value the pointer points to, this means we can manipulate memory outside our stack frame thanks to the pointer.

```go
Before: 10 0x10347fa2
Inc:    11 0x10347e39 0x10347fa2
After:  11 0x10347fa2
```

Let us look at another example:

```go
//user struct represents a user in the system
type user struct {
  name string
  email string
}

//main is the entry point for our app
func main() {
  stayOnStack()
  escapeToHeap()
}

//stayOnStack shows how the variable does not escape
func stayOnStack() user {
  u := user{
    name: "John",
    email: "john@foo.com",
  }
  return u
}

//escapeToHeap shows how the variable does escape
func escapeToHeap() *user {
  u := user{
    name: "John",
    email: "john@foo.com",
  }
  return &u
}
```

As we ca see, we make a `user` copy in `stayOnStack` and pass it (`return`) higher on the stack to the `main` function (value semantics). In `escapeToHeap` we dont return the value like in `stayOnStack`, but the address of `u` (pointer semantics). That is the value passed back up the call stack. Remember, *passed by value* means a value is always going to be copied and moved around, either the value itself or the address. You might think that after this call the `main` stack frame will now have a pointer [`*`] -> `escapeToHeap` to a value that is on the stack frame below (`escapeToHeap`). This is not the case because once we go up the stack the previous stack frames are discarded and now we have a pointer to something that is about to get erased. The memory is still there but is no longer valid, so we can't actually point to something that is not there anymore. Stacks are self-cleaning data structures. Every time we make a function call we clean the stack on the way down and initialize it with zero value. That is why it will have to be placed on the **heap**, this is what is called 'shared value' -- every stack frame can access the value on the heap. `main` is going to have a pointer to the value that is now on the heap. `escapeToHeap` function (from within the stack frame) itself is going to have a pointer to the heap. Escape analysis make the decision on what should stay on the stack and what on the heap.

Let's look at another example:

```go
func main() {
  x := 4
  fmt.Println(x)
  foo(&x) // `&` means pass the address of `x` to foo, not the `x` itself
  fmt.Println(x)
}

func foo(y *int) {
  *y := 5 // `*` is an operator here an is used to assign VALUE to that address
  fmt.Println(*y)
}

// 4
// 5 (foo prints 5 as the value of that address was changed)
// 5
```
