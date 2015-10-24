#include "qprogressdialog.h"
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QLabel>
#include <QPushButton>
#include <QObject>
#include <QProgressBar>
#include <QWidget>
#include <QProgressDialog>
#include "_cgo_export.h"

class MyQProgressDialog: public QProgressDialog {
public:
void Signal_Canceled(){callbackQProgressDialogCanceled(this->objectName().toUtf8().data());};
};

int QProgressDialog_AutoClose(QtObjectPtr ptr){
	return static_cast<QProgressDialog*>(ptr)->autoClose();
}

int QProgressDialog_AutoReset(QtObjectPtr ptr){
	return static_cast<QProgressDialog*>(ptr)->autoReset();
}

char* QProgressDialog_LabelText(QtObjectPtr ptr){
	return static_cast<QProgressDialog*>(ptr)->labelText().toUtf8().data();
}

int QProgressDialog_Maximum(QtObjectPtr ptr){
	return static_cast<QProgressDialog*>(ptr)->maximum();
}

int QProgressDialog_Minimum(QtObjectPtr ptr){
	return static_cast<QProgressDialog*>(ptr)->minimum();
}

int QProgressDialog_MinimumDuration(QtObjectPtr ptr){
	return static_cast<QProgressDialog*>(ptr)->minimumDuration();
}

void QProgressDialog_SetAutoClose(QtObjectPtr ptr, int close){
	static_cast<QProgressDialog*>(ptr)->setAutoClose(close != 0);
}

void QProgressDialog_SetAutoReset(QtObjectPtr ptr, int reset){
	static_cast<QProgressDialog*>(ptr)->setAutoReset(reset != 0);
}

void QProgressDialog_SetLabelText(QtObjectPtr ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setLabelText", Q_ARG(QString, QString(text)));
}

void QProgressDialog_SetMaximum(QtObjectPtr ptr, int maximum){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setMaximum", Q_ARG(int, maximum));
}

void QProgressDialog_SetMinimum(QtObjectPtr ptr, int minimum){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setMinimum", Q_ARG(int, minimum));
}

void QProgressDialog_SetMinimumDuration(QtObjectPtr ptr, int ms){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setMinimumDuration", Q_ARG(int, ms));
}

void QProgressDialog_SetValue(QtObjectPtr ptr, int progress){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setValue", Q_ARG(int, progress));
}

int QProgressDialog_Value(QtObjectPtr ptr){
	return static_cast<QProgressDialog*>(ptr)->value();
}

int QProgressDialog_WasCanceled(QtObjectPtr ptr){
	return static_cast<QProgressDialog*>(ptr)->wasCanceled();
}

QtObjectPtr QProgressDialog_NewQProgressDialog(QtObjectPtr parent, int f){
	return new QProgressDialog(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

QtObjectPtr QProgressDialog_NewQProgressDialog2(char* labelText, char* cancelButtonText, int minimum, int maximum, QtObjectPtr parent, int f){
	return new QProgressDialog(QString(labelText), QString(cancelButtonText), minimum, maximum, static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void QProgressDialog_Cancel(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "cancel");
}

void QProgressDialog_ConnectCanceled(QtObjectPtr ptr){
	QObject::connect(static_cast<QProgressDialog*>(ptr), static_cast<void (QProgressDialog::*)()>(&QProgressDialog::canceled), static_cast<MyQProgressDialog*>(ptr), static_cast<void (MyQProgressDialog::*)()>(&MyQProgressDialog::Signal_Canceled));;
}

void QProgressDialog_DisconnectCanceled(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QProgressDialog*>(ptr), static_cast<void (QProgressDialog::*)()>(&QProgressDialog::canceled), static_cast<MyQProgressDialog*>(ptr), static_cast<void (MyQProgressDialog::*)()>(&MyQProgressDialog::Signal_Canceled));;
}

void QProgressDialog_Open(QtObjectPtr ptr, QtObjectPtr receiver, char* member){
	static_cast<QProgressDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void QProgressDialog_Reset(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "reset");
}

void QProgressDialog_SetBar(QtObjectPtr ptr, QtObjectPtr bar){
	static_cast<QProgressDialog*>(ptr)->setBar(static_cast<QProgressBar*>(bar));
}

void QProgressDialog_SetCancelButton(QtObjectPtr ptr, QtObjectPtr cancelButton){
	static_cast<QProgressDialog*>(ptr)->setCancelButton(static_cast<QPushButton*>(cancelButton));
}

void QProgressDialog_SetCancelButtonText(QtObjectPtr ptr, char* cancelButtonText){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setCancelButtonText", Q_ARG(QString, QString(cancelButtonText)));
}

void QProgressDialog_SetLabel(QtObjectPtr ptr, QtObjectPtr label){
	static_cast<QProgressDialog*>(ptr)->setLabel(static_cast<QLabel*>(label));
}

void QProgressDialog_SetRange(QtObjectPtr ptr, int minimum, int maximum){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setRange", Q_ARG(int, minimum), Q_ARG(int, maximum));
}

void QProgressDialog_DestroyQProgressDialog(QtObjectPtr ptr){
	static_cast<QProgressDialog*>(ptr)->~QProgressDialog();
}

