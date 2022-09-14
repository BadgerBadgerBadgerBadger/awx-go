package awx

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// CredentialService implements awx credential apis.
type CredentialService struct {
	client *Client
}

// ListCredentialsResponse represents `ListCredentials` endpoint response.
type ListCredentialsResponse struct {
	Pagination
	Results []*Credential `json:"results"`
}

// ListCredentials shows a list of credentials.
func (c *CredentialService) ListCredentials(params map[string]string) ([]*Credential, *ListCredentialsResponse, error) {
	result := new(ListCredentialsResponse)
	endpoint := "/api/v2/credentials/"
	resp, err := c.client.Requester.GetJSON(endpoint, result, params)
	if err != nil {
		return nil, result, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, result, err
	}

	return result.Results, result, nil
}

// CreateCredential creates a credential
func (c *CredentialService) CreateCredential(data map[string]interface{}, params map[string]string) (*Credential, error) {
	result := new(Credential)
	mandatoryFields = []string{"name", "credential_type"}
	validate, status := ValidateParams(data, mandatoryFields)
	if !status {
		err := fmt.Errorf("Mandatory input arguments are absent: %s", validate)
		return nil, err
	}
	endpoint := "/api/v2/credentials/"
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Requester.PostJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}
	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}

// UpdateCredential updates a credential
func (c *CredentialService) UpdateCredential(id int, data map[string]interface{}, params map[string]string) (*Credential, error) {
	result := new(Credential)
	endpoint := fmt.Sprintf("/api/v2/credentials/%d", id)
	payload, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := c.client.Requester.PatchJSON(endpoint, bytes.NewReader(payload), result, params)
	if err != nil {
		return nil, err
	}
	if err := CheckResponse(resp); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteCredential deletes a credential
func (c *CredentialService) DeleteCredential(id int) (*Credential, error) {
	result := new(Credential)
	endpoint := fmt.Sprintf("/api/v2/credentials/%d", id)

	resp, err := c.client.Requester.Delete(endpoint, result, nil)
	if err != nil {
		return nil, err
	}

	if err := CheckResponse(resp); err != nil {
		return nil, err
	}

	return result, nil
}
