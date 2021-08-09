package main

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/terrServ/handlers"
	"github.com/terrServ/render"
	"log"
	"net/http"
)

type TerraformResponse struct {
	Response string `json:"response"`
	Error    error `json:"error"`
}

func createEcs(req *restful.Request, resp *restful.Response) {
	ecsParams := new(render.EcsParams)
	err := req.ReadEntity(ecsParams)
	if err != nil { // bad request
		resp.WriteErrorString(http.StatusBadRequest, err.Error())
		return
	}
	ecsScript := render.GetRenderEcsScript(ecsParams)
	log.Println(ecsScript)
	var result string
	result, err = handlers.RunUserScript(ecsScript)
	trResponse := new(TerraformResponse)
	trResponse.Error = err
	trResponse.Response = result
	resp.WriteEntity(trResponse)
}

func runUserScript(req *restful.Request, resp *restful.Response) {
	scriptFile, _, err := req.Request.FormFile("scriptFile")
	if err != nil {
		return
	}
	defer scriptFile.Close()
	var stringUserScript = handlers.GetDataFromFile(scriptFile)
	log.Println(stringUserScript)
	result, err := handlers.RunUserScript(stringUserScript)
	trResponse := new(TerraformResponse)
	trResponse.Error = err
	trResponse.Response = result
	resp.WriteEntity(trResponse)
}

func RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/terraform")
	//ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)

	ws.Route(ws.POST("/ecs/create").To(createEcs).
		Doc("Create ecs server").
		Param(ws.BodyParameter("Data", "(JSON)").DataType("main.EcsParams")))

	ws.Route(ws.POST("/run/script").To(runUserScript).
		Doc("Run user script").
		Param(ws.BodyParameter("Data", "(JSON)").DataType("text")))

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
