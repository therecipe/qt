package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
	"strings"
	"unsafe"
)

type QLocale struct {
	ptr unsafe.Pointer
}

type QLocale_ITF interface {
	QLocale_PTR() *QLocale
}

func (p *QLocale) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QLocale) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQLocale(ptr QLocale_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLocale_PTR().Pointer()
	}
	return nil
}

func NewQLocaleFromPointer(ptr unsafe.Pointer) *QLocale {
	var n = new(QLocale)
	n.SetPointer(ptr)
	return n
}

func (ptr *QLocale) QLocale_PTR() *QLocale {
	return ptr
}

//QLocale::Country
type QLocale__Country int64

const (
	QLocale__AnyCountry                             = QLocale__Country(0)
	QLocale__Afghanistan                            = QLocale__Country(1)
	QLocale__Albania                                = QLocale__Country(2)
	QLocale__Algeria                                = QLocale__Country(3)
	QLocale__AmericanSamoa                          = QLocale__Country(4)
	QLocale__Andorra                                = QLocale__Country(5)
	QLocale__Angola                                 = QLocale__Country(6)
	QLocale__Anguilla                               = QLocale__Country(7)
	QLocale__Antarctica                             = QLocale__Country(8)
	QLocale__AntiguaAndBarbuda                      = QLocale__Country(9)
	QLocale__Argentina                              = QLocale__Country(10)
	QLocale__Armenia                                = QLocale__Country(11)
	QLocale__Aruba                                  = QLocale__Country(12)
	QLocale__Australia                              = QLocale__Country(13)
	QLocale__Austria                                = QLocale__Country(14)
	QLocale__Azerbaijan                             = QLocale__Country(15)
	QLocale__Bahamas                                = QLocale__Country(16)
	QLocale__Bahrain                                = QLocale__Country(17)
	QLocale__Bangladesh                             = QLocale__Country(18)
	QLocale__Barbados                               = QLocale__Country(19)
	QLocale__Belarus                                = QLocale__Country(20)
	QLocale__Belgium                                = QLocale__Country(21)
	QLocale__Belize                                 = QLocale__Country(22)
	QLocale__Benin                                  = QLocale__Country(23)
	QLocale__Bermuda                                = QLocale__Country(24)
	QLocale__Bhutan                                 = QLocale__Country(25)
	QLocale__Bolivia                                = QLocale__Country(26)
	QLocale__BosniaAndHerzegowina                   = QLocale__Country(27)
	QLocale__Botswana                               = QLocale__Country(28)
	QLocale__BouvetIsland                           = QLocale__Country(29)
	QLocale__Brazil                                 = QLocale__Country(30)
	QLocale__BritishIndianOceanTerritory            = QLocale__Country(31)
	QLocale__Brunei                                 = QLocale__Country(32)
	QLocale__Bulgaria                               = QLocale__Country(33)
	QLocale__BurkinaFaso                            = QLocale__Country(34)
	QLocale__Burundi                                = QLocale__Country(35)
	QLocale__Cambodia                               = QLocale__Country(36)
	QLocale__Cameroon                               = QLocale__Country(37)
	QLocale__Canada                                 = QLocale__Country(38)
	QLocale__CapeVerde                              = QLocale__Country(39)
	QLocale__CaymanIslands                          = QLocale__Country(40)
	QLocale__CentralAfricanRepublic                 = QLocale__Country(41)
	QLocale__Chad                                   = QLocale__Country(42)
	QLocale__Chile                                  = QLocale__Country(43)
	QLocale__China                                  = QLocale__Country(44)
	QLocale__ChristmasIsland                        = QLocale__Country(45)
	QLocale__CocosIslands                           = QLocale__Country(46)
	QLocale__Colombia                               = QLocale__Country(47)
	QLocale__Comoros                                = QLocale__Country(48)
	QLocale__CongoKinshasa                          = QLocale__Country(49)
	QLocale__CongoBrazzaville                       = QLocale__Country(50)
	QLocale__CookIslands                            = QLocale__Country(51)
	QLocale__CostaRica                              = QLocale__Country(52)
	QLocale__IvoryCoast                             = QLocale__Country(53)
	QLocale__Croatia                                = QLocale__Country(54)
	QLocale__Cuba                                   = QLocale__Country(55)
	QLocale__Cyprus                                 = QLocale__Country(56)
	QLocale__CzechRepublic                          = QLocale__Country(57)
	QLocale__Denmark                                = QLocale__Country(58)
	QLocale__Djibouti                               = QLocale__Country(59)
	QLocale__Dominica                               = QLocale__Country(60)
	QLocale__DominicanRepublic                      = QLocale__Country(61)
	QLocale__EastTimor                              = QLocale__Country(62)
	QLocale__Ecuador                                = QLocale__Country(63)
	QLocale__Egypt                                  = QLocale__Country(64)
	QLocale__ElSalvador                             = QLocale__Country(65)
	QLocale__EquatorialGuinea                       = QLocale__Country(66)
	QLocale__Eritrea                                = QLocale__Country(67)
	QLocale__Estonia                                = QLocale__Country(68)
	QLocale__Ethiopia                               = QLocale__Country(69)
	QLocale__FalklandIslands                        = QLocale__Country(70)
	QLocale__FaroeIslands                           = QLocale__Country(71)
	QLocale__Fiji                                   = QLocale__Country(72)
	QLocale__Finland                                = QLocale__Country(73)
	QLocale__France                                 = QLocale__Country(74)
	QLocale__Guernsey                               = QLocale__Country(75)
	QLocale__FrenchGuiana                           = QLocale__Country(76)
	QLocale__FrenchPolynesia                        = QLocale__Country(77)
	QLocale__FrenchSouthernTerritories              = QLocale__Country(78)
	QLocale__Gabon                                  = QLocale__Country(79)
	QLocale__Gambia                                 = QLocale__Country(80)
	QLocale__Georgia                                = QLocale__Country(81)
	QLocale__Germany                                = QLocale__Country(82)
	QLocale__Ghana                                  = QLocale__Country(83)
	QLocale__Gibraltar                              = QLocale__Country(84)
	QLocale__Greece                                 = QLocale__Country(85)
	QLocale__Greenland                              = QLocale__Country(86)
	QLocale__Grenada                                = QLocale__Country(87)
	QLocale__Guadeloupe                             = QLocale__Country(88)
	QLocale__Guam                                   = QLocale__Country(89)
	QLocale__Guatemala                              = QLocale__Country(90)
	QLocale__Guinea                                 = QLocale__Country(91)
	QLocale__GuineaBissau                           = QLocale__Country(92)
	QLocale__Guyana                                 = QLocale__Country(93)
	QLocale__Haiti                                  = QLocale__Country(94)
	QLocale__HeardAndMcDonaldIslands                = QLocale__Country(95)
	QLocale__Honduras                               = QLocale__Country(96)
	QLocale__HongKong                               = QLocale__Country(97)
	QLocale__Hungary                                = QLocale__Country(98)
	QLocale__Iceland                                = QLocale__Country(99)
	QLocale__India                                  = QLocale__Country(100)
	QLocale__Indonesia                              = QLocale__Country(101)
	QLocale__Iran                                   = QLocale__Country(102)
	QLocale__Iraq                                   = QLocale__Country(103)
	QLocale__Ireland                                = QLocale__Country(104)
	QLocale__Israel                                 = QLocale__Country(105)
	QLocale__Italy                                  = QLocale__Country(106)
	QLocale__Jamaica                                = QLocale__Country(107)
	QLocale__Japan                                  = QLocale__Country(108)
	QLocale__Jordan                                 = QLocale__Country(109)
	QLocale__Kazakhstan                             = QLocale__Country(110)
	QLocale__Kenya                                  = QLocale__Country(111)
	QLocale__Kiribati                               = QLocale__Country(112)
	QLocale__NorthKorea                             = QLocale__Country(113)
	QLocale__SouthKorea                             = QLocale__Country(114)
	QLocale__Kuwait                                 = QLocale__Country(115)
	QLocale__Kyrgyzstan                             = QLocale__Country(116)
	QLocale__Laos                                   = QLocale__Country(117)
	QLocale__Latvia                                 = QLocale__Country(118)
	QLocale__Lebanon                                = QLocale__Country(119)
	QLocale__Lesotho                                = QLocale__Country(120)
	QLocale__Liberia                                = QLocale__Country(121)
	QLocale__Libya                                  = QLocale__Country(122)
	QLocale__Liechtenstein                          = QLocale__Country(123)
	QLocale__Lithuania                              = QLocale__Country(124)
	QLocale__Luxembourg                             = QLocale__Country(125)
	QLocale__Macau                                  = QLocale__Country(126)
	QLocale__Macedonia                              = QLocale__Country(127)
	QLocale__Madagascar                             = QLocale__Country(128)
	QLocale__Malawi                                 = QLocale__Country(129)
	QLocale__Malaysia                               = QLocale__Country(130)
	QLocale__Maldives                               = QLocale__Country(131)
	QLocale__Mali                                   = QLocale__Country(132)
	QLocale__Malta                                  = QLocale__Country(133)
	QLocale__MarshallIslands                        = QLocale__Country(134)
	QLocale__Martinique                             = QLocale__Country(135)
	QLocale__Mauritania                             = QLocale__Country(136)
	QLocale__Mauritius                              = QLocale__Country(137)
	QLocale__Mayotte                                = QLocale__Country(138)
	QLocale__Mexico                                 = QLocale__Country(139)
	QLocale__Micronesia                             = QLocale__Country(140)
	QLocale__Moldova                                = QLocale__Country(141)
	QLocale__Monaco                                 = QLocale__Country(142)
	QLocale__Mongolia                               = QLocale__Country(143)
	QLocale__Montserrat                             = QLocale__Country(144)
	QLocale__Morocco                                = QLocale__Country(145)
	QLocale__Mozambique                             = QLocale__Country(146)
	QLocale__Myanmar                                = QLocale__Country(147)
	QLocale__Namibia                                = QLocale__Country(148)
	QLocale__NauruCountry                           = QLocale__Country(149)
	QLocale__Nepal                                  = QLocale__Country(150)
	QLocale__Netherlands                            = QLocale__Country(151)
	QLocale__CuraSao                                = QLocale__Country(152)
	QLocale__NewCaledonia                           = QLocale__Country(153)
	QLocale__NewZealand                             = QLocale__Country(154)
	QLocale__Nicaragua                              = QLocale__Country(155)
	QLocale__Niger                                  = QLocale__Country(156)
	QLocale__Nigeria                                = QLocale__Country(157)
	QLocale__Niue                                   = QLocale__Country(158)
	QLocale__NorfolkIsland                          = QLocale__Country(159)
	QLocale__NorthernMarianaIslands                 = QLocale__Country(160)
	QLocale__Norway                                 = QLocale__Country(161)
	QLocale__Oman                                   = QLocale__Country(162)
	QLocale__Pakistan                               = QLocale__Country(163)
	QLocale__Palau                                  = QLocale__Country(164)
	QLocale__PalestinianTerritories                 = QLocale__Country(165)
	QLocale__Panama                                 = QLocale__Country(166)
	QLocale__PapuaNewGuinea                         = QLocale__Country(167)
	QLocale__Paraguay                               = QLocale__Country(168)
	QLocale__Peru                                   = QLocale__Country(169)
	QLocale__Philippines                            = QLocale__Country(170)
	QLocale__Pitcairn                               = QLocale__Country(171)
	QLocale__Poland                                 = QLocale__Country(172)
	QLocale__Portugal                               = QLocale__Country(173)
	QLocale__PuertoRico                             = QLocale__Country(174)
	QLocale__Qatar                                  = QLocale__Country(175)
	QLocale__Reunion                                = QLocale__Country(176)
	QLocale__Romania                                = QLocale__Country(177)
	QLocale__Russia                                 = QLocale__Country(178)
	QLocale__Rwanda                                 = QLocale__Country(179)
	QLocale__SaintKittsAndNevis                     = QLocale__Country(180)
	QLocale__SaintLucia                             = QLocale__Country(181)
	QLocale__SaintVincentAndTheGrenadines           = QLocale__Country(182)
	QLocale__Samoa                                  = QLocale__Country(183)
	QLocale__SanMarino                              = QLocale__Country(184)
	QLocale__SaoTomeAndPrincipe                     = QLocale__Country(185)
	QLocale__SaudiArabia                            = QLocale__Country(186)
	QLocale__Senegal                                = QLocale__Country(187)
	QLocale__Seychelles                             = QLocale__Country(188)
	QLocale__SierraLeone                            = QLocale__Country(189)
	QLocale__Singapore                              = QLocale__Country(190)
	QLocale__Slovakia                               = QLocale__Country(191)
	QLocale__Slovenia                               = QLocale__Country(192)
	QLocale__SolomonIslands                         = QLocale__Country(193)
	QLocale__Somalia                                = QLocale__Country(194)
	QLocale__SouthAfrica                            = QLocale__Country(195)
	QLocale__SouthGeorgiaAndTheSouthSandwichIslands = QLocale__Country(196)
	QLocale__Spain                                  = QLocale__Country(197)
	QLocale__SriLanka                               = QLocale__Country(198)
	QLocale__SaintHelena                            = QLocale__Country(199)
	QLocale__SaintPierreAndMiquelon                 = QLocale__Country(200)
	QLocale__Sudan                                  = QLocale__Country(201)
	QLocale__Suriname                               = QLocale__Country(202)
	QLocale__SvalbardAndJanMayenIslands             = QLocale__Country(203)
	QLocale__Swaziland                              = QLocale__Country(204)
	QLocale__Sweden                                 = QLocale__Country(205)
	QLocale__Switzerland                            = QLocale__Country(206)
	QLocale__Syria                                  = QLocale__Country(207)
	QLocale__Taiwan                                 = QLocale__Country(208)
	QLocale__Tajikistan                             = QLocale__Country(209)
	QLocale__Tanzania                               = QLocale__Country(210)
	QLocale__Thailand                               = QLocale__Country(211)
	QLocale__Togo                                   = QLocale__Country(212)
	QLocale__Tokelau                                = QLocale__Country(213)
	QLocale__Tonga                                  = QLocale__Country(214)
	QLocale__TrinidadAndTobago                      = QLocale__Country(215)
	QLocale__Tunisia                                = QLocale__Country(216)
	QLocale__Turkey                                 = QLocale__Country(217)
	QLocale__Turkmenistan                           = QLocale__Country(218)
	QLocale__TurksAndCaicosIslands                  = QLocale__Country(219)
	QLocale__Tuvalu                                 = QLocale__Country(220)
	QLocale__Uganda                                 = QLocale__Country(221)
	QLocale__Ukraine                                = QLocale__Country(222)
	QLocale__UnitedArabEmirates                     = QLocale__Country(223)
	QLocale__UnitedKingdom                          = QLocale__Country(224)
	QLocale__UnitedStates                           = QLocale__Country(225)
	QLocale__UnitedStatesMinorOutlyingIslands       = QLocale__Country(226)
	QLocale__Uruguay                                = QLocale__Country(227)
	QLocale__Uzbekistan                             = QLocale__Country(228)
	QLocale__Vanuatu                                = QLocale__Country(229)
	QLocale__VaticanCityState                       = QLocale__Country(230)
	QLocale__Venezuela                              = QLocale__Country(231)
	QLocale__Vietnam                                = QLocale__Country(232)
	QLocale__BritishVirginIslands                   = QLocale__Country(233)
	QLocale__UnitedStatesVirginIslands              = QLocale__Country(234)
	QLocale__WallisAndFutunaIslands                 = QLocale__Country(235)
	QLocale__WesternSahara                          = QLocale__Country(236)
	QLocale__Yemen                                  = QLocale__Country(237)
	QLocale__CanaryIslands                          = QLocale__Country(238)
	QLocale__Zambia                                 = QLocale__Country(239)
	QLocale__Zimbabwe                               = QLocale__Country(240)
	QLocale__ClippertonIsland                       = QLocale__Country(241)
	QLocale__Montenegro                             = QLocale__Country(242)
	QLocale__Serbia                                 = QLocale__Country(243)
	QLocale__SaintBarthelemy                        = QLocale__Country(244)
	QLocale__SaintMartin                            = QLocale__Country(245)
	QLocale__LatinAmericaAndTheCaribbean            = QLocale__Country(246)
	QLocale__AscensionIsland                        = QLocale__Country(247)
	QLocale__AlandIslands                           = QLocale__Country(248)
	QLocale__DiegoGarcia                            = QLocale__Country(249)
	QLocale__CeutaAndMelilla                        = QLocale__Country(250)
	QLocale__IsleOfMan                              = QLocale__Country(251)
	QLocale__Jersey                                 = QLocale__Country(252)
	QLocale__TristanDaCunha                         = QLocale__Country(253)
	QLocale__SouthSudan                             = QLocale__Country(254)
	QLocale__Bonaire                                = QLocale__Country(255)
	QLocale__SintMaarten                            = QLocale__Country(256)
	QLocale__Kosovo                                 = QLocale__Country(257)
	QLocale__DemocraticRepublicOfCongo              = QLocale__Country(QLocale__CongoKinshasa)
	QLocale__PeoplesRepublicOfCongo                 = QLocale__Country(QLocale__CongoBrazzaville)
	QLocale__DemocraticRepublicOfKorea              = QLocale__Country(QLocale__NorthKorea)
	QLocale__RepublicOfKorea                        = QLocale__Country(QLocale__SouthKorea)
	QLocale__RussianFederation                      = QLocale__Country(QLocale__Russia)
	QLocale__SyrianArabRepublic                     = QLocale__Country(QLocale__Syria)
	QLocale__LastCountry                            = QLocale__Country(QLocale__Kosovo)
)

//QLocale::CurrencySymbolFormat
type QLocale__CurrencySymbolFormat int64

const (
	QLocale__CurrencyIsoCode     = QLocale__CurrencySymbolFormat(0)
	QLocale__CurrencySymbol      = QLocale__CurrencySymbolFormat(1)
	QLocale__CurrencyDisplayName = QLocale__CurrencySymbolFormat(2)
)

//QLocale::FormatType
type QLocale__FormatType int64

const (
	QLocale__LongFormat   = QLocale__FormatType(0)
	QLocale__ShortFormat  = QLocale__FormatType(1)
	QLocale__NarrowFormat = QLocale__FormatType(2)
)

//QLocale::Language
type QLocale__Language int64

const (
	QLocale__AnyLanguage               = QLocale__Language(0)
	QLocale__C                         = QLocale__Language(1)
	QLocale__Abkhazian                 = QLocale__Language(2)
	QLocale__Oromo                     = QLocale__Language(3)
	QLocale__Afar                      = QLocale__Language(4)
	QLocale__Afrikaans                 = QLocale__Language(5)
	QLocale__Albanian                  = QLocale__Language(6)
	QLocale__Amharic                   = QLocale__Language(7)
	QLocale__Arabic                    = QLocale__Language(8)
	QLocale__Armenian                  = QLocale__Language(9)
	QLocale__Assamese                  = QLocale__Language(10)
	QLocale__Aymara                    = QLocale__Language(11)
	QLocale__Azerbaijani               = QLocale__Language(12)
	QLocale__Bashkir                   = QLocale__Language(13)
	QLocale__Basque                    = QLocale__Language(14)
	QLocale__Bengali                   = QLocale__Language(15)
	QLocale__Dzongkha                  = QLocale__Language(16)
	QLocale__Bihari                    = QLocale__Language(17)
	QLocale__Bislama                   = QLocale__Language(18)
	QLocale__Breton                    = QLocale__Language(19)
	QLocale__Bulgarian                 = QLocale__Language(20)
	QLocale__Burmese                   = QLocale__Language(21)
	QLocale__Belarusian                = QLocale__Language(22)
	QLocale__Khmer                     = QLocale__Language(23)
	QLocale__Catalan                   = QLocale__Language(24)
	QLocale__Chinese                   = QLocale__Language(25)
	QLocale__Corsican                  = QLocale__Language(26)
	QLocale__Croatian                  = QLocale__Language(27)
	QLocale__Czech                     = QLocale__Language(28)
	QLocale__Danish                    = QLocale__Language(29)
	QLocale__Dutch                     = QLocale__Language(30)
	QLocale__English                   = QLocale__Language(31)
	QLocale__Esperanto                 = QLocale__Language(32)
	QLocale__Estonian                  = QLocale__Language(33)
	QLocale__Faroese                   = QLocale__Language(34)
	QLocale__Fijian                    = QLocale__Language(35)
	QLocale__Finnish                   = QLocale__Language(36)
	QLocale__French                    = QLocale__Language(37)
	QLocale__WesternFrisian            = QLocale__Language(38)
	QLocale__Gaelic                    = QLocale__Language(39)
	QLocale__Galician                  = QLocale__Language(40)
	QLocale__Georgian                  = QLocale__Language(41)
	QLocale__German                    = QLocale__Language(42)
	QLocale__Greek                     = QLocale__Language(43)
	QLocale__Greenlandic               = QLocale__Language(44)
	QLocale__Guarani                   = QLocale__Language(45)
	QLocale__Gujarati                  = QLocale__Language(46)
	QLocale__Hausa                     = QLocale__Language(47)
	QLocale__Hebrew                    = QLocale__Language(48)
	QLocale__Hindi                     = QLocale__Language(49)
	QLocale__Hungarian                 = QLocale__Language(50)
	QLocale__Icelandic                 = QLocale__Language(51)
	QLocale__Indonesian                = QLocale__Language(52)
	QLocale__Interlingua               = QLocale__Language(53)
	QLocale__Interlingue               = QLocale__Language(54)
	QLocale__Inuktitut                 = QLocale__Language(55)
	QLocale__Inupiak                   = QLocale__Language(56)
	QLocale__Irish                     = QLocale__Language(57)
	QLocale__Italian                   = QLocale__Language(58)
	QLocale__Japanese                  = QLocale__Language(59)
	QLocale__Javanese                  = QLocale__Language(60)
	QLocale__Kannada                   = QLocale__Language(61)
	QLocale__Kashmiri                  = QLocale__Language(62)
	QLocale__Kazakh                    = QLocale__Language(63)
	QLocale__Kinyarwanda               = QLocale__Language(64)
	QLocale__Kirghiz                   = QLocale__Language(65)
	QLocale__Korean                    = QLocale__Language(66)
	QLocale__Kurdish                   = QLocale__Language(67)
	QLocale__Rundi                     = QLocale__Language(68)
	QLocale__Lao                       = QLocale__Language(69)
	QLocale__Latin                     = QLocale__Language(70)
	QLocale__Latvian                   = QLocale__Language(71)
	QLocale__Lingala                   = QLocale__Language(72)
	QLocale__Lithuanian                = QLocale__Language(73)
	QLocale__Macedonian                = QLocale__Language(74)
	QLocale__Malagasy                  = QLocale__Language(75)
	QLocale__Malay                     = QLocale__Language(76)
	QLocale__Malayalam                 = QLocale__Language(77)
	QLocale__Maltese                   = QLocale__Language(78)
	QLocale__Maori                     = QLocale__Language(79)
	QLocale__Marathi                   = QLocale__Language(80)
	QLocale__Marshallese               = QLocale__Language(81)
	QLocale__Mongolian                 = QLocale__Language(82)
	QLocale__NauruLanguage             = QLocale__Language(83)
	QLocale__Nepali                    = QLocale__Language(84)
	QLocale__NorwegianBokmal           = QLocale__Language(85)
	QLocale__Occitan                   = QLocale__Language(86)
	QLocale__Oriya                     = QLocale__Language(87)
	QLocale__Pashto                    = QLocale__Language(88)
	QLocale__Persian                   = QLocale__Language(89)
	QLocale__Polish                    = QLocale__Language(90)
	QLocale__Portuguese                = QLocale__Language(91)
	QLocale__Punjabi                   = QLocale__Language(92)
	QLocale__Quechua                   = QLocale__Language(93)
	QLocale__Romansh                   = QLocale__Language(94)
	QLocale__Romanian                  = QLocale__Language(95)
	QLocale__Russian                   = QLocale__Language(96)
	QLocale__Samoan                    = QLocale__Language(97)
	QLocale__Sango                     = QLocale__Language(98)
	QLocale__Sanskrit                  = QLocale__Language(99)
	QLocale__Serbian                   = QLocale__Language(100)
	QLocale__Ossetic                   = QLocale__Language(101)
	QLocale__SouthernSotho             = QLocale__Language(102)
	QLocale__Tswana                    = QLocale__Language(103)
	QLocale__Shona                     = QLocale__Language(104)
	QLocale__Sindhi                    = QLocale__Language(105)
	QLocale__Sinhala                   = QLocale__Language(106)
	QLocale__Swati                     = QLocale__Language(107)
	QLocale__Slovak                    = QLocale__Language(108)
	QLocale__Slovenian                 = QLocale__Language(109)
	QLocale__Somali                    = QLocale__Language(110)
	QLocale__Spanish                   = QLocale__Language(111)
	QLocale__Sundanese                 = QLocale__Language(112)
	QLocale__Swahili                   = QLocale__Language(113)
	QLocale__Swedish                   = QLocale__Language(114)
	QLocale__Sardinian                 = QLocale__Language(115)
	QLocale__Tajik                     = QLocale__Language(116)
	QLocale__Tamil                     = QLocale__Language(117)
	QLocale__Tatar                     = QLocale__Language(118)
	QLocale__Telugu                    = QLocale__Language(119)
	QLocale__Thai                      = QLocale__Language(120)
	QLocale__Tibetan                   = QLocale__Language(121)
	QLocale__Tigrinya                  = QLocale__Language(122)
	QLocale__Tongan                    = QLocale__Language(123)
	QLocale__Tsonga                    = QLocale__Language(124)
	QLocale__Turkish                   = QLocale__Language(125)
	QLocale__Turkmen                   = QLocale__Language(126)
	QLocale__Tahitian                  = QLocale__Language(127)
	QLocale__Uighur                    = QLocale__Language(128)
	QLocale__Ukrainian                 = QLocale__Language(129)
	QLocale__Urdu                      = QLocale__Language(130)
	QLocale__Uzbek                     = QLocale__Language(131)
	QLocale__Vietnamese                = QLocale__Language(132)
	QLocale__Volapuk                   = QLocale__Language(133)
	QLocale__Welsh                     = QLocale__Language(134)
	QLocale__Wolof                     = QLocale__Language(135)
	QLocale__Xhosa                     = QLocale__Language(136)
	QLocale__Yiddish                   = QLocale__Language(137)
	QLocale__Yoruba                    = QLocale__Language(138)
	QLocale__Zhuang                    = QLocale__Language(139)
	QLocale__Zulu                      = QLocale__Language(140)
	QLocale__NorwegianNynorsk          = QLocale__Language(141)
	QLocale__Bosnian                   = QLocale__Language(142)
	QLocale__Divehi                    = QLocale__Language(143)
	QLocale__Manx                      = QLocale__Language(144)
	QLocale__Cornish                   = QLocale__Language(145)
	QLocale__Akan                      = QLocale__Language(146)
	QLocale__Konkani                   = QLocale__Language(147)
	QLocale__Ga                        = QLocale__Language(148)
	QLocale__Igbo                      = QLocale__Language(149)
	QLocale__Kamba                     = QLocale__Language(150)
	QLocale__Syriac                    = QLocale__Language(151)
	QLocale__Blin                      = QLocale__Language(152)
	QLocale__Geez                      = QLocale__Language(153)
	QLocale__Koro                      = QLocale__Language(154)
	QLocale__Sidamo                    = QLocale__Language(155)
	QLocale__Atsam                     = QLocale__Language(156)
	QLocale__Tigre                     = QLocale__Language(157)
	QLocale__Jju                       = QLocale__Language(158)
	QLocale__Friulian                  = QLocale__Language(159)
	QLocale__Venda                     = QLocale__Language(160)
	QLocale__Ewe                       = QLocale__Language(161)
	QLocale__Walamo                    = QLocale__Language(162)
	QLocale__Hawaiian                  = QLocale__Language(163)
	QLocale__Tyap                      = QLocale__Language(164)
	QLocale__Nyanja                    = QLocale__Language(165)
	QLocale__Filipino                  = QLocale__Language(166)
	QLocale__SwissGerman               = QLocale__Language(167)
	QLocale__SichuanYi                 = QLocale__Language(168)
	QLocale__Kpelle                    = QLocale__Language(169)
	QLocale__LowGerman                 = QLocale__Language(170)
	QLocale__SouthNdebele              = QLocale__Language(171)
	QLocale__NorthernSotho             = QLocale__Language(172)
	QLocale__NorthernSami              = QLocale__Language(173)
	QLocale__Taroko                    = QLocale__Language(174)
	QLocale__Gusii                     = QLocale__Language(175)
	QLocale__Taita                     = QLocale__Language(176)
	QLocale__Fulah                     = QLocale__Language(177)
	QLocale__Kikuyu                    = QLocale__Language(178)
	QLocale__Samburu                   = QLocale__Language(179)
	QLocale__Sena                      = QLocale__Language(180)
	QLocale__NorthNdebele              = QLocale__Language(181)
	QLocale__Rombo                     = QLocale__Language(182)
	QLocale__Tachelhit                 = QLocale__Language(183)
	QLocale__Kabyle                    = QLocale__Language(184)
	QLocale__Nyankole                  = QLocale__Language(185)
	QLocale__Bena                      = QLocale__Language(186)
	QLocale__Vunjo                     = QLocale__Language(187)
	QLocale__Bambara                   = QLocale__Language(188)
	QLocale__Embu                      = QLocale__Language(189)
	QLocale__Cherokee                  = QLocale__Language(190)
	QLocale__Morisyen                  = QLocale__Language(191)
	QLocale__Makonde                   = QLocale__Language(192)
	QLocale__Langi                     = QLocale__Language(193)
	QLocale__Ganda                     = QLocale__Language(194)
	QLocale__Bemba                     = QLocale__Language(195)
	QLocale__Kabuverdianu              = QLocale__Language(196)
	QLocale__Meru                      = QLocale__Language(197)
	QLocale__Kalenjin                  = QLocale__Language(198)
	QLocale__Nama                      = QLocale__Language(199)
	QLocale__Machame                   = QLocale__Language(200)
	QLocale__Colognian                 = QLocale__Language(201)
	QLocale__Masai                     = QLocale__Language(202)
	QLocale__Soga                      = QLocale__Language(203)
	QLocale__Luyia                     = QLocale__Language(204)
	QLocale__Asu                       = QLocale__Language(205)
	QLocale__Teso                      = QLocale__Language(206)
	QLocale__Saho                      = QLocale__Language(207)
	QLocale__KoyraChiini               = QLocale__Language(208)
	QLocale__Rwa                       = QLocale__Language(209)
	QLocale__Luo                       = QLocale__Language(210)
	QLocale__Chiga                     = QLocale__Language(211)
	QLocale__CentralMoroccoTamazight   = QLocale__Language(212)
	QLocale__KoyraboroSenni            = QLocale__Language(213)
	QLocale__Shambala                  = QLocale__Language(214)
	QLocale__Bodo                      = QLocale__Language(215)
	QLocale__Avaric                    = QLocale__Language(216)
	QLocale__Chamorro                  = QLocale__Language(217)
	QLocale__Chechen                   = QLocale__Language(218)
	QLocale__Church                    = QLocale__Language(219)
	QLocale__Chuvash                   = QLocale__Language(220)
	QLocale__Cree                      = QLocale__Language(221)
	QLocale__Haitian                   = QLocale__Language(222)
	QLocale__Herero                    = QLocale__Language(223)
	QLocale__HiriMotu                  = QLocale__Language(224)
	QLocale__Kanuri                    = QLocale__Language(225)
	QLocale__Komi                      = QLocale__Language(226)
	QLocale__Kongo                     = QLocale__Language(227)
	QLocale__Kwanyama                  = QLocale__Language(228)
	QLocale__Limburgish                = QLocale__Language(229)
	QLocale__LubaKatanga               = QLocale__Language(230)
	QLocale__Luxembourgish             = QLocale__Language(231)
	QLocale__Navaho                    = QLocale__Language(232)
	QLocale__Ndonga                    = QLocale__Language(233)
	QLocale__Ojibwa                    = QLocale__Language(234)
	QLocale__Pali                      = QLocale__Language(235)
	QLocale__Walloon                   = QLocale__Language(236)
	QLocale__Aghem                     = QLocale__Language(237)
	QLocale__Basaa                     = QLocale__Language(238)
	QLocale__Zarma                     = QLocale__Language(239)
	QLocale__Duala                     = QLocale__Language(240)
	QLocale__JolaFonyi                 = QLocale__Language(241)
	QLocale__Ewondo                    = QLocale__Language(242)
	QLocale__Bafia                     = QLocale__Language(243)
	QLocale__MakhuwaMeetto             = QLocale__Language(244)
	QLocale__Mundang                   = QLocale__Language(245)
	QLocale__Kwasio                    = QLocale__Language(246)
	QLocale__Nuer                      = QLocale__Language(247)
	QLocale__Sakha                     = QLocale__Language(248)
	QLocale__Sangu                     = QLocale__Language(249)
	QLocale__CongoSwahili              = QLocale__Language(250)
	QLocale__Tasawaq                   = QLocale__Language(251)
	QLocale__Vai                       = QLocale__Language(252)
	QLocale__Walser                    = QLocale__Language(253)
	QLocale__Yangben                   = QLocale__Language(254)
	QLocale__Avestan                   = QLocale__Language(255)
	QLocale__Asturian                  = QLocale__Language(256)
	QLocale__Ngomba                    = QLocale__Language(257)
	QLocale__Kako                      = QLocale__Language(258)
	QLocale__Meta                      = QLocale__Language(259)
	QLocale__Ngiemboon                 = QLocale__Language(260)
	QLocale__Aragonese                 = QLocale__Language(261)
	QLocale__Akkadian                  = QLocale__Language(262)
	QLocale__AncientEgyptian           = QLocale__Language(263)
	QLocale__AncientGreek              = QLocale__Language(264)
	QLocale__Aramaic                   = QLocale__Language(265)
	QLocale__Balinese                  = QLocale__Language(266)
	QLocale__Bamun                     = QLocale__Language(267)
	QLocale__BatakToba                 = QLocale__Language(268)
	QLocale__Buginese                  = QLocale__Language(269)
	QLocale__Buhid                     = QLocale__Language(270)
	QLocale__Carian                    = QLocale__Language(271)
	QLocale__Chakma                    = QLocale__Language(272)
	QLocale__ClassicalMandaic          = QLocale__Language(273)
	QLocale__Coptic                    = QLocale__Language(274)
	QLocale__Dogri                     = QLocale__Language(275)
	QLocale__EasternCham               = QLocale__Language(276)
	QLocale__EasternKayah              = QLocale__Language(277)
	QLocale__Etruscan                  = QLocale__Language(278)
	QLocale__Gothic                    = QLocale__Language(279)
	QLocale__Hanunoo                   = QLocale__Language(280)
	QLocale__Ingush                    = QLocale__Language(281)
	QLocale__LargeFloweryMiao          = QLocale__Language(282)
	QLocale__Lepcha                    = QLocale__Language(283)
	QLocale__Limbu                     = QLocale__Language(284)
	QLocale__Lisu                      = QLocale__Language(285)
	QLocale__Lu                        = QLocale__Language(286)
	QLocale__Lycian                    = QLocale__Language(287)
	QLocale__Lydian                    = QLocale__Language(288)
	QLocale__Mandingo                  = QLocale__Language(289)
	QLocale__Manipuri                  = QLocale__Language(290)
	QLocale__Meroitic                  = QLocale__Language(291)
	QLocale__NorthernThai              = QLocale__Language(292)
	QLocale__OldIrish                  = QLocale__Language(293)
	QLocale__OldNorse                  = QLocale__Language(294)
	QLocale__OldPersian                = QLocale__Language(295)
	QLocale__OldTurkish                = QLocale__Language(296)
	QLocale__Pahlavi                   = QLocale__Language(297)
	QLocale__Parthian                  = QLocale__Language(298)
	QLocale__Phoenician                = QLocale__Language(299)
	QLocale__PrakritLanguage           = QLocale__Language(300)
	QLocale__Rejang                    = QLocale__Language(301)
	QLocale__Sabaean                   = QLocale__Language(302)
	QLocale__Samaritan                 = QLocale__Language(303)
	QLocale__Santali                   = QLocale__Language(304)
	QLocale__Saurashtra                = QLocale__Language(305)
	QLocale__Sora                      = QLocale__Language(306)
	QLocale__Sylheti                   = QLocale__Language(307)
	QLocale__Tagbanwa                  = QLocale__Language(308)
	QLocale__TaiDam                    = QLocale__Language(309)
	QLocale__TaiNua                    = QLocale__Language(310)
	QLocale__Ugaritic                  = QLocale__Language(311)
	QLocale__Akoose                    = QLocale__Language(312)
	QLocale__Lakota                    = QLocale__Language(313)
	QLocale__StandardMoroccanTamazight = QLocale__Language(314)
	QLocale__Mapuche                   = QLocale__Language(315)
	QLocale__CentralKurdish            = QLocale__Language(316)
	QLocale__LowerSorbian              = QLocale__Language(317)
	QLocale__UpperSorbian              = QLocale__Language(318)
	QLocale__Kenyang                   = QLocale__Language(319)
	QLocale__Mohawk                    = QLocale__Language(320)
	QLocale__Nko                       = QLocale__Language(321)
	QLocale__Prussian                  = QLocale__Language(322)
	QLocale__Kiche                     = QLocale__Language(323)
	QLocale__SouthernSami              = QLocale__Language(324)
	QLocale__LuleSami                  = QLocale__Language(325)
	QLocale__InariSami                 = QLocale__Language(326)
	QLocale__SkoltSami                 = QLocale__Language(327)
	QLocale__Warlpiri                  = QLocale__Language(328)
	QLocale__ManichaeanMiddlePersian   = QLocale__Language(329)
	QLocale__Mende                     = QLocale__Language(330)
	QLocale__AncientNorthArabian       = QLocale__Language(331)
	QLocale__LinearA                   = QLocale__Language(332)
	QLocale__HmongNjua                 = QLocale__Language(333)
	QLocale__Ho                        = QLocale__Language(334)
	QLocale__Lezghian                  = QLocale__Language(335)
	QLocale__Bassa                     = QLocale__Language(336)
	QLocale__Mono                      = QLocale__Language(337)
	QLocale__TedimChin                 = QLocale__Language(338)
	QLocale__Maithili                  = QLocale__Language(339)
	QLocale__Norwegian                 = QLocale__Language(QLocale__NorwegianBokmal)
	QLocale__Moldavian                 = QLocale__Language(QLocale__Romanian)
	QLocale__SerboCroatian             = QLocale__Language(QLocale__Serbian)
	QLocale__Tagalog                   = QLocale__Language(QLocale__Filipino)
	QLocale__Twi                       = QLocale__Language(QLocale__Akan)
	QLocale__Afan                      = QLocale__Language(QLocale__Oromo)
	QLocale__Byelorussian              = QLocale__Language(QLocale__Belarusian)
	QLocale__Bhutani                   = QLocale__Language(QLocale__Dzongkha)
	QLocale__Cambodian                 = QLocale__Language(QLocale__Khmer)
	QLocale__Kurundi                   = QLocale__Language(QLocale__Rundi)
	QLocale__RhaetoRomance             = QLocale__Language(QLocale__Romansh)
	QLocale__Chewa                     = QLocale__Language(QLocale__Nyanja)
	QLocale__Frisian                   = QLocale__Language(QLocale__WesternFrisian)
	QLocale__Uigur                     = QLocale__Language(QLocale__Uighur)
	QLocale__LastLanguage              = QLocale__Language(QLocale__Maithili)
)

//QLocale::MeasurementSystem
type QLocale__MeasurementSystem int64

const (
	QLocale__MetricSystem     = QLocale__MeasurementSystem(0)
	QLocale__ImperialUSSystem = QLocale__MeasurementSystem(1)
	QLocale__ImperialUKSystem = QLocale__MeasurementSystem(2)
	QLocale__ImperialSystem   = QLocale__MeasurementSystem(QLocale__ImperialUSSystem)
)

//QLocale::NumberOption
type QLocale__NumberOption int64

const (
	QLocale__OmitGroupSeparator   = QLocale__NumberOption(0x01)
	QLocale__RejectGroupSeparator = QLocale__NumberOption(0x02)
)

//QLocale::QuotationStyle
type QLocale__QuotationStyle int64

var (
	QLocale__StandardQuotation  = QLocale__QuotationStyle(0)
	QLocale__AlternateQuotation = QLocale__QuotationStyle(1)
)

//QLocale::Script
type QLocale__Script int64

const (
	QLocale__AnyScript                   = QLocale__Script(0)
	QLocale__ArabicScript                = QLocale__Script(1)
	QLocale__CyrillicScript              = QLocale__Script(2)
	QLocale__DeseretScript               = QLocale__Script(3)
	QLocale__GurmukhiScript              = QLocale__Script(4)
	QLocale__SimplifiedHanScript         = QLocale__Script(5)
	QLocale__TraditionalHanScript        = QLocale__Script(6)
	QLocale__LatinScript                 = QLocale__Script(7)
	QLocale__MongolianScript             = QLocale__Script(8)
	QLocale__TifinaghScript              = QLocale__Script(9)
	QLocale__ArmenianScript              = QLocale__Script(10)
	QLocale__BengaliScript               = QLocale__Script(11)
	QLocale__CherokeeScript              = QLocale__Script(12)
	QLocale__DevanagariScript            = QLocale__Script(13)
	QLocale__EthiopicScript              = QLocale__Script(14)
	QLocale__GeorgianScript              = QLocale__Script(15)
	QLocale__GreekScript                 = QLocale__Script(16)
	QLocale__GujaratiScript              = QLocale__Script(17)
	QLocale__HebrewScript                = QLocale__Script(18)
	QLocale__JapaneseScript              = QLocale__Script(19)
	QLocale__KhmerScript                 = QLocale__Script(20)
	QLocale__KannadaScript               = QLocale__Script(21)
	QLocale__KoreanScript                = QLocale__Script(22)
	QLocale__LaoScript                   = QLocale__Script(23)
	QLocale__MalayalamScript             = QLocale__Script(24)
	QLocale__MyanmarScript               = QLocale__Script(25)
	QLocale__OriyaScript                 = QLocale__Script(26)
	QLocale__TamilScript                 = QLocale__Script(27)
	QLocale__TeluguScript                = QLocale__Script(28)
	QLocale__ThaanaScript                = QLocale__Script(29)
	QLocale__ThaiScript                  = QLocale__Script(30)
	QLocale__TibetanScript               = QLocale__Script(31)
	QLocale__SinhalaScript               = QLocale__Script(32)
	QLocale__SyriacScript                = QLocale__Script(33)
	QLocale__YiScript                    = QLocale__Script(34)
	QLocale__VaiScript                   = QLocale__Script(35)
	QLocale__AvestanScript               = QLocale__Script(36)
	QLocale__BalineseScript              = QLocale__Script(37)
	QLocale__BamumScript                 = QLocale__Script(38)
	QLocale__BatakScript                 = QLocale__Script(39)
	QLocale__BopomofoScript              = QLocale__Script(40)
	QLocale__BrahmiScript                = QLocale__Script(41)
	QLocale__BugineseScript              = QLocale__Script(42)
	QLocale__BuhidScript                 = QLocale__Script(43)
	QLocale__CanadianAboriginalScript    = QLocale__Script(44)
	QLocale__CarianScript                = QLocale__Script(45)
	QLocale__ChakmaScript                = QLocale__Script(46)
	QLocale__ChamScript                  = QLocale__Script(47)
	QLocale__CopticScript                = QLocale__Script(48)
	QLocale__CypriotScript               = QLocale__Script(49)
	QLocale__EgyptianHieroglyphsScript   = QLocale__Script(50)
	QLocale__FraserScript                = QLocale__Script(51)
	QLocale__GlagoliticScript            = QLocale__Script(52)
	QLocale__GothicScript                = QLocale__Script(53)
	QLocale__HanScript                   = QLocale__Script(54)
	QLocale__HangulScript                = QLocale__Script(55)
	QLocale__HanunooScript               = QLocale__Script(56)
	QLocale__ImperialAramaicScript       = QLocale__Script(57)
	QLocale__InscriptionalPahlaviScript  = QLocale__Script(58)
	QLocale__InscriptionalParthianScript = QLocale__Script(59)
	QLocale__JavaneseScript              = QLocale__Script(60)
	QLocale__KaithiScript                = QLocale__Script(61)
	QLocale__KatakanaScript              = QLocale__Script(62)
	QLocale__KayahLiScript               = QLocale__Script(63)
	QLocale__KharoshthiScript            = QLocale__Script(64)
	QLocale__LannaScript                 = QLocale__Script(65)
	QLocale__LepchaScript                = QLocale__Script(66)
	QLocale__LimbuScript                 = QLocale__Script(67)
	QLocale__LinearBScript               = QLocale__Script(68)
	QLocale__LycianScript                = QLocale__Script(69)
	QLocale__LydianScript                = QLocale__Script(70)
	QLocale__MandaeanScript              = QLocale__Script(71)
	QLocale__MeiteiMayekScript           = QLocale__Script(72)
	QLocale__MeroiticScript              = QLocale__Script(73)
	QLocale__MeroiticCursiveScript       = QLocale__Script(74)
	QLocale__NkoScript                   = QLocale__Script(75)
	QLocale__NewTaiLueScript             = QLocale__Script(76)
	QLocale__OghamScript                 = QLocale__Script(77)
	QLocale__OlChikiScript               = QLocale__Script(78)
	QLocale__OldItalicScript             = QLocale__Script(79)
	QLocale__OldPersianScript            = QLocale__Script(80)
	QLocale__OldSouthArabianScript       = QLocale__Script(81)
	QLocale__OrkhonScript                = QLocale__Script(82)
	QLocale__OsmanyaScript               = QLocale__Script(83)
	QLocale__PhagsPaScript               = QLocale__Script(84)
	QLocale__PhoenicianScript            = QLocale__Script(85)
	QLocale__PollardPhoneticScript       = QLocale__Script(86)
	QLocale__RejangScript                = QLocale__Script(87)
	QLocale__RunicScript                 = QLocale__Script(88)
	QLocale__SamaritanScript             = QLocale__Script(89)
	QLocale__SaurashtraScript            = QLocale__Script(90)
	QLocale__SharadaScript               = QLocale__Script(91)
	QLocale__ShavianScript               = QLocale__Script(92)
	QLocale__SoraSompengScript           = QLocale__Script(93)
	QLocale__CuneiformScript             = QLocale__Script(94)
	QLocale__SundaneseScript             = QLocale__Script(95)
	QLocale__SylotiNagriScript           = QLocale__Script(96)
	QLocale__TagalogScript               = QLocale__Script(97)
	QLocale__TagbanwaScript              = QLocale__Script(98)
	QLocale__TaiLeScript                 = QLocale__Script(99)
	QLocale__TaiVietScript               = QLocale__Script(100)
	QLocale__TakriScript                 = QLocale__Script(101)
	QLocale__UgariticScript              = QLocale__Script(102)
	QLocale__BrailleScript               = QLocale__Script(103)
	QLocale__HiraganaScript              = QLocale__Script(104)
	QLocale__CaucasianAlbanianScript     = QLocale__Script(105)
	QLocale__BassaVahScript              = QLocale__Script(106)
	QLocale__DuployanScript              = QLocale__Script(107)
	QLocale__ElbasanScript               = QLocale__Script(108)
	QLocale__GranthaScript               = QLocale__Script(109)
	QLocale__PahawhHmongScript           = QLocale__Script(110)
	QLocale__KhojkiScript                = QLocale__Script(111)
	QLocale__LinearAScript               = QLocale__Script(112)
	QLocale__MahajaniScript              = QLocale__Script(113)
	QLocale__ManichaeanScript            = QLocale__Script(114)
	QLocale__MendeKikakuiScript          = QLocale__Script(115)
	QLocale__ModiScript                  = QLocale__Script(116)
	QLocale__MroScript                   = QLocale__Script(117)
	QLocale__OldNorthArabianScript       = QLocale__Script(118)
	QLocale__NabataeanScript             = QLocale__Script(119)
	QLocale__PalmyreneScript             = QLocale__Script(120)
	QLocale__PauCinHauScript             = QLocale__Script(121)
	QLocale__OldPermicScript             = QLocale__Script(122)
	QLocale__PsalterPahlaviScript        = QLocale__Script(123)
	QLocale__SiddhamScript               = QLocale__Script(124)
	QLocale__KhudawadiScript             = QLocale__Script(125)
	QLocale__TirhutaScript               = QLocale__Script(126)
	QLocale__VarangKshitiScript          = QLocale__Script(127)
	QLocale__SimplifiedChineseScript     = QLocale__Script(QLocale__SimplifiedHanScript)
	QLocale__TraditionalChineseScript    = QLocale__Script(QLocale__TraditionalHanScript)
	QLocale__LastScript                  = QLocale__Script(QLocale__VarangKshitiScript)
)

func NewQLocale() *QLocale {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::QLocale")
		}
	}()

	return NewQLocaleFromPointer(C.QLocale_NewQLocale())
}

func NewQLocale3(language QLocale__Language, country QLocale__Country) *QLocale {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::QLocale")
		}
	}()

	return NewQLocaleFromPointer(C.QLocale_NewQLocale3(C.int(language), C.int(country)))
}

func NewQLocale4(language QLocale__Language, script QLocale__Script, country QLocale__Country) *QLocale {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::QLocale")
		}
	}()

	return NewQLocaleFromPointer(C.QLocale_NewQLocale4(C.int(language), C.int(script), C.int(country)))
}

func NewQLocale5(other QLocale_ITF) *QLocale {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::QLocale")
		}
	}()

	return NewQLocaleFromPointer(C.QLocale_NewQLocale5(PointerFromQLocale(other)))
}

func NewQLocale2(name string) *QLocale {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::QLocale")
		}
	}()

	return NewQLocaleFromPointer(C.QLocale_NewQLocale2(C.CString(name)))
}

func (ptr *QLocale) AmText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::amText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_AmText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocale) Bcp47Name() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::bcp47Name")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_Bcp47Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocale) Country() QLocale__Country {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::country")
		}
	}()

	if ptr.Pointer() != nil {
		return QLocale__Country(C.QLocale_Country(ptr.Pointer()))
	}
	return 0
}

func QLocale_CountryToString(country QLocale__Country) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::countryToString")
		}
	}()

	return C.GoString(C.QLocale_QLocale_CountryToString(C.int(country)))
}

func (ptr *QLocale) CreateSeparatedList(list []string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::createSeparatedList")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_CreateSeparatedList(ptr.Pointer(), C.CString(strings.Join(list, ",,,"))))
	}
	return ""
}

func (ptr *QLocale) CurrencySymbol(format QLocale__CurrencySymbolFormat) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::currencySymbol")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_CurrencySymbol(ptr.Pointer(), C.int(format)))
	}
	return ""
}

func (ptr *QLocale) DateFormat(format QLocale__FormatType) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::dateFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_DateFormat(ptr.Pointer(), C.int(format)))
	}
	return ""
}

func (ptr *QLocale) DateTimeFormat(format QLocale__FormatType) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::dateTimeFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_DateTimeFormat(ptr.Pointer(), C.int(format)))
	}
	return ""
}

func (ptr *QLocale) DayName(day int, ty QLocale__FormatType) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::dayName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_DayName(ptr.Pointer(), C.int(day), C.int(ty)))
	}
	return ""
}

func (ptr *QLocale) FirstDayOfWeek() Qt__DayOfWeek {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::firstDayOfWeek")
		}
	}()

	if ptr.Pointer() != nil {
		return Qt__DayOfWeek(C.QLocale_FirstDayOfWeek(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocale) Language() QLocale__Language {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::language")
		}
	}()

	if ptr.Pointer() != nil {
		return QLocale__Language(C.QLocale_Language(ptr.Pointer()))
	}
	return 0
}

func QLocale_LanguageToString(language QLocale__Language) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::languageToString")
		}
	}()

	return C.GoString(C.QLocale_QLocale_LanguageToString(C.int(language)))
}

func (ptr *QLocale) MeasurementSystem() QLocale__MeasurementSystem {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::measurementSystem")
		}
	}()

	if ptr.Pointer() != nil {
		return QLocale__MeasurementSystem(C.QLocale_MeasurementSystem(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocale) MonthName(month int, ty QLocale__FormatType) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::monthName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_MonthName(ptr.Pointer(), C.int(month), C.int(ty)))
	}
	return ""
}

func (ptr *QLocale) Name() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::name")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocale) NativeCountryName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::nativeCountryName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_NativeCountryName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocale) NativeLanguageName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::nativeLanguageName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_NativeLanguageName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocale) NumberOptions() QLocale__NumberOption {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::numberOptions")
		}
	}()

	if ptr.Pointer() != nil {
		return QLocale__NumberOption(C.QLocale_NumberOptions(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocale) PmText() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::pmText")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_PmText(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLocale) QuoteString(str string, style QLocale__QuotationStyle) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::quoteString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_QuoteString(ptr.Pointer(), C.CString(str), C.int(style)))
	}
	return ""
}

func (ptr *QLocale) QuoteString2(str QStringRef_ITF, style QLocale__QuotationStyle) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::quoteString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_QuoteString2(ptr.Pointer(), PointerFromQStringRef(str), C.int(style)))
	}
	return ""
}

func (ptr *QLocale) Script() QLocale__Script {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::script")
		}
	}()

	if ptr.Pointer() != nil {
		return QLocale__Script(C.QLocale_Script(ptr.Pointer()))
	}
	return 0
}

func QLocale_ScriptToString(script QLocale__Script) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::scriptToString")
		}
	}()

	return C.GoString(C.QLocale_QLocale_ScriptToString(C.int(script)))
}

func QLocale_SetDefault(locale QLocale_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::setDefault")
		}
	}()

	C.QLocale_QLocale_SetDefault(PointerFromQLocale(locale))
}

func (ptr *QLocale) SetNumberOptions(options QLocale__NumberOption) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::setNumberOptions")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLocale_SetNumberOptions(ptr.Pointer(), C.int(options))
	}
}

func (ptr *QLocale) StandaloneDayName(day int, ty QLocale__FormatType) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::standaloneDayName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_StandaloneDayName(ptr.Pointer(), C.int(day), C.int(ty)))
	}
	return ""
}

func (ptr *QLocale) StandaloneMonthName(month int, ty QLocale__FormatType) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::standaloneMonthName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_StandaloneMonthName(ptr.Pointer(), C.int(month), C.int(ty)))
	}
	return ""
}

func (ptr *QLocale) TextDirection() Qt__LayoutDirection {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::textDirection")
		}
	}()

	if ptr.Pointer() != nil {
		return Qt__LayoutDirection(C.QLocale_TextDirection(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLocale) TimeFormat(format QLocale__FormatType) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::timeFormat")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_TimeFormat(ptr.Pointer(), C.int(format)))
	}
	return ""
}

func (ptr *QLocale) ToCurrencyString6(value int, symbol string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::toCurrencyString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_ToCurrencyString6(ptr.Pointer(), C.int(value), C.CString(symbol)))
	}
	return ""
}

func (ptr *QLocale) ToDateTime(stri string, format QLocale__FormatType) *QDateTime {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::toDateTime")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQDateTimeFromPointer(C.QLocale_ToDateTime(ptr.Pointer(), C.CString(stri), C.int(format)))
	}
	return nil
}

func (ptr *QLocale) ToDateTime2(stri string, format string) *QDateTime {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::toDateTime")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQDateTimeFromPointer(C.QLocale_ToDateTime2(ptr.Pointer(), C.CString(stri), C.CString(format)))
	}
	return nil
}

func (ptr *QLocale) ToInt(s string, ok bool) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::toInt")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLocale_ToInt(ptr.Pointer(), C.CString(s), C.int(qt.GoBoolToInt(ok))))
	}
	return 0
}

func (ptr *QLocale) ToInt2(s QStringRef_ITF, ok bool) int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::toInt")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QLocale_ToInt2(ptr.Pointer(), PointerFromQStringRef(s), C.int(qt.GoBoolToInt(ok))))
	}
	return 0
}

func (ptr *QLocale) ToLower(str string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::toLower")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_ToLower(ptr.Pointer(), C.CString(str)))
	}
	return ""
}

func (ptr *QLocale) ToString3(date QDate_ITF, format QLocale__FormatType) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::toString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_ToString3(ptr.Pointer(), PointerFromQDate(date), C.int(format)))
	}
	return ""
}

func (ptr *QLocale) ToString2(date QDate_ITF, format string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::toString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_ToString2(ptr.Pointer(), PointerFromQDate(date), C.CString(format)))
	}
	return ""
}

func (ptr *QLocale) ToString6(dateTime QDateTime_ITF, format QLocale__FormatType) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::toString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_ToString6(ptr.Pointer(), PointerFromQDateTime(dateTime), C.int(format)))
	}
	return ""
}

func (ptr *QLocale) ToString7(dateTime QDateTime_ITF, format string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::toString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_ToString7(ptr.Pointer(), PointerFromQDateTime(dateTime), C.CString(format)))
	}
	return ""
}

func (ptr *QLocale) ToString5(time QTime_ITF, format QLocale__FormatType) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::toString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_ToString5(ptr.Pointer(), PointerFromQTime(time), C.int(format)))
	}
	return ""
}

func (ptr *QLocale) ToString4(time QTime_ITF, format string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::toString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_ToString4(ptr.Pointer(), PointerFromQTime(time), C.CString(format)))
	}
	return ""
}

func (ptr *QLocale) ToString12(i int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::toString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_ToString12(ptr.Pointer(), C.int(i)))
	}
	return ""
}

func (ptr *QLocale) ToUpper(str string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::toUpper")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QLocale_ToUpper(ptr.Pointer(), C.CString(str)))
	}
	return ""
}

func (ptr *QLocale) UiLanguages() []string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::uiLanguages")
		}
	}()

	if ptr.Pointer() != nil {
		return strings.Split(C.GoString(C.QLocale_UiLanguages(ptr.Pointer())), ",,,")
	}
	return make([]string, 0)
}

func (ptr *QLocale) DestroyQLocale() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QLocale::~QLocale")
		}
	}()

	if ptr.Pointer() != nil {
		C.QLocale_DestroyQLocale(ptr.Pointer())
	}
}
