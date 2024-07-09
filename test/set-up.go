package test

import (
	"context"
	"os"

	"github.com/cucumber/godog"
)

func AddStepSetUp(sc *godog.ScenarioContext) {
	sc.Before(setUp)

	return
}

func setUp(ctx0 context.Context, scenario *godog.Scenario) (
	ctx context.Context, e error,
) {
	ctx = ctx0

	var (
		path string
	)

	path, e = downloadMinio()
	if e != nil {
		return
	}

	ctx = context.WithValue(ctx, ctxKeyBinPath{}, path)

	path, e = os.MkdirTemp("", "")
	if e != nil {
		return
	}

	ctx = context.WithValue(ctx, ctxKeyTempDir{}, path)

	return
}
