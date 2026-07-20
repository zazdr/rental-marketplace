.PHONY: status start stop init set-password

export PGDATA := ./data
export PGPORT := 5432
export PGHOST := 127.0.0.1
export PGUSER := postgres
export PGPASSWORD := pass
export PGDATABASE := postgres

status:
	pg_ctl status

start:
	pg_ctl -l /dev/null start

stop:
	pg_ctl stop

init:
	mkdir -p $(PGDATA)
	initdb -U $(PGUSER)
	make set-password

set-password:
	make start
	psql -c "ALTER USER $(PGUSER) PASSWORD '$(PGPASSWORD)';"
	sed -i '' 's/trust/scram-sha-256/g' $(PGDATA)/pg_hba.conf
	make stop
