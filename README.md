# GatewayService

## Tools

### 1. **go-task/task** - build tool inspired by Make *https://github.com/go-task/task*
### Commands
- **Install**
    
    It already have a command in `go generate` but you can install it manually by
    ```cmd
    go install github.com/go-task/task/v3/cmd/task@latest
    ```

- **Init default Taskfile.yml**
    ```bash
    task --init
    ```

- **Run task**
    ```cmd
    task {task_name} 
    ```

### Supported tasks
- `init` - load dependencies
- `init` - install external tools
- `run` - runs code
- `build` - builds project
- `swag-fmt` - `swag fmt`
- `swag-init` - `swag init ...`
- `swag-fmt-d` - calls `fmt` in swag docker container
- `swag-init-d` - calls `init ...` in swag docker container

### 2. **swaggo/swag** - Swagger and OpenAPI documentation generator *https://github.com/swaggo/swag*
### Commands
- **Install**
    ```bash 
    go install github.com/swaggo/swag/cmd/swag@v1.16.6
    ```
    or using docker
    ```bash
    docker run -v ${PWD}:/code ghcr.io/swaggo/swag:v1.16.6 {command_line}
    ```

- **Format comments**
    ```bash
    swag fmt
    ```

- **Generate swagger documentation**
    ```bash
    swag init -d ./cmd/,./ -o ./gen/docs
    ```

