/*
 * Copyright 2020 Google LLC.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package importer

import (
	"fmt"

	"github.com/GoogleCloudPlatform/healthcare-data-protection-suite/internal/terraform"
)

// ServiceAccount defines a struct with the necessary information for a google_service_account to be imported.
type ServiceAccount struct{}

// ImportID returns the ID of the resource to use in importing.
func (b *ServiceAccount) ImportID(rc terraform.ResourceChange, pcv ProviderConfigMap) (string, error) {
	project, err := fromConfigValues("project", rc.Change.After, pcv)
	if err != nil {
		return "", err
	}

	account_id, err := fromConfigValues("account_id", rc.Change.After, pcv)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("projects/%v/serviceAccounts/%v@%v.iam.gserviceaccount.com", project, account_id, project), nil
}