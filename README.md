# SunSpec tools for Go

This package contains Go data types representing the [SunSpec][] information model.
Subpackages implement particular use-cases that rely on this domain model:

 * [/core](./core) contains the SunSpec model block definitions and functions for working with them
 * [/modbus](#) contains utilities for talking to SunSpec devices via Modbus RTU
 * [/xml](./xml) contains utilities for exchanging SunSpec data using the XML format
 * [/generators](./generators) implements the code that transforms the [SMDX models][SMDX] into Go code for the above modules
 * [/server](#) is an example HTTP server that receives SunSpec XML data

[SunSpec]: http://sunspec.org/
[SMDX]: https://github.com/sunspec/models

## Generated code

This package uses typesafe representations of each model type.
To avoid excessive manual maintenance, structs for each type of model defined by SunSpec are generated from the [SMDX model files][SMDX].
To generate this code, first clone the SunSpec SMDX models repository into this directory:

    git clone git://github.com/sunspec/models.git

Then run the generators for each submodule that needs codegen:

    go generate ./core
