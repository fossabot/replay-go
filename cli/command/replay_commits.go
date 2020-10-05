/*
 * Copyright (c) 2020 Alex <alex@webz.asia>
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 *   @author Alex <alex@webz.asia>
 *   @copyright Copyright (c) 2020 Alex <alex@webz.asia>
 *   @since 03.10.2020
 *
 */

package command

import (
	"time"
)

type ReplayCommitsArgs struct {
	InitStartDateString *string
	InitStartDateTime   *time.Time
	NewAuthor           *string
}

func (cmd *ReplayCommitsArgs) Parse() (err error) {
	*cmd.InitStartDateTime, err = time.Parse(time.RFC3339, *cmd.InitStartDateString)

	return err
}