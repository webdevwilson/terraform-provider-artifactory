package artifactory

import "github.com/hashicorp/terraform/helper/schema"

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

func newGroupFromResource(d *schema.ResourceData) *Group {
	// Artifactory defaults to admin, let's not do that
	group := &Group{}

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
	c := m.(Client)

	user, err := c.GetGroup(d.Get("name").(string))

	if err != nil {
		return err
	}

	d.Set("name", user.Name)
	d.Set("auto_join", user.AutoJoin)
	d.Set("realm", user.Realm)
	d.Set("realm_attributes", user.RealmAttributes)

	return nil
}

func resourceGroupCreate(d *schema.ResourceData, m interface{}) error {
	group := newGroupFromResource(d)
	c := m.(Client)
	err := c.CreateGroup(group)

	if err != nil {
		return err
	}

	d.SetId(group.Name)
	return resourceGroupRead(d, m)
}

func resourceGroupUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(Client)
	group := newGroupFromResource(d)
	err := c.UpdateGroup(group)

	if err != nil {
		return err
	}

	return resourceGroupRead(d, m)
}

func resourceGroupDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(Client)
	group := newGroupFromResource(d)
	return c.DeleteGroup(group.Name)
}
