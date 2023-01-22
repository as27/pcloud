// Package api is the structure of the JSON returned by the pCloud API.
// The structs in this package can be used when calling pcloud.ApiRequest() as
// respStruct parameter. Therefore a empty variable of the struct needs to be
// created. A pointer to that variable needs to be passed to the ApiRequest() that
// the JSON response is decoded into the variable.
// The documentation to all the methods of the pCloud API can be found here:
// https://docs.pcloud.com/methods.html
package api
