//go:build ignore

// Run this file alone with: go run examples/media.go
package main
import("context";"fmt";"os";"net/url";sdk "github.com/jsonscraper/instagram-api-go")
func main(){c,e:=sdk.NewInstagram(os.Getenv("JSONSCRAPER_INSTAGRAM_LICENSE_KEY"),sdk.Options{});if e!=nil{panic(e)};r,e:=c.Request(context.Background(),"/v1/media/by/code",url.Values{"code":{"DT23hQ-EgkI"}, "cache_timeout":{"60"}});if e!=nil{panic(e)};fmt.Println(string(r.Body))}
