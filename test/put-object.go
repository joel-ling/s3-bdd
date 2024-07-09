package test

import (
	"bytes"
	"context"

	"github.com/cucumber/godog"
	"github.com/minio/minio-go/v7"
)

func AddStepPutObject(sc *godog.ScenarioContext) {
	sc.Given(`^there is in bucket "([^"]+)" an object "([^"]+)" "([^"]+)"$`,
		putObject,
	)

	sc.When(`^I put in bucket "([^"]+)" an object "([^"]+)" "([^"]+)"$`,
		putObject,
	)

	return
}

func putObject(ctx0 context.Context, bucketName, objectName, content string) (
	ctx context.Context, e error,
) {
	ctx = ctx0

	const (
		multipartPutPartSize = 5 << 20 // 5 MiB
	)

	var (
		buffer               = []byte(content)
		client *minio.Client = ctx.Value(ctxKeyMClient{}).(*minio.Client)
	)

	_, e = client.PutObject(ctx, bucketName, objectName,
		bytes.NewReader(buffer),
		-1,
		minio.PutObjectOptions{
			PartSize: multipartPutPartSize,
		},
	)
	if e != nil {
		return
	}

	return
}
