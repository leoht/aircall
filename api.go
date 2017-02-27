package aircall

import (
	"encoding/json"
	"strconv"
)

// Ping API request
func (client *Client) Ping() (*PingResponse, error) {
	data, err := client.Get("/ping", map[string]string{})
	if err != nil {
		return nil, err
	}

	response := &PingResponse{}
	json.Unmarshal(data, response)

	return response, nil
}

// Company API request
//
// res, err := client.Company()
func (client *Client) Company() (CompanyResponse, error) {
	data, err := client.Get("/company", map[string]string{})
	response := CompanyResponse{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// Users API request
func (client *Client) Users() (UsersResponse, error) {
	data, err := client.Get("/users", map[string]string{})
	response := UsersResponse{}

	if err != nil {
		return response, err
	}

	json.Unmarshal(data, &response)

	return response, nil
}

// User API request
func (client *Client) User(ID int) (*UserResponse, error) {
	data, err := client.Get("/users/"+strconv.Itoa(ID), map[string]string{})
	if err != nil {
		return nil, err
	}

	response := &UserResponse{}
	json.Unmarshal(data, response)

	return response, nil
}

// Numbers API request
func (client *Client) Numbers() (*NumbersResponse, error) {
	data, err := client.Get("/numbers", map[string]string{})
	if err != nil {
		return nil, err
	}

	response := &NumbersResponse{}
	json.Unmarshal(data, response)

	return response, nil
}

// Number API request
func (client *Client) Number(ID int) (*NumberResponse, error) {
	data, err := client.Get("/numbers/"+strconv.Itoa(ID), map[string]string{})
	if err != nil {
		return nil, err
	}

	response := &NumberResponse{}
	json.Unmarshal(data, response)

	return response, nil
}

// Calls API request
func (client *Client) Calls() (*CallsResponse, error) {
	data, err := client.Get("/calls", map[string]string{})
	if err != nil {
		return nil, err
	}

	response := &CallsResponse{}
	json.Unmarshal(data, response)

	return response, nil
}

// Call API request
func (client *Client) Call(ID int) (*CallResponse, error) {
	data, err := client.Get("/calls/"+strconv.Itoa(ID), map[string]string{})
	if err != nil {
		return nil, err
	}

	response := &CallResponse{}
	json.Unmarshal(data, response)

	return response, nil
}

// SearchCalls API request
func (client *Client) SearchCalls(search *SearchRequest) (*CallsResponse, error) {
	params := buildSearchParamsMap(search)
	data, err := client.Get("/calls", params)
	if err != nil {
		return nil, err
	}

	response := &CallsResponse{}
	json.Unmarshal(data, response)

	return response, nil
}

// TransferCall API request
func (client *Client) TransferCall(ID int, userID int) (*CallResponse, error) {
	data, err := client.Post("/calls/"+strconv.Itoa(ID)+"/transfers", TransferCallRequest{
		UserID: userID,
	})

	if err != nil {
		return nil, err
	}

	response := &CallResponse{}
	json.Unmarshal(data, response)

	return response, nil
}

// DeleteRecording API request
func (client *Client) DeleteRecording(callID int) (*Response, error) {
	data, err := client.Delete("/calls/"+strconv.Itoa(callID)+"/recording", map[string]string{})

	if err != nil {
		return nil, err
	}

	response := &Response{}
	json.Unmarshal(data, response)

	return response, nil
}

// DeleteVoicemail API request
func (client *Client) DeleteVoicemail(callID int) (*Response, error) {
	data, err := client.Delete("/calls/"+strconv.Itoa(callID)+"/voicemail", map[string]string{})

	if err != nil {
		return nil, err
	}

	response := &Response{}
	json.Unmarshal(data, response)

	return response, nil
}

// Contacts API request
func (client *Client) Contacts() (*ContactsResponse, error) {
	data, err := client.Get("/contacts", map[string]string{})
	if err != nil {
		return nil, err
	}

	response := &ContactsResponse{}
	json.Unmarshal(data, response)

	return response, nil
}

// Contact API request
func (client *Client) Contact(ID int) (*ContactResponse, error) {
	data, err := client.Get("/contacts/"+strconv.Itoa(ID), map[string]string{})
	if err != nil {
		return nil, err
	}

	response := &ContactResponse{}
	json.Unmarshal(data, response)

	return response, nil
}

// CreateContact API request
func (client *Client) CreateContact(contact *CreateContactRequest) (*ContactResponse, error) {
	data, err := client.Post("/contacts", contact)
	if err != nil {
		return nil, err
	}

	response := &ContactResponse{}
	json.Unmarshal(data, response)

	return response, nil
}

// UpdateContact API request
func (client *Client) UpdateContact(ID int, contact *CreateContactRequest) (*ContactResponse, error) {
	data, err := client.Post("/contacts/"+strconv.Itoa(ID), contact)
	if err != nil {
		return nil, err
	}

	response := &ContactResponse{}
	json.Unmarshal(data, response)

	return response, nil
}

// SearchContacts API request
func (client *Client) SearchContacts(search *SearchRequest) (*ContactsResponse, error) {

	params := buildSearchParamsMap(search)
	data, err := client.Get("/contacts/search", params)
	if err != nil {
		return nil, err
	}

	response := &ContactsResponse{}
	json.Unmarshal(data, response)

	return response, nil
}

// DeleteContact API request
func (client *Client) DeleteContact(ID int) (*Response, error) {
	data, err := client.Delete("/contacts/"+strconv.Itoa(ID), map[string]string{})

	if err != nil {
		return nil, err
	}

	response := &Response{}
	json.Unmarshal(data, response)

	return response, nil
}

func buildSearchParamsMap(search *SearchRequest) map[string]string {
	var params map[string]string

	if search.Page > 0 {
		params["page"] = strconv.Itoa(search.Page)
	}
	if search.PerPage > 0 {
		params["per_page"] = strconv.Itoa(search.Page)
	}
	if search.Order != "" {
		params["order"] = search.Order
	}
	if search.From > 0 {
		params["from"] = strconv.Itoa(search.From)
	}
	if search.To > 0 {
		params["to"] = strconv.Itoa(search.To)
	}
	if search.PhoneNumber != "" {
		params["phone_number"] = search.PhoneNumber
	}
	if search.Email != "" {
		params["email"] = search.Email
	}

	return params
}
