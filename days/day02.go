package days

import (
	"fmt"
	"strings"
)

type Day02 struct {
	data string
}

func (d *Day02) Name() string {
	return "Inventory Management System"
}

func (d *Day02) Run() (err error) {
	values := strings.Split(d.data, "\n")

	// part 1 is a simple "checksum"
	c := d.checksum(values)
	fmt.Printf("part 1 = %d\n", c)

	// part 2 is an evil n^2 string compare
	// there has to be a faster way of doing this (turn the strings into numbers somehow...) but meh, brute force it is
	for a := 0; a < len(values); a++ {
		for b := a + 1; b < len(values); b++ {
			if idx, ok := d.delta(values[a], values[b]); ok {
				fmt.Printf("(\nidx = %d\n  a = %s\n  b = %s\n)\n", idx, values[a], values[b])
				var sa, sb string
				if idx > 0 {
					sa = values[a][:idx]
				}
				if idx < len(values) {
					sb = values[b][idx+1:]
				}
				fmt.Printf("part 2 = %s%s\n", sa, sb)
				return
			}
		}
	}

	return
}

func (d *Day02) checksum(values []string) int {
	t2, t3 := 0, 0
	for _, str := range values {
		if d.count(str, 2) {
			t2++
		}
		if d.count(str, 3) {
			t3++
		}
	}
	return t2 * t3
}

func (d *Day02) count(str string, amt int) bool {
	// counting pass
	letters := map[rune]int{}
	for _, k := range strings.TrimSpace(str) {
		if _, ok := letters[k]; ok {
			letters[k]++
		} else {
			letters[k] = 1
		}
	}
	// validation pass
	for _, v := range letters {
		if v == amt {
			return true
		}
	}
	return false
}

func (d *Day02) delta(a string, b string) (idx int, ok bool) {
	for k := 0; k < len(a); k++ {
		if a[k] != b[k] {
			if ok {
				return -1, false
			} else {
				idx = k
				ok = true
			}
		}
	}
	return
}

func (d *Day02) Init() {
	d.data = `mgdsdwqulhrinkatfjoyzxcbvq
megsdwpulhrinkasfjdyzxcbvk
megsdwpulhmibkatfjoyzccbvq
megsdwtuohrinkatfjoyzecbvq
megsdwpulhdinkatfyoyzxcbvb
megszwpulhrirkatfjoyexcbvq
eegsdwpulhzinkatfjoyzxdbvq
megsdwpulhginkatfjoyqucbvq
leqsdwpulhrinkatfeoyzxcbvq
megsdwpblhrandatfjoyzxcbvq
megrdwpudhrinvatfjoyzxcbvq
mlgsdwpulhrinkatfuoyzxcbyq
mensdwpunhrinsatfjoyzxcbvq
megsdwpnlhrinqatfjoyzxcbuq
megsdwpnlhrinkatfqoysxcbvq
megsdwpulhrinkjtfjoywxobvq
megsdwpclhbinkaffjoyzxcbvq
mersdwpulhrinkattjoyzxcbdq
megsddpwohrinkatfjoyzxcbvq
mejsdwpulhrinratfjoyzxcbvl
megsdwpujhrinkatfjkbzxcbvq
megsdwpujhrinkatfjoyircbvq
megsdwgulhrinkatojoyzxcbvm
qegsdwpulhrinkatfjoyzxggvq
megsdwpflhrinkatfjsazxcbvq
megsdwpulhridkagfjoyzzcbvq
mejsdwpulhriykatfjoyyxcbvq
measdwpulhriakatfjoyzacbvq
mebsdwpulhrinkatfjoyzxmyvq
megsdwpflnrinkatfjoyzxcbxq
megsdwpulhxinuatfjoytxcbvq
megsdwsflhrinkatfjoyzxcmvq
megsdwpuuhritkarfjoyzxcbvq
megsdwpulhrinkazfjoyexcjvq
megsdwpulhrifkatfjokzxcbvc
megsdwmulvrinkatfjoyzxcblq
meggdwpulhrfnkatfjoybxcbvq
megsdwptlhrindatyjoyzxcbvq
megsdwpulhrinkauqjoyzxibvq
megsdwddlhrifkatfjoyzxcbvq
megddwpulhrinkagfjoyzxcbva
messdupulhrinkatfjoyzxqbvq
megsdwpuahrinkatfjotzxchvq
megsewpulhrankathjoyzxcbvq
megsdwpulhrinkatfjoynqcbvv
megsdwpulhrifkatfjotzxcbtq
megsdjpulcrinkitfjoyzxcbvq
fefsowpulhrinkatfjoyzxcbvq
megsdwpulapinkatfjoyzvcbvq
mejsdwpularinkatfjonzxcbvq
megsxwpllrrinkatfjoyzxcbvq
megsdlpulxvinkatfjoyzxcbvq
megsdwpulhriekztfjoyzxxbvq
megspwpulhrinuatfjoyzxcgvq
megqdwpulhrinkatfjoyzujbvq
megsdtpuehrinkatfjoyzxtbvq
megsdwjulhrmnmatfjoyzxcbvq
megjdwpulhrinkztfjoyzpcbvq
megsdwpuljrinkawfjlyzxcbvq
mewsdwpulhtinkatfjodzxcbvq
megsdwpulhrinkatfcoizxcmvq
meisdwpulhrinkatawoyzxcbvq
megshwpulhrinkatfjoyzxcbtz
megsdwpulhfinkatfjowzxcbvj
megodwpulhrcskatfjoyzxcbvq
megydxpulhrinkatfjofzxcbvq
megsdwpulhrinkatljoytxdbvq
mwgsdwpulyrinkatfjoyzxvbvq
meesdwpulhrinkatfjcwzxcbvq
cegsdwpwlhrinkatfioyzxcbvq
megsdwpulhrinkatfqhyxxcbvq
megsrrpulhrxnkatfjoyzxcbvq
megsdgpulhrinkatfjoypxcbnq
xegsdwpulhrinkaxkjoyzxcbvq
mngsdwpulhrenkatfjofzxcbvq
mecsdwpulhrinkabfjoozxcbvq
pegsdkpulhrinkatfuoyzxcbvq
megpdwpughrinkatfjoyzxzbvq
mxgrdwpulhrinyatfjoyzxcbvq
megsdwpulhpigkatfjoyzxcbfq
mhgsdwpulhrinkagfjoxzxcbvq
megsdwpblhtinkatfjqyzxcbvq
megsawvulhrinkatfjoyzxcfvq
metsuwpulhrinkatfjowzxcbvq
mexsdwpulhrinkatsjoyzxvbvq
megtdwpulhrinkatfjoizxcpvq
megidwpjlhrinkatfdoyzxcbvq
megsdwpuljrivkatfjoyqxcbvq
mespdppulhrinkatfjoyzxcbvq
megsdcpulhminkatftoyzxcbvq
megsdwpulhriskatfjayzlcbvq
mejsdwpulhrinkatfvoyzxibvq
megfdwputhriykatfjoyzxcbvq
megsdwparhrinkbtfjoyzxcbvq
mtgsdwpulhmingatfjoyzxcbvq
megsdwjulhrinkatfjtyzxtbvq
megsvwpulhrinkatfjoyzecbvf
megsdwpulprinkatyqoyzxcbvq
zegsdwpulhwinkatfjoyzxcfvq
megsdwiulhrinkatfjoyzxcjcq
aegsdwpulirinkatfjoyvxcbvq
mepsdwculbrinkatfjoyzxcbvq
msgsdcpulhrinkctfjoyzxcbvq
megsdwpulhrinkatfjoozxfevq
megscwpulhrinkatmjoyzxcbsq
megsdwpulhrinkltfjoyzxtmvq
megsddpulhrixkatfjyyzxcbvq
megsdwpulhrinkctfjoqzhcbvq
megsbapblhrinkatfjoyzxcbvq
megsdwpulhrinkatfjoyyxgpvq
mwpsdwpulhrinkutfjoyzxcbvq
megsdwpulhrinvnlfjoyzxcbvq
megsdwsulorinkakfjoyzxcbvq
mbgsdwpulhrinkktfjoyzxcbvg
megddwpulhrindatfjoyzxcbvr
memsdwpulhrinkatqjoyzxcbvo
megsdwpulbrinkarfjoyzxcbxq
megsdwpulhrilkatfjoynxcgvq
mxgsdzpulhrinkatfjoydxcbvq
megsawpulhrinkotfjoyzacbvq
measdwpulhrinkctvjoyzxcbvq
medsdwpulhtinkaqfjoyzxcbvq
megsdwsulhrknkatfjoylxcbvq
megadwpulhrinkatfjkyzxcbvh
mygsdwpulhrinkatfboyzxcavq
megsdwpulhrinkaufjxyzccbvq
mtgsdwpulhrunkatfkoyzxcbvq
megsdwpulhrinkamfwryzxcbvq
meesdwpulhrgnkatnjoyzxcbvq
megddwpolhrinkatfjoyixcbvq
megsdwpulhrinsatfjoyzicwvq
mmgsdwpulhrinkltfjoyzxcgvq
meisdwpulhhinkatfjoyzxubvq
megpdwpulhrinkatfroyzxcbkq
megszwpuhhrinkatfjoyzxcbzq
mersdwpulhrikcatfjoyzxcbvq
mjgsdwpulhrinratfjoyzxcbvl
megsdwnulkxinkatfjoyzxcbvq
megsdwpulhrinkatlaoyzxpbvq
megsdwpulhribkitfjjyzxcbvq
megsdwpzlhrinkalfjoyzkcbvq
megsdwpulhrinkutfjoyzxcdnq
mfgpdwputhrinkatfjoyzxcbvq
oegshwpclhrinkatfjoyzxcbvq
megvdwxulhrinkatfjowzxcbvq
mrgydwpulhrinkatfjoyzxcbvl
megsdwpuqhrinkaxfjoyzxcbfq
megsdwpughrinixtfjoyzxcbvq
megsdwpslxrinkatfjoyzxcbvl
megsdwpukhrinkatfjovzxnbvq
megsdwpulnrinkrtojoyzxcbvq
megsdwpuvhrinkatfjjyzxclvq
vegrdwpulhrinzatfjoyzxcbvq
melsdwwulvrinkatfjoyzxcbvq
megsdwpulhrizkbtfjoyzjcbvq
megsdwpulhrinlatfjoyzkcbvr
megsdweulhrinkaofqoyzxcbvq
megsdcpulhrifkatfdoyzxcbvq
negsdwpulhainkztfjoyzxcbvq
megsdwpklhrinkttfjoyzxvbvq
megsdwpothrdnkatfjoyzxcbvq
megsdwpulvrinkatfaoyzxcbvr
mewsdwpulhrinkatfooyzxtbvq
megtdwpulhrinkatfxoyzxcbvy
zegsdwpulhrinkamljoyzxcbvq
mkgodwpulhrcnkatfjoyzxcbvq
megsdepularinkatfjoyzxybvq
segsdwpulhrinkatfmoyzycbvq
mbgsdwpulhrsnkztfjoyzxcbvq
megsdwpmlhrinkatfjxyzhcbvq
megsdwpulhrinkatfboyqxmbvq
megsdwpulhrinkadtpoyzxcbvq
megsdmpulhrjnkatfjoyzxwbvq
megsdwpulhrivkatfroywxcbvq
megsdwdulhrinkatpjwyzxcbvq
megsdwpylhrdrkatfjoyzxcbvq
megsdwpulhrvnkptfjoyzlcbvq
megsdwpuihridkatfjuyzxcbvq
megsdwpulhbiikatfjxyzxcbvq
meggdwpulurinkatfjoyaxcbvq
mehsnwpulcrinkatfjoyzxcbvq
megsdwpulhrinkatfjoczdclvq
megsdwjulhrinkatfhkyzxcbvq
megsdwpulhrinoatfjoyzxcwvy
megpdwpelhrfnkatfjoyzxcbvq
megsjwpulhrwnkatfjdyzxcbvq
megsdwpulhrbniatfjoyzxcbpq
megsdwpulhrwnkatfjoydxcpvq
megsdwrulhrinkatfjvyixcbvq
yegvdwpulhtinkatfjoyzxcbvq
meosdwpulhrinkatfjhyzxsbvq
mjgsdwpulhrinkatfjoyyxcjvq
megsdwpulhrinkatfnoyzlmbvq
mehsdfkulhrinkatfjoyzxcbvq
megsdwpuahrinkatfyoyzxcwvq
megsdwpulhyinkatfjobzkcbvq
megsdwpjlhrinkatfnoyzscbvq
megsdwiulhrinqatfjoyzxcavq
megsdwpulurinkatfjoykxcbmq
megsdwpmlhvinkatfjotzxcbvq
megswwpulhrinfatfjoytxcbvq
mxgsdwpulhrrnkatfjoyhxcbvq
megubwpulhrinkaafjoyzxcbvq
megsdwpulhrdnsztfjoyzxcbvq
megsdwpulhronkgtfjhyzxcbvq
megsdwzulhfinratfjoyzxcbvq
megsdwpulhrinkrtfjqyixcbvq
megspwpulhrinkatfjoywxcavq
megmfjpulhrinkatfjoyzxcbvq
lelsdwpulhrinsatfjoyzxcbvq
megsdwnultrinkatfjoyzxybvq
megsdwpulhdnnkatfjpyzxcbvq
megsdwpulhzinkatfjoyzhcpvq
megoddpulhrilkatfjoyzxcbvq
megsdvpulhrwnkatfjoyzxcbdq
megsdwpulhmikkatfjoyzxcbvr
meesdwpulhrinkaofjoazxcbvq
megndwhulhrinfatfjoyzxcbvq
megsdwpulhjihkatfjoyzxcbvb
jegsdwpulhrinkattjoszxcbvq
megsdwpulhrenkatfjodzxcbvp
mefsdwpulhrinkatfjoyzxcczq
megsdwpulhrknkatfjovzxcbvb
megytwpulhrinkatfjoyzxcbvo
csgsdwpulhrinkatdjoyzxcbvq
megsbwpulhrinkazfjtyzxcbvq
mrgsdwpulhjinkatfjoyzxcbvm
megsywzulhrinkatfjofzxcbvq
megxdwpulhrndkatfjoyzxcbvq
mcgsdwpuphrixkatfjoyzxcbvq
megsdlpulxvinkatfvoyzxcbvq
megsdwpulhrinfatfjoyzicbvo
megwgwpulhrinpatfjoyzxcbvq
mggsdwpulhrinkatfjtrzxcbvq
msgsdwpulirinkatfjoyzxcbaq
megsdwpulhrineatfeoyzxcbvl
megsdnpulhrinkatojoyzxcbvb
mfgsdwpulhdinkatmjoyzxcbvq
megsdwpulhrinkawfjuyzxcbvg
megsdwpulhrincatfjoyzxcxvu
mevsdwzplhrinkatfjoyzxcbvq
megskweulhfinkatfjoyzxcbvq
mebsdvpulhrinkatfjoyzxcbkq
megddwiulhrinkatfioyzxcbvq
mygsdwpulhqinkatfooyzxcbvq
megsdweulhrfnkatfjoybxcbvq
mrgsdhpulmrinkatfjoyzxcbvq
megsdwkukhrinkatftoyzxcbvq
megsqwpulhrinfatfjoyzxebvq
megsdwpulhriskanfjoyzxctvq`
}
