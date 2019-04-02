migrate:
	mysql --login-path=local blog < migrations/initial_01.sql
seed:
	mysql --login-path=local blog < migrations/seed.sql
clean:
	mysql --login-path=local blog < migrations/clean.sql