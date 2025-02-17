package test

// Copyright (c) 2015-2019 Bitontop Technologies Inc.
// Distributed under the MIT software license, see the accompanying
// file COPYING or http://www.opensource.org/licenses/mit-license.php.

import (
	"log"
	"testing"

	"github.com/bitontop/gored/coin"
	"github.com/bitontop/gored/exchange"
	"github.com/bitontop/gored/pair"

	"github.com/bitontop/gored/exchange/bitfinex"
	"github.com/bitontop/gored/test/conf"
	// "../../exchange/bitfinex"
	// "../conf"
)

/********************Public API********************/
func Test_Bitfinex(t *testing.T) {
	e := InitBitfinex()

	pair := pair.GetPairByKey("BTC|EOS")

	// Test_Coins(e)
	// Test_Pairs(e)
	Test_Pair(e, pair)
	// Test_Orderbook(e, pair)
	Test_ConstraintFetch(e, pair)
	Test_Constraint(e, pair)

	Test_Balance(e, pair)
	// Test_Trading(e, pair, 0.0001000123, 100)
	// Test_Trading_Sell(e, pair, 0.04123456789, 0.02010123456789)
	// Test_Withdraw(e, pair.Base, 1, "ADDRESS")
	// log.Printf("Url: %v", e.GetTradingWebURL(pair))
	Test_DoWithdraw(e, pair.Target, "1", "0x37E0Fc27C6cDB5035B2a3d0682B4E7C05A4e6C46", "tag")
}

func InitBitfinex() exchange.Exchange {
	coin.Init()
	pair.Init()
	config := &exchange.Config{}
	config.Source = exchange.EXCHANGE_API
	// config.Source = exchange.JSON_FILE
	// config.SourceURI = "https://raw.githubusercontent.com/bitontop/gored/master/data"
	// utils.GetCommonDataFromJSON(config.SourceURI)
	conf.Exchange(exchange.BITFINEX, config)

	ex := bitfinex.CreateBitfinex(config)
	log.Printf("Initial [ %v ] ", ex.GetName())

	config = nil
	return ex
}
