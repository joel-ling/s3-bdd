package test

import (
	"context"
	"os/exec"

	"github.com/cucumber/godog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func AddStepNewS3Server(sc *godog.ScenarioContext) {
	sc.Given(`^there is a new S3 server running$`,
		newS3Server,
	)

	return
}

func newS3Server(ctx0 context.Context) (ctx context.Context, e error) {
	ctx = ctx0

	const (
		address     = "127.237.93.83:9383"
		defaultCred = "minioadmin"
	)

	var (
		client *minio.Client

		clientOptions = &minio.Options{
			Creds: credentials.NewStaticV4(
				defaultCred,
				defaultCred,
				"",
			),
		}

		server *exec.Cmd = exec.Command(
			ctx.Value(ctxKeyBinPath{}).(string),
			"server",
			"--address", address,
			ctx.Value(ctxKeyTempDir{}).(string),
		)
	)

	e = server.Start()
	if e != nil {
		return
	}

	client, e = minio.New(address, clientOptions)
	if e != nil {
		return
	}

	_, e = client.ListBuckets(
		context.Background(),
	)
	if e != nil {
		return
	}

	ctx = context.WithValue(ctx, ctxKeyMClient{}, client)

	ctx = context.WithValue(ctx, ctxKeySrvProc{}, server.Process)

	return
}
