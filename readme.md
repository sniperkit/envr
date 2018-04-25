# ENVentoRy

# envr
[![GitHub issues](https://img.shields.io/github/issues/dunstontc/enventory.svg)](https://github.com/dunstontc/enventory/issues)
[![License](https://img.shields.io/github/license/dunstontc/enventory.svg)](https://github.com/dunstontc/enventory/blob/master/LICENSE)

## Commands: ##
  - `-l, --list <source>`
  - `-g, --log <source>`
  - `-u, --update <source>`

## Sources ##

### Package Managers ###
  - brew
    - [braumeister](http://braumeister.org/)
  - [cargo](https://github.com/rust-lang/cargo)
    - [(crates)](https://crates.io/)
  - [composer](https://github.com/composer/composer)
  - gem
    - [(rubygems)](https://rubygems.org/)
  - go
  - luarocks
  - npm
    - [(npms)](https://npms.io/)
    - [(bower)](https://bower.io/search/)
  - [Paket](https://github.com/fsprojects/Paket)
  - pip2/3
    - [(warehouse)](https://pypi.org/)
  - vim plugins
  - [yarn](https://github.com/yarnpkg/yarn)
    - [(yarnipkg)](https://yarnpkg.com/en/)
  - [zplug](https://github.com/zplug/zplug)

#### Commands ####
  - `npm ls -g --depth 0 -json`
    - name,version(version),from(cmd),resolved(repo)
  - `pip2 list --form=json`, `"pip2", "list", "--form", "json"`
    - name, version
  - `pip3 list --form=json`, `"pip3", "list", "--form", "json"`
    - name, version
  - `brew list --versions` - `brew leaves`
    - name, version
  - `brew cask list -1`
    - name
  - `gem list`
    - name, version `([a-zA-Z0-9_-])+\s(\(.+\))`
  - `luarocks list --porcelain`
    - name version, location `[a-zA-Z0-9_-]+(?:\s+)[a-zA-Z0-9.-_]+(?:\s+installed).+`
  - `ls ~/.vscode/extensions`
  - `ls ~/.vscode-insiders/`
  - `apm ls --no-color`
    - split by empty line
    - Parse each section

### Versions ###
  - `go version`
  - `ruby --version`
  - `gem --version`
  - `nvim --version`
  - `vim --version`
  - `npm --version`
  - `node --version`
  - `cargo --version`
  - `lua -v`
  - `luarocks --version`


## [License](https://github.com/dunstontc/enventory/blob/master/LICENSE)
