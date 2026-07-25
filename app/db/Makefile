.PHONY: status start stop init

include ../env/.env.db
export

status:
	@pg_ctl status

start:
	@pg_ctl -l /dev/null start

stop:
	@pg_ctl stop

init:
	@mkdir -p $(PGDATA)
	@initdb -U $(PGUSER)
	@make start
	@psql -c "ALTER USER $(PGUSER) PASSWORD '$(PGPASSWORD)';"
	@sed -i '' 's/trust/scram-sha-256/g' $(PGDATA)/pg_hba.conf
	@make stop
