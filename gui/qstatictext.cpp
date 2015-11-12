#include "qstatictext.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextOption>
#include <QTransform>
#include <QFont>
#include <QString>
#include <QStaticText>
#include "_cgo_export.h"

class MyQStaticText: public QStaticText {
public:
};

void* QStaticText_NewQStaticText(){
	return new QStaticText();
}

void* QStaticText_NewQStaticText3(void* other){
	return new QStaticText(*static_cast<QStaticText*>(other));
}

void* QStaticText_NewQStaticText2(char* text){
	return new QStaticText(QString(text));
}

int QStaticText_PerformanceHint(void* ptr){
	return static_cast<QStaticText*>(ptr)->performanceHint();
}

void QStaticText_Prepare(void* ptr, void* matrix, void* font){
	static_cast<QStaticText*>(ptr)->prepare(*static_cast<QTransform*>(matrix), *static_cast<QFont*>(font));
}

void QStaticText_SetPerformanceHint(void* ptr, int performanceHint){
	static_cast<QStaticText*>(ptr)->setPerformanceHint(static_cast<QStaticText::PerformanceHint>(performanceHint));
}

void QStaticText_SetText(void* ptr, char* text){
	static_cast<QStaticText*>(ptr)->setText(QString(text));
}

void QStaticText_SetTextFormat(void* ptr, int textFormat){
	static_cast<QStaticText*>(ptr)->setTextFormat(static_cast<Qt::TextFormat>(textFormat));
}

void QStaticText_SetTextOption(void* ptr, void* textOption){
	static_cast<QStaticText*>(ptr)->setTextOption(*static_cast<QTextOption*>(textOption));
}

void QStaticText_SetTextWidth(void* ptr, double textWidth){
	static_cast<QStaticText*>(ptr)->setTextWidth(static_cast<qreal>(textWidth));
}

void QStaticText_Swap(void* ptr, void* other){
	static_cast<QStaticText*>(ptr)->swap(*static_cast<QStaticText*>(other));
}

char* QStaticText_Text(void* ptr){
	return static_cast<QStaticText*>(ptr)->text().toUtf8().data();
}

int QStaticText_TextFormat(void* ptr){
	return static_cast<QStaticText*>(ptr)->textFormat();
}

double QStaticText_TextWidth(void* ptr){
	return static_cast<double>(static_cast<QStaticText*>(ptr)->textWidth());
}

void QStaticText_DestroyQStaticText(void* ptr){
	static_cast<QStaticText*>(ptr)->~QStaticText();
}

