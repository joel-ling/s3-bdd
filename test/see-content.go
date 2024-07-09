package test

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

func AddStepSeeContent(sc *godog.ScenarioContext) {
	sc.Then(`^I should see content "(.*)"$`,
		seeContent,
	)

	return
}

func seeContent(ctx0 context.Context, content string) (
	ctx context.Context, e error,
) {
	ctx = ctx0

	var (
		contentBytes []byte = ctx.Value(ctxKeyContent{}).([]byte)
	)

	assert.Equal(
		godog.T(ctx),
		content,
		string(contentBytes),
	)

	return
}
