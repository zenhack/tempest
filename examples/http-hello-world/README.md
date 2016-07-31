This is a simple example app that uses go.sandstorm's `websession`
wrapper and `net/http`.

You need to build the app build outside of the vagrant VM; on your local
machine, run:

    GOOS=linux GOARCH=amd64 go build

And then do the usual vagrant-spk dance. Given how easy cross compiling
Go apps is, it is often not worth it to build in the VM.
