package rest

import ( // this line is used by starport scaffolding # 1
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"
)

const (
	// MethodGet http get method name
	MethodGet = "GET"
)

// RegisterRoutes registers ibc-identifier-related REST handlers to a router
func RegisterRoutes(clientCtx client.Context, r *mux.Router) {
	// this line is used by starport scaffolding # 2
}
