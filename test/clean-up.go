package test

import (
	"context"
	"os"

	"github.com/cucumber/godog"
)

func AddStepCleanUp(sc *godog.ScenarioContext) {
	sc.After(cleanUp)

	return
}

func cleanUp(ctx0 context.Context, scenario *godog.Scenario, e0 error) (
	ctx context.Context, e error,
) {
	ctx = ctx0

	e = ctx.Value(ctxKeySrvProc{}).(*os.Process).Kill()
	if e != nil {
		return
	}

	e = os.RemoveAll(
		ctx.Value(ctxKeyTempDir{}).(string),
	)
	if e != nil {
		return
	}

	return
}
