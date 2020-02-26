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
#         the amount of passengers to change during a stop [rand(-delta,delta)] (default 25)
#   -nAutos int
#         number of autos to run during the simulation (default 5)
#   -passengers int
#         the amount of passengers autos start with (default 50)
#   -period duration
#         The periodicity of auto stops (default 5s)
```

All options will expose an API endpoint at `http://host:port/api/state`.
So the vanilla command will expose the following endpoint: `http://localhost:8080/api/state`.

### Example Commands

```bash
# Use all the defaults
./sunflower

# Bind to the global host (0.0.0.0) on the default port (8080)
# Simulate 100 autos, where each auto starts with 45 passengers.
# Each stop will occur at `t = n (1.5 seconds)`, where k = rand(-60,60) passengers get off/on.
./sunflower -nAutos 100 -passengers 45 -period 1500ms -delta 60 0.0.0.0
```


## Output Samples
```bash
./sunflower &
curl -v localhost:8080/api/state
```

Output:

```
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
< Date: Wed, 26 Feb 2020 07:37:52 GMT
< Content-Length: 309
<
{"NumAutos":5,"Autos":[{"ID":"56cc06c9-9cc6-42fb-b9e1-0a18e989a31c","Count":82},{"ID":"66ec6c5b-7d80-493e-afca-a8b77bebeca7","Count":0},{"ID":"eb2159b5-ea6e-4049-8fa9-8d53707682ad","Count":51},{"ID":"7fb57c21-6ca0-4771-8b19-076cf217afa2","Count":58},{"ID":"2081abdf-98ea-46aa-852f-bb5f785b1e99","Count":44}]}
* Connection #0 to host localhost left intact
```

Prettified JSON:

```json
{
  "NumAutos": 5,
  "Autos": [
    {
      "ID": "56cc06c9-9cc6-42fb-b9e1-0a18e989a31c",
      "Count": 82
    },
    {
      "ID": "66ec6c5b-7d80-493e-afca-a8b77bebeca7",
      "Count": 0
    },
    {
      "ID": "eb2159b5-ea6e-4049-8fa9-8d53707682ad",
      "Count": 51
    },
    {
      "ID": "7fb57c21-6ca0-4771-8b19-076cf217afa2",
      "Count": 58
    },
    {
      "ID": "2081abdf-98ea-46aa-852f-bb5f785b1e99",
      "Count": 44
    }
  ]
}
```
