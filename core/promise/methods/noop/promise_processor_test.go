/*
 * Copyright (C) 2018 The "MysteriumNetwork/node" Authors.
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package noop

import (
	"testing"
	"time"

	"github.com/mysteriumnetwork/node/core/promise"
	"github.com/mysteriumnetwork/node/money"
	"github.com/mysteriumnetwork/node/session"
	"github.com/stretchr/testify/assert"
)

var _ session.PromiseProcessor = &PromiseProcessor{}

func TestPromiseProcessor_Start_SendsBalanceMessages(t *testing.T) {
	dialog := &fakeDialog{}

	processor := &PromiseProcessor{
		dialog:          dialog,
		balanceInterval: time.Millisecond,
	}
	err := processor.Start(proposal)
	defer processor.Stop()

	assert.NoError(t, err)
	waitForBallanceState(t, processor, balanceNotifying)

	lastMessage, err := dialog.waitSendMessage()
	assert.NoError(t, err)
	assert.Exactly(
		t,
		promise.BalanceMessage{1, true, money.NewMoney(10, money.CURRENCY_MYST)},
		lastMessage,
	)
}

func TestPromiseProcessor_Stop_StopsBalanceMessages(t *testing.T) {
	dialog := &fakeDialog{}

	processor := &PromiseProcessor{
		dialog:          dialog,
		balanceInterval: time.Millisecond,
	}
	err := processor.Start(proposal)
	assert.NoError(t, err)
	waitForBallanceState(t, processor, balanceNotifying)

	err = processor.Stop()
	assert.NoError(t, err)
	waitForBallanceState(t, processor, balanceStopped)
}

func waitForBallanceState(t *testing.T, processor *PromiseProcessor, expectedState balanceState) {
	for i := 0; i < 10; i++ {
		if processor.getBalanceState() == expectedState {
			return
		}
		time.Sleep(time.Millisecond)
	}
	assert.Fail(t, "State expected to be ", string(expectedState))
}
