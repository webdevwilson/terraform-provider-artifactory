package artifactory

import (
	"math/rand"

	"github.com/hashicorp/terraform/helper/schema"
)

const randomPasswordLength = 16

func resourceUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserCreate,
		Read:   resourceUserRead,
		Update: resourceUserUpdate,
		Delete: resourceUserDelete,
		Importer: &schema.ResourceImporter{
			State: virtualRepositoryImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"email": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"is_admin": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_editable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Set:      schema.HashString,
				Optional: true,
			},
			"realm": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func newUserFromResource(d *schema.ResourceData) *User {
	// Artifactory defaults to admin, let's not do that
	user := &User{}

	if v, ok := d.GetOk("name"); ok {
		user.Name = v.(string)
	}

	if v, ok := d.GetOk("email"); ok {
		user.Email = v.(string)
	}

	if v, ok := d.GetOk("is_admin"); ok {
		user.Admin = v.(bool)
	}

	if v, ok := d.GetOk("is_editable"); ok {
		user.ProfileUpdatable = v.(bool)
	}

	if v, ok := d.GetOk("realm"); ok {
		user.Realm = v.(string)
	}

	// create a random password
	user.Password = generatePassword()

	if v, ok := d.GetOk("groups"); ok {
		l := v.(*schema.Set).List()
		groups := make([]string, 0, len(l))
		for _, g := range l {
			groups = append(groups, g.(string))
		}
		user.Groups = groups
	}

	// Admin user's cannot have unupdatable profile
	if user.Admin {
		user.ProfileUpdatable = true
	}

	return user
}

func resourceUserRead(d *schema.ResourceData, m interface{}) error {
	c := m.(Client)

	user, err := c.GetUser(d.Get("name").(string))

	if err != nil {
		return err
	}

	d.Set("name", user.Name)
	d.Set("email", user.Email)
	d.Set("is_admin", user.Admin)
	d.Set("is_editable", user.ProfileUpdatable)
	d.Set("realm", user.Realm)

	return nil
}

func resourceUserCreate(d *schema.ResourceData, m interface{}) error {
	user := newUserFromResource(d)
	c := m.(Client)
	err := c.CreateUser(user)

	if err != nil {
		return err
	}

	d.SetId(user.Name)
	return resourceUserRead(d, m)
}

func resourceUserUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(Client)
	user := newUserFromResource(d)
	err := c.UpdateUser(user)

	if err != nil {
		return err
	}

	err = c.ExpireUserPassword(user.Name)

	if err != nil {
		return err
	}

	return resourceUserRead(d, m)
}

func resourceUserDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(Client)
	user := newUserFromResource(d)
	return c.DeleteUser(user.Name)
}

// generatePassword used as default func to generate user passwords
func generatePassword() string {
	letters := []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]byte, randomPasswordLength)
	for i := range b {
		b[i] = letters[rand.Int63()%int64(len(letters))]
	}
	return string(b)
}
