Nomad LXC Driver
==================

- Website: https://www.nomadproject.io
- Mailing list: [Google Groups](http://groups.google.com/group/nomad-tool)

Requirements
------------

- [Nomad](https://www.nomadproject.io/downloads.html) 0.9+
- [Go](https://golang.org/doc/install) 1.11 (to build the provider plugin)
- Linux host with `liblxc` and `lxc-templates` packages installed

Building The Driver
---------------------

Clone repository to: `$GOPATH/src/github.com/hashicorp/nomad-driver-lxc`

```sh
$ mkdir -p $GOPATH/src/github.com/hashicorp; cd $GOPATH/src/github.com/
$ git clone git@github.com:hashicorp/nomad-driver-lxc
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/hashicorp/nomad-driver-lxc
$ make build
```

Using the driver
----------------------

_TODO: Fill this_

Developing the Provider
---------------------------

If you wish to work on the driver, you'll first need [Go](http://www.golang.org) installed on your machine, and have have `lxc-dev` and `lxc-templates` packages installed. You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```
