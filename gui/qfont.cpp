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

char* QFont_DefaultFamily(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->defaultFamily().toUtf8().data();
}

char* QFont_LastResortFamily(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->lastResortFamily().toUtf8().data();
}

char* QFont_LastResortFont(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->lastResortFont().toUtf8().data();
}

QtObjectPtr QFont_NewQFont(){
	return new QFont();
}

QtObjectPtr QFont_NewQFont4(QtObjectPtr font){
	return new QFont(*static_cast<QFont*>(font));
}

QtObjectPtr QFont_NewQFont3(QtObjectPtr font, QtObjectPtr pd){
	return new QFont(*static_cast<QFont*>(font), static_cast<QPaintDevice*>(pd));
}

QtObjectPtr QFont_NewQFont2(char* family, int pointSize, int weight, int italic){
	return new QFont(QString(family), pointSize, weight, italic != 0);
}

int QFont_Bold(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->bold();
}

int QFont_Capitalization(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->capitalization();
}

int QFont_ExactMatch(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->exactMatch();
}

char* QFont_Family(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->family().toUtf8().data();
}

int QFont_FixedPitch(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->fixedPitch();
}

int QFont_FromString(QtObjectPtr ptr, char* descrip){
	return static_cast<QFont*>(ptr)->fromString(QString(descrip));
}

int QFont_HintingPreference(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->hintingPreference();
}

void QFont_QFont_InsertSubstitution(char* familyName, char* substituteName){
	QFont::insertSubstitution(QString(familyName), QString(substituteName));
}

void QFont_QFont_InsertSubstitutions(char* familyName, char* substituteNames){
	QFont::insertSubstitutions(QString(familyName), QString(substituteNames).split("|", QString::SkipEmptyParts));
}

int QFont_IsCopyOf(QtObjectPtr ptr, QtObjectPtr f){
	return static_cast<QFont*>(ptr)->isCopyOf(*static_cast<QFont*>(f));
}

int QFont_Italic(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->italic();
}

int QFont_Kerning(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->kerning();
}

char* QFont_Key(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->key().toUtf8().data();
}

int QFont_LetterSpacingType(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->letterSpacingType();
}

int QFont_Overline(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->overline();
}

int QFont_PixelSize(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->pixelSize();
}

int QFont_PointSize(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->pointSize();
}

void QFont_QFont_RemoveSubstitutions(char* familyName){
	QFont::removeSubstitutions(QString(familyName));
}

void QFont_SetBold(QtObjectPtr ptr, int enable){
	static_cast<QFont*>(ptr)->setBold(enable != 0);
}

void QFont_SetCapitalization(QtObjectPtr ptr, int caps){
	static_cast<QFont*>(ptr)->setCapitalization(static_cast<QFont::Capitalization>(caps));
}

void QFont_SetFamily(QtObjectPtr ptr, char* family){
	static_cast<QFont*>(ptr)->setFamily(QString(family));
}

void QFont_SetFixedPitch(QtObjectPtr ptr, int enable){
	static_cast<QFont*>(ptr)->setFixedPitch(enable != 0);
}

void QFont_SetHintingPreference(QtObjectPtr ptr, int hintingPreference){
	static_cast<QFont*>(ptr)->setHintingPreference(static_cast<QFont::HintingPreference>(hintingPreference));
}

void QFont_SetItalic(QtObjectPtr ptr, int enable){
	static_cast<QFont*>(ptr)->setItalic(enable != 0);
}

void QFont_SetKerning(QtObjectPtr ptr, int enable){
	static_cast<QFont*>(ptr)->setKerning(enable != 0);
}

void QFont_SetOverline(QtObjectPtr ptr, int enable){
	static_cast<QFont*>(ptr)->setOverline(enable != 0);
}

void QFont_SetPixelSize(QtObjectPtr ptr, int pixelSize){
	static_cast<QFont*>(ptr)->setPixelSize(pixelSize);
}

void QFont_SetPointSize(QtObjectPtr ptr, int pointSize){
	static_cast<QFont*>(ptr)->setPointSize(pointSize);
}

void QFont_SetStretch(QtObjectPtr ptr, int factor){
	static_cast<QFont*>(ptr)->setStretch(factor);
}

void QFont_SetStrikeOut(QtObjectPtr ptr, int enable){
	static_cast<QFont*>(ptr)->setStrikeOut(enable != 0);
}

void QFont_SetStyle(QtObjectPtr ptr, int style){
	static_cast<QFont*>(ptr)->setStyle(static_cast<QFont::Style>(style));
}

void QFont_SetStyleHint(QtObjectPtr ptr, int hint, int strategy){
	static_cast<QFont*>(ptr)->setStyleHint(static_cast<QFont::StyleHint>(hint), static_cast<QFont::StyleStrategy>(strategy));
}

void QFont_SetStyleName(QtObjectPtr ptr, char* styleName){
	static_cast<QFont*>(ptr)->setStyleName(QString(styleName));
}

void QFont_SetStyleStrategy(QtObjectPtr ptr, int s){
	static_cast<QFont*>(ptr)->setStyleStrategy(static_cast<QFont::StyleStrategy>(s));
}

void QFont_SetUnderline(QtObjectPtr ptr, int enable){
	static_cast<QFont*>(ptr)->setUnderline(enable != 0);
}

void QFont_SetWeight(QtObjectPtr ptr, int weight){
	static_cast<QFont*>(ptr)->setWeight(weight);
}

int QFont_Stretch(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->stretch();
}

int QFont_StrikeOut(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->strikeOut();
}

int QFont_Style(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->style();
}

int QFont_StyleHint(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->styleHint();
}

char* QFont_StyleName(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->styleName().toUtf8().data();
}

int QFont_StyleStrategy(QtObjectPtr ptr){
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

void QFont_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QFont*>(ptr)->swap(*static_cast<QFont*>(other));
}

char* QFont_ToString(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->toString().toUtf8().data();
}

int QFont_Underline(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->underline();
}

int QFont_Weight(QtObjectPtr ptr){
	return static_cast<QFont*>(ptr)->weight();
}

void QFont_DestroyQFont(QtObjectPtr ptr){
	static_cast<QFont*>(ptr)->~QFont();
}

