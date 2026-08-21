# Gateway

## Setup

1. **pre-commit hook**
    ```cmd
    git config core.hooksPath .githooks
    ```

## Tools

1.  **swaggo** - Swagger and OpenAPI documentation generator
    
    ```cmd 
    go install github.com/swaggo/swag/cmd/swag@v1.16.6
    swag init -g ./cmd/main.go
    ```
