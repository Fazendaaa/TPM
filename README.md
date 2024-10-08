# TPM

[Traefik](https://traefik.io/traefik/) Proxy Manager: alterative to [Nginx Proxy Manager](https://nginxproxymanager.com/) to easily expose already deployed systems

<div align = "center">
  <img src="./img/logo.svg" height=310>
</div>

- [TPM](#tpm)
  - [Idea](#idea)
  - [Usage](#usage)
    - [Running locally](#running-locally)
  - [TODO](#todo)
  - [Troubleshooting](#troubleshooting)
  - [References](#references)

## Idea

After deploying [CasaOS](https://casaos.zimaspace.com/) the idea to use **Nginx Proxy Manager (NPM)** to handle all my domains and certificates came to mind, even tho being a simple quest and having a plethora of tutorials about it at my disposal the challenge to do the same using Traefik seemed interesting. The perception was that it would be too difficult but it was very simple especially because Traefik's docs were impeccable; then the next logical step was to automate the whole process and generate a service to handle it all.

## Usage

In case that you are trying to deploy it to **CasaOS**, follow [this]() blogpost in which this is covered.

### Running locally

You will need to install [Docker](https://www.docker.com/) and [Docker-compose](https://docs.docker.com/compose/). Then:

1. Run:

```shell
docker-compose up --build
```

## TODO

- Improve support to handle different scenarios and providers
- Make it more beautiful
- Add stats

## Troubleshooting

- [Docker unable to bind to port 53](https://discourse.pi-hole.net/t/docker-unable-to-bind-to-port-53/45082/8)

## References

- [The 6 top Go web frameworks](https://blog.logrocket.com/6-top-go-web-frameworks/)
- [Developing Go Apps With Docker](https://www.docker.com/blog/developing-go-apps-docker/)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Is Your Container Image Really Distroless?](https://www.docker.com/blog/is-your-container-image-really-distroless/)
