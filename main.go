package main
import (
	"fmt"
	"flag"
	"github.com/labstack/echo"
	logger "EchoServer/logger"
	"EchoServer/configs"
// 	. "github.com/0x19/goesl"
// 	"github.com/gorilla/mux"
	"net/http"
	"os"
// 	"runtime"
//     "strings"
	
// 	"EchoServer/repository"
// 	adapters "EchoServer/adapters/repository"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

var (
    fshost   = flag.String("fshost", "172.28.230.128", "Freeswitch hostname. Default: localhost")
    fsport   = flag.Uint("fsport", 8021, "Freeswitch port. Default: 8021")
    password = flag.String("pass", "ClueCon", "Freeswitch password. Default: ClueCon")
    timeout  = flag.Int("timeout", 10, "Freeswitch conneciton timeout in seconds. Default: 10")
)


func handler(w http.ResponseWriter, r *http.Request) {
  var name, _ = os.Hostname()

  fmt.Fprintf(w, "<h1>This request was processed by host: %s</h1>\n", name)
}

func reloadxml(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {
		return
	}
	
	var name, _ = os.Hostname()

  fmt.Println("http.MethodPost : ",http.MethodPost);
  fmt.Println("r.Method : ",r.Method);
  
  fmt.Fprintf(w, "<h1>Reload Dialplan: %s</h1>\n", name)
}


func main(){
	e := echo.New()

	config := configs.GetConfig()	
	log := logger.NewLogger(config.Log.LogFile,config.Log.LogLevel)

	fmt.Println("API Listener Application Started")

	log1 := logger.NewLogger(config.Log.LogFile,config.Log.LogLevel);

// 	client, err := NewClient(*fshost, *fsport, *password, *timeout)
//      if err != nil {
//         Error("Error while creating new client: %s", err)
//         return
//     }

	fmt.Fprintf(os.Stdout, "Web Server started. Listening on 0.0.0.0:2000\n")
// 	r := mux.NewRouter()
// 	
// 	r.HandleFunc("/EchoServer/hostname", handler).Schemes("http")
// 	r.HandleFunc("/EchoServer/reloadxml", reloadxml).Schemes("http")
	
	
	
		
	http.HandleFunc("/EchoServer/hostname", handler)
	http.HandleFunc("/EchoServer/reloadxml", reloadxml)
	
	
	
	
	http.ListenAndServe(":2000", nil)
	
    

	if err := e.Start("0.0.0.0:10000"); err != nil {
		log.WithError(err).Fatal("echo server not able to start")
		log1.WithError(err).Fatal("Welcome")
	}
}


// func Get(uri string, timeout int, auth string, pass string, queryParam interface{}) (map[string]interface{}, int, error) {
// 	response := make(map[string]interface{})
// 	res, err := goreq.Request{
// 		Method:            "GET",
// 		Uri:               uri,
// 		ContentType:       "application/json",
// 		Accept:            "application/json",
// 		Timeout:           time.Second * time.Duration(timeout),
// 		BasicAuthUsername: auth,
// 		BasicAuthPassword: pass,
// 		QueryString:       queryParam,
// 	}.Do()
// 	return response, res.StatusCode, err
// }
// 
// func Post(uri string, timeout int, auth string, pass string, queryParam interface{}) (map[string]interface{}, int, error) {
//         response := make(map[string]interface{})
//         res, err := goreq.Request{ 
//                 Method:            "POST",
//                 Uri:               uri,
//                 ContentType:       "application/json",
//                 Accept:            "application/json",
//                 Timeout:           time.Second * time.Duration(timeout),
//                 BasicAuthUsername: auth,
//                 BasicAuthPassword: pass,
//                 Body:       queryParam,
//         }.Do()
//         return response, res.StatusCode, err
// }

