#include "qfontmetrics.h"
#include <QModelIndex>
#include <QPaintDevice>
#include <QChar>
#include <QFont>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QFontMetrics>
#include "_cgo_export.h"

class MyQFontMetrics: public QFontMetrics {
public:
};

QtObjectPtr QFontMetrics_NewQFontMetrics(QtObjectPtr font){
	return new QFontMetrics(*static_cast<QFont*>(font));
}

QtObjectPtr QFontMetrics_NewQFontMetrics2(QtObjectPtr font, QtObjectPtr paintdevice){
	return new QFontMetrics(*static_cast<QFont*>(font), static_cast<QPaintDevice*>(paintdevice));
}

QtObjectPtr QFontMetrics_NewQFontMetrics3(QtObjectPtr fm){
	return new QFontMetrics(*static_cast<QFontMetrics*>(fm));
}

int QFontMetrics_Ascent(QtObjectPtr ptr){
	return static_cast<QFontMetrics*>(ptr)->ascent();
}

int QFontMetrics_AverageCharWidth(QtObjectPtr ptr){
	return static_cast<QFontMetrics*>(ptr)->averageCharWidth();
}

int QFontMetrics_Descent(QtObjectPtr ptr){
	return static_cast<QFontMetrics*>(ptr)->descent();
}

char* QFontMetrics_ElidedText(QtObjectPtr ptr, char* text, int mode, int width, int flags){
	return static_cast<QFontMetrics*>(ptr)->elidedText(QString(text), static_cast<Qt::TextElideMode>(mode), width, flags).toUtf8().data();
}

int QFontMetrics_Height(QtObjectPtr ptr){
	return static_cast<QFontMetrics*>(ptr)->height();
}

int QFontMetrics_InFont(QtObjectPtr ptr, QtObjectPtr ch){
	return static_cast<QFontMetrics*>(ptr)->inFont(*static_cast<QChar*>(ch));
}

int QFontMetrics_Leading(QtObjectPtr ptr){
	return static_cast<QFontMetrics*>(ptr)->leading();
}

int QFontMetrics_LeftBearing(QtObjectPtr ptr, QtObjectPtr ch){
	return static_cast<QFontMetrics*>(ptr)->leftBearing(*static_cast<QChar*>(ch));
}

int QFontMetrics_LineSpacing(QtObjectPtr ptr){
	return static_cast<QFontMetrics*>(ptr)->lineSpacing();
}

int QFontMetrics_LineWidth(QtObjectPtr ptr){
	return static_cast<QFontMetrics*>(ptr)->lineWidth();
}

int QFontMetrics_MaxWidth(QtObjectPtr ptr){
	return static_cast<QFontMetrics*>(ptr)->maxWidth();
}

int QFontMetrics_MinLeftBearing(QtObjectPtr ptr){
	return static_cast<QFontMetrics*>(ptr)->minLeftBearing();
}

int QFontMetrics_MinRightBearing(QtObjectPtr ptr){
	return static_cast<QFontMetrics*>(ptr)->minRightBearing();
}

int QFontMetrics_OverlinePos(QtObjectPtr ptr){
	return static_cast<QFontMetrics*>(ptr)->overlinePos();
}

int QFontMetrics_RightBearing(QtObjectPtr ptr, QtObjectPtr ch){
	return static_cast<QFontMetrics*>(ptr)->rightBearing(*static_cast<QChar*>(ch));
}

int QFontMetrics_StrikeOutPos(QtObjectPtr ptr){
	return static_cast<QFontMetrics*>(ptr)->strikeOutPos();
}

void QFontMetrics_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QFontMetrics*>(ptr)->swap(*static_cast<QFontMetrics*>(other));
}

int QFontMetrics_UnderlinePos(QtObjectPtr ptr){
	return static_cast<QFontMetrics*>(ptr)->underlinePos();
}

int QFontMetrics_Width3(QtObjectPtr ptr, QtObjectPtr ch){
	return static_cast<QFontMetrics*>(ptr)->width(*static_cast<QChar*>(ch));
}

int QFontMetrics_Width(QtObjectPtr ptr, char* text, int len){
	return static_cast<QFontMetrics*>(ptr)->width(QString(text), len);
}

int QFontMetrics_XHeight(QtObjectPtr ptr){
	return static_cast<QFontMetrics*>(ptr)->xHeight();
}

void QFontMetrics_DestroyQFontMetrics(QtObjectPtr ptr){
	static_cast<QFontMetrics*>(ptr)->~QFontMetrics();
}

