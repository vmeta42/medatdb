/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package metadata

import (
	"configcenter/src/common/mapstr"
)

// RspID response id
type RspID struct {
	ID int64 `json:"id"`
}

// CreateResult create result
type CreateResult struct {
	BaseResp `json:",inline"`
	Data     []Association `json:"data"`
}

// UpdateResult update result
type UpdateResult struct {
	BaseResp `json:",inline"`
}

// DeleteResult delete result
type DeleteResult struct {
	BaseResp `json:",inline"`
}

// QueryObjectResult query object result
type QueryObjectResult struct {
	BaseResp `json:",inline"`
	Data     []Object `json:"data"`
}

// CreateObjectResult create object result
type CreateObjectResult struct {
	BaseResp `json:",inline"`
	Data     RspID `json:"data"`
}

// CreateObjectAttributeResult create object attribute result
type CreateObjectAttributeResult struct {
	BaseResp `json:",inline"`
	Data     RspID `json:"data"`
}

// AttributeWrapper  wrapper, expansion field
type AttributeWrapper struct {
	Attribute         `json:",inline"`
	AssoType          int    `json:"bk_asst_type"`
	AsstForward       string `json:"bk_asst_forward"`
	AssociationID     string `json:"bk_asst_obj_id"`
	PropertyGroupName string `json:"bk_property_group_name"`
}

// UpdateGroupCondition update group condition struct
type UpdateGroupCondition struct {
	ModelBizID int64 `json:"bk_biz_id"`
	Condition  struct {
		ID int64 `field:"id" json:"id,omitempty"`
	} `json:"condition"`

	Data struct {
		IsCollapse *bool   `field:"is_collapse" json:"is_collapse,omitempty"`
		Name       *string `field:"bk_group_name" json:"bk_group_name,omitempty"`
		Index      *int64  `field:"bk_group_index" json:"bk_group_index,omitempty"`
	} `json:"data"`
}

// QueryObjectAttributeWrapperResult query object attribute with association info result
type QueryObjectAttributeWrapperResult struct {
	BaseResp `json:",inline"`
	Data     []AttributeWrapper `json:"data"`
}

// QueryObjectAttributeResult query object attribute result
type QueryObjectAttributeResult struct {
	BaseResp `json:",inline"`
	Data     []Attribute `json:"data"`
}

// CreateObjectGroupResult create the object group result
type CreateObjectGroupResult struct {
	BaseResp `json:",inline"`
	Data     RspID `json:"data"`
}

// QueryObjectGroupResult query the object group result
type QueryObjectGroupResult struct {
	BaseResp `json:",inline"`
	Data     []Group `json:"data"`
}

// CreateObjectClassificationResult create the object classification result
type CreateObjectClassificationResult struct {
	BaseResp `json:",inline"`
	Data     RspID `json:"data"`
}

// QueryObjectClassificationResult query the object classification result
type QueryObjectClassificationResult struct {
	BaseResp `json:",inline"`
	Data     []Classification `json:"data"`
}

// ClassificationWithObject classification with object
type ClassificationWithObject struct {
	Classification `json:",inline"`
	Objects        []Object `json:"bk_objects"`
}

// QueryObjectClassificationWithObjectsResult query the object classification with objects result
type QueryObjectClassificationWithObjectsResult struct {
	BaseResp `json:",inline"`
	Data     []ClassificationWithObject `json:"data"`
}

// QueryObjectAssociationResult query object association result
type QueryObjectAssociationResult struct {
	BaseResp `json:",inline"`
	Data     []Association `json:"data"`
}

// InstResult inst item result
type InstResult struct {
	Count int             `json:"count" example:"1"`
	Info  []mapstr.MapStr `json:"info"`
}

// QueryInstResult query inst result
type QueryInstResult struct {
	BaseResp `json:",inline"`
	Data     InstResult `json:"data"`
}

// UpdateInstResult cache inst result
type UpdateCacheInstResult struct {
	BaseResp `json:",inline"`
	Data     mapstr.MapStr `json:"data"`
}

// CreateInstResult create inst result
type CreateInstResult struct {
	BaseResp `json:",inline"`
	Data     mapstr.MapStr `json:"data"`
}

// ObjClassificationObject define the class object class
type ObjClassificationObject struct {
	Classification `bson:",inline"`
	Objects        []Object                 `json:"bk_objects"`
	AsstObjects    map[string][]interface{} `json:"bk_asst_objects"`
}
