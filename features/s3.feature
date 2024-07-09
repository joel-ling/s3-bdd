Feature: S3
  In order to verify the effects of software interacting with the S3 API
  As a software engineer practising behaviour-driven development
  I need to establish some baseline observations

  Scenario: Create and list buckets
    Given there is a new S3 server running
    When I create a bucket "bucket0"
    And I list all the buckets I own
    Then I should see a bucket "bucket0"
    When I create a bucket "bucket1"
    And I list all the buckets I own
    Then I should see a bucket "bucket0"
    And I should see a bucket "bucket1"

  Scenario: Put and list objects in bucket
    Given there is a new S3 server running
    And there is a bucket "bucket"
    When I put in bucket "bucket" an object "object0" "Hello, "
    And I put in bucket "bucket" an object "object1" "World!"
    And I list the objects in bucket "bucket"
    Then I should see an object "object0" of size 7 bytes
    And I should see an object "object1" of size 6 bytes

  Scenario: List objects with prefix
    Given there is a new S3 server running
    And there is a bucket "bucket"
    And there is in bucket "bucket" an object "A" "Hello, "
    And there is in bucket "bucket" an object "objectB" "World!"
    And there is in bucket "bucket" an object "C" "Hello, "
    And there is in bucket "bucket" an object "objectD" "World!"
    When I list the objects in bucket "bucket" with prefix "object"
    Then I should see an object "objectB" of size 6 bytes
    And I should see an object "objectD" of size 6 bytes

  Scenario: List objects, starting after
    Given there is a new S3 server running
    And there is a bucket "bucket"
    And there is in bucket "bucket" an object "3" "World!"
    And there is in bucket "bucket" an object "1" "World!"
    And there is in bucket "bucket" an object "2" "Hello, "
    And there is in bucket "bucket" an object "0" "Hello, "
    When I list the objects in bucket "bucket" starting after "1"
    Then I should see an object "2" of size 7 bytes
    And I should see an object "3" of size 6 bytes

  Scenario: Get non-existent object
    Given there is a new S3 server running
    And there is a bucket "bucket"
    When I get from bucket "bucket" an object "object"
    Then I should see an error "The specified key does not exist."

  Scenario: Put and get object
    Given there is a new S3 server running
    And there is a bucket "bucket"
    When I put in bucket "bucket" an object "object" "Hello, World!"
    And I get from bucket "bucket" an object "object"
    Then I should see content "Hello, World!"
    And I should see no error
