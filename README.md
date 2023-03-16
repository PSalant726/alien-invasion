# Alien Invasion!

Included is the source code to simulate an Alien invasion of an arbitrary World composed of interconnected Cities. A specified number of Aliens will invade the planet, and simultaneously move randomly between connected Cities. If more than one Alien ever resides in a given City, the Aliens battle, resulting in the destruction of that City and all Aliens therein. The simulation ends when all Aliens have been either trapped by the destruction of the World's Cities (rendering them unable to move to a new City), moved 10,000 times, or themselves destroyed. Finally, the resulting state of the World is printed to the console in the same format as the input file.

## Prerequisites

1. Go (v1.20 or later). Follow the instructions [here](https://go.dev/doc/install) if necessary.

## Usage

To quickly simulate the invasion of a small test World, navigate to the root directory of this repository in your terminal, then run:

```sh
go run ./... --aliens 10 --world-map "./world_map.txt"
```

Or, if you prefer to compile a binary, run:

```sh
go build -o ./bin/alien-invasion ./...
./bin/alien-invasion --aliens 10 --world-map "./world_map.txt"
```

To simulate the invasion of a custom World, a world map must be provided via the `--world-map` option. The format of the world map must be as follows:

```
City1 north=City2 south=City3 east=City4 west=City5
City2 south=City1
City3 north=City1
City4 west=City1
City5 east=City1
```

The world map file must include one City per line, connect any two Cities with at most a single route, and provide at least as many Cities as there are Aliens (provided via the `--aliens` option). On each line of the world map file, the City's name must come first, followed by up to four directional key/value pairs, all separated by a single space. Inter-City routes may be provided in any order, but the only valid keys are the four directions shown in the above example.

### Command Line Options Reference

| Option          | Default | Description                                                     |
|:---------------:|:-------:|:----------------------------------------------------------------|
| `--aliens`      |    2    | The amount of violent Alien invaders to unleash upon the world. |
| `--world-map`   |         | The path to the world map file.                                 |
| `--help` / `-h` |         | Print usage instructions.                                       |

### Running the Tests

Unit tests are included, purely as a demonstration of my ability to write them. No attempt was made to provide complete test coverage. To execute the included tests, navigate to the root directory of this repository in your terminal, then run:

```sh
go test ./...
```

## Assumptions

The utility assumes each of the below statements to be true:

1. In addition to ending when each Alien is trapped, destroyed, or has moved 10,000 times, the invasion will end if only a single Alien remains alive/untrapped.
1. There is no functional difference between an Alien that has been trapped and an Alien that has been destroyed.
1. In the World's initial state, additional Aliens cannot invade if every City is already occupied with an Alien.
    - This prevents Cities from being prematurely destroyed.
1. If any City ever houses more than a single Alien, that City and all Aliens residing therein will be destroyed.
    - Only the first two Aliens to reach the City are credited with the City's destruction in log output.


## Omissions

1. Structured and leveled logging (see: [zerolog](https://pkg.go.dev/github.com/rs/zerolog) and/or [logrus](https://pkg.go.dev/github.com/sirupsen/logrus))
1. "Dry" or "why" run support
