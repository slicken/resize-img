# Image Resize API

A high-performance image resizing service built with Go, leveraging the standard library's image processing capabilities.

## Features

- Resize images while maintaining aspect ratio
- Support for multiple image formats (JPEG, PNG, GIF)
- RESTful API endpoints for image processing
- Concurrent image processing
- Input validation and error handling

## Prerequisites

- Go 1.23.5 or newer
- Git

## Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/resize-img.git
cd resize-img
```

2. Install dependencies:
```bash
go mod tidy
```

## Usage

1. Start the server:
```bash
go run main.go
```

The server will start on `http://localhost:8080` by default.

## API Endpoints

### Resize Image
```
POST /api/resize
```

**Request Body (multipart/form-data):**
- `image`: The image file to resize
- `width`: Desired width in pixels
- `height`: Desired height in pixels

**Response:**
- Success: 200 OK with resized image
- Error: Appropriate error status code with JSON error message

## Error Handling

The API returns appropriate HTTP status codes and error messages:

- 400: Bad Request - Invalid input parameters
- 415: Unsupported Media Type - Unsupported image format
- 500: Internal Server Error - Processing error

## Development

### Project Structure

```
resize-img/
├── main.go           # Entry point
├── handlers/         # HTTP handlers
├── services/        # Business logic
└── utils/           # Helper functions
```

### Running Tests

```bash
go test ./...
```

## License

MIT License - see LICENSE file for details

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request 