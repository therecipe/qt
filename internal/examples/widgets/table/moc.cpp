

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
#include <QAbstractItemView>
#include <QChildEvent>
#include <QEvent>
#include <QHelpEvent>
#include <QLocale>
#include <QMetaMethod>
#include <QMetaObject>
#include <QModelIndex>
#include <QObject>
#include <QPainter>
#include <QSize>
#include <QString>
#include <QStyle>
#include <QStyleOption>
#include <QStyleOptionViewItem>
#include <QStyledItemDelegate>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QWidget>

class Delegate: public QStyledItemDelegate
{
Q_OBJECT
public:
	Delegate(QObject *parent) : QStyledItemDelegate(parent) {};
	QWidget * createEditor(QWidget * parent, const QStyleOptionViewItem & option, const QModelIndex & index) const { return static_cast<QWidget*>(callbackDelegate_CreateEditor(const_cast<Delegate*>(this), parent, const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index))); };
	QString displayText(const QVariant & value, const QLocale & locale) const { return QString(callbackDelegate_DisplayText(const_cast<Delegate*>(this), const_cast<QVariant*>(&value), const_cast<QLocale*>(&locale))); };
	bool editorEvent(QEvent * event, QAbstractItemModel * model, const QStyleOptionViewItem & option, const QModelIndex & index) { return callbackDelegate_EditorEvent(this, event, model, const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index)) != 0; };
	void initStyleOption(QStyleOptionViewItem * option, const QModelIndex & index) const { callbackDelegate_InitStyleOption(const_cast<Delegate*>(this), option, const_cast<QModelIndex*>(&index)); };
	void paint(QPainter * painter, const QStyleOptionViewItem & option, const QModelIndex & index) const { callbackDelegate_Paint(const_cast<Delegate*>(this), painter, const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index)); };
	void setEditorData(QWidget * editor, const QModelIndex & index) const { callbackDelegate_SetEditorData(const_cast<Delegate*>(this), editor, const_cast<QModelIndex*>(&index)); };
	void setModelData(QWidget * editor, QAbstractItemModel * model, const QModelIndex & index) const { callbackDelegate_SetModelData(const_cast<Delegate*>(this), editor, model, const_cast<QModelIndex*>(&index)); };
	QSize sizeHint(const QStyleOptionViewItem & option, const QModelIndex & index) const { return *static_cast<QSize*>(callbackDelegate_SizeHint(const_cast<Delegate*>(this), const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index))); };
	void updateEditorGeometry(QWidget * editor, const QStyleOptionViewItem & option, const QModelIndex & index) const { callbackDelegate_UpdateEditorGeometry(const_cast<Delegate*>(this), editor, const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index)); };
	void destroyEditor(QWidget * editor, const QModelIndex & index) const { callbackDelegate_DestroyEditor(const_cast<Delegate*>(this), editor, const_cast<QModelIndex*>(&index)); };
	bool helpEvent(QHelpEvent * event, QAbstractItemView * view, const QStyleOptionViewItem & option, const QModelIndex & index) { return callbackDelegate_HelpEvent(this, event, view, const_cast<QStyleOptionViewItem*>(&option), const_cast<QModelIndex*>(&index)) != 0; };
	void timerEvent(QTimerEvent * event) { callbackDelegate_TimerEvent(this, event); };
	void childEvent(QChildEvent * event) { callbackDelegate_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackDelegate_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackDelegate_CustomEvent(this, event); };
	void deleteLater() { callbackDelegate_DeleteLater(this); };
	void disconnectNotify(const QMetaMethod & sign) { callbackDelegate_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackDelegate_Event(this, e) != 0; };
	
signals:
public slots:
};

void* Delegate_NewDelegate(void* parent)
{
	return new Delegate(static_cast<QObject*>(parent));
}

void Delegate_DestroyDelegate(void* ptr)
{
	static_cast<Delegate*>(ptr)->~Delegate();
}

void* Delegate_CreateEditor(void* ptr, void* parent, void* option, void* index)
{
	return static_cast<Delegate*>(ptr)->createEditor(static_cast<QWidget*>(parent), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void* Delegate_CreateEditorDefault(void* ptr, void* parent, void* option, void* index)
{
	return static_cast<Delegate*>(ptr)->QStyledItemDelegate::createEditor(static_cast<QWidget*>(parent), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

char* Delegate_DisplayText(void* ptr, void* value, void* locale)
{
	return const_cast<char*>(static_cast<Delegate*>(ptr)->displayText(*static_cast<QVariant*>(value), *static_cast<QLocale*>(locale)).toUtf8().constData());
}

char* Delegate_DisplayTextDefault(void* ptr, void* value, void* locale)
{
	return const_cast<char*>(static_cast<Delegate*>(ptr)->QStyledItemDelegate::displayText(*static_cast<QVariant*>(value), *static_cast<QLocale*>(locale)).toUtf8().constData());
}

char Delegate_EditorEvent(void* ptr, void* event, void* model, void* option, void* index)
{
	return static_cast<Delegate*>(ptr)->editorEvent(static_cast<QEvent*>(event), static_cast<QAbstractItemModel*>(model), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

char Delegate_EditorEventDefault(void* ptr, void* event, void* model, void* option, void* index)
{
	return static_cast<Delegate*>(ptr)->QStyledItemDelegate::editorEvent(static_cast<QEvent*>(event), static_cast<QAbstractItemModel*>(model), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void Delegate_InitStyleOption(void* ptr, void* option, void* index)
{
	static_cast<Delegate*>(ptr)->initStyleOption(static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void Delegate_InitStyleOptionDefault(void* ptr, void* option, void* index)
{
	static_cast<Delegate*>(ptr)->QStyledItemDelegate::initStyleOption(static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void Delegate_Paint(void* ptr, void* painter, void* option, void* index)
{
	static_cast<Delegate*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void Delegate_PaintDefault(void* ptr, void* painter, void* option, void* index)
{
	static_cast<Delegate*>(ptr)->QStyledItemDelegate::paint(static_cast<QPainter*>(painter), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void Delegate_SetEditorData(void* ptr, void* editor, void* index)
{
	static_cast<Delegate*>(ptr)->setEditorData(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

void Delegate_SetEditorDataDefault(void* ptr, void* editor, void* index)
{
	static_cast<Delegate*>(ptr)->QStyledItemDelegate::setEditorData(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

void Delegate_SetModelData(void* ptr, void* editor, void* model, void* index)
{
	static_cast<Delegate*>(ptr)->setModelData(static_cast<QWidget*>(editor), static_cast<QAbstractItemModel*>(model), *static_cast<QModelIndex*>(index));
}

void Delegate_SetModelDataDefault(void* ptr, void* editor, void* model, void* index)
{
	static_cast<Delegate*>(ptr)->QStyledItemDelegate::setModelData(static_cast<QWidget*>(editor), static_cast<QAbstractItemModel*>(model), *static_cast<QModelIndex*>(index));
}

void* Delegate_SizeHint(void* ptr, void* option, void* index)
{
	return ({ QSize tmpValue = static_cast<Delegate*>(ptr)->sizeHint(*static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* Delegate_SizeHintDefault(void* ptr, void* option, void* index)
{
	return ({ QSize tmpValue = static_cast<Delegate*>(ptr)->QStyledItemDelegate::sizeHint(*static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

void Delegate_UpdateEditorGeometry(void* ptr, void* editor, void* option, void* index)
{
	static_cast<Delegate*>(ptr)->updateEditorGeometry(static_cast<QWidget*>(editor), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void Delegate_UpdateEditorGeometryDefault(void* ptr, void* editor, void* option, void* index)
{
	static_cast<Delegate*>(ptr)->QStyledItemDelegate::updateEditorGeometry(static_cast<QWidget*>(editor), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void Delegate_DestroyEditor(void* ptr, void* editor, void* index)
{
	static_cast<Delegate*>(ptr)->destroyEditor(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

void Delegate_DestroyEditorDefault(void* ptr, void* editor, void* index)
{
	static_cast<Delegate*>(ptr)->QStyledItemDelegate::destroyEditor(static_cast<QWidget*>(editor), *static_cast<QModelIndex*>(index));
}

char Delegate_HelpEvent(void* ptr, void* event, void* view, void* option, void* index)
{
	return static_cast<Delegate*>(ptr)->helpEvent(static_cast<QHelpEvent*>(event), static_cast<QAbstractItemView*>(view), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

char Delegate_HelpEventDefault(void* ptr, void* event, void* view, void* option, void* index)
{
	return static_cast<Delegate*>(ptr)->QStyledItemDelegate::helpEvent(static_cast<QHelpEvent*>(event), static_cast<QAbstractItemView*>(view), *static_cast<QStyleOptionViewItem*>(option), *static_cast<QModelIndex*>(index));
}

void Delegate_TimerEvent(void* ptr, void* event)
{
	static_cast<Delegate*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void Delegate_TimerEventDefault(void* ptr, void* event)
{
	static_cast<Delegate*>(ptr)->QStyledItemDelegate::timerEvent(static_cast<QTimerEvent*>(event));
}

void Delegate_ChildEvent(void* ptr, void* event)
{
	static_cast<Delegate*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void Delegate_ChildEventDefault(void* ptr, void* event)
{
	static_cast<Delegate*>(ptr)->QStyledItemDelegate::childEvent(static_cast<QChildEvent*>(event));
}

void Delegate_ConnectNotify(void* ptr, void* sign)
{
	static_cast<Delegate*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void Delegate_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Delegate*>(ptr)->QStyledItemDelegate::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void Delegate_CustomEvent(void* ptr, void* event)
{
	static_cast<Delegate*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void Delegate_CustomEventDefault(void* ptr, void* event)
{
	static_cast<Delegate*>(ptr)->QStyledItemDelegate::customEvent(static_cast<QEvent*>(event));
}

void Delegate_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<Delegate*>(ptr), "deleteLater");
}

void Delegate_DeleteLaterDefault(void* ptr)
{
	static_cast<Delegate*>(ptr)->QStyledItemDelegate::deleteLater();
}

void Delegate_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<Delegate*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void Delegate_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Delegate*>(ptr)->QStyledItemDelegate::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char Delegate_Event(void* ptr, void* e)
{
	return static_cast<Delegate*>(ptr)->event(static_cast<QEvent*>(e));
}

char Delegate_EventDefault(void* ptr, void* e)
{
	return static_cast<Delegate*>(ptr)->QStyledItemDelegate::event(static_cast<QEvent*>(e));
}





#include "moc_moc.h"
