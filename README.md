# Using [Memcached](https://memcached.org/) with Go language.

## Introduction

The repository contains the basic code to start using Memcached with a Go client.


## Installing Memcached

Multiple ways to install Memcached.

1. On Ubuntu OS using package manager, `apt install memcached`. A drawback is the version installed might be old.

2. Using the [Memcached Docker Image](https://hub.docker.com/_/memcached).

3. Installing from source, as shown below.  

    1. Download the latest tar from [here](https://memcached.org/downloads)
    2. Follow the instructions there to install memecached.
    3. Memcached binary will now be available at `/usr/local/memcahed/bin`.
    4. Copy the `memcached.service` file from `Scripts` folder in the downloaded tar to `/etc/systemd/system`.
    5. Copy the `memcached.sysconfig` which is the environment file, to an appropriate location, eg :- `/etc/sysconfig/` and possibly rename that file to `memcached`.
    6. Edit the `memcached.service` file to point to the correct environment file and binary.
    7. Issue `sudo systemctl daemon-reload` and start the service using `sudo systemctl start memcached`.
    8. Optionally issue `sudo systemctl enable memcached` to automatically start memcached during bootup.

## Understanding this repository

- Package `cmd/basic/main.go` has the basic code to save and retrive data from memcached server.

- Package `cmd/basic_serialize/main.go` shows how to store and retrieve any type of data to memcached using [gob package](https://pkg.go.dev/encoding/gob) to serialize and deserialize the data.
    - One issue here is that `gob` fails when the byte array is decoded into an `interface{}` variable, and hence needs to be passed the variable of correct concrete type to be decoded into.

- Package `cmd/advanced_serialization/main.go` has experimental code to work out how to decode the byte array received from memcached server into an `interface{}` variable using `gob` package. Haven't been able to get a working solution.
    - Doesn't work with primitive types, but does it work with structs?

- Package `cmd/with_msgpack/main.go` uses [msgpack](https://github.com/vmihailenco/msgpack) which has no issues decoding a byte array into an `interface{}` variable.