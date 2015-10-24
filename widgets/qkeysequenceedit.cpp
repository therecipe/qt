#include "qkeysequenceedit.h"
#include <QKeySequence>
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QKeySequenceEdit>
#include "_cgo_export.h"

class MyQKeySequenceEdit: public QKeySequenceEdit {
public:
void Signal_EditingFinished(){callbackQKeySequenceEditEditingFinished(this->objectName().toUtf8().data());};
};

void QKeySequenceEdit_SetKeySequence(QtObjectPtr ptr, QtObjectPtr keySequence){
	QMetaObject::invokeMethod(static_cast<QKeySequenceEdit*>(ptr), "setKeySequence", Q_ARG(QKeySequence, *static_cast<QKeySequence*>(keySequence)));
}

QtObjectPtr QKeySequenceEdit_NewQKeySequenceEdit(QtObjectPtr parent){
	return new QKeySequenceEdit(static_cast<QWidget*>(parent));
}

QtObjectPtr QKeySequenceEdit_NewQKeySequenceEdit2(QtObjectPtr keySequence, QtObjectPtr parent){
	return new QKeySequenceEdit(*static_cast<QKeySequence*>(keySequence), static_cast<QWidget*>(parent));
}

void QKeySequenceEdit_Clear(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QKeySequenceEdit*>(ptr), "clear");
}

void QKeySequenceEdit_ConnectEditingFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QKeySequenceEdit*>(ptr), static_cast<void (QKeySequenceEdit::*)()>(&QKeySequenceEdit::editingFinished), static_cast<MyQKeySequenceEdit*>(ptr), static_cast<void (MyQKeySequenceEdit::*)()>(&MyQKeySequenceEdit::Signal_EditingFinished));;
}

void QKeySequenceEdit_DisconnectEditingFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QKeySequenceEdit*>(ptr), static_cast<void (QKeySequenceEdit::*)()>(&QKeySequenceEdit::editingFinished), static_cast<MyQKeySequenceEdit*>(ptr), static_cast<void (MyQKeySequenceEdit::*)()>(&MyQKeySequenceEdit::Signal_EditingFinished));;
}

void QKeySequenceEdit_DestroyQKeySequenceEdit(QtObjectPtr ptr){
	static_cast<QKeySequenceEdit*>(ptr)->~QKeySequenceEdit();
}

