build:
	docker build --tag stevenwal/mock-indexer .

push:
	docker push stevenwal/mock-indexer:latest