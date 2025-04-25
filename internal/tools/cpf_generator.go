package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/renanmedina/4devs-mcp/internal/services"
)

type CPFGeneratorTool struct {
	toolName        string
	toolVersion     string
	toolDescription string
	handler         func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

func NewCPFGeneratorTool() *CPFGeneratorTool {
	return &CPFGeneratorTool{
		toolName:        "CPF Generator",
		toolVersion:     "1.0.0",
		toolDescription: "Generates a valid CPF number.",
		handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
			formatted := request.Params.Arguments["formatted"].(bool)
			from_uf := request.Params.Arguments["state"].(string)

			service := services.NewCpfService()
			generated, err := service.Generate(formatted, from_uf)

			if err != nil {
				return mcp.NewToolResultError(err.Error()), nil
			}

			return mcp.NewToolResultText(*generated), nil
		},
	}
}

func (t *CPFGeneratorTool) Name() string {
	return t.toolName
}

func (t *CPFGeneratorTool) Handler() server.ToolHandlerFunc {
	return t.handler
}

func (c *CPFGeneratorTool) AsMCPTool() mcp.Tool {
	return mcp.NewTool(c.toolName, c.McpOptions()...)
}

func (t *CPFGeneratorTool) McpOptions() []mcp.ToolOption {
	return []mcp.ToolOption{
		mcp.WithDescription(t.toolDescription),
		mcp.WithBoolean(
			"formatted",
			mcp.Description("Generate CPF formatted with separator"),
			mcp.DefaultBool(true),
		),
		mcp.WithString(
			"state",
			mcp.Description("Generate CPF from a specific location state in brazil"),
		),
	}
}
