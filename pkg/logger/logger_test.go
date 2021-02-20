package logger

import (
	"context"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func Test_LoggerInContext(t *testing.T) {
	assert := assert.New(t)

	l, _ := zap.NewProduction()
	ctx := context.WithValue(context.TODO(), Key, l)

	assert.Equal(l, GetLogger(ctx))
}

func Test_LoggerNotInContext(t *testing.T) {
	assert := assert.New(t)

	assert.PanicsWithValuef(fmt.Sprintf("no logger found at key %s", Key),
		func() { GetLogger(context.TODO()) }, "failed to panic")

}

func Test_NotALogger(t *testing.T) {
	assert := assert.New(t)

	ctx := context.WithValue(context.TODO(), Key, "will_fail")

	assert.PanicsWithValuef(fmt.Sprintf("type at at key %s is not a zap logger", Key),
		func() { GetLogger(ctx) }, "failed to panic")

}
