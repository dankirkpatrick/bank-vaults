module github.com/dankirkpatrick/bank-vaults

go 1.12

require (
	cloud.google.com/go v0.43.0
	github.com/Azure/azure-sdk-for-go v23.2.0+incompatible
	github.com/Azure/go-autorest v11.7.0+incompatible
	github.com/Masterminds/semver v1.4.2
	github.com/Masterminds/sprig v2.15.0+incompatible
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20190308093441-53f19b3c6bee
	github.com/aliyun/aliyun-oss-go-sdk v0.0.0-20171213061034-52de7239022c
	github.com/aws/aws-sdk-go v1.15.31
	github.com/banzaicloud/bank-vaults v0.0.0-20190910135916-f367aad78970
	github.com/banzaicloud/bank-vaults/pkg/sdk v0.0.0-00010101000000-000000000000
	github.com/banzaicloud/k8s-objectmatcher v1.0.1-0.20190813150246-386389f72468
	github.com/coreos/etcd-operator v0.9.3
	github.com/coreos/prometheus-operator v0.29.0
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/docker/docker v1.13.1
	github.com/fsnotify/fsnotify v1.4.7
	github.com/gin-gonic/gin v1.4.0
	github.com/goph/emperror v0.17.2
	github.com/hashicorp/vault/api v1.0.4
	github.com/heroku/docker-registry-client v0.0.0-20181004091502-47ecf50fd8d4
	github.com/imdario/mergo v0.3.7
	github.com/jpillora/backoff v0.0.0-20180909062703-3050d21c67d7
	github.com/opencontainers/image-spec v1.0.1
	github.com/operator-framework/operator-sdk v0.9.0
	github.com/prometheus/client_golang v1.0.0
	github.com/sirupsen/logrus v1.4.2
	github.com/slok/kubewebhook v0.3.0
	github.com/spf13/cast v1.3.0
	github.com/spf13/cobra v0.0.4
	github.com/spf13/viper v1.4.0
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	google.golang.org/api v0.3.0
	k8s.io/api v0.0.0-20190820101039-d651a1528133
	k8s.io/apimachinery v0.0.0-20190823012420-8ca64af22337
	k8s.io/client-go v11.0.1-0.20190516230509-ae8359b20417+incompatible
	k8s.io/code-generator v0.0.0-20190808180452-d0071a119380
	k8s.io/kube-openapi v0.0.0-20190722073852-5e22f3d471e6
	sigs.k8s.io/controller-runtime v0.1.10
)

replace github.com/banzaicloud/bank-vaults/pkg/sdk => ./pkg/sdk

replace cloud.google.com/go => cloud.google.com/go v0.26.0

replace k8s.io/api => k8s.io/api v0.0.0-20181213150558-05914d821849

replace k8s.io/apimachinery => k8s.io/apimachinery v0.0.0-20181127025237-2b1284ed4c93

replace k8s.io/client-go => k8s.io/client-go v10.0.0+incompatible

replace k8s.io/code-generator => k8s.io/code-generator v0.0.0-20190416052311-01a054e913a9

replace golang.org/x/oauth2 => golang.org/x/oauth2 v0.0.0-20180821212333-d2e6202438be

replace git.apache.org/thrift.git => github.com/apache/thrift v0.12.0
