Feature: Greeter

  @story-1
  Rule: Calling GET /hello with a name returns a greeting personalized to that name

    Scenario: A named caller gets a personalized greeting
      Given the greeter service is running
      When an API Consumer calls GET /hello with name "Ada"
      Then the response is a JSON greeting with message "Hello, Ada!"

  @story-2
  Rule: Calling GET /hello without a name still succeeds with a default greeting

    Scenario: Omitting the name parameter still returns a greeting
      Given the greeter service is running
      When an API Consumer calls GET /hello without a name parameter
      Then the response is a JSON greeting with message "Hello, World!"

    Scenario: An empty name parameter still returns a greeting
      Given the greeter service is running
      When an API Consumer calls GET /hello with name ""
      Then the response is a JSON greeting with message "Hello, World!"

  @story-1 @negative
  Rule: A name longer than 100 characters is rejected

    Scenario: An overly long name is refused
      Given the greeter service is running
      When an API Consumer calls GET /hello with a name that is 101 characters long
      Then the response is a 400 error and no greeting is returned
