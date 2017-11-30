package artifactory

import (
	"encoding/json"
	"fmt"
)

// Group represents an Artifactory group
type Group struct {
	Name            string `json:"name,omitempty"`
	AutoJoin        bool   `json:"autoJoin"`
	Realm           string `json:"realm,omitempty"`
	RealmAttributes string `json:"realm,omitempty"`
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

	if err := c.validateResponse(201, resp.StatusCode, "create group"); err != nil {
		return err
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

	if err := c.validateResponse(200, resp.StatusCode, "update group"); err != nil {
		return err
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

	if err := c.validateResponse(200, resp.StatusCode, "delete group"); err != nil {
		return err
	}

	return resp.Body.Close()
}
