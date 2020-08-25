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

// Repo holds the App repository basic details
type Repo struct {
	Namespace string `json:"namespace"`
	Name string `json:"name"`
	URL string `json:"url"`
}

// RepoInternal holds the App repository details including auth and checksum
type RepoInternal struct {
	Namespace string `json:"namespace"`
	Name string `json:"name"`
	URL string `json:"url"`
	AuthorizationHeader string `bson:"-"`
	Checksum string
}