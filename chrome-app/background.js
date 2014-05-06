chrome.app.runtime.onLaunched.addListener(function() {
  chrome.app.window.create('window.html', {
    'bounds': {
      'width': window.screen.availWidth,
      'height': window.screen.availHeight
    }
  });
});
