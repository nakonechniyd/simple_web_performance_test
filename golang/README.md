# Simple performance testing for golang endpoint that generates uuid

## Environment:

```
Kernel Version: 5.3.0-46-generic
OS Type: 64-bit
Processors: 8 × Intel® Core™ i5-8350U CPU @ 1.70GHz
Memory: 15.5 GiB of RAM
```

- GoLang 1.14.2 linux/amd64
- Apache HTTP server benchmarking tool


## Install:

```
go get github.com/google/uuid

go build service.go
```

## Run:
```
./service

ab -c 100 -n 100000 http://localhost:8080/prefix
```

## Result:

```
Benchmarking localhost (be patient)
Completed 10000 requests
Completed 20000 requests
Completed 30000 requests
Completed 40000 requests
Completed 50000 requests
Completed 60000 requests
Completed 70000 requests
Completed 80000 requests
Completed 90000 requests
Completed 100000 requests
Finished 100000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8080

Document Path:          /prefix
Document Length:        50 bytes

Concurrency Level:      100
Time taken for tests:   4.864 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      16700000 bytes
HTML transferred:       5000000 bytes
Requests per second:    20559.13 [#/sec] (mean)
Time per request:       4.864 [ms] (mean)
Time per request:       0.049 [ms] (mean, across all concurrent requests)
Transfer rate:          3352.91 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        1    2   0.7      2      15
Processing:     1    2   0.7      2      18
Waiting:        0    2   0.7      2      16
Total:          3    5   0.9      5      23

Percentage of the requests served within a certain time (ms)
  50%      5
  66%      5
  75%      5
  80%      5
  90%      5
  95%      6
  98%      6
  99%      7
 100%     23 (longest request)
```