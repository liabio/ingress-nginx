#!/bin/bash
kubectl delete cm nginx-splitflow-tmpl -n kube-system
kubectl create cm nginx-splitflow-tmpl -n kube-system --from-file=nginx.tmpl=./nginx-splitflow.tmpl
kubectl patch cm nginx-splitflow-tmpl -n kube-system  --patch '{"metadata":{"labels":{"name":"ng-tmpl","type":"c"}}}'
