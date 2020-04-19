# Simple performance test for aiohttp endpoint that generates uuid

## Environment:

```
Kernel Version: 5.3.0-46-generic
OS Type: 64-bit
Processors: 8 × Intel® Core™ i5-8350U CPU @ 1.70GHz
Memory: 15.5 GiB of RAM
```

- Python 3.8.1
- Apache HTTP server benchmarking tool


## Install:

```
cd aiohttp

pip install -r requirements.txt
```

## Run:
```
python main.py

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


Server Software:        Python/3.8
Server Hostname:        localhost
Server Port:            8080

Document Path:          /prefix
Document Length:        50 bytes

Concurrency Level:      100
Time taken for tests:   29.971 seconds
Complete requests:      100000
Failed requests:        0
Total transferred:      20700000 bytes
HTML transferred:       5000000 bytes
Requests per second:    3336.57 [#/sec] (mean)
Time per request:       29.971 [ms] (mean)
Time per request:       0.300 [ms] (mean, across all concurrent requests)
Transfer rate:          674.48 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    1   0.5      1       8
Processing:    20   29   4.4     28      73
Waiting:       15   24   4.3     24      69
Total:         22   30   4.5     29      73

Percentage of the requests served within a certain time (ms)
  50%     29
  66%     30
  75%     32
  80%     33
  90%     35
  95%     38
  98%     43
  99%     46
 100%     73 (longest request)
```