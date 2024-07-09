package test

import (
	"context"
	"io"

	"github.com/cucumber/godog"
	"github.com/minio/minio-go/v7"
)

func AddStepGetObject(sc *godog.ScenarioContext) {
	sc.When(`^I get from bucket "([^"]+)" an object "([^"]+)"$`,
		getObject,
	)

	return
}

func getObject(ctx0 context.Context, bucketName, objectName string) (
	ctx context.Context, e error,
) {
	ctx = ctx0

	var (
		client *minio.Client = ctx.Value(ctxKeyMClient{}).(*minio.Client)
		object *minio.Object

		content []byte
	)

	object, e = client.GetObject(ctx, bucketName, objectName,
		minio.GetObjectOptions{},
	)
	if e != nil {
		return
	}
	// *minio.Client.GetObject returns a nil error even if object is not found.
	// Errors are only encountered later, when object data is actually read.
	//
	// https://github.com/minio/minio-go/blob/
	//   - 60eddd782a85d6b96b9a846f179b8116d4a6a51c/api-get-object.go#L242
	//   - 60eddd782a85d6b96b9a846f179b8116d4a6a51c/api-get-object.go#L61-L239

	content, e = io.ReadAll(object)
	if e != nil {
		ctx = context.WithValue(ctx, ctxKeyS3Error{}, e)

		e = nil

		return
	}

	ctx = context.WithValue(ctx, ctxKeyContent{}, content)

	return
}
