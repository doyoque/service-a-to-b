package deposithttp

import "go.uber.org/fx"

var Module = fx.Module("api-deposit", fx.Invoke(
	Create,
))
