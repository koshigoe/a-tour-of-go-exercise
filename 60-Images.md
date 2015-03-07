Exercise: Images
================

前に書いた、画像ジェネレーターを覚えてますか？ 他のものを書いてみましょう。でも今回は、データのスライスの代わりに image.Image の実装を返すようにしてみます。

Image 型を定義して、 [必要なメソッド](http://golang.org/pkg/image/#Image) を実装し、 pic.ShowImage を呼び出してみてください。

Bounds は、 image.Rect(0, 0, w, h) のように image.Rectangle を返す必要があります。

ColorModel は、 color.RGBAModel を返す必要があります。

At は、ひとつのcolorを返す必要があります。 最後の画像ジェネレーターの値 v は、 color.RGBA{v, v, 255, 255} のものと一致します。
