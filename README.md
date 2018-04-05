[![Go Report Card](https://goreportcard.com/badge/github.com/Ulbora/UlboraCmsV3)](https://goreportcard.com/report/github.com/Ulbora/UlboraCmsV3)
[![](https://img.shields.io/docker/build/mariobehling/loklak.svg)](https://hub.docker.com/r/ulboralabs/ulboracmsv3/builds/)
[UlboraCMS V3](http://www.ulboralabs.com/ulboracms)
[![Maintainability](https://api.codeclimate.com/v1/badges/39d19772f1ac0b6e2a9b/maintainability)](https://codeclimate.com/github/Ulbora/UlboraCmsV3/maintainability)
==============

UlboraCMS V3 is a framework for building dynamic web sites with Go.
Web content is saved in the cloud and cached when read.
UlboraCMS V3 uses open source micro services that provide the functionality needed. 
All micro services are available on Docker.

## Micro Services Needed (All require MySql)

### [Content Service](https://hub.docker.com/r/ulboralabs/content)
Micro service for content

### [Template Service](https://hub.docker.com/r/ulboralabs/templates)
Micro service for UlboraCMS v3 templates

### [Mail Service](https://hub.docker.com/r/ulboralabs/mail)
Mail server and mail sending micro service

### [Image Service](https://hub.docker.com/r/ulboralabs/images)
Image micro service

## Security 
UlboraCMS V3 uses enhanced JWT tokes for security. The following is needed for enhanced JWT:

### [OAuth2 Server](https://hub.docker.com/r/ulboralabs/oauth2server)
OAuth2 server 

### [User Service](https://hub.docker.com/r/ulboralabs/userservice)
User micro service





## Docker Container

The easiest way to get Ulbora CMS V3 is to use Docker [here](https://hub.docker.com/r/ulboralabs/ulboracmsv3/).


## Templates

You can build multiple templates with UlboraCMS V3 and switch between the templates through the administration screen. 

New templates can easily be installed by zipping them as tar.gz files and following a simple upload process in the administration screens.



[MIT](LICENSE)
