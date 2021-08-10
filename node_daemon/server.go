package main

import (
	"github.com/emicklei/go-restful/v3"
	"github.com/sergomyaso/node-daemon/handlers"
	"log"
	"net/http"
)

func handleHttpError(err error, resp *restful.Response, httpStatus int, stringError string) error {
	if err != nil {
		resp.WriteErrorString(httpStatus, stringError)
		log.Println(err)
	}
	return err
}

func dump(req *restful.Request, resp *restful.Response) {
	dumpParams := new(handlers.DumpParams)
	err := req.ReadEntity(dumpParams)
	if handleHttpError(err, resp, 418, "bad request") != nil {
		log.Printf("bad request, error%v\n", err)
		return
	}
	status := handlers.CreateDump(dumpParams)
	resp.WriteError(status, nil)
}

func clearPageCash(req *restful.Request, resp *restful.Response) {
	//node1Param := handlers.DumpParams{TestName: "clspg", Ip: "192.168.0.71", PodUid: "none"}
	//node2Param := handlers.DumpParams{TestName: "clspg", Ip: "192.168.0.20", PodUid: "none"}
	//route := "/dump/clspg"
	//handlers.RedirectOnNode(&node1Param, route)
	//handlers.RedirectOnNode(&node2Param, route)
	resp.WriteError(200, nil)
	handlers.ClearPageCash()
}

func RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/dump")
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)

	ws.Route(ws.POST("/clspg").To(clearPageCash).
		Doc("clear page cash on nodes").
		Param(ws.BodyParameter("Data", "(JSON)").DataType("text")))

	ws.Route(ws.POST("/memory").To(dump).
		Doc("Create dump memory").
		Param(ws.BodyParameter("Data", "(JSON)").DataType("text")))

	ws.Route(ws.POST("/cpu").To(dump).
		Doc("Get test data").
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

	log.Print("start listening on localhost" + handlers.NodePort)
	log.Fatal(http.ListenAndServe(handlers.NodePort, wsContainer))

}
