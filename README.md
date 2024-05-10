## ascii-art-web

## Authors:

Ilyas Telman (@itelman), Arshat Aitkozha (@araitkozha).

## Description/Implementation details: algorithm

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

### Output Download/Export in a File

In addition, the project consists on making sure that it is possible to export the output of the web application in `.txt` format. The file is exported with the right permissions (**read and write**) for the user.

In order to send a file as part of the HTTP response, the use of [HTTP headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers) was included.

Headers used for file transfer are [Content-Type](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Type), [Content-Length](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Length) and [Content-Disposition](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers/Content-Disposition).

### Dockerfile

- For this project the following were created:

    - one Dockerfile
    - one image
    - one container

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