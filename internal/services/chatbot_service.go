// package services

// import (
// 	"context"
// 	"log"
// 	"os"

// 	"github.com/Ayush10/tradesage-chatbot-service/internal/pb"
// 	"github.com/Ayush10/tradesage-chatbot-service/internal/utils"
// 	"github.com/sashabaranov/go-openai"
// 	"google.golang.org/grpc/codes"
// 	"google.golang.org/grpc/metadata"
// 	"google.golang.org/grpc/status"
// )

// // ChatbotServiceServer implements the ChatbotService gRPC server.
// type ChatbotServiceServer struct {
// 	pb.UnimplementedChatbotServiceServer
// 	openAIClient       *openai.Client
// 	sessionStore       map[string][]openai.ChatCompletionMessage
// 	marketDataService  *MarketDataService
// 	userProfileService *UserProfileService
// }

// // NewChatbotServiceServer initializes a new ChatbotServiceServer.
// // func NewChatbotServiceServer() (*ChatbotServiceServer, error) {
// // 	apiKey := os.Getenv("OPENAI_API_KEY")
// // 	if apiKey == "" {
// // 		return nil, status.Error(codes.Internal, "OpenAI API key not set")
// // 	}
// // 	client := openai.NewClient(apiKey)

// // 	marketDataService := NewMarketDataService()
// // 	userProfileService := NewUserProfileService()

// // 	return &ChatbotServiceServer{
// // 		openAIClient:       client,
// // 		sessionStore:       make(map[string][]openai.ChatCompletionMessage),
// // 		marketDataService:  marketDataService,
// // 		userProfileService: userProfileService,
// // 	}, nil
// // }

// func NewChatbotServiceServer() (*ChatbotServiceServer, error) {
//     apiKey := os.Getenv("OPENAI_API_KEY")
//     var client *openai.Client
//     if apiKey != "" {
//         client = openai.NewClient(apiKey)
//     } else {
//         log.Println("Warning: OpenAI API key not set. OpenAI functionalities will be disabled.")
//     }

//     marketDataService := NewMarketDataService()
//     userProfileService := NewUserProfileService()

//     return &ChatbotServiceServer{
//         openAIClient:       client,
//         sessionStore:       make(map[string][]openai.ChatCompletionMessage),
//         marketDataService:  marketDataService,
//         userProfileService: userProfileService,
//     }, nil
// }

// // SendMessage processes incoming chat messages.
// func (s *ChatbotServiceServer) SendMessage(ctx context.Context, req *pb.ChatRequest) (*pb.ChatResponse, error) {
// 	// Extract metadata.
// 	md, ok := metadata.FromIncomingContext(ctx)
// 	if !ok {
// 		return nil, status.Error(codes.Unauthenticated, "No metadata found")
// 	}

// 	// Get the authorization token.
// 	tokens := md.Get("authorization")
// 	if len(tokens) == 0 {
// 		return nil, status.Error(codes.Unauthenticated, "No auth token")
// 	}
// 	authToken := tokens[0]

// 	// Validate authToken with the Authentication Service.
// 	userID, err := utils.ValidateAuthToken(authToken)
// 	if err != nil {
// 		return nil, status.Error(codes.Unauthenticated, "Invalid auth token")
// 	}

// 	message := req.GetMessage()

// 	reply, err := s.getChatbotResponse(ctx, userID, message)
// 	if err != nil {
// 		log.Printf("Error getting chatbot response: %v", err)
// 		return nil, status.Error(codes.Internal, "Failed to get chatbot response")
// 	}

// 	return &pb.ChatResponse{Reply: reply}, nil
// }

// // getChatbotResponse communicates with the OpenAI API.
// func (s *ChatbotServiceServer) getChatbotResponse(ctx context.Context, userID, message string) (string, error) {
// 	if s.openAIClient == nil {
//         return "OpenAI functionality is currently disabled.", nil
//     }

// 	// Retrieve session history.
// 	messages := s.sessionStore[userID]

// 	// Append the new user message.
// 	messages = append(messages, openai.ChatCompletionMessage{
// 		Role:    "user",
// 		Content: message,
// 	})

// 	// Check for special commands or requests.
// 	processed, reply, err := s.processSpecialRequests(ctx, userID, message)
// 	if err != nil {
// 		return "", err
// 	}
// 	if processed {
// 		// Update session and return the reply.
// 		messages = append(messages, openai.ChatCompletionMessage{
// 			Role:    "assistant",
// 			Content: reply,
// 		})
// 		s.sessionStore[userID] = messages
// 		return reply, nil
// 	}

// 	// Create the chat completion request.
// 	req := openai.ChatCompletionRequest{
// 		Model:    openai.GPT3Dot5Turbo,
// 		Messages: messages,
// 	}

// 	// Call the OpenAI API.
// 	resp, err := s.openAIClient.CreateChatCompletion(ctx, req)
// 	if err != nil {
// 		return "", err
// 	}

// 	// Get the assistant's reply.
// 	reply = resp.Choices[0].Message.Content

// 	// Append the assistant's reply to the session.
// 	messages = append(messages, openai.ChatCompletionMessage{
// 		Role:    "assistant",
// 		Content: reply,
// 	})

// 	// Update the session store.
// 	s.sessionStore[userID] = messages

// 	return reply, nil
// }

// // processSpecialRequests handles specific user requests like market data.
// func (s *ChatbotServiceServer) processSpecialRequests(ctx context.Context, userID, message string) (bool, string, error) {

// 	// Example: Check if the message is a market data request.
// 	if isMarketDataRequest(message) {
// 		symbol := extractSymbol(message)
// 		if symbol == "" {
// 			return true, "Please provide a valid stock symbol.", nil
// 		}
// 		data, err := s.marketDataService.GetMarketData(ctx, symbol)
// 		if err != nil {
// 			return true, "Sorry, I couldn't retrieve market data at this time.", nil
// 		}
// 		reply := formatMarketDataResponse(data)
// 		return true, reply, nil
// 	}
// 	// Add more conditions for other special requests.
// 	return false, "", nil
// }

package services

import (
	"context"
	"log"
	"os"

	pb "github.com/Ayush10/tradesage-chatbot-service/internal/pb"
	userprofile "github.com/Ayush10/tradesage-chatbot-service/internal/services/user_profile"
	"github.com/Ayush10/tradesage-chatbot-service/internal/utils"
	"github.com/sashabaranov/go-openai"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

// ChatbotServiceServer implements the ChatbotService gRPC server.
type ChatbotServiceServer struct {
	pb.UnimplementedChatbotServiceServer
	openAIClient       *openai.Client
	sessionStore       map[string][]openai.ChatCompletionMessage
	marketDataService  *MarketDataService
	userProfileService *userprofile.Service
	bypassAuth         bool // Flag to bypass authentication
}

// NewChatbotServiceServer initializes a new ChatbotServiceServer.
func NewChatbotServiceServer() (*ChatbotServiceServer, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	var client *openai.Client
	if apiKey != "" {
		client = openai.NewClient(apiKey)
	} else {
		log.Println("Warning: OpenAI API key not set. OpenAI functionalities will be disabled.")
	}

	marketDataService := NewMarketDataService()
	userProfileService := userprofile.NewService()

	// Check if BYPASS_AUTH environment variable is set
	bypassAuth := os.Getenv("BYPASS_AUTH") == "true"

	return &ChatbotServiceServer{
		openAIClient:       client,
		sessionStore:       make(map[string][]openai.ChatCompletionMessage),
		marketDataService:  marketDataService,
		userProfileService: userProfileService,
		bypassAuth:         bypassAuth,
	}, nil
}

// SendMessage processes incoming chat messages.
func (s *ChatbotServiceServer) SendMessage(ctx context.Context, req *pb.ChatRequest) (*pb.ChatResponse, error) {
	var userID string
	if !s.bypassAuth {
		// Extract metadata.
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "No metadata found")
		}

		// Get the authorization token.
		tokens := md.Get("authorization")
		if len(tokens) == 0 {
			return nil, status.Error(codes.Unauthenticated, "No auth token")
		}
		authToken := tokens[0]

		// Validate authToken with the Authentication Service.
		var err error
		userID, err = utils.ValidateAuthToken(authToken)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "Invalid auth token")
		}
	} else {
		log.Println("Authentication bypassed for testing purposes.")
		userID = "test-user" // Assign a default user ID for testing
	}

	message := req.GetMessage()

	reply, err := s.getChatbotResponse(ctx, userID, message)
	if err != nil {
		log.Printf("Error getting chatbot response: %v", err)
		return nil, status.Error(codes.Internal, "Failed to get chatbot response")
	}

	return &pb.ChatResponse{Reply: reply}, nil
}

// getChatbotResponse communicates with the OpenAI API.
func (s *ChatbotServiceServer) getChatbotResponse(ctx context.Context, userID, message string) (string, error) {
	if s.openAIClient == nil {
		return "OpenAI functionality is currently disabled.", nil
	}

	// Retrieve session history.
	messages := s.sessionStore[userID]

	// Append the new user message.
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "user",
		Content: message,
	})

	// Check for special commands or requests.
	processed, reply, err := s.processSpecialRequests(ctx, userID, message)
	if err != nil {
		return "", err
	}
	if processed {
		// Update session and return the reply.
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    "assistant",
			Content: reply,
		})
		s.sessionStore[userID] = messages
		return reply, nil
	}

	// Create the chat completion request.
	req := openai.ChatCompletionRequest{
		Model:    openai.GPT3Dot5Turbo,
		Messages: messages,
	}

	// Call the OpenAI API.
	resp, err := s.openAIClient.CreateChatCompletion(ctx, req)
	if err != nil {
		return "", err
	}

	// Get the assistant's reply.
	reply = resp.Choices[0].Message.Content

	// Append the assistant's reply to the session.
	messages = append(messages, openai.ChatCompletionMessage{
		Role:    "assistant",
		Content: reply,
	})

	// Update the session store.
	s.sessionStore[userID] = messages

	return reply, nil
}

// processSpecialRequests handles specific user requests like market data.
func (s *ChatbotServiceServer) processSpecialRequests(ctx context.Context, userID, message string) (bool, string, error) {
	// Example: Check if the message is a market data request.
	if isMarketDataRequest(message) {
		symbol := extractSymbol(message)
		if symbol == "" {
			return true, "Please provide a valid stock symbol.", nil
		}
		data, err := s.marketDataService.GetMarketData(ctx, symbol)
		if err != nil {
			return true, "Sorry, I couldn't retrieve market data at this time.", nil
		}
		reply := formatMarketDataResponse(data)
		return true, reply, nil
	}
	// Add more conditions for other special requests.
	return false, "", nil
}
