#include "qtextcharformat.h"
#include <QColor>
#include <QFont>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QPen>
#include <QTextCharFormat>
#include "_cgo_export.h"

class MyQTextCharFormat: public QTextCharFormat {
public:
};

void* QTextCharFormat_NewQTextCharFormat(){
	return new QTextCharFormat();
}

char* QTextCharFormat_AnchorNames(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->anchorNames().join("|").toUtf8().data();
}

int QTextCharFormat_FontUnderline(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontUnderline();
}

void QTextCharFormat_SetUnderlineStyle(void* ptr, int style){
	static_cast<QTextCharFormat*>(ptr)->setUnderlineStyle(static_cast<QTextCharFormat::UnderlineStyle>(style));
}

char* QTextCharFormat_AnchorHref(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->anchorHref().toUtf8().data();
}

int QTextCharFormat_FontCapitalization(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontCapitalization();
}

char* QTextCharFormat_FontFamily(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontFamily().toUtf8().data();
}

int QTextCharFormat_FontFixedPitch(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontFixedPitch();
}

int QTextCharFormat_FontHintingPreference(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontHintingPreference();
}

int QTextCharFormat_FontItalic(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontItalic();
}

int QTextCharFormat_FontKerning(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontKerning();
}

double QTextCharFormat_FontLetterSpacing(void* ptr){
	return static_cast<double>(static_cast<QTextCharFormat*>(ptr)->fontLetterSpacing());
}

int QTextCharFormat_FontLetterSpacingType(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontLetterSpacingType();
}

int QTextCharFormat_FontOverline(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontOverline();
}

double QTextCharFormat_FontPointSize(void* ptr){
	return static_cast<double>(static_cast<QTextCharFormat*>(ptr)->fontPointSize());
}

int QTextCharFormat_FontStretch(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontStretch();
}

int QTextCharFormat_FontStrikeOut(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontStrikeOut();
}

int QTextCharFormat_FontStyleHint(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontStyleHint();
}

int QTextCharFormat_FontStyleStrategy(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontStyleStrategy();
}

int QTextCharFormat_FontWeight(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->fontWeight();
}

double QTextCharFormat_FontWordSpacing(void* ptr){
	return static_cast<double>(static_cast<QTextCharFormat*>(ptr)->fontWordSpacing());
}

int QTextCharFormat_IsAnchor(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->isAnchor();
}

int QTextCharFormat_IsValid(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->isValid();
}

void QTextCharFormat_SetAnchor(void* ptr, int anchor){
	static_cast<QTextCharFormat*>(ptr)->setAnchor(anchor != 0);
}

void QTextCharFormat_SetAnchorHref(void* ptr, char* value){
	static_cast<QTextCharFormat*>(ptr)->setAnchorHref(QString(value));
}

void QTextCharFormat_SetAnchorNames(void* ptr, char* names){
	static_cast<QTextCharFormat*>(ptr)->setAnchorNames(QString(names).split("|", QString::SkipEmptyParts));
}

void QTextCharFormat_SetFont2(void* ptr, void* font){
	static_cast<QTextCharFormat*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QTextCharFormat_SetFont(void* ptr, void* font, int behavior){
	static_cast<QTextCharFormat*>(ptr)->setFont(*static_cast<QFont*>(font), static_cast<QTextCharFormat::FontPropertiesInheritanceBehavior>(behavior));
}

void QTextCharFormat_SetFontCapitalization(void* ptr, int capitalization){
	static_cast<QTextCharFormat*>(ptr)->setFontCapitalization(static_cast<QFont::Capitalization>(capitalization));
}

void QTextCharFormat_SetFontFamily(void* ptr, char* family){
	static_cast<QTextCharFormat*>(ptr)->setFontFamily(QString(family));
}

void QTextCharFormat_SetFontFixedPitch(void* ptr, int fixedPitch){
	static_cast<QTextCharFormat*>(ptr)->setFontFixedPitch(fixedPitch != 0);
}

void QTextCharFormat_SetFontHintingPreference(void* ptr, int hintingPreference){
	static_cast<QTextCharFormat*>(ptr)->setFontHintingPreference(static_cast<QFont::HintingPreference>(hintingPreference));
}

void QTextCharFormat_SetFontItalic(void* ptr, int italic){
	static_cast<QTextCharFormat*>(ptr)->setFontItalic(italic != 0);
}

void QTextCharFormat_SetFontKerning(void* ptr, int enable){
	static_cast<QTextCharFormat*>(ptr)->setFontKerning(enable != 0);
}

void QTextCharFormat_SetFontLetterSpacing(void* ptr, double spacing){
	static_cast<QTextCharFormat*>(ptr)->setFontLetterSpacing(static_cast<qreal>(spacing));
}

void QTextCharFormat_SetFontLetterSpacingType(void* ptr, int letterSpacingType){
	static_cast<QTextCharFormat*>(ptr)->setFontLetterSpacingType(static_cast<QFont::SpacingType>(letterSpacingType));
}

void QTextCharFormat_SetFontOverline(void* ptr, int overline){
	static_cast<QTextCharFormat*>(ptr)->setFontOverline(overline != 0);
}

void QTextCharFormat_SetFontPointSize(void* ptr, double size){
	static_cast<QTextCharFormat*>(ptr)->setFontPointSize(static_cast<qreal>(size));
}

void QTextCharFormat_SetFontStretch(void* ptr, int factor){
	static_cast<QTextCharFormat*>(ptr)->setFontStretch(factor);
}

void QTextCharFormat_SetFontStrikeOut(void* ptr, int strikeOut){
	static_cast<QTextCharFormat*>(ptr)->setFontStrikeOut(strikeOut != 0);
}

void QTextCharFormat_SetFontStyleHint(void* ptr, int hint, int strategy){
	static_cast<QTextCharFormat*>(ptr)->setFontStyleHint(static_cast<QFont::StyleHint>(hint), static_cast<QFont::StyleStrategy>(strategy));
}

void QTextCharFormat_SetFontStyleStrategy(void* ptr, int strategy){
	static_cast<QTextCharFormat*>(ptr)->setFontStyleStrategy(static_cast<QFont::StyleStrategy>(strategy));
}

void QTextCharFormat_SetFontUnderline(void* ptr, int underline){
	static_cast<QTextCharFormat*>(ptr)->setFontUnderline(underline != 0);
}

void QTextCharFormat_SetFontWeight(void* ptr, int weight){
	static_cast<QTextCharFormat*>(ptr)->setFontWeight(weight);
}

void QTextCharFormat_SetFontWordSpacing(void* ptr, double spacing){
	static_cast<QTextCharFormat*>(ptr)->setFontWordSpacing(static_cast<qreal>(spacing));
}

void QTextCharFormat_SetTextOutline(void* ptr, void* pen){
	static_cast<QTextCharFormat*>(ptr)->setTextOutline(*static_cast<QPen*>(pen));
}

void QTextCharFormat_SetToolTip(void* ptr, char* text){
	static_cast<QTextCharFormat*>(ptr)->setToolTip(QString(text));
}

void QTextCharFormat_SetUnderlineColor(void* ptr, void* color){
	static_cast<QTextCharFormat*>(ptr)->setUnderlineColor(*static_cast<QColor*>(color));
}

void QTextCharFormat_SetVerticalAlignment(void* ptr, int alignment){
	static_cast<QTextCharFormat*>(ptr)->setVerticalAlignment(static_cast<QTextCharFormat::VerticalAlignment>(alignment));
}

char* QTextCharFormat_ToolTip(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->toolTip().toUtf8().data();
}

void* QTextCharFormat_UnderlineColor(void* ptr){
	return new QColor(static_cast<QTextCharFormat*>(ptr)->underlineColor());
}

int QTextCharFormat_UnderlineStyle(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->underlineStyle();
}

int QTextCharFormat_VerticalAlignment(void* ptr){
	return static_cast<QTextCharFormat*>(ptr)->verticalAlignment();
}

