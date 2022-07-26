package main

import ("fmt"
			"io/ioutil"
			"net/http"
		)
func main(){
response, err :=http.Get(" https://api.coinbase.com/v2/prices/spot?currency=USD")

if err !=nil {

fmt.Printf("The Http Request Failled With error %s\n",err)

} else {

data, _:= ioutil.ReadAll(response.Body)
fmt.Println(string(data))

}
jsonData :=map[string] string {"firstname":"Azizul","lastname:":"Hoq"}
jsonValue, _:= json.Marshal(jsonData)
response, err = http.Post("https://httpbin.prg/post", "application/json", bytes.NewBuffer(jsonValue))
if err !=nil {

	fmt.Printf("The Http Request Failled With error %s\n",err)
	
	} else {
	
	data, _:= ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
	
	}

}

