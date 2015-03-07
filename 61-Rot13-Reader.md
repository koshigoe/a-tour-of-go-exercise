Exercise: Rot13 Reader
======================

ある共通のパターンは、[io.Reader](http://golang.org/pkg/io/#Reader)です。 これは、何らかの方法でストリームを変更する別の io.Reader をラップしています。

例えば、 [gzip.NewReader](http://golang.org/pkg/compress/gzip/#NewReader) は、 io.Reader （gzipされたデータのストリーム）を引数で受け取り、 *gzip.Reader を返します。 その *gzip.Reader は、 io.Reader (展開されたデータのストリーム)を実装しています。

io.Reader を実装し、 io.Reader から読み出すように rot13Reader を実装してみてください。 なお、 rot13Reader は、 [ROT13](http://ja.wikipedia.org/wiki/ROT13) 換字式暗号( substitution cipher )をすべてのアルファベットの文字に適用してください。

rot13Reader 型は、みなさんに提供します。 この rot13Reader は、それの Read メソッドを実装することで io.Reader を作成してみてください。
