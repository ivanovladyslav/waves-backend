run: ; $(info $(M)Running backend app...)
	env `cat ./env/.env | xargs` go run .