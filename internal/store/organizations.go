// Copyright (C) 2020 - present MongoDB, Inc.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the Server Side Public License, version 1,
// as published by MongoDB, Inc.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// Server Side Public License for more details.
//
// You should have received a copy of the Server Side Public License
// along with this program. If not, see
// http://www.mongodb.com/licensing/server-side-public-license
//
// As a special exception, the copyright holders give permission to link the
// code of portions of this program with the OpenSSL library under certain
// conditions as described in each individual source file and distribute
// linked combinations including the program with the OpenSSL library. You
// must comply with the Server Side Public License in all respects for
// all of the code used other than as permitted herein. If you modify file(s)
// with this exception, you may extend this exception to your version of the
// file(s), but you are not obligated to do so. If you do not wish to do so,
// delete this exception statement from your version. If you delete this
// exception statement from all source files in the program, then also delete
// it in the license file.

package store

import (
	"context"
	"fmt"

	"github.com/10gen/mcli/internal/config"
	"github.com/mongodb-labs/pcgc/cloudmanager"
)

type OrganizationLister interface {
	GetAllOrganizations() (interface{}, error)
}

type OrganizationCreator interface {
	CreateOrganization(string) (interface{}, error)
}

// GetAllProjects encapsulate the logic to manage different cloud providers
func (s *Store) GetAllOrganizations() (interface{}, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		result, _, err := s.client.(*cloudmanager.Client).Organizations.GetAllOrganizations(context.Background())
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}

// CreateOrganization encapsulate the logic to manage different cloud providers
func (s *Store) CreateOrganization(name string) (interface{}, error) {
	switch s.service {
	case config.CloudManagerService, config.OpsManagerService:
		org := &cloudmanager.Organization{Name: name}
		result, _, err := s.client.(*cloudmanager.Client).Organizations.Create(context.Background(), org)
		return result, err
	default:
		return nil, fmt.Errorf("unsupported service: %s", s.service)
	}
}