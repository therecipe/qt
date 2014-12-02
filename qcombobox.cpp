#include "qcombobox.h"
#include <QComboBox>
#include "cgoexport.h"

//MyQComboBox
class MyQComboBox: public QComboBox {
Q_OBJECT
public:
void Signal_CurrentTextChanged() { callbackQComboBox(this, QString("currentTextChanged").toUtf8().data()); };
void Signal_EditTextChanged() { callbackQComboBox(this, QString("editTextChanged").toUtf8().data()); };

signals:
void Slot_Clear();
void Slot_ClearEditText();
void Slot_SetCurrentIndex(int index);

};
#include "qcombobox.moc"


//Public Functions
QtObjectPtr QComboBox_New_QWidget(QtObjectPtr parent)
{
	return new QComboBox(((QWidget*)(parent)));
}

void QComboBox_Destroy(QtObjectPtr ptr)
{
	((QComboBox*)(ptr))->~QComboBox();
}

void QComboBox_AddItems_QStringList(QtObjectPtr ptr, char* texts)
{
	((QComboBox*)(ptr))->addItems(QString(texts).split("|", QString::SkipEmptyParts));
}

int QComboBox_Count(QtObjectPtr ptr)
{
	return ((QComboBox*)(ptr))->count();
}

int QComboBox_CurrentIndex(QtObjectPtr ptr)
{
	return ((QComboBox*)(ptr))->currentIndex();
}

char* QComboBox_CurrentText(QtObjectPtr ptr)
{
	return ((QComboBox*)(ptr))->currentText().toUtf8().data();
}

int QComboBox_DuplicatesEnabled(QtObjectPtr ptr)
{
	return ((QComboBox*)(ptr))->duplicatesEnabled();
}

int QComboBox_HasFrame(QtObjectPtr ptr)
{
	return ((QComboBox*)(ptr))->hasFrame();
}

void QComboBox_HidePopup(QtObjectPtr ptr)
{
	((QComboBox*)(ptr))->hidePopup();
}

void QComboBox_InsertSeparator_Int(QtObjectPtr ptr, int index)
{
	((QComboBox*)(ptr))->insertSeparator(index);
}

int QComboBox_IsEditable(QtObjectPtr ptr)
{
	return ((QComboBox*)(ptr))->isEditable();
}

char* QComboBox_ItemText_Int(QtObjectPtr ptr, int index)
{
	return ((QComboBox*)(ptr))->itemText(index).toUtf8().data();
}

QtObjectPtr QComboBox_LineEdit(QtObjectPtr ptr)
{
	return ((QComboBox*)(ptr))->lineEdit();
}

int QComboBox_MaxCount(QtObjectPtr ptr)
{
	return ((QComboBox*)(ptr))->maxCount();
}

int QComboBox_MaxVisibleItems(QtObjectPtr ptr)
{
	return ((QComboBox*)(ptr))->maxVisibleItems();
}

int QComboBox_MinimumContentsLength(QtObjectPtr ptr)
{
	return ((QComboBox*)(ptr))->minimumContentsLength();
}

int QComboBox_ModelColumn(QtObjectPtr ptr)
{
	return ((QComboBox*)(ptr))->modelColumn();
}

void QComboBox_RemoveItem_Int(QtObjectPtr ptr, int index)
{
	((QComboBox*)(ptr))->removeItem(index);
}

void QComboBox_SetDuplicatesEnabled_Bool(QtObjectPtr ptr, int enable)
{
	((QComboBox*)(ptr))->setDuplicatesEnabled(enable != 0);
}

void QComboBox_SetEditable_Bool(QtObjectPtr ptr, int editable)
{
	((QComboBox*)(ptr))->setEditable(editable != 0);
}

void QComboBox_SetFrame_Bool(QtObjectPtr ptr, int frame)
{
	((QComboBox*)(ptr))->setFrame(frame != 0);
}

void QComboBox_SetItemText_Int_String(QtObjectPtr ptr, int index, char* text)
{
	((QComboBox*)(ptr))->setItemText(index, QString(text));
}

void QComboBox_SetLineEdit_QLineEdit(QtObjectPtr ptr, QtObjectPtr edit)
{
	((QComboBox*)(ptr))->setLineEdit(((QLineEdit*)(edit)));
}

void QComboBox_SetMaxCount_Int(QtObjectPtr ptr, int max)
{
	((QComboBox*)(ptr))->setMaxCount(max);
}

void QComboBox_SetMaxVisibleItems_Int(QtObjectPtr ptr, int maxItems)
{
	((QComboBox*)(ptr))->setMaxVisibleItems(maxItems);
}

void QComboBox_SetMinimumContentsLength_Int(QtObjectPtr ptr, int characters)
{
	((QComboBox*)(ptr))->setMinimumContentsLength(characters);
}

void QComboBox_SetModelColumn_Int(QtObjectPtr ptr, int visibleColumn)
{
	((QComboBox*)(ptr))->setModelColumn(visibleColumn);
}

void QComboBox_ShowPopup(QtObjectPtr ptr)
{
	((QComboBox*)(ptr))->showPopup();
}

//Public Slots
void QComboBox_ConnectSlotClear(QtObjectPtr ptr)
{
	QObject::connect(((MyQComboBox*)(ptr)), &MyQComboBox::Slot_Clear, ((QComboBox*)(ptr)), &QComboBox::clear, Qt::QueuedConnection);
}

void QComboBox_DisconnectSlotClear(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQComboBox*)(ptr)), &MyQComboBox::Slot_Clear, ((QComboBox*)(ptr)), &QComboBox::clear);
}

void QComboBox_Clear(QtObjectPtr ptr)
{
	((MyQComboBox*)(ptr))->Slot_Clear();
}

void QComboBox_ConnectSlotClearEditText(QtObjectPtr ptr)
{
	QObject::connect(((MyQComboBox*)(ptr)), &MyQComboBox::Slot_ClearEditText, ((QComboBox*)(ptr)), &QComboBox::clearEditText, Qt::QueuedConnection);
}

void QComboBox_DisconnectSlotClearEditText(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQComboBox*)(ptr)), &MyQComboBox::Slot_ClearEditText, ((QComboBox*)(ptr)), &QComboBox::clearEditText);
}

void QComboBox_ClearEditText(QtObjectPtr ptr)
{
	((MyQComboBox*)(ptr))->Slot_ClearEditText();
}

void QComboBox_ConnectSlotSetCurrentIndex(QtObjectPtr ptr)
{
	QObject::connect(((MyQComboBox*)(ptr)), &MyQComboBox::Slot_SetCurrentIndex, ((QComboBox*)(ptr)), &QComboBox::setCurrentIndex, Qt::QueuedConnection);
}

void QComboBox_DisconnectSlotSetCurrentIndex(QtObjectPtr ptr)
{
	QObject::disconnect(((MyQComboBox*)(ptr)), &MyQComboBox::Slot_SetCurrentIndex, ((QComboBox*)(ptr)), &QComboBox::setCurrentIndex);
}

void QComboBox_SetCurrentIndex_Int(QtObjectPtr ptr, int index)
{
	((MyQComboBox*)(ptr))->Slot_SetCurrentIndex(index);
}

//Signals
void QComboBox_ConnectSignalCurrentTextChanged(QtObjectPtr ptr)
{
	QObject::connect(((QComboBox*)(ptr)), &QComboBox::currentTextChanged, ((MyQComboBox*)(ptr)), &MyQComboBox::Signal_CurrentTextChanged, Qt::QueuedConnection);
}

void QComboBox_DisconnectSignalCurrentTextChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QComboBox*)(ptr)), &QComboBox::currentTextChanged, ((MyQComboBox*)(ptr)), &MyQComboBox::Signal_CurrentTextChanged);
}

void QComboBox_ConnectSignalEditTextChanged(QtObjectPtr ptr)
{
	QObject::connect(((QComboBox*)(ptr)), &QComboBox::editTextChanged, ((MyQComboBox*)(ptr)), &MyQComboBox::Signal_EditTextChanged, Qt::QueuedConnection);
}

void QComboBox_DisconnectSignalEditTextChanged(QtObjectPtr ptr)
{
	QObject::disconnect(((QComboBox*)(ptr)), &QComboBox::editTextChanged, ((MyQComboBox*)(ptr)), &MyQComboBox::Signal_EditTextChanged);
}

