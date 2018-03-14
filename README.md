# glowing-gopher-cli

A cli tool to get a table view of the health of all your services, add your glowing-gopher endpoint and build to see the output.

```
# Environment variables
HEALTHCHECK_ENDPOINT
```

Run the following with your own healthcheck endpoint to get started

```
go build -ldflags "-X http://127.0.0.1:3333/healthcheck=$HEALTHCHECK_ENDPOINT"
go install
glowing-gopher-cli
```
