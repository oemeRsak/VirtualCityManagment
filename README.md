# VirtualCityManagment

I just saw the control room of [Miniatur Wunderland Hamburg](https://link.com) and I want to create a virtual version of that.

## Technology

I need first simulate the enviroment. It should get the start data, maybe randomize it and than give regular output. I think it should be Object based. I would like to use Golang for it. The plan is give the output via a HTTP API or WebSocket. I need to find someway to sycronise the both faster.

Second thing that I need is the system that I monitor and control the simulation. For begin I will start with text based UI.

### Simulation

It should get the start data, maybe randomize it for begin. Simulate the thing like trafic. Output the data (position, status and more) of veheicle.
It need a map, maybe a grid-based map. I can assaign something to every cell, than set routes.
I want currently just veheicle managment.
I need to program the veheicles that they follow their router and in same time check the other veheicles and communicate. In a situlation create a warning to the system

### Managment System

The managment system is coming here to game. It will get the status, position and more from the simulation. It display it to user and e.g. in a conflict situlation between two veheiclers ask to user for a solution.
I think Python is good for it.

## Code

### Simulation Code

[main.go](Simulation/main.go) Simulation main code
[funcs.go](Simulation/funcs.go) The funcs
[structs.go](Simulation/structs.go) Object structs
[types.go](Simulation/types.go) Ready objects/types
[server.go](Simulation/server.go) HTTP handling functions
