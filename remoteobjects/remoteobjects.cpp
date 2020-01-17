// +build !minimal

#define protected public
#define private public

#include "remoteobjects.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
#include <QAbstractItemModelReplica>
#include <QAudioSystemPlugin>
#include <QByteArray>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusPendingCallWatcher>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHash>
#include <QIODevice>
#include <QItemSelectionModel>
#include <QLayout>
#include <QMap>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMediaServiceProviderPlugin>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMimeData>
#include <QModelIndex>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QPersistentModelIndex>
#include <QQuickItem>
#include <QRadioData>
#include <QRemoteObjectAbstractPersistedStore>
#include <QRemoteObjectDynamicReplica>
#include <QRemoteObjectHost>
#include <QRemoteObjectHostBase>
#include <QRemoteObjectNode>
#include <QRemoteObjectPendingCallWatcher>
#include <QRemoteObjectRegistry>
#include <QRemoteObjectRegistryHost>
#include <QRemoteObjectReplica>
#include <QRemoteObjectSettingsStore>
#include <QScriptExtensionPlugin>
#include <QSize>
#include <QString>
#include <QTimerEvent>
#include <QUrl>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QWindow>

class MyQAbstractItemModelReplica: public QAbstractItemModelReplica
{
public:
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQAbstractItemModelReplica_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackQAbstractItemModelReplica_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackQAbstractItemModelReplica_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackQAbstractItemModelReplica_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModelReplica_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackQAbstractItemModelReplica_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModelReplica_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModelReplica_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackQAbstractItemModelReplica_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModelReplica_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackQAbstractItemModelReplica_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackQAbstractItemModelReplica_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue037c88 = new QVector<int>(roles); QtRemoteObjects_PackedList { tmpValue037c88, tmpValue037c88->size() }; })); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackQAbstractItemModelReplica_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	void fetchMore(const QModelIndex & parent) { callbackQAbstractItemModelReplica_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackQAbstractItemModelReplica_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool hasChildren(const QModelIndex & parent) const { return callbackQAbstractItemModelReplica_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackQAbstractItemModelReplica_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackQAbstractItemModelReplica_HeaderDataChanged(this, orientation, first, last); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackQAbstractItemModelReplica_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackQAbstractItemModelReplica_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackQAbstractItemModelReplica_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackQAbstractItemModelReplica_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQAbstractItemModelReplica_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); QtRemoteObjects_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQAbstractItemModelReplica_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); QtRemoteObjects_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackQAbstractItemModelReplica_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackQAbstractItemModelReplica_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValuee0adf2 = new QList<QModelIndex>(indexes); QtRemoteObjects_PackedList { tmpValuee0adf2, tmpValuee0adf2->size() }; }))); };
	QStringList mimeTypes() const { return ({ QtRemoteObjects_PackedString tempVal = callbackQAbstractItemModelReplica_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackQAbstractItemModelReplica_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackQAbstractItemModelReplica_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQAbstractItemModelReplica_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQAbstractItemModelReplica_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQAbstractItemModelReplica_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackQAbstractItemModelReplica_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackQAbstractItemModelReplica_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackQAbstractItemModelReplica_ResetInternalData(this); };
	void revert() { callbackQAbstractItemModelReplica_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackQAbstractItemModelReplica_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackQAbstractItemModelReplica_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackQAbstractItemModelReplica_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackQAbstractItemModelReplica_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModelReplica_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModelReplica_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackQAbstractItemModelReplica_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModelReplica_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackQAbstractItemModelReplica_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackQAbstractItemModelReplica_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackQAbstractItemModelReplica_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue037c88 = new QMap<int, QVariant>(roles); QtRemoteObjects_PackedList { tmpValue037c88, tmpValue037c88->size() }; })) != 0; };
	QModelIndex sibling(int row, int column, const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQAbstractItemModelReplica_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&index))); };
	void sort(int column, Qt::SortOrder order) { callbackQAbstractItemModelReplica_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackQAbstractItemModelReplica_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackQAbstractItemModelReplica_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackQAbstractItemModelReplica_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackQAbstractItemModelReplica_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackQAbstractItemModelReplica_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractItemModelReplica_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAbstractItemModelReplica_CustomEvent(this, event); };
	void deleteLater() { callbackQAbstractItemModelReplica_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAbstractItemModelReplica_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractItemModelReplica_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAbstractItemModelReplica_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractItemModelReplica_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractItemModelReplica_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQAbstractItemModelReplica_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractItemModelReplica_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QAbstractItemModelReplica*)
Q_DECLARE_METATYPE(MyQAbstractItemModelReplica*)

int QAbstractItemModelReplica_QAbstractItemModelReplica_QRegisterMetaType(){qRegisterMetaType<QAbstractItemModelReplica*>(); return qRegisterMetaType<MyQAbstractItemModelReplica*>();}

int QAbstractItemModelReplica___QAbstractItemModelReplica_rolesHint_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QAbstractItemModelReplica___QAbstractItemModelReplica_rolesHint_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QAbstractItemModelReplica___QAbstractItemModelReplica_rolesHint_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

int QAbstractItemModelReplica___availableRoles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QAbstractItemModelReplica___availableRoles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QAbstractItemModelReplica___availableRoles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* QAbstractItemModelReplica___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QAbstractItemModelReplica___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* QAbstractItemModelReplica___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct QtRemoteObjects_PackedList QAbstractItemModelReplica___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue7fc3bb = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); QtRemoteObjects_PackedList { tmpValue7fc3bb, tmpValue7fc3bb->size() }; });
}

int QAbstractItemModelReplica_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractItemModelReplica_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QAbstractItemModelReplica_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* QAbstractItemModelReplica___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModelReplica___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QAbstractItemModelReplica___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QAbstractItemModelReplica___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModelReplica___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QAbstractItemModelReplica___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int QAbstractItemModelReplica___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QAbstractItemModelReplica___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QAbstractItemModelReplica___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* QAbstractItemModelReplica___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QAbstractItemModelReplica___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QAbstractItemModelReplica___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct QtRemoteObjects_PackedList QAbstractItemModelReplica___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue249128 = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtRemoteObjects_PackedList { tmpValue249128, tmpValue249128->size() }; });
}

void* QAbstractItemModelReplica___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModelReplica___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QAbstractItemModelReplica___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* QAbstractItemModelReplica___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModelReplica___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QAbstractItemModelReplica___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* QAbstractItemModelReplica___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModelReplica___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QAbstractItemModelReplica___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QAbstractItemModelReplica___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModelReplica___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QAbstractItemModelReplica___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QAbstractItemModelReplica___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModelReplica___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QAbstractItemModelReplica___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QAbstractItemModelReplica___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QAbstractItemModelReplica___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QAbstractItemModelReplica___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct QtRemoteObjects_PackedList QAbstractItemModelReplica___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue249128 = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtRemoteObjects_PackedList { tmpValue249128, tmpValue249128->size() }; });
}

int QAbstractItemModelReplica_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractItemModelReplica_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QAbstractItemModelReplica_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QAbstractItemModelReplica_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractItemModelReplica_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QAbstractItemModelReplica_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QAbstractItemModelReplica_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractItemModelReplica_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QAbstractItemModelReplica_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QAbstractItemModelReplica_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractItemModelReplica_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QAbstractItemModelReplica_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* QAbstractItemModelReplica___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractItemModelReplica___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QAbstractItemModelReplica___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QAbstractItemModelReplica___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModelReplica___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QAbstractItemModelReplica___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QAbstractItemModelReplica___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractItemModelReplica___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QAbstractItemModelReplica___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAbstractItemModelReplica___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractItemModelReplica___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QAbstractItemModelReplica___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAbstractItemModelReplica_BuddyDefault(void* ptr, void* index)
{
		return new QModelIndex(static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::buddy(*static_cast<QModelIndex*>(index)));
}

char QAbstractItemModelReplica_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModelReplica_CanFetchMoreDefault(void* ptr, void* parent)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int QAbstractItemModelReplica_ColumnCount(void* ptr, void* parent)
{
	return static_cast<QAbstractItemModelReplica*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

int QAbstractItemModelReplica_ColumnCountDefault(void* ptr, void* parent)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::columnCount(*static_cast<QModelIndex*>(parent));
}

void* QAbstractItemModelReplica_Data(void* ptr, void* index, int role)
{
	return new QVariant(static_cast<QAbstractItemModelReplica*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

void* QAbstractItemModelReplica_DataDefault(void* ptr, void* index, int role)
{
		return new QVariant(static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::data(*static_cast<QModelIndex*>(index), role));
}

char QAbstractItemModelReplica_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

void QAbstractItemModelReplica_FetchMoreDefault(void* ptr, void* parent)
{
		static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::fetchMore(*static_cast<QModelIndex*>(parent));
}

long long QAbstractItemModelReplica_FlagsDefault(void* ptr, void* index)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::flags(*static_cast<QModelIndex*>(index));
}

char QAbstractItemModelReplica_HasChildrenDefault(void* ptr, void* parent)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* QAbstractItemModelReplica_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
		return new QVariant(static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QAbstractItemModelReplica_Index(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<QAbstractItemModelReplica*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* QAbstractItemModelReplica_IndexDefault(void* ptr, int row, int column, void* parent)
{
		return new QModelIndex(static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::index(row, column, *static_cast<QModelIndex*>(parent)));
}

char QAbstractItemModelReplica_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModelReplica_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct QtRemoteObjects_PackedList QAbstractItemModelReplica_ItemDataDefault(void* ptr, void* index)
{
		return ({ QMap<int, QVariant>* tmpValue348dfc = new QMap<int, QVariant>(static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::itemData(*static_cast<QModelIndex*>(index))); QtRemoteObjects_PackedList { tmpValue348dfc, tmpValue348dfc->size() }; });
}

struct QtRemoteObjects_PackedList QAbstractItemModelReplica_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
		return ({ QList<QModelIndex>* tmpValueeeb7cc = new QList<QModelIndex>(static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtRemoteObjects_PackedList { tmpValueeeb7cc, tmpValueeeb7cc->size() }; });
}

void* QAbstractItemModelReplica_MimeDataDefault(void* ptr, void* indexes)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct QtRemoteObjects_PackedString QAbstractItemModelReplica_MimeTypesDefault(void* ptr)
{
		return ({ QByteArray* t728134 = new QByteArray(static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::mimeTypes().join("¡¦!").toUtf8()); QtRemoteObjects_PackedString { const_cast<char*>(t728134->prepend("WHITESPACE").constData()+10), t728134->size()-10, t728134 }; });
}

char QAbstractItemModelReplica_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char QAbstractItemModelReplica_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* QAbstractItemModelReplica_Parent(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QAbstractItemModelReplica*>(ptr)->parent(*static_cast<QModelIndex*>(index)));
}

void* QAbstractItemModelReplica_ParentDefault(void* ptr, void* index)
{
		return new QModelIndex(static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::parent(*static_cast<QModelIndex*>(index)));
}

char QAbstractItemModelReplica_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModelReplica_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void QAbstractItemModelReplica_ResetInternalDataDefault(void* ptr)
{
		static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::resetInternalData();
}

void QAbstractItemModelReplica_RevertDefault(void* ptr)
{
		static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::revert();
}

struct QtRemoteObjects_PackedList QAbstractItemModelReplica_RoleNamesDefault(void* ptr)
{
		return ({ QHash<int, QByteArray>* tmpValue1a6efb = new QHash<int, QByteArray>(static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::roleNames()); QtRemoteObjects_PackedList { tmpValue1a6efb, tmpValue1a6efb->size() }; });
}

int QAbstractItemModelReplica_RowCount(void* ptr, void* parent)
{
	return static_cast<QAbstractItemModelReplica*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QAbstractItemModelReplica_RowCountDefault(void* ptr, void* parent)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::rowCount(*static_cast<QModelIndex*>(parent));
}

char QAbstractItemModelReplica_SetDataDefault(void* ptr, void* index, void* value, int role)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char QAbstractItemModelReplica_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char QAbstractItemModelReplica_SetItemDataDefault(void* ptr, void* index, void* roles)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void* QAbstractItemModelReplica_SiblingDefault(void* ptr, int row, int column, void* index)
{
		return new QModelIndex(static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::sibling(row, column, *static_cast<QModelIndex*>(index)));
}

void QAbstractItemModelReplica_SortDefault(void* ptr, int column, long long order)
{
		static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::sort(column, static_cast<Qt::SortOrder>(order));
}

void* QAbstractItemModelReplica_SpanDefault(void* ptr, void* index)
{
		return ({ QSize tmpValue = static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char QAbstractItemModelReplica_SubmitDefault(void* ptr)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::submit();
}

long long QAbstractItemModelReplica_SupportedDragActionsDefault(void* ptr)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::supportedDragActions();
}

long long QAbstractItemModelReplica_SupportedDropActionsDefault(void* ptr)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::supportedDropActions();
}

void QAbstractItemModelReplica_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractItemModelReplica_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractItemModelReplica_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::customEvent(static_cast<QEvent*>(event));
}

void QAbstractItemModelReplica_DeleteLaterDefault(void* ptr)
{
		static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::deleteLater();
}

void QAbstractItemModelReplica_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAbstractItemModelReplica_EventDefault(void* ptr, void* e)
{
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::event(static_cast<QEvent*>(e));
}

char QAbstractItemModelReplica_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void* QAbstractItemModelReplica_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::metaObject());
}

void QAbstractItemModelReplica_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QAbstractItemModelReplica*>(ptr)->QAbstractItemModelReplica::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQRemoteObjectAbstractPersistedStore: public QRemoteObjectAbstractPersistedStore
{
public:
	MyQRemoteObjectAbstractPersistedStore(QObject *parent = Q_NULLPTR) : QRemoteObjectAbstractPersistedStore(parent) {QRemoteObjectAbstractPersistedStore_QRemoteObjectAbstractPersistedStore_QRegisterMetaType();};
	QList<QVariant> restoreProperties(const QString & repName, const QByteArray & repSig) { QByteArray* t7884ba = new QByteArray(repName.toUtf8()); QtRemoteObjects_PackedString repNamePacked = { const_cast<char*>(t7884ba->prepend("WHITESPACE").constData()+10), t7884ba->size()-10, t7884ba };return ({ QList<QVariant>* tmpP = static_cast<QList<QVariant>*>(callbackQRemoteObjectAbstractPersistedStore_RestoreProperties(this, repNamePacked, const_cast<QByteArray*>(&repSig))); QList<QVariant> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void saveProperties(const QString & repName, const QByteArray & repSig, const QVariantList & values) { QByteArray* t7884ba = new QByteArray(repName.toUtf8()); QtRemoteObjects_PackedString repNamePacked = { const_cast<char*>(t7884ba->prepend("WHITESPACE").constData()+10), t7884ba->size()-10, t7884ba };callbackQRemoteObjectAbstractPersistedStore_SaveProperties(this, repNamePacked, const_cast<QByteArray*>(&repSig), ({ QList<QVariant>* tmpValue048b0c = new QList<QVariant>(values); QtRemoteObjects_PackedList { tmpValue048b0c, tmpValue048b0c->size() }; })); };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectAbstractPersistedStore_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectAbstractPersistedStore_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectAbstractPersistedStore_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectAbstractPersistedStore_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectAbstractPersistedStore_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectAbstractPersistedStore_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRemoteObjectAbstractPersistedStore_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectAbstractPersistedStore_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectAbstractPersistedStore_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQRemoteObjectAbstractPersistedStore_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQRemoteObjectAbstractPersistedStore_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QRemoteObjectAbstractPersistedStore*)
Q_DECLARE_METATYPE(MyQRemoteObjectAbstractPersistedStore*)

int QRemoteObjectAbstractPersistedStore_QRemoteObjectAbstractPersistedStore_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectAbstractPersistedStore*>(); return qRegisterMetaType<MyQRemoteObjectAbstractPersistedStore*>();}

void* QRemoteObjectAbstractPersistedStore_NewQRemoteObjectAbstractPersistedStore(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectAbstractPersistedStore(static_cast<QObject*>(parent));
	}
}

struct QtRemoteObjects_PackedList QRemoteObjectAbstractPersistedStore_RestoreProperties(void* ptr, struct QtRemoteObjects_PackedString repName, void* repSig)
{
	return ({ QList<QVariant>* tmpValue53915f = new QList<QVariant>(static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->restoreProperties(QString::fromUtf8(repName.data, repName.len), *static_cast<QByteArray*>(repSig))); QtRemoteObjects_PackedList { tmpValue53915f, tmpValue53915f->size() }; });
}

void QRemoteObjectAbstractPersistedStore_SaveProperties(void* ptr, struct QtRemoteObjects_PackedString repName, void* repSig, void* values)
{
	static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->saveProperties(QString::fromUtf8(repName.data, repName.len), *static_cast<QByteArray*>(repSig), ({ QList<QVariant>* tmpP = static_cast<QList<QVariant>*>(values); QList<QVariant> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void* QRemoteObjectAbstractPersistedStore___restoreProperties_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectAbstractPersistedStore___restoreProperties_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectAbstractPersistedStore___restoreProperties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QRemoteObjectAbstractPersistedStore___saveProperties_values_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectAbstractPersistedStore___saveProperties_values_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectAbstractPersistedStore___saveProperties_values_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QRemoteObjectAbstractPersistedStore___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectAbstractPersistedStore___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectAbstractPersistedStore___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QRemoteObjectAbstractPersistedStore___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectAbstractPersistedStore___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QRemoteObjectAbstractPersistedStore___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QRemoteObjectAbstractPersistedStore___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectAbstractPersistedStore___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectAbstractPersistedStore___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QRemoteObjectAbstractPersistedStore___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectAbstractPersistedStore___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectAbstractPersistedStore___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QRemoteObjectAbstractPersistedStore_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QRemoteObjectSettingsStore*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectSettingsStore*>(ptr)->QRemoteObjectSettingsStore::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QRemoteObjectAbstractPersistedStore_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QRemoteObjectSettingsStore*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectSettingsStore*>(ptr)->QRemoteObjectSettingsStore::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QRemoteObjectAbstractPersistedStore_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QRemoteObjectSettingsStore*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectSettingsStore*>(ptr)->QRemoteObjectSettingsStore::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::customEvent(static_cast<QEvent*>(event));
	}
}

void QRemoteObjectAbstractPersistedStore_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QRemoteObjectSettingsStore*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectSettingsStore*>(ptr)->QRemoteObjectSettingsStore::deleteLater();
	} else {
		static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::deleteLater();
	}
}

void QRemoteObjectAbstractPersistedStore_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QRemoteObjectSettingsStore*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectSettingsStore*>(ptr)->QRemoteObjectSettingsStore::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

char QRemoteObjectAbstractPersistedStore_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QRemoteObjectSettingsStore*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectSettingsStore*>(ptr)->QRemoteObjectSettingsStore::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::event(static_cast<QEvent*>(e));
	}
}

char QRemoteObjectAbstractPersistedStore_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QRemoteObjectSettingsStore*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
			return static_cast<QRemoteObjectSettingsStore*>(ptr)->QRemoteObjectSettingsStore::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QRemoteObjectSettingsStore*>(ptr)->QRemoteObjectSettingsStore::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else {
		if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
			return static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	}
}

void* QRemoteObjectAbstractPersistedStore_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QRemoteObjectSettingsStore*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectSettingsStore*>(ptr)->QRemoteObjectSettingsStore::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::metaObject());
	}
}

void QRemoteObjectAbstractPersistedStore_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QRemoteObjectSettingsStore*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectSettingsStore*>(ptr)->QRemoteObjectSettingsStore::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QRemoteObjectAbstractPersistedStore*>(ptr)->QRemoteObjectAbstractPersistedStore::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

class MyQRemoteObjectDynamicReplica: public QRemoteObjectDynamicReplica
{
public:
	 ~MyQRemoteObjectDynamicReplica() { callbackQRemoteObjectDynamicReplica_DestroyQRemoteObjectDynamicReplica(this); };
	void Signal_Initialized() { callbackQRemoteObjectReplica_Initialized(this); };
	void setNode(QRemoteObjectNode * node) { callbackQRemoteObjectReplica_SetNode(this, node); };
	void Signal_StateChanged(QRemoteObjectReplica::State state, QRemoteObjectReplica::State oldState) { callbackQRemoteObjectReplica_StateChanged(this, state, oldState); };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectReplica_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectReplica_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectReplica_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectReplica_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectReplica_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectReplica_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRemoteObjectReplica_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectReplica_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectReplica_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQRemoteObjectReplica_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQRemoteObjectReplica_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QRemoteObjectDynamicReplica*)
Q_DECLARE_METATYPE(MyQRemoteObjectDynamicReplica*)

int QRemoteObjectDynamicReplica_QRemoteObjectDynamicReplica_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectDynamicReplica*>(); return qRegisterMetaType<MyQRemoteObjectDynamicReplica*>();}

void QRemoteObjectDynamicReplica_DestroyQRemoteObjectDynamicReplica(void* ptr)
{
	static_cast<QRemoteObjectDynamicReplica*>(ptr)->~QRemoteObjectDynamicReplica();
}

void QRemoteObjectDynamicReplica_DestroyQRemoteObjectDynamicReplicaDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQRemoteObjectHost: public QRemoteObjectHost
{
public:
	MyQRemoteObjectHost(QObject *parent = Q_NULLPTR) : QRemoteObjectHost(parent) {QRemoteObjectHost_QRemoteObjectHost_QRegisterMetaType();};
	MyQRemoteObjectHost(const QUrl &address, const QUrl &registryAddress = QUrl(), QRemoteObjectHostBase::AllowedSchemas allowedSchemas = BuiltInSchemasOnly, QObject *parent = Q_NULLPTR) : QRemoteObjectHost(address, registryAddress, allowedSchemas, parent) {QRemoteObjectHost_QRemoteObjectHost_QRegisterMetaType();};
	MyQRemoteObjectHost(const QUrl &address, QObject *parent) : QRemoteObjectHost(address, parent) {QRemoteObjectHost_QRemoteObjectHost_QRegisterMetaType();};
	QUrl hostUrl() const { return *static_cast<QUrl*>(callbackQRemoteObjectHost_HostUrl(const_cast<void*>(static_cast<const void*>(this)))); };
	bool setHostUrl(const QUrl & hostAddress, QRemoteObjectHostBase::AllowedSchemas allowedSchemas) { return callbackQRemoteObjectHost_SetHostUrl(this, const_cast<QUrl*>(&hostAddress), allowedSchemas) != 0; };
	void setName(const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtRemoteObjects_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };callbackQRemoteObjectNode_SetName(this, namePacked); };
	void Signal_HeartbeatIntervalChanged(int heartbeatInterval) { callbackQRemoteObjectNode_HeartbeatIntervalChanged(this, heartbeatInterval); };
	bool setRegistryUrl(const QUrl & registryAddress) { return callbackQRemoteObjectNode_SetRegistryUrl(this, const_cast<QUrl*>(&registryAddress)) != 0; };
	void timerEvent(QTimerEvent * vqt) { callbackQRemoteObjectNode_TimerEvent(this, vqt); };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectNode_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectNode_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectNode_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectNode_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRemoteObjectNode_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectNode_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectNode_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQRemoteObjectNode_ObjectNameChanged(this, objectNamePacked); };
};

Q_DECLARE_METATYPE(QRemoteObjectHost*)
Q_DECLARE_METATYPE(MyQRemoteObjectHost*)

int QRemoteObjectHost_QRemoteObjectHost_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectHost*>(); return qRegisterMetaType<MyQRemoteObjectHost*>();}

void* QRemoteObjectHost_NewQRemoteObjectHost(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectHost(static_cast<QObject*>(parent));
	}
}

void* QRemoteObjectHost_NewQRemoteObjectHost2(void* address, void* registryAddress, long long allowedSchemas, void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), *static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas), static_cast<QObject*>(parent));
	}
}

void* QRemoteObjectHost_NewQRemoteObjectHost3(void* address, void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectHost(*static_cast<QUrl*>(address), static_cast<QObject*>(parent));
	}
}

void* QRemoteObjectHost_HostUrl(void* ptr)
{
	return new QUrl(static_cast<QRemoteObjectHost*>(ptr)->hostUrl());
}

void* QRemoteObjectHost_HostUrlDefault(void* ptr)
{
		return new QUrl(static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::hostUrl());
}

char QRemoteObjectHost_SetHostUrl(void* ptr, void* hostAddress, long long allowedSchemas)
{
	return static_cast<QRemoteObjectHost*>(ptr)->setHostUrl(*static_cast<QUrl*>(hostAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas));
}

char QRemoteObjectHost_SetHostUrlDefault(void* ptr, void* hostAddress, long long allowedSchemas)
{
		return static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::setHostUrl(*static_cast<QUrl*>(hostAddress), static_cast<QRemoteObjectHostBase::AllowedSchemas>(allowedSchemas));
}

class MyQRemoteObjectHostBase: public QRemoteObjectHostBase
{
public:
	void setName(const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtRemoteObjects_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };callbackQRemoteObjectNode_SetName(this, namePacked); };
	void Signal_HeartbeatIntervalChanged(int heartbeatInterval) { callbackQRemoteObjectNode_HeartbeatIntervalChanged(this, heartbeatInterval); };
	bool setRegistryUrl(const QUrl & registryAddress) { return callbackQRemoteObjectNode_SetRegistryUrl(this, const_cast<QUrl*>(&registryAddress)) != 0; };
	void timerEvent(QTimerEvent * vqt) { callbackQRemoteObjectNode_TimerEvent(this, vqt); };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectNode_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectNode_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectNode_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectNode_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRemoteObjectNode_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectNode_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectNode_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQRemoteObjectNode_ObjectNameChanged(this, objectNamePacked); };
};

Q_DECLARE_METATYPE(QRemoteObjectHostBase*)
Q_DECLARE_METATYPE(MyQRemoteObjectHostBase*)

int QRemoteObjectHostBase_QRemoteObjectHostBase_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectHostBase*>(); return qRegisterMetaType<MyQRemoteObjectHostBase*>();}

void QRemoteObjectHostBase_AddHostSideConnection(void* ptr, void* ioDevice)
{
	static_cast<QRemoteObjectHostBase*>(ptr)->addHostSideConnection(static_cast<QIODevice*>(ioDevice));
}

char QRemoteObjectHostBase_DisableRemoting(void* ptr, void* remoteObject)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(remoteObject))) {
		return static_cast<QRemoteObjectHostBase*>(ptr)->disableRemoting(static_cast<QRemoteObjectPendingCallWatcher*>(remoteObject));
	} else {
		return static_cast<QRemoteObjectHostBase*>(ptr)->disableRemoting(static_cast<QObject*>(remoteObject));
	}
}

char QRemoteObjectHostBase_EnableRemoting2(void* ptr, void* object, struct QtRemoteObjects_PackedString name)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(object))) {
		return static_cast<QRemoteObjectHostBase*>(ptr)->enableRemoting(static_cast<QRemoteObjectPendingCallWatcher*>(object), QString::fromUtf8(name.data, name.len));
	} else {
		return static_cast<QRemoteObjectHostBase*>(ptr)->enableRemoting(static_cast<QObject*>(object), QString::fromUtf8(name.data, name.len));
	}
}

char QRemoteObjectHostBase_EnableRemoting3(void* ptr, void* model, struct QtRemoteObjects_PackedString name, void* roles, void* selectionModel)
{
	return static_cast<QRemoteObjectHostBase*>(ptr)->enableRemoting(static_cast<QAbstractItemModel*>(model), QString::fromUtf8(name.data, name.len), *static_cast<QVector<int>*>(roles), static_cast<QItemSelectionModel*>(selectionModel));
}

int QRemoteObjectHostBase___enableRemoting_roles_atList3(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QRemoteObjectHostBase___enableRemoting_roles_setList3(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QRemoteObjectHostBase___enableRemoting_roles_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

class MyQRemoteObjectNode: public QRemoteObjectNode
{
public:
	MyQRemoteObjectNode(QObject *parent = Q_NULLPTR) : QRemoteObjectNode(parent) {QRemoteObjectNode_QRemoteObjectNode_QRegisterMetaType();};
	MyQRemoteObjectNode(const QUrl &registryAddress, QObject *parent = Q_NULLPTR) : QRemoteObjectNode(registryAddress, parent) {QRemoteObjectNode_QRemoteObjectNode_QRegisterMetaType();};
	void Signal_HeartbeatIntervalChanged(int heartbeatInterval) { callbackQRemoteObjectNode_HeartbeatIntervalChanged(this, heartbeatInterval); };
	void setName(const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtRemoteObjects_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };callbackQRemoteObjectNode_SetName(this, namePacked); };
	bool setRegistryUrl(const QUrl & registryAddress) { return callbackQRemoteObjectNode_SetRegistryUrl(this, const_cast<QUrl*>(&registryAddress)) != 0; };
	void timerEvent(QTimerEvent * vqt) { callbackQRemoteObjectNode_TimerEvent(this, vqt); };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectNode_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectNode_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectNode_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectNode_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRemoteObjectNode_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectNode_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectNode_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQRemoteObjectNode_ObjectNameChanged(this, objectNamePacked); };
};

Q_DECLARE_METATYPE(QRemoteObjectNode*)
Q_DECLARE_METATYPE(MyQRemoteObjectNode*)

int QRemoteObjectNode_QRemoteObjectNode_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectNode*>(); return qRegisterMetaType<MyQRemoteObjectNode*>();}

void* QRemoteObjectNode_NewQRemoteObjectNode(void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectNode(static_cast<QObject*>(parent));
	}
}

void* QRemoteObjectNode_NewQRemoteObjectNode2(void* registryAddress, void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectNode(*static_cast<QUrl*>(registryAddress), static_cast<QObject*>(parent));
	}
}

void* QRemoteObjectNode_AcquireDynamic(void* ptr, struct QtRemoteObjects_PackedString name)
{
	return static_cast<QRemoteObjectNode*>(ptr)->acquireDynamic(QString::fromUtf8(name.data, name.len));
}

void QRemoteObjectNode_AddClientSideConnection(void* ptr, void* ioDevice)
{
	static_cast<QRemoteObjectNode*>(ptr)->addClientSideConnection(static_cast<QIODevice*>(ioDevice));
}

char QRemoteObjectNode_ConnectToNode(void* ptr, void* address)
{
	return static_cast<QRemoteObjectNode*>(ptr)->connectToNode(*static_cast<QUrl*>(address));
}

int QRemoteObjectNode_HeartbeatInterval(void* ptr)
{
	return static_cast<QRemoteObjectNode*>(ptr)->heartbeatInterval();
}

void QRemoteObjectNode_ConnectHeartbeatIntervalChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QRemoteObjectNode*>(ptr), static_cast<void (QRemoteObjectNode::*)(int)>(&QRemoteObjectNode::heartbeatIntervalChanged), static_cast<MyQRemoteObjectNode*>(ptr), static_cast<void (MyQRemoteObjectNode::*)(int)>(&MyQRemoteObjectNode::Signal_HeartbeatIntervalChanged), static_cast<Qt::ConnectionType>(t));
}

void QRemoteObjectNode_DisconnectHeartbeatIntervalChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRemoteObjectNode*>(ptr), static_cast<void (QRemoteObjectNode::*)(int)>(&QRemoteObjectNode::heartbeatIntervalChanged), static_cast<MyQRemoteObjectNode*>(ptr), static_cast<void (MyQRemoteObjectNode::*)(int)>(&MyQRemoteObjectNode::Signal_HeartbeatIntervalChanged));
}

void QRemoteObjectNode_HeartbeatIntervalChanged(void* ptr, int heartbeatInterval)
{
	static_cast<QRemoteObjectNode*>(ptr)->heartbeatIntervalChanged(heartbeatInterval);
}

struct QtRemoteObjects_PackedString QRemoteObjectNode_Instances2(void* ptr, struct QtRemoteObjects_PackedString typeName)
{
	return ({ QByteArray* t1bf5b2 = new QByteArray(static_cast<QRemoteObjectNode*>(ptr)->instances(QString::fromUtf8(typeName.data, typeName.len)).join("¡¦!").toUtf8()); QtRemoteObjects_PackedString { const_cast<char*>(t1bf5b2->prepend("WHITESPACE").constData()+10), t1bf5b2->size()-10, t1bf5b2 }; });
}

long long QRemoteObjectNode_LastError(void* ptr)
{
	return static_cast<QRemoteObjectNode*>(ptr)->lastError();
}

void* QRemoteObjectNode_PersistedStore(void* ptr)
{
	return static_cast<QRemoteObjectNode*>(ptr)->persistedStore();
}

void* QRemoteObjectNode_Registry(void* ptr)
{
	return const_cast<QRemoteObjectRegistry*>(static_cast<QRemoteObjectNode*>(ptr)->registry());
}

void* QRemoteObjectNode_RegistryUrl(void* ptr)
{
	return new QUrl(static_cast<QRemoteObjectNode*>(ptr)->registryUrl());
}

void QRemoteObjectNode_SetHeartbeatInterval(void* ptr, int interval)
{
	static_cast<QRemoteObjectNode*>(ptr)->setHeartbeatInterval(interval);
}

void QRemoteObjectNode_SetName(void* ptr, struct QtRemoteObjects_PackedString name)
{
	static_cast<QRemoteObjectNode*>(ptr)->setName(QString::fromUtf8(name.data, name.len));
}

void QRemoteObjectNode_SetNameDefault(void* ptr, struct QtRemoteObjects_PackedString name)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::setName(QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::setName(QString::fromUtf8(name.data, name.len));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::setName(QString::fromUtf8(name.data, name.len));
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::setName(QString::fromUtf8(name.data, name.len));
	}
}

void QRemoteObjectNode_SetPersistedStore(void* ptr, void* persistedStore)
{
	static_cast<QRemoteObjectNode*>(ptr)->setPersistedStore(static_cast<QRemoteObjectAbstractPersistedStore*>(persistedStore));
}

char QRemoteObjectNode_SetRegistryUrl(void* ptr, void* registryAddress)
{
	return static_cast<QRemoteObjectNode*>(ptr)->setRegistryUrl(*static_cast<QUrl*>(registryAddress));
}

char QRemoteObjectNode_SetRegistryUrlDefault(void* ptr, void* registryAddress)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::setRegistryUrl(*static_cast<QUrl*>(registryAddress));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::setRegistryUrl(*static_cast<QUrl*>(registryAddress));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::setRegistryUrl(*static_cast<QUrl*>(registryAddress));
	} else {
		return static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::setRegistryUrl(*static_cast<QUrl*>(registryAddress));
	}
}

void QRemoteObjectNode_TimerEventDefault(void* ptr, void* vqt)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::timerEvent(static_cast<QTimerEvent*>(vqt));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::timerEvent(static_cast<QTimerEvent*>(vqt));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::timerEvent(static_cast<QTimerEvent*>(vqt));
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::timerEvent(static_cast<QTimerEvent*>(vqt));
	}
}

char QRemoteObjectNode_WaitForRegistry(void* ptr, int timeout)
{
	return static_cast<QRemoteObjectNode*>(ptr)->waitForRegistry(timeout);
}

int QRemoteObjectNode___acquireModel_rolesHint_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QRemoteObjectNode___acquireModel_rolesHint_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QRemoteObjectNode___acquireModel_rolesHint_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* QRemoteObjectNode___persistProperties_props_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectNode___persistProperties_props_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectNode___persistProperties_props_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QRemoteObjectNode___retrieveProperties_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectNode___retrieveProperties_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectNode___retrieveProperties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QRemoteObjectNode___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectNode___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectNode___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QRemoteObjectNode___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectNode___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QRemoteObjectNode___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QRemoteObjectNode___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectNode___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectNode___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QRemoteObjectNode___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectNode___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectNode___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QRemoteObjectNode_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QRemoteObjectNode_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QRemoteObjectNode_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::customEvent(static_cast<QEvent*>(event));
	}
}

void QRemoteObjectNode_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::deleteLater();
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::deleteLater();
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::deleteLater();
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::deleteLater();
	}
}

void QRemoteObjectNode_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

char QRemoteObjectNode_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::event(static_cast<QEvent*>(e));
	}
}

char QRemoteObjectNode_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
			return static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
			return static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
			return static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else {
		if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
			return static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	}
}

void* QRemoteObjectNode_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QRemoteObjectRegistryHost*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectRegistryHost*>(ptr)->QRemoteObjectRegistryHost::metaObject());
	} else if (dynamic_cast<QRemoteObjectHost*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectHost*>(ptr)->QRemoteObjectHost::metaObject());
	} else if (dynamic_cast<QRemoteObjectHostBase*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectHostBase*>(ptr)->QRemoteObjectHostBase::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectNode*>(ptr)->QRemoteObjectNode::metaObject());
	}
}

class MyQRemoteObjectPendingCallWatcher: public QRemoteObjectPendingCallWatcher
{
public:
	void childEvent(QChildEvent * event) { callbackQRemoteObjectPendingCallWatcher_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectPendingCallWatcher_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectPendingCallWatcher_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectPendingCallWatcher_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectPendingCallWatcher_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectPendingCallWatcher_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRemoteObjectPendingCallWatcher_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectPendingCallWatcher_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectPendingCallWatcher_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQRemoteObjectPendingCallWatcher_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQRemoteObjectPendingCallWatcher_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QRemoteObjectPendingCallWatcher*)
Q_DECLARE_METATYPE(MyQRemoteObjectPendingCallWatcher*)

int QRemoteObjectPendingCallWatcher_QRemoteObjectPendingCallWatcher_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectPendingCallWatcher*>(); return qRegisterMetaType<MyQRemoteObjectPendingCallWatcher*>();}

void* QRemoteObjectPendingCallWatcher___children_atList(void* ptr, int i)
{
		return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectPendingCallWatcher___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectPendingCallWatcher___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject *>();
}

void* QRemoteObjectPendingCallWatcher___dynamicPropertyNames_atList(void* ptr, int i)
{
		return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectPendingCallWatcher___dynamicPropertyNames_setList(void* ptr, void* i)
{
		static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QRemoteObjectPendingCallWatcher___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QByteArray>();
}

void* QRemoteObjectPendingCallWatcher___findChildren_atList(void* ptr, int i)
{
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectPendingCallWatcher___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectPendingCallWatcher___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>();
}

void* QRemoteObjectPendingCallWatcher___findChildren_atList3(void* ptr, int i)
{
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectPendingCallWatcher___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectPendingCallWatcher___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>();
}

void QRemoteObjectPendingCallWatcher_ChildEvent(void* ptr, void* event)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QRemoteObjectPendingCallWatcher_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::childEvent(static_cast<QChildEvent*>(event));
}

void QRemoteObjectPendingCallWatcher_ConnectNotify(void* ptr, void* sign)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRemoteObjectPendingCallWatcher_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRemoteObjectPendingCallWatcher_CustomEvent(void* ptr, void* event)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QRemoteObjectPendingCallWatcher_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::customEvent(static_cast<QEvent*>(event));
}

void QRemoteObjectPendingCallWatcher_DeleteLater(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QRemoteObjectPendingCallWatcher*>(ptr), "deleteLater");
}

void QRemoteObjectPendingCallWatcher_DeleteLaterDefault(void* ptr)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::deleteLater();
}

void QRemoteObjectPendingCallWatcher_DisconnectNotify(void* ptr, void* sign)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRemoteObjectPendingCallWatcher_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QRemoteObjectPendingCallWatcher_Event(void* ptr, void* e)
{
		return static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->event(static_cast<QEvent*>(e));
}

char QRemoteObjectPendingCallWatcher_EventDefault(void* ptr, void* e)
{
		return static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::event(static_cast<QEvent*>(e));
}

char QRemoteObjectPendingCallWatcher_EventFilter(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
		return static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

char QRemoteObjectPendingCallWatcher_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
		return static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void* QRemoteObjectPendingCallWatcher_MetaObject(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->metaObject());
}

void* QRemoteObjectPendingCallWatcher_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::metaObject());
}

void QRemoteObjectPendingCallWatcher_TimerEvent(void* ptr, void* event)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QRemoteObjectPendingCallWatcher_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QRemoteObjectPendingCallWatcher*>(ptr)->QRemoteObjectPendingCallWatcher::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQRemoteObjectRegistry: public QRemoteObjectRegistry
{
public:
	 ~MyQRemoteObjectRegistry() { callbackQRemoteObjectRegistry_DestroyQRemoteObjectRegistry(this); };
	void Signal_Initialized() { callbackQRemoteObjectReplica_Initialized(this); };
	void setNode(QRemoteObjectNode * node) { callbackQRemoteObjectReplica_SetNode(this, node); };
	void Signal_StateChanged(QRemoteObjectReplica::State state, QRemoteObjectReplica::State oldState) { callbackQRemoteObjectReplica_StateChanged(this, state, oldState); };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectReplica_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectReplica_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectReplica_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectReplica_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectReplica_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectReplica_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRemoteObjectReplica_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectReplica_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectReplica_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQRemoteObjectReplica_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQRemoteObjectReplica_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QRemoteObjectRegistry*)
Q_DECLARE_METATYPE(MyQRemoteObjectRegistry*)

int QRemoteObjectRegistry_QRemoteObjectRegistry_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectRegistry*>(); return qRegisterMetaType<MyQRemoteObjectRegistry*>();}

void QRemoteObjectRegistry_DestroyQRemoteObjectRegistry(void* ptr)
{
	static_cast<QRemoteObjectRegistry*>(ptr)->~QRemoteObjectRegistry();
}

void QRemoteObjectRegistry_DestroyQRemoteObjectRegistryDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQRemoteObjectRegistryHost: public QRemoteObjectRegistryHost
{
public:
	MyQRemoteObjectRegistryHost(const QUrl &registryAddress = QUrl(), QObject *parent = Q_NULLPTR) : QRemoteObjectRegistryHost(registryAddress, parent) {QRemoteObjectRegistryHost_QRemoteObjectRegistryHost_QRegisterMetaType();};
	bool setRegistryUrl(const QUrl & registryUrl) { return callbackQRemoteObjectNode_SetRegistryUrl(this, const_cast<QUrl*>(&registryUrl)) != 0; };
	void setName(const QString & name) { QByteArray* t6ae999 = new QByteArray(name.toUtf8()); QtRemoteObjects_PackedString namePacked = { const_cast<char*>(t6ae999->prepend("WHITESPACE").constData()+10), t6ae999->size()-10, t6ae999 };callbackQRemoteObjectNode_SetName(this, namePacked); };
	void Signal_HeartbeatIntervalChanged(int heartbeatInterval) { callbackQRemoteObjectNode_HeartbeatIntervalChanged(this, heartbeatInterval); };
	void timerEvent(QTimerEvent * vqt) { callbackQRemoteObjectNode_TimerEvent(this, vqt); };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectNode_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectNode_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectNode_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectNode_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectNode_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRemoteObjectNode_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectNode_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectNode_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQRemoteObjectNode_ObjectNameChanged(this, objectNamePacked); };
};

Q_DECLARE_METATYPE(QRemoteObjectRegistryHost*)
Q_DECLARE_METATYPE(MyQRemoteObjectRegistryHost*)

int QRemoteObjectRegistryHost_QRemoteObjectRegistryHost_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectRegistryHost*>(); return qRegisterMetaType<MyQRemoteObjectRegistryHost*>();}

void* QRemoteObjectRegistryHost_NewQRemoteObjectRegistryHost(void* registryAddress, void* parent)
{
	if (dynamic_cast<QAudioSystemPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QAudioSystemPlugin*>(parent));
	} else if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QMediaServiceProviderPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QMediaServiceProviderPlugin*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QRemoteObjectPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QScriptExtensionPlugin*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QScriptExtensionPlugin*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QWindow*>(parent));
	} else {
		return new MyQRemoteObjectRegistryHost(*static_cast<QUrl*>(registryAddress), static_cast<QObject*>(parent));
	}
}

class MyQRemoteObjectReplica: public QRemoteObjectReplica
{
public:
	void Signal_Initialized() { callbackQRemoteObjectReplica_Initialized(this); };
	void setNode(QRemoteObjectNode * node) { callbackQRemoteObjectReplica_SetNode(this, node); };
	void Signal_StateChanged(QRemoteObjectReplica::State state, QRemoteObjectReplica::State oldState) { callbackQRemoteObjectReplica_StateChanged(this, state, oldState); };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectReplica_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectReplica_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectReplica_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectReplica_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectReplica_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectReplica_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRemoteObjectReplica_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectReplica_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectReplica_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQRemoteObjectReplica_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQRemoteObjectReplica_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QRemoteObjectReplica*)
Q_DECLARE_METATYPE(MyQRemoteObjectReplica*)

int QRemoteObjectReplica_QRemoteObjectReplica_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectReplica*>(); return qRegisterMetaType<MyQRemoteObjectReplica*>();}

void QRemoteObjectReplica_ConnectInitialized(void* ptr, long long t)
{
	QObject::connect(static_cast<QRemoteObjectReplica*>(ptr), static_cast<void (QRemoteObjectReplica::*)()>(&QRemoteObjectReplica::initialized), static_cast<MyQRemoteObjectReplica*>(ptr), static_cast<void (MyQRemoteObjectReplica::*)()>(&MyQRemoteObjectReplica::Signal_Initialized), static_cast<Qt::ConnectionType>(t));
}

void QRemoteObjectReplica_DisconnectInitialized(void* ptr)
{
	QObject::disconnect(static_cast<QRemoteObjectReplica*>(ptr), static_cast<void (QRemoteObjectReplica::*)()>(&QRemoteObjectReplica::initialized), static_cast<MyQRemoteObjectReplica*>(ptr), static_cast<void (MyQRemoteObjectReplica::*)()>(&MyQRemoteObjectReplica::Signal_Initialized));
}

void QRemoteObjectReplica_Initialized(void* ptr)
{
	static_cast<QRemoteObjectReplica*>(ptr)->initialized();
}

char QRemoteObjectReplica_IsInitialized(void* ptr)
{
	return static_cast<QRemoteObjectReplica*>(ptr)->isInitialized();
}

char QRemoteObjectReplica_IsReplicaValid(void* ptr)
{
	return static_cast<QRemoteObjectReplica*>(ptr)->isReplicaValid();
}

void* QRemoteObjectReplica_Node(void* ptr)
{
	return static_cast<QRemoteObjectReplica*>(ptr)->node();
}

void QRemoteObjectReplica_SetNode(void* ptr, void* node)
{
	static_cast<QRemoteObjectReplica*>(ptr)->setNode(static_cast<QRemoteObjectNode*>(node));
}

void QRemoteObjectReplica_SetNodeDefault(void* ptr, void* node)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::setNode(static_cast<QRemoteObjectNode*>(node));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::setNode(static_cast<QRemoteObjectNode*>(node));
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::setNode(static_cast<QRemoteObjectNode*>(node));
	}
}

long long QRemoteObjectReplica_State(void* ptr)
{
	return static_cast<QRemoteObjectReplica*>(ptr)->state();
}

void QRemoteObjectReplica_ConnectStateChanged(void* ptr, long long t)
{
	qRegisterMetaType<QRemoteObjectReplica::State>();
	qRegisterMetaType<QRemoteObjectReplica::State>();
	QObject::connect(static_cast<QRemoteObjectReplica*>(ptr), static_cast<void (QRemoteObjectReplica::*)(QRemoteObjectReplica::State, QRemoteObjectReplica::State)>(&QRemoteObjectReplica::stateChanged), static_cast<MyQRemoteObjectReplica*>(ptr), static_cast<void (MyQRemoteObjectReplica::*)(QRemoteObjectReplica::State, QRemoteObjectReplica::State)>(&MyQRemoteObjectReplica::Signal_StateChanged), static_cast<Qt::ConnectionType>(t));
}

void QRemoteObjectReplica_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRemoteObjectReplica*>(ptr), static_cast<void (QRemoteObjectReplica::*)(QRemoteObjectReplica::State, QRemoteObjectReplica::State)>(&QRemoteObjectReplica::stateChanged), static_cast<MyQRemoteObjectReplica*>(ptr), static_cast<void (MyQRemoteObjectReplica::*)(QRemoteObjectReplica::State, QRemoteObjectReplica::State)>(&MyQRemoteObjectReplica::Signal_StateChanged));
}

void QRemoteObjectReplica_StateChanged(void* ptr, long long state, long long oldState)
{
	static_cast<QRemoteObjectReplica*>(ptr)->stateChanged(static_cast<QRemoteObjectReplica::State>(state), static_cast<QRemoteObjectReplica::State>(oldState));
}

char QRemoteObjectReplica_WaitForSource(void* ptr, int timeout)
{
	return static_cast<QRemoteObjectReplica*>(ptr)->waitForSource(timeout);
}

void* QRemoteObjectReplica___persistProperties_props_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectReplica___persistProperties_props_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectReplica___persistProperties_props_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QRemoteObjectReplica___retrieveProperties_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectReplica___retrieveProperties_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectReplica___retrieveProperties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QRemoteObjectReplica___send_args_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectReplica___send_args_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectReplica___send_args_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QRemoteObjectReplica___sendWithReply_args_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectReplica___sendWithReply_args_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectReplica___sendWithReply_args_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QRemoteObjectReplica___setProperties_properties_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectReplica___setProperties_properties_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QRemoteObjectReplica___setProperties_properties_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QRemoteObjectReplica___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectReplica___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectReplica___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QRemoteObjectReplica___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QRemoteObjectReplica___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QRemoteObjectReplica___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QRemoteObjectReplica___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectReplica___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectReplica___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QRemoteObjectReplica___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QRemoteObjectReplica___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QRemoteObjectPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QRemoteObjectReplica___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QRemoteObjectReplica_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QRemoteObjectReplica_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QRemoteObjectReplica_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::customEvent(static_cast<QEvent*>(event));
	}
}

void QRemoteObjectReplica_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::deleteLater();
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::deleteLater();
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::deleteLater();
	}
}

void QRemoteObjectReplica_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

char QRemoteObjectReplica_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		return static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::event(static_cast<QEvent*>(e));
	}
}

char QRemoteObjectReplica_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
			return static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
			return static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else {
		if (dynamic_cast<QRemoteObjectPendingCallWatcher*>(static_cast<QObject*>(watched))) {
			return static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::eventFilter(static_cast<QRemoteObjectPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	}
}

void* QRemoteObjectReplica_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::metaObject());
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::metaObject());
	}
}

void QRemoteObjectReplica_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QRemoteObjectRegistry*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectRegistry*>(ptr)->QRemoteObjectRegistry::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QRemoteObjectDynamicReplica*>(static_cast<QObject*>(ptr))) {
		static_cast<QRemoteObjectDynamicReplica*>(ptr)->QRemoteObjectDynamicReplica::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QRemoteObjectReplica*>(ptr)->QRemoteObjectReplica::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

class MyQRemoteObjectSettingsStore: public QRemoteObjectSettingsStore
{
public:
	QList<QVariant> restoreProperties(const QString & repName, const QByteArray & repSig) { QByteArray* t7884ba = new QByteArray(repName.toUtf8()); QtRemoteObjects_PackedString repNamePacked = { const_cast<char*>(t7884ba->prepend("WHITESPACE").constData()+10), t7884ba->size()-10, t7884ba };return ({ QList<QVariant>* tmpP = static_cast<QList<QVariant>*>(callbackQRemoteObjectSettingsStore_RestoreProperties(this, repNamePacked, const_cast<QByteArray*>(&repSig))); QList<QVariant> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void saveProperties(const QString & repName, const QByteArray & repSig, const QVariantList & values) { QByteArray* t7884ba = new QByteArray(repName.toUtf8()); QtRemoteObjects_PackedString repNamePacked = { const_cast<char*>(t7884ba->prepend("WHITESPACE").constData()+10), t7884ba->size()-10, t7884ba };callbackQRemoteObjectSettingsStore_SaveProperties(this, repNamePacked, const_cast<QByteArray*>(&repSig), ({ QList<QVariant>* tmpValue048b0c = new QList<QVariant>(values); QtRemoteObjects_PackedList { tmpValue048b0c, tmpValue048b0c->size() }; })); };
	void childEvent(QChildEvent * event) { callbackQRemoteObjectAbstractPersistedStore_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRemoteObjectAbstractPersistedStore_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQRemoteObjectAbstractPersistedStore_CustomEvent(this, event); };
	void deleteLater() { callbackQRemoteObjectAbstractPersistedStore_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQRemoteObjectAbstractPersistedStore_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRemoteObjectAbstractPersistedStore_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQRemoteObjectAbstractPersistedStore_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRemoteObjectAbstractPersistedStore_EventFilter(this, watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRemoteObjectAbstractPersistedStore_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray* taa2c4f = new QByteArray(objectName.toUtf8()); QtRemoteObjects_PackedString objectNamePacked = { const_cast<char*>(taa2c4f->prepend("WHITESPACE").constData()+10), taa2c4f->size()-10, taa2c4f };callbackQRemoteObjectAbstractPersistedStore_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQRemoteObjectAbstractPersistedStore_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(QRemoteObjectSettingsStore*)
Q_DECLARE_METATYPE(MyQRemoteObjectSettingsStore*)

int QRemoteObjectSettingsStore_QRemoteObjectSettingsStore_QRegisterMetaType(){qRegisterMetaType<QRemoteObjectSettingsStore*>(); return qRegisterMetaType<MyQRemoteObjectSettingsStore*>();}

struct QtRemoteObjects_PackedList QRemoteObjectSettingsStore_RestoreProperties(void* ptr, struct QtRemoteObjects_PackedString repName, void* repSig)
{
	return ({ QList<QVariant>* tmpValue5127fd = new QList<QVariant>(static_cast<QRemoteObjectSettingsStore*>(ptr)->restoreProperties(QString::fromUtf8(repName.data, repName.len), *static_cast<QByteArray*>(repSig))); QtRemoteObjects_PackedList { tmpValue5127fd, tmpValue5127fd->size() }; });
}

struct QtRemoteObjects_PackedList QRemoteObjectSettingsStore_RestorePropertiesDefault(void* ptr, struct QtRemoteObjects_PackedString repName, void* repSig)
{
		return ({ QList<QVariant>* tmpValue873207 = new QList<QVariant>(static_cast<QRemoteObjectSettingsStore*>(ptr)->QRemoteObjectSettingsStore::restoreProperties(QString::fromUtf8(repName.data, repName.len), *static_cast<QByteArray*>(repSig))); QtRemoteObjects_PackedList { tmpValue873207, tmpValue873207->size() }; });
}

void QRemoteObjectSettingsStore_SaveProperties(void* ptr, struct QtRemoteObjects_PackedString repName, void* repSig, void* values)
{
	static_cast<QRemoteObjectSettingsStore*>(ptr)->saveProperties(QString::fromUtf8(repName.data, repName.len), *static_cast<QByteArray*>(repSig), ({ QList<QVariant>* tmpP = static_cast<QList<QVariant>*>(values); QList<QVariant> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void QRemoteObjectSettingsStore_SavePropertiesDefault(void* ptr, struct QtRemoteObjects_PackedString repName, void* repSig, void* values)
{
		static_cast<QRemoteObjectSettingsStore*>(ptr)->QRemoteObjectSettingsStore::saveProperties(QString::fromUtf8(repName.data, repName.len), *static_cast<QByteArray*>(repSig), ({ QList<QVariant>* tmpP = static_cast<QList<QVariant>*>(values); QList<QVariant> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

