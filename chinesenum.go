package chinesenum

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"strings"
)

// Option 是转换时的相关额外配置
type Option struct {
	// UseOralTwo 表示在转换中采用“两”字
	UseOralTwo bool
}

// Convertor 代表转换器对象
type Convertor interface {
	// Itoa 将一个整型转为汉字
	Itoa(int, ...Option) string
	// // Ftoa 讲一个 float64 浮点数转为汉字，prec 代表小数点后的位数。传入 -1 表示自动
	// Ftoa(float64, prec int) string
}

// Lang 表示语言文字类型
type Lang string

const (
	// ZhHans 代表简体中文小写，最大支持到千兆
	ZhHans = "zh-Hans"
	// ZhHansUpper 代表简体中文大写，最大支持到千兆
	ZhHansUpper = "zh-Hans_upper"
)

var (
	hansNumbers       = []rune{'零', '一', '二', '三', '四', '五', '六', '七', '八', '九'}
	hansCarries       = []rune{'十', '百', '千', '万', '亿', '兆'}
	hansNegative      = '负'
	hansOralTwo       = '两'
	hansUpperNumbers  = []rune{'零', '壹', '贰', '叁', '肆', '伍', '陆', '柒', '捌', '玖'}
	hansUpperCarries  = []rune{'拾', '佰', '仟', '萬', '億', '兆'}
	hansUpperNegative = '负'
	hansUpperOralTwo  = '贰'
)

const (
	n0 = iota
	n1
	n2
	n3
	n4
	n5
	n6
	n7
	n8
	n9
)

const (
	cShi = iota
	cBai
	cQian
	cWan
	cYi
	cZhao
)

// Get 获取一个转换器
func Get(lang Lang) Convertor {
	switch lang {
	default:
		return hans
	case ZhHans:
		return hans
	case ZhHansUpper:
		return hansUpper
	}
}

var (
	hans      *conv
	hansUpper *conv
)

func init() {
	hans = &conv{
		numbers:  hansNumbers,
		carries:  hansCarries,
		negative: hansNegative,
		oralTwo:  hansOralTwo,
	}
	hans.genOralTwoReplacer()

	hansUpper = &conv{
		numbers:  hansUpperNumbers,
		carries:  hansUpperCarries,
		negative: hansUpperNegative,
		oralTwo:  hansUpperOralTwo,
	}
	hansUpper.genOralTwoReplacer()

	return
}

type conv struct {
	numbers          []rune
	carries          []rune
	negative         rune
	oralTwo          rune
	oralTwoReplacerA *strings.Replacer
	oralTwoReplacerB *strings.Replacer
}

func (c *conv) Itoa(i int, opt ...Option) string {
	if 0 == i {
		// 直接返回零
		return string(c.numbers[n0])
	}
	if 2 == i && len(opt) > 0 && opt[0].UseOralTwo {
		return string(c.oralTwo)
	}

	negative := false
	if i < 0 {
		negative = true
		i = -i
	}

	// 首先转成字符串来处理
	plain := strconv.Itoa(i)
	const max = len("10000000000000000") // 一万兆，不支持
	if len(plain) >= max {
		plain = plain[0:max]
	}

	parts := partsFromStr(plain)
	s := parts.totalToStr(c)

	if negative {
		return string(c.negative) + s
	}

	// 替换“两”字
	if len(opt) > 0 && opt[0].UseOralTwo {
		s = c.convOralTwo(s)
	}
	return s
}

func (c *conv) convOralTwo(s string) string {
	s = c.oralTwoReplacerA.Replace(s)
	s = c.oralTwoReplacerB.Replace(s)
	return s
}

func (c *conv) genOralTwoReplacer() {
	// 两百
	twoBaiFrom := string(c.numbers[n2]) + string(c.carries[cBai])
	twoBaiTo := string(c.oralTwo) + string(c.carries[cBai])
	// 两千
	twoQianFrom := string(c.numbers[n2]) + string(c.carries[cQian])
	twoQianTo := string(c.oralTwo) + string(c.carries[cQian])
	// 两万
	twoWanFrom := string(c.numbers[n2]) + string(c.carries[cWan])
	twoWanTo := string(c.oralTwo) + string(c.carries[cWan])
	// 两亿
	twoYiFrom := string(c.numbers[n2]) + string(c.carries[cYi])
	twoYiTo := string(c.oralTwo) + string(c.carries[cYi])
	// 两兆
	twoZhaoFrom := string(c.numbers[n2]) + string(c.carries[cZhao])
	twoZhaoTo := string(c.oralTwo) + string(c.carries[cZhao])
	// 十二万。这个会被替换成十两万，要换回来
	twelveWanFrom := fmt.Sprintf("%c%c%c", c.carries[cShi], c.oralTwo, c.carries[cWan])
	twelveWanTo := fmt.Sprintf("%c%c%c", c.carries[cShi], c.numbers[2], c.carries[cWan])
	// 十二亿
	twelveYiFrom := fmt.Sprintf("%c%c%c", c.carries[cShi], c.oralTwo, c.carries[cYi])
	twelveYiTo := fmt.Sprintf("%c%c%c", c.carries[cShi], c.numbers[2], c.carries[cYi])
	// 十二兆
	twelveZhaoFrom := fmt.Sprintf("%c%c%c", c.carries[cShi], c.oralTwo, c.carries[cZhao])
	twelveZhaoTo := fmt.Sprintf("%c%c%c", c.carries[cShi], c.numbers[2], c.carries[cZhao])

	c.oralTwoReplacerA = strings.NewReplacer(
		twoBaiFrom, twoBaiTo,
		twoQianFrom, twoQianTo,
		twoWanFrom, twoWanTo,
		twoYiFrom, twoYiTo,
		twoZhaoFrom, twoZhaoTo,
	)
	c.oralTwoReplacerB = strings.NewReplacer(
		twelveWanFrom, twelveWanTo,
		twelveYiFrom, twelveYiTo,
		twelveZhaoFrom, twelveZhaoTo,
	)
	return
}

func (c *conv) zero() string {
	return string(c.numbers[n0])
}

type parts struct {
	// 以下从高到低按照每四个数字进位
	p []string
}

func partsFromStr(s string) *parts {
	ret := parts{}
	// 首先按照字符串尺寸切割
	l := len(s)
	switch l {
	case 1, 2, 3, 4:
		ret.p = []string{s}
	case 5, 6, 7, 8:
		ret.p = []string{
			s[0 : l-4],
			s[l-4 : l],
		}
	case 9, 10, 11, 12:
		ret.p = []string{
			s[0 : l-8],
			s[l-8 : l-4],
			s[l-4 : l],
		}
	default: // default 就是 13, 14, 15, 16，不会出现其他情况
		ret.p = []string{
			s[0 : l-12],
			s[l-12 : l-8],
			s[l-8 : l-4],
			s[l-4 : l],
		}
	}
	return &ret
}

func (p *parts) totalToStr(c *conv) string {
	partStrs := make([]string, 4)
	ten := fmt.Sprintf("%c%c", c.numbers[n1], c.carries[cShi]) // 一十
	l := len(p.p)

	for i, s := range p.p {
		off := 4 - l + i
		// log.Print("off:", off)
		chn := p.partToStr(c, s)
		chn = strings.Trim(chn, c.zero())
		if chn == "" {
			partStrs[off] = c.zero()
			continue
		}
		if i == 0 && strings.HasPrefix(chn, ten) {
			chn = trimLeftChr(chn, c.numbers[n1])
		}
		if i < l-1 {
			chn += string(c.carries[cZhao-off])
		}

		// 要不要在左边补上一个零
		if i > 0 {
			if len(chn) < 4 || '0' == chn[0] {
				chn = c.zero() + chn
			} else {
				prev := p.p[i-1]
				if prev != "" && prev[len(prev)-1] == '0' {
					chn = c.zero() + chn
				}
			}
		}
		partStrs[off] = chn
	}

	s := strings.Join(partStrs, "")
	strings.TrimLeft(s, c.zero())
	if strings.HasPrefix(s, ten) {
		s = trimLeftChr(s, c.numbers[n1])
	}

	log.Printf("total: %s", s)

	// 去掉中间多余的零
	buff := bytes.Buffer{}
	zeroCount := 0
	for _, chr := range s {
		if chr == c.numbers[n0] {
			if zeroCount == 1 {
				continue
			} else {
				zeroCount++
			}
		}
		buff.WriteRune(chr)
	}
	return strings.Trim(buff.String(), c.zero())
}

func trimLeftChr(s string, chr rune) string {
	return strings.TrimLeft(s, string(chr))
}

func trimRightChr(s string, chr rune) string {
	return strings.TrimRight(s, string(chr))
}

func (p *parts) partToStr(c *conv, s string) string {
	buff := bytes.Buffer{}
	l := len(s)

	for i, chr := range s {
		n := int(chr - '0')
		if 0 == n {
			buff.WriteRune(c.numbers[n0]) // 零
		} else {
			buff.WriteRune(c.numbers[n]) // 一、二、三、……
			if i < l-1 {
				buff.WriteRune(c.carries[l-2-i]) // 十、百、千
			}
		}
	}

	log.Printf("%s --> %s", s, buff.String())
	return buff.String() //这里的返回可能会有“一十几”的格式，也可能会有很多个零
}
