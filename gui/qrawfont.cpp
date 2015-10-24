#include "qrawfont.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QChar>
#include <QRawFont>
#include "_cgo_export.h"

class MyQRawFont: public QRawFont {
public:
};

QtObjectPtr QRawFont_NewQRawFont(){
	return new QRawFont();
}

QtObjectPtr QRawFont_NewQRawFont4(QtObjectPtr other){
	return new QRawFont(*static_cast<QRawFont*>(other));
}

char* QRawFont_FamilyName(QtObjectPtr ptr){
	return static_cast<QRawFont*>(ptr)->familyName().toUtf8().data();
}

int QRawFont_HintingPreference(QtObjectPtr ptr){
	return static_cast<QRawFont*>(ptr)->hintingPreference();
}

int QRawFont_IsValid(QtObjectPtr ptr){
	return static_cast<QRawFont*>(ptr)->isValid();
}

int QRawFont_Style(QtObjectPtr ptr){
	return static_cast<QRawFont*>(ptr)->style();
}

char* QRawFont_StyleName(QtObjectPtr ptr){
	return static_cast<QRawFont*>(ptr)->styleName().toUtf8().data();
}

int QRawFont_SupportsCharacter(QtObjectPtr ptr, QtObjectPtr character){
	return static_cast<QRawFont*>(ptr)->supportsCharacter(*static_cast<QChar*>(character));
}

void QRawFont_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QRawFont*>(ptr)->swap(*static_cast<QRawFont*>(other));
}

int QRawFont_Weight(QtObjectPtr ptr){
	return static_cast<QRawFont*>(ptr)->weight();
}

void QRawFont_DestroyQRawFont(QtObjectPtr ptr){
	static_cast<QRawFont*>(ptr)->~QRawFont();
}

