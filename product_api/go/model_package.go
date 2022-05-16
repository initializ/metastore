/*
 * Metadata Store API.
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type ModelPackage struct {

	Homepage string `json:"Homepage,omitempty"`

	ID int32 `json:"ID,omitempty"`

	Images []Image `json:"Images,omitempty"`

	Name string `json:"Name,omitempty"`

	PackageManager string `json:"PackageManager,omitempty"`

	Sources []Source `json:"Sources,omitempty"`

	Version string `json:"Version,omitempty"`

	Vulnerabilities []Vulnerability `json:"Vulnerabilities,omitempty"`
}
