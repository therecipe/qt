#include "qcombobox.h"
#include <QSize>
#include <QValidator>
#include <QString>
#include <QModelIndex>
#include <QWidget>
#include <QIcon>
#include <QEvent>
#include <QLine>
#include <QAbstractItemDelegate>
#include <QMetaObject>
#include <QUrl>
#include <QCompleter>
#include <QLineEdit>
#include <QVariant>
#include <QAbstractItemModel>
#include <QObject>
#include <QAbstractItemView>
#include <QComboBox>
#include "_cgo_export.h"

class MyQComboBox: public QComboBox {
public:
void Signal_Activated(int index){callbackQComboBoxActivated(this->objectName().toUtf8().data(), index);};
void Signal_CurrentIndexChanged(int index){callbackQComboBoxCurrentIndexChanged(this->objectName().toUtf8().data(), index);};
void Signal_CurrentTextChanged(const QString & text){callbackQComboBoxCurrentTextChanged(this->objectName().toUtf8().data(), text.toUtf8().data());};
void Signal_EditTextChanged(const QString & text){callbackQComboBoxEditTextChanged(this->objectName().toUtf8().data(), text.toUtf8().data());};
void Signal_Highlighted(int index){callbackQComboBoxHighlighted(this->objectName().toUtf8().data(), index);};
};

int QComboBox_Count(void* ptr){
	return static_cast<QComboBox*>(ptr)->count();
}

void* QComboBox_CurrentData(void* ptr, int role){
	return new QVariant(static_cast<QComboBox*>(ptr)->currentData(role));
}

int QComboBox_CurrentIndex(void* ptr){
	return static_cast<QComboBox*>(ptr)->currentIndex();
}

char* QComboBox_CurrentText(void* ptr){
	return static_cast<QComboBox*>(ptr)->currentText().toUtf8().data();
}

int QComboBox_DuplicatesEnabled(void* ptr){
	return static_cast<QComboBox*>(ptr)->duplicatesEnabled();
}

int QComboBox_HasFrame(void* ptr){
	return static_cast<QComboBox*>(ptr)->hasFrame();
}

int QComboBox_InsertPolicy(void* ptr){
	return static_cast<QComboBox*>(ptr)->insertPolicy();
}

int QComboBox_IsEditable(void* ptr){
	return static_cast<QComboBox*>(ptr)->isEditable();
}

int QComboBox_MaxCount(void* ptr){
	return static_cast<QComboBox*>(ptr)->maxCount();
}

int QComboBox_MaxVisibleItems(void* ptr){
	return static_cast<QComboBox*>(ptr)->maxVisibleItems();
}

int QComboBox_MinimumContentsLength(void* ptr){
	return static_cast<QComboBox*>(ptr)->minimumContentsLength();
}

int QComboBox_ModelColumn(void* ptr){
	return static_cast<QComboBox*>(ptr)->modelColumn();
}

void QComboBox_SetCompleter(void* ptr, void* completer){
	static_cast<QComboBox*>(ptr)->setCompleter(static_cast<QCompleter*>(completer));
}

void QComboBox_SetCurrentIndex(void* ptr, int index){
	QMetaObject::invokeMethod(static_cast<QComboBox*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void QComboBox_SetCurrentText(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QComboBox*>(ptr), "setCurrentText", Q_ARG(QString, QString(text)));
}

void QComboBox_SetDuplicatesEnabled(void* ptr, int enable){
	static_cast<QComboBox*>(ptr)->setDuplicatesEnabled(enable != 0);
}

void QComboBox_SetEditable(void* ptr, int editable){
	static_cast<QComboBox*>(ptr)->setEditable(editable != 0);
}

void QComboBox_SetFrame(void* ptr, int v){
	static_cast<QComboBox*>(ptr)->setFrame(v != 0);
}

void QComboBox_SetIconSize(void* ptr, void* size){
	static_cast<QComboBox*>(ptr)->setIconSize(*static_cast<QSize*>(size));
}

void QComboBox_SetInsertPolicy(void* ptr, int policy){
	static_cast<QComboBox*>(ptr)->setInsertPolicy(static_cast<QComboBox::InsertPolicy>(policy));
}

void QComboBox_SetMaxCount(void* ptr, int max){
	static_cast<QComboBox*>(ptr)->setMaxCount(max);
}

void QComboBox_SetMaxVisibleItems(void* ptr, int maxItems){
	static_cast<QComboBox*>(ptr)->setMaxVisibleItems(maxItems);
}

void QComboBox_SetMinimumContentsLength(void* ptr, int characters){
	static_cast<QComboBox*>(ptr)->setMinimumContentsLength(characters);
}

void QComboBox_SetModelColumn(void* ptr, int visibleColumn){
	static_cast<QComboBox*>(ptr)->setModelColumn(visibleColumn);
}

void QComboBox_SetSizeAdjustPolicy(void* ptr, int policy){
	static_cast<QComboBox*>(ptr)->setSizeAdjustPolicy(static_cast<QComboBox::SizeAdjustPolicy>(policy));
}

void QComboBox_SetValidator(void* ptr, void* validator){
	static_cast<QComboBox*>(ptr)->setValidator(static_cast<QValidator*>(validator));
}

int QComboBox_SizeAdjustPolicy(void* ptr){
	return static_cast<QComboBox*>(ptr)->sizeAdjustPolicy();
}

void* QComboBox_NewQComboBox(void* parent){
	return new QComboBox(static_cast<QWidget*>(parent));
}

void QComboBox_ConnectActivated(void* ptr){
	QObject::connect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::activated), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_Activated));;
}

void QComboBox_DisconnectActivated(void* ptr){
	QObject::disconnect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::activated), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_Activated));;
}

void QComboBox_AddItem2(void* ptr, void* icon, char* text, void* userData){
	static_cast<QComboBox*>(ptr)->addItem(*static_cast<QIcon*>(icon), QString(text), *static_cast<QVariant*>(userData));
}

void QComboBox_AddItem(void* ptr, char* text, void* userData){
	static_cast<QComboBox*>(ptr)->addItem(QString(text), *static_cast<QVariant*>(userData));
}

void QComboBox_AddItems(void* ptr, char* texts){
	static_cast<QComboBox*>(ptr)->addItems(QString(texts).split("|", QString::SkipEmptyParts));
}

void QComboBox_Clear(void* ptr){
	QMetaObject::invokeMethod(static_cast<QComboBox*>(ptr), "clear");
}

void QComboBox_ClearEditText(void* ptr){
	QMetaObject::invokeMethod(static_cast<QComboBox*>(ptr), "clearEditText");
}

void* QComboBox_Completer(void* ptr){
	return static_cast<QComboBox*>(ptr)->completer();
}

void QComboBox_ConnectCurrentIndexChanged(void* ptr){
	QObject::connect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::currentIndexChanged), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_CurrentIndexChanged));;
}

void QComboBox_DisconnectCurrentIndexChanged(void* ptr){
	QObject::disconnect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::currentIndexChanged), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_CurrentIndexChanged));;
}

void QComboBox_ConnectCurrentTextChanged(void* ptr){
	QObject::connect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(const QString &)>(&QComboBox::currentTextChanged), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(const QString &)>(&MyQComboBox::Signal_CurrentTextChanged));;
}

void QComboBox_DisconnectCurrentTextChanged(void* ptr){
	QObject::disconnect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(const QString &)>(&QComboBox::currentTextChanged), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(const QString &)>(&MyQComboBox::Signal_CurrentTextChanged));;
}

void QComboBox_ConnectEditTextChanged(void* ptr){
	QObject::connect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(const QString &)>(&QComboBox::editTextChanged), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(const QString &)>(&MyQComboBox::Signal_EditTextChanged));;
}

void QComboBox_DisconnectEditTextChanged(void* ptr){
	QObject::disconnect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(const QString &)>(&QComboBox::editTextChanged), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(const QString &)>(&MyQComboBox::Signal_EditTextChanged));;
}

int QComboBox_Event(void* ptr, void* event){
	return static_cast<QComboBox*>(ptr)->event(static_cast<QEvent*>(event));
}

int QComboBox_FindData(void* ptr, void* data, int role, int flags){
	return static_cast<QComboBox*>(ptr)->findData(*static_cast<QVariant*>(data), role, static_cast<Qt::MatchFlag>(flags));
}

int QComboBox_FindText(void* ptr, char* text, int flags){
	return static_cast<QComboBox*>(ptr)->findText(QString(text), static_cast<Qt::MatchFlag>(flags));
}

void QComboBox_HidePopup(void* ptr){
	static_cast<QComboBox*>(ptr)->hidePopup();
}

void QComboBox_ConnectHighlighted(void* ptr){
	QObject::connect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::highlighted), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_Highlighted));;
}

void QComboBox_DisconnectHighlighted(void* ptr){
	QObject::disconnect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::highlighted), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_Highlighted));;
}

void* QComboBox_InputMethodQuery(void* ptr, int query){
	return new QVariant(static_cast<QComboBox*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)));
}

void QComboBox_InsertItem2(void* ptr, int index, void* icon, char* text, void* userData){
	static_cast<QComboBox*>(ptr)->insertItem(index, *static_cast<QIcon*>(icon), QString(text), *static_cast<QVariant*>(userData));
}

void QComboBox_InsertItem(void* ptr, int index, char* text, void* userData){
	static_cast<QComboBox*>(ptr)->insertItem(index, QString(text), *static_cast<QVariant*>(userData));
}

void QComboBox_InsertItems(void* ptr, int index, char* list){
	static_cast<QComboBox*>(ptr)->insertItems(index, QString(list).split("|", QString::SkipEmptyParts));
}

void QComboBox_InsertSeparator(void* ptr, int index){
	static_cast<QComboBox*>(ptr)->insertSeparator(index);
}

void* QComboBox_ItemData(void* ptr, int index, int role){
	return new QVariant(static_cast<QComboBox*>(ptr)->itemData(index, role));
}

void* QComboBox_ItemDelegate(void* ptr){
	return static_cast<QComboBox*>(ptr)->itemDelegate();
}

char* QComboBox_ItemText(void* ptr, int index){
	return static_cast<QComboBox*>(ptr)->itemText(index).toUtf8().data();
}

void* QComboBox_LineEdit(void* ptr){
	return static_cast<QComboBox*>(ptr)->lineEdit();
}

void* QComboBox_Model(void* ptr){
	return static_cast<QComboBox*>(ptr)->model();
}

void QComboBox_RemoveItem(void* ptr, int index){
	static_cast<QComboBox*>(ptr)->removeItem(index);
}

void* QComboBox_RootModelIndex(void* ptr){
	return static_cast<QComboBox*>(ptr)->rootModelIndex().internalPointer();
}

void QComboBox_SetEditText(void* ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QComboBox*>(ptr), "setEditText", Q_ARG(QString, QString(text)));
}

void QComboBox_SetItemData(void* ptr, int index, void* value, int role){
	static_cast<QComboBox*>(ptr)->setItemData(index, *static_cast<QVariant*>(value), role);
}

void QComboBox_SetItemDelegate(void* ptr, void* delegate){
	static_cast<QComboBox*>(ptr)->setItemDelegate(static_cast<QAbstractItemDelegate*>(delegate));
}

void QComboBox_SetItemIcon(void* ptr, int index, void* icon){
	static_cast<QComboBox*>(ptr)->setItemIcon(index, *static_cast<QIcon*>(icon));
}

void QComboBox_SetItemText(void* ptr, int index, char* text){
	static_cast<QComboBox*>(ptr)->setItemText(index, QString(text));
}

void QComboBox_SetLineEdit(void* ptr, void* edit){
	static_cast<QComboBox*>(ptr)->setLineEdit(static_cast<QLineEdit*>(edit));
}

void QComboBox_SetModel(void* ptr, void* model){
	static_cast<QComboBox*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QComboBox_SetRootModelIndex(void* ptr, void* index){
	static_cast<QComboBox*>(ptr)->setRootModelIndex(*static_cast<QModelIndex*>(index));
}

void QComboBox_SetView(void* ptr, void* itemView){
	static_cast<QComboBox*>(ptr)->setView(static_cast<QAbstractItemView*>(itemView));
}

void QComboBox_ShowPopup(void* ptr){
	static_cast<QComboBox*>(ptr)->showPopup();
}

void* QComboBox_Validator(void* ptr){
	return const_cast<QValidator*>(static_cast<QComboBox*>(ptr)->validator());
}

void* QComboBox_View(void* ptr){
	return static_cast<QComboBox*>(ptr)->view();
}

void QComboBox_DestroyQComboBox(void* ptr){
	static_cast<QComboBox*>(ptr)->~QComboBox();
}

