default:
	echo "hi there!"

start:	
	make down
	make up
	make run-dev

# setup cluster of dependencies in docker
up:
	docker compose -f "compose.yml" up -d --build --wait

# run air
run-dev:
	./scripts/run_dev.sh
# (trap 'kill 0' SIGINT; make run-dev-backend & make run-dev-frontend)

# run-dev-backend:
# 	echo 'hello world
# # (trap 'kill 0' SIGINT; cd backend && air)

# run-dev-frontend:
# 	(trap 'kill 0' SIGINT; cd frontend && npm install && npm run dev -- --open)

# teardown dev setup
down:
	docker compose -f "compose.yml" down
	(echo "y" | docker volume prune)
	clear
	
get-bastion-keys: 
	./scripts/retrieve_bastion_ssh_keys.sh

connect-prod-db-tunnel:
	./scripts/connect_prod_db_tunnel.sh