module golang

// https://go.dev/doc/modules/gomod-ref
// At go 1.21 or higher
// - go directive specifies the minimum version of golang that can be used to build and run the module. 
//   to be compatible with go1.20, any feature introduced in a later version of golang cannot be used in this module
// - The go line must be greater than or equal to the go line of all dependencies.
//   Dependencies with a higher go directive than the main module will cause build failures.
go 1.22.6

// introduced in Golang 1.21, specifying the specific Golang version used to build the module. 
// This directive ensures consistent builds across different environments.
toolchain go1.22.2
