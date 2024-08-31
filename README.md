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
