package main

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/terrServ/handlers"
	"log"
	"net/http"
)

type TerraformResponse struct {
	Response string
	Error    error
}

func createEcs(req *restful.Request, resp *restful.Response) {
	ecsParams := new(handlers.EcsParams)
	err := req.ReadEntity(ecsParams)
	if err != nil { // bad request
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	ecsScript := handlers.GetRenderEcsScript(ecsParams)
	log.Println(ecsScript)
	var result string
	result, err = handlers.RunEcsScript(ecsScript)
	trResponse := new(TerraformResponse)
	trResponse.Error = err
	trResponse.Response = result
	resp.WriteEntity(trResponse)
}

func RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/sbercloud")
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)

	ws.Route(ws.POST("/ecs/create").To(createEcs).
		Doc("Create ecs server").
		Param(ws.BodyParameter("Data", "(JSON)").DataType("main.EcsParams")))

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
