package test

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/minio/minio-go/v7"
)

func AddStepListObjectsPrefix(sc *godog.ScenarioContext) {
	sc.When(`^I list the objects in bucket "([^"]+)" with prefix "([^"]+)"$`,
		listObjectsPrefix,
	)

	return
}

func listObjectsPrefix(ctx0 context.Context, bucketName, prefix string) (
	ctx context.Context, e error,
) {
	ctx = ctx0

	var (
		objects <-chan minio.ObjectInfo
		client  *minio.Client = ctx.Value(ctxKeyMClient{}).(*minio.Client)
	)

	objects = client.ListObjects(ctx, bucketName,
		minio.ListObjectsOptions{
			Prefix: prefix,
		},
	)

	ctx = context.WithValue(ctx, ctxKeyObjects{}, objects)

	return
}
