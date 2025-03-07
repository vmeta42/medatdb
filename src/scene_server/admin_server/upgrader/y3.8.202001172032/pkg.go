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
package y3_8_202001172032

import (
	"context"
	"fmt"

	"configcenter/src/common/blog"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
)

//func init() {
//	upgrader.RegistUpgrader("y3.8.202001172032", upgrade)
//}

func upgrade(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	blog.Infof("start execute y3.8.202001172032")

	if err := createAuditLogTable(ctx, db, conf); err != nil {
		blog.Errorf("[upgrade y3.8.202001172032] createAuditLogTable failed, error %s", err.Error())
		return fmt.Errorf("createAuditLogTable failed, error %s", err.Error())
	}

	if err := addAuditLogTableIndex(ctx, db, conf); err != nil {
		blog.Errorf("[upgrade y3.8.202001172032] addAuditLogTableIndex failed, error %s", err.Error())
		return fmt.Errorf("addAuditLogTableIndex failed, error %s", err.Error())
	}

	return nil
}
