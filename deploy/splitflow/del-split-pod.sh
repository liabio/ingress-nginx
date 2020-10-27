pods=$(kubectl get pod -n kube-system -owide | grep liang | awk '{print $1}')

deletepods=""
for podname in ${pods}
do   
	echo "podname: ${podname}"   
	deletepods=" ${deletepods} ${podname} "
done  
echo "deletepods: ${deletepods}"
kubectl delete pod -n kube-system ${deletepods}
