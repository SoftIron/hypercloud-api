package cloud

import (
	"fmt"
	"strings"

	"github.com/softiron/hypercloud-api/client"
)

// Service owns the /cloud methods.
type Service struct {
	ACLService
	ClusterService
	HostService
	DataCenterService
	DatastoreService
	DocumentService
	GroupService
	HookService
	ImageService
	InstanceService
	MarketService
	NetworkService
	RouterService
	SecurityGroupService
	SystemService
	TemplateService
	UserService
	ZoneService

	*service
}

type service struct {
	*client.Client
	root string
}

// NewService returns a new Service for bare-metal operations.
func NewService(c *client.Client, path string) *Service {
	s := &service{Client: c, root: path}

	return &Service{
		ACLService:           ACLService{service: s},
		ClusterService:       ClusterService{service: s},
		DataCenterService:    DataCenterService{service: s},
		DatastoreService:     DatastoreService{service: s},
		DocumentService:      DocumentService{service: s},
		GroupService:         GroupService{service: s},
		HookService:          HookService{service: s},
		HostService:          HostService{service: s},
		ImageService:         ImageService{service: s},
		InstanceService:      InstanceService{service: s},
		MarketService:        MarketService{service: s},
		NetworkService:       NetworkService{service: s},
		RouterService:        RouterService{service: s},
		SecurityGroupService: SecurityGroupService{service: s},
		SystemService:        SystemService{service: s},
		TemplateService:      TemplateService{service: s},
		UserService:          UserService{service: s},
		ZoneService:          ZoneService{service: s},

		service: s,
	}
}

// path returns a URL path with the correct prefix appended.
func (s service) path(dirs ...interface{}) string {
	p := make([]string, 2, len(dirs)+1)
	p[0] = s.root
	p[1] = PathPrefix

	for _, d := range dirs {
		p = append(p, fmt.Sprint(d))
	}

	return strings.Join(p, "/")
}
