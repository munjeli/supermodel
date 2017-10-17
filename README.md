# Supermodel

Supermodel is a development tool for distributed systems programming. One definition of a distributed project is that it has multiple repositories. This is a frequent cause of confusion and breakage during maintenance and release. A supermodel is a project file that sets a metafilesytem architecture for the project with yaml. The definition of a project might include a service, its database configuration, the server build template, the provisioning code with Chef, Ansible or another provisioner, the logging and monitoring setup. The supermodel allows developers to share a definition of a workspace to build absolute paths for automation and vendored artifacts, a common requirement of immutable build systems. The project definition can make it easy to get new developers off the ground and set up. It also cuts down on failures due to missing runtime dependencies or updates for shared resources.

## Getting Started

For local development of Supermodel, first make sure Go is properly installed and that a
[GOPATH](http://golang.org/doc/code.html#GOPATH) has been set. You will also need to add `$GOPATH/bin` to your `$PATH`.

Next, using [Git](https://git-scm.com/), clone this repository into `$GOPATH/src/github.com/munjeli/supermodel`. All the necessary dependencies are either vendored or automatically installed, so you just need to type `make`. This will compile the code and then run the tests. If this exits with exit status 0, then everything is working!

```sh
$ cd "$GOPATH/src/github.com/munjeli/supermodel"
$ make
```


### Dependencies

Supermodel stores its dependencies under `vendor/`, which [Go 1.6+ will automatically recognize and load](https://golang.org/cmd/go/#hdr-Vendor_Directories). We use [`govendor`](https://github.com/kardianos/govendor) to manage the vendored dependencies.


## Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/munjeli/supermodel/tags). 

## Authors

* **Ele Munjeli** - *Initial work* - [munjeli](https://github.com/munjeli)

## License

This project is licensed under the Apache License - see the [LICENSE.md](LICENSE.md) file for details

## Acknowledgments

* Muchas gracias to Prime 8 Consulting where I first had the opportunity to program with meta models. 

