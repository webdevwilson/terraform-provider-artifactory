package artifactory

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

// User represents an Artifactory user
type User struct {
	Name                     string   `json:"name,omitempty"`
	Email                    string   `json:"email,omitempty"`
	Password                 string   `json:"password,omitempty"`
	Admin                    bool     `json:"admin"`
	ProfileUpdatable         bool     `json:"profileUpdatable"`
	InternalPasswordDisabled bool     `json:"internalPasswordDisabled,omitempty"`
	LastLoggedIn             string   `json:"lastLoggedIn,omitempty"`
	Realm                    string   `json:"realm,omitempty"`
	Groups                   []string `json:"groups"`
}

// Group represents an Artifactory group
type Group struct {
	Name            string `json:"name,omitempty"`
	AutoJoin        bool   `json:"autoJoin"`
	Realm           string `json:"realm,omitempty"`
	RealmAttributes string `json:"realm,omitempty"`
}

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

type clientConfig struct {
	user string
	pass string
	url  string
}

// Client is used to call Artifactory REST APIs
type Client interface {
	Ping() error
	GetRepository(key string, v interface{}) error
	CreateRepository(key string, v interface{}) error
	UpdateRepository(key string, v interface{}) error
	DeleteRepository(key string) error
	GetUser(name string) (*User, error)
	CreateUser(u *User) error
	UpdateUser(u *User) error
	DeleteUser(name string) error
	GetGroup(name string) (*Group, error)
	CreateGroup(g *Group) error
	UpdateGroup(g *Group) error
	DeleteGroup(name string) error
	ExpireUserPassword(name string) error
}

var _ Client = clientConfig{}

// NewClient constructs a new artifactory client
func NewClient(username string, pass string, url string) Client {
	return clientConfig{
		username,
		pass,
		strings.TrimRight(url, "/"),
	}
}

// Ping calls the system to verify connectivity
func (c clientConfig) Ping() error {
	resp, err := c.execute("GET", "system/ping", nil)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Ping failed. Status: %s", resp.Status)
	}

	return resp.Body.Close()
}

// GetRepository fetches repository configuration from Artifactory
func (c clientConfig) GetRepository(key string, v interface{}) error {
	path := fmt.Sprintf("repositories/%s", key)
	resp, err := c.execute("GET", path, nil)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Failed to update repository. Status: %s", resp.Status)
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

	if resp.StatusCode != 200 {
		return fmt.Errorf("Failed to create repository. Status: %s", resp.Status)
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

	if resp.StatusCode != 200 {
		return fmt.Errorf("Failed to update repository. Status: %s", resp.Status)
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

	if resp.StatusCode != 200 {
		return fmt.Errorf("Failed to delete repository. Status: %s", resp.Status)
	}

	return resp.Body.Close()
}

// GetUser returns a User from Artifactory
func (c clientConfig) GetUser(name string) (*User, error) {
	path := fmt.Sprintf("security/users/%s", name)
	resp, err := c.execute("GET", path, nil)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Failed to get user. Status: %s", resp.Status)
	}

	user := &User{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(user)
	if err != nil {
		return nil, err
	}

	if err = resp.Body.Close(); err != nil {
		return nil, err
	}

	return user, nil
}

// CreateUser creates a new user in artifactory
func (c clientConfig) CreateUser(u *User) error {
	path := fmt.Sprintf("security/users/%s", u.Name)
	resp, err := c.execute("PUT", path, u)

	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		return fmt.Errorf("Failed to create user. Status: %s", resp.Status)
	}

	return resp.Body.Close()
}

// UpdateUser Updates user in Artifactory
func (c clientConfig) UpdateUser(u *User) error {
	path := fmt.Sprintf("security/users/%s", u.Name)
	resp, err := c.execute("POST", path, u)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Failed to update user. Status: %s", resp.Status)
	}

	return resp.Body.Close()
}

// DeleteUser in Artifactory
func (c clientConfig) DeleteUser(name string) error {
	path := fmt.Sprintf("security/users/%s", name)
	resp, err := c.execute("DELETE", path, nil)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Failed to delete user. Status: %s", resp.Status)
	}

	return resp.Body.Close()
}

// ExpireUserPassword
// Expires a user's password. This may initiate an email depending on configuration
// settings.
func (c clientConfig) ExpireUserPassword(name string) error {
	path := fmt.Sprintf("security/users/authorization/expirePassword/%s", name)
	resp, err := c.execute("POST", path, nil)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Error expiring password for '%s'. Got: %s", name, resp.Status)
	}

	return resp.Body.Close()
}

// GetGroup retrieves a group from Artifactory
// Returns either an error or a group, never both
func (c clientConfig) GetGroup(name string) (*Group, error) {
	path := fmt.Sprintf("security/groups/%s", name)
	resp, err := c.execute("GET", path, nil)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Error retrieving group '%s'", name)
	}

	group := &Group{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(group)
	if err != nil {
		return nil, err
	}

	if err = resp.Body.Close(); err != nil {
		return nil, err
	}

	return group, nil
}

// CreateGroup creates a new user in artifactory
func (c clientConfig) CreateGroup(g *Group) error {
	path := fmt.Sprintf("security/groups/%s", g.Name)
	resp, err := c.execute("PUT", path, g)

	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		return fmt.Errorf("Failed to create group. Status: %s", resp.Status)
	}

	return resp.Body.Close()
}

// UpdateGroup Updates group in Artifactory
func (c clientConfig) UpdateGroup(g *Group) error {
	path := fmt.Sprintf("security/groups/%s", g.Name)
	resp, err := c.execute("POST", path, g)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Failed to update group. Status: %s", resp.Status)
	}

	return resp.Body.Close()
}

// DeleteGroup in Artifactory
func (c clientConfig) DeleteGroup(name string) error {
	path := fmt.Sprintf("security/groups/%s", name)
	resp, err := c.execute("DELETE", path, nil)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Failed to delete group. Status: %s", resp.Status)
	}

	return resp.Body.Close()
}

func (c clientConfig) execute(method string, endpoint string, payload interface{}) (*http.Response, error) {
	client := &http.Client{}
	url := fmt.Sprintf("%s/api/%s", c.url, endpoint)

	var jsonpayload *bytes.Buffer
	if payload == nil {
		jsonpayload = &bytes.Buffer{}
	} else {
		var jsonbuffer []byte
		jsonpayload = bytes.NewBuffer(jsonbuffer)
		enc := json.NewEncoder(jsonpayload)
		enc.Encode(payload)
	}

	req, err := http.NewRequest(method, url, jsonpayload)
	if err != nil {
		log.Printf("[ERROR] Error creating new request: %s", err)
		return nil, err
	}
	req.SetBasicAuth(c.user, c.pass)
	req.Header.Add("content-type", "application/json")

	return client.Do(req)
}
