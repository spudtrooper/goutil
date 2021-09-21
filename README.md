# github.com/spudtrooper/goutil

Dumping ground for go libraries.

## errors
Accumulate and join errors.

## or
Logical or two values with defaults, e.g.

```
v := or.Int(a, b)
```

is the same as

```
v := a
if v == 0 {
    v = b
}
```

## cond
Emulate the ternary `? :`

```
v := cond.Int(b, a, z)
```

is the same as

```
v := a
if b {
    v = z
}
```

## lazycond
Emulate the ternary `? :` lazily

Given

```
a := func() int { return 1 }
z := func() int { return 2 }
```

```
v := lazycond.Int(b, a, z)
```

is the same as

```
v := a()
if b {
    v = z()
}
```

## internal

To generate `selenium/seleniumserver.go` run:

```
go run writeseleniumjar.go
```