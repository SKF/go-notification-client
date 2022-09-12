# go-notification-client

This is the client library for the Notification API.

## Example

Refer to [example](/example/) for examples on how to use this library.

## Error handling

Errors returned by the API are decoded into problems from the [`github.com/SKF/go-rest-utility`](https://github.com/SKF/go-rest-utility) package before being returned by the client functions. This makes it possible to use the standard [`error`](https://pkg.go.dev/errors) package to do error checking on any returned error.
