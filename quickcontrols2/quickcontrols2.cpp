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

void QQuickStyle_QQuickStyle_SetFallbackStyle(char* style)
{
	QQuickStyle::setFallbackStyle(QString(style));
}

void QQuickStyle_QQuickStyle_SetStyle(char* style)
{
	QQuickStyle::setStyle(QString(style));
}

