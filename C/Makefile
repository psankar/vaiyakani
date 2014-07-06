CFLAGS := $(shell pkg-config --cflags glib-2.0 gtk+-3.0 sqlite3) -g -ggdb -Wall -Werror -Wextra
LDFLAGS := $(shell pkg-config --libs glib-2.0 gtk+-3.0 sqlite3)

default: create-sqlite-db test-autocomplete-from-sqlite

create-sqlite-db: create-sqlite-db.c
	$(CC) $(CFLAGS) $< -o $@ $(LDFLAGS)

test-autocomplete-from-sqlite: test-autocomplete-from-sqlite.c
	$(CC) $(CFLAGS) $< -o $@ $(LDFLAGS)
