package deposit

import "context"

type Service interface {
	Create(ctx context.Context) error
}
