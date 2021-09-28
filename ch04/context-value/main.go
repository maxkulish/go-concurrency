package main

import (
	"context"
	"fmt"
)

type ctxKey int

const (
	ctxUserID ctxKey = iota
	ctxAuthToken
)

func UserID(c context.Context) string {
	return c.Value(ctxUserID).(string)
}

func AuthToken(c context.Context) string {
	return c.Value(ctxAuthToken).(string)
}

func main() {
	ProcessRequest("jane", "abc123")

	type foo int
	type bar int

	m := make(map[interface{}]int)
	m[foo(1)] = 1
	m[bar(1)] = 2

	fmt.Printf("%v\n", m)

	ProcessRequest("netstor", "gsdt1634")
}

func ProcessRequest(userID, authToken string) {
	ctx := context.WithValue(context.Background(), ctxUserID, userID)
	ctx = context.WithValue(ctx, ctxAuthToken, authToken)
	HandleResponse(ctx)
}

func HandleResponse(ctx context.Context) {
	fmt.Printf(
		"handling response for %v (%v)\n",
		UserID(ctx),
		AuthToken(ctx),
	)
}
