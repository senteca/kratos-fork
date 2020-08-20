---
id: milestones
title: Milestones and Roadmap
---

## [v0.7.0-alpha.1](https://github.com/ory/kratos/milestone/9)

*This milestone does not have a description.*

### [Enhancement](https://github.com/ory/kratos/labels/enhancement)

New feature or request

#### Issues

* [ ] Selfservice account deletion  ([kratos#596](https://github.com/ory/kratos/issues/596))
* [ ] Implement Hydra integration ([kratos#273](https://github.com/ory/kratos/issues/273))
* [ ] Self-service GDPR identity export ([kratos#658](https://github.com/ory/kratos/issues/658))
* [ ] Admin/Selfservice session management ([kratos#655](https://github.com/ory/kratos/issues/655))

## [v0.6.0-alpha.1](https://github.com/ory/kratos/milestone/8)

*This milestone does not have a description.*

### [Bug](https://github.com/ory/kratos/labels/bug)

Something isn't working

#### Issues

* [ ] Sending JSON to complete oidc/password strategy flows causes CSRF issues ([kratos#378](https://github.com/ory/kratos/issues/378))
* [ ] Unmable to use Auth0 as a generic OIDC provider ([kratos#609](https://github.com/ory/kratos/issues/609))
* [ ] Password reset emails sent twice by each of the two kratos pods in my cluster ([kratos#652](https://github.com/ory/kratos/issues/652))

### [Enhancement](https://github.com/ory/kratos/labels/enhancement)

New feature or request

#### Issues

* [ ] Implement Security Questions MFA ([kratos#469](https://github.com/ory/kratos/issues/469))
* [ ] Feature request: adjustable thresholds on how many times a password has been in a breach according to haveibeenpwned ([kratos#450](https://github.com/ory/kratos/issues/450))
* [ ] Do not send credentials to hooks ([kratos#77](https://github.com/ory/kratos/issues/77)) - [@hackerman](https://github.com/aeneasr)
* [ ] Implement immutable keyword in JSON Schema for Identity Traits ([kratos#117](https://github.com/ory/kratos/issues/117))
* [ ] Add filters to admin api ([kratos#249](https://github.com/ory/kratos/issues/249))
* [ ] Feature Request: Webhooks ([kratos#271](https://github.com/ory/kratos/issues/271))
* [ ] Support email verification paswordless login ([kratos#286](https://github.com/ory/kratos/issues/286))
* [ ] Support remote argon2 execution ([kratos#357](https://github.com/ory/kratos/issues/357)) - [@hackerman](https://github.com/aeneasr)
* [ ] Implement identity state and administrative deactivation, deletion of identities ([kratos#598](https://github.com/ory/kratos/issues/598)) - [@hackerman](https://github.com/aeneasr)
* [ ] SMTP Error spams the server logs ([kratos#402](https://github.com/ory/kratos/issues/402))
* [ ] Gracefully handle CSRF errors ([kratos#91](https://github.com/ory/kratos/issues/91)) - [@hackerman](https://github.com/aeneasr)
* [ ] How to sign in with Twitter ([kratos#517](https://github.com/ory/kratos/issues/517))
* [ ] Add ability to import user credentials ([kratos#605](https://github.com/ory/kratos/issues/605)) - [@hackerman](https://github.com/aeneasr)
* [ ] Throttling repeated login requests ([kratos#654](https://github.com/ory/kratos/issues/654))
* [ ] Require identity deactivation before administrative deletion ([kratos#657](https://github.com/ory/kratos/issues/657))

## [v0.5.0-alpha.1](https://github.com/ory/kratos/milestone/5)

This release focuses on Admin API capabilities

### [Bug](https://github.com/ory/kratos/labels/bug)

Something isn't working

#### Issues

* [ ] Logout does not use new cookie domain setting ([kratos#645](https://github.com/ory/kratos/issues/645))
* [ ] Refresh Sessions Without Having to Log In Again ([kratos#615](https://github.com/ory/kratos/issues/615)) - [@hackerman](https://github.com/aeneasr)
* [x] Generate a new UUID/token after every interaction ([kratos#236](https://github.com/ory/kratos/issues/236)) - [@hackerman](https://github.com/aeneasr)
* [x] UNIQUE constraint failure when updating identities via Admin API ([kratos#325](https://github.com/ory/kratos/issues/325)) - [@hackerman](https://github.com/aeneasr)
* [x] Can not update an identity using PUT /identities/{id} ([kratos#435](https://github.com/ory/kratos/issues/435))
* [x] Verification email is sent after password recovery ([kratos#578](https://github.com/ory/kratos/issues/578)) - [@hackerman](https://github.com/aeneasr)
* [x] Do not return expired sessions in `/sessions/whoami` ([kratos#611](https://github.com/ory/kratos/issues/611)) - [@hackerman](https://github.com/aeneasr)

### [Enhancement](https://github.com/ory/kratos/labels/enhancement)

New feature or request

#### Issues

* [x] Implement JSON capabilities in ErrorHandler ([kratos#61](https://github.com/ory/kratos/issues/61)) - [@hackerman](https://github.com/aeneasr)
* [x] Allow attaching credentials to identities in CRUD create ([kratos#200](https://github.com/ory/kratos/issues/200))
* [x] Move away from UUID-based challenges and responses ([kratos#241](https://github.com/ory/kratos/issues/241)) - [@hackerman](https://github.com/aeneasr)
* [x] Add tests to prevent duplicate migration files ([kratos#282](https://github.com/ory/kratos/issues/282)) - [@Patrik](https://github.com/zepatrik)
* [x] Session cookie (ory_kratos_session) expired time should be configurable ([kratos#326](https://github.com/ory/kratos/issues/326)) - [@hackerman](https://github.com/aeneasr)
* [x] Can not update an identity using PUT /identities/{id} ([kratos#435](https://github.com/ory/kratos/issues/435))
* [x] Make session cookie 'domain' property configurable ([kratos#516](https://github.com/ory/kratos/issues/516))
* [x] Remove one of in-memory/on-disk SQLite e2e runners and replace with faster test ([kratos#580](https://github.com/ory/kratos/issues/580)) - [@Andreas Bucksteeg](https://github.com/tricky42)
* [x] Password similarity policy is too strict ([kratos#581](https://github.com/ory/kratos/issues/581)) - [@Patrik](https://github.com/zepatrik)
* [x] Implement a test-error for implementing the Error UI ([kratos#610](https://github.com/ory/kratos/issues/610))

#### Pull Requests

* [ ] feat: implement API-based self-service flows ([kratos#624](https://github.com/ory/kratos/pull/624)) - [@hackerman](https://github.com/aeneasr)
* [x] fix: resolve identity admin api issues  ([kratos#586](https://github.com/ory/kratos/pull/586)) - [@hackerman](https://github.com/aeneasr)

### [Blocker](https://github.com/ory/kratos/labels/blocker)

#### Pull Requests

* [ ] feat: implement API-based self-service flows ([kratos#624](https://github.com/ory/kratos/pull/624)) - [@hackerman](https://github.com/aeneasr)