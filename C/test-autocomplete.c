#include <stdio.h>
#include <string.h>

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
	sqlite3_reset(stmt);

	gtk_label_set_text(GTK_LABEL(suggestions), results);
}

int main()
{
	int ret;

	ret =
	    sqlite3_open_v2("autocomplete.sqlite", &db, SQLITE_OPEN_READONLY,
			    NULL);
	if (SQLITE_OK != ret) {
		printf("Error opening database: %s\n", sqlite3_errmsg(db));
		goto end;
	}

	ret =
	    sqlite3_prepare_v2(db,
			       "SELECT tamil FROM autocomplete WHERE english LIKE ? ORDER BY score DESC LIMIT 5",
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
	sqlite3_close(db);
end:
	return ret;
}
