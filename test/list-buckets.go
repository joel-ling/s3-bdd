package test

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/minio/minio-go/v7"
)

func AddStepListBuckets(sc *godog.ScenarioContext) {
	sc.When(`^I list all the buckets I own$`,
		listBuckets,
	)

	return
}

func listBuckets(ctx0 context.Context) (ctx context.Context, e error) {
	ctx = ctx0

	var (
		buckets []minio.BucketInfo
		client  *minio.Client = ctx.Value(ctxKeyMClient{}).(*minio.Client)
	)

	buckets, e = client.ListBuckets(ctx)
	if e != nil {
		return
	}

	ctx = context.WithValue(ctx, ctxKeyBuckets{}, buckets)

	return
}
