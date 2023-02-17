This package is the entry point to Tempest's browser UI. This document
describes the high-level architecture.

The architecture is heavily inspired by [Elm][1]. In particular:

- There is a central `Model` type that defines the state of the
  application.
- There is a `Model.View()` method which returns an HTML rendering
  of the app based on the state, using [go-vdom][2].
- `message.go` defines message types that represent various events
  the UI state should respond to. Each message type has an `ApplyTo()`
  method that transforms the model. These messages are sent on a
  channel from within event handlers, and there is a loop that
  continuously receives from the channel, applies the message, and
  re-renders the UI. `ApplyTo()` must never block, as this could
  causes the UI to hang. It should always just update the model
  as it needs and spawn a goroutine for any side effects.

We use Cap'n Proto rpc over a websocket to communicate with the server.

[1]: https://elm-lang.org/
[2]: https://github.com/zenhack/go-vdom
