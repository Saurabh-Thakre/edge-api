// FIXME: golangci-lint
// nolint:revive
package config

// List of supported RHEL distributions
const distRHEL84 = "rhel-84"
const distRHEL85 = "rhel-85"
const distRHEL86 = "rhel-86"
const distRHEL87 = "rhel-87"
const distRHEL90 = "rhel-90"

// DefaultDistribution set the default image distribution in case miss it
const DefaultDistribution = distRHEL90

// ostree ref for supported distributions
const OstreeRefRHEL8 = "rhel/8/x86_64/edge"
const OstreeRefRHEL9 = "rhel/9/x86_64/edge"

// RequiredPackages contains minimun list of packages to build an image
var RequiredPackages = []string{"rhc",
	"rhc-worker-playbook",
	"subscription-manager",
	"subscription-manager-plugin-ostree",
	"insights-client"}

// RHEL8 contains additional list of packages to build an image to >= RHEL85
var RHEL8 = []string{"ansible"}

// RHEL8X contains additional list of packages to build an image to = RHEL8X
var RHEL8X = []string{"ansible-core"}

// RHEL90 contains additional list of packages to build an image to = RHEL90
var RHEL90 = []string{"ansible-core"}

// DistributionsPackages add packages by image
var DistributionsPackages = map[string][]string{
	distRHEL84: RHEL8,
	distRHEL85: RHEL8,
	distRHEL86: RHEL8X,
	distRHEL87: RHEL8X,
	distRHEL90: RHEL90,
}

// DistributionsRefs set the ref to Images
var DistributionsRefs = map[string]string{
	distRHEL84: OstreeRefRHEL8,
	distRHEL85: OstreeRefRHEL8,
	distRHEL86: OstreeRefRHEL8,
	distRHEL87: OstreeRefRHEL8,
	distRHEL90: OstreeRefRHEL9,
}
