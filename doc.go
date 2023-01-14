/*
Package pcloud provides a client for accessing the pCloud API.

The pCloud API is located in the US and EU. If you are using the API it is important that you set the correct location for your account. Otherwise the authentication will fail. The default HostURL is the EU API. To change the host use the constants HostURLEU or HostURLUS. If there will be another location in the future it is also possible to set that variable with the URL. For security reasons it is not recomended to change the protocol from HTTPS to HTTP.

For the authentication there are two different methods to get the auth-token. With AuthMethodPassword the password is send to the server as a get-parameter. This way is not recomended, because it not a good idea to send a password directly to pCloud. Even when HTTPS is used. The Default method is AuthMethodDigest. In this case first a digest is requested from the api-server. This digest is used to hash the password with the username to receive an auth-token. Cause of the extra request that way is a little bit slower. But that is just once, when the auth-token is requested.

To get the data this package does just use the JSON API from pCloud. The binary API is not supported. The JSON API is documented here: https://docs.pcloud.com/methods/
*/
package pcloud
