#include "qstatictext.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextOption>
#include <QFont>
#include <QTransform>
#include <QString>
#include <QStaticText>
#include "_cgo_export.h"

class MyQStaticText: public QStaticText {
public:
};

QtObjectPtr QStaticText_NewQStaticText(){
	return new QStaticText();
}

QtObjectPtr QStaticText_NewQStaticText3(QtObjectPtr other){
	return new QStaticText(*static_cast<QStaticText*>(other));
}

QtObjectPtr QStaticText_NewQStaticText2(char* text){
	return new QStaticText(QString(text));
}

int QStaticText_PerformanceHint(QtObjectPtr ptr){
	return static_cast<QStaticText*>(ptr)->performanceHint();
}

void QStaticText_Prepare(QtObjectPtr ptr, QtObjectPtr matrix, QtObjectPtr font){
	static_cast<QStaticText*>(ptr)->prepare(*static_cast<QTransform*>(matrix), *static_cast<QFont*>(font));
}

void QStaticText_SetPerformanceHint(QtObjectPtr ptr, int performanceHint){
	static_cast<QStaticText*>(ptr)->setPerformanceHint(static_cast<QStaticText::PerformanceHint>(performanceHint));
}

void QStaticText_SetText(QtObjectPtr ptr, char* text){
	static_cast<QStaticText*>(ptr)->setText(QString(text));
}

void QStaticText_SetTextFormat(QtObjectPtr ptr, int textFormat){
	static_cast<QStaticText*>(ptr)->setTextFormat(static_cast<Qt::TextFormat>(textFormat));
}

void QStaticText_SetTextOption(QtObjectPtr ptr, QtObjectPtr textOption){
	static_cast<QStaticText*>(ptr)->setTextOption(*static_cast<QTextOption*>(textOption));
}

void QStaticText_Swap(QtObjectPtr ptr, QtObjectPtr other){
	static_cast<QStaticText*>(ptr)->swap(*static_cast<QStaticText*>(other));
}

char* QStaticText_Text(QtObjectPtr ptr){
	return static_cast<QStaticText*>(ptr)->text().toUtf8().data();
}

int QStaticText_TextFormat(QtObjectPtr ptr){
	return static_cast<QStaticText*>(ptr)->textFormat();
}

void QStaticText_DestroyQStaticText(QtObjectPtr ptr){
	static_cast<QStaticText*>(ptr)->~QStaticText();
}

