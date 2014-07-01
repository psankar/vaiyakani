CFLAGS := $(shell pkg-config --cflags glib-2.0 gtk+-3.0 sqlite3) -g -ggdb -Wall -Werror -Wextra
LDFLAGS := $(shell pkg-config --libs glib-2.0 gtk+-3.0 sqlite3)

default: create-db test-autocomplete

create-db: create-db.c
	$(CC) $(CFLAGS) $< -o $@ $(LDFLAGS)

test-autocomplete: test-autocomplete.c
	$(CC) $(CFLAGS) $< -o $@ $(LDFLAGS)
