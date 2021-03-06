# テスト駆動開発でGO言語を学びましょう
こちらのサイトの勉強
https://andmorefine.gitbook.io/learn-go-with-tests/

## テストを作成する際のルール
 + xxx_test.goのような名前のファイルにある必要があります。
 + テスト関数はTestという単語で始まる必要があります。
 + テスト関数は1つの引数のみをとります。 t *testing.T
 + *testing.T 型を使うには、他のファイルの fmt と同じように import "testing" が必要です。


## TDDプロセスとそのステップが重要である理由
 + 失敗するテストを作成してそれを確認する要件に対応するrelevantテストを作成し、失敗の説明を簡単に理解できることを確認しました。
 + 機能するソフトウェアがあることを確認するために、最小限のコードを記述して合格させる。
 + リファクタリング、テストの安全性に裏打ちされており、操作が簡単な巧妙に作成されたコードがあることを確認します


## 同じパッケージに同じ名前の関数を作る場合
これは出来ない（エラーになる）
```
package main

func (r Rectangle) Area() float64  {
    return 0
}

func (c Circle) Area() float64  {
    return 0
}

```

パッケージを変えるか、下記のような新しく定義した型にメソッドを定義する
```
package main

func (r Rectangle) Area() float64 {
    return r.Width + r.Height
}

func (c Circle) Area() float64 {
    return 0
}

```
