package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Hostを管理するリソース
var _ = Resource("hosts", func() {
	BasePath("/hosts")

	// Hostの追加用Action
	Action("add", func() {
		Description("Hostの追加")
		Routing(
			// Endpoint -> http://localhost:8081/api/v1/hosts/add
			POST("/add"),
		)

		// パラメーターの定義
		Payload(func() {
			Member("Hostname", String, "ホスト名")
			Member("Ipaddr", String, "IPアドレス", func() {
				Default("")
			})
			Required("Hostname", "Ipaddr")
		})

		Response(OK, RespMedia)
		Response(BadRequest, ErrorMedia)
	})

	// Hostの削除用アクション
	Action("delete", func() {
		Description("Hostの削除")
		Routing(
			POST("/delete"),
		)

		// パラメーターの定義
		Payload(func() {
			Member("Hostname", String, "ホスト名")
			Member("Ipaddr", String, "IPアドレス", func() {
				Default("")
			})
			Required("Hostname", "Ipaddr")
		})
		Response(OK, RespMedia)
		Response(BadRequest, ErrorMedia)
	})
})

// Swaggerをローカルで実行するめの定義
var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET")
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger/*filepath", "public/swagger/")
})
