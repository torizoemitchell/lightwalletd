package main

import (
	"encoding/hex"
	"encoding/json"
	"github.com/zcash-hackworks/lightwalletd/common"
	"testing"
)

func Test(t *testing.T) {
	getblockchaininfoReply, err := hex.DecodeString(
		"7b22636861696e223a226d61696e222c22626c6f636b73223a3637373731" +
			"332c2268656164657273223a3637373731332c2262657374626c6f636b68" +
			"617368223a22303030303030303030303064653761373238333437316262" +
			"386362643636613935306632613332613736663232353064653033643132" +
			"37366437643036386538222c22646966666963756c7479223a3531343638" +
			"3233372e38303135313331342c22766572696669636174696f6e70726f67" +
			"72657373223a312c22636861696e776f726b223a22303030303030303030" +
			"303030303030303030303030303030303030303030303030303030303030" +
			"30303030303030303030323432326237363262663437363637222c227072" +
			"756e6564223a66616c73652c2273697a655f6f6e5f6469736b223a323339" +
			"38313131363337302c22636f6d6d69746d656e7473223a31343430313932" +
			"2c2276616c7565506f6f6c73223a5b7b226964223a227370726f7574222c" +
			"226d6f6e69746f726564223a747275652c22636861696e56616c7565223a" +
			"3134323433342e35393136333732342c22636861696e56616c75655a6174" +
			"223a31343234333435393136333732347d2c7b226964223a227361706c69" +
			"6e67222c226d6f6e69746f726564223a747275652c22636861696e56616c" +
			"7565223a3235383330362e32393830353736362c22636861696e56616c75" +
			"655a6174223a32353833303632393830353736367d5d2c22736f6674666f" +
			"726b73223a5b7b226964223a226269703334222c2276657273696f6e223a" +
			"322c22656e666f726365223a7b22737461747573223a747275652c22666f" +
			"756e64223a343030302c227265717569726564223a3735302c2277696e64" +
			"6f77223a343030307d2c2272656a656374223a7b22737461747573223a74" +
			"7275652c22666f756e64223a343030302c227265717569726564223a3935" +
			"302c2277696e646f77223a343030307d7d2c7b226964223a226269703636" +
			"222c2276657273696f6e223a332c22656e666f726365223a7b2273746174" +
			"7573223a747275652c22666f756e64223a343030302c2272657175697265" +
			"64223a3735302c2277696e646f77223a343030307d2c2272656a65637422" +
			"3a7b22737461747573223a747275652c22666f756e64223a343030302c22" +
			"7265717569726564223a3935302c2277696e646f77223a343030307d7d2c" +
			"7b226964223a226269703635222c2276657273696f6e223a342c22656e66" +
			"6f726365223a7b22737461747573223a747275652c22666f756e64223a34" +
			"3030302c227265717569726564223a3735302c2277696e646f77223a3430" +
			"30307d2c2272656a656374223a7b22737461747573223a747275652c2266" +
			"6f756e64223a343030302c227265717569726564223a3935302c2277696e" +
			"646f77223a343030307d7d5d2c227570677261646573223a7b2235626138" +
			"31623139223a7b226e616d65223a224f76657277696e746572222c226163" +
			"7469766174696f6e686569676874223a3334373530302c22737461747573" +
			"223a22616374697665222c22696e666f223a225365652068747470733a2f" +
			"2f7a2e636173682f757067726164652f6f76657277696e7465722e68746d" +
			"6c20666f722064657461696c732e227d2c223736623830396262223a7b22" +
			"6e616d65223a225361706c696e67222c2261637469766174696f6e686569" +
			"676874223a3431393230302c22737461747573223a22616374697665222c" +
			"22696e666f223a225365652068747470733a2f2f7a2e636173682f757067" +
			"726164652f7361706c696e672e68746d6c20666f722064657461696c732e" +
			"227d2c223262623430653630223a7b226e616d65223a22426c6f73736f6d" +
			"222c2261637469766174696f6e686569676874223a3635333630302c2273" +
			"7461747573223a22616374697665222c22696e666f223a22536565206874" +
			"7470733a2f2f7a2e636173682f757067726164652f626c6f73736f6d2e68" +
			"746d6c20666f722064657461696c732e227d7d2c22636f6e73656e737573" +
			"223a7b22636861696e746970223a223262623430653630222c226e657874" +
			"626c6f636b223a223262623430653630227d7d")
	if err != nil {
		t.Fatal("error decoding", err)
	}

	var f interface{}
	err = json.Unmarshal(getblockchaininfoReply, &f)
	if err != nil {
		t.Fatal("error parsing JSON getblockchaininfo response", err)
	}

	saplingHeight, blockHeight, chainName, branchID := common.GetInfoFromReply(f)
	if saplingHeight != 419200 {
		t.Error("unexpected saplingHeight", saplingHeight)
	}
	if blockHeight != 677713 {
		t.Error("unexpected blockHeight", blockHeight)
	}
	if chainName != "main" {
		t.Error("unexpected chainName", chainName)
	}
	if branchID != "2bb40e60" {
		t.Error("unexpected branchID", branchID)
	}
}
