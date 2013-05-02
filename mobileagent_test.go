package mobileagent_test

import (
	"github.com/mattn/go-mobileagent"
	"testing"
)

var docomoAgent = []string{
	`DoCoMo/1.0/633S/c20`,
	`DoCoMo/1.0/641P/c10`,
	`DoCoMo/1.0/641S/c10`,
	`DoCoMo/1.0/642S/c20`,
	`DoCoMo/1.0/Agent`,
	`DoCoMo/1.0/D209i`,
	`DoCoMo/1.0/D209i/c10`,
	`DoCoMo/1.0/D210i/c10`,
	`DoCoMo/1.0/D211i/c10`,
	`DoCoMo/1.0/D251i/c10`,
	`DoCoMo/1.0/D501i`,
	`DoCoMo/1.0/D501i/c5`,
	`DoCoMo/1.0/D502i`,
	`DoCoMo/1.0/D502i/c10`,
	`DoCoMo/1.0/D503i`,
	`DoCoMo/1.0/D503i/c10`,
	`DoCoMo/1.0/D503i/c5`,
	`DoCoMo/1.0/D503iS/c10`,
	`DoCoMo/1.0/D503iS/c5`,
	`DoCoMo/1.0/D504i/c10`,
	`DoCoMo/1.0/ER209i`,
	`DoCoMo/1.0/ER209i/c15`,
	`DoCoMo/1.0/F209i`,
	`DoCoMo/1.0/F209i/c10`,
	`DoCoMo/1.0/F210i/c10`,
	`DoCoMo/1.0/F211i/c10`,
	`DoCoMo/1.0/F251i/c10/TB`,
	`DoCoMo/1.0/F501i`,
	`DoCoMo/1.0/F502i`,
	`DoCoMo/1.0/F502i/c10`,
	`DoCoMo/1.0/F502it`,
	`DoCoMo/1.0/F502it/c10`,
	`DoCoMo/1.0/F503i`,
	`DoCoMo/1.0/F503i/c10`,
	`DoCoMo/1.0/F503i/c32`,
	`DoCoMo/1.0/F503iS`,
	`DoCoMo/1.0/F503iS/c10`,
	`DoCoMo/1.0/F504i/c10/TB`,
	`DoCoMo/1.0/F504i/c10/TJ`,
	`DoCoMo/1.0/F671i/c10`,
	`DoCoMo/1.0/GigaCode (http://gigacode.net/)`,
	`DoCoMo/1.0/KO209i`,
	`DoCoMo/1.0/KO210i`,
	`DoCoMo/1.0/KO210i/c10`,
	`DoCoMo/1.0/N209i`,
	`DoCoMo/1.0/N209i/c08`,
	`DoCoMo/1.0/N210i`,
	`DoCoMo/1.0/N210i/c10`,
	`DoCoMo/1.0/N211i/c10`,
	`DoCoMo/1.0/N501i`,
	`DoCoMo/1.0/N502i`,
	`DoCoMo/1.0/N502i/c08`,
	`DoCoMo/1.0/N502it`,
	`DoCoMo/1.0/N502it/c10`,
	`DoCoMo/1.0/N503i`,
	`DoCoMo/1.0/N503i/c10`,
	`DoCoMo/1.0/N503i/c30`,
	`DoCoMo/1.0/N503i/c5`,
	`DoCoMo/1.0/N503i/c5/serNNEBJ608187`,
	`DoCoMo/1.0/N503iS`,
	`DoCoMo/1.0/N503iS/c10`,
	`DoCoMo/1.0/N503iS/c5`,
	`DoCoMo/1.0/N504i/c10`,
	`DoCoMo/1.0/N504i/c10/TB`,
	`DoCoMo/1.0/N504i/c10/TJ`,
	`DoCoMo/1.0/N504i/c10/TJ/c0`,
	`DoCoMo/1.0/N821i`,
	`DoCoMo/1.0/N821i/c08`,
	`DoCoMo/1.0/NM502i`,
	`DoCoMo/1.0/NM502i/c10`,
	`DoCoMo/1.0/P209i`,
	`DoCoMo/1.0/P209i/c10`,
	`DoCoMo/1.0/P209is`,
	`DoCoMo/1.0/P209is (Google CHTML Proxy/1.0)`,
	`DoCoMo/1.0/P209is/c10`,
	`DoCoMo/1.0/P210i`,
	`DoCoMo/1.0/P210i/c10`,
	`DoCoMo/1.0/P211i/c10`,
	`DoCoMo/1.0/P501i`,
	`DoCoMo/1.0/P502i`,
	`DoCoMo/1.0/P502i/c10`,
	`DoCoMo/1.0/P502i/c10 (Google CHTML Proxy/1.0)`,
	`DoCoMo/1.0/P502i_mEB-PD555`,
	`DoCoMo/1.0/P503i`,
	`DoCoMo/1.0/P503i/c10`,
	`DoCoMo/1.0/P503i/c10/`,
	`DoCoMo/1.0/P503iS`,
	`DoCoMo/1.0/P503iS/c10`,
	`DoCoMo/1.0/P503iS/c10/serNMAUA482012`,
	`DoCoMo/1.0/P504i/c10`,
	`DoCoMo/1.0/P504i/c10/TB`,
	`DoCoMo/1.0/P751v/c100/s64/kPHS-K`,
	`DoCoMo/1.0/P821i`,
	`DoCoMo/1.0/P821i/c08`,
	`DoCoMo/1.0/PacketMeter/c10`,
	`DoCoMo/1.0/R209i`,
	`DoCoMo/1.0/R211i/c10`,
	`DoCoMo/1.0/R691i`,
	`DoCoMo/1.0/SH251i/c10`,
	`DoCoMo/1.0/SH712m/c10`,
	`DoCoMo/1.0/SH821i`,
	`DoCoMo/1.0/SH821i/c10`,
	`DoCoMo/1.0/SO210i/c10`,
	`DoCoMo/1.0/SO211i/c10`,
	`DoCoMo/1.0/SO502i`,
	`DoCoMo/1.0/SO502iWM/c10`,
	`DoCoMo/1.0/SO503i`,
	`DoCoMo/1.0/SO503i/c10`,
	`DoCoMo/1.0/SO503i/c10/serNSOBD506895`,
	`DoCoMo/1.0/SO503i/c10/serNSOBD597705`,
	`DoCoMo/1.0/SO503iS/c10`,
	`DoCoMo/1.0/SO504i/c10`,
	`DoCoMo/1.0/SO504i/c10/TB`,
	`DoCoMo/1.0/TEST/c10`,
	`DoCoMo/1.0/TF502i`,
	`DoCoMo/1.0/X503i/c10`,
	`DoCoMo/1.0/eggy/c300/s32/kPHS-K`,
	`DoCoMo/1.0/eggy/c300/s64/kPHS-K`,
	`DoCoMo/1.0/ex_idisplay/c10`,
	`DoCoMo/1.0/ex_ps_test00/c10`,
	`DoCoMo/1.0/iYappo`,
	`DoCoMo/1.0/p503is/c10`,
	`DoCoMo/1.0/test`,
	`DoCoMo/1.0/test/c10`,
	`DoCoMo/1.0/test/c10/TB`,
	`DoCoMo/1.1/P711m/c10`,
	`DoCoMo/2.0 D2101V(c100)`,
	`DoCoMo/2.0 MST_v_P2101V(c100)`,
	`DoCoMo/2.0 N2001(c10)`,
	`DoCoMo/2.0 N2001(c10;ser350200000307969;icc8981100000200188565F)`,
	`DoCoMo/2.0 N2002(c100)`,
	`DoCoMo/2.0 P2002(c100)`,
	`DoCoMo/2.0 P2101V`,
	`DoCoMo/2.0 P2101V(c100)`,
	`DoCoMo/2.0/N502i`,
	`DoCoMo/2.0/N502it`,
	`DoCoMo/2.0/N503i`,
	`DoCoMo/3.0/N503`,
	`DoCoMo/2.0 N06A3(c500;TB;W24H16)`,
	`DoCoMo/2.0 N04A(c100;TB;W24H16)`,
	`DoCoMo/2.0 N08A(c500;TB;W24H16)`,
}

func TestDoCoMo(t *testing.T) {
	for _, userAgent := range docomoAgent {
		if !mobileagent.IsDoCoMo(userAgent) {
			t.Fatal("IsDoCoMo Failed:", userAgent)
		}
	}
}

var jphoneAgents = []string{
	`J-PHONE/1.0`,
	`J-PHONE/2.0/J-DN02`,
	`J-PHONE/2.0/J-P02`,
	`J-PHONE/2.0/J-P03`,
	`J-PHONE/2.0/J-SA02`,
	`J-PHONE/2.0/J-SH02`,
	`J-PHONE/2.0/J-SH03`,
	`J-PHONE/2.0/J-SH03_a`,
	`J-PHONE/2.0/J-SH04`,
	`J-PHONE/2.0/J-T04`,
	`J-PHONE/2.0/J-T05`,
	`J-PHONE/3.0/J-D03`,
	`J-PHONE/3.0/J-D04`,
	`J-PHONE/3.0/J-D05`,
	`J-PHONE/3.0/J-DN03`,
	`J-PHONE/3.0/J-K03`,
	`J-PHONE/3.0/J-K04`,
	`J-PHONE/3.0/J-K05`,
	`J-PHONE/3.0/J-N03`,
	`J-PHONE/3.0/J-N03B`,
	`J-PHONE/3.0/J-N04`,
	`J-PHONE/3.0/J-N05`,
	`J-PHONE/3.0/J-NM01_a`,
	`J-PHONE/3.0/J-NM02`,
	`J-PHONE/3.0/J-PE03`,
	`J-PHONE/3.0/J-PE03_a`,
	`J-PHONE/3.0/J-SA03_a`,
	`J-PHONE/3.0/J-SA04`,
	`J-PHONE/3.0/J-SA04_a`,
	`J-PHONE/3.0/J-SH04`,
	`J-PHONE/3.0/J-SH04_a`,
	`J-PHONE/3.0/J-SH04_b`,
	`J-PHONE/3.0/J-SH04_c`,
	`J-PHONE/3.0/J-SH05`,
	`J-PHONE/3.0/J-SH05_a`,
	`J-PHONE/3.0/J-SH06`,
	`J-PHONE/3.0/J-SH07`,
	`J-PHONE/3.0/J-SH08`,
	`J-PHONE/3.0/J-T05`,
	`J-PHONE/3.0/J-T06`,
	`J-PHONE/3.0/J-T06_a`,
	`J-PHONE/3.0/J-T07`,
	`J-PHONE/4.0/J-K51/SNJKWA3001061 KW/1.00 Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`,
	`J-PHONE/4.0/J-K51/SNJKWA3040744 KW/1.00 Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`,
	`J-PHONE/4.0/J-P51/SNJMAA1036146 MA/JDP51A36 Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`,
	`J-PHONE/4.0/J-SH51 SH/0001aa Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`,
	`J-PHONE/4.0/J-SH51/SNJSHA1032366 SH/0001aa Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`,
	`J-PHONE/4.0/J-SH51/SNJSHA1041639 SH/0001aa Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`,
	`J-PHONE/4.0/J-SH51/SNJSHA2901949 SH/0001aa Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`,
	`J-PHONE/4.0/J-SH51/SNJSHA3008160 SH/0001aa Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`,
	`J-PHONE/4.0/J-SH51/SNJSHA3016183 SH/0001aa Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`,
	`J-PHONE/4.0/J-SH51/SNJSHA3029293 SH/0001aa Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`,
	`J-PHONE/4.0/J-SH51/SNXXXXXXXXX SH/0001a Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`,
	`J-PHONE/4.0/J-SH51_a/SNJSHA1045575 SH/0001aa Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`,
	`J-PHONE/4.0/J-SH51_a/SNJSHA1082487 SH/0001aa Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`,
	`J-PHONE/4.0/J-SH51_a/SNJSHA1086956 SH/0001aa Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`,
	`J-PHONE/4.0/J-SH51_a/SNJSHA3093881 SH/0001aa Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`,
	`J-PHONE/4.0/J-SH51_a/SNJSHA5081372 SH/0001aa Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`,
	`J-PHONE/4.0/J-T51/SNJTSA1077171 TS/1.00 Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`,
	`J-PHONE/4.0/J-T51/SNJTSA1082745 TS/1.00 Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`,
	`J-PHONE/4.0/J-T51/SNJTSA3001961 TS/1.00 Profile/MIDP-1.0 Configuration/CLDC-1.0 Ext-Profile/JSCL-1.1.0`,
}

func TestJPhone(t *testing.T) {
	for _, userAgent := range jphoneAgents {
		if !mobileagent.IsJPhone(userAgent) {
			t.Fatal("IsJPhone Failed:", userAgent)
		}
	}
}

var ezwebAgents = []string{
`KDDI-CA21 UP.Browser/6.0.6 (GUI) MMP/1.1`,
`KDDI-CA21 UP.Browser/6.0.7.1 (GUI) MMP/1.1`,
`KDDI-HI21 UP.Browser/6.0.2.213 (GUI) MMP/1.1`,
`KDDI-HI21 UP.Browser/6.0.2.273 (GUI) MMP/1.1`,
`KDDI-HI21 UP.Browser/6.0.6 (GUI) MMP/1.1`,
`KDDI-KC21 UP.Browser/6.0.2.273 (GUI) MMP/1.1`,
`KDDI-KC21 UP.Browser/6.0.5 (GUI) MMP/1.1`,
`KDDI-KC21 UP.Browser/6.0.6 (GUI) MMP/1.1`,
`KDDI-MA21 UP.Browser/6.0.2.276 (GUI) MMP/1.1`,
`KDDI-MA21 UP.Browser/6.0.5 (GUI) MMP/1.1`,
`KDDI-MA21 UP.Browser/6.0.6 (GUI) MMP/1.1`,
`KDDI-MA21 UP.Browser/6.0.7 (GUI) MMP/1.1`,
`KDDI-SA21 UP.Browser/6.0.6 (GUI) MMP/1.1`,
`KDDI-SA21 UP.Browser/6.0.7 (GUI) MMP/1.1`,
`KDDI-SA21 UP.Browser/6.0.7.1 (GUI) MMP/1.1`,
`KDDI-SA22 UP.Browser/6.0.7.2 (GUI) MMP/1.1`,
`KDDI-SN21 UP.Browser/6.0.7 (GUI) MMP/1.1`,
`KDDI-SN22 UP.Browser/6.0.7 (GUI) MMP/1.1`,
`KDDI-TS21 UP.Browser/6.0.2.273 (GUI) MMP/1.1`,
`KDDI-TS21 UP.Browser/6.0.2.276 (GUI) MMP/1.1`,
`KDDI-TS21 UP.Browser/6.0.5.287 (GUI) MMP/1.1`,
`KDDI-TS21 UP.Browser/6.0.6 (GUI) MMP/1.1`,
`KDDI-TS22 UP.Browser/6.0.6 (GUI) MMP/1.1`,
`KDDI-TS22 UP.Browser/6.0.7.1 (GUI) MMP/1.1`,
`KDDI-TS2A UP.Browser/6.2.0.9 (GUI) MMP/2.0`,
`UP.Browser/3.01-HI01 UP.Link/3.4.5.2`,
`UP.Browser/3.01-HI02 UP.Link/3.2.1.2`,
`UP.Browser/3.03-HI11 UP.Link/3.2.2.7c`,
`UP.Browser/3.03-HI11 UP.Link/3.4.4`,
`UP.Browser/3.03-KCT3 UP.Link/3.4.4`,
`UP.Browser/3.03-SYC1 UP.Link/3.4.4`,
`UP.Browser/3.03-TS11 UP.Link/3.2.2.7c`,
`UP.Browser/3.03-TST1 UP.Link/3.2.2.7c`,
`UP.Browser/3.04-CA11 UP.Link/3.2.2.7c`,
`UP.Browser/3.04-CA11 UP.Link/3.3.0.3`,
`UP.Browser/3.04-CA11 UP.Link/3.3.0.5`,
`UP.Browser/3.04-CA11 UP.Link/3.4.4`,
`UP.Browser/3.04-CA12 UP.Link/3.4.4`,
`UP.Browser/3.04-CA13 UP.Link/3.3.0.5`,
`UP.Browser/3.04-CA13 UP.Link/3.4.4`,
`UP.Browser/3.04-CA14 UP.Link/3.4.4`,
`UP.Browser/3.04-DN11 UP.Link/3.3.0.1`,
`UP.Browser/3.04-DN11 UP.Link/3.4.4`,
`UP.Browser/3.04-HI11 UP.Link/3.2.2.7c`,
`UP.Browser/3.04-HI11 UP.Link/3.4.4`,
`UP.Browser/3.04-HI12 UP.Link/3.2.2.7c`,
`UP.Browser/3.04-HI12 UP.Link/3.3.0.3`,
`UP.Browser/3.04-HI12 UP.Link/3.4.4`,
`UP.Browser/3.04-HI12 UP.Link/3.4.4 (Google WAP Proxy/1.0)`,
`UP.Browser/3.04-HI13 UP.Link/3.4.4`,
`UP.Browser/3.04-HI14 UP.Link/3.4.4`,
`UP.Browser/3.04-HI14 UP.Link/3.4.5.2`,
`UP.Browser/3.04-KC11 UP.Link/3.4.4`,
`UP.Browser/3.04-KC12 UP.Link/3.4.4`,
`UP.Browser/3.04-KC13 UP.Link/3.4.4`,
`UP.Browser/3.04-KC14 UP.Link/3.4.4`,
`UP.Browser/3.04-KC15 UP.Link/3.4.4`,
`UP.Browser/3.04-KCT4 UP.Link/3.4.4`,
`UP.Browser/3.04-KCT5 UP.Link/3.4.4`,
`UP.Browser/3.04-KCT6 UP.Link/3.4.4`,
`UP.Browser/3.04-KCT7 UP.Link/3.4.4`,
`UP.Browser/3.04-KCT8 UP.Link/3.4.4`,
`UP.Browser/3.04-KCT9 UP.Link/3.4.4`,
`UP.Browser/3.04-MA11 UP.Link/3.2.2.7c`,
`UP.Browser/3.04-MA11 UP.Link/3.3.0.3`,
`UP.Browser/3.04-MA11 UP.Link/3.3.0.5`,
`UP.Browser/3.04-MA11 UP.Link/3.4.4`,
`UP.Browser/3.04-MA12 UP.Link/3.2.2.7c`,
`UP.Browser/3.04-MA12 UP.Link/3.3.0.5`,
`UP.Browser/3.04-MA12 UP.Link/3.4.4`,
`UP.Browser/3.04-MA12 UP.Link/3.4.4 (Google WAP Proxy/1.0)`,
`UP.Browser/3.04-MA13 UP.Link/3.3.0.5`,
`UP.Browser/3.04-MA13 UP.Link/3.4.4`,
`UP.Browser/3.04-MA13 UP.Link/3.4.4 (Google WAP Proxy/1.0)`,
`UP.Browser/3.04-MA13 UP.Link/3.4.5.2`,
`UP.Browser/3.04-MAC2 UP.Link/3.4.4`,
`UP.Browser/3.04-MAI1 UP.Link/3.2.2.7c`,
`UP.Browser/3.04-MAI2 UP.Link/3.2.2.7c`,
`UP.Browser/3.04-MAI2 UP.Link/3.4.4`,
`UP.Browser/3.04-MAT1 UP.Link/3.3.0.3`,
`UP.Browser/3.04-MAT3 UP.Link/3.4.4`,
`UP.Browser/3.04-MIT1 UP.Link/3.3.0.3`,
`UP.Browser/3.04-MIT1 UP.Link/3.4.4`,
`UP.Browser/3.04-MIT1 UP.Link/3.4.5.2`,
`UP.Browser/3.04-SN11 UP.Link/3.2.2.7c`,
`UP.Browser/3.04-SN11 UP.Link/3.3.0.3`,
`UP.Browser/3.04-SN11 UP.Link/3.4.4`,
`UP.Browser/3.04-SN11 UP.Link/3.4.4 (Google WAP Proxy/1.0)`,
`UP.Browser/3.04-SN12 UP.Link/3.3.0.1`,
`UP.Browser/3.04-SN12 UP.Link/3.3.0.5`,
`UP.Browser/3.04-SN12 UP.Link/3.4.4`,
`UP.Browser/3.04-SN12 UP.Link/3.4.5.2`,
`UP.Browser/3.04-SN13 UP.Link/3.3.0.3`,
`UP.Browser/3.04-SN13 UP.Link/3.3.0.5`,
`UP.Browser/3.04-SN13 UP.Link/3.4.4`,
`UP.Browser/3.04-SN14 UP.Link/3.4.4`,
`UP.Browser/3.04-SN14 UP.Link/3.4.5.2`,
`UP.Browser/3.04-SN15 UP.Link/3.4.4`,
`UP.Browser/3.04-SN15 UP.Link/3.4.5.2`,
`UP.Browser/3.04-SN16 UP.Link/3.4.4`,
`UP.Browser/3.04-SN17 UP.Link/3.4.4`,
`UP.Browser/3.04-SNI1 UP.Link/3.4.4`,
`UP.Browser/3.04-ST11 UP.Link/3.3.0.1`,
`UP.Browser/3.04-ST11 UP.Link/3.3.0.5`,
`UP.Browser/3.04-ST11 UP.Link/3.4.4`,
`UP.Browser/3.04-ST12 UP.Link/3.4.4`,
`UP.Browser/3.04-ST13 UP.Link/3.4.4`,
`UP.Browser/3.04-SY11 UP.Link/3.2.2.7c`,
`UP.Browser/3.04-SY11 UP.Link/3.4.4`,
`UP.Browser/3.04-SY12 UP.Link/3.3.0.1`,
`UP.Browser/3.04-SY12 UP.Link/3.3.0.3`,
`UP.Browser/3.04-SY12 UP.Link/3.3.0.5`,
`UP.Browser/3.04-SY12 UP.Link/3.4.4`,
`UP.Browser/3.04-SY12 UP.Link/3.4.5.2`,
`UP.Browser/3.04-SY12 UP.Link/3.4.5.6`,
`UP.Browser/3.04-SY13 UP.Link/3.4.4`,
`UP.Browser/3.04-SY14 UP.Link/3.4.4`,
`UP.Browser/3.04-SY15 UP.Link/3.4.4`,
`UP.Browser/3.04-SYT3 UP.Link/3.4.4`,
`UP.Browser/3.04-SYT3 UP.Link/3.4.5.2`,
`UP.Browser/3.04-TS11 UP.Link/3.2.2.7c`,
`UP.Browser/3.04-TS11 UP.Link/3.3.0.5`,
`UP.Browser/3.04-TS11 UP.Link/3.4.4`,
`UP.Browser/3.04-TS12 UP.Link/3.2.2.7c`,
`UP.Browser/3.04-TS12 UP.Link/3.3.0.1`,
`UP.Browser/3.04-TS12 UP.Link/3.3.0.3`,
`UP.Browser/3.04-TS12 UP.Link/3.3.0.5`,
`UP.Browser/3.04-TS12 UP.Link/3.4.4`,
`UP.Browser/3.04-TS13 UP.Link/3.4.4`,
`UP.Browser/3.04-TS14 UP.Link/3.4.4`,
`UP.Browser/3.04-TS14 UP.Link/3.4.4 (Google WAP Proxy/1.0)`,
`UP.Browser/3.04-TS14 UP.Link/3.4.5.2`,
`UP.Browser/3.04-TSI1 UP.Link/3.2.2.7c`,
`UP.Browser/3.04-TST3 UP.Link/3.4.4`,
`UP.Browser/3.04-TST4 UP.Link/3.4.4`,
`UP.Browser/3.04-TST4 UP.Link/3.4.5.2`,
`UP.Browser/3.04-TST4 UP.Link/3.4.5.6`,
`UP.Browser/3.04-TST5 UP.Link/3.4.4`,
`UP.Browser/3.1-NT95 UP.Link/3.2`,
`UP.Browser/3.1-SY11 UP.Link/3.2`,
`UP.Browser/3.1-UPG1 UP.Link/3.2`,
`UP.Browser/3.2.9.1-SA12 UP.Link/3.2`,
`UP.Browser/3.2.9.1-UPG1 UP.Link/3.2`,
}

func TestEZWeb(t *testing.T) {
	for _, userAgent := range ezwebAgents {
		if !mobileagent.IsEZWeb(userAgent) {
			t.Fatal("IsEZWeb Failed:", userAgent)
		}
	}
}

var airhAgents = []string{
`Mozilla/3.0(DDIPOCKET;JRC/AH-J3001V,AH-J3002V/1.0/0100/c50)CNF/2.0`,
`Mozilla/3.0(DDIPOCKET;KYOCERA/AH-K3001V/1.4.1.67.000000/0.1/C100) Opera 7.0`,
`Mozilla/3.0(WILLCOM;KYOCERA/WX300K/1;1.0.2.8.000000/0.1/C100) Opera/7.0`,
`Mozilla/3.0(WILLCOM;SANYO/WX310SA/2;1/1/C128) NetFront/3.3`,
}

func TestAirH(t *testing.T) {
	for _, userAgent := range airhAgents {
		if !mobileagent.IsAirH(userAgent) {
			t.Fatal("IsAirH Failed:", userAgent)
		}
	}
}

var vodaphoneAgents = []string{
`Vodafone/1.0/V902SH/SHJ001/SN999999999999999 Browser/UP.Browser/7.0.2.1 Profile/MIDP-2.0 Configuration/CLDC-1.1 Ext-J-Profile/JSCL-1.2.2 Ext-V-Profile/VSCL-2.0.0,V902SH`,
`MOT-V980/80.2F.2E. MIB/2.2.1 Profile/MIDP-2.0 Configuration/CLDC-1.1,V702MO`,
}

func TestVodaphone(t *testing.T) {
	for _, userAgent := range vodaphoneAgents {
		if !mobileagent.IsVodaphone(userAgent) {
			t.Fatal("IsVodaphone Failed:", userAgent)
		}
	}
}

var softbankAgents = []string{
`SoftBank/1.0/910T/TJ001/SNXXXXXXXXX Browser/NetFront/3.3 Profile/MIDP-2.0 Configuration/CLDC-1.1`,
`SoftBank/1.0/708SC/SCJ001/SNXXXXXXXXX Browser/NetFront/3.3`,
}

func TestSoftBank(t *testing.T) {
	for _, userAgent := range softbankAgents {
		if !mobileagent.IsSoftBank(userAgent) {
			t.Fatal("IsSoftBank Failed:", userAgent)
		}
	}
}
