# can't load package: import cycle not allowed

```bash
$ go build ./...                              [一  6/15 10:14:23 2020]
can't load package: import cycle not allowed
package importcycle/a
        imports importcycle/b
        imports importcycle/a
$ go version                                  [一  6/15 10:14:29 2020]
go version go1.14.2 darwin/amd64
```

```bash
$ go get -u -x  golang.org/dl/go1.12.5        [一  6/15 10:35:33 2020]
$ go1.12.5 download                           [一  6/15 10:35:49 2020]
$ go1.12.5 build ./...                        [一  6/15 10:35:49 2020]
can't load package: import cycle not allowed
package importcycle/a
	imports importcycle/b
	imports importcycle/a
```
