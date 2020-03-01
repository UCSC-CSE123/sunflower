# Sunflower

![LogoMakr_95uczH](https://user-images.githubusercontent.com/13544676/75320402-6fbd5480-5823-11ea-8300-aefef556e1c3.png)

A bus traffic simulator

## Building

To build the code, first clone this repo using `git clone`.

```bash
git clone https://github.com/UCSC-CSE123/sunflower.git
```

Then build the code using `go build`

```bash
cd sunflower
go build
```
### Installing from pre-compiled binaries

Pre-compiled binaries are available at https://github.com/UCSC-CSE123/sunflower/releases.



## Running

For all options run `sunflower help`

```bash
sunflower --help
# Output:
# Usage: ./sunflower [host address] [port] [flags]
#         host
#                 The host address to bind to (default localhost)
#         port
#                 The port to bind to (default 8080)
# flags:
#   -delta int
#         the amount of passengers to change during a stop [rand(-delta,delta)] (default 10)
#   -duration duration
#         The length of time a bus is stopped for (default 5s)
#   -nAutos int
#         number of autos to run during the simulation (default 5)
#   -passengers int
#         the amount of passengers autos start with (default 50)
#   -period duration
#         The periodicity of auto stops (default 11s)
#   -probability int
#         the probability that a bus stops (default 75)
#   -seed int
#         the seed to pass to the RNG -- by default the seed is the current time (default 1583086368860364528)
```

All options will expose an API endpoint at `http://host:port/api/state`.
So the vanilla command will expose the following endpoint: `http://localhost:8080/api/state`.

### Example Commands

```bash
# Use all the defaults
./sunflower

# Bind to the global host (0.0.0.0) on the default port (8080)
# Simulate 100 autos, where each auto starts with 45 passengers.
# Each stop will occur at `t = n (13 seconds)` for `t = 7 seconds`, where k = rand(-60,60) passengers get off/on with a probability of 50%
# Use 06202020 as the seed
./sunflower -nAutos 100 -duration 7s -passengers 45 -period 13000ms -delta 60 -probability 50 -seed 06202020 0.0.0.0
```


## Output Samples
```bash
./sunflower &
curl -v localhost:8080/api/state
```

Output:

```
2020/03/01 18:23:20 server parameters: {Host:localhost Port:8080 StopPeriod:11s StopDuration:5s StopProbability:75 Autos:5 InitialCount:50 Delta:10 Seed:1583087000342699
201}
2020/03/01 18:23:20 starting server at http://localhost:8080/api/state
*   Trying 127.0.0.1...
* TCP_NODELAY set
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /api/state HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.58.0
> Accept: */*
>
< HTTP/1.1 200 OK
< Content-Type: application/json
< Date: Sun, 01 Mar 2020 18:23:20 GMT
< Content-Length: 505
<
{"State":{"NumAutos":5,"Autos":[{"ID":"94831eb2-6530-4369-9018-c1b3571ddc76","Count":50,"Status":"In Transit"},{"ID":"d290e86f-7bd6-4d36-8074-3b47cdb3ac3b","Count":50,"Status":"In Transit"},{"ID":"4c3dd488-0d02-4d8e-af02-abde85060a97","Count":50,"Status":"In Transit"},{"ID":"319069b7-ce3d-40c5-8096-ee0409baea51","Count":50,"Status":"In Transit"},{"ID":"d0afec42-20ec-4d60-bd52-7fb77457bc32","Count":50,"Status":"In Transit"}]},"DebugInfo":{"StopPeriodicity":"11s","InitialCount":50,"ElapsedTime":"0s"}}
* Connection #0 to host localhost left intact
```

Prettified JSON:

```json
{
  "State": {
    "NumAutos": 5,
    "Autos": [
      {
        "ID": "94831eb2-6530-4369-9018-c1b3571ddc76",
        "Count": 50,
        "Status": "In Transit"
      },
      {
        "ID": "d290e86f-7bd6-4d36-8074-3b47cdb3ac3b",
        "Count": 50,
        "Status": "In Transit"
      },
      {
        "ID": "4c3dd488-0d02-4d8e-af02-abde85060a97",
        "Count": 50,
        "Status": "In Transit"
      },
      {
        "ID": "319069b7-ce3d-40c5-8096-ee0409baea51",
        "Count": 50,
        "Status": "In Transit"
      },
      {
        "ID": "d0afec42-20ec-4d60-bd52-7fb77457bc32",
        "Count": 50,
        "Status": "In Transit"
      }
    ]
  },
  "DebugInfo": {
    "StopPeriodicity": "11s",
    "InitialCount": 50,
    "ElapsedTime": "0s"
  }
}
```
