#include "qabstractspinbox.h"
#include <QAbstractSpinBox>
#include "cgoexport.h"

//MyQAbstractSpinBox
class MyQAbstractSpinBox: public QAbstractSpinBox {
Q_OBJECT
public:
void Signal_EditingFinished() { callbackQAbstractSpinBox(this, QString("editingFinished").toUtf8().data()); };

signals:
void Slot_Clear();
void Slot_SelectAll();
void Slot_StepDown();
void Slot_StepUp();

};
#include "qabstractspinbox.moc"

//Public Types
int QAbstractSpinBox_UpDownArrows() { return QAbstractSpinBox::UpDownArrows; }
int QAbstractSpinBox_PlusMinus() { return QAbstractSpinBox::PlusMinus; }
int QAbstractSpinBox_NoButtons() { return QAbstractSpinBox::NoButtons; }
int QAbstractSpinBox_CorrectToPreviousValue() { return QAbstractSpinBox::CorrectToPreviousValue; }
int QAbstractSpinBox_CorrectToNearestValue() { return QAbstractSpinBox::CorrectToNearestValue; }
int QAbstractSpinBox_StepNone() { return QAbstractSpinBox::StepNone; }
int QAbstractSpinBox_StepUpEnabled() { return QAbstractSpinBox::StepUpEnabled; }
int QAbstractSpinBox_StepDownEnabled() { return QAbstractSpinBox::StepDownEnabled; }

//Public Functions
QtObjectPtr QAbstractSpinBox_New_QWidget(QtObjectPtr parent)
{
	return new QAbstractSpinBox(((QWidget*)(parent)));
}

void QAbstractSpinBox_Destroy(QtObjectPtr ptr)
{
	((QAbstractSpinBox*)(ptr))->~QAbstractSpinBox();
}

int QAbstractSpinBox_Alignment(QtObjectPtr ptr)
{
	return ((QAbstractSpinBox*)(ptr))->alignment();
}

int QAbstractSpinBox_ButtonSymbols(QtObjectPtr ptr)
{
	return ((QAbstractSpinBox*)(ptr))->buttonSymbols();
}

int QAbstractSpinBox_HasAcceptableInput(QtObjectPtr ptr)
{
	return ((QAbstractSpinBox*)(ptr))->hasAcceptableInput();
}

int QAbstractSpinBox_HasFrame(QtObjectPtr ptr)
{
	return ((QAbstractSpinBox*)(ptr))->hasFrame();
}

void QAbstractSpinBox_InterpretText(QtObjectPtr ptr)
{
	((QAbstractSpinBox*)(ptr))->interpretText();
}

int QAbstractSpinBox_IsAccelerated(QtObjectPtr ptr)
{
	return ((QAbstractSpinBox*)(ptr))->isAccelerated();
}

int QAbstractSpinBox_IsGroupSeparatorShown(QtObjectPtr ptr)
{
	return ((QAbstractSpinBox*)(ptr))->isGroupSeparatorShown();
}

int QAbstractSpinBox_IsReadOnly(QtObjectPtr ptr)
{
	return ((QAbstractSpinBox*)(ptr))->isReadOnly();
}

int QAbstractSpinBox_KeyboardTracking(QtObjectPtr ptr)
{
	return ((QAbstractSpinBox*)(ptr))->keyboardTracking();
}

void QAbstractSpinBox_SetAccelerated_Bool(QtObjectPtr ptr, int on)
{
	((QAbstractSpinBox*)(ptr))->setAccelerated(on != 0);
}

void QAbstractSpinBox_SetAlignment_AlignmentFlag(QtObjectPtr ptr, int flag)
{
	((QAbstractSpinBox*)(ptr))->setAlignment(((Qt::AlignmentFlag)(flag)));
}

void QAbstractSpinBox_SetButtonSymbols_ButtonSymbols(QtObjectPtr ptr, int bs)
{
	((QAbstractSpinBox*)(ptr))->setButtonSymbols(((QAbstractSpinBox::ButtonSymbols)(bs)));
}

void QAbstractSpinBox_SetFrame_Bool(QtObjectPtr ptr, int frame)
{
	((QAbstractSpinBox*)(ptr))->setFrame(frame != 0);
}

void QAbstractSpinBox_SetGroupSeparatorShown_Bool(QtObjectPtr ptr, int shown)
{
	((QAbstractSpinBox*)(ptr))->setGroupSeparatorShown(shown != 0);
}

void QAbstractSpinBox_SetKeyboardTracking_Bool(QtObjectPtr ptr, int kt)
{
	((QAbstractSpinBox*)(ptr))->setKeyboardTracking(kt != 0);
}

void QAbstractSpinBox_SetReadOnly_Bool(QtObjectPtr ptr, int r)
{
	((QAbstractSpinBox*)(ptr))->setReadOnly(r != 0);
}

void QAbstractSpinBox_SetSpecialValueText_String(QtObjectPtr ptr, char* txt)
{
	((QAbstractSpinBox*)(ptr))->setSpecialValueText(QString(txt));
}

void QAbstractSpinBox_SetWrapping_Bool(QtObjectPtr ptr, int w)
{
	((QAbstractSpinBox*)(ptr))->setWrapping(w != 0);
}

char* QAbstractSpinBox_SpecialValueText(QtObjectPtr ptr)
{
	return ((QAbstractSpinBox*)(ptr))->specialValueText().toUtf8().data();
}

void QAbstractSpinBox_StepBy_Int(QtObjectPtr ptr, int steps)
{
	((QAbstractSpinBox*)(ptr))->stepBy(steps);
}

char* QAbstractSpinBox_Text(QtObjectPtr ptr)
{
	return ((QAbstractSpinBox*)(ptr))->text().toUtf8().data();
}

int QAbstractSpinBox_Wrapping(QtObjectPtr ptr)
{
	return ((QAbstractSpinBox*)(ptr))->wrapping();
}

//Public Slots
void QAbstractSpinBox_ConnectSlotClear(QtObjectPtr ptr)
{
	QObject::connect(((MyQAbstractSpinBox*)(ptr)), &MyQAbstractSpinBox::Slot_Clear, ((QAbstractSpinBox*)(ptr)), &QAbstractSpinBox::clear, Qt::QueuedConnection);
}

void QAbstractSpinBox_DisconnectSlotClear(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAbstractSpinBox*)(ptr)), &MyQAbstractSpinBox::Slot_Clear, ((QAbstractSpinBox*)(ptr)), &QAbstractSpinBox::clear);
}

void QAbstractSpinBox_Clear(QtObjectPtr ptr)
{
	((MyQAbstractSpinBox*)(ptr))->Slot_Clear();
}

void QAbstractSpinBox_ConnectSlotSelectAll(QtObjectPtr ptr)
{
	QObject::connect(((MyQAbstractSpinBox*)(ptr)), &MyQAbstractSpinBox::Slot_SelectAll, ((QAbstractSpinBox*)(ptr)), &QAbstractSpinBox::selectAll, Qt::QueuedConnection);
}

void QAbstractSpinBox_DisconnectSlotSelectAll(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAbstractSpinBox*)(ptr)), &MyQAbstractSpinBox::Slot_SelectAll, ((QAbstractSpinBox*)(ptr)), &QAbstractSpinBox::selectAll);
}

void QAbstractSpinBox_SelectAll(QtObjectPtr ptr)
{
	((MyQAbstractSpinBox*)(ptr))->Slot_SelectAll();
}

void QAbstractSpinBox_ConnectSlotStepDown(QtObjectPtr ptr)
{
	QObject::connect(((MyQAbstractSpinBox*)(ptr)), &MyQAbstractSpinBox::Slot_StepDown, ((QAbstractSpinBox*)(ptr)), &QAbstractSpinBox::stepDown, Qt::QueuedConnection);
}

void QAbstractSpinBox_DisconnectSlotStepDown(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAbstractSpinBox*)(ptr)), &MyQAbstractSpinBox::Slot_StepDown, ((QAbstractSpinBox*)(ptr)), &QAbstractSpinBox::stepDown);
}

void QAbstractSpinBox_StepDown(QtObjectPtr ptr)
{
	((MyQAbstractSpinBox*)(ptr))->Slot_StepDown();
}

void QAbstractSpinBox_ConnectSlotStepUp(QtObjectPtr ptr)
{
	QObject::connect(((MyQAbstractSpinBox*)(ptr)), &MyQAbstractSpinBox::Slot_StepUp, ((QAbstractSpinBox*)(ptr)), &QAbstractSpinBox::stepUp, Qt::QueuedConnection);
}

void QAbstractSpinBox_DisconnectSlotStepUp(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQAbstractSpinBox*)(ptr)), &MyQAbstractSpinBox::Slot_StepUp, ((QAbstractSpinBox*)(ptr)), &QAbstractSpinBox::stepUp);
}

void QAbstractSpinBox_StepUp(QtObjectPtr ptr)
{
	((MyQAbstractSpinBox*)(ptr))->Slot_StepUp();
}

//Signals
void QAbstractSpinBox_ConnectSignalEditingFinished(QtObjectPtr ptr)
{
	QObject::connect(((QAbstractSpinBox*)(ptr)), &QAbstractSpinBox::editingFinished, ((MyQAbstractSpinBox*)(ptr)), &MyQAbstractSpinBox::Signal_EditingFinished, Qt::QueuedConnection);
}

void QAbstractSpinBox_DisconnectSignalEditingFinished(QtObjectPtr ptr)
{
	QObject::disconnect(((QAbstractSpinBox*)(ptr)), &QAbstractSpinBox::editingFinished, ((MyQAbstractSpinBox*)(ptr)), &MyQAbstractSpinBox::Signal_EditingFinished);
}

