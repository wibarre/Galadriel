package telemetry

// package name
const (
	Harvester = "harvester"
	Server    = "server"
)

// entity
const (
	TrustBundle = "trust_bundle"
	PackageName = "package_name"
	Federation  = "federation"
)

// action
const (
	Add     = "add"
	Get     = "get"
	Remove  = "remove"
	List    = "list"
	Create  = "create"
	Approve = "approve"
	Deny    = "deny"
)

// component
const (
	Endpoints           = "endpoints"
	HarvesterController = "harvester_controller"
	GaladrielServer     = "galadriel_server"
	CA                  = "ca"
	ManagerCA           = "manager_ca"
	DiskX509CA          = "disk_x509_ca"

	// SubsystemName declares a field for some subsystem name (an API, module...)
	SubsystemName = "subsystem_name"

	GaladrielServerClient = "galadriel_server_client"
)
