package controller

import (
	"context"

	"github.com/helixml/helix/api/pkg/data"
	"github.com/helixml/helix/api/pkg/types"

	openai "github.com/lukemarsden/go-openai2"
	"github.com/rs/zerolog/log"
)

type ChatCompletionOptions struct {
	AppID       string
	AssistantID string
	RAGSourceID string
}

// ChatCompletion is used by the OpenAI compatible API. Doesn't handle any historical sessions, etc.
// Runs the OpenAI with tools/app configuration and returns the response.
func (c *Controller) ChatCompletion(ctx context.Context, user *types.User, req openai.ChatCompletionRequest, opts *ChatCompletionOptions) (openai.ChatCompletionResponse, error) {

	// TODO: setup app settings
	// TODO: setup RAG prompt if source set

	resp, err := c.openAIClient.CreateChatCompletion(ctx, req)
	if err != nil {
		log.Err(err).Msg("error creating chat completion")
		return openai.ChatCompletionResponse{}, err
	}

	return resp, nil
}

// ChatCompletion is used by the OpenAI compatible API. Doesn't handle any historical sessions, etc.
// Runs the OpenAI with tools/app configuration and returns the stream.
func (c *Controller) ChatCompletionStream(ctx context.Context, user *types.User, req openai.ChatCompletionRequest, opts *ChatCompletionOptions) (*openai.ChatCompletionStream, error) {

	// TODO: setup app settings
	// TODO: setup RAG prompt if source set

	stream, err := c.openAIClient.CreateChatCompletionStream(ctx, req)
	if err != nil {
		log.Err(err).Msg("error creating chat completion stream")
		return nil, err
	}

	return stream, nil
}

// SaveChatCompletion used to persist the chat completion response to the database as a session.
func (c *Controller) SaveChatCompletion(ctx context.Context, user *types.User, req openai.ChatCompletionRequest, resp openai.ChatCompletionResponse) error {
	return nil
}

func sessionToChatCompletion(session *types.Session) (*openai.ChatCompletionRequest, error) {
	var messages []openai.ChatCompletionMessage

	// Adjust length
	var interactions []*types.Interaction
	if len(session.Interactions) > 10 {
		first, err := data.GetFirstUserInteraction(session.Interactions)
		if err != nil {
			log.Err(err).Msg("error getting first user interaction")
		} else {
			interactions = append(interactions, first)
			interactions = append(interactions, data.GetLastInteractions(session, 10)...)
		}
	} else {
		interactions = session.Interactions
	}

	// Adding the system prompt first
	if session.Metadata.SystemPrompt != "" {
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: session.Metadata.SystemPrompt,
		})
	}

	for _, interaction := range interactions {
		switch interaction.Creator {

		case types.CreatorTypeUser:
			messages = append(messages, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleUser,
				Content: interaction.Message,
			})
		case types.CreatorTypeSystem:
			messages = append(messages, openai.ChatCompletionMessage{
				Role:    openai.ChatMessageRoleAssistant,
				Content: interaction.Message,
			})
		case types.CreatorTypeTool:
			messages = append(messages, openai.ChatCompletionMessage{
				Role:       openai.ChatMessageRoleUser,
				Content:    interaction.Message,
				ToolCalls:  interaction.ToolCalls,
				ToolCallID: interaction.ToolCallID,
			})
		}
	}

	var (
		responseFormat *openai.ChatCompletionResponseFormat
		tools          []openai.Tool
		toolChoice     any
	)

	// If the last interaction has response format, use it
	last, _ := data.GetLastSystemInteraction(interactions)
	if last != nil && last.ResponseFormat.Type == types.ResponseFormatTypeJSONObject {
		responseFormat = &openai.ChatCompletionResponseFormat{
			Type:   openai.ChatCompletionResponseFormatTypeJSONObject,
			Schema: last.ResponseFormat.Schema,
		}
	}

	if last != nil && len(last.Tools) > 0 {
		tools = last.Tools
		toolChoice = last.ToolChoice
	}

	// TODO: temperature, etc.

	return &openai.ChatCompletionRequest{
		Model:          string(session.ModelName),
		Stream:         session.Metadata.Stream,
		Messages:       messages,
		ResponseFormat: responseFormat,
		Tools:          tools,
		ToolChoice:     toolChoice,
	}, nil
}
