package tools

import (
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type Tool interface {
	Name() string
	AsMCPTool() mcp.Tool
	Handler() server.ToolHandlerFunc
}
