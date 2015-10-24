#include "qcombobox.h"
#include <QAbstractItemModel>
#include <QModelIndex>
#include <QValidator>
#include <QAbstractItemView>
#include <QSize>
#include <QUrl>
#include <QVariant>
#include <QCompleter>
#include <QWidget>
#include <QLine>
#include <QLineEdit>
#include <QAbstractItemDelegate>
#include <QObject>
#include <QString>
#include <QEvent>
#include <QMetaObject>
#include <QIcon>
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

int QComboBox_Count(QtObjectPtr ptr){
	return static_cast<QComboBox*>(ptr)->count();
}

char* QComboBox_CurrentData(QtObjectPtr ptr, int role){
	return static_cast<QComboBox*>(ptr)->currentData(role).toString().toUtf8().data();
}

int QComboBox_CurrentIndex(QtObjectPtr ptr){
	return static_cast<QComboBox*>(ptr)->currentIndex();
}

char* QComboBox_CurrentText(QtObjectPtr ptr){
	return static_cast<QComboBox*>(ptr)->currentText().toUtf8().data();
}

int QComboBox_DuplicatesEnabled(QtObjectPtr ptr){
	return static_cast<QComboBox*>(ptr)->duplicatesEnabled();
}

int QComboBox_HasFrame(QtObjectPtr ptr){
	return static_cast<QComboBox*>(ptr)->hasFrame();
}

int QComboBox_InsertPolicy(QtObjectPtr ptr){
	return static_cast<QComboBox*>(ptr)->insertPolicy();
}

int QComboBox_IsEditable(QtObjectPtr ptr){
	return static_cast<QComboBox*>(ptr)->isEditable();
}

int QComboBox_MaxCount(QtObjectPtr ptr){
	return static_cast<QComboBox*>(ptr)->maxCount();
}

int QComboBox_MaxVisibleItems(QtObjectPtr ptr){
	return static_cast<QComboBox*>(ptr)->maxVisibleItems();
}

int QComboBox_MinimumContentsLength(QtObjectPtr ptr){
	return static_cast<QComboBox*>(ptr)->minimumContentsLength();
}

int QComboBox_ModelColumn(QtObjectPtr ptr){
	return static_cast<QComboBox*>(ptr)->modelColumn();
}

void QComboBox_SetCompleter(QtObjectPtr ptr, QtObjectPtr completer){
	static_cast<QComboBox*>(ptr)->setCompleter(static_cast<QCompleter*>(completer));
}

void QComboBox_SetCurrentIndex(QtObjectPtr ptr, int index){
	QMetaObject::invokeMethod(static_cast<QComboBox*>(ptr), "setCurrentIndex", Q_ARG(int, index));
}

void QComboBox_SetCurrentText(QtObjectPtr ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QComboBox*>(ptr), "setCurrentText", Q_ARG(QString, QString(text)));
}

void QComboBox_SetDuplicatesEnabled(QtObjectPtr ptr, int enable){
	static_cast<QComboBox*>(ptr)->setDuplicatesEnabled(enable != 0);
}

void QComboBox_SetEditable(QtObjectPtr ptr, int editable){
	static_cast<QComboBox*>(ptr)->setEditable(editable != 0);
}

void QComboBox_SetFrame(QtObjectPtr ptr, int v){
	static_cast<QComboBox*>(ptr)->setFrame(v != 0);
}

void QComboBox_SetIconSize(QtObjectPtr ptr, QtObjectPtr size){
	static_cast<QComboBox*>(ptr)->setIconSize(*static_cast<QSize*>(size));
}

void QComboBox_SetInsertPolicy(QtObjectPtr ptr, int policy){
	static_cast<QComboBox*>(ptr)->setInsertPolicy(static_cast<QComboBox::InsertPolicy>(policy));
}

void QComboBox_SetMaxCount(QtObjectPtr ptr, int max){
	static_cast<QComboBox*>(ptr)->setMaxCount(max);
}

void QComboBox_SetMaxVisibleItems(QtObjectPtr ptr, int maxItems){
	static_cast<QComboBox*>(ptr)->setMaxVisibleItems(maxItems);
}

void QComboBox_SetMinimumContentsLength(QtObjectPtr ptr, int characters){
	static_cast<QComboBox*>(ptr)->setMinimumContentsLength(characters);
}

void QComboBox_SetModelColumn(QtObjectPtr ptr, int visibleColumn){
	static_cast<QComboBox*>(ptr)->setModelColumn(visibleColumn);
}

void QComboBox_SetSizeAdjustPolicy(QtObjectPtr ptr, int policy){
	static_cast<QComboBox*>(ptr)->setSizeAdjustPolicy(static_cast<QComboBox::SizeAdjustPolicy>(policy));
}

void QComboBox_SetValidator(QtObjectPtr ptr, QtObjectPtr validator){
	static_cast<QComboBox*>(ptr)->setValidator(static_cast<QValidator*>(validator));
}

int QComboBox_SizeAdjustPolicy(QtObjectPtr ptr){
	return static_cast<QComboBox*>(ptr)->sizeAdjustPolicy();
}

QtObjectPtr QComboBox_NewQComboBox(QtObjectPtr parent){
	return new QComboBox(static_cast<QWidget*>(parent));
}

void QComboBox_ConnectActivated(QtObjectPtr ptr){
	QObject::connect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::activated), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_Activated));;
}

void QComboBox_DisconnectActivated(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::activated), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_Activated));;
}

void QComboBox_AddItem2(QtObjectPtr ptr, QtObjectPtr icon, char* text, char* userData){
	static_cast<QComboBox*>(ptr)->addItem(*static_cast<QIcon*>(icon), QString(text), QVariant(userData));
}

void QComboBox_AddItem(QtObjectPtr ptr, char* text, char* userData){
	static_cast<QComboBox*>(ptr)->addItem(QString(text), QVariant(userData));
}

void QComboBox_AddItems(QtObjectPtr ptr, char* texts){
	static_cast<QComboBox*>(ptr)->addItems(QString(texts).split("|", QString::SkipEmptyParts));
}

void QComboBox_Clear(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QComboBox*>(ptr), "clear");
}

void QComboBox_ClearEditText(QtObjectPtr ptr){
	QMetaObject::invokeMethod(static_cast<QComboBox*>(ptr), "clearEditText");
}

QtObjectPtr QComboBox_Completer(QtObjectPtr ptr){
	return static_cast<QComboBox*>(ptr)->completer();
}

void QComboBox_ConnectCurrentIndexChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::currentIndexChanged), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_CurrentIndexChanged));;
}

void QComboBox_DisconnectCurrentIndexChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::currentIndexChanged), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_CurrentIndexChanged));;
}

void QComboBox_ConnectCurrentTextChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(const QString &)>(&QComboBox::currentTextChanged), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(const QString &)>(&MyQComboBox::Signal_CurrentTextChanged));;
}

void QComboBox_DisconnectCurrentTextChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(const QString &)>(&QComboBox::currentTextChanged), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(const QString &)>(&MyQComboBox::Signal_CurrentTextChanged));;
}

void QComboBox_ConnectEditTextChanged(QtObjectPtr ptr){
	QObject::connect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(const QString &)>(&QComboBox::editTextChanged), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(const QString &)>(&MyQComboBox::Signal_EditTextChanged));;
}

void QComboBox_DisconnectEditTextChanged(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(const QString &)>(&QComboBox::editTextChanged), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(const QString &)>(&MyQComboBox::Signal_EditTextChanged));;
}

int QComboBox_Event(QtObjectPtr ptr, QtObjectPtr event){
	return static_cast<QComboBox*>(ptr)->event(static_cast<QEvent*>(event));
}

int QComboBox_FindData(QtObjectPtr ptr, char* data, int role, int flags){
	return static_cast<QComboBox*>(ptr)->findData(QVariant(data), role, static_cast<Qt::MatchFlag>(flags));
}

int QComboBox_FindText(QtObjectPtr ptr, char* text, int flags){
	return static_cast<QComboBox*>(ptr)->findText(QString(text), static_cast<Qt::MatchFlag>(flags));
}

void QComboBox_HidePopup(QtObjectPtr ptr){
	static_cast<QComboBox*>(ptr)->hidePopup();
}

void QComboBox_ConnectHighlighted(QtObjectPtr ptr){
	QObject::connect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::highlighted), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_Highlighted));;
}

void QComboBox_DisconnectHighlighted(QtObjectPtr ptr){
	QObject::disconnect(static_cast<QComboBox*>(ptr), static_cast<void (QComboBox::*)(int)>(&QComboBox::highlighted), static_cast<MyQComboBox*>(ptr), static_cast<void (MyQComboBox::*)(int)>(&MyQComboBox::Signal_Highlighted));;
}

char* QComboBox_InputMethodQuery(QtObjectPtr ptr, int query){
	return static_cast<QComboBox*>(ptr)->inputMethodQuery(static_cast<Qt::InputMethodQuery>(query)).toString().toUtf8().data();
}

void QComboBox_InsertItem2(QtObjectPtr ptr, int index, QtObjectPtr icon, char* text, char* userData){
	static_cast<QComboBox*>(ptr)->insertItem(index, *static_cast<QIcon*>(icon), QString(text), QVariant(userData));
}

void QComboBox_InsertItem(QtObjectPtr ptr, int index, char* text, char* userData){
	static_cast<QComboBox*>(ptr)->insertItem(index, QString(text), QVariant(userData));
}

void QComboBox_InsertItems(QtObjectPtr ptr, int index, char* list){
	static_cast<QComboBox*>(ptr)->insertItems(index, QString(list).split("|", QString::SkipEmptyParts));
}

void QComboBox_InsertSeparator(QtObjectPtr ptr, int index){
	static_cast<QComboBox*>(ptr)->insertSeparator(index);
}

char* QComboBox_ItemData(QtObjectPtr ptr, int index, int role){
	return static_cast<QComboBox*>(ptr)->itemData(index, role).toString().toUtf8().data();
}

QtObjectPtr QComboBox_ItemDelegate(QtObjectPtr ptr){
	return static_cast<QComboBox*>(ptr)->itemDelegate();
}

char* QComboBox_ItemText(QtObjectPtr ptr, int index){
	return static_cast<QComboBox*>(ptr)->itemText(index).toUtf8().data();
}

QtObjectPtr QComboBox_LineEdit(QtObjectPtr ptr){
	return static_cast<QComboBox*>(ptr)->lineEdit();
}

QtObjectPtr QComboBox_Model(QtObjectPtr ptr){
	return static_cast<QComboBox*>(ptr)->model();
}

void QComboBox_RemoveItem(QtObjectPtr ptr, int index){
	static_cast<QComboBox*>(ptr)->removeItem(index);
}

QtObjectPtr QComboBox_RootModelIndex(QtObjectPtr ptr){
	return static_cast<QComboBox*>(ptr)->rootModelIndex().internalPointer();
}

void QComboBox_SetEditText(QtObjectPtr ptr, char* text){
	QMetaObject::invokeMethod(static_cast<QComboBox*>(ptr), "setEditText", Q_ARG(QString, QString(text)));
}

void QComboBox_SetItemData(QtObjectPtr ptr, int index, char* value, int role){
	static_cast<QComboBox*>(ptr)->setItemData(index, QVariant(value), role);
}

void QComboBox_SetItemDelegate(QtObjectPtr ptr, QtObjectPtr delegate){
	static_cast<QComboBox*>(ptr)->setItemDelegate(static_cast<QAbstractItemDelegate*>(delegate));
}

void QComboBox_SetItemIcon(QtObjectPtr ptr, int index, QtObjectPtr icon){
	static_cast<QComboBox*>(ptr)->setItemIcon(index, *static_cast<QIcon*>(icon));
}

void QComboBox_SetItemText(QtObjectPtr ptr, int index, char* text){
	static_cast<QComboBox*>(ptr)->setItemText(index, QString(text));
}

void QComboBox_SetLineEdit(QtObjectPtr ptr, QtObjectPtr edit){
	static_cast<QComboBox*>(ptr)->setLineEdit(static_cast<QLineEdit*>(edit));
}

void QComboBox_SetModel(QtObjectPtr ptr, QtObjectPtr model){
	static_cast<QComboBox*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QComboBox_SetRootModelIndex(QtObjectPtr ptr, QtObjectPtr index){
	static_cast<QComboBox*>(ptr)->setRootModelIndex(*static_cast<QModelIndex*>(index));
}

void QComboBox_SetView(QtObjectPtr ptr, QtObjectPtr itemView){
	static_cast<QComboBox*>(ptr)->setView(static_cast<QAbstractItemView*>(itemView));
}

void QComboBox_ShowPopup(QtObjectPtr ptr){
	static_cast<QComboBox*>(ptr)->showPopup();
}

QtObjectPtr QComboBox_Validator(QtObjectPtr ptr){
	return const_cast<QValidator*>(static_cast<QComboBox*>(ptr)->validator());
}

QtObjectPtr QComboBox_View(QtObjectPtr ptr){
	return static_cast<QComboBox*>(ptr)->view();
}

void QComboBox_DestroyQComboBox(QtObjectPtr ptr){
	static_cast<QComboBox*>(ptr)->~QComboBox();
}

