package main

import (
	"mycms/z"
)

func main() {
	api := z.NewRouter("/api", []string{"POST"}, z.Procedure{z.SubProcedure(middle)})
	api.PushChild(
		&z.Router{Pattern: "/test", Methods: []string{"GET"}, Procedure: z.Procedure{z.SubProcedure(getTest)}},
		&z.Router{Pattern: "/test", Methods: []string{"POST"}, Procedure: z.Procedure{z.SubProcedure(postTest)}},
	)
	z.Run(api, ":3000")
}

func middle(ctx *z.ZContext) bool {
	ctx.Data = "middle"
	return true
}

func getTest(ctx *z.ZContext) bool {
	ctx.W.Write([]byte(ctx.Data.(string) + "GET TEST"))
	return true
}

func postTest(ctx *z.ZContext) bool {
	ctx.W.Write([]byte("POST TEST"))
	return true
}
