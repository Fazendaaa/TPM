# TPM

[Traefik](https://traefik.io/traefik/) Proxy Manager: alterative to [Nginx Proxy Manager](https://nginxproxymanager.com/) to easily expose already deployed systems

<div align = "center">
  <img src="./img/logo.svg" height=310>
</div>

- [TPM](#tpm)
  - [Background](#background)
    - [Idea](#idea)
    - [Opinionated project](#opinionated-project)
  - [Usage](#usage)
    - [Running locally](#running-locally)
  - [TODO](#todo)
  - [Troubleshooting](#troubleshooting)
  - [References](#references)

## Background

> This is not a production ready project! Please be careful in how and where you use it.

You **WON'T NEED** any internet access to generate certificates to last **THREE YEARS** in your local network since this project uses [mkcert](https://github.com/FiloSottile/mkcert) to handle that. The main idea is to provide a easy and automate service using a dashboard to handle all the minor details.

### Idea

After deploying [CasaOS](https://casaos.zimaspace.com/) the idea to use **Nginx Proxy Manager (NPM)** to handle all my domains and certificates came to mind, even tho being a simple quest and having a plethora of tutorials about it at my disposal the challenge to do the same using Traefik seemed interesting. The perception was that it would be too difficult but it was very simple especially because Traefik's docs were impeccable; then the next logical step was to automate the whole process and generate a service to handle it all.

### Opinionated project

- _"Why Go?"_ **Because I'm bad at it**
- _"Why Bootstrap?"_ **Because I never used before**
- _"Why PiHole?"_ **Because, at the moment, it is what I'm currently using**

## Usage

In case that you are trying to deploy it to **CasaOS**, follow [this]() blogpost in which this is covered.

### Running locally

You will need to install [Docker](https://www.docker.com/) and [Docker-compose](https://docs.docker.com/compose/). Then:

1. Run:

```shell
docker-compose up --build
```

2. Wait about 30s then open https://tpm.myhouse.localhost/
3. Accept the warning mentioning about the not valid certificate
4. Go to Certificate tab then install the generated certificate

## TODO

- Find a way to make install certificate button works
- Manage multiple domains -- currently the system supports only one but you can always deploy a dashboard to each one to "make it work"
- Add PiHole's API support to avoid using files -- this is a troublesome quest due to [lack of documentation](https://discourse.pi-hole.net/t/how-to-use-the-api/61004/9)
- Also add [Traefik's API](https://doc.traefik.io/traefik/operations/api/) support
- Add auto renew certificates
- Let's Encrypt support
- Improve support to handle different DNS Server providers
- Add docs to cover all macro use cases
- Make it more beautiful
- Add stats

## Troubleshooting

- [Docker unable to bind to port 53](https://discourse.pi-hole.net/t/docker-unable-to-bind-to-port-53/45082/8)

## References

- [The 6 top Go web frameworks](https://blog.logrocket.com/6-top-go-web-frameworks/)
- [Developing Go Apps With Docker](https://www.docker.com/blog/developing-go-apps-docker/)
- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Is Your Container Image Really Distroless?](https://www.docker.com/blog/is-your-container-image-really-distroless/)
- [golang-standards/project-layout](https://github.com/golang-standards/project-layout)
- [Bootstrap 5 Tutorial](https://www.w3schools.com/bootstrap5/index.php)
- [Golang HTML Template With Bootstrap Theme](https://www.geeksbeginner.com/golang-web-development-with-template-and-gin-framework/)
- [Docker Health Check: A Practical Guide](https://lumigo.io/container-monitoring/docker-health-check-a-practical-guide/)
- [How to convert a square SVG to all-size ICO?](https://graphicdesign.stackexchange.com/questions/77359/how-to-convert-a-square-svg-to-all-size-ico)
