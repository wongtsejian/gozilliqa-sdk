/*
 * Copyright (C) 2021 Zilliqa
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
 * along with this program.  If not, see <https://www.gnu.org/licenses/>.
 */
package core

import (
	"encoding/json"
	"github.com/wongtsejian/gozilliqa-sdk/v4/multisig"
	"github.com/wongtsejian/gozilliqa-sdk/v4/util"
	"io/ioutil"
	"strings"
	"testing"
)

func Test_DeserializeFromJsonToTxBlockT(t *testing.T) {
	txJson, err := ioutil.ReadFile("txblock.json")
	if err != nil {
		t.Fatal(err.Error())
	}

	var txBlockT TxBlockT
	err2 := json.Unmarshal(txJson, &txBlockT)
	if err2 != nil {
		t.Fatal(err2.Error())
	}

	t.Log(txBlockT)

	txBlockHeader := NewTxBlockHeaderFromTxBlockT(&txBlockT)
	headerBytes := txBlockHeader.Serialize()
	headerBytesString := strings.ToUpper(util.EncodeHex(headerBytes))
	t.Log(headerBytesString)
	if headerBytesString != "0A46080112204A458FF78CBC60B5D9CA9503BFFCF88E8042F2D787AB05B0958F4A5C472D4E271A204AA4D99C4B9C7A18D6EC3AA658EC96EF2E8170A19AE25F0B364A4711EADD51DD1090BF05180022120A100000000000000000000000000000000030033A660A20D2D85A130929A9F82A40F82B025B933628E7F77A514FFE0B80A20CC05AC17F33122000000000000000000000000000000000000000000000000000000000000000001A206929CA0C7380085057B053CB469CEC4C454C2D19E4B8DF4C7958D98960FEEEF940004A230A210200B7F73A5D4FB220EE73C2CB2C478E7CF70F6F8E8EEA5367BDCFCE265F6C0D295001" {
		t.Fail()
	}
}

func Test_DeserializeFromJsonToTxBlockT2(t *testing.T) {
	txJson, err := ioutil.ReadFile("txblock2.json")
	if err != nil {
		t.Fatal(err.Error())
	}

	var txBlockT TxBlockT
	err2 := json.Unmarshal(txJson, &txBlockT)
	if err2 != nil {
		t.Fatal(err2.Error())
	}

	t.Log(txBlockT)

	txBlockHeader := NewTxBlockHeaderFromTxBlockT(&txBlockT)
	headerBytes := txBlockHeader.Serialize()
	headerBytesString := strings.ToUpper(util.EncodeHex(headerBytes))
	t.Log(headerBytesString)
	if headerBytesString != "0A4608011220AA5764CC1646085C6DD2042BB784ED6A4E154D22F0D1D1DACDD9E86662601E071A201947718B431D25DD65C226F79F3E0A9CC96A948899DAB3422993DEF1494A9C951090BF05180022120A100000000000000000000000000000000030013A660A204853067D757551C7F0119786CC830770A5F6C4DEB0D90A0B6DA49CAD49966613122000000000000000000000000000000000000000000000000000000000000000001A201B447A8BBE57F5BD8A2B4E0EDF0BECB29BF928D007E7BB1F170754835A7AC9BE40794A230A210213D5A7F74B28F3F588FF6520748DBB541986E98F75FA78D6334B2D0AAB4C1E575001" {
		t.Fail()
	}
}

func TestVerifyTxBlock(t *testing.T) {
	txJson, err := ioutil.ReadFile("txblock3.json")
	if err != nil {
		t.Fatal(err.Error())
	}

	var txBlockT TxBlockT
	err2 := json.Unmarshal(txJson, &txBlockT)
	if err2 != nil {
		t.Fatal(err2.Error())
	}

	txBlock := NewTxBlockFromTxBlockT(&txBlockT)
	headerBytes := txBlock.Serialize()

	if "0A4608011220AA5764CC1646085C6DD2042BB784ED6A4E154D22F0D1D1DACDD9E86662601E071A201947718B431D25DD65C226F79F3E0A9CC96A948899DAB3422993DEF1494A9C951090BF05180022120A100000000000000000000000000000000030013A660A204853067D757551C7F0119786CC830770A5F6C4DEB0D90A0B6DA49CAD49966613122000000000000000000000000000000000000000000000000000000000000000001A20C35A02EC1E5A981A41E13EBF439F020BD4D6534CB3A2FF1AA0930B66BF0290A840004A230A210213D5A7F74B28F3F588FF6520748DBB541986E98F75FA78D6334B2D0AAB4C1E575001A522E0591FE3BE75A99F83192E028DF995CEDBEF356343A2D8E054D300E03AADE6D7FACB679569C7EF0E81EF202D889ABB3E58544962CF16B1648737513D0F9C000AFE00" != strings.ToUpper(util.EncodeHex(headerBytes)) {
		t.Log(strings.ToUpper(util.EncodeHex(headerBytes)))
		t.Fatal("tx bytes error")
	}

	commKeys := []string{"0213D5A7F74B28F3F588FF6520748DBB541986E98F75FA78D6334B2D0AAB4C1E57",
		"0239D4CAE39A7AC2F285796BABF7D28DC8EB7767E78409C70926D0929EA2941E36",
		"02D2D695D4A352412E0D32A8BDF6EA3A606D35FE2C2F850C54D68727D065894986",
		"02E5E1BE6C924349F2C2B20CE05A2650B3E56C7722A2E5952EE27D12DEE7A4A6E6",
		"0300AB86B413FAA64A52FB61B5A28A6C361F87A5B0871C4F01C394D261415B0989",
		"03019AF5B10FFE09FB0EE02B59195EF5E6F5BE51D17EAF5604EA452078CD465C4B",
		"0323086D473DF937B6297FB755FA8E57C0FB2760512AED7757748B597C48F797A0",
		"032AEE20CFC59EAEB7838DAC2A9BAF96C8D69CF2C866FB4A3F1DFB02BCFCA356BB",
		"033207325A3CC671034FEBA86EC8D0AA412DF60C7E8292044D510DF582787DCC05",
		"0334AA0F7CA2EAA56B6B752533F9C60777E96C6D1ABE84B463F60ADD89843794AE",
	}

	var pubKeys [][]byte
	for index, key := range commKeys {
		if txBlock.Cosigs.B2[index] {
			pubKeys = append(pubKeys, util.DecodeHex(key))
		}
	}
	aggregatedPubKey, err3 := multisig.AggregatedPubKey(pubKeys)
	if err3 != nil {
		t.Fatal(err3.Error())
	}
	t.Log("aggregated public key = ", util.EncodeHex(aggregatedPubKey))

	r, s := txBlock.GetRandS()
	if !multisig.MultiVerify(aggregatedPubKey, headerBytes, r, s) {
		t.Fail()
	}

}
