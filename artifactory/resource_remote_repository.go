package artifactory

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceRemoteRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceRemoteRepositoryCreate,
		Read:   resourceRemoteRepositoryRead,
		Update: resourceRemoteRepositoryUpdate,
		Delete: resourceRemoteRepositoryDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"package_type": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(packageTypes, true),
				Default:      "generic",
			},
			"description": &schema.Schema{
				Type:             schema.TypeString,
				Optional:         true,
				Default:          "(local file cache)",
				DiffSuppressFunc: resourceRemoteDescriptionDiffSuppress,
			},
			"notes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"includes_pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "**/*",
			},
			"excludes_pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"repo_layout_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "maven-2-default",
			},
			"handle_releases": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"handle_snapshots": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"max_unique_snapshots": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"suppress_pom_consistency_checks": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"url": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"proxy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_repo_checksum_policy_type": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringInSlice(remoteRepoChecksumPolicyTypes, true),
				Optional:     true,
			},
			"hard_fail": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"offline": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"blacked_out": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"store_artifacts_locally": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"socket_timeout_millis": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  15000,
			},
			"local_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"retrieval_cache_period_seconds": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  600,
			},
			"failed_cache_period_seconds": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"missed_cache_period_seconds": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1800,
			},
			"unused_artifacts_cleanup_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"unused_artifacts_cleanup_period_hours": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"fetch_jars_eagerly": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"fetch_sources_eagerly": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"share_configuration": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"synchronize_properties": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"property_sets": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Optional: true,
			},
			"allow_any_host_auth": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enable_cookie_management": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"bower_registry_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vcs_type": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(vcsType, true),
			},
			"vcs_git_provider": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringInSlice(vcsGitProviders, true),
			},
			"vcs_git_download_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func newRemoteRepositoryFromResource(d *schema.ResourceData) *RemoteRepositoryConfiguration {
	props := make([]string, 0, len(d.Get("property_sets").(*schema.Set).List()))

	for _, p := range d.Get("property_sets").(*schema.Set).List() {
		props = append(props, p.(string))
	}
	return &RemoteRepositoryConfiguration{
		Key:                               d.Get("key").(string),
		RClass:                            "remote",
		PackageType:                       d.Get("package_type").(string),
		URL:                               d.Get("url").(string),
		Username:                          d.Get("username").(string),
		Password:                          d.Get("password").(string),
		Proxy:                             d.Get("proxy").(string),
		Description:                       d.Get("description").(string),
		Notes:                             d.Get("notes").(string),
		IncludesPattern:                   d.Get("includes_pattern").(string),
		ExcludesPattern:                   d.Get("excludes_pattern").(string),
		RepoLayoutRef:                     d.Get("repo_layout_ref").(string),
		HandleReleases:                    d.Get("handle_releases").(bool),
		HandleSnapshots:                   d.Get("handle_snapshots").(bool),
		MaxUniqueSnapshots:                d.Get("max_unique_snapshots").(int),
		SuppressPomConsistencyChecks:      d.Get("suppress_pom_consistency_checks").(bool),
		RemoteRepoChecksumPolicyType:      d.Get("remote_repo_checksum_policy_type").(string),
		HardFail:                          d.Get("hard_fail").(bool),
		Offline:                           d.Get("offline").(bool),
		BlackedOut:                        d.Get("blacked_out").(bool),
		StoreArtifactsLocally:             d.Get("store_artifacts_locally").(bool),
		SocketTimeoutMillis:               d.Get("socket_timeout_millis").(int),
		LocalAddress:                      d.Get("local_address").(string),
		RetrievalCachePeriodSeconds:       d.Get("retrieval_cache_period_seconds").(int),
		FailedCachePeriodSeconds:          d.Get("failed_cache_period_seconds").(int),
		MissedCachePeriodSeconds:          d.Get("missed_cache_period_seconds").(int),
		UnusedArtifactsCleanupEnabled:     d.Get("unused_artifacts_cleanup_enabled").(bool),
		UnusedArtifactsCleanupPeriodHours: d.Get("unused_artifacts_cleanup_period_hours").(int),
		FetchJarsEagerly:                  d.Get("fetch_jars_eagerly").(bool),
		FetchSourcesEagerly:               d.Get("fetch_sources_eagerly").(bool),
		ShareConfiguration:                d.Get("share_configuration").(bool),
		SynchronizeProperties:             d.Get("synchronize_properties").(bool),
		PropertySets:                      props,
		AllowAnyHostAuth:                  d.Get("allow_any_host_auth").(bool),
		EnableCookieManagement:            d.Get("enable_cookie_management").(bool),
		BowerRegistryURL:                  d.Get("bower_registry_url").(string),
		VCSType:                           d.Get("vcs_type").(string),
		VCSGitProvider:                    d.Get("vcs_git_provider").(string),
		VCSGitDownloadURL:                 d.Get("vcs_git_download_url").(string),
	}
}

func resourceRemoteRepositoryCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(Client)
	repo := newRemoteRepositoryFromResource(d)

	err := c.CreateRepository(repo.Key, repo)

	if err != nil {
		return err
	}

	d.SetId(repo.Key)
	return resourceRemoteRepositoryRead(d, m)
}

func resourceRemoteRepositoryRead(d *schema.ResourceData, m interface{}) error {
	c := m.(Client)
	key := d.Id()
	var repo RemoteRepositoryConfiguration

	err := c.GetRepository(key, &repo)

	if err != nil {
		return err
	}

	d.Set("key", repo.Key)
	d.Set("type", repo.RClass)
	d.Set("package_type", repo.PackageType)
	d.Set("url", repo.URL)
	d.Set("username", repo.Username)
	d.Set("password", repo.Password)
	d.Set("proxy", repo.Proxy)
	d.Set("description", repo.Description)
	d.Set("notes", repo.Notes)
	d.Set("includes_pattern", repo.IncludesPattern)
	d.Set("excludes_pattern", repo.ExcludesPattern)
	d.Set("repo_layout_ref", repo.RepoLayoutRef)
	d.Set("handle_releases", repo.HandleReleases)
	d.Set("handle_snapshots", repo.HandleSnapshots)
	d.Set("max_unique_snapshots", repo.MaxUniqueSnapshots)
	d.Set("remote_repo_checksum_policy_type", repo.RemoteRepoChecksumPolicyType)
	d.Set("hard_fail", repo.HardFail)
	d.Set("offline", repo.Offline)
	d.Set("blacked_out", repo.BlackedOut)
	d.Set("store_artifacts_locally", repo.StoreArtifactsLocally)
	d.Set("socket_timeout_millis", repo.SocketTimeoutMillis)
	d.Set("local_address", repo.LocalAddress)
	d.Set("retrieval_cache_period_seconds", repo.RetrievalCachePeriodSeconds)
	d.Set("failed_cache_period_seconds", repo.FailedCachePeriodSeconds)
	d.Set("missed_cache_period_seconds", repo.MissedCachePeriodSeconds)
	d.Set("unused_artifacts_cleanup_enabled", repo.UnusedArtifactsCleanupEnabled)
	d.Set("unused_artifacts_cleanup_period_hours", repo.UnusedArtifactsCleanupPeriodHours)
	d.Set("fetch_jars_eagerly", repo.FetchJarsEagerly)
	d.Set("fetch_sources_eagerly", repo.FetchSourcesEagerly)
	d.Set("share_configuration", repo.ShareConfiguration)
	d.Set("synchronize_properties", repo.SynchronizeProperties)
	d.Set("allow_any_host_auth", repo.AllowAnyHostAuth)
	d.Set("enable_cookie_management", repo.EnableCookieManagement)
	d.Set("bower_registry_url", repo.BowerRegistryURL)
	d.Set("vcs_type", repo.VCSType)
	d.Set("vcs_git_provider", repo.VCSGitProvider)
	d.Set("vcs_git_download_url", repo.VCSGitDownloadURL)

	props := make([]string, 0, len(repo.PropertySets))
	for _, p := range repo.PropertySets {
		props = append(props, p)
	}
	d.Set("property_sets", props)

	return nil
}

func resourceRemoteRepositoryUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(Client)
	repo := newRemoteRepositoryFromResource(d)
	err := c.UpdateRepository(repo.Key, repo)
	if err != nil {
		return err
	}
	return resourceRemoteRepositoryRead(d, m)
}

func resourceRemoteRepositoryDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(Client)
	key := d.Get("key").(string)
	return c.DeleteRepository(key)
}

// resourceRemoteDescriptionDiffSuppress suppresses local file cache added to description
func resourceRemoteDescriptionDiffSuppress(k, old, new string, d *schema.ResourceData) bool {
	return old == fmt.Sprintf("%s (local file cache)", new)
}
