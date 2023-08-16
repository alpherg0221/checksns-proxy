## ビルド方法
1. gomobileを入れる ([参考](https://pkg.go.dev/golang.org/x/mobile/cmd/gomobile))
2. `gomobile bind -target android -androidapi 19` でandroid向けにビルドする
3. 作成されたaarファイルをAndroid側でライブラリとして読み込む