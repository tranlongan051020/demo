package constant

var ContextKey = struct {
	Language       string
	RequestTraceId string
	UserId         string
}{
	Language:       "language",
	RequestTraceId: "requestTraceId",
	UserId:         "userId",
}
