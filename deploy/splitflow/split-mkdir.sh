#!/bin/bash

containerId=$(docker run -ti --rm -d -u 0 k8s-deploy/nginx-ingress-controller:0.24.1 bash)

echo "containerId is ${containerId}"

docker cp nginx-ingress-controller ${containerId}:/nginx-ingress-controller
docker cp ingressgroup-upstream.tmpl ${containerId}:/etc/nginx/template/ingressgroup-upstream.tmpl

docker exec -ti -u 0 ${containerId} chmod 755 /nginx-ingress-controller 
docker exec -ti -u 0 ${containerId} chown www-data:www-data /nginx-ingress-controller

docker exec -ti -u 0 ${containerId} chmod 644 /etc/nginx/template/ingressgroup-upstream.tmpl
docker exec -ti -u 0 ${containerId} chown www-data:www-data /etc/nginx/template/ingressgroup-upstream.tmpl

docker exec -ti -u 0 ${containerId} mkdir -p /etc/nginx/conf.d/include-server-map/
docker exec -ti -u 0 ${containerId} chmod 766 -R /etc/nginx/conf.d
docker exec -ti -u 0 ${containerId} chown www-data:www-data -R /etc/nginx/conf.d 

echo "commit ${containerId}"
docker commit ${containerId} 192.168.26.46/k8s-deploy/nginx-ingress-controller-0.24.1-splitflow:v1

echo "push ${containerId}"
docker push  192.168.26.46/k8s-deploy/nginx-ingress-controller-0.24.1-splitflow:v1

echo "stop temp containerId  ${containerId}"
docker stop ${containerId}
