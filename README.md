# jsontagchecker
jsontagcheckerは構造体の
jsonタグがフィールド名のスネークケースで記述されているかを調べる静的解析ツールです。

例
```go

type Info struct {
	UserId string `json:id` // Bad "user_id" is correct.
	UserName string `json: user_name` // OK
}
```
## setup
```
$ git clone https://github.com/hikaru7719/jsontagchecker.git
$ go install cmd/jsontagchecker/main.go
```
でバイナリが生成されます。


`jsontagchecker [ファイル名/ディレクトリ名]`
で実行できます。

```
$ jsontagchecker sample/my.go
jsontagchecker/sample/my.go:5:21: invalid snake case json tag
```

