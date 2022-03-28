# Migrating to sqlc-gen-python
 
Starting in sqlc 1.16.0, built-in Python support has been deprecated. It will
be fully removed in 1.17.0 in favor of sqlc-gen-python.

This guide will walk you through migrating to the [sqlc-gen-python](https://github.com/tabbed/sqlc-gen-python) plugin,
which involves three steps.

1. Add the sqlc-gen-python plugin
2. Migrate each package
3. Re-generate the code

## Add the sqlc-gen-python plugin

In your configuration file, add a `plugins` array if you don't have one
already. Add the following configuration for the plugin:

```json
{
  "version": "2",
  "plugins": [
    {
      "name": "py",
      "wasm": {
        "url": "https://downloads.sqlc.dev/plugin/sqlc-gen-python_1.0.0.wasm",
        "sha256": "aca83e1f59f8ffdc604774c2f6f9eb321a2b23e07dc83fc12289d25305fa065b"
      }
    }
  ]
}
```

```yaml
version: "2"
plugins:
  - name: "py"
    wasm:
      url: "https://downloads.sqlc.dev/plugin/sqlc-gen-python_1.0.0.wasm"
      sha256: "aca83e1f59f8ffd