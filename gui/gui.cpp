#define protected public
#define private public

#include "gui.h"
#include "_cgo_export.h"

#include <QAbstractTextDocumentLayout>
#include <QAccessible>
#include <QAccessibleActionInterface>
#include <QAccessibleEditableTextInterface>
#include <QAccessibleEvent>
#include <QAccessibleInterface>
#include <QAccessibleObject>
#include <QAccessiblePlugin>
#include <QAccessibleStateChangeEvent>
#include <QAccessibleTableCellInterface>
#include <QAccessibleTableInterface>
#include <QAccessibleTableModelChangeEvent>
#include <QAccessibleTextCursorEvent>
#include <QAccessibleTextInsertEvent>
#include <QAccessibleTextInterface>
#include <QAccessibleTextRemoveEvent>
#include <QAccessibleTextSelectionEvent>
#include <QAccessibleTextUpdateEvent>
#include <QAccessibleValueChangeEvent>
#include <QAccessibleValueInterface>
#include <QAction>
#include <QActionEvent>
#include <QBackingStore>
#include <QBitmap>
#include <QBrush>
#include <QByteArray>
#include <QChar>
#include <QChildEvent>
#include <QClipboard>
#include <QCloseEvent>
#include <QColor>
#include <QConicalGradient>
#include <QContextMenuEvent>
#include <QCursor>
#include <QDesktopServices>
#include <QDoubleValidator>
#include <QDrag>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEnterEvent>
#include <QEvent>
#include <QExposeEvent>
#include <QFile>
#include <QFileOpenEvent>
#include <QFocusEvent>
#include <QFont>
#include <QFontDatabase>
#include <QFontInfo>
#include <QFontMetrics>
#include <QFontMetricsF>
#include <QGenericPlugin>
#include <QGenericPluginFactory>
#include <QGlyphRun>
#include <QGradient>
#include <QGuiApplication>
#include <QHelpEvent>
#include <QHideEvent>
#include <QHoverEvent>
#include <QIODevice>
#include <QIcon>
#include <QIconDragEvent>
#include <QIconEngine>
#include <QIconEnginePlugin>
#include <QImage>
#include <QImageIOHandler>
#include <QImageIOPlugin>
#include <QImageReader>
#include <QImageWriter>
#include <QInputEvent>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QInputMethodQueryEvent>
#include <QIntValidator>
#include <QKeyEvent>
#include <QKeySequence>
#include <QLine>
#include <QLineF>
#include <QLinearGradient>
#include <QList>
#include <QLocale>
#include <QMargins>
#include <QMarginsF>
#include <QMatrix4x4>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMimeData>
#include <QModelIndex>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QMovie>
#include <QNativeGestureEvent>
#include <QObject>
#include <QOffscreenSurface>
#include <QPageLayout>
#include <QPageSize>
#include <QPagedPaintDevice>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPaintEngine>
#include <QPaintEngineState>
#include <QPaintEvent>
#include <QPainter>
#include <QPainterPath>
#include <QPainterPathStroker>
#include <QPalette>
#include <QPdfWriter>
#include <QPen>
#include <QPicture>
#include <QPixelFormat>
#include <QPixmap>
#include <QPixmapCache>
#include <QPlatformSurfaceEvent>
#include <QPoint>
#include <QPointF>
#include <QPolygon>
#include <QPolygonF>
#include <QQuaternion>
#include <QRadialGradient>
#include <QRasterWindow>
#include <QRawFont>
#include <QRect>
#include <QRectF>
#include <QRegExp>
#include <QRegExpValidator>
#include <QRegion>
#include <QRegularExpression>
#include <QRegularExpressionValidator>
#include <QResizeEvent>
#include <QScreen>
#include <QScrollEvent>
#include <QScrollPrepareEvent>
#include <QSessionManager>
#include <QShortcut>
#include <QShortcutEvent>
#include <QShowEvent>
#include <QSize>
#include <QSizeF>
#include <QStandardItem>
#include <QStandardItemModel>
#include <QStaticText>
#include <QStatusTipEvent>
#include <QString>
#include <QStringList>
#include <QStyle>
#include <QStyleHints>
#include <QSurface>
#include <QSurfaceFormat>
#include <QSyntaxHighlighter>
#include <QTabletEvent>
#include <QTextBlock>
#include <QTextBlockFormat>
#include <QTextBlockGroup>
#include <QTextBlockUserData>
#include <QTextCharFormat>
#include <QTextCodec>
#include <QTextCursor>
#include <QTextDocument>
#include <QTextDocumentFragment>
#include <QTextDocumentWriter>
#include <QTextFormat>
#include <QTextFragment>
#include <QTextFrame>
#include <QTextFrameFormat>
#include <QTextImageFormat>
#include <QTextInlineObject>
#include <QTextItem>
#include <QTextLayout>
#include <QTextLength>
#include <QTextLine>
#include <QTextList>
#include <QTextListFormat>
#include <QTextObject>
#include <QTextObjectInterface>
#include <QTextOption>
#include <QTextTable>
#include <QTextTableCell>
#include <QTextTableCellFormat>
#include <QTextTableFormat>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QTouchDevice>
#include <QTouchEvent>
#include <QTransform>
#include <QUrl>
#include <QValidator>
#include <QVariant>
#include <QVector>
#include <QVector2D>
#include <QVector3D>
#include <QVector4D>
#include <QWhatsThis>
#include <QWhatsThisClickedEvent>
#include <QWheelEvent>
#include <QWindow>
#include <QWindowStateChangeEvent>

class MyQAbstractTextDocumentLayout: public QAbstractTextDocumentLayout
{
public:
	QRectF blockBoundingRect(const QTextBlock & block) const { return *static_cast<QRectF*>(callbackQAbstractTextDocumentLayout_BlockBoundingRect(const_cast<MyQAbstractTextDocumentLayout*>(this), this->objectName().toUtf8().data(), new QTextBlock(block))); };
	void documentChanged(int position, int charsRemoved, int charsAdded) { callbackQAbstractTextDocumentLayout_DocumentChanged(this, this->objectName().toUtf8().data(), position, charsRemoved, charsAdded); };
	QSizeF documentSize() const { return *static_cast<QSizeF*>(callbackQAbstractTextDocumentLayout_DocumentSize(const_cast<MyQAbstractTextDocumentLayout*>(this), this->objectName().toUtf8().data())); };
	void Signal_DocumentSizeChanged(const QSizeF & newSize) { callbackQAbstractTextDocumentLayout_DocumentSizeChanged(this, this->objectName().toUtf8().data(), new QSizeF(static_cast<QSizeF>(newSize).width(), static_cast<QSizeF>(newSize).height())); };
	QRectF frameBoundingRect(QTextFrame * frame) const { return *static_cast<QRectF*>(callbackQAbstractTextDocumentLayout_FrameBoundingRect(const_cast<MyQAbstractTextDocumentLayout*>(this), this->objectName().toUtf8().data(), frame)); };
	int hitTest(const QPointF & point, Qt::HitTestAccuracy accuracy) const { return callbackQAbstractTextDocumentLayout_HitTest(const_cast<MyQAbstractTextDocumentLayout*>(this), this->objectName().toUtf8().data(), new QPointF(static_cast<QPointF>(point).x(), static_cast<QPointF>(point).y()), accuracy); };
	int pageCount() const { return callbackQAbstractTextDocumentLayout_PageCount(const_cast<MyQAbstractTextDocumentLayout*>(this), this->objectName().toUtf8().data()); };
	void Signal_PageCountChanged(int newPages) { callbackQAbstractTextDocumentLayout_PageCountChanged(this, this->objectName().toUtf8().data(), newPages); };
	void Signal_Update(const QRectF & rect) { callbackQAbstractTextDocumentLayout_Update(this, this->objectName().toUtf8().data(), new QRectF(static_cast<QRectF>(rect).x(), static_cast<QRectF>(rect).y(), static_cast<QRectF>(rect).width(), static_cast<QRectF>(rect).height())); };
	void Signal_UpdateBlock(const QTextBlock & block) { callbackQAbstractTextDocumentLayout_UpdateBlock(this, this->objectName().toUtf8().data(), new QTextBlock(block)); };
	void timerEvent(QTimerEvent * event) { callbackQAbstractTextDocumentLayout_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAbstractTextDocumentLayout_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAbstractTextDocumentLayout_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQAbstractTextDocumentLayout_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQAbstractTextDocumentLayout_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAbstractTextDocumentLayout_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQAbstractTextDocumentLayout_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAbstractTextDocumentLayout_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAbstractTextDocumentLayout_MetaObject(const_cast<MyQAbstractTextDocumentLayout*>(this), this->objectName().toUtf8().data())); };
};

char* QAbstractTextDocumentLayout_AnchorAt(void* ptr, void* position)
{
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->anchorAt(*static_cast<QPointF*>(position)).toUtf8().data();
}

void* QAbstractTextDocumentLayout_BlockBoundingRect(void* ptr, void* block)
{
	return new QRectF(static_cast<QRectF>(static_cast<QAbstractTextDocumentLayout*>(ptr)->blockBoundingRect(*static_cast<QTextBlock*>(block))).x(), static_cast<QRectF>(static_cast<QAbstractTextDocumentLayout*>(ptr)->blockBoundingRect(*static_cast<QTextBlock*>(block))).y(), static_cast<QRectF>(static_cast<QAbstractTextDocumentLayout*>(ptr)->blockBoundingRect(*static_cast<QTextBlock*>(block))).width(), static_cast<QRectF>(static_cast<QAbstractTextDocumentLayout*>(ptr)->blockBoundingRect(*static_cast<QTextBlock*>(block))).height());
}

void* QAbstractTextDocumentLayout_Document(void* ptr)
{
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->document();
}

void QAbstractTextDocumentLayout_DocumentChanged(void* ptr, int position, int charsRemoved, int charsAdded)
{
	static_cast<QAbstractTextDocumentLayout*>(ptr)->documentChanged(position, charsRemoved, charsAdded);
}

void* QAbstractTextDocumentLayout_DocumentSize(void* ptr)
{
	return new QSizeF(static_cast<QSizeF>(static_cast<QAbstractTextDocumentLayout*>(ptr)->documentSize()).width(), static_cast<QSizeF>(static_cast<QAbstractTextDocumentLayout*>(ptr)->documentSize()).height());
}

void QAbstractTextDocumentLayout_ConnectDocumentSizeChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractTextDocumentLayout*>(ptr), static_cast<void (QAbstractTextDocumentLayout::*)(const QSizeF &)>(&QAbstractTextDocumentLayout::documentSizeChanged), static_cast<MyQAbstractTextDocumentLayout*>(ptr), static_cast<void (MyQAbstractTextDocumentLayout::*)(const QSizeF &)>(&MyQAbstractTextDocumentLayout::Signal_DocumentSizeChanged));
}

void QAbstractTextDocumentLayout_DisconnectDocumentSizeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractTextDocumentLayout*>(ptr), static_cast<void (QAbstractTextDocumentLayout::*)(const QSizeF &)>(&QAbstractTextDocumentLayout::documentSizeChanged), static_cast<MyQAbstractTextDocumentLayout*>(ptr), static_cast<void (MyQAbstractTextDocumentLayout::*)(const QSizeF &)>(&MyQAbstractTextDocumentLayout::Signal_DocumentSizeChanged));
}

void QAbstractTextDocumentLayout_DocumentSizeChanged(void* ptr, void* newSize)
{
	static_cast<QAbstractTextDocumentLayout*>(ptr)->documentSizeChanged(*static_cast<QSizeF*>(newSize));
}

void* QAbstractTextDocumentLayout_FrameBoundingRect(void* ptr, void* frame)
{
	return new QRectF(static_cast<QRectF>(static_cast<QAbstractTextDocumentLayout*>(ptr)->frameBoundingRect(static_cast<QTextFrame*>(frame))).x(), static_cast<QRectF>(static_cast<QAbstractTextDocumentLayout*>(ptr)->frameBoundingRect(static_cast<QTextFrame*>(frame))).y(), static_cast<QRectF>(static_cast<QAbstractTextDocumentLayout*>(ptr)->frameBoundingRect(static_cast<QTextFrame*>(frame))).width(), static_cast<QRectF>(static_cast<QAbstractTextDocumentLayout*>(ptr)->frameBoundingRect(static_cast<QTextFrame*>(frame))).height());
}

void* QAbstractTextDocumentLayout_HandlerForObject(void* ptr, int objectType)
{
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->handlerForObject(objectType);
}

int QAbstractTextDocumentLayout_HitTest(void* ptr, void* point, int accuracy)
{
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->hitTest(*static_cast<QPointF*>(point), static_cast<Qt::HitTestAccuracy>(accuracy));
}

int QAbstractTextDocumentLayout_PageCount(void* ptr)
{
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->pageCount();
}

void QAbstractTextDocumentLayout_ConnectPageCountChanged(void* ptr)
{
	QObject::connect(static_cast<QAbstractTextDocumentLayout*>(ptr), static_cast<void (QAbstractTextDocumentLayout::*)(int)>(&QAbstractTextDocumentLayout::pageCountChanged), static_cast<MyQAbstractTextDocumentLayout*>(ptr), static_cast<void (MyQAbstractTextDocumentLayout::*)(int)>(&MyQAbstractTextDocumentLayout::Signal_PageCountChanged));
}

void QAbstractTextDocumentLayout_DisconnectPageCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractTextDocumentLayout*>(ptr), static_cast<void (QAbstractTextDocumentLayout::*)(int)>(&QAbstractTextDocumentLayout::pageCountChanged), static_cast<MyQAbstractTextDocumentLayout*>(ptr), static_cast<void (MyQAbstractTextDocumentLayout::*)(int)>(&MyQAbstractTextDocumentLayout::Signal_PageCountChanged));
}

void QAbstractTextDocumentLayout_PageCountChanged(void* ptr, int newPages)
{
	static_cast<QAbstractTextDocumentLayout*>(ptr)->pageCountChanged(newPages);
}

void* QAbstractTextDocumentLayout_PaintDevice(void* ptr)
{
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->paintDevice();
}

void QAbstractTextDocumentLayout_RegisterHandler(void* ptr, int objectType, void* component)
{
	static_cast<QAbstractTextDocumentLayout*>(ptr)->registerHandler(objectType, static_cast<QObject*>(component));
}

void QAbstractTextDocumentLayout_SetPaintDevice(void* ptr, void* device)
{
	static_cast<QAbstractTextDocumentLayout*>(ptr)->setPaintDevice(static_cast<QPaintDevice*>(device));
}

void QAbstractTextDocumentLayout_UnregisterHandler(void* ptr, int objectType, void* component)
{
	static_cast<QAbstractTextDocumentLayout*>(ptr)->unregisterHandler(objectType, static_cast<QObject*>(component));
}

void QAbstractTextDocumentLayout_ConnectUpdate(void* ptr)
{
	QObject::connect(static_cast<QAbstractTextDocumentLayout*>(ptr), static_cast<void (QAbstractTextDocumentLayout::*)(const QRectF &)>(&QAbstractTextDocumentLayout::update), static_cast<MyQAbstractTextDocumentLayout*>(ptr), static_cast<void (MyQAbstractTextDocumentLayout::*)(const QRectF &)>(&MyQAbstractTextDocumentLayout::Signal_Update));
}

void QAbstractTextDocumentLayout_DisconnectUpdate(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractTextDocumentLayout*>(ptr), static_cast<void (QAbstractTextDocumentLayout::*)(const QRectF &)>(&QAbstractTextDocumentLayout::update), static_cast<MyQAbstractTextDocumentLayout*>(ptr), static_cast<void (MyQAbstractTextDocumentLayout::*)(const QRectF &)>(&MyQAbstractTextDocumentLayout::Signal_Update));
}

void QAbstractTextDocumentLayout_Update(void* ptr, void* rect)
{
	static_cast<QAbstractTextDocumentLayout*>(ptr)->update(*static_cast<QRectF*>(rect));
}

void QAbstractTextDocumentLayout_ConnectUpdateBlock(void* ptr)
{
	QObject::connect(static_cast<QAbstractTextDocumentLayout*>(ptr), static_cast<void (QAbstractTextDocumentLayout::*)(const QTextBlock &)>(&QAbstractTextDocumentLayout::updateBlock), static_cast<MyQAbstractTextDocumentLayout*>(ptr), static_cast<void (MyQAbstractTextDocumentLayout::*)(const QTextBlock &)>(&MyQAbstractTextDocumentLayout::Signal_UpdateBlock));
}

void QAbstractTextDocumentLayout_DisconnectUpdateBlock(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractTextDocumentLayout*>(ptr), static_cast<void (QAbstractTextDocumentLayout::*)(const QTextBlock &)>(&QAbstractTextDocumentLayout::updateBlock), static_cast<MyQAbstractTextDocumentLayout*>(ptr), static_cast<void (MyQAbstractTextDocumentLayout::*)(const QTextBlock &)>(&MyQAbstractTextDocumentLayout::Signal_UpdateBlock));
}

void QAbstractTextDocumentLayout_UpdateBlock(void* ptr, void* block)
{
	static_cast<QAbstractTextDocumentLayout*>(ptr)->updateBlock(*static_cast<QTextBlock*>(block));
}

void QAbstractTextDocumentLayout_TimerEvent(void* ptr, void* event)
{
	static_cast<QAbstractTextDocumentLayout*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractTextDocumentLayout_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractTextDocumentLayout*>(ptr)->QAbstractTextDocumentLayout::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAbstractTextDocumentLayout_ChildEvent(void* ptr, void* event)
{
	static_cast<QAbstractTextDocumentLayout*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractTextDocumentLayout_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractTextDocumentLayout*>(ptr)->QAbstractTextDocumentLayout::childEvent(static_cast<QChildEvent*>(event));
}

void QAbstractTextDocumentLayout_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAbstractTextDocumentLayout*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractTextDocumentLayout_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAbstractTextDocumentLayout*>(ptr)->QAbstractTextDocumentLayout::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractTextDocumentLayout_CustomEvent(void* ptr, void* event)
{
	static_cast<QAbstractTextDocumentLayout*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAbstractTextDocumentLayout_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAbstractTextDocumentLayout*>(ptr)->QAbstractTextDocumentLayout::customEvent(static_cast<QEvent*>(event));
}

void QAbstractTextDocumentLayout_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractTextDocumentLayout*>(ptr), "deleteLater");
}

void QAbstractTextDocumentLayout_DeleteLaterDefault(void* ptr)
{
	static_cast<QAbstractTextDocumentLayout*>(ptr)->QAbstractTextDocumentLayout::deleteLater();
}

void QAbstractTextDocumentLayout_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAbstractTextDocumentLayout*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAbstractTextDocumentLayout_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAbstractTextDocumentLayout*>(ptr)->QAbstractTextDocumentLayout::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QAbstractTextDocumentLayout_Event(void* ptr, void* e)
{
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->event(static_cast<QEvent*>(e));
}

int QAbstractTextDocumentLayout_EventDefault(void* ptr, void* e)
{
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->QAbstractTextDocumentLayout::event(static_cast<QEvent*>(e));
}

int QAbstractTextDocumentLayout_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QAbstractTextDocumentLayout_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAbstractTextDocumentLayout*>(ptr)->QAbstractTextDocumentLayout::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAbstractTextDocumentLayout_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAbstractTextDocumentLayout*>(ptr)->metaObject());
}

void* QAbstractTextDocumentLayout_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAbstractTextDocumentLayout*>(ptr)->QAbstractTextDocumentLayout::metaObject());
}

int QAccessible_InvalidEvent_Type()
{
	return QAccessible::InvalidEvent;
}

int QAccessible_QAccessible_IsActive()
{
	return QAccessible::isActive();
}

void* QAccessible_QAccessible_QueryAccessibleInterface(void* object)
{
	return QAccessible::queryAccessibleInterface(static_cast<QObject*>(object));
}

void QAccessible_QAccessible_SetRootObject(void* object)
{
	QAccessible::setRootObject(static_cast<QObject*>(object));
}

void QAccessible_QAccessible_UpdateAccessibility(void* event)
{
	QAccessible::updateAccessibility(static_cast<QAccessibleEvent*>(event));
}

class MyQAccessibleActionInterface: public QAccessibleActionInterface
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	QString localizedActionDescription(const QString & actionName) const { return QString(callbackQAccessibleActionInterface_LocalizedActionDescription(const_cast<MyQAccessibleActionInterface*>(this), this->objectNameAbs().toUtf8().data(), actionName.toUtf8().data())); };
	QString localizedActionName(const QString & actionName) const { return QString(callbackQAccessibleActionInterface_LocalizedActionName(const_cast<MyQAccessibleActionInterface*>(this), this->objectNameAbs().toUtf8().data(), actionName.toUtf8().data())); };
	QStringList actionNames() const { return QString(callbackQAccessibleActionInterface_ActionNames(const_cast<MyQAccessibleActionInterface*>(this), this->objectNameAbs().toUtf8().data())).split("|", QString::SkipEmptyParts); };
	void doAction(const QString & actionName) { callbackQAccessibleActionInterface_DoAction(this, this->objectNameAbs().toUtf8().data(), actionName.toUtf8().data()); };
	QStringList keyBindingsForAction(const QString & actionName) const { return QString(callbackQAccessibleActionInterface_KeyBindingsForAction(const_cast<MyQAccessibleActionInterface*>(this), this->objectNameAbs().toUtf8().data(), actionName.toUtf8().data())).split("|", QString::SkipEmptyParts); };
};

char* QAccessibleActionInterface_LocalizedActionDescription(void* ptr, char* actionName)
{
	return static_cast<QAccessibleActionInterface*>(ptr)->localizedActionDescription(QString(actionName)).toUtf8().data();
}

char* QAccessibleActionInterface_LocalizedActionDescriptionDefault(void* ptr, char* actionName)
{
	return static_cast<QAccessibleActionInterface*>(ptr)->QAccessibleActionInterface::localizedActionDescription(QString(actionName)).toUtf8().data();
}

char* QAccessibleActionInterface_LocalizedActionName(void* ptr, char* actionName)
{
	return static_cast<QAccessibleActionInterface*>(ptr)->localizedActionName(QString(actionName)).toUtf8().data();
}

char* QAccessibleActionInterface_LocalizedActionNameDefault(void* ptr, char* actionName)
{
	return static_cast<QAccessibleActionInterface*>(ptr)->QAccessibleActionInterface::localizedActionName(QString(actionName)).toUtf8().data();
}

char* QAccessibleActionInterface_ActionNames(void* ptr)
{
	return static_cast<QAccessibleActionInterface*>(ptr)->actionNames().join("|").toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_DecreaseAction()
{
	return QAccessibleActionInterface::decreaseAction().toUtf8().data();
}

void QAccessibleActionInterface_DoAction(void* ptr, char* actionName)
{
	static_cast<QAccessibleActionInterface*>(ptr)->doAction(QString(actionName));
}

char* QAccessibleActionInterface_QAccessibleActionInterface_IncreaseAction()
{
	return QAccessibleActionInterface::increaseAction().toUtf8().data();
}

char* QAccessibleActionInterface_KeyBindingsForAction(void* ptr, char* actionName)
{
	return static_cast<QAccessibleActionInterface*>(ptr)->keyBindingsForAction(QString(actionName)).join("|").toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_NextPageAction()
{
	return QAccessibleActionInterface::nextPageAction().toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_PressAction()
{
	return QAccessibleActionInterface::pressAction().toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_PreviousPageAction()
{
	return QAccessibleActionInterface::previousPageAction().toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_ScrollDownAction()
{
	return QAccessibleActionInterface::scrollDownAction().toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_ScrollLeftAction()
{
	return QAccessibleActionInterface::scrollLeftAction().toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_ScrollRightAction()
{
	return QAccessibleActionInterface::scrollRightAction().toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_ScrollUpAction()
{
	return QAccessibleActionInterface::scrollUpAction().toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_SetFocusAction()
{
	return QAccessibleActionInterface::setFocusAction().toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_ShowMenuAction()
{
	return QAccessibleActionInterface::showMenuAction().toUtf8().data();
}

char* QAccessibleActionInterface_QAccessibleActionInterface_ToggleAction()
{
	return QAccessibleActionInterface::toggleAction().toUtf8().data();
}

void QAccessibleActionInterface_DestroyQAccessibleActionInterface(void* ptr)
{
	static_cast<QAccessibleActionInterface*>(ptr)->~QAccessibleActionInterface();
}

char* QAccessibleActionInterface_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQAccessibleActionInterface*>(static_cast<QAccessibleActionInterface*>(ptr))) {
		return static_cast<MyQAccessibleActionInterface*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QAccessibleActionInterface_BASE").toUtf8().data();
}

void QAccessibleActionInterface_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQAccessibleActionInterface*>(static_cast<QAccessibleActionInterface*>(ptr))) {
		static_cast<MyQAccessibleActionInterface*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQAccessibleEditableTextInterface: public QAccessibleEditableTextInterface
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	void deleteText(int startOffset, int endOffset) { callbackQAccessibleEditableTextInterface_DeleteText(this, this->objectNameAbs().toUtf8().data(), startOffset, endOffset); };
	void insertText(int offset, const QString & text) { callbackQAccessibleEditableTextInterface_InsertText(this, this->objectNameAbs().toUtf8().data(), offset, text.toUtf8().data()); };
	void replaceText(int startOffset, int endOffset, const QString & text) { callbackQAccessibleEditableTextInterface_ReplaceText(this, this->objectNameAbs().toUtf8().data(), startOffset, endOffset, text.toUtf8().data()); };
};

void QAccessibleEditableTextInterface_DeleteText(void* ptr, int startOffset, int endOffset)
{
	static_cast<QAccessibleEditableTextInterface*>(ptr)->deleteText(startOffset, endOffset);
}

void QAccessibleEditableTextInterface_InsertText(void* ptr, int offset, char* text)
{
	static_cast<QAccessibleEditableTextInterface*>(ptr)->insertText(offset, QString(text));
}

void QAccessibleEditableTextInterface_ReplaceText(void* ptr, int startOffset, int endOffset, char* text)
{
	static_cast<QAccessibleEditableTextInterface*>(ptr)->replaceText(startOffset, endOffset, QString(text));
}

void QAccessibleEditableTextInterface_DestroyQAccessibleEditableTextInterface(void* ptr)
{
	static_cast<QAccessibleEditableTextInterface*>(ptr)->~QAccessibleEditableTextInterface();
}

char* QAccessibleEditableTextInterface_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQAccessibleEditableTextInterface*>(static_cast<QAccessibleEditableTextInterface*>(ptr))) {
		return static_cast<MyQAccessibleEditableTextInterface*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QAccessibleEditableTextInterface_BASE").toUtf8().data();
}

void QAccessibleEditableTextInterface_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQAccessibleEditableTextInterface*>(static_cast<QAccessibleEditableTextInterface*>(ptr))) {
		static_cast<MyQAccessibleEditableTextInterface*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQAccessibleEvent: public QAccessibleEvent
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQAccessibleEvent(QAccessibleInterface *interface, QAccessible::Event type) : QAccessibleEvent(interface, type) {};
	MyQAccessibleEvent(QObject *object, QAccessible::Event type) : QAccessibleEvent(object, type) {};
	QAccessibleInterface * accessibleInterface() const { return static_cast<QAccessibleInterface*>(callbackQAccessibleEvent_AccessibleInterface(const_cast<MyQAccessibleEvent*>(this), this->objectNameAbs().toUtf8().data())); };
};

void* QAccessibleEvent_NewQAccessibleEvent2(void* interfa, int ty)
{
	return new MyQAccessibleEvent(static_cast<QAccessibleInterface*>(interfa), static_cast<QAccessible::Event>(ty));
}

void* QAccessibleEvent_NewQAccessibleEvent(void* object, int ty)
{
	return new MyQAccessibleEvent(static_cast<QObject*>(object), static_cast<QAccessible::Event>(ty));
}

void* QAccessibleEvent_AccessibleInterface(void* ptr)
{
	return static_cast<QAccessibleEvent*>(ptr)->accessibleInterface();
}

void* QAccessibleEvent_AccessibleInterfaceDefault(void* ptr)
{
	return static_cast<QAccessibleEvent*>(ptr)->QAccessibleEvent::accessibleInterface();
}

int QAccessibleEvent_Child(void* ptr)
{
	return static_cast<QAccessibleEvent*>(ptr)->child();
}

void* QAccessibleEvent_Object(void* ptr)
{
	return static_cast<QAccessibleEvent*>(ptr)->object();
}

void QAccessibleEvent_SetChild(void* ptr, int child)
{
	static_cast<QAccessibleEvent*>(ptr)->setChild(child);
}

int QAccessibleEvent_Type(void* ptr)
{
	return static_cast<QAccessibleEvent*>(ptr)->type();
}

void QAccessibleEvent_DestroyQAccessibleEvent(void* ptr)
{
	static_cast<QAccessibleEvent*>(ptr)->~QAccessibleEvent();
}

char* QAccessibleEvent_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQAccessibleEvent*>(static_cast<QAccessibleEvent*>(ptr))) {
		return static_cast<MyQAccessibleEvent*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QAccessibleEvent_BASE").toUtf8().data();
}

void QAccessibleEvent_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQAccessibleEvent*>(static_cast<QAccessibleEvent*>(ptr))) {
		static_cast<MyQAccessibleEvent*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQAccessibleInterface: public QAccessibleInterface
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	QColor backgroundColor() const { return *static_cast<QColor*>(callbackQAccessibleInterface_BackgroundColor(const_cast<MyQAccessibleInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	QAccessibleInterface * child(int index) const { return static_cast<QAccessibleInterface*>(callbackQAccessibleInterface_Child(const_cast<MyQAccessibleInterface*>(this), this->objectNameAbs().toUtf8().data(), index)); };
	QAccessibleInterface * childAt(int x, int y) const { return static_cast<QAccessibleInterface*>(callbackQAccessibleInterface_ChildAt(const_cast<MyQAccessibleInterface*>(this), this->objectNameAbs().toUtf8().data(), x, y)); };
	int childCount() const { return callbackQAccessibleInterface_ChildCount(const_cast<MyQAccessibleInterface*>(this), this->objectNameAbs().toUtf8().data()); };
	QAccessibleInterface * focusChild() const { return static_cast<QAccessibleInterface*>(callbackQAccessibleInterface_FocusChild(const_cast<MyQAccessibleInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	QColor foregroundColor() const { return *static_cast<QColor*>(callbackQAccessibleInterface_ForegroundColor(const_cast<MyQAccessibleInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	int indexOfChild(const QAccessibleInterface * child) const { return callbackQAccessibleInterface_IndexOfChild(const_cast<MyQAccessibleInterface*>(this), this->objectNameAbs().toUtf8().data(), const_cast<QAccessibleInterface*>(child)); };
	void * interface_cast(QAccessible::InterfaceType ty) { return callbackQAccessibleInterface_Interface_cast(this, this->objectNameAbs().toUtf8().data(), ty); };
	bool isValid() const { return callbackQAccessibleInterface_IsValid(const_cast<MyQAccessibleInterface*>(this), this->objectNameAbs().toUtf8().data()) != 0; };
	QObject * object() const { return static_cast<QObject*>(callbackQAccessibleInterface_Object(const_cast<MyQAccessibleInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	QAccessibleInterface * parent() const { return static_cast<QAccessibleInterface*>(callbackQAccessibleInterface_Parent(const_cast<MyQAccessibleInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	QRect rect() const { return *static_cast<QRect*>(callbackQAccessibleInterface_Rect(const_cast<MyQAccessibleInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	QAccessible::Role role() const { return static_cast<QAccessible::Role>(callbackQAccessibleInterface_Role(const_cast<MyQAccessibleInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	void setText(QAccessible::Text t, const QString & text) { callbackQAccessibleInterface_SetText(this, this->objectNameAbs().toUtf8().data(), t, text.toUtf8().data()); };
	QString text(QAccessible::Text t) const { return QString(callbackQAccessibleInterface_Text(const_cast<MyQAccessibleInterface*>(this), this->objectNameAbs().toUtf8().data(), t)); };
	QWindow * window() const { return static_cast<QWindow*>(callbackQAccessibleInterface_Window(const_cast<MyQAccessibleInterface*>(this), this->objectNameAbs().toUtf8().data())); };
};

void QAccessibleInterface_DestroyQAccessibleInterface(void* ptr)
{
	static_cast<QAccessibleInterface*>(ptr)->~QAccessibleInterface();
}

void* QAccessibleInterface_ActionInterface(void* ptr)
{
	return static_cast<QAccessibleInterface*>(ptr)->actionInterface();
}

void* QAccessibleInterface_BackgroundColor(void* ptr)
{
	return new QColor(static_cast<QAccessibleInterface*>(ptr)->backgroundColor());
}

void* QAccessibleInterface_BackgroundColorDefault(void* ptr)
{
	return new QColor(static_cast<QAccessibleInterface*>(ptr)->QAccessibleInterface::backgroundColor());
}

void* QAccessibleInterface_Child(void* ptr, int index)
{
	return static_cast<QAccessibleInterface*>(ptr)->child(index);
}

void* QAccessibleInterface_ChildAt(void* ptr, int x, int y)
{
	return static_cast<QAccessibleInterface*>(ptr)->childAt(x, y);
}

int QAccessibleInterface_ChildCount(void* ptr)
{
	return static_cast<QAccessibleInterface*>(ptr)->childCount();
}

void* QAccessibleInterface_FocusChild(void* ptr)
{
	return static_cast<QAccessibleInterface*>(ptr)->focusChild();
}

void* QAccessibleInterface_FocusChildDefault(void* ptr)
{
	return static_cast<QAccessibleInterface*>(ptr)->QAccessibleInterface::focusChild();
}

void* QAccessibleInterface_ForegroundColor(void* ptr)
{
	return new QColor(static_cast<QAccessibleInterface*>(ptr)->foregroundColor());
}

void* QAccessibleInterface_ForegroundColorDefault(void* ptr)
{
	return new QColor(static_cast<QAccessibleInterface*>(ptr)->QAccessibleInterface::foregroundColor());
}

int QAccessibleInterface_IndexOfChild(void* ptr, void* child)
{
	return static_cast<QAccessibleInterface*>(ptr)->indexOfChild(static_cast<QAccessibleInterface*>(child));
}

void* QAccessibleInterface_Interface_cast(void* ptr, int ty)
{
	return static_cast<QAccessibleInterface*>(ptr)->interface_cast(static_cast<QAccessible::InterfaceType>(ty));
}

void* QAccessibleInterface_Interface_castDefault(void* ptr, int ty)
{
	return static_cast<QAccessibleInterface*>(ptr)->QAccessibleInterface::interface_cast(static_cast<QAccessible::InterfaceType>(ty));
}

int QAccessibleInterface_IsValid(void* ptr)
{
	return static_cast<QAccessibleInterface*>(ptr)->isValid();
}

void* QAccessibleInterface_Object(void* ptr)
{
	return static_cast<QAccessibleInterface*>(ptr)->object();
}

void* QAccessibleInterface_Parent(void* ptr)
{
	return static_cast<QAccessibleInterface*>(ptr)->parent();
}

void* QAccessibleInterface_Rect(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QAccessibleInterface*>(ptr)->rect()).x(), static_cast<QRect>(static_cast<QAccessibleInterface*>(ptr)->rect()).y(), static_cast<QRect>(static_cast<QAccessibleInterface*>(ptr)->rect()).width(), static_cast<QRect>(static_cast<QAccessibleInterface*>(ptr)->rect()).height());
}

int QAccessibleInterface_Role(void* ptr)
{
	return static_cast<QAccessibleInterface*>(ptr)->role();
}

void QAccessibleInterface_SetText(void* ptr, int t, char* text)
{
	static_cast<QAccessibleInterface*>(ptr)->setText(static_cast<QAccessible::Text>(t), QString(text));
}

void* QAccessibleInterface_TableCellInterface(void* ptr)
{
	return static_cast<QAccessibleInterface*>(ptr)->tableCellInterface();
}

void* QAccessibleInterface_TableInterface(void* ptr)
{
	return static_cast<QAccessibleInterface*>(ptr)->tableInterface();
}

char* QAccessibleInterface_Text(void* ptr, int t)
{
	return static_cast<QAccessibleInterface*>(ptr)->text(static_cast<QAccessible::Text>(t)).toUtf8().data();
}

void* QAccessibleInterface_TextInterface(void* ptr)
{
	return static_cast<QAccessibleInterface*>(ptr)->textInterface();
}

void* QAccessibleInterface_ValueInterface(void* ptr)
{
	return static_cast<QAccessibleInterface*>(ptr)->valueInterface();
}

void* QAccessibleInterface_Window(void* ptr)
{
	return static_cast<QAccessibleInterface*>(ptr)->window();
}

void* QAccessibleInterface_WindowDefault(void* ptr)
{
	return static_cast<QAccessibleInterface*>(ptr)->QAccessibleInterface::window();
}

char* QAccessibleInterface_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQAccessibleInterface*>(static_cast<QAccessibleInterface*>(ptr))) {
		return static_cast<MyQAccessibleInterface*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QAccessibleInterface_BASE").toUtf8().data();
}

void QAccessibleInterface_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQAccessibleInterface*>(static_cast<QAccessibleInterface*>(ptr))) {
		static_cast<MyQAccessibleInterface*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQAccessibleObject: public QAccessibleObject
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	QAccessibleInterface * childAt(int x, int y) const { return static_cast<QAccessibleInterface*>(callbackQAccessibleObject_ChildAt(const_cast<MyQAccessibleObject*>(this), this->objectNameAbs().toUtf8().data(), x, y)); };
	bool isValid() const { return callbackQAccessibleObject_IsValid(const_cast<MyQAccessibleObject*>(this), this->objectNameAbs().toUtf8().data()) != 0; };
	QObject * object() const { return static_cast<QObject*>(callbackQAccessibleObject_Object(const_cast<MyQAccessibleObject*>(this), this->objectNameAbs().toUtf8().data())); };
	QRect rect() const { return *static_cast<QRect*>(callbackQAccessibleObject_Rect(const_cast<MyQAccessibleObject*>(this), this->objectNameAbs().toUtf8().data())); };
	void setText(QAccessible::Text t, const QString & text) { callbackQAccessibleObject_SetText(this, this->objectNameAbs().toUtf8().data(), t, text.toUtf8().data()); };
	QColor backgroundColor() const { return *static_cast<QColor*>(callbackQAccessibleObject_BackgroundColor(const_cast<MyQAccessibleObject*>(this), this->objectNameAbs().toUtf8().data())); };
	QAccessibleInterface * child(int index) const { return static_cast<QAccessibleInterface*>(callbackQAccessibleObject_Child(const_cast<MyQAccessibleObject*>(this), this->objectNameAbs().toUtf8().data(), index)); };
	int childCount() const { return callbackQAccessibleObject_ChildCount(const_cast<MyQAccessibleObject*>(this), this->objectNameAbs().toUtf8().data()); };
	QAccessibleInterface * focusChild() const { return static_cast<QAccessibleInterface*>(callbackQAccessibleObject_FocusChild(const_cast<MyQAccessibleObject*>(this), this->objectNameAbs().toUtf8().data())); };
	QColor foregroundColor() const { return *static_cast<QColor*>(callbackQAccessibleObject_ForegroundColor(const_cast<MyQAccessibleObject*>(this), this->objectNameAbs().toUtf8().data())); };
	int indexOfChild(const QAccessibleInterface * child) const { return callbackQAccessibleObject_IndexOfChild(const_cast<MyQAccessibleObject*>(this), this->objectNameAbs().toUtf8().data(), const_cast<QAccessibleInterface*>(child)); };
	void * interface_cast(QAccessible::InterfaceType ty) { return callbackQAccessibleObject_Interface_cast(this, this->objectNameAbs().toUtf8().data(), ty); };
	QAccessibleInterface * parent() const { return static_cast<QAccessibleInterface*>(callbackQAccessibleObject_Parent(const_cast<MyQAccessibleObject*>(this), this->objectNameAbs().toUtf8().data())); };
	QAccessible::Role role() const { return static_cast<QAccessible::Role>(callbackQAccessibleObject_Role(const_cast<MyQAccessibleObject*>(this), this->objectNameAbs().toUtf8().data())); };
	QString text(QAccessible::Text t) const { return QString(callbackQAccessibleObject_Text(const_cast<MyQAccessibleObject*>(this), this->objectNameAbs().toUtf8().data(), t)); };
	QWindow * window() const { return static_cast<QWindow*>(callbackQAccessibleObject_Window(const_cast<MyQAccessibleObject*>(this), this->objectNameAbs().toUtf8().data())); };
};

void* QAccessibleObject_ChildAt(void* ptr, int x, int y)
{
	return static_cast<QAccessibleObject*>(ptr)->childAt(x, y);
}

void* QAccessibleObject_ChildAtDefault(void* ptr, int x, int y)
{
	return static_cast<QAccessibleObject*>(ptr)->QAccessibleObject::childAt(x, y);
}

int QAccessibleObject_IsValid(void* ptr)
{
	return static_cast<QAccessibleObject*>(ptr)->isValid();
}

int QAccessibleObject_IsValidDefault(void* ptr)
{
	return static_cast<QAccessibleObject*>(ptr)->QAccessibleObject::isValid();
}

void* QAccessibleObject_Object(void* ptr)
{
	return static_cast<QAccessibleObject*>(ptr)->object();
}

void* QAccessibleObject_ObjectDefault(void* ptr)
{
	return static_cast<QAccessibleObject*>(ptr)->QAccessibleObject::object();
}

void* QAccessibleObject_Rect(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QAccessibleObject*>(ptr)->rect()).x(), static_cast<QRect>(static_cast<QAccessibleObject*>(ptr)->rect()).y(), static_cast<QRect>(static_cast<QAccessibleObject*>(ptr)->rect()).width(), static_cast<QRect>(static_cast<QAccessibleObject*>(ptr)->rect()).height());
}

void* QAccessibleObject_RectDefault(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QAccessibleObject*>(ptr)->QAccessibleObject::rect()).x(), static_cast<QRect>(static_cast<QAccessibleObject*>(ptr)->QAccessibleObject::rect()).y(), static_cast<QRect>(static_cast<QAccessibleObject*>(ptr)->QAccessibleObject::rect()).width(), static_cast<QRect>(static_cast<QAccessibleObject*>(ptr)->QAccessibleObject::rect()).height());
}

void QAccessibleObject_SetText(void* ptr, int t, char* text)
{
	static_cast<QAccessibleObject*>(ptr)->setText(static_cast<QAccessible::Text>(t), QString(text));
}

void QAccessibleObject_SetTextDefault(void* ptr, int t, char* text)
{
	static_cast<QAccessibleObject*>(ptr)->QAccessibleObject::setText(static_cast<QAccessible::Text>(t), QString(text));
}

void QAccessibleObject_DestroyQAccessibleObject(void* ptr)
{
	static_cast<QAccessibleObject*>(ptr)->~QAccessibleObject();
}

char* QAccessibleObject_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQAccessibleObject*>(static_cast<QAccessibleObject*>(ptr))) {
		return static_cast<MyQAccessibleObject*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QAccessibleObject_BASE").toUtf8().data();
}

void QAccessibleObject_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQAccessibleObject*>(static_cast<QAccessibleObject*>(ptr))) {
		static_cast<MyQAccessibleObject*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QAccessibleObject_BackgroundColor(void* ptr)
{
	return new QColor(static_cast<QAccessibleObject*>(ptr)->backgroundColor());
}

void* QAccessibleObject_BackgroundColorDefault(void* ptr)
{
	return new QColor(static_cast<QAccessibleObject*>(ptr)->QAccessibleObject::backgroundColor());
}

void* QAccessibleObject_Child(void* ptr, int index)
{
	return static_cast<QAccessibleObject*>(ptr)->child(index);
}



int QAccessibleObject_ChildCount(void* ptr)
{
	return static_cast<QAccessibleObject*>(ptr)->childCount();
}



void* QAccessibleObject_FocusChild(void* ptr)
{
	return static_cast<QAccessibleObject*>(ptr)->focusChild();
}

void* QAccessibleObject_FocusChildDefault(void* ptr)
{
	return static_cast<QAccessibleObject*>(ptr)->QAccessibleObject::focusChild();
}

void* QAccessibleObject_ForegroundColor(void* ptr)
{
	return new QColor(static_cast<QAccessibleObject*>(ptr)->foregroundColor());
}

void* QAccessibleObject_ForegroundColorDefault(void* ptr)
{
	return new QColor(static_cast<QAccessibleObject*>(ptr)->QAccessibleObject::foregroundColor());
}

int QAccessibleObject_IndexOfChild(void* ptr, void* child)
{
	return static_cast<QAccessibleObject*>(ptr)->indexOfChild(static_cast<QAccessibleInterface*>(child));
}



void* QAccessibleObject_Interface_cast(void* ptr, int ty)
{
	return static_cast<QAccessibleObject*>(ptr)->interface_cast(static_cast<QAccessible::InterfaceType>(ty));
}

void* QAccessibleObject_Interface_castDefault(void* ptr, int ty)
{
	return static_cast<QAccessibleObject*>(ptr)->QAccessibleObject::interface_cast(static_cast<QAccessible::InterfaceType>(ty));
}

void* QAccessibleObject_Parent(void* ptr)
{
	return static_cast<QAccessibleObject*>(ptr)->parent();
}



int QAccessibleObject_Role(void* ptr)
{
	return static_cast<QAccessibleObject*>(ptr)->role();
}



char* QAccessibleObject_Text(void* ptr, int t)
{
	return static_cast<QAccessibleObject*>(ptr)->text(static_cast<QAccessible::Text>(t)).toUtf8().data();
}



void* QAccessibleObject_Window(void* ptr)
{
	return static_cast<QAccessibleObject*>(ptr)->window();
}

void* QAccessibleObject_WindowDefault(void* ptr)
{
	return static_cast<QAccessibleObject*>(ptr)->QAccessibleObject::window();
}

class MyQAccessiblePlugin: public QAccessiblePlugin
{
public:
	MyQAccessiblePlugin(QObject *parent) : QAccessiblePlugin(parent) {};
	QAccessibleInterface * create(const QString & key, QObject * object) { return static_cast<QAccessibleInterface*>(callbackQAccessiblePlugin_Create(this, this->objectName().toUtf8().data(), key.toUtf8().data(), object)); };
	void timerEvent(QTimerEvent * event) { callbackQAccessiblePlugin_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQAccessiblePlugin_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAccessiblePlugin_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQAccessiblePlugin_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQAccessiblePlugin_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAccessiblePlugin_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQAccessiblePlugin_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAccessiblePlugin_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQAccessiblePlugin_MetaObject(const_cast<MyQAccessiblePlugin*>(this), this->objectName().toUtf8().data())); };
};

void* QAccessiblePlugin_NewQAccessiblePlugin(void* parent)
{
	return new MyQAccessiblePlugin(static_cast<QObject*>(parent));
}

void* QAccessiblePlugin_Create(void* ptr, char* key, void* object)
{
	return static_cast<QAccessiblePlugin*>(ptr)->create(QString(key), static_cast<QObject*>(object));
}

void QAccessiblePlugin_DestroyQAccessiblePlugin(void* ptr)
{
	static_cast<QAccessiblePlugin*>(ptr)->~QAccessiblePlugin();
}

void QAccessiblePlugin_TimerEvent(void* ptr, void* event)
{
	static_cast<QAccessiblePlugin*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QAccessiblePlugin_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAccessiblePlugin*>(ptr)->QAccessiblePlugin::timerEvent(static_cast<QTimerEvent*>(event));
}

void QAccessiblePlugin_ChildEvent(void* ptr, void* event)
{
	static_cast<QAccessiblePlugin*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QAccessiblePlugin_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAccessiblePlugin*>(ptr)->QAccessiblePlugin::childEvent(static_cast<QChildEvent*>(event));
}

void QAccessiblePlugin_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QAccessiblePlugin*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAccessiblePlugin_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAccessiblePlugin*>(ptr)->QAccessiblePlugin::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAccessiblePlugin_CustomEvent(void* ptr, void* event)
{
	static_cast<QAccessiblePlugin*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QAccessiblePlugin_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAccessiblePlugin*>(ptr)->QAccessiblePlugin::customEvent(static_cast<QEvent*>(event));
}

void QAccessiblePlugin_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAccessiblePlugin*>(ptr), "deleteLater");
}

void QAccessiblePlugin_DeleteLaterDefault(void* ptr)
{
	static_cast<QAccessiblePlugin*>(ptr)->QAccessiblePlugin::deleteLater();
}

void QAccessiblePlugin_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QAccessiblePlugin*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAccessiblePlugin_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAccessiblePlugin*>(ptr)->QAccessiblePlugin::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QAccessiblePlugin_Event(void* ptr, void* e)
{
	return static_cast<QAccessiblePlugin*>(ptr)->event(static_cast<QEvent*>(e));
}

int QAccessiblePlugin_EventDefault(void* ptr, void* e)
{
	return static_cast<QAccessiblePlugin*>(ptr)->QAccessiblePlugin::event(static_cast<QEvent*>(e));
}

int QAccessiblePlugin_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QAccessiblePlugin*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QAccessiblePlugin_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAccessiblePlugin*>(ptr)->QAccessiblePlugin::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QAccessiblePlugin_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAccessiblePlugin*>(ptr)->metaObject());
}

void* QAccessiblePlugin_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QAccessiblePlugin*>(ptr)->QAccessiblePlugin::metaObject());
}

void* QAccessibleStateChangeEvent_AccessibleInterface(void* ptr)
{
	return static_cast<QAccessibleStateChangeEvent*>(ptr)->accessibleInterface();
}

void* QAccessibleStateChangeEvent_AccessibleInterfaceDefault(void* ptr)
{
	return static_cast<QAccessibleStateChangeEvent*>(ptr)->QAccessibleStateChangeEvent::accessibleInterface();
}

class MyQAccessibleTableCellInterface: public QAccessibleTableCellInterface
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	int columnExtent() const { return callbackQAccessibleTableCellInterface_ColumnExtent(const_cast<MyQAccessibleTableCellInterface*>(this), this->objectNameAbs().toUtf8().data()); };
	int columnIndex() const { return callbackQAccessibleTableCellInterface_ColumnIndex(const_cast<MyQAccessibleTableCellInterface*>(this), this->objectNameAbs().toUtf8().data()); };
	bool isSelected() const { return callbackQAccessibleTableCellInterface_IsSelected(const_cast<MyQAccessibleTableCellInterface*>(this), this->objectNameAbs().toUtf8().data()) != 0; };
	int rowExtent() const { return callbackQAccessibleTableCellInterface_RowExtent(const_cast<MyQAccessibleTableCellInterface*>(this), this->objectNameAbs().toUtf8().data()); };
	int rowIndex() const { return callbackQAccessibleTableCellInterface_RowIndex(const_cast<MyQAccessibleTableCellInterface*>(this), this->objectNameAbs().toUtf8().data()); };
	QAccessibleInterface * table() const { return static_cast<QAccessibleInterface*>(callbackQAccessibleTableCellInterface_Table(const_cast<MyQAccessibleTableCellInterface*>(this), this->objectNameAbs().toUtf8().data())); };
};

int QAccessibleTableCellInterface_ColumnExtent(void* ptr)
{
	return static_cast<QAccessibleTableCellInterface*>(ptr)->columnExtent();
}

int QAccessibleTableCellInterface_ColumnIndex(void* ptr)
{
	return static_cast<QAccessibleTableCellInterface*>(ptr)->columnIndex();
}

int QAccessibleTableCellInterface_IsSelected(void* ptr)
{
	return static_cast<QAccessibleTableCellInterface*>(ptr)->isSelected();
}

int QAccessibleTableCellInterface_RowExtent(void* ptr)
{
	return static_cast<QAccessibleTableCellInterface*>(ptr)->rowExtent();
}

int QAccessibleTableCellInterface_RowIndex(void* ptr)
{
	return static_cast<QAccessibleTableCellInterface*>(ptr)->rowIndex();
}

void* QAccessibleTableCellInterface_Table(void* ptr)
{
	return static_cast<QAccessibleTableCellInterface*>(ptr)->table();
}

void QAccessibleTableCellInterface_DestroyQAccessibleTableCellInterface(void* ptr)
{
	static_cast<QAccessibleTableCellInterface*>(ptr)->~QAccessibleTableCellInterface();
}

char* QAccessibleTableCellInterface_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQAccessibleTableCellInterface*>(static_cast<QAccessibleTableCellInterface*>(ptr))) {
		return static_cast<MyQAccessibleTableCellInterface*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QAccessibleTableCellInterface_BASE").toUtf8().data();
}

void QAccessibleTableCellInterface_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQAccessibleTableCellInterface*>(static_cast<QAccessibleTableCellInterface*>(ptr))) {
		static_cast<MyQAccessibleTableCellInterface*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQAccessibleTableInterface: public QAccessibleTableInterface
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	QAccessibleInterface * caption() const { return static_cast<QAccessibleInterface*>(callbackQAccessibleTableInterface_Caption(const_cast<MyQAccessibleTableInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	QAccessibleInterface * cellAt(int row, int column) const { return static_cast<QAccessibleInterface*>(callbackQAccessibleTableInterface_CellAt(const_cast<MyQAccessibleTableInterface*>(this), this->objectNameAbs().toUtf8().data(), row, column)); };
	int columnCount() const { return callbackQAccessibleTableInterface_ColumnCount(const_cast<MyQAccessibleTableInterface*>(this), this->objectNameAbs().toUtf8().data()); };
	QString columnDescription(int column) const { return QString(callbackQAccessibleTableInterface_ColumnDescription(const_cast<MyQAccessibleTableInterface*>(this), this->objectNameAbs().toUtf8().data(), column)); };
	bool isColumnSelected(int column) const { return callbackQAccessibleTableInterface_IsColumnSelected(const_cast<MyQAccessibleTableInterface*>(this), this->objectNameAbs().toUtf8().data(), column) != 0; };
	bool isRowSelected(int row) const { return callbackQAccessibleTableInterface_IsRowSelected(const_cast<MyQAccessibleTableInterface*>(this), this->objectNameAbs().toUtf8().data(), row) != 0; };
	void modelChange(QAccessibleTableModelChangeEvent * event) { callbackQAccessibleTableInterface_ModelChange(this, this->objectNameAbs().toUtf8().data(), event); };
	int rowCount() const { return callbackQAccessibleTableInterface_RowCount(const_cast<MyQAccessibleTableInterface*>(this), this->objectNameAbs().toUtf8().data()); };
	QString rowDescription(int row) const { return QString(callbackQAccessibleTableInterface_RowDescription(const_cast<MyQAccessibleTableInterface*>(this), this->objectNameAbs().toUtf8().data(), row)); };
	bool selectColumn(int column) { return callbackQAccessibleTableInterface_SelectColumn(this, this->objectNameAbs().toUtf8().data(), column) != 0; };
	bool selectRow(int row) { return callbackQAccessibleTableInterface_SelectRow(this, this->objectNameAbs().toUtf8().data(), row) != 0; };
	int selectedCellCount() const { return callbackQAccessibleTableInterface_SelectedCellCount(const_cast<MyQAccessibleTableInterface*>(this), this->objectNameAbs().toUtf8().data()); };
	int selectedColumnCount() const { return callbackQAccessibleTableInterface_SelectedColumnCount(const_cast<MyQAccessibleTableInterface*>(this), this->objectNameAbs().toUtf8().data()); };
	int selectedRowCount() const { return callbackQAccessibleTableInterface_SelectedRowCount(const_cast<MyQAccessibleTableInterface*>(this), this->objectNameAbs().toUtf8().data()); };
	QAccessibleInterface * summary() const { return static_cast<QAccessibleInterface*>(callbackQAccessibleTableInterface_Summary(const_cast<MyQAccessibleTableInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	bool unselectColumn(int column) { return callbackQAccessibleTableInterface_UnselectColumn(this, this->objectNameAbs().toUtf8().data(), column) != 0; };
	bool unselectRow(int row) { return callbackQAccessibleTableInterface_UnselectRow(this, this->objectNameAbs().toUtf8().data(), row) != 0; };
};

void* QAccessibleTableInterface_Caption(void* ptr)
{
	return static_cast<QAccessibleTableInterface*>(ptr)->caption();
}

void* QAccessibleTableInterface_CellAt(void* ptr, int row, int column)
{
	return static_cast<QAccessibleTableInterface*>(ptr)->cellAt(row, column);
}

int QAccessibleTableInterface_ColumnCount(void* ptr)
{
	return static_cast<QAccessibleTableInterface*>(ptr)->columnCount();
}

char* QAccessibleTableInterface_ColumnDescription(void* ptr, int column)
{
	return static_cast<QAccessibleTableInterface*>(ptr)->columnDescription(column).toUtf8().data();
}

int QAccessibleTableInterface_IsColumnSelected(void* ptr, int column)
{
	return static_cast<QAccessibleTableInterface*>(ptr)->isColumnSelected(column);
}

int QAccessibleTableInterface_IsRowSelected(void* ptr, int row)
{
	return static_cast<QAccessibleTableInterface*>(ptr)->isRowSelected(row);
}

void QAccessibleTableInterface_ModelChange(void* ptr, void* event)
{
	static_cast<QAccessibleTableInterface*>(ptr)->modelChange(static_cast<QAccessibleTableModelChangeEvent*>(event));
}

int QAccessibleTableInterface_RowCount(void* ptr)
{
	return static_cast<QAccessibleTableInterface*>(ptr)->rowCount();
}

char* QAccessibleTableInterface_RowDescription(void* ptr, int row)
{
	return static_cast<QAccessibleTableInterface*>(ptr)->rowDescription(row).toUtf8().data();
}

int QAccessibleTableInterface_SelectColumn(void* ptr, int column)
{
	return static_cast<QAccessibleTableInterface*>(ptr)->selectColumn(column);
}

int QAccessibleTableInterface_SelectRow(void* ptr, int row)
{
	return static_cast<QAccessibleTableInterface*>(ptr)->selectRow(row);
}

int QAccessibleTableInterface_SelectedCellCount(void* ptr)
{
	return static_cast<QAccessibleTableInterface*>(ptr)->selectedCellCount();
}

int QAccessibleTableInterface_SelectedColumnCount(void* ptr)
{
	return static_cast<QAccessibleTableInterface*>(ptr)->selectedColumnCount();
}

int QAccessibleTableInterface_SelectedRowCount(void* ptr)
{
	return static_cast<QAccessibleTableInterface*>(ptr)->selectedRowCount();
}

void* QAccessibleTableInterface_Summary(void* ptr)
{
	return static_cast<QAccessibleTableInterface*>(ptr)->summary();
}

int QAccessibleTableInterface_UnselectColumn(void* ptr, int column)
{
	return static_cast<QAccessibleTableInterface*>(ptr)->unselectColumn(column);
}

int QAccessibleTableInterface_UnselectRow(void* ptr, int row)
{
	return static_cast<QAccessibleTableInterface*>(ptr)->unselectRow(row);
}

void QAccessibleTableInterface_DestroyQAccessibleTableInterface(void* ptr)
{
	static_cast<QAccessibleTableInterface*>(ptr)->~QAccessibleTableInterface();
}

char* QAccessibleTableInterface_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQAccessibleTableInterface*>(static_cast<QAccessibleTableInterface*>(ptr))) {
		return static_cast<MyQAccessibleTableInterface*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QAccessibleTableInterface_BASE").toUtf8().data();
}

void QAccessibleTableInterface_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQAccessibleTableInterface*>(static_cast<QAccessibleTableInterface*>(ptr))) {
		static_cast<MyQAccessibleTableInterface*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QAccessibleTableModelChangeEvent_NewQAccessibleTableModelChangeEvent2(void* iface, int changeType)
{
	return new QAccessibleTableModelChangeEvent(static_cast<QAccessibleInterface*>(iface), static_cast<QAccessibleTableModelChangeEvent::ModelChangeType>(changeType));
}

void* QAccessibleTableModelChangeEvent_NewQAccessibleTableModelChangeEvent(void* object, int changeType)
{
	return new QAccessibleTableModelChangeEvent(static_cast<QObject*>(object), static_cast<QAccessibleTableModelChangeEvent::ModelChangeType>(changeType));
}

int QAccessibleTableModelChangeEvent_FirstColumn(void* ptr)
{
	return static_cast<QAccessibleTableModelChangeEvent*>(ptr)->firstColumn();
}

int QAccessibleTableModelChangeEvent_FirstRow(void* ptr)
{
	return static_cast<QAccessibleTableModelChangeEvent*>(ptr)->firstRow();
}

int QAccessibleTableModelChangeEvent_LastColumn(void* ptr)
{
	return static_cast<QAccessibleTableModelChangeEvent*>(ptr)->lastColumn();
}

int QAccessibleTableModelChangeEvent_LastRow(void* ptr)
{
	return static_cast<QAccessibleTableModelChangeEvent*>(ptr)->lastRow();
}

int QAccessibleTableModelChangeEvent_ModelChangeType(void* ptr)
{
	return static_cast<QAccessibleTableModelChangeEvent*>(ptr)->modelChangeType();
}

void QAccessibleTableModelChangeEvent_SetFirstColumn(void* ptr, int column)
{
	static_cast<QAccessibleTableModelChangeEvent*>(ptr)->setFirstColumn(column);
}

void QAccessibleTableModelChangeEvent_SetFirstRow(void* ptr, int row)
{
	static_cast<QAccessibleTableModelChangeEvent*>(ptr)->setFirstRow(row);
}

void QAccessibleTableModelChangeEvent_SetLastColumn(void* ptr, int column)
{
	static_cast<QAccessibleTableModelChangeEvent*>(ptr)->setLastColumn(column);
}

void QAccessibleTableModelChangeEvent_SetLastRow(void* ptr, int row)
{
	static_cast<QAccessibleTableModelChangeEvent*>(ptr)->setLastRow(row);
}

void QAccessibleTableModelChangeEvent_SetModelChangeType(void* ptr, int changeType)
{
	static_cast<QAccessibleTableModelChangeEvent*>(ptr)->setModelChangeType(static_cast<QAccessibleTableModelChangeEvent::ModelChangeType>(changeType));
}

void* QAccessibleTableModelChangeEvent_AccessibleInterface(void* ptr)
{
	return static_cast<QAccessibleTableModelChangeEvent*>(ptr)->accessibleInterface();
}

void* QAccessibleTableModelChangeEvent_AccessibleInterfaceDefault(void* ptr)
{
	return static_cast<QAccessibleTableModelChangeEvent*>(ptr)->QAccessibleTableModelChangeEvent::accessibleInterface();
}

void* QAccessibleTextCursorEvent_NewQAccessibleTextCursorEvent2(void* iface, int cursorPos)
{
	return new QAccessibleTextCursorEvent(static_cast<QAccessibleInterface*>(iface), cursorPos);
}

void* QAccessibleTextCursorEvent_NewQAccessibleTextCursorEvent(void* object, int cursorPos)
{
	return new QAccessibleTextCursorEvent(static_cast<QObject*>(object), cursorPos);
}

int QAccessibleTextCursorEvent_CursorPosition(void* ptr)
{
	return static_cast<QAccessibleTextCursorEvent*>(ptr)->cursorPosition();
}

void QAccessibleTextCursorEvent_SetCursorPosition(void* ptr, int position)
{
	static_cast<QAccessibleTextCursorEvent*>(ptr)->setCursorPosition(position);
}

void* QAccessibleTextCursorEvent_AccessibleInterface(void* ptr)
{
	return static_cast<QAccessibleTextCursorEvent*>(ptr)->accessibleInterface();
}

void* QAccessibleTextCursorEvent_AccessibleInterfaceDefault(void* ptr)
{
	return static_cast<QAccessibleTextCursorEvent*>(ptr)->QAccessibleTextCursorEvent::accessibleInterface();
}

void* QAccessibleTextInsertEvent_NewQAccessibleTextInsertEvent2(void* iface, int position, char* text)
{
	return new QAccessibleTextInsertEvent(static_cast<QAccessibleInterface*>(iface), position, QString(text));
}

void* QAccessibleTextInsertEvent_NewQAccessibleTextInsertEvent(void* object, int position, char* text)
{
	return new QAccessibleTextInsertEvent(static_cast<QObject*>(object), position, QString(text));
}

int QAccessibleTextInsertEvent_ChangePosition(void* ptr)
{
	return static_cast<QAccessibleTextInsertEvent*>(ptr)->changePosition();
}

char* QAccessibleTextInsertEvent_TextInserted(void* ptr)
{
	return static_cast<QAccessibleTextInsertEvent*>(ptr)->textInserted().toUtf8().data();
}

void* QAccessibleTextInsertEvent_AccessibleInterface(void* ptr)
{
	return static_cast<QAccessibleTextInsertEvent*>(ptr)->accessibleInterface();
}

void* QAccessibleTextInsertEvent_AccessibleInterfaceDefault(void* ptr)
{
	return static_cast<QAccessibleTextInsertEvent*>(ptr)->QAccessibleTextInsertEvent::accessibleInterface();
}

class MyQAccessibleTextInterface: public QAccessibleTextInterface
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	void addSelection(int startOffset, int endOffset) { callbackQAccessibleTextInterface_AddSelection(this, this->objectNameAbs().toUtf8().data(), startOffset, endOffset); };
	QString attributes(int offset, int * startOffset, int * endOffset) const { return QString(callbackQAccessibleTextInterface_Attributes(const_cast<MyQAccessibleTextInterface*>(this), this->objectNameAbs().toUtf8().data(), offset, *startOffset, *endOffset)); };
	int characterCount() const { return callbackQAccessibleTextInterface_CharacterCount(const_cast<MyQAccessibleTextInterface*>(this), this->objectNameAbs().toUtf8().data()); };
	QRect characterRect(int offset) const { return *static_cast<QRect*>(callbackQAccessibleTextInterface_CharacterRect(const_cast<MyQAccessibleTextInterface*>(this), this->objectNameAbs().toUtf8().data(), offset)); };
	int cursorPosition() const { return callbackQAccessibleTextInterface_CursorPosition(const_cast<MyQAccessibleTextInterface*>(this), this->objectNameAbs().toUtf8().data()); };
	int offsetAtPoint(const QPoint & point) const { return callbackQAccessibleTextInterface_OffsetAtPoint(const_cast<MyQAccessibleTextInterface*>(this), this->objectNameAbs().toUtf8().data(), new QPoint(static_cast<QPoint>(point).x(), static_cast<QPoint>(point).y())); };
	void removeSelection(int selectionIndex) { callbackQAccessibleTextInterface_RemoveSelection(this, this->objectNameAbs().toUtf8().data(), selectionIndex); };
	void scrollToSubstring(int startIndex, int endIndex) { callbackQAccessibleTextInterface_ScrollToSubstring(this, this->objectNameAbs().toUtf8().data(), startIndex, endIndex); };
	void selection(int selectionIndex, int * startOffset, int * endOffset) const { callbackQAccessibleTextInterface_Selection(const_cast<MyQAccessibleTextInterface*>(this), this->objectNameAbs().toUtf8().data(), selectionIndex, *startOffset, *endOffset); };
	int selectionCount() const { return callbackQAccessibleTextInterface_SelectionCount(const_cast<MyQAccessibleTextInterface*>(this), this->objectNameAbs().toUtf8().data()); };
	void setCursorPosition(int position) { callbackQAccessibleTextInterface_SetCursorPosition(this, this->objectNameAbs().toUtf8().data(), position); };
	void setSelection(int selectionIndex, int startOffset, int endOffset) { callbackQAccessibleTextInterface_SetSelection(this, this->objectNameAbs().toUtf8().data(), selectionIndex, startOffset, endOffset); };
	QString text(int startOffset, int endOffset) const { return QString(callbackQAccessibleTextInterface_Text(const_cast<MyQAccessibleTextInterface*>(this), this->objectNameAbs().toUtf8().data(), startOffset, endOffset)); };
	QString textAfterOffset(int offset, QAccessible::TextBoundaryType boundaryType, int * startOffset, int * endOffset) const { return QString(callbackQAccessibleTextInterface_TextAfterOffset(const_cast<MyQAccessibleTextInterface*>(this), this->objectNameAbs().toUtf8().data(), offset, boundaryType, *startOffset, *endOffset)); };
	QString textAtOffset(int offset, QAccessible::TextBoundaryType boundaryType, int * startOffset, int * endOffset) const { return QString(callbackQAccessibleTextInterface_TextAtOffset(const_cast<MyQAccessibleTextInterface*>(this), this->objectNameAbs().toUtf8().data(), offset, boundaryType, *startOffset, *endOffset)); };
	QString textBeforeOffset(int offset, QAccessible::TextBoundaryType boundaryType, int * startOffset, int * endOffset) const { return QString(callbackQAccessibleTextInterface_TextBeforeOffset(const_cast<MyQAccessibleTextInterface*>(this), this->objectNameAbs().toUtf8().data(), offset, boundaryType, *startOffset, *endOffset)); };
};

void QAccessibleTextInterface_AddSelection(void* ptr, int startOffset, int endOffset)
{
	static_cast<QAccessibleTextInterface*>(ptr)->addSelection(startOffset, endOffset);
}

char* QAccessibleTextInterface_Attributes(void* ptr, int offset, int startOffset, int endOffset)
{
	return static_cast<QAccessibleTextInterface*>(ptr)->attributes(offset, &startOffset, &endOffset).toUtf8().data();
}

int QAccessibleTextInterface_CharacterCount(void* ptr)
{
	return static_cast<QAccessibleTextInterface*>(ptr)->characterCount();
}

void* QAccessibleTextInterface_CharacterRect(void* ptr, int offset)
{
	return new QRect(static_cast<QRect>(static_cast<QAccessibleTextInterface*>(ptr)->characterRect(offset)).x(), static_cast<QRect>(static_cast<QAccessibleTextInterface*>(ptr)->characterRect(offset)).y(), static_cast<QRect>(static_cast<QAccessibleTextInterface*>(ptr)->characterRect(offset)).width(), static_cast<QRect>(static_cast<QAccessibleTextInterface*>(ptr)->characterRect(offset)).height());
}

int QAccessibleTextInterface_CursorPosition(void* ptr)
{
	return static_cast<QAccessibleTextInterface*>(ptr)->cursorPosition();
}

int QAccessibleTextInterface_OffsetAtPoint(void* ptr, void* point)
{
	return static_cast<QAccessibleTextInterface*>(ptr)->offsetAtPoint(*static_cast<QPoint*>(point));
}

void QAccessibleTextInterface_RemoveSelection(void* ptr, int selectionIndex)
{
	static_cast<QAccessibleTextInterface*>(ptr)->removeSelection(selectionIndex);
}

void QAccessibleTextInterface_ScrollToSubstring(void* ptr, int startIndex, int endIndex)
{
	static_cast<QAccessibleTextInterface*>(ptr)->scrollToSubstring(startIndex, endIndex);
}

void QAccessibleTextInterface_Selection(void* ptr, int selectionIndex, int startOffset, int endOffset)
{
	static_cast<QAccessibleTextInterface*>(ptr)->selection(selectionIndex, &startOffset, &endOffset);
}

int QAccessibleTextInterface_SelectionCount(void* ptr)
{
	return static_cast<QAccessibleTextInterface*>(ptr)->selectionCount();
}

void QAccessibleTextInterface_SetCursorPosition(void* ptr, int position)
{
	static_cast<QAccessibleTextInterface*>(ptr)->setCursorPosition(position);
}

void QAccessibleTextInterface_SetSelection(void* ptr, int selectionIndex, int startOffset, int endOffset)
{
	static_cast<QAccessibleTextInterface*>(ptr)->setSelection(selectionIndex, startOffset, endOffset);
}

char* QAccessibleTextInterface_Text(void* ptr, int startOffset, int endOffset)
{
	return static_cast<QAccessibleTextInterface*>(ptr)->text(startOffset, endOffset).toUtf8().data();
}

char* QAccessibleTextInterface_TextAfterOffset(void* ptr, int offset, int boundaryType, int startOffset, int endOffset)
{
	return static_cast<QAccessibleTextInterface*>(ptr)->textAfterOffset(offset, static_cast<QAccessible::TextBoundaryType>(boundaryType), &startOffset, &endOffset).toUtf8().data();
}

char* QAccessibleTextInterface_TextAfterOffsetDefault(void* ptr, int offset, int boundaryType, int startOffset, int endOffset)
{
	return static_cast<QAccessibleTextInterface*>(ptr)->QAccessibleTextInterface::textAfterOffset(offset, static_cast<QAccessible::TextBoundaryType>(boundaryType), &startOffset, &endOffset).toUtf8().data();
}

char* QAccessibleTextInterface_TextAtOffset(void* ptr, int offset, int boundaryType, int startOffset, int endOffset)
{
	return static_cast<QAccessibleTextInterface*>(ptr)->textAtOffset(offset, static_cast<QAccessible::TextBoundaryType>(boundaryType), &startOffset, &endOffset).toUtf8().data();
}

char* QAccessibleTextInterface_TextAtOffsetDefault(void* ptr, int offset, int boundaryType, int startOffset, int endOffset)
{
	return static_cast<QAccessibleTextInterface*>(ptr)->QAccessibleTextInterface::textAtOffset(offset, static_cast<QAccessible::TextBoundaryType>(boundaryType), &startOffset, &endOffset).toUtf8().data();
}

char* QAccessibleTextInterface_TextBeforeOffset(void* ptr, int offset, int boundaryType, int startOffset, int endOffset)
{
	return static_cast<QAccessibleTextInterface*>(ptr)->textBeforeOffset(offset, static_cast<QAccessible::TextBoundaryType>(boundaryType), &startOffset, &endOffset).toUtf8().data();
}

char* QAccessibleTextInterface_TextBeforeOffsetDefault(void* ptr, int offset, int boundaryType, int startOffset, int endOffset)
{
	return static_cast<QAccessibleTextInterface*>(ptr)->QAccessibleTextInterface::textBeforeOffset(offset, static_cast<QAccessible::TextBoundaryType>(boundaryType), &startOffset, &endOffset).toUtf8().data();
}

void QAccessibleTextInterface_DestroyQAccessibleTextInterface(void* ptr)
{
	static_cast<QAccessibleTextInterface*>(ptr)->~QAccessibleTextInterface();
}

char* QAccessibleTextInterface_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQAccessibleTextInterface*>(static_cast<QAccessibleTextInterface*>(ptr))) {
		return static_cast<MyQAccessibleTextInterface*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QAccessibleTextInterface_BASE").toUtf8().data();
}

void QAccessibleTextInterface_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQAccessibleTextInterface*>(static_cast<QAccessibleTextInterface*>(ptr))) {
		static_cast<MyQAccessibleTextInterface*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QAccessibleTextRemoveEvent_NewQAccessibleTextRemoveEvent2(void* iface, int position, char* text)
{
	return new QAccessibleTextRemoveEvent(static_cast<QAccessibleInterface*>(iface), position, QString(text));
}

void* QAccessibleTextRemoveEvent_NewQAccessibleTextRemoveEvent(void* object, int position, char* text)
{
	return new QAccessibleTextRemoveEvent(static_cast<QObject*>(object), position, QString(text));
}

int QAccessibleTextRemoveEvent_ChangePosition(void* ptr)
{
	return static_cast<QAccessibleTextRemoveEvent*>(ptr)->changePosition();
}

char* QAccessibleTextRemoveEvent_TextRemoved(void* ptr)
{
	return static_cast<QAccessibleTextRemoveEvent*>(ptr)->textRemoved().toUtf8().data();
}

void* QAccessibleTextRemoveEvent_AccessibleInterface(void* ptr)
{
	return static_cast<QAccessibleTextRemoveEvent*>(ptr)->accessibleInterface();
}

void* QAccessibleTextRemoveEvent_AccessibleInterfaceDefault(void* ptr)
{
	return static_cast<QAccessibleTextRemoveEvent*>(ptr)->QAccessibleTextRemoveEvent::accessibleInterface();
}

void* QAccessibleTextSelectionEvent_NewQAccessibleTextSelectionEvent2(void* iface, int start, int end)
{
	return new QAccessibleTextSelectionEvent(static_cast<QAccessibleInterface*>(iface), start, end);
}

void* QAccessibleTextSelectionEvent_NewQAccessibleTextSelectionEvent(void* object, int start, int end)
{
	return new QAccessibleTextSelectionEvent(static_cast<QObject*>(object), start, end);
}

int QAccessibleTextSelectionEvent_SelectionEnd(void* ptr)
{
	return static_cast<QAccessibleTextSelectionEvent*>(ptr)->selectionEnd();
}

int QAccessibleTextSelectionEvent_SelectionStart(void* ptr)
{
	return static_cast<QAccessibleTextSelectionEvent*>(ptr)->selectionStart();
}

void QAccessibleTextSelectionEvent_SetSelection(void* ptr, int start, int end)
{
	static_cast<QAccessibleTextSelectionEvent*>(ptr)->setSelection(start, end);
}

void* QAccessibleTextSelectionEvent_AccessibleInterface(void* ptr)
{
	return static_cast<QAccessibleTextSelectionEvent*>(ptr)->accessibleInterface();
}

void* QAccessibleTextSelectionEvent_AccessibleInterfaceDefault(void* ptr)
{
	return static_cast<QAccessibleTextSelectionEvent*>(ptr)->QAccessibleTextSelectionEvent::accessibleInterface();
}

void* QAccessibleTextUpdateEvent_NewQAccessibleTextUpdateEvent2(void* iface, int position, char* oldText, char* text)
{
	return new QAccessibleTextUpdateEvent(static_cast<QAccessibleInterface*>(iface), position, QString(oldText), QString(text));
}

void* QAccessibleTextUpdateEvent_NewQAccessibleTextUpdateEvent(void* object, int position, char* oldText, char* text)
{
	return new QAccessibleTextUpdateEvent(static_cast<QObject*>(object), position, QString(oldText), QString(text));
}

int QAccessibleTextUpdateEvent_ChangePosition(void* ptr)
{
	return static_cast<QAccessibleTextUpdateEvent*>(ptr)->changePosition();
}

char* QAccessibleTextUpdateEvent_TextInserted(void* ptr)
{
	return static_cast<QAccessibleTextUpdateEvent*>(ptr)->textInserted().toUtf8().data();
}

char* QAccessibleTextUpdateEvent_TextRemoved(void* ptr)
{
	return static_cast<QAccessibleTextUpdateEvent*>(ptr)->textRemoved().toUtf8().data();
}

void* QAccessibleTextUpdateEvent_AccessibleInterface(void* ptr)
{
	return static_cast<QAccessibleTextUpdateEvent*>(ptr)->accessibleInterface();
}

void* QAccessibleTextUpdateEvent_AccessibleInterfaceDefault(void* ptr)
{
	return static_cast<QAccessibleTextUpdateEvent*>(ptr)->QAccessibleTextUpdateEvent::accessibleInterface();
}

void* QAccessibleValueChangeEvent_NewQAccessibleValueChangeEvent2(void* iface, void* val)
{
	return new QAccessibleValueChangeEvent(static_cast<QAccessibleInterface*>(iface), *static_cast<QVariant*>(val));
}

void* QAccessibleValueChangeEvent_NewQAccessibleValueChangeEvent(void* object, void* value)
{
	return new QAccessibleValueChangeEvent(static_cast<QObject*>(object), *static_cast<QVariant*>(value));
}

void QAccessibleValueChangeEvent_SetValue(void* ptr, void* value)
{
	static_cast<QAccessibleValueChangeEvent*>(ptr)->setValue(*static_cast<QVariant*>(value));
}

void* QAccessibleValueChangeEvent_Value(void* ptr)
{
	return new QVariant(static_cast<QAccessibleValueChangeEvent*>(ptr)->value());
}

void* QAccessibleValueChangeEvent_AccessibleInterface(void* ptr)
{
	return static_cast<QAccessibleValueChangeEvent*>(ptr)->accessibleInterface();
}

void* QAccessibleValueChangeEvent_AccessibleInterfaceDefault(void* ptr)
{
	return static_cast<QAccessibleValueChangeEvent*>(ptr)->QAccessibleValueChangeEvent::accessibleInterface();
}

class MyQAccessibleValueInterface: public QAccessibleValueInterface
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	QVariant currentValue() const { return *static_cast<QVariant*>(callbackQAccessibleValueInterface_CurrentValue(const_cast<MyQAccessibleValueInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	QVariant maximumValue() const { return *static_cast<QVariant*>(callbackQAccessibleValueInterface_MaximumValue(const_cast<MyQAccessibleValueInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	QVariant minimumStepSize() const { return *static_cast<QVariant*>(callbackQAccessibleValueInterface_MinimumStepSize(const_cast<MyQAccessibleValueInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	QVariant minimumValue() const { return *static_cast<QVariant*>(callbackQAccessibleValueInterface_MinimumValue(const_cast<MyQAccessibleValueInterface*>(this), this->objectNameAbs().toUtf8().data())); };
	void setCurrentValue(const QVariant & value) { callbackQAccessibleValueInterface_SetCurrentValue(this, this->objectNameAbs().toUtf8().data(), new QVariant(value)); };
};

void* QAccessibleValueInterface_CurrentValue(void* ptr)
{
	return new QVariant(static_cast<QAccessibleValueInterface*>(ptr)->currentValue());
}

void* QAccessibleValueInterface_MaximumValue(void* ptr)
{
	return new QVariant(static_cast<QAccessibleValueInterface*>(ptr)->maximumValue());
}

void* QAccessibleValueInterface_MinimumStepSize(void* ptr)
{
	return new QVariant(static_cast<QAccessibleValueInterface*>(ptr)->minimumStepSize());
}

void* QAccessibleValueInterface_MinimumValue(void* ptr)
{
	return new QVariant(static_cast<QAccessibleValueInterface*>(ptr)->minimumValue());
}

void QAccessibleValueInterface_SetCurrentValue(void* ptr, void* value)
{
	static_cast<QAccessibleValueInterface*>(ptr)->setCurrentValue(*static_cast<QVariant*>(value));
}

void QAccessibleValueInterface_DestroyQAccessibleValueInterface(void* ptr)
{
	static_cast<QAccessibleValueInterface*>(ptr)->~QAccessibleValueInterface();
}

char* QAccessibleValueInterface_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQAccessibleValueInterface*>(static_cast<QAccessibleValueInterface*>(ptr))) {
		return static_cast<MyQAccessibleValueInterface*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QAccessibleValueInterface_BASE").toUtf8().data();
}

void QAccessibleValueInterface_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQAccessibleValueInterface*>(static_cast<QAccessibleValueInterface*>(ptr))) {
		static_cast<MyQAccessibleValueInterface*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QActionEvent_NewQActionEvent(int ty, void* action, void* before)
{
	return new QActionEvent(ty, static_cast<QAction*>(action), static_cast<QAction*>(before));
}

void* QActionEvent_Action(void* ptr)
{
	return static_cast<QActionEvent*>(ptr)->action();
}

void* QActionEvent_Before(void* ptr)
{
	return static_cast<QActionEvent*>(ptr)->before();
}

void* QBackingStore_PaintDevice(void* ptr)
{
	return static_cast<QBackingStore*>(ptr)->paintDevice();
}

void* QBackingStore_NewQBackingStore(void* window)
{
	return new QBackingStore(static_cast<QWindow*>(window));
}

void QBackingStore_BeginPaint(void* ptr, void* region)
{
	static_cast<QBackingStore*>(ptr)->beginPaint(*static_cast<QRegion*>(region));
}

void QBackingStore_EndPaint(void* ptr)
{
	static_cast<QBackingStore*>(ptr)->endPaint();
}

void QBackingStore_Flush(void* ptr, void* region, void* win, void* offset)
{
	static_cast<QBackingStore*>(ptr)->flush(*static_cast<QRegion*>(region), static_cast<QWindow*>(win), *static_cast<QPoint*>(offset));
}

int QBackingStore_HasStaticContents(void* ptr)
{
	return static_cast<QBackingStore*>(ptr)->hasStaticContents();
}

void QBackingStore_Resize(void* ptr, void* size)
{
	static_cast<QBackingStore*>(ptr)->resize(*static_cast<QSize*>(size));
}

int QBackingStore_Scroll(void* ptr, void* area, int dx, int dy)
{
	return static_cast<QBackingStore*>(ptr)->scroll(*static_cast<QRegion*>(area), dx, dy);
}

void QBackingStore_SetStaticContents(void* ptr, void* region)
{
	static_cast<QBackingStore*>(ptr)->setStaticContents(*static_cast<QRegion*>(region));
}

void* QBackingStore_Size(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QBackingStore*>(ptr)->size()).width(), static_cast<QSize>(static_cast<QBackingStore*>(ptr)->size()).height());
}

void* QBackingStore_StaticContents(void* ptr)
{
	return new QRegion(static_cast<QBackingStore*>(ptr)->staticContents());
}

void* QBackingStore_Window(void* ptr)
{
	return static_cast<QBackingStore*>(ptr)->window();
}

void QBackingStore_DestroyQBackingStore(void* ptr)
{
	static_cast<QBackingStore*>(ptr)->~QBackingStore();
}

void* QBitmap_NewQBitmap3(int width, int height)
{
	return new QBitmap(width, height);
}

void* QBitmap_NewQBitmap()
{
	return new QBitmap();
}

void* QBitmap_NewQBitmap2(void* pixmap)
{
	return new QBitmap(*static_cast<QPixmap*>(pixmap));
}

void* QBitmap_NewQBitmap4(void* size)
{
	return new QBitmap(*static_cast<QSize*>(size));
}

void* QBitmap_NewQBitmap5(char* fileName, char* format)
{
	return new QBitmap(QString(fileName), const_cast<const char*>(format));
}

void QBitmap_Clear(void* ptr)
{
	static_cast<QBitmap*>(ptr)->clear();
}

void QBitmap_Swap(void* ptr, void* other)
{
	static_cast<QBitmap*>(ptr)->swap(*static_cast<QBitmap*>(other));
}

void QBitmap_DestroyQBitmap(void* ptr)
{
	static_cast<QBitmap*>(ptr)->~QBitmap();
}

int QBitmap_Metric(void* ptr, int metric)
{
	return static_cast<QBitmap*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

int QBitmap_MetricDefault(void* ptr, int metric)
{
	return static_cast<QBitmap*>(ptr)->QBitmap::metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

void* QBitmap_PaintEngine(void* ptr)
{
	return static_cast<QBitmap*>(ptr)->paintEngine();
}

void* QBitmap_PaintEngineDefault(void* ptr)
{
	return static_cast<QBitmap*>(ptr)->QBitmap::paintEngine();
}

void* QBrush_NewQBrush4(int color, int style)
{
	return new QBrush(static_cast<Qt::GlobalColor>(color), static_cast<Qt::BrushStyle>(style));
}

void QBrush_SetColor(void* ptr, void* color)
{
	static_cast<QBrush*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void* QBrush_Texture(void* ptr)
{
	return new QPixmap(static_cast<QBrush*>(ptr)->texture());
}

void* QBrush_NewQBrush()
{
	return new QBrush();
}

void* QBrush_NewQBrush2(int style)
{
	return new QBrush(static_cast<Qt::BrushStyle>(style));
}

void* QBrush_NewQBrush6(int color, void* pixmap)
{
	return new QBrush(static_cast<Qt::GlobalColor>(color), *static_cast<QPixmap*>(pixmap));
}

void* QBrush_NewQBrush9(void* other)
{
	return new QBrush(*static_cast<QBrush*>(other));
}

void* QBrush_NewQBrush3(void* color, int style)
{
	return new QBrush(*static_cast<QColor*>(color), static_cast<Qt::BrushStyle>(style));
}

void* QBrush_NewQBrush5(void* color, void* pixmap)
{
	return new QBrush(*static_cast<QColor*>(color), *static_cast<QPixmap*>(pixmap));
}

void* QBrush_NewQBrush10(void* gradient)
{
	return new QBrush(*static_cast<QGradient*>(gradient));
}

void* QBrush_NewQBrush8(void* image)
{
	return new QBrush(*static_cast<QImage*>(image));
}

void* QBrush_NewQBrush7(void* pixmap)
{
	return new QBrush(*static_cast<QPixmap*>(pixmap));
}

void* QBrush_Color(void* ptr)
{
	return new QColor(static_cast<QBrush*>(ptr)->color());
}

void* QBrush_Gradient(void* ptr)
{
	return const_cast<QGradient*>(static_cast<QBrush*>(ptr)->gradient());
}

int QBrush_IsOpaque(void* ptr)
{
	return static_cast<QBrush*>(ptr)->isOpaque();
}

void QBrush_SetColor2(void* ptr, int color)
{
	static_cast<QBrush*>(ptr)->setColor(static_cast<Qt::GlobalColor>(color));
}

void QBrush_SetStyle(void* ptr, int style)
{
	static_cast<QBrush*>(ptr)->setStyle(static_cast<Qt::BrushStyle>(style));
}

void QBrush_SetTexture(void* ptr, void* pixmap)
{
	static_cast<QBrush*>(ptr)->setTexture(*static_cast<QPixmap*>(pixmap));
}

void QBrush_SetTextureImage(void* ptr, void* image)
{
	static_cast<QBrush*>(ptr)->setTextureImage(*static_cast<QImage*>(image));
}

void QBrush_SetTransform(void* ptr, void* matrix)
{
	static_cast<QBrush*>(ptr)->setTransform(*static_cast<QTransform*>(matrix));
}

int QBrush_Style(void* ptr)
{
	return static_cast<QBrush*>(ptr)->style();
}

void QBrush_Swap(void* ptr, void* other)
{
	static_cast<QBrush*>(ptr)->swap(*static_cast<QBrush*>(other));
}

void* QBrush_TextureImage(void* ptr)
{
	return new QImage(static_cast<QBrush*>(ptr)->textureImage());
}

void QBrush_DestroyQBrush(void* ptr)
{
	static_cast<QBrush*>(ptr)->~QBrush();
}

class MyQClipboard: public QClipboard
{
public:
	void Signal_Changed(QClipboard::Mode mode) { callbackQClipboard_Changed(this, this->objectName().toUtf8().data(), mode); };
	void Signal_DataChanged() { callbackQClipboard_DataChanged(this, this->objectName().toUtf8().data()); };
	void Signal_FindBufferChanged() { callbackQClipboard_FindBufferChanged(this, this->objectName().toUtf8().data()); };
	void Signal_SelectionChanged() { callbackQClipboard_SelectionChanged(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQClipboard_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQClipboard_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQClipboard_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQClipboard_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQClipboard_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQClipboard_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQClipboard_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQClipboard_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQClipboard_MetaObject(const_cast<MyQClipboard*>(this), this->objectName().toUtf8().data())); };
};

void QClipboard_Clear(void* ptr, int mode)
{
	static_cast<QClipboard*>(ptr)->clear(static_cast<QClipboard::Mode>(mode));
}

void* QClipboard_MimeData(void* ptr, int mode)
{
	return const_cast<QMimeData*>(static_cast<QClipboard*>(ptr)->mimeData(static_cast<QClipboard::Mode>(mode)));
}

void QClipboard_SetMimeData(void* ptr, void* src, int mode)
{
	static_cast<QClipboard*>(ptr)->setMimeData(static_cast<QMimeData*>(src), static_cast<QClipboard::Mode>(mode));
}

void QClipboard_ConnectChanged(void* ptr)
{
	QObject::connect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)(QClipboard::Mode)>(&QClipboard::changed), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)(QClipboard::Mode)>(&MyQClipboard::Signal_Changed));
}

void QClipboard_DisconnectChanged(void* ptr)
{
	QObject::disconnect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)(QClipboard::Mode)>(&QClipboard::changed), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)(QClipboard::Mode)>(&MyQClipboard::Signal_Changed));
}

void QClipboard_Changed(void* ptr, int mode)
{
	static_cast<QClipboard*>(ptr)->changed(static_cast<QClipboard::Mode>(mode));
}

void QClipboard_ConnectDataChanged(void* ptr)
{
	QObject::connect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::dataChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_DataChanged));
}

void QClipboard_DisconnectDataChanged(void* ptr)
{
	QObject::disconnect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::dataChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_DataChanged));
}

void QClipboard_DataChanged(void* ptr)
{
	static_cast<QClipboard*>(ptr)->dataChanged();
}

void QClipboard_ConnectFindBufferChanged(void* ptr)
{
	QObject::connect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::findBufferChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_FindBufferChanged));
}

void QClipboard_DisconnectFindBufferChanged(void* ptr)
{
	QObject::disconnect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::findBufferChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_FindBufferChanged));
}

void QClipboard_FindBufferChanged(void* ptr)
{
	static_cast<QClipboard*>(ptr)->findBufferChanged();
}

void* QClipboard_Image(void* ptr, int mode)
{
	return new QImage(static_cast<QClipboard*>(ptr)->image(static_cast<QClipboard::Mode>(mode)));
}

int QClipboard_OwnsClipboard(void* ptr)
{
	return static_cast<QClipboard*>(ptr)->ownsClipboard();
}

int QClipboard_OwnsFindBuffer(void* ptr)
{
	return static_cast<QClipboard*>(ptr)->ownsFindBuffer();
}

int QClipboard_OwnsSelection(void* ptr)
{
	return static_cast<QClipboard*>(ptr)->ownsSelection();
}

void* QClipboard_Pixmap(void* ptr, int mode)
{
	return new QPixmap(static_cast<QClipboard*>(ptr)->pixmap(static_cast<QClipboard::Mode>(mode)));
}

void QClipboard_ConnectSelectionChanged(void* ptr)
{
	QObject::connect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::selectionChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_SelectionChanged));
}

void QClipboard_DisconnectSelectionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QClipboard*>(ptr), static_cast<void (QClipboard::*)()>(&QClipboard::selectionChanged), static_cast<MyQClipboard*>(ptr), static_cast<void (MyQClipboard::*)()>(&MyQClipboard::Signal_SelectionChanged));
}

void QClipboard_SelectionChanged(void* ptr)
{
	static_cast<QClipboard*>(ptr)->selectionChanged();
}

void QClipboard_SetImage(void* ptr, void* image, int mode)
{
	static_cast<QClipboard*>(ptr)->setImage(*static_cast<QImage*>(image), static_cast<QClipboard::Mode>(mode));
}

void QClipboard_SetPixmap(void* ptr, void* pixmap, int mode)
{
	static_cast<QClipboard*>(ptr)->setPixmap(*static_cast<QPixmap*>(pixmap), static_cast<QClipboard::Mode>(mode));
}

void QClipboard_SetText(void* ptr, char* text, int mode)
{
	static_cast<QClipboard*>(ptr)->setText(QString(text), static_cast<QClipboard::Mode>(mode));
}

int QClipboard_SupportsFindBuffer(void* ptr)
{
	return static_cast<QClipboard*>(ptr)->supportsFindBuffer();
}

int QClipboard_SupportsSelection(void* ptr)
{
	return static_cast<QClipboard*>(ptr)->supportsSelection();
}

char* QClipboard_Text(void* ptr, int mode)
{
	return static_cast<QClipboard*>(ptr)->text(static_cast<QClipboard::Mode>(mode)).toUtf8().data();
}

void QClipboard_TimerEvent(void* ptr, void* event)
{
	static_cast<QClipboard*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QClipboard_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QClipboard*>(ptr)->QClipboard::timerEvent(static_cast<QTimerEvent*>(event));
}

void QClipboard_ChildEvent(void* ptr, void* event)
{
	static_cast<QClipboard*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QClipboard_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QClipboard*>(ptr)->QClipboard::childEvent(static_cast<QChildEvent*>(event));
}

void QClipboard_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QClipboard*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QClipboard_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QClipboard*>(ptr)->QClipboard::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QClipboard_CustomEvent(void* ptr, void* event)
{
	static_cast<QClipboard*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QClipboard_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QClipboard*>(ptr)->QClipboard::customEvent(static_cast<QEvent*>(event));
}

void QClipboard_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QClipboard*>(ptr), "deleteLater");
}

void QClipboard_DeleteLaterDefault(void* ptr)
{
	static_cast<QClipboard*>(ptr)->QClipboard::deleteLater();
}

void QClipboard_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QClipboard*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QClipboard_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QClipboard*>(ptr)->QClipboard::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QClipboard_Event(void* ptr, void* e)
{
	return static_cast<QClipboard*>(ptr)->event(static_cast<QEvent*>(e));
}

int QClipboard_EventDefault(void* ptr, void* e)
{
	return static_cast<QClipboard*>(ptr)->QClipboard::event(static_cast<QEvent*>(e));
}

int QClipboard_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QClipboard*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QClipboard_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QClipboard*>(ptr)->QClipboard::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QClipboard_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QClipboard*>(ptr)->metaObject());
}

void* QClipboard_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QClipboard*>(ptr)->QClipboard::metaObject());
}

void* QCloseEvent_NewQCloseEvent()
{
	return new QCloseEvent();
}

void* QColor_ConvertTo(void* ptr, int colorSpec)
{
	return new QColor(static_cast<QColor*>(ptr)->convertTo(static_cast<QColor::Spec>(colorSpec)));
}

void QColor_SetRgbF(void* ptr, double r, double g, double b, double a)
{
	static_cast<QColor*>(ptr)->setRgbF(static_cast<double>(r), static_cast<double>(g), static_cast<double>(b), static_cast<double>(a));
}

void* QColor_NewQColor()
{
	return new QColor();
}

void* QColor_NewQColor8(int color)
{
	return new QColor(static_cast<Qt::GlobalColor>(color));
}

void* QColor_NewQColor6(void* color)
{
	return new QColor(*static_cast<QColor*>(color));
}

void* QColor_NewQColor4(char* name)
{
	return new QColor(QString(name));
}

void* QColor_NewQColor5(char* name)
{
	return new QColor(const_cast<const char*>(name));
}

void* QColor_NewQColor2(int r, int g, int b, int a)
{
	return new QColor(r, g, b, a);
}

int QColor_Alpha(void* ptr)
{
	return static_cast<QColor*>(ptr)->alpha();
}

double QColor_AlphaF(void* ptr)
{
	return static_cast<double>(static_cast<QColor*>(ptr)->alphaF());
}

int QColor_Black(void* ptr)
{
	return static_cast<QColor*>(ptr)->black();
}

double QColor_BlackF(void* ptr)
{
	return static_cast<double>(static_cast<QColor*>(ptr)->blackF());
}

int QColor_Blue(void* ptr)
{
	return static_cast<QColor*>(ptr)->blue();
}

double QColor_BlueF(void* ptr)
{
	return static_cast<double>(static_cast<QColor*>(ptr)->blueF());
}

char* QColor_QColor_ColorNames()
{
	return QColor::colorNames().join("|").toUtf8().data();
}

int QColor_Cyan(void* ptr)
{
	return static_cast<QColor*>(ptr)->cyan();
}

double QColor_CyanF(void* ptr)
{
	return static_cast<double>(static_cast<QColor*>(ptr)->cyanF());
}

void* QColor_Darker(void* ptr, int factor)
{
	return new QColor(static_cast<QColor*>(ptr)->darker(factor));
}

void* QColor_QColor_FromCmyk(int c, int m, int y, int k, int a)
{
	return new QColor(QColor::fromCmyk(c, m, y, k, a));
}

void* QColor_QColor_FromCmykF(double c, double m, double y, double k, double a)
{
	return new QColor(QColor::fromCmykF(static_cast<double>(c), static_cast<double>(m), static_cast<double>(y), static_cast<double>(k), static_cast<double>(a)));
}

void* QColor_QColor_FromHsl(int h, int s, int l, int a)
{
	return new QColor(QColor::fromHsl(h, s, l, a));
}

void* QColor_QColor_FromHslF(double h, double s, double l, double a)
{
	return new QColor(QColor::fromHslF(static_cast<double>(h), static_cast<double>(s), static_cast<double>(l), static_cast<double>(a)));
}

void* QColor_QColor_FromHsv(int h, int s, int v, int a)
{
	return new QColor(QColor::fromHsv(h, s, v, a));
}

void* QColor_QColor_FromHsvF(double h, double s, double v, double a)
{
	return new QColor(QColor::fromHsvF(static_cast<double>(h), static_cast<double>(s), static_cast<double>(v), static_cast<double>(a)));
}

void* QColor_QColor_FromRgb2(int r, int g, int b, int a)
{
	return new QColor(QColor::fromRgb(r, g, b, a));
}

void* QColor_QColor_FromRgbF(double r, double g, double b, double a)
{
	return new QColor(QColor::fromRgbF(static_cast<double>(r), static_cast<double>(g), static_cast<double>(b), static_cast<double>(a)));
}

void QColor_GetCmyk(void* ptr, int c, int m, int y, int k, int a)
{
	static_cast<QColor*>(ptr)->getCmyk(&c, &m, &y, &k, &a);
}

void QColor_GetHsl(void* ptr, int h, int s, int l, int a)
{
	static_cast<QColor*>(ptr)->getHsl(&h, &s, &l, &a);
}

void QColor_GetHsv(void* ptr, int h, int s, int v, int a)
{
	static_cast<QColor*>(ptr)->getHsv(&h, &s, &v, &a);
}

void QColor_GetRgb(void* ptr, int r, int g, int b, int a)
{
	static_cast<QColor*>(ptr)->getRgb(&r, &g, &b, &a);
}

int QColor_Green(void* ptr)
{
	return static_cast<QColor*>(ptr)->green();
}

double QColor_GreenF(void* ptr)
{
	return static_cast<double>(static_cast<QColor*>(ptr)->greenF());
}

int QColor_HslHue(void* ptr)
{
	return static_cast<QColor*>(ptr)->hslHue();
}

double QColor_HslHueF(void* ptr)
{
	return static_cast<double>(static_cast<QColor*>(ptr)->hslHueF());
}

int QColor_HslSaturation(void* ptr)
{
	return static_cast<QColor*>(ptr)->hslSaturation();
}

double QColor_HslSaturationF(void* ptr)
{
	return static_cast<double>(static_cast<QColor*>(ptr)->hslSaturationF());
}

int QColor_HsvHue(void* ptr)
{
	return static_cast<QColor*>(ptr)->hsvHue();
}

double QColor_HsvHueF(void* ptr)
{
	return static_cast<double>(static_cast<QColor*>(ptr)->hsvHueF());
}

int QColor_HsvSaturation(void* ptr)
{
	return static_cast<QColor*>(ptr)->hsvSaturation();
}

double QColor_HsvSaturationF(void* ptr)
{
	return static_cast<double>(static_cast<QColor*>(ptr)->hsvSaturationF());
}

int QColor_Hue(void* ptr)
{
	return static_cast<QColor*>(ptr)->hue();
}

double QColor_HueF(void* ptr)
{
	return static_cast<double>(static_cast<QColor*>(ptr)->hueF());
}

int QColor_IsValid(void* ptr)
{
	return static_cast<QColor*>(ptr)->isValid();
}

int QColor_QColor_IsValidColor(char* name)
{
	return QColor::isValidColor(QString(name));
}

void* QColor_Lighter(void* ptr, int factor)
{
	return new QColor(static_cast<QColor*>(ptr)->lighter(factor));
}

int QColor_Lightness(void* ptr)
{
	return static_cast<QColor*>(ptr)->lightness();
}

double QColor_LightnessF(void* ptr)
{
	return static_cast<double>(static_cast<QColor*>(ptr)->lightnessF());
}

int QColor_Magenta(void* ptr)
{
	return static_cast<QColor*>(ptr)->magenta();
}

double QColor_MagentaF(void* ptr)
{
	return static_cast<double>(static_cast<QColor*>(ptr)->magentaF());
}

char* QColor_Name(void* ptr)
{
	return static_cast<QColor*>(ptr)->name().toUtf8().data();
}

char* QColor_Name2(void* ptr, int format)
{
	return static_cast<QColor*>(ptr)->name(static_cast<QColor::NameFormat>(format)).toUtf8().data();
}

int QColor_Red(void* ptr)
{
	return static_cast<QColor*>(ptr)->red();
}

double QColor_RedF(void* ptr)
{
	return static_cast<double>(static_cast<QColor*>(ptr)->redF());
}

int QColor_Saturation(void* ptr)
{
	return static_cast<QColor*>(ptr)->saturation();
}

double QColor_SaturationF(void* ptr)
{
	return static_cast<double>(static_cast<QColor*>(ptr)->saturationF());
}

void QColor_SetAlpha(void* ptr, int alpha)
{
	static_cast<QColor*>(ptr)->setAlpha(alpha);
}

void QColor_SetAlphaF(void* ptr, double alpha)
{
	static_cast<QColor*>(ptr)->setAlphaF(static_cast<double>(alpha));
}

void QColor_SetBlue(void* ptr, int blue)
{
	static_cast<QColor*>(ptr)->setBlue(blue);
}

void QColor_SetBlueF(void* ptr, double blue)
{
	static_cast<QColor*>(ptr)->setBlueF(static_cast<double>(blue));
}

void QColor_SetCmyk(void* ptr, int c, int m, int y, int k, int a)
{
	static_cast<QColor*>(ptr)->setCmyk(c, m, y, k, a);
}

void QColor_SetCmykF(void* ptr, double c, double m, double y, double k, double a)
{
	static_cast<QColor*>(ptr)->setCmykF(static_cast<double>(c), static_cast<double>(m), static_cast<double>(y), static_cast<double>(k), static_cast<double>(a));
}

void QColor_SetGreen(void* ptr, int green)
{
	static_cast<QColor*>(ptr)->setGreen(green);
}

void QColor_SetGreenF(void* ptr, double green)
{
	static_cast<QColor*>(ptr)->setGreenF(static_cast<double>(green));
}

void QColor_SetHsl(void* ptr, int h, int s, int l, int a)
{
	static_cast<QColor*>(ptr)->setHsl(h, s, l, a);
}

void QColor_SetHslF(void* ptr, double h, double s, double l, double a)
{
	static_cast<QColor*>(ptr)->setHslF(static_cast<double>(h), static_cast<double>(s), static_cast<double>(l), static_cast<double>(a));
}

void QColor_SetHsv(void* ptr, int h, int s, int v, int a)
{
	static_cast<QColor*>(ptr)->setHsv(h, s, v, a);
}

void QColor_SetHsvF(void* ptr, double h, double s, double v, double a)
{
	static_cast<QColor*>(ptr)->setHsvF(static_cast<double>(h), static_cast<double>(s), static_cast<double>(v), static_cast<double>(a));
}

void QColor_SetNamedColor(void* ptr, char* name)
{
	static_cast<QColor*>(ptr)->setNamedColor(QString(name));
}

void QColor_SetRed(void* ptr, int red)
{
	static_cast<QColor*>(ptr)->setRed(red);
}

void QColor_SetRedF(void* ptr, double red)
{
	static_cast<QColor*>(ptr)->setRedF(static_cast<double>(red));
}

void QColor_SetRgb(void* ptr, int r, int g, int b, int a)
{
	static_cast<QColor*>(ptr)->setRgb(r, g, b, a);
}

int QColor_Spec(void* ptr)
{
	return static_cast<QColor*>(ptr)->spec();
}

void* QColor_ToCmyk(void* ptr)
{
	return new QColor(static_cast<QColor*>(ptr)->toCmyk());
}

void* QColor_ToHsl(void* ptr)
{
	return new QColor(static_cast<QColor*>(ptr)->toHsl());
}

void* QColor_ToHsv(void* ptr)
{
	return new QColor(static_cast<QColor*>(ptr)->toHsv());
}

void* QColor_ToRgb(void* ptr)
{
	return new QColor(static_cast<QColor*>(ptr)->toRgb());
}

int QColor_Value(void* ptr)
{
	return static_cast<QColor*>(ptr)->value();
}

double QColor_ValueF(void* ptr)
{
	return static_cast<double>(static_cast<QColor*>(ptr)->valueF());
}

int QColor_Yellow(void* ptr)
{
	return static_cast<QColor*>(ptr)->yellow();
}

double QColor_YellowF(void* ptr)
{
	return static_cast<double>(static_cast<QColor*>(ptr)->yellowF());
}

void* QConicalGradient_NewQConicalGradient()
{
	return new QConicalGradient();
}

void* QConicalGradient_NewQConicalGradient2(void* center, double angle)
{
	return new QConicalGradient(*static_cast<QPointF*>(center), static_cast<double>(angle));
}

void* QConicalGradient_NewQConicalGradient3(double cx, double cy, double angle)
{
	return new QConicalGradient(static_cast<double>(cx), static_cast<double>(cy), static_cast<double>(angle));
}

double QConicalGradient_Angle(void* ptr)
{
	return static_cast<double>(static_cast<QConicalGradient*>(ptr)->angle());
}

void* QConicalGradient_Center(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QConicalGradient*>(ptr)->center()).x(), static_cast<QPointF>(static_cast<QConicalGradient*>(ptr)->center()).y());
}

void QConicalGradient_SetAngle(void* ptr, double angle)
{
	static_cast<QConicalGradient*>(ptr)->setAngle(static_cast<double>(angle));
}

void QConicalGradient_SetCenter(void* ptr, void* center)
{
	static_cast<QConicalGradient*>(ptr)->setCenter(*static_cast<QPointF*>(center));
}

void QConicalGradient_SetCenter2(void* ptr, double x, double y)
{
	static_cast<QConicalGradient*>(ptr)->setCenter(static_cast<double>(x), static_cast<double>(y));
}

void* QContextMenuEvent_NewQContextMenuEvent3(int reason, void* pos)
{
	return new QContextMenuEvent(static_cast<QContextMenuEvent::Reason>(reason), *static_cast<QPoint*>(pos));
}

void* QContextMenuEvent_NewQContextMenuEvent2(int reason, void* pos, void* globalPos)
{
	return new QContextMenuEvent(static_cast<QContextMenuEvent::Reason>(reason), *static_cast<QPoint*>(pos), *static_cast<QPoint*>(globalPos));
}

void* QContextMenuEvent_NewQContextMenuEvent(int reason, void* pos, void* globalPos, int modifiers)
{
	return new QContextMenuEvent(static_cast<QContextMenuEvent::Reason>(reason), *static_cast<QPoint*>(pos), *static_cast<QPoint*>(globalPos), static_cast<Qt::KeyboardModifier>(modifiers));
}

void* QContextMenuEvent_GlobalPos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QContextMenuEvent*>(ptr)->globalPos()).x(), static_cast<QPoint>(static_cast<QContextMenuEvent*>(ptr)->globalPos()).y());
}

int QContextMenuEvent_GlobalX(void* ptr)
{
	return static_cast<QContextMenuEvent*>(ptr)->globalX();
}

int QContextMenuEvent_GlobalY(void* ptr)
{
	return static_cast<QContextMenuEvent*>(ptr)->globalY();
}

void* QContextMenuEvent_Pos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QContextMenuEvent*>(ptr)->pos()).x(), static_cast<QPoint>(static_cast<QContextMenuEvent*>(ptr)->pos()).y());
}

int QContextMenuEvent_Reason(void* ptr)
{
	return static_cast<QContextMenuEvent*>(ptr)->reason();
}

int QContextMenuEvent_X(void* ptr)
{
	return static_cast<QContextMenuEvent*>(ptr)->x();
}

int QContextMenuEvent_Y(void* ptr)
{
	return static_cast<QContextMenuEvent*>(ptr)->y();
}

void* QCursor_QCursor_Pos()
{
	return new QPoint(static_cast<QPoint>(QCursor::pos()).x(), static_cast<QPoint>(QCursor::pos()).y());
}

void* QCursor_QCursor_Pos2(void* screen)
{
	return new QPoint(static_cast<QPoint>(QCursor::pos(static_cast<QScreen*>(screen))).x(), static_cast<QPoint>(QCursor::pos(static_cast<QScreen*>(screen))).y());
}

void QCursor_QCursor_SetPos2(void* screen, int x, int y)
{
	QCursor::setPos(static_cast<QScreen*>(screen), x, y);
}

void QCursor_QCursor_SetPos(int x, int y)
{
	QCursor::setPos(x, y);
}

void* QCursor_NewQCursor()
{
	return new QCursor();
}

void* QCursor_NewQCursor6(void* other)
{
	return new QCursor(*static_cast<QCursor*>(other));
}

void* QCursor_NewQCursor2(int shape)
{
	return new QCursor(static_cast<Qt::CursorShape>(shape));
}

void* QCursor_NewQCursor3(void* bitmap, void* mask, int hotX, int hotY)
{
	return new QCursor(*static_cast<QBitmap*>(bitmap), *static_cast<QBitmap*>(mask), hotX, hotY);
}

void* QCursor_NewQCursor5(void* c)
{
	return new QCursor(*static_cast<QCursor*>(c));
}

void* QCursor_NewQCursor4(void* pixmap, int hotX, int hotY)
{
	return new QCursor(*static_cast<QPixmap*>(pixmap), hotX, hotY);
}

void* QCursor_Bitmap(void* ptr)
{
	return const_cast<QBitmap*>(static_cast<QCursor*>(ptr)->bitmap());
}

void* QCursor_HotSpot(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QCursor*>(ptr)->hotSpot()).x(), static_cast<QPoint>(static_cast<QCursor*>(ptr)->hotSpot()).y());
}

void* QCursor_Mask(void* ptr)
{
	return const_cast<QBitmap*>(static_cast<QCursor*>(ptr)->mask());
}

void* QCursor_Pixmap(void* ptr)
{
	return new QPixmap(static_cast<QCursor*>(ptr)->pixmap());
}

void QCursor_QCursor_SetPos4(void* screen, void* p)
{
	QCursor::setPos(static_cast<QScreen*>(screen), *static_cast<QPoint*>(p));
}

void QCursor_QCursor_SetPos3(void* p)
{
	QCursor::setPos(*static_cast<QPoint*>(p));
}

void QCursor_SetShape(void* ptr, int shape)
{
	static_cast<QCursor*>(ptr)->setShape(static_cast<Qt::CursorShape>(shape));
}

int QCursor_Shape(void* ptr)
{
	return static_cast<QCursor*>(ptr)->shape();
}

void QCursor_DestroyQCursor(void* ptr)
{
	static_cast<QCursor*>(ptr)->~QCursor();
}

int QDesktopServices_QDesktopServices_OpenUrl(void* url)
{
	return QDesktopServices::openUrl(*static_cast<QUrl*>(url));
}

void QDesktopServices_QDesktopServices_SetUrlHandler(char* scheme, void* receiver, char* method)
{
	QDesktopServices::setUrlHandler(QString(scheme), static_cast<QObject*>(receiver), const_cast<const char*>(method));
}

void QDesktopServices_QDesktopServices_UnsetUrlHandler(char* scheme)
{
	QDesktopServices::unsetUrlHandler(QString(scheme));
}

class MyQDoubleValidator: public QDoubleValidator
{
public:
	MyQDoubleValidator(QObject *parent) : QDoubleValidator(parent) {};
	QValidator::State validate(QString & input, int & pos) const { return static_cast<QValidator::State>(callbackQDoubleValidator_Validate(const_cast<MyQDoubleValidator*>(this), this->objectName().toUtf8().data(), input.toUtf8().data(), pos)); };
	void fixup(QString & input) const { callbackQDoubleValidator_Fixup(const_cast<MyQDoubleValidator*>(this), this->objectName().toUtf8().data(), input.toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQDoubleValidator_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQDoubleValidator_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDoubleValidator_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQDoubleValidator_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQDoubleValidator_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDoubleValidator_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQDoubleValidator_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDoubleValidator_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDoubleValidator_MetaObject(const_cast<MyQDoubleValidator*>(this), this->objectName().toUtf8().data())); };
};

int QDoubleValidator_Notation(void* ptr)
{
	return static_cast<QDoubleValidator*>(ptr)->notation();
}

void QDoubleValidator_SetDecimals(void* ptr, int vin)
{
	static_cast<QDoubleValidator*>(ptr)->setDecimals(vin);
}

void QDoubleValidator_SetNotation(void* ptr, int vno)
{
	static_cast<QDoubleValidator*>(ptr)->setNotation(static_cast<QDoubleValidator::Notation>(vno));
}





void* QDoubleValidator_NewQDoubleValidator(void* parent)
{
	return new MyQDoubleValidator(static_cast<QObject*>(parent));
}

int QDoubleValidator_Decimals(void* ptr)
{
	return static_cast<QDoubleValidator*>(ptr)->decimals();
}

void QDoubleValidator_DestroyQDoubleValidator(void* ptr)
{
	static_cast<QDoubleValidator*>(ptr)->~QDoubleValidator();
}





void QDoubleValidator_TimerEvent(void* ptr, void* event)
{
	static_cast<QDoubleValidator*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDoubleValidator_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDoubleValidator*>(ptr)->QDoubleValidator::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDoubleValidator_ChildEvent(void* ptr, void* event)
{
	static_cast<QDoubleValidator*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDoubleValidator_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDoubleValidator*>(ptr)->QDoubleValidator::childEvent(static_cast<QChildEvent*>(event));
}

void QDoubleValidator_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDoubleValidator*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDoubleValidator_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDoubleValidator*>(ptr)->QDoubleValidator::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDoubleValidator_CustomEvent(void* ptr, void* event)
{
	static_cast<QDoubleValidator*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDoubleValidator_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDoubleValidator*>(ptr)->QDoubleValidator::customEvent(static_cast<QEvent*>(event));
}

void QDoubleValidator_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDoubleValidator*>(ptr), "deleteLater");
}

void QDoubleValidator_DeleteLaterDefault(void* ptr)
{
	static_cast<QDoubleValidator*>(ptr)->QDoubleValidator::deleteLater();
}

void QDoubleValidator_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDoubleValidator*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDoubleValidator_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDoubleValidator*>(ptr)->QDoubleValidator::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QDoubleValidator_Event(void* ptr, void* e)
{
	return static_cast<QDoubleValidator*>(ptr)->event(static_cast<QEvent*>(e));
}

int QDoubleValidator_EventDefault(void* ptr, void* e)
{
	return static_cast<QDoubleValidator*>(ptr)->QDoubleValidator::event(static_cast<QEvent*>(e));
}

int QDoubleValidator_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDoubleValidator*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QDoubleValidator_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDoubleValidator*>(ptr)->QDoubleValidator::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDoubleValidator_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDoubleValidator*>(ptr)->metaObject());
}

void* QDoubleValidator_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDoubleValidator*>(ptr)->QDoubleValidator::metaObject());
}

class MyQDrag: public QDrag
{
public:
	MyQDrag(QObject *dragSource) : QDrag(dragSource) {};
	void Signal_ActionChanged(Qt::DropAction action) { callbackQDrag_ActionChanged(this, this->objectName().toUtf8().data(), action); };
	void Signal_TargetChanged(QObject * newTarget) { callbackQDrag_TargetChanged(this, this->objectName().toUtf8().data(), newTarget); };
	void timerEvent(QTimerEvent * event) { callbackQDrag_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQDrag_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDrag_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQDrag_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQDrag_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDrag_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQDrag_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDrag_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDrag_MetaObject(const_cast<MyQDrag*>(this), this->objectName().toUtf8().data())); };
};

void* QDrag_NewQDrag(void* dragSource)
{
	return new MyQDrag(static_cast<QObject*>(dragSource));
}

void QDrag_ConnectActionChanged(void* ptr)
{
	QObject::connect(static_cast<QDrag*>(ptr), static_cast<void (QDrag::*)(Qt::DropAction)>(&QDrag::actionChanged), static_cast<MyQDrag*>(ptr), static_cast<void (MyQDrag::*)(Qt::DropAction)>(&MyQDrag::Signal_ActionChanged));
}

void QDrag_DisconnectActionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDrag*>(ptr), static_cast<void (QDrag::*)(Qt::DropAction)>(&QDrag::actionChanged), static_cast<MyQDrag*>(ptr), static_cast<void (MyQDrag::*)(Qt::DropAction)>(&MyQDrag::Signal_ActionChanged));
}

void QDrag_ActionChanged(void* ptr, int action)
{
	static_cast<QDrag*>(ptr)->actionChanged(static_cast<Qt::DropAction>(action));
}

int QDrag_DefaultAction(void* ptr)
{
	return static_cast<QDrag*>(ptr)->defaultAction();
}

void* QDrag_DragCursor(void* ptr, int action)
{
	return new QPixmap(static_cast<QDrag*>(ptr)->dragCursor(static_cast<Qt::DropAction>(action)));
}

int QDrag_Exec(void* ptr, int supportedActions)
{
	return static_cast<QDrag*>(ptr)->exec(static_cast<Qt::DropAction>(supportedActions));
}

int QDrag_Exec2(void* ptr, int supportedActions, int defaultDropAction)
{
	return static_cast<QDrag*>(ptr)->exec(static_cast<Qt::DropAction>(supportedActions), static_cast<Qt::DropAction>(defaultDropAction));
}

void* QDrag_HotSpot(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QDrag*>(ptr)->hotSpot()).x(), static_cast<QPoint>(static_cast<QDrag*>(ptr)->hotSpot()).y());
}

void* QDrag_MimeData(void* ptr)
{
	return static_cast<QDrag*>(ptr)->mimeData();
}

void* QDrag_Pixmap(void* ptr)
{
	return new QPixmap(static_cast<QDrag*>(ptr)->pixmap());
}

void QDrag_SetDragCursor(void* ptr, void* cursor, int action)
{
	static_cast<QDrag*>(ptr)->setDragCursor(*static_cast<QPixmap*>(cursor), static_cast<Qt::DropAction>(action));
}

void QDrag_SetHotSpot(void* ptr, void* hotspot)
{
	static_cast<QDrag*>(ptr)->setHotSpot(*static_cast<QPoint*>(hotspot));
}

void QDrag_SetMimeData(void* ptr, void* data)
{
	static_cast<QDrag*>(ptr)->setMimeData(static_cast<QMimeData*>(data));
}

void QDrag_SetPixmap(void* ptr, void* pixmap)
{
	static_cast<QDrag*>(ptr)->setPixmap(*static_cast<QPixmap*>(pixmap));
}

void* QDrag_Source(void* ptr)
{
	return static_cast<QDrag*>(ptr)->source();
}

int QDrag_SupportedActions(void* ptr)
{
	return static_cast<QDrag*>(ptr)->supportedActions();
}

void* QDrag_Target(void* ptr)
{
	return static_cast<QDrag*>(ptr)->target();
}

void QDrag_ConnectTargetChanged(void* ptr)
{
	QObject::connect(static_cast<QDrag*>(ptr), static_cast<void (QDrag::*)(QObject *)>(&QDrag::targetChanged), static_cast<MyQDrag*>(ptr), static_cast<void (MyQDrag::*)(QObject *)>(&MyQDrag::Signal_TargetChanged));
}

void QDrag_DisconnectTargetChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDrag*>(ptr), static_cast<void (QDrag::*)(QObject *)>(&QDrag::targetChanged), static_cast<MyQDrag*>(ptr), static_cast<void (MyQDrag::*)(QObject *)>(&MyQDrag::Signal_TargetChanged));
}

void QDrag_TargetChanged(void* ptr, void* newTarget)
{
	static_cast<QDrag*>(ptr)->targetChanged(static_cast<QObject*>(newTarget));
}

void QDrag_DestroyQDrag(void* ptr)
{
	static_cast<QDrag*>(ptr)->~QDrag();
}

void QDrag_TimerEvent(void* ptr, void* event)
{
	static_cast<QDrag*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDrag_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QDrag*>(ptr)->QDrag::timerEvent(static_cast<QTimerEvent*>(event));
}

void QDrag_ChildEvent(void* ptr, void* event)
{
	static_cast<QDrag*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDrag_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QDrag*>(ptr)->QDrag::childEvent(static_cast<QChildEvent*>(event));
}

void QDrag_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QDrag*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDrag_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDrag*>(ptr)->QDrag::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDrag_CustomEvent(void* ptr, void* event)
{
	static_cast<QDrag*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDrag_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QDrag*>(ptr)->QDrag::customEvent(static_cast<QEvent*>(event));
}

void QDrag_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QDrag*>(ptr), "deleteLater");
}

void QDrag_DeleteLaterDefault(void* ptr)
{
	static_cast<QDrag*>(ptr)->QDrag::deleteLater();
}

void QDrag_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QDrag*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDrag_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QDrag*>(ptr)->QDrag::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QDrag_Event(void* ptr, void* e)
{
	return static_cast<QDrag*>(ptr)->event(static_cast<QEvent*>(e));
}

int QDrag_EventDefault(void* ptr, void* e)
{
	return static_cast<QDrag*>(ptr)->QDrag::event(static_cast<QEvent*>(e));
}

int QDrag_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QDrag*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QDrag_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QDrag*>(ptr)->QDrag::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QDrag_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDrag*>(ptr)->metaObject());
}

void* QDrag_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QDrag*>(ptr)->QDrag::metaObject());
}

void* QDragEnterEvent_NewQDragEnterEvent(void* point, int actions, void* data, int buttons, int modifiers)
{
	return new QDragEnterEvent(*static_cast<QPoint*>(point), static_cast<Qt::DropAction>(actions), static_cast<QMimeData*>(data), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

void* QDragLeaveEvent_NewQDragLeaveEvent()
{
	return new QDragLeaveEvent();
}

void* QDragMoveEvent_NewQDragMoveEvent(void* pos, int actions, void* data, int buttons, int modifiers, int ty)
{
	return new QDragMoveEvent(*static_cast<QPoint*>(pos), static_cast<Qt::DropAction>(actions), static_cast<QMimeData*>(data), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<QEvent::Type>(ty));
}

void QDragMoveEvent_Accept2(void* ptr)
{
	static_cast<QDragMoveEvent*>(ptr)->accept();
}

void QDragMoveEvent_Accept(void* ptr, void* rectangle)
{
	static_cast<QDragMoveEvent*>(ptr)->accept(*static_cast<QRect*>(rectangle));
}

void* QDragMoveEvent_AnswerRect(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QDragMoveEvent*>(ptr)->answerRect()).x(), static_cast<QRect>(static_cast<QDragMoveEvent*>(ptr)->answerRect()).y(), static_cast<QRect>(static_cast<QDragMoveEvent*>(ptr)->answerRect()).width(), static_cast<QRect>(static_cast<QDragMoveEvent*>(ptr)->answerRect()).height());
}

void QDragMoveEvent_Ignore2(void* ptr)
{
	static_cast<QDragMoveEvent*>(ptr)->ignore();
}

void QDragMoveEvent_Ignore(void* ptr, void* rectangle)
{
	static_cast<QDragMoveEvent*>(ptr)->ignore(*static_cast<QRect*>(rectangle));
}

void QDragMoveEvent_DestroyQDragMoveEvent(void* ptr)
{
	static_cast<QDragMoveEvent*>(ptr)->~QDragMoveEvent();
}

void QDropEvent_SetDropAction(void* ptr, int action)
{
	static_cast<QDropEvent*>(ptr)->setDropAction(static_cast<Qt::DropAction>(action));
}

void* QDropEvent_NewQDropEvent(void* pos, int actions, void* data, int buttons, int modifiers, int ty)
{
	return new QDropEvent(*static_cast<QPointF*>(pos), static_cast<Qt::DropAction>(actions), static_cast<QMimeData*>(data), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<QEvent::Type>(ty));
}

void QDropEvent_AcceptProposedAction(void* ptr)
{
	static_cast<QDropEvent*>(ptr)->acceptProposedAction();
}

int QDropEvent_DropAction(void* ptr)
{
	return static_cast<QDropEvent*>(ptr)->dropAction();
}

int QDropEvent_KeyboardModifiers(void* ptr)
{
	return static_cast<QDropEvent*>(ptr)->keyboardModifiers();
}

void* QDropEvent_MimeData(void* ptr)
{
	return const_cast<QMimeData*>(static_cast<QDropEvent*>(ptr)->mimeData());
}

int QDropEvent_MouseButtons(void* ptr)
{
	return static_cast<QDropEvent*>(ptr)->mouseButtons();
}

void* QDropEvent_Pos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QDropEvent*>(ptr)->pos()).x(), static_cast<QPoint>(static_cast<QDropEvent*>(ptr)->pos()).y());
}

void* QDropEvent_PosF(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QDropEvent*>(ptr)->posF()).x(), static_cast<QPointF>(static_cast<QDropEvent*>(ptr)->posF()).y());
}

int QDropEvent_PossibleActions(void* ptr)
{
	return static_cast<QDropEvent*>(ptr)->possibleActions();
}

int QDropEvent_ProposedAction(void* ptr)
{
	return static_cast<QDropEvent*>(ptr)->proposedAction();
}

void* QDropEvent_Source(void* ptr)
{
	return static_cast<QDropEvent*>(ptr)->source();
}

void* QEnterEvent_NewQEnterEvent(void* localPos, void* windowPos, void* screenPos)
{
	return new QEnterEvent(*static_cast<QPointF*>(localPos), *static_cast<QPointF*>(windowPos), *static_cast<QPointF*>(screenPos));
}

void* QEnterEvent_GlobalPos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QEnterEvent*>(ptr)->globalPos()).x(), static_cast<QPoint>(static_cast<QEnterEvent*>(ptr)->globalPos()).y());
}

int QEnterEvent_GlobalX(void* ptr)
{
	return static_cast<QEnterEvent*>(ptr)->globalX();
}

int QEnterEvent_GlobalY(void* ptr)
{
	return static_cast<QEnterEvent*>(ptr)->globalY();
}

void* QEnterEvent_LocalPos(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QEnterEvent*>(ptr)->localPos()).x(), static_cast<QPointF>(static_cast<QEnterEvent*>(ptr)->localPos()).y());
}

void* QEnterEvent_Pos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QEnterEvent*>(ptr)->pos()).x(), static_cast<QPoint>(static_cast<QEnterEvent*>(ptr)->pos()).y());
}

void* QEnterEvent_ScreenPos(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QEnterEvent*>(ptr)->screenPos()).x(), static_cast<QPointF>(static_cast<QEnterEvent*>(ptr)->screenPos()).y());
}

void* QEnterEvent_WindowPos(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QEnterEvent*>(ptr)->windowPos()).x(), static_cast<QPointF>(static_cast<QEnterEvent*>(ptr)->windowPos()).y());
}

int QEnterEvent_X(void* ptr)
{
	return static_cast<QEnterEvent*>(ptr)->x();
}

int QEnterEvent_Y(void* ptr)
{
	return static_cast<QEnterEvent*>(ptr)->y();
}

void* QExposeEvent_NewQExposeEvent(void* exposeRegion)
{
	return new QExposeEvent(*static_cast<QRegion*>(exposeRegion));
}

void* QExposeEvent_Region(void* ptr)
{
	return new QRegion(static_cast<QExposeEvent*>(ptr)->region());
}

int QFileOpenEvent_OpenFile(void* ptr, void* file, int flags)
{
	return static_cast<QFileOpenEvent*>(ptr)->openFile(*static_cast<QFile*>(file), static_cast<QIODevice::OpenModeFlag>(flags));
}

char* QFileOpenEvent_File(void* ptr)
{
	return static_cast<QFileOpenEvent*>(ptr)->file().toUtf8().data();
}

void* QFileOpenEvent_Url(void* ptr)
{
	return new QUrl(static_cast<QFileOpenEvent*>(ptr)->url());
}

void* QFocusEvent_NewQFocusEvent(int ty, int reason)
{
	return new QFocusEvent(static_cast<QEvent::Type>(ty), static_cast<Qt::FocusReason>(reason));
}

int QFocusEvent_GotFocus(void* ptr)
{
	return static_cast<QFocusEvent*>(ptr)->gotFocus();
}

int QFocusEvent_LostFocus(void* ptr)
{
	return static_cast<QFocusEvent*>(ptr)->lostFocus();
}

int QFocusEvent_Reason(void* ptr)
{
	return static_cast<QFocusEvent*>(ptr)->reason();
}

int QFont_Times_Type()
{
	return QFont::Times;
}

int QFont_Courier_Type()
{
	return QFont::Courier;
}

int QFont_OldEnglish_Type()
{
	return QFont::OldEnglish;
}

int QFont_System_Type()
{
	return QFont::System;
}

int QFont_AnyStyle_Type()
{
	return QFont::AnyStyle;
}

int QFont_Cursive_Type()
{
	return QFont::Cursive;
}

int QFont_Monospace_Type()
{
	return QFont::Monospace;
}

int QFont_Fantasy_Type()
{
	return QFont::Fantasy;
}

char* QFont_DefaultFamily(void* ptr)
{
	return static_cast<QFont*>(ptr)->defaultFamily().toUtf8().data();
}

char* QFont_LastResortFamily(void* ptr)
{
	return static_cast<QFont*>(ptr)->lastResortFamily().toUtf8().data();
}

char* QFont_LastResortFont(void* ptr)
{
	return static_cast<QFont*>(ptr)->lastResortFont().toUtf8().data();
}

void* QFont_NewQFont()
{
	return new QFont();
}

void* QFont_NewQFont4(void* font)
{
	return new QFont(*static_cast<QFont*>(font));
}

void* QFont_NewQFont3(void* font, void* pd)
{
	return new QFont(*static_cast<QFont*>(font), static_cast<QPaintDevice*>(pd));
}

void* QFont_NewQFont2(char* family, int pointSize, int weight, int italic)
{
	return new QFont(QString(family), pointSize, weight, italic != 0);
}

int QFont_Bold(void* ptr)
{
	return static_cast<QFont*>(ptr)->bold();
}

int QFont_Capitalization(void* ptr)
{
	return static_cast<QFont*>(ptr)->capitalization();
}

int QFont_ExactMatch(void* ptr)
{
	return static_cast<QFont*>(ptr)->exactMatch();
}

char* QFont_Family(void* ptr)
{
	return static_cast<QFont*>(ptr)->family().toUtf8().data();
}

int QFont_FixedPitch(void* ptr)
{
	return static_cast<QFont*>(ptr)->fixedPitch();
}

int QFont_FromString(void* ptr, char* descrip)
{
	return static_cast<QFont*>(ptr)->fromString(QString(descrip));
}

int QFont_HintingPreference(void* ptr)
{
	return static_cast<QFont*>(ptr)->hintingPreference();
}

void QFont_QFont_InsertSubstitution(char* familyName, char* substituteName)
{
	QFont::insertSubstitution(QString(familyName), QString(substituteName));
}

void QFont_QFont_InsertSubstitutions(char* familyName, char* substituteNames)
{
	QFont::insertSubstitutions(QString(familyName), QString(substituteNames).split("|", QString::SkipEmptyParts));
}

int QFont_IsCopyOf(void* ptr, void* f)
{
	return static_cast<QFont*>(ptr)->isCopyOf(*static_cast<QFont*>(f));
}

int QFont_Italic(void* ptr)
{
	return static_cast<QFont*>(ptr)->italic();
}

int QFont_Kerning(void* ptr)
{
	return static_cast<QFont*>(ptr)->kerning();
}

char* QFont_Key(void* ptr)
{
	return static_cast<QFont*>(ptr)->key().toUtf8().data();
}

double QFont_LetterSpacing(void* ptr)
{
	return static_cast<double>(static_cast<QFont*>(ptr)->letterSpacing());
}

int QFont_LetterSpacingType(void* ptr)
{
	return static_cast<QFont*>(ptr)->letterSpacingType();
}

int QFont_Overline(void* ptr)
{
	return static_cast<QFont*>(ptr)->overline();
}

int QFont_PixelSize(void* ptr)
{
	return static_cast<QFont*>(ptr)->pixelSize();
}

int QFont_PointSize(void* ptr)
{
	return static_cast<QFont*>(ptr)->pointSize();
}

double QFont_PointSizeF(void* ptr)
{
	return static_cast<double>(static_cast<QFont*>(ptr)->pointSizeF());
}

void QFont_QFont_RemoveSubstitutions(char* familyName)
{
	QFont::removeSubstitutions(QString(familyName));
}

void* QFont_Resolve(void* ptr, void* other)
{
	return new QFont(static_cast<QFont*>(ptr)->resolve(*static_cast<QFont*>(other)));
}

void QFont_SetBold(void* ptr, int enable)
{
	static_cast<QFont*>(ptr)->setBold(enable != 0);
}

void QFont_SetCapitalization(void* ptr, int caps)
{
	static_cast<QFont*>(ptr)->setCapitalization(static_cast<QFont::Capitalization>(caps));
}

void QFont_SetFamily(void* ptr, char* family)
{
	static_cast<QFont*>(ptr)->setFamily(QString(family));
}

void QFont_SetFixedPitch(void* ptr, int enable)
{
	static_cast<QFont*>(ptr)->setFixedPitch(enable != 0);
}

void QFont_SetHintingPreference(void* ptr, int hintingPreference)
{
	static_cast<QFont*>(ptr)->setHintingPreference(static_cast<QFont::HintingPreference>(hintingPreference));
}

void QFont_SetItalic(void* ptr, int enable)
{
	static_cast<QFont*>(ptr)->setItalic(enable != 0);
}

void QFont_SetKerning(void* ptr, int enable)
{
	static_cast<QFont*>(ptr)->setKerning(enable != 0);
}

void QFont_SetLetterSpacing(void* ptr, int ty, double spacing)
{
	static_cast<QFont*>(ptr)->setLetterSpacing(static_cast<QFont::SpacingType>(ty), static_cast<double>(spacing));
}

void QFont_SetOverline(void* ptr, int enable)
{
	static_cast<QFont*>(ptr)->setOverline(enable != 0);
}

void QFont_SetPixelSize(void* ptr, int pixelSize)
{
	static_cast<QFont*>(ptr)->setPixelSize(pixelSize);
}

void QFont_SetPointSize(void* ptr, int pointSize)
{
	static_cast<QFont*>(ptr)->setPointSize(pointSize);
}

void QFont_SetPointSizeF(void* ptr, double pointSize)
{
	static_cast<QFont*>(ptr)->setPointSizeF(static_cast<double>(pointSize));
}

void QFont_SetStretch(void* ptr, int factor)
{
	static_cast<QFont*>(ptr)->setStretch(factor);
}

void QFont_SetStrikeOut(void* ptr, int enable)
{
	static_cast<QFont*>(ptr)->setStrikeOut(enable != 0);
}

void QFont_SetStyle(void* ptr, int style)
{
	static_cast<QFont*>(ptr)->setStyle(static_cast<QFont::Style>(style));
}

void QFont_SetStyleHint(void* ptr, int hint, int strategy)
{
	static_cast<QFont*>(ptr)->setStyleHint(static_cast<QFont::StyleHint>(hint), static_cast<QFont::StyleStrategy>(strategy));
}

void QFont_SetStyleName(void* ptr, char* styleName)
{
	static_cast<QFont*>(ptr)->setStyleName(QString(styleName));
}

void QFont_SetStyleStrategy(void* ptr, int s)
{
	static_cast<QFont*>(ptr)->setStyleStrategy(static_cast<QFont::StyleStrategy>(s));
}

void QFont_SetUnderline(void* ptr, int enable)
{
	static_cast<QFont*>(ptr)->setUnderline(enable != 0);
}

void QFont_SetWeight(void* ptr, int weight)
{
	static_cast<QFont*>(ptr)->setWeight(weight);
}

void QFont_SetWordSpacing(void* ptr, double spacing)
{
	static_cast<QFont*>(ptr)->setWordSpacing(static_cast<double>(spacing));
}

int QFont_Stretch(void* ptr)
{
	return static_cast<QFont*>(ptr)->stretch();
}

int QFont_StrikeOut(void* ptr)
{
	return static_cast<QFont*>(ptr)->strikeOut();
}

int QFont_Style(void* ptr)
{
	return static_cast<QFont*>(ptr)->style();
}

int QFont_StyleHint(void* ptr)
{
	return static_cast<QFont*>(ptr)->styleHint();
}

char* QFont_StyleName(void* ptr)
{
	return static_cast<QFont*>(ptr)->styleName().toUtf8().data();
}

int QFont_StyleStrategy(void* ptr)
{
	return static_cast<QFont*>(ptr)->styleStrategy();
}

char* QFont_QFont_Substitute(char* familyName)
{
	return QFont::substitute(QString(familyName)).toUtf8().data();
}

char* QFont_QFont_Substitutes(char* familyName)
{
	return QFont::substitutes(QString(familyName)).join("|").toUtf8().data();
}

char* QFont_QFont_Substitutions()
{
	return QFont::substitutions().join("|").toUtf8().data();
}

void QFont_Swap(void* ptr, void* other)
{
	static_cast<QFont*>(ptr)->swap(*static_cast<QFont*>(other));
}

char* QFont_ToString(void* ptr)
{
	return static_cast<QFont*>(ptr)->toString().toUtf8().data();
}

int QFont_Underline(void* ptr)
{
	return static_cast<QFont*>(ptr)->underline();
}

int QFont_Weight(void* ptr)
{
	return static_cast<QFont*>(ptr)->weight();
}

double QFont_WordSpacing(void* ptr)
{
	return static_cast<double>(static_cast<QFont*>(ptr)->wordSpacing());
}

void QFont_DestroyQFont(void* ptr)
{
	static_cast<QFont*>(ptr)->~QFont();
}

int QFontDatabase_Ogham_Type()
{
	return QFontDatabase::Ogham;
}

int QFontDatabase_Runic_Type()
{
	return QFontDatabase::Runic;
}

int QFontDatabase_Nko_Type()
{
	return QFontDatabase::Nko;
}

int QFontDatabase_WritingSystemsCount_Type()
{
	return QFontDatabase::WritingSystemsCount;
}

int QFontDatabase_QFontDatabase_RemoveAllApplicationFonts()
{
	return QFontDatabase::removeAllApplicationFonts();
}

int QFontDatabase_QFontDatabase_RemoveApplicationFont(int id)
{
	return QFontDatabase::removeApplicationFont(id);
}

void* QFontDatabase_NewQFontDatabase()
{
	return new QFontDatabase();
}

int QFontDatabase_QFontDatabase_AddApplicationFont(char* fileName)
{
	return QFontDatabase::addApplicationFont(QString(fileName));
}

int QFontDatabase_QFontDatabase_AddApplicationFontFromData(char* fontData)
{
	return QFontDatabase::addApplicationFontFromData(QByteArray(fontData));
}

char* QFontDatabase_QFontDatabase_ApplicationFontFamilies(int id)
{
	return QFontDatabase::applicationFontFamilies(id).join("|").toUtf8().data();
}

int QFontDatabase_Bold(void* ptr, char* family, char* style)
{
	return static_cast<QFontDatabase*>(ptr)->bold(QString(family), QString(style));
}

char* QFontDatabase_Families(void* ptr, int writingSystem)
{
	return static_cast<QFontDatabase*>(ptr)->families(static_cast<QFontDatabase::WritingSystem>(writingSystem)).join("|").toUtf8().data();
}

void* QFontDatabase_Font(void* ptr, char* family, char* style, int pointSize)
{
	return new QFont(static_cast<QFontDatabase*>(ptr)->font(QString(family), QString(style), pointSize));
}

int QFontDatabase_IsBitmapScalable(void* ptr, char* family, char* style)
{
	return static_cast<QFontDatabase*>(ptr)->isBitmapScalable(QString(family), QString(style));
}

int QFontDatabase_IsFixedPitch(void* ptr, char* family, char* style)
{
	return static_cast<QFontDatabase*>(ptr)->isFixedPitch(QString(family), QString(style));
}

int QFontDatabase_IsPrivateFamily(void* ptr, char* family)
{
	return static_cast<QFontDatabase*>(ptr)->isPrivateFamily(QString(family));
}

int QFontDatabase_IsScalable(void* ptr, char* family, char* style)
{
	return static_cast<QFontDatabase*>(ptr)->isScalable(QString(family), QString(style));
}

int QFontDatabase_IsSmoothlyScalable(void* ptr, char* family, char* style)
{
	return static_cast<QFontDatabase*>(ptr)->isSmoothlyScalable(QString(family), QString(style));
}

int QFontDatabase_Italic(void* ptr, char* family, char* style)
{
	return static_cast<QFontDatabase*>(ptr)->italic(QString(family), QString(style));
}

char* QFontDatabase_StyleString(void* ptr, void* font)
{
	return static_cast<QFontDatabase*>(ptr)->styleString(*static_cast<QFont*>(font)).toUtf8().data();
}

char* QFontDatabase_StyleString2(void* ptr, void* fontInfo)
{
	return static_cast<QFontDatabase*>(ptr)->styleString(*static_cast<QFontInfo*>(fontInfo)).toUtf8().data();
}

char* QFontDatabase_Styles(void* ptr, char* family)
{
	return static_cast<QFontDatabase*>(ptr)->styles(QString(family)).join("|").toUtf8().data();
}

void* QFontDatabase_QFontDatabase_SystemFont(int ty)
{
	return new QFont(QFontDatabase::systemFont(static_cast<QFontDatabase::SystemFont>(ty)));
}

int QFontDatabase_Weight(void* ptr, char* family, char* style)
{
	return static_cast<QFontDatabase*>(ptr)->weight(QString(family), QString(style));
}

char* QFontDatabase_QFontDatabase_WritingSystemName(int writingSystem)
{
	return QFontDatabase::writingSystemName(static_cast<QFontDatabase::WritingSystem>(writingSystem)).toUtf8().data();
}

char* QFontDatabase_QFontDatabase_WritingSystemSample(int writingSystem)
{
	return QFontDatabase::writingSystemSample(static_cast<QFontDatabase::WritingSystem>(writingSystem)).toUtf8().data();
}

void* QFontInfo_NewQFontInfo(void* font)
{
	return new QFontInfo(*static_cast<QFont*>(font));
}

void* QFontInfo_NewQFontInfo2(void* fi)
{
	return new QFontInfo(*static_cast<QFontInfo*>(fi));
}

int QFontInfo_Bold(void* ptr)
{
	return static_cast<QFontInfo*>(ptr)->bold();
}

int QFontInfo_ExactMatch(void* ptr)
{
	return static_cast<QFontInfo*>(ptr)->exactMatch();
}

char* QFontInfo_Family(void* ptr)
{
	return static_cast<QFontInfo*>(ptr)->family().toUtf8().data();
}

int QFontInfo_FixedPitch(void* ptr)
{
	return static_cast<QFontInfo*>(ptr)->fixedPitch();
}

int QFontInfo_Italic(void* ptr)
{
	return static_cast<QFontInfo*>(ptr)->italic();
}

int QFontInfo_PixelSize(void* ptr)
{
	return static_cast<QFontInfo*>(ptr)->pixelSize();
}

int QFontInfo_PointSize(void* ptr)
{
	return static_cast<QFontInfo*>(ptr)->pointSize();
}

double QFontInfo_PointSizeF(void* ptr)
{
	return static_cast<double>(static_cast<QFontInfo*>(ptr)->pointSizeF());
}

int QFontInfo_Style(void* ptr)
{
	return static_cast<QFontInfo*>(ptr)->style();
}

char* QFontInfo_StyleName(void* ptr)
{
	return static_cast<QFontInfo*>(ptr)->styleName().toUtf8().data();
}

void QFontInfo_Swap(void* ptr, void* other)
{
	static_cast<QFontInfo*>(ptr)->swap(*static_cast<QFontInfo*>(other));
}

int QFontInfo_StyleHint(void* ptr)
{
	return static_cast<QFontInfo*>(ptr)->styleHint();
}

int QFontInfo_Weight(void* ptr)
{
	return static_cast<QFontInfo*>(ptr)->weight();
}

void QFontInfo_DestroyQFontInfo(void* ptr)
{
	static_cast<QFontInfo*>(ptr)->~QFontInfo();
}

void* QFontMetrics_NewQFontMetrics(void* font)
{
	return new QFontMetrics(*static_cast<QFont*>(font));
}

void* QFontMetrics_NewQFontMetrics2(void* font, void* paintdevice)
{
	return new QFontMetrics(*static_cast<QFont*>(font), static_cast<QPaintDevice*>(paintdevice));
}

void* QFontMetrics_NewQFontMetrics3(void* fm)
{
	return new QFontMetrics(*static_cast<QFontMetrics*>(fm));
}

int QFontMetrics_Ascent(void* ptr)
{
	return static_cast<QFontMetrics*>(ptr)->ascent();
}

int QFontMetrics_AverageCharWidth(void* ptr)
{
	return static_cast<QFontMetrics*>(ptr)->averageCharWidth();
}

void* QFontMetrics_BoundingRect(void* ptr, void* ch)
{
	return new QRect(static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->boundingRect(*static_cast<QChar*>(ch))).x(), static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->boundingRect(*static_cast<QChar*>(ch))).y(), static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->boundingRect(*static_cast<QChar*>(ch))).width(), static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->boundingRect(*static_cast<QChar*>(ch))).height());
}

void* QFontMetrics_BoundingRect4(void* ptr, void* rect, int flags, char* text, int tabStops, int tabArray)
{
	return new QRect(static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->boundingRect(*static_cast<QRect*>(rect), flags, QString(text), tabStops, &tabArray)).x(), static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->boundingRect(*static_cast<QRect*>(rect), flags, QString(text), tabStops, &tabArray)).y(), static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->boundingRect(*static_cast<QRect*>(rect), flags, QString(text), tabStops, &tabArray)).width(), static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->boundingRect(*static_cast<QRect*>(rect), flags, QString(text), tabStops, &tabArray)).height());
}

void* QFontMetrics_BoundingRect2(void* ptr, char* text)
{
	return new QRect(static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->boundingRect(QString(text))).x(), static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->boundingRect(QString(text))).y(), static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->boundingRect(QString(text))).width(), static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->boundingRect(QString(text))).height());
}

void* QFontMetrics_BoundingRect3(void* ptr, int x, int y, int width, int height, int flags, char* text, int tabStops, int tabArray)
{
	return new QRect(static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->boundingRect(x, y, width, height, flags, QString(text), tabStops, &tabArray)).x(), static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->boundingRect(x, y, width, height, flags, QString(text), tabStops, &tabArray)).y(), static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->boundingRect(x, y, width, height, flags, QString(text), tabStops, &tabArray)).width(), static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->boundingRect(x, y, width, height, flags, QString(text), tabStops, &tabArray)).height());
}

int QFontMetrics_Descent(void* ptr)
{
	return static_cast<QFontMetrics*>(ptr)->descent();
}

char* QFontMetrics_ElidedText(void* ptr, char* text, int mode, int width, int flags)
{
	return static_cast<QFontMetrics*>(ptr)->elidedText(QString(text), static_cast<Qt::TextElideMode>(mode), width, flags).toUtf8().data();
}

int QFontMetrics_Height(void* ptr)
{
	return static_cast<QFontMetrics*>(ptr)->height();
}

int QFontMetrics_InFont(void* ptr, void* ch)
{
	return static_cast<QFontMetrics*>(ptr)->inFont(*static_cast<QChar*>(ch));
}

int QFontMetrics_Leading(void* ptr)
{
	return static_cast<QFontMetrics*>(ptr)->leading();
}

int QFontMetrics_LeftBearing(void* ptr, void* ch)
{
	return static_cast<QFontMetrics*>(ptr)->leftBearing(*static_cast<QChar*>(ch));
}

int QFontMetrics_LineSpacing(void* ptr)
{
	return static_cast<QFontMetrics*>(ptr)->lineSpacing();
}

int QFontMetrics_LineWidth(void* ptr)
{
	return static_cast<QFontMetrics*>(ptr)->lineWidth();
}

int QFontMetrics_MaxWidth(void* ptr)
{
	return static_cast<QFontMetrics*>(ptr)->maxWidth();
}

int QFontMetrics_MinLeftBearing(void* ptr)
{
	return static_cast<QFontMetrics*>(ptr)->minLeftBearing();
}

int QFontMetrics_MinRightBearing(void* ptr)
{
	return static_cast<QFontMetrics*>(ptr)->minRightBearing();
}

int QFontMetrics_OverlinePos(void* ptr)
{
	return static_cast<QFontMetrics*>(ptr)->overlinePos();
}

int QFontMetrics_RightBearing(void* ptr, void* ch)
{
	return static_cast<QFontMetrics*>(ptr)->rightBearing(*static_cast<QChar*>(ch));
}

void* QFontMetrics_Size(void* ptr, int flags, char* text, int tabStops, int tabArray)
{
	return new QSize(static_cast<QSize>(static_cast<QFontMetrics*>(ptr)->size(flags, QString(text), tabStops, &tabArray)).width(), static_cast<QSize>(static_cast<QFontMetrics*>(ptr)->size(flags, QString(text), tabStops, &tabArray)).height());
}

int QFontMetrics_StrikeOutPos(void* ptr)
{
	return static_cast<QFontMetrics*>(ptr)->strikeOutPos();
}

void QFontMetrics_Swap(void* ptr, void* other)
{
	static_cast<QFontMetrics*>(ptr)->swap(*static_cast<QFontMetrics*>(other));
}

void* QFontMetrics_TightBoundingRect(void* ptr, char* text)
{
	return new QRect(static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->tightBoundingRect(QString(text))).x(), static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->tightBoundingRect(QString(text))).y(), static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->tightBoundingRect(QString(text))).width(), static_cast<QRect>(static_cast<QFontMetrics*>(ptr)->tightBoundingRect(QString(text))).height());
}

int QFontMetrics_UnderlinePos(void* ptr)
{
	return static_cast<QFontMetrics*>(ptr)->underlinePos();
}

int QFontMetrics_Width3(void* ptr, void* ch)
{
	return static_cast<QFontMetrics*>(ptr)->width(*static_cast<QChar*>(ch));
}

int QFontMetrics_Width(void* ptr, char* text, int len)
{
	return static_cast<QFontMetrics*>(ptr)->width(QString(text), len);
}

int QFontMetrics_XHeight(void* ptr)
{
	return static_cast<QFontMetrics*>(ptr)->xHeight();
}

void QFontMetrics_DestroyQFontMetrics(void* ptr)
{
	static_cast<QFontMetrics*>(ptr)->~QFontMetrics();
}

void* QFontMetricsF_NewQFontMetricsF(void* font)
{
	return new QFontMetricsF(*static_cast<QFont*>(font));
}

void* QFontMetricsF_NewQFontMetricsF2(void* font, void* paintdevice)
{
	return new QFontMetricsF(*static_cast<QFont*>(font), static_cast<QPaintDevice*>(paintdevice));
}

void* QFontMetricsF_NewQFontMetricsF3(void* fontMetrics)
{
	return new QFontMetricsF(*static_cast<QFontMetrics*>(fontMetrics));
}

void* QFontMetricsF_NewQFontMetricsF4(void* fm)
{
	return new QFontMetricsF(*static_cast<QFontMetricsF*>(fm));
}

double QFontMetricsF_Ascent(void* ptr)
{
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->ascent());
}

double QFontMetricsF_AverageCharWidth(void* ptr)
{
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->averageCharWidth());
}

void* QFontMetricsF_BoundingRect2(void* ptr, void* ch)
{
	return new QRectF(static_cast<QRectF>(static_cast<QFontMetricsF*>(ptr)->boundingRect(*static_cast<QChar*>(ch))).x(), static_cast<QRectF>(static_cast<QFontMetricsF*>(ptr)->boundingRect(*static_cast<QChar*>(ch))).y(), static_cast<QRectF>(static_cast<QFontMetricsF*>(ptr)->boundingRect(*static_cast<QChar*>(ch))).width(), static_cast<QRectF>(static_cast<QFontMetricsF*>(ptr)->boundingRect(*static_cast<QChar*>(ch))).height());
}

void* QFontMetricsF_BoundingRect3(void* ptr, void* rect, int flags, char* text, int tabStops, int tabArray)
{
	return new QRectF(static_cast<QRectF>(static_cast<QFontMetricsF*>(ptr)->boundingRect(*static_cast<QRectF*>(rect), flags, QString(text), tabStops, &tabArray)).x(), static_cast<QRectF>(static_cast<QFontMetricsF*>(ptr)->boundingRect(*static_cast<QRectF*>(rect), flags, QString(text), tabStops, &tabArray)).y(), static_cast<QRectF>(static_cast<QFontMetricsF*>(ptr)->boundingRect(*static_cast<QRectF*>(rect), flags, QString(text), tabStops, &tabArray)).width(), static_cast<QRectF>(static_cast<QFontMetricsF*>(ptr)->boundingRect(*static_cast<QRectF*>(rect), flags, QString(text), tabStops, &tabArray)).height());
}

void* QFontMetricsF_BoundingRect(void* ptr, char* text)
{
	return new QRectF(static_cast<QRectF>(static_cast<QFontMetricsF*>(ptr)->boundingRect(QString(text))).x(), static_cast<QRectF>(static_cast<QFontMetricsF*>(ptr)->boundingRect(QString(text))).y(), static_cast<QRectF>(static_cast<QFontMetricsF*>(ptr)->boundingRect(QString(text))).width(), static_cast<QRectF>(static_cast<QFontMetricsF*>(ptr)->boundingRect(QString(text))).height());
}

double QFontMetricsF_Descent(void* ptr)
{
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->descent());
}

char* QFontMetricsF_ElidedText(void* ptr, char* text, int mode, double width, int flags)
{
	return static_cast<QFontMetricsF*>(ptr)->elidedText(QString(text), static_cast<Qt::TextElideMode>(mode), static_cast<double>(width), flags).toUtf8().data();
}

double QFontMetricsF_Height(void* ptr)
{
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->height());
}

int QFontMetricsF_InFont(void* ptr, void* ch)
{
	return static_cast<QFontMetricsF*>(ptr)->inFont(*static_cast<QChar*>(ch));
}

double QFontMetricsF_Leading(void* ptr)
{
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->leading());
}

double QFontMetricsF_LeftBearing(void* ptr, void* ch)
{
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->leftBearing(*static_cast<QChar*>(ch)));
}

double QFontMetricsF_LineSpacing(void* ptr)
{
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->lineSpacing());
}

double QFontMetricsF_LineWidth(void* ptr)
{
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->lineWidth());
}

double QFontMetricsF_MaxWidth(void* ptr)
{
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->maxWidth());
}

double QFontMetricsF_MinLeftBearing(void* ptr)
{
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->minLeftBearing());
}

double QFontMetricsF_MinRightBearing(void* ptr)
{
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->minRightBearing());
}

double QFontMetricsF_OverlinePos(void* ptr)
{
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->overlinePos());
}

double QFontMetricsF_RightBearing(void* ptr, void* ch)
{
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->rightBearing(*static_cast<QChar*>(ch)));
}

void* QFontMetricsF_Size(void* ptr, int flags, char* text, int tabStops, int tabArray)
{
	return new QSizeF(static_cast<QSizeF>(static_cast<QFontMetricsF*>(ptr)->size(flags, QString(text), tabStops, &tabArray)).width(), static_cast<QSizeF>(static_cast<QFontMetricsF*>(ptr)->size(flags, QString(text), tabStops, &tabArray)).height());
}

double QFontMetricsF_StrikeOutPos(void* ptr)
{
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->strikeOutPos());
}

void QFontMetricsF_Swap(void* ptr, void* other)
{
	static_cast<QFontMetricsF*>(ptr)->swap(*static_cast<QFontMetricsF*>(other));
}

void* QFontMetricsF_TightBoundingRect(void* ptr, char* text)
{
	return new QRectF(static_cast<QRectF>(static_cast<QFontMetricsF*>(ptr)->tightBoundingRect(QString(text))).x(), static_cast<QRectF>(static_cast<QFontMetricsF*>(ptr)->tightBoundingRect(QString(text))).y(), static_cast<QRectF>(static_cast<QFontMetricsF*>(ptr)->tightBoundingRect(QString(text))).width(), static_cast<QRectF>(static_cast<QFontMetricsF*>(ptr)->tightBoundingRect(QString(text))).height());
}

double QFontMetricsF_UnderlinePos(void* ptr)
{
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->underlinePos());
}

double QFontMetricsF_Width2(void* ptr, void* ch)
{
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->width(*static_cast<QChar*>(ch)));
}

double QFontMetricsF_Width(void* ptr, char* text)
{
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->width(QString(text)));
}

double QFontMetricsF_XHeight(void* ptr)
{
	return static_cast<double>(static_cast<QFontMetricsF*>(ptr)->xHeight());
}

void QFontMetricsF_DestroyQFontMetricsF(void* ptr)
{
	static_cast<QFontMetricsF*>(ptr)->~QFontMetricsF();
}

class MyQGenericPlugin: public QGenericPlugin
{
public:
	MyQGenericPlugin(QObject *parent) : QGenericPlugin(parent) {};
	QObject * create(const QString & key, const QString & specification) { return static_cast<QObject*>(callbackQGenericPlugin_Create(this, this->objectName().toUtf8().data(), key.toUtf8().data(), specification.toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQGenericPlugin_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQGenericPlugin_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGenericPlugin_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQGenericPlugin_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQGenericPlugin_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGenericPlugin_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQGenericPlugin_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGenericPlugin_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGenericPlugin_MetaObject(const_cast<MyQGenericPlugin*>(this), this->objectName().toUtf8().data())); };
};

void* QGenericPlugin_NewQGenericPlugin(void* parent)
{
	return new MyQGenericPlugin(static_cast<QObject*>(parent));
}

void* QGenericPlugin_Create(void* ptr, char* key, char* specification)
{
	return static_cast<QGenericPlugin*>(ptr)->create(QString(key), QString(specification));
}

void QGenericPlugin_DestroyQGenericPlugin(void* ptr)
{
	static_cast<QGenericPlugin*>(ptr)->~QGenericPlugin();
}

void QGenericPlugin_TimerEvent(void* ptr, void* event)
{
	static_cast<QGenericPlugin*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGenericPlugin_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGenericPlugin*>(ptr)->QGenericPlugin::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGenericPlugin_ChildEvent(void* ptr, void* event)
{
	static_cast<QGenericPlugin*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGenericPlugin_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QGenericPlugin*>(ptr)->QGenericPlugin::childEvent(static_cast<QChildEvent*>(event));
}

void QGenericPlugin_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QGenericPlugin*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGenericPlugin_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGenericPlugin*>(ptr)->QGenericPlugin::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGenericPlugin_CustomEvent(void* ptr, void* event)
{
	static_cast<QGenericPlugin*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGenericPlugin_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QGenericPlugin*>(ptr)->QGenericPlugin::customEvent(static_cast<QEvent*>(event));
}

void QGenericPlugin_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGenericPlugin*>(ptr), "deleteLater");
}

void QGenericPlugin_DeleteLaterDefault(void* ptr)
{
	static_cast<QGenericPlugin*>(ptr)->QGenericPlugin::deleteLater();
}

void QGenericPlugin_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QGenericPlugin*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGenericPlugin_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGenericPlugin*>(ptr)->QGenericPlugin::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QGenericPlugin_Event(void* ptr, void* e)
{
	return static_cast<QGenericPlugin*>(ptr)->event(static_cast<QEvent*>(e));
}

int QGenericPlugin_EventDefault(void* ptr, void* e)
{
	return static_cast<QGenericPlugin*>(ptr)->QGenericPlugin::event(static_cast<QEvent*>(e));
}

int QGenericPlugin_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGenericPlugin*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QGenericPlugin_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGenericPlugin*>(ptr)->QGenericPlugin::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGenericPlugin_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGenericPlugin*>(ptr)->metaObject());
}

void* QGenericPlugin_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGenericPlugin*>(ptr)->QGenericPlugin::metaObject());
}

void* QGenericPluginFactory_QGenericPluginFactory_Create(char* key, char* specification)
{
	return QGenericPluginFactory::create(QString(key), QString(specification));
}

char* QGenericPluginFactory_QGenericPluginFactory_Keys()
{
	return QGenericPluginFactory::keys().join("|").toUtf8().data();
}

void* QGlyphRun_NewQGlyphRun()
{
	return new QGlyphRun();
}

void* QGlyphRun_NewQGlyphRun2(void* other)
{
	return new QGlyphRun(*static_cast<QGlyphRun*>(other));
}

void* QGlyphRun_BoundingRect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QGlyphRun*>(ptr)->boundingRect()).x(), static_cast<QRectF>(static_cast<QGlyphRun*>(ptr)->boundingRect()).y(), static_cast<QRectF>(static_cast<QGlyphRun*>(ptr)->boundingRect()).width(), static_cast<QRectF>(static_cast<QGlyphRun*>(ptr)->boundingRect()).height());
}

void QGlyphRun_Clear(void* ptr)
{
	static_cast<QGlyphRun*>(ptr)->clear();
}

int QGlyphRun_Flags(void* ptr)
{
	return static_cast<QGlyphRun*>(ptr)->flags();
}

int QGlyphRun_IsEmpty(void* ptr)
{
	return static_cast<QGlyphRun*>(ptr)->isEmpty();
}

int QGlyphRun_IsRightToLeft(void* ptr)
{
	return static_cast<QGlyphRun*>(ptr)->isRightToLeft();
}

int QGlyphRun_Overline(void* ptr)
{
	return static_cast<QGlyphRun*>(ptr)->overline();
}

void* QGlyphRun_RawFont(void* ptr)
{
	return new QRawFont(static_cast<QGlyphRun*>(ptr)->rawFont());
}

void QGlyphRun_SetBoundingRect(void* ptr, void* boundingRect)
{
	static_cast<QGlyphRun*>(ptr)->setBoundingRect(*static_cast<QRectF*>(boundingRect));
}

void QGlyphRun_SetFlag(void* ptr, int flag, int enabled)
{
	static_cast<QGlyphRun*>(ptr)->setFlag(static_cast<QGlyphRun::GlyphRunFlag>(flag), enabled != 0);
}

void QGlyphRun_SetFlags(void* ptr, int flags)
{
	static_cast<QGlyphRun*>(ptr)->setFlags(static_cast<QGlyphRun::GlyphRunFlag>(flags));
}

void QGlyphRun_SetOverline(void* ptr, int overline)
{
	static_cast<QGlyphRun*>(ptr)->setOverline(overline != 0);
}

void QGlyphRun_SetRawFont(void* ptr, void* rawFont)
{
	static_cast<QGlyphRun*>(ptr)->setRawFont(*static_cast<QRawFont*>(rawFont));
}

void QGlyphRun_SetRightToLeft(void* ptr, int rightToLeft)
{
	static_cast<QGlyphRun*>(ptr)->setRightToLeft(rightToLeft != 0);
}

void QGlyphRun_SetStrikeOut(void* ptr, int strikeOut)
{
	static_cast<QGlyphRun*>(ptr)->setStrikeOut(strikeOut != 0);
}

void QGlyphRun_SetUnderline(void* ptr, int underline)
{
	static_cast<QGlyphRun*>(ptr)->setUnderline(underline != 0);
}

int QGlyphRun_StrikeOut(void* ptr)
{
	return static_cast<QGlyphRun*>(ptr)->strikeOut();
}

void QGlyphRun_Swap(void* ptr, void* other)
{
	static_cast<QGlyphRun*>(ptr)->swap(*static_cast<QGlyphRun*>(other));
}

int QGlyphRun_Underline(void* ptr)
{
	return static_cast<QGlyphRun*>(ptr)->underline();
}

void QGlyphRun_DestroyQGlyphRun(void* ptr)
{
	static_cast<QGlyphRun*>(ptr)->~QGlyphRun();
}

void QGradient_SetColorAt(void* ptr, double position, void* color)
{
	static_cast<QGradient*>(ptr)->setColorAt(static_cast<double>(position), *static_cast<QColor*>(color));
}

int QGradient_CoordinateMode(void* ptr)
{
	return static_cast<QGradient*>(ptr)->coordinateMode();
}

void QGradient_SetCoordinateMode(void* ptr, int mode)
{
	static_cast<QGradient*>(ptr)->setCoordinateMode(static_cast<QGradient::CoordinateMode>(mode));
}

void QGradient_SetSpread(void* ptr, int method)
{
	static_cast<QGradient*>(ptr)->setSpread(static_cast<QGradient::Spread>(method));
}

int QGradient_Spread(void* ptr)
{
	return static_cast<QGradient*>(ptr)->spread();
}

int QGradient_Type(void* ptr)
{
	return static_cast<QGradient*>(ptr)->type();
}

class MyQGuiApplication: public QGuiApplication
{
public:
	MyQGuiApplication(int &argc, char **argv) : QGuiApplication(argc, argv) {};
	void Signal_ApplicationStateChanged(Qt::ApplicationState state) { callbackQGuiApplication_ApplicationStateChanged(this, this->objectName().toUtf8().data(), state); };
	bool event(QEvent * e) { return callbackQGuiApplication_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	void Signal_FocusObjectChanged(QObject * focusObject) { callbackQGuiApplication_FocusObjectChanged(this, this->objectName().toUtf8().data(), focusObject); };
	void Signal_FocusWindowChanged(QWindow * focusWindow) { callbackQGuiApplication_FocusWindowChanged(this, this->objectName().toUtf8().data(), focusWindow); };
	void Signal_FontDatabaseChanged() { callbackQGuiApplication_FontDatabaseChanged(this, this->objectName().toUtf8().data()); };
	void Signal_LastWindowClosed() { callbackQGuiApplication_LastWindowClosed(this, this->objectName().toUtf8().data()); };
	void Signal_LayoutDirectionChanged(Qt::LayoutDirection direction) { callbackQGuiApplication_LayoutDirectionChanged(this, this->objectName().toUtf8().data(), direction); };
	void Signal_PaletteChanged(const QPalette & palette) { callbackQGuiApplication_PaletteChanged(this, this->objectName().toUtf8().data(), new QPalette(palette)); };
	void Signal_ScreenAdded(QScreen * screen) { callbackQGuiApplication_ScreenAdded(this, this->objectName().toUtf8().data(), screen); };
	void Signal_ScreenRemoved(QScreen * screen) { callbackQGuiApplication_ScreenRemoved(this, this->objectName().toUtf8().data(), screen); };
	void quit() { callbackQGuiApplication_Quit(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQGuiApplication_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQGuiApplication_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGuiApplication_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQGuiApplication_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQGuiApplication_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGuiApplication_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGuiApplication_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQGuiApplication_MetaObject(const_cast<MyQGuiApplication*>(this), this->objectName().toUtf8().data())); };
};

char* QGuiApplication_QGuiApplication_ApplicationDisplayName()
{
	return QGuiApplication::applicationDisplayName().toUtf8().data();
}

int QGuiApplication_QGuiApplication_ApplicationState()
{
	return QGuiApplication::applicationState();
}

int QGuiApplication_IsSavingSession(void* ptr)
{
	return static_cast<QGuiApplication*>(ptr)->isSavingSession();
}

int QGuiApplication_IsSessionRestored(void* ptr)
{
	return static_cast<QGuiApplication*>(ptr)->isSessionRestored();
}

int QGuiApplication_QGuiApplication_LayoutDirection()
{
	return QGuiApplication::layoutDirection();
}

void* QGuiApplication_QGuiApplication_OverrideCursor()
{
	return QGuiApplication::overrideCursor();
}

char* QGuiApplication_QGuiApplication_PlatformName()
{
	return QGuiApplication::platformName().toUtf8().data();
}

int QGuiApplication_QGuiApplication_QueryKeyboardModifiers()
{
	return QGuiApplication::queryKeyboardModifiers();
}

int QGuiApplication_QGuiApplication_QuitOnLastWindowClosed()
{
	return QGuiApplication::quitOnLastWindowClosed();
}

void QGuiApplication_QGuiApplication_RestoreOverrideCursor()
{
	QGuiApplication::restoreOverrideCursor();
}

char* QGuiApplication_SessionId(void* ptr)
{
	return static_cast<QGuiApplication*>(ptr)->sessionId().toUtf8().data();
}

char* QGuiApplication_SessionKey(void* ptr)
{
	return static_cast<QGuiApplication*>(ptr)->sessionKey().toUtf8().data();
}

void QGuiApplication_QGuiApplication_SetApplicationDisplayName(char* name)
{
	QGuiApplication::setApplicationDisplayName(QString(name));
}

void QGuiApplication_QGuiApplication_SetLayoutDirection(int direction)
{
	QGuiApplication::setLayoutDirection(static_cast<Qt::LayoutDirection>(direction));
}

void QGuiApplication_QGuiApplication_SetOverrideCursor(void* cursor)
{
	QGuiApplication::setOverrideCursor(*static_cast<QCursor*>(cursor));
}

void QGuiApplication_QGuiApplication_SetQuitOnLastWindowClosed(int quit)
{
	QGuiApplication::setQuitOnLastWindowClosed(quit != 0);
}

void QGuiApplication_QGuiApplication_SetWindowIcon(void* icon)
{
	QGuiApplication::setWindowIcon(*static_cast<QIcon*>(icon));
}

void* QGuiApplication_QGuiApplication_WindowIcon()
{
	return new QIcon(QGuiApplication::windowIcon());
}

void* QGuiApplication_NewQGuiApplication(int argc, char* argv)
{
	QList<QByteArray> aList = QByteArray(argv).split('|');
	char *argvs[argc];
	static int argcs = argc;
	for (int i = 0; i < argc; i++)
		argvs[i] = aList[i].data();

	return new MyQGuiApplication(argcs, argvs);
}

void QGuiApplication_ConnectApplicationStateChanged(void* ptr)
{
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(Qt::ApplicationState)>(&QGuiApplication::applicationStateChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(Qt::ApplicationState)>(&MyQGuiApplication::Signal_ApplicationStateChanged));
}

void QGuiApplication_DisconnectApplicationStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(Qt::ApplicationState)>(&QGuiApplication::applicationStateChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(Qt::ApplicationState)>(&MyQGuiApplication::Signal_ApplicationStateChanged));
}

void QGuiApplication_ApplicationStateChanged(void* ptr, int state)
{
	static_cast<QGuiApplication*>(ptr)->applicationStateChanged(static_cast<Qt::ApplicationState>(state));
}

void QGuiApplication_QGuiApplication_ChangeOverrideCursor(void* cursor)
{
	QGuiApplication::changeOverrideCursor(*static_cast<QCursor*>(cursor));
}

void* QGuiApplication_QGuiApplication_Clipboard()
{
	return QGuiApplication::clipboard();
}

int QGuiApplication_QGuiApplication_DesktopSettingsAware()
{
	return QGuiApplication::desktopSettingsAware();
}

double QGuiApplication_DevicePixelRatio(void* ptr)
{
	return static_cast<double>(static_cast<QGuiApplication*>(ptr)->devicePixelRatio());
}

int QGuiApplication_Event(void* ptr, void* e)
{
	return static_cast<QGuiApplication*>(ptr)->event(static_cast<QEvent*>(e));
}

int QGuiApplication_EventDefault(void* ptr, void* e)
{
	return static_cast<QGuiApplication*>(ptr)->QGuiApplication::event(static_cast<QEvent*>(e));
}

int QGuiApplication_QGuiApplication_Exec()
{
	return QGuiApplication::exec();
}

void* QGuiApplication_QGuiApplication_FocusObject()
{
	return QGuiApplication::focusObject();
}

void QGuiApplication_ConnectFocusObjectChanged(void* ptr)
{
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QObject *)>(&QGuiApplication::focusObjectChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QObject *)>(&MyQGuiApplication::Signal_FocusObjectChanged));
}

void QGuiApplication_DisconnectFocusObjectChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QObject *)>(&QGuiApplication::focusObjectChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QObject *)>(&MyQGuiApplication::Signal_FocusObjectChanged));
}

void QGuiApplication_FocusObjectChanged(void* ptr, void* focusObject)
{
	static_cast<QGuiApplication*>(ptr)->focusObjectChanged(static_cast<QObject*>(focusObject));
}

void* QGuiApplication_QGuiApplication_FocusWindow()
{
	return QGuiApplication::focusWindow();
}

void QGuiApplication_ConnectFocusWindowChanged(void* ptr)
{
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QWindow *)>(&QGuiApplication::focusWindowChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QWindow *)>(&MyQGuiApplication::Signal_FocusWindowChanged));
}

void QGuiApplication_DisconnectFocusWindowChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QWindow *)>(&QGuiApplication::focusWindowChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QWindow *)>(&MyQGuiApplication::Signal_FocusWindowChanged));
}

void QGuiApplication_FocusWindowChanged(void* ptr, void* focusWindow)
{
	static_cast<QGuiApplication*>(ptr)->focusWindowChanged(static_cast<QWindow*>(focusWindow));
}

void* QGuiApplication_QGuiApplication_Font()
{
	return new QFont(QGuiApplication::font());
}

void QGuiApplication_ConnectFontDatabaseChanged(void* ptr)
{
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)()>(&QGuiApplication::fontDatabaseChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)()>(&MyQGuiApplication::Signal_FontDatabaseChanged));
}

void QGuiApplication_DisconnectFontDatabaseChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)()>(&QGuiApplication::fontDatabaseChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)()>(&MyQGuiApplication::Signal_FontDatabaseChanged));
}

void QGuiApplication_FontDatabaseChanged(void* ptr)
{
	static_cast<QGuiApplication*>(ptr)->fontDatabaseChanged();
}

void* QGuiApplication_QGuiApplication_InputMethod()
{
	return QGuiApplication::inputMethod();
}

int QGuiApplication_QGuiApplication_IsLeftToRight()
{
	return QGuiApplication::isLeftToRight();
}

int QGuiApplication_QGuiApplication_IsRightToLeft()
{
	return QGuiApplication::isRightToLeft();
}

int QGuiApplication_QGuiApplication_KeyboardModifiers()
{
	return QGuiApplication::keyboardModifiers();
}

void QGuiApplication_ConnectLastWindowClosed(void* ptr)
{
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)()>(&QGuiApplication::lastWindowClosed), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)()>(&MyQGuiApplication::Signal_LastWindowClosed));
}

void QGuiApplication_DisconnectLastWindowClosed(void* ptr)
{
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)()>(&QGuiApplication::lastWindowClosed), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)()>(&MyQGuiApplication::Signal_LastWindowClosed));
}

void QGuiApplication_LastWindowClosed(void* ptr)
{
	static_cast<QGuiApplication*>(ptr)->lastWindowClosed();
}

void QGuiApplication_ConnectLayoutDirectionChanged(void* ptr)
{
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(Qt::LayoutDirection)>(&QGuiApplication::layoutDirectionChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(Qt::LayoutDirection)>(&MyQGuiApplication::Signal_LayoutDirectionChanged));
}

void QGuiApplication_DisconnectLayoutDirectionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(Qt::LayoutDirection)>(&QGuiApplication::layoutDirectionChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(Qt::LayoutDirection)>(&MyQGuiApplication::Signal_LayoutDirectionChanged));
}

void QGuiApplication_LayoutDirectionChanged(void* ptr, int direction)
{
	static_cast<QGuiApplication*>(ptr)->layoutDirectionChanged(static_cast<Qt::LayoutDirection>(direction));
}

void* QGuiApplication_QGuiApplication_ModalWindow()
{
	return QGuiApplication::modalWindow();
}

int QGuiApplication_QGuiApplication_MouseButtons()
{
	return QGuiApplication::mouseButtons();
}

void* QGuiApplication_QGuiApplication_Palette()
{
	return new QPalette(QGuiApplication::palette());
}

void QGuiApplication_ConnectPaletteChanged(void* ptr)
{
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(const QPalette &)>(&QGuiApplication::paletteChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(const QPalette &)>(&MyQGuiApplication::Signal_PaletteChanged));
}

void QGuiApplication_DisconnectPaletteChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(const QPalette &)>(&QGuiApplication::paletteChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(const QPalette &)>(&MyQGuiApplication::Signal_PaletteChanged));
}

void QGuiApplication_PaletteChanged(void* ptr, void* palette)
{
	static_cast<QGuiApplication*>(ptr)->paletteChanged(*static_cast<QPalette*>(palette));
}

void* QGuiApplication_QGuiApplication_PrimaryScreen()
{
	return QGuiApplication::primaryScreen();
}

void QGuiApplication_ConnectScreenAdded(void* ptr)
{
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QScreen *)>(&QGuiApplication::screenAdded), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QScreen *)>(&MyQGuiApplication::Signal_ScreenAdded));
}

void QGuiApplication_DisconnectScreenAdded(void* ptr)
{
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QScreen *)>(&QGuiApplication::screenAdded), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QScreen *)>(&MyQGuiApplication::Signal_ScreenAdded));
}

void QGuiApplication_ScreenAdded(void* ptr, void* screen)
{
	static_cast<QGuiApplication*>(ptr)->screenAdded(static_cast<QScreen*>(screen));
}

void QGuiApplication_ConnectScreenRemoved(void* ptr)
{
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QScreen *)>(&QGuiApplication::screenRemoved), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QScreen *)>(&MyQGuiApplication::Signal_ScreenRemoved));
}

void QGuiApplication_DisconnectScreenRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(QScreen *)>(&QGuiApplication::screenRemoved), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(QScreen *)>(&MyQGuiApplication::Signal_ScreenRemoved));
}

void QGuiApplication_ScreenRemoved(void* ptr, void* screen)
{
	static_cast<QGuiApplication*>(ptr)->screenRemoved(static_cast<QScreen*>(screen));
}

void QGuiApplication_QGuiApplication_SetDesktopSettingsAware(int on)
{
	QGuiApplication::setDesktopSettingsAware(on != 0);
}

void QGuiApplication_QGuiApplication_SetFont(void* font)
{
	QGuiApplication::setFont(*static_cast<QFont*>(font));
}

void QGuiApplication_QGuiApplication_SetPalette(void* pal)
{
	QGuiApplication::setPalette(*static_cast<QPalette*>(pal));
}

void* QGuiApplication_QGuiApplication_StyleHints()
{
	return QGuiApplication::styleHints();
}

void QGuiApplication_QGuiApplication_Sync()
{
	QGuiApplication::sync();
}

void* QGuiApplication_QGuiApplication_TopLevelAt(void* pos)
{
	return QGuiApplication::topLevelAt(*static_cast<QPoint*>(pos));
}

void QGuiApplication_DestroyQGuiApplication(void* ptr)
{
	static_cast<QGuiApplication*>(ptr)->~QGuiApplication();
}

void QGuiApplication_Quit(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGuiApplication*>(ptr), "quit");
}

void QGuiApplication_QuitDefault(void* ptr)
{
	static_cast<QGuiApplication*>(ptr)->QGuiApplication::quit();
}

void QGuiApplication_TimerEvent(void* ptr, void* event)
{
	static_cast<QGuiApplication*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QGuiApplication_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QGuiApplication*>(ptr)->QGuiApplication::timerEvent(static_cast<QTimerEvent*>(event));
}

void QGuiApplication_ChildEvent(void* ptr, void* event)
{
	static_cast<QGuiApplication*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QGuiApplication_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QGuiApplication*>(ptr)->QGuiApplication::childEvent(static_cast<QChildEvent*>(event));
}

void QGuiApplication_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QGuiApplication*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGuiApplication_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGuiApplication*>(ptr)->QGuiApplication::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGuiApplication_CustomEvent(void* ptr, void* event)
{
	static_cast<QGuiApplication*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QGuiApplication_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QGuiApplication*>(ptr)->QGuiApplication::customEvent(static_cast<QEvent*>(event));
}

void QGuiApplication_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QGuiApplication*>(ptr), "deleteLater");
}

void QGuiApplication_DeleteLaterDefault(void* ptr)
{
	static_cast<QGuiApplication*>(ptr)->QGuiApplication::deleteLater();
}

void QGuiApplication_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QGuiApplication*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGuiApplication_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QGuiApplication*>(ptr)->QGuiApplication::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QGuiApplication_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QGuiApplication*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QGuiApplication_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QGuiApplication*>(ptr)->QGuiApplication::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QGuiApplication_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGuiApplication*>(ptr)->metaObject());
}

void* QGuiApplication_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QGuiApplication*>(ptr)->QGuiApplication::metaObject());
}

void* QHelpEvent_NewQHelpEvent(int ty, void* pos, void* globalPos)
{
	return new QHelpEvent(static_cast<QEvent::Type>(ty), *static_cast<QPoint*>(pos), *static_cast<QPoint*>(globalPos));
}

void* QHelpEvent_GlobalPos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QHelpEvent*>(ptr)->globalPos()).x(), static_cast<QPoint>(static_cast<QHelpEvent*>(ptr)->globalPos()).y());
}

int QHelpEvent_GlobalX(void* ptr)
{
	return static_cast<QHelpEvent*>(ptr)->globalX();
}

int QHelpEvent_GlobalY(void* ptr)
{
	return static_cast<QHelpEvent*>(ptr)->globalY();
}

void* QHelpEvent_Pos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QHelpEvent*>(ptr)->pos()).x(), static_cast<QPoint>(static_cast<QHelpEvent*>(ptr)->pos()).y());
}

int QHelpEvent_X(void* ptr)
{
	return static_cast<QHelpEvent*>(ptr)->x();
}

int QHelpEvent_Y(void* ptr)
{
	return static_cast<QHelpEvent*>(ptr)->y();
}

void* QHideEvent_NewQHideEvent()
{
	return new QHideEvent();
}

void* QHoverEvent_NewQHoverEvent(int ty, void* pos, void* oldPos, int modifiers)
{
	return new QHoverEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(pos), *static_cast<QPointF*>(oldPos), static_cast<Qt::KeyboardModifier>(modifiers));
}

void* QHoverEvent_OldPos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QHoverEvent*>(ptr)->oldPos()).x(), static_cast<QPoint>(static_cast<QHoverEvent*>(ptr)->oldPos()).y());
}

void* QHoverEvent_OldPosF(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QHoverEvent*>(ptr)->oldPosF()).x(), static_cast<QPointF>(static_cast<QHoverEvent*>(ptr)->oldPosF()).y());
}

void* QHoverEvent_Pos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QHoverEvent*>(ptr)->pos()).x(), static_cast<QPoint>(static_cast<QHoverEvent*>(ptr)->pos()).y());
}

void* QHoverEvent_PosF(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QHoverEvent*>(ptr)->posF()).x(), static_cast<QPointF>(static_cast<QHoverEvent*>(ptr)->posF()).y());
}

void* QIcon_NewQIcon()
{
	return new QIcon();
}

void* QIcon_NewQIcon4(void* other)
{
	return new QIcon(*static_cast<QIcon*>(other));
}

void* QIcon_NewQIcon6(void* engine)
{
	return new QIcon(static_cast<QIconEngine*>(engine));
}

void* QIcon_NewQIcon3(void* other)
{
	return new QIcon(*static_cast<QIcon*>(other));
}

void* QIcon_NewQIcon2(void* pixmap)
{
	return new QIcon(*static_cast<QPixmap*>(pixmap));
}

void* QIcon_NewQIcon5(char* fileName)
{
	return new QIcon(QString(fileName));
}

void* QIcon_ActualSize2(void* ptr, void* window, void* size, int mode, int state)
{
	return new QSize(static_cast<QSize>(static_cast<QIcon*>(ptr)->actualSize(static_cast<QWindow*>(window), *static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state))).width(), static_cast<QSize>(static_cast<QIcon*>(ptr)->actualSize(static_cast<QWindow*>(window), *static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state))).height());
}

void* QIcon_ActualSize(void* ptr, void* size, int mode, int state)
{
	return new QSize(static_cast<QSize>(static_cast<QIcon*>(ptr)->actualSize(*static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state))).width(), static_cast<QSize>(static_cast<QIcon*>(ptr)->actualSize(*static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state))).height());
}

void QIcon_AddFile(void* ptr, char* fileName, void* size, int mode, int state)
{
	static_cast<QIcon*>(ptr)->addFile(QString(fileName), *static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void QIcon_AddPixmap(void* ptr, void* pixmap, int mode, int state)
{
	static_cast<QIcon*>(ptr)->addPixmap(*static_cast<QPixmap*>(pixmap), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

long long QIcon_CacheKey(void* ptr)
{
	return static_cast<long long>(static_cast<QIcon*>(ptr)->cacheKey());
}

void* QIcon_QIcon_FromTheme(char* name, void* fallback)
{
	return new QIcon(QIcon::fromTheme(QString(name), *static_cast<QIcon*>(fallback)));
}

int QIcon_QIcon_HasThemeIcon(char* name)
{
	return QIcon::hasThemeIcon(QString(name));
}

int QIcon_IsNull(void* ptr)
{
	return static_cast<QIcon*>(ptr)->isNull();
}

char* QIcon_Name(void* ptr)
{
	return static_cast<QIcon*>(ptr)->name().toUtf8().data();
}

void QIcon_Paint(void* ptr, void* painter, void* rect, int alignment, int mode, int state)
{
	static_cast<QIcon*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QRect*>(rect), static_cast<Qt::AlignmentFlag>(alignment), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void QIcon_Paint2(void* ptr, void* painter, int x, int y, int w, int h, int alignment, int mode, int state)
{
	static_cast<QIcon*>(ptr)->paint(static_cast<QPainter*>(painter), x, y, w, h, static_cast<Qt::AlignmentFlag>(alignment), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void* QIcon_Pixmap2(void* ptr, void* window, void* size, int mode, int state)
{
	return new QPixmap(static_cast<QIcon*>(ptr)->pixmap(static_cast<QWindow*>(window), *static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state)));
}

void* QIcon_Pixmap(void* ptr, void* size, int mode, int state)
{
	return new QPixmap(static_cast<QIcon*>(ptr)->pixmap(*static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state)));
}

void* QIcon_Pixmap4(void* ptr, int extent, int mode, int state)
{
	return new QPixmap(static_cast<QIcon*>(ptr)->pixmap(extent, static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state)));
}

void* QIcon_Pixmap3(void* ptr, int w, int h, int mode, int state)
{
	return new QPixmap(static_cast<QIcon*>(ptr)->pixmap(w, h, static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state)));
}

void QIcon_QIcon_SetThemeName(char* name)
{
	QIcon::setThemeName(QString(name));
}

void QIcon_QIcon_SetThemeSearchPaths(char* paths)
{
	QIcon::setThemeSearchPaths(QString(paths).split("|", QString::SkipEmptyParts));
}

void QIcon_Swap(void* ptr, void* other)
{
	static_cast<QIcon*>(ptr)->swap(*static_cast<QIcon*>(other));
}

char* QIcon_QIcon_ThemeName()
{
	return QIcon::themeName().toUtf8().data();
}

char* QIcon_QIcon_ThemeSearchPaths()
{
	return QIcon::themeSearchPaths().join("|").toUtf8().data();
}

void QIcon_DestroyQIcon(void* ptr)
{
	static_cast<QIcon*>(ptr)->~QIcon();
}

void* QIconDragEvent_NewQIconDragEvent()
{
	return new QIconDragEvent();
}

class MyQIconEngine: public QIconEngine
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	QSize actualSize(const QSize & size, QIcon::Mode mode, QIcon::State state) { return *static_cast<QSize*>(callbackQIconEngine_ActualSize(this, this->objectNameAbs().toUtf8().data(), new QSize(static_cast<QSize>(size).width(), static_cast<QSize>(size).height()), mode, state)); };
	void addFile(const QString & fileName, const QSize & size, QIcon::Mode mode, QIcon::State state) { callbackQIconEngine_AddFile(this, this->objectNameAbs().toUtf8().data(), fileName.toUtf8().data(), new QSize(static_cast<QSize>(size).width(), static_cast<QSize>(size).height()), mode, state); };
	void addPixmap(const QPixmap & pixmap, QIcon::Mode mode, QIcon::State state) { callbackQIconEngine_AddPixmap(this, this->objectNameAbs().toUtf8().data(), new QPixmap(pixmap), mode, state); };
	QIconEngine * clone() const { return static_cast<QIconEngine*>(callbackQIconEngine_Clone(const_cast<MyQIconEngine*>(this), this->objectNameAbs().toUtf8().data())); };
	QString iconName() const { return QString(callbackQIconEngine_IconName(const_cast<MyQIconEngine*>(this), this->objectNameAbs().toUtf8().data())); };
	QString key() const { return QString(callbackQIconEngine_Key(const_cast<MyQIconEngine*>(this), this->objectNameAbs().toUtf8().data())); };
	void paint(QPainter * painter, const QRect & rect, QIcon::Mode mode, QIcon::State state) { callbackQIconEngine_Paint(this, this->objectNameAbs().toUtf8().data(), painter, new QRect(static_cast<QRect>(rect).x(), static_cast<QRect>(rect).y(), static_cast<QRect>(rect).width(), static_cast<QRect>(rect).height()), mode, state); };
	QPixmap pixmap(const QSize & size, QIcon::Mode mode, QIcon::State state) { return *static_cast<QPixmap*>(callbackQIconEngine_Pixmap(this, this->objectNameAbs().toUtf8().data(), new QSize(static_cast<QSize>(size).width(), static_cast<QSize>(size).height()), mode, state)); };
	void virtual_hook(int id, void * data) { callbackQIconEngine_Virtual_hook(this, this->objectNameAbs().toUtf8().data(), id, data); };
};

void* QIconEngine_ActualSize(void* ptr, void* size, int mode, int state)
{
	return new QSize(static_cast<QSize>(static_cast<QIconEngine*>(ptr)->actualSize(*static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state))).width(), static_cast<QSize>(static_cast<QIconEngine*>(ptr)->actualSize(*static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state))).height());
}

void* QIconEngine_ActualSizeDefault(void* ptr, void* size, int mode, int state)
{
	return new QSize(static_cast<QSize>(static_cast<QIconEngine*>(ptr)->QIconEngine::actualSize(*static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state))).width(), static_cast<QSize>(static_cast<QIconEngine*>(ptr)->QIconEngine::actualSize(*static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state))).height());
}

void QIconEngine_AddFile(void* ptr, char* fileName, void* size, int mode, int state)
{
	static_cast<QIconEngine*>(ptr)->addFile(QString(fileName), *static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void QIconEngine_AddFileDefault(void* ptr, char* fileName, void* size, int mode, int state)
{
	static_cast<QIconEngine*>(ptr)->QIconEngine::addFile(QString(fileName), *static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void QIconEngine_AddPixmap(void* ptr, void* pixmap, int mode, int state)
{
	static_cast<QIconEngine*>(ptr)->addPixmap(*static_cast<QPixmap*>(pixmap), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void QIconEngine_AddPixmapDefault(void* ptr, void* pixmap, int mode, int state)
{
	static_cast<QIconEngine*>(ptr)->QIconEngine::addPixmap(*static_cast<QPixmap*>(pixmap), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void* QIconEngine_Clone(void* ptr)
{
	return static_cast<QIconEngine*>(ptr)->clone();
}

char* QIconEngine_IconName(void* ptr)
{
	return static_cast<QIconEngine*>(ptr)->iconName().toUtf8().data();
}

char* QIconEngine_IconNameDefault(void* ptr)
{
	return static_cast<QIconEngine*>(ptr)->QIconEngine::iconName().toUtf8().data();
}

char* QIconEngine_Key(void* ptr)
{
	return static_cast<QIconEngine*>(ptr)->key().toUtf8().data();
}

char* QIconEngine_KeyDefault(void* ptr)
{
	return static_cast<QIconEngine*>(ptr)->QIconEngine::key().toUtf8().data();
}

void QIconEngine_Paint(void* ptr, void* painter, void* rect, int mode, int state)
{
	static_cast<QIconEngine*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QRect*>(rect), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void* QIconEngine_Pixmap(void* ptr, void* size, int mode, int state)
{
	return new QPixmap(static_cast<QIconEngine*>(ptr)->pixmap(*static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state)));
}

void* QIconEngine_PixmapDefault(void* ptr, void* size, int mode, int state)
{
	return new QPixmap(static_cast<QIconEngine*>(ptr)->QIconEngine::pixmap(*static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state)));
}

void QIconEngine_Virtual_hook(void* ptr, int id, void* data)
{
	static_cast<QIconEngine*>(ptr)->virtual_hook(id, data);
}

void QIconEngine_Virtual_hookDefault(void* ptr, int id, void* data)
{
	static_cast<QIconEngine*>(ptr)->QIconEngine::virtual_hook(id, data);
}

void QIconEngine_DestroyQIconEngine(void* ptr)
{
	static_cast<QIconEngine*>(ptr)->~QIconEngine();
}

char* QIconEngine_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQIconEngine*>(static_cast<QIconEngine*>(ptr))) {
		return static_cast<MyQIconEngine*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QIconEngine_BASE").toUtf8().data();
}

void QIconEngine_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQIconEngine*>(static_cast<QIconEngine*>(ptr))) {
		static_cast<MyQIconEngine*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQIconEnginePlugin: public QIconEnginePlugin
{
public:
	MyQIconEnginePlugin(QObject *parent) : QIconEnginePlugin(parent) {};
	QIconEngine * create(const QString & filename) { return static_cast<QIconEngine*>(callbackQIconEnginePlugin_Create(this, this->objectName().toUtf8().data(), filename.toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQIconEnginePlugin_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQIconEnginePlugin_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQIconEnginePlugin_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQIconEnginePlugin_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQIconEnginePlugin_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQIconEnginePlugin_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQIconEnginePlugin_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQIconEnginePlugin_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQIconEnginePlugin_MetaObject(const_cast<MyQIconEnginePlugin*>(this), this->objectName().toUtf8().data())); };
};

void* QIconEnginePlugin_NewQIconEnginePlugin(void* parent)
{
	return new MyQIconEnginePlugin(static_cast<QObject*>(parent));
}

void* QIconEnginePlugin_Create(void* ptr, char* filename)
{
	return static_cast<QIconEnginePlugin*>(ptr)->create(QString(filename));
}

void QIconEnginePlugin_DestroyQIconEnginePlugin(void* ptr)
{
	static_cast<QIconEnginePlugin*>(ptr)->~QIconEnginePlugin();
}

void QIconEnginePlugin_TimerEvent(void* ptr, void* event)
{
	static_cast<QIconEnginePlugin*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QIconEnginePlugin_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QIconEnginePlugin*>(ptr)->QIconEnginePlugin::timerEvent(static_cast<QTimerEvent*>(event));
}

void QIconEnginePlugin_ChildEvent(void* ptr, void* event)
{
	static_cast<QIconEnginePlugin*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QIconEnginePlugin_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QIconEnginePlugin*>(ptr)->QIconEnginePlugin::childEvent(static_cast<QChildEvent*>(event));
}

void QIconEnginePlugin_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QIconEnginePlugin*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QIconEnginePlugin_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QIconEnginePlugin*>(ptr)->QIconEnginePlugin::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QIconEnginePlugin_CustomEvent(void* ptr, void* event)
{
	static_cast<QIconEnginePlugin*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QIconEnginePlugin_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QIconEnginePlugin*>(ptr)->QIconEnginePlugin::customEvent(static_cast<QEvent*>(event));
}

void QIconEnginePlugin_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QIconEnginePlugin*>(ptr), "deleteLater");
}

void QIconEnginePlugin_DeleteLaterDefault(void* ptr)
{
	static_cast<QIconEnginePlugin*>(ptr)->QIconEnginePlugin::deleteLater();
}

void QIconEnginePlugin_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QIconEnginePlugin*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QIconEnginePlugin_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QIconEnginePlugin*>(ptr)->QIconEnginePlugin::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QIconEnginePlugin_Event(void* ptr, void* e)
{
	return static_cast<QIconEnginePlugin*>(ptr)->event(static_cast<QEvent*>(e));
}

int QIconEnginePlugin_EventDefault(void* ptr, void* e)
{
	return static_cast<QIconEnginePlugin*>(ptr)->QIconEnginePlugin::event(static_cast<QEvent*>(e));
}

int QIconEnginePlugin_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QIconEnginePlugin*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QIconEnginePlugin_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QIconEnginePlugin*>(ptr)->QIconEnginePlugin::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QIconEnginePlugin_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QIconEnginePlugin*>(ptr)->metaObject());
}

void* QIconEnginePlugin_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QIconEnginePlugin*>(ptr)->QIconEnginePlugin::metaObject());
}

int QImage_ColorCount(void* ptr)
{
	return static_cast<QImage*>(ptr)->colorCount();
}

void* QImage_Copy(void* ptr, void* rectangle)
{
	return new QImage(static_cast<QImage*>(ptr)->copy(*static_cast<QRect*>(rectangle)));
}

void QImage_Fill2(void* ptr, int color)
{
	static_cast<QImage*>(ptr)->fill(static_cast<Qt::GlobalColor>(color));
}

void QImage_Fill3(void* ptr, void* color)
{
	static_cast<QImage*>(ptr)->fill(*static_cast<QColor*>(color));
}

int QImage_Height(void* ptr)
{
	return static_cast<QImage*>(ptr)->height();
}

int QImage_IsNull(void* ptr)
{
	return static_cast<QImage*>(ptr)->isNull();
}

void* QImage_Offset(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QImage*>(ptr)->offset()).x(), static_cast<QPoint>(static_cast<QImage*>(ptr)->offset()).y());
}

void* QImage_Rect(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QImage*>(ptr)->rect()).x(), static_cast<QRect>(static_cast<QImage*>(ptr)->rect()).y(), static_cast<QRect>(static_cast<QImage*>(ptr)->rect()).width(), static_cast<QRect>(static_cast<QImage*>(ptr)->rect()).height());
}

void* QImage_Scaled(void* ptr, void* size, int aspectRatioMode, int transformMode)
{
	return new QImage(static_cast<QImage*>(ptr)->scaled(*static_cast<QSize*>(size), static_cast<Qt::AspectRatioMode>(aspectRatioMode), static_cast<Qt::TransformationMode>(transformMode)));
}

void* QImage_ScaledToHeight(void* ptr, int height, int mode)
{
	return new QImage(static_cast<QImage*>(ptr)->scaledToHeight(height, static_cast<Qt::TransformationMode>(mode)));
}

void* QImage_ScaledToWidth(void* ptr, int width, int mode)
{
	return new QImage(static_cast<QImage*>(ptr)->scaledToWidth(width, static_cast<Qt::TransformationMode>(mode)));
}

void QImage_SetOffset(void* ptr, void* offset)
{
	static_cast<QImage*>(ptr)->setOffset(*static_cast<QPoint*>(offset));
}

void QImage_SetText(void* ptr, char* key, char* text)
{
	static_cast<QImage*>(ptr)->setText(QString(key), QString(text));
}

void* QImage_Size(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QImage*>(ptr)->size()).width(), static_cast<QSize>(static_cast<QImage*>(ptr)->size()).height());
}

int QImage_Width(void* ptr)
{
	return static_cast<QImage*>(ptr)->width();
}

void* QImage_NewQImage()
{
	return new QImage();
}

void* QImage_NewQImage11(void* other)
{
	return new QImage(*static_cast<QImage*>(other));
}

void* QImage_NewQImage10(void* image)
{
	return new QImage(*static_cast<QImage*>(image));
}

void* QImage_NewQImage2(void* size, int format)
{
	return new QImage(*static_cast<QSize*>(size), static_cast<QImage::Format>(format));
}

void* QImage_NewQImage9(char* fileName, char* format)
{
	return new QImage(QString(fileName), const_cast<const char*>(format));
}

void* QImage_NewQImage3(int width, int height, int format)
{
	return new QImage(width, height, static_cast<QImage::Format>(format));
}

int QImage_AllGray(void* ptr)
{
	return static_cast<QImage*>(ptr)->allGray();
}

int QImage_BitPlaneCount(void* ptr)
{
	return static_cast<QImage*>(ptr)->bitPlaneCount();
}

int QImage_ByteCount(void* ptr)
{
	return static_cast<QImage*>(ptr)->byteCount();
}

int QImage_BytesPerLine(void* ptr)
{
	return static_cast<QImage*>(ptr)->bytesPerLine();
}

long long QImage_CacheKey(void* ptr)
{
	return static_cast<long long>(static_cast<QImage*>(ptr)->cacheKey());
}

void* QImage_Copy2(void* ptr, int x, int y, int width, int height)
{
	return new QImage(static_cast<QImage*>(ptr)->copy(x, y, width, height));
}

void* QImage_CreateAlphaMask(void* ptr, int flags)
{
	return new QImage(static_cast<QImage*>(ptr)->createAlphaMask(static_cast<Qt::ImageConversionFlag>(flags)));
}

void* QImage_CreateHeuristicMask(void* ptr, int clipTight)
{
	return new QImage(static_cast<QImage*>(ptr)->createHeuristicMask(clipTight != 0));
}

int QImage_Depth(void* ptr)
{
	return static_cast<QImage*>(ptr)->depth();
}

double QImage_DevicePixelRatio(void* ptr)
{
	return static_cast<double>(static_cast<QImage*>(ptr)->devicePixelRatio());
}

int QImage_DotsPerMeterX(void* ptr)
{
	return static_cast<QImage*>(ptr)->dotsPerMeterX();
}

int QImage_DotsPerMeterY(void* ptr)
{
	return static_cast<QImage*>(ptr)->dotsPerMeterY();
}

int QImage_Format(void* ptr)
{
	return static_cast<QImage*>(ptr)->format();
}

void* QImage_QImage_FromData2(char* data, char* format)
{
	return new QImage(QImage::fromData(QByteArray(data), const_cast<const char*>(format)));
}

int QImage_HasAlphaChannel(void* ptr)
{
	return static_cast<QImage*>(ptr)->hasAlphaChannel();
}

void QImage_InvertPixels(void* ptr, int mode)
{
	static_cast<QImage*>(ptr)->invertPixels(static_cast<QImage::InvertMode>(mode));
}

int QImage_IsGrayscale(void* ptr)
{
	return static_cast<QImage*>(ptr)->isGrayscale();
}

int QImage_Load2(void* ptr, void* device, char* format)
{
	return static_cast<QImage*>(ptr)->load(static_cast<QIODevice*>(device), const_cast<const char*>(format));
}

int QImage_Load(void* ptr, char* fileName, char* format)
{
	return static_cast<QImage*>(ptr)->load(QString(fileName), const_cast<const char*>(format));
}

int QImage_LoadFromData2(void* ptr, char* data, char* format)
{
	return static_cast<QImage*>(ptr)->loadFromData(QByteArray(data), const_cast<const char*>(format));
}

int QImage_PixelIndex(void* ptr, void* position)
{
	return static_cast<QImage*>(ptr)->pixelIndex(*static_cast<QPoint*>(position));
}

int QImage_PixelIndex2(void* ptr, int x, int y)
{
	return static_cast<QImage*>(ptr)->pixelIndex(x, y);
}

int QImage_Save2(void* ptr, void* device, char* format, int quality)
{
	return static_cast<QImage*>(ptr)->save(static_cast<QIODevice*>(device), const_cast<const char*>(format), quality);
}

int QImage_Save(void* ptr, char* fileName, char* format, int quality)
{
	return static_cast<QImage*>(ptr)->save(QString(fileName), const_cast<const char*>(format), quality);
}

void* QImage_Scaled2(void* ptr, int width, int height, int aspectRatioMode, int transformMode)
{
	return new QImage(static_cast<QImage*>(ptr)->scaled(width, height, static_cast<Qt::AspectRatioMode>(aspectRatioMode), static_cast<Qt::TransformationMode>(transformMode)));
}

void QImage_SetColorCount(void* ptr, int colorCount)
{
	static_cast<QImage*>(ptr)->setColorCount(colorCount);
}

void QImage_SetDevicePixelRatio(void* ptr, double scaleFactor)
{
	static_cast<QImage*>(ptr)->setDevicePixelRatio(static_cast<double>(scaleFactor));
}

void QImage_SetDotsPerMeterX(void* ptr, int x)
{
	static_cast<QImage*>(ptr)->setDotsPerMeterX(x);
}

void QImage_SetDotsPerMeterY(void* ptr, int y)
{
	static_cast<QImage*>(ptr)->setDotsPerMeterY(y);
}

void* QImage_SmoothScaled(void* ptr, int w, int h)
{
	return new QImage(static_cast<QImage*>(ptr)->smoothScaled(w, h));
}

void QImage_Swap(void* ptr, void* other)
{
	static_cast<QImage*>(ptr)->swap(*static_cast<QImage*>(other));
}

char* QImage_Text(void* ptr, char* key)
{
	return static_cast<QImage*>(ptr)->text(QString(key)).toUtf8().data();
}

char* QImage_TextKeys(void* ptr)
{
	return static_cast<QImage*>(ptr)->textKeys().join("|").toUtf8().data();
}

int QImage_QImage_ToImageFormat(void* format)
{
	return QImage::toImageFormat(*static_cast<QPixelFormat*>(format));
}

void* QImage_Transformed2(void* ptr, void* matrix, int mode)
{
	return new QImage(static_cast<QImage*>(ptr)->transformed(*static_cast<QTransform*>(matrix), static_cast<Qt::TransformationMode>(mode)));
}

int QImage_Valid(void* ptr, void* pos)
{
	return static_cast<QImage*>(ptr)->valid(*static_cast<QPoint*>(pos));
}

int QImage_Valid2(void* ptr, int x, int y)
{
	return static_cast<QImage*>(ptr)->valid(x, y);
}

void QImage_DestroyQImage(void* ptr)
{
	static_cast<QImage*>(ptr)->~QImage();
}

void* QImage_ConvertToFormat(void* ptr, int format, int flags)
{
	return new QImage(static_cast<QImage*>(ptr)->convertToFormat(static_cast<QImage::Format>(format), static_cast<Qt::ImageConversionFlag>(flags)));
}

void* QImage_Mirrored(void* ptr, int horizontal, int vertical)
{
	return new QImage(static_cast<QImage*>(ptr)->mirrored(horizontal != 0, vertical != 0));
}

void* QImage_RgbSwapped(void* ptr)
{
	return new QImage(static_cast<QImage*>(ptr)->rgbSwapped());
}

int QImage_Metric(void* ptr, int metric)
{
	return static_cast<QImage*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

int QImage_MetricDefault(void* ptr, int metric)
{
	return static_cast<QImage*>(ptr)->QImage::metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

void* QImage_PaintEngine(void* ptr)
{
	return static_cast<QImage*>(ptr)->paintEngine();
}

void* QImage_PaintEngineDefault(void* ptr)
{
	return static_cast<QImage*>(ptr)->QImage::paintEngine();
}

class MyQImageIOHandler: public QImageIOHandler
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQImageIOHandler() : QImageIOHandler() {};
	bool canRead() const { return callbackQImageIOHandler_CanRead(const_cast<MyQImageIOHandler*>(this), this->objectNameAbs().toUtf8().data()) != 0; };
	int currentImageNumber() const { return callbackQImageIOHandler_CurrentImageNumber(const_cast<MyQImageIOHandler*>(this), this->objectNameAbs().toUtf8().data()); };
	QRect currentImageRect() const { return *static_cast<QRect*>(callbackQImageIOHandler_CurrentImageRect(const_cast<MyQImageIOHandler*>(this), this->objectNameAbs().toUtf8().data())); };
	int imageCount() const { return callbackQImageIOHandler_ImageCount(const_cast<MyQImageIOHandler*>(this), this->objectNameAbs().toUtf8().data()); };
	bool jumpToImage(int imageNumber) { return callbackQImageIOHandler_JumpToImage(this, this->objectNameAbs().toUtf8().data(), imageNumber) != 0; };
	bool jumpToNextImage() { return callbackQImageIOHandler_JumpToNextImage(this, this->objectNameAbs().toUtf8().data()) != 0; };
	int loopCount() const { return callbackQImageIOHandler_LoopCount(const_cast<MyQImageIOHandler*>(this), this->objectNameAbs().toUtf8().data()); };
	int nextImageDelay() const { return callbackQImageIOHandler_NextImageDelay(const_cast<MyQImageIOHandler*>(this), this->objectNameAbs().toUtf8().data()); };
	QVariant option(QImageIOHandler::ImageOption option) const { return *static_cast<QVariant*>(callbackQImageIOHandler_Option(const_cast<MyQImageIOHandler*>(this), this->objectNameAbs().toUtf8().data(), option)); };
	bool read(QImage * image) { return callbackQImageIOHandler_Read(this, this->objectNameAbs().toUtf8().data(), image) != 0; };
	void setOption(QImageIOHandler::ImageOption option, const QVariant & value) { callbackQImageIOHandler_SetOption(this, this->objectNameAbs().toUtf8().data(), option, new QVariant(value)); };
	bool supportsOption(QImageIOHandler::ImageOption option) const { return callbackQImageIOHandler_SupportsOption(const_cast<MyQImageIOHandler*>(this), this->objectNameAbs().toUtf8().data(), option) != 0; };
	bool write(const QImage & image) { return callbackQImageIOHandler_Write(this, this->objectNameAbs().toUtf8().data(), new QImage(image)) != 0; };
};

void* QImageIOHandler_NewQImageIOHandler()
{
	return new MyQImageIOHandler();
}

int QImageIOHandler_CanRead(void* ptr)
{
	return static_cast<QImageIOHandler*>(ptr)->canRead();
}

int QImageIOHandler_CurrentImageNumber(void* ptr)
{
	return static_cast<QImageIOHandler*>(ptr)->currentImageNumber();
}

int QImageIOHandler_CurrentImageNumberDefault(void* ptr)
{
	return static_cast<QImageIOHandler*>(ptr)->QImageIOHandler::currentImageNumber();
}

void* QImageIOHandler_CurrentImageRect(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QImageIOHandler*>(ptr)->currentImageRect()).x(), static_cast<QRect>(static_cast<QImageIOHandler*>(ptr)->currentImageRect()).y(), static_cast<QRect>(static_cast<QImageIOHandler*>(ptr)->currentImageRect()).width(), static_cast<QRect>(static_cast<QImageIOHandler*>(ptr)->currentImageRect()).height());
}

void* QImageIOHandler_CurrentImageRectDefault(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QImageIOHandler*>(ptr)->QImageIOHandler::currentImageRect()).x(), static_cast<QRect>(static_cast<QImageIOHandler*>(ptr)->QImageIOHandler::currentImageRect()).y(), static_cast<QRect>(static_cast<QImageIOHandler*>(ptr)->QImageIOHandler::currentImageRect()).width(), static_cast<QRect>(static_cast<QImageIOHandler*>(ptr)->QImageIOHandler::currentImageRect()).height());
}

void* QImageIOHandler_Device(void* ptr)
{
	return static_cast<QImageIOHandler*>(ptr)->device();
}

char* QImageIOHandler_Format(void* ptr)
{
	return QString(static_cast<QImageIOHandler*>(ptr)->format()).toUtf8().data();
}

int QImageIOHandler_ImageCount(void* ptr)
{
	return static_cast<QImageIOHandler*>(ptr)->imageCount();
}

int QImageIOHandler_ImageCountDefault(void* ptr)
{
	return static_cast<QImageIOHandler*>(ptr)->QImageIOHandler::imageCount();
}

int QImageIOHandler_JumpToImage(void* ptr, int imageNumber)
{
	return static_cast<QImageIOHandler*>(ptr)->jumpToImage(imageNumber);
}

int QImageIOHandler_JumpToImageDefault(void* ptr, int imageNumber)
{
	return static_cast<QImageIOHandler*>(ptr)->QImageIOHandler::jumpToImage(imageNumber);
}

int QImageIOHandler_JumpToNextImage(void* ptr)
{
	return static_cast<QImageIOHandler*>(ptr)->jumpToNextImage();
}

int QImageIOHandler_JumpToNextImageDefault(void* ptr)
{
	return static_cast<QImageIOHandler*>(ptr)->QImageIOHandler::jumpToNextImage();
}

int QImageIOHandler_LoopCount(void* ptr)
{
	return static_cast<QImageIOHandler*>(ptr)->loopCount();
}

int QImageIOHandler_LoopCountDefault(void* ptr)
{
	return static_cast<QImageIOHandler*>(ptr)->QImageIOHandler::loopCount();
}

int QImageIOHandler_NextImageDelay(void* ptr)
{
	return static_cast<QImageIOHandler*>(ptr)->nextImageDelay();
}

int QImageIOHandler_NextImageDelayDefault(void* ptr)
{
	return static_cast<QImageIOHandler*>(ptr)->QImageIOHandler::nextImageDelay();
}

void* QImageIOHandler_Option(void* ptr, int option)
{
	return new QVariant(static_cast<QImageIOHandler*>(ptr)->option(static_cast<QImageIOHandler::ImageOption>(option)));
}

void* QImageIOHandler_OptionDefault(void* ptr, int option)
{
	return new QVariant(static_cast<QImageIOHandler*>(ptr)->QImageIOHandler::option(static_cast<QImageIOHandler::ImageOption>(option)));
}

int QImageIOHandler_Read(void* ptr, void* image)
{
	return static_cast<QImageIOHandler*>(ptr)->read(static_cast<QImage*>(image));
}

void QImageIOHandler_SetDevice(void* ptr, void* device)
{
	static_cast<QImageIOHandler*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QImageIOHandler_SetFormat(void* ptr, char* format)
{
	static_cast<QImageIOHandler*>(ptr)->setFormat(QByteArray(format));
}

void QImageIOHandler_SetFormat2(void* ptr, char* format)
{
	static_cast<QImageIOHandler*>(ptr)->setFormat(QByteArray(format));
}

void QImageIOHandler_SetOption(void* ptr, int option, void* value)
{
	static_cast<QImageIOHandler*>(ptr)->setOption(static_cast<QImageIOHandler::ImageOption>(option), *static_cast<QVariant*>(value));
}

void QImageIOHandler_SetOptionDefault(void* ptr, int option, void* value)
{
	static_cast<QImageIOHandler*>(ptr)->QImageIOHandler::setOption(static_cast<QImageIOHandler::ImageOption>(option), *static_cast<QVariant*>(value));
}

int QImageIOHandler_SupportsOption(void* ptr, int option)
{
	return static_cast<QImageIOHandler*>(ptr)->supportsOption(static_cast<QImageIOHandler::ImageOption>(option));
}

int QImageIOHandler_SupportsOptionDefault(void* ptr, int option)
{
	return static_cast<QImageIOHandler*>(ptr)->QImageIOHandler::supportsOption(static_cast<QImageIOHandler::ImageOption>(option));
}

int QImageIOHandler_Write(void* ptr, void* image)
{
	return static_cast<QImageIOHandler*>(ptr)->write(*static_cast<QImage*>(image));
}

int QImageIOHandler_WriteDefault(void* ptr, void* image)
{
	return static_cast<QImageIOHandler*>(ptr)->QImageIOHandler::write(*static_cast<QImage*>(image));
}

void QImageIOHandler_DestroyQImageIOHandler(void* ptr)
{
	static_cast<QImageIOHandler*>(ptr)->~QImageIOHandler();
}

char* QImageIOHandler_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQImageIOHandler*>(static_cast<QImageIOHandler*>(ptr))) {
		return static_cast<MyQImageIOHandler*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QImageIOHandler_BASE").toUtf8().data();
}

void QImageIOHandler_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQImageIOHandler*>(static_cast<QImageIOHandler*>(ptr))) {
		static_cast<MyQImageIOHandler*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQImageIOPlugin: public QImageIOPlugin
{
public:
	MyQImageIOPlugin(QObject *parent) : QImageIOPlugin(parent) {};
	Capabilities capabilities(QIODevice * device, const QByteArray & format) const { return static_cast<QImageIOPlugin::Capability>(callbackQImageIOPlugin_Capabilities(const_cast<MyQImageIOPlugin*>(this), this->objectName().toUtf8().data(), device, QString(format).toUtf8().data())); };
	QImageIOHandler * create(QIODevice * device, const QByteArray & format) const { return static_cast<QImageIOHandler*>(callbackQImageIOPlugin_Create(const_cast<MyQImageIOPlugin*>(this), this->objectName().toUtf8().data(), device, QString(format).toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQImageIOPlugin_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQImageIOPlugin_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQImageIOPlugin_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQImageIOPlugin_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQImageIOPlugin_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQImageIOPlugin_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQImageIOPlugin_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQImageIOPlugin_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQImageIOPlugin_MetaObject(const_cast<MyQImageIOPlugin*>(this), this->objectName().toUtf8().data())); };
};

void* QImageIOPlugin_NewQImageIOPlugin(void* parent)
{
	return new MyQImageIOPlugin(static_cast<QObject*>(parent));
}

int QImageIOPlugin_Capabilities(void* ptr, void* device, char* format)
{
	return static_cast<QImageIOPlugin*>(ptr)->capabilities(static_cast<QIODevice*>(device), QByteArray(format));
}

void* QImageIOPlugin_Create(void* ptr, void* device, char* format)
{
	return static_cast<QImageIOPlugin*>(ptr)->create(static_cast<QIODevice*>(device), QByteArray(format));
}

void QImageIOPlugin_DestroyQImageIOPlugin(void* ptr)
{
	static_cast<QImageIOPlugin*>(ptr)->~QImageIOPlugin();
}

void QImageIOPlugin_TimerEvent(void* ptr, void* event)
{
	static_cast<QImageIOPlugin*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QImageIOPlugin_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QImageIOPlugin*>(ptr)->QImageIOPlugin::timerEvent(static_cast<QTimerEvent*>(event));
}

void QImageIOPlugin_ChildEvent(void* ptr, void* event)
{
	static_cast<QImageIOPlugin*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QImageIOPlugin_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QImageIOPlugin*>(ptr)->QImageIOPlugin::childEvent(static_cast<QChildEvent*>(event));
}

void QImageIOPlugin_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QImageIOPlugin*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QImageIOPlugin_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QImageIOPlugin*>(ptr)->QImageIOPlugin::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QImageIOPlugin_CustomEvent(void* ptr, void* event)
{
	static_cast<QImageIOPlugin*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QImageIOPlugin_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QImageIOPlugin*>(ptr)->QImageIOPlugin::customEvent(static_cast<QEvent*>(event));
}

void QImageIOPlugin_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QImageIOPlugin*>(ptr), "deleteLater");
}

void QImageIOPlugin_DeleteLaterDefault(void* ptr)
{
	static_cast<QImageIOPlugin*>(ptr)->QImageIOPlugin::deleteLater();
}

void QImageIOPlugin_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QImageIOPlugin*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QImageIOPlugin_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QImageIOPlugin*>(ptr)->QImageIOPlugin::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QImageIOPlugin_Event(void* ptr, void* e)
{
	return static_cast<QImageIOPlugin*>(ptr)->event(static_cast<QEvent*>(e));
}

int QImageIOPlugin_EventDefault(void* ptr, void* e)
{
	return static_cast<QImageIOPlugin*>(ptr)->QImageIOPlugin::event(static_cast<QEvent*>(e));
}

int QImageIOPlugin_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QImageIOPlugin*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QImageIOPlugin_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QImageIOPlugin*>(ptr)->QImageIOPlugin::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QImageIOPlugin_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QImageIOPlugin*>(ptr)->metaObject());
}

void* QImageIOPlugin_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QImageIOPlugin*>(ptr)->QImageIOPlugin::metaObject());
}

void* QImageReader_NewQImageReader()
{
	return new QImageReader();
}

void* QImageReader_NewQImageReader2(void* device, char* format)
{
	return new QImageReader(static_cast<QIODevice*>(device), QByteArray(format));
}

void* QImageReader_NewQImageReader3(char* fileName, char* format)
{
	return new QImageReader(QString(fileName), QByteArray(format));
}

int QImageReader_AutoDetectImageFormat(void* ptr)
{
	return static_cast<QImageReader*>(ptr)->autoDetectImageFormat();
}

int QImageReader_AutoTransform(void* ptr)
{
	return static_cast<QImageReader*>(ptr)->autoTransform();
}

void* QImageReader_BackgroundColor(void* ptr)
{
	return new QColor(static_cast<QImageReader*>(ptr)->backgroundColor());
}

int QImageReader_CanRead(void* ptr)
{
	return static_cast<QImageReader*>(ptr)->canRead();
}

void* QImageReader_ClipRect(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QImageReader*>(ptr)->clipRect()).x(), static_cast<QRect>(static_cast<QImageReader*>(ptr)->clipRect()).y(), static_cast<QRect>(static_cast<QImageReader*>(ptr)->clipRect()).width(), static_cast<QRect>(static_cast<QImageReader*>(ptr)->clipRect()).height());
}

int QImageReader_CurrentImageNumber(void* ptr)
{
	return static_cast<QImageReader*>(ptr)->currentImageNumber();
}

void* QImageReader_CurrentImageRect(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QImageReader*>(ptr)->currentImageRect()).x(), static_cast<QRect>(static_cast<QImageReader*>(ptr)->currentImageRect()).y(), static_cast<QRect>(static_cast<QImageReader*>(ptr)->currentImageRect()).width(), static_cast<QRect>(static_cast<QImageReader*>(ptr)->currentImageRect()).height());
}

int QImageReader_DecideFormatFromContent(void* ptr)
{
	return static_cast<QImageReader*>(ptr)->decideFormatFromContent();
}

void* QImageReader_Device(void* ptr)
{
	return static_cast<QImageReader*>(ptr)->device();
}

int QImageReader_Error(void* ptr)
{
	return static_cast<QImageReader*>(ptr)->error();
}

char* QImageReader_ErrorString(void* ptr)
{
	return static_cast<QImageReader*>(ptr)->errorString().toUtf8().data();
}

char* QImageReader_FileName(void* ptr)
{
	return static_cast<QImageReader*>(ptr)->fileName().toUtf8().data();
}

char* QImageReader_Format(void* ptr)
{
	return QString(static_cast<QImageReader*>(ptr)->format()).toUtf8().data();
}

int QImageReader_ImageCount(void* ptr)
{
	return static_cast<QImageReader*>(ptr)->imageCount();
}

char* QImageReader_QImageReader_ImageFormat3(void* device)
{
	return QString(QImageReader::imageFormat(static_cast<QIODevice*>(device))).toUtf8().data();
}

char* QImageReader_QImageReader_ImageFormat2(char* fileName)
{
	return QString(QImageReader::imageFormat(QString(fileName))).toUtf8().data();
}

int QImageReader_ImageFormat(void* ptr)
{
	return static_cast<QImageReader*>(ptr)->imageFormat();
}

int QImageReader_JumpToImage(void* ptr, int imageNumber)
{
	return static_cast<QImageReader*>(ptr)->jumpToImage(imageNumber);
}

int QImageReader_JumpToNextImage(void* ptr)
{
	return static_cast<QImageReader*>(ptr)->jumpToNextImage();
}

int QImageReader_LoopCount(void* ptr)
{
	return static_cast<QImageReader*>(ptr)->loopCount();
}

int QImageReader_NextImageDelay(void* ptr)
{
	return static_cast<QImageReader*>(ptr)->nextImageDelay();
}

int QImageReader_Quality(void* ptr)
{
	return static_cast<QImageReader*>(ptr)->quality();
}

void* QImageReader_Read(void* ptr)
{
	return new QImage(static_cast<QImageReader*>(ptr)->read());
}

int QImageReader_Read2(void* ptr, void* image)
{
	return static_cast<QImageReader*>(ptr)->read(static_cast<QImage*>(image));
}

void* QImageReader_ScaledClipRect(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QImageReader*>(ptr)->scaledClipRect()).x(), static_cast<QRect>(static_cast<QImageReader*>(ptr)->scaledClipRect()).y(), static_cast<QRect>(static_cast<QImageReader*>(ptr)->scaledClipRect()).width(), static_cast<QRect>(static_cast<QImageReader*>(ptr)->scaledClipRect()).height());
}

void* QImageReader_ScaledSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QImageReader*>(ptr)->scaledSize()).width(), static_cast<QSize>(static_cast<QImageReader*>(ptr)->scaledSize()).height());
}

void QImageReader_SetAutoDetectImageFormat(void* ptr, int enabled)
{
	static_cast<QImageReader*>(ptr)->setAutoDetectImageFormat(enabled != 0);
}

void QImageReader_SetAutoTransform(void* ptr, int enabled)
{
	static_cast<QImageReader*>(ptr)->setAutoTransform(enabled != 0);
}

void QImageReader_SetBackgroundColor(void* ptr, void* color)
{
	static_cast<QImageReader*>(ptr)->setBackgroundColor(*static_cast<QColor*>(color));
}

void QImageReader_SetClipRect(void* ptr, void* rect)
{
	static_cast<QImageReader*>(ptr)->setClipRect(*static_cast<QRect*>(rect));
}

void QImageReader_SetDecideFormatFromContent(void* ptr, int ignored)
{
	static_cast<QImageReader*>(ptr)->setDecideFormatFromContent(ignored != 0);
}

void QImageReader_SetDevice(void* ptr, void* device)
{
	static_cast<QImageReader*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QImageReader_SetFileName(void* ptr, char* fileName)
{
	static_cast<QImageReader*>(ptr)->setFileName(QString(fileName));
}

void QImageReader_SetFormat(void* ptr, char* format)
{
	static_cast<QImageReader*>(ptr)->setFormat(QByteArray(format));
}

void QImageReader_SetQuality(void* ptr, int quality)
{
	static_cast<QImageReader*>(ptr)->setQuality(quality);
}

void QImageReader_SetScaledClipRect(void* ptr, void* rect)
{
	static_cast<QImageReader*>(ptr)->setScaledClipRect(*static_cast<QRect*>(rect));
}

void QImageReader_SetScaledSize(void* ptr, void* size)
{
	static_cast<QImageReader*>(ptr)->setScaledSize(*static_cast<QSize*>(size));
}

void* QImageReader_Size(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QImageReader*>(ptr)->size()).width(), static_cast<QSize>(static_cast<QImageReader*>(ptr)->size()).height());
}

char* QImageReader_SubType(void* ptr)
{
	return QString(static_cast<QImageReader*>(ptr)->subType()).toUtf8().data();
}

int QImageReader_SupportsAnimation(void* ptr)
{
	return static_cast<QImageReader*>(ptr)->supportsAnimation();
}

int QImageReader_SupportsOption(void* ptr, int option)
{
	return static_cast<QImageReader*>(ptr)->supportsOption(static_cast<QImageIOHandler::ImageOption>(option));
}

char* QImageReader_Text(void* ptr, char* key)
{
	return static_cast<QImageReader*>(ptr)->text(QString(key)).toUtf8().data();
}

char* QImageReader_TextKeys(void* ptr)
{
	return static_cast<QImageReader*>(ptr)->textKeys().join("|").toUtf8().data();
}

int QImageReader_Transformation(void* ptr)
{
	return static_cast<QImageReader*>(ptr)->transformation();
}

void QImageReader_DestroyQImageReader(void* ptr)
{
	static_cast<QImageReader*>(ptr)->~QImageReader();
}

void* QImageWriter_NewQImageWriter()
{
	return new QImageWriter();
}

void* QImageWriter_NewQImageWriter2(void* device, char* format)
{
	return new QImageWriter(static_cast<QIODevice*>(device), QByteArray(format));
}

void* QImageWriter_NewQImageWriter3(char* fileName, char* format)
{
	return new QImageWriter(QString(fileName), QByteArray(format));
}

int QImageWriter_CanWrite(void* ptr)
{
	return static_cast<QImageWriter*>(ptr)->canWrite();
}

int QImageWriter_Compression(void* ptr)
{
	return static_cast<QImageWriter*>(ptr)->compression();
}

void* QImageWriter_Device(void* ptr)
{
	return static_cast<QImageWriter*>(ptr)->device();
}

int QImageWriter_Error(void* ptr)
{
	return static_cast<QImageWriter*>(ptr)->error();
}

char* QImageWriter_ErrorString(void* ptr)
{
	return static_cast<QImageWriter*>(ptr)->errorString().toUtf8().data();
}

char* QImageWriter_FileName(void* ptr)
{
	return static_cast<QImageWriter*>(ptr)->fileName().toUtf8().data();
}

char* QImageWriter_Format(void* ptr)
{
	return QString(static_cast<QImageWriter*>(ptr)->format()).toUtf8().data();
}

int QImageWriter_OptimizedWrite(void* ptr)
{
	return static_cast<QImageWriter*>(ptr)->optimizedWrite();
}

int QImageWriter_ProgressiveScanWrite(void* ptr)
{
	return static_cast<QImageWriter*>(ptr)->progressiveScanWrite();
}

int QImageWriter_Quality(void* ptr)
{
	return static_cast<QImageWriter*>(ptr)->quality();
}

void QImageWriter_SetCompression(void* ptr, int compression)
{
	static_cast<QImageWriter*>(ptr)->setCompression(compression);
}

void QImageWriter_SetDevice(void* ptr, void* device)
{
	static_cast<QImageWriter*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QImageWriter_SetFileName(void* ptr, char* fileName)
{
	static_cast<QImageWriter*>(ptr)->setFileName(QString(fileName));
}

void QImageWriter_SetFormat(void* ptr, char* format)
{
	static_cast<QImageWriter*>(ptr)->setFormat(QByteArray(format));
}

void QImageWriter_SetOptimizedWrite(void* ptr, int optimize)
{
	static_cast<QImageWriter*>(ptr)->setOptimizedWrite(optimize != 0);
}

void QImageWriter_SetProgressiveScanWrite(void* ptr, int progressive)
{
	static_cast<QImageWriter*>(ptr)->setProgressiveScanWrite(progressive != 0);
}

void QImageWriter_SetQuality(void* ptr, int quality)
{
	static_cast<QImageWriter*>(ptr)->setQuality(quality);
}

void QImageWriter_SetSubType(void* ptr, char* ty)
{
	static_cast<QImageWriter*>(ptr)->setSubType(QByteArray(ty));
}

void QImageWriter_SetText(void* ptr, char* key, char* text)
{
	static_cast<QImageWriter*>(ptr)->setText(QString(key), QString(text));
}

void QImageWriter_SetTransformation(void* ptr, int transform)
{
	static_cast<QImageWriter*>(ptr)->setTransformation(static_cast<QImageIOHandler::Transformation>(transform));
}

char* QImageWriter_SubType(void* ptr)
{
	return QString(static_cast<QImageWriter*>(ptr)->subType()).toUtf8().data();
}

int QImageWriter_SupportsOption(void* ptr, int option)
{
	return static_cast<QImageWriter*>(ptr)->supportsOption(static_cast<QImageIOHandler::ImageOption>(option));
}

int QImageWriter_Transformation(void* ptr)
{
	return static_cast<QImageWriter*>(ptr)->transformation();
}

int QImageWriter_Write(void* ptr, void* image)
{
	return static_cast<QImageWriter*>(ptr)->write(*static_cast<QImage*>(image));
}

void QImageWriter_DestroyQImageWriter(void* ptr)
{
	static_cast<QImageWriter*>(ptr)->~QImageWriter();
}

int QInputEvent_Modifiers(void* ptr)
{
	return static_cast<QInputEvent*>(ptr)->modifiers();
}

class MyQInputMethod: public QInputMethod
{
public:
	void Signal_AnimatingChanged() { callbackQInputMethod_AnimatingChanged(this, this->objectName().toUtf8().data()); };
	void commit() { callbackQInputMethod_Commit(this, this->objectName().toUtf8().data()); };
	void Signal_CursorRectangleChanged() { callbackQInputMethod_CursorRectangleChanged(this, this->objectName().toUtf8().data()); };
	void hide() { callbackQInputMethod_Hide(this, this->objectName().toUtf8().data()); };
	void Signal_InputDirectionChanged(Qt::LayoutDirection newDirection) { callbackQInputMethod_InputDirectionChanged(this, this->objectName().toUtf8().data(), newDirection); };
	void invokeAction(QInputMethod::Action a, int cursorPosition) { callbackQInputMethod_InvokeAction(this, this->objectName().toUtf8().data(), a, cursorPosition); };
	void Signal_KeyboardRectangleChanged() { callbackQInputMethod_KeyboardRectangleChanged(this, this->objectName().toUtf8().data()); };
	void Signal_LocaleChanged() { callbackQInputMethod_LocaleChanged(this, this->objectName().toUtf8().data()); };
	void reset() { callbackQInputMethod_Reset(this, this->objectName().toUtf8().data()); };
	void show() { callbackQInputMethod_Show(this, this->objectName().toUtf8().data()); };
	void update(Qt::InputMethodQueries queries) { callbackQInputMethod_Update(this, this->objectName().toUtf8().data(), queries); };
	void Signal_VisibleChanged() { callbackQInputMethod_VisibleChanged(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQInputMethod_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQInputMethod_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQInputMethod_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQInputMethod_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQInputMethod_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQInputMethod_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQInputMethod_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQInputMethod_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQInputMethod_MetaObject(const_cast<MyQInputMethod*>(this), this->objectName().toUtf8().data())); };
};

void* QInputMethod_CursorRectangle(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QInputMethod*>(ptr)->cursorRectangle()).x(), static_cast<QRectF>(static_cast<QInputMethod*>(ptr)->cursorRectangle()).y(), static_cast<QRectF>(static_cast<QInputMethod*>(ptr)->cursorRectangle()).width(), static_cast<QRectF>(static_cast<QInputMethod*>(ptr)->cursorRectangle()).height());
}

int QInputMethod_InputDirection(void* ptr)
{
	return static_cast<QInputMethod*>(ptr)->inputDirection();
}

int QInputMethod_IsAnimating(void* ptr)
{
	return static_cast<QInputMethod*>(ptr)->isAnimating();
}

int QInputMethod_IsVisible(void* ptr)
{
	return static_cast<QInputMethod*>(ptr)->isVisible();
}

void* QInputMethod_KeyboardRectangle(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QInputMethod*>(ptr)->keyboardRectangle()).x(), static_cast<QRectF>(static_cast<QInputMethod*>(ptr)->keyboardRectangle()).y(), static_cast<QRectF>(static_cast<QInputMethod*>(ptr)->keyboardRectangle()).width(), static_cast<QRectF>(static_cast<QInputMethod*>(ptr)->keyboardRectangle()).height());
}

void* QInputMethod_Locale(void* ptr)
{
	return new QLocale(static_cast<QInputMethod*>(ptr)->locale());
}

void QInputMethod_ConnectAnimatingChanged(void* ptr)
{
	QObject::connect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::animatingChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_AnimatingChanged));
}

void QInputMethod_DisconnectAnimatingChanged(void* ptr)
{
	QObject::disconnect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::animatingChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_AnimatingChanged));
}

void QInputMethod_AnimatingChanged(void* ptr)
{
	static_cast<QInputMethod*>(ptr)->animatingChanged();
}

void QInputMethod_Commit(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "commit");
}

void QInputMethod_ConnectCursorRectangleChanged(void* ptr)
{
	QObject::connect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::cursorRectangleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_CursorRectangleChanged));
}

void QInputMethod_DisconnectCursorRectangleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::cursorRectangleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_CursorRectangleChanged));
}

void QInputMethod_CursorRectangleChanged(void* ptr)
{
	static_cast<QInputMethod*>(ptr)->cursorRectangleChanged();
}

void QInputMethod_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "hide");
}

void QInputMethod_ConnectInputDirectionChanged(void* ptr)
{
	QObject::connect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)(Qt::LayoutDirection)>(&QInputMethod::inputDirectionChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)(Qt::LayoutDirection)>(&MyQInputMethod::Signal_InputDirectionChanged));
}

void QInputMethod_DisconnectInputDirectionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)(Qt::LayoutDirection)>(&QInputMethod::inputDirectionChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)(Qt::LayoutDirection)>(&MyQInputMethod::Signal_InputDirectionChanged));
}

void QInputMethod_InputDirectionChanged(void* ptr, int newDirection)
{
	static_cast<QInputMethod*>(ptr)->inputDirectionChanged(static_cast<Qt::LayoutDirection>(newDirection));
}

void* QInputMethod_InputItemRectangle(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QInputMethod*>(ptr)->inputItemRectangle()).x(), static_cast<QRectF>(static_cast<QInputMethod*>(ptr)->inputItemRectangle()).y(), static_cast<QRectF>(static_cast<QInputMethod*>(ptr)->inputItemRectangle()).width(), static_cast<QRectF>(static_cast<QInputMethod*>(ptr)->inputItemRectangle()).height());
}

void QInputMethod_InvokeAction(void* ptr, int a, int cursorPosition)
{
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "invokeAction", Q_ARG(QInputMethod::Action, static_cast<QInputMethod::Action>(a)), Q_ARG(int, cursorPosition));
}

void QInputMethod_ConnectKeyboardRectangleChanged(void* ptr)
{
	QObject::connect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::keyboardRectangleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_KeyboardRectangleChanged));
}

void QInputMethod_DisconnectKeyboardRectangleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::keyboardRectangleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_KeyboardRectangleChanged));
}

void QInputMethod_KeyboardRectangleChanged(void* ptr)
{
	static_cast<QInputMethod*>(ptr)->keyboardRectangleChanged();
}

void QInputMethod_ConnectLocaleChanged(void* ptr)
{
	QObject::connect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::localeChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_LocaleChanged));
}

void QInputMethod_DisconnectLocaleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::localeChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_LocaleChanged));
}

void QInputMethod_LocaleChanged(void* ptr)
{
	static_cast<QInputMethod*>(ptr)->localeChanged();
}

void* QInputMethod_QInputMethod_QueryFocusObject(int query, void* argument)
{
	return new QVariant(QInputMethod::queryFocusObject(static_cast<Qt::InputMethodQuery>(query), *static_cast<QVariant*>(argument)));
}

void QInputMethod_Reset(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "reset");
}

void QInputMethod_SetInputItemRectangle(void* ptr, void* rect)
{
	static_cast<QInputMethod*>(ptr)->setInputItemRectangle(*static_cast<QRectF*>(rect));
}

void QInputMethod_SetInputItemTransform(void* ptr, void* transform)
{
	static_cast<QInputMethod*>(ptr)->setInputItemTransform(*static_cast<QTransform*>(transform));
}

void QInputMethod_SetVisible(void* ptr, int visible)
{
	static_cast<QInputMethod*>(ptr)->setVisible(visible != 0);
}

void QInputMethod_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "show");
}

void QInputMethod_Update(void* ptr, int queries)
{
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "update", Q_ARG(Qt::InputMethodQuery, static_cast<Qt::InputMethodQuery>(queries)));
}

void QInputMethod_ConnectVisibleChanged(void* ptr)
{
	QObject::connect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::visibleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_VisibleChanged));
}

void QInputMethod_DisconnectVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::visibleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_VisibleChanged));
}

void QInputMethod_VisibleChanged(void* ptr)
{
	static_cast<QInputMethod*>(ptr)->visibleChanged();
}

void QInputMethod_TimerEvent(void* ptr, void* event)
{
	static_cast<QInputMethod*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QInputMethod_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QInputMethod*>(ptr)->QInputMethod::timerEvent(static_cast<QTimerEvent*>(event));
}

void QInputMethod_ChildEvent(void* ptr, void* event)
{
	static_cast<QInputMethod*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QInputMethod_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QInputMethod*>(ptr)->QInputMethod::childEvent(static_cast<QChildEvent*>(event));
}

void QInputMethod_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QInputMethod*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QInputMethod_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QInputMethod*>(ptr)->QInputMethod::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QInputMethod_CustomEvent(void* ptr, void* event)
{
	static_cast<QInputMethod*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QInputMethod_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QInputMethod*>(ptr)->QInputMethod::customEvent(static_cast<QEvent*>(event));
}

void QInputMethod_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "deleteLater");
}

void QInputMethod_DeleteLaterDefault(void* ptr)
{
	static_cast<QInputMethod*>(ptr)->QInputMethod::deleteLater();
}

void QInputMethod_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QInputMethod*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QInputMethod_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QInputMethod*>(ptr)->QInputMethod::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QInputMethod_Event(void* ptr, void* e)
{
	return static_cast<QInputMethod*>(ptr)->event(static_cast<QEvent*>(e));
}

int QInputMethod_EventDefault(void* ptr, void* e)
{
	return static_cast<QInputMethod*>(ptr)->QInputMethod::event(static_cast<QEvent*>(e));
}

int QInputMethod_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QInputMethod*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QInputMethod_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QInputMethod*>(ptr)->QInputMethod::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QInputMethod_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QInputMethod*>(ptr)->metaObject());
}

void* QInputMethod_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QInputMethod*>(ptr)->QInputMethod::metaObject());
}

void* QInputMethodEvent_NewQInputMethodEvent()
{
	return new QInputMethodEvent();
}

void* QInputMethodEvent_NewQInputMethodEvent3(void* other)
{
	return new QInputMethodEvent(*static_cast<QInputMethodEvent*>(other));
}

char* QInputMethodEvent_CommitString(void* ptr)
{
	return static_cast<QInputMethodEvent*>(ptr)->commitString().toUtf8().data();
}

char* QInputMethodEvent_PreeditString(void* ptr)
{
	return static_cast<QInputMethodEvent*>(ptr)->preeditString().toUtf8().data();
}

int QInputMethodEvent_ReplacementLength(void* ptr)
{
	return static_cast<QInputMethodEvent*>(ptr)->replacementLength();
}

int QInputMethodEvent_ReplacementStart(void* ptr)
{
	return static_cast<QInputMethodEvent*>(ptr)->replacementStart();
}

void QInputMethodEvent_SetCommitString(void* ptr, char* commitString, int replaceFrom, int replaceLength)
{
	static_cast<QInputMethodEvent*>(ptr)->setCommitString(QString(commitString), replaceFrom, replaceLength);
}

void* QInputMethodQueryEvent_NewQInputMethodQueryEvent(int queries)
{
	return new QInputMethodQueryEvent(static_cast<Qt::InputMethodQuery>(queries));
}

int QInputMethodQueryEvent_Queries(void* ptr)
{
	return static_cast<QInputMethodQueryEvent*>(ptr)->queries();
}

void QInputMethodQueryEvent_SetValue(void* ptr, int query, void* value)
{
	static_cast<QInputMethodQueryEvent*>(ptr)->setValue(static_cast<Qt::InputMethodQuery>(query), *static_cast<QVariant*>(value));
}

void* QInputMethodQueryEvent_Value(void* ptr, int query)
{
	return new QVariant(static_cast<QInputMethodQueryEvent*>(ptr)->value(static_cast<Qt::InputMethodQuery>(query)));
}

class MyQIntValidator: public QIntValidator
{
public:
	MyQIntValidator(QObject *parent) : QIntValidator(parent) {};
	MyQIntValidator(int minimum, int maximum, QObject *parent) : QIntValidator(minimum, maximum, parent) {};
	QValidator::State validate(QString & input, int & pos) const { return static_cast<QValidator::State>(callbackQIntValidator_Validate(const_cast<MyQIntValidator*>(this), this->objectName().toUtf8().data(), input.toUtf8().data(), pos)); };
	void fixup(QString & input) const { callbackQIntValidator_Fixup(const_cast<MyQIntValidator*>(this), this->objectName().toUtf8().data(), input.toUtf8().data()); };
	void setRange(int bottom, int top) { callbackQIntValidator_SetRange(this, this->objectName().toUtf8().data(), bottom, top); };
	void timerEvent(QTimerEvent * event) { callbackQIntValidator_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQIntValidator_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQIntValidator_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQIntValidator_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQIntValidator_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQIntValidator_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQIntValidator_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQIntValidator_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQIntValidator_MetaObject(const_cast<MyQIntValidator*>(this), this->objectName().toUtf8().data())); };
};

void QIntValidator_SetBottom(void* ptr, int vin)
{
	static_cast<QIntValidator*>(ptr)->setBottom(vin);
}

void QIntValidator_SetTop(void* ptr, int vin)
{
	static_cast<QIntValidator*>(ptr)->setTop(vin);
}





void* QIntValidator_NewQIntValidator(void* parent)
{
	return new MyQIntValidator(static_cast<QObject*>(parent));
}

void* QIntValidator_NewQIntValidator2(int minimum, int maximum, void* parent)
{
	return new MyQIntValidator(minimum, maximum, static_cast<QObject*>(parent));
}

int QIntValidator_Bottom(void* ptr)
{
	return static_cast<QIntValidator*>(ptr)->bottom();
}





void QIntValidator_SetRange(void* ptr, int bottom, int top)
{
	static_cast<QIntValidator*>(ptr)->setRange(bottom, top);
}

void QIntValidator_SetRangeDefault(void* ptr, int bottom, int top)
{
	static_cast<QIntValidator*>(ptr)->QIntValidator::setRange(bottom, top);
}

int QIntValidator_Top(void* ptr)
{
	return static_cast<QIntValidator*>(ptr)->top();
}

void QIntValidator_DestroyQIntValidator(void* ptr)
{
	static_cast<QIntValidator*>(ptr)->~QIntValidator();
}

void QIntValidator_TimerEvent(void* ptr, void* event)
{
	static_cast<QIntValidator*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QIntValidator_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QIntValidator*>(ptr)->QIntValidator::timerEvent(static_cast<QTimerEvent*>(event));
}

void QIntValidator_ChildEvent(void* ptr, void* event)
{
	static_cast<QIntValidator*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QIntValidator_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QIntValidator*>(ptr)->QIntValidator::childEvent(static_cast<QChildEvent*>(event));
}

void QIntValidator_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QIntValidator*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QIntValidator_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QIntValidator*>(ptr)->QIntValidator::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QIntValidator_CustomEvent(void* ptr, void* event)
{
	static_cast<QIntValidator*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QIntValidator_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QIntValidator*>(ptr)->QIntValidator::customEvent(static_cast<QEvent*>(event));
}

void QIntValidator_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QIntValidator*>(ptr), "deleteLater");
}

void QIntValidator_DeleteLaterDefault(void* ptr)
{
	static_cast<QIntValidator*>(ptr)->QIntValidator::deleteLater();
}

void QIntValidator_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QIntValidator*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QIntValidator_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QIntValidator*>(ptr)->QIntValidator::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QIntValidator_Event(void* ptr, void* e)
{
	return static_cast<QIntValidator*>(ptr)->event(static_cast<QEvent*>(e));
}

int QIntValidator_EventDefault(void* ptr, void* e)
{
	return static_cast<QIntValidator*>(ptr)->QIntValidator::event(static_cast<QEvent*>(e));
}

int QIntValidator_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QIntValidator*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QIntValidator_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QIntValidator*>(ptr)->QIntValidator::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QIntValidator_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QIntValidator*>(ptr)->metaObject());
}

void* QIntValidator_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QIntValidator*>(ptr)->QIntValidator::metaObject());
}

int QKeyEvent_Matches(void* ptr, int key)
{
	return static_cast<QKeyEvent*>(ptr)->matches(static_cast<QKeySequence::StandardKey>(key));
}

int QKeyEvent_Count(void* ptr)
{
	return static_cast<QKeyEvent*>(ptr)->count();
}

int QKeyEvent_IsAutoRepeat(void* ptr)
{
	return static_cast<QKeyEvent*>(ptr)->isAutoRepeat();
}

int QKeyEvent_Key(void* ptr)
{
	return static_cast<QKeyEvent*>(ptr)->key();
}

int QKeyEvent_Modifiers(void* ptr)
{
	return static_cast<QKeyEvent*>(ptr)->modifiers();
}

char* QKeyEvent_Text(void* ptr)
{
	return static_cast<QKeyEvent*>(ptr)->text().toUtf8().data();
}

void* QKeySequence_NewQKeySequence()
{
	return new QKeySequence();
}

void* QKeySequence_NewQKeySequence5(int key)
{
	return new QKeySequence(static_cast<QKeySequence::StandardKey>(key));
}

void* QKeySequence_NewQKeySequence4(void* keysequence)
{
	return new QKeySequence(*static_cast<QKeySequence*>(keysequence));
}

void* QKeySequence_NewQKeySequence2(char* key, int format)
{
	return new QKeySequence(QString(key), static_cast<QKeySequence::SequenceFormat>(format));
}

void* QKeySequence_NewQKeySequence3(int k1, int k2, int k3, int k4)
{
	return new QKeySequence(k1, k2, k3, k4);
}

int QKeySequence_Count(void* ptr)
{
	return static_cast<QKeySequence*>(ptr)->count();
}

void* QKeySequence_QKeySequence_FromString(char* str, int format)
{
	return new QKeySequence(QKeySequence::fromString(QString(str), static_cast<QKeySequence::SequenceFormat>(format)));
}

int QKeySequence_IsEmpty(void* ptr)
{
	return static_cast<QKeySequence*>(ptr)->isEmpty();
}

int QKeySequence_Matches(void* ptr, void* seq)
{
	return static_cast<QKeySequence*>(ptr)->matches(*static_cast<QKeySequence*>(seq));
}

void* QKeySequence_QKeySequence_Mnemonic(char* text)
{
	return new QKeySequence(QKeySequence::mnemonic(QString(text)));
}

void QKeySequence_Swap(void* ptr, void* other)
{
	static_cast<QKeySequence*>(ptr)->swap(*static_cast<QKeySequence*>(other));
}

char* QKeySequence_ToString(void* ptr, int format)
{
	return static_cast<QKeySequence*>(ptr)->toString(static_cast<QKeySequence::SequenceFormat>(format)).toUtf8().data();
}

void QKeySequence_DestroyQKeySequence(void* ptr)
{
	static_cast<QKeySequence*>(ptr)->~QKeySequence();
}

void* QLinearGradient_NewQLinearGradient3(double x1, double y1, double x2, double y2)
{
	return new QLinearGradient(static_cast<double>(x1), static_cast<double>(y1), static_cast<double>(x2), static_cast<double>(y2));
}

void* QLinearGradient_NewQLinearGradient()
{
	return new QLinearGradient();
}

void* QLinearGradient_NewQLinearGradient2(void* start, void* finalStop)
{
	return new QLinearGradient(*static_cast<QPointF*>(start), *static_cast<QPointF*>(finalStop));
}

void* QLinearGradient_FinalStop(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QLinearGradient*>(ptr)->finalStop()).x(), static_cast<QPointF>(static_cast<QLinearGradient*>(ptr)->finalStop()).y());
}

void QLinearGradient_SetFinalStop(void* ptr, void* stop)
{
	static_cast<QLinearGradient*>(ptr)->setFinalStop(*static_cast<QPointF*>(stop));
}

void QLinearGradient_SetFinalStop2(void* ptr, double x, double y)
{
	static_cast<QLinearGradient*>(ptr)->setFinalStop(static_cast<double>(x), static_cast<double>(y));
}

void QLinearGradient_SetStart(void* ptr, void* start)
{
	static_cast<QLinearGradient*>(ptr)->setStart(*static_cast<QPointF*>(start));
}

void QLinearGradient_SetStart2(void* ptr, double x, double y)
{
	static_cast<QLinearGradient*>(ptr)->setStart(static_cast<double>(x), static_cast<double>(y));
}

void* QLinearGradient_Start(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QLinearGradient*>(ptr)->start()).x(), static_cast<QPointF>(static_cast<QLinearGradient*>(ptr)->start()).y());
}

void* QMatrix4x4_NewQMatrix4x4()
{
	return new QMatrix4x4();
}

void* QMatrix4x4_NewQMatrix4x47(void* transform)
{
	return new QMatrix4x4(*static_cast<QTransform*>(transform));
}

int QMatrix4x4_IsAffine(void* ptr)
{
	return static_cast<QMatrix4x4*>(ptr)->isAffine();
}

int QMatrix4x4_IsIdentity(void* ptr)
{
	return static_cast<QMatrix4x4*>(ptr)->isIdentity();
}

void QMatrix4x4_LookAt(void* ptr, void* eye, void* center, void* up)
{
	static_cast<QMatrix4x4*>(ptr)->lookAt(*static_cast<QVector3D*>(eye), *static_cast<QVector3D*>(center), *static_cast<QVector3D*>(up));
}

void* QMatrix4x4_Map(void* ptr, void* point)
{
	return new QPoint(static_cast<QPoint>(static_cast<QMatrix4x4*>(ptr)->map(*static_cast<QPoint*>(point))).x(), static_cast<QPoint>(static_cast<QMatrix4x4*>(ptr)->map(*static_cast<QPoint*>(point))).y());
}

void* QMatrix4x4_Map2(void* ptr, void* point)
{
	return new QPointF(static_cast<QPointF>(static_cast<QMatrix4x4*>(ptr)->map(*static_cast<QPointF*>(point))).x(), static_cast<QPointF>(static_cast<QMatrix4x4*>(ptr)->map(*static_cast<QPointF*>(point))).y());
}

void* QMatrix4x4_MapRect(void* ptr, void* rect)
{
	return new QRect(static_cast<QRect>(static_cast<QMatrix4x4*>(ptr)->mapRect(*static_cast<QRect*>(rect))).x(), static_cast<QRect>(static_cast<QMatrix4x4*>(ptr)->mapRect(*static_cast<QRect*>(rect))).y(), static_cast<QRect>(static_cast<QMatrix4x4*>(ptr)->mapRect(*static_cast<QRect*>(rect))).width(), static_cast<QRect>(static_cast<QMatrix4x4*>(ptr)->mapRect(*static_cast<QRect*>(rect))).height());
}

void* QMatrix4x4_MapRect2(void* ptr, void* rect)
{
	return new QRectF(static_cast<QRectF>(static_cast<QMatrix4x4*>(ptr)->mapRect(*static_cast<QRectF*>(rect))).x(), static_cast<QRectF>(static_cast<QMatrix4x4*>(ptr)->mapRect(*static_cast<QRectF*>(rect))).y(), static_cast<QRectF>(static_cast<QMatrix4x4*>(ptr)->mapRect(*static_cast<QRectF*>(rect))).width(), static_cast<QRectF>(static_cast<QMatrix4x4*>(ptr)->mapRect(*static_cast<QRectF*>(rect))).height());
}

void QMatrix4x4_Optimize(void* ptr)
{
	static_cast<QMatrix4x4*>(ptr)->optimize();
}

void QMatrix4x4_Ortho2(void* ptr, void* rect)
{
	static_cast<QMatrix4x4*>(ptr)->ortho(*static_cast<QRect*>(rect));
}

void QMatrix4x4_Ortho3(void* ptr, void* rect)
{
	static_cast<QMatrix4x4*>(ptr)->ortho(*static_cast<QRectF*>(rect));
}

void QMatrix4x4_Rotate2(void* ptr, void* quaternion)
{
	static_cast<QMatrix4x4*>(ptr)->rotate(*static_cast<QQuaternion*>(quaternion));
}

void QMatrix4x4_Scale(void* ptr, void* vector)
{
	static_cast<QMatrix4x4*>(ptr)->scale(*static_cast<QVector3D*>(vector));
}

void QMatrix4x4_SetColumn(void* ptr, int index, void* value)
{
	static_cast<QMatrix4x4*>(ptr)->setColumn(index, *static_cast<QVector4D*>(value));
}

void QMatrix4x4_SetRow(void* ptr, int index, void* value)
{
	static_cast<QMatrix4x4*>(ptr)->setRow(index, *static_cast<QVector4D*>(value));
}

void QMatrix4x4_SetToIdentity(void* ptr)
{
	static_cast<QMatrix4x4*>(ptr)->setToIdentity();
}

void QMatrix4x4_Translate(void* ptr, void* vector)
{
	static_cast<QMatrix4x4*>(ptr)->translate(*static_cast<QVector3D*>(vector));
}

void QMatrix4x4_Viewport2(void* ptr, void* rect)
{
	static_cast<QMatrix4x4*>(ptr)->viewport(*static_cast<QRectF*>(rect));
}

void* QMouseEvent_NewQMouseEvent(int ty, void* localPos, int button, int buttons, int modifiers)
{
	return new QMouseEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(localPos), static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

void* QMouseEvent_NewQMouseEvent2(int ty, void* localPos, void* screenPos, int button, int buttons, int modifiers)
{
	return new QMouseEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(localPos), *static_cast<QPointF*>(screenPos), static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

void* QMouseEvent_NewQMouseEvent3(int ty, void* localPos, void* windowPos, void* screenPos, int button, int buttons, int modifiers)
{
	return new QMouseEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(localPos), *static_cast<QPointF*>(windowPos), *static_cast<QPointF*>(screenPos), static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

int QMouseEvent_Button(void* ptr)
{
	return static_cast<QMouseEvent*>(ptr)->button();
}

int QMouseEvent_Buttons(void* ptr)
{
	return static_cast<QMouseEvent*>(ptr)->buttons();
}

int QMouseEvent_Flags(void* ptr)
{
	return static_cast<QMouseEvent*>(ptr)->flags();
}

void* QMouseEvent_GlobalPos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QMouseEvent*>(ptr)->globalPos()).x(), static_cast<QPoint>(static_cast<QMouseEvent*>(ptr)->globalPos()).y());
}

int QMouseEvent_GlobalX(void* ptr)
{
	return static_cast<QMouseEvent*>(ptr)->globalX();
}

int QMouseEvent_GlobalY(void* ptr)
{
	return static_cast<QMouseEvent*>(ptr)->globalY();
}

void* QMouseEvent_LocalPos(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QMouseEvent*>(ptr)->localPos()).x(), static_cast<QPointF>(static_cast<QMouseEvent*>(ptr)->localPos()).y());
}

void* QMouseEvent_Pos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QMouseEvent*>(ptr)->pos()).x(), static_cast<QPoint>(static_cast<QMouseEvent*>(ptr)->pos()).y());
}

void* QMouseEvent_ScreenPos(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QMouseEvent*>(ptr)->screenPos()).x(), static_cast<QPointF>(static_cast<QMouseEvent*>(ptr)->screenPos()).y());
}

int QMouseEvent_Source(void* ptr)
{
	return static_cast<QMouseEvent*>(ptr)->source();
}

void* QMouseEvent_WindowPos(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QMouseEvent*>(ptr)->windowPos()).x(), static_cast<QPointF>(static_cast<QMouseEvent*>(ptr)->windowPos()).y());
}

int QMouseEvent_X(void* ptr)
{
	return static_cast<QMouseEvent*>(ptr)->x();
}

int QMouseEvent_Y(void* ptr)
{
	return static_cast<QMouseEvent*>(ptr)->y();
}

void* QMoveEvent_NewQMoveEvent(void* pos, void* oldPos)
{
	return new QMoveEvent(*static_cast<QPoint*>(pos), *static_cast<QPoint*>(oldPos));
}

void* QMoveEvent_OldPos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QMoveEvent*>(ptr)->oldPos()).x(), static_cast<QPoint>(static_cast<QMoveEvent*>(ptr)->oldPos()).y());
}

void* QMoveEvent_Pos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QMoveEvent*>(ptr)->pos()).x(), static_cast<QPoint>(static_cast<QMoveEvent*>(ptr)->pos()).y());
}

class MyQMovie: public QMovie
{
public:
	MyQMovie(QIODevice *device, const QByteArray &format, QObject *parent) : QMovie(device, format, parent) {};
	MyQMovie(QObject *parent) : QMovie(parent) {};
	MyQMovie(const QString &fileName, const QByteArray &format, QObject *parent) : QMovie(fileName, format, parent) {};
	void setSpeed(int percentSpeed) { callbackQMovie_SetSpeed(this, this->objectName().toUtf8().data(), percentSpeed); };
	void Signal_Error(QImageReader::ImageReaderError error) { callbackQMovie_Error(this, this->objectName().toUtf8().data(), error); };
	void Signal_Finished() { callbackQMovie_Finished(this, this->objectName().toUtf8().data()); };
	void Signal_FrameChanged(int frameNumber) { callbackQMovie_FrameChanged(this, this->objectName().toUtf8().data(), frameNumber); };
	bool jumpToNextFrame() { return callbackQMovie_JumpToNextFrame(this, this->objectName().toUtf8().data()) != 0; };
	void Signal_Resized(const QSize & size) { callbackQMovie_Resized(this, this->objectName().toUtf8().data(), new QSize(static_cast<QSize>(size).width(), static_cast<QSize>(size).height())); };
	void setPaused(bool paused) { callbackQMovie_SetPaused(this, this->objectName().toUtf8().data(), paused); };
	void start() { callbackQMovie_Start(this, this->objectName().toUtf8().data()); };
	void Signal_Started() { callbackQMovie_Started(this, this->objectName().toUtf8().data()); };
	void Signal_StateChanged(QMovie::MovieState state) { callbackQMovie_StateChanged(this, this->objectName().toUtf8().data(), state); };
	void stop() { callbackQMovie_Stop(this, this->objectName().toUtf8().data()); };
	void Signal_Updated(const QRect & rect) { callbackQMovie_Updated(this, this->objectName().toUtf8().data(), new QRect(static_cast<QRect>(rect).x(), static_cast<QRect>(rect).y(), static_cast<QRect>(rect).width(), static_cast<QRect>(rect).height())); };
	void timerEvent(QTimerEvent * event) { callbackQMovie_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQMovie_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQMovie_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQMovie_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQMovie_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQMovie_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQMovie_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQMovie_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQMovie_MetaObject(const_cast<MyQMovie*>(this), this->objectName().toUtf8().data())); };
};

int QMovie_CacheMode(void* ptr)
{
	return static_cast<QMovie*>(ptr)->cacheMode();
}

void QMovie_SetCacheMode(void* ptr, int mode)
{
	static_cast<QMovie*>(ptr)->setCacheMode(static_cast<QMovie::CacheMode>(mode));
}

void QMovie_SetSpeed(void* ptr, int percentSpeed)
{
	QMetaObject::invokeMethod(static_cast<QMovie*>(ptr), "setSpeed", Q_ARG(int, percentSpeed));
}

int QMovie_Speed(void* ptr)
{
	return static_cast<QMovie*>(ptr)->speed();
}

void* QMovie_NewQMovie2(void* device, char* format, void* parent)
{
	return new MyQMovie(static_cast<QIODevice*>(device), QByteArray(format), static_cast<QObject*>(parent));
}

void* QMovie_NewQMovie(void* parent)
{
	return new MyQMovie(static_cast<QObject*>(parent));
}

void* QMovie_NewQMovie3(char* fileName, char* format, void* parent)
{
	return new MyQMovie(QString(fileName), QByteArray(format), static_cast<QObject*>(parent));
}

void* QMovie_BackgroundColor(void* ptr)
{
	return new QColor(static_cast<QMovie*>(ptr)->backgroundColor());
}

int QMovie_CurrentFrameNumber(void* ptr)
{
	return static_cast<QMovie*>(ptr)->currentFrameNumber();
}

void* QMovie_CurrentImage(void* ptr)
{
	return new QImage(static_cast<QMovie*>(ptr)->currentImage());
}

void* QMovie_CurrentPixmap(void* ptr)
{
	return new QPixmap(static_cast<QMovie*>(ptr)->currentPixmap());
}

void* QMovie_Device(void* ptr)
{
	return static_cast<QMovie*>(ptr)->device();
}

void QMovie_ConnectError(void* ptr)
{
	QObject::connect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(QImageReader::ImageReaderError)>(&QMovie::error), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(QImageReader::ImageReaderError)>(&MyQMovie::Signal_Error));
}

void QMovie_DisconnectError(void* ptr)
{
	QObject::disconnect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(QImageReader::ImageReaderError)>(&QMovie::error), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(QImageReader::ImageReaderError)>(&MyQMovie::Signal_Error));
}

void QMovie_Error(void* ptr, int error)
{
	static_cast<QMovie*>(ptr)->error(static_cast<QImageReader::ImageReaderError>(error));
}

char* QMovie_FileName(void* ptr)
{
	return static_cast<QMovie*>(ptr)->fileName().toUtf8().data();
}

void QMovie_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)()>(&QMovie::finished), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)()>(&MyQMovie::Signal_Finished));
}

void QMovie_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)()>(&QMovie::finished), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)()>(&MyQMovie::Signal_Finished));
}

void QMovie_Finished(void* ptr)
{
	static_cast<QMovie*>(ptr)->finished();
}

char* QMovie_Format(void* ptr)
{
	return QString(static_cast<QMovie*>(ptr)->format()).toUtf8().data();
}

void QMovie_ConnectFrameChanged(void* ptr)
{
	QObject::connect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(int)>(&QMovie::frameChanged), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(int)>(&MyQMovie::Signal_FrameChanged));
}

void QMovie_DisconnectFrameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(int)>(&QMovie::frameChanged), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(int)>(&MyQMovie::Signal_FrameChanged));
}

void QMovie_FrameChanged(void* ptr, int frameNumber)
{
	static_cast<QMovie*>(ptr)->frameChanged(frameNumber);
}

int QMovie_FrameCount(void* ptr)
{
	return static_cast<QMovie*>(ptr)->frameCount();
}

void* QMovie_FrameRect(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QMovie*>(ptr)->frameRect()).x(), static_cast<QRect>(static_cast<QMovie*>(ptr)->frameRect()).y(), static_cast<QRect>(static_cast<QMovie*>(ptr)->frameRect()).width(), static_cast<QRect>(static_cast<QMovie*>(ptr)->frameRect()).height());
}

int QMovie_IsValid(void* ptr)
{
	return static_cast<QMovie*>(ptr)->isValid();
}

int QMovie_JumpToFrame(void* ptr, int frameNumber)
{
	return static_cast<QMovie*>(ptr)->jumpToFrame(frameNumber);
}

int QMovie_JumpToNextFrame(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QMovie*>(ptr), "jumpToNextFrame", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QMovie_LoopCount(void* ptr)
{
	return static_cast<QMovie*>(ptr)->loopCount();
}

int QMovie_NextFrameDelay(void* ptr)
{
	return static_cast<QMovie*>(ptr)->nextFrameDelay();
}

void QMovie_ConnectResized(void* ptr)
{
	QObject::connect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(const QSize &)>(&QMovie::resized), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(const QSize &)>(&MyQMovie::Signal_Resized));
}

void QMovie_DisconnectResized(void* ptr)
{
	QObject::disconnect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(const QSize &)>(&QMovie::resized), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(const QSize &)>(&MyQMovie::Signal_Resized));
}

void QMovie_Resized(void* ptr, void* size)
{
	static_cast<QMovie*>(ptr)->resized(*static_cast<QSize*>(size));
}

void* QMovie_ScaledSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QMovie*>(ptr)->scaledSize()).width(), static_cast<QSize>(static_cast<QMovie*>(ptr)->scaledSize()).height());
}

void QMovie_SetBackgroundColor(void* ptr, void* color)
{
	static_cast<QMovie*>(ptr)->setBackgroundColor(*static_cast<QColor*>(color));
}

void QMovie_SetDevice(void* ptr, void* device)
{
	static_cast<QMovie*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QMovie_SetFileName(void* ptr, char* fileName)
{
	static_cast<QMovie*>(ptr)->setFileName(QString(fileName));
}

void QMovie_SetFormat(void* ptr, char* format)
{
	static_cast<QMovie*>(ptr)->setFormat(QByteArray(format));
}

void QMovie_SetPaused(void* ptr, int paused)
{
	QMetaObject::invokeMethod(static_cast<QMovie*>(ptr), "setPaused", Q_ARG(bool, paused != 0));
}

void QMovie_SetScaledSize(void* ptr, void* size)
{
	static_cast<QMovie*>(ptr)->setScaledSize(*static_cast<QSize*>(size));
}

void QMovie_Start(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMovie*>(ptr), "start");
}

void QMovie_ConnectStarted(void* ptr)
{
	QObject::connect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)()>(&QMovie::started), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)()>(&MyQMovie::Signal_Started));
}

void QMovie_DisconnectStarted(void* ptr)
{
	QObject::disconnect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)()>(&QMovie::started), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)()>(&MyQMovie::Signal_Started));
}

void QMovie_Started(void* ptr)
{
	static_cast<QMovie*>(ptr)->started();
}

int QMovie_State(void* ptr)
{
	return static_cast<QMovie*>(ptr)->state();
}

void QMovie_ConnectStateChanged(void* ptr)
{
	QObject::connect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(QMovie::MovieState)>(&QMovie::stateChanged), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(QMovie::MovieState)>(&MyQMovie::Signal_StateChanged));
}

void QMovie_DisconnectStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(QMovie::MovieState)>(&QMovie::stateChanged), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(QMovie::MovieState)>(&MyQMovie::Signal_StateChanged));
}

void QMovie_StateChanged(void* ptr, int state)
{
	static_cast<QMovie*>(ptr)->stateChanged(static_cast<QMovie::MovieState>(state));
}

void QMovie_Stop(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMovie*>(ptr), "stop");
}

void QMovie_ConnectUpdated(void* ptr)
{
	QObject::connect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(const QRect &)>(&QMovie::updated), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(const QRect &)>(&MyQMovie::Signal_Updated));
}

void QMovie_DisconnectUpdated(void* ptr)
{
	QObject::disconnect(static_cast<QMovie*>(ptr), static_cast<void (QMovie::*)(const QRect &)>(&QMovie::updated), static_cast<MyQMovie*>(ptr), static_cast<void (MyQMovie::*)(const QRect &)>(&MyQMovie::Signal_Updated));
}

void QMovie_Updated(void* ptr, void* rect)
{
	static_cast<QMovie*>(ptr)->updated(*static_cast<QRect*>(rect));
}

void QMovie_DestroyQMovie(void* ptr)
{
	static_cast<QMovie*>(ptr)->~QMovie();
}

void QMovie_TimerEvent(void* ptr, void* event)
{
	static_cast<QMovie*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QMovie_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QMovie*>(ptr)->QMovie::timerEvent(static_cast<QTimerEvent*>(event));
}

void QMovie_ChildEvent(void* ptr, void* event)
{
	static_cast<QMovie*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QMovie_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QMovie*>(ptr)->QMovie::childEvent(static_cast<QChildEvent*>(event));
}

void QMovie_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QMovie*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMovie_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMovie*>(ptr)->QMovie::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMovie_CustomEvent(void* ptr, void* event)
{
	static_cast<QMovie*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QMovie_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QMovie*>(ptr)->QMovie::customEvent(static_cast<QEvent*>(event));
}

void QMovie_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QMovie*>(ptr), "deleteLater");
}

void QMovie_DeleteLaterDefault(void* ptr)
{
	static_cast<QMovie*>(ptr)->QMovie::deleteLater();
}

void QMovie_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QMovie*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QMovie_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QMovie*>(ptr)->QMovie::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QMovie_Event(void* ptr, void* e)
{
	return static_cast<QMovie*>(ptr)->event(static_cast<QEvent*>(e));
}

int QMovie_EventDefault(void* ptr, void* e)
{
	return static_cast<QMovie*>(ptr)->QMovie::event(static_cast<QEvent*>(e));
}

int QMovie_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QMovie*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QMovie_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QMovie*>(ptr)->QMovie::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QMovie_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMovie*>(ptr)->metaObject());
}

void* QMovie_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QMovie*>(ptr)->QMovie::metaObject());
}

int QNativeGestureEvent_GestureType(void* ptr)
{
	return static_cast<QNativeGestureEvent*>(ptr)->gestureType();
}

void* QNativeGestureEvent_GlobalPos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QNativeGestureEvent*>(ptr)->globalPos()).x(), static_cast<QPoint>(static_cast<QNativeGestureEvent*>(ptr)->globalPos()).y());
}

void* QNativeGestureEvent_LocalPos(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QNativeGestureEvent*>(ptr)->localPos()).x(), static_cast<QPointF>(static_cast<QNativeGestureEvent*>(ptr)->localPos()).y());
}

void* QNativeGestureEvent_Pos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QNativeGestureEvent*>(ptr)->pos()).x(), static_cast<QPoint>(static_cast<QNativeGestureEvent*>(ptr)->pos()).y());
}

void* QNativeGestureEvent_ScreenPos(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QNativeGestureEvent*>(ptr)->screenPos()).x(), static_cast<QPointF>(static_cast<QNativeGestureEvent*>(ptr)->screenPos()).y());
}

double QNativeGestureEvent_Value(void* ptr)
{
	return static_cast<double>(static_cast<QNativeGestureEvent*>(ptr)->value());
}

void* QNativeGestureEvent_WindowPos(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QNativeGestureEvent*>(ptr)->windowPos()).x(), static_cast<QPointF>(static_cast<QNativeGestureEvent*>(ptr)->windowPos()).y());
}

class MyQOffscreenSurface: public QOffscreenSurface
{
public:
	MyQOffscreenSurface(QScreen *targetScreen) : QOffscreenSurface(targetScreen) {};
	QSurfaceFormat format() const { return *static_cast<QSurfaceFormat*>(callbackQOffscreenSurface_Format(const_cast<MyQOffscreenSurface*>(this), this->objectName().toUtf8().data())); };
	void Signal_ScreenChanged(QScreen * screen) { callbackQOffscreenSurface_ScreenChanged(this, this->objectName().toUtf8().data(), screen); };
	QSize size() const { return *static_cast<QSize*>(callbackQOffscreenSurface_Size(const_cast<MyQOffscreenSurface*>(this), this->objectName().toUtf8().data())); };
	SurfaceType surfaceType() const { return static_cast<QSurface::SurfaceType>(callbackQOffscreenSurface_SurfaceType(const_cast<MyQOffscreenSurface*>(this), this->objectName().toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQOffscreenSurface_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQOffscreenSurface_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQOffscreenSurface_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQOffscreenSurface_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQOffscreenSurface_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQOffscreenSurface_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQOffscreenSurface_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQOffscreenSurface_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQOffscreenSurface_MetaObject(const_cast<MyQOffscreenSurface*>(this), this->objectName().toUtf8().data())); };
};

void* QOffscreenSurface_NewQOffscreenSurface(void* targetScreen)
{
	return new MyQOffscreenSurface(static_cast<QScreen*>(targetScreen));
}

void QOffscreenSurface_Create(void* ptr)
{
	static_cast<QOffscreenSurface*>(ptr)->create();
}

void QOffscreenSurface_Destroy(void* ptr)
{
	static_cast<QOffscreenSurface*>(ptr)->destroy();
}

void* QOffscreenSurface_Format(void* ptr)
{
	return new QSurfaceFormat(static_cast<QOffscreenSurface*>(ptr)->format());
}

void* QOffscreenSurface_FormatDefault(void* ptr)
{
	return new QSurfaceFormat(static_cast<QOffscreenSurface*>(ptr)->QOffscreenSurface::format());
}

int QOffscreenSurface_IsValid(void* ptr)
{
	return static_cast<QOffscreenSurface*>(ptr)->isValid();
}

void* QOffscreenSurface_RequestedFormat(void* ptr)
{
	return new QSurfaceFormat(static_cast<QOffscreenSurface*>(ptr)->requestedFormat());
}

void* QOffscreenSurface_Screen(void* ptr)
{
	return static_cast<QOffscreenSurface*>(ptr)->screen();
}

void QOffscreenSurface_ConnectScreenChanged(void* ptr)
{
	QObject::connect(static_cast<QOffscreenSurface*>(ptr), static_cast<void (QOffscreenSurface::*)(QScreen *)>(&QOffscreenSurface::screenChanged), static_cast<MyQOffscreenSurface*>(ptr), static_cast<void (MyQOffscreenSurface::*)(QScreen *)>(&MyQOffscreenSurface::Signal_ScreenChanged));
}

void QOffscreenSurface_DisconnectScreenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QOffscreenSurface*>(ptr), static_cast<void (QOffscreenSurface::*)(QScreen *)>(&QOffscreenSurface::screenChanged), static_cast<MyQOffscreenSurface*>(ptr), static_cast<void (MyQOffscreenSurface::*)(QScreen *)>(&MyQOffscreenSurface::Signal_ScreenChanged));
}

void QOffscreenSurface_ScreenChanged(void* ptr, void* screen)
{
	static_cast<QOffscreenSurface*>(ptr)->screenChanged(static_cast<QScreen*>(screen));
}

void QOffscreenSurface_SetFormat(void* ptr, void* format)
{
	static_cast<QOffscreenSurface*>(ptr)->setFormat(*static_cast<QSurfaceFormat*>(format));
}

void QOffscreenSurface_SetScreen(void* ptr, void* newScreen)
{
	static_cast<QOffscreenSurface*>(ptr)->setScreen(static_cast<QScreen*>(newScreen));
}

void* QOffscreenSurface_Size(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QOffscreenSurface*>(ptr)->size()).width(), static_cast<QSize>(static_cast<QOffscreenSurface*>(ptr)->size()).height());
}

void* QOffscreenSurface_SizeDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QOffscreenSurface*>(ptr)->QOffscreenSurface::size()).width(), static_cast<QSize>(static_cast<QOffscreenSurface*>(ptr)->QOffscreenSurface::size()).height());
}

int QOffscreenSurface_SurfaceType(void* ptr)
{
	return static_cast<QOffscreenSurface*>(ptr)->surfaceType();
}

int QOffscreenSurface_SurfaceTypeDefault(void* ptr)
{
	return static_cast<QOffscreenSurface*>(ptr)->QOffscreenSurface::surfaceType();
}

void QOffscreenSurface_DestroyQOffscreenSurface(void* ptr)
{
	static_cast<QOffscreenSurface*>(ptr)->~QOffscreenSurface();
}

void QOffscreenSurface_TimerEvent(void* ptr, void* event)
{
	static_cast<QOffscreenSurface*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QOffscreenSurface_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QOffscreenSurface*>(ptr)->QOffscreenSurface::timerEvent(static_cast<QTimerEvent*>(event));
}

void QOffscreenSurface_ChildEvent(void* ptr, void* event)
{
	static_cast<QOffscreenSurface*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QOffscreenSurface_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QOffscreenSurface*>(ptr)->QOffscreenSurface::childEvent(static_cast<QChildEvent*>(event));
}

void QOffscreenSurface_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QOffscreenSurface*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QOffscreenSurface_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QOffscreenSurface*>(ptr)->QOffscreenSurface::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QOffscreenSurface_CustomEvent(void* ptr, void* event)
{
	static_cast<QOffscreenSurface*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QOffscreenSurface_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QOffscreenSurface*>(ptr)->QOffscreenSurface::customEvent(static_cast<QEvent*>(event));
}

void QOffscreenSurface_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QOffscreenSurface*>(ptr), "deleteLater");
}

void QOffscreenSurface_DeleteLaterDefault(void* ptr)
{
	static_cast<QOffscreenSurface*>(ptr)->QOffscreenSurface::deleteLater();
}

void QOffscreenSurface_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QOffscreenSurface*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QOffscreenSurface_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QOffscreenSurface*>(ptr)->QOffscreenSurface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QOffscreenSurface_Event(void* ptr, void* e)
{
	return static_cast<QOffscreenSurface*>(ptr)->event(static_cast<QEvent*>(e));
}

int QOffscreenSurface_EventDefault(void* ptr, void* e)
{
	return static_cast<QOffscreenSurface*>(ptr)->QOffscreenSurface::event(static_cast<QEvent*>(e));
}

int QOffscreenSurface_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QOffscreenSurface*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QOffscreenSurface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QOffscreenSurface*>(ptr)->QOffscreenSurface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QOffscreenSurface_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QOffscreenSurface*>(ptr)->metaObject());
}

void* QOffscreenSurface_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QOffscreenSurface*>(ptr)->QOffscreenSurface::metaObject());
}

void* QPageLayout_NewQPageLayout()
{
	return new QPageLayout();
}

void* QPageLayout_NewQPageLayout3(void* other)
{
	return new QPageLayout(*static_cast<QPageLayout*>(other));
}

void* QPageLayout_NewQPageLayout2(void* pageSize, int orientation, void* margins, int units, void* minMargins)
{
	return new QPageLayout(*static_cast<QPageSize*>(pageSize), static_cast<QPageLayout::Orientation>(orientation), *static_cast<QMarginsF*>(margins), static_cast<QPageLayout::Unit>(units), *static_cast<QMarginsF*>(minMargins));
}

void* QPageLayout_FullRect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QPageLayout*>(ptr)->fullRect()).x(), static_cast<QRectF>(static_cast<QPageLayout*>(ptr)->fullRect()).y(), static_cast<QRectF>(static_cast<QPageLayout*>(ptr)->fullRect()).width(), static_cast<QRectF>(static_cast<QPageLayout*>(ptr)->fullRect()).height());
}

void* QPageLayout_FullRect2(void* ptr, int units)
{
	return new QRectF(static_cast<QRectF>(static_cast<QPageLayout*>(ptr)->fullRect(static_cast<QPageLayout::Unit>(units))).x(), static_cast<QRectF>(static_cast<QPageLayout*>(ptr)->fullRect(static_cast<QPageLayout::Unit>(units))).y(), static_cast<QRectF>(static_cast<QPageLayout*>(ptr)->fullRect(static_cast<QPageLayout::Unit>(units))).width(), static_cast<QRectF>(static_cast<QPageLayout*>(ptr)->fullRect(static_cast<QPageLayout::Unit>(units))).height());
}

void* QPageLayout_FullRectPixels(void* ptr, int resolution)
{
	return new QRect(static_cast<QRect>(static_cast<QPageLayout*>(ptr)->fullRectPixels(resolution)).x(), static_cast<QRect>(static_cast<QPageLayout*>(ptr)->fullRectPixels(resolution)).y(), static_cast<QRect>(static_cast<QPageLayout*>(ptr)->fullRectPixels(resolution)).width(), static_cast<QRect>(static_cast<QPageLayout*>(ptr)->fullRectPixels(resolution)).height());
}

void* QPageLayout_FullRectPoints(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QPageLayout*>(ptr)->fullRectPoints()).x(), static_cast<QRect>(static_cast<QPageLayout*>(ptr)->fullRectPoints()).y(), static_cast<QRect>(static_cast<QPageLayout*>(ptr)->fullRectPoints()).width(), static_cast<QRect>(static_cast<QPageLayout*>(ptr)->fullRectPoints()).height());
}

int QPageLayout_IsEquivalentTo(void* ptr, void* other)
{
	return static_cast<QPageLayout*>(ptr)->isEquivalentTo(*static_cast<QPageLayout*>(other));
}

int QPageLayout_IsValid(void* ptr)
{
	return static_cast<QPageLayout*>(ptr)->isValid();
}

void* QPageLayout_Margins(void* ptr)
{
	return new QMarginsF(static_cast<QMarginsF>(static_cast<QPageLayout*>(ptr)->margins()).left(), static_cast<QMarginsF>(static_cast<QPageLayout*>(ptr)->margins()).top(), static_cast<QMarginsF>(static_cast<QPageLayout*>(ptr)->margins()).right(), static_cast<QMarginsF>(static_cast<QPageLayout*>(ptr)->margins()).bottom());
}

void* QPageLayout_Margins2(void* ptr, int units)
{
	return new QMarginsF(static_cast<QMarginsF>(static_cast<QPageLayout*>(ptr)->margins(static_cast<QPageLayout::Unit>(units))).left(), static_cast<QMarginsF>(static_cast<QPageLayout*>(ptr)->margins(static_cast<QPageLayout::Unit>(units))).top(), static_cast<QMarginsF>(static_cast<QPageLayout*>(ptr)->margins(static_cast<QPageLayout::Unit>(units))).right(), static_cast<QMarginsF>(static_cast<QPageLayout*>(ptr)->margins(static_cast<QPageLayout::Unit>(units))).bottom());
}

void* QPageLayout_MarginsPixels(void* ptr, int resolution)
{
	return new QMargins(static_cast<QMargins>(static_cast<QPageLayout*>(ptr)->marginsPixels(resolution)).left(), static_cast<QMargins>(static_cast<QPageLayout*>(ptr)->marginsPixels(resolution)).top(), static_cast<QMargins>(static_cast<QPageLayout*>(ptr)->marginsPixels(resolution)).right(), static_cast<QMargins>(static_cast<QPageLayout*>(ptr)->marginsPixels(resolution)).bottom());
}

void* QPageLayout_MarginsPoints(void* ptr)
{
	return new QMargins(static_cast<QMargins>(static_cast<QPageLayout*>(ptr)->marginsPoints()).left(), static_cast<QMargins>(static_cast<QPageLayout*>(ptr)->marginsPoints()).top(), static_cast<QMargins>(static_cast<QPageLayout*>(ptr)->marginsPoints()).right(), static_cast<QMargins>(static_cast<QPageLayout*>(ptr)->marginsPoints()).bottom());
}

void* QPageLayout_MaximumMargins(void* ptr)
{
	return new QMarginsF(static_cast<QMarginsF>(static_cast<QPageLayout*>(ptr)->maximumMargins()).left(), static_cast<QMarginsF>(static_cast<QPageLayout*>(ptr)->maximumMargins()).top(), static_cast<QMarginsF>(static_cast<QPageLayout*>(ptr)->maximumMargins()).right(), static_cast<QMarginsF>(static_cast<QPageLayout*>(ptr)->maximumMargins()).bottom());
}

void* QPageLayout_MinimumMargins(void* ptr)
{
	return new QMarginsF(static_cast<QMarginsF>(static_cast<QPageLayout*>(ptr)->minimumMargins()).left(), static_cast<QMarginsF>(static_cast<QPageLayout*>(ptr)->minimumMargins()).top(), static_cast<QMarginsF>(static_cast<QPageLayout*>(ptr)->minimumMargins()).right(), static_cast<QMarginsF>(static_cast<QPageLayout*>(ptr)->minimumMargins()).bottom());
}

int QPageLayout_Mode(void* ptr)
{
	return static_cast<QPageLayout*>(ptr)->mode();
}

int QPageLayout_Orientation(void* ptr)
{
	return static_cast<QPageLayout*>(ptr)->orientation();
}

void* QPageLayout_PageSize(void* ptr)
{
	return new QPageSize(static_cast<QPageLayout*>(ptr)->pageSize());
}

void* QPageLayout_PaintRect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QPageLayout*>(ptr)->paintRect()).x(), static_cast<QRectF>(static_cast<QPageLayout*>(ptr)->paintRect()).y(), static_cast<QRectF>(static_cast<QPageLayout*>(ptr)->paintRect()).width(), static_cast<QRectF>(static_cast<QPageLayout*>(ptr)->paintRect()).height());
}

void* QPageLayout_PaintRect2(void* ptr, int units)
{
	return new QRectF(static_cast<QRectF>(static_cast<QPageLayout*>(ptr)->paintRect(static_cast<QPageLayout::Unit>(units))).x(), static_cast<QRectF>(static_cast<QPageLayout*>(ptr)->paintRect(static_cast<QPageLayout::Unit>(units))).y(), static_cast<QRectF>(static_cast<QPageLayout*>(ptr)->paintRect(static_cast<QPageLayout::Unit>(units))).width(), static_cast<QRectF>(static_cast<QPageLayout*>(ptr)->paintRect(static_cast<QPageLayout::Unit>(units))).height());
}

void* QPageLayout_PaintRectPixels(void* ptr, int resolution)
{
	return new QRect(static_cast<QRect>(static_cast<QPageLayout*>(ptr)->paintRectPixels(resolution)).x(), static_cast<QRect>(static_cast<QPageLayout*>(ptr)->paintRectPixels(resolution)).y(), static_cast<QRect>(static_cast<QPageLayout*>(ptr)->paintRectPixels(resolution)).width(), static_cast<QRect>(static_cast<QPageLayout*>(ptr)->paintRectPixels(resolution)).height());
}

void* QPageLayout_PaintRectPoints(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QPageLayout*>(ptr)->paintRectPoints()).x(), static_cast<QRect>(static_cast<QPageLayout*>(ptr)->paintRectPoints()).y(), static_cast<QRect>(static_cast<QPageLayout*>(ptr)->paintRectPoints()).width(), static_cast<QRect>(static_cast<QPageLayout*>(ptr)->paintRectPoints()).height());
}

int QPageLayout_SetBottomMargin(void* ptr, double bottomMargin)
{
	return static_cast<QPageLayout*>(ptr)->setBottomMargin(static_cast<double>(bottomMargin));
}

int QPageLayout_SetLeftMargin(void* ptr, double leftMargin)
{
	return static_cast<QPageLayout*>(ptr)->setLeftMargin(static_cast<double>(leftMargin));
}

void QPageLayout_SetMinimumMargins(void* ptr, void* minMargins)
{
	static_cast<QPageLayout*>(ptr)->setMinimumMargins(*static_cast<QMarginsF*>(minMargins));
}

void QPageLayout_SetMode(void* ptr, int mode)
{
	static_cast<QPageLayout*>(ptr)->setMode(static_cast<QPageLayout::Mode>(mode));
}

void QPageLayout_SetOrientation(void* ptr, int orientation)
{
	static_cast<QPageLayout*>(ptr)->setOrientation(static_cast<QPageLayout::Orientation>(orientation));
}

void QPageLayout_SetPageSize(void* ptr, void* pageSize, void* minMargins)
{
	static_cast<QPageLayout*>(ptr)->setPageSize(*static_cast<QPageSize*>(pageSize), *static_cast<QMarginsF*>(minMargins));
}

int QPageLayout_SetRightMargin(void* ptr, double rightMargin)
{
	return static_cast<QPageLayout*>(ptr)->setRightMargin(static_cast<double>(rightMargin));
}

int QPageLayout_SetTopMargin(void* ptr, double topMargin)
{
	return static_cast<QPageLayout*>(ptr)->setTopMargin(static_cast<double>(topMargin));
}

void QPageLayout_SetUnits(void* ptr, int units)
{
	static_cast<QPageLayout*>(ptr)->setUnits(static_cast<QPageLayout::Unit>(units));
}

void QPageLayout_Swap(void* ptr, void* other)
{
	static_cast<QPageLayout*>(ptr)->swap(*static_cast<QPageLayout*>(other));
}

int QPageLayout_Units(void* ptr)
{
	return static_cast<QPageLayout*>(ptr)->units();
}

void QPageLayout_DestroyQPageLayout(void* ptr)
{
	static_cast<QPageLayout*>(ptr)->~QPageLayout();
}

void* QPageSize_NewQPageSize()
{
	return new QPageSize();
}

void* QPageSize_NewQPageSize2(int pageSize)
{
	return new QPageSize(static_cast<QPageSize::PageSizeId>(pageSize));
}

void* QPageSize_NewQPageSize5(void* other)
{
	return new QPageSize(*static_cast<QPageSize*>(other));
}

void* QPageSize_NewQPageSize3(void* pointSize, char* name, int matchPolicy)
{
	return new QPageSize(*static_cast<QSize*>(pointSize), QString(name), static_cast<QPageSize::SizeMatchPolicy>(matchPolicy));
}

void* QPageSize_NewQPageSize4(void* size, int units, char* name, int matchPolicy)
{
	return new QPageSize(*static_cast<QSizeF*>(size), static_cast<QPageSize::Unit>(units), QString(name), static_cast<QPageSize::SizeMatchPolicy>(matchPolicy));
}

void* QPageSize_QPageSize_DefinitionSize2(int pageSizeId)
{
	return new QSizeF(static_cast<QSizeF>(QPageSize::definitionSize(static_cast<QPageSize::PageSizeId>(pageSizeId))).width(), static_cast<QSizeF>(QPageSize::definitionSize(static_cast<QPageSize::PageSizeId>(pageSizeId))).height());
}

void* QPageSize_DefinitionSize(void* ptr)
{
	return new QSizeF(static_cast<QSizeF>(static_cast<QPageSize*>(ptr)->definitionSize()).width(), static_cast<QSizeF>(static_cast<QPageSize*>(ptr)->definitionSize()).height());
}

int QPageSize_QPageSize_DefinitionUnits2(int pageSizeId)
{
	return QPageSize::definitionUnits(static_cast<QPageSize::PageSizeId>(pageSizeId));
}

int QPageSize_DefinitionUnits(void* ptr)
{
	return static_cast<QPageSize*>(ptr)->definitionUnits();
}

int QPageSize_QPageSize_Id2(void* pointSize, int matchPolicy)
{
	return QPageSize::id(*static_cast<QSize*>(pointSize), static_cast<QPageSize::SizeMatchPolicy>(matchPolicy));
}

int QPageSize_QPageSize_Id3(void* size, int units, int matchPolicy)
{
	return QPageSize::id(*static_cast<QSizeF*>(size), static_cast<QPageSize::Unit>(units), static_cast<QPageSize::SizeMatchPolicy>(matchPolicy));
}

int QPageSize_QPageSize_Id4(int windowsId)
{
	return QPageSize::id(windowsId);
}

int QPageSize_Id(void* ptr)
{
	return static_cast<QPageSize*>(ptr)->id();
}

int QPageSize_IsEquivalentTo(void* ptr, void* other)
{
	return static_cast<QPageSize*>(ptr)->isEquivalentTo(*static_cast<QPageSize*>(other));
}

int QPageSize_IsValid(void* ptr)
{
	return static_cast<QPageSize*>(ptr)->isValid();
}

char* QPageSize_QPageSize_Key2(int pageSizeId)
{
	return QPageSize::key(static_cast<QPageSize::PageSizeId>(pageSizeId)).toUtf8().data();
}

char* QPageSize_Key(void* ptr)
{
	return static_cast<QPageSize*>(ptr)->key().toUtf8().data();
}

char* QPageSize_QPageSize_Name2(int pageSizeId)
{
	return QPageSize::name(static_cast<QPageSize::PageSizeId>(pageSizeId)).toUtf8().data();
}

char* QPageSize_Name(void* ptr)
{
	return static_cast<QPageSize*>(ptr)->name().toUtf8().data();
}

void* QPageSize_Rect(void* ptr, int units)
{
	return new QRectF(static_cast<QRectF>(static_cast<QPageSize*>(ptr)->rect(static_cast<QPageSize::Unit>(units))).x(), static_cast<QRectF>(static_cast<QPageSize*>(ptr)->rect(static_cast<QPageSize::Unit>(units))).y(), static_cast<QRectF>(static_cast<QPageSize*>(ptr)->rect(static_cast<QPageSize::Unit>(units))).width(), static_cast<QRectF>(static_cast<QPageSize*>(ptr)->rect(static_cast<QPageSize::Unit>(units))).height());
}

void* QPageSize_RectPixels(void* ptr, int resolution)
{
	return new QRect(static_cast<QRect>(static_cast<QPageSize*>(ptr)->rectPixels(resolution)).x(), static_cast<QRect>(static_cast<QPageSize*>(ptr)->rectPixels(resolution)).y(), static_cast<QRect>(static_cast<QPageSize*>(ptr)->rectPixels(resolution)).width(), static_cast<QRect>(static_cast<QPageSize*>(ptr)->rectPixels(resolution)).height());
}

void* QPageSize_RectPoints(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QPageSize*>(ptr)->rectPoints()).x(), static_cast<QRect>(static_cast<QPageSize*>(ptr)->rectPoints()).y(), static_cast<QRect>(static_cast<QPageSize*>(ptr)->rectPoints()).width(), static_cast<QRect>(static_cast<QPageSize*>(ptr)->rectPoints()).height());
}

void* QPageSize_QPageSize_Size2(int pageSizeId, int units)
{
	return new QSizeF(static_cast<QSizeF>(QPageSize::size(static_cast<QPageSize::PageSizeId>(pageSizeId), static_cast<QPageSize::Unit>(units))).width(), static_cast<QSizeF>(QPageSize::size(static_cast<QPageSize::PageSizeId>(pageSizeId), static_cast<QPageSize::Unit>(units))).height());
}

void* QPageSize_Size(void* ptr, int units)
{
	return new QSizeF(static_cast<QSizeF>(static_cast<QPageSize*>(ptr)->size(static_cast<QPageSize::Unit>(units))).width(), static_cast<QSizeF>(static_cast<QPageSize*>(ptr)->size(static_cast<QPageSize::Unit>(units))).height());
}

void* QPageSize_QPageSize_SizePixels2(int pageSizeId, int resolution)
{
	return new QSize(static_cast<QSize>(QPageSize::sizePixels(static_cast<QPageSize::PageSizeId>(pageSizeId), resolution)).width(), static_cast<QSize>(QPageSize::sizePixels(static_cast<QPageSize::PageSizeId>(pageSizeId), resolution)).height());
}

void* QPageSize_SizePixels(void* ptr, int resolution)
{
	return new QSize(static_cast<QSize>(static_cast<QPageSize*>(ptr)->sizePixels(resolution)).width(), static_cast<QSize>(static_cast<QPageSize*>(ptr)->sizePixels(resolution)).height());
}

void* QPageSize_QPageSize_SizePoints2(int pageSizeId)
{
	return new QSize(static_cast<QSize>(QPageSize::sizePoints(static_cast<QPageSize::PageSizeId>(pageSizeId))).width(), static_cast<QSize>(QPageSize::sizePoints(static_cast<QPageSize::PageSizeId>(pageSizeId))).height());
}

void* QPageSize_SizePoints(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QPageSize*>(ptr)->sizePoints()).width(), static_cast<QSize>(static_cast<QPageSize*>(ptr)->sizePoints()).height());
}

void QPageSize_Swap(void* ptr, void* other)
{
	static_cast<QPageSize*>(ptr)->swap(*static_cast<QPageSize*>(other));
}

int QPageSize_QPageSize_WindowsId2(int pageSizeId)
{
	return QPageSize::windowsId(static_cast<QPageSize::PageSizeId>(pageSizeId));
}

int QPageSize_WindowsId(void* ptr)
{
	return static_cast<QPageSize*>(ptr)->windowsId();
}

void QPageSize_DestroyQPageSize(void* ptr)
{
	static_cast<QPageSize*>(ptr)->~QPageSize();
}

class MyQPagedPaintDevice: public QPagedPaintDevice
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQPagedPaintDevice() : QPagedPaintDevice() {};
	bool newPage() { return callbackQPagedPaintDevice_NewPage(this, this->objectNameAbs().toUtf8().data()) != 0; };
	void setPageSize(QPagedPaintDevice::PageSize size) { callbackQPagedPaintDevice_SetPageSize2(this, this->objectNameAbs().toUtf8().data(), size); };
	void setPageSizeMM(const QSizeF & size) { callbackQPagedPaintDevice_SetPageSizeMM(this, this->objectNameAbs().toUtf8().data(), new QSizeF(static_cast<QSizeF>(size).width(), static_cast<QSizeF>(size).height())); };
	int metric(QPaintDevice::PaintDeviceMetric metric) const { return callbackQPagedPaintDevice_Metric(const_cast<MyQPagedPaintDevice*>(this), this->objectNameAbs().toUtf8().data(), metric); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPagedPaintDevice_PaintEngine(const_cast<MyQPagedPaintDevice*>(this), this->objectNameAbs().toUtf8().data())); };
};

void* QPagedPaintDevice_NewQPagedPaintDevice()
{
	return new MyQPagedPaintDevice();
}

int QPagedPaintDevice_NewPage(void* ptr)
{
	return static_cast<QPagedPaintDevice*>(ptr)->newPage();
}

void* QPagedPaintDevice_PageLayout(void* ptr)
{
	return new QPageLayout(static_cast<QPagedPaintDevice*>(ptr)->pageLayout());
}

int QPagedPaintDevice_PageSize(void* ptr)
{
	return static_cast<QPagedPaintDevice*>(ptr)->pageSize();
}

void* QPagedPaintDevice_PageSizeMM(void* ptr)
{
	return new QSizeF(static_cast<QSizeF>(static_cast<QPagedPaintDevice*>(ptr)->pageSizeMM()).width(), static_cast<QSizeF>(static_cast<QPagedPaintDevice*>(ptr)->pageSizeMM()).height());
}

int QPagedPaintDevice_SetPageLayout(void* ptr, void* newPageLayout)
{
	return static_cast<QPagedPaintDevice*>(ptr)->setPageLayout(*static_cast<QPageLayout*>(newPageLayout));
}

int QPagedPaintDevice_SetPageMargins(void* ptr, void* margins)
{
	return static_cast<QPagedPaintDevice*>(ptr)->setPageMargins(*static_cast<QMarginsF*>(margins));
}

int QPagedPaintDevice_SetPageMargins2(void* ptr, void* margins, int units)
{
	return static_cast<QPagedPaintDevice*>(ptr)->setPageMargins(*static_cast<QMarginsF*>(margins), static_cast<QPageLayout::Unit>(units));
}

int QPagedPaintDevice_SetPageOrientation(void* ptr, int orientation)
{
	return static_cast<QPagedPaintDevice*>(ptr)->setPageOrientation(static_cast<QPageLayout::Orientation>(orientation));
}

int QPagedPaintDevice_SetPageSize(void* ptr, void* pageSize)
{
	return static_cast<QPagedPaintDevice*>(ptr)->setPageSize(*static_cast<QPageSize*>(pageSize));
}

void QPagedPaintDevice_SetPageSize2(void* ptr, int size)
{
	static_cast<QPagedPaintDevice*>(ptr)->setPageSize(static_cast<QPagedPaintDevice::PageSize>(size));
}

void QPagedPaintDevice_SetPageSize2Default(void* ptr, int size)
{
	static_cast<QPagedPaintDevice*>(ptr)->QPagedPaintDevice::setPageSize(static_cast<QPagedPaintDevice::PageSize>(size));
}

void QPagedPaintDevice_SetPageSizeMM(void* ptr, void* size)
{
	static_cast<QPagedPaintDevice*>(ptr)->setPageSizeMM(*static_cast<QSizeF*>(size));
}

void QPagedPaintDevice_SetPageSizeMMDefault(void* ptr, void* size)
{
	static_cast<QPagedPaintDevice*>(ptr)->QPagedPaintDevice::setPageSizeMM(*static_cast<QSizeF*>(size));
}

void QPagedPaintDevice_DestroyQPagedPaintDevice(void* ptr)
{
	static_cast<QPagedPaintDevice*>(ptr)->~QPagedPaintDevice();
}

char* QPagedPaintDevice_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQPagedPaintDevice*>(static_cast<QPagedPaintDevice*>(ptr))) {
		return static_cast<MyQPagedPaintDevice*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QPagedPaintDevice_BASE").toUtf8().data();
}

void QPagedPaintDevice_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQPagedPaintDevice*>(static_cast<QPagedPaintDevice*>(ptr))) {
		static_cast<MyQPagedPaintDevice*>(ptr)->setObjectNameAbs(QString(name));
	}
}

int QPagedPaintDevice_Metric(void* ptr, int metric)
{
	return static_cast<QPagedPaintDevice*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

int QPagedPaintDevice_MetricDefault(void* ptr, int metric)
{
	return static_cast<QPagedPaintDevice*>(ptr)->QPagedPaintDevice::metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

void* QPagedPaintDevice_PaintEngine(void* ptr)
{
	return static_cast<QPagedPaintDevice*>(ptr)->paintEngine();
}



class MyQPaintDevice: public QPaintDevice
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQPaintDevice() : QPaintDevice() {};
	int metric(QPaintDevice::PaintDeviceMetric metric) const { return callbackQPaintDevice_Metric(const_cast<MyQPaintDevice*>(this), this->objectNameAbs().toUtf8().data(), metric); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPaintDevice_PaintEngine(const_cast<MyQPaintDevice*>(this), this->objectNameAbs().toUtf8().data())); };
};

void* QPaintDevice_NewQPaintDevice()
{
	return new MyQPaintDevice();
}

int QPaintDevice_Metric(void* ptr, int metric)
{
	return static_cast<QPaintDevice*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

int QPaintDevice_MetricDefault(void* ptr, int metric)
{
	return static_cast<QPaintDevice*>(ptr)->QPaintDevice::metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

void QPaintDevice_DestroyQPaintDevice(void* ptr)
{
	static_cast<QPaintDevice*>(ptr)->~QPaintDevice();
}

int QPaintDevice_ColorCount(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->colorCount();
}

int QPaintDevice_Depth(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->depth();
}

int QPaintDevice_DevicePixelRatio(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->devicePixelRatio();
}

int QPaintDevice_Height(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->height();
}

int QPaintDevice_HeightMM(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->heightMM();
}

int QPaintDevice_LogicalDpiX(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->logicalDpiX();
}

int QPaintDevice_LogicalDpiY(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->logicalDpiY();
}

void* QPaintDevice_PaintEngine(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->paintEngine();
}

int QPaintDevice_PaintingActive(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->paintingActive();
}

int QPaintDevice_PhysicalDpiX(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->physicalDpiX();
}

int QPaintDevice_PhysicalDpiY(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->physicalDpiY();
}

int QPaintDevice_Width(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->width();
}

int QPaintDevice_WidthMM(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->widthMM();
}

char* QPaintDevice_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQPaintDevice*>(static_cast<QPaintDevice*>(ptr))) {
		return static_cast<MyQPaintDevice*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QPaintDevice_BASE").toUtf8().data();
}

void QPaintDevice_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQPaintDevice*>(static_cast<QPaintDevice*>(ptr))) {
		static_cast<MyQPaintDevice*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQPaintDeviceWindow: public QPaintDeviceWindow
{
public:
	void paintEvent(QPaintEvent * event) { callbackQPaintDeviceWindow_PaintEvent(this, this->objectName().toUtf8().data(), event); };
	void update() { callbackQPaintDeviceWindow_Update3(this, this->objectName().toUtf8().data()); };
	void setHeight(int arg) { callbackQPaintDeviceWindow_SetHeight(this, this->objectName().toUtf8().data(), arg); };
	void setMaximumHeight(int h) { callbackQPaintDeviceWindow_SetMaximumHeight(this, this->objectName().toUtf8().data(), h); };
	void setMaximumWidth(int w) { callbackQPaintDeviceWindow_SetMaximumWidth(this, this->objectName().toUtf8().data(), w); };
	void setMinimumHeight(int h) { callbackQPaintDeviceWindow_SetMinimumHeight(this, this->objectName().toUtf8().data(), h); };
	void setMinimumWidth(int w) { callbackQPaintDeviceWindow_SetMinimumWidth(this, this->objectName().toUtf8().data(), w); };
	void setTitle(const QString & vqs) { callbackQPaintDeviceWindow_SetTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void setVisible(bool visible) { callbackQPaintDeviceWindow_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWidth(int arg) { callbackQPaintDeviceWindow_SetWidth(this, this->objectName().toUtf8().data(), arg); };
	void setX(int arg) { callbackQPaintDeviceWindow_SetX(this, this->objectName().toUtf8().data(), arg); };
	void setY(int arg) { callbackQPaintDeviceWindow_SetY(this, this->objectName().toUtf8().data(), arg); };
	void alert(int msec) { callbackQPaintDeviceWindow_Alert(this, this->objectName().toUtf8().data(), msec); };
	bool close() { return callbackQPaintDeviceWindow_Close(this, this->objectName().toUtf8().data()) != 0; };
	bool event(QEvent * ev) { return callbackQPaintDeviceWindow_Event(this, this->objectName().toUtf8().data(), ev) != 0; };
	void exposeEvent(QExposeEvent * ev) { callbackQPaintDeviceWindow_ExposeEvent(this, this->objectName().toUtf8().data(), ev); };
	void focusInEvent(QFocusEvent * ev) { callbackQPaintDeviceWindow_FocusInEvent(this, this->objectName().toUtf8().data(), ev); };
	QObject * focusObject() const { return static_cast<QObject*>(callbackQPaintDeviceWindow_FocusObject(const_cast<MyQPaintDeviceWindow*>(this), this->objectName().toUtf8().data())); };
	void focusOutEvent(QFocusEvent * ev) { callbackQPaintDeviceWindow_FocusOutEvent(this, this->objectName().toUtf8().data(), ev); };
	QSurfaceFormat format() const { return *static_cast<QSurfaceFormat*>(callbackQPaintDeviceWindow_Format(const_cast<MyQPaintDeviceWindow*>(this), this->objectName().toUtf8().data())); };
	void hide() { callbackQPaintDeviceWindow_Hide(this, this->objectName().toUtf8().data()); };
	void hideEvent(QHideEvent * ev) { callbackQPaintDeviceWindow_HideEvent(this, this->objectName().toUtf8().data(), ev); };
	void keyPressEvent(QKeyEvent * ev) { callbackQPaintDeviceWindow_KeyPressEvent(this, this->objectName().toUtf8().data(), ev); };
	void keyReleaseEvent(QKeyEvent * ev) { callbackQPaintDeviceWindow_KeyReleaseEvent(this, this->objectName().toUtf8().data(), ev); };
	void lower() { callbackQPaintDeviceWindow_Lower(this, this->objectName().toUtf8().data()); };
	void mouseDoubleClickEvent(QMouseEvent * ev) { callbackQPaintDeviceWindow_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), ev); };
	void mouseMoveEvent(QMouseEvent * ev) { callbackQPaintDeviceWindow_MouseMoveEvent(this, this->objectName().toUtf8().data(), ev); };
	void mousePressEvent(QMouseEvent * ev) { callbackQPaintDeviceWindow_MousePressEvent(this, this->objectName().toUtf8().data(), ev); };
	void mouseReleaseEvent(QMouseEvent * ev) { callbackQPaintDeviceWindow_MouseReleaseEvent(this, this->objectName().toUtf8().data(), ev); };
	void moveEvent(QMoveEvent * ev) { callbackQPaintDeviceWindow_MoveEvent(this, this->objectName().toUtf8().data(), ev); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQPaintDeviceWindow_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void raise() { callbackQPaintDeviceWindow_Raise(this, this->objectName().toUtf8().data()); };
	void requestActivate() { callbackQPaintDeviceWindow_RequestActivate(this, this->objectName().toUtf8().data()); };
	void requestUpdate() { callbackQPaintDeviceWindow_RequestUpdate(this, this->objectName().toUtf8().data()); };
	void resizeEvent(QResizeEvent * ev) { callbackQPaintDeviceWindow_ResizeEvent(this, this->objectName().toUtf8().data(), ev); };
	void show() { callbackQPaintDeviceWindow_Show(this, this->objectName().toUtf8().data()); };
	void showEvent(QShowEvent * ev) { callbackQPaintDeviceWindow_ShowEvent(this, this->objectName().toUtf8().data(), ev); };
	void showFullScreen() { callbackQPaintDeviceWindow_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackQPaintDeviceWindow_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackQPaintDeviceWindow_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackQPaintDeviceWindow_ShowNormal(this, this->objectName().toUtf8().data()); };
	QSize size() const { return *static_cast<QSize*>(callbackQPaintDeviceWindow_Size(const_cast<MyQPaintDeviceWindow*>(this), this->objectName().toUtf8().data())); };
	SurfaceType surfaceType() const { return static_cast<QSurface::SurfaceType>(callbackQPaintDeviceWindow_SurfaceType(const_cast<MyQPaintDeviceWindow*>(this), this->objectName().toUtf8().data())); };
	void tabletEvent(QTabletEvent * ev) { callbackQPaintDeviceWindow_TabletEvent(this, this->objectName().toUtf8().data(), ev); };
	void touchEvent(QTouchEvent * ev) { callbackQPaintDeviceWindow_TouchEvent(this, this->objectName().toUtf8().data(), ev); };
	void wheelEvent(QWheelEvent * ev) { callbackQPaintDeviceWindow_WheelEvent(this, this->objectName().toUtf8().data(), ev); };
	void timerEvent(QTimerEvent * event) { callbackQPaintDeviceWindow_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQPaintDeviceWindow_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQPaintDeviceWindow_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQPaintDeviceWindow_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQPaintDeviceWindow_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQPaintDeviceWindow_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQPaintDeviceWindow_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQPaintDeviceWindow_MetaObject(const_cast<MyQPaintDeviceWindow*>(this), this->objectName().toUtf8().data())); };
	int metric(QPaintDevice::PaintDeviceMetric metric) const { return callbackQPaintDeviceWindow_Metric(const_cast<MyQPaintDeviceWindow*>(this), this->objectName().toUtf8().data(), metric); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPaintDeviceWindow_PaintEngine(const_cast<MyQPaintDeviceWindow*>(this), this->objectName().toUtf8().data())); };
};

void QPaintDeviceWindow_PaintEvent(void* ptr, void* event)
{
	static_cast<QPaintDeviceWindow*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QPaintDeviceWindow_PaintEventDefault(void* ptr, void* event)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::paintEvent(static_cast<QPaintEvent*>(event));
}

void QPaintDeviceWindow_Update3(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "update");
}

void QPaintDeviceWindow_Update(void* ptr, void* rect)
{
	static_cast<QPaintDeviceWindow*>(ptr)->update(*static_cast<QRect*>(rect));
}

void QPaintDeviceWindow_Update2(void* ptr, void* region)
{
	static_cast<QPaintDeviceWindow*>(ptr)->update(*static_cast<QRegion*>(region));
}

void QPaintDeviceWindow_SetHeight(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "setHeight", Q_ARG(int, arg));
}

void QPaintDeviceWindow_SetHeightDefault(void* ptr, int arg)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::setHeight(arg);
}

void QPaintDeviceWindow_SetMaximumHeight(void* ptr, int h)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "setMaximumHeight", Q_ARG(int, h));
}

void QPaintDeviceWindow_SetMaximumHeightDefault(void* ptr, int h)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::setMaximumHeight(h);
}

void QPaintDeviceWindow_SetMaximumWidth(void* ptr, int w)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "setMaximumWidth", Q_ARG(int, w));
}

void QPaintDeviceWindow_SetMaximumWidthDefault(void* ptr, int w)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::setMaximumWidth(w);
}

void QPaintDeviceWindow_SetMinimumHeight(void* ptr, int h)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "setMinimumHeight", Q_ARG(int, h));
}

void QPaintDeviceWindow_SetMinimumHeightDefault(void* ptr, int h)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::setMinimumHeight(h);
}

void QPaintDeviceWindow_SetMinimumWidth(void* ptr, int w)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "setMinimumWidth", Q_ARG(int, w));
}

void QPaintDeviceWindow_SetMinimumWidthDefault(void* ptr, int w)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::setMinimumWidth(w);
}

void QPaintDeviceWindow_SetTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "setTitle", Q_ARG(QString, QString(vqs)));
}

void QPaintDeviceWindow_SetTitleDefault(void* ptr, char* vqs)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::setTitle(QString(vqs));
}

void QPaintDeviceWindow_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QPaintDeviceWindow_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::setVisible(visible != 0);
}

void QPaintDeviceWindow_SetWidth(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "setWidth", Q_ARG(int, arg));
}

void QPaintDeviceWindow_SetWidthDefault(void* ptr, int arg)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::setWidth(arg);
}

void QPaintDeviceWindow_SetX(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "setX", Q_ARG(int, arg));
}

void QPaintDeviceWindow_SetXDefault(void* ptr, int arg)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::setX(arg);
}

void QPaintDeviceWindow_SetY(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "setY", Q_ARG(int, arg));
}

void QPaintDeviceWindow_SetYDefault(void* ptr, int arg)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::setY(arg);
}

void QPaintDeviceWindow_Alert(void* ptr, int msec)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "alert", Q_ARG(int, msec));
}

void QPaintDeviceWindow_AlertDefault(void* ptr, int msec)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::alert(msec);
}

int QPaintDeviceWindow_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QPaintDeviceWindow_CloseDefault(void* ptr)
{
	return static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::close();
}

int QPaintDeviceWindow_Event(void* ptr, void* ev)
{
	return static_cast<QPaintDeviceWindow*>(ptr)->event(static_cast<QEvent*>(ev));
}

int QPaintDeviceWindow_EventDefault(void* ptr, void* ev)
{
	return static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::event(static_cast<QEvent*>(ev));
}

void QPaintDeviceWindow_ExposeEvent(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->exposeEvent(static_cast<QExposeEvent*>(ev));
}

void QPaintDeviceWindow_ExposeEventDefault(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::exposeEvent(static_cast<QExposeEvent*>(ev));
}

void QPaintDeviceWindow_FocusInEvent(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(ev));
}

void QPaintDeviceWindow_FocusInEventDefault(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::focusInEvent(static_cast<QFocusEvent*>(ev));
}

void* QPaintDeviceWindow_FocusObject(void* ptr)
{
	return static_cast<QPaintDeviceWindow*>(ptr)->focusObject();
}

void* QPaintDeviceWindow_FocusObjectDefault(void* ptr)
{
	return static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::focusObject();
}

void QPaintDeviceWindow_FocusOutEvent(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void QPaintDeviceWindow_FocusOutEventDefault(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void* QPaintDeviceWindow_Format(void* ptr)
{
	return new QSurfaceFormat(static_cast<QPaintDeviceWindow*>(ptr)->format());
}

void* QPaintDeviceWindow_FormatDefault(void* ptr)
{
	return new QSurfaceFormat(static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::format());
}

void QPaintDeviceWindow_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "hide");
}

void QPaintDeviceWindow_HideDefault(void* ptr)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::hide();
}

void QPaintDeviceWindow_HideEvent(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->hideEvent(static_cast<QHideEvent*>(ev));
}

void QPaintDeviceWindow_HideEventDefault(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::hideEvent(static_cast<QHideEvent*>(ev));
}

void QPaintDeviceWindow_KeyPressEvent(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(ev));
}

void QPaintDeviceWindow_KeyPressEventDefault(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::keyPressEvent(static_cast<QKeyEvent*>(ev));
}

void QPaintDeviceWindow_KeyReleaseEvent(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(ev));
}

void QPaintDeviceWindow_KeyReleaseEventDefault(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::keyReleaseEvent(static_cast<QKeyEvent*>(ev));
}

void QPaintDeviceWindow_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "lower");
}

void QPaintDeviceWindow_LowerDefault(void* ptr)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::lower();
}

void QPaintDeviceWindow_MouseDoubleClickEvent(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(ev));
}

void QPaintDeviceWindow_MouseDoubleClickEventDefault(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::mouseDoubleClickEvent(static_cast<QMouseEvent*>(ev));
}

void QPaintDeviceWindow_MouseMoveEvent(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(ev));
}

void QPaintDeviceWindow_MouseMoveEventDefault(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::mouseMoveEvent(static_cast<QMouseEvent*>(ev));
}

void QPaintDeviceWindow_MousePressEvent(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(ev));
}

void QPaintDeviceWindow_MousePressEventDefault(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::mousePressEvent(static_cast<QMouseEvent*>(ev));
}

void QPaintDeviceWindow_MouseReleaseEvent(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(ev));
}

void QPaintDeviceWindow_MouseReleaseEventDefault(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::mouseReleaseEvent(static_cast<QMouseEvent*>(ev));
}

void QPaintDeviceWindow_MoveEvent(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->moveEvent(static_cast<QMoveEvent*>(ev));
}

void QPaintDeviceWindow_MoveEventDefault(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::moveEvent(static_cast<QMoveEvent*>(ev));
}

int QPaintDeviceWindow_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QPaintDeviceWindow*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int QPaintDeviceWindow_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::nativeEvent(QByteArray(eventType), message, &result);
}

void QPaintDeviceWindow_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "raise");
}

void QPaintDeviceWindow_RaiseDefault(void* ptr)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::raise();
}

void QPaintDeviceWindow_RequestActivate(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "requestActivate");
}

void QPaintDeviceWindow_RequestActivateDefault(void* ptr)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::requestActivate();
}

void QPaintDeviceWindow_RequestUpdate(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "requestUpdate");
}

void QPaintDeviceWindow_RequestUpdateDefault(void* ptr)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::requestUpdate();
}

void QPaintDeviceWindow_ResizeEvent(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(ev));
}

void QPaintDeviceWindow_ResizeEventDefault(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::resizeEvent(static_cast<QResizeEvent*>(ev));
}

void QPaintDeviceWindow_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "show");
}

void QPaintDeviceWindow_ShowDefault(void* ptr)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::show();
}

void QPaintDeviceWindow_ShowEvent(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->showEvent(static_cast<QShowEvent*>(ev));
}

void QPaintDeviceWindow_ShowEventDefault(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::showEvent(static_cast<QShowEvent*>(ev));
}

void QPaintDeviceWindow_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "showFullScreen");
}

void QPaintDeviceWindow_ShowFullScreenDefault(void* ptr)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::showFullScreen();
}

void QPaintDeviceWindow_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "showMaximized");
}

void QPaintDeviceWindow_ShowMaximizedDefault(void* ptr)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::showMaximized();
}

void QPaintDeviceWindow_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "showMinimized");
}

void QPaintDeviceWindow_ShowMinimizedDefault(void* ptr)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::showMinimized();
}

void QPaintDeviceWindow_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "showNormal");
}

void QPaintDeviceWindow_ShowNormalDefault(void* ptr)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::showNormal();
}

void* QPaintDeviceWindow_Size(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QPaintDeviceWindow*>(ptr)->size()).width(), static_cast<QSize>(static_cast<QPaintDeviceWindow*>(ptr)->size()).height());
}

void* QPaintDeviceWindow_SizeDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::size()).width(), static_cast<QSize>(static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::size()).height());
}

int QPaintDeviceWindow_SurfaceType(void* ptr)
{
	return static_cast<QPaintDeviceWindow*>(ptr)->surfaceType();
}

int QPaintDeviceWindow_SurfaceTypeDefault(void* ptr)
{
	return static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::surfaceType();
}

void QPaintDeviceWindow_TabletEvent(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(ev));
}

void QPaintDeviceWindow_TabletEventDefault(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::tabletEvent(static_cast<QTabletEvent*>(ev));
}

void QPaintDeviceWindow_TouchEvent(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->touchEvent(static_cast<QTouchEvent*>(ev));
}

void QPaintDeviceWindow_TouchEventDefault(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::touchEvent(static_cast<QTouchEvent*>(ev));
}

void QPaintDeviceWindow_WheelEvent(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(ev));
}

void QPaintDeviceWindow_WheelEventDefault(void* ptr, void* ev)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::wheelEvent(static_cast<QWheelEvent*>(ev));
}

void QPaintDeviceWindow_TimerEvent(void* ptr, void* event)
{
	static_cast<QPaintDeviceWindow*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QPaintDeviceWindow_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::timerEvent(static_cast<QTimerEvent*>(event));
}

void QPaintDeviceWindow_ChildEvent(void* ptr, void* event)
{
	static_cast<QPaintDeviceWindow*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QPaintDeviceWindow_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::childEvent(static_cast<QChildEvent*>(event));
}

void QPaintDeviceWindow_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QPaintDeviceWindow*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPaintDeviceWindow_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPaintDeviceWindow_CustomEvent(void* ptr, void* event)
{
	static_cast<QPaintDeviceWindow*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QPaintDeviceWindow_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::customEvent(static_cast<QEvent*>(event));
}

void QPaintDeviceWindow_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPaintDeviceWindow*>(ptr), "deleteLater");
}

void QPaintDeviceWindow_DeleteLaterDefault(void* ptr)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::deleteLater();
}

void QPaintDeviceWindow_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QPaintDeviceWindow*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPaintDeviceWindow_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QPaintDeviceWindow_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QPaintDeviceWindow*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QPaintDeviceWindow_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QPaintDeviceWindow_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QPaintDeviceWindow*>(ptr)->metaObject());
}

void* QPaintDeviceWindow_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::metaObject());
}

int QPaintDeviceWindow_Metric(void* ptr, int metric)
{
	return static_cast<QPaintDeviceWindow*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

int QPaintDeviceWindow_MetricDefault(void* ptr, int metric)
{
	return static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

void* QPaintDeviceWindow_PaintEngine(void* ptr)
{
	return static_cast<QPaintDeviceWindow*>(ptr)->paintEngine();
}

void* QPaintDeviceWindow_PaintEngineDefault(void* ptr)
{
	return static_cast<QPaintDeviceWindow*>(ptr)->QPaintDeviceWindow::paintEngine();
}

class MyQPaintEngine: public QPaintEngine
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	void drawEllipse(const QRectF & rect) { callbackQPaintEngine_DrawEllipse(this, this->objectNameAbs().toUtf8().data(), new QRectF(static_cast<QRectF>(rect).x(), static_cast<QRectF>(rect).y(), static_cast<QRectF>(rect).width(), static_cast<QRectF>(rect).height())); };
	void drawImage(const QRectF & rectangle, const QImage & image, const QRectF & sr, Qt::ImageConversionFlags flags) { callbackQPaintEngine_DrawImage(this, this->objectNameAbs().toUtf8().data(), new QRectF(static_cast<QRectF>(rectangle).x(), static_cast<QRectF>(rectangle).y(), static_cast<QRectF>(rectangle).width(), static_cast<QRectF>(rectangle).height()), new QImage(image), new QRectF(static_cast<QRectF>(sr).x(), static_cast<QRectF>(sr).y(), static_cast<QRectF>(sr).width(), static_cast<QRectF>(sr).height()), flags); };
	void drawPolygon(const QPointF * points, int pointCount, QPaintEngine::PolygonDrawMode mode) { callbackQPaintEngine_DrawPolygon(this, this->objectNameAbs().toUtf8().data(), const_cast<QPointF*>(points), pointCount, mode); };
	bool begin(QPaintDevice * pdev) { return callbackQPaintEngine_Begin(this, this->objectNameAbs().toUtf8().data(), pdev) != 0; };
	void drawEllipse(const QRect & rect) { callbackQPaintEngine_DrawEllipse2(this, this->objectNameAbs().toUtf8().data(), new QRect(static_cast<QRect>(rect).x(), static_cast<QRect>(rect).y(), static_cast<QRect>(rect).width(), static_cast<QRect>(rect).height())); };
	void drawLines(const QLine * lines, int lineCount) { callbackQPaintEngine_DrawLines2(this, this->objectNameAbs().toUtf8().data(), const_cast<QLine*>(lines), lineCount); };
	void drawLines(const QLineF * lines, int lineCount) { callbackQPaintEngine_DrawLines(this, this->objectNameAbs().toUtf8().data(), const_cast<QLineF*>(lines), lineCount); };
	void drawPath(const QPainterPath & path) { callbackQPaintEngine_DrawPath(this, this->objectNameAbs().toUtf8().data(), new QPainterPath(path)); };
	void drawPixmap(const QRectF & r, const QPixmap & pm, const QRectF & sr) { callbackQPaintEngine_DrawPixmap(this, this->objectNameAbs().toUtf8().data(), new QRectF(static_cast<QRectF>(r).x(), static_cast<QRectF>(r).y(), static_cast<QRectF>(r).width(), static_cast<QRectF>(r).height()), new QPixmap(pm), new QRectF(static_cast<QRectF>(sr).x(), static_cast<QRectF>(sr).y(), static_cast<QRectF>(sr).width(), static_cast<QRectF>(sr).height())); };
	void drawPoints(const QPoint * points, int pointCount) { callbackQPaintEngine_DrawPoints2(this, this->objectNameAbs().toUtf8().data(), const_cast<QPoint*>(points), pointCount); };
	void drawPoints(const QPointF * points, int pointCount) { callbackQPaintEngine_DrawPoints(this, this->objectNameAbs().toUtf8().data(), const_cast<QPointF*>(points), pointCount); };
	void drawPolygon(const QPoint * points, int pointCount, QPaintEngine::PolygonDrawMode mode) { callbackQPaintEngine_DrawPolygon2(this, this->objectNameAbs().toUtf8().data(), const_cast<QPoint*>(points), pointCount, mode); };
	void drawRects(const QRect * rects, int rectCount) { callbackQPaintEngine_DrawRects2(this, this->objectNameAbs().toUtf8().data(), const_cast<QRect*>(rects), rectCount); };
	void drawRects(const QRectF * rects, int rectCount) { callbackQPaintEngine_DrawRects(this, this->objectNameAbs().toUtf8().data(), const_cast<QRectF*>(rects), rectCount); };
	void drawTiledPixmap(const QRectF & rect, const QPixmap & pixmap, const QPointF & p) { callbackQPaintEngine_DrawTiledPixmap(this, this->objectNameAbs().toUtf8().data(), new QRectF(static_cast<QRectF>(rect).x(), static_cast<QRectF>(rect).y(), static_cast<QRectF>(rect).width(), static_cast<QRectF>(rect).height()), new QPixmap(pixmap), new QPointF(static_cast<QPointF>(p).x(), static_cast<QPointF>(p).y())); };
	bool end() { return callbackQPaintEngine_End(this, this->objectNameAbs().toUtf8().data()) != 0; };
	Type type() const { return static_cast<QPaintEngine::Type>(callbackQPaintEngine_Type(const_cast<MyQPaintEngine*>(this), this->objectNameAbs().toUtf8().data())); };
};

void QPaintEngine_DrawEllipse(void* ptr, void* rect)
{
	static_cast<QPaintEngine*>(ptr)->drawEllipse(*static_cast<QRectF*>(rect));
}

void QPaintEngine_DrawEllipseDefault(void* ptr, void* rect)
{
	static_cast<QPaintEngine*>(ptr)->QPaintEngine::drawEllipse(*static_cast<QRectF*>(rect));
}

void QPaintEngine_DrawImage(void* ptr, void* rectangle, void* image, void* sr, int flags)
{
	static_cast<QPaintEngine*>(ptr)->drawImage(*static_cast<QRectF*>(rectangle), *static_cast<QImage*>(image), *static_cast<QRectF*>(sr), static_cast<Qt::ImageConversionFlag>(flags));
}

void QPaintEngine_DrawImageDefault(void* ptr, void* rectangle, void* image, void* sr, int flags)
{
	static_cast<QPaintEngine*>(ptr)->QPaintEngine::drawImage(*static_cast<QRectF*>(rectangle), *static_cast<QImage*>(image), *static_cast<QRectF*>(sr), static_cast<Qt::ImageConversionFlag>(flags));
}

void QPaintEngine_DrawPolygon(void* ptr, void* points, int pointCount, int mode)
{
	static_cast<QPaintEngine*>(ptr)->drawPolygon(static_cast<QPointF*>(points), pointCount, static_cast<QPaintEngine::PolygonDrawMode>(mode));
}

void QPaintEngine_DrawPolygonDefault(void* ptr, void* points, int pointCount, int mode)
{
	static_cast<QPaintEngine*>(ptr)->QPaintEngine::drawPolygon(static_cast<QPointF*>(points), pointCount, static_cast<QPaintEngine::PolygonDrawMode>(mode));
}

int QPaintEngine_Begin(void* ptr, void* pdev)
{
	return static_cast<QPaintEngine*>(ptr)->begin(static_cast<QPaintDevice*>(pdev));
}

void QPaintEngine_DrawEllipse2(void* ptr, void* rect)
{
	static_cast<QPaintEngine*>(ptr)->drawEllipse(*static_cast<QRect*>(rect));
}

void QPaintEngine_DrawEllipse2Default(void* ptr, void* rect)
{
	static_cast<QPaintEngine*>(ptr)->QPaintEngine::drawEllipse(*static_cast<QRect*>(rect));
}

void QPaintEngine_DrawLines2(void* ptr, void* lines, int lineCount)
{
	static_cast<QPaintEngine*>(ptr)->drawLines(static_cast<QLine*>(lines), lineCount);
}

void QPaintEngine_DrawLines2Default(void* ptr, void* lines, int lineCount)
{
	static_cast<QPaintEngine*>(ptr)->QPaintEngine::drawLines(static_cast<QLine*>(lines), lineCount);
}

void QPaintEngine_DrawLines(void* ptr, void* lines, int lineCount)
{
	static_cast<QPaintEngine*>(ptr)->drawLines(static_cast<QLineF*>(lines), lineCount);
}

void QPaintEngine_DrawLinesDefault(void* ptr, void* lines, int lineCount)
{
	static_cast<QPaintEngine*>(ptr)->QPaintEngine::drawLines(static_cast<QLineF*>(lines), lineCount);
}

void QPaintEngine_DrawPath(void* ptr, void* path)
{
	static_cast<QPaintEngine*>(ptr)->drawPath(*static_cast<QPainterPath*>(path));
}

void QPaintEngine_DrawPathDefault(void* ptr, void* path)
{
	static_cast<QPaintEngine*>(ptr)->QPaintEngine::drawPath(*static_cast<QPainterPath*>(path));
}

void QPaintEngine_DrawPixmap(void* ptr, void* r, void* pm, void* sr)
{
	static_cast<QPaintEngine*>(ptr)->drawPixmap(*static_cast<QRectF*>(r), *static_cast<QPixmap*>(pm), *static_cast<QRectF*>(sr));
}

void QPaintEngine_DrawPoints2(void* ptr, void* points, int pointCount)
{
	static_cast<QPaintEngine*>(ptr)->drawPoints(static_cast<QPoint*>(points), pointCount);
}

void QPaintEngine_DrawPoints2Default(void* ptr, void* points, int pointCount)
{
	static_cast<QPaintEngine*>(ptr)->QPaintEngine::drawPoints(static_cast<QPoint*>(points), pointCount);
}

void QPaintEngine_DrawPoints(void* ptr, void* points, int pointCount)
{
	static_cast<QPaintEngine*>(ptr)->drawPoints(static_cast<QPointF*>(points), pointCount);
}

void QPaintEngine_DrawPointsDefault(void* ptr, void* points, int pointCount)
{
	static_cast<QPaintEngine*>(ptr)->QPaintEngine::drawPoints(static_cast<QPointF*>(points), pointCount);
}

void QPaintEngine_DrawPolygon2(void* ptr, void* points, int pointCount, int mode)
{
	static_cast<QPaintEngine*>(ptr)->drawPolygon(static_cast<QPoint*>(points), pointCount, static_cast<QPaintEngine::PolygonDrawMode>(mode));
}

void QPaintEngine_DrawPolygon2Default(void* ptr, void* points, int pointCount, int mode)
{
	static_cast<QPaintEngine*>(ptr)->QPaintEngine::drawPolygon(static_cast<QPoint*>(points), pointCount, static_cast<QPaintEngine::PolygonDrawMode>(mode));
}

void QPaintEngine_DrawRects2(void* ptr, void* rects, int rectCount)
{
	static_cast<QPaintEngine*>(ptr)->drawRects(static_cast<QRect*>(rects), rectCount);
}

void QPaintEngine_DrawRects2Default(void* ptr, void* rects, int rectCount)
{
	static_cast<QPaintEngine*>(ptr)->QPaintEngine::drawRects(static_cast<QRect*>(rects), rectCount);
}

void QPaintEngine_DrawRects(void* ptr, void* rects, int rectCount)
{
	static_cast<QPaintEngine*>(ptr)->drawRects(static_cast<QRectF*>(rects), rectCount);
}

void QPaintEngine_DrawRectsDefault(void* ptr, void* rects, int rectCount)
{
	static_cast<QPaintEngine*>(ptr)->QPaintEngine::drawRects(static_cast<QRectF*>(rects), rectCount);
}

void QPaintEngine_DrawTiledPixmap(void* ptr, void* rect, void* pixmap, void* p)
{
	static_cast<QPaintEngine*>(ptr)->drawTiledPixmap(*static_cast<QRectF*>(rect), *static_cast<QPixmap*>(pixmap), *static_cast<QPointF*>(p));
}

void QPaintEngine_DrawTiledPixmapDefault(void* ptr, void* rect, void* pixmap, void* p)
{
	static_cast<QPaintEngine*>(ptr)->QPaintEngine::drawTiledPixmap(*static_cast<QRectF*>(rect), *static_cast<QPixmap*>(pixmap), *static_cast<QPointF*>(p));
}

int QPaintEngine_End(void* ptr)
{
	return static_cast<QPaintEngine*>(ptr)->end();
}

int QPaintEngine_HasFeature(void* ptr, int feature)
{
	return static_cast<QPaintEngine*>(ptr)->hasFeature(static_cast<QPaintEngine::PaintEngineFeature>(feature));
}

int QPaintEngine_IsActive(void* ptr)
{
	return static_cast<QPaintEngine*>(ptr)->isActive();
}

void* QPaintEngine_PaintDevice(void* ptr)
{
	return static_cast<QPaintEngine*>(ptr)->paintDevice();
}

void* QPaintEngine_Painter(void* ptr)
{
	return static_cast<QPaintEngine*>(ptr)->painter();
}

void QPaintEngine_SetActive(void* ptr, int state)
{
	static_cast<QPaintEngine*>(ptr)->setActive(state != 0);
}

int QPaintEngine_Type(void* ptr)
{
	return static_cast<QPaintEngine*>(ptr)->type();
}

void QPaintEngine_DestroyQPaintEngine(void* ptr)
{
	static_cast<QPaintEngine*>(ptr)->~QPaintEngine();
}

char* QPaintEngine_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQPaintEngine*>(static_cast<QPaintEngine*>(ptr))) {
		return static_cast<MyQPaintEngine*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QPaintEngine_BASE").toUtf8().data();
}

void QPaintEngine_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQPaintEngine*>(static_cast<QPaintEngine*>(ptr))) {
		static_cast<MyQPaintEngine*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QPaintEngineState_BackgroundBrush(void* ptr)
{
	return new QBrush(static_cast<QPaintEngineState*>(ptr)->backgroundBrush());
}

int QPaintEngineState_BackgroundMode(void* ptr)
{
	return static_cast<QPaintEngineState*>(ptr)->backgroundMode();
}

void* QPaintEngineState_Brush(void* ptr)
{
	return new QBrush(static_cast<QPaintEngineState*>(ptr)->brush());
}

int QPaintEngineState_BrushNeedsResolving(void* ptr)
{
	return static_cast<QPaintEngineState*>(ptr)->brushNeedsResolving();
}

void* QPaintEngineState_BrushOrigin(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QPaintEngineState*>(ptr)->brushOrigin()).x(), static_cast<QPointF>(static_cast<QPaintEngineState*>(ptr)->brushOrigin()).y());
}

int QPaintEngineState_ClipOperation(void* ptr)
{
	return static_cast<QPaintEngineState*>(ptr)->clipOperation();
}

void* QPaintEngineState_ClipPath(void* ptr)
{
	return new QPainterPath(static_cast<QPaintEngineState*>(ptr)->clipPath());
}

void* QPaintEngineState_ClipRegion(void* ptr)
{
	return new QRegion(static_cast<QPaintEngineState*>(ptr)->clipRegion());
}

int QPaintEngineState_CompositionMode(void* ptr)
{
	return static_cast<QPaintEngineState*>(ptr)->compositionMode();
}

void* QPaintEngineState_Font(void* ptr)
{
	return new QFont(static_cast<QPaintEngineState*>(ptr)->font());
}

int QPaintEngineState_IsClipEnabled(void* ptr)
{
	return static_cast<QPaintEngineState*>(ptr)->isClipEnabled();
}

double QPaintEngineState_Opacity(void* ptr)
{
	return static_cast<double>(static_cast<QPaintEngineState*>(ptr)->opacity());
}

void* QPaintEngineState_Painter(void* ptr)
{
	return static_cast<QPaintEngineState*>(ptr)->painter();
}

void* QPaintEngineState_Pen(void* ptr)
{
	return new QPen(static_cast<QPaintEngineState*>(ptr)->pen());
}

int QPaintEngineState_PenNeedsResolving(void* ptr)
{
	return static_cast<QPaintEngineState*>(ptr)->penNeedsResolving();
}

int QPaintEngineState_RenderHints(void* ptr)
{
	return static_cast<QPaintEngineState*>(ptr)->renderHints();
}

int QPaintEngineState_State(void* ptr)
{
	return static_cast<QPaintEngineState*>(ptr)->state();
}

void* QPaintEvent_NewQPaintEvent2(void* paintRect)
{
	return new QPaintEvent(*static_cast<QRect*>(paintRect));
}

void* QPaintEvent_NewQPaintEvent(void* paintRegion)
{
	return new QPaintEvent(*static_cast<QRegion*>(paintRegion));
}

void* QPaintEvent_Rect(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QPaintEvent*>(ptr)->rect()).x(), static_cast<QRect>(static_cast<QPaintEvent*>(ptr)->rect()).y(), static_cast<QRect>(static_cast<QPaintEvent*>(ptr)->rect()).width(), static_cast<QRect>(static_cast<QPaintEvent*>(ptr)->rect()).height());
}

void* QPaintEvent_Region(void* ptr)
{
	return new QRegion(static_cast<QPaintEvent*>(ptr)->region());
}

void* QPainter_NewQPainter2(void* device)
{
	return new QPainter(static_cast<QPaintDevice*>(device));
}

int QPainter_Begin(void* ptr, void* device)
{
	return static_cast<QPainter*>(ptr)->begin(static_cast<QPaintDevice*>(device));
}

void* QPainter_BoundingRect2(void* ptr, void* rectangle, int flags, char* text)
{
	return new QRect(static_cast<QRect>(static_cast<QPainter*>(ptr)->boundingRect(*static_cast<QRect*>(rectangle), flags, QString(text))).x(), static_cast<QRect>(static_cast<QPainter*>(ptr)->boundingRect(*static_cast<QRect*>(rectangle), flags, QString(text))).y(), static_cast<QRect>(static_cast<QPainter*>(ptr)->boundingRect(*static_cast<QRect*>(rectangle), flags, QString(text))).width(), static_cast<QRect>(static_cast<QPainter*>(ptr)->boundingRect(*static_cast<QRect*>(rectangle), flags, QString(text))).height());
}

void* QPainter_BoundingRect4(void* ptr, void* rectangle, char* text, void* option)
{
	return new QRectF(static_cast<QRectF>(static_cast<QPainter*>(ptr)->boundingRect(*static_cast<QRectF*>(rectangle), QString(text), *static_cast<QTextOption*>(option))).x(), static_cast<QRectF>(static_cast<QPainter*>(ptr)->boundingRect(*static_cast<QRectF*>(rectangle), QString(text), *static_cast<QTextOption*>(option))).y(), static_cast<QRectF>(static_cast<QPainter*>(ptr)->boundingRect(*static_cast<QRectF*>(rectangle), QString(text), *static_cast<QTextOption*>(option))).width(), static_cast<QRectF>(static_cast<QPainter*>(ptr)->boundingRect(*static_cast<QRectF*>(rectangle), QString(text), *static_cast<QTextOption*>(option))).height());
}

void* QPainter_BoundingRect(void* ptr, void* rectangle, int flags, char* text)
{
	return new QRectF(static_cast<QRectF>(static_cast<QPainter*>(ptr)->boundingRect(*static_cast<QRectF*>(rectangle), flags, QString(text))).x(), static_cast<QRectF>(static_cast<QPainter*>(ptr)->boundingRect(*static_cast<QRectF*>(rectangle), flags, QString(text))).y(), static_cast<QRectF>(static_cast<QPainter*>(ptr)->boundingRect(*static_cast<QRectF*>(rectangle), flags, QString(text))).width(), static_cast<QRectF>(static_cast<QPainter*>(ptr)->boundingRect(*static_cast<QRectF*>(rectangle), flags, QString(text))).height());
}

void QPainter_DrawArc(void* ptr, void* rectangle, int startAngle, int spanAngle)
{
	static_cast<QPainter*>(ptr)->drawArc(*static_cast<QRectF*>(rectangle), startAngle, spanAngle);
}

void QPainter_DrawChord(void* ptr, void* rectangle, int startAngle, int spanAngle)
{
	static_cast<QPainter*>(ptr)->drawChord(*static_cast<QRectF*>(rectangle), startAngle, spanAngle);
}

void QPainter_DrawConvexPolygon2(void* ptr, void* points, int pointCount)
{
	static_cast<QPainter*>(ptr)->drawConvexPolygon(static_cast<QPoint*>(points), pointCount);
}

void QPainter_DrawConvexPolygon(void* ptr, void* points, int pointCount)
{
	static_cast<QPainter*>(ptr)->drawConvexPolygon(static_cast<QPointF*>(points), pointCount);
}

void QPainter_DrawEllipse2(void* ptr, void* rectangle)
{
	static_cast<QPainter*>(ptr)->drawEllipse(*static_cast<QRect*>(rectangle));
}

void QPainter_DrawEllipse(void* ptr, void* rectangle)
{
	static_cast<QPainter*>(ptr)->drawEllipse(*static_cast<QRectF*>(rectangle));
}

void QPainter_DrawGlyphRun(void* ptr, void* position, void* glyphs)
{
	static_cast<QPainter*>(ptr)->drawGlyphRun(*static_cast<QPointF*>(position), *static_cast<QGlyphRun*>(glyphs));
}

void QPainter_DrawImage3(void* ptr, void* point, void* image)
{
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QPointF*>(point), *static_cast<QImage*>(image));
}

void QPainter_DrawImage(void* ptr, void* target, void* image, void* source, int flags)
{
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QRectF*>(target), *static_cast<QImage*>(image), *static_cast<QRectF*>(source), static_cast<Qt::ImageConversionFlag>(flags));
}

void QPainter_DrawLines2(void* ptr, void* lines, int lineCount)
{
	static_cast<QPainter*>(ptr)->drawLines(static_cast<QLine*>(lines), lineCount);
}

void QPainter_DrawPicture(void* ptr, void* point, void* picture)
{
	static_cast<QPainter*>(ptr)->drawPicture(*static_cast<QPointF*>(point), *static_cast<QPicture*>(picture));
}

void QPainter_DrawPie(void* ptr, void* rectangle, int startAngle, int spanAngle)
{
	static_cast<QPainter*>(ptr)->drawPie(*static_cast<QRectF*>(rectangle), startAngle, spanAngle);
}

void QPainter_DrawPixmap5(void* ptr, void* point, void* pixmap)
{
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QPointF*>(point), *static_cast<QPixmap*>(pixmap));
}

void QPainter_DrawPixmap(void* ptr, void* target, void* pixmap, void* source)
{
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QRectF*>(target), *static_cast<QPixmap*>(pixmap), *static_cast<QRectF*>(source));
}

void QPainter_DrawRects2(void* ptr, void* rectangles, int rectCount)
{
	static_cast<QPainter*>(ptr)->drawRects(static_cast<QRect*>(rectangles), rectCount);
}

void QPainter_DrawRects(void* ptr, void* rectangles, int rectCount)
{
	static_cast<QPainter*>(ptr)->drawRects(static_cast<QRectF*>(rectangles), rectCount);
}

void QPainter_DrawText(void* ptr, void* position, char* text)
{
	static_cast<QPainter*>(ptr)->drawText(*static_cast<QPointF*>(position), QString(text));
}

void QPainter_DrawText5(void* ptr, void* rectangle, int flags, char* text, void* boundingRect)
{
	static_cast<QPainter*>(ptr)->drawText(*static_cast<QRect*>(rectangle), flags, QString(text), static_cast<QRect*>(boundingRect));
}

void QPainter_DrawText8(void* ptr, void* rectangle, char* text, void* option)
{
	static_cast<QPainter*>(ptr)->drawText(*static_cast<QRectF*>(rectangle), QString(text), *static_cast<QTextOption*>(option));
}

void QPainter_DrawText4(void* ptr, void* rectangle, int flags, char* text, void* boundingRect)
{
	static_cast<QPainter*>(ptr)->drawText(*static_cast<QRectF*>(rectangle), flags, QString(text), static_cast<QRectF*>(boundingRect));
}

void QPainter_DrawTiledPixmap(void* ptr, void* rectangle, void* pixmap, void* position)
{
	static_cast<QPainter*>(ptr)->drawTiledPixmap(*static_cast<QRectF*>(rectangle), *static_cast<QPixmap*>(pixmap), *static_cast<QPointF*>(position));
}

void QPainter_EraseRect(void* ptr, void* rectangle)
{
	static_cast<QPainter*>(ptr)->eraseRect(*static_cast<QRectF*>(rectangle));
}

void QPainter_FillRect5(void* ptr, void* rectangle, void* brush)
{
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRect*>(rectangle), *static_cast<QBrush*>(brush));
}

void QPainter_FillRect6(void* ptr, void* rectangle, void* color)
{
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRect*>(rectangle), *static_cast<QColor*>(color));
}

void QPainter_FillRect(void* ptr, void* rectangle, void* brush)
{
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRectF*>(rectangle), *static_cast<QBrush*>(brush));
}

void QPainter_FillRect7(void* ptr, void* rectangle, void* color)
{
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRectF*>(rectangle), *static_cast<QColor*>(color));
}

void QPainter_Rotate(void* ptr, double angle)
{
	static_cast<QPainter*>(ptr)->rotate(static_cast<double>(angle));
}

void QPainter_SetBackground(void* ptr, void* brush)
{
	static_cast<QPainter*>(ptr)->setBackground(*static_cast<QBrush*>(brush));
}

void QPainter_SetBrushOrigin(void* ptr, void* position)
{
	static_cast<QPainter*>(ptr)->setBrushOrigin(*static_cast<QPointF*>(position));
}

void QPainter_SetClipPath(void* ptr, void* path, int operation)
{
	static_cast<QPainter*>(ptr)->setClipPath(*static_cast<QPainterPath*>(path), static_cast<Qt::ClipOperation>(operation));
}

void QPainter_SetClipRect3(void* ptr, void* rectangle, int operation)
{
	static_cast<QPainter*>(ptr)->setClipRect(*static_cast<QRect*>(rectangle), static_cast<Qt::ClipOperation>(operation));
}

void QPainter_SetClipRect(void* ptr, void* rectangle, int operation)
{
	static_cast<QPainter*>(ptr)->setClipRect(*static_cast<QRectF*>(rectangle), static_cast<Qt::ClipOperation>(operation));
}

void QPainter_SetClipRegion(void* ptr, void* region, int operation)
{
	static_cast<QPainter*>(ptr)->setClipRegion(*static_cast<QRegion*>(region), static_cast<Qt::ClipOperation>(operation));
}

void QPainter_SetViewport(void* ptr, void* rectangle)
{
	static_cast<QPainter*>(ptr)->setViewport(*static_cast<QRect*>(rectangle));
}

void QPainter_SetWindow(void* ptr, void* rectangle)
{
	static_cast<QPainter*>(ptr)->setWindow(*static_cast<QRect*>(rectangle));
}

void* QPainter_NewQPainter()
{
	return new QPainter();
}

void* QPainter_Background(void* ptr)
{
	return new QBrush(static_cast<QPainter*>(ptr)->background());
}

int QPainter_BackgroundMode(void* ptr)
{
	return static_cast<QPainter*>(ptr)->backgroundMode();
}

void QPainter_BeginNativePainting(void* ptr)
{
	static_cast<QPainter*>(ptr)->beginNativePainting();
}

void* QPainter_BoundingRect3(void* ptr, int x, int y, int w, int h, int flags, char* text)
{
	return new QRect(static_cast<QRect>(static_cast<QPainter*>(ptr)->boundingRect(x, y, w, h, flags, QString(text))).x(), static_cast<QRect>(static_cast<QPainter*>(ptr)->boundingRect(x, y, w, h, flags, QString(text))).y(), static_cast<QRect>(static_cast<QPainter*>(ptr)->boundingRect(x, y, w, h, flags, QString(text))).width(), static_cast<QRect>(static_cast<QPainter*>(ptr)->boundingRect(x, y, w, h, flags, QString(text))).height());
}

void* QPainter_Brush(void* ptr)
{
	return new QBrush(static_cast<QPainter*>(ptr)->brush());
}

void* QPainter_BrushOrigin(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QPainter*>(ptr)->brushOrigin()).x(), static_cast<QPoint>(static_cast<QPainter*>(ptr)->brushOrigin()).y());
}

void* QPainter_ClipBoundingRect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QPainter*>(ptr)->clipBoundingRect()).x(), static_cast<QRectF>(static_cast<QPainter*>(ptr)->clipBoundingRect()).y(), static_cast<QRectF>(static_cast<QPainter*>(ptr)->clipBoundingRect()).width(), static_cast<QRectF>(static_cast<QPainter*>(ptr)->clipBoundingRect()).height());
}

void* QPainter_ClipPath(void* ptr)
{
	return new QPainterPath(static_cast<QPainter*>(ptr)->clipPath());
}

void* QPainter_ClipRegion(void* ptr)
{
	return new QRegion(static_cast<QPainter*>(ptr)->clipRegion());
}

int QPainter_CompositionMode(void* ptr)
{
	return static_cast<QPainter*>(ptr)->compositionMode();
}

void* QPainter_Device(void* ptr)
{
	return static_cast<QPainter*>(ptr)->device();
}

void QPainter_DrawArc2(void* ptr, void* rectangle, int startAngle, int spanAngle)
{
	static_cast<QPainter*>(ptr)->drawArc(*static_cast<QRect*>(rectangle), startAngle, spanAngle);
}

void QPainter_DrawArc3(void* ptr, int x, int y, int width, int height, int startAngle, int spanAngle)
{
	static_cast<QPainter*>(ptr)->drawArc(x, y, width, height, startAngle, spanAngle);
}

void QPainter_DrawChord2(void* ptr, void* rectangle, int startAngle, int spanAngle)
{
	static_cast<QPainter*>(ptr)->drawChord(*static_cast<QRect*>(rectangle), startAngle, spanAngle);
}

void QPainter_DrawChord3(void* ptr, int x, int y, int width, int height, int startAngle, int spanAngle)
{
	static_cast<QPainter*>(ptr)->drawChord(x, y, width, height, startAngle, spanAngle);
}

void QPainter_DrawConvexPolygon4(void* ptr, void* polygon)
{
	static_cast<QPainter*>(ptr)->drawConvexPolygon(*static_cast<QPolygon*>(polygon));
}

void QPainter_DrawConvexPolygon3(void* ptr, void* polygon)
{
	static_cast<QPainter*>(ptr)->drawConvexPolygon(*static_cast<QPolygonF*>(polygon));
}

void QPainter_DrawEllipse5(void* ptr, void* center, int rx, int ry)
{
	static_cast<QPainter*>(ptr)->drawEllipse(*static_cast<QPoint*>(center), rx, ry);
}

void QPainter_DrawEllipse4(void* ptr, void* center, double rx, double ry)
{
	static_cast<QPainter*>(ptr)->drawEllipse(*static_cast<QPointF*>(center), static_cast<double>(rx), static_cast<double>(ry));
}

void QPainter_DrawEllipse3(void* ptr, int x, int y, int width, int height)
{
	static_cast<QPainter*>(ptr)->drawEllipse(x, y, width, height);
}

void QPainter_DrawImage4(void* ptr, void* point, void* image)
{
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QPoint*>(point), *static_cast<QImage*>(image));
}

void QPainter_DrawImage6(void* ptr, void* point, void* image, void* source, int flags)
{
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QPoint*>(point), *static_cast<QImage*>(image), *static_cast<QRect*>(source), static_cast<Qt::ImageConversionFlag>(flags));
}

void QPainter_DrawImage5(void* ptr, void* point, void* image, void* source, int flags)
{
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QPointF*>(point), *static_cast<QImage*>(image), *static_cast<QRectF*>(source), static_cast<Qt::ImageConversionFlag>(flags));
}

void QPainter_DrawImage8(void* ptr, void* rectangle, void* image)
{
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QRect*>(rectangle), *static_cast<QImage*>(image));
}

void QPainter_DrawImage2(void* ptr, void* target, void* image, void* source, int flags)
{
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QRect*>(target), *static_cast<QImage*>(image), *static_cast<QRect*>(source), static_cast<Qt::ImageConversionFlag>(flags));
}

void QPainter_DrawImage7(void* ptr, void* rectangle, void* image)
{
	static_cast<QPainter*>(ptr)->drawImage(*static_cast<QRectF*>(rectangle), *static_cast<QImage*>(image));
}

void QPainter_DrawImage9(void* ptr, int x, int y, void* image, int sx, int sy, int sw, int sh, int flags)
{
	static_cast<QPainter*>(ptr)->drawImage(x, y, *static_cast<QImage*>(image), sx, sy, sw, sh, static_cast<Qt::ImageConversionFlag>(flags));
}

void QPainter_DrawLine2(void* ptr, void* line)
{
	static_cast<QPainter*>(ptr)->drawLine(*static_cast<QLine*>(line));
}

void QPainter_DrawLine(void* ptr, void* line)
{
	static_cast<QPainter*>(ptr)->drawLine(*static_cast<QLineF*>(line));
}

void QPainter_DrawLine3(void* ptr, void* p1, void* p2)
{
	static_cast<QPainter*>(ptr)->drawLine(*static_cast<QPoint*>(p1), *static_cast<QPoint*>(p2));
}

void QPainter_DrawLine4(void* ptr, void* p1, void* p2)
{
	static_cast<QPainter*>(ptr)->drawLine(*static_cast<QPointF*>(p1), *static_cast<QPointF*>(p2));
}

void QPainter_DrawLine5(void* ptr, int x1, int y1, int x2, int y2)
{
	static_cast<QPainter*>(ptr)->drawLine(x1, y1, x2, y2);
}

void QPainter_DrawLines(void* ptr, void* lines, int lineCount)
{
	static_cast<QPainter*>(ptr)->drawLines(static_cast<QLineF*>(lines), lineCount);
}

void QPainter_DrawLines4(void* ptr, void* pointPairs, int lineCount)
{
	static_cast<QPainter*>(ptr)->drawLines(static_cast<QPoint*>(pointPairs), lineCount);
}

void QPainter_DrawLines3(void* ptr, void* pointPairs, int lineCount)
{
	static_cast<QPainter*>(ptr)->drawLines(static_cast<QPointF*>(pointPairs), lineCount);
}

void QPainter_DrawPath(void* ptr, void* path)
{
	static_cast<QPainter*>(ptr)->drawPath(*static_cast<QPainterPath*>(path));
}

void QPainter_DrawPicture2(void* ptr, void* point, void* picture)
{
	static_cast<QPainter*>(ptr)->drawPicture(*static_cast<QPoint*>(point), *static_cast<QPicture*>(picture));
}

void QPainter_DrawPicture3(void* ptr, int x, int y, void* picture)
{
	static_cast<QPainter*>(ptr)->drawPicture(x, y, *static_cast<QPicture*>(picture));
}

void QPainter_DrawPie2(void* ptr, void* rectangle, int startAngle, int spanAngle)
{
	static_cast<QPainter*>(ptr)->drawPie(*static_cast<QRect*>(rectangle), startAngle, spanAngle);
}

void QPainter_DrawPie3(void* ptr, int x, int y, int width, int height, int startAngle, int spanAngle)
{
	static_cast<QPainter*>(ptr)->drawPie(x, y, width, height, startAngle, spanAngle);
}

void QPainter_DrawPixmap6(void* ptr, void* point, void* pixmap)
{
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QPoint*>(point), *static_cast<QPixmap*>(pixmap));
}

void QPainter_DrawPixmap4(void* ptr, void* point, void* pixmap, void* source)
{
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QPoint*>(point), *static_cast<QPixmap*>(pixmap), *static_cast<QRect*>(source));
}

void QPainter_DrawPixmap3(void* ptr, void* point, void* pixmap, void* source)
{
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QPointF*>(point), *static_cast<QPixmap*>(pixmap), *static_cast<QRectF*>(source));
}

void QPainter_DrawPixmap8(void* ptr, void* rectangle, void* pixmap)
{
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QRect*>(rectangle), *static_cast<QPixmap*>(pixmap));
}

void QPainter_DrawPixmap2(void* ptr, void* target, void* pixmap, void* source)
{
	static_cast<QPainter*>(ptr)->drawPixmap(*static_cast<QRect*>(target), *static_cast<QPixmap*>(pixmap), *static_cast<QRect*>(source));
}

void QPainter_DrawPixmap7(void* ptr, int x, int y, void* pixmap)
{
	static_cast<QPainter*>(ptr)->drawPixmap(x, y, *static_cast<QPixmap*>(pixmap));
}

void QPainter_DrawPixmap11(void* ptr, int x, int y, void* pixmap, int sx, int sy, int sw, int sh)
{
	static_cast<QPainter*>(ptr)->drawPixmap(x, y, *static_cast<QPixmap*>(pixmap), sx, sy, sw, sh);
}

void QPainter_DrawPixmap10(void* ptr, int x, int y, int w, int h, void* pixmap, int sx, int sy, int sw, int sh)
{
	static_cast<QPainter*>(ptr)->drawPixmap(x, y, w, h, *static_cast<QPixmap*>(pixmap), sx, sy, sw, sh);
}

void QPainter_DrawPixmap9(void* ptr, int x, int y, int width, int height, void* pixmap)
{
	static_cast<QPainter*>(ptr)->drawPixmap(x, y, width, height, *static_cast<QPixmap*>(pixmap));
}

void QPainter_DrawPoint2(void* ptr, void* position)
{
	static_cast<QPainter*>(ptr)->drawPoint(*static_cast<QPoint*>(position));
}

void QPainter_DrawPoint(void* ptr, void* position)
{
	static_cast<QPainter*>(ptr)->drawPoint(*static_cast<QPointF*>(position));
}

void QPainter_DrawPoint3(void* ptr, int x, int y)
{
	static_cast<QPainter*>(ptr)->drawPoint(x, y);
}

void QPainter_DrawPoints2(void* ptr, void* points, int pointCount)
{
	static_cast<QPainter*>(ptr)->drawPoints(static_cast<QPoint*>(points), pointCount);
}

void QPainter_DrawPoints(void* ptr, void* points, int pointCount)
{
	static_cast<QPainter*>(ptr)->drawPoints(static_cast<QPointF*>(points), pointCount);
}

void QPainter_DrawPoints4(void* ptr, void* points)
{
	static_cast<QPainter*>(ptr)->drawPoints(*static_cast<QPolygon*>(points));
}

void QPainter_DrawPoints3(void* ptr, void* points)
{
	static_cast<QPainter*>(ptr)->drawPoints(*static_cast<QPolygonF*>(points));
}

void QPainter_DrawPolygon2(void* ptr, void* points, int pointCount, int fillRule)
{
	static_cast<QPainter*>(ptr)->drawPolygon(static_cast<QPoint*>(points), pointCount, static_cast<Qt::FillRule>(fillRule));
}

void QPainter_DrawPolygon(void* ptr, void* points, int pointCount, int fillRule)
{
	static_cast<QPainter*>(ptr)->drawPolygon(static_cast<QPointF*>(points), pointCount, static_cast<Qt::FillRule>(fillRule));
}

void QPainter_DrawPolygon4(void* ptr, void* points, int fillRule)
{
	static_cast<QPainter*>(ptr)->drawPolygon(*static_cast<QPolygon*>(points), static_cast<Qt::FillRule>(fillRule));
}

void QPainter_DrawPolygon3(void* ptr, void* points, int fillRule)
{
	static_cast<QPainter*>(ptr)->drawPolygon(*static_cast<QPolygonF*>(points), static_cast<Qt::FillRule>(fillRule));
}

void QPainter_DrawPolyline2(void* ptr, void* points, int pointCount)
{
	static_cast<QPainter*>(ptr)->drawPolyline(static_cast<QPoint*>(points), pointCount);
}

void QPainter_DrawPolyline(void* ptr, void* points, int pointCount)
{
	static_cast<QPainter*>(ptr)->drawPolyline(static_cast<QPointF*>(points), pointCount);
}

void QPainter_DrawPolyline4(void* ptr, void* points)
{
	static_cast<QPainter*>(ptr)->drawPolyline(*static_cast<QPolygon*>(points));
}

void QPainter_DrawPolyline3(void* ptr, void* points)
{
	static_cast<QPainter*>(ptr)->drawPolyline(*static_cast<QPolygonF*>(points));
}

void QPainter_DrawRect2(void* ptr, void* rectangle)
{
	static_cast<QPainter*>(ptr)->drawRect(*static_cast<QRect*>(rectangle));
}

void QPainter_DrawRect(void* ptr, void* rectangle)
{
	static_cast<QPainter*>(ptr)->drawRect(*static_cast<QRectF*>(rectangle));
}

void QPainter_DrawRect3(void* ptr, int x, int y, int width, int height)
{
	static_cast<QPainter*>(ptr)->drawRect(x, y, width, height);
}

void QPainter_DrawRoundedRect2(void* ptr, void* rect, double xRadius, double yRadius, int mode)
{
	static_cast<QPainter*>(ptr)->drawRoundedRect(*static_cast<QRect*>(rect), static_cast<double>(xRadius), static_cast<double>(yRadius), static_cast<Qt::SizeMode>(mode));
}

void QPainter_DrawRoundedRect(void* ptr, void* rect, double xRadius, double yRadius, int mode)
{
	static_cast<QPainter*>(ptr)->drawRoundedRect(*static_cast<QRectF*>(rect), static_cast<double>(xRadius), static_cast<double>(yRadius), static_cast<Qt::SizeMode>(mode));
}

void QPainter_DrawRoundedRect3(void* ptr, int x, int y, int w, int h, double xRadius, double yRadius, int mode)
{
	static_cast<QPainter*>(ptr)->drawRoundedRect(x, y, w, h, static_cast<double>(xRadius), static_cast<double>(yRadius), static_cast<Qt::SizeMode>(mode));
}

void QPainter_DrawStaticText2(void* ptr, void* topLeftPosition, void* staticText)
{
	static_cast<QPainter*>(ptr)->drawStaticText(*static_cast<QPoint*>(topLeftPosition), *static_cast<QStaticText*>(staticText));
}

void QPainter_DrawStaticText(void* ptr, void* topLeftPosition, void* staticText)
{
	static_cast<QPainter*>(ptr)->drawStaticText(*static_cast<QPointF*>(topLeftPosition), *static_cast<QStaticText*>(staticText));
}

void QPainter_DrawStaticText3(void* ptr, int left, int top, void* staticText)
{
	static_cast<QPainter*>(ptr)->drawStaticText(left, top, *static_cast<QStaticText*>(staticText));
}

void QPainter_DrawText3(void* ptr, void* position, char* text)
{
	static_cast<QPainter*>(ptr)->drawText(*static_cast<QPoint*>(position), QString(text));
}

void QPainter_DrawText6(void* ptr, int x, int y, char* text)
{
	static_cast<QPainter*>(ptr)->drawText(x, y, QString(text));
}

void QPainter_DrawText7(void* ptr, int x, int y, int width, int height, int flags, char* text, void* boundingRect)
{
	static_cast<QPainter*>(ptr)->drawText(x, y, width, height, flags, QString(text), static_cast<QRect*>(boundingRect));
}

void QPainter_DrawTiledPixmap2(void* ptr, void* rectangle, void* pixmap, void* position)
{
	static_cast<QPainter*>(ptr)->drawTiledPixmap(*static_cast<QRect*>(rectangle), *static_cast<QPixmap*>(pixmap), *static_cast<QPoint*>(position));
}

void QPainter_DrawTiledPixmap3(void* ptr, int x, int y, int width, int height, void* pixmap, int sx, int sy)
{
	static_cast<QPainter*>(ptr)->drawTiledPixmap(x, y, width, height, *static_cast<QPixmap*>(pixmap), sx, sy);
}

int QPainter_End(void* ptr)
{
	return static_cast<QPainter*>(ptr)->end();
}

void QPainter_EndNativePainting(void* ptr)
{
	static_cast<QPainter*>(ptr)->endNativePainting();
}

void QPainter_EraseRect2(void* ptr, void* rectangle)
{
	static_cast<QPainter*>(ptr)->eraseRect(*static_cast<QRect*>(rectangle));
}

void QPainter_EraseRect3(void* ptr, int x, int y, int width, int height)
{
	static_cast<QPainter*>(ptr)->eraseRect(x, y, width, height);
}

void QPainter_FillPath(void* ptr, void* path, void* brush)
{
	static_cast<QPainter*>(ptr)->fillPath(*static_cast<QPainterPath*>(path), *static_cast<QBrush*>(brush));
}

void QPainter_FillRect3(void* ptr, void* rectangle, int style)
{
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRect*>(rectangle), static_cast<Qt::BrushStyle>(style));
}

void QPainter_FillRect11(void* ptr, void* rectangle, int color)
{
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRect*>(rectangle), static_cast<Qt::GlobalColor>(color));
}

void QPainter_FillRect4(void* ptr, void* rectangle, int style)
{
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRectF*>(rectangle), static_cast<Qt::BrushStyle>(style));
}

void QPainter_FillRect12(void* ptr, void* rectangle, int color)
{
	static_cast<QPainter*>(ptr)->fillRect(*static_cast<QRectF*>(rectangle), static_cast<Qt::GlobalColor>(color));
}

void QPainter_FillRect2(void* ptr, int x, int y, int width, int height, int style)
{
	static_cast<QPainter*>(ptr)->fillRect(x, y, width, height, static_cast<Qt::BrushStyle>(style));
}

void QPainter_FillRect10(void* ptr, int x, int y, int width, int height, int color)
{
	static_cast<QPainter*>(ptr)->fillRect(x, y, width, height, static_cast<Qt::GlobalColor>(color));
}

void QPainter_FillRect8(void* ptr, int x, int y, int width, int height, void* brush)
{
	static_cast<QPainter*>(ptr)->fillRect(x, y, width, height, *static_cast<QBrush*>(brush));
}

void QPainter_FillRect9(void* ptr, int x, int y, int width, int height, void* color)
{
	static_cast<QPainter*>(ptr)->fillRect(x, y, width, height, *static_cast<QColor*>(color));
}

void* QPainter_Font(void* ptr)
{
	return new QFont(static_cast<QPainter*>(ptr)->font());
}

void* QPainter_FontInfo(void* ptr)
{
	return new QFontInfo(static_cast<QPainter*>(ptr)->fontInfo());
}

void* QPainter_FontMetrics(void* ptr)
{
	return new QFontMetrics(static_cast<QPainter*>(ptr)->fontMetrics());
}

int QPainter_HasClipping(void* ptr)
{
	return static_cast<QPainter*>(ptr)->hasClipping();
}

int QPainter_IsActive(void* ptr)
{
	return static_cast<QPainter*>(ptr)->isActive();
}

int QPainter_LayoutDirection(void* ptr)
{
	return static_cast<QPainter*>(ptr)->layoutDirection();
}

double QPainter_Opacity(void* ptr)
{
	return static_cast<double>(static_cast<QPainter*>(ptr)->opacity());
}

void* QPainter_PaintEngine(void* ptr)
{
	return static_cast<QPainter*>(ptr)->paintEngine();
}

void* QPainter_Pen(void* ptr)
{
	return new QPen(static_cast<QPainter*>(ptr)->pen());
}

int QPainter_RenderHints(void* ptr)
{
	return static_cast<QPainter*>(ptr)->renderHints();
}

void QPainter_ResetTransform(void* ptr)
{
	static_cast<QPainter*>(ptr)->resetTransform();
}

void QPainter_Restore(void* ptr)
{
	static_cast<QPainter*>(ptr)->restore();
}

void QPainter_Save(void* ptr)
{
	static_cast<QPainter*>(ptr)->save();
}

void QPainter_Scale(void* ptr, double sx, double sy)
{
	static_cast<QPainter*>(ptr)->scale(static_cast<double>(sx), static_cast<double>(sy));
}

void QPainter_SetBackgroundMode(void* ptr, int mode)
{
	static_cast<QPainter*>(ptr)->setBackgroundMode(static_cast<Qt::BGMode>(mode));
}

void QPainter_SetBrush2(void* ptr, int style)
{
	static_cast<QPainter*>(ptr)->setBrush(static_cast<Qt::BrushStyle>(style));
}

void QPainter_SetBrush(void* ptr, void* brush)
{
	static_cast<QPainter*>(ptr)->setBrush(*static_cast<QBrush*>(brush));
}

void QPainter_SetBrushOrigin2(void* ptr, void* position)
{
	static_cast<QPainter*>(ptr)->setBrushOrigin(*static_cast<QPoint*>(position));
}

void QPainter_SetBrushOrigin3(void* ptr, int x, int y)
{
	static_cast<QPainter*>(ptr)->setBrushOrigin(x, y);
}

void QPainter_SetClipRect2(void* ptr, int x, int y, int width, int height, int operation)
{
	static_cast<QPainter*>(ptr)->setClipRect(x, y, width, height, static_cast<Qt::ClipOperation>(operation));
}

void QPainter_SetClipping(void* ptr, int enable)
{
	static_cast<QPainter*>(ptr)->setClipping(enable != 0);
}

void QPainter_SetCompositionMode(void* ptr, int mode)
{
	static_cast<QPainter*>(ptr)->setCompositionMode(static_cast<QPainter::CompositionMode>(mode));
}

void QPainter_SetFont(void* ptr, void* font)
{
	static_cast<QPainter*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QPainter_SetLayoutDirection(void* ptr, int direction)
{
	static_cast<QPainter*>(ptr)->setLayoutDirection(static_cast<Qt::LayoutDirection>(direction));
}

void QPainter_SetOpacity(void* ptr, double opacity)
{
	static_cast<QPainter*>(ptr)->setOpacity(static_cast<double>(opacity));
}

void QPainter_SetPen3(void* ptr, int style)
{
	static_cast<QPainter*>(ptr)->setPen(static_cast<Qt::PenStyle>(style));
}

void QPainter_SetPen2(void* ptr, void* color)
{
	static_cast<QPainter*>(ptr)->setPen(*static_cast<QColor*>(color));
}

void QPainter_SetPen(void* ptr, void* pen)
{
	static_cast<QPainter*>(ptr)->setPen(*static_cast<QPen*>(pen));
}

void QPainter_SetRenderHint(void* ptr, int hint, int on)
{
	static_cast<QPainter*>(ptr)->setRenderHint(static_cast<QPainter::RenderHint>(hint), on != 0);
}

void QPainter_SetRenderHints(void* ptr, int hints, int on)
{
	static_cast<QPainter*>(ptr)->setRenderHints(static_cast<QPainter::RenderHint>(hints), on != 0);
}

void QPainter_SetTransform(void* ptr, void* transform, int combine)
{
	static_cast<QPainter*>(ptr)->setTransform(*static_cast<QTransform*>(transform), combine != 0);
}

void QPainter_SetViewTransformEnabled(void* ptr, int enable)
{
	static_cast<QPainter*>(ptr)->setViewTransformEnabled(enable != 0);
}

void QPainter_SetViewport2(void* ptr, int x, int y, int width, int height)
{
	static_cast<QPainter*>(ptr)->setViewport(x, y, width, height);
}

void QPainter_SetWindow2(void* ptr, int x, int y, int width, int height)
{
	static_cast<QPainter*>(ptr)->setWindow(x, y, width, height);
}

void QPainter_SetWorldMatrixEnabled(void* ptr, int enable)
{
	static_cast<QPainter*>(ptr)->setWorldMatrixEnabled(enable != 0);
}

void QPainter_SetWorldTransform(void* ptr, void* matrix, int combine)
{
	static_cast<QPainter*>(ptr)->setWorldTransform(*static_cast<QTransform*>(matrix), combine != 0);
}

void QPainter_Shear(void* ptr, double sh, double sv)
{
	static_cast<QPainter*>(ptr)->shear(static_cast<double>(sh), static_cast<double>(sv));
}

void QPainter_StrokePath(void* ptr, void* path, void* pen)
{
	static_cast<QPainter*>(ptr)->strokePath(*static_cast<QPainterPath*>(path), *static_cast<QPen*>(pen));
}

int QPainter_TestRenderHint(void* ptr, int hint)
{
	return static_cast<QPainter*>(ptr)->testRenderHint(static_cast<QPainter::RenderHint>(hint));
}

void QPainter_Translate2(void* ptr, void* offset)
{
	static_cast<QPainter*>(ptr)->translate(*static_cast<QPoint*>(offset));
}

void QPainter_Translate(void* ptr, void* offset)
{
	static_cast<QPainter*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

void QPainter_Translate3(void* ptr, double dx, double dy)
{
	static_cast<QPainter*>(ptr)->translate(static_cast<double>(dx), static_cast<double>(dy));
}

int QPainter_ViewTransformEnabled(void* ptr)
{
	return static_cast<QPainter*>(ptr)->viewTransformEnabled();
}

void* QPainter_Viewport(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QPainter*>(ptr)->viewport()).x(), static_cast<QRect>(static_cast<QPainter*>(ptr)->viewport()).y(), static_cast<QRect>(static_cast<QPainter*>(ptr)->viewport()).width(), static_cast<QRect>(static_cast<QPainter*>(ptr)->viewport()).height());
}

void* QPainter_Window(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QPainter*>(ptr)->window()).x(), static_cast<QRect>(static_cast<QPainter*>(ptr)->window()).y(), static_cast<QRect>(static_cast<QPainter*>(ptr)->window()).width(), static_cast<QRect>(static_cast<QPainter*>(ptr)->window()).height());
}

int QPainter_WorldMatrixEnabled(void* ptr)
{
	return static_cast<QPainter*>(ptr)->worldMatrixEnabled();
}

void QPainter_DestroyQPainter(void* ptr)
{
	static_cast<QPainter*>(ptr)->~QPainter();
}

void* QPainterPath_NewQPainterPath3(void* path)
{
	return new QPainterPath(*static_cast<QPainterPath*>(path));
}

void QPainterPath_AddEllipse(void* ptr, void* boundingRectangle)
{
	static_cast<QPainterPath*>(ptr)->addEllipse(*static_cast<QRectF*>(boundingRectangle));
}

void QPainterPath_AddPath(void* ptr, void* path)
{
	static_cast<QPainterPath*>(ptr)->addPath(*static_cast<QPainterPath*>(path));
}

void QPainterPath_AddRect(void* ptr, void* rectangle)
{
	static_cast<QPainterPath*>(ptr)->addRect(*static_cast<QRectF*>(rectangle));
}

void QPainterPath_AddText(void* ptr, void* point, void* font, char* text)
{
	static_cast<QPainterPath*>(ptr)->addText(*static_cast<QPointF*>(point), *static_cast<QFont*>(font), QString(text));
}

void QPainterPath_ArcMoveTo(void* ptr, void* rectangle, double angle)
{
	static_cast<QPainterPath*>(ptr)->arcMoveTo(*static_cast<QRectF*>(rectangle), static_cast<double>(angle));
}

void QPainterPath_ArcTo(void* ptr, void* rectangle, double startAngle, double sweepLength)
{
	static_cast<QPainterPath*>(ptr)->arcTo(*static_cast<QRectF*>(rectangle), static_cast<double>(startAngle), static_cast<double>(sweepLength));
}

void QPainterPath_ConnectPath(void* ptr, void* path)
{
	static_cast<QPainterPath*>(ptr)->connectPath(*static_cast<QPainterPath*>(path));
}

int QPainterPath_Contains(void* ptr, void* point)
{
	return static_cast<QPainterPath*>(ptr)->contains(*static_cast<QPointF*>(point));
}

int QPainterPath_Contains2(void* ptr, void* rectangle)
{
	return static_cast<QPainterPath*>(ptr)->contains(*static_cast<QRectF*>(rectangle));
}

void QPainterPath_CubicTo(void* ptr, void* c1, void* c2, void* endPoint)
{
	static_cast<QPainterPath*>(ptr)->cubicTo(*static_cast<QPointF*>(c1), *static_cast<QPointF*>(c2), *static_cast<QPointF*>(endPoint));
}

void* QPainterPath_CurrentPosition(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QPainterPath*>(ptr)->currentPosition()).x(), static_cast<QPointF>(static_cast<QPainterPath*>(ptr)->currentPosition()).y());
}

int QPainterPath_ElementCount(void* ptr)
{
	return static_cast<QPainterPath*>(ptr)->elementCount();
}

int QPainterPath_Intersects(void* ptr, void* rectangle)
{
	return static_cast<QPainterPath*>(ptr)->intersects(*static_cast<QRectF*>(rectangle));
}

int QPainterPath_IsEmpty(void* ptr)
{
	return static_cast<QPainterPath*>(ptr)->isEmpty();
}

void QPainterPath_LineTo(void* ptr, void* endPoint)
{
	static_cast<QPainterPath*>(ptr)->lineTo(*static_cast<QPointF*>(endPoint));
}

void QPainterPath_MoveTo(void* ptr, void* point)
{
	static_cast<QPainterPath*>(ptr)->moveTo(*static_cast<QPointF*>(point));
}

void QPainterPath_QuadTo(void* ptr, void* c, void* endPoint)
{
	static_cast<QPainterPath*>(ptr)->quadTo(*static_cast<QPointF*>(c), *static_cast<QPointF*>(endPoint));
}

void QPainterPath_SetElementPositionAt(void* ptr, int index, double x, double y)
{
	static_cast<QPainterPath*>(ptr)->setElementPositionAt(index, static_cast<double>(x), static_cast<double>(y));
}

void QPainterPath_SetFillRule(void* ptr, int fillRule)
{
	static_cast<QPainterPath*>(ptr)->setFillRule(static_cast<Qt::FillRule>(fillRule));
}

void* QPainterPath_NewQPainterPath()
{
	return new QPainterPath();
}

void* QPainterPath_NewQPainterPath2(void* startPoint)
{
	return new QPainterPath(*static_cast<QPointF*>(startPoint));
}

void QPainterPath_AddEllipse3(void* ptr, void* center, double rx, double ry)
{
	static_cast<QPainterPath*>(ptr)->addEllipse(*static_cast<QPointF*>(center), static_cast<double>(rx), static_cast<double>(ry));
}

void QPainterPath_AddEllipse2(void* ptr, double x, double y, double width, double height)
{
	static_cast<QPainterPath*>(ptr)->addEllipse(static_cast<double>(x), static_cast<double>(y), static_cast<double>(width), static_cast<double>(height));
}

void QPainterPath_AddPolygon(void* ptr, void* polygon)
{
	static_cast<QPainterPath*>(ptr)->addPolygon(*static_cast<QPolygonF*>(polygon));
}

void QPainterPath_AddRect2(void* ptr, double x, double y, double width, double height)
{
	static_cast<QPainterPath*>(ptr)->addRect(static_cast<double>(x), static_cast<double>(y), static_cast<double>(width), static_cast<double>(height));
}

void QPainterPath_AddRegion(void* ptr, void* region)
{
	static_cast<QPainterPath*>(ptr)->addRegion(*static_cast<QRegion*>(region));
}

void QPainterPath_AddRoundedRect(void* ptr, void* rect, double xRadius, double yRadius, int mode)
{
	static_cast<QPainterPath*>(ptr)->addRoundedRect(*static_cast<QRectF*>(rect), static_cast<double>(xRadius), static_cast<double>(yRadius), static_cast<Qt::SizeMode>(mode));
}

void QPainterPath_AddRoundedRect2(void* ptr, double x, double y, double w, double h, double xRadius, double yRadius, int mode)
{
	static_cast<QPainterPath*>(ptr)->addRoundedRect(static_cast<double>(x), static_cast<double>(y), static_cast<double>(w), static_cast<double>(h), static_cast<double>(xRadius), static_cast<double>(yRadius), static_cast<Qt::SizeMode>(mode));
}

void QPainterPath_AddText2(void* ptr, double x, double y, void* font, char* text)
{
	static_cast<QPainterPath*>(ptr)->addText(static_cast<double>(x), static_cast<double>(y), *static_cast<QFont*>(font), QString(text));
}

double QPainterPath_AngleAtPercent(void* ptr, double t)
{
	return static_cast<double>(static_cast<QPainterPath*>(ptr)->angleAtPercent(static_cast<double>(t)));
}

void QPainterPath_ArcMoveTo2(void* ptr, double x, double y, double width, double height, double angle)
{
	static_cast<QPainterPath*>(ptr)->arcMoveTo(static_cast<double>(x), static_cast<double>(y), static_cast<double>(width), static_cast<double>(height), static_cast<double>(angle));
}

void QPainterPath_ArcTo2(void* ptr, double x, double y, double width, double height, double startAngle, double sweepLength)
{
	static_cast<QPainterPath*>(ptr)->arcTo(static_cast<double>(x), static_cast<double>(y), static_cast<double>(width), static_cast<double>(height), static_cast<double>(startAngle), static_cast<double>(sweepLength));
}

void* QPainterPath_BoundingRect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QPainterPath*>(ptr)->boundingRect()).x(), static_cast<QRectF>(static_cast<QPainterPath*>(ptr)->boundingRect()).y(), static_cast<QRectF>(static_cast<QPainterPath*>(ptr)->boundingRect()).width(), static_cast<QRectF>(static_cast<QPainterPath*>(ptr)->boundingRect()).height());
}

void QPainterPath_CloseSubpath(void* ptr)
{
	static_cast<QPainterPath*>(ptr)->closeSubpath();
}

int QPainterPath_Contains3(void* ptr, void* p)
{
	return static_cast<QPainterPath*>(ptr)->contains(*static_cast<QPainterPath*>(p));
}

void* QPainterPath_ControlPointRect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QPainterPath*>(ptr)->controlPointRect()).x(), static_cast<QRectF>(static_cast<QPainterPath*>(ptr)->controlPointRect()).y(), static_cast<QRectF>(static_cast<QPainterPath*>(ptr)->controlPointRect()).width(), static_cast<QRectF>(static_cast<QPainterPath*>(ptr)->controlPointRect()).height());
}

void QPainterPath_CubicTo2(void* ptr, double c1X, double c1Y, double c2X, double c2Y, double endPointX, double endPointY)
{
	static_cast<QPainterPath*>(ptr)->cubicTo(static_cast<double>(c1X), static_cast<double>(c1Y), static_cast<double>(c2X), static_cast<double>(c2Y), static_cast<double>(endPointX), static_cast<double>(endPointY));
}

int QPainterPath_FillRule(void* ptr)
{
	return static_cast<QPainterPath*>(ptr)->fillRule();
}

void* QPainterPath_Intersected(void* ptr, void* p)
{
	return new QPainterPath(static_cast<QPainterPath*>(ptr)->intersected(*static_cast<QPainterPath*>(p)));
}

int QPainterPath_Intersects2(void* ptr, void* p)
{
	return static_cast<QPainterPath*>(ptr)->intersects(*static_cast<QPainterPath*>(p));
}

double QPainterPath_Length(void* ptr)
{
	return static_cast<double>(static_cast<QPainterPath*>(ptr)->length());
}

void QPainterPath_LineTo2(void* ptr, double x, double y)
{
	static_cast<QPainterPath*>(ptr)->lineTo(static_cast<double>(x), static_cast<double>(y));
}

void QPainterPath_MoveTo2(void* ptr, double x, double y)
{
	static_cast<QPainterPath*>(ptr)->moveTo(static_cast<double>(x), static_cast<double>(y));
}

double QPainterPath_PercentAtLength(void* ptr, double len)
{
	return static_cast<double>(static_cast<QPainterPath*>(ptr)->percentAtLength(static_cast<double>(len)));
}

void* QPainterPath_PointAtPercent(void* ptr, double t)
{
	return new QPointF(static_cast<QPointF>(static_cast<QPainterPath*>(ptr)->pointAtPercent(static_cast<double>(t))).x(), static_cast<QPointF>(static_cast<QPainterPath*>(ptr)->pointAtPercent(static_cast<double>(t))).y());
}

void QPainterPath_QuadTo2(void* ptr, double cx, double cy, double endPointX, double endPointY)
{
	static_cast<QPainterPath*>(ptr)->quadTo(static_cast<double>(cx), static_cast<double>(cy), static_cast<double>(endPointX), static_cast<double>(endPointY));
}

void* QPainterPath_Simplified(void* ptr)
{
	return new QPainterPath(static_cast<QPainterPath*>(ptr)->simplified());
}

double QPainterPath_SlopeAtPercent(void* ptr, double t)
{
	return static_cast<double>(static_cast<QPainterPath*>(ptr)->slopeAtPercent(static_cast<double>(t)));
}

void* QPainterPath_Subtracted(void* ptr, void* p)
{
	return new QPainterPath(static_cast<QPainterPath*>(ptr)->subtracted(*static_cast<QPainterPath*>(p)));
}

void QPainterPath_Swap(void* ptr, void* other)
{
	static_cast<QPainterPath*>(ptr)->swap(*static_cast<QPainterPath*>(other));
}

void* QPainterPath_ToFillPolygon(void* ptr, void* matrix)
{
	return new QPolygonF(static_cast<QPainterPath*>(ptr)->toFillPolygon(*static_cast<QTransform*>(matrix)));
}

void* QPainterPath_ToReversed(void* ptr)
{
	return new QPainterPath(static_cast<QPainterPath*>(ptr)->toReversed());
}

void QPainterPath_Translate2(void* ptr, void* offset)
{
	static_cast<QPainterPath*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

void QPainterPath_Translate(void* ptr, double dx, double dy)
{
	static_cast<QPainterPath*>(ptr)->translate(static_cast<double>(dx), static_cast<double>(dy));
}

void* QPainterPath_Translated2(void* ptr, void* offset)
{
	return new QPainterPath(static_cast<QPainterPath*>(ptr)->translated(*static_cast<QPointF*>(offset)));
}

void* QPainterPath_Translated(void* ptr, double dx, double dy)
{
	return new QPainterPath(static_cast<QPainterPath*>(ptr)->translated(static_cast<double>(dx), static_cast<double>(dy)));
}

void* QPainterPath_United(void* ptr, void* p)
{
	return new QPainterPath(static_cast<QPainterPath*>(ptr)->united(*static_cast<QPainterPath*>(p)));
}

void QPainterPath_DestroyQPainterPath(void* ptr)
{
	static_cast<QPainterPath*>(ptr)->~QPainterPath();
}

void* QPainterPathStroker_NewQPainterPathStroker()
{
	return new QPainterPathStroker();
}

void* QPainterPathStroker_NewQPainterPathStroker2(void* pen)
{
	return new QPainterPathStroker(*static_cast<QPen*>(pen));
}

int QPainterPathStroker_CapStyle(void* ptr)
{
	return static_cast<QPainterPathStroker*>(ptr)->capStyle();
}

void* QPainterPathStroker_CreateStroke(void* ptr, void* path)
{
	return new QPainterPath(static_cast<QPainterPathStroker*>(ptr)->createStroke(*static_cast<QPainterPath*>(path)));
}

double QPainterPathStroker_CurveThreshold(void* ptr)
{
	return static_cast<double>(static_cast<QPainterPathStroker*>(ptr)->curveThreshold());
}

double QPainterPathStroker_DashOffset(void* ptr)
{
	return static_cast<double>(static_cast<QPainterPathStroker*>(ptr)->dashOffset());
}

int QPainterPathStroker_JoinStyle(void* ptr)
{
	return static_cast<QPainterPathStroker*>(ptr)->joinStyle();
}

double QPainterPathStroker_MiterLimit(void* ptr)
{
	return static_cast<double>(static_cast<QPainterPathStroker*>(ptr)->miterLimit());
}

void QPainterPathStroker_SetCapStyle(void* ptr, int style)
{
	static_cast<QPainterPathStroker*>(ptr)->setCapStyle(static_cast<Qt::PenCapStyle>(style));
}

void QPainterPathStroker_SetCurveThreshold(void* ptr, double threshold)
{
	static_cast<QPainterPathStroker*>(ptr)->setCurveThreshold(static_cast<double>(threshold));
}

void QPainterPathStroker_SetDashOffset(void* ptr, double offset)
{
	static_cast<QPainterPathStroker*>(ptr)->setDashOffset(static_cast<double>(offset));
}

void QPainterPathStroker_SetDashPattern(void* ptr, int style)
{
	static_cast<QPainterPathStroker*>(ptr)->setDashPattern(static_cast<Qt::PenStyle>(style));
}

void QPainterPathStroker_SetJoinStyle(void* ptr, int style)
{
	static_cast<QPainterPathStroker*>(ptr)->setJoinStyle(static_cast<Qt::PenJoinStyle>(style));
}

void QPainterPathStroker_SetMiterLimit(void* ptr, double limit)
{
	static_cast<QPainterPathStroker*>(ptr)->setMiterLimit(static_cast<double>(limit));
}

void QPainterPathStroker_SetWidth(void* ptr, double width)
{
	static_cast<QPainterPathStroker*>(ptr)->setWidth(static_cast<double>(width));
}

double QPainterPathStroker_Width(void* ptr)
{
	return static_cast<double>(static_cast<QPainterPathStroker*>(ptr)->width());
}

void QPainterPathStroker_DestroyQPainterPathStroker(void* ptr)
{
	static_cast<QPainterPathStroker*>(ptr)->~QPainterPathStroker();
}

int QPalette_NColorRoles_Type()
{
	return QPalette::NColorRoles;
}

void* QPalette_Brush(void* ptr, int group, int role)
{
	return new QBrush(static_cast<QPalette*>(ptr)->brush(static_cast<QPalette::ColorGroup>(group), static_cast<QPalette::ColorRole>(role)));
}

int QPalette_IsEqual(void* ptr, int cg1, int cg2)
{
	return static_cast<QPalette*>(ptr)->isEqual(static_cast<QPalette::ColorGroup>(cg1), static_cast<QPalette::ColorGroup>(cg2));
}

void QPalette_SetBrush2(void* ptr, int group, int role, void* brush)
{
	static_cast<QPalette*>(ptr)->setBrush(static_cast<QPalette::ColorGroup>(group), static_cast<QPalette::ColorRole>(role), *static_cast<QBrush*>(brush));
}

void* QPalette_NewQPalette()
{
	return new QPalette();
}

void* QPalette_NewQPalette8(void* other)
{
	return new QPalette(*static_cast<QPalette*>(other));
}

void* QPalette_NewQPalette3(int button)
{
	return new QPalette(static_cast<Qt::GlobalColor>(button));
}

void* QPalette_NewQPalette5(void* windowText, void* button, void* light, void* dark, void* mid, void* text, void* bright_text, void* base, void* window)
{
	return new QPalette(*static_cast<QBrush*>(windowText), *static_cast<QBrush*>(button), *static_cast<QBrush*>(light), *static_cast<QBrush*>(dark), *static_cast<QBrush*>(mid), *static_cast<QBrush*>(text), *static_cast<QBrush*>(bright_text), *static_cast<QBrush*>(base), *static_cast<QBrush*>(window));
}

void* QPalette_NewQPalette2(void* button)
{
	return new QPalette(*static_cast<QColor*>(button));
}

void* QPalette_NewQPalette4(void* button, void* window)
{
	return new QPalette(*static_cast<QColor*>(button), *static_cast<QColor*>(window));
}

void* QPalette_NewQPalette7(void* p)
{
	return new QPalette(*static_cast<QPalette*>(p));
}

void* QPalette_AlternateBase(void* ptr)
{
	return new QBrush(static_cast<QPalette*>(ptr)->alternateBase());
}

void* QPalette_Base(void* ptr)
{
	return new QBrush(static_cast<QPalette*>(ptr)->base());
}

void* QPalette_BrightText(void* ptr)
{
	return new QBrush(static_cast<QPalette*>(ptr)->brightText());
}

void* QPalette_Brush2(void* ptr, int role)
{
	return new QBrush(static_cast<QPalette*>(ptr)->brush(static_cast<QPalette::ColorRole>(role)));
}

void* QPalette_Button(void* ptr)
{
	return new QBrush(static_cast<QPalette*>(ptr)->button());
}

void* QPalette_ButtonText(void* ptr)
{
	return new QBrush(static_cast<QPalette*>(ptr)->buttonText());
}

long long QPalette_CacheKey(void* ptr)
{
	return static_cast<long long>(static_cast<QPalette*>(ptr)->cacheKey());
}

void* QPalette_Color(void* ptr, int group, int role)
{
	return new QColor(static_cast<QPalette*>(ptr)->color(static_cast<QPalette::ColorGroup>(group), static_cast<QPalette::ColorRole>(role)));
}

void* QPalette_Color2(void* ptr, int role)
{
	return new QColor(static_cast<QPalette*>(ptr)->color(static_cast<QPalette::ColorRole>(role)));
}

int QPalette_CurrentColorGroup(void* ptr)
{
	return static_cast<QPalette*>(ptr)->currentColorGroup();
}

void* QPalette_Dark(void* ptr)
{
	return new QBrush(static_cast<QPalette*>(ptr)->dark());
}

void* QPalette_Highlight(void* ptr)
{
	return new QBrush(static_cast<QPalette*>(ptr)->highlight());
}

void* QPalette_HighlightedText(void* ptr)
{
	return new QBrush(static_cast<QPalette*>(ptr)->highlightedText());
}

int QPalette_IsBrushSet(void* ptr, int cg, int cr)
{
	return static_cast<QPalette*>(ptr)->isBrushSet(static_cast<QPalette::ColorGroup>(cg), static_cast<QPalette::ColorRole>(cr));
}

int QPalette_IsCopyOf(void* ptr, void* p)
{
	return static_cast<QPalette*>(ptr)->isCopyOf(*static_cast<QPalette*>(p));
}

void* QPalette_Light(void* ptr)
{
	return new QBrush(static_cast<QPalette*>(ptr)->light());
}

void* QPalette_Link(void* ptr)
{
	return new QBrush(static_cast<QPalette*>(ptr)->link());
}

void* QPalette_LinkVisited(void* ptr)
{
	return new QBrush(static_cast<QPalette*>(ptr)->linkVisited());
}

void* QPalette_Mid(void* ptr)
{
	return new QBrush(static_cast<QPalette*>(ptr)->mid());
}

void* QPalette_Midlight(void* ptr)
{
	return new QBrush(static_cast<QPalette*>(ptr)->midlight());
}

void* QPalette_Resolve(void* ptr, void* other)
{
	return new QPalette(static_cast<QPalette*>(ptr)->resolve(*static_cast<QPalette*>(other)));
}

void QPalette_SetBrush(void* ptr, int role, void* brush)
{
	static_cast<QPalette*>(ptr)->setBrush(static_cast<QPalette::ColorRole>(role), *static_cast<QBrush*>(brush));
}

void QPalette_SetColor(void* ptr, int group, int role, void* color)
{
	static_cast<QPalette*>(ptr)->setColor(static_cast<QPalette::ColorGroup>(group), static_cast<QPalette::ColorRole>(role), *static_cast<QColor*>(color));
}

void QPalette_SetColor2(void* ptr, int role, void* color)
{
	static_cast<QPalette*>(ptr)->setColor(static_cast<QPalette::ColorRole>(role), *static_cast<QColor*>(color));
}

void QPalette_SetColorGroup(void* ptr, int cg, void* windowText, void* button, void* light, void* dark, void* mid, void* text, void* bright_text, void* base, void* window)
{
	static_cast<QPalette*>(ptr)->setColorGroup(static_cast<QPalette::ColorGroup>(cg), *static_cast<QBrush*>(windowText), *static_cast<QBrush*>(button), *static_cast<QBrush*>(light), *static_cast<QBrush*>(dark), *static_cast<QBrush*>(mid), *static_cast<QBrush*>(text), *static_cast<QBrush*>(bright_text), *static_cast<QBrush*>(base), *static_cast<QBrush*>(window));
}

void QPalette_SetCurrentColorGroup(void* ptr, int cg)
{
	static_cast<QPalette*>(ptr)->setCurrentColorGroup(static_cast<QPalette::ColorGroup>(cg));
}

void* QPalette_Shadow(void* ptr)
{
	return new QBrush(static_cast<QPalette*>(ptr)->shadow());
}

void QPalette_Swap(void* ptr, void* other)
{
	static_cast<QPalette*>(ptr)->swap(*static_cast<QPalette*>(other));
}

void* QPalette_Text(void* ptr)
{
	return new QBrush(static_cast<QPalette*>(ptr)->text());
}

void* QPalette_ToolTipBase(void* ptr)
{
	return new QBrush(static_cast<QPalette*>(ptr)->toolTipBase());
}

void* QPalette_ToolTipText(void* ptr)
{
	return new QBrush(static_cast<QPalette*>(ptr)->toolTipText());
}

void* QPalette_Window(void* ptr)
{
	return new QBrush(static_cast<QPalette*>(ptr)->window());
}

void* QPalette_WindowText(void* ptr)
{
	return new QBrush(static_cast<QPalette*>(ptr)->windowText());
}

void QPalette_DestroyQPalette(void* ptr)
{
	static_cast<QPalette*>(ptr)->~QPalette();
}

class MyQPdfWriter: public QPdfWriter
{
public:
	MyQPdfWriter(QIODevice *device) : QPdfWriter(device) {};
	MyQPdfWriter(const QString &filename) : QPdfWriter(filename) {};
	bool newPage() { return callbackQPdfWriter_NewPage(this, this->objectName().toUtf8().data()) != 0; };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPdfWriter_PaintEngine(const_cast<MyQPdfWriter*>(this), this->objectName().toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQPdfWriter_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQPdfWriter_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQPdfWriter_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQPdfWriter_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQPdfWriter_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQPdfWriter_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQPdfWriter_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQPdfWriter_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQPdfWriter_MetaObject(const_cast<MyQPdfWriter*>(this), this->objectName().toUtf8().data())); };
	void setPageSize(QPagedPaintDevice::PageSize size) { callbackQPdfWriter_SetPageSize2(this, this->objectName().toUtf8().data(), size); };
	void setPageSizeMM(const QSizeF & size) { callbackQPdfWriter_SetPageSizeMM(this, this->objectName().toUtf8().data(), new QSizeF(static_cast<QSizeF>(size).width(), static_cast<QSizeF>(size).height())); };
	int metric(QPaintDevice::PaintDeviceMetric metric) const { return callbackQPdfWriter_Metric(const_cast<MyQPdfWriter*>(this), this->objectName().toUtf8().data(), metric); };
};

void* QPdfWriter_NewQPdfWriter2(void* device)
{
	return new MyQPdfWriter(static_cast<QIODevice*>(device));
}

void* QPdfWriter_NewQPdfWriter(char* filename)
{
	return new MyQPdfWriter(QString(filename));
}

char* QPdfWriter_Creator(void* ptr)
{
	return static_cast<QPdfWriter*>(ptr)->creator().toUtf8().data();
}

int QPdfWriter_NewPage(void* ptr)
{
	return static_cast<QPdfWriter*>(ptr)->newPage();
}

int QPdfWriter_NewPageDefault(void* ptr)
{
	return static_cast<QPdfWriter*>(ptr)->QPdfWriter::newPage();
}

void* QPdfWriter_PageLayout(void* ptr)
{
	return new QPageLayout(static_cast<QPdfWriter*>(ptr)->pageLayout());
}

void* QPdfWriter_PaintEngine(void* ptr)
{
	return static_cast<QPdfWriter*>(ptr)->paintEngine();
}

void* QPdfWriter_PaintEngineDefault(void* ptr)
{
	return static_cast<QPdfWriter*>(ptr)->QPdfWriter::paintEngine();
}

int QPdfWriter_Resolution(void* ptr)
{
	return static_cast<QPdfWriter*>(ptr)->resolution();
}

void QPdfWriter_SetCreator(void* ptr, char* creator)
{
	static_cast<QPdfWriter*>(ptr)->setCreator(QString(creator));
}

int QPdfWriter_SetPageLayout(void* ptr, void* newPageLayout)
{
	return static_cast<QPdfWriter*>(ptr)->setPageLayout(*static_cast<QPageLayout*>(newPageLayout));
}

int QPdfWriter_SetPageMargins(void* ptr, void* margins)
{
	return static_cast<QPdfWriter*>(ptr)->setPageMargins(*static_cast<QMarginsF*>(margins));
}

int QPdfWriter_SetPageMargins2(void* ptr, void* margins, int units)
{
	return static_cast<QPdfWriter*>(ptr)->setPageMargins(*static_cast<QMarginsF*>(margins), static_cast<QPageLayout::Unit>(units));
}

int QPdfWriter_SetPageOrientation(void* ptr, int orientation)
{
	return static_cast<QPdfWriter*>(ptr)->setPageOrientation(static_cast<QPageLayout::Orientation>(orientation));
}

int QPdfWriter_SetPageSize(void* ptr, void* pageSize)
{
	return static_cast<QPdfWriter*>(ptr)->setPageSize(*static_cast<QPageSize*>(pageSize));
}

void QPdfWriter_SetResolution(void* ptr, int resolution)
{
	static_cast<QPdfWriter*>(ptr)->setResolution(resolution);
}

void QPdfWriter_SetTitle(void* ptr, char* title)
{
	static_cast<QPdfWriter*>(ptr)->setTitle(QString(title));
}

char* QPdfWriter_Title(void* ptr)
{
	return static_cast<QPdfWriter*>(ptr)->title().toUtf8().data();
}

void QPdfWriter_DestroyQPdfWriter(void* ptr)
{
	static_cast<QPdfWriter*>(ptr)->~QPdfWriter();
}

void QPdfWriter_TimerEvent(void* ptr, void* event)
{
	static_cast<QPdfWriter*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QPdfWriter_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QPdfWriter*>(ptr)->QPdfWriter::timerEvent(static_cast<QTimerEvent*>(event));
}

void QPdfWriter_ChildEvent(void* ptr, void* event)
{
	static_cast<QPdfWriter*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QPdfWriter_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QPdfWriter*>(ptr)->QPdfWriter::childEvent(static_cast<QChildEvent*>(event));
}

void QPdfWriter_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QPdfWriter*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPdfWriter_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QPdfWriter*>(ptr)->QPdfWriter::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPdfWriter_CustomEvent(void* ptr, void* event)
{
	static_cast<QPdfWriter*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QPdfWriter_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QPdfWriter*>(ptr)->QPdfWriter::customEvent(static_cast<QEvent*>(event));
}

void QPdfWriter_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QPdfWriter*>(ptr), "deleteLater");
}

void QPdfWriter_DeleteLaterDefault(void* ptr)
{
	static_cast<QPdfWriter*>(ptr)->QPdfWriter::deleteLater();
}

void QPdfWriter_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QPdfWriter*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QPdfWriter_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QPdfWriter*>(ptr)->QPdfWriter::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QPdfWriter_Event(void* ptr, void* e)
{
	return static_cast<QPdfWriter*>(ptr)->event(static_cast<QEvent*>(e));
}

int QPdfWriter_EventDefault(void* ptr, void* e)
{
	return static_cast<QPdfWriter*>(ptr)->QPdfWriter::event(static_cast<QEvent*>(e));
}

int QPdfWriter_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QPdfWriter*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QPdfWriter_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QPdfWriter*>(ptr)->QPdfWriter::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QPdfWriter_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QPdfWriter*>(ptr)->metaObject());
}

void* QPdfWriter_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QPdfWriter*>(ptr)->QPdfWriter::metaObject());
}

void QPdfWriter_SetPageSize2(void* ptr, int size)
{
	static_cast<QPdfWriter*>(ptr)->setPageSize(static_cast<QPagedPaintDevice::PageSize>(size));
}

void QPdfWriter_SetPageSize2Default(void* ptr, int size)
{
	static_cast<QPdfWriter*>(ptr)->QPdfWriter::setPageSize(static_cast<QPagedPaintDevice::PageSize>(size));
}

void QPdfWriter_SetPageSizeMM(void* ptr, void* size)
{
	static_cast<QPdfWriter*>(ptr)->setPageSizeMM(*static_cast<QSizeF*>(size));
}

void QPdfWriter_SetPageSizeMMDefault(void* ptr, void* size)
{
	static_cast<QPdfWriter*>(ptr)->QPdfWriter::setPageSizeMM(*static_cast<QSizeF*>(size));
}

int QPdfWriter_Metric(void* ptr, int metric)
{
	return static_cast<QPdfWriter*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

int QPdfWriter_MetricDefault(void* ptr, int metric)
{
	return static_cast<QPdfWriter*>(ptr)->QPdfWriter::metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

void* QPen_NewQPen4(void* brush, double width, int style, int cap, int join)
{
	return new QPen(*static_cast<QBrush*>(brush), static_cast<double>(width), static_cast<Qt::PenStyle>(style), static_cast<Qt::PenCapStyle>(cap), static_cast<Qt::PenJoinStyle>(join));
}

void* QPen_NewQPen5(void* pen)
{
	return new QPen(*static_cast<QPen*>(pen));
}

void* QPen_Color(void* ptr)
{
	return new QColor(static_cast<QPen*>(ptr)->color());
}

void QPen_SetCapStyle(void* ptr, int style)
{
	static_cast<QPen*>(ptr)->setCapStyle(static_cast<Qt::PenCapStyle>(style));
}

void QPen_SetColor(void* ptr, void* color)
{
	static_cast<QPen*>(ptr)->setColor(*static_cast<QColor*>(color));
}

void QPen_SetJoinStyle(void* ptr, int style)
{
	static_cast<QPen*>(ptr)->setJoinStyle(static_cast<Qt::PenJoinStyle>(style));
}

void QPen_SetStyle(void* ptr, int style)
{
	static_cast<QPen*>(ptr)->setStyle(static_cast<Qt::PenStyle>(style));
}

void QPen_SetWidth(void* ptr, int width)
{
	static_cast<QPen*>(ptr)->setWidth(width);
}

int QPen_Style(void* ptr)
{
	return static_cast<QPen*>(ptr)->style();
}

int QPen_Width(void* ptr)
{
	return static_cast<QPen*>(ptr)->width();
}

double QPen_WidthF(void* ptr)
{
	return static_cast<double>(static_cast<QPen*>(ptr)->widthF());
}

void* QPen_NewQPen()
{
	return new QPen();
}

void* QPen_NewQPen6(void* pen)
{
	return new QPen(*static_cast<QPen*>(pen));
}

void* QPen_NewQPen2(int style)
{
	return new QPen(static_cast<Qt::PenStyle>(style));
}

void* QPen_NewQPen3(void* color)
{
	return new QPen(*static_cast<QColor*>(color));
}

void* QPen_Brush(void* ptr)
{
	return new QBrush(static_cast<QPen*>(ptr)->brush());
}

int QPen_CapStyle(void* ptr)
{
	return static_cast<QPen*>(ptr)->capStyle();
}

double QPen_DashOffset(void* ptr)
{
	return static_cast<double>(static_cast<QPen*>(ptr)->dashOffset());
}

int QPen_IsCosmetic(void* ptr)
{
	return static_cast<QPen*>(ptr)->isCosmetic();
}

int QPen_IsSolid(void* ptr)
{
	return static_cast<QPen*>(ptr)->isSolid();
}

int QPen_JoinStyle(void* ptr)
{
	return static_cast<QPen*>(ptr)->joinStyle();
}

double QPen_MiterLimit(void* ptr)
{
	return static_cast<double>(static_cast<QPen*>(ptr)->miterLimit());
}

void QPen_SetBrush(void* ptr, void* brush)
{
	static_cast<QPen*>(ptr)->setBrush(*static_cast<QBrush*>(brush));
}

void QPen_SetCosmetic(void* ptr, int cosmetic)
{
	static_cast<QPen*>(ptr)->setCosmetic(cosmetic != 0);
}

void QPen_SetDashOffset(void* ptr, double offset)
{
	static_cast<QPen*>(ptr)->setDashOffset(static_cast<double>(offset));
}

void QPen_SetMiterLimit(void* ptr, double limit)
{
	static_cast<QPen*>(ptr)->setMiterLimit(static_cast<double>(limit));
}

void QPen_SetWidthF(void* ptr, double width)
{
	static_cast<QPen*>(ptr)->setWidthF(static_cast<double>(width));
}

void QPen_Swap(void* ptr, void* other)
{
	static_cast<QPen*>(ptr)->swap(*static_cast<QPen*>(other));
}

void QPen_DestroyQPen(void* ptr)
{
	static_cast<QPen*>(ptr)->~QPen();
}

class MyQPicture: public QPicture
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQPicture(const QPicture &pic) : QPicture(pic) {};
	MyQPicture(int formatVersion) : QPicture(formatVersion) {};
	int metric(QPaintDevice::PaintDeviceMetric metric) const { return callbackQPicture_Metric(const_cast<MyQPicture*>(this), this->objectNameAbs().toUtf8().data(), metric); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPicture_PaintEngine(const_cast<MyQPicture*>(this), this->objectNameAbs().toUtf8().data())); };
};

char* QPicture_Data(void* ptr)
{
	return QString(static_cast<QPicture*>(ptr)->data()).toUtf8().data();
}

int QPicture_IsNull(void* ptr)
{
	return static_cast<QPicture*>(ptr)->isNull();
}

void* QPicture_NewQPicture2(void* pic)
{
	return new MyQPicture(*static_cast<QPicture*>(pic));
}

void* QPicture_NewQPicture(int formatVersion)
{
	return new MyQPicture(formatVersion);
}

void* QPicture_BoundingRect(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QPicture*>(ptr)->boundingRect()).x(), static_cast<QRect>(static_cast<QPicture*>(ptr)->boundingRect()).y(), static_cast<QRect>(static_cast<QPicture*>(ptr)->boundingRect()).width(), static_cast<QRect>(static_cast<QPicture*>(ptr)->boundingRect()).height());
}

int QPicture_Load2(void* ptr, void* dev, char* format)
{
	return static_cast<QPicture*>(ptr)->load(static_cast<QIODevice*>(dev), const_cast<const char*>(format));
}

int QPicture_Load(void* ptr, char* fileName, char* format)
{
	return static_cast<QPicture*>(ptr)->load(QString(fileName), const_cast<const char*>(format));
}

int QPicture_Play(void* ptr, void* painter)
{
	return static_cast<QPicture*>(ptr)->play(static_cast<QPainter*>(painter));
}

int QPicture_Save2(void* ptr, void* dev, char* format)
{
	return static_cast<QPicture*>(ptr)->save(static_cast<QIODevice*>(dev), const_cast<const char*>(format));
}

int QPicture_Save(void* ptr, char* fileName, char* format)
{
	return static_cast<QPicture*>(ptr)->save(QString(fileName), const_cast<const char*>(format));
}

void QPicture_SetBoundingRect(void* ptr, void* r)
{
	static_cast<QPicture*>(ptr)->setBoundingRect(*static_cast<QRect*>(r));
}

void QPicture_Swap(void* ptr, void* other)
{
	static_cast<QPicture*>(ptr)->swap(*static_cast<QPicture*>(other));
}

void QPicture_DestroyQPicture(void* ptr)
{
	static_cast<QPicture*>(ptr)->~QPicture();
}

char* QPicture_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQPicture*>(static_cast<QPicture*>(ptr))) {
		return static_cast<MyQPicture*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QPicture_BASE").toUtf8().data();
}

void QPicture_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQPicture*>(static_cast<QPicture*>(ptr))) {
		static_cast<MyQPicture*>(ptr)->setObjectNameAbs(QString(name));
	}
}

int QPicture_Metric(void* ptr, int metric)
{
	return static_cast<QPicture*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

int QPicture_MetricDefault(void* ptr, int metric)
{
	return static_cast<QPicture*>(ptr)->QPicture::metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

void* QPicture_PaintEngine(void* ptr)
{
	return static_cast<QPicture*>(ptr)->paintEngine();
}

void* QPicture_PaintEngineDefault(void* ptr)
{
	return static_cast<QPicture*>(ptr)->QPicture::paintEngine();
}

void* QPixelFormat_NewQPixelFormat()
{
	return new QPixelFormat();
}

int QPixelFormat_AlphaPosition(void* ptr)
{
	return static_cast<QPixelFormat*>(ptr)->alphaPosition();
}

int QPixelFormat_AlphaUsage(void* ptr)
{
	return static_cast<QPixelFormat*>(ptr)->alphaUsage();
}

int QPixelFormat_ByteOrder(void* ptr)
{
	return static_cast<QPixelFormat*>(ptr)->byteOrder();
}

int QPixelFormat_ColorModel(void* ptr)
{
	return static_cast<QPixelFormat*>(ptr)->colorModel();
}

int QPixelFormat_Premultiplied(void* ptr)
{
	return static_cast<QPixelFormat*>(ptr)->premultiplied();
}

int QPixelFormat_TypeInterpretation(void* ptr)
{
	return static_cast<QPixelFormat*>(ptr)->typeInterpretation();
}

int QPixelFormat_YuvLayout(void* ptr)
{
	return static_cast<QPixelFormat*>(ptr)->yuvLayout();
}

void* QPixmap_NewQPixmap3(int width, int height)
{
	return new QPixmap(width, height);
}

void* QPixmap_Copy(void* ptr, void* rectangle)
{
	return new QPixmap(static_cast<QPixmap*>(ptr)->copy(*static_cast<QRect*>(rectangle)));
}

int QPixmap_Depth(void* ptr)
{
	return static_cast<QPixmap*>(ptr)->depth();
}

void* QPixmap_QPixmap_FromImage(void* image, int flags)
{
	return new QPixmap(QPixmap::fromImage(*static_cast<QImage*>(image), static_cast<Qt::ImageConversionFlag>(flags)));
}

void* QPixmap_QPixmap_FromImageReader(void* imageReader, int flags)
{
	return new QPixmap(QPixmap::fromImageReader(static_cast<QImageReader*>(imageReader), static_cast<Qt::ImageConversionFlag>(flags)));
}

int QPixmap_Height(void* ptr)
{
	return static_cast<QPixmap*>(ptr)->height();
}

int QPixmap_IsNull(void* ptr)
{
	return static_cast<QPixmap*>(ptr)->isNull();
}

int QPixmap_IsQBitmap(void* ptr)
{
	return static_cast<QPixmap*>(ptr)->isQBitmap();
}

void* QPixmap_Rect(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QPixmap*>(ptr)->rect()).x(), static_cast<QRect>(static_cast<QPixmap*>(ptr)->rect()).y(), static_cast<QRect>(static_cast<QPixmap*>(ptr)->rect()).width(), static_cast<QRect>(static_cast<QPixmap*>(ptr)->rect()).height());
}

void* QPixmap_Scaled(void* ptr, void* size, int aspectRatioMode, int transformMode)
{
	return new QPixmap(static_cast<QPixmap*>(ptr)->scaled(*static_cast<QSize*>(size), static_cast<Qt::AspectRatioMode>(aspectRatioMode), static_cast<Qt::TransformationMode>(transformMode)));
}

void* QPixmap_ScaledToHeight(void* ptr, int height, int mode)
{
	return new QPixmap(static_cast<QPixmap*>(ptr)->scaledToHeight(height, static_cast<Qt::TransformationMode>(mode)));
}

void* QPixmap_ScaledToWidth(void* ptr, int width, int mode)
{
	return new QPixmap(static_cast<QPixmap*>(ptr)->scaledToWidth(width, static_cast<Qt::TransformationMode>(mode)));
}

void* QPixmap_Size(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QPixmap*>(ptr)->size()).width(), static_cast<QSize>(static_cast<QPixmap*>(ptr)->size()).height());
}

int QPixmap_Width(void* ptr)
{
	return static_cast<QPixmap*>(ptr)->width();
}

void* QPixmap_NewQPixmap()
{
	return new QPixmap();
}

void* QPixmap_NewQPixmap6(void* pixmap)
{
	return new QPixmap(*static_cast<QPixmap*>(pixmap));
}

void* QPixmap_NewQPixmap8(void* size)
{
	return new QPixmap(*static_cast<QSize*>(size));
}

void* QPixmap_NewQPixmap4(char* fileName, char* format, int flags)
{
	return new QPixmap(QString(fileName), const_cast<const char*>(format), static_cast<Qt::ImageConversionFlag>(flags));
}

long long QPixmap_CacheKey(void* ptr)
{
	return static_cast<long long>(static_cast<QPixmap*>(ptr)->cacheKey());
}

int QPixmap_ConvertFromImage(void* ptr, void* image, int flags)
{
	return static_cast<QPixmap*>(ptr)->convertFromImage(*static_cast<QImage*>(image), static_cast<Qt::ImageConversionFlag>(flags));
}

void* QPixmap_Copy2(void* ptr, int x, int y, int width, int height)
{
	return new QPixmap(static_cast<QPixmap*>(ptr)->copy(x, y, width, height));
}

int QPixmap_QPixmap_DefaultDepth()
{
	return QPixmap::defaultDepth();
}

void QPixmap_Detach(void* ptr)
{
	static_cast<QPixmap*>(ptr)->detach();
}

double QPixmap_DevicePixelRatio(void* ptr)
{
	return static_cast<double>(static_cast<QPixmap*>(ptr)->devicePixelRatio());
}

void QPixmap_Fill(void* ptr, void* color)
{
	static_cast<QPixmap*>(ptr)->fill(*static_cast<QColor*>(color));
}

void* QPixmap_QPixmap_FromImage2(void* image, int flags)
{
	return new QPixmap(QPixmap::fromImage(*static_cast<QImage*>(image), static_cast<Qt::ImageConversionFlag>(flags)));
}

int QPixmap_HasAlpha(void* ptr)
{
	return static_cast<QPixmap*>(ptr)->hasAlpha();
}

int QPixmap_HasAlphaChannel(void* ptr)
{
	return static_cast<QPixmap*>(ptr)->hasAlphaChannel();
}

int QPixmap_Load(void* ptr, char* fileName, char* format, int flags)
{
	return static_cast<QPixmap*>(ptr)->load(QString(fileName), const_cast<const char*>(format), static_cast<Qt::ImageConversionFlag>(flags));
}

int QPixmap_LoadFromData2(void* ptr, char* data, char* format, int flags)
{
	return static_cast<QPixmap*>(ptr)->loadFromData(QByteArray(data), const_cast<const char*>(format), static_cast<Qt::ImageConversionFlag>(flags));
}

int QPixmap_Save2(void* ptr, void* device, char* format, int quality)
{
	return static_cast<QPixmap*>(ptr)->save(static_cast<QIODevice*>(device), const_cast<const char*>(format), quality);
}

int QPixmap_Save(void* ptr, char* fileName, char* format, int quality)
{
	return static_cast<QPixmap*>(ptr)->save(QString(fileName), const_cast<const char*>(format), quality);
}

void* QPixmap_Scaled2(void* ptr, int width, int height, int aspectRatioMode, int transformMode)
{
	return new QPixmap(static_cast<QPixmap*>(ptr)->scaled(width, height, static_cast<Qt::AspectRatioMode>(aspectRatioMode), static_cast<Qt::TransformationMode>(transformMode)));
}

void QPixmap_Scroll2(void* ptr, int dx, int dy, void* rect, void* exposed)
{
	static_cast<QPixmap*>(ptr)->scroll(dx, dy, *static_cast<QRect*>(rect), static_cast<QRegion*>(exposed));
}

void QPixmap_Scroll(void* ptr, int dx, int dy, int x, int y, int width, int height, void* exposed)
{
	static_cast<QPixmap*>(ptr)->scroll(dx, dy, x, y, width, height, static_cast<QRegion*>(exposed));
}

void QPixmap_SetDevicePixelRatio(void* ptr, double scaleFactor)
{
	static_cast<QPixmap*>(ptr)->setDevicePixelRatio(static_cast<double>(scaleFactor));
}

void QPixmap_SetMask(void* ptr, void* mask)
{
	static_cast<QPixmap*>(ptr)->setMask(*static_cast<QBitmap*>(mask));
}

void QPixmap_Swap(void* ptr, void* other)
{
	static_cast<QPixmap*>(ptr)->swap(*static_cast<QPixmap*>(other));
}

void* QPixmap_ToImage(void* ptr)
{
	return new QImage(static_cast<QPixmap*>(ptr)->toImage());
}

void* QPixmap_Transformed(void* ptr, void* transform, int mode)
{
	return new QPixmap(static_cast<QPixmap*>(ptr)->transformed(*static_cast<QTransform*>(transform), static_cast<Qt::TransformationMode>(mode)));
}

void QPixmap_DestroyQPixmap(void* ptr)
{
	static_cast<QPixmap*>(ptr)->~QPixmap();
}

int QPixmap_Metric(void* ptr, int metric)
{
	return static_cast<QPixmap*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

int QPixmap_MetricDefault(void* ptr, int metric)
{
	return static_cast<QPixmap*>(ptr)->QPixmap::metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

void* QPixmap_PaintEngine(void* ptr)
{
	return static_cast<QPixmap*>(ptr)->paintEngine();
}

void* QPixmap_PaintEngineDefault(void* ptr)
{
	return static_cast<QPixmap*>(ptr)->QPixmap::paintEngine();
}

int QPixmapCache_QPixmapCache_CacheLimit()
{
	return QPixmapCache::cacheLimit();
}

void QPixmapCache_QPixmapCache_Clear()
{
	QPixmapCache::clear();
}

int QPixmapCache_QPixmapCache_Find(char* key, void* pixmap)
{
	return QPixmapCache::find(QString(key), static_cast<QPixmap*>(pixmap));
}

int QPixmapCache_QPixmapCache_Insert(char* key, void* pixmap)
{
	return QPixmapCache::insert(QString(key), *static_cast<QPixmap*>(pixmap));
}

void QPixmapCache_QPixmapCache_Remove(char* key)
{
	QPixmapCache::remove(QString(key));
}

void QPixmapCache_QPixmapCache_SetCacheLimit(int n)
{
	QPixmapCache::setCacheLimit(n);
}

void* QPlatformSurfaceEvent_NewQPlatformSurfaceEvent(int surfaceEventType)
{
	return new QPlatformSurfaceEvent(static_cast<QPlatformSurfaceEvent::SurfaceEventType>(surfaceEventType));
}

int QPlatformSurfaceEvent_SurfaceEventType(void* ptr)
{
	return static_cast<QPlatformSurfaceEvent*>(ptr)->surfaceEventType();
}

void* QPolygon_NewQPolygon5(void* rectangle, int closed)
{
	return new QPolygon(*static_cast<QRect*>(rectangle), closed != 0);
}

int QPolygon_ContainsPoint(void* ptr, void* point, int fillRule)
{
	return static_cast<QPolygon*>(ptr)->containsPoint(*static_cast<QPoint*>(point), static_cast<Qt::FillRule>(fillRule));
}

void QPolygon_PutPoints3(void* ptr, int index, int nPoints, void* fromPolygon, int fromIndex)
{
	static_cast<QPolygon*>(ptr)->putPoints(index, nPoints, *static_cast<QPolygon*>(fromPolygon), fromIndex);
}

void* QPolygon_NewQPolygon()
{
	return new QPolygon();
}

void* QPolygon_NewQPolygon3(void* polygon)
{
	return new QPolygon(*static_cast<QPolygon*>(polygon));
}

void* QPolygon_NewQPolygon2(int size)
{
	return new QPolygon(size);
}

void* QPolygon_BoundingRect(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QPolygon*>(ptr)->boundingRect()).x(), static_cast<QRect>(static_cast<QPolygon*>(ptr)->boundingRect()).y(), static_cast<QRect>(static_cast<QPolygon*>(ptr)->boundingRect()).width(), static_cast<QRect>(static_cast<QPolygon*>(ptr)->boundingRect()).height());
}

void* QPolygon_Intersected(void* ptr, void* r)
{
	return new QPolygon(static_cast<QPolygon*>(ptr)->intersected(*static_cast<QPolygon*>(r)));
}

void* QPolygon_Point2(void* ptr, int index)
{
	return new QPoint(static_cast<QPoint>(static_cast<QPolygon*>(ptr)->point(index)).x(), static_cast<QPoint>(static_cast<QPolygon*>(ptr)->point(index)).y());
}

void QPolygon_Point(void* ptr, int index, int x, int y)
{
	static_cast<QPolygon*>(ptr)->point(index, &x, &y);
}

void QPolygon_SetPoint2(void* ptr, int index, void* point)
{
	static_cast<QPolygon*>(ptr)->setPoint(index, *static_cast<QPoint*>(point));
}

void QPolygon_SetPoint(void* ptr, int index, int x, int y)
{
	static_cast<QPolygon*>(ptr)->setPoint(index, x, y);
}

void QPolygon_SetPoints(void* ptr, int nPoints, int points)
{
	static_cast<QPolygon*>(ptr)->setPoints(nPoints, &points);
}

void* QPolygon_Subtracted(void* ptr, void* r)
{
	return new QPolygon(static_cast<QPolygon*>(ptr)->subtracted(*static_cast<QPolygon*>(r)));
}

void QPolygon_Swap(void* ptr, void* other)
{
	static_cast<QPolygon*>(ptr)->swap(*static_cast<QPolygon*>(other));
}

void QPolygon_Translate2(void* ptr, void* offset)
{
	static_cast<QPolygon*>(ptr)->translate(*static_cast<QPoint*>(offset));
}

void QPolygon_Translate(void* ptr, int dx, int dy)
{
	static_cast<QPolygon*>(ptr)->translate(dx, dy);
}

void* QPolygon_Translated2(void* ptr, void* offset)
{
	return new QPolygon(static_cast<QPolygon*>(ptr)->translated(*static_cast<QPoint*>(offset)));
}

void* QPolygon_Translated(void* ptr, int dx, int dy)
{
	return new QPolygon(static_cast<QPolygon*>(ptr)->translated(dx, dy));
}

void* QPolygon_United(void* ptr, void* r)
{
	return new QPolygon(static_cast<QPolygon*>(ptr)->united(*static_cast<QPolygon*>(r)));
}

void QPolygon_DestroyQPolygon(void* ptr)
{
	static_cast<QPolygon*>(ptr)->~QPolygon();
}

void* QPolygonF_NewQPolygonF6(void* polygon)
{
	return new QPolygonF(*static_cast<QPolygon*>(polygon));
}

void* QPolygonF_NewQPolygonF5(void* rectangle)
{
	return new QPolygonF(*static_cast<QRectF*>(rectangle));
}

int QPolygonF_ContainsPoint(void* ptr, void* point, int fillRule)
{
	return static_cast<QPolygonF*>(ptr)->containsPoint(*static_cast<QPointF*>(point), static_cast<Qt::FillRule>(fillRule));
}

void* QPolygonF_NewQPolygonF()
{
	return new QPolygonF();
}

void* QPolygonF_NewQPolygonF3(void* polygon)
{
	return new QPolygonF(*static_cast<QPolygonF*>(polygon));
}

void* QPolygonF_NewQPolygonF2(int size)
{
	return new QPolygonF(size);
}

void* QPolygonF_BoundingRect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QPolygonF*>(ptr)->boundingRect()).x(), static_cast<QRectF>(static_cast<QPolygonF*>(ptr)->boundingRect()).y(), static_cast<QRectF>(static_cast<QPolygonF*>(ptr)->boundingRect()).width(), static_cast<QRectF>(static_cast<QPolygonF*>(ptr)->boundingRect()).height());
}

void* QPolygonF_Intersected(void* ptr, void* r)
{
	return new QPolygonF(static_cast<QPolygonF*>(ptr)->intersected(*static_cast<QPolygonF*>(r)));
}

int QPolygonF_IsClosed(void* ptr)
{
	return static_cast<QPolygonF*>(ptr)->isClosed();
}

void* QPolygonF_Subtracted(void* ptr, void* r)
{
	return new QPolygonF(static_cast<QPolygonF*>(ptr)->subtracted(*static_cast<QPolygonF*>(r)));
}

void QPolygonF_Swap(void* ptr, void* other)
{
	static_cast<QPolygonF*>(ptr)->swap(*static_cast<QPolygonF*>(other));
}

void* QPolygonF_ToPolygon(void* ptr)
{
	return new QPolygon(static_cast<QPolygonF*>(ptr)->toPolygon());
}

void QPolygonF_Translate(void* ptr, void* offset)
{
	static_cast<QPolygonF*>(ptr)->translate(*static_cast<QPointF*>(offset));
}

void QPolygonF_Translate2(void* ptr, double dx, double dy)
{
	static_cast<QPolygonF*>(ptr)->translate(static_cast<double>(dx), static_cast<double>(dy));
}

void* QPolygonF_Translated(void* ptr, void* offset)
{
	return new QPolygonF(static_cast<QPolygonF*>(ptr)->translated(*static_cast<QPointF*>(offset)));
}

void* QPolygonF_Translated2(void* ptr, double dx, double dy)
{
	return new QPolygonF(static_cast<QPolygonF*>(ptr)->translated(static_cast<double>(dx), static_cast<double>(dy)));
}

void* QPolygonF_United(void* ptr, void* r)
{
	return new QPolygonF(static_cast<QPolygonF*>(ptr)->united(*static_cast<QPolygonF*>(r)));
}

void QPolygonF_DestroyQPolygonF(void* ptr)
{
	static_cast<QPolygonF*>(ptr)->~QPolygonF();
}

void* QQuaternion_NewQQuaternion()
{
	return new QQuaternion();
}

void* QQuaternion_NewQQuaternion5(void* vector)
{
	return new QQuaternion(*static_cast<QVector4D*>(vector));
}

void QQuaternion_GetAxes(void* ptr, void* xAxis, void* yAxis, void* zAxis)
{
	static_cast<QQuaternion*>(ptr)->getAxes(static_cast<QVector3D*>(xAxis), static_cast<QVector3D*>(yAxis), static_cast<QVector3D*>(zAxis));
}

int QQuaternion_IsIdentity(void* ptr)
{
	return static_cast<QQuaternion*>(ptr)->isIdentity();
}

int QQuaternion_IsNull(void* ptr)
{
	return static_cast<QQuaternion*>(ptr)->isNull();
}

void QQuaternion_Normalize(void* ptr)
{
	static_cast<QQuaternion*>(ptr)->normalize();
}

void QQuaternion_SetVector(void* ptr, void* vector)
{
	static_cast<QQuaternion*>(ptr)->setVector(*static_cast<QVector3D*>(vector));
}

void* QRadialGradient_NewQRadialGradient()
{
	return new QRadialGradient();
}

void* QRadialGradient_NewQRadialGradient6(void* center, double centerRadius, void* focalPoint, double focalRadius)
{
	return new QRadialGradient(*static_cast<QPointF*>(center), static_cast<double>(centerRadius), *static_cast<QPointF*>(focalPoint), static_cast<double>(focalRadius));
}

void* QRadialGradient_NewQRadialGradient4(void* center, double radius)
{
	return new QRadialGradient(*static_cast<QPointF*>(center), static_cast<double>(radius));
}

void* QRadialGradient_NewQRadialGradient2(void* center, double radius, void* focalPoint)
{
	return new QRadialGradient(*static_cast<QPointF*>(center), static_cast<double>(radius), *static_cast<QPointF*>(focalPoint));
}

void* QRadialGradient_NewQRadialGradient7(double cx, double cy, double centerRadius, double fx, double fy, double focalRadius)
{
	return new QRadialGradient(static_cast<double>(cx), static_cast<double>(cy), static_cast<double>(centerRadius), static_cast<double>(fx), static_cast<double>(fy), static_cast<double>(focalRadius));
}

void* QRadialGradient_NewQRadialGradient5(double cx, double cy, double radius)
{
	return new QRadialGradient(static_cast<double>(cx), static_cast<double>(cy), static_cast<double>(radius));
}

void* QRadialGradient_NewQRadialGradient3(double cx, double cy, double radius, double fx, double fy)
{
	return new QRadialGradient(static_cast<double>(cx), static_cast<double>(cy), static_cast<double>(radius), static_cast<double>(fx), static_cast<double>(fy));
}

void* QRadialGradient_Center(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QRadialGradient*>(ptr)->center()).x(), static_cast<QPointF>(static_cast<QRadialGradient*>(ptr)->center()).y());
}

double QRadialGradient_CenterRadius(void* ptr)
{
	return static_cast<double>(static_cast<QRadialGradient*>(ptr)->centerRadius());
}

void* QRadialGradient_FocalPoint(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QRadialGradient*>(ptr)->focalPoint()).x(), static_cast<QPointF>(static_cast<QRadialGradient*>(ptr)->focalPoint()).y());
}

double QRadialGradient_FocalRadius(void* ptr)
{
	return static_cast<double>(static_cast<QRadialGradient*>(ptr)->focalRadius());
}

double QRadialGradient_Radius(void* ptr)
{
	return static_cast<double>(static_cast<QRadialGradient*>(ptr)->radius());
}

void QRadialGradient_SetCenter(void* ptr, void* center)
{
	static_cast<QRadialGradient*>(ptr)->setCenter(*static_cast<QPointF*>(center));
}

void QRadialGradient_SetCenter2(void* ptr, double x, double y)
{
	static_cast<QRadialGradient*>(ptr)->setCenter(static_cast<double>(x), static_cast<double>(y));
}

void QRadialGradient_SetCenterRadius(void* ptr, double radius)
{
	static_cast<QRadialGradient*>(ptr)->setCenterRadius(static_cast<double>(radius));
}

void QRadialGradient_SetFocalPoint(void* ptr, void* focalPoint)
{
	static_cast<QRadialGradient*>(ptr)->setFocalPoint(*static_cast<QPointF*>(focalPoint));
}

void QRadialGradient_SetFocalPoint2(void* ptr, double x, double y)
{
	static_cast<QRadialGradient*>(ptr)->setFocalPoint(static_cast<double>(x), static_cast<double>(y));
}

void QRadialGradient_SetFocalRadius(void* ptr, double radius)
{
	static_cast<QRadialGradient*>(ptr)->setFocalRadius(static_cast<double>(radius));
}

void QRadialGradient_SetRadius(void* ptr, double radius)
{
	static_cast<QRadialGradient*>(ptr)->setRadius(static_cast<double>(radius));
}

void* QRasterWindow_NewQRasterWindow(void* parent)
{
	return new QRasterWindow(static_cast<QWindow*>(parent));
}

void QRasterWindow_PaintEvent(void* ptr, void* event)
{
	static_cast<QRasterWindow*>(ptr)->paintEvent(static_cast<QPaintEvent*>(event));
}

void QRasterWindow_PaintEventDefault(void* ptr, void* event)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::paintEvent(static_cast<QPaintEvent*>(event));
}

void QRasterWindow_Update3(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "update");
}

void QRasterWindow_Update3Default(void* ptr)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::update();
}

void QRasterWindow_SetHeight(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "setHeight", Q_ARG(int, arg));
}

void QRasterWindow_SetHeightDefault(void* ptr, int arg)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::setHeight(arg);
}

void QRasterWindow_SetMaximumHeight(void* ptr, int h)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "setMaximumHeight", Q_ARG(int, h));
}

void QRasterWindow_SetMaximumHeightDefault(void* ptr, int h)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::setMaximumHeight(h);
}

void QRasterWindow_SetMaximumWidth(void* ptr, int w)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "setMaximumWidth", Q_ARG(int, w));
}

void QRasterWindow_SetMaximumWidthDefault(void* ptr, int w)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::setMaximumWidth(w);
}

void QRasterWindow_SetMinimumHeight(void* ptr, int h)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "setMinimumHeight", Q_ARG(int, h));
}

void QRasterWindow_SetMinimumHeightDefault(void* ptr, int h)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::setMinimumHeight(h);
}

void QRasterWindow_SetMinimumWidth(void* ptr, int w)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "setMinimumWidth", Q_ARG(int, w));
}

void QRasterWindow_SetMinimumWidthDefault(void* ptr, int w)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::setMinimumWidth(w);
}

void QRasterWindow_SetTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "setTitle", Q_ARG(QString, QString(vqs)));
}

void QRasterWindow_SetTitleDefault(void* ptr, char* vqs)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::setTitle(QString(vqs));
}

void QRasterWindow_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QRasterWindow_SetVisibleDefault(void* ptr, int visible)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::setVisible(visible != 0);
}

void QRasterWindow_SetWidth(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "setWidth", Q_ARG(int, arg));
}

void QRasterWindow_SetWidthDefault(void* ptr, int arg)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::setWidth(arg);
}

void QRasterWindow_SetX(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "setX", Q_ARG(int, arg));
}

void QRasterWindow_SetXDefault(void* ptr, int arg)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::setX(arg);
}

void QRasterWindow_SetY(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "setY", Q_ARG(int, arg));
}

void QRasterWindow_SetYDefault(void* ptr, int arg)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::setY(arg);
}

void QRasterWindow_Alert(void* ptr, int msec)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "alert", Q_ARG(int, msec));
}

void QRasterWindow_AlertDefault(void* ptr, int msec)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::alert(msec);
}

int QRasterWindow_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QRasterWindow_CloseDefault(void* ptr)
{
	return static_cast<QRasterWindow*>(ptr)->QRasterWindow::close();
}

int QRasterWindow_Event(void* ptr, void* ev)
{
	return static_cast<QRasterWindow*>(ptr)->event(static_cast<QEvent*>(ev));
}

int QRasterWindow_EventDefault(void* ptr, void* ev)
{
	return static_cast<QRasterWindow*>(ptr)->QRasterWindow::event(static_cast<QEvent*>(ev));
}

void QRasterWindow_ExposeEvent(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->exposeEvent(static_cast<QExposeEvent*>(ev));
}

void QRasterWindow_ExposeEventDefault(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::exposeEvent(static_cast<QExposeEvent*>(ev));
}

void QRasterWindow_FocusInEvent(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(ev));
}

void QRasterWindow_FocusInEventDefault(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::focusInEvent(static_cast<QFocusEvent*>(ev));
}

void* QRasterWindow_FocusObject(void* ptr)
{
	return static_cast<QRasterWindow*>(ptr)->focusObject();
}

void* QRasterWindow_FocusObjectDefault(void* ptr)
{
	return static_cast<QRasterWindow*>(ptr)->QRasterWindow::focusObject();
}

void QRasterWindow_FocusOutEvent(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void QRasterWindow_FocusOutEventDefault(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void* QRasterWindow_Format(void* ptr)
{
	return new QSurfaceFormat(static_cast<QRasterWindow*>(ptr)->format());
}

void* QRasterWindow_FormatDefault(void* ptr)
{
	return new QSurfaceFormat(static_cast<QRasterWindow*>(ptr)->QRasterWindow::format());
}

void QRasterWindow_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "hide");
}

void QRasterWindow_HideDefault(void* ptr)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::hide();
}

void QRasterWindow_HideEvent(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->hideEvent(static_cast<QHideEvent*>(ev));
}

void QRasterWindow_HideEventDefault(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::hideEvent(static_cast<QHideEvent*>(ev));
}

void QRasterWindow_KeyPressEvent(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(ev));
}

void QRasterWindow_KeyPressEventDefault(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::keyPressEvent(static_cast<QKeyEvent*>(ev));
}

void QRasterWindow_KeyReleaseEvent(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(ev));
}

void QRasterWindow_KeyReleaseEventDefault(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::keyReleaseEvent(static_cast<QKeyEvent*>(ev));
}

void QRasterWindow_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "lower");
}

void QRasterWindow_LowerDefault(void* ptr)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::lower();
}

void QRasterWindow_MouseDoubleClickEvent(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(ev));
}

void QRasterWindow_MouseDoubleClickEventDefault(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::mouseDoubleClickEvent(static_cast<QMouseEvent*>(ev));
}

void QRasterWindow_MouseMoveEvent(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(ev));
}

void QRasterWindow_MouseMoveEventDefault(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::mouseMoveEvent(static_cast<QMouseEvent*>(ev));
}

void QRasterWindow_MousePressEvent(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(ev));
}

void QRasterWindow_MousePressEventDefault(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::mousePressEvent(static_cast<QMouseEvent*>(ev));
}

void QRasterWindow_MouseReleaseEvent(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(ev));
}

void QRasterWindow_MouseReleaseEventDefault(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::mouseReleaseEvent(static_cast<QMouseEvent*>(ev));
}

void QRasterWindow_MoveEvent(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->moveEvent(static_cast<QMoveEvent*>(ev));
}

void QRasterWindow_MoveEventDefault(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::moveEvent(static_cast<QMoveEvent*>(ev));
}

int QRasterWindow_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QRasterWindow*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int QRasterWindow_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QRasterWindow*>(ptr)->QRasterWindow::nativeEvent(QByteArray(eventType), message, &result);
}

void QRasterWindow_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "raise");
}

void QRasterWindow_RaiseDefault(void* ptr)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::raise();
}

void QRasterWindow_RequestActivate(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "requestActivate");
}

void QRasterWindow_RequestActivateDefault(void* ptr)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::requestActivate();
}

void QRasterWindow_RequestUpdate(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "requestUpdate");
}

void QRasterWindow_RequestUpdateDefault(void* ptr)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::requestUpdate();
}

void QRasterWindow_ResizeEvent(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(ev));
}

void QRasterWindow_ResizeEventDefault(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::resizeEvent(static_cast<QResizeEvent*>(ev));
}

void QRasterWindow_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "show");
}

void QRasterWindow_ShowDefault(void* ptr)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::show();
}

void QRasterWindow_ShowEvent(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->showEvent(static_cast<QShowEvent*>(ev));
}

void QRasterWindow_ShowEventDefault(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::showEvent(static_cast<QShowEvent*>(ev));
}

void QRasterWindow_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "showFullScreen");
}

void QRasterWindow_ShowFullScreenDefault(void* ptr)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::showFullScreen();
}

void QRasterWindow_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "showMaximized");
}

void QRasterWindow_ShowMaximizedDefault(void* ptr)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::showMaximized();
}

void QRasterWindow_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "showMinimized");
}

void QRasterWindow_ShowMinimizedDefault(void* ptr)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::showMinimized();
}

void QRasterWindow_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "showNormal");
}

void QRasterWindow_ShowNormalDefault(void* ptr)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::showNormal();
}

void* QRasterWindow_Size(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QRasterWindow*>(ptr)->size()).width(), static_cast<QSize>(static_cast<QRasterWindow*>(ptr)->size()).height());
}

void* QRasterWindow_SizeDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QRasterWindow*>(ptr)->QRasterWindow::size()).width(), static_cast<QSize>(static_cast<QRasterWindow*>(ptr)->QRasterWindow::size()).height());
}

int QRasterWindow_SurfaceType(void* ptr)
{
	return static_cast<QRasterWindow*>(ptr)->surfaceType();
}

int QRasterWindow_SurfaceTypeDefault(void* ptr)
{
	return static_cast<QRasterWindow*>(ptr)->QRasterWindow::surfaceType();
}

void QRasterWindow_TabletEvent(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(ev));
}

void QRasterWindow_TabletEventDefault(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::tabletEvent(static_cast<QTabletEvent*>(ev));
}

void QRasterWindow_TouchEvent(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->touchEvent(static_cast<QTouchEvent*>(ev));
}

void QRasterWindow_TouchEventDefault(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::touchEvent(static_cast<QTouchEvent*>(ev));
}

void QRasterWindow_WheelEvent(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(ev));
}

void QRasterWindow_WheelEventDefault(void* ptr, void* ev)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::wheelEvent(static_cast<QWheelEvent*>(ev));
}

void QRasterWindow_TimerEvent(void* ptr, void* event)
{
	static_cast<QRasterWindow*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QRasterWindow_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::timerEvent(static_cast<QTimerEvent*>(event));
}

void QRasterWindow_ChildEvent(void* ptr, void* event)
{
	static_cast<QRasterWindow*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QRasterWindow_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::childEvent(static_cast<QChildEvent*>(event));
}

void QRasterWindow_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QRasterWindow*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRasterWindow_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRasterWindow_CustomEvent(void* ptr, void* event)
{
	static_cast<QRasterWindow*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QRasterWindow_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::customEvent(static_cast<QEvent*>(event));
}

void QRasterWindow_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRasterWindow*>(ptr), "deleteLater");
}

void QRasterWindow_DeleteLaterDefault(void* ptr)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::deleteLater();
}

void QRasterWindow_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QRasterWindow*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRasterWindow_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRasterWindow*>(ptr)->QRasterWindow::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QRasterWindow_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QRasterWindow*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QRasterWindow_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QRasterWindow*>(ptr)->QRasterWindow::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QRasterWindow_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QRasterWindow*>(ptr)->metaObject());
}

void* QRasterWindow_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QRasterWindow*>(ptr)->QRasterWindow::metaObject());
}

int QRasterWindow_Metric(void* ptr, int metric)
{
	return static_cast<QRasterWindow*>(ptr)->metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

int QRasterWindow_MetricDefault(void* ptr, int metric)
{
	return static_cast<QRasterWindow*>(ptr)->QRasterWindow::metric(static_cast<QPaintDevice::PaintDeviceMetric>(metric));
}

void* QRasterWindow_PaintEngine(void* ptr)
{
	return static_cast<QRasterWindow*>(ptr)->paintEngine();
}

void* QRasterWindow_PaintEngineDefault(void* ptr)
{
	return static_cast<QRasterWindow*>(ptr)->QRasterWindow::paintEngine();
}

void* QRawFont_NewQRawFont()
{
	return new QRawFont();
}

void* QRawFont_NewQRawFont3(char* fontData, double pixelSize, int hintingPreference)
{
	return new QRawFont(QByteArray(fontData), static_cast<double>(pixelSize), static_cast<QFont::HintingPreference>(hintingPreference));
}

void* QRawFont_NewQRawFont4(void* other)
{
	return new QRawFont(*static_cast<QRawFont*>(other));
}

void* QRawFont_NewQRawFont2(char* fileName, double pixelSize, int hintingPreference)
{
	return new QRawFont(QString(fileName), static_cast<double>(pixelSize), static_cast<QFont::HintingPreference>(hintingPreference));
}

double QRawFont_Ascent(void* ptr)
{
	return static_cast<double>(static_cast<QRawFont*>(ptr)->ascent());
}

double QRawFont_AverageCharWidth(void* ptr)
{
	return static_cast<double>(static_cast<QRawFont*>(ptr)->averageCharWidth());
}

double QRawFont_Descent(void* ptr)
{
	return static_cast<double>(static_cast<QRawFont*>(ptr)->descent());
}

char* QRawFont_FamilyName(void* ptr)
{
	return static_cast<QRawFont*>(ptr)->familyName().toUtf8().data();
}

char* QRawFont_FontTable(void* ptr, char* tagName)
{
	return QString(static_cast<QRawFont*>(ptr)->fontTable(const_cast<const char*>(tagName))).toUtf8().data();
}

void* QRawFont_QRawFont_FromFont(void* font, int writingSystem)
{
	return new QRawFont(QRawFont::fromFont(*static_cast<QFont*>(font), static_cast<QFontDatabase::WritingSystem>(writingSystem)));
}

int QRawFont_HintingPreference(void* ptr)
{
	return static_cast<QRawFont*>(ptr)->hintingPreference();
}

int QRawFont_IsValid(void* ptr)
{
	return static_cast<QRawFont*>(ptr)->isValid();
}

double QRawFont_Leading(void* ptr)
{
	return static_cast<double>(static_cast<QRawFont*>(ptr)->leading());
}

double QRawFont_LineThickness(void* ptr)
{
	return static_cast<double>(static_cast<QRawFont*>(ptr)->lineThickness());
}

void QRawFont_LoadFromData(void* ptr, char* fontData, double pixelSize, int hintingPreference)
{
	static_cast<QRawFont*>(ptr)->loadFromData(QByteArray(fontData), static_cast<double>(pixelSize), static_cast<QFont::HintingPreference>(hintingPreference));
}

void QRawFont_LoadFromFile(void* ptr, char* fileName, double pixelSize, int hintingPreference)
{
	static_cast<QRawFont*>(ptr)->loadFromFile(QString(fileName), static_cast<double>(pixelSize), static_cast<QFont::HintingPreference>(hintingPreference));
}

double QRawFont_MaxCharWidth(void* ptr)
{
	return static_cast<double>(static_cast<QRawFont*>(ptr)->maxCharWidth());
}

double QRawFont_PixelSize(void* ptr)
{
	return static_cast<double>(static_cast<QRawFont*>(ptr)->pixelSize());
}

void QRawFont_SetPixelSize(void* ptr, double pixelSize)
{
	static_cast<QRawFont*>(ptr)->setPixelSize(static_cast<double>(pixelSize));
}

int QRawFont_Style(void* ptr)
{
	return static_cast<QRawFont*>(ptr)->style();
}

char* QRawFont_StyleName(void* ptr)
{
	return static_cast<QRawFont*>(ptr)->styleName().toUtf8().data();
}

int QRawFont_SupportsCharacter(void* ptr, void* character)
{
	return static_cast<QRawFont*>(ptr)->supportsCharacter(*static_cast<QChar*>(character));
}

void QRawFont_Swap(void* ptr, void* other)
{
	static_cast<QRawFont*>(ptr)->swap(*static_cast<QRawFont*>(other));
}

double QRawFont_UnderlinePosition(void* ptr)
{
	return static_cast<double>(static_cast<QRawFont*>(ptr)->underlinePosition());
}

double QRawFont_UnitsPerEm(void* ptr)
{
	return static_cast<double>(static_cast<QRawFont*>(ptr)->unitsPerEm());
}

int QRawFont_Weight(void* ptr)
{
	return static_cast<QRawFont*>(ptr)->weight();
}

double QRawFont_XHeight(void* ptr)
{
	return static_cast<double>(static_cast<QRawFont*>(ptr)->xHeight());
}

void QRawFont_DestroyQRawFont(void* ptr)
{
	static_cast<QRawFont*>(ptr)->~QRawFont();
}

class MyQRegExpValidator: public QRegExpValidator
{
public:
	MyQRegExpValidator(QObject *parent) : QRegExpValidator(parent) {};
	MyQRegExpValidator(const QRegExp &rx, QObject *parent) : QRegExpValidator(rx, parent) {};
	QValidator::State validate(QString & input, int & pos) const { return static_cast<QValidator::State>(callbackQRegExpValidator_Validate(const_cast<MyQRegExpValidator*>(this), this->objectName().toUtf8().data(), input.toUtf8().data(), pos)); };
	void fixup(QString & input) const { callbackQRegExpValidator_Fixup(const_cast<MyQRegExpValidator*>(this), this->objectName().toUtf8().data(), input.toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQRegExpValidator_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQRegExpValidator_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRegExpValidator_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQRegExpValidator_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQRegExpValidator_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRegExpValidator_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQRegExpValidator_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRegExpValidator_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRegExpValidator_MetaObject(const_cast<MyQRegExpValidator*>(this), this->objectName().toUtf8().data())); };
};

void QRegExpValidator_SetRegExp(void* ptr, void* rx)
{
	static_cast<QRegExpValidator*>(ptr)->setRegExp(*static_cast<QRegExp*>(rx));
}

void* QRegExpValidator_NewQRegExpValidator(void* parent)
{
	return new MyQRegExpValidator(static_cast<QObject*>(parent));
}

void* QRegExpValidator_NewQRegExpValidator2(void* rx, void* parent)
{
	return new MyQRegExpValidator(*static_cast<QRegExp*>(rx), static_cast<QObject*>(parent));
}

void* QRegExpValidator_RegExp(void* ptr)
{
	return new QRegExp(static_cast<QRegExpValidator*>(ptr)->regExp());
}





void QRegExpValidator_DestroyQRegExpValidator(void* ptr)
{
	static_cast<QRegExpValidator*>(ptr)->~QRegExpValidator();
}





void QRegExpValidator_TimerEvent(void* ptr, void* event)
{
	static_cast<QRegExpValidator*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QRegExpValidator_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QRegExpValidator*>(ptr)->QRegExpValidator::timerEvent(static_cast<QTimerEvent*>(event));
}

void QRegExpValidator_ChildEvent(void* ptr, void* event)
{
	static_cast<QRegExpValidator*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QRegExpValidator_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QRegExpValidator*>(ptr)->QRegExpValidator::childEvent(static_cast<QChildEvent*>(event));
}

void QRegExpValidator_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QRegExpValidator*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRegExpValidator_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRegExpValidator*>(ptr)->QRegExpValidator::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRegExpValidator_CustomEvent(void* ptr, void* event)
{
	static_cast<QRegExpValidator*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QRegExpValidator_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QRegExpValidator*>(ptr)->QRegExpValidator::customEvent(static_cast<QEvent*>(event));
}

void QRegExpValidator_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRegExpValidator*>(ptr), "deleteLater");
}

void QRegExpValidator_DeleteLaterDefault(void* ptr)
{
	static_cast<QRegExpValidator*>(ptr)->QRegExpValidator::deleteLater();
}

void QRegExpValidator_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QRegExpValidator*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRegExpValidator_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRegExpValidator*>(ptr)->QRegExpValidator::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QRegExpValidator_Event(void* ptr, void* e)
{
	return static_cast<QRegExpValidator*>(ptr)->event(static_cast<QEvent*>(e));
}

int QRegExpValidator_EventDefault(void* ptr, void* e)
{
	return static_cast<QRegExpValidator*>(ptr)->QRegExpValidator::event(static_cast<QEvent*>(e));
}

int QRegExpValidator_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QRegExpValidator*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QRegExpValidator_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QRegExpValidator*>(ptr)->QRegExpValidator::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QRegExpValidator_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QRegExpValidator*>(ptr)->metaObject());
}

void* QRegExpValidator_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QRegExpValidator*>(ptr)->QRegExpValidator::metaObject());
}

void* QRegion_NewQRegion()
{
	return new QRegion();
}

void* QRegion_NewQRegion5(void* bm)
{
	return new QRegion(*static_cast<QBitmap*>(bm));
}

void* QRegion_NewQRegion3(void* a, int fillRule)
{
	return new QRegion(*static_cast<QPolygon*>(a), static_cast<Qt::FillRule>(fillRule));
}

void* QRegion_NewQRegion6(void* r, int t)
{
	return new QRegion(*static_cast<QRect*>(r), static_cast<QRegion::RegionType>(t));
}

void* QRegion_NewQRegion4(void* r)
{
	return new QRegion(*static_cast<QRegion*>(r));
}

void* QRegion_BoundingRect(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QRegion*>(ptr)->boundingRect()).x(), static_cast<QRect>(static_cast<QRegion*>(ptr)->boundingRect()).y(), static_cast<QRect>(static_cast<QRegion*>(ptr)->boundingRect()).width(), static_cast<QRect>(static_cast<QRegion*>(ptr)->boundingRect()).height());
}

int QRegion_Contains(void* ptr, void* p)
{
	return static_cast<QRegion*>(ptr)->contains(*static_cast<QPoint*>(p));
}

int QRegion_Contains2(void* ptr, void* r)
{
	return static_cast<QRegion*>(ptr)->contains(*static_cast<QRect*>(r));
}

void* QRegion_Intersected2(void* ptr, void* rect)
{
	return new QRegion(static_cast<QRegion*>(ptr)->intersected(*static_cast<QRect*>(rect)));
}

void* QRegion_Intersected(void* ptr, void* r)
{
	return new QRegion(static_cast<QRegion*>(ptr)->intersected(*static_cast<QRegion*>(r)));
}

int QRegion_Intersects2(void* ptr, void* rect)
{
	return static_cast<QRegion*>(ptr)->intersects(*static_cast<QRect*>(rect));
}

int QRegion_IsEmpty(void* ptr)
{
	return static_cast<QRegion*>(ptr)->isEmpty();
}

int QRegion_IsNull(void* ptr)
{
	return static_cast<QRegion*>(ptr)->isNull();
}

int QRegion_RectCount(void* ptr)
{
	return static_cast<QRegion*>(ptr)->rectCount();
}

void QRegion_SetRects(void* ptr, void* rects, int number)
{
	static_cast<QRegion*>(ptr)->setRects(static_cast<QRect*>(rects), number);
}

void* QRegion_Subtracted(void* ptr, void* r)
{
	return new QRegion(static_cast<QRegion*>(ptr)->subtracted(*static_cast<QRegion*>(r)));
}

void QRegion_Translate(void* ptr, int dx, int dy)
{
	static_cast<QRegion*>(ptr)->translate(dx, dy);
}

void* QRegion_United2(void* ptr, void* rect)
{
	return new QRegion(static_cast<QRegion*>(ptr)->united(*static_cast<QRect*>(rect)));
}

void* QRegion_United(void* ptr, void* r)
{
	return new QRegion(static_cast<QRegion*>(ptr)->united(*static_cast<QRegion*>(r)));
}

void* QRegion_Xored(void* ptr, void* r)
{
	return new QRegion(static_cast<QRegion*>(ptr)->xored(*static_cast<QRegion*>(r)));
}

void* QRegion_NewQRegion2(int x, int y, int w, int h, int t)
{
	return new QRegion(x, y, w, h, static_cast<QRegion::RegionType>(t));
}

int QRegion_Intersects(void* ptr, void* region)
{
	return static_cast<QRegion*>(ptr)->intersects(*static_cast<QRegion*>(region));
}

void QRegion_Swap(void* ptr, void* other)
{
	static_cast<QRegion*>(ptr)->swap(*static_cast<QRegion*>(other));
}

void QRegion_Translate2(void* ptr, void* point)
{
	static_cast<QRegion*>(ptr)->translate(*static_cast<QPoint*>(point));
}

void* QRegion_Translated2(void* ptr, void* p)
{
	return new QRegion(static_cast<QRegion*>(ptr)->translated(*static_cast<QPoint*>(p)));
}

void* QRegion_Translated(void* ptr, int dx, int dy)
{
	return new QRegion(static_cast<QRegion*>(ptr)->translated(dx, dy));
}

class MyQRegularExpressionValidator: public QRegularExpressionValidator
{
public:
	MyQRegularExpressionValidator(QObject *parent) : QRegularExpressionValidator(parent) {};
	MyQRegularExpressionValidator(const QRegularExpression &re, QObject *parent) : QRegularExpressionValidator(re, parent) {};
	void setRegularExpression(const QRegularExpression & re) { callbackQRegularExpressionValidator_SetRegularExpression(this, this->objectName().toUtf8().data(), new QRegularExpression(re)); };
	void Signal_RegularExpressionChanged(const QRegularExpression & re) { callbackQRegularExpressionValidator_RegularExpressionChanged(this, this->objectName().toUtf8().data(), new QRegularExpression(re)); };
	QValidator::State validate(QString & input, int & pos) const { return static_cast<QValidator::State>(callbackQRegularExpressionValidator_Validate(const_cast<MyQRegularExpressionValidator*>(this), this->objectName().toUtf8().data(), input.toUtf8().data(), pos)); };
	void fixup(QString & input) const { callbackQRegularExpressionValidator_Fixup(const_cast<MyQRegularExpressionValidator*>(this), this->objectName().toUtf8().data(), input.toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQRegularExpressionValidator_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQRegularExpressionValidator_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQRegularExpressionValidator_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQRegularExpressionValidator_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQRegularExpressionValidator_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQRegularExpressionValidator_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQRegularExpressionValidator_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQRegularExpressionValidator_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQRegularExpressionValidator_MetaObject(const_cast<MyQRegularExpressionValidator*>(this), this->objectName().toUtf8().data())); };
};

void* QRegularExpressionValidator_RegularExpression(void* ptr)
{
	return new QRegularExpression(static_cast<QRegularExpressionValidator*>(ptr)->regularExpression());
}

void QRegularExpressionValidator_SetRegularExpression(void* ptr, void* re)
{
	QMetaObject::invokeMethod(static_cast<QRegularExpressionValidator*>(ptr), "setRegularExpression", Q_ARG(QRegularExpression, *static_cast<QRegularExpression*>(re)));
}

void* QRegularExpressionValidator_NewQRegularExpressionValidator(void* parent)
{
	return new MyQRegularExpressionValidator(static_cast<QObject*>(parent));
}

void* QRegularExpressionValidator_NewQRegularExpressionValidator2(void* re, void* parent)
{
	return new MyQRegularExpressionValidator(*static_cast<QRegularExpression*>(re), static_cast<QObject*>(parent));
}

void QRegularExpressionValidator_ConnectRegularExpressionChanged(void* ptr)
{
	QObject::connect(static_cast<QRegularExpressionValidator*>(ptr), static_cast<void (QRegularExpressionValidator::*)(const QRegularExpression &)>(&QRegularExpressionValidator::regularExpressionChanged), static_cast<MyQRegularExpressionValidator*>(ptr), static_cast<void (MyQRegularExpressionValidator::*)(const QRegularExpression &)>(&MyQRegularExpressionValidator::Signal_RegularExpressionChanged));
}

void QRegularExpressionValidator_DisconnectRegularExpressionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QRegularExpressionValidator*>(ptr), static_cast<void (QRegularExpressionValidator::*)(const QRegularExpression &)>(&QRegularExpressionValidator::regularExpressionChanged), static_cast<MyQRegularExpressionValidator*>(ptr), static_cast<void (MyQRegularExpressionValidator::*)(const QRegularExpression &)>(&MyQRegularExpressionValidator::Signal_RegularExpressionChanged));
}

void QRegularExpressionValidator_RegularExpressionChanged(void* ptr, void* re)
{
	static_cast<QRegularExpressionValidator*>(ptr)->regularExpressionChanged(*static_cast<QRegularExpression*>(re));
}





void QRegularExpressionValidator_DestroyQRegularExpressionValidator(void* ptr)
{
	static_cast<QRegularExpressionValidator*>(ptr)->~QRegularExpressionValidator();
}





void QRegularExpressionValidator_TimerEvent(void* ptr, void* event)
{
	static_cast<QRegularExpressionValidator*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QRegularExpressionValidator_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QRegularExpressionValidator*>(ptr)->QRegularExpressionValidator::timerEvent(static_cast<QTimerEvent*>(event));
}

void QRegularExpressionValidator_ChildEvent(void* ptr, void* event)
{
	static_cast<QRegularExpressionValidator*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QRegularExpressionValidator_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QRegularExpressionValidator*>(ptr)->QRegularExpressionValidator::childEvent(static_cast<QChildEvent*>(event));
}

void QRegularExpressionValidator_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QRegularExpressionValidator*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRegularExpressionValidator_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRegularExpressionValidator*>(ptr)->QRegularExpressionValidator::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRegularExpressionValidator_CustomEvent(void* ptr, void* event)
{
	static_cast<QRegularExpressionValidator*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QRegularExpressionValidator_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QRegularExpressionValidator*>(ptr)->QRegularExpressionValidator::customEvent(static_cast<QEvent*>(event));
}

void QRegularExpressionValidator_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QRegularExpressionValidator*>(ptr), "deleteLater");
}

void QRegularExpressionValidator_DeleteLaterDefault(void* ptr)
{
	static_cast<QRegularExpressionValidator*>(ptr)->QRegularExpressionValidator::deleteLater();
}

void QRegularExpressionValidator_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QRegularExpressionValidator*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QRegularExpressionValidator_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QRegularExpressionValidator*>(ptr)->QRegularExpressionValidator::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QRegularExpressionValidator_Event(void* ptr, void* e)
{
	return static_cast<QRegularExpressionValidator*>(ptr)->event(static_cast<QEvent*>(e));
}

int QRegularExpressionValidator_EventDefault(void* ptr, void* e)
{
	return static_cast<QRegularExpressionValidator*>(ptr)->QRegularExpressionValidator::event(static_cast<QEvent*>(e));
}

int QRegularExpressionValidator_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QRegularExpressionValidator*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QRegularExpressionValidator_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QRegularExpressionValidator*>(ptr)->QRegularExpressionValidator::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QRegularExpressionValidator_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QRegularExpressionValidator*>(ptr)->metaObject());
}

void* QRegularExpressionValidator_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QRegularExpressionValidator*>(ptr)->QRegularExpressionValidator::metaObject());
}

void* QResizeEvent_NewQResizeEvent(void* size, void* oldSize)
{
	return new QResizeEvent(*static_cast<QSize*>(size), *static_cast<QSize*>(oldSize));
}

void* QResizeEvent_OldSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QResizeEvent*>(ptr)->oldSize()).width(), static_cast<QSize>(static_cast<QResizeEvent*>(ptr)->oldSize()).height());
}

void* QResizeEvent_Size(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QResizeEvent*>(ptr)->size()).width(), static_cast<QSize>(static_cast<QResizeEvent*>(ptr)->size()).height());
}

class MyQScreen: public QScreen
{
public:
	void Signal_AvailableGeometryChanged(const QRect & geometry) { callbackQScreen_AvailableGeometryChanged(this, this->objectName().toUtf8().data(), new QRect(static_cast<QRect>(geometry).x(), static_cast<QRect>(geometry).y(), static_cast<QRect>(geometry).width(), static_cast<QRect>(geometry).height())); };
	void Signal_GeometryChanged(const QRect & geometry) { callbackQScreen_GeometryChanged(this, this->objectName().toUtf8().data(), new QRect(static_cast<QRect>(geometry).x(), static_cast<QRect>(geometry).y(), static_cast<QRect>(geometry).width(), static_cast<QRect>(geometry).height())); };
	void Signal_LogicalDotsPerInchChanged(qreal dpi) { callbackQScreen_LogicalDotsPerInchChanged(this, this->objectName().toUtf8().data(), static_cast<double>(dpi)); };
	void Signal_OrientationChanged(Qt::ScreenOrientation orientation) { callbackQScreen_OrientationChanged(this, this->objectName().toUtf8().data(), orientation); };
	void Signal_PhysicalDotsPerInchChanged(qreal dpi) { callbackQScreen_PhysicalDotsPerInchChanged(this, this->objectName().toUtf8().data(), static_cast<double>(dpi)); };
	void Signal_PhysicalSizeChanged(const QSizeF & size) { callbackQScreen_PhysicalSizeChanged(this, this->objectName().toUtf8().data(), new QSizeF(static_cast<QSizeF>(size).width(), static_cast<QSizeF>(size).height())); };
	void Signal_PrimaryOrientationChanged(Qt::ScreenOrientation orientation) { callbackQScreen_PrimaryOrientationChanged(this, this->objectName().toUtf8().data(), orientation); };
	void Signal_RefreshRateChanged(qreal refreshRate) { callbackQScreen_RefreshRateChanged(this, this->objectName().toUtf8().data(), static_cast<double>(refreshRate)); };
	void Signal_VirtualGeometryChanged(const QRect & rect) { callbackQScreen_VirtualGeometryChanged(this, this->objectName().toUtf8().data(), new QRect(static_cast<QRect>(rect).x(), static_cast<QRect>(rect).y(), static_cast<QRect>(rect).width(), static_cast<QRect>(rect).height())); };
	void timerEvent(QTimerEvent * event) { callbackQScreen_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQScreen_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScreen_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQScreen_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQScreen_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScreen_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQScreen_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScreen_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQScreen_MetaObject(const_cast<MyQScreen*>(this), this->objectName().toUtf8().data())); };
};

void* QScreen_AvailableGeometry(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QScreen*>(ptr)->availableGeometry()).x(), static_cast<QRect>(static_cast<QScreen*>(ptr)->availableGeometry()).y(), static_cast<QRect>(static_cast<QScreen*>(ptr)->availableGeometry()).width(), static_cast<QRect>(static_cast<QScreen*>(ptr)->availableGeometry()).height());
}

void* QScreen_AvailableSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QScreen*>(ptr)->availableSize()).width(), static_cast<QSize>(static_cast<QScreen*>(ptr)->availableSize()).height());
}

void* QScreen_AvailableVirtualGeometry(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QScreen*>(ptr)->availableVirtualGeometry()).x(), static_cast<QRect>(static_cast<QScreen*>(ptr)->availableVirtualGeometry()).y(), static_cast<QRect>(static_cast<QScreen*>(ptr)->availableVirtualGeometry()).width(), static_cast<QRect>(static_cast<QScreen*>(ptr)->availableVirtualGeometry()).height());
}

void* QScreen_AvailableVirtualSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QScreen*>(ptr)->availableVirtualSize()).width(), static_cast<QSize>(static_cast<QScreen*>(ptr)->availableVirtualSize()).height());
}

int QScreen_Depth(void* ptr)
{
	return static_cast<QScreen*>(ptr)->depth();
}

double QScreen_DevicePixelRatio(void* ptr)
{
	return static_cast<double>(static_cast<QScreen*>(ptr)->devicePixelRatio());
}

void* QScreen_Geometry(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QScreen*>(ptr)->geometry()).x(), static_cast<QRect>(static_cast<QScreen*>(ptr)->geometry()).y(), static_cast<QRect>(static_cast<QScreen*>(ptr)->geometry()).width(), static_cast<QRect>(static_cast<QScreen*>(ptr)->geometry()).height());
}

double QScreen_LogicalDotsPerInch(void* ptr)
{
	return static_cast<double>(static_cast<QScreen*>(ptr)->logicalDotsPerInch());
}

double QScreen_LogicalDotsPerInchX(void* ptr)
{
	return static_cast<double>(static_cast<QScreen*>(ptr)->logicalDotsPerInchX());
}

double QScreen_LogicalDotsPerInchY(void* ptr)
{
	return static_cast<double>(static_cast<QScreen*>(ptr)->logicalDotsPerInchY());
}

char* QScreen_Name(void* ptr)
{
	return static_cast<QScreen*>(ptr)->name().toUtf8().data();
}

int QScreen_NativeOrientation(void* ptr)
{
	return static_cast<QScreen*>(ptr)->nativeOrientation();
}

int QScreen_Orientation(void* ptr)
{
	return static_cast<QScreen*>(ptr)->orientation();
}

double QScreen_PhysicalDotsPerInch(void* ptr)
{
	return static_cast<double>(static_cast<QScreen*>(ptr)->physicalDotsPerInch());
}

double QScreen_PhysicalDotsPerInchX(void* ptr)
{
	return static_cast<double>(static_cast<QScreen*>(ptr)->physicalDotsPerInchX());
}

double QScreen_PhysicalDotsPerInchY(void* ptr)
{
	return static_cast<double>(static_cast<QScreen*>(ptr)->physicalDotsPerInchY());
}

void* QScreen_PhysicalSize(void* ptr)
{
	return new QSizeF(static_cast<QSizeF>(static_cast<QScreen*>(ptr)->physicalSize()).width(), static_cast<QSizeF>(static_cast<QScreen*>(ptr)->physicalSize()).height());
}

int QScreen_PrimaryOrientation(void* ptr)
{
	return static_cast<QScreen*>(ptr)->primaryOrientation();
}

double QScreen_RefreshRate(void* ptr)
{
	return static_cast<double>(static_cast<QScreen*>(ptr)->refreshRate());
}

void* QScreen_Size(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QScreen*>(ptr)->size()).width(), static_cast<QSize>(static_cast<QScreen*>(ptr)->size()).height());
}

void* QScreen_VirtualGeometry(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QScreen*>(ptr)->virtualGeometry()).x(), static_cast<QRect>(static_cast<QScreen*>(ptr)->virtualGeometry()).y(), static_cast<QRect>(static_cast<QScreen*>(ptr)->virtualGeometry()).width(), static_cast<QRect>(static_cast<QScreen*>(ptr)->virtualGeometry()).height());
}

void* QScreen_VirtualSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QScreen*>(ptr)->virtualSize()).width(), static_cast<QSize>(static_cast<QScreen*>(ptr)->virtualSize()).height());
}

int QScreen_AngleBetween(void* ptr, int a, int b)
{
	return static_cast<QScreen*>(ptr)->angleBetween(static_cast<Qt::ScreenOrientation>(a), static_cast<Qt::ScreenOrientation>(b));
}

void QScreen_ConnectAvailableGeometryChanged(void* ptr)
{
	QObject::connect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(const QRect &)>(&QScreen::availableGeometryChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(const QRect &)>(&MyQScreen::Signal_AvailableGeometryChanged));
}

void QScreen_DisconnectAvailableGeometryChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(const QRect &)>(&QScreen::availableGeometryChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(const QRect &)>(&MyQScreen::Signal_AvailableGeometryChanged));
}

void QScreen_AvailableGeometryChanged(void* ptr, void* geometry)
{
	static_cast<QScreen*>(ptr)->availableGeometryChanged(*static_cast<QRect*>(geometry));
}

void QScreen_ConnectGeometryChanged(void* ptr)
{
	QObject::connect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(const QRect &)>(&QScreen::geometryChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(const QRect &)>(&MyQScreen::Signal_GeometryChanged));
}

void QScreen_DisconnectGeometryChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(const QRect &)>(&QScreen::geometryChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(const QRect &)>(&MyQScreen::Signal_GeometryChanged));
}

void QScreen_GeometryChanged(void* ptr, void* geometry)
{
	static_cast<QScreen*>(ptr)->geometryChanged(*static_cast<QRect*>(geometry));
}

void* QScreen_GrabWindow(void* ptr, unsigned long long window, int x, int y, int width, int height)
{
	return new QPixmap(static_cast<QScreen*>(ptr)->grabWindow(static_cast<unsigned long long>(window), x, y, width, height));
}

int QScreen_IsLandscape(void* ptr, int o)
{
	return static_cast<QScreen*>(ptr)->isLandscape(static_cast<Qt::ScreenOrientation>(o));
}

int QScreen_IsPortrait(void* ptr, int o)
{
	return static_cast<QScreen*>(ptr)->isPortrait(static_cast<Qt::ScreenOrientation>(o));
}

void QScreen_ConnectLogicalDotsPerInchChanged(void* ptr)
{
	QObject::connect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(qreal)>(&QScreen::logicalDotsPerInchChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(qreal)>(&MyQScreen::Signal_LogicalDotsPerInchChanged));
}

void QScreen_DisconnectLogicalDotsPerInchChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(qreal)>(&QScreen::logicalDotsPerInchChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(qreal)>(&MyQScreen::Signal_LogicalDotsPerInchChanged));
}

void QScreen_LogicalDotsPerInchChanged(void* ptr, double dpi)
{
	static_cast<QScreen*>(ptr)->logicalDotsPerInchChanged(static_cast<double>(dpi));
}

void* QScreen_MapBetween(void* ptr, int a, int b, void* rect)
{
	return new QRect(static_cast<QRect>(static_cast<QScreen*>(ptr)->mapBetween(static_cast<Qt::ScreenOrientation>(a), static_cast<Qt::ScreenOrientation>(b), *static_cast<QRect*>(rect))).x(), static_cast<QRect>(static_cast<QScreen*>(ptr)->mapBetween(static_cast<Qt::ScreenOrientation>(a), static_cast<Qt::ScreenOrientation>(b), *static_cast<QRect*>(rect))).y(), static_cast<QRect>(static_cast<QScreen*>(ptr)->mapBetween(static_cast<Qt::ScreenOrientation>(a), static_cast<Qt::ScreenOrientation>(b), *static_cast<QRect*>(rect))).width(), static_cast<QRect>(static_cast<QScreen*>(ptr)->mapBetween(static_cast<Qt::ScreenOrientation>(a), static_cast<Qt::ScreenOrientation>(b), *static_cast<QRect*>(rect))).height());
}

void QScreen_ConnectOrientationChanged(void* ptr)
{
	QObject::connect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(Qt::ScreenOrientation)>(&QScreen::orientationChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(Qt::ScreenOrientation)>(&MyQScreen::Signal_OrientationChanged));
}

void QScreen_DisconnectOrientationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(Qt::ScreenOrientation)>(&QScreen::orientationChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(Qt::ScreenOrientation)>(&MyQScreen::Signal_OrientationChanged));
}

void QScreen_OrientationChanged(void* ptr, int orientation)
{
	static_cast<QScreen*>(ptr)->orientationChanged(static_cast<Qt::ScreenOrientation>(orientation));
}

int QScreen_OrientationUpdateMask(void* ptr)
{
	return static_cast<QScreen*>(ptr)->orientationUpdateMask();
}

void QScreen_ConnectPhysicalDotsPerInchChanged(void* ptr)
{
	QObject::connect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(qreal)>(&QScreen::physicalDotsPerInchChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(qreal)>(&MyQScreen::Signal_PhysicalDotsPerInchChanged));
}

void QScreen_DisconnectPhysicalDotsPerInchChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(qreal)>(&QScreen::physicalDotsPerInchChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(qreal)>(&MyQScreen::Signal_PhysicalDotsPerInchChanged));
}

void QScreen_PhysicalDotsPerInchChanged(void* ptr, double dpi)
{
	static_cast<QScreen*>(ptr)->physicalDotsPerInchChanged(static_cast<double>(dpi));
}

void QScreen_ConnectPhysicalSizeChanged(void* ptr)
{
	QObject::connect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(const QSizeF &)>(&QScreen::physicalSizeChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(const QSizeF &)>(&MyQScreen::Signal_PhysicalSizeChanged));
}

void QScreen_DisconnectPhysicalSizeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(const QSizeF &)>(&QScreen::physicalSizeChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(const QSizeF &)>(&MyQScreen::Signal_PhysicalSizeChanged));
}

void QScreen_PhysicalSizeChanged(void* ptr, void* size)
{
	static_cast<QScreen*>(ptr)->physicalSizeChanged(*static_cast<QSizeF*>(size));
}

void QScreen_ConnectPrimaryOrientationChanged(void* ptr)
{
	QObject::connect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(Qt::ScreenOrientation)>(&QScreen::primaryOrientationChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(Qt::ScreenOrientation)>(&MyQScreen::Signal_PrimaryOrientationChanged));
}

void QScreen_DisconnectPrimaryOrientationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(Qt::ScreenOrientation)>(&QScreen::primaryOrientationChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(Qt::ScreenOrientation)>(&MyQScreen::Signal_PrimaryOrientationChanged));
}

void QScreen_PrimaryOrientationChanged(void* ptr, int orientation)
{
	static_cast<QScreen*>(ptr)->primaryOrientationChanged(static_cast<Qt::ScreenOrientation>(orientation));
}

void QScreen_ConnectRefreshRateChanged(void* ptr)
{
	QObject::connect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(qreal)>(&QScreen::refreshRateChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(qreal)>(&MyQScreen::Signal_RefreshRateChanged));
}

void QScreen_DisconnectRefreshRateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(qreal)>(&QScreen::refreshRateChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(qreal)>(&MyQScreen::Signal_RefreshRateChanged));
}

void QScreen_RefreshRateChanged(void* ptr, double refreshRate)
{
	static_cast<QScreen*>(ptr)->refreshRateChanged(static_cast<double>(refreshRate));
}

void QScreen_SetOrientationUpdateMask(void* ptr, int mask)
{
	static_cast<QScreen*>(ptr)->setOrientationUpdateMask(static_cast<Qt::ScreenOrientation>(mask));
}

void QScreen_ConnectVirtualGeometryChanged(void* ptr)
{
	QObject::connect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(const QRect &)>(&QScreen::virtualGeometryChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(const QRect &)>(&MyQScreen::Signal_VirtualGeometryChanged));
}

void QScreen_DisconnectVirtualGeometryChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(const QRect &)>(&QScreen::virtualGeometryChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(const QRect &)>(&MyQScreen::Signal_VirtualGeometryChanged));
}

void QScreen_VirtualGeometryChanged(void* ptr, void* rect)
{
	static_cast<QScreen*>(ptr)->virtualGeometryChanged(*static_cast<QRect*>(rect));
}

void QScreen_DestroyQScreen(void* ptr)
{
	static_cast<QScreen*>(ptr)->~QScreen();
}

void QScreen_TimerEvent(void* ptr, void* event)
{
	static_cast<QScreen*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QScreen_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QScreen*>(ptr)->QScreen::timerEvent(static_cast<QTimerEvent*>(event));
}

void QScreen_ChildEvent(void* ptr, void* event)
{
	static_cast<QScreen*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QScreen_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QScreen*>(ptr)->QScreen::childEvent(static_cast<QChildEvent*>(event));
}

void QScreen_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QScreen*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScreen_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScreen*>(ptr)->QScreen::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScreen_CustomEvent(void* ptr, void* event)
{
	static_cast<QScreen*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QScreen_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QScreen*>(ptr)->QScreen::customEvent(static_cast<QEvent*>(event));
}

void QScreen_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QScreen*>(ptr), "deleteLater");
}

void QScreen_DeleteLaterDefault(void* ptr)
{
	static_cast<QScreen*>(ptr)->QScreen::deleteLater();
}

void QScreen_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QScreen*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScreen_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QScreen*>(ptr)->QScreen::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QScreen_Event(void* ptr, void* e)
{
	return static_cast<QScreen*>(ptr)->event(static_cast<QEvent*>(e));
}

int QScreen_EventDefault(void* ptr, void* e)
{
	return static_cast<QScreen*>(ptr)->QScreen::event(static_cast<QEvent*>(e));
}

int QScreen_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QScreen*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QScreen_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QScreen*>(ptr)->QScreen::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QScreen_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScreen*>(ptr)->metaObject());
}

void* QScreen_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QScreen*>(ptr)->QScreen::metaObject());
}

void* QScrollEvent_NewQScrollEvent(void* contentPos, void* overshootDistance, int scrollState)
{
	return new QScrollEvent(*static_cast<QPointF*>(contentPos), *static_cast<QPointF*>(overshootDistance), static_cast<QScrollEvent::ScrollState>(scrollState));
}

void* QScrollEvent_ContentPos(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QScrollEvent*>(ptr)->contentPos()).x(), static_cast<QPointF>(static_cast<QScrollEvent*>(ptr)->contentPos()).y());
}

void* QScrollEvent_OvershootDistance(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QScrollEvent*>(ptr)->overshootDistance()).x(), static_cast<QPointF>(static_cast<QScrollEvent*>(ptr)->overshootDistance()).y());
}

int QScrollEvent_ScrollState(void* ptr)
{
	return static_cast<QScrollEvent*>(ptr)->scrollState();
}

void QScrollEvent_DestroyQScrollEvent(void* ptr)
{
	static_cast<QScrollEvent*>(ptr)->~QScrollEvent();
}

void* QScrollPrepareEvent_NewQScrollPrepareEvent(void* startPos)
{
	return new QScrollPrepareEvent(*static_cast<QPointF*>(startPos));
}

void* QScrollPrepareEvent_ContentPos(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QScrollPrepareEvent*>(ptr)->contentPos()).x(), static_cast<QPointF>(static_cast<QScrollPrepareEvent*>(ptr)->contentPos()).y());
}

void* QScrollPrepareEvent_ContentPosRange(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QScrollPrepareEvent*>(ptr)->contentPosRange()).x(), static_cast<QRectF>(static_cast<QScrollPrepareEvent*>(ptr)->contentPosRange()).y(), static_cast<QRectF>(static_cast<QScrollPrepareEvent*>(ptr)->contentPosRange()).width(), static_cast<QRectF>(static_cast<QScrollPrepareEvent*>(ptr)->contentPosRange()).height());
}

void QScrollPrepareEvent_SetContentPos(void* ptr, void* pos)
{
	static_cast<QScrollPrepareEvent*>(ptr)->setContentPos(*static_cast<QPointF*>(pos));
}

void QScrollPrepareEvent_SetContentPosRange(void* ptr, void* rect)
{
	static_cast<QScrollPrepareEvent*>(ptr)->setContentPosRange(*static_cast<QRectF*>(rect));
}

void QScrollPrepareEvent_SetViewportSize(void* ptr, void* size)
{
	static_cast<QScrollPrepareEvent*>(ptr)->setViewportSize(*static_cast<QSizeF*>(size));
}

void* QScrollPrepareEvent_StartPos(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QScrollPrepareEvent*>(ptr)->startPos()).x(), static_cast<QPointF>(static_cast<QScrollPrepareEvent*>(ptr)->startPos()).y());
}

void* QScrollPrepareEvent_ViewportSize(void* ptr)
{
	return new QSizeF(static_cast<QSizeF>(static_cast<QScrollPrepareEvent*>(ptr)->viewportSize()).width(), static_cast<QSizeF>(static_cast<QScrollPrepareEvent*>(ptr)->viewportSize()).height());
}

void QScrollPrepareEvent_DestroyQScrollPrepareEvent(void* ptr)
{
	static_cast<QScrollPrepareEvent*>(ptr)->~QScrollPrepareEvent();
}

int QSessionManager_RestartHint(void* ptr)
{
	return static_cast<QSessionManager*>(ptr)->restartHint();
}

char* QSessionManager_SessionKey(void* ptr)
{
	return static_cast<QSessionManager*>(ptr)->sessionKey().toUtf8().data();
}

int QSessionManager_AllowsErrorInteraction(void* ptr)
{
	return static_cast<QSessionManager*>(ptr)->allowsErrorInteraction();
}

int QSessionManager_AllowsInteraction(void* ptr)
{
	return static_cast<QSessionManager*>(ptr)->allowsInteraction();
}

void QSessionManager_Cancel(void* ptr)
{
	static_cast<QSessionManager*>(ptr)->cancel();
}

char* QSessionManager_DiscardCommand(void* ptr)
{
	return static_cast<QSessionManager*>(ptr)->discardCommand().join("|").toUtf8().data();
}

int QSessionManager_IsPhase2(void* ptr)
{
	return static_cast<QSessionManager*>(ptr)->isPhase2();
}

void QSessionManager_Release(void* ptr)
{
	static_cast<QSessionManager*>(ptr)->release();
}

void QSessionManager_RequestPhase2(void* ptr)
{
	static_cast<QSessionManager*>(ptr)->requestPhase2();
}

char* QSessionManager_RestartCommand(void* ptr)
{
	return static_cast<QSessionManager*>(ptr)->restartCommand().join("|").toUtf8().data();
}

char* QSessionManager_SessionId(void* ptr)
{
	return static_cast<QSessionManager*>(ptr)->sessionId().toUtf8().data();
}

void QSessionManager_SetDiscardCommand(void* ptr, char* command)
{
	static_cast<QSessionManager*>(ptr)->setDiscardCommand(QString(command).split("|", QString::SkipEmptyParts));
}

void QSessionManager_SetManagerProperty2(void* ptr, char* name, char* value)
{
	static_cast<QSessionManager*>(ptr)->setManagerProperty(QString(name), QString(value));
}

void QSessionManager_SetManagerProperty(void* ptr, char* name, char* value)
{
	static_cast<QSessionManager*>(ptr)->setManagerProperty(QString(name), QString(value).split("|", QString::SkipEmptyParts));
}

void QSessionManager_SetRestartCommand(void* ptr, char* command)
{
	static_cast<QSessionManager*>(ptr)->setRestartCommand(QString(command).split("|", QString::SkipEmptyParts));
}

void QSessionManager_SetRestartHint(void* ptr, int hint)
{
	static_cast<QSessionManager*>(ptr)->setRestartHint(static_cast<QSessionManager::RestartHint>(hint));
}

void QSessionManager_TimerEvent(void* ptr, void* event)
{
	static_cast<QSessionManager*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSessionManager_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSessionManager*>(ptr)->QSessionManager::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSessionManager_ChildEvent(void* ptr, void* event)
{
	static_cast<QSessionManager*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSessionManager_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSessionManager*>(ptr)->QSessionManager::childEvent(static_cast<QChildEvent*>(event));
}

void QSessionManager_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSessionManager*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSessionManager_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSessionManager*>(ptr)->QSessionManager::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSessionManager_CustomEvent(void* ptr, void* event)
{
	static_cast<QSessionManager*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSessionManager_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSessionManager*>(ptr)->QSessionManager::customEvent(static_cast<QEvent*>(event));
}

void QSessionManager_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSessionManager*>(ptr), "deleteLater");
}

void QSessionManager_DeleteLaterDefault(void* ptr)
{
	static_cast<QSessionManager*>(ptr)->QSessionManager::deleteLater();
}

void QSessionManager_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSessionManager*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSessionManager_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSessionManager*>(ptr)->QSessionManager::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QSessionManager_Event(void* ptr, void* e)
{
	return static_cast<QSessionManager*>(ptr)->event(static_cast<QEvent*>(e));
}

int QSessionManager_EventDefault(void* ptr, void* e)
{
	return static_cast<QSessionManager*>(ptr)->QSessionManager::event(static_cast<QEvent*>(e));
}

int QSessionManager_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSessionManager*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QSessionManager_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSessionManager*>(ptr)->QSessionManager::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSessionManager_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSessionManager*>(ptr)->metaObject());
}

void* QSessionManager_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSessionManager*>(ptr)->QSessionManager::metaObject());
}

void* QShortcutEvent_NewQShortcutEvent(void* key, int id, int ambiguous)
{
	return new QShortcutEvent(*static_cast<QKeySequence*>(key), id, ambiguous != 0);
}

int QShortcutEvent_IsAmbiguous(void* ptr)
{
	return static_cast<QShortcutEvent*>(ptr)->isAmbiguous();
}

void* QShortcutEvent_Key(void* ptr)
{
	return new QKeySequence(static_cast<QShortcutEvent*>(ptr)->key());
}

int QShortcutEvent_ShortcutId(void* ptr)
{
	return static_cast<QShortcutEvent*>(ptr)->shortcutId();
}

void QShortcutEvent_DestroyQShortcutEvent(void* ptr)
{
	static_cast<QShortcutEvent*>(ptr)->~QShortcutEvent();
}

void* QShowEvent_NewQShowEvent()
{
	return new QShowEvent();
}

class MyQStandardItem: public QStandardItem
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	MyQStandardItem() : QStandardItem() {};
	MyQStandardItem(const QIcon &icon, const QString &text) : QStandardItem(icon, text) {};
	MyQStandardItem(const QString &text) : QStandardItem(text) {};
	MyQStandardItem(int rows, int columns) : QStandardItem(rows, columns) {};
	MyQStandardItem(const QStandardItem &other) : QStandardItem(other) {};
	QStandardItem * clone() const { return static_cast<QStandardItem*>(callbackQStandardItem_Clone(const_cast<MyQStandardItem*>(this), this->objectNameAbs().toUtf8().data())); };
	QVariant data(int role) const { return *static_cast<QVariant*>(callbackQStandardItem_Data(const_cast<MyQStandardItem*>(this), this->objectNameAbs().toUtf8().data(), role)); };
	void setData(const QVariant & value, int role) { callbackQStandardItem_SetData(this, this->objectNameAbs().toUtf8().data(), new QVariant(value), role); };
	int type() const { return callbackQStandardItem_Type(const_cast<MyQStandardItem*>(this), this->objectNameAbs().toUtf8().data()); };
};

void* QStandardItem_NewQStandardItem()
{
	return new MyQStandardItem();
}

void* QStandardItem_NewQStandardItem3(void* icon, char* text)
{
	return new MyQStandardItem(*static_cast<QIcon*>(icon), QString(text));
}

void* QStandardItem_NewQStandardItem2(char* text)
{
	return new MyQStandardItem(QString(text));
}

void* QStandardItem_NewQStandardItem4(int rows, int columns)
{
	return new MyQStandardItem(rows, columns);
}

void* QStandardItem_NewQStandardItem5(void* other)
{
	return new MyQStandardItem(*static_cast<QStandardItem*>(other));
}

char* QStandardItem_AccessibleDescription(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->accessibleDescription().toUtf8().data();
}

char* QStandardItem_AccessibleText(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->accessibleText().toUtf8().data();
}

void QStandardItem_AppendRow2(void* ptr, void* item)
{
	static_cast<QStandardItem*>(ptr)->appendRow(static_cast<QStandardItem*>(item));
}

void* QStandardItem_Background(void* ptr)
{
	return new QBrush(static_cast<QStandardItem*>(ptr)->background());
}

int QStandardItem_CheckState(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->checkState();
}

void* QStandardItem_Child(void* ptr, int row, int column)
{
	return static_cast<QStandardItem*>(ptr)->child(row, column);
}

void* QStandardItem_Clone(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->clone();
}

void* QStandardItem_CloneDefault(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->QStandardItem::clone();
}

int QStandardItem_Column(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->column();
}

int QStandardItem_ColumnCount(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->columnCount();
}

void* QStandardItem_Data(void* ptr, int role)
{
	return new QVariant(static_cast<QStandardItem*>(ptr)->data(role));
}

void* QStandardItem_DataDefault(void* ptr, int role)
{
	return new QVariant(static_cast<QStandardItem*>(ptr)->QStandardItem::data(role));
}

void QStandardItem_EmitDataChanged(void* ptr)
{
	static_cast<QStandardItem*>(ptr)->emitDataChanged();
}

int QStandardItem_Flags(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->flags();
}

void* QStandardItem_Font(void* ptr)
{
	return new QFont(static_cast<QStandardItem*>(ptr)->font());
}

void* QStandardItem_Foreground(void* ptr)
{
	return new QBrush(static_cast<QStandardItem*>(ptr)->foreground());
}

int QStandardItem_HasChildren(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->hasChildren();
}

void* QStandardItem_Icon(void* ptr)
{
	return new QIcon(static_cast<QStandardItem*>(ptr)->icon());
}

void* QStandardItem_Index(void* ptr)
{
	return new QModelIndex(static_cast<QStandardItem*>(ptr)->index());
}

void QStandardItem_InsertColumns(void* ptr, int column, int count)
{
	static_cast<QStandardItem*>(ptr)->insertColumns(column, count);
}

void QStandardItem_InsertRow2(void* ptr, int row, void* item)
{
	static_cast<QStandardItem*>(ptr)->insertRow(row, static_cast<QStandardItem*>(item));
}

void QStandardItem_InsertRows2(void* ptr, int row, int count)
{
	static_cast<QStandardItem*>(ptr)->insertRows(row, count);
}

int QStandardItem_IsCheckable(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->isCheckable();
}

int QStandardItem_IsDragEnabled(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->isDragEnabled();
}

int QStandardItem_IsDropEnabled(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->isDropEnabled();
}

int QStandardItem_IsEditable(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->isEditable();
}

int QStandardItem_IsEnabled(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->isEnabled();
}

int QStandardItem_IsSelectable(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->isSelectable();
}

int QStandardItem_IsTristate(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->isTristate();
}

void* QStandardItem_Model(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->model();
}

void* QStandardItem_Parent(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->parent();
}

void QStandardItem_RemoveColumn(void* ptr, int column)
{
	static_cast<QStandardItem*>(ptr)->removeColumn(column);
}

void QStandardItem_RemoveColumns(void* ptr, int column, int count)
{
	static_cast<QStandardItem*>(ptr)->removeColumns(column, count);
}

void QStandardItem_RemoveRow(void* ptr, int row)
{
	static_cast<QStandardItem*>(ptr)->removeRow(row);
}

void QStandardItem_RemoveRows(void* ptr, int row, int count)
{
	static_cast<QStandardItem*>(ptr)->removeRows(row, count);
}

int QStandardItem_Row(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->row();
}

int QStandardItem_RowCount(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->rowCount();
}

void QStandardItem_SetAccessibleDescription(void* ptr, char* accessibleDescription)
{
	static_cast<QStandardItem*>(ptr)->setAccessibleDescription(QString(accessibleDescription));
}

void QStandardItem_SetAccessibleText(void* ptr, char* accessibleText)
{
	static_cast<QStandardItem*>(ptr)->setAccessibleText(QString(accessibleText));
}

void QStandardItem_SetBackground(void* ptr, void* brush)
{
	static_cast<QStandardItem*>(ptr)->setBackground(*static_cast<QBrush*>(brush));
}

void QStandardItem_SetCheckState(void* ptr, int state)
{
	static_cast<QStandardItem*>(ptr)->setCheckState(static_cast<Qt::CheckState>(state));
}

void QStandardItem_SetCheckable(void* ptr, int checkable)
{
	static_cast<QStandardItem*>(ptr)->setCheckable(checkable != 0);
}

void QStandardItem_SetChild2(void* ptr, int row, void* item)
{
	static_cast<QStandardItem*>(ptr)->setChild(row, static_cast<QStandardItem*>(item));
}

void QStandardItem_SetChild(void* ptr, int row, int column, void* item)
{
	static_cast<QStandardItem*>(ptr)->setChild(row, column, static_cast<QStandardItem*>(item));
}

void QStandardItem_SetColumnCount(void* ptr, int columns)
{
	static_cast<QStandardItem*>(ptr)->setColumnCount(columns);
}

void QStandardItem_SetData(void* ptr, void* value, int role)
{
	static_cast<QStandardItem*>(ptr)->setData(*static_cast<QVariant*>(value), role);
}

void QStandardItem_SetDataDefault(void* ptr, void* value, int role)
{
	static_cast<QStandardItem*>(ptr)->QStandardItem::setData(*static_cast<QVariant*>(value), role);
}

void QStandardItem_SetDragEnabled(void* ptr, int dragEnabled)
{
	static_cast<QStandardItem*>(ptr)->setDragEnabled(dragEnabled != 0);
}

void QStandardItem_SetDropEnabled(void* ptr, int dropEnabled)
{
	static_cast<QStandardItem*>(ptr)->setDropEnabled(dropEnabled != 0);
}

void QStandardItem_SetEditable(void* ptr, int editable)
{
	static_cast<QStandardItem*>(ptr)->setEditable(editable != 0);
}

void QStandardItem_SetEnabled(void* ptr, int enabled)
{
	static_cast<QStandardItem*>(ptr)->setEnabled(enabled != 0);
}

void QStandardItem_SetFlags(void* ptr, int flags)
{
	static_cast<QStandardItem*>(ptr)->setFlags(static_cast<Qt::ItemFlag>(flags));
}

void QStandardItem_SetFont(void* ptr, void* font)
{
	static_cast<QStandardItem*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QStandardItem_SetForeground(void* ptr, void* brush)
{
	static_cast<QStandardItem*>(ptr)->setForeground(*static_cast<QBrush*>(brush));
}

void QStandardItem_SetIcon(void* ptr, void* icon)
{
	static_cast<QStandardItem*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QStandardItem_SetRowCount(void* ptr, int rows)
{
	static_cast<QStandardItem*>(ptr)->setRowCount(rows);
}

void QStandardItem_SetSelectable(void* ptr, int selectable)
{
	static_cast<QStandardItem*>(ptr)->setSelectable(selectable != 0);
}

void QStandardItem_SetSizeHint(void* ptr, void* size)
{
	static_cast<QStandardItem*>(ptr)->setSizeHint(*static_cast<QSize*>(size));
}

void QStandardItem_SetStatusTip(void* ptr, char* statusTip)
{
	static_cast<QStandardItem*>(ptr)->setStatusTip(QString(statusTip));
}

void QStandardItem_SetText(void* ptr, char* text)
{
	static_cast<QStandardItem*>(ptr)->setText(QString(text));
}

void QStandardItem_SetTextAlignment(void* ptr, int alignment)
{
	static_cast<QStandardItem*>(ptr)->setTextAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QStandardItem_SetToolTip(void* ptr, char* toolTip)
{
	static_cast<QStandardItem*>(ptr)->setToolTip(QString(toolTip));
}

void QStandardItem_SetTristate(void* ptr, int tristate)
{
	static_cast<QStandardItem*>(ptr)->setTristate(tristate != 0);
}

void QStandardItem_SetWhatsThis(void* ptr, char* whatsThis)
{
	static_cast<QStandardItem*>(ptr)->setWhatsThis(QString(whatsThis));
}

void* QStandardItem_SizeHint(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QStandardItem*>(ptr)->sizeHint()).width(), static_cast<QSize>(static_cast<QStandardItem*>(ptr)->sizeHint()).height());
}

void QStandardItem_SortChildren(void* ptr, int column, int order)
{
	static_cast<QStandardItem*>(ptr)->sortChildren(column, static_cast<Qt::SortOrder>(order));
}

char* QStandardItem_StatusTip(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->statusTip().toUtf8().data();
}

void* QStandardItem_TakeChild(void* ptr, int row, int column)
{
	return static_cast<QStandardItem*>(ptr)->takeChild(row, column);
}

char* QStandardItem_Text(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->text().toUtf8().data();
}

int QStandardItem_TextAlignment(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->textAlignment();
}

char* QStandardItem_ToolTip(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->toolTip().toUtf8().data();
}

int QStandardItem_Type(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->type();
}

int QStandardItem_TypeDefault(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->QStandardItem::type();
}

char* QStandardItem_WhatsThis(void* ptr)
{
	return static_cast<QStandardItem*>(ptr)->whatsThis().toUtf8().data();
}

void QStandardItem_DestroyQStandardItem(void* ptr)
{
	static_cast<QStandardItem*>(ptr)->~QStandardItem();
}

char* QStandardItem_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQStandardItem*>(static_cast<QStandardItem*>(ptr))) {
		return static_cast<MyQStandardItem*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QStandardItem_BASE").toUtf8().data();
}

void QStandardItem_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQStandardItem*>(static_cast<QStandardItem*>(ptr))) {
		static_cast<MyQStandardItem*>(ptr)->setObjectNameAbs(QString(name));
	}
}

class MyQStandardItemModel: public QStandardItemModel
{
public:
	MyQStandardItemModel(QObject *parent) : QStandardItemModel(parent) {};
	MyQStandardItemModel(int rows, int columns, QObject *parent) : QStandardItemModel(rows, columns, parent) {};
	int columnCount(const QModelIndex & parent) const { return callbackQStandardItemModel_ColumnCount(const_cast<MyQStandardItemModel*>(this), this->objectName().toUtf8().data(), new QModelIndex(parent)); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackQStandardItemModel_Data(const_cast<MyQStandardItemModel*>(this), this->objectName().toUtf8().data(), new QModelIndex(index), role)); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackQStandardItemModel_DropMimeData(this, this->objectName().toUtf8().data(), const_cast<QMimeData*>(data), action, row, column, new QModelIndex(parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackQStandardItemModel_Flags(const_cast<MyQStandardItemModel*>(this), this->objectName().toUtf8().data(), new QModelIndex(index))); };
	bool hasChildren(const QModelIndex & parent) const { return callbackQStandardItemModel_HasChildren(const_cast<MyQStandardItemModel*>(this), this->objectName().toUtf8().data(), new QModelIndex(parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackQStandardItemModel_HeaderData(const_cast<MyQStandardItemModel*>(this), this->objectName().toUtf8().data(), section, orientation, role)); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackQStandardItemModel_Index(const_cast<MyQStandardItemModel*>(this), this->objectName().toUtf8().data(), row, column, new QModelIndex(parent))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackQStandardItemModel_InsertColumns(this, this->objectName().toUtf8().data(), column, count, new QModelIndex(parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackQStandardItemModel_InsertRows(this, this->objectName().toUtf8().data(), row, count, new QModelIndex(parent)) != 0; };
	void Signal_ItemChanged(QStandardItem * item) { callbackQStandardItemModel_ItemChanged(this, this->objectName().toUtf8().data(), item); };
	QStringList mimeTypes() const { return QString(callbackQStandardItemModel_MimeTypes(const_cast<MyQStandardItemModel*>(this), this->objectName().toUtf8().data())).split("|", QString::SkipEmptyParts); };
	QModelIndex parent(const QModelIndex & child) const { return *static_cast<QModelIndex*>(callbackQStandardItemModel_Parent(const_cast<MyQStandardItemModel*>(this), this->objectName().toUtf8().data(), new QModelIndex(child))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackQStandardItemModel_RemoveColumns(this, this->objectName().toUtf8().data(), column, count, new QModelIndex(parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackQStandardItemModel_RemoveRows(this, this->objectName().toUtf8().data(), row, count, new QModelIndex(parent)) != 0; };
	int rowCount(const QModelIndex & parent) const { return callbackQStandardItemModel_RowCount(const_cast<MyQStandardItemModel*>(this), this->objectName().toUtf8().data(), new QModelIndex(parent)); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackQStandardItemModel_SetData(this, this->objectName().toUtf8().data(), new QModelIndex(index), new QVariant(value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackQStandardItemModel_SetHeaderData(this, this->objectName().toUtf8().data(), section, orientation, new QVariant(value), role) != 0; };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackQStandardItemModel_Sibling(const_cast<MyQStandardItemModel*>(this), this->objectName().toUtf8().data(), row, column, new QModelIndex(idx))); };
	void sort(int column, Qt::SortOrder order) { callbackQStandardItemModel_Sort(this, this->objectName().toUtf8().data(), column, order); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackQStandardItemModel_SupportedDropActions(const_cast<MyQStandardItemModel*>(this), this->objectName().toUtf8().data())); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQStandardItemModel_Buddy(const_cast<MyQStandardItemModel*>(this), this->objectName().toUtf8().data(), new QModelIndex(index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackQStandardItemModel_CanDropMimeData(const_cast<MyQStandardItemModel*>(this), this->objectName().toUtf8().data(), const_cast<QMimeData*>(data), action, row, column, new QModelIndex(parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackQStandardItemModel_CanFetchMore(const_cast<MyQStandardItemModel*>(this), this->objectName().toUtf8().data(), new QModelIndex(parent)) != 0; };
	void fetchMore(const QModelIndex & parent) { callbackQStandardItemModel_FetchMore(this, this->objectName().toUtf8().data(), new QModelIndex(parent)); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQStandardItemModel_MoveColumns(this, this->objectName().toUtf8().data(), new QModelIndex(sourceParent), sourceColumn, count, new QModelIndex(destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQStandardItemModel_MoveRows(this, this->objectName().toUtf8().data(), new QModelIndex(sourceParent), sourceRow, count, new QModelIndex(destinationParent), destinationChild) != 0; };
	void resetInternalData() { callbackQStandardItemModel_ResetInternalData(this, this->objectName().toUtf8().data()); };
	void revert() { callbackQStandardItemModel_Revert(this, this->objectName().toUtf8().data()); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackQStandardItemModel_Span(const_cast<MyQStandardItemModel*>(this), this->objectName().toUtf8().data(), new QModelIndex(index))); };
	bool submit() { return callbackQStandardItemModel_Submit(this, this->objectName().toUtf8().data()) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackQStandardItemModel_SupportedDragActions(const_cast<MyQStandardItemModel*>(this), this->objectName().toUtf8().data())); };
	void timerEvent(QTimerEvent * event) { callbackQStandardItemModel_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQStandardItemModel_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQStandardItemModel_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQStandardItemModel_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQStandardItemModel_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQStandardItemModel_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQStandardItemModel_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQStandardItemModel_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQStandardItemModel_MetaObject(const_cast<MyQStandardItemModel*>(this), this->objectName().toUtf8().data())); };
};

void QStandardItemModel_SetSortRole(void* ptr, int role)
{
	static_cast<QStandardItemModel*>(ptr)->setSortRole(role);
}

int QStandardItemModel_SortRole(void* ptr)
{
	return static_cast<QStandardItemModel*>(ptr)->sortRole();
}

void* QStandardItemModel_NewQStandardItemModel(void* parent)
{
	return new MyQStandardItemModel(static_cast<QObject*>(parent));
}

void* QStandardItemModel_NewQStandardItemModel2(int rows, int columns, void* parent)
{
	return new MyQStandardItemModel(rows, columns, static_cast<QObject*>(parent));
}

void QStandardItemModel_AppendRow2(void* ptr, void* item)
{
	static_cast<QStandardItemModel*>(ptr)->appendRow(static_cast<QStandardItem*>(item));
}

void QStandardItemModel_Clear(void* ptr)
{
	static_cast<QStandardItemModel*>(ptr)->clear();
}

int QStandardItemModel_ColumnCount(void* ptr, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* QStandardItemModel_Data(void* ptr, void* index, int role)
{
	return new QVariant(static_cast<QStandardItemModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

void* QStandardItemModel_DataDefault(void* ptr, void* index, int role)
{
	return new QVariant(static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::data(*static_cast<QModelIndex*>(index), role));
}

int QStandardItemModel_DropMimeData(void* ptr, void* data, int action, int row, int column, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_DropMimeDataDefault(void* ptr, void* data, int action, int row, int column, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_Flags(void* ptr, void* index)
{
	return static_cast<QStandardItemModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

int QStandardItemModel_FlagsDefault(void* ptr, void* index)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::flags(*static_cast<QModelIndex*>(index));
}

int QStandardItemModel_HasChildren(void* ptr, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* QStandardItemModel_HeaderData(void* ptr, int section, int orientation, int role)
{
	return new QVariant(static_cast<QStandardItemModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QStandardItemModel_HeaderDataDefault(void* ptr, int section, int orientation, int role)
{
	return new QVariant(static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QStandardItemModel_HorizontalHeaderItem(void* ptr, int column)
{
	return static_cast<QStandardItemModel*>(ptr)->horizontalHeaderItem(column);
}

void* QStandardItemModel_Index(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<QStandardItemModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* QStandardItemModel_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* QStandardItemModel_IndexFromItem(void* ptr, void* item)
{
	return new QModelIndex(static_cast<QStandardItemModel*>(ptr)->indexFromItem(static_cast<QStandardItem*>(item)));
}

int QStandardItemModel_InsertColumn2(void* ptr, int column, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->insertColumn(column, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_InsertColumns(void* ptr, int column, int count, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_InsertRow2(void* ptr, int row, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->insertRow(row, *static_cast<QModelIndex*>(parent));
}

void QStandardItemModel_InsertRow3(void* ptr, int row, void* item)
{
	static_cast<QStandardItemModel*>(ptr)->insertRow(row, static_cast<QStandardItem*>(item));
}

int QStandardItemModel_InsertRows(void* ptr, int row, int count, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

void* QStandardItemModel_InvisibleRootItem(void* ptr)
{
	return static_cast<QStandardItemModel*>(ptr)->invisibleRootItem();
}

void* QStandardItemModel_Item(void* ptr, int row, int column)
{
	return static_cast<QStandardItemModel*>(ptr)->item(row, column);
}

void QStandardItemModel_ConnectItemChanged(void* ptr)
{
	QObject::connect(static_cast<QStandardItemModel*>(ptr), static_cast<void (QStandardItemModel::*)(QStandardItem *)>(&QStandardItemModel::itemChanged), static_cast<MyQStandardItemModel*>(ptr), static_cast<void (MyQStandardItemModel::*)(QStandardItem *)>(&MyQStandardItemModel::Signal_ItemChanged));
}

void QStandardItemModel_DisconnectItemChanged(void* ptr)
{
	QObject::disconnect(static_cast<QStandardItemModel*>(ptr), static_cast<void (QStandardItemModel::*)(QStandardItem *)>(&QStandardItemModel::itemChanged), static_cast<MyQStandardItemModel*>(ptr), static_cast<void (MyQStandardItemModel::*)(QStandardItem *)>(&MyQStandardItemModel::Signal_ItemChanged));
}

void QStandardItemModel_ItemChanged(void* ptr, void* item)
{
	static_cast<QStandardItemModel*>(ptr)->itemChanged(static_cast<QStandardItem*>(item));
}

void* QStandardItemModel_ItemFromIndex(void* ptr, void* index)
{
	return static_cast<QStandardItemModel*>(ptr)->itemFromIndex(*static_cast<QModelIndex*>(index));
}

void* QStandardItemModel_ItemPrototype(void* ptr)
{
	return const_cast<QStandardItem*>(static_cast<QStandardItemModel*>(ptr)->itemPrototype());
}

char* QStandardItemModel_MimeTypes(void* ptr)
{
	return static_cast<QStandardItemModel*>(ptr)->mimeTypes().join("|").toUtf8().data();
}

char* QStandardItemModel_MimeTypesDefault(void* ptr)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::mimeTypes().join("|").toUtf8().data();
}

void* QStandardItemModel_Parent(void* ptr, void* child)
{
	return new QModelIndex(static_cast<QStandardItemModel*>(ptr)->parent(*static_cast<QModelIndex*>(child)));
}

void* QStandardItemModel_ParentDefault(void* ptr, void* child)
{
	return new QModelIndex(static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::parent(*static_cast<QModelIndex*>(child)));
}

int QStandardItemModel_RemoveColumns(void* ptr, int column, int count, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_RemoveRows(void* ptr, int row, int count, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_RowCount(void* ptr, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_RowCountDefault(void* ptr, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::rowCount(*static_cast<QModelIndex*>(parent));
}

void QStandardItemModel_SetColumnCount(void* ptr, int columns)
{
	static_cast<QStandardItemModel*>(ptr)->setColumnCount(columns);
}

int QStandardItemModel_SetData(void* ptr, void* index, void* value, int role)
{
	return static_cast<QStandardItemModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

int QStandardItemModel_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

int QStandardItemModel_SetHeaderData(void* ptr, int section, int orientation, void* value, int role)
{
	return static_cast<QStandardItemModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

int QStandardItemModel_SetHeaderDataDefault(void* ptr, int section, int orientation, void* value, int role)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

void QStandardItemModel_SetHorizontalHeaderItem(void* ptr, int column, void* item)
{
	static_cast<QStandardItemModel*>(ptr)->setHorizontalHeaderItem(column, static_cast<QStandardItem*>(item));
}

void QStandardItemModel_SetHorizontalHeaderLabels(void* ptr, char* labels)
{
	static_cast<QStandardItemModel*>(ptr)->setHorizontalHeaderLabels(QString(labels).split("|", QString::SkipEmptyParts));
}

void QStandardItemModel_SetItem2(void* ptr, int row, void* item)
{
	static_cast<QStandardItemModel*>(ptr)->setItem(row, static_cast<QStandardItem*>(item));
}

void QStandardItemModel_SetItem(void* ptr, int row, int column, void* item)
{
	static_cast<QStandardItemModel*>(ptr)->setItem(row, column, static_cast<QStandardItem*>(item));
}

void QStandardItemModel_SetItemPrototype(void* ptr, void* item)
{
	static_cast<QStandardItemModel*>(ptr)->setItemPrototype(static_cast<QStandardItem*>(item));
}

void QStandardItemModel_SetRowCount(void* ptr, int rows)
{
	static_cast<QStandardItemModel*>(ptr)->setRowCount(rows);
}

void QStandardItemModel_SetVerticalHeaderItem(void* ptr, int row, void* item)
{
	static_cast<QStandardItemModel*>(ptr)->setVerticalHeaderItem(row, static_cast<QStandardItem*>(item));
}

void QStandardItemModel_SetVerticalHeaderLabels(void* ptr, char* labels)
{
	static_cast<QStandardItemModel*>(ptr)->setVerticalHeaderLabels(QString(labels).split("|", QString::SkipEmptyParts));
}

void* QStandardItemModel_Sibling(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<QStandardItemModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* QStandardItemModel_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void QStandardItemModel_Sort(void* ptr, int column, int order)
{
	static_cast<QStandardItemModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void QStandardItemModel_SortDefault(void* ptr, int column, int order)
{
	static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::sort(column, static_cast<Qt::SortOrder>(order));
}

int QStandardItemModel_SupportedDropActions(void* ptr)
{
	return static_cast<QStandardItemModel*>(ptr)->supportedDropActions();
}

int QStandardItemModel_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::supportedDropActions();
}

void* QStandardItemModel_TakeHorizontalHeaderItem(void* ptr, int column)
{
	return static_cast<QStandardItemModel*>(ptr)->takeHorizontalHeaderItem(column);
}

void* QStandardItemModel_TakeItem(void* ptr, int row, int column)
{
	return static_cast<QStandardItemModel*>(ptr)->takeItem(row, column);
}

void* QStandardItemModel_TakeVerticalHeaderItem(void* ptr, int row)
{
	return static_cast<QStandardItemModel*>(ptr)->takeVerticalHeaderItem(row);
}

void* QStandardItemModel_VerticalHeaderItem(void* ptr, int row)
{
	return static_cast<QStandardItemModel*>(ptr)->verticalHeaderItem(row);
}

void QStandardItemModel_DestroyQStandardItemModel(void* ptr)
{
	static_cast<QStandardItemModel*>(ptr)->~QStandardItemModel();
}

void* QStandardItemModel_Buddy(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QStandardItemModel*>(ptr)->buddy(*static_cast<QModelIndex*>(index)));
}

void* QStandardItemModel_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::buddy(*static_cast<QModelIndex*>(index)));
}

int QStandardItemModel_CanDropMimeData(void* ptr, void* data, int action, int row, int column, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_CanDropMimeDataDefault(void* ptr, void* data, int action, int row, int column, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_CanFetchMore(void* ptr, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

void QStandardItemModel_FetchMore(void* ptr, void* parent)
{
	static_cast<QStandardItemModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

void QStandardItemModel_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

int QStandardItemModel_MoveColumns(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QStandardItemModel*>(ptr)->moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

int QStandardItemModel_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

int QStandardItemModel_MoveRows(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QStandardItemModel*>(ptr)->moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

int QStandardItemModel_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void QStandardItemModel_ResetInternalData(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QStandardItemModel*>(ptr), "resetInternalData");
}

void QStandardItemModel_ResetInternalDataDefault(void* ptr)
{
	static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::resetInternalData();
}

void QStandardItemModel_Revert(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QStandardItemModel*>(ptr), "revert");
}

void QStandardItemModel_RevertDefault(void* ptr)
{
	static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::revert();
}

void* QStandardItemModel_Span(void* ptr, void* index)
{
	return new QSize(static_cast<QSize>(static_cast<QStandardItemModel*>(ptr)->span(*static_cast<QModelIndex*>(index))).width(), static_cast<QSize>(static_cast<QStandardItemModel*>(ptr)->span(*static_cast<QModelIndex*>(index))).height());
}

void* QStandardItemModel_SpanDefault(void* ptr, void* index)
{
	return new QSize(static_cast<QSize>(static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::span(*static_cast<QModelIndex*>(index))).width(), static_cast<QSize>(static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::span(*static_cast<QModelIndex*>(index))).height());
}

int QStandardItemModel_Submit(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QStandardItemModel*>(ptr), "submit", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

int QStandardItemModel_SubmitDefault(void* ptr)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::submit();
}

int QStandardItemModel_SupportedDragActions(void* ptr)
{
	return static_cast<QStandardItemModel*>(ptr)->supportedDragActions();
}

int QStandardItemModel_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::supportedDragActions();
}

void QStandardItemModel_TimerEvent(void* ptr, void* event)
{
	static_cast<QStandardItemModel*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QStandardItemModel_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void QStandardItemModel_ChildEvent(void* ptr, void* event)
{
	static_cast<QStandardItemModel*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QStandardItemModel_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::childEvent(static_cast<QChildEvent*>(event));
}

void QStandardItemModel_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QStandardItemModel*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QStandardItemModel_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QStandardItemModel_CustomEvent(void* ptr, void* event)
{
	static_cast<QStandardItemModel*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QStandardItemModel_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::customEvent(static_cast<QEvent*>(event));
}

void QStandardItemModel_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QStandardItemModel*>(ptr), "deleteLater");
}

void QStandardItemModel_DeleteLaterDefault(void* ptr)
{
	static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::deleteLater();
}

void QStandardItemModel_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QStandardItemModel*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QStandardItemModel_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QStandardItemModel_Event(void* ptr, void* e)
{
	return static_cast<QStandardItemModel*>(ptr)->event(static_cast<QEvent*>(e));
}

int QStandardItemModel_EventDefault(void* ptr, void* e)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::event(static_cast<QEvent*>(e));
}

int QStandardItemModel_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QStandardItemModel*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QStandardItemModel_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QStandardItemModel_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QStandardItemModel*>(ptr)->metaObject());
}

void* QStandardItemModel_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QStandardItemModel*>(ptr)->QStandardItemModel::metaObject());
}

void* QStaticText_NewQStaticText()
{
	return new QStaticText();
}

void* QStaticText_NewQStaticText3(void* other)
{
	return new QStaticText(*static_cast<QStaticText*>(other));
}

void* QStaticText_NewQStaticText2(char* text)
{
	return new QStaticText(QString(text));
}

int QStaticText_PerformanceHint(void* ptr)
{
	return static_cast<QStaticText*>(ptr)->performanceHint();
}

void QStaticText_Prepare(void* ptr, void* matrix, void* font)
{
	static_cast<QStaticText*>(ptr)->prepare(*static_cast<QTransform*>(matrix), *static_cast<QFont*>(font));
}

void QStaticText_SetPerformanceHint(void* ptr, int performanceHint)
{
	static_cast<QStaticText*>(ptr)->setPerformanceHint(static_cast<QStaticText::PerformanceHint>(performanceHint));
}

void QStaticText_SetText(void* ptr, char* text)
{
	static_cast<QStaticText*>(ptr)->setText(QString(text));
}

void QStaticText_SetTextFormat(void* ptr, int textFormat)
{
	static_cast<QStaticText*>(ptr)->setTextFormat(static_cast<Qt::TextFormat>(textFormat));
}

void QStaticText_SetTextOption(void* ptr, void* textOption)
{
	static_cast<QStaticText*>(ptr)->setTextOption(*static_cast<QTextOption*>(textOption));
}

void QStaticText_SetTextWidth(void* ptr, double textWidth)
{
	static_cast<QStaticText*>(ptr)->setTextWidth(static_cast<double>(textWidth));
}

void* QStaticText_Size(void* ptr)
{
	return new QSizeF(static_cast<QSizeF>(static_cast<QStaticText*>(ptr)->size()).width(), static_cast<QSizeF>(static_cast<QStaticText*>(ptr)->size()).height());
}

void QStaticText_Swap(void* ptr, void* other)
{
	static_cast<QStaticText*>(ptr)->swap(*static_cast<QStaticText*>(other));
}

char* QStaticText_Text(void* ptr)
{
	return static_cast<QStaticText*>(ptr)->text().toUtf8().data();
}

int QStaticText_TextFormat(void* ptr)
{
	return static_cast<QStaticText*>(ptr)->textFormat();
}

void* QStaticText_TextOption(void* ptr)
{
	return new QTextOption(static_cast<QStaticText*>(ptr)->textOption());
}

double QStaticText_TextWidth(void* ptr)
{
	return static_cast<double>(static_cast<QStaticText*>(ptr)->textWidth());
}

void QStaticText_DestroyQStaticText(void* ptr)
{
	static_cast<QStaticText*>(ptr)->~QStaticText();
}

void* QStatusTipEvent_NewQStatusTipEvent(char* tip)
{
	return new QStatusTipEvent(QString(tip));
}

char* QStatusTipEvent_Tip(void* ptr)
{
	return static_cast<QStatusTipEvent*>(ptr)->tip().toUtf8().data();
}

class MyQStyleHints: public QStyleHints
{
public:
	void Signal_CursorFlashTimeChanged(int cursorFlashTime) { callbackQStyleHints_CursorFlashTimeChanged(this, this->objectName().toUtf8().data(), cursorFlashTime); };
	void Signal_KeyboardInputIntervalChanged(int keyboardInputInterval) { callbackQStyleHints_KeyboardInputIntervalChanged(this, this->objectName().toUtf8().data(), keyboardInputInterval); };
	void Signal_MouseDoubleClickIntervalChanged(int mouseDoubleClickInterval) { callbackQStyleHints_MouseDoubleClickIntervalChanged(this, this->objectName().toUtf8().data(), mouseDoubleClickInterval); };
	void Signal_StartDragDistanceChanged(int startDragDistance) { callbackQStyleHints_StartDragDistanceChanged(this, this->objectName().toUtf8().data(), startDragDistance); };
	void Signal_StartDragTimeChanged(int startDragTime) { callbackQStyleHints_StartDragTimeChanged(this, this->objectName().toUtf8().data(), startDragTime); };
	void timerEvent(QTimerEvent * event) { callbackQStyleHints_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQStyleHints_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQStyleHints_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQStyleHints_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQStyleHints_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQStyleHints_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQStyleHints_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQStyleHints_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQStyleHints_MetaObject(const_cast<MyQStyleHints*>(this), this->objectName().toUtf8().data())); };
};

int QStyleHints_CursorFlashTime(void* ptr)
{
	return static_cast<QStyleHints*>(ptr)->cursorFlashTime();
}

double QStyleHints_FontSmoothingGamma(void* ptr)
{
	return static_cast<double>(static_cast<QStyleHints*>(ptr)->fontSmoothingGamma());
}

int QStyleHints_KeyboardAutoRepeatRate(void* ptr)
{
	return static_cast<QStyleHints*>(ptr)->keyboardAutoRepeatRate();
}

int QStyleHints_KeyboardInputInterval(void* ptr)
{
	return static_cast<QStyleHints*>(ptr)->keyboardInputInterval();
}

int QStyleHints_MouseDoubleClickInterval(void* ptr)
{
	return static_cast<QStyleHints*>(ptr)->mouseDoubleClickInterval();
}

int QStyleHints_MousePressAndHoldInterval(void* ptr)
{
	return static_cast<QStyleHints*>(ptr)->mousePressAndHoldInterval();
}

int QStyleHints_PasswordMaskDelay(void* ptr)
{
	return static_cast<QStyleHints*>(ptr)->passwordMaskDelay();
}

int QStyleHints_SetFocusOnTouchRelease(void* ptr)
{
	return static_cast<QStyleHints*>(ptr)->setFocusOnTouchRelease();
}

int QStyleHints_ShowIsFullScreen(void* ptr)
{
	return static_cast<QStyleHints*>(ptr)->showIsFullScreen();
}

int QStyleHints_SingleClickActivation(void* ptr)
{
	return static_cast<QStyleHints*>(ptr)->singleClickActivation();
}

int QStyleHints_StartDragDistance(void* ptr)
{
	return static_cast<QStyleHints*>(ptr)->startDragDistance();
}

int QStyleHints_StartDragTime(void* ptr)
{
	return static_cast<QStyleHints*>(ptr)->startDragTime();
}

int QStyleHints_StartDragVelocity(void* ptr)
{
	return static_cast<QStyleHints*>(ptr)->startDragVelocity();
}

int QStyleHints_TabFocusBehavior(void* ptr)
{
	return static_cast<QStyleHints*>(ptr)->tabFocusBehavior();
}

int QStyleHints_UseRtlExtensions(void* ptr)
{
	return static_cast<QStyleHints*>(ptr)->useRtlExtensions();
}

void QStyleHints_ConnectCursorFlashTimeChanged(void* ptr)
{
	QObject::connect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::cursorFlashTimeChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_CursorFlashTimeChanged));
}

void QStyleHints_DisconnectCursorFlashTimeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::cursorFlashTimeChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_CursorFlashTimeChanged));
}

void QStyleHints_CursorFlashTimeChanged(void* ptr, int cursorFlashTime)
{
	static_cast<QStyleHints*>(ptr)->cursorFlashTimeChanged(cursorFlashTime);
}

void QStyleHints_ConnectKeyboardInputIntervalChanged(void* ptr)
{
	QObject::connect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::keyboardInputIntervalChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_KeyboardInputIntervalChanged));
}

void QStyleHints_DisconnectKeyboardInputIntervalChanged(void* ptr)
{
	QObject::disconnect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::keyboardInputIntervalChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_KeyboardInputIntervalChanged));
}

void QStyleHints_KeyboardInputIntervalChanged(void* ptr, int keyboardInputInterval)
{
	static_cast<QStyleHints*>(ptr)->keyboardInputIntervalChanged(keyboardInputInterval);
}

void QStyleHints_ConnectMouseDoubleClickIntervalChanged(void* ptr)
{
	QObject::connect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::mouseDoubleClickIntervalChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_MouseDoubleClickIntervalChanged));
}

void QStyleHints_DisconnectMouseDoubleClickIntervalChanged(void* ptr)
{
	QObject::disconnect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::mouseDoubleClickIntervalChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_MouseDoubleClickIntervalChanged));
}

void QStyleHints_MouseDoubleClickIntervalChanged(void* ptr, int mouseDoubleClickInterval)
{
	static_cast<QStyleHints*>(ptr)->mouseDoubleClickIntervalChanged(mouseDoubleClickInterval);
}

void QStyleHints_ConnectStartDragDistanceChanged(void* ptr)
{
	QObject::connect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::startDragDistanceChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_StartDragDistanceChanged));
}

void QStyleHints_DisconnectStartDragDistanceChanged(void* ptr)
{
	QObject::disconnect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::startDragDistanceChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_StartDragDistanceChanged));
}

void QStyleHints_StartDragDistanceChanged(void* ptr, int startDragDistance)
{
	static_cast<QStyleHints*>(ptr)->startDragDistanceChanged(startDragDistance);
}

void QStyleHints_ConnectStartDragTimeChanged(void* ptr)
{
	QObject::connect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::startDragTimeChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_StartDragTimeChanged));
}

void QStyleHints_DisconnectStartDragTimeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QStyleHints*>(ptr), static_cast<void (QStyleHints::*)(int)>(&QStyleHints::startDragTimeChanged), static_cast<MyQStyleHints*>(ptr), static_cast<void (MyQStyleHints::*)(int)>(&MyQStyleHints::Signal_StartDragTimeChanged));
}

void QStyleHints_StartDragTimeChanged(void* ptr, int startDragTime)
{
	static_cast<QStyleHints*>(ptr)->startDragTimeChanged(startDragTime);
}

void QStyleHints_TimerEvent(void* ptr, void* event)
{
	static_cast<QStyleHints*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QStyleHints_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QStyleHints*>(ptr)->QStyleHints::timerEvent(static_cast<QTimerEvent*>(event));
}

void QStyleHints_ChildEvent(void* ptr, void* event)
{
	static_cast<QStyleHints*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QStyleHints_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QStyleHints*>(ptr)->QStyleHints::childEvent(static_cast<QChildEvent*>(event));
}

void QStyleHints_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QStyleHints*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QStyleHints_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QStyleHints*>(ptr)->QStyleHints::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QStyleHints_CustomEvent(void* ptr, void* event)
{
	static_cast<QStyleHints*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QStyleHints_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QStyleHints*>(ptr)->QStyleHints::customEvent(static_cast<QEvent*>(event));
}

void QStyleHints_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QStyleHints*>(ptr), "deleteLater");
}

void QStyleHints_DeleteLaterDefault(void* ptr)
{
	static_cast<QStyleHints*>(ptr)->QStyleHints::deleteLater();
}

void QStyleHints_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QStyleHints*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QStyleHints_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QStyleHints*>(ptr)->QStyleHints::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QStyleHints_Event(void* ptr, void* e)
{
	return static_cast<QStyleHints*>(ptr)->event(static_cast<QEvent*>(e));
}

int QStyleHints_EventDefault(void* ptr, void* e)
{
	return static_cast<QStyleHints*>(ptr)->QStyleHints::event(static_cast<QEvent*>(e));
}

int QStyleHints_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QStyleHints*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QStyleHints_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QStyleHints*>(ptr)->QStyleHints::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QStyleHints_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QStyleHints*>(ptr)->metaObject());
}

void* QStyleHints_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QStyleHints*>(ptr)->QStyleHints::metaObject());
}

class MyQSurface: public QSurface
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	QSurfaceFormat format() const { return *static_cast<QSurfaceFormat*>(callbackQSurface_Format(const_cast<MyQSurface*>(this), this->objectNameAbs().toUtf8().data())); };
	QSize size() const { return *static_cast<QSize*>(callbackQSurface_Size(const_cast<MyQSurface*>(this), this->objectNameAbs().toUtf8().data())); };
	SurfaceType surfaceType() const { return static_cast<QSurface::SurfaceType>(callbackQSurface_SurfaceType(const_cast<MyQSurface*>(this), this->objectNameAbs().toUtf8().data())); };
};

void* QSurface_Format(void* ptr)
{
	return new QSurfaceFormat(static_cast<QSurface*>(ptr)->format());
}

void* QSurface_Size(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QSurface*>(ptr)->size()).width(), static_cast<QSize>(static_cast<QSurface*>(ptr)->size()).height());
}

int QSurface_SupportsOpenGL(void* ptr)
{
	return static_cast<QSurface*>(ptr)->supportsOpenGL();
}

int QSurface_SurfaceClass(void* ptr)
{
	return static_cast<QSurface*>(ptr)->surfaceClass();
}

int QSurface_SurfaceType(void* ptr)
{
	return static_cast<QSurface*>(ptr)->surfaceType();
}

void QSurface_DestroyQSurface(void* ptr)
{
	static_cast<QSurface*>(ptr)->~QSurface();
}

char* QSurface_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQSurface*>(static_cast<QSurface*>(ptr))) {
		return static_cast<MyQSurface*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QSurface_BASE").toUtf8().data();
}

void QSurface_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQSurface*>(static_cast<QSurface*>(ptr))) {
		static_cast<MyQSurface*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QSurfaceFormat_NewQSurfaceFormat()
{
	return new QSurfaceFormat();
}

void* QSurfaceFormat_NewQSurfaceFormat2(int options)
{
	return new QSurfaceFormat(static_cast<QSurfaceFormat::FormatOption>(options));
}

void* QSurfaceFormat_NewQSurfaceFormat3(void* other)
{
	return new QSurfaceFormat(*static_cast<QSurfaceFormat*>(other));
}

int QSurfaceFormat_AlphaBufferSize(void* ptr)
{
	return static_cast<QSurfaceFormat*>(ptr)->alphaBufferSize();
}

int QSurfaceFormat_BlueBufferSize(void* ptr)
{
	return static_cast<QSurfaceFormat*>(ptr)->blueBufferSize();
}

void* QSurfaceFormat_QSurfaceFormat_DefaultFormat()
{
	return new QSurfaceFormat(QSurfaceFormat::defaultFormat());
}

int QSurfaceFormat_DepthBufferSize(void* ptr)
{
	return static_cast<QSurfaceFormat*>(ptr)->depthBufferSize();
}

int QSurfaceFormat_GreenBufferSize(void* ptr)
{
	return static_cast<QSurfaceFormat*>(ptr)->greenBufferSize();
}

int QSurfaceFormat_HasAlpha(void* ptr)
{
	return static_cast<QSurfaceFormat*>(ptr)->hasAlpha();
}

int QSurfaceFormat_MajorVersion(void* ptr)
{
	return static_cast<QSurfaceFormat*>(ptr)->majorVersion();
}

int QSurfaceFormat_MinorVersion(void* ptr)
{
	return static_cast<QSurfaceFormat*>(ptr)->minorVersion();
}

int QSurfaceFormat_Options(void* ptr)
{
	return static_cast<QSurfaceFormat*>(ptr)->options();
}

int QSurfaceFormat_Profile(void* ptr)
{
	return static_cast<QSurfaceFormat*>(ptr)->profile();
}

int QSurfaceFormat_RedBufferSize(void* ptr)
{
	return static_cast<QSurfaceFormat*>(ptr)->redBufferSize();
}

int QSurfaceFormat_RenderableType(void* ptr)
{
	return static_cast<QSurfaceFormat*>(ptr)->renderableType();
}

int QSurfaceFormat_Samples(void* ptr)
{
	return static_cast<QSurfaceFormat*>(ptr)->samples();
}

void QSurfaceFormat_SetAlphaBufferSize(void* ptr, int size)
{
	static_cast<QSurfaceFormat*>(ptr)->setAlphaBufferSize(size);
}

void QSurfaceFormat_SetBlueBufferSize(void* ptr, int size)
{
	static_cast<QSurfaceFormat*>(ptr)->setBlueBufferSize(size);
}

void QSurfaceFormat_QSurfaceFormat_SetDefaultFormat(void* format)
{
	QSurfaceFormat::setDefaultFormat(*static_cast<QSurfaceFormat*>(format));
}

void QSurfaceFormat_SetDepthBufferSize(void* ptr, int size)
{
	static_cast<QSurfaceFormat*>(ptr)->setDepthBufferSize(size);
}

void QSurfaceFormat_SetGreenBufferSize(void* ptr, int size)
{
	static_cast<QSurfaceFormat*>(ptr)->setGreenBufferSize(size);
}

void QSurfaceFormat_SetMajorVersion(void* ptr, int major)
{
	static_cast<QSurfaceFormat*>(ptr)->setMajorVersion(major);
}

void QSurfaceFormat_SetMinorVersion(void* ptr, int minor)
{
	static_cast<QSurfaceFormat*>(ptr)->setMinorVersion(minor);
}

void QSurfaceFormat_SetOption(void* ptr, int option, int on)
{
	static_cast<QSurfaceFormat*>(ptr)->setOption(static_cast<QSurfaceFormat::FormatOption>(option), on != 0);
}

void QSurfaceFormat_SetOptions(void* ptr, int options)
{
	static_cast<QSurfaceFormat*>(ptr)->setOptions(static_cast<QSurfaceFormat::FormatOption>(options));
}

void QSurfaceFormat_SetProfile(void* ptr, int profile)
{
	static_cast<QSurfaceFormat*>(ptr)->setProfile(static_cast<QSurfaceFormat::OpenGLContextProfile>(profile));
}

void QSurfaceFormat_SetRedBufferSize(void* ptr, int size)
{
	static_cast<QSurfaceFormat*>(ptr)->setRedBufferSize(size);
}

void QSurfaceFormat_SetRenderableType(void* ptr, int ty)
{
	static_cast<QSurfaceFormat*>(ptr)->setRenderableType(static_cast<QSurfaceFormat::RenderableType>(ty));
}

void QSurfaceFormat_SetSamples(void* ptr, int numSamples)
{
	static_cast<QSurfaceFormat*>(ptr)->setSamples(numSamples);
}

void QSurfaceFormat_SetStencilBufferSize(void* ptr, int size)
{
	static_cast<QSurfaceFormat*>(ptr)->setStencilBufferSize(size);
}

void QSurfaceFormat_SetStereo(void* ptr, int enable)
{
	static_cast<QSurfaceFormat*>(ptr)->setStereo(enable != 0);
}

void QSurfaceFormat_SetSwapBehavior(void* ptr, int behavior)
{
	static_cast<QSurfaceFormat*>(ptr)->setSwapBehavior(static_cast<QSurfaceFormat::SwapBehavior>(behavior));
}

void QSurfaceFormat_SetSwapInterval(void* ptr, int interval)
{
	static_cast<QSurfaceFormat*>(ptr)->setSwapInterval(interval);
}

void QSurfaceFormat_SetVersion(void* ptr, int major, int minor)
{
	static_cast<QSurfaceFormat*>(ptr)->setVersion(major, minor);
}

int QSurfaceFormat_StencilBufferSize(void* ptr)
{
	return static_cast<QSurfaceFormat*>(ptr)->stencilBufferSize();
}

int QSurfaceFormat_Stereo(void* ptr)
{
	return static_cast<QSurfaceFormat*>(ptr)->stereo();
}

int QSurfaceFormat_SwapBehavior(void* ptr)
{
	return static_cast<QSurfaceFormat*>(ptr)->swapBehavior();
}

int QSurfaceFormat_SwapInterval(void* ptr)
{
	return static_cast<QSurfaceFormat*>(ptr)->swapInterval();
}

int QSurfaceFormat_TestOption(void* ptr, int option)
{
	return static_cast<QSurfaceFormat*>(ptr)->testOption(static_cast<QSurfaceFormat::FormatOption>(option));
}

void QSurfaceFormat_DestroyQSurfaceFormat(void* ptr)
{
	static_cast<QSurfaceFormat*>(ptr)->~QSurfaceFormat();
}

class MyQSyntaxHighlighter: public QSyntaxHighlighter
{
public:
	MyQSyntaxHighlighter(QObject *parent) : QSyntaxHighlighter(parent) {};
	MyQSyntaxHighlighter(QTextDocument *parent) : QSyntaxHighlighter(parent) {};
	void highlightBlock(const QString & text) { callbackQSyntaxHighlighter_HighlightBlock(this, this->objectName().toUtf8().data(), text.toUtf8().data()); };
	void rehighlight() { callbackQSyntaxHighlighter_Rehighlight(this, this->objectName().toUtf8().data()); };
	void rehighlightBlock(const QTextBlock & block) { callbackQSyntaxHighlighter_RehighlightBlock(this, this->objectName().toUtf8().data(), new QTextBlock(block)); };
	void timerEvent(QTimerEvent * event) { callbackQSyntaxHighlighter_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQSyntaxHighlighter_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQSyntaxHighlighter_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQSyntaxHighlighter_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQSyntaxHighlighter_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQSyntaxHighlighter_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQSyntaxHighlighter_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQSyntaxHighlighter_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQSyntaxHighlighter_MetaObject(const_cast<MyQSyntaxHighlighter*>(this), this->objectName().toUtf8().data())); };
};

void* QSyntaxHighlighter_NewQSyntaxHighlighter(void* parent)
{
	return new MyQSyntaxHighlighter(static_cast<QObject*>(parent));
}

void* QSyntaxHighlighter_NewQSyntaxHighlighter2(void* parent)
{
	return new MyQSyntaxHighlighter(static_cast<QTextDocument*>(parent));
}

void* QSyntaxHighlighter_CurrentBlock(void* ptr)
{
	return new QTextBlock(static_cast<QSyntaxHighlighter*>(ptr)->currentBlock());
}

int QSyntaxHighlighter_CurrentBlockState(void* ptr)
{
	return static_cast<QSyntaxHighlighter*>(ptr)->currentBlockState();
}

void* QSyntaxHighlighter_CurrentBlockUserData(void* ptr)
{
	return static_cast<QSyntaxHighlighter*>(ptr)->currentBlockUserData();
}

void* QSyntaxHighlighter_Document(void* ptr)
{
	return static_cast<QSyntaxHighlighter*>(ptr)->document();
}

void QSyntaxHighlighter_HighlightBlock(void* ptr, char* text)
{
	static_cast<QSyntaxHighlighter*>(ptr)->highlightBlock(QString(text));
}

int QSyntaxHighlighter_PreviousBlockState(void* ptr)
{
	return static_cast<QSyntaxHighlighter*>(ptr)->previousBlockState();
}

void QSyntaxHighlighter_Rehighlight(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSyntaxHighlighter*>(ptr), "rehighlight");
}

void QSyntaxHighlighter_RehighlightBlock(void* ptr, void* block)
{
	QMetaObject::invokeMethod(static_cast<QSyntaxHighlighter*>(ptr), "rehighlightBlock", Q_ARG(QTextBlock, *static_cast<QTextBlock*>(block)));
}

void QSyntaxHighlighter_SetCurrentBlockState(void* ptr, int newState)
{
	static_cast<QSyntaxHighlighter*>(ptr)->setCurrentBlockState(newState);
}

void QSyntaxHighlighter_SetCurrentBlockUserData(void* ptr, void* data)
{
	static_cast<QSyntaxHighlighter*>(ptr)->setCurrentBlockUserData(static_cast<QTextBlockUserData*>(data));
}

void QSyntaxHighlighter_SetDocument(void* ptr, void* doc)
{
	static_cast<QSyntaxHighlighter*>(ptr)->setDocument(static_cast<QTextDocument*>(doc));
}

void QSyntaxHighlighter_SetFormat2(void* ptr, int start, int count, void* color)
{
	static_cast<QSyntaxHighlighter*>(ptr)->setFormat(start, count, *static_cast<QColor*>(color));
}

void QSyntaxHighlighter_SetFormat3(void* ptr, int start, int count, void* font)
{
	static_cast<QSyntaxHighlighter*>(ptr)->setFormat(start, count, *static_cast<QFont*>(font));
}

void QSyntaxHighlighter_SetFormat(void* ptr, int start, int count, void* format)
{
	static_cast<QSyntaxHighlighter*>(ptr)->setFormat(start, count, *static_cast<QTextCharFormat*>(format));
}

void QSyntaxHighlighter_DestroyQSyntaxHighlighter(void* ptr)
{
	static_cast<QSyntaxHighlighter*>(ptr)->~QSyntaxHighlighter();
}

void QSyntaxHighlighter_TimerEvent(void* ptr, void* event)
{
	static_cast<QSyntaxHighlighter*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QSyntaxHighlighter_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QSyntaxHighlighter*>(ptr)->QSyntaxHighlighter::timerEvent(static_cast<QTimerEvent*>(event));
}

void QSyntaxHighlighter_ChildEvent(void* ptr, void* event)
{
	static_cast<QSyntaxHighlighter*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QSyntaxHighlighter_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QSyntaxHighlighter*>(ptr)->QSyntaxHighlighter::childEvent(static_cast<QChildEvent*>(event));
}

void QSyntaxHighlighter_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QSyntaxHighlighter*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSyntaxHighlighter_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSyntaxHighlighter*>(ptr)->QSyntaxHighlighter::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSyntaxHighlighter_CustomEvent(void* ptr, void* event)
{
	static_cast<QSyntaxHighlighter*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QSyntaxHighlighter_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QSyntaxHighlighter*>(ptr)->QSyntaxHighlighter::customEvent(static_cast<QEvent*>(event));
}

void QSyntaxHighlighter_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QSyntaxHighlighter*>(ptr), "deleteLater");
}

void QSyntaxHighlighter_DeleteLaterDefault(void* ptr)
{
	static_cast<QSyntaxHighlighter*>(ptr)->QSyntaxHighlighter::deleteLater();
}

void QSyntaxHighlighter_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QSyntaxHighlighter*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QSyntaxHighlighter_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QSyntaxHighlighter*>(ptr)->QSyntaxHighlighter::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QSyntaxHighlighter_Event(void* ptr, void* e)
{
	return static_cast<QSyntaxHighlighter*>(ptr)->event(static_cast<QEvent*>(e));
}

int QSyntaxHighlighter_EventDefault(void* ptr, void* e)
{
	return static_cast<QSyntaxHighlighter*>(ptr)->QSyntaxHighlighter::event(static_cast<QEvent*>(e));
}

int QSyntaxHighlighter_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QSyntaxHighlighter*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QSyntaxHighlighter_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QSyntaxHighlighter*>(ptr)->QSyntaxHighlighter::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QSyntaxHighlighter_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSyntaxHighlighter*>(ptr)->metaObject());
}

void* QSyntaxHighlighter_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QSyntaxHighlighter*>(ptr)->QSyntaxHighlighter::metaObject());
}

void* QTabletEvent_NewQTabletEvent(int ty, void* pos, void* globalPos, int device, int pointerType, double pressure, int xTilt, int yTilt, double tangentialPressure, double rotation, int z, int keyState, long long uniqueID, int button, int buttons)
{
	return new QTabletEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(pos), *static_cast<QPointF*>(globalPos), device, pointerType, static_cast<double>(pressure), xTilt, yTilt, static_cast<double>(tangentialPressure), static_cast<double>(rotation), z, static_cast<Qt::KeyboardModifier>(keyState), static_cast<long long>(uniqueID), static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButton>(buttons));
}

int QTabletEvent_Button(void* ptr)
{
	return static_cast<QTabletEvent*>(ptr)->button();
}

int QTabletEvent_Buttons(void* ptr)
{
	return static_cast<QTabletEvent*>(ptr)->buttons();
}

int QTabletEvent_Device(void* ptr)
{
	return static_cast<QTabletEvent*>(ptr)->device();
}

void* QTabletEvent_GlobalPos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QTabletEvent*>(ptr)->globalPos()).x(), static_cast<QPoint>(static_cast<QTabletEvent*>(ptr)->globalPos()).y());
}

void* QTabletEvent_GlobalPosF(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QTabletEvent*>(ptr)->globalPosF()).x(), static_cast<QPointF>(static_cast<QTabletEvent*>(ptr)->globalPosF()).y());
}

int QTabletEvent_GlobalX(void* ptr)
{
	return static_cast<QTabletEvent*>(ptr)->globalX();
}

int QTabletEvent_GlobalY(void* ptr)
{
	return static_cast<QTabletEvent*>(ptr)->globalY();
}

double QTabletEvent_HiResGlobalX(void* ptr)
{
	return static_cast<double>(static_cast<QTabletEvent*>(ptr)->hiResGlobalX());
}

double QTabletEvent_HiResGlobalY(void* ptr)
{
	return static_cast<double>(static_cast<QTabletEvent*>(ptr)->hiResGlobalY());
}

int QTabletEvent_PointerType(void* ptr)
{
	return static_cast<QTabletEvent*>(ptr)->pointerType();
}

void* QTabletEvent_Pos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QTabletEvent*>(ptr)->pos()).x(), static_cast<QPoint>(static_cast<QTabletEvent*>(ptr)->pos()).y());
}

void* QTabletEvent_PosF(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QTabletEvent*>(ptr)->posF()).x(), static_cast<QPointF>(static_cast<QTabletEvent*>(ptr)->posF()).y());
}

double QTabletEvent_Pressure(void* ptr)
{
	return static_cast<double>(static_cast<QTabletEvent*>(ptr)->pressure());
}

double QTabletEvent_Rotation(void* ptr)
{
	return static_cast<double>(static_cast<QTabletEvent*>(ptr)->rotation());
}

double QTabletEvent_TangentialPressure(void* ptr)
{
	return static_cast<double>(static_cast<QTabletEvent*>(ptr)->tangentialPressure());
}

long long QTabletEvent_UniqueId(void* ptr)
{
	return static_cast<long long>(static_cast<QTabletEvent*>(ptr)->uniqueId());
}

int QTabletEvent_X(void* ptr)
{
	return static_cast<QTabletEvent*>(ptr)->x();
}

int QTabletEvent_XTilt(void* ptr)
{
	return static_cast<QTabletEvent*>(ptr)->xTilt();
}

int QTabletEvent_Y(void* ptr)
{
	return static_cast<QTabletEvent*>(ptr)->y();
}

int QTabletEvent_YTilt(void* ptr)
{
	return static_cast<QTabletEvent*>(ptr)->yTilt();
}

int QTabletEvent_Z(void* ptr)
{
	return static_cast<QTabletEvent*>(ptr)->z();
}

int QTextBlock_IsValid(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->isValid();
}

void* QTextBlock_NewQTextBlock(void* other)
{
	return new QTextBlock(*static_cast<QTextBlock*>(other));
}

int QTextBlock_BlockFormatIndex(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->blockFormatIndex();
}

int QTextBlock_CharFormatIndex(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->charFormatIndex();
}

void QTextBlock_ClearLayout(void* ptr)
{
	static_cast<QTextBlock*>(ptr)->clearLayout();
}

int QTextBlock_Contains(void* ptr, int position)
{
	return static_cast<QTextBlock*>(ptr)->contains(position);
}

int QTextBlock_BlockNumber(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->blockNumber();
}

void* QTextBlock_Document(void* ptr)
{
	return const_cast<QTextDocument*>(static_cast<QTextBlock*>(ptr)->document());
}

int QTextBlock_FirstLineNumber(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->firstLineNumber();
}

int QTextBlock_IsVisible(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->isVisible();
}

void* QTextBlock_Layout(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->layout();
}

int QTextBlock_Length(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->length();
}

int QTextBlock_LineCount(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->lineCount();
}

void* QTextBlock_Next(void* ptr)
{
	return new QTextBlock(static_cast<QTextBlock*>(ptr)->next());
}

int QTextBlock_Position(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->position();
}

void* QTextBlock_Previous(void* ptr)
{
	return new QTextBlock(static_cast<QTextBlock*>(ptr)->previous());
}

int QTextBlock_Revision(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->revision();
}

void QTextBlock_SetLineCount(void* ptr, int count)
{
	static_cast<QTextBlock*>(ptr)->setLineCount(count);
}

void QTextBlock_SetRevision(void* ptr, int rev)
{
	static_cast<QTextBlock*>(ptr)->setRevision(rev);
}

void QTextBlock_SetUserData(void* ptr, void* data)
{
	static_cast<QTextBlock*>(ptr)->setUserData(static_cast<QTextBlockUserData*>(data));
}

void QTextBlock_SetUserState(void* ptr, int state)
{
	static_cast<QTextBlock*>(ptr)->setUserState(state);
}

void QTextBlock_SetVisible(void* ptr, int visible)
{
	static_cast<QTextBlock*>(ptr)->setVisible(visible != 0);
}

char* QTextBlock_Text(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->text().toUtf8().data();
}

int QTextBlock_TextDirection(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->textDirection();
}

void* QTextBlock_TextList(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->textList();
}

void* QTextBlock_UserData(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->userData();
}

int QTextBlock_UserState(void* ptr)
{
	return static_cast<QTextBlock*>(ptr)->userState();
}

void* QTextBlockFormat_NewQTextBlockFormat()
{
	return new QTextBlockFormat();
}

int QTextBlockFormat_Alignment(void* ptr)
{
	return static_cast<QTextBlockFormat*>(ptr)->alignment();
}

double QTextBlockFormat_BottomMargin(void* ptr)
{
	return static_cast<double>(static_cast<QTextBlockFormat*>(ptr)->bottomMargin());
}

int QTextBlockFormat_Indent(void* ptr)
{
	return static_cast<QTextBlockFormat*>(ptr)->indent();
}

int QTextBlockFormat_IsValid(void* ptr)
{
	return static_cast<QTextBlockFormat*>(ptr)->isValid();
}

double QTextBlockFormat_LeftMargin(void* ptr)
{
	return static_cast<double>(static_cast<QTextBlockFormat*>(ptr)->leftMargin());
}

double QTextBlockFormat_LineHeight2(void* ptr)
{
	return static_cast<double>(static_cast<QTextBlockFormat*>(ptr)->lineHeight());
}

double QTextBlockFormat_LineHeight(void* ptr, double scriptLineHeight, double scaling)
{
	return static_cast<double>(static_cast<QTextBlockFormat*>(ptr)->lineHeight(static_cast<double>(scriptLineHeight), static_cast<double>(scaling)));
}

int QTextBlockFormat_LineHeightType(void* ptr)
{
	return static_cast<QTextBlockFormat*>(ptr)->lineHeightType();
}

int QTextBlockFormat_NonBreakableLines(void* ptr)
{
	return static_cast<QTextBlockFormat*>(ptr)->nonBreakableLines();
}

int QTextBlockFormat_PageBreakPolicy(void* ptr)
{
	return static_cast<QTextBlockFormat*>(ptr)->pageBreakPolicy();
}

double QTextBlockFormat_RightMargin(void* ptr)
{
	return static_cast<double>(static_cast<QTextBlockFormat*>(ptr)->rightMargin());
}

void QTextBlockFormat_SetAlignment(void* ptr, int alignment)
{
	static_cast<QTextBlockFormat*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QTextBlockFormat_SetBottomMargin(void* ptr, double margin)
{
	static_cast<QTextBlockFormat*>(ptr)->setBottomMargin(static_cast<double>(margin));
}

void QTextBlockFormat_SetIndent(void* ptr, int indentation)
{
	static_cast<QTextBlockFormat*>(ptr)->setIndent(indentation);
}

void QTextBlockFormat_SetLeftMargin(void* ptr, double margin)
{
	static_cast<QTextBlockFormat*>(ptr)->setLeftMargin(static_cast<double>(margin));
}

void QTextBlockFormat_SetLineHeight(void* ptr, double height, int heightType)
{
	static_cast<QTextBlockFormat*>(ptr)->setLineHeight(static_cast<double>(height), heightType);
}

void QTextBlockFormat_SetNonBreakableLines(void* ptr, int b)
{
	static_cast<QTextBlockFormat*>(ptr)->setNonBreakableLines(b != 0);
}

void QTextBlockFormat_SetPageBreakPolicy(void* ptr, int policy)
{
	static_cast<QTextBlockFormat*>(ptr)->setPageBreakPolicy(static_cast<QTextFormat::PageBreakFlag>(policy));
}

void QTextBlockFormat_SetRightMargin(void* ptr, double margin)
{
	static_cast<QTextBlockFormat*>(ptr)->setRightMargin(static_cast<double>(margin));
}

void QTextBlockFormat_SetTextIndent(void* ptr, double indent)
{
	static_cast<QTextBlockFormat*>(ptr)->setTextIndent(static_cast<double>(indent));
}

void QTextBlockFormat_SetTopMargin(void* ptr, double margin)
{
	static_cast<QTextBlockFormat*>(ptr)->setTopMargin(static_cast<double>(margin));
}

double QTextBlockFormat_TextIndent(void* ptr)
{
	return static_cast<double>(static_cast<QTextBlockFormat*>(ptr)->textIndent());
}

double QTextBlockFormat_TopMargin(void* ptr)
{
	return static_cast<double>(static_cast<QTextBlockFormat*>(ptr)->topMargin());
}

class MyQTextBlockGroup: public QTextBlockGroup
{
public:
	MyQTextBlockGroup(QTextDocument *document) : QTextBlockGroup(document) {};
	void blockFormatChanged(const QTextBlock & block) { callbackQTextBlockGroup_BlockFormatChanged(this, this->objectName().toUtf8().data(), new QTextBlock(block)); };
	void blockInserted(const QTextBlock & block) { callbackQTextBlockGroup_BlockInserted(this, this->objectName().toUtf8().data(), new QTextBlock(block)); };
	void blockRemoved(const QTextBlock & block) { callbackQTextBlockGroup_BlockRemoved(this, this->objectName().toUtf8().data(), new QTextBlock(block)); };
	void timerEvent(QTimerEvent * event) { callbackQTextBlockGroup_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQTextBlockGroup_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQTextBlockGroup_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQTextBlockGroup_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQTextBlockGroup_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQTextBlockGroup_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQTextBlockGroup_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQTextBlockGroup_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQTextBlockGroup_MetaObject(const_cast<MyQTextBlockGroup*>(this), this->objectName().toUtf8().data())); };
};

void* QTextBlockGroup_NewQTextBlockGroup(void* document)
{
	return new MyQTextBlockGroup(static_cast<QTextDocument*>(document));
}

void QTextBlockGroup_BlockFormatChanged(void* ptr, void* block)
{
	static_cast<QTextBlockGroup*>(ptr)->blockFormatChanged(*static_cast<QTextBlock*>(block));
}

void QTextBlockGroup_BlockFormatChangedDefault(void* ptr, void* block)
{
	static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::blockFormatChanged(*static_cast<QTextBlock*>(block));
}

void QTextBlockGroup_BlockInserted(void* ptr, void* block)
{
	static_cast<QTextBlockGroup*>(ptr)->blockInserted(*static_cast<QTextBlock*>(block));
}

void QTextBlockGroup_BlockInsertedDefault(void* ptr, void* block)
{
	static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::blockInserted(*static_cast<QTextBlock*>(block));
}

void QTextBlockGroup_BlockRemoved(void* ptr, void* block)
{
	static_cast<QTextBlockGroup*>(ptr)->blockRemoved(*static_cast<QTextBlock*>(block));
}

void QTextBlockGroup_BlockRemovedDefault(void* ptr, void* block)
{
	static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::blockRemoved(*static_cast<QTextBlock*>(block));
}

void QTextBlockGroup_DestroyQTextBlockGroup(void* ptr)
{
	static_cast<QTextBlockGroup*>(ptr)->~QTextBlockGroup();
}

void QTextBlockGroup_TimerEvent(void* ptr, void* event)
{
	static_cast<QTextBlockGroup*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTextBlockGroup_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTextBlockGroup_ChildEvent(void* ptr, void* event)
{
	static_cast<QTextBlockGroup*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTextBlockGroup_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::childEvent(static_cast<QChildEvent*>(event));
}

void QTextBlockGroup_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QTextBlockGroup*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextBlockGroup_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextBlockGroup_CustomEvent(void* ptr, void* event)
{
	static_cast<QTextBlockGroup*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTextBlockGroup_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::customEvent(static_cast<QEvent*>(event));
}

void QTextBlockGroup_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTextBlockGroup*>(ptr), "deleteLater");
}

void QTextBlockGroup_DeleteLaterDefault(void* ptr)
{
	static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::deleteLater();
}

void QTextBlockGroup_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QTextBlockGroup*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextBlockGroup_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QTextBlockGroup_Event(void* ptr, void* e)
{
	return static_cast<QTextBlockGroup*>(ptr)->event(static_cast<QEvent*>(e));
}

int QTextBlockGroup_EventDefault(void* ptr, void* e)
{
	return static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::event(static_cast<QEvent*>(e));
}

int QTextBlockGroup_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QTextBlockGroup*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QTextBlockGroup_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QTextBlockGroup_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTextBlockGroup*>(ptr)->metaObject());
}

void* QTextBlockGroup_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTextBlockGroup*>(ptr)->QTextBlockGroup::metaObject());
}

class MyQTextBlockUserData: public QTextBlockUserData
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
};

void QTextBlockUserData_DestroyQTextBlockUserData(void* ptr)
{
	static_cast<QTextBlockUserData*>(ptr)->~QTextBlockUserData();
}

char* QTextBlockUserData_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQTextBlockUserData*>(static_cast<QTextBlockUserData*>(ptr))) {
		return static_cast<MyQTextBlockUserData*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QTextBlockUserData_BASE").toUtf8().data();
}

void QTextBlockUserData_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQTextBlockUserData*>(static_cast<QTextBlockUserData*>(ptr))) {
		static_cast<MyQTextBlockUserData*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QTextCharFormat_NewQTextCharFormat()
{
	return new QTextCharFormat();
}

char* QTextCharFormat_AnchorNames(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->anchorNames().join("|").toUtf8().data();
}

int QTextCharFormat_FontUnderline(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->fontUnderline();
}

void QTextCharFormat_SetUnderlineStyle(void* ptr, int style)
{
	static_cast<QTextCharFormat*>(ptr)->setUnderlineStyle(static_cast<QTextCharFormat::UnderlineStyle>(style));
}

char* QTextCharFormat_AnchorHref(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->anchorHref().toUtf8().data();
}

void* QTextCharFormat_Font(void* ptr)
{
	return new QFont(static_cast<QTextCharFormat*>(ptr)->font());
}

int QTextCharFormat_FontCapitalization(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->fontCapitalization();
}

char* QTextCharFormat_FontFamily(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->fontFamily().toUtf8().data();
}

int QTextCharFormat_FontFixedPitch(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->fontFixedPitch();
}

int QTextCharFormat_FontHintingPreference(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->fontHintingPreference();
}

int QTextCharFormat_FontItalic(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->fontItalic();
}

int QTextCharFormat_FontKerning(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->fontKerning();
}

double QTextCharFormat_FontLetterSpacing(void* ptr)
{
	return static_cast<double>(static_cast<QTextCharFormat*>(ptr)->fontLetterSpacing());
}

int QTextCharFormat_FontLetterSpacingType(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->fontLetterSpacingType();
}

int QTextCharFormat_FontOverline(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->fontOverline();
}

double QTextCharFormat_FontPointSize(void* ptr)
{
	return static_cast<double>(static_cast<QTextCharFormat*>(ptr)->fontPointSize());
}

int QTextCharFormat_FontStretch(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->fontStretch();
}

int QTextCharFormat_FontStrikeOut(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->fontStrikeOut();
}

int QTextCharFormat_FontStyleHint(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->fontStyleHint();
}

int QTextCharFormat_FontStyleStrategy(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->fontStyleStrategy();
}

int QTextCharFormat_FontWeight(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->fontWeight();
}

double QTextCharFormat_FontWordSpacing(void* ptr)
{
	return static_cast<double>(static_cast<QTextCharFormat*>(ptr)->fontWordSpacing());
}

int QTextCharFormat_IsAnchor(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->isAnchor();
}

int QTextCharFormat_IsValid(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->isValid();
}

void QTextCharFormat_SetAnchor(void* ptr, int anchor)
{
	static_cast<QTextCharFormat*>(ptr)->setAnchor(anchor != 0);
}

void QTextCharFormat_SetAnchorHref(void* ptr, char* value)
{
	static_cast<QTextCharFormat*>(ptr)->setAnchorHref(QString(value));
}

void QTextCharFormat_SetAnchorNames(void* ptr, char* names)
{
	static_cast<QTextCharFormat*>(ptr)->setAnchorNames(QString(names).split("|", QString::SkipEmptyParts));
}

void QTextCharFormat_SetFont2(void* ptr, void* font)
{
	static_cast<QTextCharFormat*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QTextCharFormat_SetFont(void* ptr, void* font, int behavior)
{
	static_cast<QTextCharFormat*>(ptr)->setFont(*static_cast<QFont*>(font), static_cast<QTextCharFormat::FontPropertiesInheritanceBehavior>(behavior));
}

void QTextCharFormat_SetFontCapitalization(void* ptr, int capitalization)
{
	static_cast<QTextCharFormat*>(ptr)->setFontCapitalization(static_cast<QFont::Capitalization>(capitalization));
}

void QTextCharFormat_SetFontFamily(void* ptr, char* family)
{
	static_cast<QTextCharFormat*>(ptr)->setFontFamily(QString(family));
}

void QTextCharFormat_SetFontFixedPitch(void* ptr, int fixedPitch)
{
	static_cast<QTextCharFormat*>(ptr)->setFontFixedPitch(fixedPitch != 0);
}

void QTextCharFormat_SetFontHintingPreference(void* ptr, int hintingPreference)
{
	static_cast<QTextCharFormat*>(ptr)->setFontHintingPreference(static_cast<QFont::HintingPreference>(hintingPreference));
}

void QTextCharFormat_SetFontItalic(void* ptr, int italic)
{
	static_cast<QTextCharFormat*>(ptr)->setFontItalic(italic != 0);
}

void QTextCharFormat_SetFontKerning(void* ptr, int enable)
{
	static_cast<QTextCharFormat*>(ptr)->setFontKerning(enable != 0);
}

void QTextCharFormat_SetFontLetterSpacing(void* ptr, double spacing)
{
	static_cast<QTextCharFormat*>(ptr)->setFontLetterSpacing(static_cast<double>(spacing));
}

void QTextCharFormat_SetFontLetterSpacingType(void* ptr, int letterSpacingType)
{
	static_cast<QTextCharFormat*>(ptr)->setFontLetterSpacingType(static_cast<QFont::SpacingType>(letterSpacingType));
}

void QTextCharFormat_SetFontOverline(void* ptr, int overline)
{
	static_cast<QTextCharFormat*>(ptr)->setFontOverline(overline != 0);
}

void QTextCharFormat_SetFontPointSize(void* ptr, double size)
{
	static_cast<QTextCharFormat*>(ptr)->setFontPointSize(static_cast<double>(size));
}

void QTextCharFormat_SetFontStretch(void* ptr, int factor)
{
	static_cast<QTextCharFormat*>(ptr)->setFontStretch(factor);
}

void QTextCharFormat_SetFontStrikeOut(void* ptr, int strikeOut)
{
	static_cast<QTextCharFormat*>(ptr)->setFontStrikeOut(strikeOut != 0);
}

void QTextCharFormat_SetFontStyleHint(void* ptr, int hint, int strategy)
{
	static_cast<QTextCharFormat*>(ptr)->setFontStyleHint(static_cast<QFont::StyleHint>(hint), static_cast<QFont::StyleStrategy>(strategy));
}

void QTextCharFormat_SetFontStyleStrategy(void* ptr, int strategy)
{
	static_cast<QTextCharFormat*>(ptr)->setFontStyleStrategy(static_cast<QFont::StyleStrategy>(strategy));
}

void QTextCharFormat_SetFontUnderline(void* ptr, int underline)
{
	static_cast<QTextCharFormat*>(ptr)->setFontUnderline(underline != 0);
}

void QTextCharFormat_SetFontWeight(void* ptr, int weight)
{
	static_cast<QTextCharFormat*>(ptr)->setFontWeight(weight);
}

void QTextCharFormat_SetFontWordSpacing(void* ptr, double spacing)
{
	static_cast<QTextCharFormat*>(ptr)->setFontWordSpacing(static_cast<double>(spacing));
}

void QTextCharFormat_SetTextOutline(void* ptr, void* pen)
{
	static_cast<QTextCharFormat*>(ptr)->setTextOutline(*static_cast<QPen*>(pen));
}

void QTextCharFormat_SetToolTip(void* ptr, char* text)
{
	static_cast<QTextCharFormat*>(ptr)->setToolTip(QString(text));
}

void QTextCharFormat_SetUnderlineColor(void* ptr, void* color)
{
	static_cast<QTextCharFormat*>(ptr)->setUnderlineColor(*static_cast<QColor*>(color));
}

void QTextCharFormat_SetVerticalAlignment(void* ptr, int alignment)
{
	static_cast<QTextCharFormat*>(ptr)->setVerticalAlignment(static_cast<QTextCharFormat::VerticalAlignment>(alignment));
}

void* QTextCharFormat_TextOutline(void* ptr)
{
	return new QPen(static_cast<QTextCharFormat*>(ptr)->textOutline());
}

char* QTextCharFormat_ToolTip(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->toolTip().toUtf8().data();
}

void* QTextCharFormat_UnderlineColor(void* ptr)
{
	return new QColor(static_cast<QTextCharFormat*>(ptr)->underlineColor());
}

int QTextCharFormat_UnderlineStyle(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->underlineStyle();
}

int QTextCharFormat_VerticalAlignment(void* ptr)
{
	return static_cast<QTextCharFormat*>(ptr)->verticalAlignment();
}

void QTextCursor_InsertBlock3(void* ptr, void* format, void* charFormat)
{
	static_cast<QTextCursor*>(ptr)->insertBlock(*static_cast<QTextBlockFormat*>(format), *static_cast<QTextCharFormat*>(charFormat));
}

void* QTextCursor_InsertTable2(void* ptr, int rows, int columns)
{
	return static_cast<QTextCursor*>(ptr)->insertTable(rows, columns);
}

void* QTextCursor_InsertTable(void* ptr, int rows, int columns, void* format)
{
	return static_cast<QTextCursor*>(ptr)->insertTable(rows, columns, *static_cast<QTextTableFormat*>(format));
}

void QTextCursor_InsertText2(void* ptr, char* text, void* format)
{
	static_cast<QTextCursor*>(ptr)->insertText(QString(text), *static_cast<QTextCharFormat*>(format));
}

int QTextCursor_MovePosition(void* ptr, int operation, int mode, int n)
{
	return static_cast<QTextCursor*>(ptr)->movePosition(static_cast<QTextCursor::MoveOperation>(operation), static_cast<QTextCursor::MoveMode>(mode), n);
}

void* QTextCursor_NewQTextCursor()
{
	return new QTextCursor();
}

void* QTextCursor_NewQTextCursor2(void* document)
{
	return new QTextCursor(static_cast<QTextDocument*>(document));
}

void* QTextCursor_NewQTextCursor4(void* frame)
{
	return new QTextCursor(static_cast<QTextFrame*>(frame));
}

void* QTextCursor_NewQTextCursor5(void* block)
{
	return new QTextCursor(*static_cast<QTextBlock*>(block));
}

void* QTextCursor_NewQTextCursor7(void* cursor)
{
	return new QTextCursor(*static_cast<QTextCursor*>(cursor));
}

int QTextCursor_Anchor(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->anchor();
}

int QTextCursor_AtBlockEnd(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->atBlockEnd();
}

int QTextCursor_AtBlockStart(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->atBlockStart();
}

int QTextCursor_AtEnd(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->atEnd();
}

int QTextCursor_AtStart(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->atStart();
}

void QTextCursor_BeginEditBlock(void* ptr)
{
	static_cast<QTextCursor*>(ptr)->beginEditBlock();
}

void* QTextCursor_Block(void* ptr)
{
	return new QTextBlock(static_cast<QTextCursor*>(ptr)->block());
}

int QTextCursor_BlockNumber(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->blockNumber();
}

void QTextCursor_ClearSelection(void* ptr)
{
	static_cast<QTextCursor*>(ptr)->clearSelection();
}

int QTextCursor_ColumnNumber(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->columnNumber();
}

void* QTextCursor_CreateList2(void* ptr, int style)
{
	return static_cast<QTextCursor*>(ptr)->createList(static_cast<QTextListFormat::Style>(style));
}

void* QTextCursor_CreateList(void* ptr, void* format)
{
	return static_cast<QTextCursor*>(ptr)->createList(*static_cast<QTextListFormat*>(format));
}

void* QTextCursor_CurrentFrame(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->currentFrame();
}

void* QTextCursor_CurrentList(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->currentList();
}

void* QTextCursor_CurrentTable(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->currentTable();
}

void QTextCursor_DeleteChar(void* ptr)
{
	static_cast<QTextCursor*>(ptr)->deleteChar();
}

void QTextCursor_DeletePreviousChar(void* ptr)
{
	static_cast<QTextCursor*>(ptr)->deletePreviousChar();
}

void* QTextCursor_Document(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->document();
}

void QTextCursor_EndEditBlock(void* ptr)
{
	static_cast<QTextCursor*>(ptr)->endEditBlock();
}

int QTextCursor_HasComplexSelection(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->hasComplexSelection();
}

int QTextCursor_HasSelection(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->hasSelection();
}

void QTextCursor_InsertBlock(void* ptr)
{
	static_cast<QTextCursor*>(ptr)->insertBlock();
}

void QTextCursor_InsertBlock2(void* ptr, void* format)
{
	static_cast<QTextCursor*>(ptr)->insertBlock(*static_cast<QTextBlockFormat*>(format));
}

void QTextCursor_InsertFragment(void* ptr, void* fragment)
{
	static_cast<QTextCursor*>(ptr)->insertFragment(*static_cast<QTextDocumentFragment*>(fragment));
}

void* QTextCursor_InsertFrame(void* ptr, void* format)
{
	return static_cast<QTextCursor*>(ptr)->insertFrame(*static_cast<QTextFrameFormat*>(format));
}

void QTextCursor_InsertHtml(void* ptr, char* html)
{
	static_cast<QTextCursor*>(ptr)->insertHtml(QString(html));
}

void QTextCursor_InsertImage4(void* ptr, void* image, char* name)
{
	static_cast<QTextCursor*>(ptr)->insertImage(*static_cast<QImage*>(image), QString(name));
}

void QTextCursor_InsertImage3(void* ptr, char* name)
{
	static_cast<QTextCursor*>(ptr)->insertImage(QString(name));
}

void QTextCursor_InsertImage(void* ptr, void* format)
{
	static_cast<QTextCursor*>(ptr)->insertImage(*static_cast<QTextImageFormat*>(format));
}

void QTextCursor_InsertImage2(void* ptr, void* format, int alignment)
{
	static_cast<QTextCursor*>(ptr)->insertImage(*static_cast<QTextImageFormat*>(format), static_cast<QTextFrameFormat::Position>(alignment));
}

void* QTextCursor_InsertList2(void* ptr, int style)
{
	return static_cast<QTextCursor*>(ptr)->insertList(static_cast<QTextListFormat::Style>(style));
}

void* QTextCursor_InsertList(void* ptr, void* format)
{
	return static_cast<QTextCursor*>(ptr)->insertList(*static_cast<QTextListFormat*>(format));
}

void QTextCursor_InsertText(void* ptr, char* text)
{
	static_cast<QTextCursor*>(ptr)->insertText(QString(text));
}

int QTextCursor_IsCopyOf(void* ptr, void* other)
{
	return static_cast<QTextCursor*>(ptr)->isCopyOf(*static_cast<QTextCursor*>(other));
}

int QTextCursor_IsNull(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->isNull();
}

void QTextCursor_JoinPreviousEditBlock(void* ptr)
{
	static_cast<QTextCursor*>(ptr)->joinPreviousEditBlock();
}

int QTextCursor_KeepPositionOnInsert(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->keepPositionOnInsert();
}

void QTextCursor_MergeBlockCharFormat(void* ptr, void* modifier)
{
	static_cast<QTextCursor*>(ptr)->mergeBlockCharFormat(*static_cast<QTextCharFormat*>(modifier));
}

void QTextCursor_MergeBlockFormat(void* ptr, void* modifier)
{
	static_cast<QTextCursor*>(ptr)->mergeBlockFormat(*static_cast<QTextBlockFormat*>(modifier));
}

void QTextCursor_MergeCharFormat(void* ptr, void* modifier)
{
	static_cast<QTextCursor*>(ptr)->mergeCharFormat(*static_cast<QTextCharFormat*>(modifier));
}

int QTextCursor_Position(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->position();
}

int QTextCursor_PositionInBlock(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->positionInBlock();
}

void QTextCursor_RemoveSelectedText(void* ptr)
{
	static_cast<QTextCursor*>(ptr)->removeSelectedText();
}

void QTextCursor_Select(void* ptr, int selection)
{
	static_cast<QTextCursor*>(ptr)->select(static_cast<QTextCursor::SelectionType>(selection));
}

void QTextCursor_SelectedTableCells(void* ptr, int firstRow, int numRows, int firstColumn, int numColumns)
{
	static_cast<QTextCursor*>(ptr)->selectedTableCells(&firstRow, &numRows, &firstColumn, &numColumns);
}

char* QTextCursor_SelectedText(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->selectedText().toUtf8().data();
}

void* QTextCursor_Selection(void* ptr)
{
	return new QTextDocumentFragment(static_cast<QTextCursor*>(ptr)->selection());
}

int QTextCursor_SelectionEnd(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->selectionEnd();
}

int QTextCursor_SelectionStart(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->selectionStart();
}

void QTextCursor_SetBlockCharFormat(void* ptr, void* format)
{
	static_cast<QTextCursor*>(ptr)->setBlockCharFormat(*static_cast<QTextCharFormat*>(format));
}

void QTextCursor_SetBlockFormat(void* ptr, void* format)
{
	static_cast<QTextCursor*>(ptr)->setBlockFormat(*static_cast<QTextBlockFormat*>(format));
}

void QTextCursor_SetCharFormat(void* ptr, void* format)
{
	static_cast<QTextCursor*>(ptr)->setCharFormat(*static_cast<QTextCharFormat*>(format));
}

void QTextCursor_SetKeepPositionOnInsert(void* ptr, int b)
{
	static_cast<QTextCursor*>(ptr)->setKeepPositionOnInsert(b != 0);
}

void QTextCursor_SetPosition(void* ptr, int pos, int m)
{
	static_cast<QTextCursor*>(ptr)->setPosition(pos, static_cast<QTextCursor::MoveMode>(m));
}

void QTextCursor_SetVerticalMovementX(void* ptr, int x)
{
	static_cast<QTextCursor*>(ptr)->setVerticalMovementX(x);
}

void QTextCursor_SetVisualNavigation(void* ptr, int b)
{
	static_cast<QTextCursor*>(ptr)->setVisualNavigation(b != 0);
}

void QTextCursor_Swap(void* ptr, void* other)
{
	static_cast<QTextCursor*>(ptr)->swap(*static_cast<QTextCursor*>(other));
}

int QTextCursor_VerticalMovementX(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->verticalMovementX();
}

int QTextCursor_VisualNavigation(void* ptr)
{
	return static_cast<QTextCursor*>(ptr)->visualNavigation();
}

void QTextCursor_DestroyQTextCursor(void* ptr)
{
	static_cast<QTextCursor*>(ptr)->~QTextCursor();
}

class MyQTextDocument: public QTextDocument
{
public:
	MyQTextDocument(QObject *parent) : QTextDocument(parent) {};
	MyQTextDocument(const QString &text, QObject *parent) : QTextDocument(text, parent) {};
	QTextObject * createObject(const QTextFormat & format) { return static_cast<QTextObject*>(callbackQTextDocument_CreateObject(this, this->objectName().toUtf8().data(), new QTextFormat(format))); };
	void setModified(bool m) { callbackQTextDocument_SetModified(this, this->objectName().toUtf8().data(), m); };
	void Signal_BaseUrlChanged(const QUrl & url) { callbackQTextDocument_BaseUrlChanged(this, this->objectName().toUtf8().data(), new QUrl(url)); };
	void Signal_BlockCountChanged(int newBlockCount) { callbackQTextDocument_BlockCountChanged(this, this->objectName().toUtf8().data(), newBlockCount); };
	void clear() { callbackQTextDocument_Clear(this, this->objectName().toUtf8().data()); };
	void Signal_ContentsChange(int position, int charsRemoved, int charsAdded) { callbackQTextDocument_ContentsChange(this, this->objectName().toUtf8().data(), position, charsRemoved, charsAdded); };
	void Signal_ContentsChanged() { callbackQTextDocument_ContentsChanged(this, this->objectName().toUtf8().data()); };
	void Signal_CursorPositionChanged(const QTextCursor & cursor) { callbackQTextDocument_CursorPositionChanged(this, this->objectName().toUtf8().data(), new QTextCursor(cursor)); };
	void Signal_DocumentLayoutChanged() { callbackQTextDocument_DocumentLayoutChanged(this, this->objectName().toUtf8().data()); };
	QVariant loadResource(int ty, const QUrl & name) { return *static_cast<QVariant*>(callbackQTextDocument_LoadResource(this, this->objectName().toUtf8().data(), ty, new QUrl(name))); };
	void Signal_ModificationChanged(bool changed) { callbackQTextDocument_ModificationChanged(this, this->objectName().toUtf8().data(), changed); };
	void redo() { callbackQTextDocument_Redo2(this, this->objectName().toUtf8().data()); };
	void Signal_RedoAvailable(bool available) { callbackQTextDocument_RedoAvailable(this, this->objectName().toUtf8().data(), available); };
	void undo() { callbackQTextDocument_Undo2(this, this->objectName().toUtf8().data()); };
	void Signal_UndoAvailable(bool available) { callbackQTextDocument_UndoAvailable(this, this->objectName().toUtf8().data(), available); };
	void Signal_UndoCommandAdded() { callbackQTextDocument_UndoCommandAdded(this, this->objectName().toUtf8().data()); };
	void timerEvent(QTimerEvent * event) { callbackQTextDocument_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQTextDocument_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQTextDocument_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQTextDocument_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQTextDocument_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQTextDocument_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQTextDocument_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQTextDocument_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQTextDocument_MetaObject(const_cast<MyQTextDocument*>(this), this->objectName().toUtf8().data())); };
};

void* QTextDocument_BaseUrl(void* ptr)
{
	return new QUrl(static_cast<QTextDocument*>(ptr)->baseUrl());
}

int QTextDocument_BlockCount(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->blockCount();
}

void* QTextDocument_CreateObject(void* ptr, void* format)
{
	return static_cast<QTextDocument*>(ptr)->createObject(*static_cast<QTextFormat*>(format));
}

void* QTextDocument_CreateObjectDefault(void* ptr, void* format)
{
	return static_cast<QTextDocument*>(ptr)->QTextDocument::createObject(*static_cast<QTextFormat*>(format));
}

char* QTextDocument_DefaultStyleSheet(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->defaultStyleSheet().toUtf8().data();
}

double QTextDocument_DocumentMargin(void* ptr)
{
	return static_cast<double>(static_cast<QTextDocument*>(ptr)->documentMargin());
}

void* QTextDocument_Find4(void* ptr, char* subString, int position, int options)
{
	return new QTextCursor(static_cast<QTextDocument*>(ptr)->find(QString(subString), position, static_cast<QTextDocument::FindFlag>(options)));
}

double QTextDocument_IndentWidth(void* ptr)
{
	return static_cast<double>(static_cast<QTextDocument*>(ptr)->indentWidth());
}

int QTextDocument_IsModified(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->isModified();
}

int QTextDocument_IsUndoRedoEnabled(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->isUndoRedoEnabled();
}

void QTextDocument_MarkContentsDirty(void* ptr, int position, int length)
{
	static_cast<QTextDocument*>(ptr)->markContentsDirty(position, length);
}

int QTextDocument_MaximumBlockCount(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->maximumBlockCount();
}

void* QTextDocument_PageSize(void* ptr)
{
	return new QSizeF(static_cast<QSizeF>(static_cast<QTextDocument*>(ptr)->pageSize()).width(), static_cast<QSizeF>(static_cast<QTextDocument*>(ptr)->pageSize()).height());
}

void QTextDocument_SetBaseUrl(void* ptr, void* url)
{
	static_cast<QTextDocument*>(ptr)->setBaseUrl(*static_cast<QUrl*>(url));
}

void QTextDocument_SetDefaultStyleSheet(void* ptr, char* sheet)
{
	static_cast<QTextDocument*>(ptr)->setDefaultStyleSheet(QString(sheet));
}

void QTextDocument_SetDocumentMargin(void* ptr, double margin)
{
	static_cast<QTextDocument*>(ptr)->setDocumentMargin(static_cast<double>(margin));
}

void QTextDocument_SetMaximumBlockCount(void* ptr, int maximum)
{
	static_cast<QTextDocument*>(ptr)->setMaximumBlockCount(maximum);
}

void QTextDocument_SetModified(void* ptr, int m)
{
	QMetaObject::invokeMethod(static_cast<QTextDocument*>(ptr), "setModified", Q_ARG(bool, m != 0));
}

void QTextDocument_SetPageSize(void* ptr, void* size)
{
	static_cast<QTextDocument*>(ptr)->setPageSize(*static_cast<QSizeF*>(size));
}

void QTextDocument_SetTextWidth(void* ptr, double width)
{
	static_cast<QTextDocument*>(ptr)->setTextWidth(static_cast<double>(width));
}

void QTextDocument_SetUndoRedoEnabled(void* ptr, int enable)
{
	static_cast<QTextDocument*>(ptr)->setUndoRedoEnabled(enable != 0);
}

void QTextDocument_SetUseDesignMetrics(void* ptr, int b)
{
	static_cast<QTextDocument*>(ptr)->setUseDesignMetrics(b != 0);
}

void* QTextDocument_Size(void* ptr)
{
	return new QSizeF(static_cast<QSizeF>(static_cast<QTextDocument*>(ptr)->size()).width(), static_cast<QSizeF>(static_cast<QTextDocument*>(ptr)->size()).height());
}

double QTextDocument_TextWidth(void* ptr)
{
	return static_cast<double>(static_cast<QTextDocument*>(ptr)->textWidth());
}

int QTextDocument_UseDesignMetrics(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->useDesignMetrics();
}

void* QTextDocument_NewQTextDocument(void* parent)
{
	return new MyQTextDocument(static_cast<QObject*>(parent));
}

void* QTextDocument_NewQTextDocument2(char* text, void* parent)
{
	return new MyQTextDocument(QString(text), static_cast<QObject*>(parent));
}

void QTextDocument_AddResource(void* ptr, int ty, void* name, void* resource)
{
	static_cast<QTextDocument*>(ptr)->addResource(ty, *static_cast<QUrl*>(name), *static_cast<QVariant*>(resource));
}

void QTextDocument_AdjustSize(void* ptr)
{
	static_cast<QTextDocument*>(ptr)->adjustSize();
}

int QTextDocument_AvailableRedoSteps(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->availableRedoSteps();
}

int QTextDocument_AvailableUndoSteps(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->availableUndoSteps();
}

void QTextDocument_ConnectBaseUrlChanged(void* ptr)
{
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(const QUrl &)>(&QTextDocument::baseUrlChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(const QUrl &)>(&MyQTextDocument::Signal_BaseUrlChanged));
}

void QTextDocument_DisconnectBaseUrlChanged(void* ptr)
{
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(const QUrl &)>(&QTextDocument::baseUrlChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(const QUrl &)>(&MyQTextDocument::Signal_BaseUrlChanged));
}

void QTextDocument_BaseUrlChanged(void* ptr, void* url)
{
	static_cast<QTextDocument*>(ptr)->baseUrlChanged(*static_cast<QUrl*>(url));
}

void* QTextDocument_Begin(void* ptr)
{
	return new QTextBlock(static_cast<QTextDocument*>(ptr)->begin());
}

void QTextDocument_ConnectBlockCountChanged(void* ptr)
{
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(int)>(&QTextDocument::blockCountChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(int)>(&MyQTextDocument::Signal_BlockCountChanged));
}

void QTextDocument_DisconnectBlockCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(int)>(&QTextDocument::blockCountChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(int)>(&MyQTextDocument::Signal_BlockCountChanged));
}

void QTextDocument_BlockCountChanged(void* ptr, int newBlockCount)
{
	static_cast<QTextDocument*>(ptr)->blockCountChanged(newBlockCount);
}

int QTextDocument_CharacterCount(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->characterCount();
}

void QTextDocument_Clear(void* ptr)
{
	static_cast<QTextDocument*>(ptr)->clear();
}

void QTextDocument_ClearDefault(void* ptr)
{
	static_cast<QTextDocument*>(ptr)->QTextDocument::clear();
}

void QTextDocument_ClearUndoRedoStacks(void* ptr, int stacksToClear)
{
	static_cast<QTextDocument*>(ptr)->clearUndoRedoStacks(static_cast<QTextDocument::Stacks>(stacksToClear));
}

void* QTextDocument_Clone(void* ptr, void* parent)
{
	return static_cast<QTextDocument*>(ptr)->clone(static_cast<QObject*>(parent));
}

void QTextDocument_ConnectContentsChange(void* ptr)
{
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(int, int, int)>(&QTextDocument::contentsChange), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(int, int, int)>(&MyQTextDocument::Signal_ContentsChange));
}

void QTextDocument_DisconnectContentsChange(void* ptr)
{
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(int, int, int)>(&QTextDocument::contentsChange), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(int, int, int)>(&MyQTextDocument::Signal_ContentsChange));
}

void QTextDocument_ContentsChange(void* ptr, int position, int charsRemoved, int charsAdded)
{
	static_cast<QTextDocument*>(ptr)->contentsChange(position, charsRemoved, charsAdded);
}

void QTextDocument_ConnectContentsChanged(void* ptr)
{
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)()>(&QTextDocument::contentsChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)()>(&MyQTextDocument::Signal_ContentsChanged));
}

void QTextDocument_DisconnectContentsChanged(void* ptr)
{
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)()>(&QTextDocument::contentsChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)()>(&MyQTextDocument::Signal_ContentsChanged));
}

void QTextDocument_ContentsChanged(void* ptr)
{
	static_cast<QTextDocument*>(ptr)->contentsChanged();
}

void QTextDocument_ConnectCursorPositionChanged(void* ptr)
{
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(const QTextCursor &)>(&QTextDocument::cursorPositionChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(const QTextCursor &)>(&MyQTextDocument::Signal_CursorPositionChanged));
}

void QTextDocument_DisconnectCursorPositionChanged(void* ptr)
{
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(const QTextCursor &)>(&QTextDocument::cursorPositionChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(const QTextCursor &)>(&MyQTextDocument::Signal_CursorPositionChanged));
}

void QTextDocument_CursorPositionChanged(void* ptr, void* cursor)
{
	static_cast<QTextDocument*>(ptr)->cursorPositionChanged(*static_cast<QTextCursor*>(cursor));
}

int QTextDocument_DefaultCursorMoveStyle(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->defaultCursorMoveStyle();
}

void* QTextDocument_DefaultFont(void* ptr)
{
	return new QFont(static_cast<QTextDocument*>(ptr)->defaultFont());
}

void* QTextDocument_DefaultTextOption(void* ptr)
{
	return new QTextOption(static_cast<QTextDocument*>(ptr)->defaultTextOption());
}

void* QTextDocument_DocumentLayout(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->documentLayout();
}

void QTextDocument_ConnectDocumentLayoutChanged(void* ptr)
{
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)()>(&QTextDocument::documentLayoutChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)()>(&MyQTextDocument::Signal_DocumentLayoutChanged));
}

void QTextDocument_DisconnectDocumentLayoutChanged(void* ptr)
{
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)()>(&QTextDocument::documentLayoutChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)()>(&MyQTextDocument::Signal_DocumentLayoutChanged));
}

void QTextDocument_DocumentLayoutChanged(void* ptr)
{
	static_cast<QTextDocument*>(ptr)->documentLayoutChanged();
}

void QTextDocument_DrawContents(void* ptr, void* p, void* rect)
{
	static_cast<QTextDocument*>(ptr)->drawContents(static_cast<QPainter*>(p), *static_cast<QRectF*>(rect));
}

void* QTextDocument_End(void* ptr)
{
	return new QTextBlock(static_cast<QTextDocument*>(ptr)->end());
}

void* QTextDocument_Find6(void* ptr, void* expr, void* cursor, int options)
{
	return new QTextCursor(static_cast<QTextDocument*>(ptr)->find(*static_cast<QRegExp*>(expr), *static_cast<QTextCursor*>(cursor), static_cast<QTextDocument::FindFlag>(options)));
}

void* QTextDocument_Find5(void* ptr, void* expr, int from, int options)
{
	return new QTextCursor(static_cast<QTextDocument*>(ptr)->find(*static_cast<QRegExp*>(expr), from, static_cast<QTextDocument::FindFlag>(options)));
}

void* QTextDocument_Find3(void* ptr, void* expr, void* cursor, int options)
{
	return new QTextCursor(static_cast<QTextDocument*>(ptr)->find(*static_cast<QRegularExpression*>(expr), *static_cast<QTextCursor*>(cursor), static_cast<QTextDocument::FindFlag>(options)));
}

void* QTextDocument_Find2(void* ptr, void* expr, int from, int options)
{
	return new QTextCursor(static_cast<QTextDocument*>(ptr)->find(*static_cast<QRegularExpression*>(expr), from, static_cast<QTextDocument::FindFlag>(options)));
}

void* QTextDocument_Find(void* ptr, char* subString, void* cursor, int options)
{
	return new QTextCursor(static_cast<QTextDocument*>(ptr)->find(QString(subString), *static_cast<QTextCursor*>(cursor), static_cast<QTextDocument::FindFlag>(options)));
}

void* QTextDocument_FindBlock(void* ptr, int pos)
{
	return new QTextBlock(static_cast<QTextDocument*>(ptr)->findBlock(pos));
}

void* QTextDocument_FindBlockByLineNumber(void* ptr, int lineNumber)
{
	return new QTextBlock(static_cast<QTextDocument*>(ptr)->findBlockByLineNumber(lineNumber));
}

void* QTextDocument_FindBlockByNumber(void* ptr, int blockNumber)
{
	return new QTextBlock(static_cast<QTextDocument*>(ptr)->findBlockByNumber(blockNumber));
}

void* QTextDocument_FirstBlock(void* ptr)
{
	return new QTextBlock(static_cast<QTextDocument*>(ptr)->firstBlock());
}

double QTextDocument_IdealWidth(void* ptr)
{
	return static_cast<double>(static_cast<QTextDocument*>(ptr)->idealWidth());
}

int QTextDocument_IsEmpty(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->isEmpty();
}

int QTextDocument_IsRedoAvailable(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->isRedoAvailable();
}

int QTextDocument_IsUndoAvailable(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->isUndoAvailable();
}

void* QTextDocument_LastBlock(void* ptr)
{
	return new QTextBlock(static_cast<QTextDocument*>(ptr)->lastBlock());
}

int QTextDocument_LineCount(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->lineCount();
}

void* QTextDocument_LoadResource(void* ptr, int ty, void* name)
{
	return new QVariant(static_cast<QTextDocument*>(ptr)->loadResource(ty, *static_cast<QUrl*>(name)));
}

void* QTextDocument_LoadResourceDefault(void* ptr, int ty, void* name)
{
	return new QVariant(static_cast<QTextDocument*>(ptr)->QTextDocument::loadResource(ty, *static_cast<QUrl*>(name)));
}

char* QTextDocument_MetaInformation(void* ptr, int info)
{
	return static_cast<QTextDocument*>(ptr)->metaInformation(static_cast<QTextDocument::MetaInformation>(info)).toUtf8().data();
}

void QTextDocument_ConnectModificationChanged(void* ptr)
{
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(bool)>(&QTextDocument::modificationChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(bool)>(&MyQTextDocument::Signal_ModificationChanged));
}

void QTextDocument_DisconnectModificationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(bool)>(&QTextDocument::modificationChanged), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(bool)>(&MyQTextDocument::Signal_ModificationChanged));
}

void QTextDocument_ModificationChanged(void* ptr, int changed)
{
	static_cast<QTextDocument*>(ptr)->modificationChanged(changed != 0);
}

void* QTextDocument_Object(void* ptr, int objectIndex)
{
	return static_cast<QTextDocument*>(ptr)->object(objectIndex);
}

void* QTextDocument_ObjectForFormat(void* ptr, void* f)
{
	return static_cast<QTextDocument*>(ptr)->objectForFormat(*static_cast<QTextFormat*>(f));
}

int QTextDocument_PageCount(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->pageCount();
}

void QTextDocument_Print(void* ptr, void* printer)
{
	static_cast<QTextDocument*>(ptr)->print(static_cast<QPagedPaintDevice*>(printer));
}

void QTextDocument_Redo2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTextDocument*>(ptr), "redo");
}

void QTextDocument_Redo(void* ptr, void* cursor)
{
	static_cast<QTextDocument*>(ptr)->redo(static_cast<QTextCursor*>(cursor));
}

void QTextDocument_ConnectRedoAvailable(void* ptr)
{
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(bool)>(&QTextDocument::redoAvailable), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(bool)>(&MyQTextDocument::Signal_RedoAvailable));
}

void QTextDocument_DisconnectRedoAvailable(void* ptr)
{
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(bool)>(&QTextDocument::redoAvailable), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(bool)>(&MyQTextDocument::Signal_RedoAvailable));
}

void QTextDocument_RedoAvailable(void* ptr, int available)
{
	static_cast<QTextDocument*>(ptr)->redoAvailable(available != 0);
}

void* QTextDocument_Resource(void* ptr, int ty, void* name)
{
	return new QVariant(static_cast<QTextDocument*>(ptr)->resource(ty, *static_cast<QUrl*>(name)));
}

int QTextDocument_Revision(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->revision();
}

void* QTextDocument_RootFrame(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->rootFrame();
}

void QTextDocument_SetDefaultCursorMoveStyle(void* ptr, int style)
{
	static_cast<QTextDocument*>(ptr)->setDefaultCursorMoveStyle(static_cast<Qt::CursorMoveStyle>(style));
}

void QTextDocument_SetDefaultFont(void* ptr, void* font)
{
	static_cast<QTextDocument*>(ptr)->setDefaultFont(*static_cast<QFont*>(font));
}

void QTextDocument_SetDefaultTextOption(void* ptr, void* option)
{
	static_cast<QTextDocument*>(ptr)->setDefaultTextOption(*static_cast<QTextOption*>(option));
}

void QTextDocument_SetDocumentLayout(void* ptr, void* layout)
{
	static_cast<QTextDocument*>(ptr)->setDocumentLayout(static_cast<QAbstractTextDocumentLayout*>(layout));
}

void QTextDocument_SetHtml(void* ptr, char* html)
{
	static_cast<QTextDocument*>(ptr)->setHtml(QString(html));
}

void QTextDocument_SetIndentWidth(void* ptr, double width)
{
	static_cast<QTextDocument*>(ptr)->setIndentWidth(static_cast<double>(width));
}

void QTextDocument_SetMetaInformation(void* ptr, int info, char* stri)
{
	static_cast<QTextDocument*>(ptr)->setMetaInformation(static_cast<QTextDocument::MetaInformation>(info), QString(stri));
}

void QTextDocument_SetPlainText(void* ptr, char* text)
{
	static_cast<QTextDocument*>(ptr)->setPlainText(QString(text));
}

char* QTextDocument_ToHtml(void* ptr, char* encoding)
{
	return static_cast<QTextDocument*>(ptr)->toHtml(QByteArray(encoding)).toUtf8().data();
}

char* QTextDocument_ToPlainText(void* ptr)
{
	return static_cast<QTextDocument*>(ptr)->toPlainText().toUtf8().data();
}

void QTextDocument_Undo2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTextDocument*>(ptr), "undo");
}

void QTextDocument_Undo(void* ptr, void* cursor)
{
	static_cast<QTextDocument*>(ptr)->undo(static_cast<QTextCursor*>(cursor));
}

void QTextDocument_ConnectUndoAvailable(void* ptr)
{
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(bool)>(&QTextDocument::undoAvailable), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(bool)>(&MyQTextDocument::Signal_UndoAvailable));
}

void QTextDocument_DisconnectUndoAvailable(void* ptr)
{
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)(bool)>(&QTextDocument::undoAvailable), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)(bool)>(&MyQTextDocument::Signal_UndoAvailable));
}

void QTextDocument_UndoAvailable(void* ptr, int available)
{
	static_cast<QTextDocument*>(ptr)->undoAvailable(available != 0);
}

void QTextDocument_ConnectUndoCommandAdded(void* ptr)
{
	QObject::connect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)()>(&QTextDocument::undoCommandAdded), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)()>(&MyQTextDocument::Signal_UndoCommandAdded));
}

void QTextDocument_DisconnectUndoCommandAdded(void* ptr)
{
	QObject::disconnect(static_cast<QTextDocument*>(ptr), static_cast<void (QTextDocument::*)()>(&QTextDocument::undoCommandAdded), static_cast<MyQTextDocument*>(ptr), static_cast<void (MyQTextDocument::*)()>(&MyQTextDocument::Signal_UndoCommandAdded));
}

void QTextDocument_UndoCommandAdded(void* ptr)
{
	static_cast<QTextDocument*>(ptr)->undoCommandAdded();
}

void QTextDocument_DestroyQTextDocument(void* ptr)
{
	static_cast<QTextDocument*>(ptr)->~QTextDocument();
}

void QTextDocument_TimerEvent(void* ptr, void* event)
{
	static_cast<QTextDocument*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTextDocument_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QTextDocument*>(ptr)->QTextDocument::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTextDocument_ChildEvent(void* ptr, void* event)
{
	static_cast<QTextDocument*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTextDocument_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QTextDocument*>(ptr)->QTextDocument::childEvent(static_cast<QChildEvent*>(event));
}

void QTextDocument_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QTextDocument*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextDocument_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTextDocument*>(ptr)->QTextDocument::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextDocument_CustomEvent(void* ptr, void* event)
{
	static_cast<QTextDocument*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTextDocument_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QTextDocument*>(ptr)->QTextDocument::customEvent(static_cast<QEvent*>(event));
}

void QTextDocument_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTextDocument*>(ptr), "deleteLater");
}

void QTextDocument_DeleteLaterDefault(void* ptr)
{
	static_cast<QTextDocument*>(ptr)->QTextDocument::deleteLater();
}

void QTextDocument_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QTextDocument*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextDocument_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTextDocument*>(ptr)->QTextDocument::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QTextDocument_Event(void* ptr, void* e)
{
	return static_cast<QTextDocument*>(ptr)->event(static_cast<QEvent*>(e));
}

int QTextDocument_EventDefault(void* ptr, void* e)
{
	return static_cast<QTextDocument*>(ptr)->QTextDocument::event(static_cast<QEvent*>(e));
}

int QTextDocument_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QTextDocument*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QTextDocument_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QTextDocument*>(ptr)->QTextDocument::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QTextDocument_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTextDocument*>(ptr)->metaObject());
}

void* QTextDocument_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTextDocument*>(ptr)->QTextDocument::metaObject());
}

void* QTextDocumentFragment_NewQTextDocumentFragment4(void* other)
{
	return new QTextDocumentFragment(*static_cast<QTextDocumentFragment*>(other));
}

void* QTextDocumentFragment_QTextDocumentFragment_FromHtml(char* text)
{
	return new QTextDocumentFragment(QTextDocumentFragment::fromHtml(QString(text)));
}

void* QTextDocumentFragment_QTextDocumentFragment_FromHtml2(char* text, void* resourceProvider)
{
	return new QTextDocumentFragment(QTextDocumentFragment::fromHtml(QString(text), static_cast<QTextDocument*>(resourceProvider)));
}

void* QTextDocumentFragment_NewQTextDocumentFragment()
{
	return new QTextDocumentFragment();
}

void* QTextDocumentFragment_NewQTextDocumentFragment3(void* cursor)
{
	return new QTextDocumentFragment(*static_cast<QTextCursor*>(cursor));
}

void* QTextDocumentFragment_NewQTextDocumentFragment2(void* document)
{
	return new QTextDocumentFragment(static_cast<QTextDocument*>(document));
}

void* QTextDocumentFragment_QTextDocumentFragment_FromPlainText(char* plainText)
{
	return new QTextDocumentFragment(QTextDocumentFragment::fromPlainText(QString(plainText)));
}

int QTextDocumentFragment_IsEmpty(void* ptr)
{
	return static_cast<QTextDocumentFragment*>(ptr)->isEmpty();
}

char* QTextDocumentFragment_ToHtml(void* ptr, char* encoding)
{
	return static_cast<QTextDocumentFragment*>(ptr)->toHtml(QByteArray(encoding)).toUtf8().data();
}

char* QTextDocumentFragment_ToPlainText(void* ptr)
{
	return static_cast<QTextDocumentFragment*>(ptr)->toPlainText().toUtf8().data();
}

void QTextDocumentFragment_DestroyQTextDocumentFragment(void* ptr)
{
	static_cast<QTextDocumentFragment*>(ptr)->~QTextDocumentFragment();
}

void* QTextDocumentWriter_NewQTextDocumentWriter()
{
	return new QTextDocumentWriter();
}

void* QTextDocumentWriter_NewQTextDocumentWriter2(void* device, char* format)
{
	return new QTextDocumentWriter(static_cast<QIODevice*>(device), QByteArray(format));
}

void* QTextDocumentWriter_NewQTextDocumentWriter3(char* fileName, char* format)
{
	return new QTextDocumentWriter(QString(fileName), QByteArray(format));
}

void* QTextDocumentWriter_Codec(void* ptr)
{
	return static_cast<QTextDocumentWriter*>(ptr)->codec();
}

void* QTextDocumentWriter_Device(void* ptr)
{
	return static_cast<QTextDocumentWriter*>(ptr)->device();
}

char* QTextDocumentWriter_FileName(void* ptr)
{
	return static_cast<QTextDocumentWriter*>(ptr)->fileName().toUtf8().data();
}

char* QTextDocumentWriter_Format(void* ptr)
{
	return QString(static_cast<QTextDocumentWriter*>(ptr)->format()).toUtf8().data();
}

void QTextDocumentWriter_SetCodec(void* ptr, void* codec)
{
	static_cast<QTextDocumentWriter*>(ptr)->setCodec(static_cast<QTextCodec*>(codec));
}

void QTextDocumentWriter_SetDevice(void* ptr, void* device)
{
	static_cast<QTextDocumentWriter*>(ptr)->setDevice(static_cast<QIODevice*>(device));
}

void QTextDocumentWriter_SetFileName(void* ptr, char* fileName)
{
	static_cast<QTextDocumentWriter*>(ptr)->setFileName(QString(fileName));
}

void QTextDocumentWriter_SetFormat(void* ptr, char* format)
{
	static_cast<QTextDocumentWriter*>(ptr)->setFormat(QByteArray(format));
}

int QTextDocumentWriter_Write(void* ptr, void* document)
{
	return static_cast<QTextDocumentWriter*>(ptr)->write(static_cast<QTextDocument*>(document));
}

int QTextDocumentWriter_Write2(void* ptr, void* fragment)
{
	return static_cast<QTextDocumentWriter*>(ptr)->write(*static_cast<QTextDocumentFragment*>(fragment));
}

void QTextDocumentWriter_DestroyQTextDocumentWriter(void* ptr)
{
	static_cast<QTextDocumentWriter*>(ptr)->~QTextDocumentWriter();
}

void* QTextFormat_NewQTextFormat3(void* other)
{
	return new QTextFormat(*static_cast<QTextFormat*>(other));
}

void QTextFormat_SetObjectIndex(void* ptr, int index)
{
	static_cast<QTextFormat*>(ptr)->setObjectIndex(index);
}

void* QTextFormat_NewQTextFormat()
{
	return new QTextFormat();
}

void* QTextFormat_NewQTextFormat2(int ty)
{
	return new QTextFormat(ty);
}

void* QTextFormat_Background(void* ptr)
{
	return new QBrush(static_cast<QTextFormat*>(ptr)->background());
}

int QTextFormat_BoolProperty(void* ptr, int propertyId)
{
	return static_cast<QTextFormat*>(ptr)->boolProperty(propertyId);
}

void* QTextFormat_BrushProperty(void* ptr, int propertyId)
{
	return new QBrush(static_cast<QTextFormat*>(ptr)->brushProperty(propertyId));
}

void QTextFormat_ClearBackground(void* ptr)
{
	static_cast<QTextFormat*>(ptr)->clearBackground();
}

void QTextFormat_ClearForeground(void* ptr)
{
	static_cast<QTextFormat*>(ptr)->clearForeground();
}

void QTextFormat_ClearProperty(void* ptr, int propertyId)
{
	static_cast<QTextFormat*>(ptr)->clearProperty(propertyId);
}

void* QTextFormat_ColorProperty(void* ptr, int propertyId)
{
	return new QColor(static_cast<QTextFormat*>(ptr)->colorProperty(propertyId));
}

double QTextFormat_DoubleProperty(void* ptr, int propertyId)
{
	return static_cast<double>(static_cast<QTextFormat*>(ptr)->doubleProperty(propertyId));
}

void* QTextFormat_Foreground(void* ptr)
{
	return new QBrush(static_cast<QTextFormat*>(ptr)->foreground());
}

int QTextFormat_HasProperty(void* ptr, int propertyId)
{
	return static_cast<QTextFormat*>(ptr)->hasProperty(propertyId);
}

int QTextFormat_IntProperty(void* ptr, int propertyId)
{
	return static_cast<QTextFormat*>(ptr)->intProperty(propertyId);
}

int QTextFormat_IsBlockFormat(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->isBlockFormat();
}

int QTextFormat_IsCharFormat(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->isCharFormat();
}

int QTextFormat_IsEmpty(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->isEmpty();
}

int QTextFormat_IsFrameFormat(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->isFrameFormat();
}

int QTextFormat_IsImageFormat(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->isImageFormat();
}

int QTextFormat_IsListFormat(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->isListFormat();
}

int QTextFormat_IsTableCellFormat(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->isTableCellFormat();
}

int QTextFormat_IsTableFormat(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->isTableFormat();
}

int QTextFormat_IsValid(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->isValid();
}

int QTextFormat_LayoutDirection(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->layoutDirection();
}

void QTextFormat_Merge(void* ptr, void* other)
{
	static_cast<QTextFormat*>(ptr)->merge(*static_cast<QTextFormat*>(other));
}

int QTextFormat_ObjectIndex(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->objectIndex();
}

int QTextFormat_ObjectType(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->objectType();
}

void* QTextFormat_PenProperty(void* ptr, int propertyId)
{
	return new QPen(static_cast<QTextFormat*>(ptr)->penProperty(propertyId));
}

void* QTextFormat_Property(void* ptr, int propertyId)
{
	return new QVariant(static_cast<QTextFormat*>(ptr)->property(propertyId));
}

int QTextFormat_PropertyCount(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->propertyCount();
}

void QTextFormat_SetBackground(void* ptr, void* brush)
{
	static_cast<QTextFormat*>(ptr)->setBackground(*static_cast<QBrush*>(brush));
}

void QTextFormat_SetForeground(void* ptr, void* brush)
{
	static_cast<QTextFormat*>(ptr)->setForeground(*static_cast<QBrush*>(brush));
}

void QTextFormat_SetLayoutDirection(void* ptr, int direction)
{
	static_cast<QTextFormat*>(ptr)->setLayoutDirection(static_cast<Qt::LayoutDirection>(direction));
}

void QTextFormat_SetObjectType(void* ptr, int ty)
{
	static_cast<QTextFormat*>(ptr)->setObjectType(ty);
}

void QTextFormat_SetProperty(void* ptr, int propertyId, void* value)
{
	static_cast<QTextFormat*>(ptr)->setProperty(propertyId, *static_cast<QVariant*>(value));
}

char* QTextFormat_StringProperty(void* ptr, int propertyId)
{
	return static_cast<QTextFormat*>(ptr)->stringProperty(propertyId).toUtf8().data();
}

void QTextFormat_Swap(void* ptr, void* other)
{
	static_cast<QTextFormat*>(ptr)->swap(*static_cast<QTextFormat*>(other));
}

int QTextFormat_Type(void* ptr)
{
	return static_cast<QTextFormat*>(ptr)->type();
}

void QTextFormat_DestroyQTextFormat(void* ptr)
{
	static_cast<QTextFormat*>(ptr)->~QTextFormat();
}

void* QTextFragment_NewQTextFragment()
{
	return new QTextFragment();
}

void* QTextFragment_NewQTextFragment3(void* other)
{
	return new QTextFragment(*static_cast<QTextFragment*>(other));
}

int QTextFragment_CharFormatIndex(void* ptr)
{
	return static_cast<QTextFragment*>(ptr)->charFormatIndex();
}

int QTextFragment_Contains(void* ptr, int position)
{
	return static_cast<QTextFragment*>(ptr)->contains(position);
}

int QTextFragment_IsValid(void* ptr)
{
	return static_cast<QTextFragment*>(ptr)->isValid();
}

int QTextFragment_Length(void* ptr)
{
	return static_cast<QTextFragment*>(ptr)->length();
}

int QTextFragment_Position(void* ptr)
{
	return static_cast<QTextFragment*>(ptr)->position();
}

char* QTextFragment_Text(void* ptr)
{
	return static_cast<QTextFragment*>(ptr)->text().toUtf8().data();
}

void* QTextFrame_NewQTextFrame(void* document)
{
	return new QTextFrame(static_cast<QTextDocument*>(document));
}

void* QTextFrame_FirstCursorPosition(void* ptr)
{
	return new QTextCursor(static_cast<QTextFrame*>(ptr)->firstCursorPosition());
}

int QTextFrame_FirstPosition(void* ptr)
{
	return static_cast<QTextFrame*>(ptr)->firstPosition();
}

void* QTextFrame_LastCursorPosition(void* ptr)
{
	return new QTextCursor(static_cast<QTextFrame*>(ptr)->lastCursorPosition());
}

int QTextFrame_LastPosition(void* ptr)
{
	return static_cast<QTextFrame*>(ptr)->lastPosition();
}

void* QTextFrame_ParentFrame(void* ptr)
{
	return static_cast<QTextFrame*>(ptr)->parentFrame();
}

void QTextFrame_SetFrameFormat(void* ptr, void* format)
{
	static_cast<QTextFrame*>(ptr)->setFrameFormat(*static_cast<QTextFrameFormat*>(format));
}

void QTextFrame_DestroyQTextFrame(void* ptr)
{
	static_cast<QTextFrame*>(ptr)->~QTextFrame();
}

void QTextFrame_TimerEvent(void* ptr, void* event)
{
	static_cast<QTextFrame*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTextFrame_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QTextFrame*>(ptr)->QTextFrame::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTextFrame_ChildEvent(void* ptr, void* event)
{
	static_cast<QTextFrame*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTextFrame_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QTextFrame*>(ptr)->QTextFrame::childEvent(static_cast<QChildEvent*>(event));
}

void QTextFrame_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QTextFrame*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextFrame_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTextFrame*>(ptr)->QTextFrame::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextFrame_CustomEvent(void* ptr, void* event)
{
	static_cast<QTextFrame*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTextFrame_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QTextFrame*>(ptr)->QTextFrame::customEvent(static_cast<QEvent*>(event));
}

void QTextFrame_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTextFrame*>(ptr), "deleteLater");
}

void QTextFrame_DeleteLaterDefault(void* ptr)
{
	static_cast<QTextFrame*>(ptr)->QTextFrame::deleteLater();
}

void QTextFrame_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QTextFrame*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextFrame_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTextFrame*>(ptr)->QTextFrame::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QTextFrame_Event(void* ptr, void* e)
{
	return static_cast<QTextFrame*>(ptr)->event(static_cast<QEvent*>(e));
}

int QTextFrame_EventDefault(void* ptr, void* e)
{
	return static_cast<QTextFrame*>(ptr)->QTextFrame::event(static_cast<QEvent*>(e));
}

int QTextFrame_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QTextFrame*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QTextFrame_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QTextFrame*>(ptr)->QTextFrame::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QTextFrame_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTextFrame*>(ptr)->metaObject());
}

void* QTextFrame_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTextFrame*>(ptr)->QTextFrame::metaObject());
}

void* QTextFrameFormat_NewQTextFrameFormat()
{
	return new QTextFrameFormat();
}

double QTextFrameFormat_BottomMargin(void* ptr)
{
	return static_cast<double>(static_cast<QTextFrameFormat*>(ptr)->bottomMargin());
}

double QTextFrameFormat_LeftMargin(void* ptr)
{
	return static_cast<double>(static_cast<QTextFrameFormat*>(ptr)->leftMargin());
}

double QTextFrameFormat_RightMargin(void* ptr)
{
	return static_cast<double>(static_cast<QTextFrameFormat*>(ptr)->rightMargin());
}

void QTextFrameFormat_SetMargin(void* ptr, double margin)
{
	static_cast<QTextFrameFormat*>(ptr)->setMargin(static_cast<double>(margin));
}

double QTextFrameFormat_TopMargin(void* ptr)
{
	return static_cast<double>(static_cast<QTextFrameFormat*>(ptr)->topMargin());
}

double QTextFrameFormat_Border(void* ptr)
{
	return static_cast<double>(static_cast<QTextFrameFormat*>(ptr)->border());
}

void* QTextFrameFormat_BorderBrush(void* ptr)
{
	return new QBrush(static_cast<QTextFrameFormat*>(ptr)->borderBrush());
}

int QTextFrameFormat_BorderStyle(void* ptr)
{
	return static_cast<QTextFrameFormat*>(ptr)->borderStyle();
}

int QTextFrameFormat_IsValid(void* ptr)
{
	return static_cast<QTextFrameFormat*>(ptr)->isValid();
}

double QTextFrameFormat_Margin(void* ptr)
{
	return static_cast<double>(static_cast<QTextFrameFormat*>(ptr)->margin());
}

double QTextFrameFormat_Padding(void* ptr)
{
	return static_cast<double>(static_cast<QTextFrameFormat*>(ptr)->padding());
}

int QTextFrameFormat_PageBreakPolicy(void* ptr)
{
	return static_cast<QTextFrameFormat*>(ptr)->pageBreakPolicy();
}

int QTextFrameFormat_Position(void* ptr)
{
	return static_cast<QTextFrameFormat*>(ptr)->position();
}

void QTextFrameFormat_SetBorder(void* ptr, double width)
{
	static_cast<QTextFrameFormat*>(ptr)->setBorder(static_cast<double>(width));
}

void QTextFrameFormat_SetBorderBrush(void* ptr, void* brush)
{
	static_cast<QTextFrameFormat*>(ptr)->setBorderBrush(*static_cast<QBrush*>(brush));
}

void QTextFrameFormat_SetBorderStyle(void* ptr, int style)
{
	static_cast<QTextFrameFormat*>(ptr)->setBorderStyle(static_cast<QTextFrameFormat::BorderStyle>(style));
}

void QTextFrameFormat_SetBottomMargin(void* ptr, double margin)
{
	static_cast<QTextFrameFormat*>(ptr)->setBottomMargin(static_cast<double>(margin));
}

void QTextFrameFormat_SetHeight(void* ptr, void* height)
{
	static_cast<QTextFrameFormat*>(ptr)->setHeight(*static_cast<QTextLength*>(height));
}

void QTextFrameFormat_SetHeight2(void* ptr, double height)
{
	static_cast<QTextFrameFormat*>(ptr)->setHeight(static_cast<double>(height));
}

void QTextFrameFormat_SetLeftMargin(void* ptr, double margin)
{
	static_cast<QTextFrameFormat*>(ptr)->setLeftMargin(static_cast<double>(margin));
}

void QTextFrameFormat_SetPadding(void* ptr, double width)
{
	static_cast<QTextFrameFormat*>(ptr)->setPadding(static_cast<double>(width));
}

void QTextFrameFormat_SetPageBreakPolicy(void* ptr, int policy)
{
	static_cast<QTextFrameFormat*>(ptr)->setPageBreakPolicy(static_cast<QTextFormat::PageBreakFlag>(policy));
}

void QTextFrameFormat_SetPosition(void* ptr, int policy)
{
	static_cast<QTextFrameFormat*>(ptr)->setPosition(static_cast<QTextFrameFormat::Position>(policy));
}

void QTextFrameFormat_SetRightMargin(void* ptr, double margin)
{
	static_cast<QTextFrameFormat*>(ptr)->setRightMargin(static_cast<double>(margin));
}

void QTextFrameFormat_SetTopMargin(void* ptr, double margin)
{
	static_cast<QTextFrameFormat*>(ptr)->setTopMargin(static_cast<double>(margin));
}

void QTextFrameFormat_SetWidth(void* ptr, void* width)
{
	static_cast<QTextFrameFormat*>(ptr)->setWidth(*static_cast<QTextLength*>(width));
}

void QTextFrameFormat_SetWidth2(void* ptr, double width)
{
	static_cast<QTextFrameFormat*>(ptr)->setWidth(static_cast<double>(width));
}

void* QTextImageFormat_NewQTextImageFormat()
{
	return new QTextImageFormat();
}

double QTextImageFormat_Height(void* ptr)
{
	return static_cast<double>(static_cast<QTextImageFormat*>(ptr)->height());
}

int QTextImageFormat_IsValid(void* ptr)
{
	return static_cast<QTextImageFormat*>(ptr)->isValid();
}

char* QTextImageFormat_Name(void* ptr)
{
	return static_cast<QTextImageFormat*>(ptr)->name().toUtf8().data();
}

void QTextImageFormat_SetHeight(void* ptr, double height)
{
	static_cast<QTextImageFormat*>(ptr)->setHeight(static_cast<double>(height));
}

void QTextImageFormat_SetName(void* ptr, char* name)
{
	static_cast<QTextImageFormat*>(ptr)->setName(QString(name));
}

void QTextImageFormat_SetWidth(void* ptr, double width)
{
	static_cast<QTextImageFormat*>(ptr)->setWidth(static_cast<double>(width));
}

double QTextImageFormat_Width(void* ptr)
{
	return static_cast<double>(static_cast<QTextImageFormat*>(ptr)->width());
}

double QTextInlineObject_Ascent(void* ptr)
{
	return static_cast<double>(static_cast<QTextInlineObject*>(ptr)->ascent());
}

double QTextInlineObject_Descent(void* ptr)
{
	return static_cast<double>(static_cast<QTextInlineObject*>(ptr)->descent());
}

void* QTextInlineObject_Format(void* ptr)
{
	return new QTextFormat(static_cast<QTextInlineObject*>(ptr)->format());
}

int QTextInlineObject_FormatIndex(void* ptr)
{
	return static_cast<QTextInlineObject*>(ptr)->formatIndex();
}

double QTextInlineObject_Height(void* ptr)
{
	return static_cast<double>(static_cast<QTextInlineObject*>(ptr)->height());
}

int QTextInlineObject_IsValid(void* ptr)
{
	return static_cast<QTextInlineObject*>(ptr)->isValid();
}

void* QTextInlineObject_Rect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QTextInlineObject*>(ptr)->rect()).x(), static_cast<QRectF>(static_cast<QTextInlineObject*>(ptr)->rect()).y(), static_cast<QRectF>(static_cast<QTextInlineObject*>(ptr)->rect()).width(), static_cast<QRectF>(static_cast<QTextInlineObject*>(ptr)->rect()).height());
}

void QTextInlineObject_SetAscent(void* ptr, double a)
{
	static_cast<QTextInlineObject*>(ptr)->setAscent(static_cast<double>(a));
}

void QTextInlineObject_SetDescent(void* ptr, double d)
{
	static_cast<QTextInlineObject*>(ptr)->setDescent(static_cast<double>(d));
}

void QTextInlineObject_SetWidth(void* ptr, double w)
{
	static_cast<QTextInlineObject*>(ptr)->setWidth(static_cast<double>(w));
}

int QTextInlineObject_TextDirection(void* ptr)
{
	return static_cast<QTextInlineObject*>(ptr)->textDirection();
}

int QTextInlineObject_TextPosition(void* ptr)
{
	return static_cast<QTextInlineObject*>(ptr)->textPosition();
}

double QTextInlineObject_Width(void* ptr)
{
	return static_cast<double>(static_cast<QTextInlineObject*>(ptr)->width());
}

double QTextItem_Ascent(void* ptr)
{
	return static_cast<double>(static_cast<QTextItem*>(ptr)->ascent());
}

double QTextItem_Descent(void* ptr)
{
	return static_cast<double>(static_cast<QTextItem*>(ptr)->descent());
}

void* QTextItem_Font(void* ptr)
{
	return new QFont(static_cast<QTextItem*>(ptr)->font());
}

int QTextItem_RenderFlags(void* ptr)
{
	return static_cast<QTextItem*>(ptr)->renderFlags();
}

char* QTextItem_Text(void* ptr)
{
	return static_cast<QTextItem*>(ptr)->text().toUtf8().data();
}

double QTextItem_Width(void* ptr)
{
	return static_cast<double>(static_cast<QTextItem*>(ptr)->width());
}

void QTextLayout_DrawCursor2(void* ptr, void* painter, void* position, int cursorPosition)
{
	static_cast<QTextLayout*>(ptr)->drawCursor(static_cast<QPainter*>(painter), *static_cast<QPointF*>(position), cursorPosition);
}

void QTextLayout_DrawCursor(void* ptr, void* painter, void* position, int cursorPosition, int width)
{
	static_cast<QTextLayout*>(ptr)->drawCursor(static_cast<QPainter*>(painter), *static_cast<QPointF*>(position), cursorPosition, width);
}

void* QTextLayout_NewQTextLayout()
{
	return new QTextLayout();
}

void* QTextLayout_NewQTextLayout2(char* text)
{
	return new QTextLayout(QString(text));
}

void* QTextLayout_NewQTextLayout3(char* text, void* font, void* paintdevice)
{
	return new QTextLayout(QString(text), *static_cast<QFont*>(font), static_cast<QPaintDevice*>(paintdevice));
}

void QTextLayout_BeginLayout(void* ptr)
{
	static_cast<QTextLayout*>(ptr)->beginLayout();
}

void* QTextLayout_BoundingRect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QTextLayout*>(ptr)->boundingRect()).x(), static_cast<QRectF>(static_cast<QTextLayout*>(ptr)->boundingRect()).y(), static_cast<QRectF>(static_cast<QTextLayout*>(ptr)->boundingRect()).width(), static_cast<QRectF>(static_cast<QTextLayout*>(ptr)->boundingRect()).height());
}

int QTextLayout_CacheEnabled(void* ptr)
{
	return static_cast<QTextLayout*>(ptr)->cacheEnabled();
}

void QTextLayout_ClearAdditionalFormats(void* ptr)
{
	static_cast<QTextLayout*>(ptr)->clearAdditionalFormats();
}

void QTextLayout_ClearLayout(void* ptr)
{
	static_cast<QTextLayout*>(ptr)->clearLayout();
}

int QTextLayout_CursorMoveStyle(void* ptr)
{
	return static_cast<QTextLayout*>(ptr)->cursorMoveStyle();
}

void QTextLayout_EndLayout(void* ptr)
{
	static_cast<QTextLayout*>(ptr)->endLayout();
}

void* QTextLayout_Font(void* ptr)
{
	return new QFont(static_cast<QTextLayout*>(ptr)->font());
}

int QTextLayout_IsValidCursorPosition(void* ptr, int pos)
{
	return static_cast<QTextLayout*>(ptr)->isValidCursorPosition(pos);
}

int QTextLayout_LeftCursorPosition(void* ptr, int oldPos)
{
	return static_cast<QTextLayout*>(ptr)->leftCursorPosition(oldPos);
}

int QTextLayout_LineCount(void* ptr)
{
	return static_cast<QTextLayout*>(ptr)->lineCount();
}

double QTextLayout_MaximumWidth(void* ptr)
{
	return static_cast<double>(static_cast<QTextLayout*>(ptr)->maximumWidth());
}

double QTextLayout_MinimumWidth(void* ptr)
{
	return static_cast<double>(static_cast<QTextLayout*>(ptr)->minimumWidth());
}

int QTextLayout_NextCursorPosition(void* ptr, int oldPos, int mode)
{
	return static_cast<QTextLayout*>(ptr)->nextCursorPosition(oldPos, static_cast<QTextLayout::CursorMode>(mode));
}

void* QTextLayout_Position(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QTextLayout*>(ptr)->position()).x(), static_cast<QPointF>(static_cast<QTextLayout*>(ptr)->position()).y());
}

int QTextLayout_PreeditAreaPosition(void* ptr)
{
	return static_cast<QTextLayout*>(ptr)->preeditAreaPosition();
}

char* QTextLayout_PreeditAreaText(void* ptr)
{
	return static_cast<QTextLayout*>(ptr)->preeditAreaText().toUtf8().data();
}

int QTextLayout_PreviousCursorPosition(void* ptr, int oldPos, int mode)
{
	return static_cast<QTextLayout*>(ptr)->previousCursorPosition(oldPos, static_cast<QTextLayout::CursorMode>(mode));
}

int QTextLayout_RightCursorPosition(void* ptr, int oldPos)
{
	return static_cast<QTextLayout*>(ptr)->rightCursorPosition(oldPos);
}

void QTextLayout_SetCacheEnabled(void* ptr, int enable)
{
	static_cast<QTextLayout*>(ptr)->setCacheEnabled(enable != 0);
}

void QTextLayout_SetCursorMoveStyle(void* ptr, int style)
{
	static_cast<QTextLayout*>(ptr)->setCursorMoveStyle(static_cast<Qt::CursorMoveStyle>(style));
}

void QTextLayout_SetFont(void* ptr, void* font)
{
	static_cast<QTextLayout*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QTextLayout_SetPosition(void* ptr, void* p)
{
	static_cast<QTextLayout*>(ptr)->setPosition(*static_cast<QPointF*>(p));
}

void QTextLayout_SetPreeditArea(void* ptr, int position, char* text)
{
	static_cast<QTextLayout*>(ptr)->setPreeditArea(position, QString(text));
}

void QTextLayout_SetText(void* ptr, char* stri)
{
	static_cast<QTextLayout*>(ptr)->setText(QString(stri));
}

void QTextLayout_SetTextOption(void* ptr, void* option)
{
	static_cast<QTextLayout*>(ptr)->setTextOption(*static_cast<QTextOption*>(option));
}

char* QTextLayout_Text(void* ptr)
{
	return static_cast<QTextLayout*>(ptr)->text().toUtf8().data();
}

void* QTextLayout_TextOption(void* ptr)
{
	return new QTextOption(static_cast<QTextLayout*>(ptr)->textOption());
}

void QTextLayout_DestroyQTextLayout(void* ptr)
{
	static_cast<QTextLayout*>(ptr)->~QTextLayout();
}

void* QTextLength_NewQTextLength()
{
	return new QTextLength();
}

void* QTextLength_NewQTextLength2(int ty, double value)
{
	return new QTextLength(static_cast<QTextLength::Type>(ty), static_cast<double>(value));
}

double QTextLength_RawValue(void* ptr)
{
	return static_cast<double>(static_cast<QTextLength*>(ptr)->rawValue());
}

int QTextLength_Type(void* ptr)
{
	return static_cast<QTextLength*>(ptr)->type();
}

double QTextLength_Value(void* ptr, double maximumLength)
{
	return static_cast<double>(static_cast<QTextLength*>(ptr)->value(static_cast<double>(maximumLength)));
}

int QTextLine_XToCursor(void* ptr, double x, int cpos)
{
	return static_cast<QTextLine*>(ptr)->xToCursor(static_cast<double>(x), static_cast<QTextLine::CursorPosition>(cpos));
}

void* QTextLine_NewQTextLine()
{
	return new QTextLine();
}

double QTextLine_Ascent(void* ptr)
{
	return static_cast<double>(static_cast<QTextLine*>(ptr)->ascent());
}

double QTextLine_CursorToX(void* ptr, int cursorPos, int edge)
{
	return static_cast<double>(static_cast<QTextLine*>(ptr)->cursorToX(&cursorPos, static_cast<QTextLine::Edge>(edge)));
}

double QTextLine_CursorToX2(void* ptr, int cursorPos, int edge)
{
	return static_cast<double>(static_cast<QTextLine*>(ptr)->cursorToX(cursorPos, static_cast<QTextLine::Edge>(edge)));
}

double QTextLine_Descent(void* ptr)
{
	return static_cast<double>(static_cast<QTextLine*>(ptr)->descent());
}

double QTextLine_Height(void* ptr)
{
	return static_cast<double>(static_cast<QTextLine*>(ptr)->height());
}

double QTextLine_HorizontalAdvance(void* ptr)
{
	return static_cast<double>(static_cast<QTextLine*>(ptr)->horizontalAdvance());
}

int QTextLine_IsValid(void* ptr)
{
	return static_cast<QTextLine*>(ptr)->isValid();
}

double QTextLine_Leading(void* ptr)
{
	return static_cast<double>(static_cast<QTextLine*>(ptr)->leading());
}

int QTextLine_LeadingIncluded(void* ptr)
{
	return static_cast<QTextLine*>(ptr)->leadingIncluded();
}

int QTextLine_LineNumber(void* ptr)
{
	return static_cast<QTextLine*>(ptr)->lineNumber();
}

void* QTextLine_NaturalTextRect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QTextLine*>(ptr)->naturalTextRect()).x(), static_cast<QRectF>(static_cast<QTextLine*>(ptr)->naturalTextRect()).y(), static_cast<QRectF>(static_cast<QTextLine*>(ptr)->naturalTextRect()).width(), static_cast<QRectF>(static_cast<QTextLine*>(ptr)->naturalTextRect()).height());
}

double QTextLine_NaturalTextWidth(void* ptr)
{
	return static_cast<double>(static_cast<QTextLine*>(ptr)->naturalTextWidth());
}

void* QTextLine_Position(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QTextLine*>(ptr)->position()).x(), static_cast<QPointF>(static_cast<QTextLine*>(ptr)->position()).y());
}

void* QTextLine_Rect(void* ptr)
{
	return new QRectF(static_cast<QRectF>(static_cast<QTextLine*>(ptr)->rect()).x(), static_cast<QRectF>(static_cast<QTextLine*>(ptr)->rect()).y(), static_cast<QRectF>(static_cast<QTextLine*>(ptr)->rect()).width(), static_cast<QRectF>(static_cast<QTextLine*>(ptr)->rect()).height());
}

void QTextLine_SetLeadingIncluded(void* ptr, int included)
{
	static_cast<QTextLine*>(ptr)->setLeadingIncluded(included != 0);
}

void QTextLine_SetLineWidth(void* ptr, double width)
{
	static_cast<QTextLine*>(ptr)->setLineWidth(static_cast<double>(width));
}

void QTextLine_SetNumColumns(void* ptr, int numColumns)
{
	static_cast<QTextLine*>(ptr)->setNumColumns(numColumns);
}

void QTextLine_SetNumColumns2(void* ptr, int numColumns, double alignmentWidth)
{
	static_cast<QTextLine*>(ptr)->setNumColumns(numColumns, static_cast<double>(alignmentWidth));
}

void QTextLine_SetPosition(void* ptr, void* pos)
{
	static_cast<QTextLine*>(ptr)->setPosition(*static_cast<QPointF*>(pos));
}

int QTextLine_TextLength(void* ptr)
{
	return static_cast<QTextLine*>(ptr)->textLength();
}

int QTextLine_TextStart(void* ptr)
{
	return static_cast<QTextLine*>(ptr)->textStart();
}

double QTextLine_Width(void* ptr)
{
	return static_cast<double>(static_cast<QTextLine*>(ptr)->width());
}

double QTextLine_X(void* ptr)
{
	return static_cast<double>(static_cast<QTextLine*>(ptr)->x());
}

double QTextLine_Y(void* ptr)
{
	return static_cast<double>(static_cast<QTextLine*>(ptr)->y());
}

int QTextList_ItemNumber(void* ptr, void* block)
{
	return static_cast<QTextList*>(ptr)->itemNumber(*static_cast<QTextBlock*>(block));
}

char* QTextList_ItemText(void* ptr, void* block)
{
	return static_cast<QTextList*>(ptr)->itemText(*static_cast<QTextBlock*>(block)).toUtf8().data();
}

void QTextList_Add(void* ptr, void* block)
{
	static_cast<QTextList*>(ptr)->add(*static_cast<QTextBlock*>(block));
}

int QTextList_Count(void* ptr)
{
	return static_cast<QTextList*>(ptr)->count();
}

void* QTextList_Item(void* ptr, int i)
{
	return new QTextBlock(static_cast<QTextList*>(ptr)->item(i));
}

void QTextList_Remove(void* ptr, void* block)
{
	static_cast<QTextList*>(ptr)->remove(*static_cast<QTextBlock*>(block));
}

void QTextList_RemoveItem(void* ptr, int i)
{
	static_cast<QTextList*>(ptr)->removeItem(i);
}

void QTextList_SetFormat(void* ptr, void* format)
{
	static_cast<QTextList*>(ptr)->setFormat(*static_cast<QTextListFormat*>(format));
}

void QTextList_BlockFormatChanged(void* ptr, void* block)
{
	static_cast<QTextList*>(ptr)->blockFormatChanged(*static_cast<QTextBlock*>(block));
}

void QTextList_BlockFormatChangedDefault(void* ptr, void* block)
{
	static_cast<QTextList*>(ptr)->QTextList::blockFormatChanged(*static_cast<QTextBlock*>(block));
}

void QTextList_BlockInserted(void* ptr, void* block)
{
	static_cast<QTextList*>(ptr)->blockInserted(*static_cast<QTextBlock*>(block));
}

void QTextList_BlockInsertedDefault(void* ptr, void* block)
{
	static_cast<QTextList*>(ptr)->QTextList::blockInserted(*static_cast<QTextBlock*>(block));
}

void QTextList_BlockRemoved(void* ptr, void* block)
{
	static_cast<QTextList*>(ptr)->blockRemoved(*static_cast<QTextBlock*>(block));
}

void QTextList_BlockRemovedDefault(void* ptr, void* block)
{
	static_cast<QTextList*>(ptr)->QTextList::blockRemoved(*static_cast<QTextBlock*>(block));
}

void QTextList_TimerEvent(void* ptr, void* event)
{
	static_cast<QTextList*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTextList_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QTextList*>(ptr)->QTextList::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTextList_ChildEvent(void* ptr, void* event)
{
	static_cast<QTextList*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTextList_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QTextList*>(ptr)->QTextList::childEvent(static_cast<QChildEvent*>(event));
}

void QTextList_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QTextList*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextList_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTextList*>(ptr)->QTextList::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextList_CustomEvent(void* ptr, void* event)
{
	static_cast<QTextList*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTextList_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QTextList*>(ptr)->QTextList::customEvent(static_cast<QEvent*>(event));
}

void QTextList_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTextList*>(ptr), "deleteLater");
}

void QTextList_DeleteLaterDefault(void* ptr)
{
	static_cast<QTextList*>(ptr)->QTextList::deleteLater();
}

void QTextList_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QTextList*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextList_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTextList*>(ptr)->QTextList::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QTextList_Event(void* ptr, void* e)
{
	return static_cast<QTextList*>(ptr)->event(static_cast<QEvent*>(e));
}

int QTextList_EventDefault(void* ptr, void* e)
{
	return static_cast<QTextList*>(ptr)->QTextList::event(static_cast<QEvent*>(e));
}

int QTextList_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QTextList*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QTextList_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QTextList*>(ptr)->QTextList::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QTextList_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTextList*>(ptr)->metaObject());
}

void* QTextList_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTextList*>(ptr)->QTextList::metaObject());
}

void* QTextListFormat_NewQTextListFormat()
{
	return new QTextListFormat();
}

int QTextListFormat_Indent(void* ptr)
{
	return static_cast<QTextListFormat*>(ptr)->indent();
}

int QTextListFormat_IsValid(void* ptr)
{
	return static_cast<QTextListFormat*>(ptr)->isValid();
}

char* QTextListFormat_NumberPrefix(void* ptr)
{
	return static_cast<QTextListFormat*>(ptr)->numberPrefix().toUtf8().data();
}

char* QTextListFormat_NumberSuffix(void* ptr)
{
	return static_cast<QTextListFormat*>(ptr)->numberSuffix().toUtf8().data();
}

void QTextListFormat_SetIndent(void* ptr, int indentation)
{
	static_cast<QTextListFormat*>(ptr)->setIndent(indentation);
}

void QTextListFormat_SetNumberPrefix(void* ptr, char* numberPrefix)
{
	static_cast<QTextListFormat*>(ptr)->setNumberPrefix(QString(numberPrefix));
}

void QTextListFormat_SetNumberSuffix(void* ptr, char* numberSuffix)
{
	static_cast<QTextListFormat*>(ptr)->setNumberSuffix(QString(numberSuffix));
}

void QTextListFormat_SetStyle(void* ptr, int style)
{
	static_cast<QTextListFormat*>(ptr)->setStyle(static_cast<QTextListFormat::Style>(style));
}

int QTextListFormat_Style(void* ptr)
{
	return static_cast<QTextListFormat*>(ptr)->style();
}

void* QTextObject_NewQTextObject(void* document)
{
	return new QTextObject(static_cast<QTextDocument*>(document));
}

void* QTextObject_Document(void* ptr)
{
	return static_cast<QTextObject*>(ptr)->document();
}

void* QTextObject_Format(void* ptr)
{
	return new QTextFormat(static_cast<QTextObject*>(ptr)->format());
}

int QTextObject_FormatIndex(void* ptr)
{
	return static_cast<QTextObject*>(ptr)->formatIndex();
}

int QTextObject_ObjectIndex(void* ptr)
{
	return static_cast<QTextObject*>(ptr)->objectIndex();
}

void QTextObject_SetFormat(void* ptr, void* format)
{
	static_cast<QTextObject*>(ptr)->setFormat(*static_cast<QTextFormat*>(format));
}

void QTextObject_DestroyQTextObject(void* ptr)
{
	static_cast<QTextObject*>(ptr)->~QTextObject();
}

void QTextObject_TimerEvent(void* ptr, void* event)
{
	static_cast<QTextObject*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTextObject_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QTextObject*>(ptr)->QTextObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTextObject_ChildEvent(void* ptr, void* event)
{
	static_cast<QTextObject*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTextObject_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QTextObject*>(ptr)->QTextObject::childEvent(static_cast<QChildEvent*>(event));
}

void QTextObject_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QTextObject*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextObject_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTextObject*>(ptr)->QTextObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextObject_CustomEvent(void* ptr, void* event)
{
	static_cast<QTextObject*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTextObject_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QTextObject*>(ptr)->QTextObject::customEvent(static_cast<QEvent*>(event));
}

void QTextObject_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTextObject*>(ptr), "deleteLater");
}

void QTextObject_DeleteLaterDefault(void* ptr)
{
	static_cast<QTextObject*>(ptr)->QTextObject::deleteLater();
}

void QTextObject_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QTextObject*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextObject_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTextObject*>(ptr)->QTextObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QTextObject_Event(void* ptr, void* e)
{
	return static_cast<QTextObject*>(ptr)->event(static_cast<QEvent*>(e));
}

int QTextObject_EventDefault(void* ptr, void* e)
{
	return static_cast<QTextObject*>(ptr)->QTextObject::event(static_cast<QEvent*>(e));
}

int QTextObject_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QTextObject*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QTextObject_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QTextObject*>(ptr)->QTextObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QTextObject_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTextObject*>(ptr)->metaObject());
}

void* QTextObject_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTextObject*>(ptr)->QTextObject::metaObject());
}

class MyQTextObjectInterface: public QTextObjectInterface
{
public:
	QString _objectName;
	QString objectNameAbs() const { return this->_objectName; };
	void setObjectNameAbs(const QString &name) { this->_objectName = name; };
	void drawObject(QPainter * painter, const QRectF & rect, QTextDocument * doc, int posInDocument, const QTextFormat & format) { callbackQTextObjectInterface_DrawObject(this, this->objectNameAbs().toUtf8().data(), painter, new QRectF(static_cast<QRectF>(rect).x(), static_cast<QRectF>(rect).y(), static_cast<QRectF>(rect).width(), static_cast<QRectF>(rect).height()), doc, posInDocument, new QTextFormat(format)); };
	QSizeF intrinsicSize(QTextDocument * doc, int posInDocument, const QTextFormat & format) { return *static_cast<QSizeF*>(callbackQTextObjectInterface_IntrinsicSize(this, this->objectNameAbs().toUtf8().data(), doc, posInDocument, new QTextFormat(format))); };
};

void QTextObjectInterface_DrawObject(void* ptr, void* painter, void* rect, void* doc, int posInDocument, void* format)
{
	static_cast<QTextObjectInterface*>(ptr)->drawObject(static_cast<QPainter*>(painter), *static_cast<QRectF*>(rect), static_cast<QTextDocument*>(doc), posInDocument, *static_cast<QTextFormat*>(format));
}

void* QTextObjectInterface_IntrinsicSize(void* ptr, void* doc, int posInDocument, void* format)
{
	return new QSizeF(static_cast<QSizeF>(static_cast<QTextObjectInterface*>(ptr)->intrinsicSize(static_cast<QTextDocument*>(doc), posInDocument, *static_cast<QTextFormat*>(format))).width(), static_cast<QSizeF>(static_cast<QTextObjectInterface*>(ptr)->intrinsicSize(static_cast<QTextDocument*>(doc), posInDocument, *static_cast<QTextFormat*>(format))).height());
}

void QTextObjectInterface_DestroyQTextObjectInterface(void* ptr)
{
	static_cast<QTextObjectInterface*>(ptr)->~QTextObjectInterface();
}

char* QTextObjectInterface_ObjectNameAbs(void* ptr)
{
	if (dynamic_cast<MyQTextObjectInterface*>(static_cast<QTextObjectInterface*>(ptr))) {
		return static_cast<MyQTextObjectInterface*>(ptr)->objectNameAbs().toUtf8().data();
	}
	return QString("QTextObjectInterface_BASE").toUtf8().data();
}

void QTextObjectInterface_SetObjectNameAbs(void* ptr, char* name)
{
	if (dynamic_cast<MyQTextObjectInterface*>(static_cast<QTextObjectInterface*>(ptr))) {
		static_cast<MyQTextObjectInterface*>(ptr)->setObjectNameAbs(QString(name));
	}
}

void* QTextOption_NewQTextOption3(void* other)
{
	return new QTextOption(*static_cast<QTextOption*>(other));
}

void* QTextOption_NewQTextOption()
{
	return new QTextOption();
}

void* QTextOption_NewQTextOption2(int alignment)
{
	return new QTextOption(static_cast<Qt::AlignmentFlag>(alignment));
}

int QTextOption_Alignment(void* ptr)
{
	return static_cast<QTextOption*>(ptr)->alignment();
}

int QTextOption_Flags(void* ptr)
{
	return static_cast<QTextOption*>(ptr)->flags();
}

void QTextOption_SetAlignment(void* ptr, int alignment)
{
	static_cast<QTextOption*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QTextOption_SetFlags(void* ptr, int flags)
{
	static_cast<QTextOption*>(ptr)->setFlags(static_cast<QTextOption::Flag>(flags));
}

void QTextOption_SetTabStop(void* ptr, double tabStop)
{
	static_cast<QTextOption*>(ptr)->setTabStop(static_cast<double>(tabStop));
}

void QTextOption_SetTextDirection(void* ptr, int direction)
{
	static_cast<QTextOption*>(ptr)->setTextDirection(static_cast<Qt::LayoutDirection>(direction));
}

void QTextOption_SetUseDesignMetrics(void* ptr, int enable)
{
	static_cast<QTextOption*>(ptr)->setUseDesignMetrics(enable != 0);
}

void QTextOption_SetWrapMode(void* ptr, int mode)
{
	static_cast<QTextOption*>(ptr)->setWrapMode(static_cast<QTextOption::WrapMode>(mode));
}

double QTextOption_TabStop(void* ptr)
{
	return static_cast<double>(static_cast<QTextOption*>(ptr)->tabStop());
}

int QTextOption_TextDirection(void* ptr)
{
	return static_cast<QTextOption*>(ptr)->textDirection();
}

int QTextOption_UseDesignMetrics(void* ptr)
{
	return static_cast<QTextOption*>(ptr)->useDesignMetrics();
}

int QTextOption_WrapMode(void* ptr)
{
	return static_cast<QTextOption*>(ptr)->wrapMode();
}

void QTextOption_DestroyQTextOption(void* ptr)
{
	static_cast<QTextOption*>(ptr)->~QTextOption();
}

void* QTextTable_CellAt3(void* ptr, void* cursor)
{
	return new QTextTableCell(static_cast<QTextTable*>(ptr)->cellAt(*static_cast<QTextCursor*>(cursor)));
}

void* QTextTable_CellAt(void* ptr, int row, int column)
{
	return new QTextTableCell(static_cast<QTextTable*>(ptr)->cellAt(row, column));
}

void QTextTable_InsertColumns(void* ptr, int index, int columns)
{
	static_cast<QTextTable*>(ptr)->insertColumns(index, columns);
}

void QTextTable_InsertRows(void* ptr, int index, int rows)
{
	static_cast<QTextTable*>(ptr)->insertRows(index, rows);
}

void QTextTable_RemoveColumns(void* ptr, int index, int columns)
{
	static_cast<QTextTable*>(ptr)->removeColumns(index, columns);
}

void QTextTable_RemoveRows(void* ptr, int index, int rows)
{
	static_cast<QTextTable*>(ptr)->removeRows(index, rows);
}

void QTextTable_Resize(void* ptr, int rows, int columns)
{
	static_cast<QTextTable*>(ptr)->resize(rows, columns);
}

void* QTextTable_RowEnd(void* ptr, void* cursor)
{
	return new QTextCursor(static_cast<QTextTable*>(ptr)->rowEnd(*static_cast<QTextCursor*>(cursor)));
}

void* QTextTable_RowStart(void* ptr, void* cursor)
{
	return new QTextCursor(static_cast<QTextTable*>(ptr)->rowStart(*static_cast<QTextCursor*>(cursor)));
}

void QTextTable_SetFormat(void* ptr, void* format)
{
	static_cast<QTextTable*>(ptr)->setFormat(*static_cast<QTextTableFormat*>(format));
}

void QTextTable_AppendColumns(void* ptr, int count)
{
	static_cast<QTextTable*>(ptr)->appendColumns(count);
}

void QTextTable_AppendRows(void* ptr, int count)
{
	static_cast<QTextTable*>(ptr)->appendRows(count);
}

void* QTextTable_CellAt2(void* ptr, int position)
{
	return new QTextTableCell(static_cast<QTextTable*>(ptr)->cellAt(position));
}

int QTextTable_Columns(void* ptr)
{
	return static_cast<QTextTable*>(ptr)->columns();
}

void QTextTable_MergeCells2(void* ptr, void* cursor)
{
	static_cast<QTextTable*>(ptr)->mergeCells(*static_cast<QTextCursor*>(cursor));
}

void QTextTable_MergeCells(void* ptr, int row, int column, int numRows, int numCols)
{
	static_cast<QTextTable*>(ptr)->mergeCells(row, column, numRows, numCols);
}

int QTextTable_Rows(void* ptr)
{
	return static_cast<QTextTable*>(ptr)->rows();
}

void QTextTable_SplitCell(void* ptr, int row, int column, int numRows, int numCols)
{
	static_cast<QTextTable*>(ptr)->splitCell(row, column, numRows, numCols);
}

void QTextTable_TimerEvent(void* ptr, void* event)
{
	static_cast<QTextTable*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QTextTable_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QTextTable*>(ptr)->QTextTable::timerEvent(static_cast<QTimerEvent*>(event));
}

void QTextTable_ChildEvent(void* ptr, void* event)
{
	static_cast<QTextTable*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QTextTable_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QTextTable*>(ptr)->QTextTable::childEvent(static_cast<QChildEvent*>(event));
}

void QTextTable_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QTextTable*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextTable_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTextTable*>(ptr)->QTextTable::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextTable_CustomEvent(void* ptr, void* event)
{
	static_cast<QTextTable*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QTextTable_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QTextTable*>(ptr)->QTextTable::customEvent(static_cast<QEvent*>(event));
}

void QTextTable_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTextTable*>(ptr), "deleteLater");
}

void QTextTable_DeleteLaterDefault(void* ptr)
{
	static_cast<QTextTable*>(ptr)->QTextTable::deleteLater();
}

void QTextTable_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QTextTable*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QTextTable_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QTextTable*>(ptr)->QTextTable::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QTextTable_Event(void* ptr, void* e)
{
	return static_cast<QTextTable*>(ptr)->event(static_cast<QEvent*>(e));
}

int QTextTable_EventDefault(void* ptr, void* e)
{
	return static_cast<QTextTable*>(ptr)->QTextTable::event(static_cast<QEvent*>(e));
}

int QTextTable_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QTextTable*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QTextTable_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QTextTable*>(ptr)->QTextTable::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QTextTable_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTextTable*>(ptr)->metaObject());
}

void* QTextTable_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QTextTable*>(ptr)->QTextTable::metaObject());
}

void* QTextTableCell_NewQTextTableCell()
{
	return new QTextTableCell();
}

void* QTextTableCell_NewQTextTableCell2(void* other)
{
	return new QTextTableCell(*static_cast<QTextTableCell*>(other));
}

int QTextTableCell_Column(void* ptr)
{
	return static_cast<QTextTableCell*>(ptr)->column();
}

int QTextTableCell_ColumnSpan(void* ptr)
{
	return static_cast<QTextTableCell*>(ptr)->columnSpan();
}

void* QTextTableCell_FirstCursorPosition(void* ptr)
{
	return new QTextCursor(static_cast<QTextTableCell*>(ptr)->firstCursorPosition());
}

int QTextTableCell_IsValid(void* ptr)
{
	return static_cast<QTextTableCell*>(ptr)->isValid();
}

void* QTextTableCell_LastCursorPosition(void* ptr)
{
	return new QTextCursor(static_cast<QTextTableCell*>(ptr)->lastCursorPosition());
}

int QTextTableCell_Row(void* ptr)
{
	return static_cast<QTextTableCell*>(ptr)->row();
}

int QTextTableCell_RowSpan(void* ptr)
{
	return static_cast<QTextTableCell*>(ptr)->rowSpan();
}

void QTextTableCell_SetFormat(void* ptr, void* format)
{
	static_cast<QTextTableCell*>(ptr)->setFormat(*static_cast<QTextCharFormat*>(format));
}

int QTextTableCell_TableCellFormatIndex(void* ptr)
{
	return static_cast<QTextTableCell*>(ptr)->tableCellFormatIndex();
}

void QTextTableCell_DestroyQTextTableCell(void* ptr)
{
	static_cast<QTextTableCell*>(ptr)->~QTextTableCell();
}

void* QTextTableCellFormat_NewQTextTableCellFormat()
{
	return new QTextTableCellFormat();
}

double QTextTableCellFormat_BottomPadding(void* ptr)
{
	return static_cast<double>(static_cast<QTextTableCellFormat*>(ptr)->bottomPadding());
}

int QTextTableCellFormat_IsValid(void* ptr)
{
	return static_cast<QTextTableCellFormat*>(ptr)->isValid();
}

double QTextTableCellFormat_LeftPadding(void* ptr)
{
	return static_cast<double>(static_cast<QTextTableCellFormat*>(ptr)->leftPadding());
}

double QTextTableCellFormat_RightPadding(void* ptr)
{
	return static_cast<double>(static_cast<QTextTableCellFormat*>(ptr)->rightPadding());
}

void QTextTableCellFormat_SetBottomPadding(void* ptr, double padding)
{
	static_cast<QTextTableCellFormat*>(ptr)->setBottomPadding(static_cast<double>(padding));
}

void QTextTableCellFormat_SetLeftPadding(void* ptr, double padding)
{
	static_cast<QTextTableCellFormat*>(ptr)->setLeftPadding(static_cast<double>(padding));
}

void QTextTableCellFormat_SetPadding(void* ptr, double padding)
{
	static_cast<QTextTableCellFormat*>(ptr)->setPadding(static_cast<double>(padding));
}

void QTextTableCellFormat_SetRightPadding(void* ptr, double padding)
{
	static_cast<QTextTableCellFormat*>(ptr)->setRightPadding(static_cast<double>(padding));
}

void QTextTableCellFormat_SetTopPadding(void* ptr, double padding)
{
	static_cast<QTextTableCellFormat*>(ptr)->setTopPadding(static_cast<double>(padding));
}

double QTextTableCellFormat_TopPadding(void* ptr)
{
	return static_cast<double>(static_cast<QTextTableCellFormat*>(ptr)->topPadding());
}

void* QTextTableFormat_NewQTextTableFormat()
{
	return new QTextTableFormat();
}

int QTextTableFormat_Alignment(void* ptr)
{
	return static_cast<QTextTableFormat*>(ptr)->alignment();
}

double QTextTableFormat_CellPadding(void* ptr)
{
	return static_cast<double>(static_cast<QTextTableFormat*>(ptr)->cellPadding());
}

double QTextTableFormat_CellSpacing(void* ptr)
{
	return static_cast<double>(static_cast<QTextTableFormat*>(ptr)->cellSpacing());
}

void QTextTableFormat_ClearColumnWidthConstraints(void* ptr)
{
	static_cast<QTextTableFormat*>(ptr)->clearColumnWidthConstraints();
}

int QTextTableFormat_Columns(void* ptr)
{
	return static_cast<QTextTableFormat*>(ptr)->columns();
}

int QTextTableFormat_HeaderRowCount(void* ptr)
{
	return static_cast<QTextTableFormat*>(ptr)->headerRowCount();
}

int QTextTableFormat_IsValid(void* ptr)
{
	return static_cast<QTextTableFormat*>(ptr)->isValid();
}

void QTextTableFormat_SetAlignment(void* ptr, int alignment)
{
	static_cast<QTextTableFormat*>(ptr)->setAlignment(static_cast<Qt::AlignmentFlag>(alignment));
}

void QTextTableFormat_SetCellPadding(void* ptr, double padding)
{
	static_cast<QTextTableFormat*>(ptr)->setCellPadding(static_cast<double>(padding));
}

void QTextTableFormat_SetCellSpacing(void* ptr, double spacing)
{
	static_cast<QTextTableFormat*>(ptr)->setCellSpacing(static_cast<double>(spacing));
}

void QTextTableFormat_SetHeaderRowCount(void* ptr, int count)
{
	static_cast<QTextTableFormat*>(ptr)->setHeaderRowCount(count);
}

void* QTouchDevice_NewQTouchDevice()
{
	return new QTouchDevice();
}

int QTouchDevice_Capabilities(void* ptr)
{
	return static_cast<QTouchDevice*>(ptr)->capabilities();
}

int QTouchDevice_MaximumTouchPoints(void* ptr)
{
	return static_cast<QTouchDevice*>(ptr)->maximumTouchPoints();
}

char* QTouchDevice_Name(void* ptr)
{
	return static_cast<QTouchDevice*>(ptr)->name().toUtf8().data();
}

void QTouchDevice_SetCapabilities(void* ptr, int caps)
{
	static_cast<QTouchDevice*>(ptr)->setCapabilities(static_cast<QTouchDevice::CapabilityFlag>(caps));
}

void QTouchDevice_SetMaximumTouchPoints(void* ptr, int max)
{
	static_cast<QTouchDevice*>(ptr)->setMaximumTouchPoints(max);
}

void QTouchDevice_SetName(void* ptr, char* name)
{
	static_cast<QTouchDevice*>(ptr)->setName(QString(name));
}

void QTouchDevice_SetType(void* ptr, int devType)
{
	static_cast<QTouchDevice*>(ptr)->setType(static_cast<QTouchDevice::DeviceType>(devType));
}

int QTouchDevice_Type(void* ptr)
{
	return static_cast<QTouchDevice*>(ptr)->type();
}

void QTouchDevice_DestroyQTouchDevice(void* ptr)
{
	static_cast<QTouchDevice*>(ptr)->~QTouchDevice();
}

void* QTouchEvent_Device(void* ptr)
{
	return static_cast<QTouchEvent*>(ptr)->device();
}

void* QTouchEvent_Target(void* ptr)
{
	return static_cast<QTouchEvent*>(ptr)->target();
}

int QTouchEvent_TouchPointStates(void* ptr)
{
	return static_cast<QTouchEvent*>(ptr)->touchPointStates();
}

void* QTouchEvent_Window(void* ptr)
{
	return static_cast<QTouchEvent*>(ptr)->window();
}

void QTouchEvent_DestroyQTouchEvent(void* ptr)
{
	static_cast<QTouchEvent*>(ptr)->~QTouchEvent();
}

void* QTransform_NewQTransform3(double m11, double m12, double m13, double m21, double m22, double m23, double m31, double m32, double m33)
{
	return new QTransform(static_cast<double>(m11), static_cast<double>(m12), static_cast<double>(m13), static_cast<double>(m21), static_cast<double>(m22), static_cast<double>(m23), static_cast<double>(m31), static_cast<double>(m32), static_cast<double>(m33));
}

void* QTransform_NewQTransform4(double m11, double m12, double m21, double m22, double dx, double dy)
{
	return new QTransform(static_cast<double>(m11), static_cast<double>(m12), static_cast<double>(m21), static_cast<double>(m22), static_cast<double>(dx), static_cast<double>(dy));
}

void* QTransform_Map5(void* ptr, void* line)
{
	return new QLineF(static_cast<QLineF>(static_cast<QTransform*>(ptr)->map(*static_cast<QLineF*>(line))).p1(), static_cast<QLineF>(static_cast<QTransform*>(ptr)->map(*static_cast<QLineF*>(line))).p2());
}

void* QTransform_Map3(void* ptr, void* point)
{
	return new QPoint(static_cast<QPoint>(static_cast<QTransform*>(ptr)->map(*static_cast<QPoint*>(point))).x(), static_cast<QPoint>(static_cast<QTransform*>(ptr)->map(*static_cast<QPoint*>(point))).y());
}

void* QTransform_Map7(void* ptr, void* polygon)
{
	return new QPolygon(static_cast<QTransform*>(ptr)->map(*static_cast<QPolygon*>(polygon)));
}

void* QTransform_Map6(void* ptr, void* polygon)
{
	return new QPolygonF(static_cast<QTransform*>(ptr)->map(*static_cast<QPolygonF*>(polygon)));
}

void* QTransform_Map8(void* ptr, void* region)
{
	return new QRegion(static_cast<QTransform*>(ptr)->map(*static_cast<QRegion*>(region)));
}

void* QTransform_MapRect2(void* ptr, void* rectangle)
{
	return new QRect(static_cast<QRect>(static_cast<QTransform*>(ptr)->mapRect(*static_cast<QRect*>(rectangle))).x(), static_cast<QRect>(static_cast<QTransform*>(ptr)->mapRect(*static_cast<QRect*>(rectangle))).y(), static_cast<QRect>(static_cast<QTransform*>(ptr)->mapRect(*static_cast<QRect*>(rectangle))).width(), static_cast<QRect>(static_cast<QTransform*>(ptr)->mapRect(*static_cast<QRect*>(rectangle))).height());
}

void* QTransform_MapRect(void* ptr, void* rectangle)
{
	return new QRectF(static_cast<QRectF>(static_cast<QTransform*>(ptr)->mapRect(*static_cast<QRectF*>(rectangle))).x(), static_cast<QRectF>(static_cast<QTransform*>(ptr)->mapRect(*static_cast<QRectF*>(rectangle))).y(), static_cast<QRectF>(static_cast<QTransform*>(ptr)->mapRect(*static_cast<QRectF*>(rectangle))).width(), static_cast<QRectF>(static_cast<QTransform*>(ptr)->mapRect(*static_cast<QRectF*>(rectangle))).height());
}

void* QTransform_MapToPolygon(void* ptr, void* rectangle)
{
	return new QPolygon(static_cast<QTransform*>(ptr)->mapToPolygon(*static_cast<QRect*>(rectangle)));
}

int QTransform_QTransform_QuadToSquare(void* quad, void* trans)
{
	return QTransform::quadToSquare(*static_cast<QPolygonF*>(quad), *static_cast<QTransform*>(trans));
}

void* QTransform_NewQTransform()
{
	return new QTransform();
}

double QTransform_Determinant(void* ptr)
{
	return static_cast<double>(static_cast<QTransform*>(ptr)->determinant());
}

double QTransform_Dx(void* ptr)
{
	return static_cast<double>(static_cast<QTransform*>(ptr)->dx());
}

double QTransform_Dy(void* ptr)
{
	return static_cast<double>(static_cast<QTransform*>(ptr)->dy());
}

int QTransform_IsAffine(void* ptr)
{
	return static_cast<QTransform*>(ptr)->isAffine();
}

int QTransform_IsIdentity(void* ptr)
{
	return static_cast<QTransform*>(ptr)->isIdentity();
}

int QTransform_IsInvertible(void* ptr)
{
	return static_cast<QTransform*>(ptr)->isInvertible();
}

int QTransform_IsRotating(void* ptr)
{
	return static_cast<QTransform*>(ptr)->isRotating();
}

int QTransform_IsScaling(void* ptr)
{
	return static_cast<QTransform*>(ptr)->isScaling();
}

int QTransform_IsTranslating(void* ptr)
{
	return static_cast<QTransform*>(ptr)->isTranslating();
}

double QTransform_M11(void* ptr)
{
	return static_cast<double>(static_cast<QTransform*>(ptr)->m11());
}

double QTransform_M12(void* ptr)
{
	return static_cast<double>(static_cast<QTransform*>(ptr)->m12());
}

double QTransform_M13(void* ptr)
{
	return static_cast<double>(static_cast<QTransform*>(ptr)->m13());
}

double QTransform_M21(void* ptr)
{
	return static_cast<double>(static_cast<QTransform*>(ptr)->m21());
}

double QTransform_M22(void* ptr)
{
	return static_cast<double>(static_cast<QTransform*>(ptr)->m22());
}

double QTransform_M23(void* ptr)
{
	return static_cast<double>(static_cast<QTransform*>(ptr)->m23());
}

double QTransform_M31(void* ptr)
{
	return static_cast<double>(static_cast<QTransform*>(ptr)->m31());
}

double QTransform_M32(void* ptr)
{
	return static_cast<double>(static_cast<QTransform*>(ptr)->m32());
}

double QTransform_M33(void* ptr)
{
	return static_cast<double>(static_cast<QTransform*>(ptr)->m33());
}

void* QTransform_Map4(void* ptr, void* l)
{
	return new QLine(static_cast<QLine>(static_cast<QTransform*>(ptr)->map(*static_cast<QLine*>(l))).p1(), static_cast<QLine>(static_cast<QTransform*>(ptr)->map(*static_cast<QLine*>(l))).p2());
}

void* QTransform_Map9(void* ptr, void* path)
{
	return new QPainterPath(static_cast<QTransform*>(ptr)->map(*static_cast<QPainterPath*>(path)));
}

void* QTransform_Map2(void* ptr, void* p)
{
	return new QPointF(static_cast<QPointF>(static_cast<QTransform*>(ptr)->map(*static_cast<QPointF*>(p))).x(), static_cast<QPointF>(static_cast<QTransform*>(ptr)->map(*static_cast<QPointF*>(p))).y());
}

void QTransform_Map10(void* ptr, int x, int y, int tx, int ty)
{
	static_cast<QTransform*>(ptr)->map(x, y, &tx, &ty);
}

int QTransform_QTransform_QuadToQuad(void* one, void* two, void* trans)
{
	return QTransform::quadToQuad(*static_cast<QPolygonF*>(one), *static_cast<QPolygonF*>(two), *static_cast<QTransform*>(trans));
}

void QTransform_Reset(void* ptr)
{
	static_cast<QTransform*>(ptr)->reset();
}

void QTransform_SetMatrix(void* ptr, double m11, double m12, double m13, double m21, double m22, double m23, double m31, double m32, double m33)
{
	static_cast<QTransform*>(ptr)->setMatrix(static_cast<double>(m11), static_cast<double>(m12), static_cast<double>(m13), static_cast<double>(m21), static_cast<double>(m22), static_cast<double>(m23), static_cast<double>(m31), static_cast<double>(m32), static_cast<double>(m33));
}

int QTransform_QTransform_SquareToQuad(void* quad, void* trans)
{
	return QTransform::squareToQuad(*static_cast<QPolygonF*>(quad), *static_cast<QTransform*>(trans));
}

int QTransform_Type(void* ptr)
{
	return static_cast<QTransform*>(ptr)->type();
}

class MyQValidator: public QValidator
{
public:
	void fixup(QString & input) const { callbackQValidator_Fixup(const_cast<MyQValidator*>(this), this->objectName().toUtf8().data(), input.toUtf8().data()); };
	void Signal_Changed() { callbackQValidator_Changed(this, this->objectName().toUtf8().data()); };
	State validate(QString & input, int & pos) const { return static_cast<QValidator::State>(callbackQValidator_Validate(const_cast<MyQValidator*>(this), this->objectName().toUtf8().data(), input.toUtf8().data(), pos)); };
	void timerEvent(QTimerEvent * event) { callbackQValidator_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQValidator_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQValidator_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQValidator_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQValidator_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQValidator_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool event(QEvent * e) { return callbackQValidator_Event(this, this->objectName().toUtf8().data(), e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQValidator_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQValidator_MetaObject(const_cast<MyQValidator*>(this), this->objectName().toUtf8().data())); };
};





void QValidator_ConnectChanged(void* ptr)
{
	QObject::connect(static_cast<QValidator*>(ptr), static_cast<void (QValidator::*)()>(&QValidator::changed), static_cast<MyQValidator*>(ptr), static_cast<void (MyQValidator::*)()>(&MyQValidator::Signal_Changed));
}

void QValidator_DisconnectChanged(void* ptr)
{
	QObject::disconnect(static_cast<QValidator*>(ptr), static_cast<void (QValidator::*)()>(&QValidator::changed), static_cast<MyQValidator*>(ptr), static_cast<void (MyQValidator::*)()>(&MyQValidator::Signal_Changed));
}

void QValidator_Changed(void* ptr)
{
	static_cast<QValidator*>(ptr)->changed();
}

void* QValidator_Locale(void* ptr)
{
	return new QLocale(static_cast<QValidator*>(ptr)->locale());
}

void QValidator_SetLocale(void* ptr, void* locale)
{
	static_cast<QValidator*>(ptr)->setLocale(*static_cast<QLocale*>(locale));
}



void QValidator_DestroyQValidator(void* ptr)
{
	static_cast<QValidator*>(ptr)->~QValidator();
}

void QValidator_TimerEvent(void* ptr, void* event)
{
	static_cast<QValidator*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QValidator_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QValidator*>(ptr)->QValidator::timerEvent(static_cast<QTimerEvent*>(event));
}

void QValidator_ChildEvent(void* ptr, void* event)
{
	static_cast<QValidator*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QValidator_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QValidator*>(ptr)->QValidator::childEvent(static_cast<QChildEvent*>(event));
}

void QValidator_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QValidator*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QValidator_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QValidator*>(ptr)->QValidator::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QValidator_CustomEvent(void* ptr, void* event)
{
	static_cast<QValidator*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QValidator_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QValidator*>(ptr)->QValidator::customEvent(static_cast<QEvent*>(event));
}

void QValidator_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QValidator*>(ptr), "deleteLater");
}

void QValidator_DeleteLaterDefault(void* ptr)
{
	static_cast<QValidator*>(ptr)->QValidator::deleteLater();
}

void QValidator_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QValidator*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QValidator_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QValidator*>(ptr)->QValidator::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QValidator_Event(void* ptr, void* e)
{
	return static_cast<QValidator*>(ptr)->event(static_cast<QEvent*>(e));
}

int QValidator_EventDefault(void* ptr, void* e)
{
	return static_cast<QValidator*>(ptr)->QValidator::event(static_cast<QEvent*>(e));
}

int QValidator_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QValidator*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QValidator_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QValidator*>(ptr)->QValidator::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QValidator_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QValidator*>(ptr)->metaObject());
}

void* QValidator_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QValidator*>(ptr)->QValidator::metaObject());
}

void* QVector2D_NewQVector2D()
{
	return new QVector2D();
}

void* QVector2D_NewQVector2D4(void* point)
{
	return new QVector2D(*static_cast<QPoint*>(point));
}

void* QVector2D_NewQVector2D5(void* point)
{
	return new QVector2D(*static_cast<QPointF*>(point));
}

void* QVector2D_NewQVector2D6(void* vector)
{
	return new QVector2D(*static_cast<QVector3D*>(vector));
}

void* QVector2D_NewQVector2D7(void* vector)
{
	return new QVector2D(*static_cast<QVector4D*>(vector));
}

int QVector2D_IsNull(void* ptr)
{
	return static_cast<QVector2D*>(ptr)->isNull();
}

void QVector2D_Normalize(void* ptr)
{
	static_cast<QVector2D*>(ptr)->normalize();
}

void* QVector2D_ToPoint(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QVector2D*>(ptr)->toPoint()).x(), static_cast<QPoint>(static_cast<QVector2D*>(ptr)->toPoint()).y());
}

void* QVector2D_ToPointF(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QVector2D*>(ptr)->toPointF()).x(), static_cast<QPointF>(static_cast<QVector2D*>(ptr)->toPointF()).y());
}

void* QVector3D_NewQVector3D()
{
	return new QVector3D();
}

void* QVector3D_NewQVector3D4(void* point)
{
	return new QVector3D(*static_cast<QPoint*>(point));
}

void* QVector3D_NewQVector3D5(void* point)
{
	return new QVector3D(*static_cast<QPointF*>(point));
}

void* QVector3D_NewQVector3D6(void* vector)
{
	return new QVector3D(*static_cast<QVector2D*>(vector));
}

void* QVector3D_NewQVector3D8(void* vector)
{
	return new QVector3D(*static_cast<QVector4D*>(vector));
}

int QVector3D_IsNull(void* ptr)
{
	return static_cast<QVector3D*>(ptr)->isNull();
}

void QVector3D_Normalize(void* ptr)
{
	static_cast<QVector3D*>(ptr)->normalize();
}

void* QVector3D_ToPoint(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QVector3D*>(ptr)->toPoint()).x(), static_cast<QPoint>(static_cast<QVector3D*>(ptr)->toPoint()).y());
}

void* QVector3D_ToPointF(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QVector3D*>(ptr)->toPointF()).x(), static_cast<QPointF>(static_cast<QVector3D*>(ptr)->toPointF()).y());
}

void* QVector4D_NewQVector4D()
{
	return new QVector4D();
}

void* QVector4D_NewQVector4D4(void* point)
{
	return new QVector4D(*static_cast<QPoint*>(point));
}

void* QVector4D_NewQVector4D5(void* point)
{
	return new QVector4D(*static_cast<QPointF*>(point));
}

void* QVector4D_NewQVector4D6(void* vector)
{
	return new QVector4D(*static_cast<QVector2D*>(vector));
}

void* QVector4D_NewQVector4D8(void* vector)
{
	return new QVector4D(*static_cast<QVector3D*>(vector));
}

int QVector4D_IsNull(void* ptr)
{
	return static_cast<QVector4D*>(ptr)->isNull();
}

void QVector4D_Normalize(void* ptr)
{
	static_cast<QVector4D*>(ptr)->normalize();
}

void* QVector4D_ToPoint(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QVector4D*>(ptr)->toPoint()).x(), static_cast<QPoint>(static_cast<QVector4D*>(ptr)->toPoint()).y());
}

void* QVector4D_ToPointF(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QVector4D*>(ptr)->toPointF()).x(), static_cast<QPointF>(static_cast<QVector4D*>(ptr)->toPointF()).y());
}

void* QWhatsThisClickedEvent_NewQWhatsThisClickedEvent(char* href)
{
	return new QWhatsThisClickedEvent(QString(href));
}

char* QWhatsThisClickedEvent_Href(void* ptr)
{
	return static_cast<QWhatsThisClickedEvent*>(ptr)->href().toUtf8().data();
}

void* QWheelEvent_NewQWheelEvent(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, int qt4Delta, int qt4Orientation, int buttons, int modifiers)
{
	return new QWheelEvent(*static_cast<QPointF*>(pos), *static_cast<QPointF*>(globalPos), *static_cast<QPoint*>(pixelDelta), *static_cast<QPoint*>(angleDelta), qt4Delta, static_cast<Qt::Orientation>(qt4Orientation), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

void* QWheelEvent_NewQWheelEvent4(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, int qt4Delta, int qt4Orientation, int buttons, int modifiers, int phase)
{
	return new QWheelEvent(*static_cast<QPointF*>(pos), *static_cast<QPointF*>(globalPos), *static_cast<QPoint*>(pixelDelta), *static_cast<QPoint*>(angleDelta), qt4Delta, static_cast<Qt::Orientation>(qt4Orientation), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<Qt::ScrollPhase>(phase));
}

void* QWheelEvent_NewQWheelEvent5(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, int qt4Delta, int qt4Orientation, int buttons, int modifiers, int phase, int source)
{
	return new QWheelEvent(*static_cast<QPointF*>(pos), *static_cast<QPointF*>(globalPos), *static_cast<QPoint*>(pixelDelta), *static_cast<QPoint*>(angleDelta), qt4Delta, static_cast<Qt::Orientation>(qt4Orientation), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<Qt::ScrollPhase>(phase), static_cast<Qt::MouseEventSource>(source));
}

void* QWheelEvent_AngleDelta(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QWheelEvent*>(ptr)->angleDelta()).x(), static_cast<QPoint>(static_cast<QWheelEvent*>(ptr)->angleDelta()).y());
}

int QWheelEvent_Buttons(void* ptr)
{
	return static_cast<QWheelEvent*>(ptr)->buttons();
}

void* QWheelEvent_GlobalPos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QWheelEvent*>(ptr)->globalPos()).x(), static_cast<QPoint>(static_cast<QWheelEvent*>(ptr)->globalPos()).y());
}

void* QWheelEvent_GlobalPosF(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QWheelEvent*>(ptr)->globalPosF()).x(), static_cast<QPointF>(static_cast<QWheelEvent*>(ptr)->globalPosF()).y());
}

int QWheelEvent_GlobalX(void* ptr)
{
	return static_cast<QWheelEvent*>(ptr)->globalX();
}

int QWheelEvent_GlobalY(void* ptr)
{
	return static_cast<QWheelEvent*>(ptr)->globalY();
}

int QWheelEvent_Phase(void* ptr)
{
	return static_cast<QWheelEvent*>(ptr)->phase();
}

void* QWheelEvent_PixelDelta(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QWheelEvent*>(ptr)->pixelDelta()).x(), static_cast<QPoint>(static_cast<QWheelEvent*>(ptr)->pixelDelta()).y());
}

void* QWheelEvent_Pos(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QWheelEvent*>(ptr)->pos()).x(), static_cast<QPoint>(static_cast<QWheelEvent*>(ptr)->pos()).y());
}

void* QWheelEvent_PosF(void* ptr)
{
	return new QPointF(static_cast<QPointF>(static_cast<QWheelEvent*>(ptr)->posF()).x(), static_cast<QPointF>(static_cast<QWheelEvent*>(ptr)->posF()).y());
}

int QWheelEvent_Source(void* ptr)
{
	return static_cast<QWheelEvent*>(ptr)->source();
}

int QWheelEvent_X(void* ptr)
{
	return static_cast<QWheelEvent*>(ptr)->x();
}

int QWheelEvent_Y(void* ptr)
{
	return static_cast<QWheelEvent*>(ptr)->y();
}

class MyQWindow: public QWindow
{
public:
	MyQWindow(QScreen *targetScreen) : QWindow(targetScreen) {};
	MyQWindow(QWindow *parent) : QWindow(parent) {};
	void setHeight(int arg) { callbackQWindow_SetHeight(this, this->objectName().toUtf8().data(), arg); };
	void setMaximumHeight(int h) { callbackQWindow_SetMaximumHeight(this, this->objectName().toUtf8().data(), h); };
	void setMaximumWidth(int w) { callbackQWindow_SetMaximumWidth(this, this->objectName().toUtf8().data(), w); };
	void setMinimumHeight(int h) { callbackQWindow_SetMinimumHeight(this, this->objectName().toUtf8().data(), h); };
	void setMinimumWidth(int w) { callbackQWindow_SetMinimumWidth(this, this->objectName().toUtf8().data(), w); };
	void setTitle(const QString & vqs) { callbackQWindow_SetTitle(this, this->objectName().toUtf8().data(), vqs.toUtf8().data()); };
	void setVisible(bool visible) { callbackQWindow_SetVisible(this, this->objectName().toUtf8().data(), visible); };
	void setWidth(int arg) { callbackQWindow_SetWidth(this, this->objectName().toUtf8().data(), arg); };
	void setX(int arg) { callbackQWindow_SetX(this, this->objectName().toUtf8().data(), arg); };
	void setY(int arg) { callbackQWindow_SetY(this, this->objectName().toUtf8().data(), arg); };
	void Signal_ActiveChanged() { callbackQWindow_ActiveChanged(this, this->objectName().toUtf8().data()); };
	void alert(int msec) { callbackQWindow_Alert(this, this->objectName().toUtf8().data(), msec); };
	bool close() { return callbackQWindow_Close(this, this->objectName().toUtf8().data()) != 0; };
	void Signal_ContentOrientationChanged(Qt::ScreenOrientation orientation) { callbackQWindow_ContentOrientationChanged(this, this->objectName().toUtf8().data(), orientation); };
	bool event(QEvent * ev) { return callbackQWindow_Event(this, this->objectName().toUtf8().data(), ev) != 0; };
	void exposeEvent(QExposeEvent * ev) { callbackQWindow_ExposeEvent(this, this->objectName().toUtf8().data(), ev); };
	void focusInEvent(QFocusEvent * ev) { callbackQWindow_FocusInEvent(this, this->objectName().toUtf8().data(), ev); };
	QObject * focusObject() const { return static_cast<QObject*>(callbackQWindow_FocusObject(const_cast<MyQWindow*>(this), this->objectName().toUtf8().data())); };
	void Signal_FocusObjectChanged(QObject * object) { callbackQWindow_FocusObjectChanged(this, this->objectName().toUtf8().data(), object); };
	void focusOutEvent(QFocusEvent * ev) { callbackQWindow_FocusOutEvent(this, this->objectName().toUtf8().data(), ev); };
	QSurfaceFormat format() const { return *static_cast<QSurfaceFormat*>(callbackQWindow_Format(const_cast<MyQWindow*>(this), this->objectName().toUtf8().data())); };
	void Signal_HeightChanged(int arg) { callbackQWindow_HeightChanged(this, this->objectName().toUtf8().data(), arg); };
	void hide() { callbackQWindow_Hide(this, this->objectName().toUtf8().data()); };
	void hideEvent(QHideEvent * ev) { callbackQWindow_HideEvent(this, this->objectName().toUtf8().data(), ev); };
	void keyPressEvent(QKeyEvent * ev) { callbackQWindow_KeyPressEvent(this, this->objectName().toUtf8().data(), ev); };
	void keyReleaseEvent(QKeyEvent * ev) { callbackQWindow_KeyReleaseEvent(this, this->objectName().toUtf8().data(), ev); };
	void lower() { callbackQWindow_Lower(this, this->objectName().toUtf8().data()); };
	void Signal_MaximumHeightChanged(int arg) { callbackQWindow_MaximumHeightChanged(this, this->objectName().toUtf8().data(), arg); };
	void Signal_MaximumWidthChanged(int arg) { callbackQWindow_MaximumWidthChanged(this, this->objectName().toUtf8().data(), arg); };
	void Signal_MinimumHeightChanged(int arg) { callbackQWindow_MinimumHeightChanged(this, this->objectName().toUtf8().data(), arg); };
	void Signal_MinimumWidthChanged(int arg) { callbackQWindow_MinimumWidthChanged(this, this->objectName().toUtf8().data(), arg); };
	void Signal_ModalityChanged(Qt::WindowModality modality) { callbackQWindow_ModalityChanged(this, this->objectName().toUtf8().data(), modality); };
	void mouseDoubleClickEvent(QMouseEvent * ev) { callbackQWindow_MouseDoubleClickEvent(this, this->objectName().toUtf8().data(), ev); };
	void mouseMoveEvent(QMouseEvent * ev) { callbackQWindow_MouseMoveEvent(this, this->objectName().toUtf8().data(), ev); };
	void mousePressEvent(QMouseEvent * ev) { callbackQWindow_MousePressEvent(this, this->objectName().toUtf8().data(), ev); };
	void mouseReleaseEvent(QMouseEvent * ev) { callbackQWindow_MouseReleaseEvent(this, this->objectName().toUtf8().data(), ev); };
	void moveEvent(QMoveEvent * ev) { callbackQWindow_MoveEvent(this, this->objectName().toUtf8().data(), ev); };
	bool nativeEvent(const QByteArray & eventType, void * message, long * result) { return callbackQWindow_NativeEvent(this, this->objectName().toUtf8().data(), QString(eventType).toUtf8().data(), message, *result) != 0; };
	void Signal_OpacityChanged(qreal opacity) { callbackQWindow_OpacityChanged(this, this->objectName().toUtf8().data(), static_cast<double>(opacity)); };
	void raise() { callbackQWindow_Raise(this, this->objectName().toUtf8().data()); };
	void requestActivate() { callbackQWindow_RequestActivate(this, this->objectName().toUtf8().data()); };
	void requestUpdate() { callbackQWindow_RequestUpdate(this, this->objectName().toUtf8().data()); };
	void resizeEvent(QResizeEvent * ev) { callbackQWindow_ResizeEvent(this, this->objectName().toUtf8().data(), ev); };
	void Signal_ScreenChanged(QScreen * screen) { callbackQWindow_ScreenChanged(this, this->objectName().toUtf8().data(), screen); };
	void show() { callbackQWindow_Show(this, this->objectName().toUtf8().data()); };
	void showEvent(QShowEvent * ev) { callbackQWindow_ShowEvent(this, this->objectName().toUtf8().data(), ev); };
	void showFullScreen() { callbackQWindow_ShowFullScreen(this, this->objectName().toUtf8().data()); };
	void showMaximized() { callbackQWindow_ShowMaximized(this, this->objectName().toUtf8().data()); };
	void showMinimized() { callbackQWindow_ShowMinimized(this, this->objectName().toUtf8().data()); };
	void showNormal() { callbackQWindow_ShowNormal(this, this->objectName().toUtf8().data()); };
	QSize size() const { return *static_cast<QSize*>(callbackQWindow_Size(const_cast<MyQWindow*>(this), this->objectName().toUtf8().data())); };
	SurfaceType surfaceType() const { return static_cast<QSurface::SurfaceType>(callbackQWindow_SurfaceType(const_cast<MyQWindow*>(this), this->objectName().toUtf8().data())); };
	void tabletEvent(QTabletEvent * ev) { callbackQWindow_TabletEvent(this, this->objectName().toUtf8().data(), ev); };
	void touchEvent(QTouchEvent * ev) { callbackQWindow_TouchEvent(this, this->objectName().toUtf8().data(), ev); };
	void Signal_VisibilityChanged(QWindow::Visibility visibility) { callbackQWindow_VisibilityChanged(this, this->objectName().toUtf8().data(), visibility); };
	void Signal_VisibleChanged(bool arg) { callbackQWindow_VisibleChanged(this, this->objectName().toUtf8().data(), arg); };
	void wheelEvent(QWheelEvent * ev) { callbackQWindow_WheelEvent(this, this->objectName().toUtf8().data(), ev); };
	void Signal_WidthChanged(int arg) { callbackQWindow_WidthChanged(this, this->objectName().toUtf8().data(), arg); };
	void Signal_WindowStateChanged(Qt::WindowState windowState) { callbackQWindow_WindowStateChanged(this, this->objectName().toUtf8().data(), windowState); };
	void Signal_WindowTitleChanged(const QString & title) { callbackQWindow_WindowTitleChanged(this, this->objectName().toUtf8().data(), title.toUtf8().data()); };
	void Signal_XChanged(int arg) { callbackQWindow_XChanged(this, this->objectName().toUtf8().data(), arg); };
	void Signal_YChanged(int arg) { callbackQWindow_YChanged(this, this->objectName().toUtf8().data(), arg); };
	void timerEvent(QTimerEvent * event) { callbackQWindow_TimerEvent(this, this->objectName().toUtf8().data(), event); };
	void childEvent(QChildEvent * event) { callbackQWindow_ChildEvent(this, this->objectName().toUtf8().data(), event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWindow_ConnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	void customEvent(QEvent * event) { callbackQWindow_CustomEvent(this, this->objectName().toUtf8().data(), event); };
	void deleteLater() { callbackQWindow_DeleteLater(this, this->objectName().toUtf8().data()); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWindow_DisconnectNotify(this, this->objectName().toUtf8().data(), new QMetaMethod(sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWindow_EventFilter(this, this->objectName().toUtf8().data(), watched, event) != 0; };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQWindow_MetaObject(const_cast<MyQWindow*>(this), this->objectName().toUtf8().data())); };
};

int QWindow_ContentOrientation(void* ptr)
{
	return static_cast<QWindow*>(ptr)->contentOrientation();
}

int QWindow_Flags(void* ptr)
{
	return static_cast<QWindow*>(ptr)->flags();
}

int QWindow_IsVisible(void* ptr)
{
	return static_cast<QWindow*>(ptr)->isVisible();
}

void* QWindow_MapFromGlobal(void* ptr, void* pos)
{
	return new QPoint(static_cast<QPoint>(static_cast<QWindow*>(ptr)->mapFromGlobal(*static_cast<QPoint*>(pos))).x(), static_cast<QPoint>(static_cast<QWindow*>(ptr)->mapFromGlobal(*static_cast<QPoint*>(pos))).y());
}

void* QWindow_MapToGlobal(void* ptr, void* pos)
{
	return new QPoint(static_cast<QPoint>(static_cast<QWindow*>(ptr)->mapToGlobal(*static_cast<QPoint*>(pos))).x(), static_cast<QPoint>(static_cast<QWindow*>(ptr)->mapToGlobal(*static_cast<QPoint*>(pos))).y());
}

int QWindow_Modality(void* ptr)
{
	return static_cast<QWindow*>(ptr)->modality();
}

double QWindow_Opacity(void* ptr)
{
	return static_cast<double>(static_cast<QWindow*>(ptr)->opacity());
}

void QWindow_ReportContentOrientationChange(void* ptr, int orientation)
{
	static_cast<QWindow*>(ptr)->reportContentOrientationChange(static_cast<Qt::ScreenOrientation>(orientation));
}

void QWindow_SetFlags(void* ptr, int flags)
{
	static_cast<QWindow*>(ptr)->setFlags(static_cast<Qt::WindowType>(flags));
}

void QWindow_SetHeight(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setHeight", Q_ARG(int, arg));
}

void QWindow_SetMaximumHeight(void* ptr, int h)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setMaximumHeight", Q_ARG(int, h));
}

void QWindow_SetMaximumWidth(void* ptr, int w)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setMaximumWidth", Q_ARG(int, w));
}

void QWindow_SetMinimumHeight(void* ptr, int h)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setMinimumHeight", Q_ARG(int, h));
}

void QWindow_SetMinimumWidth(void* ptr, int w)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setMinimumWidth", Q_ARG(int, w));
}

void QWindow_SetModality(void* ptr, int modality)
{
	static_cast<QWindow*>(ptr)->setModality(static_cast<Qt::WindowModality>(modality));
}

void QWindow_SetOpacity(void* ptr, double level)
{
	static_cast<QWindow*>(ptr)->setOpacity(static_cast<double>(level));
}

void QWindow_SetTitle(void* ptr, char* vqs)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setTitle", Q_ARG(QString, QString(vqs)));
}

void QWindow_SetVisibility(void* ptr, int v)
{
	static_cast<QWindow*>(ptr)->setVisibility(static_cast<QWindow::Visibility>(v));
}

void QWindow_SetVisible(void* ptr, int visible)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setVisible", Q_ARG(bool, visible != 0));
}

void QWindow_SetWidth(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setWidth", Q_ARG(int, arg));
}

void QWindow_SetX(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setX", Q_ARG(int, arg));
}

void QWindow_SetY(void* ptr, int arg)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setY", Q_ARG(int, arg));
}

char* QWindow_Title(void* ptr)
{
	return static_cast<QWindow*>(ptr)->title().toUtf8().data();
}

int QWindow_Visibility(void* ptr)
{
	return static_cast<QWindow*>(ptr)->visibility();
}

void* QWindow_NewQWindow(void* targetScreen)
{
	return new MyQWindow(static_cast<QScreen*>(targetScreen));
}

void* QWindow_NewQWindow2(void* parent)
{
	return new MyQWindow(static_cast<QWindow*>(parent));
}

void QWindow_ConnectActiveChanged(void* ptr)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)()>(&QWindow::activeChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)()>(&MyQWindow::Signal_ActiveChanged));
}

void QWindow_DisconnectActiveChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)()>(&QWindow::activeChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)()>(&MyQWindow::Signal_ActiveChanged));
}

void QWindow_ActiveChanged(void* ptr)
{
	static_cast<QWindow*>(ptr)->activeChanged();
}

void QWindow_Alert(void* ptr, int msec)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "alert", Q_ARG(int, msec));
}

void* QWindow_BaseSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QWindow*>(ptr)->baseSize()).width(), static_cast<QSize>(static_cast<QWindow*>(ptr)->baseSize()).height());
}

int QWindow_Close(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "close", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

void QWindow_ConnectContentOrientationChanged(void* ptr)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(Qt::ScreenOrientation)>(&QWindow::contentOrientationChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(Qt::ScreenOrientation)>(&MyQWindow::Signal_ContentOrientationChanged));
}

void QWindow_DisconnectContentOrientationChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(Qt::ScreenOrientation)>(&QWindow::contentOrientationChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(Qt::ScreenOrientation)>(&MyQWindow::Signal_ContentOrientationChanged));
}

void QWindow_ContentOrientationChanged(void* ptr, int orientation)
{
	static_cast<QWindow*>(ptr)->contentOrientationChanged(static_cast<Qt::ScreenOrientation>(orientation));
}

void QWindow_Create(void* ptr)
{
	static_cast<QWindow*>(ptr)->create();
}

void* QWindow_Cursor(void* ptr)
{
	return new QCursor(static_cast<QWindow*>(ptr)->cursor());
}

void QWindow_Destroy(void* ptr)
{
	static_cast<QWindow*>(ptr)->destroy();
}

double QWindow_DevicePixelRatio(void* ptr)
{
	return static_cast<double>(static_cast<QWindow*>(ptr)->devicePixelRatio());
}

int QWindow_Event(void* ptr, void* ev)
{
	return static_cast<QWindow*>(ptr)->event(static_cast<QEvent*>(ev));
}

int QWindow_EventDefault(void* ptr, void* ev)
{
	return static_cast<QWindow*>(ptr)->QWindow::event(static_cast<QEvent*>(ev));
}

void QWindow_ExposeEvent(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->exposeEvent(static_cast<QExposeEvent*>(ev));
}

void QWindow_ExposeEventDefault(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->QWindow::exposeEvent(static_cast<QExposeEvent*>(ev));
}

char* QWindow_FilePath(void* ptr)
{
	return static_cast<QWindow*>(ptr)->filePath().toUtf8().data();
}

void QWindow_FocusInEvent(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(ev));
}

void QWindow_FocusInEventDefault(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->QWindow::focusInEvent(static_cast<QFocusEvent*>(ev));
}

void* QWindow_FocusObject(void* ptr)
{
	return static_cast<QWindow*>(ptr)->focusObject();
}

void* QWindow_FocusObjectDefault(void* ptr)
{
	return static_cast<QWindow*>(ptr)->QWindow::focusObject();
}

void QWindow_ConnectFocusObjectChanged(void* ptr)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(QObject *)>(&QWindow::focusObjectChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(QObject *)>(&MyQWindow::Signal_FocusObjectChanged));
}

void QWindow_DisconnectFocusObjectChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(QObject *)>(&QWindow::focusObjectChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(QObject *)>(&MyQWindow::Signal_FocusObjectChanged));
}

void QWindow_FocusObjectChanged(void* ptr, void* object)
{
	static_cast<QWindow*>(ptr)->focusObjectChanged(static_cast<QObject*>(object));
}

void QWindow_FocusOutEvent(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void QWindow_FocusOutEventDefault(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->QWindow::focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void* QWindow_Format(void* ptr)
{
	return new QSurfaceFormat(static_cast<QWindow*>(ptr)->format());
}

void* QWindow_FormatDefault(void* ptr)
{
	return new QSurfaceFormat(static_cast<QWindow*>(ptr)->QWindow::format());
}

void* QWindow_FrameGeometry(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QWindow*>(ptr)->frameGeometry()).x(), static_cast<QRect>(static_cast<QWindow*>(ptr)->frameGeometry()).y(), static_cast<QRect>(static_cast<QWindow*>(ptr)->frameGeometry()).width(), static_cast<QRect>(static_cast<QWindow*>(ptr)->frameGeometry()).height());
}

void* QWindow_FrameMargins(void* ptr)
{
	return new QMargins(static_cast<QMargins>(static_cast<QWindow*>(ptr)->frameMargins()).left(), static_cast<QMargins>(static_cast<QWindow*>(ptr)->frameMargins()).top(), static_cast<QMargins>(static_cast<QWindow*>(ptr)->frameMargins()).right(), static_cast<QMargins>(static_cast<QWindow*>(ptr)->frameMargins()).bottom());
}

void* QWindow_FramePosition(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QWindow*>(ptr)->framePosition()).x(), static_cast<QPoint>(static_cast<QWindow*>(ptr)->framePosition()).y());
}

void* QWindow_QWindow_FromWinId(unsigned long long id)
{
	return QWindow::fromWinId(static_cast<unsigned long long>(id));
}

void* QWindow_Geometry(void* ptr)
{
	return new QRect(static_cast<QRect>(static_cast<QWindow*>(ptr)->geometry()).x(), static_cast<QRect>(static_cast<QWindow*>(ptr)->geometry()).y(), static_cast<QRect>(static_cast<QWindow*>(ptr)->geometry()).width(), static_cast<QRect>(static_cast<QWindow*>(ptr)->geometry()).height());
}

int QWindow_Height(void* ptr)
{
	return static_cast<QWindow*>(ptr)->height();
}

void QWindow_ConnectHeightChanged(void* ptr)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::heightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_HeightChanged));
}

void QWindow_DisconnectHeightChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::heightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_HeightChanged));
}

void QWindow_HeightChanged(void* ptr, int arg)
{
	static_cast<QWindow*>(ptr)->heightChanged(arg);
}

void QWindow_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "hide");
}

void QWindow_HideEvent(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->hideEvent(static_cast<QHideEvent*>(ev));
}

void QWindow_HideEventDefault(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->QWindow::hideEvent(static_cast<QHideEvent*>(ev));
}

void* QWindow_Icon(void* ptr)
{
	return new QIcon(static_cast<QWindow*>(ptr)->icon());
}

int QWindow_IsActive(void* ptr)
{
	return static_cast<QWindow*>(ptr)->isActive();
}

int QWindow_IsAncestorOf(void* ptr, void* child, int mode)
{
	return static_cast<QWindow*>(ptr)->isAncestorOf(static_cast<QWindow*>(child), static_cast<QWindow::AncestorMode>(mode));
}

int QWindow_IsExposed(void* ptr)
{
	return static_cast<QWindow*>(ptr)->isExposed();
}

int QWindow_IsModal(void* ptr)
{
	return static_cast<QWindow*>(ptr)->isModal();
}

int QWindow_IsTopLevel(void* ptr)
{
	return static_cast<QWindow*>(ptr)->isTopLevel();
}

void QWindow_KeyPressEvent(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(ev));
}

void QWindow_KeyPressEventDefault(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->QWindow::keyPressEvent(static_cast<QKeyEvent*>(ev));
}

void QWindow_KeyReleaseEvent(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(ev));
}

void QWindow_KeyReleaseEventDefault(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->QWindow::keyReleaseEvent(static_cast<QKeyEvent*>(ev));
}

void QWindow_Lower(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "lower");
}

void* QWindow_Mask(void* ptr)
{
	return new QRegion(static_cast<QWindow*>(ptr)->mask());
}

int QWindow_MaximumHeight(void* ptr)
{
	return static_cast<QWindow*>(ptr)->maximumHeight();
}

void QWindow_ConnectMaximumHeightChanged(void* ptr)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::maximumHeightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MaximumHeightChanged));
}

void QWindow_DisconnectMaximumHeightChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::maximumHeightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MaximumHeightChanged));
}

void QWindow_MaximumHeightChanged(void* ptr, int arg)
{
	static_cast<QWindow*>(ptr)->maximumHeightChanged(arg);
}

void* QWindow_MaximumSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QWindow*>(ptr)->maximumSize()).width(), static_cast<QSize>(static_cast<QWindow*>(ptr)->maximumSize()).height());
}

int QWindow_MaximumWidth(void* ptr)
{
	return static_cast<QWindow*>(ptr)->maximumWidth();
}

void QWindow_ConnectMaximumWidthChanged(void* ptr)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::maximumWidthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MaximumWidthChanged));
}

void QWindow_DisconnectMaximumWidthChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::maximumWidthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MaximumWidthChanged));
}

void QWindow_MaximumWidthChanged(void* ptr, int arg)
{
	static_cast<QWindow*>(ptr)->maximumWidthChanged(arg);
}

int QWindow_MinimumHeight(void* ptr)
{
	return static_cast<QWindow*>(ptr)->minimumHeight();
}

void QWindow_ConnectMinimumHeightChanged(void* ptr)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::minimumHeightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MinimumHeightChanged));
}

void QWindow_DisconnectMinimumHeightChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::minimumHeightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MinimumHeightChanged));
}

void QWindow_MinimumHeightChanged(void* ptr, int arg)
{
	static_cast<QWindow*>(ptr)->minimumHeightChanged(arg);
}

void* QWindow_MinimumSize(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QWindow*>(ptr)->minimumSize()).width(), static_cast<QSize>(static_cast<QWindow*>(ptr)->minimumSize()).height());
}

int QWindow_MinimumWidth(void* ptr)
{
	return static_cast<QWindow*>(ptr)->minimumWidth();
}

void QWindow_ConnectMinimumWidthChanged(void* ptr)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::minimumWidthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MinimumWidthChanged));
}

void QWindow_DisconnectMinimumWidthChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::minimumWidthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_MinimumWidthChanged));
}

void QWindow_MinimumWidthChanged(void* ptr, int arg)
{
	static_cast<QWindow*>(ptr)->minimumWidthChanged(arg);
}

void QWindow_ConnectModalityChanged(void* ptr)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(Qt::WindowModality)>(&QWindow::modalityChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(Qt::WindowModality)>(&MyQWindow::Signal_ModalityChanged));
}

void QWindow_DisconnectModalityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(Qt::WindowModality)>(&QWindow::modalityChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(Qt::WindowModality)>(&MyQWindow::Signal_ModalityChanged));
}

void QWindow_ModalityChanged(void* ptr, int modality)
{
	static_cast<QWindow*>(ptr)->modalityChanged(static_cast<Qt::WindowModality>(modality));
}

void QWindow_MouseDoubleClickEvent(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MouseDoubleClickEventDefault(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->QWindow::mouseDoubleClickEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MouseMoveEvent(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MouseMoveEventDefault(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->QWindow::mouseMoveEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MousePressEvent(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MousePressEventDefault(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->QWindow::mousePressEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MouseReleaseEvent(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MouseReleaseEventDefault(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->QWindow::mouseReleaseEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MoveEvent(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->moveEvent(static_cast<QMoveEvent*>(ev));
}

void QWindow_MoveEventDefault(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->QWindow::moveEvent(static_cast<QMoveEvent*>(ev));
}

int QWindow_NativeEvent(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QWindow*>(ptr)->nativeEvent(QByteArray(eventType), message, &result);
}

int QWindow_NativeEventDefault(void* ptr, char* eventType, void* message, long result)
{
	return static_cast<QWindow*>(ptr)->QWindow::nativeEvent(QByteArray(eventType), message, &result);
}

void QWindow_ConnectOpacityChanged(void* ptr)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(qreal)>(&QWindow::opacityChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(qreal)>(&MyQWindow::Signal_OpacityChanged));
}

void QWindow_DisconnectOpacityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(qreal)>(&QWindow::opacityChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(qreal)>(&MyQWindow::Signal_OpacityChanged));
}

void QWindow_OpacityChanged(void* ptr, double opacity)
{
	static_cast<QWindow*>(ptr)->opacityChanged(static_cast<double>(opacity));
}

void* QWindow_Parent(void* ptr)
{
	return static_cast<QWindow*>(ptr)->parent();
}

void* QWindow_Position(void* ptr)
{
	return new QPoint(static_cast<QPoint>(static_cast<QWindow*>(ptr)->position()).x(), static_cast<QPoint>(static_cast<QWindow*>(ptr)->position()).y());
}

void QWindow_Raise(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "raise");
}

void QWindow_RequestActivate(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "requestActivate");
}

void QWindow_RequestUpdate(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "requestUpdate");
}

void* QWindow_RequestedFormat(void* ptr)
{
	return new QSurfaceFormat(static_cast<QWindow*>(ptr)->requestedFormat());
}

void QWindow_Resize(void* ptr, void* newSize)
{
	static_cast<QWindow*>(ptr)->resize(*static_cast<QSize*>(newSize));
}

void QWindow_Resize2(void* ptr, int w, int h)
{
	static_cast<QWindow*>(ptr)->resize(w, h);
}

void QWindow_ResizeEvent(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->resizeEvent(static_cast<QResizeEvent*>(ev));
}

void QWindow_ResizeEventDefault(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->QWindow::resizeEvent(static_cast<QResizeEvent*>(ev));
}

void* QWindow_Screen(void* ptr)
{
	return static_cast<QWindow*>(ptr)->screen();
}

void QWindow_ConnectScreenChanged(void* ptr)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(QScreen *)>(&QWindow::screenChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(QScreen *)>(&MyQWindow::Signal_ScreenChanged));
}

void QWindow_DisconnectScreenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(QScreen *)>(&QWindow::screenChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(QScreen *)>(&MyQWindow::Signal_ScreenChanged));
}

void QWindow_ScreenChanged(void* ptr, void* screen)
{
	static_cast<QWindow*>(ptr)->screenChanged(static_cast<QScreen*>(screen));
}

void QWindow_SetBaseSize(void* ptr, void* size)
{
	static_cast<QWindow*>(ptr)->setBaseSize(*static_cast<QSize*>(size));
}

void QWindow_SetCursor(void* ptr, void* cursor)
{
	static_cast<QWindow*>(ptr)->setCursor(*static_cast<QCursor*>(cursor));
}

void QWindow_SetFilePath(void* ptr, char* filePath)
{
	static_cast<QWindow*>(ptr)->setFilePath(QString(filePath));
}

void QWindow_SetFormat(void* ptr, void* format)
{
	static_cast<QWindow*>(ptr)->setFormat(*static_cast<QSurfaceFormat*>(format));
}

void QWindow_SetFramePosition(void* ptr, void* point)
{
	static_cast<QWindow*>(ptr)->setFramePosition(*static_cast<QPoint*>(point));
}

void QWindow_SetGeometry2(void* ptr, void* rect)
{
	static_cast<QWindow*>(ptr)->setGeometry(*static_cast<QRect*>(rect));
}

void QWindow_SetGeometry(void* ptr, int posx, int posy, int w, int h)
{
	static_cast<QWindow*>(ptr)->setGeometry(posx, posy, w, h);
}

void QWindow_SetIcon(void* ptr, void* icon)
{
	static_cast<QWindow*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

int QWindow_SetKeyboardGrabEnabled(void* ptr, int grab)
{
	return static_cast<QWindow*>(ptr)->setKeyboardGrabEnabled(grab != 0);
}

void QWindow_SetMask(void* ptr, void* region)
{
	static_cast<QWindow*>(ptr)->setMask(*static_cast<QRegion*>(region));
}

void QWindow_SetMaximumSize(void* ptr, void* size)
{
	static_cast<QWindow*>(ptr)->setMaximumSize(*static_cast<QSize*>(size));
}

void QWindow_SetMinimumSize(void* ptr, void* size)
{
	static_cast<QWindow*>(ptr)->setMinimumSize(*static_cast<QSize*>(size));
}

int QWindow_SetMouseGrabEnabled(void* ptr, int grab)
{
	return static_cast<QWindow*>(ptr)->setMouseGrabEnabled(grab != 0);
}

void QWindow_SetParent(void* ptr, void* parent)
{
	static_cast<QWindow*>(ptr)->setParent(static_cast<QWindow*>(parent));
}

void QWindow_SetPosition(void* ptr, void* pt)
{
	static_cast<QWindow*>(ptr)->setPosition(*static_cast<QPoint*>(pt));
}

void QWindow_SetPosition2(void* ptr, int posx, int posy)
{
	static_cast<QWindow*>(ptr)->setPosition(posx, posy);
}

void QWindow_SetScreen(void* ptr, void* newScreen)
{
	static_cast<QWindow*>(ptr)->setScreen(static_cast<QScreen*>(newScreen));
}

void QWindow_SetSizeIncrement(void* ptr, void* size)
{
	static_cast<QWindow*>(ptr)->setSizeIncrement(*static_cast<QSize*>(size));
}

void QWindow_SetSurfaceType(void* ptr, int surfaceType)
{
	static_cast<QWindow*>(ptr)->setSurfaceType(static_cast<QSurface::SurfaceType>(surfaceType));
}

void QWindow_SetTransientParent(void* ptr, void* parent)
{
	static_cast<QWindow*>(ptr)->setTransientParent(static_cast<QWindow*>(parent));
}

void QWindow_SetWindowState(void* ptr, int state)
{
	static_cast<QWindow*>(ptr)->setWindowState(static_cast<Qt::WindowState>(state));
}

void QWindow_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "show");
}

void QWindow_ShowEvent(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->showEvent(static_cast<QShowEvent*>(ev));
}

void QWindow_ShowEventDefault(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->QWindow::showEvent(static_cast<QShowEvent*>(ev));
}

void QWindow_ShowFullScreen(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "showFullScreen");
}

void QWindow_ShowMaximized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "showMaximized");
}

void QWindow_ShowMinimized(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "showMinimized");
}

void QWindow_ShowNormal(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "showNormal");
}

void* QWindow_Size(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QWindow*>(ptr)->size()).width(), static_cast<QSize>(static_cast<QWindow*>(ptr)->size()).height());
}

void* QWindow_SizeDefault(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QWindow*>(ptr)->QWindow::size()).width(), static_cast<QSize>(static_cast<QWindow*>(ptr)->QWindow::size()).height());
}

void* QWindow_SizeIncrement(void* ptr)
{
	return new QSize(static_cast<QSize>(static_cast<QWindow*>(ptr)->sizeIncrement()).width(), static_cast<QSize>(static_cast<QWindow*>(ptr)->sizeIncrement()).height());
}

int QWindow_SurfaceType(void* ptr)
{
	return static_cast<QWindow*>(ptr)->surfaceType();
}

int QWindow_SurfaceTypeDefault(void* ptr)
{
	return static_cast<QWindow*>(ptr)->QWindow::surfaceType();
}

void QWindow_TabletEvent(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->tabletEvent(static_cast<QTabletEvent*>(ev));
}

void QWindow_TabletEventDefault(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->QWindow::tabletEvent(static_cast<QTabletEvent*>(ev));
}

void QWindow_TouchEvent(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->touchEvent(static_cast<QTouchEvent*>(ev));
}

void QWindow_TouchEventDefault(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->QWindow::touchEvent(static_cast<QTouchEvent*>(ev));
}

void* QWindow_TransientParent(void* ptr)
{
	return static_cast<QWindow*>(ptr)->transientParent();
}

int QWindow_Type(void* ptr)
{
	return static_cast<QWindow*>(ptr)->type();
}

void QWindow_UnsetCursor(void* ptr)
{
	static_cast<QWindow*>(ptr)->unsetCursor();
}

void QWindow_ConnectVisibilityChanged(void* ptr)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(QWindow::Visibility)>(&QWindow::visibilityChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(QWindow::Visibility)>(&MyQWindow::Signal_VisibilityChanged));
}

void QWindow_DisconnectVisibilityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(QWindow::Visibility)>(&QWindow::visibilityChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(QWindow::Visibility)>(&MyQWindow::Signal_VisibilityChanged));
}

void QWindow_VisibilityChanged(void* ptr, int visibility)
{
	static_cast<QWindow*>(ptr)->visibilityChanged(static_cast<QWindow::Visibility>(visibility));
}

void QWindow_ConnectVisibleChanged(void* ptr)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(bool)>(&QWindow::visibleChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(bool)>(&MyQWindow::Signal_VisibleChanged));
}

void QWindow_DisconnectVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(bool)>(&QWindow::visibleChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(bool)>(&MyQWindow::Signal_VisibleChanged));
}

void QWindow_VisibleChanged(void* ptr, int arg)
{
	static_cast<QWindow*>(ptr)->visibleChanged(arg != 0);
}

void QWindow_WheelEvent(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(ev));
}

void QWindow_WheelEventDefault(void* ptr, void* ev)
{
	static_cast<QWindow*>(ptr)->QWindow::wheelEvent(static_cast<QWheelEvent*>(ev));
}

int QWindow_Width(void* ptr)
{
	return static_cast<QWindow*>(ptr)->width();
}

void QWindow_ConnectWidthChanged(void* ptr)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::widthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_WidthChanged));
}

void QWindow_DisconnectWidthChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::widthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_WidthChanged));
}

void QWindow_WidthChanged(void* ptr, int arg)
{
	static_cast<QWindow*>(ptr)->widthChanged(arg);
}

unsigned long long QWindow_WinId(void* ptr)
{
	return static_cast<unsigned long long>(static_cast<QWindow*>(ptr)->winId());
}

int QWindow_WindowState(void* ptr)
{
	return static_cast<QWindow*>(ptr)->windowState();
}

void QWindow_ConnectWindowStateChanged(void* ptr)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(Qt::WindowState)>(&QWindow::windowStateChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(Qt::WindowState)>(&MyQWindow::Signal_WindowStateChanged));
}

void QWindow_DisconnectWindowStateChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(Qt::WindowState)>(&QWindow::windowStateChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(Qt::WindowState)>(&MyQWindow::Signal_WindowStateChanged));
}

void QWindow_WindowStateChanged(void* ptr, int windowState)
{
	static_cast<QWindow*>(ptr)->windowStateChanged(static_cast<Qt::WindowState>(windowState));
}

void QWindow_ConnectWindowTitleChanged(void* ptr)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(const QString &)>(&QWindow::windowTitleChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(const QString &)>(&MyQWindow::Signal_WindowTitleChanged));
}

void QWindow_DisconnectWindowTitleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(const QString &)>(&QWindow::windowTitleChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(const QString &)>(&MyQWindow::Signal_WindowTitleChanged));
}

void QWindow_WindowTitleChanged(void* ptr, char* title)
{
	static_cast<QWindow*>(ptr)->windowTitleChanged(QString(title));
}

int QWindow_X(void* ptr)
{
	return static_cast<QWindow*>(ptr)->x();
}

void QWindow_ConnectXChanged(void* ptr)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::xChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_XChanged));
}

void QWindow_DisconnectXChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::xChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_XChanged));
}

void QWindow_XChanged(void* ptr, int arg)
{
	static_cast<QWindow*>(ptr)->xChanged(arg);
}

int QWindow_Y(void* ptr)
{
	return static_cast<QWindow*>(ptr)->y();
}

void QWindow_ConnectYChanged(void* ptr)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::yChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_YChanged));
}

void QWindow_DisconnectYChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::yChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_YChanged));
}

void QWindow_YChanged(void* ptr, int arg)
{
	static_cast<QWindow*>(ptr)->yChanged(arg);
}

void QWindow_DestroyQWindow(void* ptr)
{
	static_cast<QWindow*>(ptr)->~QWindow();
}

void QWindow_TimerEvent(void* ptr, void* event)
{
	static_cast<QWindow*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QWindow_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QWindow*>(ptr)->QWindow::timerEvent(static_cast<QTimerEvent*>(event));
}

void QWindow_ChildEvent(void* ptr, void* event)
{
	static_cast<QWindow*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QWindow_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QWindow*>(ptr)->QWindow::childEvent(static_cast<QChildEvent*>(event));
}

void QWindow_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QWindow*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWindow_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWindow*>(ptr)->QWindow::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWindow_CustomEvent(void* ptr, void* event)
{
	static_cast<QWindow*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QWindow_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QWindow*>(ptr)->QWindow::customEvent(static_cast<QEvent*>(event));
}

void QWindow_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "deleteLater");
}

void QWindow_DeleteLaterDefault(void* ptr)
{
	static_cast<QWindow*>(ptr)->QWindow::deleteLater();
}

void QWindow_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QWindow*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWindow_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWindow*>(ptr)->QWindow::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

int QWindow_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QWindow*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

int QWindow_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QWindow*>(ptr)->QWindow::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void* QWindow_MetaObject(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QWindow*>(ptr)->metaObject());
}

void* QWindow_MetaObjectDefault(void* ptr)
{
	return const_cast<QMetaObject*>(static_cast<QWindow*>(ptr)->QWindow::metaObject());
}

int QWindowStateChangeEvent_OldState(void* ptr)
{
	return static_cast<QWindowStateChangeEvent*>(ptr)->oldState();
}

