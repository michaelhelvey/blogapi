migrate:
	mysql --login-path=local blog < migrations/initial_01.sql
clean:
	mysql --login-path=local blog < migrations/clean.sql