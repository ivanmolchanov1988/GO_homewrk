package main
 
import (
"fmt"
"io/ioutil"
"log"
"net/http"
//"encoding/json"
"strings"
)
 
var oBool = false

func main() {
http.HandleFunc("/GO/connect", echoPayload)
http.HandleFunc("/GO/get", Goget)
http.HandleFunc("/GO/id", Id)
//http.HandleFunc("/GO/avia", avia)
http.HandleFunc("/GO/other", Other)
log.Printf("Go Backend: { HTTPVersion = 1 }; serving on https://localhost:9191/GO/connect")
log.Fatal(http.ListenAndServeTLS(":9191", "./cert/server.crt", "./cert/server.key", nil))
}
 
func echoPayload(w http.ResponseWriter, req *http.Request) {

new_body_content := "New content."
req.Body = ioutil.NopCloser(strings.NewReader(new_body_content))



log.Printf("Request connection: %s, path: %s", req.Proto, req.URL.Path[1:])
defer req.Body.Close()
//w.Write(bytes)
contents, err := ioutil.ReadAll(req.Body)
if err != nil {
log.Fatalf("Oops! Failed reading body of the request.\n %s", err)
http.Error(w, err.Error(), 500)

} else {
	oBool = true
}
fmt.Println(string(contents))
fmt.Println(string(string(http.StatusForbidden))) //WTF?
fmt.Fprintf(w, "%s\n", string(contents))
}

func Goget(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	if oBool == true {
		fmt.Println(" GET bool OK")

	respGet :=[]byte(`
	{
		"id":1,
		"name":"test"
	}
	`)
    bytes := respGet
    // if err != nil {
    //     fmt.Println("Can't serislize", respGet)
    // }
    fmt.Printf("%v => %v, '%v'\n", respGet, bytes, string(bytes))
    w.Write(bytes)



	} // bool == true GET end
}

func Id(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()
	contents, err := ioutil.ReadAll(req.Body)
	if err != nil {
		log.Fatalf("Oops! Failed reading body of the request.\n %s", err)
		http.Error(w, err.Error(), 500)

	} else {
		oBool = true
	}
	fmt.Println(string(contents))

	// if ID ... pass

	// Бронирование билета у авиакомпании
	// resp, err = client.Post pass
	var avia bool
	avia = true //false?
	if avia {
		bytes := []byte(" true")
		w.Write(bytes)
	}

	// предложение доп опций

	
}

func Other(w http.ResponseWriter, req *http.Request) {
	defer req.Body.Close()

	respOther :=[]byte(`
	{
		"id":2,
		"name":"test2"
	}
	`)
	fmt.Printf("%v => %v, '%v'\n", respOther, respOther, string(respOther))
    w.Write(respOther)
}

// оплата в банке




