package main
import (
 "context"
 "fmt"
 "os"
 "net/url"
 sdk "github.com/jsonscraper/jsonscraper-go"
)
func main() {
 client, err := sdk.NewTikTok(os.Getenv("JSONSCRAPER_TIKTOK_LICENSE_KEY"), sdk.Options{})
 if err != nil { panic(err) }
 response, err := client.GetVideoById(context.Background(), url.Values{"video_id":{"7106855913906081070"},"cache_timeout":{"60"}})
 if err != nil { panic(err) }
 fmt.Println(string(response.Body))
}
