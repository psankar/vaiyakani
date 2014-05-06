chrome.extension.onMessage.addListener(
	function(request, sender, sendResponse) {
	chrome.pageAction.show(sender.tab.id);
	sendResponse();
});

chrome.browserAction.onClicked.addListener(
	function(newTab) {
	chrome.tabs.create ( {url: chrome.extension.getURL('src/editor/editor.html') });
});
