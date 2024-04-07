# Kubernetes enviroments using docker-desktop(Single node cluster)

# Part. Configmaps
* Register DB Scheme
	- kubectl create configmap init-db-script --from-file=DB/db.sql
* Register Simulator config file & app code file
	- kubectl create configmap simulator-config-app --from-file=Simulation/line_station.json --from-file=Simulation/simulator.py