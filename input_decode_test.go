package eth_decode_lib

import (
	"encoding/hex"
	"fmt"
	"testing"
)

// tx 0xf81bd8bf20567486a863cc0a8223d415836ee8f4f73fc50e33b80b0af56f8462
const txInput = "0x4e43e4950000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000012000000000000000000000000000000000000000000000000000000000000000c0000000000000000000000000f0245f6251bef9447a08766b9da2b07b28ad80b0000000000000000000000000000000000000000000000000000000000157efa3000000000000000000000000000000000000000000000000000000000157f3a2201dcae949d7a898a8565b9b872f15b41197e3ac680f4e49d9fa2bac58f59df1d44de474d0fcfc0a7323be0302e8df89d37aeb422c734033ab394bbc227687340000000000000000000000000000000000000000000000000000000000000089000000000000000000000000000000000000000000000000000000000000005f541fa67ddd0cc7a9fa10d2adbb6c490f9c112e73a1c3ecd785f81f8dac9758b85310d21909e4c57aadcb150dfe9b90f0beb3b5012d04186078843fc37656efac000000000000000000000000000000000000000000000000000000000000001c91efe3fdb23ca9977858c04417d3ef43016604ec48e271f1fe79197b5dce90992a4374a07805c87367a08b2aaa40fabdbbeeb91b012ad34fee24c45d10ba3af1000000000000000000000000000000000000000000000000000000000000001c0632aaff273c0b7ea7bd60e1e76207d1fcb13a791a1fb9a74594e55a9e6b186c605981575cf8aeaa12411d2c684a35e7e38ce434d93cd08123883316bfde1eaa000000000000000000000000000000000000000000000000000000000000001b028d0fe699b766ac0704597acb2966a021c41d70815271ffdaf05f9bf67b41e76e541c6588bcf7895429244cf889e57031efc75a74d606e9cc831c7c7eb5e676000000000000000000000000000000000000000000000000000000000000001c1b35f1bac9ba122bbaf73a62c455b7d53a6c492a424de100c1c08486382ce7b003e31424ad1138b95f0c9570167dd72e4d09b82b778be1cc2eb0599344225e42000000000000000000000000000000000000000000000000000000000000001b656c20fbdbd2d8dc98b803ba163f586f10f8a419df43c32e1a45c4d8daa57fc8372c27938f594ed166f8a2f0da2963183415e6709c9196af08c02723237c0fd6000000000000000000000000000000000000000000000000000000000000001bbc12f0809641c7f9769c126cc8d6ce86f70e08a589a8c3ebff66a79e19cde6304f206e9fb16918027c20acf667860adf6e7bfc06dea52e2b78ea156c4c6a7498000000000000000000000000000000000000000000000000000000000000001c1309504227ea84e3cce949df28d7e74c269db81b7d99e2f3948f7e53705cceb41c5a9de5c4b11c8146a9bb93dfd5105fe809a2d33daf0ef4ee0a8fe772341d09000000000000000000000000000000000000000000000000000000000000001cfff3ce7d2e8b57737c34a12d4f8d8199b045b28a78931a4fb5adcdd0d56207892494dc0caf959fef00f3c2d89fb3d1d1d095f358f1092cf2d32d4e841c6a12a4000000000000000000000000000000000000000000000000000000000000001b00d58ce317b5763b084a7c16ba4884d0973e09121d5dd8fe6b6b1d8ad73b55f822591cbbe3b02e6f694f3544eb91cbe20e783bb9edb031d4806e43567a4ccdf7000000000000000000000000000000000000000000000000000000000000001b920119fd6e2308de20d20ebcaada1cecdd9b6887e42f6670b8682c9ccf5310a6275c66b71520771daf1619dda3502a149ba56e9f381655a127cc5cacaab5d204000000000000000000000000000000000000000000000000000000000000001b7ec9bc8cba0c7662ff0dfd06e8dc0da995c7bbbbbdc59f083a768e71fefe127f6cf2f4a6dd7e0c35e6ef78897aaaa3b54c036969f2448ca63d91230ef0160803000000000000000000000000000000000000000000000000000000000000001c0be770d79b3fe0adc7c39c96e73602520fb4ca4b7eb4eb17a2dcbd454c831b976650620f6109fc6e0acf05dbc7e6c4bff4db91ed97c0e604aea718066a881ec0000000000000000000000000000000000000000000000000000000000000001ce9efa0a4b4780f0f5ba678fc40467d2e97060253c8459e6e339b156e387b92c051f1ba92859254df404ffdd644ead356c47c8cc5ea99fc7d3f7baac834b1ecc4000000000000000000000000000000000000000000000000000000000000001bba9828ccd96c44f3be39cca4099f778b8f017b930855a94334a259486fc07f1819ea73742ae109b033ba2510d7f51cbb15baf2165ef54531ecb71ca060e326cd000000000000000000000000000000000000000000000000000000000000001cc2dcec2f98303cd9e12527a8922809e5e8863330ff218c9808eacce4d337f1520f8a67ef826f926f299c105ede1bf8cefdd6903183516adc9206c84cd5b81fc3000000000000000000000000000000000000000000000000000000000000001b7dde5c90e92697df5ddd74c7bdd162e3da03b4dd3022f3d179e33baa30a3b32267ddadae0db81d946e2c0df658bcd95000a8e5ad9217aa887078801fe6c4e1a0000000000000000000000000000000000000000000000000000000000000001b1d16025e3dae55a30548a72c5f128c04d8573b9ecf89c3d8da36357fdb0d670307a37c3cf158f0def9626fc62df88718ffe039701bfa3d4916d50cb29678f3cd000000000000000000000000000000000000000000000000000000000000001c8631f0fd177ced8d460bb4ca82941fcc8526ea4e1b92875a07950b5a9b3a3ce078d2d269d78f0dd60984759f29e40d2c58d1b59c3fe4e400f5654ad8f5e9e961000000000000000000000000000000000000000000000000000000000000001b16fad4378310a264245eec2e3731fba7c825ffb771fdd42d834e8f0ab19590a54cdac99441956ee4b9fc88d1b065a069b2f4250dde34b497540a7e334ed3968b000000000000000000000000000000000000000000000000000000000000001b0f9a67c9d10d35436437a33b6dcf32568353f313abe4a6cd031d17af99748693467f2a298489b6d6cbd151e6a2c4c24797704d836af30a58c5b1a3736f0b57e5000000000000000000000000000000000000000000000000000000000000001cd2562b1c40cb875c73801ce92eec6593033d5422818645a4f8db9df3db8df3c477f946f08b056c55395b106c451bf38791c62ba76d2376d8a2702a03e3bfb12c000000000000000000000000000000000000000000000000000000000000001b4dd1b0b8fe8e197847d8844d746c0f820921f98ba88b6feb41129c53cf958f29112078c28c18b33ffc49e872c846447acd0d20dea88325eb03e3d9693f1e9335000000000000000000000000000000000000000000000000000000000000001cd1cb2d0cb68c92ad04046e685ea0c4d58b3a6b550590569d75af3aa2ce9dc7071edf05543e54566d8a0b6e30927b9d540a816d3a34b36694de6fb5a29f0e38b5000000000000000000000000000000000000000000000000000000000000001c7d1dac8ffd53adec1f80b837e6f2536424d4b13844ccfe5e03978a05e942d0450e422c906a2f325d170be948d58e6770a5bef7ce551316f2244e3b99d5e0f6cd000000000000000000000000000000000000000000000000000000000000001b1ffb95e5d59421dda717973958d71ec359b017b59bc1816aba8f6bc8ffef588e4ebe8f999b9a54de65237e44910a0ae9e83ddf1be9ec0613e5ab781fee4807fa000000000000000000000000000000000000000000000000000000000000001b9a0bcd93624dde2b896bf7f24235d392b564090966adb397423ec8b2f1916eac5f768a55780ce004eab10b2e3f7a666abdd17b9f2eb7bc8d484506ca23b48da5000000000000000000000000000000000000000000000000000000000000001b2fa8e9144b551297c4292e6d849399e48851922be2831d49bf754011692be9b261c720e034e07bfafe21d2d823f634454d0b66d440794844c073e3d02dd07792000000000000000000000000000000000000000000000000000000000000001ca6a443779050c67a3d26379fbe004110b98e2a94d8f8e740fe36c8d7446d55ab7586297a0067fa15dc2d6a4d473d28521136064fd3359418753c5d4712988785000000000000000000000000000000000000000000000000000000000000001b5d89d0562579ef279032e44fbfd63cd0122bfe7140361d8ffa167389ee05a69211d7759bc29a022488252058ef4701ad26752b97ac6a3b2e498e5235c12ab1e8000000000000000000000000000000000000000000000000000000000000001b484f3c156ad2ee20851f8b044bd7474225f3bdc67ecce429130cd00c2bba8bc07501438ef28047f6d4e2521ea9645d7a8209804d1630eb34123dd81490251878000000000000000000000000000000000000000000000000000000000000001b420317816896690e4da38731200989d71e6f3e839345b216f1e1441d6c4c53b274c032967653e2dbaba805e21cedb317be40f9ec6b00c8c9e1619e3316742739000000000000000000000000000000000000000000000000000000000000001b326e349e9654db229aefa8238a5cc2776e6980d2f0df081b8472b795ad68e6387240184f81b38e3feffb147bbe5f2ab5f6175ff02b61772aace919dcc1fe4b7b000000000000000000000000000000000000000000000000000000000000001bef3d87fff59962650075ad331be59da45d59279a61d120ee1429a79a2fe34e1c6ba311a5ce77a9be85941e644b2a5b5adb1ba382e691f941048df2ee7552498f000000000000000000000000000000000000000000000000000000000000001b91166e4d0d88376d4e87b0a66927779fac91b83143eee2d19e429880acadc07323c055717acf4841c6b12f958673bdad1dc4154c74ab880a36942f2700cc89d0000000000000000000000000000000000000000000000000000000000000001cd110fab0e1b9c70941150872b8b852dde6e57a223c0c3a398b63e98258981d6c643b911de60b9b0df76bccc80890592fde0000daa647d16a0a9737cdb1dc8479000000000000000000000000000000000000000000000000000000000000001ccf62a91187a4943b714323ed93870b1ceaf5110e7dadc3b64bde924169cd8e340ac26462a0c28306c55ccd31e4190e5922b6730a21d1b7bf66212d6cdf16bbd1000000000000000000000000000000000000000000000000000000000000001c3d22e4c92b4bff9bbc70dbbec94d1d0f5700eade0d11076ab4c218de6b8c6a261b00937e254dde8877baae5e2f16b152a91fd5c813c9fac7bc37ff2d958224fc000000000000000000000000000000000000000000000000000000000000001bd15588fa088ec41283f1e0cece6f4257cf73437d94547514e7d2e122cca7e33c529b92a959e96968f3f0b713116222760a3fd0b48b66ee068fac72d8a0fc70d6000000000000000000000000000000000000000000000000000000000000001b777f1869d796d4279275a354833c7c15ec9d6aa5ae8c81fe05ec3d30a3b0c9044b2b3c0a3521f20dfe4328e7ae96ae0928665b9f5b0ffe5b51847bc3443f4f95000000000000000000000000000000000000000000000000000000000000001b811cc1642af28c56353b92d9a7fa27fc7bfc0e8d740a9b4af0c1c38f7dd3c80a46489fc1d28903fa54b7974f294621d4632f5b7140f7b56cbb6fa2c144697f3d000000000000000000000000000000000000000000000000000000000000001ba73518b69cedaa706492f78ce23701dc6af1267df4c47249f2e6c2f92f44584d49c00125a057f4c85f08d173b2fb6d394ea43114914fc602180eb814f0cf3eaf000000000000000000000000000000000000000000000000000000000000001cb3b9c1a8d4076fed0e767d2336188c1939c22ccefe54f62d14bee0a2b5659d9e0a63d17b011d3b918348110dadbab276dfc0c2465ec3c53b770513c7eaa56596000000000000000000000000000000000000000000000000000000000000001b937e93fa8044ab32f68439e9177c89dbc8f533d10d018347fd3dd9affab500f72c0f64004667d17d19fd68716aa2d41133b2a3d44007d0100d12e46ef60e69cc000000000000000000000000000000000000000000000000000000000000001ceab4d55cbf02f94eb5e6da3cf9069cafe49b768583f176ecf3dd3ac8dc347464337e0ec4fc98a2b10a90fd959bb706bca4ffea5815248f0cb4a756b22ffee1a0000000000000000000000000000000000000000000000000000000000000001c866134f65f70de6be3f3379c23a081ab7112b5fee34de49bc883c2b03da37eb220fb68c19d5e7d0d83dac3086ede06ca531252a5a2c7fa3b089096850cf9cc23000000000000000000000000000000000000000000000000000000000000001c94c1c72b039e6ac7e6a45bc5165a255061a648f9626336d060c892b14e2e195e413706e6acb657a776c219b9885f07c9eb5aa5f56ad0ba3b14fb1e12a2caecfa000000000000000000000000000000000000000000000000000000000000001bacf1aa8387d942fc99913929939166939f0888a0ef2676c559471077c68ebcf97587bec3f9cd8e915a7857b7b4e8f96c2bc5c849e6318c270d8268c0c03c7a79000000000000000000000000000000000000000000000000000000000000001bada8c4e6e8e4bfa22d22a045aa2c7b64628628d20d2f81b1fafddcab0446147e686b113031ee93add543b4bee43f9541d01b77b43b13c247eb23cc4d06469d04000000000000000000000000000000000000000000000000000000000000001cd50bf7e27c1b98dfe910e98be29ad31f6306bf67abf1cd7c00fbf72ba1ff4b4d7e1bbce9f515657f3b9187dfa2dc07fc218be5a6815817f779aeddf9bb6fc3c1000000000000000000000000000000000000000000000000000000000000001b36fc1471b93dbcb9d6f53fc5de4ce6a611fed97504657dae1d2df309d4df70d216414da34b57cd89ead64c9c0820342ecb482a06934590c01640d12537896b61000000000000000000000000000000000000000000000000000000000000001bf3c9dcd9189f2da3edff9ad3beeddfa75308b351413cc38f055df66b7cbc0c2b565c1112a5356b3f618cf2e2392d6b0dd7738dff650fa2f836a11407ac2472d4000000000000000000000000000000000000000000000000000000000000001ce5f4c991c6b5f814805cc247200c6650cbfb42e1cebe9c9013682bebda9269f61c63ae05a5df4c4976ce598828a6785e0746bd3e438262847cd35441d37ed6b7000000000000000000000000000000000000000000000000000000000000001b84567d31d214dc447eefcee084084ac47d0dbc600c55acaad9924b61978e4e5d70d9a9e402f19b71c25e413989aeb5ead0ad5b312cb4a176a6030cf3ae44dde5000000000000000000000000000000000000000000000000000000000000001c55e73937c61ea5960cf35175f89125048c862299dc3c1763b0a6f781f596c2387a09a3b18457bf9a105e5aca93f7dade3cd926fffb0f579d2298d3c36952c1fc000000000000000000000000000000000000000000000000000000000000001b8511e45ac28946b13a6e12a66698fe8b3260c32189d79fd18226c95e2975b05e0c97d725c0b006ba24cadd1da98c8b49f11a1dd3db70ba801edf23a32d382334000000000000000000000000000000000000000000000000000000000000001cb3415faf2e62a55883adaa113d978d09ca69d9aace67cb003fbede2f9fbc7a70043e02be50b36eb3a1847ade4713d7d57d521f57cbd9f77eb713fbe3558eec1b000000000000000000000000000000000000000000000000000000000000001b21862c9fe1e78802fc06ed9019235435375f035d4d2fefd7def22fdc773e896666fa3ca740bfbb63cc1e32d17b5a1ebd8231de97051d04b04029c79959cd9992000000000000000000000000000000000000000000000000000000000000001b76f811af0bdf14f48f306fab91a123170af402e9b19745f06f4d6317b20a6aeb6a7f9fac0d89451a5304e758678051fcd228822f7ff7aa60d1e4f5a3962690e1000000000000000000000000000000000000000000000000000000000000001b92fc52272c87a5f95ffc84b3db05b35b948e366b9393387ce0e1effacb88ef0e519ee47b3692b11cb347e372b97eb1dc964440d9c06dce29e7b3bb4bd4d9f0ca000000000000000000000000000000000000000000000000000000000000001cb476dfe21f66a17ea4eac80de374045dd29dd8a1bf0f9f2258b8ca3abc4af81256de4755eba7f59d0420751aa1c7b09b41043099eb45e4793b42c998bcf963f6000000000000000000000000000000000000000000000000000000000000001b5c2d235abced9f4460952ad88e1e6affec52ea94338c8075f7271acd2bc5a3292c221ef040611082d192b0e860ed8403216a60ed7eac994862957293385336fb000000000000000000000000000000000000000000000000000000000000001c5bf81d753d079ca40aad5cd7a36fd4f59112d62cb9ad1fbf16dd3bdae881097d1b355e6bb11e611c56abf4451f4cbd631a470121554598dd6a7466a7235428a2000000000000000000000000000000000000000000000000000000000000001b857d0bd2d1baeb2bb06d77527e20e710af0ba7fa38edbc54e2076eff48a478a7551bd08b7155bcf5c47d8451cb964861fb7d3b197fa16cce32dc01fb1e7f293e000000000000000000000000000000000000000000000000000000000000001c2c51b99b0528933ff13675bb98958501eaec0000908337158583632dd189b2de4d0bdf8e73b3166ca98e5e633aba7a416f281d798df1110aae09a1a8fb2ff1cc000000000000000000000000000000000000000000000000000000000000001b35d3d5b2d9cace7ddfedc43bd1c22a139e5afc30a60bb0b6fac1e18239a0aaa458e254257ef96668dcf5a6d2bdfceb4b37d052f1c42318a7118ebc920e9b0051000000000000000000000000000000000000000000000000000000000000001c11910adba5781ebad89240ec01236b5f18064922a3d07c8eef6731d63c8bc0783292b1ed5098cd2a5497c9f6a7dd599d4628ec97fefa77b593d07d97ae491c7e000000000000000000000000000000000000000000000000000000000000001c33395f772708cf18979cd0ce726178fe11980aac3a67448a70a2d6c2f83697464e8d4ec13b287a3ac1da8b3053f3cb7e2df4a572656829f6ac74b082ea5ee0ea000000000000000000000000000000000000000000000000000000000000001b192628338806756f0bcd7e5cb2ede5707dd40985a217f50b3f1a7d3a9d4bb43851e9a6a1ed1341d39d5da4de87172e44bebfd77bb5c947dfbc0844acc823ffa1000000000000000000000000000000000000000000000000000000000000001ce5759fbb53d3e10eea9267d375a47b6068f507c7d1c015db01b585eea148adc86f503f4480ecdfb6dd25389add9dab996702a85bad2311ff4fc16dd7be6eba1e000000000000000000000000000000000000000000000000000000000000001b12a66b3523941781be1616bc7327a7e3c0818172bb508b7401c538bf662d61af0f7b86d8d241ad988263a9c76649a0d97cc6062b26de1a2ae1f495e2d47ecc82000000000000000000000000000000000000000000000000000000000000001c189f6f04af1dd959f0553fe328976c274bf631fdfffb564e130ab58c14c0237560f009ef89e6bd798957060b5fcd07f4099a47fb59a776e4a5292878f4124c64000000000000000000000000000000000000000000000000000000000000001bb3c4b0a6d9fc7af60eda27fb9a386e6a4cffde372062915de4a94df61373ab5a39b0be9ed28730bb650ea291e1fbca4b7705ed30aac379590b4a2727359fdbd6000000000000000000000000000000000000000000000000000000000000001bc18ef16a38994da1a949557b41e82bf427426815efc1fd56e49ba170dc4cce30070c11289ec0cf69e33295cd0eb1e6626858e5879c5efd4a57877ad26c199904000000000000000000000000000000000000000000000000000000000000001c558411517e350a2db5d3a5e2cedf28ef958182dca6e6c6fdc57278743f207a7f63c9e082899d3439e2e293d386def8bb62cdc7a83e384ca569850e50b100992d000000000000000000000000000000000000000000000000000000000000001b8fa6c498f24f862399747d530ddf5e8c3b641b9bb9a83b92c476661b8b32a2d02fd3b5959d6ece769c940f7b15f011ddfba5ead39eb914d40b33ef9be6fd192e000000000000000000000000000000000000000000000000000000000000001bf3ecaacbe4c5428a6e144bd964985e7a6db152db90b6f405f830db4ce90afa360eae542cb5f5fac14a42bd7b2addaa5c070ca0dbd216db78f464b056ded10766000000000000000000000000000000000000000000000000000000000000001b485f0f8eb399b8a4254ccee9239556a1c80f108b468711305fbb272d71536d4e5ab954aa1f13d6e4d6b634961c01f046c2d414521ab3d2a6e1941c5d0accfa92000000000000000000000000000000000000000000000000000000000000001cf41128c51823b64ac27e16ecba00294ec43d7e31eeb0aa8523502c638640051b37836af5d0c815aa1d5de61109df9f69ea107df0f4e5683fbdeb6edd9c719997000000000000000000000000000000000000000000000000000000000000001b3de92267c03ce63868d6ae4e5c1bcfa5a097cc6890691d429cbe68eef7f782ec765706a0ab3bb8f8ff2c3f04fd3d13cf57ec40dd978f72ba6b64710c63ab8c77000000000000000000000000000000000000000000000000000000000000001b67f89acf06f9339a566289d28648ad90b096f6255b0e608c6a177b35394f5e85299b35bb30c5fea5e801a986c1b6ba26625894234e47af1aa93f49e9d5052347000000000000000000000000000000000000000000000000000000000000001b37371cbdcb38da7cd80ec3cf449f2e73197b0424126c41bb9d9f049a76c72af073d583ab114e2c37b272a204a71ff2e46b4ef481988ca1d5385b2cf08ac5968d000000000000000000000000000000000000000000000000000000000000001b891e81f9b5d51f4b85168cc4603f68ea0adad3249af47d2951206318708f315b47f9a70a4f05ec36cd960d33895c66eed47d5b5edb148b71e107c6601288dba4000000000000000000000000000000000000000000000000000000000000001c2772eb862994d0f83f3a119151a23b0bfab563ade229a248eb4ddc05cc07f14c300f4fb15c26d3bc53a55e6986da8a39bbb0034bcf8dce60e501340ca840eed8000000000000000000000000000000000000000000000000000000000000001cf68e7e5f2a734bcaeb128021965fe69a6232fdc8bb87b818f469c2380916f6b9212ab548bd6b8a50a28b24c78432e34fdf9e16e3a1452d754d802baef26b2336000000000000000000000000000000000000000000000000000000000000001c19d52313c1c1c520105e3983a0a321fd436d61cc0be0832f0c42dc6464249cda779ea732bddd4b79e423dcc49089d29fc02aa17bacca84d7d0472743f3ae1369000000000000000000000000000000000000000000000000000000000000001b4fea24a7bb76c8475853139ed0c59d53d5f87708d15597eee90c8001e5dffdfe7d64cb19b42e5ae48457aca8dcc90ef1f6edd648f443bfad6e4e278b9e87b11a000000000000000000000000000000000000000000000000000000000000001b464b7dc654ba90f274ff4ce63720fef17e8061a971475ba3cc260f6c1db5af91589f78bc0dba0bd9e2fbae3f587effe013cd7423fefbd35768ccf7eb4a3715eb000000000000000000000000000000000000000000000000000000000000001c41c9d363ea650b7b109de89d99be81019e93eab4f210cae93641eebdf83adfcd22615bcf3b04ba4b42b9d63dad885199757848f8e909da8d25808a0edf538717000000000000000000000000000000000000000000000000000000000000001c9594d9267584e5a728c047a243fc4d9eec2c063e9003430606b94ff7931d0f45301904e3502af1366c57581d21200ebc1f2ed17842c035cbbff00ae9e0899533000000000000000000000000000000000000000000000000000000000000001bd6787b77ab3c3b343a025e72c2bd8d8d9863c11cff1e6e629d5fd7dfc8f4269742ec1deb79abc8f8821792c7448c16517a83f68dc1c982cc30d5f442e04e270b000000000000000000000000000000000000000000000000000000000000001bdc07940549238f448545184f66a4c9166bf2ff1a178a9c9ba79f9313b00a57f0606bfcb99c9059c19f8c89cb9ba7307b27cf8ab553519359ceef00cef8bb055d000000000000000000000000000000000000000000000000000000000000001bfe2df784b3fe54cf4d9739c156a3fa22d95bef2290f5e3e8bf0676427405663b7a8458ceb198cdbcbd3d78219323542be2e6b2b25cf4c66e45656c931954201b000000000000000000000000000000000000000000000000000000000000001cd606f705004c84de85a2f760b840c08abfb8d840076d4d51bd539c479ba5dc9d0b899f23951579f7b5b7bad4ad59c3bc34ced705bc82d4b12bb66d1f39e3b7e1000000000000000000000000000000000000000000000000000000000000001be28b590645ba454a73e312917573926293e0eafbcdfc66b1e618a9930f7bef1e79ceba9162138fce7e377505e96909af0ed0e4a2c382fb01637cd2fbd0cee9fb000000000000000000000000000000000000000000000000000000000000001c"
const abiJsonStr = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"headerBlockId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"}],\"name\":\"NewHeaderBlock\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"headerBlockId\",\"type\":\"uint256\"}],\"name\":\"ResetHeaderBlock\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"CHAINID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"VOTE_TYPE\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_nextHeaderBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentHeaderBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLastChildBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"headerBlocks\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"root\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"createdAt\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"heimdallId\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"networkId\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"_heimdallId\",\"type\":\"string\"}],\"name\":\"setHeimdallId\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"setNextHeaderBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"slash\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256[3][]\",\"name\":\"sigs\",\"type\":\"uint256[3][]\"}],\"name\":\"submitCheckpoint\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"sigs\",\"type\":\"bytes\"}],\"name\":\"submitHeaderBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"numDeposits\",\"type\":\"uint256\"}],\"name\":\"updateDepositId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

func TestDecodeTxInputsBytes(t *testing.T) {
	txInputBytes, err := hex.DecodeString(txInput[2:])
	if err != nil {
		fmt.Println(err)
		return
	}
	funcBs, data, err := DecodeTxInputsBytes(txInputBytes, abiJsonStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("input function sig:", hex.EncodeToString(funcBs))
	fmt.Println("input function data:", TxInputDataToJsonStr(data))
}

func TestDecodeTxInputsHex(t *testing.T) {
	funcBs, data, err := DecodeTxInputsHex(txInput, abiJsonStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("input function sig:", hex.EncodeToString(funcBs))
	fmt.Println("input function data:", TxInputDataToJsonStr(data))
}
