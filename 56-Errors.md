Exercise: Errors
================

Sqrt 関数を以前の演習からコピーし、 error の値を返すように修正してみてください。

Sqrt は、複素数をサポートしていないので、負の値が与えられたとき、nil以外のエラー値を返す必要があります。

新しいタイプの

```go
type ErrNegativeSqrt float64
```
を作成してみてください。

そして、 error ErrNegativeSqrt(-2).String.Error() で、 "cannot Sqrt negative number: -2" を返すような

```go
func (e ErrNegativeSqrt) Error() string
```
メソッドを作ります。

**注意**： Error メソッド中で、 fmt.Print(e) を呼び出すことは、無限ループにプログラムが陥ることでしょう。

最初に、 fmt.Print(float64(e)) として e を変換することより、これを避けることができます。 なぜでしょうか？

負の値が与えられたとき、 ErrNegativeSqrt の値を返すように Sqrt 関数を修正してみてください。
