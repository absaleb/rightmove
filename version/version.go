package version

import (
	"gitlab.okta-solutions.com/mashroom/backend/common/health"
)

const major = 0
const minor = 0
const patch = 1
const relinfo = "master"

const status = health.HealthStatus_RUNNING
const description = "rightmove running nominal"

func NewHealthServer() health.Server {
	return health.NewServer(major, minor, patch, relinfo)
}
