// This is a WIP script to interact with the test app via the UI, for testing
// things that can't be accessed via api requests.
//
// Right now it just logs in as a dev account and opens the first grain on
// the list. We'll end up integrating this into the test suite and having it
// actually poke the right things soonish.

import Nightmare from 'nightmare';

const nightmare = Nightmare({
	// TODO: would be nice not to hard-code this, but on my(zenhack's) machine,
	// nightmare can't seem to find the binary from the npm package it depends
	// on.
	electronPath: "/usr/bin/electron",
})

// Monkey-patch nightmare to provide a wait then click combo; we use this a
// lot and I'm surprised it isn't already defined.
nightmare.waitclick = function(sel) {
	return this.wait(sel).click(sel)
}

nightmare
	.goto('http://local.sandstorm.io:6080')
	.waitclick('button.dev.login.expand')
	.waitclick('button[data-name="Alice Dev Admin"]')
	.waitclick('li.navitem-open-grain.current > a')
	.waitclick('.grain-list-table .grain-name > a')
	.end()
	.then(() => console.log('then'))
	.catch((err) => {
		console.error('catch:', err)
	})
