// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package convert

import (
	"strings"

	atlas "github.com/mongodb/go-client-mongodb-atlas/mongodbatlas"
)

//Public constants
const (
	AdminDB = "admin"
)

//Private constants
const (
	roleSep = "@"
)

// BuildRoles converts the roles inside the array of string in an array of Atlas.Role Objects
// r contains roles in the format roleName@dbName
func BuildRoles(r []string) []atlas.Role {
	roles := make([]atlas.Role, len(r))
	for i, roleP := range r {
		role := strings.Split(roleP, roleSep)
		roleName := role[0]
		databaseName := AdminDB
		if len(role) > 1 {
			databaseName = role[1]
		}

		roles[i] = atlas.Role{
			RoleName:     roleName,
			DatabaseName: databaseName,
		}
	}
	return roles
}