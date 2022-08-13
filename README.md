# prof course finder

name not finalized

[design doc](https://docs.google.com/document/d/18EV5vSysP4g-dQlOz8RPAOX1CN95_Bku6ngtJE6O48w/edit#heading=h.ng2zz6cp2tz0)

# git workflow
1. create issue for problem working on
2. create branch based off issue(s)
3. create pull request, code review if necessary

# resources 

frontend: [ html, css, javascript ]

backend: [go lang](https://go.dev/learn/)

webscraping: [python](https://docs.python.org/3/)

reverse proxy: [relayd](https://man.openbsd.org/relayd.8)

# dependencies

- golang
- python
- python-pip3

## go dependencies

```
go get github.com/julienschmidt/httprouter
```

[http router docs](https://pkg.go.dev/github.com/julienschmidt/httprouter)

## python dependencies

```
pip install "requests==2.25.1" beautifulsoup4 RateMyProfessorAPI
```

# deployment 
## openbsd additions
```
ln -s '/usr/local/bin/python3' '/bin/python3'
```

