// Copyright 2018 The Nakama Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package server

import (
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/heroiclabs/nakama/api"
	"golang.org/x/net/context"
)

func (s *ApiServer) ListStorageEntities(ctx context.Context, in *api.ListStorageEntitiesRequest) (*api.StorageEntityList, error) {

	return nil, nil
}

func (s *ApiServer) GetStorageEntities(ctx context.Context, in *api.GetStorageEntitiesRequest) (*api.StorageEntities, error) {

	return nil, nil
}

func (s *ApiServer) WriteStorageEntities(ctx context.Context, in *api.WriteStorageEntitiesRequest) (*api.StorageEntityAcks, error) {

	return nil, nil
}

func (s *ApiServer) DeleteStorageEntities(ctx context.Context, in *api.DeleteStorageEntitiesRequest) (*empty.Empty, error) {

	return &empty.Empty{}, nil
}