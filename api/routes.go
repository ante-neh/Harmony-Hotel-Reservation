package api 
import ("net/http")
func router() *http.ServeMux{
	mux := http.NewServeMux()

	return mux
}