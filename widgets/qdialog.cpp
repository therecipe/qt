#include "qdialog.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QWidget>
#include <QObject>
#include <QDial>
#include <QDialog>
#include "_cgo_export.h"

class MyQDialog: public QDialog {
public:
void Signal_Accepted(){callbackQDialogAccepted(this->objectName().toUtf8().data());};
void Signal_Finished(int result){callbackQDialogFinished(this->objectName().toUtf8().data(), result);};
void Signal_Rejected(){callbackQDialogRejected(this->objectName().toUtf8().data());};
};

int QDialog_IsSizeGripEnabled(void* ptr){
	return static_cast<QDialog*>(ptr)->isSizeGripEnabled();
}

void QDialog_SetModal(void* ptr, int modal){
	static_cast<QDialog*>(ptr)->setModal(modal != 0);
}

void QDialog_SetResult(void* ptr, int i){
	static_cast<QDialog*>(ptr)->setResult(i);
}

void QDialog_SetSizeGripEnabled(void* ptr, int v){
	static_cast<QDialog*>(ptr)->setSizeGripEnabled(v != 0);
}

void* QDialog_NewQDialog(void* parent, int f){
	return new QDialog(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void QDialog_Accept(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDialog*>(ptr), "accept");
}

void QDialog_ConnectAccepted(void* ptr){
	QObject::connect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)()>(&QDialog::accepted), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)()>(&MyQDialog::Signal_Accepted));;
}

void QDialog_DisconnectAccepted(void* ptr){
	QObject::disconnect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)()>(&QDialog::accepted), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)()>(&MyQDialog::Signal_Accepted));;
}

void QDialog_Done(void* ptr, int r){
	QMetaObject::invokeMethod(static_cast<QDialog*>(ptr), "done", Q_ARG(int, r));
}

int QDialog_Exec(void* ptr){
	return QMetaObject::invokeMethod(static_cast<QDialog*>(ptr), "exec");
}

void QDialog_ConnectFinished(void* ptr){
	QObject::connect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)(int)>(&QDialog::finished), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)(int)>(&MyQDialog::Signal_Finished));;
}

void QDialog_DisconnectFinished(void* ptr){
	QObject::disconnect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)(int)>(&QDialog::finished), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)(int)>(&MyQDialog::Signal_Finished));;
}

void QDialog_Open(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDialog*>(ptr), "open");
}

void QDialog_Reject(void* ptr){
	QMetaObject::invokeMethod(static_cast<QDialog*>(ptr), "reject");
}

void QDialog_ConnectRejected(void* ptr){
	QObject::connect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)()>(&QDialog::rejected), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)()>(&MyQDialog::Signal_Rejected));;
}

void QDialog_DisconnectRejected(void* ptr){
	QObject::disconnect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)()>(&QDialog::rejected), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)()>(&MyQDialog::Signal_Rejected));;
}

int QDialog_Result(void* ptr){
	return static_cast<QDialog*>(ptr)->result();
}

void QDialog_SetVisible(void* ptr, int visible){
	static_cast<QDialog*>(ptr)->setVisible(visible != 0);
}

void QDialog_DestroyQDialog(void* ptr){
	static_cast<QDialog*>(ptr)->~QDialog();
}

