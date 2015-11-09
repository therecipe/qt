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

void* QFontInfo_NewQFontInfo(void* font){
	return new QFontInfo(*static_cast<QFont*>(font));
}

void* QFontInfo_NewQFontInfo2(void* fi){
	return new QFontInfo(*static_cast<QFontInfo*>(fi));
}

int QFontInfo_Bold(void* ptr){
	return static_cast<QFontInfo*>(ptr)->bold();
}

int QFontInfo_ExactMatch(void* ptr){
	return static_cast<QFontInfo*>(ptr)->exactMatch();
}

char* QFontInfo_Family(void* ptr){
	return static_cast<QFontInfo*>(ptr)->family().toUtf8().data();
}

int QFontInfo_FixedPitch(void* ptr){
	return static_cast<QFontInfo*>(ptr)->fixedPitch();
}

int QFontInfo_Italic(void* ptr){
	return static_cast<QFontInfo*>(ptr)->italic();
}

int QFontInfo_PixelSize(void* ptr){
	return static_cast<QFontInfo*>(ptr)->pixelSize();
}

int QFontInfo_PointSize(void* ptr){
	return static_cast<QFontInfo*>(ptr)->pointSize();
}

double QFontInfo_PointSizeF(void* ptr){
	return static_cast<double>(static_cast<QFontInfo*>(ptr)->pointSizeF());
}

int QFontInfo_Style(void* ptr){
	return static_cast<QFontInfo*>(ptr)->style();
}

char* QFontInfo_StyleName(void* ptr){
	return static_cast<QFontInfo*>(ptr)->styleName().toUtf8().data();
}

void QFontInfo_Swap(void* ptr, void* other){
	static_cast<QFontInfo*>(ptr)->swap(*static_cast<QFontInfo*>(other));
}

int QFontInfo_StyleHint(void* ptr){
	return static_cast<QFontInfo*>(ptr)->styleHint();
}

int QFontInfo_Weight(void* ptr){
	return static_cast<QFontInfo*>(ptr)->weight();
}

void QFontInfo_DestroyQFontInfo(void* ptr){
	static_cast<QFontInfo*>(ptr)->~QFontInfo();
}

