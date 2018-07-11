# Supermodel

Supermodel is a development tool for distributed systems programming. 

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


### Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/munjeli/supermodel/tags). 

## Usage


## Why
One definition of a distributed project is that it has multiple repositories. This is a frequent cause of confusion and breakage during maintenance and release. A supermodel is a project file that sets a metafilesytem architecture for the project with yaml. The definition of a project might include a service, its database configuration, the server build template, the provisioning code with Chef, Ansible or another provisioner, the logging and monitoring setup. The supermodel allows developers to share a definition of a workspace to build absolute paths for automation and vendored artifacts, a common requirement of immutable build systems. The project definition can make it easy to get new developers off the ground and set up. It also might help failures due to missing runtime dependencies or complicated updates for shared resources.

If you think about it, when you're working in an environment with end to end automation and immutable servers the whole infrastructure could be regarded as a single file system including the configuration of networking on a cloud and all the application code. Then every 'project' is really just a subset of the infrastructure. Setting up a repository of supermodels for your organization can help get everyone on the same page in discussions about systems. 

## Models

A model for Supermodel is a yaml file that describes a file system made of namespaces and repositories. This model will be created from the file by creating folders for the namespaces and git cloning the repositories. The model has a root which then can be a defined workspace for absolute paths in automation or vendoring artifacts. There are only two keys for a model hash: namespaces and repos.

Models are a logical definition and not enforced in any way. You could make models for various views of the source. For example, if a project manager wanted to see the git metrics or git logs of all the projects they oversee, that could be scripted against a custom model. Likewise models give us a reference for runtime dependencies that could be leveraged for troubleshooting and debugging. 

## Authors

* **Ele Munjeli** - *Initial work* - [munjeli](https://github.com/munjeli)

## License

This project is licensed under the Apache License - see the [LICENSE](LICENSE) file for details

## Acknowledgments

* Muchas gracias to Prime 8 Consulting where I first had the opportunity to program with meta models. 

