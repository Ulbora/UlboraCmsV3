[![Go Report Card](https://goreportcard.com/badge/github.com/Ulbora/UlboraCmsV3)](https://goreportcard.com/report/github.com/Ulbora/UlboraCmsV3)
[![](https://img.shields.io/docker/build/mariobehling/loklak.svg)](https://hub.docker.com/r/ulboralabs/ulboracmsv3/builds/)


[UlboraCMS UI](http://www.ulboralabs.com/ulboracms)
==============

UlboraCMS UI is a framework for building dynamic web sites using [Ulbora Headless CMS](https://github.com/Ulbora/UlboraContentService).
Web content is saved in the cloud and cached when read.
UlboraCMS UI uses open source Ulbora Headless CMS for content. 
Ulbora Headless CMS on available on Docker.

# Easy Way To Get Started

## [Free Account](http://www.register.myapigateway.com/?ulboraCms=true)

Get a free acount and get started now. Then:

docker run -p 8090:80 --name yourdomin.com --log-opt max-size=50m --env GATEWAY_API_KEY=yourApiKey --env AUTH_CODE_CLIENT_ID=YourClientNumber --env AUTH_CODE_CLIENT_SECRET=YourClientSecret --env AUTH_REDIRECT_HOST=http://oauth2.myapigateway.com:3000 --env AUTH_HOST=http://oauth2.myapigateway.com:3000 --env CAPTCHA_SECRET=Your Captch Secret --env API_GATEWAY_HOST=http://gateway.myapigateway.com -v /static/templates --restart=always -d ulboralabs/ulboracmsv3 sh

# Hard Way To Get Started


## Build Your Own System

Building your own system gives you total control over your system, but you must host all the needed micro services.

## Micro Services Needed (All require MySql)

### [Ulbora Headless CMS](https://hub.docker.com/r/ulboralabs/content)
Ulbora Headless CMS for content

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
