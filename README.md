buildinfo
=========

A small library for go that can be used to include optional build information in your go binaries. 

Usage:

    //First step compiles the buildinfo library with the correct variables set
    go install -ldflags '-X github.com/sqdk/buildinfo.App APP_NAME -X github.com/sqdk/buildinfo.Branch DEVELOPMENT'
    //You can now build the binary that needs build information embedded
    go build

Please note that all future binaries will be compiled with the embedded information unless the buildinfo package is recompiled with new values.
