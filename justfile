# Run the Go tests in watch mode
test:
    cd code && gotestsum --watch ./...

# Start the Slidev presentation
present:
    cd presentation && npm start
