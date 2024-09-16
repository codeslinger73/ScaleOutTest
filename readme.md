# ScaleOutTest
## Go app that is built to go inside of an Azure Function

# Build - written for go Go 1.21
## Clone repo
## cd cmd
## go build -o ../api
## cd ..

# run locally - Need AF core tools installed
## func start

# urls
## http://localhost:7071/api/ping
## http://localhost:7071/api/testshort?iterations=4000
## http://localhost:7071/api/testlong?iterations=4000

# iterations
## The app will generate random strings.  The testshort will generate a 500 byte string.  Testlong will generate a 3k string.  The iterations querystring param will instruct the app how many times to do that in a loop.