## 1.2.0

- Tokaido now uses a globally-resolvable *.local.tokaido.io domain. Manipulation of your local hosts file is no longer necessary. Thanks @garrettw! 
- Version-locked Yamamnote to 1.0.0
- settings.tok.php no longer includes a ?> php closing tag
- settings.tok.php now includes a random Drupal hash salt
- settings.tok.php now includes a fixed path for default's sites files
- settings.tok.php now includes memcache settings for Drupal 8 sites
- Tokaido will warn Linux users if their max_user_watches setting is probably too low



## 1.1.2

- Bugfix: Fixed a bug where the proxy network configuration was duplicated with each `tok up` run

## 1.1.1

- Bugfix: Dots in project names won't cause docker-compose networks to fail (#34)

## 1.1.0 

- Tokaido can now be run from subdirectory inside a Git repo.
- Bugfix (linux): Tok will now create the systemd path if it doesn't exist
- Bugfix: Set the base_url for Drupal7 sites to the proxy (tokaido.local) URL, so assets load properly
- Moved xdebug to it's own image. When not enabled, Tokaido will be nearly twice as fast
- Prevent an unhandled exception when a container has not started
- Temporary workaround for installing the unison-fsmonitor utility via pip, while homebrew is broken (#29)
- SSH agent now uses identity file only, fixing a bug where the user has more than 5 keys in ssh agent

## 1.0.0

- Remove version and customcompose as persistent global flags
- Audit of command short and long descriptions
- Fix missing Drupal version detection
- Tokaido will now prompt for a Drupal version and path if it can't detect one automatically
- Add the /private and sites/\*/files folders to gitignore automatically
- Add Adminer service (enable with: `tok config-set service adminer enabled true`)
- Add Redis service (enable with: `tok config-set service redis enabled true`)
- Add MailHog support (enable with: `tok config-set service mailhog enabled true`)
- 'Open' can now be used to open secondary services like mailhog, adminer (eg `tok open adminer`)
- 'ports' now dynamically returns ports only for services which are enabled
- Add dependency checks for Linux
- MySQL max_allowed_packet is now set to 64MB by default. 
- Fixed a bug with Compose file used 'cmd' instead of 'command' parameters. 
- Add the Tokaido Proxy service and HTTPS certificate support

