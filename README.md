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
``