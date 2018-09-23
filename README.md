# nfl-cli
> CLI for NFL Game center

Work in progress cli for NFL game stats, with live updates in the game is in progress.  Currently the the cli displays game play-by-play and scoring drives. More on the way...

### Prerequisites 
[Sportradar NFL API](https://sportradar.us/nfl-api/) API Key


### Environments

```
SR_API_KEY=YOUR_API_KEY
```

### Example usage:
```
./nfl-cli -w=3 -y=2018 -st=REG
```

### Flags:
| Option | Description                           | Default |
| ------ |-------------------------------------- | ------- |
| w      | schedule week                         |       1 |
| y      | year                                  |    2018 |
| st     | season type                           |       1 |

