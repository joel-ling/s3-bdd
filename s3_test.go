package s3bdd

import (
	"testing"

	"github.com/cucumber/godog"

	"github.com/joel-ling/s3-bdd/test"
)

func TestS3(t *testing.T) {
	var (
		scenarioInitializer = func(sc *godog.ScenarioContext) {
			test.AddStepSetUp(sc)
			test.AddStepNewS3Server(sc)
			test.AddStepCreateBucket(sc)
			test.AddStepListBuckets(sc)
			test.AddStepPutObject(sc)
			test.AddStepListObjects(sc)
			test.AddStepListObjectsPrefix(sc)
			test.AddStepListObjectsAfter(sc)
			test.AddStepGetObject(sc)
			test.AddStepSeeBucket(sc)
			test.AddStepSeeObject(sc)
			test.AddStepSeeContent(sc)
			test.AddStepSeeError(sc)
			test.AddStepSeeNoError(sc)
			test.AddStepCleanUp(sc)

			return
		}

		options = &godog.Options{
			Format:   "pretty",
			Paths:    []string{"features"},
			TestingT: t,
		}

		suite = godog.TestSuite{
			ScenarioInitializer: scenarioInitializer,
			Options:             options,
		}
	)

	if suite.Run() != 0 {
		t.Fatal()
	}

	return
}
