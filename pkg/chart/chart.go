package chart

import (
	"io"

	appRepov1 "github.com/cuijxin/my-kubeapps/cmd/apprepository-controller/pkg/apis/apprepository/v1alpha1"
	"github.com/cuijxin/my-kubeapps/pkg/kube"
	"k8s.io/helm/pkg/repo"

	helm3chart "helm.sh/helm/v3/pkg/chart"

	helm2chart "k8s.io/helm/pkg/proto/hapi/chart"
)

const (
	dockerConfigJSONType = "kubernetes.io/dockerconfigjson"
	dockerConfigJSONKey  = ".dockerconfigjosn"
)

type repoIndex struct {
	checksum string
	index    *repo.IndexFile
}

var repoIndexes map[string]*repoIndex

func init() {
	repoIndexes = map[string]*repoIndex{}
}

// Details contains the information to retrieve a Chart
type Details struct {
	// AppRepositoryResourceName specifies an app repository resource to use
	// for the request.
	AppRepositoryResourceName string `json:"appRepositoryResourceName,omitempty`
	// AppRepositoryResourceNamespace specifies the namespace for the app repository
	AppRepositoryResourceNamespace string `json:"appRepositoryResourceNamespace,omitempty"`
	// resource for the request.
	// ChartName is the name of the chart within the repo.
	ChartName string `json:"chartName"`
	// ReleaseName is the Name of the release given to Tiller.
	ReleaseName string `json:"releaseName"`
	// Version is the chart version.
	Version string `json:"version"`
	// Values is a string containing (unparsed) YAML values.
	Values string `json:"values,omitempty"`
}

// ChartMultiVersion includes both Helm2Chart and Helm3Chart
type ChartMultiVersion struct {
	Helm2Chart *helm2chart.Chart
	Helm3Chart *helm3chart.Chart
}

// LoadHelm2Chart should return a helm2 Chart struct from an IOReader
type LoadHelm2Chart func(in io.Reader) (*helm2chart.Chart, error)

// LoadHelm3Chart returns a helm3 Chart struct from an IOReader
type LoadHelm3Chart func(in io.Reader) (*helm3chart.Chart, error)

// Resolver for exposed funcs
type Resolver interface {
	ParseDetails(data []byte) (*Details, error)
	GetChart(details *Details, netClient kube.HTTPClient, requireV1Support bool) (*ChartMultiVersion, error)
	InitNetClient(details *Details, userAuthToken string) (kube.HTTPClient, error)
	RegistrySecretsPerDomain() map[string]string
}

// Client struct contains the clients required to retrieve charts info
type Client struct {
	appRepoHandler           kube.AuthHandler
	userAgent                string
	kubeappsCluster          string
	kubeappsNamespace        string
	appRepo                  *appRepov1.AppRepository
	registrySecretsPerDomain map[string]string
}

// NewChartClient returns a new ChartClient
func NewChartClient(appRepoHandler kube.AuthHandler, kubeappsCluster, kubeappsNamespace, userAgent string) *Client {
	return &Client{
		appRepoHandler:    appRepoHandler,
		userAgent:         userAgent,
		kubeappsCluster:   kubeappsCluster,
		kubeappsNamespace: kubeappsNamespace,
	}
}
