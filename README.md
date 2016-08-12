go-overlay2
===========

[![GoDoc](https://godoc.org/github.com/barracks510/go-overlay2?status.svg)](https://godoc.org/github.com/barracks510/go-overlay2)

Copyright (c) 2016 Dennis Chen <barracks510@gmail.com>

Description
-----------

Port of the Docker Overlay2 GraphDriver to a non-Docker implementation.

OverlayFS is a union filesystem that allows the presentation of a single filesystem that results from the overlay of one unto another.

Installation
------------

This package can be installed with the `go get` command:

    go get github.com/barracks510/go-overlay2

_go-overlay2_ depends on Linux kernel 4.0 or newer, compiled with OverlayFS. 
RedHat has backported the latest OverlayFS to kernel 3.10 for RHEL7.1 and CentOS 7.1.
Using go-overlay2 on these backported kernels has been minimally tested.
If you want to build your app with go-overlay2, you will need to cross compile for Linux.

This package currently is a work-in-progress. I do not forsee the API becoming stable in the near future.
It is not recommended to use this package without locking in or vendoring it as a dependency.

Documentation
-------------

API documentation may be found here: https://godoc.org/github.com/barracks510/go-overlay2  
Kernel documentation may be found here: https://www.kernel.org/doc/Documentation/filesystems/overlayfs.txt  
Docker docmentation may be found here: https://docs.docker.com/engine/userguide/storagedriver/overlayfs-driver/  


Examples can be found under the `./_example` directory

License
-------

Apache Software License 2.0. See COPYING for details.  
This repository contains code from other projects. See NOTICE for details.
