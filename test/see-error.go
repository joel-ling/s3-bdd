package test

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

func AddStepSeeError(sc *godog.ScenarioContext) {
	sc.Then(`^I should see an error "([^"]+)"$`,
		seeError,
	)

	return
}

func seeError(ctx0 context.Context, errorString string) (
	ctx context.Context, e error,
) {
	ctx = ctx0

	assert.Equal(
		godog.T(ctx),
		errorString,
		ctx.Value(ctxKeyS3Error{}).(error).Error(),
	)

	return
}
