#include "qitemeditorfactory.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QItemEditorCreatorBase>
#include <QWidget>
#include <QItemEditorCreator>
#include <QString>
#include <QItemEditorFactory>
#include "_cgo_export.h"

class MyQItemEditorFactory: public QItemEditorFactory {
public:
};

QtObjectPtr QItemEditorFactory_NewQItemEditorFactory(){
	return new QItemEditorFactory();
}

QtObjectPtr QItemEditorFactory_CreateEditor(QtObjectPtr ptr, int userType, QtObjectPtr parent){
	return static_cast<QItemEditorFactory*>(ptr)->createEditor(userType, static_cast<QWidget*>(parent));
}

QtObjectPtr QItemEditorFactory_QItemEditorFactory_DefaultFactory(){
	return const_cast<QItemEditorFactory*>(QItemEditorFactory::defaultFactory());
}

void QItemEditorFactory_RegisterEditor(QtObjectPtr ptr, int userType, QtObjectPtr creator){
	static_cast<QItemEditorFactory*>(ptr)->registerEditor(userType, static_cast<QItemEditorCreatorBase*>(creator));
}

void QItemEditorFactory_QItemEditorFactory_SetDefaultFactory(QtObjectPtr factory){
	QItemEditorFactory::setDefaultFactory(static_cast<QItemEditorFactory*>(factory));
}

void QItemEditorFactory_DestroyQItemEditorFactory(QtObjectPtr ptr){
	static_cast<QItemEditorFactory*>(ptr)->~QItemEditorFactory();
}

