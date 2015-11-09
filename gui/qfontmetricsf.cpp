#include "qfontmetricsf.h"
#include <QModelIndex>
#include <QFontMetrics>
#include <QChar>
#include <QPaintDevice>
#include <QFont>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QFontMetricsF>
#include "_cgo_export.h"

class MyQFontMetricsF: public QFontMetricsF {
public:
};

void* QFontMetricsF_NewQFontMetricsF(void* font){
	return new QFontMetricsF(*static_cast<QFont*>(font));
}

void* QFontMetricsF_NewQFontMetricsF2(void* font, void* paintdevice){
	return new QFontMetricsF(*static_cast<QFont*>(font), static_cast<QPaintDevice*>(paintdevice));
}

void* QFontMetricsF_NewQFontMetricsF3(void* fontMetrics){
	return new QFontMetricsF(*static_cast<QFontMetrics*>(fontMetrics));
}

void* QFontMetricsF_NewQFontMetricsF4(void* fm){
	return new QFontMetricsF(*static_cast<QFontMetricsF*>(fm));
}

double QFontMetricsF_Ascent(void* ptr){
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->ascent());
}

double QFontMetricsF_AverageCharWidth(void* ptr){
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->averageCharWidth());
}

double QFontMetricsF_Descent(void* ptr){
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->descent());
}

char* QFontMetricsF_ElidedText(void* ptr, char* text, int mode, double width, int flags){
	return static_cast<QFontMetricsF*>(ptr)->elidedText(QString(text), static_cast<Qt::TextElideMode>(mode), static_cast<qreal>(width), flags).toUtf8().data();
}

double QFontMetricsF_Height(void* ptr){
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->height());
}

int QFontMetricsF_InFont(void* ptr, void* ch){
	return static_cast<QFontMetricsF*>(ptr)->inFont(*static_cast<QChar*>(ch));
}

double QFontMetricsF_Leading(void* ptr){
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->leading());
}

double QFontMetricsF_LeftBearing(void* ptr, void* ch){
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->leftBearing(*static_cast<QChar*>(ch)));
}

double QFontMetricsF_LineSpacing(void* ptr){
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->lineSpacing());
}

double QFontMetricsF_LineWidth(void* ptr){
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->lineWidth());
}

double QFontMetricsF_MaxWidth(void* ptr){
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->maxWidth());
}

double QFontMetricsF_MinLeftBearing(void* ptr){
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->minLeftBearing());
}

double QFontMetricsF_MinRightBearing(void* ptr){
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->minRightBearing());
}

double QFontMetricsF_OverlinePos(void* ptr){
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->overlinePos());
}

double QFontMetricsF_RightBearing(void* ptr, void* ch){
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->rightBearing(*static_cast<QChar*>(ch)));
}

double QFontMetricsF_StrikeOutPos(void* ptr){
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->strikeOutPos());
}

void QFontMetricsF_Swap(void* ptr, void* other){
	static_cast<QFontMetricsF*>(ptr)->swap(*static_cast<QFontMetricsF*>(other));
}

double QFontMetricsF_UnderlinePos(void* ptr){
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->underlinePos());
}

double QFontMetricsF_Width2(void* ptr, void* ch){
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->width(*static_cast<QChar*>(ch)));
}

double QFontMetricsF_Width(void* ptr, char* text){
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->width(QString(text)));
}

double QFontMetricsF_XHeight(void* ptr){
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->xHeight());
}

void QFontMetricsF_DestroyQFontMetricsF(void* ptr){
	static_cast<QFontMetricsF*>(ptr)->~QFontMetricsF();
}

