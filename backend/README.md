# The backend api for [skyzar](https://skyzar.app/)

## Setup
* Run `cp .env.example .env` into this directory
* Configure your [Hypixel API Key](https://developer.hypixel.net/), [MongoDB URI](https://mongodb.com/), [Cloudflare Turnstile Secret](https://www.cloudflare.com/products/turnstile/), and an API Token.
* Run `go mod download` to install the dependencies
* To run the worker, either build the exectuable via `go build` or `go run .`