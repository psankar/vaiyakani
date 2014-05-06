chrome.extension.sendMessage({}, function(response) {
	var readyStateCheckInterval = setInterval(function() {
	if (document.readyState === "complete") {
		clearInterval(readyStateCheckInterval);

		// ----------------------------------------------------------
		// This part of the script triggers when page is done loading.
		// We can use this implement vaiyakani auto-complete support for the
		// text input boxes in the html pages.
		console.log("Hello. This message was sent from vaiyakani/scripts/inject.js");
		// ----------------------------------------------------------

	}
	}, 10);
});
