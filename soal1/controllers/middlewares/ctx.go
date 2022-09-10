package middlewares

type ctxKey int

const (
	CtxKeyUserID   ctxKey = iota
	CtxKeyUserName ctxKey = iota
	// ...
)
