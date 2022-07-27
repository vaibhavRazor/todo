.PHONY: deploy-roles
deploy-roles:
	kubectl apply -f prom_cluster_role.yaml
	kubectl apply -f prom_service.yaml
	kubectl apply -f prom_cluster_role_binding.yaml
 
.PHONY: deploy-prometheus
deploy-prometheus:
	kubectl apply -f prom_config_map.yaml
	kubectl apply -f prom_deployment.yaml
	kubectl apply -f prom_service.yaml
