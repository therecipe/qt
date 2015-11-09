#include "qfont.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPaintDevice>
#include <QFont>
#include "_cgo_export.h"

class MyQFont: public QFont {
public:
};

int QFont_Times_Type(){
	return QFont::Times;
}

int QFont_Courier_Type(){
	return QFont::Courier;
}

int QFont_OldEnglish_Type(){
	return QFont::OldEnglish;
}

int QFont_System_Type(){
	return QFont::System;
}

int QFont_AnyStyle_Type(){
	return QFont::AnyStyle;
}

int QFont_Cursive_Type(){
	return QFont::Cursive;
}

int QFont_Monospace_Type(){
	return QFont::Monospace;
}

int QFont_Fantasy_Type(){
	return QFont::Fantasy;
}

char* QFont_DefaultFamily(void* ptr){
	return static_cast<QFont*>(ptr)->defaultFamily().toUtf8().data();
}

char* QFont_LastResortFamily(void* ptr){
	return static_cast<QFont*>(ptr)->lastResortFamily().toUtf8().data();
}

char* QFont_LastResortFont(void* ptr){
	return static_cast<QFont*>(ptr)->lastResortFont().toUtf8().data();
}

void* QFont_NewQFont(){
	return new QFont();
}

void* QFont_NewQFont4(void* font){
	return new QFont(*static_cast<QFont*>(font));
}

void* QFont_NewQFont3(void* font, void* pd){
	return new QFont(*static_cast<QFont*>(font), static_cast<QPaintDevice*>(pd));
}

void* QFont_NewQFont2(char* family, int pointSize, int weight, int italic){
	return new QFont(QString(family), pointSize, weight, italic != 0);
}

int QFont_Bold(void* ptr){
	return static_cast<QFont*>(ptr)->bold();
}

int QFont_Capitalization(void* ptr){
	return static_cast<QFont*>(ptr)->capitalization();
}

int QFont_ExactMatch(void* ptr){
	return static_cast<QFont*>(ptr)->exactMatch();
}

char* QFont_Family(void* ptr){
	return static_cast<QFont*>(ptr)->family().toUtf8().data();
}

int QFont_FixedPitch(void* ptr){
	return static_cast<QFont*>(ptr)->fixedPitch();
}

int QFont_FromString(void* ptr, char* descrip){
	return static_cast<QFont*>(ptr)->fromString(QString(descrip));
}

int QFont_HintingPreference(void* ptr){
	return static_cast<QFont*>(ptr)->hintingPreference();
}

void QFont_QFont_InsertSubstitution(char* familyName, char* substituteName){
	QFont::insertSubstitution(QString(familyName), QString(substituteName));
}

void QFont_QFont_InsertSubstitutions(char* familyName, char* substituteNames){
	QFont::insertSubstitutions(QString(familyName), QString(substituteNames).split("|", QString::SkipEmptyParts));
}

int QFont_IsCopyOf(void* ptr, void* f){
	return static_cast<QFont*>(ptr)->isCopyOf(*static_cast<QFont*>(f));
}

int QFont_Italic(void* ptr){
	return static_cast<QFont*>(ptr)->italic();
}

int QFont_Kerning(void* ptr){
	return static_cast<QFont*>(ptr)->kerning();
}

char* QFont_Key(void* ptr){
	return static_cast<QFont*>(ptr)->key().toUtf8().data();
}

double QFont_LetterSpacing(void* ptr){
	return static_cast<double>(static_cast<QFont*>(ptr)->letterSpacing());
}

int QFont_LetterSpacingType(void* ptr){
	return static_cast<QFont*>(ptr)->letterSpacingType();
}

int QFont_Overline(void* ptr){
	return static_cast<QFont*>(ptr)->overline();
}

int QFont_PixelSize(void* ptr){
	return static_cast<QFont*>(ptr)->pixelSize();
}

int QFont_PointSize(void* ptr){
	return static_cast<QFont*>(ptr)->pointSize();
}

double QFont_PointSizeF(void* ptr){
	return static_cast<double>(static_cast<QFont*>(ptr)->pointSizeF());
}

void QFont_QFont_RemoveSubstitutions(char* familyName){
	QFont::removeSubstitutions(QString(familyName));
}

void QFont_SetBold(void* ptr, int enable){
	static_cast<QFont*>(ptr)->setBold(enable != 0);
}

void QFont_SetCapitalization(void* ptr, int caps){
	static_cast<QFont*>(ptr)->setCapitalization(static_cast<QFont::Capitalization>(caps));
}

void QFont_SetFamily(void* ptr, char* family){
	static_cast<QFont*>(ptr)->setFamily(QString(family));
}

void QFont_SetFixedPitch(void* ptr, int enable){
	static_cast<QFont*>(ptr)->setFixedPitch(enable != 0);
}

void QFont_SetHintingPreference(void* ptr, int hintingPreference){
	static_cast<QFont*>(ptr)->setHintingPreference(static_cast<QFont::HintingPreference>(hintingPreference));
}

void QFont_SetItalic(void* ptr, int enable){
	static_cast<QFont*>(ptr)->setItalic(enable != 0);
}

void QFont_SetKerning(void* ptr, int enable){
	static_cast<QFont*>(ptr)->setKerning(enable != 0);
}

void QFont_SetLetterSpacing(void* ptr, int ty, double spacing){
	static_cast<QFont*>(ptr)->setLetterSpacing(static_cast<QFont::SpacingType>(ty), static_cast<qreal>(spacing));
}

void QFont_SetOverline(void* ptr, int enable){
	static_cast<QFont*>(ptr)->setOverline(enable != 0);
}

void QFont_SetPixelSize(void* ptr, int pixelSize){
	static_cast<QFont*>(ptr)->setPixelSize(pixelSize);
}

void QFont_SetPointSize(void* ptr, int pointSize){
	static_cast<QFont*>(ptr)->setPointSize(pointSize);
}

void QFont_SetPointSizeF(void* ptr, double pointSize){
	static_cast<QFont*>(ptr)->setPointSizeF(static_cast<qreal>(pointSize));
}

void QFont_SetStretch(void* ptr, int factor){
	static_cast<QFont*>(ptr)->setStretch(factor);
}

void QFont_SetStrikeOut(void* ptr, int enable){
	static_cast<QFont*>(ptr)->setStrikeOut(enable != 0);
}

void QFont_SetStyle(void* ptr, int style){
	static_cast<QFont*>(ptr)->setStyle(static_cast<QFont::Style>(style));
}

void QFont_SetStyleHint(void* ptr, int hint, int strategy){
	static_cast<QFont*>(ptr)->setStyleHint(static_cast<QFont::StyleHint>(hint), static_cast<QFont::StyleStrategy>(strategy));
}

void QFont_SetStyleName(void* ptr, char* styleName){
	static_cast<QFont*>(ptr)->setStyleName(QString(styleName));
}

void QFont_SetStyleStrategy(void* ptr, int s){
	static_cast<QFont*>(ptr)->setStyleStrategy(static_cast<QFont::StyleStrategy>(s));
}

void QFont_SetUnderline(void* ptr, int enable){
	static_cast<QFont*>(ptr)->setUnderline(enable != 0);
}

void QFont_SetWeight(void* ptr, int weight){
	static_cast<QFont*>(ptr)->setWeight(weight);
}

void QFont_SetWordSpacing(void* ptr, double spacing){
	static_cast<QFont*>(ptr)->setWordSpacing(static_cast<qreal>(spacing));
}

int QFont_Stretch(void* ptr){
	return static_cast<QFont*>(ptr)->stretch();
}

int QFont_StrikeOut(void* ptr){
	return static_cast<QFont*>(ptr)->strikeOut();
}

int QFont_Style(void* ptr){
	return static_cast<QFont*>(ptr)->style();
}

int QFont_StyleHint(void* ptr){
	return static_cast<QFont*>(ptr)->styleHint();
}

char* QFont_StyleName(void* ptr){
	return static_cast<QFont*>(ptr)->styleName().toUtf8().data();
}

int QFont_StyleStrategy(void* ptr){
	return static_cast<QFont*>(ptr)->styleStrategy();
}

char* QFont_QFont_Substitute(char* familyName){
	return QFont::substitute(QString(familyName)).toUtf8().data();
}

char* QFont_QFont_Substitutes(char* familyName){
	return QFont::substitutes(QString(familyName)).join("|").toUtf8().data();
}

char* QFont_QFont_Substitutions(){
	return QFont::substitutions().join("|").toUtf8().data();
}

void QFont_Swap(void* ptr, void* other){
	static_cast<QFont*>(ptr)->swap(*static_cast<QFont*>(other));
}

char* QFont_ToString(void* ptr){
	return static_cast<QFont*>(ptr)->toString().toUtf8().data();
}

int QFont_Underline(void* ptr){
	return static_cast<QFont*>(ptr)->underline();
}

int QFont_Weight(void* ptr){
	return static_cast<QFont*>(ptr)->weight();
}

double QFont_WordSpacing(void* ptr){
	return static_cast<double>(static_cast<QFont*>(ptr)->wordSpacing());
}

void QFont_DestroyQFont(void* ptr){
	static_cast<QFont*>(ptr)->~QFont();
}

