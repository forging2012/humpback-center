package ctrl

import "github.com/humpback/gounits/logger"
import "humpback-center/cluster"
import "humpback-center/etc"
import "humpback-center/repository"

// Controller is exprted
type Controller struct {
	Configuration   *etc.Configuration
	Cluster         *cluster.Cluster
	RepositoryCache *repository.RepositoryCache
}

// NewController is exported
func NewController(configuration *etc.Configuration) (*Controller, error) {

	cluster, err := createCluster(configuration)
	if err != nil {
		return nil, err
	}

	repositorycache, err := createRepositoryCache(configuration)
	if err != nil {
		return nil, err
	}

	return &Controller{
		Configuration:   configuration,
		Cluster:         cluster,
		RepositoryCache: repositorycache,
	}, nil
}

// Initialize is exported
// init cluster
func (c *Controller) Initialize() error {

	logger.INFO("[#ctrl#] controller initialize.....")
	return c.startCluster()
}

// UnInitialize is exported
// uninit cluster
func (c *Controller) UnInitialize() {

	c.stopCluster()
	logger.INFO("[#ctrl#] controller uninitialized.")
}
