.PHONY: image
image:
	docker build -t app1:0.0.1 .
	docker build -t app1:0.0.2 .
	docker build -t app1:0.0.3 .

	docker build -t app2:0.0.1 .
	docker build -t app2:0.0.2 .
	docker build -t app2:0.0.3 .