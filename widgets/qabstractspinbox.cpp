#include "qabstractspinbox.h"
#include <QVariant>
#include <QUrl>
#include <QModelIndex>
#include <QMetaObject>
#include <QWidget>
#include <QObject>
#include <QEvent>
#include <QString>
#include <QAbstractSpinBox>
#include "_cgo_export.h"

class MyQAbstractSpinBox: public QAbstractSpinBox {
public:
void Signal_EditingFinished(){callbackQAbstractSpinBoxEditingFinished(this->objectName().toUtf8().data());};
};

int QAbstractSpinBox_Alignment(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->alignment();
}

int QAbstractSpinBox_ButtonSymbols(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->buttonSymbols();
}

int QAbstractSpinBox_CorrectionMode(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->correctionMode();
}

int QAbstractSpinBox_HasAcceptableInput(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->hasAcceptableInput();
}

int QAbstractSpinBox_HasFrame(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->hasFrame();
}

int QAbstractSpinBox_IsAccelerated(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->isAccelerated();
}

int QAbstractSpinBox_IsGroupSeparatorShown(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->isGroupSeparatorShown();
}

int QAbstractSpinBox_IsReadOnly(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->isReadOnly();
}

int QAbstractSpinBox_KeyboardTracking(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->keyboardTracking();
}

void QAbstractSpinBox_SetAccelerated(void* ptr, int on){
	static_cast<QAbstractSpinBox*>(ptr)->setAccelerated(on != 0);
}

void QAbstractSpinBox_SetAlignment(void* ptr, int flag){
	static_cast<QAbstractSpinBox*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(flag));
}

void QAbstractSpinBox_SetButtonSymbols(void* ptr, int bs){
	static_cast<QAbstractSpinBox*>(ptr)->setButtonSymbols(static_cast<QAbstractSpinBox::ButtonSymbols>(bs));
}

void QAbstractSpinBox_SetCorrectionMode(void* ptr, int cm){
	static_cast<QAbstractSpinBox*>(ptr)->setCorrectionMode(static_cast<QAbstractSpinBox::CorrectionMode>(cm));
}

void QAbstractSpinBox_SetFrame(void* ptr, int v){
	static_cast<QAbstractSpinBox*>(ptr)->setFrame(v != 0);
}

void QAbstractSpinBox_SetGroupSeparatorShown(void* ptr, int shown){
	static_cast<QAbstractSpinBox*>(ptr)->setGroupSeparatorShown(shown != 0);
}

void QAbstractSpinBox_SetKeyboardTracking(void* ptr, int kt){
	static_cast<QAbstractSpinBox*>(ptr)->setKeyboardTracking(kt != 0);
}

void QAbstractSpinBox_SetReadOnly(void* ptr, int r){
	static_cast<QAbstractSpinBox*>(ptr)->setReadOnly(r != 0);
}

void QAbstractSpinBox_SetSpecialValueText(void* ptr, char* txt){
	static_cast<QAbstractSpinBox*>(ptr)->setSpecialValueText(QString(txt));
}

void QAbstractSpinBox_SetWrapping(void* ptr, int w){
	static_cast<QAbstractSpinBox*>(ptr)->setWrapping(w != 0);
}

char* QAbstractSpinBox_SpecialValueText(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->specialValueText().toUtf8().data();
}

char* QAbstractSpinBox_Text(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->text().toUtf8().data();
}

int QAbstractSpinBox_Wrapping(void* ptr){
	return static_cast<QAbstractSpinBox*>(ptr)->wrapping();
}

void* QAbstractSpinBox_NewQAbstractSpinBox(void* parent){
	return new QAbstractSpinBox(static_cast<QWidget*>(parent));
}

void QAbstractSpinBox_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractSpinBox*>(ptr), "clear");
}

void QAbstractSpinBox_ConnectEditingFinished(void* ptr){
	QObject::connect(static_cast<QAbstractSpinBox*>(ptr), static_cast<void (QAbstractSpinBox::*)()>(&QAbstractSpinBox::editingFinished), static_cast<MyQAbstractSpinBox*>(ptr), static_cast<void (MyQAbstractSpinBox::*)()>(&MyQAbstractSpinBox::Signal_EditingFinished));;
}

void QAbstractSpinBox_DisconnectEditingFinished(void* ptr){
	QObject::disconnect(static_cast<QAbstractSpinBox*>(ptr), static_cast<void (QAbstractSpinBox::*)()>(&QAbstractSpinBox::editingFinished), static_cast<MyQAbstractSpinBox*>(ptr), static_cast<void (MyQAbstractSpinBox::*)()>(&MyQAbstractSpinBox::Signal_EditingFinished));;
}

int QAbstractSpinBox_Event(void* ptr, void* event){
	return static_cast<QAbstractSpinBox*>(ptr)->event(static_cast<QEvent*>(event));
}

void* QAbstractSpinBox_InputMethodQuery(void* ptr, int query){
	return new QVariant(static_cast<QAbstractSpinBox*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QAbstractSpinBox_InterpretText(void* ptr){
	static_cast<QAbstractSpinBox*>(ptr)->interpretText();
}

void QAbstractSpinBox_SelectAll(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractSpinBox*>(ptr), "selectAll");
}

void QAbstractSpinBox_StepBy(void* ptr, int steps){
	static_cast<QAbstractSpinBox*>(ptr)->stepBy(steps);
}

void QAbstractSpinBox_StepDown(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractSpinBox*>(ptr), "stepDown");
}

void QAbstractSpinBox_StepUp(void* ptr){
	QMetaObject::invokeMethod(static_cast<QAbstractSpinBox*>(ptr), "stepUp");
}

void QAbstractSpinBox_DestroyQAbstractSpinBox(void* ptr){
	static_cast<QAbstractSpinBox*>(ptr)->~QAbstractSpinBox();
}

