package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("goa Practice", func() {
	Title("imyslx/goa")
	Description("goaの練習用サンプルです")
	Docs(func() {
		Description("wiki")
		URL("https://github.com/imyslx/go-practice/goa/README.md")
	})

	// ホストの設定
	Host("localhost:8081")
	Scheme("http")
	BasePath("/api/v1")

	//　CORSポリシーの定義
	Origin("http://localhost:8081/swagger", func() {
		Methods("GET", "POST", "PUT", "DELETE")
		//プリフライト要求応答をキャッシュする時間
		MaxAge(600)
		// Access-Control-Allow-Credentialsヘッダーを設定する
		Credentials()
	})
})

// RespMedia : レスポンスデータの定義
var RespMedia = MediaType("application/vnd.result+json", func() {
	Description("Practice for goa API.")
	// どのような値があるか（複数定義出来る）
	Attributes(func() {
		Attribute("status", Integer, "status")
		Attribute("message", String, "message")
		Required("status", "message")
	})
	// 返すレスポンスのフォーマット（別記事で紹介予定）
	View("default", func() {
		Attribute("status")
		Attribute("message")
	})
})
