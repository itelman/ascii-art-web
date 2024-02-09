## ascii-art-web

### Authors:

Arshat Aitkozha (@araitkozha), Ilyas Telman (@itelman).

### Description/Implementation details: algorithm

Ascii-art-web consists in creating and running a server, in which it will be possible to use a web **GUI** (graphical user interface) version of [ascii-art](../ascii-art) project.

The webpage allows the use of the following banners:

- [shadow](../ascii-art/shadow.txt)
- [standard](../ascii-art/standard.txt)
- [thinkertoy](../ascii-art/thinkertoy.txt)

The following HTTP endpoints were implemented:

1. GET request to route `/`: Sends HTML response, the main page.
2. POST request to route `/ascii-art`: that sends data to Go server (text and a banner).

The server displays the result of the provided POST request in the route `/ascii-art`, which applies the same template as the main page.

The main page has:

- a textarea for text input
- a dropdown menu for switching between banners
- a submit button, which sends a POST request and outputs the result on the page.

The website also returns the following HTTP status codes:

- OK (200), if everything went without errors.
- Not Found (404), if a resource is not found; for example a template, banner or a page.
- Bad Request (400), for requests with invalid text input.
- Method Not Allowed (405), for unsupported requests (GET requests, invalid request values, etc.).
- Internal Server Error (500), for unhandled errors.

## Usage: how to run

- Download the repository to your local machine.
- Open the repository via Terminal.
- Run the following command:
```console
go run main.go
```
- Run the server on:
```console
http://localhost:8080/
```