package artifactory

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/webdevwilson/go-artifactory/artifactory"
)

func resourceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceGroupCreate,
		Read:   resourceGroupRead,
		Update: resourceGroupUpdate,
		Delete: resourceGroupDelete,
		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"auto_join": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"realm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"realm_attributes": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func newGroupFromResource(d *schema.ResourceData) *artifactory.Group {
	// Artifactory defaults to admin, let's not do that
	group := &artifactory.Group{}

	if v, ok := d.GetOk("name"); ok {
		group.Name = v.(string)
	}

	if v, ok := d.GetOk("auto_join"); ok {
		group.AutoJoin = v.(bool)
	}

	if v, ok := d.GetOk("realm"); ok {
		group.Realm = v.(string)
	}

	if v, ok := d.GetOk("realm_attributes"); ok {
		group.RealmAttributes = v.(string)
	}

	return group
}

func resourceGroupRead(d *schema.ResourceData, m interface{}) error {
	c := m.(artifactory.Client)

	group, err := c.GetGroup(d.Get("name").(string))

	if err != nil {
		return err
	}

	d.Set("name", group.Name)
	d.Set("auto_join", group.AutoJoin)
	d.Set("realm", group.Realm)
	d.Set("realm_attributes", group.RealmAttributes)

	return nil
}

func resourceGroupCreate(d *schema.ResourceData, m interface{}) error {
	group := newGroupFromResource(d)
	c := m.(artifactory.Client)
	err := c.CreateGroup(group)

	if err != nil {
		return err
	}

	d.SetId(group.Name)
	return resourceGroupRead(d, m)
}

func resourceGroupUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(artifactory.Client)
	group := newGroupFromResource(d)
	err := c.UpdateGroup(group)

	if err != nil {
		return err
	}

	return resourceGroupRead(d, m)
}

func resourceGroupDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(artifactory.Client)
	group := newGroupFromResource(d)
	return c.DeleteGroup(group.Name)
}
