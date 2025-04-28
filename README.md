# Postgirl

![example_screenshot](https://github.com/user-attachments/assets/ac9b319e-260c-45ad-8d7a-b48b02f1f7f4)

Postgirl is a simple API testing tool, designed to send HTTP requests easily from your terminal.

## Features

- Send Http `GET`,`POST`,`DELETE`,`PUT` and more method.
- Set query parameter, headers and request body
- Auto save request to a local cache file
- Automatically load previous request from cache file on startup

## Installation

### release page

You can visit to [releases page](https://github.com/insanXYZ/postgirl/releases), select the latest version and download it.

### build from source

```bash
git clone https://github.com/insanXYZ/postgirl
cd postgirl/
go build
chmod a+x ./postgirl
./postgirl
```

## FAQ

### How to copy text from text area(body, headers, params, response) ?

You can select/block the text and use combination ctrl + b for copy it.

### Why ctrl + b ?

Most terminals use ctrl + c as a signal to stop an application.
