![Iris from scratch](https://guides.nanobox.io/assets/quickstart-icons/iris.png)

# Iris from scratch

Run a Iris app locally, install nothing besides nanobox.

<a href="https://nanobox.io/download"><img src="https://guides.nanobox.io/assets/quickstart-icons/download.png" /></a>

## Clone the repo

```bash
# clone the code
git clone https://github.com/nanobox-quickstarts/nanobox-iris.git

# cd into the iris app
cd nanobox-iris
```

## Run the app

```bash
# Run Iris as you would normally, with Nanobox
nanobox run go run main.go
```

## Check it out

```bash
# Add a convenient way to access your app from the browser
nanobox dns add local iris.dev
```

Visit your app at <a href="http://iris.dev:8080" target="\_blank">iris.dev:8080</a>

## Explore
With Nanobox, you don't have to have anything installed on your machine to run your app:

```bash
# drop into a Nanobox console
nanobox run

# where golang is installed,
go version

# your packages are available,
git -v

# and your code is mounted
ls
```

## Now What?
For more details about running Iris apps with nanobox visit [guides.nanobox.io/golang/iris/](https://guides.nanobox.io/golang/iris/)

<a href="https://nanobox.io"><img src="https://guides.nanobox.io/assets/quickstart-icons/footer.png" /></a>
