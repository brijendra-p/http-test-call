# http-test-call
Go tool to test API connection or call API for a specified number of times within a specified interval.


Parameters:
* url: Target API URL (required)
* data: JSON payload (required)
* n: Number of requests (required)
* wait: Wait time in seconds between requests (required)

How to run:

```
 ./http-test-call  -url "http://yout_url/api" -n 5 -wait 2 -data '{
    "id": 531
}'
```
