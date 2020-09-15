/*
Copyright (c) 2017 The Helm Authors
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
package models

import (
	"encoding/json"
	"time"

	"database/sql/driver"

	"k8s.io/helm/pkg/proto/hapi/chart"
)

// Repo holds the App repository basic details
type Repo struct {
	Namespace string `json:"namespace"`
	Name      string `json:"name"`
	URL       string `json:"url"`
}

// RepoInternal holds the App repository details including auth and checksum
type RepoInternal struct {
	Namespace           string `json:"namespace"`
	Name                string `json:"name"`
	URL                 string `json:"url"`
	AuthorizationHeader string `bson:"-"`
	Checksum            string
}

// Chart is a higher-level representation of a chart package
type Chart struct {
	ID            string              `json:"ID" bson:"chart_id"`
	Name          string              `json:"name"`
	Repo          *Repo               `json:"repo"`
	Description   string              `json:"description"`
	Home          string              `json:"home"`
	Keywords      []string            `json:"keywords"`
	Maintainers   []chart.Maintainers `json:"maintainers"`
	Sources       []string            `json:"sources"`
	Icon          string              `json:"icon"`
	RawIcon       []byte              `json:"icon_content_type" bson:"icon_content_type,omitempty"`
	Category      string              `json:"category"`
	ChartVersions []ChartVersion      `json:"chartVersions"`
}

// ChartIconString is a higher-level representation of a chart package
// TODO(andresmgot) Replace this type when the icon is stored as a binary
type ChartIconString struct {
	Chart
	RawIcon string `json:"raw_icon" bson:"raw_icon"`
}

// ChartVersion is a representation of a specific version of a chart
type ChartVersion struct {
	Version    string    `json:"version"`
	AppVersion string    `json:"app_version"`
	Created    time.Time `json:"created"`
	Digest     string    `json:"digest"`
	URLs       []string  `json:"urls"`
	// The following three fields get set with the URL paths to the respective
	// chart files (as opposed to the similar fields on ChartFiles which
	// contain the actual content).
	Readme string `json:"readme" bson:"-"`
	Values string `json:"values" bson:"-"`
	Schema string `json:"schema" bson:"-"`
}

// ChartFiles holds the README and values for a given chart version
type ChartFiles struct {
	ID     string `bson:"file_id"`
	Readme string
	Values string
	Schema string
	Repo   *Repo
	Digest string
}

// Allow to convert ChartFiles to a sql JSON
func (a ChartFiles) Value() (driver.Value, error) {
	return json.Marshal(a)
}

type RepoCheck struct {
	ID         string    `bson:"_id"`
	LastUpdate time.Time `bson:"last_update"`
	Checksum   string    `bson:"checksum"`
}
