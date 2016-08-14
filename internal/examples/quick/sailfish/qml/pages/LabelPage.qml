/****************************************************************************************
**
** Copyright (C) 2013 Jolla Ltd.
** Contact: Matt Vogt <matthew.vogt@jollamobile.com>
** All rights reserved.
** 
** This file is part of Sailfish Silica UI component package.
**
** You may use this file under the terms of BSD license as follows:
**
** Redistribution and use in source and binary forms, with or without
** modification, are permitted provided that the following conditions are met:
**     * Redistributions of source code must retain the above copyright
**       notice, this list of conditions and the following disclaimer.
**     * Redistributions in binary form must reproduce the above copyright
**       notice, this list of conditions and the following disclaimer in the
**       documentation and/or other materials provided with the distribution.
**     * Neither the name of the Jolla Ltd nor the
**       names of its contributors may be used to endorse or promote products
**       derived from this software without specific prior written permission.
** 
** THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
** ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
** WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
** DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDERS OR CONTRIBUTORS BE LIABLE FOR
** ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES
** (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES;
** LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND
** ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
** (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE OF THIS
** SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
**
****************************************************************************************/

import QtQuick 2.0
import Sailfish.Silica 1.0

Page {
    id: page
    SilicaFlickable {
        anchors.fill: parent
        contentWidth: parent.width
        contentHeight: col.height + Theme.paddingLarge

        VerticalScrollDecorator {}

        Column {
            id: col
            spacing: Theme.paddingLarge
            width: parent.width
            PageHeader {
                title: "Labels"
            }
            SectionHeader {
                text: "Text styling"
            }
            Label {
                text: "Header text"
                anchors.horizontalCenter: parent.horizontalCenter
                font.family: Theme.fontFamilyHeading
            }
            Label {
                text: "Normal text"
                anchors.horizontalCenter: parent.horizontalCenter
            }
            Label {
                text: "Secondary text"
                color: Theme.secondaryColor
                anchors.horizontalCenter: parent.horizontalCenter
            }
            Label {
                text: "Highlighted text"
                color: Theme.highlightColor
                anchors.horizontalCenter: parent.horizontalCenter
            }
            Label {
                text: "Secondary highlighted text"
                color: Theme.secondaryHighlightColor
                anchors.horizontalCenter: parent.horizontalCenter
            }
            Rectangle {
                color: Theme.highlightBackgroundColor
                anchors.horizontalCenter: parent.horizontalCenter
                height: Theme.itemSizeSmall
                width: page.width
                Label {
                    text: "Normal text with background"
                    anchors.centerIn: parent
                }
            }
            SectionHeader {
                text: "Text eliding"
            }
            Label {
                text: "Text that should not be elided or faded out"
                width: Math.min(page.width - 2*Theme.horizontalPageMargin, implicitWidth)
                anchors.horizontalCenter: parent.horizontalCenter
                horizontalAlignment: Text.AlignLeft
                color: Theme.secondaryHighlightColor
            }
            Label {
                text: "Text that should be elided off the right end"
                width: Math.min(page.width - 2*Theme.horizontalPageMargin, implicitWidth*0.9)
                anchors.horizontalCenter: parent.horizontalCenter
                horizontalAlignment: Text.AlignLeft
                truncationMode: TruncationMode.Elide
                color: Theme.secondaryHighlightColor
            }
            Label {
                text: "Text that should be elided off the left end"
                width: Math.min(page.width - 2*Theme.horizontalPageMargin, implicitWidth*0.9)
                anchors.horizontalCenter: parent.horizontalCenter
                horizontalAlignment: Text.AlignRight
                truncationMode: TruncationMode.Elide
                color: Theme.secondaryHighlightColor
            }
            Label {
                text: "Text that should be faded out rather than elided"
                width: Math.min(page.width - 2*Theme.horizontalPageMargin, implicitWidth*0.9)
                anchors.horizontalCenter: parent.horizontalCenter
                horizontalAlignment: Text.AlignLeft
                truncationMode: TruncationMode.Fade
                color: Theme.secondaryHighlightColor
            }
            Label {
                text: "Text that should be faded out from the left"
                width: Math.min(page.width - 2*Theme.horizontalPageMargin, implicitWidth*0.9)
                anchors.horizontalCenter: parent.horizontalCenter
                horizontalAlignment: Text.AlignRight
                truncationMode: TruncationMode.Fade
                color: Theme.secondaryHighlightColor
            }

            SectionHeader {
                text: "Scripts"
            }
            Column {
                x: Theme.paddingLarge
                width: parent.width - 2*x
                spacing: Theme.paddingLarge
                Label {
                    // Latin
                    text: "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."
                    color: Theme.highlightColor
                    wrapMode: Text.Wrap
                    width: parent.width
                }
                Label {
                    // Arabic
                    text: "من ناحية أخرى، ندين باستياء الصالحين ويكره الرجال الذين سحرت جدا ومحبط من سحر متعة لحظة، يعميها الرغبة، وأنها لا يمكن أن نتوقع من الألم والمتاعب التي لا بد أن تترتب على ذلك. واللوم على قدم المساواة ينتمي إلى أولئك الذين يفشلون في أداء واجبها من خلال ضعف الإرادة، والذي هو نفسه قوله من خلال تقلص من الكدح والألم."
                    color: Theme.highlightColor
                    wrapMode: Text.Wrap
                    width: parent.width
                }
                Label {
                    // Simplified Chinese
                    text: "在另一方面，我们谴责与义愤和不喜欢的人谁是如此引诱，士气低落的时刻愉悦的魅力，被欲望蒙蔽这样，他们无法预见的痛苦，而且也必然随之而来的麻烦;和平等的怪属于那些软弱的意志，这是一样的，通过从辛劳和痛苦缩小说在他们的责任谁失败。"
                    color: Theme.highlightColor
                    wrapMode: Text.Wrap
                    width: parent.width
                }
                Label {
                    // Cyrillic
                    text: "С другой стороны, мы осуждаем с праведным негодованием и не нравится мужчинам, которые так развлекали и деморализованных прелестями удовольствия данный момент, так ослеплен желанием, что они не могут предвидеть боль и проблемы, которые связаны с наступить; и равна вина принадлежит тем, кто не в их обязанности через слабости воли, что то же самое, как говорят через сокращение от тяжелого труда и боли."
                    color: Theme.highlightColor
                    wrapMode: Text.Wrap
                    width: parent.width
                }
                Label {
                    // Hindi
                    text: "दूसरी ओर, हम तो वे दर्द और पीछा करने के लिए बाध्य कर रहे हैं कि मुसीबत उम्मीद नहीं कर सकते कि, इतनी इच्छा से अंधे, पल की खुशी का आकर्षण द्वारा और हतोत्साहित कर रहे हैं, जो पुरुषों धर्मी रोष के साथ निंदा और नापसंद; और बराबर दोष परिश्रम और दर्द से सिकुड़ने के माध्यम से यह कहते हुए एक ही है जो इच्छाशक्ति की कमजोरी, के माध्यम से अपने कर्तव्य में असफल, जो उन लोगों के अंतर्गत आता है।"
                    color: Theme.highlightColor
                    wrapMode: Text.Wrap
                    width: parent.width
                }
                Label {
                    // Bengali
                    text: "সমস্ত মানুষ স্বাধীনভাবে সমান মর্যাদা এবং অধিকার নিয়ে জন্মগ্রহণ করে। তাঁদের বিবেক এবং বুদ্ধি আছে; সুতরাং সকলেরই একে অপরের প্রতি ভ্রাতৃত্বসুলভ মনোভাব নিয়ে আচরণ করা উচিৎ।"
                    color: Theme.highlightColor
                    wrapMode: Text.Wrap
                    width: parent.width
                }
                Label {
                    // Tamil
                    text: "தமிழ் அரிச்சுவடி"
                    color: Theme.highlightColor
                    wrapMode: Text.Wrap
                    width: parent.width
                }
                Label {
                    // Telugu
                    text: "తెలుగు లిపి"
                    color: Theme.highlightColor
                    wrapMode: Text.Wrap
                    width: parent.width
                }
                Label {
                    // Punjabi / gurmughi
                    text: "ਲਹੌਰ ਪਾਕਿਸਤਾਨੀ ਪੰਜਾਬ ਦੀ ਰਾਜਧਾਨੀ ਹੈ । ਲੋਕ ਗਿਣਤੀ ਦੇ ਨਾਲ ਕਰਾਚੀ ਤੋਂ ਬਾਅਦ ਲਹੌਰ ਦੂਜਾ ਸਭ ਤੋਂ ਵੱਡਾ ਸ਼ਹਿਰ ਹੈ । ਲਹੌਰ ਪਾਕਿਸਤਾਨ ਦਾ ਸਿਆਸੀ"
                    color: Theme.highlightColor
                    wrapMode: Text.Wrap
                    width: parent.width
                }

                Label {
                    // kannada
                    text: "ಕನ್ನಡ"
                    color: Theme.highlightColor
                    wrapMode: Text.Wrap
                    width: parent.width
                }
                Label {
                    // malayalam
                    text: "മലയാളം"
                    color: Theme.highlightColor
                    wrapMode: Text.Wrap
                    width: parent.width
                }
                Label {
                    // Emoticon
                    text: "😀 😁 😂 😃 😄 😅 😆 😇 😈 😉 😊 😋 😌 😍 😎 😏 😐 😑 😒 😓 😔 😕 😖 😗 😘 😙 😚 😛 😜 😝 😞 😟 😠 😡 😢 😣 😤 😥 😦 😧 😨 😩 😪 😫 😬 😭 😮 😯 😰 😱 😲 😳 😴 😵 😶 😷 😸 😹 😺 😻 😼 😽 😾 😿 🙀"
                    color: Theme.highlightColor
                    wrapMode: Text.Wrap
                    width: parent.width
                }
            }
        }
    }
}
