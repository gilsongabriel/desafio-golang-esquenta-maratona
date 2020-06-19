# Desafio Golang - Esquenta 3ª Maratona FullCycle
## [Code Education](https://code.education) / [School Of Net](https://schoolofnet.com)

### Docker Hub
[gilsongabriel/desafio-golang-esquenta-maratona-fullcycle](https://hub.docker.com/r/gilsongabriel/desafio-golang-esquenta-maratona-fullcycle)

#### Porta web 8087

#### Docker Compose

```
version: "3.8"
services:
  app:
    image: gilsongabriel/desafio-golang-esquenta-maratona-fullcycle
    container_name: app_name
    ports:
    - 8087:8087
    restart: always
    tty: true
```