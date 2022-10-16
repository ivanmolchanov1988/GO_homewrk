package main
 
import (
"bytes"
"crypto/tls"
"crypto/x509"
"fmt"
"io/ioutil"
"log"
"net/http"
"encoding/json"
)

type Iot struct {

    Id      int             `json:"id"`
	Name    string          `json:"name"`
}
 
func main() {
client := &http.Client{}
 
// Create a pool with the server certificate since it is not signed
// by a known CA
caCert, err := ioutil.ReadFile("./cert/server.crt")
if err != nil {
log.Fatalf("Reading server certificate: %s", err)
}
caCertPool := x509.NewCertPool()
caCertPool.AppendCertsFromPEM(caCert)
 
// Create TLS configuration with the certificate of the server
//tlsConfig := &tls.Config{
//RootCAs: caCertPool,
//}
//____DEV____Выключить проверку
tlsConfig := &tls.Config{
	InsecureSkipVerify: true,
}
 
// Use the proper transport in the client
client.Transport = &http.Transport{
TLSClientConfig: tlsConfig,
}
 
// Perform the request
resp, err := client.Post("https://localhost:9191/GO/connect", "text/plain", bytes.NewBufferString("Connect"))
if err != nil {
		log.Fatalf("Failed get: %s", err)
	}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
if err != nil {
		log.Fatalf("Failed reading response body: %s", err)
	}
fmt.Printf("Got response %d: %s %s", resp.StatusCode, resp.Proto, string(body))

// ___ DEV___
// var name string
// fmt.Scanf("%s\n", &name)


if resp.StatusCode == 200 {
	fmt.Println("OK")

	resp, err := client.Get("https://localhost:9191/GO/get")
	if err != nil {
		log.Fatalf("Failed GET: %s", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Failed reading response body from GET: %s", err)
	}
	fmt.Printf("Got response from GET %d \n", resp.StatusCode)
	fmt.Println(string(body))
	
	var iot Iot
	err = json.Unmarshal(body, &iot)
	if err != nil {
		panic(err)
	}

    out, err := json.Marshal(&iot)
    if err != nil {
        fmt.Println("Can't deserislize", body)
    }
	fmt.Println(string(out))


	var oId string
	fmt.Println("Выбрать id")
	fmt.Scanf("%s\n", &oId)

	// if oId ... pass

	resp, err = client.Post("https://localhost:9191/GO/id", "text/plain", bytes.NewBufferString(oId))
	if err != nil {
		log.Fatalf("ID Failed get: %s", err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("ID Failed reading response body: %s", err)
	}
	fmt.Printf("Got response from ID %d: %s %s", resp.StatusCode, resp.Proto, string(body))

	if resp.StatusCode == 200 {
		fmt.Println("ТУТ?")
		resp, err = client.Get("https://localhost:9191/GO/other")
		if err != nil {
			log.Fatalf("Other Failed get: %s", err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("Failed reading response body from Other: %s", err)
		}
		fmt.Printf("Got response from ID %d: %s %s", resp.StatusCode, resp.Proto, string(body))
	}
}


}



