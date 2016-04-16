#define protected public

#include "help.h"
#include "_cgo_export.h"

#include <QAbstractItemDelegate>
#include <QAbstractItemModel>
#include <QAction>
#include <QActionEvent>
#include <QChildEvent>
#include <QCloseEvent>
#include <QContextMenuEvent>
#include <QDrag>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QHelpContentItem>
#include <QHelpContentModel>
#include <QHelpContentWidget>
#include <QHelpEngine>
#include <QHelpEngineCore>
#include <QHelpIndexModel>
#include <QHelpIndexWidget>
#include <QHelpSearchEngine>
#include <QHelpSearchQuery>
#include <QHelpSearchQueryWidget>
#include <QHelpSearchResultWidget>
#include <QHideEvent>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QItemSelection>
#include <QItemSelectionModel>
#include <QKeyEvent>
#include <QMetaObject>
#include <QModelIndex>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPaintEvent>
#include <QPainter>
#include <QPoint>
#include <QRect>
#include <QResizeEvent>
#include <QShowEvent>
#include <QString>
#include <QTabletEvent>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QUrl>
#include <QVariant>
#include <QWheelEvent>
#include <QWidget>

void* QHelpContentItem_Child(void* ptr, int row){
	return static_cast<QHelpContentItem*>(ptr)->child(row);
}

int QHelpContentItem_ChildCount(void* ptr){
	return static_cast<QHelpContentItem*>(ptr)->childCount();
}

int QHelpContentItem_ChildPosition(void* ptr, void* child){
	return static_cast<QHelpContentItem*>(ptr)->childPosition(static_cast<QHelpContentItem*>(child));
}

void* QHelpContentItem_Parent(void* ptr){
	return static_cast<QHelpContentItem*>(ptr)->parent();
}

int QHelpContentItem_Row(void* ptr){
	return static_cast<QHelpContentItem*>(ptr)->row();
}

char* QHelpContentItem_Title(void* ptr){
	return static_cast<QHelpContentItem*>(ptr)->title().toUtf8().data();
}

void* QHelpContentItem_Url(void* ptr){
	return new QUrl(static_cast<QHelpContentItem*>(ptr)->url());
}

void QHelpContentItem_DestroyQHelpContentItem(void* ptr){
	static_cast<QHelpContentItem*>(ptr)->~QHelpContentItem();
}

class MyQHelpContentModel: public QHelpContentModel {
public:
	void Signal_ContentsCreated() { callbackQHelpContentModelContentsCreated(this, this->objectName().toUtf8().data()); };
	void Signal_ContentsCreationStarted() { callbackQHelpContentModelContentsCreationStarted(this, this->objectName().toUtf8().data()); };
	void revert() { if (!callbackQHelpContentModelRevert(this, this->objectName().toUtf8().data())) { QHelpContentModel::revert(); }; };
	void sort(int column, Qt::SortOrder order) { callbackQHelpContentModelSort(this, this->objectName().toUtf8().data(), column, order); };
	void timerEvent(QTimerEvent * event) { callbackQHelpContentModelTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQHelpContentModelChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQHelpContentModelCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QHelpContentModel_ColumnCount(void* ptr, void* parent){
	return static_cast<QHelpContentModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

void* QHelpContentModel_ContentItemAt(void* ptr, void* index){
	return static_cast<QHelpContentModel*>(ptr)->contentItemAt(*static_cast<QModelIndex*>(index));
}

void QHelpContentModel_ConnectContentsCreated(void* ptr){
	QObject::connect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreated), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreated));;
}

void QHelpContentModel_DisconnectContentsCreated(void* ptr){
	QObject::disconnect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreated), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreated));;
}

void QHelpContentModel_ContentsCreated(void* ptr){
	static_cast<QHelpContentModel*>(ptr)->contentsCreated();
}

void QHelpContentModel_ConnectContentsCreationStarted(void* ptr){
	QObject::connect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreationStarted), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreationStarted));;
}

void QHelpContentModel_DisconnectContentsCreationStarted(void* ptr){
	QObject::disconnect(static_cast<QHelpContentModel*>(ptr), static_cast<void (QHelpContentModel::*)()>(&QHelpContentModel::contentsCreationStarted), static_cast<MyQHelpContentModel*>(ptr), static_cast<void (MyQHelpContentModel::*)()>(&MyQHelpContentModel::Signal_ContentsCreationStarted));;
}

void QHelpContentModel_ContentsCreationStarted(void* ptr){
	static_cast<QHelpContentModel*>(ptr)->contentsCreationStarted();
}

void QHelpContentModel_CreateContents(void* ptr, char* customFilterName){
	static_cast<QHelpContentModel*>(ptr)->createContents(QString(customFilterName));
}

void* QHelpContentModel_Data(void* ptr, void* index, int role){
	return new QVariant(static_cast<QHelpContentModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

int QHelpContentModel_IsCreatingContents(void* ptr){
	return static_cast<QHelpContentModel*>(ptr)->isCreatingContents();
}

int QHelpContentModel_RowCount(void* ptr, void* parent){
	return static_cast<QHelpContentModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

void QHelpContentModel_DestroyQHelpContentModel(void* ptr){
	static_cast<QHelpContentModel*>(ptr)->~QHelpContentModel();
}

void QHelpContentModel_Revert(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQHelpContentModel*>(ptr), "revert");
}

void QHelpContentModel_RevertDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpContentModel*>(ptr), "revert");
}

void QHelpContentModel_Sort(void* ptr, int column, int order){
	static_cast<MyQHelpContentModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void QHelpContentModel_SortDefault(void* ptr, int column, int order){
	static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void QHelpContentModel_TimerEvent(void* ptr, void* event){
	static_cast<MyQHelpContentModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpContentModel_TimerEventDefault(void* ptr, void* event){
	static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpContentModel_ChildEvent(void* ptr, void* event){
	static_cast<MyQHelpContentModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHelpContentModel_ChildEventDefault(void* ptr, void* event){
	static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpContentModel_CustomEvent(void* ptr, void* event){
	static_cast<MyQHelpContentModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHelpContentModel_CustomEventDefault(void* ptr, void* event){
	static_cast<QHelpContentModel*>(ptr)->QHelpContentModel::customEvent(static_cast<QEvent*>(event));
}

class MyQHelpContentWidget: public QHelpContentWidget {
public:
	void Signal_LinkActivated(const QUrl & link) { callbackQHelpContentWidgetLinkActivated(this, this->objectName().toUtf8().data(), new QUrl(link)); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQHelpContentWidgetDragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void keyPressEvent(QKeyEvent * event) { callbackQHelpContentWidgetKeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyboardSearch(const QString & search) { callbackQHelpContentWidgetKeyboardSearch(this, this->objectName().toUtf8().data(), search.toUtf8().data()); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQHelpContentWidgetMouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQHelpContentWidgetMouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQHelpContentWidgetMousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQHelpContentWidgetMouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void paintEvent(QPaintEvent * event) { callbackQHelpContentWidgetPaintEvent(this, this->objectName().toUtf8().data(), event); };
	void reset() { if (!callbackQHelpContentWidgetReset(this, this->objectName().toUtf8().data())) { QHelpContentWidget::reset(); }; };
	void scrollContentsBy(int dx, int dy) { callbackQHelpContentWidgetScrollContentsBy(this, this->objectName().toUtf8().data(), dx, dy); };
	void selectAll() { if (!callbackQHelpContentWidgetSelectAll(this, this->objectName().toUtf8().data())) { QHelpContentWidget::selectAll(); }; };
	void setModel(QAbstractItemModel * model) { callbackQHelpContentWidgetSetModel(this, this->objectName().toUtf8().data(), model); };
	void setSelection(const QRect & rect, QItemSelectionModel::SelectionFlags command) { callbackQHelpContentWidgetSetSelection(this, this->objectName().toUtf8().data(), new QRect(static_cast<QRect>(rect).x(), static_cast<QRect>(rect).y(), static_cast<QRect>(rect).width(), static_cast<QRect>(rect).height()), command); };
	void setSelectionModel(QItemSelectionModel * selectionModel) { callbackQHelpContentWidgetSetSelectionModel(this, this->objectName().toUtf8().data(), selectionModel); };
	void timerEvent(QTimerEvent * event) { callbackQHelpContentWidgetTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void updateGeometries() { if (!callbackQHelpContentWidgetUpdateGeometries(this, this->objectName().toUtf8().data())) { QHelpContentWidget::updateGeometries(); }; };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQHelpContentWidgetDragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void closeEditor(QWidget * editor, QAbstractItemDelegate::EndEditHint hint) { if (!callbackQHelpContentWidgetCloseEditor(this, this->objectName().toUtf8().data(), editor, hint)) { QHelpContentWidget::closeEditor(editor, hint); }; };
	void commitData(QWidget * editor) { if (!callbackQHelpContentWidgetCommitData(this, this->objectName().toUtf8().data(), editor)) { QHelpContentWidget::commitData(editor); }; };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQHelpContentWidgetDragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQHelpContentWidgetDropEvent(this, this->objectName().toUtf8().data(), event); };
	void editorDestroyed(QObject * editor) { if (!callbackQHelpContentWidgetEditorDestroyed(this, this->objectName().toUtf8().data(), editor)) { QHelpContentWidget::editorDestroyed(editor); }; };
	void focusInEvent(QFocusEvent * event) { callbackQHelpContentWidgetFocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQHelpContentWidgetFocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQHelpContentWidgetInputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	void resizeEvent(QResizeEvent * event) { callbackQHelpContentWidgetResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void startDrag(Qt::DropActions supportedActions) { callbackQHelpContentWidgetStartDrag(this, this->objectName().toUtf8().data(), supportedActions); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQHelpContentWidgetContextMenuEvent(this, this->objectName().toUtf8().data(), e); };
	void setupViewport(QWidget * viewport) { callbackQHelpContentWidgetSetupViewport(this, this->objectName().toUtf8().data(), viewport); };
	void wheelEvent(QWheelEvent * e) { callbackQHelpContentWidgetWheelEvent(this, this->objectName().toUtf8().data(), e); };
	void changeEvent(QEvent * ev) { callbackQHelpContentWidgetChangeEvent(this, this->objectName().toUtf8().data(), ev); };
	void actionEvent(QActionEvent * event) { callbackQHelpContentWidgetActionEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQHelpContentWidgetEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void hideEvent(QHideEvent * event) { callbackQHelpContentWidgetHideEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQHelpContentWidgetLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void moveEvent(QMoveEvent * event) { callbackQHelpContentWidgetMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void setVisible(bool visible) { if (!callbackQHelpContentWidgetSetVisible(this, this->objectName().toUtf8().data(), visible)) { QHelpContentWidget::setVisible(visible); }; };
	void showEvent(QShowEvent * event) { callbackQHelpContentWidgetShowEvent(this, this->objectName().toUtf8().data(), event); };
	void closeEvent(QCloseEvent * event) { callbackQHelpContentWidgetCloseEvent(this, this->objectName().toUtf8().data(), event); };
	void initPainter(QPainter * painter) const { callbackQHelpContentWidgetInitPainter(const_cast<MyQHelpContentWidget*>(this), this->objectName().toUtf8().data(), painter); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQHelpContentWidgetKeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void tabletEvent(QTabletEvent * event) { callbackQHelpContentWidgetTabletEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQHelpContentWidgetChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQHelpContentWidgetCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QHelpContentWidget_ConnectLinkActivated(void* ptr){
	QObject::connect(static_cast<QHelpContentWidget*>(ptr), static_cast<void (QHelpContentWidget::*)(const QUrl &)>(&QHelpContentWidget::linkActivated), static_cast<MyQHelpContentWidget*>(ptr), static_cast<void (MyQHelpContentWidget::*)(const QUrl &)>(&MyQHelpContentWidget::Signal_LinkActivated));;
}

void QHelpContentWidget_DisconnectLinkActivated(void* ptr){
	QObject::disconnect(static_cast<QHelpContentWidget*>(ptr), static_cast<void (QHelpContentWidget::*)(const QUrl &)>(&QHelpContentWidget::linkActivated), static_cast<MyQHelpContentWidget*>(ptr), static_cast<void (MyQHelpContentWidget::*)(const QUrl &)>(&MyQHelpContentWidget::Signal_LinkActivated));;
}

void QHelpContentWidget_LinkActivated(void* ptr, void* link){
	static_cast<QHelpContentWidget*>(ptr)->linkActivated(*static_cast<QUrl*>(link));
}

void QHelpContentWidget_DragMoveEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QHelpContentWidget_DragMoveEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QHelpContentWidget_KeyPressEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpContentWidget_KeyPressEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpContentWidget_KeyboardSearch(void* ptr, char* search){
	static_cast<MyQHelpContentWidget*>(ptr)->keyboardSearch(QString(search));
}

void QHelpContentWidget_KeyboardSearchDefault(void* ptr, char* search){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::keyboardSearch(QString(search));
}

void QHelpContentWidget_MouseDoubleClickEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpContentWidget_MouseDoubleClickEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpContentWidget_MouseMoveEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QHelpContentWidget_MouseMoveEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QHelpContentWidget_MousePressEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpContentWidget_MousePressEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpContentWidget_MouseReleaseEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QHelpContentWidget_MouseReleaseEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QHelpContentWidget_PaintEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QHelpContentWidget_PaintEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void QHelpContentWidget_Reset(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQHelpContentWidget*>(ptr), "reset");
}

void QHelpContentWidget_ResetDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "reset");
}

void QHelpContentWidget_ScrollContentsBy(void* ptr, int dx, int dy){
	static_cast<MyQHelpContentWidget*>(ptr)->scrollContentsBy(dx, dy);
}

void QHelpContentWidget_ScrollContentsByDefault(void* ptr, int dx, int dy){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::scrollContentsBy(dx, dy);
}

void QHelpContentWidget_SelectAll(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQHelpContentWidget*>(ptr), "selectAll");
}

void QHelpContentWidget_SelectAllDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "selectAll");
}

void QHelpContentWidget_SetModel(void* ptr, void* model){
	static_cast<MyQHelpContentWidget*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QHelpContentWidget_SetModelDefault(void* ptr, void* model){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setModel(static_cast<QAbstractItemModel*>(model));
}

void QHelpContentWidget_SetSelection(void* ptr, void* rect, int command){
	static_cast<MyQHelpContentWidget*>(ptr)->setSelection(*static_cast<QRect*>(rect), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QHelpContentWidget_SetSelectionDefault(void* ptr, void* rect, int command){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setSelection(*static_cast<QRect*>(rect), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QHelpContentWidget_SetSelectionModel(void* ptr, void* selectionModel){
	static_cast<MyQHelpContentWidget*>(ptr)->setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

void QHelpContentWidget_SetSelectionModelDefault(void* ptr, void* selectionModel){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

void QHelpContentWidget_TimerEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpContentWidget_TimerEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpContentWidget_UpdateGeometries(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQHelpContentWidget*>(ptr), "updateGeometries");
}

void QHelpContentWidget_UpdateGeometriesDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "updateGeometries");
}

void QHelpContentWidget_DragLeaveEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QHelpContentWidget_DragLeaveEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QHelpContentWidget_CloseEditor(void* ptr, void* editor, int hint){
	QMetaObject::invokeMethod(static_cast<MyQHelpContentWidget*>(ptr), "closeEditor", Q_ARG(QWidget*, static_cast<QWidget*>(editor)), Q_ARG(QAbstractItemDelegate::EndEditHint, static_cast<QAbstractItemDelegate::EndEditHint>(hint)));
}

void QHelpContentWidget_CloseEditorDefault(void* ptr, void* editor, int hint){
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "closeEditor", Q_ARG(QWidget*, static_cast<QWidget*>(editor)), Q_ARG(QAbstractItemDelegate::EndEditHint, static_cast<QAbstractItemDelegate::EndEditHint>(hint)));
}

void QHelpContentWidget_CommitData(void* ptr, void* editor){
	QMetaObject::invokeMethod(static_cast<MyQHelpContentWidget*>(ptr), "commitData", Q_ARG(QWidget*, static_cast<QWidget*>(editor)));
}

void QHelpContentWidget_CommitDataDefault(void* ptr, void* editor){
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "commitData", Q_ARG(QWidget*, static_cast<QWidget*>(editor)));
}

void QHelpContentWidget_DragEnterEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpContentWidget_DragEnterEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpContentWidget_DropEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QHelpContentWidget_DropEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void QHelpContentWidget_EditorDestroyed(void* ptr, void* editor){
	QMetaObject::invokeMethod(static_cast<MyQHelpContentWidget*>(ptr), "editorDestroyed", Q_ARG(QObject*, static_cast<QObject*>(editor)));
}

void QHelpContentWidget_EditorDestroyedDefault(void* ptr, void* editor){
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "editorDestroyed", Q_ARG(QObject*, static_cast<QObject*>(editor)));
}

void QHelpContentWidget_FocusInEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QHelpContentWidget_FocusInEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QHelpContentWidget_FocusOutEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpContentWidget_FocusOutEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpContentWidget_InputMethodEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QHelpContentWidget_InputMethodEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QHelpContentWidget_ResizeEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QHelpContentWidget_ResizeEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QHelpContentWidget_StartDrag(void* ptr, int supportedActions){
	static_cast<MyQHelpContentWidget*>(ptr)->startDrag(static_cast<Qt::DropAction>(supportedActions));
}

void QHelpContentWidget_StartDragDefault(void* ptr, int supportedActions){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::startDrag(static_cast<Qt::DropAction>(supportedActions));
}

void QHelpContentWidget_ContextMenuEvent(void* ptr, void* e){
	static_cast<MyQHelpContentWidget*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void QHelpContentWidget_ContextMenuEventDefault(void* ptr, void* e){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void QHelpContentWidget_SetupViewport(void* ptr, void* viewport){
	static_cast<MyQHelpContentWidget*>(ptr)->setupViewport(static_cast<QWidget*>(viewport));
}

void QHelpContentWidget_SetupViewportDefault(void* ptr, void* viewport){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::setupViewport(static_cast<QWidget*>(viewport));
}

void QHelpContentWidget_WheelEvent(void* ptr, void* e){
	static_cast<MyQHelpContentWidget*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(e));
}

void QHelpContentWidget_WheelEventDefault(void* ptr, void* e){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::wheelEvent(static_cast<QWheelEvent*>(e));
}

void QHelpContentWidget_ChangeEvent(void* ptr, void* ev){
	static_cast<MyQHelpContentWidget*>(ptr)->changeEvent(static_cast<QEvent*>(ev));
}

void QHelpContentWidget_ChangeEventDefault(void* ptr, void* ev){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::changeEvent(static_cast<QEvent*>(ev));
}

void QHelpContentWidget_ActionEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpContentWidget_ActionEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpContentWidget_EnterEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QHelpContentWidget_EnterEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::enterEvent(static_cast<QEvent*>(event));
}

void QHelpContentWidget_HideEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpContentWidget_HideEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpContentWidget_LeaveEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QHelpContentWidget_LeaveEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::leaveEvent(static_cast<QEvent*>(event));
}

void QHelpContentWidget_MoveEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpContentWidget_MoveEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpContentWidget_SetVisible(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<MyQHelpContentWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QHelpContentWidget_SetVisibleDefault(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QHelpContentWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QHelpContentWidget_ShowEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QHelpContentWidget_ShowEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::showEvent(static_cast<QShowEvent*>(event));
}

void QHelpContentWidget_CloseEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QHelpContentWidget_CloseEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QHelpContentWidget_InitPainter(void* ptr, void* painter){
	static_cast<MyQHelpContentWidget*>(ptr)->initPainter(static_cast<QPainter*>(painter));
}

void QHelpContentWidget_InitPainterDefault(void* ptr, void* painter){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::initPainter(static_cast<QPainter*>(painter));
}

void QHelpContentWidget_KeyReleaseEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpContentWidget_KeyReleaseEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpContentWidget_TabletEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpContentWidget_TabletEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpContentWidget_ChildEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHelpContentWidget_ChildEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpContentWidget_CustomEvent(void* ptr, void* event){
	static_cast<MyQHelpContentWidget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHelpContentWidget_CustomEventDefault(void* ptr, void* event){
	static_cast<QHelpContentWidget*>(ptr)->QHelpContentWidget::customEvent(static_cast<QEvent*>(event));
}

void* QHelpEngine_NewQHelpEngine(char* collectionFile, void* parent){
	return new QHelpEngine(QString(collectionFile), static_cast<QObject*>(parent));
}

void* QHelpEngine_ContentModel(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->contentModel();
}

void* QHelpEngine_ContentWidget(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->contentWidget();
}

void* QHelpEngine_IndexModel(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->indexModel();
}

void* QHelpEngine_IndexWidget(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->indexWidget();
}

void* QHelpEngine_SearchEngine(void* ptr){
	return static_cast<QHelpEngine*>(ptr)->searchEngine();
}

void QHelpEngine_DestroyQHelpEngine(void* ptr){
	static_cast<QHelpEngine*>(ptr)->~QHelpEngine();
}

void QHelpEngine_TimerEvent(void* ptr, void* event){
	static_cast<QHelpEngine*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpEngine_TimerEventDefault(void* ptr, void* event){
	static_cast<QHelpEngine*>(ptr)->QHelpEngine::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpEngine_ChildEvent(void* ptr, void* event){
	static_cast<QHelpEngine*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHelpEngine_ChildEventDefault(void* ptr, void* event){
	static_cast<QHelpEngine*>(ptr)->QHelpEngine::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpEngine_CustomEvent(void* ptr, void* event){
	static_cast<QHelpEngine*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHelpEngine_CustomEventDefault(void* ptr, void* event){
	static_cast<QHelpEngine*>(ptr)->QHelpEngine::customEvent(static_cast<QEvent*>(event));
}

class MyQHelpEngineCore: public QHelpEngineCore {
public:
	MyQHelpEngineCore(const QString &collectionFile, QObject *parent) : QHelpEngineCore(collectionFile, parent) {};
	void Signal_CurrentFilterChanged(const QString & newFilter) { callbackQHelpEngineCoreCurrentFilterChanged(this, this->objectName().toUtf8().data(), newFilter.toUtf8().data()); };
	void Signal_ReadersAboutToBeInvalidated() { callbackQHelpEngineCoreReadersAboutToBeInvalidated(this, this->objectName().toUtf8().data()); };
	void Signal_SetupFinished() { callbackQHelpEngineCoreSetupFinished(this, this->objectName().toUtf8().data()); };
	void Signal_SetupStarted() { callbackQHelpEngineCoreSetupStarted(this, this->objectName().toUtf8().data()); };
	void Signal_Warning(const QString & msg) { callbackQHelpEngineCoreWarning(this, this->objectName().toUtf8().data(), msg.toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQHelpEngineCoreTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQHelpEngineCoreChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQHelpEngineCoreCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QHelpEngineCore_AutoSaveFilter(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->autoSaveFilter();
}

char* QHelpEngineCore_CollectionFile(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->collectionFile().toUtf8().data();
}

char* QHelpEngineCore_CurrentFilter(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->currentFilter().toUtf8().data();
}

void QHelpEngineCore_SetAutoSaveFilter(void* ptr, int save){
	static_cast<QHelpEngineCore*>(ptr)->setAutoSaveFilter(save != 0);
}

void QHelpEngineCore_SetCollectionFile(void* ptr, char* fileName){
	static_cast<QHelpEngineCore*>(ptr)->setCollectionFile(QString(fileName));
}

void QHelpEngineCore_SetCurrentFilter(void* ptr, char* filterName){
	static_cast<QHelpEngineCore*>(ptr)->setCurrentFilter(QString(filterName));
}

void* QHelpEngineCore_NewQHelpEngineCore(char* collectionFile, void* parent){
	return new MyQHelpEngineCore(QString(collectionFile), static_cast<QObject*>(parent));
}

int QHelpEngineCore_AddCustomFilter(void* ptr, char* filterName, char* attributes){
	return static_cast<QHelpEngineCore*>(ptr)->addCustomFilter(QString(filterName), QString(attributes).split("|", QString::SkipEmptyParts));
}

int QHelpEngineCore_CopyCollectionFile(void* ptr, char* fileName){
	return static_cast<QHelpEngineCore*>(ptr)->copyCollectionFile(QString(fileName));
}

void QHelpEngineCore_ConnectCurrentFilterChanged(void* ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::currentFilterChanged), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_CurrentFilterChanged));;
}

void QHelpEngineCore_DisconnectCurrentFilterChanged(void* ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::currentFilterChanged), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_CurrentFilterChanged));;
}

void QHelpEngineCore_CurrentFilterChanged(void* ptr, char* newFilter){
	static_cast<QHelpEngineCore*>(ptr)->currentFilterChanged(QString(newFilter));
}

char* QHelpEngineCore_CustomFilters(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->customFilters().join("|").toUtf8().data();
}

void* QHelpEngineCore_CustomValue(void* ptr, char* key, void* defaultValue){
	return new QVariant(static_cast<QHelpEngineCore*>(ptr)->customValue(QString(key), *static_cast<QVariant*>(defaultValue)));
}

char* QHelpEngineCore_DocumentationFileName(void* ptr, char* namespaceName){
	return static_cast<QHelpEngineCore*>(ptr)->documentationFileName(QString(namespaceName)).toUtf8().data();
}

char* QHelpEngineCore_Error(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->error().toUtf8().data();
}

char* QHelpEngineCore_FileData(void* ptr, void* url){
	return QString(static_cast<QHelpEngineCore*>(ptr)->fileData(*static_cast<QUrl*>(url))).toUtf8().data();
}

char* QHelpEngineCore_FilterAttributes(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->filterAttributes().join("|").toUtf8().data();
}

char* QHelpEngineCore_FilterAttributes2(void* ptr, char* filterName){
	return static_cast<QHelpEngineCore*>(ptr)->filterAttributes(QString(filterName)).join("|").toUtf8().data();
}

void* QHelpEngineCore_FindFile(void* ptr, void* url){
	return new QUrl(static_cast<QHelpEngineCore*>(ptr)->findFile(*static_cast<QUrl*>(url)));
}

void* QHelpEngineCore_QHelpEngineCore_MetaData(char* documentationFileName, char* name){
	return new QVariant(QHelpEngineCore::metaData(QString(documentationFileName), QString(name)));
}

char* QHelpEngineCore_QHelpEngineCore_NamespaceName(char* documentationFileName){
	return QHelpEngineCore::namespaceName(QString(documentationFileName)).toUtf8().data();
}

void QHelpEngineCore_ConnectReadersAboutToBeInvalidated(void* ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::readersAboutToBeInvalidated), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_ReadersAboutToBeInvalidated));;
}

void QHelpEngineCore_DisconnectReadersAboutToBeInvalidated(void* ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::readersAboutToBeInvalidated), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_ReadersAboutToBeInvalidated));;
}

void QHelpEngineCore_ReadersAboutToBeInvalidated(void* ptr){
	static_cast<QHelpEngineCore*>(ptr)->readersAboutToBeInvalidated();
}

int QHelpEngineCore_RegisterDocumentation(void* ptr, char* documentationFileName){
	return static_cast<QHelpEngineCore*>(ptr)->registerDocumentation(QString(documentationFileName));
}

char* QHelpEngineCore_RegisteredDocumentations(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->registeredDocumentations().join("|").toUtf8().data();
}

int QHelpEngineCore_RemoveCustomFilter(void* ptr, char* filterName){
	return static_cast<QHelpEngineCore*>(ptr)->removeCustomFilter(QString(filterName));
}

int QHelpEngineCore_RemoveCustomValue(void* ptr, char* key){
	return static_cast<QHelpEngineCore*>(ptr)->removeCustomValue(QString(key));
}

int QHelpEngineCore_SetCustomValue(void* ptr, char* key, void* value){
	return static_cast<QHelpEngineCore*>(ptr)->setCustomValue(QString(key), *static_cast<QVariant*>(value));
}

int QHelpEngineCore_SetupData(void* ptr){
	return static_cast<QHelpEngineCore*>(ptr)->setupData();
}

void QHelpEngineCore_ConnectSetupFinished(void* ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupFinished), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupFinished));;
}

void QHelpEngineCore_DisconnectSetupFinished(void* ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupFinished), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupFinished));;
}

void QHelpEngineCore_SetupFinished(void* ptr){
	static_cast<QHelpEngineCore*>(ptr)->setupFinished();
}

void QHelpEngineCore_ConnectSetupStarted(void* ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupStarted), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupStarted));;
}

void QHelpEngineCore_DisconnectSetupStarted(void* ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)()>(&QHelpEngineCore::setupStarted), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)()>(&MyQHelpEngineCore::Signal_SetupStarted));;
}

void QHelpEngineCore_SetupStarted(void* ptr){
	static_cast<QHelpEngineCore*>(ptr)->setupStarted();
}

int QHelpEngineCore_UnregisterDocumentation(void* ptr, char* namespaceName){
	return static_cast<QHelpEngineCore*>(ptr)->unregisterDocumentation(QString(namespaceName));
}

void QHelpEngineCore_ConnectWarning(void* ptr){
	QObject::connect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::warning), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_Warning));;
}

void QHelpEngineCore_DisconnectWarning(void* ptr){
	QObject::disconnect(static_cast<QHelpEngineCore*>(ptr), static_cast<void (QHelpEngineCore::*)(const QString &)>(&QHelpEngineCore::warning), static_cast<MyQHelpEngineCore*>(ptr), static_cast<void (MyQHelpEngineCore::*)(const QString &)>(&MyQHelpEngineCore::Signal_Warning));;
}

void QHelpEngineCore_Warning(void* ptr, char* msg){
	static_cast<QHelpEngineCore*>(ptr)->warning(QString(msg));
}

void QHelpEngineCore_DestroyQHelpEngineCore(void* ptr){
	static_cast<QHelpEngineCore*>(ptr)->~QHelpEngineCore();
}

void QHelpEngineCore_TimerEvent(void* ptr, void* event){
	static_cast<MyQHelpEngineCore*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpEngineCore_TimerEventDefault(void* ptr, void* event){
	static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpEngineCore_ChildEvent(void* ptr, void* event){
	static_cast<MyQHelpEngineCore*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHelpEngineCore_ChildEventDefault(void* ptr, void* event){
	static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpEngineCore_CustomEvent(void* ptr, void* event){
	static_cast<MyQHelpEngineCore*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHelpEngineCore_CustomEventDefault(void* ptr, void* event){
	static_cast<QHelpEngineCore*>(ptr)->QHelpEngineCore::customEvent(static_cast<QEvent*>(event));
}

class MyQHelpIndexModel: public QHelpIndexModel {
public:
	void Signal_IndexCreated() { callbackQHelpIndexModelIndexCreated(this, this->objectName().toUtf8().data()); };
	void Signal_IndexCreationStarted() { callbackQHelpIndexModelIndexCreationStarted(this, this->objectName().toUtf8().data()); };
	void sort(int column, Qt::SortOrder order) { callbackQHelpIndexModelSort(this, this->objectName().toUtf8().data(), column, order); };
	void revert() { if (!callbackQHelpIndexModelRevert(this, this->objectName().toUtf8().data())) { QHelpIndexModel::revert(); }; };
	void timerEvent(QTimerEvent * event) { callbackQHelpIndexModelTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQHelpIndexModelChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQHelpIndexModelCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QHelpIndexModel_CreateIndex(void* ptr, char* customFilterName){
	static_cast<QHelpIndexModel*>(ptr)->createIndex(QString(customFilterName));
}

void QHelpIndexModel_ConnectIndexCreated(void* ptr){
	QObject::connect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreated), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreated));;
}

void QHelpIndexModel_DisconnectIndexCreated(void* ptr){
	QObject::disconnect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreated), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreated));;
}

void QHelpIndexModel_IndexCreated(void* ptr){
	static_cast<QHelpIndexModel*>(ptr)->indexCreated();
}

void QHelpIndexModel_ConnectIndexCreationStarted(void* ptr){
	QObject::connect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreationStarted), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreationStarted));;
}

void QHelpIndexModel_DisconnectIndexCreationStarted(void* ptr){
	QObject::disconnect(static_cast<QHelpIndexModel*>(ptr), static_cast<void (QHelpIndexModel::*)()>(&QHelpIndexModel::indexCreationStarted), static_cast<MyQHelpIndexModel*>(ptr), static_cast<void (MyQHelpIndexModel::*)()>(&MyQHelpIndexModel::Signal_IndexCreationStarted));;
}

void QHelpIndexModel_IndexCreationStarted(void* ptr){
	static_cast<QHelpIndexModel*>(ptr)->indexCreationStarted();
}

int QHelpIndexModel_IsCreatingIndex(void* ptr){
	return static_cast<QHelpIndexModel*>(ptr)->isCreatingIndex();
}

void QHelpIndexModel_Sort(void* ptr, int column, int order){
	static_cast<MyQHelpIndexModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void QHelpIndexModel_SortDefault(void* ptr, int column, int order){
	static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void QHelpIndexModel_Revert(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQHelpIndexModel*>(ptr), "revert");
}

void QHelpIndexModel_RevertDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpIndexModel*>(ptr), "revert");
}

void QHelpIndexModel_TimerEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpIndexModel_TimerEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpIndexModel_ChildEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHelpIndexModel_ChildEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpIndexModel_CustomEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHelpIndexModel_CustomEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexModel*>(ptr)->QHelpIndexModel::customEvent(static_cast<QEvent*>(event));
}

class MyQHelpIndexWidget: public QHelpIndexWidget {
public:
	void Signal_LinkActivated(const QUrl & link, const QString & keyword) { callbackQHelpIndexWidgetLinkActivated(this, this->objectName().toUtf8().data(), new QUrl(link), keyword.toUtf8().data()); };
	void dragLeaveEvent(QDragLeaveEvent * e) { callbackQHelpIndexWidgetDragLeaveEvent(this, this->objectName().toUtf8().data(), e); };
	void dragMoveEvent(QDragMoveEvent * e) { callbackQHelpIndexWidgetDragMoveEvent(this, this->objectName().toUtf8().data(), e); };
	void dropEvent(QDropEvent * e) { callbackQHelpIndexWidgetDropEvent(this, this->objectName().toUtf8().data(), e); };
	void mouseMoveEvent(QMouseEvent * e) { callbackQHelpIndexWidgetMouseMoveEvent(this, this->objectName().toUtf8().data(), e); };
	void mouseReleaseEvent(QMouseEvent * e) { callbackQHelpIndexWidgetMouseReleaseEvent(this, this->objectName().toUtf8().data(), e); };
	void paintEvent(QPaintEvent * e) { callbackQHelpIndexWidgetPaintEvent(this, this->objectName().toUtf8().data(), e); };
	void resizeEvent(QResizeEvent * e) { callbackQHelpIndexWidgetResizeEvent(this, this->objectName().toUtf8().data(), e); };
	void setSelection(const QRect & rect, QItemSelectionModel::SelectionFlags command) { callbackQHelpIndexWidgetSetSelection(this, this->objectName().toUtf8().data(), new QRect(static_cast<QRect>(rect).x(), static_cast<QRect>(rect).y(), static_cast<QRect>(rect).width(), static_cast<QRect>(rect).height()), command); };
	void startDrag(Qt::DropActions supportedActions) { callbackQHelpIndexWidgetStartDrag(this, this->objectName().toUtf8().data(), supportedActions); };
	void timerEvent(QTimerEvent * e) { callbackQHelpIndexWidgetTimerEvent(this, this->objectName().toUtf8().data(), e); };
	void updateGeometries() { if (!callbackQHelpIndexWidgetUpdateGeometries(this, this->objectName().toUtf8().data())) { QHelpIndexWidget::updateGeometries(); }; };
	void closeEditor(QWidget * editor, QAbstractItemDelegate::EndEditHint hint) { if (!callbackQHelpIndexWidgetCloseEditor(this, this->objectName().toUtf8().data(), editor, hint)) { QHelpIndexWidget::closeEditor(editor, hint); }; };
	void commitData(QWidget * editor) { if (!callbackQHelpIndexWidgetCommitData(this, this->objectName().toUtf8().data(), editor)) { QHelpIndexWidget::commitData(editor); }; };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQHelpIndexWidgetDragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void editorDestroyed(QObject * editor) { if (!callbackQHelpIndexWidgetEditorDestroyed(this, this->objectName().toUtf8().data(), editor)) { QHelpIndexWidget::editorDestroyed(editor); }; };
	void focusInEvent(QFocusEvent * event) { callbackQHelpIndexWidgetFocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQHelpIndexWidgetFocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQHelpIndexWidgetInputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	void keyPressEvent(QKeyEvent * event) { callbackQHelpIndexWidgetKeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyboardSearch(const QString & search) { callbackQHelpIndexWidgetKeyboardSearch(this, this->objectName().toUtf8().data(), search.toUtf8().data()); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQHelpIndexWidgetMouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQHelpIndexWidgetMousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void reset() { if (!callbackQHelpIndexWidgetReset(this, this->objectName().toUtf8().data())) { QHelpIndexWidget::reset(); }; };
	void selectAll() { if (!callbackQHelpIndexWidgetSelectAll(this, this->objectName().toUtf8().data())) { QHelpIndexWidget::selectAll(); }; };
	void setModel(QAbstractItemModel * model) { callbackQHelpIndexWidgetSetModel(this, this->objectName().toUtf8().data(), model); };
	void setSelectionModel(QItemSelectionModel * selectionModel) { callbackQHelpIndexWidgetSetSelectionModel(this, this->objectName().toUtf8().data(), selectionModel); };
	void contextMenuEvent(QContextMenuEvent * e) { callbackQHelpIndexWidgetContextMenuEvent(this, this->objectName().toUtf8().data(), e); };
	void scrollContentsBy(int dx, int dy) { callbackQHelpIndexWidgetScrollContentsBy(this, this->objectName().toUtf8().data(), dx, dy); };
	void setupViewport(QWidget * viewport) { callbackQHelpIndexWidgetSetupViewport(this, this->objectName().toUtf8().data(), viewport); };
	void wheelEvent(QWheelEvent * e) { callbackQHelpIndexWidgetWheelEvent(this, this->objectName().toUtf8().data(), e); };
	void changeEvent(QEvent * ev) { callbackQHelpIndexWidgetChangeEvent(this, this->objectName().toUtf8().data(), ev); };
	void actionEvent(QActionEvent * event) { callbackQHelpIndexWidgetActionEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQHelpIndexWidgetEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void hideEvent(QHideEvent * event) { callbackQHelpIndexWidgetHideEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQHelpIndexWidgetLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void moveEvent(QMoveEvent * event) { callbackQHelpIndexWidgetMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void setVisible(bool visible) { if (!callbackQHelpIndexWidgetSetVisible(this, this->objectName().toUtf8().data(), visible)) { QHelpIndexWidget::setVisible(visible); }; };
	void showEvent(QShowEvent * event) { callbackQHelpIndexWidgetShowEvent(this, this->objectName().toUtf8().data(), event); };
	void closeEvent(QCloseEvent * event) { callbackQHelpIndexWidgetCloseEvent(this, this->objectName().toUtf8().data(), event); };
	void initPainter(QPainter * painter) const { callbackQHelpIndexWidgetInitPainter(const_cast<MyQHelpIndexWidget*>(this), this->objectName().toUtf8().data(), painter); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQHelpIndexWidgetKeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void tabletEvent(QTabletEvent * event) { callbackQHelpIndexWidgetTabletEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQHelpIndexWidgetChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQHelpIndexWidgetCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void QHelpIndexWidget_ActivateCurrentItem(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "activateCurrentItem");
}

void QHelpIndexWidget_FilterIndices(void* ptr, char* filter, char* wildcard){
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "filterIndices", Q_ARG(QString, QString(filter)), Q_ARG(QString, QString(wildcard)));
}

void QHelpIndexWidget_ConnectLinkActivated(void* ptr){
	QObject::connect(static_cast<QHelpIndexWidget*>(ptr), static_cast<void (QHelpIndexWidget::*)(const QUrl &, const QString &)>(&QHelpIndexWidget::linkActivated), static_cast<MyQHelpIndexWidget*>(ptr), static_cast<void (MyQHelpIndexWidget::*)(const QUrl &, const QString &)>(&MyQHelpIndexWidget::Signal_LinkActivated));;
}

void QHelpIndexWidget_DisconnectLinkActivated(void* ptr){
	QObject::disconnect(static_cast<QHelpIndexWidget*>(ptr), static_cast<void (QHelpIndexWidget::*)(const QUrl &, const QString &)>(&QHelpIndexWidget::linkActivated), static_cast<MyQHelpIndexWidget*>(ptr), static_cast<void (MyQHelpIndexWidget::*)(const QUrl &, const QString &)>(&MyQHelpIndexWidget::Signal_LinkActivated));;
}

void QHelpIndexWidget_LinkActivated(void* ptr, void* link, char* keyword){
	static_cast<QHelpIndexWidget*>(ptr)->linkActivated(*static_cast<QUrl*>(link), QString(keyword));
}

void QHelpIndexWidget_DragLeaveEvent(void* ptr, void* e){
	static_cast<MyQHelpIndexWidget*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(e));
}

void QHelpIndexWidget_DragLeaveEventDefault(void* ptr, void* e){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(e));
}

void QHelpIndexWidget_DragMoveEvent(void* ptr, void* e){
	static_cast<MyQHelpIndexWidget*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(e));
}

void QHelpIndexWidget_DragMoveEventDefault(void* ptr, void* e){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(e));
}

void QHelpIndexWidget_DropEvent(void* ptr, void* e){
	static_cast<MyQHelpIndexWidget*>(ptr)->dropEvent(static_cast<QDropEvent*>(e));
}

void QHelpIndexWidget_DropEventDefault(void* ptr, void* e){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::dropEvent(static_cast<QDropEvent*>(e));
}

void QHelpIndexWidget_MouseMoveEvent(void* ptr, void* e){
	static_cast<MyQHelpIndexWidget*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(e));
}

void QHelpIndexWidget_MouseMoveEventDefault(void* ptr, void* e){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::mouseMoveEvent(static_cast<QMouseEvent*>(e));
}

void QHelpIndexWidget_MouseReleaseEvent(void* ptr, void* e){
	static_cast<MyQHelpIndexWidget*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(e));
}

void QHelpIndexWidget_MouseReleaseEventDefault(void* ptr, void* e){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(e));
}

void QHelpIndexWidget_PaintEvent(void* ptr, void* e){
	static_cast<MyQHelpIndexWidget*>(ptr)->paintEvent(static_cast<QPaintEvent*>(e));
}

void QHelpIndexWidget_PaintEventDefault(void* ptr, void* e){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::paintEvent(static_cast<QPaintEvent*>(e));
}

void QHelpIndexWidget_ResizeEvent(void* ptr, void* e){
	static_cast<MyQHelpIndexWidget*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(e));
}

void QHelpIndexWidget_ResizeEventDefault(void* ptr, void* e){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::resizeEvent(static_cast<QResizeEvent*>(e));
}

void QHelpIndexWidget_SetSelection(void* ptr, void* rect, int command){
	static_cast<MyQHelpIndexWidget*>(ptr)->setSelection(*static_cast<QRect*>(rect), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QHelpIndexWidget_SetSelectionDefault(void* ptr, void* rect, int command){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setSelection(*static_cast<QRect*>(rect), static_cast<QItemSelectionModel::SelectionFlag>(command));
}

void QHelpIndexWidget_StartDrag(void* ptr, int supportedActions){
	static_cast<MyQHelpIndexWidget*>(ptr)->startDrag(static_cast<Qt::DropAction>(supportedActions));
}

void QHelpIndexWidget_StartDragDefault(void* ptr, int supportedActions){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::startDrag(static_cast<Qt::DropAction>(supportedActions));
}

void QHelpIndexWidget_TimerEvent(void* ptr, void* e){
	static_cast<MyQHelpIndexWidget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(e));
}

void QHelpIndexWidget_TimerEventDefault(void* ptr, void* e){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::timerEvent(static_cast<QTimerEvent*>(e));
}

void QHelpIndexWidget_UpdateGeometries(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQHelpIndexWidget*>(ptr), "updateGeometries");
}

void QHelpIndexWidget_UpdateGeometriesDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "updateGeometries");
}

void QHelpIndexWidget_CloseEditor(void* ptr, void* editor, int hint){
	QMetaObject::invokeMethod(static_cast<MyQHelpIndexWidget*>(ptr), "closeEditor", Q_ARG(QWidget*, static_cast<QWidget*>(editor)), Q_ARG(QAbstractItemDelegate::EndEditHint, static_cast<QAbstractItemDelegate::EndEditHint>(hint)));
}

void QHelpIndexWidget_CloseEditorDefault(void* ptr, void* editor, int hint){
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "closeEditor", Q_ARG(QWidget*, static_cast<QWidget*>(editor)), Q_ARG(QAbstractItemDelegate::EndEditHint, static_cast<QAbstractItemDelegate::EndEditHint>(hint)));
}

void QHelpIndexWidget_CommitData(void* ptr, void* editor){
	QMetaObject::invokeMethod(static_cast<MyQHelpIndexWidget*>(ptr), "commitData", Q_ARG(QWidget*, static_cast<QWidget*>(editor)));
}

void QHelpIndexWidget_CommitDataDefault(void* ptr, void* editor){
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "commitData", Q_ARG(QWidget*, static_cast<QWidget*>(editor)));
}

void QHelpIndexWidget_DragEnterEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexWidget*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpIndexWidget_DragEnterEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpIndexWidget_EditorDestroyed(void* ptr, void* editor){
	QMetaObject::invokeMethod(static_cast<MyQHelpIndexWidget*>(ptr), "editorDestroyed", Q_ARG(QObject*, static_cast<QObject*>(editor)));
}

void QHelpIndexWidget_EditorDestroyedDefault(void* ptr, void* editor){
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "editorDestroyed", Q_ARG(QObject*, static_cast<QObject*>(editor)));
}

void QHelpIndexWidget_FocusInEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexWidget*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QHelpIndexWidget_FocusInEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QHelpIndexWidget_FocusOutEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexWidget*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpIndexWidget_FocusOutEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpIndexWidget_InputMethodEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexWidget*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QHelpIndexWidget_InputMethodEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QHelpIndexWidget_KeyPressEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexWidget*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpIndexWidget_KeyPressEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpIndexWidget_KeyboardSearch(void* ptr, char* search){
	static_cast<MyQHelpIndexWidget*>(ptr)->keyboardSearch(QString(search));
}

void QHelpIndexWidget_KeyboardSearchDefault(void* ptr, char* search){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::keyboardSearch(QString(search));
}

void QHelpIndexWidget_MouseDoubleClickEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexWidget*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpIndexWidget_MouseDoubleClickEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpIndexWidget_MousePressEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexWidget*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpIndexWidget_MousePressEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpIndexWidget_Reset(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQHelpIndexWidget*>(ptr), "reset");
}

void QHelpIndexWidget_ResetDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "reset");
}

void QHelpIndexWidget_SelectAll(void* ptr){
	QMetaObject::invokeMethod(static_cast<MyQHelpIndexWidget*>(ptr), "selectAll");
}

void QHelpIndexWidget_SelectAllDefault(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "selectAll");
}

void QHelpIndexWidget_SetModel(void* ptr, void* model){
	static_cast<MyQHelpIndexWidget*>(ptr)->setModel(static_cast<QAbstractItemModel*>(model));
}

void QHelpIndexWidget_SetModelDefault(void* ptr, void* model){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setModel(static_cast<QAbstractItemModel*>(model));
}

void QHelpIndexWidget_SetSelectionModel(void* ptr, void* selectionModel){
	static_cast<MyQHelpIndexWidget*>(ptr)->setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

void QHelpIndexWidget_SetSelectionModelDefault(void* ptr, void* selectionModel){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setSelectionModel(static_cast<QItemSelectionModel*>(selectionModel));
}

void QHelpIndexWidget_ContextMenuEvent(void* ptr, void* e){
	static_cast<MyQHelpIndexWidget*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void QHelpIndexWidget_ContextMenuEventDefault(void* ptr, void* e){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(e));
}

void QHelpIndexWidget_ScrollContentsBy(void* ptr, int dx, int dy){
	static_cast<MyQHelpIndexWidget*>(ptr)->scrollContentsBy(dx, dy);
}

void QHelpIndexWidget_ScrollContentsByDefault(void* ptr, int dx, int dy){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::scrollContentsBy(dx, dy);
}

void QHelpIndexWidget_SetupViewport(void* ptr, void* viewport){
	static_cast<MyQHelpIndexWidget*>(ptr)->setupViewport(static_cast<QWidget*>(viewport));
}

void QHelpIndexWidget_SetupViewportDefault(void* ptr, void* viewport){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::setupViewport(static_cast<QWidget*>(viewport));
}

void QHelpIndexWidget_WheelEvent(void* ptr, void* e){
	static_cast<MyQHelpIndexWidget*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(e));
}

void QHelpIndexWidget_WheelEventDefault(void* ptr, void* e){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::wheelEvent(static_cast<QWheelEvent*>(e));
}

void QHelpIndexWidget_ChangeEvent(void* ptr, void* ev){
	static_cast<MyQHelpIndexWidget*>(ptr)->changeEvent(static_cast<QEvent*>(ev));
}

void QHelpIndexWidget_ChangeEventDefault(void* ptr, void* ev){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::changeEvent(static_cast<QEvent*>(ev));
}

void QHelpIndexWidget_ActionEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexWidget*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpIndexWidget_ActionEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpIndexWidget_EnterEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexWidget*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QHelpIndexWidget_EnterEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::enterEvent(static_cast<QEvent*>(event));
}

void QHelpIndexWidget_HideEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexWidget*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpIndexWidget_HideEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpIndexWidget_LeaveEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexWidget*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QHelpIndexWidget_LeaveEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::leaveEvent(static_cast<QEvent*>(event));
}

void QHelpIndexWidget_MoveEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexWidget*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpIndexWidget_MoveEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpIndexWidget_SetVisible(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<MyQHelpIndexWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QHelpIndexWidget_SetVisibleDefault(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QHelpIndexWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QHelpIndexWidget_ShowEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexWidget*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QHelpIndexWidget_ShowEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::showEvent(static_cast<QShowEvent*>(event));
}

void QHelpIndexWidget_CloseEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexWidget*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QHelpIndexWidget_CloseEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QHelpIndexWidget_InitPainter(void* ptr, void* painter){
	static_cast<MyQHelpIndexWidget*>(ptr)->initPainter(static_cast<QPainter*>(painter));
}

void QHelpIndexWidget_InitPainterDefault(void* ptr, void* painter){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::initPainter(static_cast<QPainter*>(painter));
}

void QHelpIndexWidget_KeyReleaseEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexWidget*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpIndexWidget_KeyReleaseEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpIndexWidget_TabletEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexWidget*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpIndexWidget_TabletEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpIndexWidget_ChildEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexWidget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHelpIndexWidget_ChildEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpIndexWidget_CustomEvent(void* ptr, void* event){
	static_cast<MyQHelpIndexWidget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHelpIndexWidget_CustomEventDefault(void* ptr, void* event){
	static_cast<QHelpIndexWidget*>(ptr)->QHelpIndexWidget::customEvent(static_cast<QEvent*>(event));
}

class MyQHelpSearchEngine: public QHelpSearchEngine {
public:
	void Signal_IndexingFinished() { callbackQHelpSearchEngineIndexingFinished(this, this->objectName().toUtf8().data()); };
	void Signal_IndexingStarted() { callbackQHelpSearchEngineIndexingStarted(this, this->objectName().toUtf8().data()); };
	void Signal_SearchingFinished(int hits) { callbackQHelpSearchEngineSearchingFinished(this, this->objectName().toUtf8().data(), hits); };
	void Signal_SearchingStarted() { callbackQHelpSearchEngineSearchingStarted(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQHelpSearchEngineTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQHelpSearchEngineChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQHelpSearchEngineCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QHelpSearchEngine_NewQHelpSearchEngine(void* helpEngine, void* parent){
	return new QHelpSearchEngine(static_cast<QHelpEngineCore*>(helpEngine), static_cast<QObject*>(parent));
}

void QHelpSearchEngine_CancelIndexing(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "cancelIndexing");
}

void QHelpSearchEngine_CancelSearching(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "cancelSearching");
}

int QHelpSearchEngine_HitCount(void* ptr){
	return static_cast<QHelpSearchEngine*>(ptr)->hitCount();
}

void QHelpSearchEngine_ConnectIndexingFinished(void* ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingFinished));;
}

void QHelpSearchEngine_DisconnectIndexingFinished(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingFinished));;
}

void QHelpSearchEngine_IndexingFinished(void* ptr){
	static_cast<QHelpSearchEngine*>(ptr)->indexingFinished();
}

void QHelpSearchEngine_ConnectIndexingStarted(void* ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingStarted));;
}

void QHelpSearchEngine_DisconnectIndexingStarted(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::indexingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_IndexingStarted));;
}

void QHelpSearchEngine_IndexingStarted(void* ptr){
	static_cast<QHelpSearchEngine*>(ptr)->indexingStarted();
}

void* QHelpSearchEngine_QueryWidget(void* ptr){
	return static_cast<QHelpSearchEngine*>(ptr)->queryWidget();
}

void QHelpSearchEngine_ReindexDocumentation(void* ptr){
	QMetaObject::invokeMethod(static_cast<QHelpSearchEngine*>(ptr), "reindexDocumentation");
}

void* QHelpSearchEngine_ResultWidget(void* ptr){
	return static_cast<QHelpSearchEngine*>(ptr)->resultWidget();
}

void QHelpSearchEngine_ConnectSearchingFinished(void* ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)(int)>(&QHelpSearchEngine::searchingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)(int)>(&MyQHelpSearchEngine::Signal_SearchingFinished));;
}

void QHelpSearchEngine_DisconnectSearchingFinished(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)(int)>(&QHelpSearchEngine::searchingFinished), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)(int)>(&MyQHelpSearchEngine::Signal_SearchingFinished));;
}

void QHelpSearchEngine_SearchingFinished(void* ptr, int hits){
	static_cast<QHelpSearchEngine*>(ptr)->searchingFinished(hits);
}

void QHelpSearchEngine_ConnectSearchingStarted(void* ptr){
	QObject::connect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::searchingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_SearchingStarted));;
}

void QHelpSearchEngine_DisconnectSearchingStarted(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchEngine*>(ptr), static_cast<void (QHelpSearchEngine::*)()>(&QHelpSearchEngine::searchingStarted), static_cast<MyQHelpSearchEngine*>(ptr), static_cast<void (MyQHelpSearchEngine::*)()>(&MyQHelpSearchEngine::Signal_SearchingStarted));;
}

void QHelpSearchEngine_SearchingStarted(void* ptr){
	static_cast<QHelpSearchEngine*>(ptr)->searchingStarted();
}

void QHelpSearchEngine_DestroyQHelpSearchEngine(void* ptr){
	static_cast<QHelpSearchEngine*>(ptr)->~QHelpSearchEngine();
}

void QHelpSearchEngine_TimerEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchEngine*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpSearchEngine_TimerEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpSearchEngine_ChildEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchEngine*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHelpSearchEngine_ChildEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpSearchEngine_CustomEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchEngine*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHelpSearchEngine_CustomEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchEngine*>(ptr)->QHelpSearchEngine::customEvent(static_cast<QEvent*>(event));
}

void* QHelpSearchQuery_NewQHelpSearchQuery(){
	return new QHelpSearchQuery();
}

void* QHelpSearchQuery_NewQHelpSearchQuery2(int field, char* wordList){
	return new QHelpSearchQuery(static_cast<QHelpSearchQuery::FieldName>(field), QString(wordList).split("|", QString::SkipEmptyParts));
}

class MyQHelpSearchQueryWidget: public QHelpSearchQueryWidget {
public:
	void Signal_Search() { callbackQHelpSearchQueryWidgetSearch(this, this->objectName().toUtf8().data()); };
	void actionEvent(QActionEvent * event) { callbackQHelpSearchQueryWidgetActionEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQHelpSearchQueryWidgetDragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQHelpSearchQueryWidgetDragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQHelpSearchQueryWidgetDragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQHelpSearchQueryWidgetDropEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQHelpSearchQueryWidgetEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQHelpSearchQueryWidgetFocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void hideEvent(QHideEvent * event) { callbackQHelpSearchQueryWidgetHideEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQHelpSearchQueryWidgetLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void moveEvent(QMoveEvent * event) { callbackQHelpSearchQueryWidgetMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void paintEvent(QPaintEvent * event) { callbackQHelpSearchQueryWidgetPaintEvent(this, this->objectName().toUtf8().data(), event); };
	void setVisible(bool visible) { if (!callbackQHelpSearchQueryWidgetSetVisible(this, this->objectName().toUtf8().data(), visible)) { QHelpSearchQueryWidget::setVisible(visible); }; };
	void showEvent(QShowEvent * event) { callbackQHelpSearchQueryWidgetShowEvent(this, this->objectName().toUtf8().data(), event); };
	void closeEvent(QCloseEvent * event) { callbackQHelpSearchQueryWidgetCloseEvent(this, this->objectName().toUtf8().data(), event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQHelpSearchQueryWidgetContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	void initPainter(QPainter * painter) const { callbackQHelpSearchQueryWidgetInitPainter(const_cast<MyQHelpSearchQueryWidget*>(this), this->objectName().toUtf8().data(), painter); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQHelpSearchQueryWidgetInputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	void keyPressEvent(QKeyEvent * event) { callbackQHelpSearchQueryWidgetKeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQHelpSearchQueryWidgetKeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQHelpSearchQueryWidgetMouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQHelpSearchQueryWidgetMouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQHelpSearchQueryWidgetMousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQHelpSearchQueryWidgetMouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void resizeEvent(QResizeEvent * event) { callbackQHelpSearchQueryWidgetResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void tabletEvent(QTabletEvent * event) { callbackQHelpSearchQueryWidgetTabletEvent(this, this->objectName().toUtf8().data(), event); };
	void wheelEvent(QWheelEvent * event) { callbackQHelpSearchQueryWidgetWheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQHelpSearchQueryWidgetTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQHelpSearchQueryWidgetChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQHelpSearchQueryWidgetCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

int QHelpSearchQueryWidget_IsCompactMode(void* ptr){
	return static_cast<QHelpSearchQueryWidget*>(ptr)->isCompactMode();
}

void* QHelpSearchQueryWidget_NewQHelpSearchQueryWidget(void* parent){
	return new QHelpSearchQueryWidget(static_cast<QWidget*>(parent));
}

void QHelpSearchQueryWidget_CollapseExtendedSearch(void* ptr){
	static_cast<QHelpSearchQueryWidget*>(ptr)->collapseExtendedSearch();
}

void QHelpSearchQueryWidget_ExpandExtendedSearch(void* ptr){
	static_cast<QHelpSearchQueryWidget*>(ptr)->expandExtendedSearch();
}

void QHelpSearchQueryWidget_ConnectSearch(void* ptr){
	QObject::connect(static_cast<QHelpSearchQueryWidget*>(ptr), static_cast<void (QHelpSearchQueryWidget::*)()>(&QHelpSearchQueryWidget::search), static_cast<MyQHelpSearchQueryWidget*>(ptr), static_cast<void (MyQHelpSearchQueryWidget::*)()>(&MyQHelpSearchQueryWidget::Signal_Search));;
}

void QHelpSearchQueryWidget_DisconnectSearch(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchQueryWidget*>(ptr), static_cast<void (QHelpSearchQueryWidget::*)()>(&QHelpSearchQueryWidget::search), static_cast<MyQHelpSearchQueryWidget*>(ptr), static_cast<void (MyQHelpSearchQueryWidget::*)()>(&MyQHelpSearchQueryWidget::Signal_Search));;
}

void QHelpSearchQueryWidget_Search(void* ptr){
	static_cast<QHelpSearchQueryWidget*>(ptr)->search();
}

void QHelpSearchQueryWidget_DestroyQHelpSearchQueryWidget(void* ptr){
	static_cast<QHelpSearchQueryWidget*>(ptr)->~QHelpSearchQueryWidget();
}

void QHelpSearchQueryWidget_ActionEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpSearchQueryWidget_ActionEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpSearchQueryWidget_DragEnterEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpSearchQueryWidget_DragEnterEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpSearchQueryWidget_DragLeaveEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QHelpSearchQueryWidget_DragLeaveEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QHelpSearchQueryWidget_DragMoveEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QHelpSearchQueryWidget_DragMoveEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QHelpSearchQueryWidget_DropEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QHelpSearchQueryWidget_DropEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void QHelpSearchQueryWidget_EnterEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QHelpSearchQueryWidget_EnterEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::enterEvent(static_cast<QEvent*>(event));
}

void QHelpSearchQueryWidget_FocusOutEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpSearchQueryWidget_FocusOutEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpSearchQueryWidget_HideEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpSearchQueryWidget_HideEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpSearchQueryWidget_LeaveEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QHelpSearchQueryWidget_LeaveEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::leaveEvent(static_cast<QEvent*>(event));
}

void QHelpSearchQueryWidget_MoveEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpSearchQueryWidget_MoveEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpSearchQueryWidget_PaintEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QHelpSearchQueryWidget_PaintEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void QHelpSearchQueryWidget_SetVisible(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<MyQHelpSearchQueryWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QHelpSearchQueryWidget_SetVisibleDefault(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QHelpSearchQueryWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QHelpSearchQueryWidget_ShowEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QHelpSearchQueryWidget_ShowEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::showEvent(static_cast<QShowEvent*>(event));
}

void QHelpSearchQueryWidget_CloseEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QHelpSearchQueryWidget_CloseEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QHelpSearchQueryWidget_ContextMenuEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QHelpSearchQueryWidget_ContextMenuEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QHelpSearchQueryWidget_InitPainter(void* ptr, void* painter){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->initPainter(static_cast<QPainter*>(painter));
}

void QHelpSearchQueryWidget_InitPainterDefault(void* ptr, void* painter){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::initPainter(static_cast<QPainter*>(painter));
}

void QHelpSearchQueryWidget_InputMethodEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QHelpSearchQueryWidget_InputMethodEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QHelpSearchQueryWidget_KeyPressEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchQueryWidget_KeyPressEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchQueryWidget_KeyReleaseEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchQueryWidget_KeyReleaseEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchQueryWidget_MouseDoubleClickEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchQueryWidget_MouseDoubleClickEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchQueryWidget_MouseMoveEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchQueryWidget_MouseMoveEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchQueryWidget_MousePressEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchQueryWidget_MousePressEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchQueryWidget_MouseReleaseEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchQueryWidget_MouseReleaseEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchQueryWidget_ResizeEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QHelpSearchQueryWidget_ResizeEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QHelpSearchQueryWidget_TabletEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpSearchQueryWidget_TabletEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpSearchQueryWidget_WheelEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QHelpSearchQueryWidget_WheelEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QHelpSearchQueryWidget_TimerEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpSearchQueryWidget_TimerEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpSearchQueryWidget_ChildEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHelpSearchQueryWidget_ChildEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpSearchQueryWidget_CustomEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchQueryWidget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHelpSearchQueryWidget_CustomEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchQueryWidget*>(ptr)->QHelpSearchQueryWidget::customEvent(static_cast<QEvent*>(event));
}

class MyQHelpSearchResultWidget: public QHelpSearchResultWidget {
public:
	void Signal_RequestShowLink(const QUrl & link) { callbackQHelpSearchResultWidgetRequestShowLink(this, this->objectName().toUtf8().data(), new QUrl(link)); };
	void actionEvent(QActionEvent * event) { callbackQHelpSearchResultWidgetActionEvent(this, this->objectName().toUtf8().data(), event); };
	void dragEnterEvent(QDragEnterEvent * event) { callbackQHelpSearchResultWidgetDragEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void dragLeaveEvent(QDragLeaveEvent * event) { callbackQHelpSearchResultWidgetDragLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void dragMoveEvent(QDragMoveEvent * event) { callbackQHelpSearchResultWidgetDragMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void dropEvent(QDropEvent * event) { callbackQHelpSearchResultWidgetDropEvent(this, this->objectName().toUtf8().data(), event); };
	void enterEvent(QEvent * event) { callbackQHelpSearchResultWidgetEnterEvent(this, this->objectName().toUtf8().data(), event); };
	void focusInEvent(QFocusEvent * event) { callbackQHelpSearchResultWidgetFocusInEvent(this, this->objectName().toUtf8().data(), event); };
	void focusOutEvent(QFocusEvent * event) { callbackQHelpSearchResultWidgetFocusOutEvent(this, this->objectName().toUtf8().data(), event); };
	void hideEvent(QHideEvent * event) { callbackQHelpSearchResultWidgetHideEvent(this, this->objectName().toUtf8().data(), event); };
	void leaveEvent(QEvent * event) { callbackQHelpSearchResultWidgetLeaveEvent(this, this->objectName().toUtf8().data(), event); };
	void moveEvent(QMoveEvent * event) { callbackQHelpSearchResultWidgetMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void paintEvent(QPaintEvent * event) { callbackQHelpSearchResultWidgetPaintEvent(this, this->objectName().toUtf8().data(), event); };
	void setVisible(bool visible) { if (!callbackQHelpSearchResultWidgetSetVisible(this, this->objectName().toUtf8().data(), visible)) { QHelpSearchResultWidget::setVisible(visible); }; };
	void showEvent(QShowEvent * event) { callbackQHelpSearchResultWidgetShowEvent(this, this->objectName().toUtf8().data(), event); };
	void closeEvent(QCloseEvent * event) { callbackQHelpSearchResultWidgetCloseEvent(this, this->objectName().toUtf8().data(), event); };
	void contextMenuEvent(QContextMenuEvent * event) { callbackQHelpSearchResultWidgetContextMenuEvent(this, this->objectName().toUtf8().data(), event); };
	void initPainter(QPainter * painter) const { callbackQHelpSearchResultWidgetInitPainter(const_cast<MyQHelpSearchResultWidget*>(this), this->objectName().toUtf8().data(), painter); };
	void inputMethodEvent(QInputMethodEvent * event) { callbackQHelpSearchResultWidgetInputMethodEvent(this, this->objectName().toUtf8().data(), event); };
	void keyPressEvent(QKeyEvent * event) { callbackQHelpSearchResultWidgetKeyPressEvent(this, this->objectName().toUtf8().data(), event); };
	void keyReleaseEvent(QKeyEvent * event) { callbackQHelpSearchResultWidgetKeyReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseDoubleClickEvent(QMouseEvent * event) { callbackQHelpSearchResultWidgetMouseDoubleClickEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseMoveEvent(QMouseEvent * event) { callbackQHelpSearchResultWidgetMouseMoveEvent(this, this->objectName().toUtf8().data(), event); };
	void mousePressEvent(QMouseEvent * event) { callbackQHelpSearchResultWidgetMousePressEvent(this, this->objectName().toUtf8().data(), event); };
	void mouseReleaseEvent(QMouseEvent * event) { callbackQHelpSearchResultWidgetMouseReleaseEvent(this, this->objectName().toUtf8().data(), event); };
	void resizeEvent(QResizeEvent * event) { callbackQHelpSearchResultWidgetResizeEvent(this, this->objectName().toUtf8().data(), event); };
	void tabletEvent(QTabletEvent * event) { callbackQHelpSearchResultWidgetTabletEvent(this, this->objectName().toUtf8().data(), event); };
	void wheelEvent(QWheelEvent * event) { callbackQHelpSearchResultWidgetWheelEvent(this, this->objectName().toUtf8().data(), event); };
	void timerEvent(QTimerEvent * event) { callbackQHelpSearchResultWidgetTimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQHelpSearchResultWidgetChildEvent(this, this->objectName().toUtf8().data(), event); };
	void customEvent(QEvent * event) { callbackQHelpSearchResultWidgetCustomEvent(this, this->objectName().toUtf8().data(), event); };
};

void* QHelpSearchResultWidget_LinkAt(void* ptr, void* point){
	return new QUrl(static_cast<QHelpSearchResultWidget*>(ptr)->linkAt(*static_cast<QPoint*>(point)));
}

void QHelpSearchResultWidget_ConnectRequestShowLink(void* ptr){
	QObject::connect(static_cast<QHelpSearchResultWidget*>(ptr), static_cast<void (QHelpSearchResultWidget::*)(const QUrl &)>(&QHelpSearchResultWidget::requestShowLink), static_cast<MyQHelpSearchResultWidget*>(ptr), static_cast<void (MyQHelpSearchResultWidget::*)(const QUrl &)>(&MyQHelpSearchResultWidget::Signal_RequestShowLink));;
}

void QHelpSearchResultWidget_DisconnectRequestShowLink(void* ptr){
	QObject::disconnect(static_cast<QHelpSearchResultWidget*>(ptr), static_cast<void (QHelpSearchResultWidget::*)(const QUrl &)>(&QHelpSearchResultWidget::requestShowLink), static_cast<MyQHelpSearchResultWidget*>(ptr), static_cast<void (MyQHelpSearchResultWidget::*)(const QUrl &)>(&MyQHelpSearchResultWidget::Signal_RequestShowLink));;
}

void QHelpSearchResultWidget_RequestShowLink(void* ptr, void* link){
	static_cast<QHelpSearchResultWidget*>(ptr)->requestShowLink(*static_cast<QUrl*>(link));
}

void QHelpSearchResultWidget_DestroyQHelpSearchResultWidget(void* ptr){
	static_cast<QHelpSearchResultWidget*>(ptr)->~QHelpSearchResultWidget();
}

void QHelpSearchResultWidget_ActionEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpSearchResultWidget_ActionEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::actionEvent(static_cast<QActionEvent*>(event));
}

void QHelpSearchResultWidget_DragEnterEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpSearchResultWidget_DragEnterEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::dragEnterEvent(static_cast<QDragEnterEvent*>(event));
}

void QHelpSearchResultWidget_DragLeaveEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QHelpSearchResultWidget_DragLeaveEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::dragLeaveEvent(static_cast<QDragLeaveEvent*>(event));
}

void QHelpSearchResultWidget_DragMoveEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QHelpSearchResultWidget_DragMoveEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::dragMoveEvent(static_cast<QDragMoveEvent*>(event));
}

void QHelpSearchResultWidget_DropEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->dropEvent(static_cast<QDropEvent*>(event));
}

void QHelpSearchResultWidget_DropEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::dropEvent(static_cast<QDropEvent*>(event));
}

void QHelpSearchResultWidget_EnterEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->enterEvent(static_cast<QEvent*>(event));
}

void QHelpSearchResultWidget_EnterEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::enterEvent(static_cast<QEvent*>(event));
}

void QHelpSearchResultWidget_FocusInEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(event));
}

void QHelpSearchResultWidget_FocusInEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::focusInEvent(static_cast<QFocusEvent*>(event));
}

void QHelpSearchResultWidget_FocusOutEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpSearchResultWidget_FocusOutEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::focusOutEvent(static_cast<QFocusEvent*>(event));
}

void QHelpSearchResultWidget_HideEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpSearchResultWidget_HideEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::hideEvent(static_cast<QHideEvent*>(event));
}

void QHelpSearchResultWidget_LeaveEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->leaveEvent(static_cast<QEvent*>(event));
}

void QHelpSearchResultWidget_LeaveEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::leaveEvent(static_cast<QEvent*>(event));
}

void QHelpSearchResultWidget_MoveEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpSearchResultWidget_MoveEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::moveEvent(static_cast<QMoveEvent*>(event));
}

void QHelpSearchResultWidget_PaintEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QHelpSearchResultWidget_PaintEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::paintEvent(static_cast<QPaintEvent*>(event));
}

void QHelpSearchResultWidget_SetVisible(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<MyQHelpSearchResultWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QHelpSearchResultWidget_SetVisibleDefault(void* ptr, int visible){
	QMetaObject::invokeMethod(static_cast<QHelpSearchResultWidget*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QHelpSearchResultWidget_ShowEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->showEvent(static_cast<QShowEvent*>(event));
}

void QHelpSearchResultWidget_ShowEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::showEvent(static_cast<QShowEvent*>(event));
}

void QHelpSearchResultWidget_CloseEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->closeEvent(static_cast<QCloseEvent*>(event));
}

void QHelpSearchResultWidget_CloseEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::closeEvent(static_cast<QCloseEvent*>(event));
}

void QHelpSearchResultWidget_ContextMenuEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QHelpSearchResultWidget_ContextMenuEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::contextMenuEvent(static_cast<QContextMenuEvent*>(event));
}

void QHelpSearchResultWidget_InitPainter(void* ptr, void* painter){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->initPainter(static_cast<QPainter*>(painter));
}

void QHelpSearchResultWidget_InitPainterDefault(void* ptr, void* painter){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::initPainter(static_cast<QPainter*>(painter));
}

void QHelpSearchResultWidget_InputMethodEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QHelpSearchResultWidget_InputMethodEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::inputMethodEvent(static_cast<QInputMethodEvent*>(event));
}

void QHelpSearchResultWidget_KeyPressEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchResultWidget_KeyPressEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::keyPressEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchResultWidget_KeyReleaseEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchResultWidget_KeyReleaseEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::keyReleaseEvent(static_cast<QKeyEvent*>(event));
}

void QHelpSearchResultWidget_MouseDoubleClickEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchResultWidget_MouseDoubleClickEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::mouseDoubleClickEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchResultWidget_MouseMoveEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchResultWidget_MouseMoveEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::mouseMoveEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchResultWidget_MousePressEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchResultWidget_MousePressEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::mousePressEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchResultWidget_MouseReleaseEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchResultWidget_MouseReleaseEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::mouseReleaseEvent(static_cast<QMouseEvent*>(event));
}

void QHelpSearchResultWidget_ResizeEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(event));
}

void QHelpSearchResultWidget_ResizeEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::resizeEvent(static_cast<QResizeEvent*>(event));
}

void QHelpSearchResultWidget_TabletEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpSearchResultWidget_TabletEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::tabletEvent(static_cast<QTabletEvent*>(event));
}

void QHelpSearchResultWidget_WheelEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(event));
}

void QHelpSearchResultWidget_WheelEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::wheelEvent(static_cast<QWheelEvent*>(event));
}

void QHelpSearchResultWidget_TimerEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpSearchResultWidget_TimerEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::timerEvent(static_cast<QTimerEvent*>(event));
}

void QHelpSearchResultWidget_ChildEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QHelpSearchResultWidget_ChildEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::childEvent(static_cast<QChildEvent*>(event));
}

void QHelpSearchResultWidget_CustomEvent(void* ptr, void* event){
	static_cast<MyQHelpSearchResultWidget*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QHelpSearchResultWidget_CustomEventDefault(void* ptr, void* event){
	static_cast<QHelpSearchResultWidget*>(ptr)->QHelpSearchResultWidget::customEvent(static_cast<QEvent*>(event));
}

