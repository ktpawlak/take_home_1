# Take home exercise 1

## Purpose
A simple http server providing RestAPI enpoints for getting server's version and system startup times.

### Compilation:
```
go build system_info.go
```

### Output:
```
$ ./system_info &
Server ready, endpoints: /version and /duration

$ curl http://localhost:8080/
/version: returns server version
/duration: returns startup duration

$ curl http://localhost:8080/version
version: 1.0.0

$ curl http://localhost:8080/duration
kernel: 1.897s
userspace: 8.162s
total: 10.059s
```

## JSON output
It is possible to return output in JSON format. Toggle const printJSON variable to true.

### Output:
```
$ curl http://localhost:8080/
{"/duration":"returns startup duration","/version":"returns server version"}

$ curl http://localhost:8080/duration
{"kernel":"1.897s","total":"10.059s","userspace":"8.162s"}

$ curl http://localhost:8080/version
{"version":"1.0.0"}
```