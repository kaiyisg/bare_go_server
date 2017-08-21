# REPO NAME

## Preconditions
A `secure` instance must be running to connect to.

## Setup
Your go path should already be setup as per the engineering onboarding documents.

**Clone this repo into your go path**

```bash
mkdir -p ~/go/src/github.com/ORGANIZATION/
cd ~/go/src/github.com/ORGANIZATION/
git clone git@github.com:ORGANIZATION/bare_go_server.git
```

**Install the dependencies**
This repo uses the glide dependency manager to manage dependencies. To install the dependencies:

```bash
glide install
```

## Development

**Run the tests**

**Run the server**

*In `~/go/src/github.com/ORGANIZATION/bare_go_server*

```bash
go run main.go
```

**Adding new golang libraries**

*In `~/go/src/github.com/ORGANIZATION/bare_go_server*

```bash
glide get github.com/AUTHOR/REPO
```
