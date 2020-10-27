pods=$(kubectl get pod -n kube-system -owide | grep liang | awk '{print $1}')
echo "args is $1"
logspod=""
echo "logspod is ${logspod}"
for i in ${!pods[@]}
do   
	echo "podname: ${pods[$i]}"  

        if [ "$1" = $i  ]; then
		logspod=${pods[$i]}
		break	
	fi
done 
echo "logspod is ${logspod}"
if [ "${logspod}"z != ""z  ]; then
	kubectl logs -f -n kube-system ${logspod}
fi

