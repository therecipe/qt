#include "qdialog.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QMetaObject>
#include <QDial>
#include <QObject>
#include <QDialog>
#include "_cgo_export.h"

class MyQDialog: public QDialog {
public:
void Signal_Accepted(){callbackQDialogAccepted(this->objectName().toUtf8().data());};
void Signal_Finished(int result){callbackQDialogFinished(this->objectName().toUtf8().data(), result);};
void Signal_Rejected(){callbackQDialogRejected(this->objectName().toUtf8().data());};
};

int QDialog_IsSizeGripEnabled(QtObjectPtr ptr){
	return static_cast<QDialog*>(ptr)->isSizeGripEnabled();
}

void QDialog_SetModal(QtObjectPtr ptr, int modal){
	static_cast<QDialog*>(ptr)->setModal(modal != 0);
}

void QDialog_SetResult(QtObjectPtr ptr, int i){
	static_cast<QDialog*>(ptr)->setResult(i);
}

void QDialog_SetSizeGripEnabled(QtObjectPtr ptr, int v){
	static_cast<QDialog*>(ptr)->setSizeGripEnabled(v != 0);
}

QtObjectPtr QDialog_NewQDialog(QtObjectPtr parent, int f){
	return new QDialog(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void QDialog_Accept(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QDialog*>(ptr), "accept");
}

void QDialog_ConnectAccepted(QtObjectPtr ptr){
	QObject::connect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)()>(&QDialog::accepted), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)()>(&MyQDialog::Signal_Accepted));;
}

void QDialog_DisconnectAccepted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)()>(&QDialog::accepted), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)()>(&MyQDialog::Signal_Accepted));;
}

void QDialog_Done(QtObjectPtr ptr, int r){
	QMetaObject::invokeMethod(static_cast<QDialog*>(ptr), "done", Q_ARG(int, r));
}

int QDialog_Exec(QtObjectPtr ptr){
	return QMetaObject::invokeMethod(static_cast<QDialog*>(ptr), "exec");
}

void QDialog_ConnectFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)(int)>(&QDialog::finished), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)(int)>(&MyQDialog::Signal_Finished));;
}

void QDialog_DisconnectFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)(int)>(&QDialog::finished), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)(int)>(&MyQDialog::Signal_Finished));;
}

void QDialog_Open(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QDialog*>(ptr), "open");
}

void QDialog_Reject(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QDialog*>(ptr), "reject");
}

void QDialog_ConnectRejected(QtObjectPtr ptr){
	QObject::connect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)()>(&QDialog::rejected), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)()>(&MyQDialog::Signal_Rejected));;
}

void QDialog_DisconnectRejected(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QDialog*>(ptr), static_cast<void (QDialog::*)()>(&QDialog::rejected), static_cast<MyQDialog*>(ptr), static_cast<void (MyQDialog::*)()>(&MyQDialog::Signal_Rejected));;
}

int QDialog_Result(QtObjectPtr ptr){
	return static_cast<QDialog*>(ptr)->result();
}

void QDialog_SetVisible(QtObjectPtr ptr, int visible){
	static_cast<QDialog*>(ptr)->setVisible(visible != 0);
}

void QDialog_DestroyQDialog(QtObjectPtr ptr){
	static_cast<QDialog*>(ptr)->~QDialog();
}

