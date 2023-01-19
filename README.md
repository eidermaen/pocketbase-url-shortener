# Pocketbase URL shortener

> Ready-to-use url shortener built with [Go](https://go.dev/) and [Pocketbase.io](https://pocketbase.io/)


**Disclaimer: As of the commit day, Pocketbase is still under development! You might
check out their docs for any changes!**

---

## Setup

In order to use the url shortener, do the following:

### Step 1: Create collection

First, you have to create the needed collection

- Checkout project
- Run `go mod download`
- Run `go run main.go serve`
- Head over to `localhost:8090/_` and create an account
- Login with the used credentials
- Create a new collection named `link` with the following fields
  - slug: type text [nonempty, unique]
  - url: type url [nonempty]
  - clicks: type number [min: 0]


### Step 2: Add new entries

Once the collection is created, you can start adding new links, like
- slug: `gh`
- url: `https://github.com`
- clicks: `0`

### Step 3: Use it!

Now you can open `http://localhost:8090/gh` in your browser and should be redirected to Github.

## Features

- SQlite powered database
- Usage of used short urls will be tracked (counter which get's increased every time the short url is being accessed)
- Ready-to-use Dockerfile for deploying the app anywhere you want
