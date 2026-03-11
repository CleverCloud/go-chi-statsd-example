# Go Chi + StatsD Example on Clever Cloud

[![Clever Cloud - PaaS](https://img.shields.io/badge/Clever%20Cloud-PaaS-orange)](https://clever-cloud.com)

This is a simple Go application that demonstrates how to deploy a web server with [Chi](https://github.com/go-chi/chi) router and [StatsD](https://github.com/statsd/statsd) metrics to Clever Cloud.

This is a modern replacement for the [goji-statsd-example](https://github.com/CleverCloud/demo-go-goji-statsd), using Chi which is the idiomatic, lightweight router for Go.

For a complete guide on deploying Go applications, see the [Clever Cloud Go documentation](https://www.clever-cloud.com/developers/doc/applications/golang/).

## About the Application

This application uses Chi router with request logging middleware and reports metrics to a StatsD server. It listens on `0.0.0.0:8080`.

On Clever Cloud, a StatsD server is available at `127.0.0.1:8125` on every application instance — no addon required. Each request to `/hello/{name}` increments a `hello.request.{name}` counter. Once deployed, your custom metrics are accessible through Warp 10 under the `statsd.*` prefix. Learn more about [publishing your own metrics](https://www.clever.cloud/developers/doc/metrics/#publish-your-own-metrics).

### Routes

| Route | Description |
|-------|-------------|
| `GET /hello/{name}` | Returns a greeting with the given name and sends a StatsD metric |

## Technology Stack

- [Go](https://go.dev/) 1.21+ with [Chi](https://github.com/go-chi/chi) v5 router
- [go-statsd-client](https://github.com/cactus/go-statsd-client) v5 for StatsD metrics

## Prerequisites

- Go 1.21+

## Running the Application Locally

```bash
go run .
```

The application will be accessible at http://localhost:8080/hello/world.

## Deploying on Clever Cloud

You have two options to deploy your application on Clever Cloud: using the Web Console or using the Clever Tools CLI.

### Option 1: Deploy using the Web Console

#### 1. Create an account on Clever Cloud

If you don't already have an account, go to the [Clever Cloud console](https://console.clever-cloud.com/) and follow the registration instructions.

#### 2. Set up your application on Clever Cloud

1. Log in to the [Clever Cloud console](https://console.clever-cloud.com/)
2. Click on "Create" and select "An application"
3. Choose "Go" as the runtime environment
4. Configure your application settings (name, region, etc.)

#### 3. Configure Environment Variables

Add the following environment variables in the Clever Cloud console:

| Variable | Value | Description |
|----------|-------|-------------|
| `CC_GO_BUILD_TOOL` | `gomod` | Use Go Modules for building (reads `go.mod`) |

#### 4. Deploy Your Application

You can deploy your application using Git:

```bash
# Add Clever Cloud as a remote repository
git remote add clever git+ssh://git@push-par-clevercloud-customers.services.clever-cloud.com/app_<your-app-id>.git

# Push your code to deploy
git push clever master
```

### Option 2: Deploy using Clever Tools CLI

#### 1. Install Clever Tools

Install the Clever Tools CLI following the [official documentation](https://www.clever-cloud.com/doc/clever-tools/getting_started/):

```bash
# Using npm
npm install -g clever-tools

# Or using Homebrew (macOS)
brew install clever-tools
```

#### 2. Log in to your Clever Cloud account

```bash
clever login
```

#### 3. Create a new application

```bash
# Initialize the current directory as a Clever Cloud application
clever create --type go <YOUR_APP_NAME>

# Set the required environment variables
clever env set CC_GO_BUILD_TOOL gomod
```

#### 4. Deploy your application

```bash
clever deploy
```

#### 5. Open your application in a browser

Once deployed, you can access your application at the provided domain.

### Monitoring Your Application

Once deployed, you can monitor your application through:

- **Web Console**: The Clever Cloud console provides logs, metrics, and other tools to help you manage your application.
- **CLI**: Use `clever logs` to view application logs and `clever status` to check the status of your application.

## Additional Resources

- [Go Documentation](https://go.dev/doc/)
- [Chi Documentation](https://github.com/go-chi/chi)
- [Clever Cloud Documentation](https://www.clever-cloud.com/doc/)
- [Clever Cloud Go Deployment](https://www.clever-cloud.com/developers/doc/applications/golang/)
- [Clever Cloud Metrics & Custom Metrics](https://www.clever.cloud/developers/doc/metrics/)
