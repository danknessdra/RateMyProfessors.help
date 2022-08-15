# RateMyProfessors.help
> one line summary: rate my professor helper to help find the best professors for your courses

> more detailed description

[design doc](https://docs.google.com/document/d/18EV5vSysP4g-dQlOz8RPAOX1CN95_Bku6ngtJE6O48w/edit#heading=h.ng2zz6cp2tz0)

# development dependencies

- [golang](https://go.dev/)
- [python](https://www.python.org/)
- [python pip](https://github.com/pypa/pip)

## OS dependencies 
### openbsd 

[openbsd ports tree](https://www.openbsd.org/faq/ports/ports.html)
```
pkg_add python3

# needs the ports tree to compile binaries
cd /usr/ports/devel/py-pip/
make install clean

cd /usr/ports/lang/go/
make install clean
```

### ubuntu 

> ubuntu's golang package is outdated, follow the golang doc instructions

[go lang install](https://tip.golang.org/doc/install)

```
sudo apt-get install python3 python3-pip 
```


## go dependencies


```
go get github.com/julienschmidt/httprouter
```

## python dependencies

```
pip install "requests==2.25.1" beautifulsoup4 RateMyProfessorAPI
```


# deployment 

## openbsd 

[openbsd ports tree](https://www.openbsd.org/faq/ports/ports.html)

```
ln -s '/usr/local/bin/python3' '/bin/python3'

# need the ports tree to install
cd /usr/ports/lang/go/
make install clean

```

# development workflow
1. create issue for problem working on
2. create branch based off issue(s)
3. create pull request, code review if necessary

# technology stack

__frontend__ : [ html, css , javascript ]
- [reactjs](https://reactjs.org/docs/)
- [babeljs](https://babeljs.io/docs/en)

__backend__ : [go lang](https://go.dev/learn/)
- [julienschmidt/httprouter](https://pkg.go.dev/github.com/julienschmidt/httprouter)

__webscraping__ : [python](https://docs.python.org/3/)
- [RateMyProfessorAPI](https://pypi.org/project/RateMyProfessorAPI/)
- [beautifulsoup](https://pypi.org/project/beautifulsoup4/)

__reverse proxy__ : [relayd](https://man.openbsd.org/relayd.8)

__tls__ : [ [acme-client](https://man.openbsd.org/acme-client.1) , [httpd](https://man.openbsd.org/httpd.8) ]

