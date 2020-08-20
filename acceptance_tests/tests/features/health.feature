Feature: Test the healthcheck

	Scenario: Happy path healtcheck
		Given Redis is available
		When I make a "GET" request to "/health"
		Then I receive a status code of "200"
		And the response body field ".ready" boolean value equals "true"

	Scenario: Redis unavailable
		Given Redis is unavailable
		When I make a "GET" request to "/health"
		Then I receive a status code of "503"
		And the response body field ".ready" boolean value equals "true"