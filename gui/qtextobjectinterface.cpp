#include "qtextobjectinterface.h"
#include <QUrl>
#include <QTextFormat>
#include <QRectF>
#include <QPainter>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QTextDocument>
#include <QRect>
#include <QTextObject>
#include <QTextObjectInterface>
#include "_cgo_export.h"

class MyQTextObjectInterface: public QTextObjectInterface {
public:
};

void QTextObjectInterface_DrawObject(void* ptr, void* painter, void* rect, void* doc, int posInDocument, void* format){
	static_cast<QTextObjectInterface*>(ptr)->drawObject(static_cast<QPainter*>(painter), *static_cast<QRectF*>(rect), static_cast<QTextDocument*>(doc), posInDocument, *static_cast<QTextFormat*>(format));
}

void QTextObjectInterface_DestroyQTextObjectInterface(void* ptr){
	static_cast<QTextObjectInterface*>(ptr)->~QTextObjectInterface();
}

