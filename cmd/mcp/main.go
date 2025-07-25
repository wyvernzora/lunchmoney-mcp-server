package main

import (
	"log"
	"os"

	"github.com/mark3labs/mcp-go/server"
	"github.com/wyvernzora/lunchmoney-mcp-server/internal"
	"github.com/wyvernzora/lunchmoney-mcp-server/pkg/tools"
)

func main() {
	bindAddr := os.Getenv("BIND_ADDRESS")
	if bindAddr == "" {
		bindAddr = "0.0.0.0"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	lmToken := os.Getenv("LUNCHMONEY_TOKEN")
	if lmToken == "" {
		log.Fatal("LUNCHMONEY_TOKEN is required but absent")
	}

	httpServer := server.NewStreamableHTTPServer(
		createMCPServer(),
		internal.WithLunchMoneyClient(lmToken),
	)
	log.Printf("HTTP server listening on %s:%s/mcp", bindAddr, port)
	if err := httpServer.Start(bindAddr + ":" + port); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}

func createMCPServer() *server.MCPServer {
	mcpServer := server.NewMCPServer(
		"lunchmoney-mcp-server",
		"1.0.0",
		server.WithToolCapabilities(true),
		server.WithLogging(),
	)
	mcpServer.AddTools(tools.ListCategoriesTool)

	return mcpServer
}
