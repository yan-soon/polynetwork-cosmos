/*
 * Copyright (C) 2020 The poly network Authors
 * This file is part of The poly network library.
 *
 * The  poly network  is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The  poly network  is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 * You should have received a copy of the GNU Lesser General Public License
 * along with The poly network .  If not, see <http://www.gnu.org/licenses/>.
 */

package keeper_test

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	polycommon "github.com/polynetwork/poly/common"
	vconfig "github.com/polynetwork/poly/consensus/vbft/config"
	polytype "github.com/polynetwork/poly/core/types"
	"github.com/stretchr/testify/assert"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/Switcheo/polynetwork-cosmos/simapp"
	"github.com/Switcheo/polynetwork-cosmos/x/headersync/types"
)

var (
	header0     = "00000000ffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000c25191cf7592b15c04ca1bdcb07677a1bf3c995353ee4e68e35f798ee83fdcee00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000008e305f000000001dac2b7c00000000fd1a057b226c6561646572223a343239343936373239352c227672665f76616c7565223a22484a675171706769355248566745716354626e6443456c384d516837446172364e4e646f6f79553051666f67555634764d50675851524171384d6f38373853426a2b38577262676c2b36714d7258686b667a72375751343d222c227672665f70726f6f66223a22785864422b5451454c4c6a59734965305378596474572f442f39542f746e5854624e436667354e62364650596370382f55706a524c572f536a5558643552576b75646632646f4c5267727052474b76305566385a69413d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a343239343936373239352c226e65775f636861696e5f636f6e666967223a7b2276657273696f6e223a312c2276696577223a312c226e223a372c2263223a322c22626c6f636b5f6d73675f64656c6179223a31303030303030303030302c22686173685f6d73675f64656c6179223a31303030303030303030302c22706565725f68616e647368616b655f74696d656f7574223a31303030303030303030302c227065657273223a5b7b22696e646578223a312c226964223a2231323035303365663434626562613834343232626437366135393935333163396665353039363961393239613066656533356466363636393066333730636531396661386330227d2c7b22696e646578223a322c226964223a2231323035303338323437656663666561653066646637363036383564316163316330383362653366663565396134613534386263336132653938663034333466303932343833227d2c7b22696e646578223a332c226964223a2231323035303232303932653334653031373664636366386162623439366238333364353931643235353333343639623363616630653237396239373432393535646438666333227d2c7b22696e646578223a342c226964223a2231323035303237626437373165363861646238383339383238326532316138623033633132663634633233353165613439613262613036613033323763383362323339636139227d2c7b22696e646578223a352c226964223a2231323035303264306430653838336337336438323536636634333134383232646464393733633031373962373364386564336466383561616433386433366138623262306337227d2c7b22696e646578223a362c226964223a2231323035303361346634346464363563626363353262316431616335313734373337386137663834373533623566376266323736306361323133393063656436623137326262227d2c7b22696e646578223a372c226964223a2231323035303236393663306362653734663031656538356533633065626534656264633562656134303466313939643032363266313934316664333966663064313030323537227d5d2c22706f735f7461626c65223a5b362c362c312c352c322c322c372c332c362c352c362c322c372c322c372c322c362c362c322c312c352c312c312c352c332c372c372c332c332c332c342c352c332c332c342c322c352c342c352c322c312c372c372c372c312c332c342c362c372c342c372c362c352c332c352c372c312c322c332c352c372c342c342c362c332c332c342c322c352c352c322c322c342c342c332c312c352c362c342c322c372c322c312c362c312c342c362c312c332c312c372c332c312c312c342c362c352c342c362c362c352c372c312c342c325d2c226d61785f626c6f636b5f6368616e67655f76696577223a36303030307d7d3cc22b9403d96ee5c9422ca9d502e0907617ccb20000"
	header1     = "00000000ffffffffffffff7ffa946154bfc4e885a9cdaa98bbde21b0b909c584c9fc1189e4547ac1ce91845b0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000078fc2e71206220ec1a222b1d0ea9f4c95e843c400f39ca8cffe983e4a7ca2ec47afa345f0100000096a8b699d5806503fd0c017b226c6561646572223a332c227672665f76616c7565223a2242473563534435627a63353447312b6977756c4e72576c4c4b3079646d59304452356c4937714f41325a51672f65515a76346d4c5777386f2f67304c677a735a586c6e4f697a6e35424b6e36474c447a4a762f666750773d222c227672665f70726f6f66223a225a674242615578396f6d6e694b6362565a6d6433386174704c6f7a4e684a4d2b44425a41786755674f352b33504a4f4c7151525355724d6551444e70452f3378774c4f346c62326165515448736b764a4a7644486a413d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a302c226e65775f636861696e5f636f6e666967223a6e756c6c7d000000000000000000000000000000000000000005231205022092e34e0176dccf8abb496b833d591d25533469b3caf0e279b9742955dd8fc323120503ef44beba84422bd76a599531c9fe50969a929a0fee35df66690f370ce19fa8c0231205027bd771e68adb88398282e21a8b03c12f64c2351ea49a2ba06a0327c83b239ca923120503a4f44dd65cbcc52b1d1ac51747378a7f84753b5f7bf2760ca21390ced6b172bb23120502d0d0e883c73d8256cf4314822ddd973c0179b73d8ed3df85aad38d36a8b2b0c70542011cd4e5cd55eff0dc21b9d915ce7c7132dd79dff4af70f5ddf732b667b57c163a9b72e0c04a31cd7bd825ce46c0895b2ca2866222bb8c36c39029279fa7b0ced3e542011b9a1a8b6795ec5e1bdb8366d718d76b6021a76faf018fa18ac19409838e98b1332a00e26d97c502917fa855cf351c6e0d17205cec47204a94d6627d00dd474cad42011bad1b608402de149b2efd5f4bd00bf65ef58c268601601a47b6275149f3271bc853236b7b41b7e7996dd9d8704fbad8f11fa19e7b7ed902ac8739ad45e979484d42011b73e860ec2ac329554b800ec12d4e362b515695a5e445bafc416b043efb08241627db846bc8cbb205871d0e28de46cde714c728c1b3a2332295cb7099345060f142011b550f1ead1491093ce2a44c49fb0a76ab88ff398e3238722f80bf0c741eda97687648db9a704c59459a7121f57034fe7237d007c3b24b41651e6e319e9908e01e"
	header100   = "00000000ffffffffffffff7fce48d13809875fac375b133a7c77105bd83e11e51c45590f965d434190a800f9000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004104292335b2884b77995402cef2b463c45d40df809dc49c99e47c288703c7d51806355f6400000015ffe8fa04990b9ffd0c017b226c6561646572223a352c227672665f76616c7565223a22424a6646636f2f74704641524c6166617a73705a7634735a596c4745455570664b36567a4a31722f463749485835594c5757416d574f586d3374736c6f7268356269585659336c31696f37634a3149776b3169553438733d222c227672665f70726f6f66223a22687147726b68664c716a7a756b6244625644396d656968712b346c69654931752f65306b4433755957785947384f7747653856464767695731386c436b474b66354f30455938755a31634932756575736f48577341773d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a302c226e65775f636861696e5f636f6e666967223a6e756c6c7d00000000000000000000000000000000000000000623120502d0d0e883c73d8256cf4314822ddd973c0179b73d8ed3df85aad38d36a8b2b0c7231205022092e34e0176dccf8abb496b833d591d25533469b3caf0e279b9742955dd8fc3231205027bd771e68adb88398282e21a8b03c12f64c2351ea49a2ba06a0327c83b239ca923120502696c0cbe74f01ee85e3c0ebe4ebdc5bea404f199d0262f1941fd39ff0d100257231205038247efcfeae0fdf760685d1ac1c083be3ff5e9a4a548bc3a2e98f0434f09248323120503a4f44dd65cbcc52b1d1ac51747378a7f84753b5f7bf2760ca21390ced6b172bb0642011c3fe1389bfca65bcaae9d333c383f6fc73c1f423db47a22902b50f134a6f50c3e5f33743a6b208ddccaa6e0155e84e7c1241f21cab2a903e021b9ad0d98ea783b42011b177120ab2ff8aedd234466c5c1285dc9b851f1682fc3738383c1299dbc8db0ea4e3d840b3ceca297629710b2ea8ba30243e9de0e57a56f0b2547a444b5aeca5142011c854876a8e31070568c63658e2d6246b37433a8036f086aa9617d1a598bf16e683369ae2fc5a228d97ea71ae1c8b288b7457eb05935462b1edc073f2c50f587e042011b06fd7bb87cadbc671f3d636a19cb979d1bf856b6d855b3c31793dd6edb8edab01d5db710a1b633b0ef6993f3c7c025322453c2f5d9f4d9f97e8ba2f88958b9f942011bd94bcc40024bff94ca962f1a98dd16b3f5cda3c3354e560954927d86b583f86034159f90c01c594bfdb4b66e5bb3ee08a0adfdd5109df05594ca1bf0b439411742011c1cbf2f630af7ee565ea76e22d63b811b39a7d836d57201855ca8324d0708ac9d2eb938f8c6af5031c042c4dd3506f6ac94fb17f51055bcf06b998adf9c6a0990"
	header789   = "00000000ffffffffffffff7f24b2a5614cfb124e1418a8d8ef12acc56dcf9c14da85b3fdc974c7a3ecacbf3800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000128f71ea0c41792493b790576a617cf535f7f2628afe9cbcf8a8b5d9fd1eeecb7e22355f150300004a9c2857be05fd76fd0c017b226c6561646572223a332c227672665f76616c7565223a224247724137745176555447676843376339634e6e446c397877386e315172564d43486a362f365064506946792f694470544747664b7464724b4b657332557751414a4655634138497751433067684f657147526f5552413d222c227672665f70726f6f66223a227a6d4830446a646d5a536a6f7079336e58795876736c435641616734636b49773935744b3658504d387748765556506d57386d4f727838704363345534756668573647586e54574447724f6e334f6a356e76304244513d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a302c226e65775f636861696e5f636f6e666967223a6e756c6c7d000000000000000000000000000000000000000005231205022092e34e0176dccf8abb496b833d591d25533469b3caf0e279b9742955dd8fc323120503a4f44dd65cbcc52b1d1ac51747378a7f84753b5f7bf2760ca21390ced6b172bb231205038247efcfeae0fdf760685d1ac1c083be3ff5e9a4a548bc3a2e98f0434f092483231205027bd771e68adb88398282e21a8b03c12f64c2351ea49a2ba06a0327c83b239ca923120502d0d0e883c73d8256cf4314822ddd973c0179b73d8ed3df85aad38d36a8b2b0c70542011beaa584d6fe77426ab5cb7601b2254d646bb22da8f871351ed603fcf99f7c8e1f7d680c595990df183ec56172e9c03fec8f68b8a824f3268fbaaeca8428a3237242011bbabd655f0a871327a0518860cd4bcc809a1a457b4b4027f6002d39ad4a98b29a5e49f91d509b6665b977ad553c0cb1e338b9f39a6991d300dcde3875e6b2459542011b2d95435c800c269747f08b17afc3953451befca1b1da3660c40f662efdda8a156a29c5a03e66edd582adae641ea2e5f235893d1a8066bd4686295a7ef9ab6f3e42011bcde39bd54d6cc53af6870c50ac6e6560081a2cce0fb61538d17bb5e748649cea72069be01dac2704c0c9234b14663ac4696bbc48a30b02db4195c0615cadb03842011c2d927338308e11ea6ffc2c0c901227cdb760c9022b80665cf2ff4e57f110c1334f912a5943d04228c354e6c4d180e713557f34f2a38b43e9637bdf8195e94ca3"
	header60000 = "00000000ffffffffffffff7f721adaf73aaedee422b623581572078e895f7a47b63999b301d8a78c67ebf1fec31614565201089bf5c639d923ebfc63541c91522b8c29b0ff605bebac309b0e0000000000000000000000000000000000000000000000000000000000000000e7f41653fdaf63c4366db84674af2b614fcc40d5f37af5af4ccfad543fb318bc94bd3a5f60ea0000bbcb0265fa485dc6fd0c057b226c6561646572223a312c227672665f76616c7565223a22424d4644766a765a495737577578626148522b6568736b7a316d426d385a577065483758687a536632374b4b6a752f676b446b564f6578524b3142306d6d31322b68364175304264393963384f6335474b3843785858733d222c227672665f70726f6f66223a2254412b66757778345a4b796b4a724974633165685944587372424a35682f4841494744484261786f5569644e765634654d6437684345325872342b56456d414539573249326834486e35794c6a4a4f674858376167513d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a36303030302c226e65775f636861696e5f636f6e666967223a7b2276657273696f6e223a312c2276696577223a322c226e223a372c2263223a322c22626c6f636b5f6d73675f64656c6179223a31303030303030303030302c22686173685f6d73675f64656c6179223a31303030303030303030302c22706565725f68616e647368616b655f74696d656f7574223a31303030303030303030302c227065657273223a5b7b22696e646578223a312c226964223a2231323035303365663434626562613834343232626437366135393935333163396665353039363961393239613066656533356466363636393066333730636531396661386330227d2c7b22696e646578223a362c226964223a2231323035303361346634346464363563626363353262316431616335313734373337386137663834373533623566376266323736306361323133393063656436623137326262227d2c7b22696e646578223a322c226964223a2231323035303338323437656663666561653066646637363036383564316163316330383362653366663565396134613534386263336132653938663034333466303932343833227d2c7b22696e646578223a352c226964223a2231323035303264306430653838336337336438323536636634333134383232646464393733633031373962373364386564336466383561616433386433366138623262306337227d2c7b22696e646578223a342c226964223a2231323035303237626437373165363861646238383339383238326532316138623033633132663634633233353165613439613262613036613033323763383362323339636139227d2c7b22696e646578223a372c226964223a2231323035303236393663306362653734663031656538356533633065626534656264633562656134303466313939643032363266313934316664333966663064313030323537227d2c7b22696e646578223a332c226964223a2231323035303232303932653334653031373664636366386162623439366238333364353931643235353333343639623363616630653237396239373432393535646438666333227d5d2c22706f735f7461626c65223a5b322c362c322c322c312c352c352c372c372c332c362c332c352c342c332c352c332c362c352c362c362c312c342c322c362c352c332c372c332c372c332c362c312c342c312c362c362c322c342c352c312c312c312c372c352c342c332c332c332c352c342c312c342c342c322c342c312c322c322c312c342c362c352c342c342c362c352c372c372c362c332c322c322c352c352c332c372c372c372c322c332c332c372c372c372c312c312c332c312c372c352c322c312c342c362c322c342c312c322c362c322c352c342c362c375d2c226d61785f626c6f636b5f6368616e67655f76696577223a36303030307d7d3cc22b9403d96ee5c9422ca9d502e0907617ccb20623120503ef44beba84422bd76a599531c9fe50969a929a0fee35df66690f370ce19fa8c023120503a4f44dd65cbcc52b1d1ac51747378a7f84753b5f7bf2760ca21390ced6b172bb231205027bd771e68adb88398282e21a8b03c12f64c2351ea49a2ba06a0327c83b239ca9231205022092e34e0176dccf8abb496b833d591d25533469b3caf0e279b9742955dd8fc3231205038247efcfeae0fdf760685d1ac1c083be3ff5e9a4a548bc3a2e98f0434f09248323120502696c0cbe74f01ee85e3c0ebe4ebdc5bea404f199d0262f1941fd39ff0d1002570642011ce631bea110252971770367cf76e7b8255ca0bfcaa5bc35468d31c3b72eac364d76bd89b73879f30c7bd08326558d072e19e6f96cbb808dcbd40e4a209fe7f15742011bf1376babf31495fbe2433887cdeee92eefd3eb1d31360370ab9d2727161d6bb27594ffd3568452e0e514d929b6d0f7fedc7e776b6f7cb034e462441a855a500842011cac274dc007ba01c568d3c2c928f4c7878f5f473dee7c1440472f6fb18aba9d673e41e33f2aed0d3b7dbb12ff11f7b0ce59b3233a699618a74d64eee9475bf6f842011ccffc25fd448eaf274372d0b54b51ac95e5a62559fcc41dfd30a3ac9e73ca930a6083a9c59bf10baca68be238aa9f9bb0d08238674bbbe9992d919e8e3b78740942011c393897c4fb77264caddc74f4361a482c7841d5316599e2a492b889c2db3d2611419947a1da60e81934307c73b59fbf58aa89cdf0048f97e6550a08772f073ecb42011c0194499ac67d079ed6363c118f6a86cb90f70415183a192b63ca6a8c499516c44eb486dddac7b8110e555dae9010e2b9409f07a47f5fc782abe80623cfd5b511"
	header60005 = "00000000ffffffffffffff7f9c91404cf8eb56143b5a93db1501f91639d09582266965455eabc2955d88fd7bd8654a2ed24967403d20bc736349b3c2bd4c9ed22c96f63247ad0a90fa9ec4b000000000000000000000000000000000000000000000000000000000000000002c455f61e6750403d09c3cf8886d53e15e65bd09fca59757ae6ef9059350a990a2bd3a5f65ea0000d10ef17ddf2e9984fd10017b226c6561646572223a372c227672665f76616c7565223a22424b366b545578793442493144482b304d79326b663244585a5a4337424b675572766b63535a4d68474a55796a342b4f2f645357386b6430624e75504d705034516741626f6564336b367343394a636a2f6c4e503677553d222c227672665f70726f6f66223a2275644a36374c455746464c456b6535785657496e44542f4f6e396654544679327632544b367664704766496271517a775265666b504c45306f3577396f462f4777754e2b444f6e7838347968516147695850494a4a513d3d222c226c6173745f636f6e6669675f626c6f636b5f6e756d223a36303030302c226e65775f636861696e5f636f6e666967223a6e756c6c7d00000000000000000000000000000000000000000623120502696c0cbe74f01ee85e3c0ebe4ebdc5bea404f199d0262f1941fd39ff0d100257231205038247efcfeae0fdf760685d1ac1c083be3ff5e9a4a548bc3a2e98f0434f092483231205022092e34e0176dccf8abb496b833d591d25533469b3caf0e279b9742955dd8fc323120503ef44beba84422bd76a599531c9fe50969a929a0fee35df66690f370ce19fa8c023120503a4f44dd65cbcc52b1d1ac51747378a7f84753b5f7bf2760ca21390ced6b172bb23120502d0d0e883c73d8256cf4314822ddd973c0179b73d8ed3df85aad38d36a8b2b0c70642011c64f6913231218f17f16fecf461cffcce3355d32b57104233c0798cb2d05fbcf435142c251cb5427694e44894dcd2596ee98fe0b3d763c47a51e171f241c7b6cf42011ce6978c74b370f162bed5c21a81299f0b9943a273ef7e6926df1b459cd3822b804a39e2e90467516138c148197358f3d052178c9c0e682670f9c69f245393ffb442011b141b295c0e0b4566571ddeb25b20e7eaded8a447f0188a8902225dd3908b5b3b6a964af58f212385f3e8298a53324fa0ac5eb760add82cb5e4a36df3445eb4c542011bd996c10033dd73323d4fd1626a0a15d66ab3b74264578636dc470988cbd78f862f72a805011090de8ae86034cada4b443653a5dcee7ddd409bedcbe9531e7cd542011b514ff160eb1fe41720c4be232de64a3e777685dd42b546ded1697f0cdc6c5aa520dd6de4e0a6a68dba00ec84c654c628d985de18b7647f29c7b968e0699d978c42011cd9ff4f3283de762affbdf98923456db0a563dec5aac37763a7302b771eddcbc810851a3894d57546ae5ed97be373d3d29881fe755118141fafb408066597cb7b"
)

// returns context and an app with updated mint keeper
func createTestApp(isCheckTx bool) (*simapp.SimApp, sdk.Context, types.QueryClient) {
	app := simapp.Setup(isCheckTx)

	ctx := app.BaseApp.NewContext(isCheckTx, tmproto.Header{})
	queryHelper := baseapp.NewQueryServerTestHelper(ctx, app.InterfaceRegistry())
	types.RegisterQueryServer(queryHelper, app.HeadersyncKeeper)
	queryClient := types.NewQueryClient(queryHelper)

	return app, ctx, queryClient
}

func Test_headersync_Serialize_PolyHeader(t *testing.T) {
	var header polytype.Header
	h0s, _ := hex.DecodeString(header0)
	source := polycommon.NewZeroCopySource(h0s)
	err := header.Deserialization(source)
	assert.Nil(t, err)
	sink := polycommon.NewZeroCopySink(nil)
	err = header.Serialization(sink)
	assert.Nil(t, err)
	assert.Equal(t, h0s, sink.Bytes())
}

func Test_headersync_SyncHeaders(t *testing.T) {
	app, ctx, _ := createTestApp(true)

	h0s, _ := hex.DecodeString(header0)
	header := new(polytype.Header)
	err := header.Deserialization(polycommon.NewZeroCopySource(h0s))
	assert.Nil(t, err)

	err = app.HeadersyncKeeper.SyncGenesisHeader(ctx, header0)
	assert.Nil(t, err, "Sync genesis header fail")

	var chainID uint64 = header.ChainID
	consensusPeers, err := app.HeadersyncKeeper.GetConsensusPeers(ctx, chainID)
	assert.Nil(t, err)
	fmt.Printf("consensusPeers are %s\n", consensusPeers.String())
	// the genesis header should contain the consensus peers info, the height may be not ZERO but should be the header contains consensus address of next cycle
	assert.Equal(t, header.ChainID, consensusPeers.ChainId)
	assert.Equal(t, header.Height, consensusPeers.Height)
	cpBs, err := ExtractChainConfig(header)
	assert.Nil(t, err)
	resSink := polycommon.NewZeroCopySink(nil)
	consensusPeers.Serialize(resSink)
	assert.Equal(t, cpBs, resSink.Bytes())

	err = app.HeadersyncKeeper.SyncBlockHeaders(ctx, []string{header1, header789, header100, header60000, header60005})
	assert.Nil(t, err, "Sync Poly Chain block headers fail")

	consensusPeers, err = app.HeadersyncKeeper.GetConsensusPeers(ctx, chainID)
	assert.Nil(t, err)
	fmt.Printf("consensusPeers are %s\n", consensusPeers.String())
	assert.Equal(t, chainID, consensusPeers.ChainId)
	assert.Equal(t, uint32(60000), consensusPeers.Height)
	resSink = polycommon.NewZeroCopySink(nil)
	consensusPeers.Serialize(resSink)

	h180000s, err := hex.DecodeString(header60000)
	assert.Nil(t, err, "header180000 hex to header180000 bytes error")
	header180000 := new(polytype.Header)
	err = header180000.Deserialization(polycommon.NewZeroCopySource(h180000s))
	assert.Nil(t, err)
	cpBs, err = ExtractChainConfig(header180000)
	assert.Nil(t, err)
	assert.Equal(t, cpBs, resSink.Bytes())
}

func ExtractChainConfig(header *polytype.Header) ([]byte, error) {
	blkInfo := &vconfig.VbftBlockInfo{}
	if err := json.Unmarshal(header.ConsensusPayload, blkInfo); err != nil {
		return nil, fmt.Errorf("Unmarshal error")
	}
	if blkInfo.NewChainConfig != nil {
		consensusPeers := &types.ConsensusPeers{
			ChainId: header.ChainID,
			Height:  header.Height,
			Peers:   make(map[string]*types.Peer),
		}
		for _, p := range blkInfo.NewChainConfig.Peers {
			consensusPeers.Peers[p.ID] = &types.Peer{Index: p.Index, Pubkey: p.ID}
		}
		sink := polycommon.NewZeroCopySink(nil)
		consensusPeers.Serialize(sink)
		return sink.Bytes(), nil
	}

	return nil, fmt.Errorf("No new chain config")
}
