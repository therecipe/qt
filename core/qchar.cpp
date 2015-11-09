#include "qchar.h"
#include <QLatin1Char>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QChar>
#include "_cgo_export.h"

class MyQChar: public QChar {
public:
};

void* QChar_NewQChar(){
	return new QChar();
}

void* QChar_NewQChar8(void* ch){
	return new QChar(*static_cast<QLatin1Char*>(ch));
}

void* QChar_NewQChar7(int ch){
	return new QChar(static_cast<QChar::SpecialCharacter>(ch));
}

void* QChar_NewQChar9(char* ch){
	return new QChar(*ch);
}

void* QChar_NewQChar6(int code){
	return new QChar(code);
}

int QChar_Category(void* ptr){
	return static_cast<QChar*>(ptr)->category();
}

int QChar_QChar_CurrentUnicodeVersion(){
	return QChar::currentUnicodeVersion();
}

char* QChar_Decomposition(void* ptr){
	return static_cast<QChar*>(ptr)->decomposition().toUtf8().data();
}

int QChar_DecompositionTag(void* ptr){
	return static_cast<QChar*>(ptr)->decompositionTag();
}

int QChar_DigitValue(void* ptr){
	return static_cast<QChar*>(ptr)->digitValue();
}

int QChar_Direction(void* ptr){
	return static_cast<QChar*>(ptr)->direction();
}

int QChar_HasMirrored(void* ptr){
	return static_cast<QChar*>(ptr)->hasMirrored();
}

int QChar_IsDigit(void* ptr){
	return static_cast<QChar*>(ptr)->isDigit();
}

int QChar_IsHighSurrogate(void* ptr){
	return static_cast<QChar*>(ptr)->isHighSurrogate();
}

int QChar_IsLetter(void* ptr){
	return static_cast<QChar*>(ptr)->isLetter();
}

int QChar_IsLetterOrNumber(void* ptr){
	return static_cast<QChar*>(ptr)->isLetterOrNumber();
}

int QChar_IsLower(void* ptr){
	return static_cast<QChar*>(ptr)->isLower();
}

int QChar_IsLowSurrogate(void* ptr){
	return static_cast<QChar*>(ptr)->isLowSurrogate();
}

int QChar_IsMark(void* ptr){
	return static_cast<QChar*>(ptr)->isMark();
}

int QChar_IsNonCharacter(void* ptr){
	return static_cast<QChar*>(ptr)->isNonCharacter();
}

int QChar_IsNull(void* ptr){
	return static_cast<QChar*>(ptr)->isNull();
}

int QChar_IsNumber(void* ptr){
	return static_cast<QChar*>(ptr)->isNumber();
}

int QChar_IsPrint(void* ptr){
	return static_cast<QChar*>(ptr)->isPrint();
}

int QChar_IsPunct(void* ptr){
	return static_cast<QChar*>(ptr)->isPunct();
}

int QChar_IsSpace(void* ptr){
	return static_cast<QChar*>(ptr)->isSpace();
}

int QChar_IsSurrogate(void* ptr){
	return static_cast<QChar*>(ptr)->isSurrogate();
}

int QChar_IsSymbol(void* ptr){
	return static_cast<QChar*>(ptr)->isSymbol();
}

int QChar_IsTitleCase(void* ptr){
	return static_cast<QChar*>(ptr)->isTitleCase();
}

int QChar_IsUpper(void* ptr){
	return static_cast<QChar*>(ptr)->isUpper();
}

int QChar_JoiningType(void* ptr){
	return static_cast<QChar*>(ptr)->joiningType();
}

int QChar_Script(void* ptr){
	return static_cast<QChar*>(ptr)->script();
}

int QChar_UnicodeVersion(void* ptr){
	return static_cast<QChar*>(ptr)->unicodeVersion();
}

