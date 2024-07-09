package test

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/minio/minio-go/v7"
)

func AddStepCreateBucket(sc *godog.ScenarioContext) {
	sc.Given(`^there is a bucket "([^"]+)"$`,
		createBucket,
	)

	sc.When(`^I create a bucket "([^"]+)"$`,
		createBucket,
	)

	return
}

func createBucket(ctx0 context.Context, bucketName string) (
	ctx context.Context, e error,
) {
	ctx = ctx0

	var (
		client *minio.Client = ctx.Value(ctxKeyMClient{}).(*minio.Client)
	)

	e = client.MakeBucket(ctx, bucketName,
		minio.MakeBucketOptions{},
	)
	if e != nil {
		return
	}

	return
}
