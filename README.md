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

## tcpconn

* テーマ
  * TCPコネクション貼りたいだけのアプリケーション
  * Server - Client
  * データの送受信の仕組みが何となく掴めるレベル感で。

その名の通りTCPコネクションを張るためのアプリケーション。
Server側はもう少し手を入れないと使えないかな。エラー処理とかすごく心配。

## goa

* テーマ
  * goaによるAPI構築の練習
  * とりあえず適当にJSONをPOSTする仕組みと、レスポンスを返す仕組み作り。

そのまんまですが、goaの練習用です。
JSONをPOSTして、そのデータを元に何か処理するってところまで作りたかったので一旦そこまで。
よくわかっていないので色々ゴミも残ってる予感。
