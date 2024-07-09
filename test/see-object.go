package test

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/minio/minio-go/v7"
	"github.com/stretchr/testify/assert"
)

func AddStepSeeObject(sc *godog.ScenarioContext) {
	sc.Then(`^I should see an object "([^"]+)" of size (\d+) bytes$`,
		seeObject,
	)

	return
}

func seeObject(ctx0 context.Context, objectName string, objectSize int64) (
	ctx context.Context, e error,
) {
	ctx = ctx0

	var (
		object  minio.ObjectInfo
		objects = ctx.Value(ctxKeyObjects{}).(<-chan minio.ObjectInfo)
	)

	object = <-objects

	assert.Equal(
		godog.T(ctx),
		objectName,
		object.Key,
	)

	assert.Equal(
		godog.T(ctx),
		objectSize,
		object.Size,
	)

	return
}
