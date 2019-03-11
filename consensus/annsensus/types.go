// Copyright © 2019 Annchain Authors <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package annsensus

import (
	"fmt"
	"github.com/annchain/OG/og"
	"github.com/annchain/OG/types"
	"time"
)

const (
	TimeoutPropose   = time.Duration(5) * time.Second
	TimeoutPreVote   = time.Duration(5) * time.Second
	TimeoutPreCommit = time.Duration(5) * time.Second
	TimeoutDelta     = time.Duration(1) * time.Second
)

type ValueIdMatchType int

const (
	MatchTypeAny ValueIdMatchType = iota
	MatchTypeByValue
	MatchTypeNil
)

type StepType int

const (
	StepTypePropose StepType = iota
	StepTypePreVote
	StepTypePreCommit
)

func (m StepType) String() string {
	switch m {
	case StepTypePropose:
		return "Proposal"
	case StepTypePreVote:
		return "PreVote"
	case StepTypePreCommit:
		return "PreCommit"
	default:
		return "Unknown"
	}
}

func (m StepType) IsAfter(o StepType) bool {
	return m > o
}

type Message struct {
	Type    og.MessageType
	Payload interface{}
}

func (m *Message) String() string {
	return fmt.Sprintf("%s %+v", m.Type.String(), m.Payload)
}

type ChangeStateEvent struct {
	NewStepType StepType
	HeightRound types.HeightRound
}

type TendermintContext struct {
	HeightRound types.HeightRound
	StepType    StepType
}

func (t *TendermintContext) Equal(w WaiterContext) bool {
	v, ok := w.(*TendermintContext)
	if !ok {
		return false
	}
	return t.HeightRound == v.HeightRound && t.StepType == v.StepType
}

func (t *TendermintContext) IsAfter(w WaiterContext) bool {
	v, ok := w.(*TendermintContext)
	if !ok {
		return false
	}
	return t.HeightRound.IsAfter(v.HeightRound) || (t.HeightRound == v.HeightRound && t.StepType.IsAfter(v.StepType))
}
