# Start
```bash
git clone --recursive https://github.com/biolee/swagger-ui-go-embed.git
cd swagger-ui-go-embed
glide install
go run ./example/main.go
```

# Usage

```
import (
	"net/http"

	"github.com/biolee/swagger-ui-go-embed/ui"
)
func main(){
	http.Handle("/",ui.Handler)
	http.ListenAndServe(":8081",nil)
}
```