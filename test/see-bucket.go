package test

import (
	"context"

	"github.com/cucumber/godog"
	"github.com/minio/minio-go/v7"
)

func AddStepSeeBucket(sc *godog.ScenarioContext) {
	sc.Then(`^I should see a bucket "([^"]+)"$`,
		seeBucket,
	)

	return
}

func seeBucket(ctx0 context.Context, bucketName string) (
	ctx context.Context, e error,
) {
	ctx = ctx0

	var (
		bucket  minio.BucketInfo
		buckets []minio.BucketInfo = ctx.Value(ctxKeyBuckets{}).([]minio.BucketInfo)
	)

	for _, bucket = range buckets {
		if bucket.Name == bucketName {
			return
		}
	}

	godog.T(ctx).Fail()

	return
}
