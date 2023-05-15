/*
 * Copyright 2020 Anton Globa
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
package gotspl

import "bytes"

const REPRINT_NAME = "REPRINT"

type ReprintImpl struct {
}

type ReprintBuilder interface {
	TSPLCommand
}

func ReprintCmd() ReprintBuilder {
	return ReprintImpl{}
}

func (h ReprintImpl) GetMessage() ([]byte, error) {
	buf := bytes.NewBufferString(REPRINT_NAME)
	buf.WriteString(EMPTY_SPACE)
	buf.WriteString("ON")
	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}
