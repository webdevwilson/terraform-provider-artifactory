package artifactory

import (
	"encoding/json"
	"fmt"
)

// LocalRepositoryConfiguration contains items present in local repository requests
type LocalRepositoryConfiguration struct {
	Key                          string   `json:"key,omitempty"`
	RClass                       string   `json:"rclass,omitempty"`
	PackageType                  string   `json:"packageType,omitempty"`
	Description                  string   `json:"description,omitempty"`
	Notes                        string   `json:"notes,omitempty"`
	IncludesPattern              string   `json:"includesPattern,omitempty"`
	ExcludesPattern              string   `json:"excludesPattern,omitempty"`
	RepoLayoutRef                string   `json:"repoLayoutRef,omitempty"`
	HandleReleases               bool     `json:"handleReleases,omitempty"`
	HandleSnapshots              bool     `json:"handleSnapshots,omitempty"`
	MaxUniqueSnapshots           int      `json:"maxUniqueSnapshots,omitempty"`
	DebianTrivialLayout          bool     `json:"debianTrivialLayout,omitempty"`
	ChecksumPolicyType           string   `json:"checksumPolicyType,omitempty"`
	MaxUniqueTags                int      `json:"maxUniqueTags,omitempty"`
	SnapshotVersionBehavior      string   `json:"snapshotVersionBehavior,omitempty"`
	SuppressPomConsistencyChecks bool     `json:"suppressPomConsistencyChecks,omitempty"`
	BlackedOut                   bool     `json:"blackedOut,omitempty"`
	PropertySets                 []string `json:"propertySets,omitempty"`
	ArchiveBrowsingEnabled       bool     `json:"archiveBrowsingEnabled,omitempty"`
	CalculateYumMetadata         bool     `json:"calculateYumMetadata,omitempty"`
	YumRootDepth                 int      `json:"yumRootDepth,omitempty"`
	DockerAPIVersion             string   `json:"dockerApiVersion,omitempty"`
	EnableFileListsIndexing      bool     `json:"enableFileListsIndexing,omitempty"`
}

// RemoteRepositoryConfiguration for configuring a remote repository
type RemoteRepositoryConfiguration struct {
	Key                               string   `json:"key,omitempty"`
	RClass                            string   `json:"rclass,omitempty"`
	PackageType                       string   `json:"packageType,omitempty"`
	URL                               string   `json:"url,omitempty"`
	Username                          string   `json:"username,omitempty"`
	Password                          string   `json:"password,omitempty"`
	Proxy                             string   `json:"proxy,omitempty"`
	Description                       string   `json:"description,omitempty"`
	Notes                             string   `json:"notes,omitempty"`
	IncludesPattern                   string   `json:"includesPattern,omitempty"`
	ExcludesPattern                   string   `json:"excludesPattern,omitempty"`
	RepoLayoutRef                     string   `json:"repoLayoutRef,omitempty"`
	RemoteRepoChecksumPolicyType      string   `json:"remoteRepoChecksumPolicyType,omitempty"`
	HandleReleases                    bool     `json:"handleReleases,omitempty"`
	HandleSnapshots                   bool     `json:"handleSnapshots,omitempty"`
	MaxUniqueSnapshots                int      `json:"maxUniqueSnapshots,omitempty"`
	SuppressPomConsistencyChecks      bool     `json:"suppressPomConsistencyChecks,omitempty"`
	HardFail                          bool     `json:"hardFail,omitempty"`
	Offline                           bool     `json:"offline,omitempty"`
	BlackedOut                        bool     `json:"blackedOut,omitempty"`
	StoreArtifactsLocally             bool     `json:"storeArtifactsLocally,omitempty"`
	SocketTimeoutMillis               int      `json:"socketTimeoutMillis,omitempty"`
	LocalAddress                      string   `json:"localAddress,omitempty"`
	RetrievalCachePeriodSeconds       int      `json:"retrievalCachePeriodSecs,omitempty"`
	FailedCachePeriodSeconds          int      `json:"failedRetrievalCachePeriodSecs,omitempty"`
	MissedCachePeriodSeconds          int      `json:"missedRetrievalCachePeriodSecs,omitempty"`
	UnusedArtifactsCleanupEnabled     bool     `json:"unusedArtifactsCleanupEnabled,omitempty"`
	UnusedArtifactsCleanupPeriodHours int      `json:"unusedArtifactsCleanupPeriodHours,omitempty"`
	FetchJarsEagerly                  bool     `json:"fetchJarsEagerly,omitempty"`
	FetchSourcesEagerly               bool     `json:"fetchSourcesEagerly,omitempty"`
	ShareConfiguration                bool     `json:"shareConfiguration,omitempty"`
	SynchronizeProperties             bool     `json:"synchronizeProperties,omitempty"`
	PropertySets                      []string `json:"propertySets,omitempty"`
	AllowAnyHostAuth                  bool     `json:"allowAnyHostAuth,omitempty"`
	EnableCookieManagement            bool     `json:"enableCookieManagement,omitempty"`
	BowerRegistryURL                  string   `json:"bowerRegistryUrl,omitempty"`
	VCSType                           string   `json:"vcsType,omitempty"`
	VCSGitProvider                    string   `json:"vcsGitProvider,omitempty"`
	VCSGitDownloadURL                 string   `json:"vcsGitDownloadUrl,omitempty"`
}

// VirtualRepositoryConfiguration for
type VirtualRepositoryConfiguration struct {
	Key                                           string   `json:"key,omitempty"`
	RClass                                        string   `json:"rclass,omitempty"`
	PackageType                                   string   `json:"packageType,omitempty"`
	Description                                   string   `json:"description,omitempty"`
	Notes                                         string   `json:"notes,omitempty"`
	IncludesPattern                               string   `json:"includesPattern,omitempty"`
	ExcludesPattern                               string   `json:"excludesPattern,omitempty"`
	ArtifactoryRequestsCanRetrieveRemoteArtifacts bool     `json:"artifactoryRequestsCanRetrieveRemoteArtifacts,omitempty"`
	KeyPair                                       string   `json:"keyPair,omitempty"`
	PomRepositoryReferencesCleanupPolicy          string   `json:"pomRepositoryReferencesCleanupPolicy,omitempty"`
	DefaultDeploymentRepo                         string   `json:"defaultDeploymentRepo,omitempty"`
	Repositories                                  []string `json:"repositories,omitempty"`
}

// GetRepository fetches repository configuration from Artifactory
func (c clientConfig) GetRepository(key string, v interface{}) error {
	path := fmt.Sprintf("repositories/%s", key)
	resp, err := c.execute("GET", path, nil)

	if err != nil {
		return err
	}

	if err := c.validateResponse(200, resp.StatusCode, "read repository"); err != nil {
		return err
	}

	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(v)
	if err != nil {
		return err
	}

	return resp.Body.Close()
}

// CreateRepository Creates a repository in Artifactory
func (c clientConfig) CreateRepository(key string, v interface{}) error {
	path := fmt.Sprintf("repositories/%s", key)
	resp, err := c.execute("PUT", path, v)

	if err != nil {
		return err
	}

	if err := c.validateResponse(200, resp.StatusCode, "create repository"); err != nil {
		return err
	}

	return resp.Body.Close()
}

// UpdateRepository Updates an Artifactory repository
func (c clientConfig) UpdateRepository(key string, v interface{}) error {
	path := fmt.Sprintf("repositories/%s", key)
	resp, err := c.execute("POST", path, v)

	if err != nil {
		return err
	}

	if err := c.validateResponse(200, resp.StatusCode, "update repository"); err != nil {
		return err
	}

	return resp.Body.Close()
}

// DeleteRepository deletes a repository from Artifactory
func (c clientConfig) DeleteRepository(key string) error {
	path := fmt.Sprintf("repositories/%s", key)
	resp, err := c.execute("DELETE", path, nil)

	if err != nil {
		return err
	}

	if err := c.validateResponse(200, resp.StatusCode, "delete repository"); err != nil {
		return err
	}

	return resp.Body.Close()
}
