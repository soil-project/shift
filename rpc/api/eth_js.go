// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package api

// JS api provided by web3.js
// shf_sign not standard

const Eth_JS = `
web3._extend(
    // To be figured out // SHIFT
    {
	property: 'shf',
	methods:
	[
		new web3._extend.Method({
			name: 'sign',
			call: 'shf_sign',
			params: 2,
			inputFormatter: [web3._extend.utils.toAddress, null]
		}),
		new web3._extend.Method({
			name: 'resend',
			call: 'shf_resend',
			params: 3,
			inputFormatter: [web3._extend.formatters.inputTransactionFormatter, web3._extend.utils.fromDecimal, web3._extend.utils.fromDecimal]
		})
	],
	properties:
	[
		new web3._extend.Property({
			name: 'pendingTransactions',
			getter: 'shf_pendingTransactions'
		})
	]
});
`
