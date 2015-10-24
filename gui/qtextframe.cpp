#include "qtextframe.h"
#include <QTextFrameFormat>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QTextDocument>
#include <QTextFrame>
#include "_cgo_export.h"

class MyQTextFrame: public QTextFrame {
public:
};

QtObjectPtr QTextFrame_NewQTextFrame(QtObjectPtr document){
	return new QTextFrame(static_cast<QTextDocument*>(document));
}

int QTextFrame_FirstPosition(QtObjectPtr ptr){
	return static_cast<QTextFrame*>(ptr)->firstPosition();
}

int QTextFrame_LastPosition(QtObjectPtr ptr){
	return static_cast<QTextFrame*>(ptr)->lastPosition();
}

QtObjectPtr QTextFrame_ParentFrame(QtObjectPtr ptr){
	return static_cast<QTextFrame*>(ptr)->parentFrame();
}

void QTextFrame_SetFrameFormat(QtObjectPtr ptr, QtObjectPtr format){
	static_cast<QTextFrame*>(ptr)->setFrameFormat(*static_cast<QTextFrameFormat*>(format));
}

void QTextFrame_DestroyQTextFrame(QtObjectPtr ptr){
	static_cast<QTextFrame*>(ptr)->~QTextFrame();
}

