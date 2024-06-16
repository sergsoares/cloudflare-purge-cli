# Cloudflare Purge CLI

A single binary for Clean cache inside Cloudflare.

- Great for CICD pipeline usage after frontend deploy.
- Good for infrastructure runbook when things go wrong.
- Support with flexible environment variables.

# Usage

```
$ TOKEN=XXXX DOMAIN=YYYY cloudflare-purge-cli
2024/06/16 16:18:41 [loadOptions] Loading options...
2024/06/16 16:18:41 [purge] Purging cache...
2024/06/16 16:18:42 [purge] Cache cleaned
```