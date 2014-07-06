#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include <gtk/gtk.h>

#include <sqlite3.h>

GtkWidget *suggestions;
GtkWidget *entry;

sqlite3 *db;
sqlite3_stmt *stmt;

void kb()
{
	int ret;

	/* Ideally these two should be allocated dynamically,
	 * but for the time being, stack is enough */
	char results[1024] = "";
	char search_string[24] = "";

	const gchar *prefix = gtk_entry_get_text(GTK_ENTRY(entry));
	if (strlen(prefix) < 1)
		return;
	strncpy(search_string, (char *)prefix, 23);
	strcat(search_string, "%%");
	fprintf(stderr, "Prefix is: [%s]\tString searched is: [%s]\n", prefix,
		search_string);

	ret = sqlite3_bind_text(stmt, 1, search_string, -1, SQLITE_STATIC);
	if (SQLITE_OK != ret) {
		printf("Error binding string to the prepared statement: %s\n",
		       sqlite3_errmsg(db));
		return;
	}

	time_t start;
	start = time(NULL);
	while (1) {
		ret = sqlite3_step(stmt);
		if (ret == SQLITE_ROW) {
			//fprintf(stderr, "%s\n", sqlite3_column_text(stmt, 0));
			strncat(results, (char *)sqlite3_column_text(stmt, 0),
				1024 - strlen(results) - 1);
			strcat(results, "\n");
		} else {
			//fprintf(stderr, "%d\n", ret);
			break;
		}
	}
	time_t end;
	end = time(NULL);
	fprintf(stderr, "Query took [%f] seconds\n", difftime(end, start));
	sqlite3_reset(stmt);

	gtk_label_set_text(GTK_LABEL(suggestions), results);
}

#define UNUSED(x) (void)(x)

static int disk_to_mem_cb(void *data, int argc, char **argv, char **azColName)
{
	UNUSED(data);
	UNUSED(argc);
	UNUSED(azColName);

	sqlite3_reset(stmt);
	sqlite3_bind_text(stmt, 1, argv[0], -1, SQLITE_STATIC);
	sqlite3_bind_text(stmt, 2, argv[1], -1, SQLITE_STATIC);
	sqlite3_bind_int(stmt, 3, atoi(argv[2]));

	if (SQLITE_DONE != sqlite3_step(stmt)) {
		printf
		    ("Error executing the prepared statement: %s\n",
		     sqlite3_errmsg(db));
		return -1;
	}

	return 0;
}

int main()
{
	int ret;
	sqlite3 *disk_db = NULL;

	fprintf(stderr,
		"Opening disk database and trying to create the in memory database\n");
	ret =
	    sqlite3_open_v2("ta-wiki.sqlite", &disk_db, SQLITE_OPEN_READONLY,
			    NULL);
	if (SQLITE_OK != ret) {
		printf("Error opening database: %s\n", sqlite3_errmsg(disk_db));
		goto end;
	}

	ret = sqlite3_open(":memory:", &db);
	if (SQLITE_OK != ret) {
		if (db == NULL)
			fprintf(stderr, "Insufficient Memory\n");
		else
			printf("Error opening database: %s\n",
			       sqlite3_errmsg(db));
		goto close_and_end;
	}

	char *err;
	ret =
	    sqlite3_exec(db,
			 "CREATE TABLE autocomplete (english varchar2, tamil varchar2, score int)",
			 NULL, NULL, &err);
	if (SQLITE_OK != ret) {
		printf("Error creating sqlite table: %s\n", sqlite3_errmsg(db));
		goto close_and_end;
	}

	ret =
	    sqlite3_prepare_v2(db, "INSERT INTO autocomplete VALUES (?, ?, ?)",
			       -1, &stmt, NULL);
	if (SQLITE_OK != ret) {
		printf("Error creating prepared statement: %s\n",
		       sqlite3_errmsg(db));
		goto close_and_end;
	}

	ret = sqlite3_exec(db, "BEGIN TRANSACTION", NULL, NULL, &err);
	if (SQLITE_OK != ret) {
		printf("Error beginning a transaction: %s\n",
		       sqlite3_errmsg(db));
		goto close_and_end;
	}

	/* Select each record from the disk db and insert into the in-memory table */
	ret =
	    sqlite3_exec(disk_db,
			 "SELECT english, tamil, score FROM autocomplete",
			 disk_to_mem_cb, NULL, &err);
	if (SQLITE_OK != ret) {
		printf("Error conversion of disk_to_mem: %s\n",
		       sqlite3_errmsg(db));
		goto close_and_end;
	}

	ret = sqlite3_exec(db, "COMMIT", NULL, NULL, &err);
	if (SQLITE_OK != ret) {
		printf("Error committing a transaction: %s\n",
		       sqlite3_errmsg(db));
		goto close_and_end;
	}
	sqlite3_finalize(stmt);
	sqlite3_close(disk_db);

	fprintf(stderr, "In-memory database creation complete\n");

	ret =
	    sqlite3_prepare_v2(db,
			       "SELECT DISTINCT tamil FROM autocomplete WHERE english LIKE ? ORDER BY score DESC LIMIT 5",
			       -1, &stmt, NULL);

	if (SQLITE_OK != ret) {
		printf("Error creating prepared statement: %s\n",
		       sqlite3_errmsg(db));
		goto close_and_end;
	}

	GtkWidget *window;

	gtk_init(NULL, NULL);
	window = gtk_window_new(GTK_WINDOW_TOPLEVEL);
	gtk_window_set_title(GTK_WINDOW(window), "Vaiyakani");

	entry = gtk_entry_new();
	suggestions = gtk_label_new("");

	GtkWidget *box;
	box = gtk_box_new(GTK_ORIENTATION_HORIZONTAL, 10);

	gtk_container_add(GTK_CONTAINER(box), entry);
	gtk_container_add(GTK_CONTAINER(box), suggestions);

	gtk_container_add(GTK_CONTAINER(window), box);
	g_signal_connect(entry, "changed", kb, NULL);

	gtk_widget_show_all(window);
	gtk_main();

	sqlite3_finalize(stmt);

close_and_end:
	if (disk_db)
		sqlite3_close(disk_db);
	if (db)
		sqlite3_close(db);
end:
	return ret;
}
