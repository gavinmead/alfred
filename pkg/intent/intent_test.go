package intent

import (
	"context"
	"testing"

	"github.com/gavinmead/alfred/pkg/log"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
)

type IntentTestSuite struct {
	suite.Suite
	ctx context.Context
}

func (s *IntentTestSuite) SetupTest() {
	logger, _ := zap.NewProduction()
	s.ctx = context.WithValue(context.TODO(), log.Key, logger)
}

func (s *IntentTestSuite) TestIntentBuild() {
	var testTable = []struct {
		statement string
		expected  Request
	}{
		{
			"@alfred please create the tags: test1, test2 and test3",
			Request{
				Action:      Create,
				Entity:      "tag",
				ContextType: List,
				Context:     []string{"test1", "test2", "test3"},
				Raw:         "@alfred please create the tags: test1, test2 and test3",
			},
		},
		{
			"@alfred create the tags: test1, test2 and test3",
			Request{
				Action:      Create,
				Entity:      "tag",
				ContextType: List,
				Context:     []string{"test1", "test2", "test3"},
				Raw:         "@alfred create the tags: test1, test2 and test3",
			},
		},
		{
			"@alfred please create tag test1",
			Request{
				Action:      Create,
				Entity:      "tag",
				ContextType: Word,
				Context:     "test1",
				Raw:         "@alfred please create tag test1",
			},
		},
		{
			"@alfred create the tag test1",
			Request{
				Action:      Create,
				Entity:      "tag",
				ContextType: Word,
				Context:     "test1",
				Raw:         "@alfred create the tag test1",
			},
		},
		{
			"@alfred create the tags test1 and test2",
			Request{
				Action:      Create,
				Entity:      "tag",
				ContextType: List,
				Context:     []string{"test1", "test2"},
				Raw:         "@alfred create the tags test1 and test2",
			},
		},
		{
			"@alfred please delete the tag test1",
			Request{
				Action:      Delete,
				Entity:      "tag",
				ContextType: Word,
				Context:     "test1",
				Raw:         "@alfred please delete the tag test1",
			},
		},
		{
			"@alfred show tags",
			Request{
				Action:      Get,
				Entity:      "tag",
				ContextType: None,
				Context:     nil,
				Raw:         "@alfred show tags",
			},
		},
	}

	assert := assert.New(s.T())
	for _, v := range testTable {
		actual, err := Handle(s.ctx, v.statement)
		assert.NoError(err)
		assert.Equal(v.expected, actual)
	}

}

func Test_IntentSuite(t *testing.T) {
	suite.Run(t, new(IntentTestSuite))
}
