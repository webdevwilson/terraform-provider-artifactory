package artifactory

import (
	"log"
	"os"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/helper/validation"
)

func resourceVirtualRepository() *schema.Resource {
	return &schema.Resource{
		Create: resourceVirtualRepositoryCreate,
		Read:   resourceVirtualRepositoryRead,
		Update: resourceVirtualRepositoryUpdate,
		Delete: resourceVirtualRepositoryDelete,
		Importer: &schema.ResourceImporter{
			State: virtualRepositoryImportStatePassthrough,
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
			"repositories": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Optional: true,
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
			"debian_trivial_layout": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"repo_layout_ref": &schema.Schema{
				Type:     schema.TypeString,
				Default:  "maven-2-default",
				Optional: true,
			},
			"artifactory_requests_can_retrieve_remote_artifacts": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"key_pair": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pom_repository_references_cleanup_policy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "discard_active_reference",
			},
			"default_deployment_repo": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func newVirtualRepositoryFromResource(d *schema.ResourceData) *VirtualRepositoryConfiguration {
	repos := make([]string, 0, len(d.Get("repositories").(*schema.Set).List()))

	for _, r := range d.Get("repositories").(*schema.Set).List() {
		repos = append(repos, r.(string))
	}
	os.Stdout.WriteString("newVirtualRepositoryFromResource")

	return &VirtualRepositoryConfiguration{
		Key:                                           d.Get("key").(string),
		RClass:                                        "virtual",
		PackageType:                                   d.Get("package_type").(string),
		Repositories:                                  repos,
		Description:                                   d.Get("description").(string),
		Notes:                                         d.Get("notes").(string),
		IncludesPattern:                               d.Get("includes_pattern").(string),
		ExcludesPattern:                               d.Get("excludes_pattern").(string),
		ArtifactoryRequestsCanRetrieveRemoteArtifacts: d.Get("artifactory_requests_can_retrieve_remote_artifacts").(bool),
		KeyPair: d.Get("key_pair").(string),
		PomRepositoryReferencesCleanupPolicy: d.Get("pom_repository_references_cleanup_policy").(string),
		DefaultDeploymentRepo:                d.Get("default_deployment_repo").(string),
	}
}

func resourceVirtualRepositoryCreate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[TRACE] Creating artifactory.virtual_repository Id=%s\n", d.Get("key"))
	c := m.(Client)
	repo := newVirtualRepositoryFromResource(d)
	err := c.CreateRepository(repo.Key, repo)

	if err != nil {
		return err
	}

	d.SetId(repo.Key)
	return resourceVirtualRepositoryRead(d, m)
}

func resourceVirtualRepositoryRead(d *schema.ResourceData, m interface{}) error {
	log.Printf("[TRACE] Reading artifactory.virtual_repository Id=%s\n", d.Id())
	c := m.(Client)
	key := d.Id()
	var repo VirtualRepositoryConfiguration

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
	d.Set("artifactory_requests_can_retrieve_remote_artifacts", repo.ArtifactoryRequestsCanRetrieveRemoteArtifacts)
	d.Set("key_pair", repo.KeyPair)
	d.Set("pom_repository_references_cleanup_policy", repo.PomRepositoryReferencesCleanupPolicy)
	d.Set("default_deployment_repo", repo.DefaultDeploymentRepo)

	repos := make([]string, 0, len(repo.Repositories))
	for _, r := range repo.Repositories {
		repos = append(repos, r)
	}
	d.Set("repositories", repos)

	return nil
}

func resourceVirtualRepositoryUpdate(d *schema.ResourceData, m interface{}) error {
	log.Printf("[TRACE] Updating artifactory.virtual_repository Id=%s\n", d.Id())
	c := m.(Client)
	repo := newVirtualRepositoryFromResource(d)
	err := c.UpdateRepository(repo.Key, repo)
	if err != nil {
		return err
	}
	return resourceVirtualRepositoryRead(d, m)
}

func resourceVirtualRepositoryDelete(d *schema.ResourceData, m interface{}) error {
	log.Printf("[TRACE] Deleting artifactory.virtual_repository Id=%s\n", d.Id())
	c := m.(Client)
	key := d.Id()
	return c.DeleteRepository(key)
}

func virtualRepositoryImportStatePassthrough(d *schema.ResourceData, m interface{}) (s []*schema.ResourceData, err error) {
	log.Printf("[DEBUG] Importing state!")
	s, err = schema.ImportStatePassthrough(d, m)
	log.Printf("[DEBUG] Done importing state!")
	return
}
