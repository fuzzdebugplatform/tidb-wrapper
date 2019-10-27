# tidb-wrapper


## 编译该项目

```
# first you should install go-bindata
go get -u github.com/jteeuwen/go-bindata/...
# then make the project
make
```

## 打包示例项目


```
./tidb-wrapper -D example-project/
```

它将生成`mywrapper`子目录，里面就是生成的代码，运行该项目。

默认情况下，统计功能是关闭的，需要使用`switch`接口开启（统计接口监听在43222）：

```
http://localhost:43222/switch
```

访问示例项目的接口：

```
http://localhost:4333/test?key=9
# return 10
```


查看当前所有在统计的数据与摘要：

```
http://localhost:4333/status
```

```json
{
  "248be6a1": "9"
}
```

查看刚刚 key = 9 所覆盖的文件路径:

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