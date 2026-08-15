// Copyright IBM Corp. 2021, 2026
// SPDX-License-Identifier: MPL-2.0

package provider


type HttpProviderConfig struct {
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/http/3.6.1/docs#alias HttpProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
}

