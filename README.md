[![made-with-Go](https://img.shields.io/badge/Made%20with-Go-orange)](http://golang.org) [![proc-arch](https://img.shields.io/badge/Arch-x86%20%7C%20AMD64%20%7C%20ARM5%20%7C%20ARM7-blue)](http://golang.org) [![os](https://img.shields.io/badge/OS-Linux%20%7C%20Windows%20%7C%20Darwin-yellowgreen)](http://golang.org)


# Web interface for sending Wake-on-lan (magic packet)

A GoLang based HTTP server which will send a Wake-on-lan package (magic packet) on local network. The request can be send using web interface or directly using HTTP request with mapped device name in the URL. You can bookmark direct link to device(s) on your browsers to wake them using single HTTP call for ease of access.

## This is a fork of wolweb for Kubernetes
* Removes Vdir Support because I want it to run at index all times.
* Shuffled files around for easier kubernetes setup.
* Does not fail when config.json and devices.json files does not exists.

### Will Do
* Fix UI Bug on Editing an entry.

## Bootstrap UI with JS Grid for editing data

![Screenshot](wolweb_ui.png)

The UI features CRUD operation implemented using [js-grid.com](https://github.com/tabalinas/jsgrid) plugin. 

### Wake-up directly using HTTP Request

/wake/**&lt;hostname&gt;** -  Returns a JSON object

```json
{
  "success":true,
  "message":"Sent magic packet to device Server with Mac 34:E6:D7:33:12:71 on Broadcast IP 192.168.1.255:9",
  "error":null
}
```

## Configure the app

The application will use the following default values if they are not explicitly configured as explained in sections below.

| Config                | Description                                                                               | Default             | Variable Name |
| --------------------- | ----------------------------------------------------------------------------------------- | ------------------- | ------------- |
| Port                  | Define the port on which the webserver will listen                                        | **8089**            | WOLWEBPORT    |

You can override the default application configuration by setting environment variables. 

## Using with Docker Container

This project includes [Dockerfile (based on Alpine)](./Dockerfile) and [docker-compose.yml](./docker-compose.yml) files which you can use to build the image for your platform and run it using the docker compose file.

**Build Docker Image:**

```
docker build -t wolweb .
```
**Run Docker Image:**

```
docker-compose up -d
```

> I could not get this to run using Docker's bridged network. The only way I was able to make it work was to use host network for the docker container. See this [https://github.com/docker/for-linux/issues/637](https://github.com/docker/for-linux/issues/637) for details.

## Build on Windows
(Original Author not me)
I use VS Code with Go extension. To build this project on windows
```
go build -o wolweb.exe .
```

## Build for ASUS Routers (ARM v5)
(Original Author not me)
I initially thought of running this application on my router, so I needed to build the application without having to install build tool on my router. I use the following **PowerShell** one liner to build targeting the ARM v5 platform on my Windows machine with VS Code:
```powershell
 $Env:GOOS = "linux"; $Env:GOARCH = "arm"; $Env:GOARM = "5"; go build -o wolweb .
```
Copy the file over to router and make it executable.
```sh
chmod +x wolweb
```

To see detailed instructions on how to run this application as service on ASUS router with custom firmware [asuswrt-merlin](https://www.asuswrt-merlin.net/) see this [Wiki guide](https://github.com/alchemistake/wolweb/wiki/Run-on-asuswrt-merlin)

## Credits
Thank you to David Baumann's project https://github.com/dabondi/go-rest-wol for providing the framework which I modified a little to work within constraints of environment.
