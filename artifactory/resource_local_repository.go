package artifactory

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceLocalRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceLocalRepositoryCreate,
		Read:   resourceLocalRepositoryRead,
		Update: resourceLocalRepositoryUpdate,
		Delete: resourceLocalRepositoryDelete,
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
				Default:      "generic",
				ValidateFunc: validation.StringInSlice(packageTypes, true),
				ForceNew:     true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"notes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"includes_pattern": &schema.Schema{
				Type:     schema.TypeString,
				Default:  "**/*",
				Optional: true,
			},
			"excludes_pattern": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"repo_layout_ref": &schema.Schema{
				Type:     schema.TypeString,
				Default:  "maven-2-default",
				Optional: true,
			},
			"handle_releases": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"handle_snapshots": &schema.Schema{
				Type:     schema.TypeBool,
				Default:  true,
				Optional: true,
			},
			"max_unique_snapshots": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"debian_trivial_layout": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"checksum_policy_type": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "client-checksums",
				ValidateFunc: validation.StringInSlice(checksumPolicyTypes, true),
			},
			"max_unique_tags": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"snapshot_version_behavior": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "non-unique",
				ValidateFunc: validation.StringInSlice(snapshotVersionBehaviors, true),
			},
			"suppress_pom_consistency_checks": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"blacked_out": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"property_sets": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Optional: true,
			},
			"archive_browsing_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"calculate_yum_metadata": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"yum_root_depth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"docker_api_version": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "V2",
				ValidateFunc: validation.StringInSlice([]string{"V1", "V2"}, true),
			},
		},
	}
}

func newLocalRepositoryFromResource(d *schema.ResourceData) *LocalRepositoryConfiguration {

	props := make([]string, 0, len(d.Get("property_sets").(*schema.Set).List()))

	for _, p := range d.Get("property_sets").(*schema.Set).List() {
		props = append(props, p.(string))
	}

	return &LocalRepositoryConfiguration{
		Key:                          d.Get("key").(string),
		RClass:                       "local",
		PackageType:                  d.Get("package_type").(string),
		Description:                  d.Get("description").(string),
		Notes:                        d.Get("notes").(string),
		IncludesPattern:              d.Get("includes_pattern").(string),
		ExcludesPattern:              d.Get("excludes_pattern").(string),
		RepoLayoutRef:                d.Get("repo_layout_ref").(string),
		HandleReleases:               d.Get("handle_releases").(bool),
		HandleSnapshots:              d.Get("handle_snapshots").(bool),
		MaxUniqueSnapshots:           d.Get("max_unique_snapshots").(int),
		DebianTrivialLayout:          d.Get("debian_trivial_layout").(bool),
		ChecksumPolicyType:           d.Get("checksum_policy_type").(string),
		MaxUniqueTags:                d.Get("max_unique_tags").(int),
		SnapshotVersionBehavior:      d.Get("snapshot_version_behavior").(string),
		SuppressPomConsistencyChecks: d.Get("suppress_pom_consistency_checks").(bool),
		BlackedOut:                   d.Get("blacked_out").(bool),
		ArchiveBrowsingEnabled:       d.Get("archive_browsing_enabled").(bool),
		CalculateYumMetadata:         d.Get("calculate_yum_metadata").(bool),
		YumRootDepth:                 d.Get("yum_root_depth").(int),
		DockerAPIVersion:             d.Get("docker_api_version").(string),
		PropertySets:                 props,
	}
}

func resourceLocalRepositoryCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(Client)
	repo := newLocalRepositoryFromResource(d)

	err := c.CreateRepository(repo.Key, repo)

	if err != nil {
		return err
	}

	d.SetId(repo.Key)
	return resourceLocalRepositoryRead(d, m)
}

func resourceLocalRepositoryRead(d *schema.ResourceData, m interface{}) error {
	c := m.(Client)
	key := d.Id()

	var repo LocalRepositoryConfiguration

	err := c.GetRepository(key, &repo)

	if err != nil {
		return err
	}

	d.Set("key", repo.Key)
	d.Set("type", repo.RClass)
	d.Set("package_type", repo.PackageType)
	d.Set("description", repo.Description)
	d.Set("notes", repo.Notes)
	d.Set("includes_pattern", repo.IncludesPattern)
	d.Set("excludes_pattern", repo.ExcludesPattern)
	d.Set("repo_layout_ref", repo.RepoLayoutRef)
	d.Set("handle_releases", repo.HandleReleases)
	d.Set("handle_snapshots", repo.HandleSnapshots)
	d.Set("max_unique_snapshots", repo.MaxUniqueSnapshots)
	d.Set("debian_trivial_layout", repo.DebianTrivialLayout)
	d.Set("checksum_policy_type", repo.ChecksumPolicyType)
	d.Set("max_unique_tags", repo.MaxUniqueTags)
	d.Set("snapshot_version_behavior", repo.SnapshotVersionBehavior)
	d.Set("suppress_pom_consistency_checks", repo.SuppressPomConsistencyChecks)
	d.Set("blacked_out", repo.BlackedOut)
	d.Set("archive_browsing_enabled", repo.ArchiveBrowsingEnabled)
	d.Set("calculate_yum_metadata", repo.CalculateYumMetadata)
	d.Set("yum_root_depth", repo.YumRootDepth)
	d.Set("docker_api_version", repo.DockerAPIVersion)

	props := make([]string, 0, len(repo.PropertySets))
	for _, p := range repo.PropertySets {
		props = append(props, p)
	}
	d.Set("property_sets", props)

	return nil
}

func resourceLocalRepositoryUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(Client)
	repo := newLocalRepositoryFromResource(d)
	err := c.UpdateRepository(repo.Key, repo)
	if err != nil {
		return err
	}
	return resourceLocalRepositoryRead(d, m)
}

func resourceLocalRepositoryDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(Client)
	key := d.Id()
	return c.DeleteRepository(key)
}
