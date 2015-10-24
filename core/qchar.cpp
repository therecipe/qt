#include "qchar.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QLatin1Char>
#include <QChar>
#include "_cgo_export.h"

class MyQChar: public QChar {
public:
};

QtObjectPtr QChar_NewQChar(){
	return new QChar();
}

QtObjectPtr QChar_NewQChar8(QtObjectPtr ch){
	return new QChar(*static_cast<QLatin1Char*>(ch));
}

QtObjectPtr QChar_NewQChar7(int ch){
	return new QChar(static_cast<QChar::SpecialCharacter>(ch));
}

QtObjectPtr QChar_NewQChar9(char* ch){
	return new QChar(*ch);
}

QtObjectPtr QChar_NewQChar6(int code){
	return new QChar(code);
}

int QChar_Category(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->category();
}

int QChar_QChar_CurrentUnicodeVersion(){
	return QChar::currentUnicodeVersion();
}

char* QChar_Decomposition(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->decomposition().toUtf8().data();
}

int QChar_DecompositionTag(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->decompositionTag();
}

int QChar_DigitValue(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->digitValue();
}

int QChar_Direction(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->direction();
}

int QChar_HasMirrored(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->hasMirrored();
}

int QChar_IsDigit(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->isDigit();
}

int QChar_IsHighSurrogate(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->isHighSurrogate();
}

int QChar_IsLetter(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->isLetter();
}

int QChar_IsLetterOrNumber(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->isLetterOrNumber();
}

int QChar_IsLower(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->isLower();
}

int QChar_IsLowSurrogate(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->isLowSurrogate();
}

int QChar_IsMark(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->isMark();
}

int QChar_IsNonCharacter(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->isNonCharacter();
}

int QChar_IsNull(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->isNull();
}

int QChar_IsNumber(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->isNumber();
}

int QChar_IsPrint(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->isPrint();
}

int QChar_IsPunct(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->isPunct();
}

int QChar_IsSpace(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->isSpace();
}

int QChar_IsSurrogate(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->isSurrogate();
}

int QChar_IsSymbol(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->isSymbol();
}

int QChar_IsTitleCase(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->isTitleCase();
}

int QChar_IsUpper(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->isUpper();
}

int QChar_JoiningType(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->joiningType();
}

int QChar_Script(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->script();
}

int QChar_UnicodeVersion(QtObjectPtr ptr){
	return static_cast<QChar*>(ptr)->unicodeVersion();
}

