# go-utils

## Description

Multiple tools and example for golang

- database : example to use database in golang
- fichier : example to use file creation and JSON reading
- mail : example on mail

## Install go linux

1. get go

- > curl -O https://go.dev/dl/go1.17.5.linux-amd64.tar.gz

2. extract the package

- > sudo tar -C /usr/local -xzf go1.15.4.linux-amd64.tar.gz

3. add the exec command to *~/.profile*
    
- > vim ~/.profile

4. add the following line

```sh
export GOPATH=$HOME/work
export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin
```

5. refresh the profile

- > source ~/.profile
