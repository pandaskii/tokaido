package dockertmpl

var TokComposeYml = `# START Generated by Tokaido
version: "2"
services:
  unison:
    image: onnimonni/unison
    environment:
      - UNISON_DIR=/tokaido/site
      - UNISON_UID=1001
      - UNISON_GID=1001
    ports:
      - "5000"
    volumes:
      - /tokaido/site
  syslog:
    image: tokaido/syslog:latest
    volumes:
      - /tokaido/logs
  haproxy:
    user: "1005"
    image: tokaido/haproxy:latest
    ports:
      - "8080"
      - "8443"
    depends_on:
      - varnish
      - nginx
  varnish:
    user: "1004"
    image: tokaido/varnish:latest
    depends_on:
      - nginx
    volumes_from:
      - syslog
  nginx:
    user: "1002"
    image: tokaido/nginx:latest
    volumes_from:
      - unison
      - syslog
    depends_on:
      - fpm
    ports:
      - "8082"
  fpm:
    user: "1001"
    image: tokaido/fpm:latest
    working_dir: /tokaido/site/
    volumes_from:
      - unison
      - syslog
    depends_on:
      - syslog
    ports:
      - "9000"
    environment:
      PHP_DISPLAY_ERRORS: "yes"
  memcache:
    image: memcached:1.5-alpine
  mysql:
    image: mysql:5.7
    volumes_from:
      - syslog
    ports:
      - "3306"
    environment:
      MYSQL_DATABASE: tokaido
      MYSQL_USER: tokaido
      MYSQL_PASSWORD: tokaido
      MYSQL_ROOT_PASSWORD: tokaido
  drush:
    image: tokaido/drush:latest
    hostname: 'tokaido'
    ports:
      - "22"
    working_dir: /tokaido/site
    volumes_from:
      - unison
      - syslog
    environment:
      SSH_AUTH_SOCK: /ssh/auth/sock

`