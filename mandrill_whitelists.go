// Copyright 2013 Matthew Baird
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//     http://www.apache.org/licenses/LICENSE-2.0
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gochimp

import (
	"errors"
	"log"
)

const whitelists_list_endpoint string = "/whitelists/list.json"     // Retrieves your email rejection whitelist. You can provide an email address or search prefix to limit the results. Returns up to 1000 results.
const whitelists_add_endpoint string = "/whitelists/add.json"       // Adds an email to your email rejection whitelist. If the address is currently on your blacklist, that blacklist entry will be removed automatically.
const whitelists_delete_endpoint string = "/whitelists/delete.json" // Removes an email address from the whitelist.

// WhitelistsList retrieves your email rejection whitelist. You can provide an email address or search prefix to limit the results. Returns up to 1000 results.
func (a *MandrillAPI) WhitelistsList(email string) ([]Whitelist, error) {
	var response []Whitelist
	if email == "" {
		return response, errors.New("email cannot be blank")
	}
	var params map[string]interface{} = make(map[string]interface{})
	params["email"] = email
	err := parseMandrillJson(a, whitelists_list_endpoint, params, &response)
	return response, err
}

// WhitelistsAdd can error with one of the following: Invalid_Key, ValidationError, GeneralError
func (a *MandrillAPI) WhitelistsAdd(email string, comment string) (bool, error) {
	var response map[string]interface{}
	retval := false
	if email == "" {
		return retval, errors.New("email cannot be blank")
	}
	if comment == "" {
		return retval, errors.New("comment cannot be blank")
	}
	var params map[string]interface{} = make(map[string]interface{})
	params["email"] = email
	params["comment"] = comment

	err := parseMandrillJson(a, whitelists_add_endpoint, params, &response)
	ok := false
	if err == nil {
		retval, ok = response["added"].(bool)
		if ok != true {
			log.Fatal("Received response with added parameter, however type was not bool, this should not happen")
		}
	}
	return retval, err
}

// WhitelistsDelete removes an email address from the whitelist.
// can error with one of the following: Invalid_Reject, Invalid_Key, ValidationError, GeneralError
func (a *MandrillAPI) WhitelistsDelete(email string) (bool, error) {
	var response map[string]interface{}
	retval := false
	if email == "" {
		return retval, errors.New("email cannot be blank")
	}
	var params map[string]interface{} = make(map[string]interface{})
	params["email"] = email
	err := parseMandrillJson(a, whitelists_delete_endpoint, params, &response)
	ok := false
	if err == nil {
		retval, ok = response["deleted"].(bool)
		if ok != true {
			log.Fatal("Received response with deleted parameter, however type was not bool, this should not happen")
		}
	}
	return retval, err
}

type Whitelist struct {
	Email     string  `json:"email"`
	Detail    string  `json:"detail"`
	CreatedAt APITime `json:"created_at"`
}
