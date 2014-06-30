#include <stdio.h>

#include <sqlite3.h>

int main()
{
	sqlite3 *db;
	int ret;

	ret =
	    sqlite3_open_v2("autocomplete.sqlite", &db, SQLITE_OPEN_READONLY,
			    NULL);
	if (SQLITE_OK != ret) {
		printf("Error opening database: %s\n", sqlite3_errmsg(db));
		goto end;
	}

	sqlite3_stmt *stmt;
	ret =
	    sqlite3_prepare_v2(db,
			       "SELECT tamil FROM autocomplete WHERE english LIKE 'a%%' ORDER BY score DESC LIMIT 5",
			       -1, &stmt, NULL);

	if (SQLITE_OK != ret) {
		printf("Error creating prepared statement: %s\n",
		       sqlite3_errmsg(db));
		goto close_and_end;
	}

	while ((ret = sqlite3_step(stmt)) == SQLITE_ROW) {
			printf("%s\n", sqlite3_column_text(stmt, 0));
	}
	sqlite3_finalize(stmt);

close_and_end:
	sqlite3_close(db);
end:
	return ret;
}
