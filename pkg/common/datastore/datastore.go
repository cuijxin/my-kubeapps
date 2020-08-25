/*
Copyright (c) 2016-2017 Bitnami
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

// Package datastore implements an interface on top of the mgo mongo client
package datastore

import (
	"time"

	"github.com/globalsign/mgo"
)

const defaultTimeout = 30 * time.Second

// Config configures the database connection
type Config struct {
	URL      string
	Database string
	Username string
	Password string
	Timeout  time.Duration
}

// Session is an interface for a MongoDB session
type Session interface {
	DB() (Database, func())
}

// Database is an interface for accessing a MongoDB database
type Database interface {
	C(name string) Collection
}

// Bulk is an interface for running Bulk queries on a MongoDB collection
type Bulk interface {
	Upsert(pairs ...interface{})
	RemoveAll(selectors ...interface{})
	Run() (*mgo.BulkResult, error)
}

// Query is an interface for querying a MongoDB collection
type Query interface {
	All(result interface{}) error
	One(result interface{}) error
	Sort(fields ...string) Query
	Select(selector interface{}) Query
}

// Pipe is an interface for MongoDB aggregation
type Pipe interface {
	All(result interface{}) error
	One(result interface{}) error
}
