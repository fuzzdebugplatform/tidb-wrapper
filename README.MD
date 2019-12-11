# tidb-wrapper


## compile this project

```
# first you should install go-bindata
go get -u github.com/jteeuwen/go-bindata/...
# then make the project
make
```

## wrap the example project in this repo


```
./tidb-wrapper -D example-project/
```

It will generates a subdirectory named `mywrapper`, in which is the generated code. Run this project,

Trace is closed in default, you need open it with `switch` interface (trace interfaces listen on 43222):

```
http://localhost:43222/switch
```

Visit app interface of example project:

```
http://localhost:4333/test?key=9
# return 10
```

See all traced inputs and corresponding digest:

```
http://localhost:43222/status
```

```json
{
  "248be6a1": "9"
}
```

See trace of input `key = 9`:

```
http://localhost:43222/trace/248be6a1
```

```json
{
  "sql": "9",
  "trace": [
    {
      "file": "cmd/server_start.go",
      "line": null
    },
    {
      "file": "protocol/protocol.go",
      "line": null
    },
    {
      "file": "server/server.go",
      "line": [
        [9, 11],
        [15,15],
        [17,19],
        [21,21]
      ]
    }
  ]
}
```
