Rate Limiter Exercise - Auth0
Instructions
We provide you with a basic "Hello World" api boilerplate in a golang 1.17 environment (), as well as a basic integration test that calls to that existing endpoint. To support the testing we use stretchr testify.
The boilerplate is meant to accelerate you but you are free to change it as you see fit.
This exercise represents a MVP / POC feature, that we may expand upon further in the following interview stages.
There are no hidden tests or code that will automatically score your exercise. It will be manually reviewed and graded by our team's engineers, with expectations being adjusted based to your seniority.
Implement your solution as if you were developing code to be deployed and maintained by you and your team as best as you can given the limited time frame, the sandbox environment limitations and the requirements asked, and document any compromises, gaps or possible follow-ups in a README.md file.
Exercise
To protect our customers and our APIs, we use a rate limits mechanism.

Your task today will be that of implementing a small standalone service to provide that capability to our other services that clients directly use.

Token bucket algorithm
The way we express rate limiting is in the form of Burst limit (higher) and sustained limit (lower), for instance you could have a rate limit defined for the route template GET users/:id of 100 req burst and 50 req/min sustained.

This means you can accumulate requests' "credit", or tokens, up to a maximum of 100, but they only refresh at a rate of a total of 50 per minute. Every time you check to allow a request for a certain route template, you deduct one token. If you run out of tokens the request should be rejected, until tokens replenish at the defined sustained rate.

An algorithm that seems fit for this approach is what's called Token Bucket. Feel free to use it in your solution or not, as long as the requirements presented below are met.

Requirements
Please implement a small service that:

Loads the configuration represented in config.json every time it starts. In there you will find an array at rateLimitsPerEndpoint with a configuration for each route template the service should provide a rate limit for:
endpoint: route template being limited, acts as a key provided by the caller to check and consume its request tokens
burst: the number of burst requests allowed
sustained: the number of sustained requested per minute.
Exposes a single endpoint at /take/, which receives the route template to check the rate limit for from the caller via query string, route params or body.
This service is called by other client facing services (i.e. as part of their middleware pipeline) to determine if a given request should be accepted or not for the specified route template:
If the limit for the specified route is not exceeded, the endpoint should consume one token and return a response with the number of remaining tokens and confirmation the request should be accepted.
If the limit for the specified route has already been met (empty bucket), the endpoint’s response should have a result of 0 remaining tokens and inform the consumer the request should be rejected.
The bucket for each endpoint starts containing a burst number of tokens, and refills with a sustained number of tokens per minute up to a max of burst number of tokens.
The refilling of tokens should happen as soon as possible, i.e. a sustained rate of 120 req/min should make available 1 token every 0.5 seconds.
Assumptions
The configuration json will always fit in memory.
Assume this service will run in a single instance.
Service state can be kept in memory only (can be lost if service crashes, etc).
The service runs in our internal network, with no exposure to the internet, and without any need to identify or authenticate the caller.
Good luck!

[execution time limit] 30 seconds

[memory limit] 1 GB
