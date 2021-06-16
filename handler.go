package handlers

import (
	"context"
	"fmt"
)

func HandlerFunc(ctx context.Context, input string) (string, error) {
	return fmt.Sprintf("Path %s!", input), nil
}
