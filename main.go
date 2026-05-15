package main
import (

	"fmt"

	"net/http"
)
func handlerAPI(w http.responseWriter, r *http.Request){
	
}

func main(){
	
	http.HandleFunc("/api", handlerAPI)
	http.ListenAndServe(":8080", nil)
	fmt.Println("Server started on port :8080"
}
