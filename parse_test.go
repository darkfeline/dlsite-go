// Copyright (C) 2018 Allen Li
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

package dlsite

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func testdataPath(p string, c RJCode) string {
	return filepath.Join("testdata", p, fmt.Sprintf("%s.html", c))
}

func TestParseWorkWithSeries(t *testing.T) {
	t.Parallel()
	c := "RJ189758"
	exp := &Work{
		RJCode:      RJCode(c),
		Name:        "意地悪な機械人形に完全支配される音声 地獄級射精禁止オナニーサポート4 ヘルエグゼキューション",
		Maker:       "B-bishop",
		Series:      "地獄級オナニーサポート",
		Description: "皆様、こんにちは。サークルB-bishopのpawnlank7と申します。\n今作は、地獄級に激しいオナニーサポート作品第4弾です。\n\n機械人形の『シコシコ』というボイスにあわせてオナニーをすることになり、全てを機械人形に支配されます。\nシコシコの声はセリフとは別で流れますので、言葉責めも絶え間なくお楽しみいただけます。\n\n今作では機械人形に実験と称した地獄を味わわされます。\n『ヘルエグゼキューション』は地獄の執行を意味します。\n\nなお、今作では終盤に複数の女性が登場し、オナニーを見学されるシーンがございます。\nモードによっての難易度に大きな差を設けてありますので、お好みでご試聴くださいませ。\n特にヘルモードでは極めて厳しい内容になっております。\n\n☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆3種類の難易度設定☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆\n\n本作品では、ハードモード、ベリーハードモード、ヘルモードのレベルをご用意しております。\nこれはオナニーパートのトラック数の違いです。\nトラックが進むにつれて段階的に厳しくなりますので、軽い難易度のものほど後半のトラックは削除されております。\n\n●ハードモード\n基本的なプレイが中心のベーシックな内容ですが、射精を強く禁止され、言葉責めも多量となっております。\nシーン08,09を削除\n\n●ベリーハードモード\nさらに視聴者を追い込むトラックが追加されており、内容はとても過酷です。\nシーン10を削除\n\n●ヘルモード\nすべてのシナリオをご堪能いただけますが、大変厳しい内容で、極めて残酷な内容となっております。\nシーン11は特別仕様\n\n重複するシナリオはありますが、ハードモードから徐々に慣らしていくことをオススメいたします。\n\n特にヘルモードは、機械人形は手加減なしの実験を執行するので、容赦は一切ありません。\n\n詳しくは00注意事項をご視聴下さい。\n\n\n☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆収録シーン&難易度☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆\n\n■00\u3000注意事項\n免責事項や説明。\n\n■01\u3000プロローグ\nオナニーサポートについての説明。\n\n■02\u3000実験準備\u3000\u3000★☆☆☆☆☆☆☆☆☆\n勃起の強制を行います。\n耳元で執拗に唾液の淫らな音を流し込みます。\n\n■03\u3000我慢実験\u3000\u3000★★☆☆☆☆☆☆☆☆\n扱きたい気持ちへの我慢を実験します。\nじっくり嬲るようなスピードでのオナニーを強制します。\n\n■04\u3000性交実験\u3000\u3000★★★☆☆☆☆☆☆☆\n女性との性行為で興奮するかを実験します。\n本格的に扱かせ、ペニスへ負荷をかけていきます。\n\n■05\u3000水音実験\u3000\u3000★★★★☆☆☆☆☆☆\nフェラチオの真似で興奮するかを実験します。\n水音に連動した激しいオナニーで責め立てます。\n\n■06\u3000羞恥実験\u3000\u3000★★★★★☆☆☆☆☆\n恥ずかしい行為で興奮するかを実験します。\n卑猥で情けない言葉を連呼させながらいたぶります。\n\n■07\u3000罵倒実験\u3000\u3000★★★★★★☆☆☆☆\n見下したような罵倒で興奮するかを実験します。\n感情もなく厳しい言葉で激しく追い打ちをかけます。\n\n■08\u3000耐久実験\u3000\u3000★★★★★★★★☆☆\n射精させるつもりで行う追加実験です。\n延々と長時間扱かせ、ペニスを狂わせます。\n(ベリーハードモード、ヘルモードのみ)\n\n■09\u3000絶望実験\u3000\u3000★★★★★★★★☆☆\n射精させるつもりで行う追加実験です。\n何度も最高速で扱かせ、ペニスを射精させようとします。\n(ベリーハードモード、ヘルモードのみ)\n\n■10\u3000最終確認\u3000\u3000★★★★★★★★★★\n極めて絶望的な執行を行います。\n内容はご自身でお確かめくださいませ。\n女性1、2が登場します。\n(ヘルモードのみ)\n\n■11\u3000射精許可\u3000\u3000測定不能\n内容はご自身でお確かめくださいませ。\n女性1、2、3、4が登場します。\n(ヘルモードのみ特別仕様)\n\n\n視聴時間は、\n●ハードモード\u3000\u3000\u3000\u3000: 63:15\n●ベリーハードモード\u3000: 75:45\n●ヘルモード\u3000\u3000\u3000\u3000\u3000:107:15\n(いずれも『00注意事項』の時間は含まない)\nとなっております。\n\n\n体験版には、\n00注意事項\n01プロローグ\n作中シーンの一部を3つ\nを収録。\n\n\n声優\n機械人形:ゅかにゃん様\n女性1:柚凪様\n女性2:井上果林様\n女性3:大山チロル様\n女性4:西浦のどか様\n\nイラスト\nHobby様\n\n本作品の販売\nB-bishop http://pawnlank7.blog.fc2.com",
	}
	testParseWork(t, c, "work", exp)
}

func TestParseWorkWithoutSeries(t *testing.T) {
	t.Parallel()
	c := "RJ173248"
	exp := &Work{
		RJCode:      RJCode(c),
		Name:        "搾精天使ピュアミルク 背後からバイノーラルでいじめられる音声",
		Maker:       "B-bishop",
		Description: "皆様、こんにちは。サークルB-bishopのpawnlank7と申します。\n\n今作は、バイノーラル収録作品となります。\n\n背後から抱きかかえられて囁かれる表現にこだわりました。\n左に右に、近くと遠くで、その変化をお楽しみください。\nとことんまで寸止めされ、焦らされ、馬鹿にされ、いじめられましょう。\n\nSEなし版あり\n\n☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆収録トラック☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆☆\n\n■1.プロローグ\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u300009:39\n貴方がえっちな音声を聞いているところに現れたピュアミルク…。\n彼女はザーメンミルクを貴方に捧げるようにお願いする。\n生意気な彼女に反抗的な貴方、しかし、天使の唾液であっという間に…。\n\n■2.手コキ編\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u300025:08\n背後から抱きかかえられながらの手コキです。\n媚薬効果のある天使の唾液をたっぷり塗りこまれながらのスロー手コキをされます。\n\n■3.足コキ編\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u300014:54\n背後からホールドされながらの足コキです。\n情けなくも足裏で挟み込まれながら、たっぷり可愛がられます。\n\n■4.○○コキ編\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u300024:55\n背後から拘束されて股を強制的に開かされながらの○○コキです。\n搾精天使のマジカルアイテムで一撃で射精まで追い込まれます。\n\n■5.エピローグ\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u300001:47\nザーメンミルクを回収し終わったピュアミルクからのお言葉です。\n\n■おまけ\u3000冒頭えっちボイス\u3000\u3000\u3000\u300005:10\n冒頭で流れるえっちな音声を全編です。\n\n■おまけ\u3000両耳耳舐め\u3000\u3000\u3000\u3000\u3000\u3000\u300000:48\u3000\u3000\u3000\u3000\u3000\n両方の耳を舐め続ける音声です。\nループ再生に対応しています。\n\n\n合計\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u300082:21\n\n\n\n声優\n餅よもぎ様\n七海うと様\n\nイラスト\nふれいむタン様\n\n効果音\nシロクマの嫁様\n\n\n\n本作品の販売\nB-bishop http://pawnlank7.blog.fc2.com/",
	}
	testParseWork(t, c, "work", exp)
}

func TestParseWorkWithTracklist(t *testing.T) {
	t.Parallel()
	c := "RJ126928"
	exp := &Work{
		RJCode: RJCode(c),
		Name:   "まじこスハロウィン -可愛い彼女は吸血鬼!? 妖しく光る魅了の魔眼の巻-",
		Maker:  "クッキーボイス",
		Series: "ハロウィンパーティー",
		TrackList: []Track{
			Track{
				Name: "WELCOME\u3000TO\u3000HALLOWEEN",
				Text: "『ハロウィンの招待状』\u3000約1分\u3000※BGMの有無選択可能",
			},
			Track{
				Name: "魅了吸血お漏らし",
				Text: "『ヒロインの来訪、魅了束縛でのイタズラ』\u3000約18分",
			},
			Track{
				Name: "フェラチオきば愛情",
				Text: "『精液を摂取する、吸血鬼』\t約16分",
			},
			Track{
				Name: "あなたもう゛ぁんぷ",
				Text: "『童貞卒業、牙で噛み合い』\t約15分",
			},
			Track{
				Name: "突発ラジオ",
				Text: "『ラブラブ朝のグッドモーニング放送局』\t約6分",
			},
		},
		Description: "☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆\n\n月の魔力の影響を色濃く見受けられるこの街で\u3000不定期に開催される『7日間ハロウィン』\n\nそして\u3000その期間中にだけ出回る\u3000”魔法の仮装衣装”\u3000通称『まじこス』\n\n\u3000\u3000\u3000『魔女が仕立ててるんだって』\n\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000『一晩だけ本物の魔物になっちゃうの』\n\u3000\u3000\u3000\u3000\u3000\u3000\u3000『着ると幸運が訪れるらしいよ』\n\n\u3000\u3000\u3000\u3000\u3000\u3000『だけど、お菓子がもらえないと\u3000えっちな呪いにかかっちゃうんだよ』\n\n\u3000\u3000\u3000『ほんと?\u3000なんだか怖い……』\n\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000『んふふ\u3000それがいいんじゃない』\n\n☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆\n\n\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000今宵もまた\u3000楽しい楽しいハロウィーン\n\n\u3000\u3000\u3000\u3000\u3000\u3000魔法に身を包んだ少女が一人\u3000あなたの家の扉を叩き訪ねます\n\n\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000― Trick or Treat? ―\n\n☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆\n\n◆内容紹介◆\n\n\u3000普段は素っ気ない気だるげな美少女月城ルナは、\n\u3000ハロウィーンの魔力でホンモノの吸血鬼になっていました。\n\u3000あなたは、決して抗えない魅了の魔眼で束縛され、\n\u3000その首に甘い牙の口づけを受け、全身をとろとろにとかされていきます……。\n\n◆ドキドキえっちなセリフがいっぱい◆\n\n「あなたは、もう、吸血鬼の私のチャームに完全にかかってしまいました。\n\u3000どうしてか、胸がドキドキとして……嬉しくなってきました。\n\u3000これは、ご褒美です……どうぞ、欲情をもっとかき立てられてください」\n\n「あむ、吸血鬼になっているためか、あなたの血なら、いくらでも吸えそうなほど、甘く感じます……。\n\u3000ん、ちゅ、ちゅ……。\n\u3000ん、大人しくしていてください……。痛みはほとんどないのに、恥ずかしかったりして、動くのは……往生際が悪いというか……」\n\n「んっ……はぁ……見てください、全部、奥まで、入りましたよ……。童貞卒業、おめでとうございます。\n\u3000わかりますよね、嬉しそうにおち〇ちんが、奥のほうでひくひくと震えていますよ。催促するように、ひくひくって、ほら、また……。\n\u3000もっと、キュっとしめてみましょうか……。\n\u3000んっ……完全に私の虜ですね、おち〇ちんどころか、あなたからまで、嬉しそうな悲鳴があがりました」\n\n☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆;:*:;☆\n\nCV \u3000\u3000\u3000:分倍河原シホ\u3000\u3000\u3000\u3000\u3000\u3000\u3000http://shiho.moe.in/v/\nCG \u3000\u3000\u3000:皐月みすず\u3000\u3000\u3000\u3000\u3000\u3000\u3000\u3000http://satukinchi.xxxxxxxx.jp/\nBGM :フリー音楽素材こんとどぅふぇ\u3000http://conte-de-fees.com/\n\n企画 \u3000\u3000:クッキーボイス",
	}
	testParseWork(t, c, "work", exp)
}

func TestParseWorkAnnounce(t *testing.T) {
	t.Parallel()
	c := "RJ189666"
	exp := &Work{
		RJCode:      RJCode(c),
		Name:        "強気な妹に連射させられる!? ～即ヌキ淫語16～",
		Maker:       "S彼女",
		Series:      "即ヌキ淫語",
		Description: "即ヌキ淫語シリーズ第十六弾は、強気な妹があなたを性のオモチャにしてるお話。\n\n兄(あなた)が射精する様が大好きな妹は、一回射精させるだけでは飽き足らず、毎日のように連続で射精させたがる。\nあなたは早漏だけど、連続射精はできてしまう。\nそれが楽しくて、妹は今日もあなたのペニスを様々な方法で弄り倒して、兄チ○ポを射精させまくる。\n\n\n01\n\u3000見られてオナニー、オッパイ吸ってオナニー、手コキ、パンツ手コキ、の4連続射精\n\n02\n\u3000淫語でオナニー、オッパイに擦り付け、69体勢で軽いフェラ、クンニしながらオナニー、の4連続射精\n\n03\n\u3000フェラチオ、素股、正常位、騎乗位、の4連続射精\n\n\n妹にオモチャにされる時間は合計で\u300000:54:22\n製品版は、SEあり版(A)/SEなし版(B)の2種類がセットになっていますので、お好きな方をチョイスしてお聞きいただけます。\n※内容によってSEのない場合もあります。\n内容のイメージは体験版(約6分)でご確認ください。\n\n声優 : 今谷皆美\nイラスト : 神瀬あから",
	}
	testParseWork(t, c, "announce", exp)
}

func testParseWork(t *testing.T, rjcode string, p string, exp *Work) {
	c := Parse(rjcode)
	f, err := os.Open(testdataPath(p, c))
	if err != nil {
		t.Fatalf("Error opening test file: %s", err)
	}
	defer f.Close()
	w, err := parseWork(RJCode(c), f)
	if err != nil {
		t.Fatalf("Error parsing work: %s", err)
	}
	if w.RJCode != RJCode(c) {
		t.Errorf("Expected RJCode %#v, got %#v", c, w.RJCode)
	}
	if exp.Name != w.Name {
		t.Errorf("Expected Name %#v, got %#v", exp.Name, w.Name)
	}
	if exp.Maker != w.Maker {
		t.Errorf("Expected Maker %#v, got %#v", exp.Maker, w.Maker)
	}
	if exp.Series != w.Series {
		t.Errorf("Expected Series %#v, got %#v", exp.Series, w.Series)
	}
	if exp.Description != w.Description {
		t.Errorf("Expected Description %#v, got %#v", exp.Description, w.Description)
	}
	if !reflect.DeepEqual(exp.TrackList, w.TrackList) {
		t.Errorf("Expected TrackList %#v, got %#v", exp.TrackList, w.TrackList)
	}
}