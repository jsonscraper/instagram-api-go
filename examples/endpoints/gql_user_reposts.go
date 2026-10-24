//go:build ignore

// Get user reposts. GET /gql/user/reposts
// Run this standalone example: go run examples/endpoints/gql_user_reposts.go --help
package main
import("context";"encoding/json";"errors";"flag";"fmt";neturl "net/url";"os";"strings";"time";sdk "github.com/jsonscraper/instagram-api-go")
var (
 timeout=flag.Duration("timeout",30*time.Second,"Request timeout, e.g. 30s")
 retries=flag.Int("retries",0,"Optional retries, 0 to 5; may consume quota")
 output=flag.String("output","","Save JSON or binary response to a file")
 user_id=flag.String("user_id","1107766105","API query parameter: user_id")
 repost_next_max_id=flag.String("repost_next_max_id","","API query parameter: repost_next_max_id")
 flat=flag.String("flat","","API query parameter: flat")
 cache_timeout=flag.String("cache_timeout","60","API query parameter: cache_timeout")
)
func run() error {
 flag.Parse()
 key:=strings.TrimSpace(os.Getenv("JSONSCRAPER_INSTAGRAM_LICENSE_KEY"));if key==""{return errors.New("set JSONSCRAPER_INSTAGRAM_LICENSE_KEY before making a licensed request")}
 if *timeout<=0||*retries<0||*retries>5{return errors.New("use a positive timeout and 0 to 5 retries")}
 client,err:=sdk.NewInstagram(key,sdk.Options{Timeout:*timeout,Retries:*retries});if err!=nil{return err}
 ctx,cancel:=context.WithTimeout(context.Background(),*timeout*time.Duration(*retries+1)+time.Minute);defer cancel()
 query:=neturl.Values{"user_id":{*user_id}, "repost_next_max_id":{*repost_next_max_id}, "flat":{*flat}, "cache_timeout":{*cache_timeout}}
 supplied:=map[string]bool{};flag.Visit(func(f *flag.Flag){supplied[f.Name]=true})
 if !supplied["repost_next_max_id"]{query.Del("repost_next_max_id")}; if !supplied["flat"]{query.Del("flat")}
 response,err:=client.Request(ctx,"/gql/user/reposts",query);if err!=nil{return err}
 if !strings.Contains(response.Header.Get("Content-Type"),"json"){
  if *output==""{return errors.New("binary response: provide --output=filename")}
  return os.WriteFile(*output,response.Body,0600)
 }
 var envelope map[string]any
 if err=response.Decode(&envelope);err!=nil{return fmt.Errorf("invalid JSON response: %w",err)}
 if envelope["status"]!="ok"||envelope["rate_limit_reached"]==true{return errors.New("API application error or quota limit; inspect the response envelope in your application")}
 rendered,err:=json.MarshalIndent(envelope,"","  ");if err!=nil{return err}
 if *output!=""{return os.WriteFile(*output,append(rendered,'\n'),0600)}
 fmt.Println(string(rendered));return nil
}
func main(){if err:=run();err!=nil{var apiError *sdk.APIError;if errors.As(err,&apiError){fmt.Fprintf(os.Stderr,"HTTP %d; Retry-After: %s\n",apiError.Status,apiError.RetryAfter)}else{fmt.Fprintln(os.Stderr,err)};os.Exit(1)}}
