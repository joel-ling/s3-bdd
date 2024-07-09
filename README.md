```bash
$ go test
```
```gherkin
Feature: S3
  In order to verify the effects of software interacting with the S3 API
  As a software engineer practising behaviour-driven development
  I need to establish some baseline observations

  Scenario: Create and list buckets        # features/s3.feature:6
    Given there is a new S3 server running # new-s3-server.go:20 -> github.com/joel-ling/s3-bdd/test.newS3Server
    When I create a bucket "bucket0"       # create-bucket.go:22 -> github.com/joel-ling/s3-bdd/test.createBucket
    And I list all the buckets I own       # list-buckets.go:18 -> github.com/joel-ling/s3-bdd/test.listBuckets
    Then I should see a bucket "bucket0"   # see-bucket.go:18 -> github.com/joel-ling/s3-bdd/test.seeBucket
    When I create a bucket "bucket1"       # create-bucket.go:22 -> github.com/joel-ling/s3-bdd/test.createBucket
    And I list all the buckets I own       # list-buckets.go:18 -> github.com/joel-ling/s3-bdd/test.listBuckets
    Then I should see a bucket "bucket0"   # see-bucket.go:18 -> github.com/joel-ling/s3-bdd/test.seeBucket
    And I should see a bucket "bucket1"    # see-bucket.go:18 -> github.com/joel-ling/s3-bdd/test.seeBucket

  Scenario: Put and list objects in bucket                      # features/s3.feature:16
    Given there is a new S3 server running                      # new-s3-server.go:20 -> github.com/joel-ling/s3-bdd/test.newS3Server
    And there is a bucket "bucket"                              # create-bucket.go:22 -> github.com/joel-ling/s3-bdd/test.createBucket
    When I put in bucket "bucket" an object "object0" "Hello, " # put-object.go:23 -> github.com/joel-ling/s3-bdd/test.putObject
    And I put in bucket "bucket" an object "object1" "World!"   # put-object.go:23 -> github.com/joel-ling/s3-bdd/test.putObject
    And I list the objects in bucket "bucket"                   # list-objects.go:18 -> github.com/joel-ling/s3-bdd/test.listObjects
    Then I should see an object "object0" of size 7 bytes       # see-object.go:19 -> github.com/joel-ling/s3-bdd/test.seeObject
    And I should see an object "object1" of size 6 bytes        # see-object.go:19 -> github.com/joel-ling/s3-bdd/test.seeObject

  Scenario: List objects with prefix                                # features/s3.feature:25
    Given there is a new S3 server running                          # new-s3-server.go:20 -> github.com/joel-ling/s3-bdd/test.newS3Server
    And there is a bucket "bucket"                                  # create-bucket.go:22 -> github.com/joel-ling/s3-bdd/test.createBucket
    And there is in bucket "bucket" an object "A" "Hello, "         # put-object.go:23 -> github.com/joel-ling/s3-bdd/test.putObject
    And there is in bucket "bucket" an object "objectB" "World!"    # put-object.go:23 -> github.com/joel-ling/s3-bdd/test.putObject
    And there is in bucket "bucket" an object "C" "Hello, "         # put-object.go:23 -> github.com/joel-ling/s3-bdd/test.putObject
    And there is in bucket "bucket" an object "objectD" "World!"    # put-object.go:23 -> github.com/joel-ling/s3-bdd/test.putObject
    When I list the objects in bucket "bucket" with prefix "object" # list-objects-prefix.go:18 -> github.com/joel-ling/s3-bdd/test.listObjectsPrefix
    Then I should see an object "objectB" of size 6 bytes           # see-object.go:19 -> github.com/joel-ling/s3-bdd/test.seeObject
    And I should see an object "objectD" of size 6 bytes            # see-object.go:19 -> github.com/joel-ling/s3-bdd/test.seeObject

  Scenario: List objects, starting after                          # features/s3.feature:36
    Given there is a new S3 server running                        # new-s3-server.go:20 -> github.com/joel-ling/s3-bdd/test.newS3Server
    And there is a bucket "bucket"                                # create-bucket.go:22 -> github.com/joel-ling/s3-bdd/test.createBucket
    And there is in bucket "bucket" an object "3" "World!"        # put-object.go:23 -> github.com/joel-ling/s3-bdd/test.putObject
    And there is in bucket "bucket" an object "1" "World!"        # put-object.go:23 -> github.com/joel-ling/s3-bdd/test.putObject
    And there is in bucket "bucket" an object "2" "Hello, "       # put-object.go:23 -> github.com/joel-ling/s3-bdd/test.putObject
    And there is in bucket "bucket" an object "0" "Hello, "       # put-object.go:23 -> github.com/joel-ling/s3-bdd/test.putObject
    When I list the objects in bucket "bucket" starting after "1" # list-objects-after.go:19 -> github.com/joel-ling/s3-bdd/test.listObjectsAfter
    Then I should see an object "2" of size 7 bytes               # see-object.go:19 -> github.com/joel-ling/s3-bdd/test.seeObject
    And I should see an object "3" of size 6 bytes                # see-object.go:19 -> github.com/joel-ling/s3-bdd/test.seeObject

  Scenario: Get non-existent object                                # features/s3.feature:47
    Given there is a new S3 server running                         # new-s3-server.go:20 -> github.com/joel-ling/s3-bdd/test.newS3Server
    And there is a bucket "bucket"                                 # create-bucket.go:22 -> github.com/joel-ling/s3-bdd/test.createBucket
    When I get from bucket "bucket" an object "object"             # get-object.go:19 -> github.com/joel-ling/s3-bdd/test.getObject
    Then I should see an error "The specified key does not exist." # see-error.go:18 -> github.com/joel-ling/s3-bdd/test.seeError

  Scenario: Put and get object                                       # features/s3.feature:53
    Given there is a new S3 server running                           # new-s3-server.go:20 -> github.com/joel-ling/s3-bdd/test.newS3Server
    And there is a bucket "bucket"                                   # create-bucket.go:22 -> github.com/joel-ling/s3-bdd/test.createBucket
    When I put in bucket "bucket" an object "object" "Hello, World!" # put-object.go:23 -> github.com/joel-ling/s3-bdd/test.putObject
    And I get from bucket "bucket" an object "object"                # get-object.go:19 -> github.com/joel-ling/s3-bdd/test.getObject
    Then I should see content "Hello, World!"                        # see-content.go:18 -> github.com/joel-ling/s3-bdd/test.seeContent
    And I should see no error                                        # see-no-error.go:18 -> github.com/joel-ling/s3-bdd/test.seeNoError
```
```txt
6 scenarios (6 passed)
43 steps (43 passed)
2.969250143s
PASS
ok  	github.com/joel-ling/s3-bdd	2.976s
```
