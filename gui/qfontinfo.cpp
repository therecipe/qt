#include "qfontinfo.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QFont>
#include <QFontInfo>
#include "_cgo_export.h"

class MyQFontInfo: public QFontInfo {
public:
};

QtObjectPtr QFontInfo_NewQFontInfo(QtObjectPtr font){
	return new QFontInfo(*static_cast<QFont*>(font));
}

QtObjectPtr QFontInfo_NewQFontInfo2(QtObjectPtr fi){
	return new QFontInfo(*static_cast<QFontInfo*>(fi));
}

int QFontInfo_Bold(QtObjectPtr ptr){
	return static_cast<QFontInfo*>(ptr)->bold();
}

int QFontInfo_ExactMatch(QtObjectPtr ptr){
	return static_cast<QFontInfo*>(ptr)->exactMatch();
}

char* QFontInfo_Family(QtObjectPtr ptr){
	return static_cast<QFontInfo*>(ptr)->family().toUtf8().data();
}

int QFontInfo_FixedPitch(QtObjectPtr ptr){
	return static_cast<QFontInfo*>(ptr)->fixedPitch();
}

int QFontInfo_Italic(QtObjectPtr ptr){
	return static_cast<QFontInfo*>(ptr)->italic();
}

int QFontInfo_PixelSize(QtObjectPtr ptr){
	return static_cast<QFontInfo*>(ptr)->pixelSize();
}

int QFontInfo_PointSize(QtObjectPtr ptr){
	return static_cast<QFontInfo*>(ptr)->pointSize();
}

int QFontInfo_Style(QtObjectPtr ptr){
	return static_cast<QFontInfo*>(ptr)->style();
}

char* QFontInfo_StyleName(QtObjectPtr ptr){
	return static_cast<QFontInfo*>(ptr)->styleName().toUtf8().data();
}

void QFontInfo_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QFontInfo*>(ptr)->swap(*static_cast<QFontInfo*>(other));
}

int QFontInfo_StyleHint(QtObjectPtr ptr){
	return static_cast<QFontInfo*>(ptr)->styleHint();
}

int QFontInfo_Weight(QtObjectPtr ptr){
	return static_cast<QFontInfo*>(ptr)->weight();
}

void QFontInfo_DestroyQFontInfo(QtObjectPtr ptr){
	static_cast<QFontInfo*>(ptr)->~QFontInfo();
}

