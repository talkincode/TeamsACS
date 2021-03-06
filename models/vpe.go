/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.0
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

// Vpe
// VPE is also a BRAS system
type Vpe  = DataObject

func (v DataObject) GetSecret() string {
	return v.GetStringValue("secret","")
}

func (v DataObject) GetVendorCode() string {
	return v.GetStringValue("vendor_code","")
}

// VpeManager
type VpeManager struct{ *ModelManager }

func (m *ModelManager) GetVpeManager() *VpeManager {
	store, _ := m.ManagerMap.Get("VpeManager")
	return store.(*VpeManager)
}


// GetVpeByIpaddr
func (m *VpeManager) GetVpeByIpaddr(ip string) (*Vpe, error) {
	coll := m.GetTeamsAcsCollection(TeamsacsVpe)
	doc := coll.FindOne(context.TODO(), bson.M{"ipaddr": ip})
	err := doc.Err()
	if err != nil {
		return nil, err
	}
	var result = new(Vpe)
	err = doc.Decode(result)
	return result, err
}

// GetVpeByIdentifier
func (m *VpeManager) GetVpeByIdentifier(identifier string) (*Vpe, error) {
	coll := m.GetTeamsAcsCollection(TeamsacsVpe)
	doc := coll.FindOne(context.TODO(), bson.M{"identifier": identifier})
	err := doc.Err()
	if err != nil {
		return nil, err
	}
	var result = new(Vpe)
	err = doc.Decode(result)
	return result, err
}
