# Pearbutter

```
                      __        __  __
   ___  ___ ___ _____/ /  __ __/ /_/ /____ ____
  / _ \/ -_) _ `/ __/ _ \/ // / __/ __/ -_) __/
 / .__/\__/\_,_/_/ /_.__/\_,_/\__/\__/\__/_/
/_/
```

RSS feed to IRC messages translator.

## Features

- Custom formatting
- Multiple {servers, channels, feeds} support
- Onconnect command

## Installation

### Docker (recommended)

1. `git clone https://github.com/X3NOOO/pearbutter`
2. `cd ./pearbutter/`
3. `bash ./build.sh config`
4. `docker build -t pearbutter .`
5. `docker run -d --restart unless-stopped --name pearbutter pearbutter`

### Raw binary

1. `git clone https://github.com/X3NOOO/pearbutter`
2. `cd ./pearbutter/`
3. `bash ./build.sh config release run`
