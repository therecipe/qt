#include "qitemeditorcreatorbase.h"
#include <QModelIndex>
#include <QByteArray>
#include <QWidget>
#include <QItemEditorCreator>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QItemEditorCreatorBase>
#include "_cgo_export.h"

class MyQItemEditorCreatorBase: public QItemEditorCreatorBase {
public:
};

void QItemEditorCreatorBase_DestroyQItemEditorCreatorBase(void* ptr){
	static_cast<QItemEditorCreatorBase*>(ptr)->~QItemEditorCreatorBase();
}

void* QItemEditorCreatorBase_CreateWidget(void* ptr, void* parent){
	return static_cast<QItemEditorCreatorBase*>(ptr)->createWidget(static_cast<QWidget*>(parent));
}

void* QItemEditorCreatorBase_ValuePropertyName(void* ptr){
	return new QByteArray(static_cast<QItemEditorCreatorBase*>(ptr)->valuePropertyName());
}

