import "net/http"

func main() {
	http.Handler("/helloworld", helloWorldHandler)

	log.Fatal(http.ListenAndServe(":8080", nil)
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World")
}