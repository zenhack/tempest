This document contains a high-level roadmap for tempest development.
It is mostly unordered; a lot of stuff can in principle proceed in
parallel.

# App platform/APIs

Things that interact with running apps that still need implementing

## Sandbox

A few pieces of the sandbox still need to be added to reach parity with
Sandstorm's isolation:

- [x] Load a seccomp-bpf filter in the sandbox launcher. We can just
  copy the BPF assembly directly from Sandstorm. This will give us
  isolation equivalent to Sandstorm with `USE_EXPERIMENTAL_SECCOMP_FILTER=true`;
  this is supposed to displace the old filter eventually anyway.
- [x] Set max file limits in sandbox setup (`setrlimit(RLIMIT_NOFILE, ...)`).
- [x] Set CSP and any other security headers on HTTP responses.

See also `BLOCKERS` at the root of the repository for a checklist of things that
need to  be done before anyone should rely on Tempest for security.

## Platform API Features

These are things that are not related to isolation of apps, but are
necessary for the functioning of some existing apps.

- [ ] Pass all necessary information through to grains when spawning a
  session.
  - [ ] Permissions
  - [ ] audit the implementation to make sure there's nothing else.
- [ ] Implement methods of `SandstormApi` and `SessionContext` from
  `grain.capnp`; right now all of these just throw "unimplemented"
  exceptions.
- [ ] Implement the postMessage APIs provided on the client side.
  - [ ] Powerbox requests
  - [ ] Offer iframes
  - [ ] Calls for mucking with the title/setting the URL bar
  - [ ] any others? Go through sandstorm docs/implementation to generate
    a complete list, which can replace this.

# User Interface

Things to do in the user interface. This section is mostly high-level
ideas, but in addition to what's here, as we add features those will
of course need corresponding interfaces to be built.

## A11Y

Tempest should be accessible from the start; Sandstorm was reportedly
pretty good on this, we should not regress.

Some of this comes down to following standard best-practices for
authoring HTML, but we should also devise a strategy for
testing/validating this, rather than just waiting for bug reports from
potential users who have a11y needs.

## I18n

Tempest should be internationalized, ideally from early on. Exactly what
this looks like requires a bit of research.

## Styling

- [ ] We should pick out pallets to use for color, size, and fonts early
  on. There is already a little bit of this in style.css, but it will
  need further definition.
- [ ] We should think about how to do dark mode support early on.

# User Features

Things users should be able to do, other than just interact with apps

- [ ] Install apps
- [ ] Create/delete grains
- [ ] Sharing stuff

# Administrative Features

Features used by admins

- [ ] HTTPS setup
- [ ] User management
