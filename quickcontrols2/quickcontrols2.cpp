// +build !minimal

#define protected public
#define private public

#include "quickcontrols2.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QQuickStyle>
#include <QString>

struct QtQuickControls2_PackedString QQuickStyle_QQuickStyle_Name()
{
	return ({ QByteArray tca3084 = QQuickStyle::name().toUtf8(); QtQuickControls2_PackedString { const_cast<char*>(tca3084.prepend("WHITESPACE").constData()+10), tca3084.size()-10 }; });
}

struct QtQuickControls2_PackedString QQuickStyle_QQuickStyle_Path()
{
	return ({ QByteArray t432960 = QQuickStyle::path().toUtf8(); QtQuickControls2_PackedString { const_cast<char*>(t432960.prepend("WHITESPACE").constData()+10), t432960.size()-10 }; });
}

struct QtQuickControls2_PackedString QQuickStyle_QQuickStyle_AvailableStyles()
{
	return ({ QByteArray td4aa30 = QQuickStyle::availableStyles().join("|").toUtf8(); QtQuickControls2_PackedString { const_cast<char*>(td4aa30.prepend("WHITESPACE").constData()+10), td4aa30.size()-10 }; });
}

void QQuickStyle_QQuickStyle_SetFallbackStyle(struct QtQuickControls2_PackedString style)
{
	QQuickStyle::setFallbackStyle(QString::fromUtf8(style.data, style.len));
}

void QQuickStyle_QQuickStyle_SetStyle(struct QtQuickControls2_PackedString style)
{
	QQuickStyle::setStyle(QString::fromUtf8(style.data, style.len));
}

