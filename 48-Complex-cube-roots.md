Advanced Exercise: Complex cube roots
=====================================

complex64 と complex128 の型による、 Go言語の複素数( complex )について見てみましょう。 立方根( cube root )の場合、ニュートン法は、以下の式を繰り返すことになります：

![](http://go-tour-jp.appspot.com/static/newton3.png)

"2"の立方根を探し、アルゴリズムが正しく動作するか確認してください。 math/cmplx パッケージに [Pow](http://golang.org/pkg/math/cmplx/#Pow) 関数がありますので、結果を確認してみましょう。
