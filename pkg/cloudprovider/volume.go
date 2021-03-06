/*
Copyright 2016 The Archon Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cloudprovider

import (
	"encoding/json"
)

// VolumeError is the json response of an volume operation
type VolumeError struct {
	Message string `json:"message"`
	Status  string `json:"status"`
	Device  string `json:"device"`
}

func (v VolumeError) Error() string {
	return v.Message
}

// NewVolumeError creates failure error with given message
func NewVolumeError(msg string) VolumeError {
	return VolumeError{msg, "Failure", ""}
}

var (
	VolumeSuccess = VolumeError{"", "Success", ""}
)

func (v VolumeError) ToJson() string {
	ret, _ := json.Marshal(&v)
	return string(ret)
}

// Json data
type VolumeOptions map[string]interface{}

// Volume interface of flex volume plugin
type Volume interface {
	Init() error
	Attach(options VolumeOptions) error
	Detach(device string) error
	Mount(dir, device string, options VolumeOptions) error
	Unmount(dir string) error
}
