//go:build ignore

package main
import (
 "context"
 "fmt"
 "os"
 "net/url"
 sdk "github.com/jsonscraper/instagram-api-go"
)
func main() {
 client, err := sdk.NewInstagram(os.Getenv("JSONSCRAPER_INSTAGRAM_LICENSE_KEY"), sdk.Options{})
 if err != nil { panic(err) }
 response, err := client.V1UserByUsername(context.Background(), url.Values{"username":{"apple"}, "cache_timeout":{"60"}})
 if err != nil { panic(err) }
 fmt.Println(string(response.Body))
}
