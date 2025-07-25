package tools

import (
	"context"
	"encoding/json"

	lm "github.com/icco/lunchmoney"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
	"github.com/wyvernzora/lunchmoney-mcp-server/internal"
)

// Category defines a transaction category, including its name,
// an optional description, and any nested subcategories.
type Category struct {
	Name          string     `json:"name"`
	Description   string     `json:"description,omitempty"`
	Subcategories []Category `json:"subcategories,omitempty"`
}

// ListCategoriesTool is an MCP server tool that lists transaction categories
// and their descriptions in a hierarchical format.
var ListCategoriesTool = server.ServerTool{
	Tool: mcp.NewTool("list_categories",
		mcp.WithDescription("List transaction categories and their descriptions"),
	),
	Handler: func(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
		client := internal.LunchMoneyClientFromContext(ctx)
		categories, err := client.GetCategories(ctx)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to call LunchMoney API", err), nil
		}
		results := map[string]any{
			"categories": transformCategories(categories),
		}
		str, err := json.Marshal(results)
		if err != nil {
			return mcp.NewToolResultErrorFromErr("failed to marshal the response", err), nil
		}

		return mcp.NewToolResultText(string(str)), nil
	},
}

// transformCategories converts LunchMoney API response into a structure that is easier for an LLM
// to consume
func transformCategories(data []*lm.Category) []Category {
	// Group categories by their parent ID
	childrenMap := make(map[int64][]*lm.Category)
	for _, cat := range data {
		if cat.GroupID != 0 {
			childrenMap[cat.GroupID] = append(childrenMap[cat.GroupID], cat)
		}
	}

	// Construct the category tree
	var result []Category
	for _, cat := range data {
		if cat.GroupID == 0 {
			var subs []Category
			for _, child := range childrenMap[cat.ID] {
				subs = append(subs, Category{
					Name:        child.Name,
					Description: augmentDescription(child),
				})
			}
			result = append(result, Category{
				Name:          cat.Name,
				Description:   augmentDescription(cat),
				Subcategories: subs,
			})
		}
	}
	return result
}

// augmentDescription adds additional information about category flags to the
// response description to be consumed by LLM
func augmentDescription(c *lm.Category) string {
	desc := c.Description
	if c.IsIncome {
		desc += " Transactions in this category are income, not expenses."
	}
	if c.ExcludeFromBudget {
		desc += " Transactions here are excluded from budget calculations."
	}
	if c.ExcludeFromTotals {
		desc += " Transactions here are excluded from totals calculations."
	}
	return desc
}
