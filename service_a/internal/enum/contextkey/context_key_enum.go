package contextkeyenum

type ContextKey string

const (
	Authorization ContextKey = "AUTHORIZATION_CONTEXT"
	Resource      ContextKey = "RESOURCE_CONTEXT"
)
