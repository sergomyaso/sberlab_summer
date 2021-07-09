package main

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/terrServ/handlers"
	"log"
	"net/http"
)

var configScriptProvider = ""

func createEcs(req *restful.Request, resp *restful.Response) {
	ecsParams := new(handlers.EcsParams)
	err := req.ReadEntity(ecsParams)
	if err != nil { // bad request
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	ecsScript := handlers.GetRenderEcsScript(ecsParams)
	if configScriptProvider != "" {
		handlers.RunEcsScript(configScriptProvider, ecsScript)
		resp.WriteEntity(ecsParams)
		return
	}
	resp.WriteErrorString(404,"configure failed")
}

func createConfig(req *restful.Request, resp *restful.Response) {
	configParams := new(handlers.ProviderConfig)
	err := req.ReadEntity(configParams)
	if err != nil { // bad request
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	configScriptProvider = handlers.GetRenderConfigScript(configParams)
	log.Println("config script:" + configScriptProvider)

	resp.WriteEntity(configParams)
}

func RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/sbercloud")
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)

	ws.Route(ws.POST("/create/ecs").To(createEcs).
		Doc("Create ecs server").
		Param(ws.BodyParameter("Data", "(JSON)").DataType("main.EcsParams")))

	ws.Route(ws.POST("/config").To(createConfig).
		Doc("Config provider").
		Param(ws.BodyParameter("Data", "(JSON)").DataType("main.ConfigParams")))

	container.Add(ws)
}

func CORSFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	resp.AddHeader(restful.HEADER_AccessControlAllowOrigin, "*")
	chain.ProcessFilter(req, resp)
}

func main() {
	wsContainer := restful.NewContainer()
	RegisterTo(wsContainer)
	// Add container filter to enable CORS
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"PUT", "POST", "GET", "DELETE"},
		AllowedDomains: []string{"*"},
		CookiesAllowed: false,
		Container:      wsContainer}
	wsContainer.Filter(cors.Filter)

	// Add container filter to respond to OPTIONS
	wsContainer.Filter(wsContainer.OPTIONSFilter)
	wsContainer.Filter(CORSFilter)

	log.Print("start listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", wsContainer))
}
