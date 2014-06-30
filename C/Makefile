CFLAGS := $(shell pkg-config --cflags glib-2.0 gtk+-2.0 sqlite3) -g -ggdb -Wall -Werror -Wextra
LDFLAGS := $(shell pkg-config --libs glib-2.0 gtk+-2.0 sqlite3)

create-db: create-db.c
	$(CC) $(CFLAGS) $< -o $@ $(LDFLAGS)
