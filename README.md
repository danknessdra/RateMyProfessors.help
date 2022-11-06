# RateMyProfessors.help
> The RateMyProfessor helper to help find the best professors for your courses

[RateMyProfessors.help](https://ratemyprofessors.help)

Finds the best teacher for your courses using [RateMyProfessors.com](https://ratemyprofessors.com) to rank the best teacher for your specified course.

# development dependencies

- [golang 1.17-7](https://go.dev/)
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
pip install "requests==2.25.1" beautifulsoup4 RateMyProfessorAPI psycopg2 python-dotenv
```


# deployment 

## linux (docker) 

1. install docker
2. setup docker-compose.yml "./examples/docker-compose.yml"
3. docker-compose up -d


## openbsd 

[openbsd ports tree](https://www.openbsd.org/faq/ports/ports.html)

```
# need the ports tree to install
cd /usr/ports/lang/go/
make install clean

git clone <repository>
cd <repository>
go build
./webserver
```

for setting up tls with httpd, relayd, and acme-client view ./examples

### security practices

create a service user and service group with own dir in /var (`_rmp`)
have rc.d service run as service user

for ssl support, create ssl-cert group chown keys/certs
add service user to group

for postgresql, follow pkg-readme for best practice

### SEE ALSO

```
httpd(8), httpd.conf(5), acme-client(1), relayd(8), relayd.conf(5)
```

### resources

[generating certs](https://libredd.it/r/openbsd/comments/wotu2v/solved_unable_to_generate_a_crt_with_acmeclient/ikgtwuv/?context=3)

[relayd](https://gist.github.com/anon987654321/4532cf8d6c59c1f43ec8973faa031103)

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

__ci/cd__: [jenkins](jenkins.io)

test commit
