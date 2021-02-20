package skill

import (
	"context"

	"github.com/gavinmead/alfred/pkg/intent"
)

// Skill is the core interface for Alfred to do something when a request
// is received
type Skill interface {

	// Supports will return true if this skill can process this request
	Supports(context.Context, intent.Request) bool
}
