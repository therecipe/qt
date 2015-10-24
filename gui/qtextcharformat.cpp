#include "qtextcharformat.h"
#include <QFont>
#include <QColor>
#include <QPen>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextCharFormat>
#include "_cgo_export.h"

class MyQTextCharFormat: public QTextCharFormat {
public:
};

QtObjectPtr QTextCharFormat_NewQTextCharFormat(){
	return new QTextCharFormat();
}

char* QTextCharFormat_AnchorNames(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->anchorNames().join("|").toUtf8().data();
}

int QTextCharFormat_FontUnderline(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontUnderline();
}

void QTextCharFormat_SetUnderlineStyle(QtObjectPtr ptr, int style){
	static_cast<QTextCharFormat*>(ptr)->setUnderlineStyle(static_cast<QTextCharFormat::UnderlineStyle>(style));
}

char* QTextCharFormat_AnchorHref(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->anchorHref().toUtf8().data();
}

int QTextCharFormat_FontCapitalization(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontCapitalization();
}

char* QTextCharFormat_FontFamily(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontFamily().toUtf8().data();
}

int QTextCharFormat_FontFixedPitch(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontFixedPitch();
}

int QTextCharFormat_FontHintingPreference(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontHintingPreference();
}

int QTextCharFormat_FontItalic(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontItalic();
}

int QTextCharFormat_FontKerning(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontKerning();
}

int QTextCharFormat_FontLetterSpacingType(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontLetterSpacingType();
}

int QTextCharFormat_FontOverline(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontOverline();
}

int QTextCharFormat_FontStretch(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontStretch();
}

int QTextCharFormat_FontStrikeOut(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontStrikeOut();
}

int QTextCharFormat_FontStyleHint(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontStyleHint();
}

int QTextCharFormat_FontStyleStrategy(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontStyleStrategy();
}

int QTextCharFormat_FontWeight(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontWeight();
}

int QTextCharFormat_IsAnchor(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->isAnchor();
}

int QTextCharFormat_IsValid(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->isValid();
}

void QTextCharFormat_SetAnchor(QtObjectPtr ptr, int anchor){
	static_cast<QTextCharFormat*>(ptr)->setAnchor(anchor != 0);
}

void QTextCharFormat_SetAnchorHref(QtObjectPtr ptr, char* value){
	static_cast<QTextCharFormat*>(ptr)->setAnchorHref(QString(value));
}

void QTextCharFormat_SetAnchorNames(QtObjectPtr ptr, char* names){
	static_cast<QTextCharFormat*>(ptr)->setAnchorNames(QString(names).split("|", QString::SkipEmptyParts));
}

void QTextCharFormat_SetFont2(QtObjectPtr ptr, QtObjectPtr font){
	static_cast<QTextCharFormat*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QTextCharFormat_SetFont(QtObjectPtr ptr, QtObjectPtr font, int behavior){
	static_cast<QTextCharFormat*>(ptr)->setFont(*static_cast<QFont*>(font), static_cast<QTextCharFormat::FontPropertiesInheritanceBehavior>(behavior));
}

void QTextCharFormat_SetFontCapitalization(QtObjectPtr ptr, int capitalization){
	static_cast<QTextCharFormat*>(ptr)->setFontCapitalization(static_cast<QFont::Capitalization>(capitalization));
}

void QTextCharFormat_SetFontFamily(QtObjectPtr ptr, char* family){
	static_cast<QTextCharFormat*>(ptr)->setFontFamily(QString(family));
}

void QTextCharFormat_SetFontFixedPitch(QtObjectPtr ptr, int fixedPitch){
	static_cast<QTextCharFormat*>(ptr)->setFontFixedPitch(fixedPitch != 0);
}

void QTextCharFormat_SetFontHintingPreference(QtObjectPtr ptr, int hintingPreference){
	static_cast<QTextCharFormat*>(ptr)->setFontHintingPreference(static_cast<QFont::HintingPreference>(hintingPreference));
}

void QTextCharFormat_SetFontItalic(QtObjectPtr ptr, int italic){
	static_cast<QTextCharFormat*>(ptr)->setFontItalic(italic != 0);
}

void QTextCharFormat_SetFontKerning(QtObjectPtr ptr, int enable){
	static_cast<QTextCharFormat*>(ptr)->setFontKerning(enable != 0);
}

void QTextCharFormat_SetFontLetterSpacingType(QtObjectPtr ptr, int letterSpacingType){
	static_cast<QTextCharFormat*>(ptr)->setFontLetterSpacingType(static_cast<QFont::SpacingType>(letterSpacingType));
}

void QTextCharFormat_SetFontOverline(QtObjectPtr ptr, int overline){
	static_cast<QTextCharFormat*>(ptr)->setFontOverline(overline != 0);
}

void QTextCharFormat_SetFontStretch(QtObjectPtr ptr, int factor){
	static_cast<QTextCharFormat*>(ptr)->setFontStretch(factor);
}

void QTextCharFormat_SetFontStrikeOut(QtObjectPtr ptr, int strikeOut){
	static_cast<QTextCharFormat*>(ptr)->setFontStrikeOut(strikeOut != 0);
}

void QTextCharFormat_SetFontStyleHint(QtObjectPtr ptr, int hint, int strategy){
	static_cast<QTextCharFormat*>(ptr)->setFontStyleHint(static_cast<QFont::StyleHint>(hint), static_cast<QFont::StyleStrategy>(strategy));
}

void QTextCharFormat_SetFontStyleStrategy(QtObjectPtr ptr, int strategy){
	static_cast<QTextCharFormat*>(ptr)->setFontStyleStrategy(static_cast<QFont::StyleStrategy>(strategy));
}

void QTextCharFormat_SetFontUnderline(QtObjectPtr ptr, int underline){
	static_cast<QTextCharFormat*>(ptr)->setFontUnderline(underline != 0);
}

void QTextCharFormat_SetFontWeight(QtObjectPtr ptr, int weight){
	static_cast<QTextCharFormat*>(ptr)->setFontWeight(weight);
}

void QTextCharFormat_SetTextOutline(QtObjectPtr ptr, QtObjectPtr pen){
	static_cast<QTextCharFormat*>(ptr)->setTextOutline(*static_cast<QPen*>(pen));
}

void QTextCharFormat_SetToolTip(QtObjectPtr ptr, char* text){
	static_cast<QTextCharFormat*>(ptr)->setToolTip(QString(text));
}

void QTextCharFormat_SetUnderlineColor(QtObjectPtr ptr, QtObjectPtr color){
	static_cast<QTextCharFormat*>(ptr)->setUnderlineColor(*static_cast<QColor*>(color));
}

void QTextCharFormat_SetVerticalAlignment(QtObjectPtr ptr, int alignment){
	static_cast<QTextCharFormat*>(ptr)->setVerticalAlignment(static_cast<QTextCharFormat::VerticalAlignment>(alignment));
}

char* QTextCharFormat_ToolTip(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->toolTip().toUtf8().data();
}

int QTextCharFormat_UnderlineStyle(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->underlineStyle();
}

int QTextCharFormat_VerticalAlignment(QtObjectPtr ptr){
	return static_cast<QTextCharFormat*>(ptr)->verticalAlignment();
}

