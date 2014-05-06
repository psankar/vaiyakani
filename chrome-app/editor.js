var chooseFileButton = document.querySelector('#chooseFile');
var fileDetailsDiv = document.querySelector('#fileDetails');
var textarea = document.querySelector('textarea');

chooseFileButton.addEventListener('click', function(e) {
	var accepts = [{
		mimeTypes: ['text/*'],
		extensions: ['js', 'css', 'txt', 'html', 'xml', 'tsv', 'csv', 'rtf']
	}];
	chrome.fileSystem.chooseEntry({type: 'openFile', accepts: accepts}, function(theEntry) {
		if (!theEntry) {
			return;
		}
		loadFileEntry(theEntry);
	});
});

function loadFileEntry(_chosenEntry) {
  chosenEntry = _chosenEntry;
  chosenEntry.file(function(file) {
    readAsText(chosenEntry, function(result) {
      textarea.value = result;
      textarea.hidden = false;
    });
    // Update display.
    fileDetailsDiv.hidden = false;
    displayEntryData(chosenEntry);
  });
}

function readAsText(fileEntry, callback) {
  fileEntry.file(function(file) {
    var reader = new FileReader();

    reader.onerror = errorHandler;
    reader.onload = function(e) {
      callback(e.target.result);
    };

    reader.readAsText(file);
  });
}

function displayEntryData(theEntry) {
  if (theEntry.isFile) {
    chrome.fileSystem.getDisplayPath(theEntry, function(path) {
      document.querySelector('#filePath').textContent = path;
    });
    theEntry.getMetadata(function(data) {
      document.querySelector('#fileSize').textContent = data.size;
    });
  }
}

function errorHandler(e) {
  console.error(e);
}


