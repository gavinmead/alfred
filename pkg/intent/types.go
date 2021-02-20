package intent

import (
	"context"
)

// ContextType identifies the type of information included in the context of the Request
type ContextType string

const (

	//Word is the context type of a single word
	Word ContextType = "word"

	//List is a context of a comma separated list
	List ContextType = "list"

	//Phase is a set of words
	Phrase ContextType = "phrase"
)

// Request contains the
type Request struct {
	Action string
	Entity string

	//ContentType helps the skill convert the Context to a specific type
	ContextType ContextType

	//Context is useful information needed by skill like the name of an Entity
	Context interface{}

	//Raw is the original statement
	Raw string
}

// Handler is responsible for taking a statement and breaking into the core components
// necessary for Alfred to invoke a skill.  The context depends on the action
// but an example would be having a list of things separated by a comma.
type Handler interface {

	// Build will take a statement and construct the intent
	// If no request can be created then an error is returned
	Build(context.Context, string) (Request, error)
}
