# TradeSage AI - Chatbot Integration Service

## Overview

The **TradeSage AI - Chatbot Integration Service** is a microservice designed to provide an AI-powered trading assistant for the TradeSage AI platform. It leverages OpenAI's GPT-3.5 Turbo model to deliver conversational responses, market analysis, and personalized trading advice. The service integrates with real-time market data and user profiles to enhance user experience and trading efficiency.

## Features

- **AI Chatbot**: Provides intelligent, conversational AI responses to user queries.
- **Real-Time Market Data Integration**: Fetches up-to-date stock prices and market insights.
- **User Authentication**: Validates JWT tokens with the existing Authentication Service.
- **Session Management**: Maintains conversation context for personalized interactions.
- **Microservices Architecture**: Designed for seamless integration with other services and frontend applications.

## Prerequisites

- **Go** version 1.18 or higher
- **Protobuf Compiler (`protoc`)** with Go plugins
- **OpenAI API Key**
- **Market Data API Key** (e.g., from Alpha Vantage)
- **Authentication Service** (completed and running, using JWT in Go)
- **Frontend Application** (e.g., Next.js) ready for integration

## Setup Instructions

### 1. Clone the Repository

```bash
git clone https://github.com/yourusername/tradesage-chatbot-service.git
cd tradesage-chatbot-service
```

# TradeSage Chatbot Service Setup Guide

## Environment Configuration

### Setting Up Environment Variables

1. Create a `.env` file in the root directory:

```bash
touch .env
```

2. Add the following content to `.env`:

```dotenv
OPENAI_API_KEY=your_openai_api_key
MARKET_DATA_API_KEY=your_market_data_api_key
CHATBOT_SERVICE_PORT=50051
AUTH_SERVICE_URL=http://localhost:5000/validate-token
MARKET_DATA_API_URL=https://www.alphavantage.co/query
```

> **Note**: Replace `your_openai_api_key` and `your_market_data_api_key` with your actual API keys. Ensure that `AUTH_SERVICE_URL` points to your Authentication Service's token validation endpoint.

## Installation and Setup

### Installing Dependencies

```bash
go mod tidy
```

### Generating Protobuf Files

Ensure you have `protoc` and the Go plugins installed, then generate the protobuf files:

```bash
protoc --go_out=internal/pb --go-grpc_out=internal/pb internal/proto/chatbot.proto
```

### Running the Service

```bash
go run cmd/tradesage-chatbot-service/main.go
```

## Service Operation Guide

### Running Requirements

- Ensure all environment variables are correctly set
- Start the service using the command above
- Service will run on the port specified in `CHATBOT_SERVICE_PORT` (default: 50051)
- Monitor logs for startup errors

## Testing

### Using grpcurl

```bash
grpcurl -plaintext \
  -d '{"message": "What is the price of AAPL?"}' \
  -H 'authorization: Bearer your_jwt_token' \
  localhost:50051 \
  chatbot.ChatbotService/SendMessage
```

Expected Response:

```json
{
  "reply": "The current price of AAPL is $123.45."
}
```

> **Note**: The actual price will vary based on real-time data.

## Integration Guide

### Authentication Service Integration

The Chatbot Service validates JWT tokens through the Authentication Service. Ensure `AUTH_SERVICE_URL` points to the correct endpoint. Token validation is handled by the `auth.go` utility.

### Market Data Service Integration

Real-time market data is fetched using the configured API key and URL from the `.env` file. The service can be adapted for different market data providers by updating `MARKET_DATA_API_URL` and modifying the MarketDataService implementation.

## Frontend Integration Guide

### 1. Generate gRPC Client Code

Install protoc-gen-grpc-web:

```bash
npm install -g protoc-gen-grpc-web
```

Generate client code:

```bash
protoc -I=./internal/proto \
  --js_out=import_style=commonjs:./frontend/src/generated \
  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./frontend/src/generated \
  internal/proto/chatbot.proto
```

### 2. gRPC-Web Proxy Setup

Set up a proxy (like Envoy) or use grpc-web library for browser communication with the gRPC service.

### 3. Frontend Implementation

Example TypeScript (React) implementation:

```typescript
import { ChatbotServiceClient } from "./generated/chatbot_pb_service";
import { ChatRequest } from "./generated/chatbot_pb";
import { grpc } from "@improbable-eng/grpc-web";

const client = new ChatbotServiceClient("http://localhost:8080"); // Proxy URL

const metadata = new grpc.Metadata();
metadata.set("authorization", "Bearer your_jwt_token");

const request = new ChatRequest();
request.setMessage("Tell me about TSLA.");

client.sendMessage(request, metadata, (err, response) => {
  if (err) {
    console.error("Error:", err.message);
  } else {
    console.log("Reply:", response.getReply());
  }
});
```

### 4. Session Management

- Implement proper session state management
- Handle reconnections and errors gracefully

### 5. Deployment and Testing

- Deploy the frontend and verify communication with the Chatbot Service
- Test the complete interaction flow

## Security Considerations

- **JWT Validation**: Implement secure token validation through the Authentication Service
- **API Keys**: Keep `.env` file and sensitive information out of version control
- **Secure Communication**: Use TLS for all service communication in production
- **Input Sanitization**: Implement proper sanitization for all user inputs
