## imgvuln
A simple image vulnerability scanner for programmers to scan kubernetes pods and workloads.

# How to use
step 1: fork or clone repo
`https:github.com/adeniyistephen/imgvuln`
step 2: change directory to the imgvuln path
`cd cmd     cd imgvuln`
step 3: build project
`go build`
step 4: start a kubernetes cluster
eg: `kind create cluster --name mycluster --image kindest/node:v1.20.2`
step 5: deploy the PolicyReport resource
`cd .. | cd .. | kubectl create -f kubernetes/crd/v1alpha2/wgpolicyk8s.io_clusterpolicyreports.yaml`
step 6: deploy the vulnerabilityReport crd
`cd cmd | cd imgvuln | ./imgvuln init`
step 7: deploy a simple image that has little vulnerability
`kubectl create deployment zipkin --image openzipkin/zipkin:latest`
to check if deployment is ready: `kubectl get pod --all-namespaces`
step 8: scan deployed resources
`./imgvuln scan policyreports zipkin-gdhd988748hd`
step 9: get clusterreport
`kubectl get clusterpolicyreports zipkin-gdhg78938hd -o yaml`

Thanks. you can go through the codes Jim, please kindly check the pkg/cmd/write.go file and pkg/plugin/trivy/plugin.go function(ParsePolicyReportData)

still in my personal repo, will delete after building, will transfer code to the wg-policy repo when done.

Stephen.