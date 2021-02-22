package intent

import (
	"context"
)

// Action contains the base action for a skill and
// additional words that are similar to the action.
type Action struct {
	Name  string
	Words []string
}

// ActionFinder searches its database for the Action associated
// with the provided word.  Returns an error if no action was found
type ActionFinder interface {
	Find(context.Context, string) (Action, error)
}

// ContextType identifies the type of information included in the context of the Request
type ContextType string

const (

	//Word is the context type of a single word
	Word ContextType = "word"

	//List is a context of a comma separated list
	List ContextType = "list"

	//Phrase is a set of words
	Phrase ContextType = "phrase"

	None ContextType = "none"
)

// Request contains the
type Request struct {
	Action Action
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
