## server 1

Listen and counts the request on /count

## server 2

Prints out HEADER info to the server

EX: 

``` 
GET / HTTP/1.1
Header["Accept"] = ["text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9"]
Header["Sec-Fetch-Mode"] = ["navigate"]
Header["Connection"] = ["keep-alive"]
Header["Sec-Ch-Ua-Mobile"] = ["?0"]
Header["Dnt"] = ["1"]
Header["User-Agent"] = ["Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/ Safari/"]
Header["Upgrade-Insecure-Requests"] = ["1"]
Header["Sec-Fetch-Site"] = ["none"]
Header["Accept-Encoding"] = ["gzip, deflate, br"]
Header["Accept-Language"] = ["en,no;q=0.9,nb;q=0.8,la;q=0.7"]
Header["Sec-Ch-Ua"] = ["\"Chromium\";v=\"88\", \"Google Chrome\";v=\"88\", \";Not A Brand\";v=\"99\""]
Header["Cache-Control"] = ["max-age=0"]
Header["Sec-Fetch-User"] = ["?1"]
Header["Sec-Fetch-Dest"] = ["document"]
Host = "localhost:8090"
RemoteAddr = "[::1]:638789"

```