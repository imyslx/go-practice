# About this

This is for practice to write golang.  
You can use this code freely, but it will have a lot of bugs.

Goの練習用コードなので必要なら好きに使っていただいて結構ですが保証はしません。

# Codes

## timer

* テーマ
  * とりあえず書いてみる。
  * gorotine使ってみる。 
  * 量産しようとしたら構造体必要になった。

構造体にTickerとMethodを組み込んで内部でgoroutineを生み出す構造。

mainで構造体MonTickerを量産し、個々のMonTickerのworkメソッドを起動させる。  
このworkメソッドの中でgoroutineを生み出し、MonTickがもつTickerの通知を基に処理を行う。

