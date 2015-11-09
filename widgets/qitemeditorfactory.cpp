#include "qitemeditorfactory.h"
#include <QUrl>
#include <QModelIndex>
#include <QItemEditorCreator>
#include <QByteArray>
#include <QItemEditorCreatorBase>
#include <QWidget>
#include <QString>
#include <QVariant>
#include <QItemEditorFactory>
#include "_cgo_export.h"

class MyQItemEditorFactory: public QItemEditorFactory {
public:
};

void* QItemEditorFactory_NewQItemEditorFactory(){
	return new QItemEditorFactory();
}

void* QItemEditorFactory_CreateEditor(void* ptr, int userType, void* parent){
	return static_cast<QItemEditorFactory*>(ptr)->createEditor(userType, static_cast<QWidget*>(parent));
}

void* QItemEditorFactory_QItemEditorFactory_DefaultFactory(){
	return const_cast<QItemEditorFactory*>(QItemEditorFactory::defaultFactory());
}

void QItemEditorFactory_RegisterEditor(void* ptr, int userType, void* creator){
	static_cast<QItemEditorFactory*>(ptr)->registerEditor(userType, static_cast<QItemEditorCreatorBase*>(creator));
}

void QItemEditorFactory_QItemEditorFactory_SetDefaultFactory(void* factory){
	QItemEditorFactory::setDefaultFactory(static_cast<QItemEditorFactory*>(factory));
}

void* QItemEditorFactory_ValuePropertyName(void* ptr, int userType){
	return new QByteArray(static_cast<QItemEditorFactory*>(ptr)->valuePropertyName(userType));
}

void QItemEditorFactory_DestroyQItemEditorFactory(void* ptr){
	static_cast<QItemEditorFactory*>(ptr)->~QItemEditorFactory();
}

