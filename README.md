# go-smtp
a mailjet package implementation

<h6> Define these environment variables </h6>

PORT='http server port' <br>
staticapikey='static key for x-api header' <br>
MJ_APIKEY_PUBLIC='mailerjet public key' <br>
MJ_APIKEY_PRIVATE='mailerjet private key' <br>

generate api key from here(for x-api header):
<a href='https://codepen.io/corenominal/pen/rxOmMJ?__cf_chl_captcha_tk__=dcc83d9775c5bcdb7e4a214fc450e5088bee529b-1609791166-0-AaCrGQVFzVYAHyyvpMR-T4-jUJua_WtzDRQtiLrouRohRZQ7ouwIYN2S0e5Q4qXN6cPOkkWYaCHV9cerIo2DV0eDjrt13uvt9ew4jYpf_gqbsuvTqU2q6ibFOJ3PO7CUW_bl83qMjKROgYvFg2p1OSvWrKPIKg5H9FjM3Am1zlRUB71P6df4thhoCw_oUAttBN3uTsnsB6oAvAsjdoASYTjdOqipEzcNdU8IL8Yv5P4j5d5N7dOmSo9cTNOXiC53KOaZbD6dK7nV4vVrGWj6lSavMBiUyPw0cdspvP8x60CwrMsXEhl5J-Qg5suiJ7vHSwDFYi1qaxl5S3W_7mGAOdgEaa1oS2QaiEhem9a93RT9t5BIRdoX3PHmez0qvYiepuNSZUof4d6Zp4ZeOhWzZG-xhOMX4dSiP28ZzOPgIyrablt_REK7kkxRyuZIOqTO4EB2NTBHDd7qauwMyrF7oa3FMqWGuPeiYW5HHauCHIKnL9cIPoml9EkGqV8lbUU4MeLJxzb2QR4HwhmGmruIMoDafqXMcZOK6o_lQARjNErjKZya-NIvdPVGE2S-2jcP-l5OGUJLxWQATSN8_8qaFbt3-0yLbPkAM-FypVyzbw56'>API KEY Generator </a>

for executing the program in local copy into go workspace and in terminal type this command when pwd is project directory:

go run cmd\gosmtp-server\main.go
