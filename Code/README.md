# Projekt

Check the other folders for further READMES.

## Systemd

In `/etc/systemd/system/` there is a file called `api.service` which will launch the server.

```
systemctl start api.service
```

To make the server persistent:

```
systemctl enable api.service
```

We might need envirionment variables later for the JWT, and these can be written in the service file

## Golang
The backend of the system is developed in go.

## Python
The frontend of the system is developed in python there is also another part of the frontend which is webbased and therefore not python.

### Makefile

We made a makefile which makes it easy to deploy the server directly after creating it `make deploy` 


