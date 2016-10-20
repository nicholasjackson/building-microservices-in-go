@search
Feature: As a user when I call the search endpoint, I would like to receive a list of kittens

  Scenario: Invalid query
    Given I have no search criteria
    When I call the search endpoint
    Then I should receive a bad request message

  Scenario: Valid query
    Given I have a valid search criteria
    When I call the search endpoint
    Then I should receive a list of kittens
