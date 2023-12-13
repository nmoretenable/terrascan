/*
    Copyright (C) 2022 Tenable, Inc.

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

package config

import (
	"github.com/awslabs/goformation/v7/cloudformation/ssm"
)

// SSMParameterConfig holds config for SSMParameter
type SSMParameterConfig struct {
	Config
	Name           string `json:"name"`
	Description    string `json:"description"`
	Type           string `json:"type"`
	Value          string `json:"value"`
	Tier           string `json:"tier"`
	Policies       string `json:"policies"`
	AllowedPattern string `json:"allowed_pattern"`
}

// GetSSMParameterConfig returns config for SSM Parameter
func GetSSMParameterConfig(b *ssm.Parameter) []AWSResourceConfig {
	cf := SSMParameterConfig{
		Config: Config{
			Name: *b.Name,
			Tags: b.Tags,
		},
		Name:  *b.Name,
		Type:  b.Type,
		Value: b.Value,
	}
	if b.Description != nil {
		cf.Description = *b.Description
	}
	if b.Tier != nil {
		cf.Tier = *b.Tier
	}
	if b.Policies != nil {
		cf.Policies = *b.Policies
	}
	if b.AllowedPattern != nil {
		cf.AllowedPattern = *b.AllowedPattern
	}

	return []AWSResourceConfig{{
		Resource: cf,
		Metadata: b.AWSCloudFormationMetadata,
	}}
}
