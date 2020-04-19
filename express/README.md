# Simple performance testing for express endpoint that generates uuid

## Environment:

```
Kernel Version: 5.3.0-46-generic
OS Type: 64-bit
Processors: 8 × Intel® Core™ i5-8350U CPU @ 1.70GHz
Memory: 15.5 GiB of RAM
```

- Node Version Manager
- NodeJS 12.16
- Apache HTTP server benchmarking tool


## Install:

```
nvm install 12.16.2
nvm use 12.16.2

cd express
npm i
```

## Run:
```
nvm use 12.16.2

node index.js

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
Document Length:        49 bytes

Concurrency Level:      100
Time taken for tests:   10.388 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      25600000 bytes
HTML transferred:       4900000 bytes
Requests per second:    9626.35 [#/sec] (mean)
Time per request:       10.388 [ms] (mean)
Time per request:       0.104 [ms] (mean, across all concurrent requests)
Transfer rate:          2406.59 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.4      0       8
Processing:     6   10   3.2      9      68
Waiting:        3    7   2.4      7      63
Total:          7   10   3.4      9      77
ERROR: The median and mean for the initial connection time are more than twice the standard
       deviation apart. These results are NOT reliable.

Percentage of the requests served within a certain time (ms)
  50%      9
  66%     10
  75%     10
  80%     11
  90%     12
  95%     15
  98%     20
  99%     24
 100%     77 (longest request)
```