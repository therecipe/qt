#include "qtextobjectinterface.h"
#include <QTextObject>
#include <QTextFormat>
#include <QPainter>
#include <QString>
#include <QVariant>
#include <QModelIndex>
#include <QTextDocument>
#include <QUrl>
#include <QRect>
#include <QRectF>
#include <QTextObjectInterface>
#include "_cgo_export.h"

class MyQTextObjectInterface: public QTextObjectInterface {
public:
};

void QTextObjectInterface_DrawObject(QtObjectPtr ptr, QtObjectPtr painter, QtObjectPtr rect, QtObjectPtr doc, int posInDocument, QtObjectPtr format){
	static_cast<QTextObjectInterface*>(ptr)->drawObject(static_cast<QPainter*>(painter), *static_cast<QRectF*>(rect), static_cast<QTextDocument*>(doc), posInDocument, *static_cast<QTextFormat*>(format));
}

void QTextObjectInterface_DestroyQTextObjectInterface(QtObjectPtr ptr){
	static_cast<QTextObjectInterface*>(ptr)->~QTextObjectInterface();
}

