## DATABASE UTILITIES

#### THESE SCRIPTS ALLOW BACKUPS OF DATABASES TO BE PERFORMED. FOLLOW DIRECTIONS CAREFULLY>

To restore the database to a recent snapshot, run `./restore.sh`. This will replace your
database with the data stowed in the `sql/restore_database.sql` script.

To take a snapshot of your existing database for future recovery, run `./clone.sh`. A `sql`
file will then be generated and placed in the `dumps/` directory. Rename this file to 
`restore_database.sql` and drop it in into the `sql/` directory overwriting the existing
copy. Or you can stash various copies as you like. 