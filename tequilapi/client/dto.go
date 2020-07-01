/*
 * Copyright (C) 2017 The "MysteriumNetwork/node" Authors.
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

package client

// Fees represents the transactor fee
type Fees struct {
	Registration uint64 `json:"registration"`
	Settlement   uint64 `json:"settlement"`
	Accountant   uint16 `json:"accountant"`
}

// RegistrationDataDTO holds input data required to register new myst identity on blockchain smart contract
type RegistrationDataDTO struct {
	Status     string `json:"status"`
	Registered bool   `json:"registered"`
}

// NATStatusDTO gives information about NAT traversal success or failure
type NATStatusDTO struct {
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

// SettleRequest represents the request to settle accountant promises
type SettleRequest struct {
	AccountantID string `json:"accountant_id"`
	ProviderID   string `json:"provider_id"`
}

// SettleWithBeneficiaryRequest represent the request to settle with new beneficiary address.
type SettleWithBeneficiaryRequest struct {
	SettleRequest
	Beneficiary string `json:"beneficiary"`
}
