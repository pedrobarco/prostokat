<h1 align="center">
  Prostokat
</h1>

<h4 align="center">A "polished" GNU/Linux tilling utility.</h4>

<p align="center"> <a href="https://github.com/pedrobarco/prostokat/issues">
    <img src="https://img.shields.io/github/issues/pedrobarco/prostokat"
         alt="issues">
  </a>
 <a href="https://github.com/pedrobarco/prostokat/network/members">
    <img src="https://img.shields.io/github/forks/pedrobarco/prostokat"
         alt="forks">
  </a>
 <a href="https://github.com/pedrobarco/prostokat/stargazers">
    <img src="https://img.shields.io/github/stars/pedrobarco/prostokat"
         alt="stargazers">
  </a>
 <a href="https://github.com/pedrobarco/prostokat/blob/main/LICENSE">
    <img src="https://img.shields.io/github/license/pedrobarco/prostokat"
         alt="LICENSE">
  </a>
</p> <p align="center">
  <a href="#key-features">Key Features</a> •
  <a href="#intro">Install</a> •
  <a href="#install">Install</a> •
  <a href="#how-to-use">How To Use</a> •
  <a href="#configuration">Configuration</a> •
  <a href="#download">Download</a> •
  <a href="#credits">Credits</a> •
  <a href="#contribute">Contribute</a> •
  <a href="#license">License</a>
</p>

## Key Features

- [x] Tile using a layout and a grid
- [x] Use a global mousebind to tile windows
- [ ] Tile by snapping windows to areas
- [ ] Choose your daemon: CLI vs GUI
- [x] Compatible with most GNU/Linux desktop environments
- [x] No dependencies, just a binary

## Intro

Prostokat uses profiles to manage different tilling configurations.
With pk one could create, edit, delete and activate profiles.

A **profile** is a configuration for your tilling: it has a **grid** and a **layout**.

A **grid** represents the way the screen is divided in equally sized cells.
- A 3x1 grid would divide a screen into 3, using 3 columns and 1 row.
- A 4x2 grid would divide a screen into 8, using 4 columns and 2 rows. 

```bash
# 3x1 grid
|-+-+-|
| | | |
|-+-+-|

# 4x2 grid
|-+-+-+-|
| | | | |
|-+-+-+-|
| | | | |
|-+-+-+-|
```

A **layout** is an array of areas in your grid which will be used to tile your windows to.
You can have as many areas in your layout as grid cells, so for a 2x2 grid you could not have more than 4 areas.

- A layout in a 2x2 grid could have 4 areas, one per grid cell
- A layout in a 3x1 grid could have 3 areas, dividing the screen in thirds
- A layout in a 4x1 grid could have 3 areas: 1 big one at the middle, and 2 smaller ones at the sides

```bash
# 2x2 grid with 4 areas 
|-+-|    |-+-|
| | |    |1|2|
|-+-| -> |-+-|
| | |    |3|4|
|-+-|    |-+-|

# 3x1 grid with 3 areas
|-+-+-|    |-+-+-|
| | | | -> |1|2|3|
|-+-+-|    |-+-+-|

# 4x1 grid with 3 areas
|-+-+-+-|    |-+-+-+-|
| | | | | -> |1| 2 |3|
|-+-+-+-|    |-+-+-+-|
```

An **area** is a tilling region that has coordinates and dimensions.
- posx: x position in grid
- posy: y position in grid
- height: number of rows to span
- width: number of columns to span

```yaml
# profiles/default.yaml
grid:
  cols: 3
  rows: 1
# splits monitor in thirds (grid: 3x1)
# |-+-+-|    |-+-+-|
# | | | | -> |1|2|3|
# |-+-+-|    |-+-+-|
layouts:
# area 1: pos(0, 0) dim(1x1)
- posx: 0
  posy: 0
  width: 1
  height: 1
# area 2: pos(1, 0) dim(1x1)
- posx: 1
  posy: 0
  width: 1
  height: 1
# area 3: pos(2, 0) dim(1x1)
- posx: 2
  posy: 0
  width: 1
  height: 1
```

## Install 

To clone and run this application, you'll need [Git](https://git-scm.com) and [Golang](https://golang.org/) installed on your computer. From your command line:

```bash
# Clone this repository
$ git clone https://github.com/pedrobarco/prostokat.git

# Go into the repository
$ cd prostokat

# Build the app
$ make clean && make build

# Install it
$ sudo make install

# Run the CLI
$ pk -h
```

## How to Use

Start prostokat via its CLI by running `pk start`. 
You should now be able to use `Shift + Mouse 3` to tile an active window to the closest tilling area.
> Note: Prostokat uses the current mouse position to detect which area should the window be tilled to. 

```
Usage:
  pk [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  help        Help about any command
  init        Initialize or reinitialize pk
  profiles    Manage prostokat profiles
  start       Start prostokat window manager

Flags:
  -h, --help   help for pk

Use "pk [command] --help" for more information about a command.
```

The default configuration splits the monitor in thirds.
However, you can have as many profiles as you would like.

### Configuring Profiles

Using `pk profiles` you can easily manage your profiles.

Suppose we want to have 2 new profiles: a `dev` and `games`.

For the `dev` profile: 
- tile to corners
- divides our monitor in 4

For the `games` profile: 
- a large area at the center 
- two equally sized areas at the sides

```
# profile/dev
|-+-|    |-+-|
| | |    |1|2|
|-+-| -> |-+-|
| | |    |3|4|
|-+-|    |-+-|

# profile/games
|-+-+-+-|    |-+-+-+-|
| | | | | -> |1| 2 |3|
|-+-+-+-|    |-+-+-+-|
```

Let's start by creating profiles.
> Note: pk uses the default profile template to create new profiles.

```bash
$ pk profiles create dev
+ $HOME/.config/prostokat/profiles/dev.yaml 
Profile dev was successfully created 

$ pk profiles create games
+ $HOME/.config/prostokat/profiles/profiles.yaml 
Profile games was successfully created 
```

Now we can edit the profiles. 
- For the `dev` profile we want a 2x2 grid with a 4 area layout.
- For the `games` profile we want a 4x1 grid with a 3 area layout.

Open and edit your profiles with `pk profiles edit [profile]` 

`$ pk profiles edit dev`

```yaml
# profiles/dev.yaml
grid:
  cols: 2
  rows: 2
layouts:
- posx: 0
  posy: 0
  width: 1
  height: 1
- posx: 1
  posy: 0
  width: 1
  height: 1
- posx: 0
  posy: 1
  width: 1
  height: 1
- posx: 1
  posy: 1
  width: 1
  height: 1
```

`$ pk profiles edit games`

```yaml
# profiles/games.yaml
grid:
  cols: 4
  rows: 1
layouts:
- posx: 0
  posy: 0
  width: 1
  height: 1
- posx: 1
  posy: 0
  width: 2
  height: 1
- posx: 3
  posy: 0
  width: 1
  height: 1
```

Now we can run `$ pk profiles activate dev` to set our default profile for `$ pk start`.
When playing games, we can run prostokat with a custom profile: `$ pk start -p games`.

Have fun!

## Download (WIP)

You can [download](https://github.com/pedrobarco/prostokat/releases) the latest installable version of Prostokat for GNU/Linux distros.

## Credits

This software uses the following open source packages:

- [xgbutils](https://github.com/BurntSushi/xgbutil)
- [xgb](https://github.com/BurntSushi/xgb)
- [cobra](https://github.com/spf13/cobra)
- [viper](https://github.com/spf13/viper)

## Contribute

1. [Fork it](https://github.com/pedrobarco/prostokat/fork)
2. Create your feature branch (`git checkout -b feature/myFeature`)
3. Commit your changes (`git commit -am 'add something'`)
4. Push to the branch (`git push origin feature/myFeature`)
5. Create a new [Pull Request](https://github.com/pedrobarco/prostokat/pulls)

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/pedrobarco/prostokat/blob/main/LICENSE) file for details

---

> [pedrobarco.github.io](https://pedrobarco.github.io) &nbsp;&middot;&nbsp;
> GitHub [@pedrobarco](https://github.com/pedrobarco)

