# lc500

![logo](logo.png)

lc500 is a multi-tenant JavaScript hosting web server. It leverages V8 Isolates to provide isolated environments for each HTTP request, allowing them to be processed individually using JavaScript and return responses accordingly.

## Key Features

- **Multi-tenancy**: Supports multiple isolated JavaScript environments.
- **V8 Isolates**: Uses V8 engine's isolates for secure and efficient JavaScript execution.
- **Per-request Isolation**: Each HTTP request is processed in its own isolated environment.
- **JavaScript Processing**: Allows custom JavaScript code to handle requests and generate responses.
- **Flexible Response Handling**: Enables dynamic generation of HTTP responses based on request data.

## How It Works

1. When an HTTP request is received, lc500 creates a new V8 Isolate.
2. The request data is made available to the JavaScript environment.
3. User-defined JavaScript code processes the request within the isolated environment.
4. The JavaScript code generates a response, which is then sent back to the client.
5. After processing, the isolate is disposed of, ensuring a clean slate for the next request.

This architecture ensures that each request is handled in a secure, isolated manner, preventing cross-request interference and enhancing overall system security and stability.

## Request Processing Flow

```mermaid
sequenceDiagram

    actor User
    participant Gateway
    participant Worker
    participant Script Storage as Script Storage <br>(S3-compatible object storage)
    participant Context Storage as Context Storage <br>(called blueprint here)
    User->>Gateway: Send HTTP request

    par Initialize V8 isolate context and Fetch script and context
        Gateway->>Worker: Initialize V8 isolate context
        Gateway->>Script Storage: Fetch script
        Gateway->>Context Storage: Request context
    end

    Worker->>Gateway: V8 isolate context initialized
    Script Storage->>Gateway: Return script
    Gateway->>Worker: Pass script and ready to run
    Context Storage->>Gateway: Return context
    Gateway->>Worker: Pass context
    Worker->>Worker: Execute script
    Worker->>Gateway: Return execution result
    Gateway->>User: Send HTTP response
```

