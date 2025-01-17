package test

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/minio/minio-go/v7"
)

func AddStepListObjectsAfter(sc *godog.ScenarioContext) {
	sc.When(
		`^I list the objects in bucket "([^"]+)" starting after "([^"]+)"$`,
		listObjectsAfter,
	)

	return
}

func listObjectsAfter(ctx0 context.Context, bucketName, startAfter string) (
	ctx context.Context, e error,
) {
	ctx = ctx0

	var (
		objects <-chan minio.ObjectInfo
		client  *minio.Client = ctx.Value(ctxKeyMClient{}).(*minio.Client)
	)

	objects = client.ListObjects(ctx, bucketName,
		minio.ListObjectsOptions{
			StartAfter: startAfter,
		},
	)

	ctx = context.WithValue(ctx, ctxKeyObjects{}, objects)

	return
}
