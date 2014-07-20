Use `make` to build all the binaries. Install glib2-devel gtk+-3-devel sqlite3-devel packages before you build.

* **create-db**: This will take a .englishwords file generated from the previous folder and will generate a .sqlite database.
* **test-autocomplete-from-sqlite**: This will autocomplete from the .sqlite database created from the previous step. The time taken for each keypress auto-completion will be printed in the terminal where you launch this.
* **test-autocomplete-from-inmem-sqlite**: This will create an in-memory db of the .sqlite database created by the create-db binary. Then auto-completion will be performed on this in-memory database. The time taken for each keypress auto-completion will be printed in the terminal for this binary also.


###Lessons Learnt:
* In my Thinkpad T430 with 8 GB RAM and QUAD core processor, both the programs took approximately same time, in the order of milli-seconds for ta-wiki.sqlite which has 8019817 records, with 3 columns.
* in-memory vs disk does not matter if you have sufficient RAM. The linux kernel page mapping and/or the sqlite caching layer work really well, that choosing in-memory may never be useful. No wonder mongodb chose to let the linux kernel take care of caching.
