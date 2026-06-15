package main
import ("encoding/json";"fmt";"net/http";"os")
const svcName = "cache-store-ddf6e7"
type HealthResponse struct{Service string `json:"service"`;Status string `json:"status"`;Version string `json:"version"`}
func healthHandler(w http.ResponseWriter, r *http.Request){w.Header().Set("Content-Type","application/json");json.NewEncoder(w).Encode(HealthResponse{Service:svcName,Status:"ok",Version:"1.0.0"})}
func main(){port:=os.Getenv("PORT");if port==""{port="8080"};http.HandleFunc("/health",healthHandler);fmt.Printf("[%s] Listening on :%s\n",svcName,port)}
