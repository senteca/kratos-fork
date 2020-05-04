---
id: selfservice-flow-completion
title: Self-Service Flow Completion
---

Self-Service flows such as Login, Registration, Updating Settings support two
successful response modes:

- For browsers, the response will be a [redirection](#redirection).
- For API clients (this includes AJAX) the response will be in [JSON](#json).

## Redirection

Browser requests, identified by the `Accept: text/html` header, complete with a
redirection flow. If no redirection URL is set for the flow, the Default
Redirect URL will be used:

```yaml file="path/to/my/kratos.config.yml"
urls:
  default_redirect_to: https://always-end-up-here-per-default/
```

It is possible to specify a redirect URL per Self-Service Flow:

```yaml file="path/to/my/kratos.config.yml"
selfservice:
  login:
    after:
      default_redirect_to: https://end-up-here-after-login/
  registration:
    after:
      default_redirect_to: https://end-up-here-after-registration/
  # ...
```

You may also set redirect URLs per strategy (overrides
`selfservice.<login|registration|...>.default_return_to`):

```yaml file="path/to/my/kratos.config.yml"
selfservice:
  login:
    after:
      default_redirect_to: https://this-is-overridden-by-password/
      password:
        default_redirect_to: https://end-up-here-after-login-with-password/
    # ...
```

It is also possible to redirect someone back to the original URL. For example,
if a user requests `https://www.myapp.com/blog/write` but is not logged in, we
want the user to end up at that page after login. To achieve that, you append
`?return_to=https://www.myapp.com/blog/write` when initializing the Login /
Registration /Settings flow.

Because ORY Kratos prevents Open Redirect Attacks, you need to whitelist the
domain in your ORY Kratos config:

```yaml file="path/to/my/kratos.config.yml"
urls:
  whitelisted_return_to_urls:
    - https://www.myapp.com/
```

## JSON

This feature is currently in prototype phase and will be documented at a later
stage.
