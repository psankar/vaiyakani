#!/usr/bin/env node
var trie = {};
var line_no = 0;

var lineReader = require('line-reader');
var fs = require('fs');

//lineReader.eachLine('psankar.englishwords', function(line, last) {
lineReader.eachLine('ta-wikipedia.englishwords', function(line, last) {
	console.log(line_no);
	line_no ++;

	var words = line.toString().split(",");
	var score = words[0];
	var tamilWord = words[1];
	var englishWords = words.slice(2);

	for (var engIndex = 0, englishWord; englishWord = englishWords[engIndex]; engIndex++) {
		var letters = englishWord.split("");
		var active = trie;

		for (var j = 0; j < letters.length; j++ ) {
			var letter = letters[j];
			var pos = active[letter];

			if (pos == null) {
				active[letter] = {};
			}
			active = active[letter];
		}
		/* TODO: Add the score too */
		active.value = tamilWord;
	} /* End of English words loop */

	if(last) {
		fs.writeFile("trie.json", JSON.stringify(trie, undefined, 4));
	}
});
