#include "qkeysequenceedit.h"
#include <QWidget>
#include <QObject>
#include <QKeySequence>
#include <QMetaObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QKeySequenceEdit>
#include "_cgo_export.h"

class MyQKeySequenceEdit: public QKeySequenceEdit {
public:
void Signal_EditingFinished(){callbackQKeySequenceEditEditingFinished(this->objectName().toUtf8().data());};
};

void QKeySequenceEdit_SetKeySequence(void* ptr, void* keySequence){
	QMetaObject::invokeMethod(static_cast<QKeySequenceEdit*>(ptr), "setKeySequence", Q_ARG(QKeySequence, *static_cast<QKeySequence*>(keySequence)));
}

void* QKeySequenceEdit_NewQKeySequenceEdit(void* parent){
	return new QKeySequenceEdit(static_cast<QWidget*>(parent));
}

void* QKeySequenceEdit_NewQKeySequenceEdit2(void* keySequence, void* parent){
	return new QKeySequenceEdit(*static_cast<QKeySequence*>(keySequence), static_cast<QWidget*>(parent));
}

void QKeySequenceEdit_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QKeySequenceEdit*>(ptr), "clear");
}

void QKeySequenceEdit_ConnectEditingFinished(void* ptr){
	QObject::connect(static_cast<QKeySequenceEdit*>(ptr), static_cast<void (QKeySequenceEdit::*)()>(&QKeySequenceEdit::editingFinished), static_cast<MyQKeySequenceEdit*>(ptr), static_cast<void (MyQKeySequenceEdit::*)()>(&MyQKeySequenceEdit::Signal_EditingFinished));;
}

void QKeySequenceEdit_DisconnectEditingFinished(void* ptr){
	QObject::disconnect(static_cast<QKeySequenceEdit*>(ptr), static_cast<void (QKeySequenceEdit::*)()>(&QKeySequenceEdit::editingFinished), static_cast<MyQKeySequenceEdit*>(ptr), static_cast<void (MyQKeySequenceEdit::*)()>(&MyQKeySequenceEdit::Signal_EditingFinished));;
}

void QKeySequenceEdit_DestroyQKeySequenceEdit(void* ptr){
	static_cast<QKeySequenceEdit*>(ptr)->~QKeySequenceEdit();
}

