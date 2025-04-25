package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
	"github.com/renanmedina/4devs-mcp/internal/tools"
	"github.com/renanmedina/4devs-mcp/observability"
)

func main() {
	logger := observability.GetLogger()

	logger.Info("Starting 4Devs MCP server...")
	s := server.NewMCPServer(
		"4Devs - MCP",
		"1.0.0",
	)

	toolsList := getTools()
	logger.Info("Registering tools...")
	for _, tool := range toolsList {
		logger.Info(fmt.Sprintf("Registering %s tool ...", tool.Name()))
		s.AddTool(tool.AsMCPTool(), tool.Handler())
		logger.Info(fmt.Sprintf("%s tool registered successfully ...", tool.Name()))
	}

	if err := server.ServeStdio(s); err != nil {
		logger.Error(err.Error())
		panic(err)
	}
}

func getTools() []tools.Tool {
	return []tools.Tool{
		tools.NewCPFGeneratorTool(),
	}
}
