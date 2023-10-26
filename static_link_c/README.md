ref

https://stackoverflow.com/questions/70835519/using-a-cgo-shared-library-in-a-go-program

https://github.com/hajsf/tutorial/tree/master/ffi/goCallCLib-staticLink

```
[~/go/src/github.com/ludwig125/cgo-sample/static_link_c] $gcc -shared -fpic lib.c -o ./libadd.a
[~/go/src/github.com/ludwig125/cgo-sample/static_link_c] $
```

```
[~/go/src/github.com/ludwig125/cgo-sample/static_link_c] $go run main.go
Welcome from external C function
3
```
