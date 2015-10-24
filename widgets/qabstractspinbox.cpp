#include "qabstractspinbox.h"
#include <QMetaObject>
#include <QObject>
#include <QString>
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QWidget>
#include <QEvent>
#include <QAbstractSpinBox>
#include "_cgo_export.h"

class MyQAbstractSpinBox: public QAbstractSpinBox {
public:
void Signal_EditingFinished(){callbackQAbstractSpinBoxEditingFinished(this->objectName().toUtf8().data());};
};

int QAbstractSpinBox_Alignment(QtObjectPtr ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->alignment();
}

int QAbstractSpinBox_ButtonSymbols(QtObjectPtr ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->buttonSymbols();
}

int QAbstractSpinBox_CorrectionMode(QtObjectPtr ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->correctionMode();
}

int QAbstractSpinBox_HasAcceptableInput(QtObjectPtr ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->hasAcceptableInput();
}

int QAbstractSpinBox_HasFrame(QtObjectPtr ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->hasFrame();
}

int QAbstractSpinBox_IsAccelerated(QtObjectPtr ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->isAccelerated();
}

int QAbstractSpinBox_IsGroupSeparatorShown(QtObjectPtr ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->isGroupSeparatorShown();
}

int QAbstractSpinBox_IsReadOnly(QtObjectPtr ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->isReadOnly();
}

int QAbstractSpinBox_KeyboardTracking(QtObjectPtr ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->keyboardTracking();
}

void QAbstractSpinBox_SetAccelerated(QtObjectPtr ptr, int on){
	static_cast<QAbstractSpinBox*>(ptr)->setAccelerated(on != 0);
}

void QAbstractSpinBox_SetAlignment(QtObjectPtr ptr, int flag){
	static_cast<QAbstractSpinBox*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(flag));
}

void QAbstractSpinBox_SetButtonSymbols(QtObjectPtr ptr, int bs){
	static_cast<QAbstractSpinBox*>(ptr)->setButtonSymbols(static_cast<QAbstractSpinBox::ButtonSymbols>(bs));
}

void QAbstractSpinBox_SetCorrectionMode(QtObjectPtr ptr, int cm){
	static_cast<QAbstractSpinBox*>(ptr)->setCorrectionMode(static_cast<QAbstractSpinBox::CorrectionMode>(cm));
}

void QAbstractSpinBox_SetFrame(QtObjectPtr ptr, int v){
	static_cast<QAbstractSpinBox*>(ptr)->setFrame(v != 0);
}

void QAbstractSpinBox_SetGroupSeparatorShown(QtObjectPtr ptr, int shown){
	static_cast<QAbstractSpinBox*>(ptr)->setGroupSeparatorShown(shown != 0);
}

void QAbstractSpinBox_SetKeyboardTracking(QtObjectPtr ptr, int kt){
	static_cast<QAbstractSpinBox*>(ptr)->setKeyboardTracking(kt != 0);
}

void QAbstractSpinBox_SetReadOnly(QtObjectPtr ptr, int r){
	static_cast<QAbstractSpinBox*>(ptr)->setReadOnly(r != 0);
}

void QAbstractSpinBox_SetSpecialValueText(QtObjectPtr ptr, char* txt){
	static_cast<QAbstractSpinBox*>(ptr)->setSpecialValueText(QString(txt));
}

void QAbstractSpinBox_SetWrapping(QtObjectPtr ptr, int w){
	static_cast<QAbstractSpinBox*>(ptr)->setWrapping(w != 0);
}

char* QAbstractSpinBox_SpecialValueText(QtObjectPtr ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->specialValueText().toUtf8().data();
}

char* QAbstractSpinBox_Text(QtObjectPtr ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->text().toUtf8().data();
}

int QAbstractSpinBox_Wrapping(QtObjectPtr ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->wrapping();
}

QtObjectPtr QAbstractSpinBox_NewQAbstractSpinBox(QtObjectPtr parent){
	return new QAbstractSpinBox(static_cast<QWidget*>(parent));
}

void QAbstractSpinBox_Clear(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractSpinBox*>(ptr), "clear");
}

void QAbstractSpinBox_ConnectEditingFinished(QtObjectPtr ptr){
	QObject::connect(static_cast<QAbstractSpinBox*>(ptr), static_cast<void (QAbstractSpinBox::*)()>(&QAbstractSpinBox::editingFinished), static_cast<MyQAbstractSpinBox*>(ptr), static_cast<void (MyQAbstractSpinBox::*)()>(&MyQAbstractSpinBox::Signal_EditingFinished));;
}

void QAbstractSpinBox_DisconnectEditingFinished(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QAbstractSpinBox*>(ptr), static_cast<void (QAbstractSpinBox::*)()>(&QAbstractSpinBox::editingFinished), static_cast<MyQAbstractSpinBox*>(ptr), static_cast<void (MyQAbstractSpinBox::*)()>(&MyQAbstractSpinBox::Signal_EditingFinished));;
}

int QAbstractSpinBox_Event(QtObjectPtr ptr, QtObjectPtr event){
	return static_cast<QAbstractSpinBox*>(ptr)->event(static_cast<QEvent*>(event));
}

char* QAbstractSpinBox_InputMethodQuery(QtObjectPtr ptr, int query){
	return static_cast<QAbstractSpinBox*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)).toString().toUtf8().data();
}

void QAbstractSpinBox_InterpretText(QtObjectPtr ptr){
	static_cast<QAbstractSpinBox*>(ptr)->interpretText();
}

void QAbstractSpinBox_SelectAll(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractSpinBox*>(ptr), "selectAll");
}

void QAbstractSpinBox_StepBy(QtObjectPtr ptr, int steps){
	static_cast<QAbstractSpinBox*>(ptr)->stepBy(steps);
}

void QAbstractSpinBox_StepDown(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractSpinBox*>(ptr), "stepDown");
}

void QAbstractSpinBox_StepUp(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractSpinBox*>(ptr), "stepUp");
}

void QAbstractSpinBox_DestroyQAbstractSpinBox(QtObjectPtr ptr){
	static_cast<QAbstractSpinBox*>(ptr)->~QAbstractSpinBox();
}

