package intent

import (
	"context"
	"fmt"
	"strings"

	"github.com/aaaton/golem/v4"
	"github.com/aaaton/golem/v4/dicts/en"
	"github.com/gavinmead/alfred/pkg/log"
	"github.com/jdkato/prose/v2"
	"go.uber.org/zap"
)

var Get = Action{
	Name:  "get",
	Words: []string{"get", "show", "display", "fetch"},
}

var Create = Action{
	Name:  "create",
	Words: []string{"create", "new", "build"},
}

var Delete = Action{
	Name:  "delete",
	Words: []string{"delete", "remove"},
}

var Update = Action{
	Name:  "update",
	Words: []string{"update"},
}

// DefaultActionFinder searches a preset action definitions
type DefaultActionFinder struct {
	Actions []Action
}

// Find will return the action based on provided word.  The word should be normalized to
// lowercase
func (d *DefaultActionFinder) Find(ctx context.Context, word string) (Action, error) {
	var toReturn Action

	normalizedWord := strings.ToLower(word)

	for _, action := range d.Actions {
		for _, word := range action.Words {
			if strings.Compare(normalizedWord, word) == 0 {
				return action, nil
			}
		}
	}

	return toReturn, fmt.Errorf("action not found")
}

//DefaultHandler is the default implementation of the intent request Handler
type DefaultHandler struct {
	ActionFinder ActionFinder
}

// Build will extract the statement into its constitutent parts
// and return the request
func (d *DefaultHandler) Build(ctx context.Context, statement string) (Request, error) {
	logger := log.GetLogger(ctx)
	logger.Info("processing", zap.String("statement", statement))

	doc, err := prose.NewDocument(statement, prose.WithExtraction(false))
	if err != nil {
		return Request{}, err
	}

	//Process tokens
	actionIdx := -1
	entityIdx := -1

	var request Request
	request.Raw = statement

	for i, t := range doc.Tokens() {
		//Look for a verb first
		if actionIdx == -1 {
			switch t.Tag {
			case "VBP":
				fallthrough
			case "VB":
				fallthrough
			case "NN":
				action, _ := d.ActionFinder.Find(ctx, t.Text)
				if len(action.Name) != 0 {
					actionIdx = i
					request.Action = action
				}
				continue
			default:
				continue
			}
		}

		if entityIdx == -1 {
			if strings.Contains(t.Tag, "NN") {
				entityIdx = i
				//Lemmatize the entity
				request.Entity = lemmatizer.Lemma(t.Text)
				logger.Debug("found entity", zap.Int("tokenIdx", entityIdx), zap.String("entity", request.Entity))

				if strings.Compare(t.Tag, "NNS") == 0 {
					logger.Debug("plural entity detected", zap.String("entity", request.Entity),
						zap.String("contextType", string(List)))

					if strings.Compare(request.Action.Name, Get.Name) == 0 {
						request.ContextType = None
					} else {
						request.ContextType = List
					}

				} else {
					request.ContextType = Word
				}
				continue
			} else {
				continue
			}
		}

		switch l := request.ContextType; l {
		case List:
			contextEntries := make([]string, 0)
			for _, entry := range doc.Tokens()[i:] {
				//go through all of the NN
				if strings.Compare(entry.Tag, "NN") == 0 {
					contextEntries = append(contextEntries, entry.Text)
				}
			}
			request.Context = contextEntries
			break
		case Word:
			//Get the next token if it exists
			request.Context = t.Text
		default:
			return request, fmt.Errorf("invalid statement: no context type found")
		}
		break
	}

	return request, nil
}

var actionFinder = &DefaultActionFinder{
	Actions: []Action{
		Get,
		Delete,
		Create,
		Update,
	},
}

var handler = &DefaultHandler{
	ActionFinder: actionFinder,
}

var lemmatizer = getLemmatizer()

// Handle is a wrapper around DefaultHandler to build a request
// from the provided statement
func Handle(ctx context.Context, statement string) (Request, error) {
	return handler.Build(ctx, statement)
}

func getLemmatizer() *golem.Lemmatizer {
	lemmatizer, err := golem.New(en.New())
	if err != nil {
		panic(err)
	}
	return lemmatizer
}
