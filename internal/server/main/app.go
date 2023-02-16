package servermain

import (
	"net/http"

	websessioncp "zenhack.net/go/tempest/capnp/web-session"
	"zenhack.net/go/tempest/pkg/exp/websession"
)

func ServeApp(
	webSession websessioncp.WebSession,
	w http.ResponseWriter,
	req *http.Request,
	rootHost string,
) {
	w.Header().Set(
		"Content-Security-Policy",
		// TODO(perf): refactor so we can call this once on startup,
		// and not have to reconstruct the string on every request:
		uiContentSecurityPolicy(
			req.URL.Scheme == "https",
			rootHost,
		),
	)

	websession.Handler{
		Session: webSession,
	}.ServeHTTP(w, req)
}

// uiContentSecurityPolicy returns a content security policy that disallows
// loading of remote resources. It accepts as arguments whether or not we're
// using https, and the root domain name.
//
// Note the following:
//
// - Currently there are still exceptions for images and media, as these have
//   some legitimate use cases (e.g. embedding images in feeds in ttrss) and
//   we want to provide a way for a user to allow these via the UI before we
//   block them by default
// - The unsafe-* directives are currently necessary to avoid breaking many
//   apps. They make CSP not particularly useful in mitating XSS attacks,
//   but do not present an information-leaking hazard.
// - In the future, we should provide a way for apps to opt-in to more
//   restrictive policies, as a useful mitigation for things like XSS vulns.
//   in the apps.
func uiContentSecurityPolicy(isSecure bool, rootHost string) string {
	const unsafe = "'unsafe-inline' 'unsafe-eval' data: blob:; "
	rootHttpHost := "http"
	wsHost := "ws"
	if isSecure {
		rootHttpHost += "s"
		wsHost += "s"
	}
	rootHttpHost += "://" + rootHost
	wsHost += "://" + rootHost
	return "default-src 'none'; " +
		"webrtc 'block'; " +

		// TODO: change these to 'self' when we're ready, per above.
		"img-src * " + unsafe +
		"media-src * " + unsafe +
		"script-src 'self' " + unsafe +
		"style-src 'self' " + unsafe +
		"child-src 'self' " + unsafe +
		"font-src 'self' " + unsafe +

		// frame-src needs to allow references to rootHost, because
		// we allow apps to pull the content of offer-iframes from
		// there:
		"frame-src 'self' " + rootHttpHost + " " + unsafe +

		// Service workers can intercept http requests and muck with
		// response headers, possibly overriding our security settings,
		// so we need to disable them.
		"worker-src 'none'; " +

		// 'self' alone does not allow websocket connections; see:
		// https://github.com/w3c/webappsec-csp/issues/7
		"connect-src 'self' " + wsHost + ";"
}
