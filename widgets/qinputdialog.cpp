#include "qinputdialog.h"
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QLine>
#include <QLineEdit>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QInputDialog>
#include "_cgo_export.h"

class MyQInputDialog: public QInputDialog {
public:
void Signal_IntValueChanged(int value){callbackQInputDialogIntValueChanged(this->objectName().toUtf8().data(), value);};
void Signal_IntValueSelected(int value){callbackQInputDialogIntValueSelected(this->objectName().toUtf8().data(), value);};
void Signal_TextValueChanged(const QString & text){callbackQInputDialogTextValueChanged(this->objectName().toUtf8().data(), text.toUtf8().data());};
void Signal_TextValueSelected(const QString & text){callbackQInputDialogTextValueSelected(this->objectName().toUtf8().data(), text.toUtf8().data());};
};

char* QInputDialog_CancelButtonText(QtObjectPtr ptr){
	return static_cast<QInputDialog*>(ptr)->cancelButtonText().toUtf8().data();
}

char* QInputDialog_ComboBoxItems(QtObjectPtr ptr){
	return static_cast<QInputDialog*>(ptr)->comboBoxItems().join("|").toUtf8().data();
}

int QInputDialog_DoubleDecimals(QtObjectPtr ptr){
	return static_cast<QInputDialog*>(ptr)->doubleDecimals();
}

int QInputDialog_InputMode(QtObjectPtr ptr){
	return static_cast<QInputDialog*>(ptr)->inputMode();
}

int QInputDialog_IntMaximum(QtObjectPtr ptr){
	return static_cast<QInputDialog*>(ptr)->intMaximum();
}

int QInputDialog_IntMinimum(QtObjectPtr ptr){
	return static_cast<QInputDialog*>(ptr)->intMinimum();
}

int QInputDialog_IntStep(QtObjectPtr ptr){
	return static_cast<QInputDialog*>(ptr)->intStep();
}

int QInputDialog_IntValue(QtObjectPtr ptr){
	return static_cast<QInputDialog*>(ptr)->intValue();
}

int QInputDialog_IsComboBoxEditable(QtObjectPtr ptr){
	return static_cast<QInputDialog*>(ptr)->isComboBoxEditable();
}

char* QInputDialog_LabelText(QtObjectPtr ptr){
	return static_cast<QInputDialog*>(ptr)->labelText().toUtf8().data();
}

char* QInputDialog_OkButtonText(QtObjectPtr ptr){
	return static_cast<QInputDialog*>(ptr)->okButtonText().toUtf8().data();
}

int QInputDialog_Options(QtObjectPtr ptr){
	return static_cast<QInputDialog*>(ptr)->options();
}

void QInputDialog_SetCancelButtonText(QtObjectPtr ptr, char* text){
	static_cast<QInputDialog*>(ptr)->setCancelButtonText(QString(text));
}

void QInputDialog_SetComboBoxEditable(QtObjectPtr ptr, int editable){
	static_cast<QInputDialog*>(ptr)->setComboBoxEditable(editable != 0);
}

void QInputDialog_SetComboBoxItems(QtObjectPtr ptr, char* items){
	static_cast<QInputDialog*>(ptr)->setComboBoxItems(QString(items).split("|", QString::SkipEmptyParts));
}

void QInputDialog_SetDoubleDecimals(QtObjectPtr ptr, int decimals){
	static_cast<QInputDialog*>(ptr)->setDoubleDecimals(decimals);
}

void QInputDialog_SetInputMode(QtObjectPtr ptr, int mode){
	static_cast<QInputDialog*>(ptr)->setInputMode(static_cast<QInputDialog::InputMode>(mode));
}

void QInputDialog_SetIntMaximum(QtObjectPtr ptr, int max){
	static_cast<QInputDialog*>(ptr)->setIntMaximum(max);
}

void QInputDialog_SetIntMinimum(QtObjectPtr ptr, int min){
	static_cast<QInputDialog*>(ptr)->setIntMinimum(min);
}

void QInputDialog_SetIntStep(QtObjectPtr ptr, int step){
	static_cast<QInputDialog*>(ptr)->setIntStep(step);
}

void QInputDialog_SetIntValue(QtObjectPtr ptr, int value){
	static_cast<QInputDialog*>(ptr)->setIntValue(value);
}

void QInputDialog_SetLabelText(QtObjectPtr ptr, char* text){
	static_cast<QInputDialog*>(ptr)->setLabelText(QString(text));
}

void QInputDialog_SetOkButtonText(QtObjectPtr ptr, char* text){
	static_cast<QInputDialog*>(ptr)->setOkButtonText(QString(text));
}

void QInputDialog_SetOptions(QtObjectPtr ptr, int options){
	static_cast<QInputDialog*>(ptr)->setOptions(static_cast<QInputDialog::InputDialogOption>(options));
}

void QInputDialog_SetTextEchoMode(QtObjectPtr ptr, int mode){
	static_cast<QInputDialog*>(ptr)->setTextEchoMode(static_cast<QLineEdit::EchoMode>(mode));
}

void QInputDialog_SetTextValue(QtObjectPtr ptr, char* text){
	static_cast<QInputDialog*>(ptr)->setTextValue(QString(text));
}

int QInputDialog_TextEchoMode(QtObjectPtr ptr){
	return static_cast<QInputDialog*>(ptr)->textEchoMode();
}

char* QInputDialog_TextValue(QtObjectPtr ptr){
	return static_cast<QInputDialog*>(ptr)->textValue().toUtf8().data();
}

QtObjectPtr QInputDialog_NewQInputDialog(QtObjectPtr parent, int flags){
	return new QInputDialog(static_cast<QWidget*>(parent), static_cast<Qt::WindowType>(flags));
}

void QInputDialog_Done(QtObjectPtr ptr, int result){
	static_cast<QInputDialog*>(ptr)->done(result);
}

int QInputDialog_QInputDialog_GetInt(QtObjectPtr parent, char* title, char* label, int value, int min, int max, int step, int ok, int flags){
	return QInputDialog::getInt(static_cast<QWidget*>(parent), QString(title), QString(label), value, min, max, step, NULL, static_cast<Qt::WindowType>(flags));
}

char* QInputDialog_QInputDialog_GetItem(QtObjectPtr parent, char* title, char* label, char* items, int current, int editable, int ok, int flags, int inputMethodHints){
	return QInputDialog::getItem(static_cast<QWidget*>(parent), QString(title), QString(label), QString(items).split("|", QString::SkipEmptyParts), current, editable != 0, NULL, static_cast<Qt::WindowType>(flags), static_cast<Qt::InputMethodHint>(inputMethodHints)).toUtf8().data();
}

char* QInputDialog_QInputDialog_GetMultiLineText(QtObjectPtr parent, char* title, char* label, char* text, int ok, int flags, int inputMethodHints){
	return QInputDialog::getMultiLineText(static_cast<QWidget*>(parent), QString(title), QString(label), QString(text), NULL, static_cast<Qt::WindowType>(flags), static_cast<Qt::InputMethodHint>(inputMethodHints)).toUtf8().data();
}

char* QInputDialog_QInputDialog_GetText(QtObjectPtr parent, char* title, char* label, int mode, char* text, int ok, int flags, int inputMethodHints){
	return QInputDialog::getText(static_cast<QWidget*>(parent), QString(title), QString(label), static_cast<QLineEdit::EchoMode>(mode), QString(text), NULL, static_cast<Qt::WindowType>(flags), static_cast<Qt::InputMethodHint>(inputMethodHints)).toUtf8().data();
}

void QInputDialog_ConnectIntValueChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QInputDialog*>(ptr), static_cast<void (QInputDialog::*)(int)>(&QInputDialog::intValueChanged), static_cast<MyQInputDialog*>(ptr), static_cast<void (MyQInputDialog::*)(int)>(&MyQInputDialog::Signal_IntValueChanged));;
}

void QInputDialog_DisconnectIntValueChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QInputDialog*>(ptr), static_cast<void (QInputDialog::*)(int)>(&QInputDialog::intValueChanged), static_cast<MyQInputDialog*>(ptr), static_cast<void (MyQInputDialog::*)(int)>(&MyQInputDialog::Signal_IntValueChanged));;
}

void QInputDialog_ConnectIntValueSelected(QtObjectPtr ptr){
	QObject::connect(static_cast<QInputDialog*>(ptr), static_cast<void (QInputDialog::*)(int)>(&QInputDialog::intValueSelected), static_cast<MyQInputDialog*>(ptr), static_cast<void (MyQInputDialog::*)(int)>(&MyQInputDialog::Signal_IntValueSelected));;
}

void QInputDialog_DisconnectIntValueSelected(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QInputDialog*>(ptr), static_cast<void (QInputDialog::*)(int)>(&QInputDialog::intValueSelected), static_cast<MyQInputDialog*>(ptr), static_cast<void (MyQInputDialog::*)(int)>(&MyQInputDialog::Signal_IntValueSelected));;
}

void QInputDialog_Open(QtObjectPtr ptr, QtObjectPtr receiver, char* member){
	static_cast<QInputDialog*>(ptr)->open(static_cast<QObject*>(receiver), const_cast<const char*>(member));
}

void QInputDialog_SetIntRange(QtObjectPtr ptr, int min, int max){
	static_cast<QInputDialog*>(ptr)->setIntRange(min, max);
}

void QInputDialog_SetOption(QtObjectPtr ptr, int option, int on){
	static_cast<QInputDialog*>(ptr)->setOption(static_cast<QInputDialog::InputDialogOption>(option), on != 0);
}

void QInputDialog_SetVisible(QtObjectPtr ptr, int visible){
	static_cast<QInputDialog*>(ptr)->setVisible(visible != 0);
}

int QInputDialog_TestOption(QtObjectPtr ptr, int option){
	return static_cast<QInputDialog*>(ptr)->testOption(static_cast<QInputDialog::InputDialogOption>(option));
}

void QInputDialog_ConnectTextValueChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QInputDialog*>(ptr), static_cast<void (QInputDialog::*)(const QString &)>(&QInputDialog::textValueChanged), static_cast<MyQInputDialog*>(ptr), static_cast<void (MyQInputDialog::*)(const QString &)>(&MyQInputDialog::Signal_TextValueChanged));;
}

void QInputDialog_DisconnectTextValueChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QInputDialog*>(ptr), static_cast<void (QInputDialog::*)(const QString &)>(&QInputDialog::textValueChanged), static_cast<MyQInputDialog*>(ptr), static_cast<void (MyQInputDialog::*)(const QString &)>(&MyQInputDialog::Signal_TextValueChanged));;
}

void QInputDialog_ConnectTextValueSelected(QtObjectPtr ptr){
	QObject::connect(static_cast<QInputDialog*>(ptr), static_cast<void (QInputDialog::*)(const QString &)>(&QInputDialog::textValueSelected), static_cast<MyQInputDialog*>(ptr), static_cast<void (MyQInputDialog::*)(const QString &)>(&MyQInputDialog::Signal_TextValueSelected));;
}

void QInputDialog_DisconnectTextValueSelected(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QInputDialog*>(ptr), static_cast<void (QInputDialog::*)(const QString &)>(&QInputDialog::textValueSelected), static_cast<MyQInputDialog*>(ptr), static_cast<void (MyQInputDialog::*)(const QString &)>(&MyQInputDialog::Signal_TextValueSelected));;
}

void QInputDialog_DestroyQInputDialog(QtObjectPtr ptr){
	static_cast<QInputDialog*>(ptr)->~QInputDialog();
}

