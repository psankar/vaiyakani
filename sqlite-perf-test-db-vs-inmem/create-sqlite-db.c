#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include <glib.h>
#include <sqlite3.h>

int main(int argc, char *argv[])
{

	sqlite3 *db;
	int ret = -1;

	if (2 != argc) {
		printf("Usage: create-db <filename.englishwords>\n");
		goto end;
	}

	GError *gerr;
	gchar *contents;
	if (!g_file_get_contents(argv[1], &contents, NULL, &gerr)) {
		printf("Error opening the .englishwords file: [%s]\n",
		       gerr->message);
		goto end;
	}

	ret = sqlite3_open("autocomplete.sqlite", &db);
	if (SQLITE_OK != ret) {
		printf("Error opening sqlite database: %s\n",
		       sqlite3_errmsg(db));
		goto end;
	}

	char *err;
	ret =
	    sqlite3_exec(db,
			 "CREATE TABLE IF NOT EXISTS autocomplete (english TEXT COLLATE NOCASE, tamil COLLATE NOCASE, score int)",
			 NULL, NULL, &err);
	if (SQLITE_OK != ret) {
		printf("Error creating sqlite table: %s\n", sqlite3_errmsg(db));
		goto close_and_end;
	}
	fprintf(stderr, "Table created. Inserting data\n");

	ret = sqlite3_exec(db, "BEGIN TRANSACTION", NULL, NULL, &err);
	if (SQLITE_OK != ret) {
		printf("Error beginning a transaction: %s\n",
		       sqlite3_errmsg(db));
		goto close_and_end;
	}

	sqlite3_stmt *stmt;
	ret =
	    sqlite3_prepare_v2(db, "INSERT INTO autocomplete VALUES (?, ?, ?)",
			       -1, &stmt, NULL);
	if (SQLITE_OK != ret) {
		printf("Error creating prepared statement: %s\n",
		       sqlite3_errmsg(db));
		goto close_and_end;
	}

	gchar **lines;
	lines = g_strsplit(contents, "\n", -1);
	g_free(contents);

	gchar *line;
	int i = 0;
	while (1) {
		line = lines[i++];

		if (!line || !strlen(line))
			break;

		gchar **tokens;
		tokens = g_strsplit(line, ",", -1);

		int j;
		j = 2;
		while (tokens[j]) {
			sqlite3_reset(stmt);

			//fprintf(stderr, "Inserting [%s]\n", tokens[j]);
			if (SQLITE_OK !=
			    sqlite3_bind_text(stmt, 1, tokens[j], -1,
					      SQLITE_STATIC)) {
				printf
				    ("Error binding the first paramater: %s\n",
				     sqlite3_errmsg(db));
				goto close_and_end;
			}

			if (SQLITE_OK !=
			    sqlite3_bind_text(stmt, 2, tokens[1], -1,
					      SQLITE_STATIC)) {
				printf
				    ("Error binding the second paramater: %s\n",
				     sqlite3_errmsg(db));
				goto close_and_end;
			}

			if (SQLITE_OK !=
			    sqlite3_bind_int(stmt, 3, atoi(tokens[0]))) {
				printf
				    ("Error binding the third paramater: %s\n",
				     sqlite3_errmsg(db));
				goto close_and_end;
			}

			if (SQLITE_DONE != sqlite3_step(stmt)) {
				printf
				    ("Error executing the prepared statement: %s\n",
				     sqlite3_errmsg(db));
				goto close_and_end;
			}
			j++;
		}
		g_strfreev(tokens);

	}
	g_strfreev(lines);

	ret = sqlite3_exec(db, "COMMIT", NULL, NULL, &err);
	if (SQLITE_OK != ret) {
		printf("Error committing a transaction: %s\n",
		       sqlite3_errmsg(db));
	}

	/* Create Indexes */
	fprintf(stderr, "Data inserted into tables. About to create english index\n");
	ret = sqlite3_exec(db, "CREATE INDEX english_idx ON autocomplete (english)", NULL, NULL, &err);
	if (SQLITE_OK != ret) {
		printf("Error creating an index for the english column: %s\n",
		       sqlite3_errmsg(db));
	}

	fprintf(stderr, "Index created for english. Now creating index for english + tamil\n");
	ret = sqlite3_exec(db, "CREATE INDEX eng_tam_idx ON autocomplete (english, tamil)", NULL, NULL, &err);
	if (SQLITE_OK != ret) {
		printf("Error creating an index for the english+tamil columns: %s\n",
		       sqlite3_errmsg(db));
	}

	fprintf(stderr, "Index created for english + tamil. Now creating index for english + tamil + score\n");
	ret = sqlite3_exec(db, "CREATE INDEX eng_tam_score_idx ON autocomplete (english, tamil, score)", NULL, NULL, &err);
	if (SQLITE_OK != ret) {
		printf("Error creating an index for the english+tamil+score columns: %s\n",
		       sqlite3_errmsg(db));
	}

	fprintf(stderr, "Index created for english + tamil + score. Now creating index for english + score\n");
	ret = sqlite3_exec(db, "CREATE INDEX english_score_idx ON autocomplete (english, score)", NULL, NULL, &err);
	if (SQLITE_OK != ret) {
		printf("Error creating an index for the english+tamil+score columns: %s\n",
		       sqlite3_errmsg(db));
	}

close_and_end:
	sqlite3_close(db);

end:
	return ret;
}
