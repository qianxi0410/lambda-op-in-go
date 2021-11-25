# lambda-op-in-go

[wip] lambda func impl in go

> 一见到函数是一等公民，立刻想到函数式，立刻想到不可变数据，立刻想到闭包，立刻想到高阶函数，立刻想到lambda演算，立刻想到数学，立刻想到自己是小天才，某些程序员的想像惟在这一层能够如此跃进。
>  
>  
> by Treeman Zhou

## doc

### map

```go
s := stream.Init("i love you").Map(func(idx int, e stream.Elem) stream.Elem {
    return unicode.ToUpper(e.(rune))
}).String()
// expect to be: I LOVE YOU

ints := stream.Init([]int{1, 2, 3}).Map(func(idx int, e stream.Elem) stream.Elem {
    return e.(int) * 3
}).IntSlice()
// expect to be: [3, 6, 9]
```

### foeach

```go
stream.Init([]int{1, 2, 3, 4}).Foreach(func(idx int, e stream.Elem) {
    fmt.Println(fmt.Sprintf("this is %d", e.(int)))
})
```

### reverse

```go
stream.Init([]int{1, 2, 3, 4}).Reverse().IntSlice()
// expect to be [4, 3, 2, 1]
```

### filter

```go
is := stream.Init([]int{1, 2, 3, 4}).Filter(func(idx int, e stream.Elem) bool {
    return e.(int)%2 == 0
}).IntSlice()
// expect to be [2, 4]
```
