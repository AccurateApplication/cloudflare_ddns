## Dynamic DNS with go
This adds my external IP as a DNS record at cloudflare.
It adds the 'subdomain' in the config file as a subdomain record
It uses ipify.org to get external IP.

## Config
Config file is used for domain, subdomain, refresh rate (in minutes), and cloudflare email.
API key is stored as an environment variable. (`export CF_API_KEY='key'`).
