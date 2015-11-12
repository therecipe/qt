#include "qrawfont.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QByteArray>
#include <QChar>
#include <QFont>
#include <QString>
#include <QRawFont>
#include "_cgo_export.h"

class MyQRawFont: public QRawFont {
public:
};

void* QRawFont_NewQRawFont(){
	return new QRawFont();
}

void* QRawFont_NewQRawFont3(void* fontData, double pixelSize, int hintingPreference){
	return new QRawFont(*static_cast<QByteArray*>(fontData), static_cast<qreal>(pixelSize), static_cast<QFont::HintingPreference>(hintingPreference));
}

void* QRawFont_NewQRawFont4(void* other){
	return new QRawFont(*static_cast<QRawFont*>(other));
}

void* QRawFont_NewQRawFont2(char* fileName, double pixelSize, int hintingPreference){
	return new QRawFont(QString(fileName), static_cast<qreal>(pixelSize), static_cast<QFont::HintingPreference>(hintingPreference));
}

double QRawFont_Ascent(void* ptr){
	return static_cast<double>(static_cast<QRawFont*>(ptr)->ascent());
}

double QRawFont_AverageCharWidth(void* ptr){
	return static_cast<double>(static_cast<QRawFont*>(ptr)->averageCharWidth());
}

double QRawFont_Descent(void* ptr){
	return static_cast<double>(static_cast<QRawFont*>(ptr)->descent());
}

char* QRawFont_FamilyName(void* ptr){
	return static_cast<QRawFont*>(ptr)->familyName().toUtf8().data();
}

void* QRawFont_FontTable(void* ptr, char* tagName){
	return new QByteArray(static_cast<QRawFont*>(ptr)->fontTable(const_cast<const char*>(tagName)));
}

int QRawFont_HintingPreference(void* ptr){
	return static_cast<QRawFont*>(ptr)->hintingPreference();
}

int QRawFont_IsValid(void* ptr){
	return static_cast<QRawFont*>(ptr)->isValid();
}

double QRawFont_Leading(void* ptr){
	return static_cast<double>(static_cast<QRawFont*>(ptr)->leading());
}

double QRawFont_LineThickness(void* ptr){
	return static_cast<double>(static_cast<QRawFont*>(ptr)->lineThickness());
}

void QRawFont_LoadFromData(void* ptr, void* fontData, double pixelSize, int hintingPreference){
	static_cast<QRawFont*>(ptr)->loadFromData(*static_cast<QByteArray*>(fontData), static_cast<qreal>(pixelSize), static_cast<QFont::HintingPreference>(hintingPreference));
}

void QRawFont_LoadFromFile(void* ptr, char* fileName, double pixelSize, int hintingPreference){
	static_cast<QRawFont*>(ptr)->loadFromFile(QString(fileName), static_cast<qreal>(pixelSize), static_cast<QFont::HintingPreference>(hintingPreference));
}

double QRawFont_MaxCharWidth(void* ptr){
	return static_cast<double>(static_cast<QRawFont*>(ptr)->maxCharWidth());
}

double QRawFont_PixelSize(void* ptr){
	return static_cast<double>(static_cast<QRawFont*>(ptr)->pixelSize());
}

void QRawFont_SetPixelSize(void* ptr, double pixelSize){
	static_cast<QRawFont*>(ptr)->setPixelSize(static_cast<qreal>(pixelSize));
}

int QRawFont_Style(void* ptr){
	return static_cast<QRawFont*>(ptr)->style();
}

char* QRawFont_StyleName(void* ptr){
	return static_cast<QRawFont*>(ptr)->styleName().toUtf8().data();
}

int QRawFont_SupportsCharacter(void* ptr, void* character){
	return static_cast<QRawFont*>(ptr)->supportsCharacter(*static_cast<QChar*>(character));
}

void QRawFont_Swap(void* ptr, void* other){
	static_cast<QRawFont*>(ptr)->swap(*static_cast<QRawFont*>(other));
}

double QRawFont_UnderlinePosition(void* ptr){
	return static_cast<double>(static_cast<QRawFont*>(ptr)->underlinePosition());
}

double QRawFont_UnitsPerEm(void* ptr){
	return static_cast<double>(static_cast<QRawFont*>(ptr)->unitsPerEm());
}

int QRawFont_Weight(void* ptr){
	return static_cast<QRawFont*>(ptr)->weight();
}

double QRawFont_XHeight(void* ptr){
	return static_cast<double>(static_cast<QRawFont*>(ptr)->xHeight());
}

void QRawFont_DestroyQRawFont(void* ptr){
	static_cast<QRawFont*>(ptr)->~QRawFont();
}

