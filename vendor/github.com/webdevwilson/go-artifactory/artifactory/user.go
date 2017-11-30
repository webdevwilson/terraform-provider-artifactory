package artifactory

import (
	"encoding/json"
	"fmt"
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

	if err := c.validateResponse(201, resp.StatusCode, "create user"); err != nil {
		return err
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

	if err := c.validateResponse(200, resp.StatusCode, "update user"); err != nil {
		return err
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

	if err := c.validateResponse(200, resp.StatusCode, "delete user"); err != nil {
		return err
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
