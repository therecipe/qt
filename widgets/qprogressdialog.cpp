#include "qprogressdialog.h"
#include <QUrl>
#include <QModelIndex>
#include <QPushButton>
#include <QMetaObject>
#include <QLabel>
#include <QVariant>
#include <QProgressBar>
#include <QWidget>
#include <QObject>
#include <QString>
#include <QProgressDialog>
#include "_cgo_export.h"

class MyQProgressDialog: public QProgressDialog {
public:
void Signal_Canceled(){callbackQProgressDialogCanceled(this->objectName().toUtf8().data());};
};

int QProgressDialog_AutoClose(void* ptr){
	return static_cast<QProgressDialog*>(ptr)->autoClose();
}

int QProgressDialog_AutoReset(void* ptr){
	return static_cast<QProgressDialog*>(ptr)->autoReset();
}

char* QProgressDialog_LabelText(void* ptr){
	return static_cast<QProgressDialog*>(ptr)->labelText().toUtf8().data();
}

int QProgressDialog_Maximum(void* ptr){
	return static_cast<QProgressDialog*>(ptr)->maximum();
}

int QProgressDialog_Minimum(void* ptr){
	return static_cast<QProgressDialog*>(ptr)->minimum();
}

int QProgressDialog_MinimumDuration(void* ptr){
	return static_cast<QProgressDialog*>(ptr)->minimumDuration();
}

void QProgressDialog_SetAutoClose(void* ptr, int close){
	static_cast<QProgressDialog*>(ptr)->setAutoClose(close != 0);
}

void QProgressDialog_SetAutoReset(void* ptr, int reset){
	static_cast<QProgressDialog*>(ptr)->setAutoReset(reset != 0);
}

void QProgressDialog_SetLabelText(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setLabelText", Q_ARG(QString, QString(text)));
}

void QProgressDialog_SetMaximum(void* ptr, int maximum){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setMaximum", Q_ARG(int, maximum));
}

void QProgressDialog_SetMinimum(void* ptr, int minimum){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setMinimum", Q_ARG(int, minimum));
}

void QProgressDialog_SetMinimumDuration(void* ptr, int ms){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setMinimumDuration", Q_ARG(int, ms));
}

void QProgressDialog_SetValue(void* ptr, int progress){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setValue", Q_ARG(int, progress));
}

int QProgressDialog_Value(void* ptr){
	return static_cast<QProgressDialog*>(ptr)->value();
}

int QProgressDialog_WasCanceled(void* ptr){
	return static_cast<QProgressDialog*>(ptr)->wasCanceled();
}

void* QProgressDialog_NewQProgressDialog(void* parent, int f){
	return new QProgressDialog(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void* QProgressDialog_NewQProgressDialog2(char* labelText, char* cancelButtonText, int minimum, int maximum, void* parent, int f){
	return new QProgressDialog(QString(labelText), QString(cancelButtonText), minimum, maximum, static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(f));
}

void QProgressDialog_Cancel(void* ptr){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "cancel");
}

void QProgressDialog_ConnectCanceled(void* ptr){
	QObject::connect(static_cast<QProgressDialog*>(ptr), static_cast<void (QProgressDialog::*)()>(&QProgressDialog::canceled), static_cast<MyQProgressDialog*>(ptr), static_cast<void (MyQProgressDialog::*)()>(&MyQProgressDialog::Signal_Canceled));;
}

void QProgressDialog_DisconnectCanceled(void* ptr){
	QObject::disconnect(static_cast<QProgressDialog*>(ptr), static_cast<void (QProgressDialog::*)()>(&QProgressDialog::canceled), static_cast<MyQProgressDialog*>(ptr), static_cast<void (MyQProgressDialog::*)()>(&MyQProgressDialog::Signal_Canceled));;
}

void QProgressDialog_Open(void* ptr, void* receiver, char* member){
	static_cast<QProgressDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void QProgressDialog_Reset(void* ptr){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "reset");
}

void QProgressDialog_SetBar(void* ptr, void* bar){
	static_cast<QProgressDialog*>(ptr)->setBar(static_cast<QProgressBar*>(bar));
}

void QProgressDialog_SetCancelButton(void* ptr, void* cancelButton){
	static_cast<QProgressDialog*>(ptr)->setCancelButton(static_cast<QPushButton*>(cancelButton));
}

void QProgressDialog_SetCancelButtonText(void* ptr, char* cancelButtonText){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setCancelButtonText", Q_ARG(QString, QString(cancelButtonText)));
}

void QProgressDialog_SetLabel(void* ptr, void* label){
	static_cast<QProgressDialog*>(ptr)->setLabel(static_cast<QLabel*>(label));
}

void QProgressDialog_SetRange(void* ptr, int minimum, int maximum){
	QMetaObject::invokeMethod(static_cast<QProgressDialog*>(ptr), "setRange", Q_ARG(int, minimum), Q_ARG(int, maximum));
}

void QProgressDialog_DestroyQProgressDialog(void* ptr){
	static_cast<QProgressDialog*>(ptr)->~QProgressDialog();
}

