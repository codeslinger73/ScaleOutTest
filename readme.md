# Description
This app came about because of a request to test the scale out capabilities of an Azure region.  In the past our requests for new instances based on out scale out plan had been delayed hours or even days.  We received information that the region had been upgraded.  I looked for an app but didn't find anything that seemed like it would work well so I created this one and wired up the customer handler for an Azure Funcation interface.  Nothing special.  Just creates random strings in a loop to tie up some CPU and force scale out.

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
