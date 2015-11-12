#include "qtextframe.h"
#include <QModelIndex>
#include <QTextDocument>
#include <QTextFrameFormat>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QTextFrame>
#include "_cgo_export.h"

class MyQTextFrame: public QTextFrame {
public:
};

void* QTextFrame_NewQTextFrame(void* document){
	return new QTextFrame(static_cast<QTextDocument*>(document));
}

int QTextFrame_FirstPosition(void* ptr){
	return static_cast<QTextFrame*>(ptr)->firstPosition();
}

int QTextFrame_LastPosition(void* ptr){
	return static_cast<QTextFrame*>(ptr)->lastPosition();
}

void* QTextFrame_ParentFrame(void* ptr){
	return static_cast<QTextFrame*>(ptr)->parentFrame();
}

void QTextFrame_SetFrameFormat(void* ptr, void* format){
	static_cast<QTextFrame*>(ptr)->setFrameFormat(*static_cast<QTextFrameFormat*>(format));
}

void QTextFrame_DestroyQTextFrame(void* ptr){
	static_cast<QTextFrame*>(ptr)->~QTextFrame();
}

