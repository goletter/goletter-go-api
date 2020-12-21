package bootstrap

import (
	"goletter-go-api/pkg/route"
	"goletter-go-api/routes"

	"github.com/gorilla/mux"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)
	routes.RegisterApiRoutes(router)

	route.SetRoute(router)

	return router
}
