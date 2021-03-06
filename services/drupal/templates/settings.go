package drupaltmpl

// SettingsD7Append - (Append) docroot/sites/default/settings.php for Drupal 7
var SettingsD7Append = []byte(`/*
  * Generated by Tokaido
  */

  if (file_exists(DRUPAL_ROOT . '/sites/default/settings.tok.php')) {
    include DRUPAL_ROOT . '/sites/default/settings.tok.php';
  }

  /*
  * END Generated by Tokaido
  */

`)

// SettingsD7Tok - docroot/sites/default/settings.tok.php for Drupal 7
func SettingsD7Tok(hashSalt, projectName string) []byte {
	return []byte(`<?php

    /**
     * @file
     * Configuration file for Tokaido local dev environments. Add this to .gitignore
     *
     * Check out https://tokaido.io/docs for help managing your Tokaido environment
     *
     * Generated by Tokaido
     */
  
     $databases['default']['default'] = [
      'host' => 'mysql',
      'database' => 'tokaido',
      'username' => 'tokaido',
      'password' => 'tokaido',
      'port' => 3306,
      'driver' => 'mysql',
      'prefix' => '',
    ];
  
    $conf['file_private_path'] = '/tokaido/site/private/default';
    $conf['file_temporary_path'] = '/tmp';
  
    $drupal_hash_salt = '` + hashSalt + `';

    $base_url = 'https://` + projectName + `.local.tokaido.io:5154';
        
    /*
    * END Generated by Tokaido
    */
  
  `)
}

// SettingsD8Append - (Append) docroot/sites/default/settings.php for Drupal 8
var SettingsD8Append = []byte(`/*
* Generated by Tokaido
*/

if (file_exists($app_root . '/' . $site_path . '/settings.tok.php')) {
  include $app_root . '/' . $site_path . '/settings.tok.php';
}

/*
* END Generated by Tokaido
*/

`)

// SettingsD8Tok - docroot/sites/default/settings.tok.php for Drupal 8
func SettingsD8Tok(hashSalt string) []byte {
	return []byte(`<?php

/**
 * @file
 * Configuration file for Tokaido local dev environments. Add this to .gitignore
 *
 * Check out https://tokaido.io/docs for help managing your Tokaido environment
 *
 * Generated by Tokaido
 */

$databases['default']['default'] = [
  'host' => 'mysql',
  'database' => 'tokaido',
  'username' => 'tokaido',
  'password' => 'tokaido',
  'port' => 3306,
  'driver' => 'mysql',
  'namespace' => 'Drupal\\Core\\Database\\Driver\\mysql',
  'prefix' => '',
];

$settings['file_private_path'] = '/tokaido/site/private/default/';
$settings['file_temporary_path'] = '/tmp';

$settings['hash_salt'] = '` + hashSalt + `';

if (\Drupal::hasService('cache.backend.memcache')) {
  $settings['cache']['default'] = 'cache.backend.memcache';
  $settings['memcache_storage']['key_prefix'] = '';
  $settings['memcache_storage']['memcached_servers'] = ['memcache:11211' => 'default'];
}

/*
* END Generated by Tokaido
*/
`)
}
