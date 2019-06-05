## Getting Started

### what is NGINX
------------------
NGINX is open source software for web serving, reverse proxying, caching, load balancing, media streaming, and more. It started out as a web server designed for maximum performance and stability. In addition to its HTTP server capabilities, NGINX can also function as a proxy server for email (IMAP, POP3, and SMTP) and a reverse proxy and load balancer for HTTP, TCP, and UDP servers.

### Beginners Guide

This guide describes how to start and stop nginx, and reload its configuration, explains the structure of the configuration file and describes how to set up nginx to serve out static content, how to configure nginx as a proxy server, and how to connect it with a FastCGI application. **nginx** has one master process and several worker processes. The main purpose of the master process is to read and evaluate configuration, and maintain worker processes. Worker processes do actual processing of requests. nginx employs event-based model and OS-dependent mechanisms to efficiently distribute requests among worker processes. The number of worker processes is defined in the configuration file and may be fixed for a given configuration or automatically adjusted to the number of available CPU cores. The way **nginx** and its modules work is determined in the configuration file. By default, the configuration file is named `nginx.conf` and placed in the directory `/usr/local/nginx/conf`, `/etc/nginx`, or /`usr/local/etc/nginx`.

### Starting, Stopping, and Reloading Configuration
----------------------------------------
To start nginx, run the executable file. In ubuntu it can be as wasy as typing nginx in command line

 Once nginx is started, it can be controlled by invoking the executable with the -s parameter. Use the following syntax:
 > ``` nginx -s signal```

 Where signal may be one of the following:
 1. **stop** — fast shutdown
 2. **quit** — graceful shutdown
 3. **reload** — reloading the configuration file
 4. **reopen** — reopening the log files

For getting the list of all running nginx processes, the ps utility may be used, for example, in the following way

> ```ps -ax | grep nginx```

### Deeper Look at nginx.conf
-------------------
nginx consists of modules which are controlled by directives specified in the configuration file. Directives are divided into simple directives and block directives. A simple directive consists of the name and parameters separated by spaces and ends with a semicolon (;). A block directive has the same structure as a simple directive, but instead of the semicolon it ends with a set of additional instructions surrounded by braces ({ and }). If a block directive can have other directives inside braces, it is called a context (examples: events, http, server, and location).
Directives placed in the configuration file outside of any contexts are considered to be in the main context. The events and http directives reside in the main context, server in http, and location in server.
The rest of a line after the # sign is considered a comment.

Example:
```
user www-data;
worker_processes 4;
pid /var/run/nginx.pid;

events {
        worker_connections 768;
        # multi_accept on;
}

http {

        ##
        # Basic Settings
        ##

        sendfile on;
        tcp_nopush on;
        tcp_nodelay on;
        keepalive_timeout 65;
        types_hash_max_size 2048;
        # server_tokens off;

        # server_names_hash_bucket_size 64;
        # server_name_in_redirect off;

        include /etc/nginx/mime.types;
        default_type application/octet-stream;

        ##
        # Logging Settings
        ##

        access_log /var/log/nginx/access.log;
        error_log /var/log/nginx/error.log;

        ##
        # Gzip Settings
        ##

        gzip on;
        gzip_disable "msie6";

        # gzip_vary on;
        # gzip_proxied any;
        # gzip_comp_level 6;
        # gzip_buffers 16 8k;
        # gzip_http_version 1.1;
        # gzip_types text/plain text/css application/json application/x-javascript text/xml application/xml application/xml+rss text/javascript;

        ##
        # nginx-naxsi config
        ##
        # Uncomment it if you installed nginx-naxsi
        ##

        #include /etc/nginx/naxsi_core.rules;

        ##
        # nginx-passenger config
        ##
        # Uncomment it if you installed nginx-passenger
        ##

        #passenger_root /usr;
        #passenger_ruby /usr/bin/ruby;

        ##
        # Virtual Host Configs
        ##

        include /etc/nginx/conf.d/*.conf;
        include /etc/nginx/sites-enabled/*;
}


#mail {
#       # See sample authentication script at:
#       # http://wiki.nginx.org/ImapAuthenticateWithApachePhpScript
#
#       # auth_http localhost/auth.php;
#       # pop3_capabilities "TOP" "USER";
#       # imap_capabilities "IMAP4rev1" "UIDPLUS";
#
#       server {
#               listen     localhost:110;
#               protocol   pop3;
#               proxy      on;
#       }
#
#       server {
#               listen     localhost:143;
#               protocol   imap;
#               proxy      on;
#       }
#}
```
a `ps` comman to list nginx process show. With One master and 4 worker process as configured in nginx.conf file.

```
root@precise64:~# ps -ax | grep nginx
Warning: bad ps syntax, perhaps a bogus '-'? See http://procps.sf.net/faq.html
  1956 ?        Ss     0:00 nginx: master process nginx
  1957 ?        S      0:00 nginx: worker process
  1958 ?        S      0:00 nginx: worker process
  1959 ?        S      0:00 nginx: worker process
  1960 ?        S      0:00 nginx: worker process
  1967 pts/0    S+     0:00 grep --color=auto nginx
```
### Serving Static Content
--------------------------------------
Project step
Files will be served from different local directories: /data/www (which may contain HTML files) and /data/images (containing images)
1. create the /data/www directory and put an index.html
2. create the /data/images directory and place some images in it
3. configuration file

  ```
  user www-data;
  worker_processes 4;
  pid /var/run/nginx.pid;

  events {
          worker_connections 768;
          # multi_accept on;
  }

  http {
      server {
      location / {
          root /data/www;
      }

      location /images/ {
          root /data;
      }
      }

  }
  ```

  - Generally, the configuration file may include several server blocks distinguished by ports on which they listen to and by server names. Once nginx decides which server processes a request, it tests the URI specified in the request’s header against the parameters of the location directives defined inside the server block. This location block specifies the “/” prefix compared with the URI from the request. For matching requests, the URI will be added to the path specified in the root directive, that is, to /data/www, to form the path to the requested file on the local file system.

  This is already a working configuration of a server that listens on the standard port 80 and is accessible on the local machine at http://localhost/. In response to requests with URIs starting with /images/, the server will send files from the /data/images directory.
4. `nginx -s reload`

running Curl on localhost port 80 show.
```
root@precise64:/data/images# curl http://localhost:80/
<html>
<header><title>This is title</title></header>
<body>
Hello world
</body>
</html>
```

### Reference

1. [nginx](https://nginx.org/en/docs/beginners_guide.html)
