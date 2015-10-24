#include "qfontmetricsf.h"
#include <QFont>
#include <QFontMetrics>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPaintDevice>
#include <QChar>
#include <QFontMetricsF>
#include "_cgo_export.h"

class MyQFontMetricsF: public QFontMetricsF {
public:
};

QtObjectPtr QFontMetricsF_NewQFontMetricsF(QtObjectPtr font){
	return new QFontMetricsF(*static_cast<QFont*>(font));
}

QtObjectPtr QFontMetricsF_NewQFontMetricsF2(QtObjectPtr font, QtObjectPtr paintdevice){
	return new QFontMetricsF(*static_cast<QFont*>(font), static_cast<QPaintDevice*>(paintdevice));
}

QtObjectPtr QFontMetricsF_NewQFontMetricsF3(QtObjectPtr fontMetrics){
	return new QFontMetricsF(*static_cast<QFontMetrics*>(fontMetrics));
}

QtObjectPtr QFontMetricsF_NewQFontMetricsF4(QtObjectPtr fm){
	return new QFontMetricsF(*static_cast<QFontMetricsF*>(fm));
}

int QFontMetricsF_InFont(QtObjectPtr ptr, QtObjectPtr ch){
	return static_cast<QFontMetricsF*>(ptr)->inFont(*static_cast<QChar*>(ch));
}

void QFontMetricsF_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QFontMetricsF*>(ptr)->swap(*static_cast<QFontMetricsF*>(other));
}

void QFontMetricsF_DestroyQFontMetricsF(QtObjectPtr ptr){
	static_cast<QFontMetricsF*>(ptr)->~QFontMetricsF();
}

