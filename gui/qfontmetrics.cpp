#include "qfontmetrics.h"
#include <QChar>
#include <QPaintDevice>
#include <QFont>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QFontMetrics>
#include "_cgo_export.h"

class MyQFontMetrics: public QFontMetrics {
public:
};

void* QFontMetrics_NewQFontMetrics(void* font){
	return new QFontMetrics(*static_cast<QFont*>(font));
}

void* QFontMetrics_NewQFontMetrics2(void* font, void* paintdevice){
	return new QFontMetrics(*static_cast<QFont*>(font), static_cast<QPaintDevice*>(paintdevice));
}

void* QFontMetrics_NewQFontMetrics3(void* fm){
	return new QFontMetrics(*static_cast<QFontMetrics*>(fm));
}

int QFontMetrics_Ascent(void* ptr){
	return static_cast<QFontMetrics*>(ptr)->ascent();
}

int QFontMetrics_AverageCharWidth(void* ptr){
	return static_cast<QFontMetrics*>(ptr)->averageCharWidth();
}

int QFontMetrics_Descent(void* ptr){
	return static_cast<QFontMetrics*>(ptr)->descent();
}

char* QFontMetrics_ElidedText(void* ptr, char* text, int mode, int width, int flags){
	return static_cast<QFontMetrics*>(ptr)->elidedText(QString(text), static_cast<Qt::TextElideMode>(mode), width, flags).toUtf8().data();
}

int QFontMetrics_Height(void* ptr){
	return static_cast<QFontMetrics*>(ptr)->height();
}

int QFontMetrics_InFont(void* ptr, void* ch){
	return static_cast<QFontMetrics*>(ptr)->inFont(*static_cast<QChar*>(ch));
}

int QFontMetrics_Leading(void* ptr){
	return static_cast<QFontMetrics*>(ptr)->leading();
}

int QFontMetrics_LeftBearing(void* ptr, void* ch){
	return static_cast<QFontMetrics*>(ptr)->leftBearing(*static_cast<QChar*>(ch));
}

int QFontMetrics_LineSpacing(void* ptr){
	return static_cast<QFontMetrics*>(ptr)->lineSpacing();
}

int QFontMetrics_LineWidth(void* ptr){
	return static_cast<QFontMetrics*>(ptr)->lineWidth();
}

int QFontMetrics_MaxWidth(void* ptr){
	return static_cast<QFontMetrics*>(ptr)->maxWidth();
}

int QFontMetrics_MinLeftBearing(void* ptr){
	return static_cast<QFontMetrics*>(ptr)->minLeftBearing();
}

int QFontMetrics_MinRightBearing(void* ptr){
	return static_cast<QFontMetrics*>(ptr)->minRightBearing();
}

int QFontMetrics_OverlinePos(void* ptr){
	return static_cast<QFontMetrics*>(ptr)->overlinePos();
}

int QFontMetrics_RightBearing(void* ptr, void* ch){
	return static_cast<QFontMetrics*>(ptr)->rightBearing(*static_cast<QChar*>(ch));
}

int QFontMetrics_StrikeOutPos(void* ptr){
	return static_cast<QFontMetrics*>(ptr)->strikeOutPos();
}

void QFontMetrics_Swap(void* ptr, void* other){
	static_cast<QFontMetrics*>(ptr)->swap(*static_cast<QFontMetrics*>(other));
}

int QFontMetrics_UnderlinePos(void* ptr){
	return static_cast<QFontMetrics*>(ptr)->underlinePos();
}

int QFontMetrics_Width3(void* ptr, void* ch){
	return static_cast<QFontMetrics*>(ptr)->width(*static_cast<QChar*>(ch));
}

int QFontMetrics_Width(void* ptr, char* text, int len){
	return static_cast<QFontMetrics*>(ptr)->width(QString(text), len);
}

int QFontMetrics_XHeight(void* ptr){
	return static_cast<QFontMetrics*>(ptr)->xHeight();
}

void QFontMetrics_DestroyQFontMetrics(void* ptr){
	static_cast<QFontMetrics*>(ptr)->~QFontMetrics();
}

