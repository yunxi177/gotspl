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

import (
	"bytes"
	"errors"
)

const AUTODETECT_NAME = "AUTODETECT"

type AutoDetectImpl struct {
	paperLength *float64
	gapLength   *float64
}

type AutoDetectBuilder interface {
	TSPLCommand
	PaperLength(paperLength float64) AutoDetectBuilder
	GapLength(gapLength float64) AutoDetectBuilder
}

func AutoDetectCmd() AutoDetectBuilder {
	return AutoDetectImpl{}
}

func (gi AutoDetectImpl) PaperLength(paperLength float64) AutoDetectBuilder {
	if gi.paperLength == nil {
		gi.paperLength = new(float64)
	}
	*gi.paperLength = paperLength
	return gi
}

func (gi AutoDetectImpl) GapLength(gapLength float64) AutoDetectBuilder {
	if gi.gapLength == nil {
		gi.gapLength = new(float64)
	}
	*gi.gapLength = gapLength
	return gi
}

func (gi AutoDetectImpl) GetMessage() ([]byte, error) {

	if !((gi.gapLength == gi.paperLength && gi.paperLength == nil) || (gi.gapLength != nil && gi.paperLength != nil)) {
		return nil, errors.New("ParseError AUTODETECT Command: gapLength, paperLength should be specified together")
	}

	buf := bytes.NewBufferString(AUTODETECT_NAME)

	if gi.gapLength != nil && gi.paperLength != nil {
		buf.WriteString(EMPTY_SPACE)
		buf.Write(formatFloatWithUnits(*gi.paperLength, false))
		buf.WriteString(VALUE_SEPARATOR)
		buf.Write(formatFloatWithUnits(*gi.gapLength, false))
	}

	buf.Write(LINE_ENDING_BYTES)
	return buf.Bytes(), nil
}
