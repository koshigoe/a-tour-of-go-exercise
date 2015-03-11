Exercise: HTTP Handlers
=======================

以下の型を実装し、それらにServeHTTPメソッドを定義してみてください。 自分のWebサーバで、特定のパスを処理（ハンドル）できるように登録してみてください。

```go
type String string

type Struct struct {
    Greeting string
    Punct    string
    Who      string
}
```
例えば、以下のようなハンドラを用いて登録できるようにする必要があります。

```go
http.Handle("/string", String("I'm a frayed knot."))
http.Handle("/struct", &Struct{"Hello", ":", "Gophers!"})
```
