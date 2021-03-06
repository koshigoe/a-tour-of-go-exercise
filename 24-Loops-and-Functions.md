Exercise: Loops and Functions
=============================

関数とループを使った簡単な練習として、[ニュートン法](http://ja.wikipedia.org/wiki/%E3%83%8B%E3%83%A5%E3%83%BC%E3%83%88%E3%83%B3%E6%B3%95)を使った平方根の計算を実装してみましょう。

この問題では、ニュートン法は、 開始点 z を選び、以下の式を繰り返すことによって、 Sqrt(x) を近似します。

![](http://go-tour-jp.appspot.com/static/newton.png)

最初は、その計算式を10回だけ繰り返し、 x を(1, 2, 3, ...)と様々な値に対する結果がどれだけ正解値に近いかを確認してみてください。

次に、ループを回すときの直前に求めたzの値がこれ以上変化しなくなったとき （または、差がとても小さくなったとき） に停止するようにループを変更してみてください。 この変更により、ループ回数が多くなったか、少なくなったのか見てみてください。 [math.Sqrt](http://golang.org/pkg/math/#Sqrt) と比べてどれくらい近似できましたか？

ヒント：浮動小数点を宣言し、値を初期化するには、型のキャストか、浮動小数点を使ってください：

```go
z := float64(1)
z := 1.0
````
