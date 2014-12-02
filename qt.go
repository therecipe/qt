package qt

//#include "qt.h"
import "C"

type AlignmentFlag int

var (
	ALIGNLEFT            = AlignmentFlag(C.Qt_AlignLeft())
	ALIGNRIGHT           = AlignmentFlag(C.Qt_AlignRight())
	ALIGNHCENTER         = AlignmentFlag(C.Qt_AlignHCenter())
	ALIGNJUSTIFY         = AlignmentFlag(C.Qt_AlignJustify())
	ALIGNTOP             = AlignmentFlag(C.Qt_AlignTop())
	ALIGNBOTTOM          = AlignmentFlag(C.Qt_AlignBottom())
	ALIGNVCENTER         = AlignmentFlag(C.Qt_AlignVCenter())
	ALIGNBASELINE        = AlignmentFlag(C.Qt_AlignBaseline())
	ALIGNCENTER          = AlignmentFlag(C.Qt_AlignCenter())
	ALIGNABSOLUTE        = AlignmentFlag(C.Qt_AlignAbsolute())
	ALIGNLEADING         = AlignmentFlag(C.Qt_AlignLeading())
	ALIGNTRAILING        = AlignmentFlag(C.Qt_AlignTrailing())
	ALIGNHORIZONTAL_MASK = AlignmentFlag(C.Qt_AlignHorizontal_Mask())
	ALIGNVERTICAL_MASK   = AlignmentFlag(C.Qt_AlignVertical_Mask())
)

type AnchorPoint int

var (
	ANCHORLEFT             = AnchorPoint(C.Qt_AnchorLeft())
	ANCHORHORIZONTALCENTER = AnchorPoint(C.Qt_AnchorHorizontalCenter())
	ANCHORRIGHT            = AnchorPoint(C.Qt_AnchorRight())
	ANCHORTOP              = AnchorPoint(C.Qt_AnchorTop())
	ANCHORVERTICALCENTER   = AnchorPoint(C.Qt_AnchorVerticalCenter())
	ANCHORBOTTOM           = AnchorPoint(C.Qt_AnchorBottom())
)

type ApplicationAttribute int

var (
	AA_DONTSHOWICONSINMENUS                   = ApplicationAttribute(C.Qt_AA_DontShowIconsInMenus())
	AA_NATIVEWINDOWS                          = ApplicationAttribute(C.Qt_AA_NativeWindows())
	AA_DONTCREATENATIVEWIDGETSIBLINGS         = ApplicationAttribute(C.Qt_AA_DontCreateNativeWidgetSiblings())
	AA_MACPLUGINAPPLICATION                   = ApplicationAttribute(C.Qt_AA_MacPluginApplication())
	AA_DONTUSENATIVEMENUBAR                   = ApplicationAttribute(C.Qt_AA_DontUseNativeMenuBar())
	AA_MACDONTSWAPCTRLANDMETA                 = ApplicationAttribute(C.Qt_AA_MacDontSwapCtrlAndMeta())
	AA_USE96DPI                               = ApplicationAttribute(C.Qt_AA_Use96Dpi())
	AA_X11INITTHREADS                         = ApplicationAttribute(C.Qt_AA_X11InitThreads())
	AA_SYNTHESIZETOUCHFORUNHANDLEDMOUSEEVENTS = ApplicationAttribute(C.Qt_AA_SynthesizeTouchForUnhandledMouseEvents())
	AA_SYNTHESIZEMOUSEFORUNHANDLEDTOUCHEVENTS = ApplicationAttribute(C.Qt_AA_SynthesizeMouseForUnhandledTouchEvents())
	AA_USEHIGHDPIPIXMAPS                      = ApplicationAttribute(C.Qt_AA_UseHighDpiPixmaps())
	AA_FORCERASTERWIDGETS                     = ApplicationAttribute(C.Qt_AA_ForceRasterWidgets())
	AA_USEDESKTOPOPENGL                       = ApplicationAttribute(C.Qt_AA_UseDesktopOpenGL())
	AA_USEOPENGLES                            = ApplicationAttribute(C.Qt_AA_UseOpenGLES())
	AA_IMMEDIATEWIDGETCREATION                = ApplicationAttribute(C.Qt_AA_ImmediateWidgetCreation())
	AA_MSWINDOWSUSEDIRECT3DBYDEFAULT          = ApplicationAttribute(C.Qt_AA_MSWindowsUseDirect3DByDefault())
)

type ApplicationState int

var (
	APPLICATIONSUSPENDED = ApplicationState(C.Qt_ApplicationSuspended())
	APPLICATIONHIDDEN    = ApplicationState(C.Qt_ApplicationHidden())
	APPLICATIONINACTIVE  = ApplicationState(C.Qt_ApplicationInactive())
	APPLICATIONACTIVE    = ApplicationState(C.Qt_ApplicationActive())
)

type ArrowType int

var (
	NOARROW    = ArrowType(C.Qt_NoArrow())
	UPARROW    = ArrowType(C.Qt_UpArrow())
	DOWNARROW  = ArrowType(C.Qt_DownArrow())
	LEFTARROW  = ArrowType(C.Qt_LeftArrow())
	RIGHTARROW = ArrowType(C.Qt_RightArrow())
)

type AspectRatioMode int

var (
	IGNOREASPECTRATIO          = AspectRatioMode(C.Qt_IgnoreAspectRatio())
	KEEPASPECTRATIO            = AspectRatioMode(C.Qt_KeepAspectRatio())
	KEEPASPECTRATIOBYEXPANDING = AspectRatioMode(C.Qt_KeepAspectRatioByExpanding())
)

type Axis int

var (
	XAXIS = Axis(C.Qt_XAxis())
	YAXIS = Axis(C.Qt_YAxis())
	ZAXIS = Axis(C.Qt_ZAxis())
)

type BGMode int

var (
	TRANSPARENTMODE = BGMode(C.Qt_TransparentMode())
	OPAQUEMODE      = BGMode(C.Qt_OpaqueMode())
)

type BrushStyle int

var (
	NOBRUSH                = BrushStyle(C.Qt_NoBrush())
	SOLIDPATTERN           = BrushStyle(C.Qt_SolidPattern())
	DENSE1PATTERN          = BrushStyle(C.Qt_Dense1Pattern())
	DENSE2PATTERN          = BrushStyle(C.Qt_Dense2Pattern())
	DENSE3PATTERN          = BrushStyle(C.Qt_Dense3Pattern())
	DENSE4PATTERN          = BrushStyle(C.Qt_Dense4Pattern())
	DENSE5PATTERN          = BrushStyle(C.Qt_Dense5Pattern())
	DENSE6PATTERN          = BrushStyle(C.Qt_Dense6Pattern())
	DENSE7PATTERN          = BrushStyle(C.Qt_Dense7Pattern())
	HORPATTERN             = BrushStyle(C.Qt_HorPattern())
	VERPATTERN             = BrushStyle(C.Qt_VerPattern())
	CROSSPATTERN           = BrushStyle(C.Qt_CrossPattern())
	BDIAGPATTERN           = BrushStyle(C.Qt_BDiagPattern())
	FDIAGPATTERN           = BrushStyle(C.Qt_FDiagPattern())
	DIAGCROSSPATTERN       = BrushStyle(C.Qt_DiagCrossPattern())
	LINEARGRADIENTPATTERN  = BrushStyle(C.Qt_LinearGradientPattern())
	CONICALGRADIENTPATTERN = BrushStyle(C.Qt_ConicalGradientPattern())
	RADIALGRADIENTPATTERN  = BrushStyle(C.Qt_RadialGradientPattern())
	TEXTUREPATTERN         = BrushStyle(C.Qt_TexturePattern())
)

type CaseSensitivity int

var (
	CASEINSENSITIVE = CaseSensitivity(C.Qt_CaseInsensitive())
	CASESENSITIVE   = CaseSensitivity(C.Qt_CaseSensitive())
)

type CheckState int

var (
	UNCHECKED        = CheckState(C.Qt_Unchecked())
	PARTIALLYCHECKED = CheckState(C.Qt_PartiallyChecked())
	CHECKED          = CheckState(C.Qt_Checked())
)

type ClipOperation int

var (
	NOCLIP        = ClipOperation(C.Qt_NoClip())
	REPLACECLIP   = ClipOperation(C.Qt_ReplaceClip())
	INTERSECTCLIP = ClipOperation(C.Qt_IntersectClip())
)

type ConnectionType int

var (
	AUTOCONNECTION           = ConnectionType(C.Qt_AutoConnection())
	DIRECTCONNECTION         = ConnectionType(C.Qt_DirectConnection())
	QUEUEDCONNECTION         = ConnectionType(C.Qt_QueuedConnection())
	BLOCKINGQUEUEDCONNECTION = ConnectionType(C.Qt_BlockingQueuedConnection())
	UNIQUECONNECTION         = ConnectionType(C.Qt_UniqueConnection())
)

type ContextMenuPolicy int

var (
	NOCONTEXTMENU      = ContextMenuPolicy(C.Qt_NoContextMenu())
	PREVENTCONTEXTMENU = ContextMenuPolicy(C.Qt_PreventContextMenu())
	DEFAULTCONTEXTMENU = ContextMenuPolicy(C.Qt_DefaultContextMenu())
	ACTIONSCONTEXTMENU = ContextMenuPolicy(C.Qt_ActionsContextMenu())
	CUSTOMCONTEXTMENU  = ContextMenuPolicy(C.Qt_CustomContextMenu())
)

type CoordinateSystem int

var (
	DEVICECOORDINATES  = CoordinateSystem(C.Qt_DeviceCoordinates())
	LOGICALCOORDINATES = CoordinateSystem(C.Qt_LogicalCoordinates())
)

type Corner int

var (
	TOPLEFTCORNER     = Corner(C.Qt_TopLeftCorner())
	TOPRIGHTCORNER    = Corner(C.Qt_TopRightCorner())
	BOTTOMLEFTCORNER  = Corner(C.Qt_BottomLeftCorner())
	BOTTOMRIGHTCORNER = Corner(C.Qt_BottomRightCorner())
)

type CursorMoveStyle int

var (
	LOGICALMOVESTYLE = CursorMoveStyle(C.Qt_LogicalMoveStyle())
	VISUALMOVESTYLE  = CursorMoveStyle(C.Qt_VisualMoveStyle())
)

type CursorShape int

var (
	ARROWCURSOR        = CursorShape(C.Qt_ArrowCursor())
	UPARROWCURSOR      = CursorShape(C.Qt_UpArrowCursor())
	CROSSCURSOR        = CursorShape(C.Qt_CrossCursor())
	WAITCURSOR         = CursorShape(C.Qt_WaitCursor())
	IBEAMCURSOR        = CursorShape(C.Qt_IBeamCursor())
	SIZEVERCURSOR      = CursorShape(C.Qt_SizeVerCursor())
	SIZEHORCURSOR      = CursorShape(C.Qt_SizeHorCursor())
	SIZEBDIAGCURSOR    = CursorShape(C.Qt_SizeBDiagCursor())
	SIZEFDIAGCURSOR    = CursorShape(C.Qt_SizeFDiagCursor())
	SIZEALLCURSOR      = CursorShape(C.Qt_SizeAllCursor())
	BLANKCURSOR        = CursorShape(C.Qt_BlankCursor())
	SPLITVCURSOR       = CursorShape(C.Qt_SplitVCursor())
	SPLITHCURSOR       = CursorShape(C.Qt_SplitHCursor())
	POINTINGHANDCURSOR = CursorShape(C.Qt_PointingHandCursor())
	FORBIDDENCURSOR    = CursorShape(C.Qt_ForbiddenCursor())
	OPENHANDCURSOR     = CursorShape(C.Qt_OpenHandCursor())
	CLOSEDHANDCURSOR   = CursorShape(C.Qt_ClosedHandCursor())
	WHATSTHISCURSOR    = CursorShape(C.Qt_WhatsThisCursor())
	BUSYCURSOR         = CursorShape(C.Qt_BusyCursor())
	DRAGMOVECURSOR     = CursorShape(C.Qt_DragMoveCursor())
	DRAGCOPYCURSOR     = CursorShape(C.Qt_DragCopyCursor())
	DRAGLINKCURSOR     = CursorShape(C.Qt_DragLinkCursor())
	BITMAPCURSOR       = CursorShape(C.Qt_BitmapCursor())
)

type DateFormat int

var (
	TEXTDATE               = DateFormat(C.Qt_TextDate())
	ISODATE                = DateFormat(C.Qt_ISODate())
	SYSTEMLOCALESHORTDATE  = DateFormat(C.Qt_SystemLocaleShortDate())
	SYSTEMLOCALELONGDATE   = DateFormat(C.Qt_SystemLocaleLongDate())
	DEFAULTLOCALESHORTDATE = DateFormat(C.Qt_DefaultLocaleShortDate())
	DEFAULTLOCALELONGDATE  = DateFormat(C.Qt_DefaultLocaleLongDate())
	SYSTEMLOCALEDATE       = DateFormat(C.Qt_SystemLocaleDate())
	LOCALEDATE             = DateFormat(C.Qt_LocaleDate())
	LOCALDATE              = DateFormat(C.Qt_LocalDate())
	RFC2822DATE            = DateFormat(C.Qt_RFC2822Date())
)

type DayOfWeek int

var (
	MONDAY    = DayOfWeek(C.Qt_Monday())
	TUESDAY   = DayOfWeek(C.Qt_Tuesday())
	WEDNESDAY = DayOfWeek(C.Qt_Wednesday())
	THURSDAY  = DayOfWeek(C.Qt_Thursday())
	FRIDAY    = DayOfWeek(C.Qt_Friday())
	SATURDAY  = DayOfWeek(C.Qt_Saturday())
	SUNDAY    = DayOfWeek(C.Qt_Sunday())
)

type DockWidgetArea int

var (
	LEFTDOCKWIDGETAREA   = DockWidgetArea(C.Qt_LeftDockWidgetArea())
	RIGHTDOCKWIDGETAREA  = DockWidgetArea(C.Qt_RightDockWidgetArea())
	TOPDOCKWIDGETAREA    = DockWidgetArea(C.Qt_TopDockWidgetArea())
	BOTTOMDOCKWIDGETAREA = DockWidgetArea(C.Qt_BottomDockWidgetArea())
	ALLDOCKWIDGETAREAS   = DockWidgetArea(C.Qt_AllDockWidgetAreas())
	NODOCKWIDGETAREA     = DockWidgetArea(C.Qt_NoDockWidgetArea())
)

type DropAction int

var (
	COPYACTION       = DropAction(C.Qt_CopyAction())
	MOVEACTION       = DropAction(C.Qt_MoveAction())
	LINKACTION       = DropAction(C.Qt_LinkAction())
	ACTIONMASK       = DropAction(C.Qt_ActionMask())
	IGNOREACTION     = DropAction(C.Qt_IgnoreAction())
	TARGETMOVEACTION = DropAction(C.Qt_TargetMoveAction())
)

type Edge int

var (
	TOPEDGE    = Edge(C.Qt_TopEdge())
	LEFTEDGE   = Edge(C.Qt_LeftEdge())
	RIGHTEDGE  = Edge(C.Qt_RightEdge())
	BOTTOMEDGE = Edge(C.Qt_BottomEdge())
)

type EventPriority int

var (
	HIGHEVENTPRIORITY   = EventPriority(C.Qt_HighEventPriority())
	NORMALEVENTPRIORITY = EventPriority(C.Qt_NormalEventPriority())
	LOWEVENTPRIORITY    = EventPriority(C.Qt_LowEventPriority())
)

type FillRule int

var (
	ODDEVENFILL = FillRule(C.Qt_OddEvenFill())
	WINDINGFILL = FillRule(C.Qt_WindingFill())
)

type FindChildOption int

var (
	FINDDIRECTCHILDRENONLY  = FindChildOption(C.Qt_FindDirectChildrenOnly())
	FINDCHILDRENRECURSIVELY = FindChildOption(C.Qt_FindChildrenRecursively())
)

type FocusPolicy int

var (
	TABFOCUS    = FocusPolicy(C.Qt_TabFocus())
	CLICKFOCUS  = FocusPolicy(C.Qt_ClickFocus())
	STRONGFOCUS = FocusPolicy(C.Qt_StrongFocus())
	WHEELFOCUS  = FocusPolicy(C.Qt_WheelFocus())
	NOFOCUS     = FocusPolicy(C.Qt_NoFocus())
)

type FocusReason int

var (
	MOUSEFOCUSREASON        = FocusReason(C.Qt_MouseFocusReason())
	TABFOCUSREASON          = FocusReason(C.Qt_TabFocusReason())
	BACKTABFOCUSREASON      = FocusReason(C.Qt_BacktabFocusReason())
	ACTIVEWINDOWFOCUSREASON = FocusReason(C.Qt_ActiveWindowFocusReason())
	POPUPFOCUSREASON        = FocusReason(C.Qt_PopupFocusReason())
	SHORTCUTFOCUSREASON     = FocusReason(C.Qt_ShortcutFocusReason())
	MENUBARFOCUSREASON      = FocusReason(C.Qt_MenuBarFocusReason())
	OTHERFOCUSREASON        = FocusReason(C.Qt_OtherFocusReason())
)

type GestureFlag int

var (
	DONTSTARTGESTUREONCHILDREN       = GestureFlag(C.Qt_DontStartGestureOnChildren())
	RECEIVEPARTIALGESTURES           = GestureFlag(C.Qt_ReceivePartialGestures())
	IGNOREDGESTURESPROPAGATETOPARENT = GestureFlag(C.Qt_IgnoredGesturesPropagateToParent())
)

type GestureState int

var (
	GESTURESTARTED  = GestureState(C.Qt_GestureStarted())
	GESTUREUPDATED  = GestureState(C.Qt_GestureUpdated())
	GESTUREFINISHED = GestureState(C.Qt_GestureFinished())
	GESTURECANCELED = GestureState(C.Qt_GestureCanceled())
)

type GestureType int

var (
	TAPGESTURE        = GestureType(C.Qt_TapGesture())
	TAPANDHOLDGESTURE = GestureType(C.Qt_TapAndHoldGesture())
	PANGESTURE        = GestureType(C.Qt_PanGesture())
	PINCHGESTURE      = GestureType(C.Qt_PinchGesture())
	SWIPEGESTURE      = GestureType(C.Qt_SwipeGesture())
	CUSTOMGESTURE     = GestureType(C.Qt_CustomGesture())
)

type GlobalColor int

var (
	WHITE       = GlobalColor(C.Qt_white())
	BLACK       = GlobalColor(C.Qt_black())
	RED         = GlobalColor(C.Qt_red())
	DARKRED     = GlobalColor(C.Qt_darkRed())
	GREEN       = GlobalColor(C.Qt_green())
	DARKGREEN   = GlobalColor(C.Qt_darkGreen())
	BLUE        = GlobalColor(C.Qt_blue())
	DARKBLUE    = GlobalColor(C.Qt_darkBlue())
	CYAN        = GlobalColor(C.Qt_cyan())
	DARKCYAN    = GlobalColor(C.Qt_darkCyan())
	MAGENTA     = GlobalColor(C.Qt_magenta())
	DARKMAGENTA = GlobalColor(C.Qt_darkMagenta())
	YELLOW      = GlobalColor(C.Qt_yellow())
	DARKYELLOW  = GlobalColor(C.Qt_darkYellow())
	GRAY        = GlobalColor(C.Qt_gray())
	DARKGRAY    = GlobalColor(C.Qt_darkGray())
	LIGHTGRAY   = GlobalColor(C.Qt_lightGray())
	TRANSPARENT = GlobalColor(C.Qt_transparent())
	COLOR0      = GlobalColor(C.Qt_color0())
	COLOR1      = GlobalColor(C.Qt_color1())
)

type HitTestAccuracy int

var (
	EXACTHIT = HitTestAccuracy(C.Qt_ExactHit())
	FUZZYHIT = HitTestAccuracy(C.Qt_FuzzyHit())
)

type ImageConversionFlag int

var (
	AUTOCOLOR            = ImageConversionFlag(C.Qt_AutoColor())
	COLORONLY            = ImageConversionFlag(C.Qt_ColorOnly())
	MONOONLY             = ImageConversionFlag(C.Qt_MonoOnly())
	DIFFUSEDITHER        = ImageConversionFlag(C.Qt_DiffuseDither())
	ORDEREDDITHER        = ImageConversionFlag(C.Qt_OrderedDither())
	THRESHOLDDITHER      = ImageConversionFlag(C.Qt_ThresholdDither())
	THRESHOLDALPHADITHER = ImageConversionFlag(C.Qt_ThresholdAlphaDither())
	ORDEREDALPHADITHER   = ImageConversionFlag(C.Qt_OrderedAlphaDither())
	DIFFUSEALPHADITHER   = ImageConversionFlag(C.Qt_DiffuseAlphaDither())
	PREFERDITHER         = ImageConversionFlag(C.Qt_PreferDither())
	AVOIDDITHER          = ImageConversionFlag(C.Qt_AvoidDither())
	NOOPAQUEDETECTION    = ImageConversionFlag(C.Qt_NoOpaqueDetection())
	NOFORMATCONVERSION   = ImageConversionFlag(C.Qt_NoFormatConversion())
)

type InputMethodHint int

var (
	IMHNONE                   = InputMethodHint(C.Qt_ImhNone())
	IMHHIDDENTEXT             = InputMethodHint(C.Qt_ImhHiddenText())
	IMHSENSITIVEDATA          = InputMethodHint(C.Qt_ImhSensitiveData())
	IMHNOAUTOUPPERCASE        = InputMethodHint(C.Qt_ImhNoAutoUppercase())
	IMHPREFERNUMBERS          = InputMethodHint(C.Qt_ImhPreferNumbers())
	IMHPREFERUPPERCASE        = InputMethodHint(C.Qt_ImhPreferUppercase())
	IMHPREFERLOWERCASE        = InputMethodHint(C.Qt_ImhPreferLowercase())
	IMHNOPREDICTIVETEXT       = InputMethodHint(C.Qt_ImhNoPredictiveText())
	IMHDATE                   = InputMethodHint(C.Qt_ImhDate())
	IMHTIME                   = InputMethodHint(C.Qt_ImhTime())
	IMHPREFERLATIN            = InputMethodHint(C.Qt_ImhPreferLatin())
	IMHMULTILINE              = InputMethodHint(C.Qt_ImhMultiLine())
	IMHDIGITSONLY             = InputMethodHint(C.Qt_ImhDigitsOnly())
	IMHFORMATTEDNUMBERSONLY   = InputMethodHint(C.Qt_ImhFormattedNumbersOnly())
	IMHUPPERCASEONLY          = InputMethodHint(C.Qt_ImhUppercaseOnly())
	IMHLOWERCASEONLY          = InputMethodHint(C.Qt_ImhLowercaseOnly())
	IMHDIALABLECHARACTERSONLY = InputMethodHint(C.Qt_ImhDialableCharactersOnly())
	IMHEMAILCHARACTERSONLY    = InputMethodHint(C.Qt_ImhEmailCharactersOnly())
	IMHURLCHARACTERSONLY      = InputMethodHint(C.Qt_ImhUrlCharactersOnly())
	IMHLATINONLY              = InputMethodHint(C.Qt_ImhLatinOnly())
	IMHEXCLUSIVEINPUTMASK     = InputMethodHint(C.Qt_ImhExclusiveInputMask())
)

type InputMethodQuery int

var (
	IMENABLED           = InputMethodQuery(C.Qt_ImEnabled())
	IMMICROFOCUS        = InputMethodQuery(C.Qt_ImMicroFocus())
	IMCURSORRECTANGLE   = InputMethodQuery(C.Qt_ImCursorRectangle())
	IMFONT              = InputMethodQuery(C.Qt_ImFont())
	IMCURSORPOSITION    = InputMethodQuery(C.Qt_ImCursorPosition())
	IMSURROUNDINGTEXT   = InputMethodQuery(C.Qt_ImSurroundingText())
	IMCURRENTSELECTION  = InputMethodQuery(C.Qt_ImCurrentSelection())
	IMMAXIMUMTEXTLENGTH = InputMethodQuery(C.Qt_ImMaximumTextLength())
	IMANCHORPOSITION    = InputMethodQuery(C.Qt_ImAnchorPosition())
	IMHINTS             = InputMethodQuery(C.Qt_ImHints())
	IMPREFERREDLANGUAGE = InputMethodQuery(C.Qt_ImPreferredLanguage())
	IMPLATFORMDATA      = InputMethodQuery(C.Qt_ImPlatformData())
	IMABSOLUTEPOSITION  = InputMethodQuery(C.Qt_ImAbsolutePosition())
	IMTEXTBEFORECURSOR  = InputMethodQuery(C.Qt_ImTextBeforeCursor())
	IMTEXTAFTERCURSOR   = InputMethodQuery(C.Qt_ImTextAfterCursor())
	IMQUERYINPUT        = InputMethodQuery(C.Qt_ImQueryInput())
	IMQUERYALL          = InputMethodQuery(C.Qt_ImQueryAll())
)

type ItemDataRole int

var (
	DISPLAYROLE               = ItemDataRole(C.Qt_DisplayRole())
	DECORATIONROLE            = ItemDataRole(C.Qt_DecorationRole())
	EDITROLE                  = ItemDataRole(C.Qt_EditRole())
	TOOLTIPROLE               = ItemDataRole(C.Qt_ToolTipRole())
	STATUSTIPROLE             = ItemDataRole(C.Qt_StatusTipRole())
	WHATSTHISROLE             = ItemDataRole(C.Qt_WhatsThisRole())
	SIZEHINTROLE              = ItemDataRole(C.Qt_SizeHintRole())
	FONTROLE                  = ItemDataRole(C.Qt_FontRole())
	TEXTALIGNMENTROLE         = ItemDataRole(C.Qt_TextAlignmentRole())
	BACKGROUNDROLE            = ItemDataRole(C.Qt_BackgroundRole())
	BACKGROUNDCOLORROLE       = ItemDataRole(C.Qt_BackgroundColorRole())
	FOREGROUNDROLE            = ItemDataRole(C.Qt_ForegroundRole())
	TEXTCOLORROLE             = ItemDataRole(C.Qt_TextColorRole())
	CHECKSTATEROLE            = ItemDataRole(C.Qt_CheckStateRole())
	INITIALSORTORDERROLE      = ItemDataRole(C.Qt_InitialSortOrderRole())
	ACCESSIBLETEXTROLE        = ItemDataRole(C.Qt_AccessibleTextRole())
	ACCESSIBLEDESCRIPTIONROLE = ItemDataRole(C.Qt_AccessibleDescriptionRole())
	USERROLE                  = ItemDataRole(C.Qt_UserRole())
)

type ItemFlag int

var (
	NOITEMFLAGS          = ItemFlag(C.Qt_NoItemFlags())
	ITEMISSELECTABLE     = ItemFlag(C.Qt_ItemIsSelectable())
	ITEMISEDITABLE       = ItemFlag(C.Qt_ItemIsEditable())
	ITEMISDRAGENABLED    = ItemFlag(C.Qt_ItemIsDragEnabled())
	ITEMISDROPENABLED    = ItemFlag(C.Qt_ItemIsDropEnabled())
	ITEMISUSERCHECKABLE  = ItemFlag(C.Qt_ItemIsUserCheckable())
	ITEMISENABLED        = ItemFlag(C.Qt_ItemIsEnabled())
	ITEMISTRISTATE       = ItemFlag(C.Qt_ItemIsTristate())
	ITEMNEVERHASCHILDREN = ItemFlag(C.Qt_ItemNeverHasChildren())
)

type ItemSelectionMode int

var (
	CONTAINSITEMSHAPE          = ItemSelectionMode(C.Qt_ContainsItemShape())
	INTERSECTSITEMSHAPE        = ItemSelectionMode(C.Qt_IntersectsItemShape())
	CONTAINSITEMBOUNDINGRECT   = ItemSelectionMode(C.Qt_ContainsItemBoundingRect())
	INTERSECTSITEMBOUNDINGRECT = ItemSelectionMode(C.Qt_IntersectsItemBoundingRect())
)

type Key int

var (
	KEY_ESCAPE                 = Key(C.Qt_Key_Escape())
	KEY_TAB                    = Key(C.Qt_Key_Tab())
	KEY_BACKTAB                = Key(C.Qt_Key_Backtab())
	KEY_BACKSPACE              = Key(C.Qt_Key_Backspace())
	KEY_RETURN                 = Key(C.Qt_Key_Return())
	KEY_ENTER                  = Key(C.Qt_Key_Enter())
	KEY_INSERT                 = Key(C.Qt_Key_Insert())
	KEY_DELETE                 = Key(C.Qt_Key_Delete())
	KEY_PAUSE                  = Key(C.Qt_Key_Pause())
	KEY_PRINT                  = Key(C.Qt_Key_Print())
	KEY_SYSREQ                 = Key(C.Qt_Key_SysReq())
	KEY_CLEAR                  = Key(C.Qt_Key_Clear())
	KEY_HOME                   = Key(C.Qt_Key_Home())
	KEY_END                    = Key(C.Qt_Key_End())
	KEY_LEFT                   = Key(C.Qt_Key_Left())
	KEY_UP                     = Key(C.Qt_Key_Up())
	KEY_RIGHT                  = Key(C.Qt_Key_Right())
	KEY_DOWN                   = Key(C.Qt_Key_Down())
	KEY_PAGEUP                 = Key(C.Qt_Key_PageUp())
	KEY_PAGEDOWN               = Key(C.Qt_Key_PageDown())
	KEY_SHIFT                  = Key(C.Qt_Key_Shift())
	KEY_CONTROL                = Key(C.Qt_Key_Control())
	KEY_META                   = Key(C.Qt_Key_Meta())
	KEY_ALT                    = Key(C.Qt_Key_Alt())
	KEY_ALTGR                  = Key(C.Qt_Key_AltGr())
	KEY_CAPSLOCK               = Key(C.Qt_Key_CapsLock())
	KEY_NUMLOCK                = Key(C.Qt_Key_NumLock())
	KEY_SCROLLLOCK             = Key(C.Qt_Key_ScrollLock())
	KEY_F1                     = Key(C.Qt_Key_F1())
	KEY_F2                     = Key(C.Qt_Key_F2())
	KEY_F3                     = Key(C.Qt_Key_F3())
	KEY_F4                     = Key(C.Qt_Key_F4())
	KEY_F5                     = Key(C.Qt_Key_F5())
	KEY_F6                     = Key(C.Qt_Key_F6())
	KEY_F7                     = Key(C.Qt_Key_F7())
	KEY_F8                     = Key(C.Qt_Key_F8())
	KEY_F9                     = Key(C.Qt_Key_F9())
	KEY_F10                    = Key(C.Qt_Key_F10())
	KEY_F11                    = Key(C.Qt_Key_F11())
	KEY_F12                    = Key(C.Qt_Key_F12())
	KEY_F13                    = Key(C.Qt_Key_F13())
	KEY_F14                    = Key(C.Qt_Key_F14())
	KEY_F15                    = Key(C.Qt_Key_F15())
	KEY_F16                    = Key(C.Qt_Key_F16())
	KEY_F17                    = Key(C.Qt_Key_F17())
	KEY_F18                    = Key(C.Qt_Key_F18())
	KEY_F19                    = Key(C.Qt_Key_F19())
	KEY_F20                    = Key(C.Qt_Key_F20())
	KEY_F21                    = Key(C.Qt_Key_F21())
	KEY_F22                    = Key(C.Qt_Key_F22())
	KEY_F23                    = Key(C.Qt_Key_F23())
	KEY_F24                    = Key(C.Qt_Key_F24())
	KEY_F25                    = Key(C.Qt_Key_F25())
	KEY_F26                    = Key(C.Qt_Key_F26())
	KEY_F27                    = Key(C.Qt_Key_F27())
	KEY_F28                    = Key(C.Qt_Key_F28())
	KEY_F29                    = Key(C.Qt_Key_F29())
	KEY_F30                    = Key(C.Qt_Key_F30())
	KEY_F31                    = Key(C.Qt_Key_F31())
	KEY_F32                    = Key(C.Qt_Key_F32())
	KEY_F33                    = Key(C.Qt_Key_F33())
	KEY_F34                    = Key(C.Qt_Key_F34())
	KEY_F35                    = Key(C.Qt_Key_F35())
	KEY_SUPER_L                = Key(C.Qt_Key_Super_L())
	KEY_SUPER_R                = Key(C.Qt_Key_Super_R())
	KEY_MENU                   = Key(C.Qt_Key_Menu())
	KEY_HYPER_L                = Key(C.Qt_Key_Hyper_L())
	KEY_HYPER_R                = Key(C.Qt_Key_Hyper_R())
	KEY_HELP                   = Key(C.Qt_Key_Help())
	KEY_DIRECTION_L            = Key(C.Qt_Key_Direction_L())
	KEY_DIRECTION_R            = Key(C.Qt_Key_Direction_R())
	KEY_SPACE                  = Key(C.Qt_Key_Space())
	KEY_ANY                    = Key(C.Qt_Key_Any())
	KEY_EXCLAM                 = Key(C.Qt_Key_Exclam())
	KEY_QUOTEDBL               = Key(C.Qt_Key_QuoteDbl())
	KEY_NUMBERSIGN             = Key(C.Qt_Key_NumberSign())
	KEY_DOLLAR                 = Key(C.Qt_Key_Dollar())
	KEY_PERCENT                = Key(C.Qt_Key_Percent())
	KEY_AMPERSAND              = Key(C.Qt_Key_Ampersand())
	KEY_APOSTROPHE             = Key(C.Qt_Key_Apostrophe())
	KEY_PARENLEFT              = Key(C.Qt_Key_ParenLeft())
	KEY_PARENRIGHT             = Key(C.Qt_Key_ParenRight())
	KEY_ASTERISK               = Key(C.Qt_Key_Asterisk())
	KEY_PLUS                   = Key(C.Qt_Key_Plus())
	KEY_COMMA                  = Key(C.Qt_Key_Comma())
	KEY_MINUS                  = Key(C.Qt_Key_Minus())
	KEY_PERIOD                 = Key(C.Qt_Key_Period())
	KEY_SLASH                  = Key(C.Qt_Key_Slash())
	KEY_0                      = Key(C.Qt_Key_0())
	KEY_1                      = Key(C.Qt_Key_1())
	KEY_2                      = Key(C.Qt_Key_2())
	KEY_3                      = Key(C.Qt_Key_3())
	KEY_4                      = Key(C.Qt_Key_4())
	KEY_5                      = Key(C.Qt_Key_5())
	KEY_6                      = Key(C.Qt_Key_6())
	KEY_7                      = Key(C.Qt_Key_7())
	KEY_8                      = Key(C.Qt_Key_8())
	KEY_9                      = Key(C.Qt_Key_9())
	KEY_COLON                  = Key(C.Qt_Key_Colon())
	KEY_SEMICOLON              = Key(C.Qt_Key_Semicolon())
	KEY_LESS                   = Key(C.Qt_Key_Less())
	KEY_EQUAL                  = Key(C.Qt_Key_Equal())
	KEY_GREATER                = Key(C.Qt_Key_Greater())
	KEY_QUESTION               = Key(C.Qt_Key_Question())
	KEY_AT                     = Key(C.Qt_Key_At())
	KEY_A                      = Key(C.Qt_Key_A())
	KEY_B                      = Key(C.Qt_Key_B())
	KEY_C                      = Key(C.Qt_Key_C())
	KEY_D                      = Key(C.Qt_Key_D())
	KEY_E                      = Key(C.Qt_Key_E())
	KEY_F                      = Key(C.Qt_Key_F())
	KEY_G                      = Key(C.Qt_Key_G())
	KEY_H                      = Key(C.Qt_Key_H())
	KEY_I                      = Key(C.Qt_Key_I())
	KEY_J                      = Key(C.Qt_Key_J())
	KEY_K                      = Key(C.Qt_Key_K())
	KEY_L                      = Key(C.Qt_Key_L())
	KEY_M                      = Key(C.Qt_Key_M())
	KEY_N                      = Key(C.Qt_Key_N())
	KEY_O                      = Key(C.Qt_Key_O())
	KEY_P                      = Key(C.Qt_Key_P())
	KEY_Q                      = Key(C.Qt_Key_Q())
	KEY_R                      = Key(C.Qt_Key_R())
	KEY_S                      = Key(C.Qt_Key_S())
	KEY_T                      = Key(C.Qt_Key_T())
	KEY_U                      = Key(C.Qt_Key_U())
	KEY_V                      = Key(C.Qt_Key_V())
	KEY_W                      = Key(C.Qt_Key_W())
	KEY_X                      = Key(C.Qt_Key_X())
	KEY_Y                      = Key(C.Qt_Key_Y())
	KEY_Z                      = Key(C.Qt_Key_Z())
	KEY_BRACKETLEFT            = Key(C.Qt_Key_BracketLeft())
	KEY_BACKSLASH              = Key(C.Qt_Key_Backslash())
	KEY_BRACKETRIGHT           = Key(C.Qt_Key_BracketRight())
	KEY_ASCIICIRCUM            = Key(C.Qt_Key_AsciiCircum())
	KEY_UNDERSCORE             = Key(C.Qt_Key_Underscore())
	KEY_QUOTELEFT              = Key(C.Qt_Key_QuoteLeft())
	KEY_BRACELEFT              = Key(C.Qt_Key_BraceLeft())
	KEY_BAR                    = Key(C.Qt_Key_Bar())
	KEY_BRACERIGHT             = Key(C.Qt_Key_BraceRight())
	KEY_ASCIITILDE             = Key(C.Qt_Key_AsciiTilde())
	KEY_NOBREAKSPACE           = Key(C.Qt_Key_nobreakspace())
	KEY_EXCLAMDOWN             = Key(C.Qt_Key_exclamdown())
	KEY_CENT                   = Key(C.Qt_Key_cent())
	KEY_STERLING               = Key(C.Qt_Key_sterling())
	KEY_CURRENCY               = Key(C.Qt_Key_currency())
	KEY_YEN                    = Key(C.Qt_Key_yen())
	KEY_BROKENBAR              = Key(C.Qt_Key_brokenbar())
	KEY_SECTION                = Key(C.Qt_Key_section())
	KEY_DIAERESIS              = Key(C.Qt_Key_diaeresis())
	KEY_COPYRIGHT              = Key(C.Qt_Key_copyright())
	KEY_ORDFEMININE            = Key(C.Qt_Key_ordfeminine())
	KEY_GUILLEMOTLEFT          = Key(C.Qt_Key_guillemotleft())
	KEY_NOTSIGN                = Key(C.Qt_Key_notsign())
	KEY_HYPHEN                 = Key(C.Qt_Key_hyphen())
	KEY_REGISTERED             = Key(C.Qt_Key_registered())
	KEY_MACRON                 = Key(C.Qt_Key_macron())
	KEY_DEGREE                 = Key(C.Qt_Key_degree())
	KEY_PLUSMINUS              = Key(C.Qt_Key_plusminus())
	KEY_TWOSUPERIOR            = Key(C.Qt_Key_twosuperior())
	KEY_THREESUPERIOR          = Key(C.Qt_Key_threesuperior())
	KEY_ACUTE                  = Key(C.Qt_Key_acute())
	KEY_MU                     = Key(C.Qt_Key_mu())
	KEY_PARAGRAPH              = Key(C.Qt_Key_paragraph())
	KEY_PERIODCENTERED         = Key(C.Qt_Key_periodcentered())
	KEY_CEDILLA                = Key(C.Qt_Key_cedilla())
	KEY_ONESUPERIOR            = Key(C.Qt_Key_onesuperior())
	KEY_MASCULINE              = Key(C.Qt_Key_masculine())
	KEY_GUILLEMOTRIGHT         = Key(C.Qt_Key_guillemotright())
	KEY_ONEQUARTER             = Key(C.Qt_Key_onequarter())
	KEY_ONEHALF                = Key(C.Qt_Key_onehalf())
	KEY_THREEQUARTERS          = Key(C.Qt_Key_threequarters())
	KEY_QUESTIONDOWN           = Key(C.Qt_Key_questiondown())
	KEY_AGRAVE                 = Key(C.Qt_Key_Agrave())
	KEY_AACUTE                 = Key(C.Qt_Key_Aacute())
	KEY_ACIRCUMFLEX            = Key(C.Qt_Key_Acircumflex())
	KEY_ATILDE                 = Key(C.Qt_Key_Atilde())
	KEY_ADIAERESIS             = Key(C.Qt_Key_Adiaeresis())
	KEY_ARING                  = Key(C.Qt_Key_Aring())
	KEY_AE                     = Key(C.Qt_Key_AE())
	KEY_CCEDILLA               = Key(C.Qt_Key_Ccedilla())
	KEY_EGRAVE                 = Key(C.Qt_Key_Egrave())
	KEY_EACUTE                 = Key(C.Qt_Key_Eacute())
	KEY_ECIRCUMFLEX            = Key(C.Qt_Key_Ecircumflex())
	KEY_EDIAERESIS             = Key(C.Qt_Key_Ediaeresis())
	KEY_IGRAVE                 = Key(C.Qt_Key_Igrave())
	KEY_IACUTE                 = Key(C.Qt_Key_Iacute())
	KEY_ICIRCUMFLEX            = Key(C.Qt_Key_Icircumflex())
	KEY_IDIAERESIS             = Key(C.Qt_Key_Idiaeresis())
	KEY_ETH                    = Key(C.Qt_Key_ETH())
	KEY_NTILDE                 = Key(C.Qt_Key_Ntilde())
	KEY_OGRAVE                 = Key(C.Qt_Key_Ograve())
	KEY_OACUTE                 = Key(C.Qt_Key_Oacute())
	KEY_OCIRCUMFLEX            = Key(C.Qt_Key_Ocircumflex())
	KEY_OTILDE                 = Key(C.Qt_Key_Otilde())
	KEY_ODIAERESIS             = Key(C.Qt_Key_Odiaeresis())
	KEY_MULTIPLY               = Key(C.Qt_Key_multiply())
	KEY_OOBLIQUE               = Key(C.Qt_Key_Ooblique())
	KEY_UGRAVE                 = Key(C.Qt_Key_Ugrave())
	KEY_UACUTE                 = Key(C.Qt_Key_Uacute())
	KEY_UCIRCUMFLEX            = Key(C.Qt_Key_Ucircumflex())
	KEY_UDIAERESIS             = Key(C.Qt_Key_Udiaeresis())
	KEY_YACUTE                 = Key(C.Qt_Key_Yacute())
	KEY_THORN                  = Key(C.Qt_Key_THORN())
	KEY_SSHARP                 = Key(C.Qt_Key_ssharp())
	KEY_DIVISION               = Key(C.Qt_Key_division())
	KEY_YDIAERESIS             = Key(C.Qt_Key_ydiaeresis())
	KEY_MULTI_KEY              = Key(C.Qt_Key_Multi_key())
	KEY_CODEINPUT              = Key(C.Qt_Key_Codeinput())
	KEY_SINGLECANDIDATE        = Key(C.Qt_Key_SingleCandidate())
	KEY_MULTIPLECANDIDATE      = Key(C.Qt_Key_MultipleCandidate())
	KEY_PREVIOUSCANDIDATE      = Key(C.Qt_Key_PreviousCandidate())
	KEY_MODE_SWITCH            = Key(C.Qt_Key_Mode_switch())
	KEY_KANJI                  = Key(C.Qt_Key_Kanji())
	KEY_MUHENKAN               = Key(C.Qt_Key_Muhenkan())
	KEY_HENKAN                 = Key(C.Qt_Key_Henkan())
	KEY_ROMAJI                 = Key(C.Qt_Key_Romaji())
	KEY_HIRAGANA               = Key(C.Qt_Key_Hiragana())
	KEY_KATAKANA               = Key(C.Qt_Key_Katakana())
	KEY_HIRAGANA_KATAKANA      = Key(C.Qt_Key_Hiragana_Katakana())
	KEY_ZENKAKU                = Key(C.Qt_Key_Zenkaku())
	KEY_HANKAKU                = Key(C.Qt_Key_Hankaku())
	KEY_ZENKAKU_HANKAKU        = Key(C.Qt_Key_Zenkaku_Hankaku())
	KEY_TOUROKU                = Key(C.Qt_Key_Touroku())
	KEY_MASSYO                 = Key(C.Qt_Key_Massyo())
	KEY_KANA_LOCK              = Key(C.Qt_Key_Kana_Lock())
	KEY_KANA_SHIFT             = Key(C.Qt_Key_Kana_Shift())
	KEY_EISU_SHIFT             = Key(C.Qt_Key_Eisu_Shift())
	KEY_EISU_TOGGLE            = Key(C.Qt_Key_Eisu_toggle())
	KEY_HANGUL                 = Key(C.Qt_Key_Hangul())
	KEY_HANGUL_START           = Key(C.Qt_Key_Hangul_Start())
	KEY_HANGUL_END             = Key(C.Qt_Key_Hangul_End())
	KEY_HANGUL_HANJA           = Key(C.Qt_Key_Hangul_Hanja())
	KEY_HANGUL_JAMO            = Key(C.Qt_Key_Hangul_Jamo())
	KEY_HANGUL_ROMAJA          = Key(C.Qt_Key_Hangul_Romaja())
	KEY_HANGUL_JEONJA          = Key(C.Qt_Key_Hangul_Jeonja())
	KEY_HANGUL_BANJA           = Key(C.Qt_Key_Hangul_Banja())
	KEY_HANGUL_PREHANJA        = Key(C.Qt_Key_Hangul_PreHanja())
	KEY_HANGUL_POSTHANJA       = Key(C.Qt_Key_Hangul_PostHanja())
	KEY_HANGUL_SPECIAL         = Key(C.Qt_Key_Hangul_Special())
	KEY_DEAD_GRAVE             = Key(C.Qt_Key_Dead_Grave())
	KEY_DEAD_ACUTE             = Key(C.Qt_Key_Dead_Acute())
	KEY_DEAD_CIRCUMFLEX        = Key(C.Qt_Key_Dead_Circumflex())
	KEY_DEAD_TILDE             = Key(C.Qt_Key_Dead_Tilde())
	KEY_DEAD_MACRON            = Key(C.Qt_Key_Dead_Macron())
	KEY_DEAD_BREVE             = Key(C.Qt_Key_Dead_Breve())
	KEY_DEAD_ABOVEDOT          = Key(C.Qt_Key_Dead_Abovedot())
	KEY_DEAD_DIAERESIS         = Key(C.Qt_Key_Dead_Diaeresis())
	KEY_DEAD_ABOVERING         = Key(C.Qt_Key_Dead_Abovering())
	KEY_DEAD_DOUBLEACUTE       = Key(C.Qt_Key_Dead_Doubleacute())
	KEY_DEAD_CARON             = Key(C.Qt_Key_Dead_Caron())
	KEY_DEAD_CEDILLA           = Key(C.Qt_Key_Dead_Cedilla())
	KEY_DEAD_OGONEK            = Key(C.Qt_Key_Dead_Ogonek())
	KEY_DEAD_IOTA              = Key(C.Qt_Key_Dead_Iota())
	KEY_DEAD_VOICED_SOUND      = Key(C.Qt_Key_Dead_Voiced_Sound())
	KEY_DEAD_SEMIVOICED_SOUND  = Key(C.Qt_Key_Dead_Semivoiced_Sound())
	KEY_DEAD_BELOWDOT          = Key(C.Qt_Key_Dead_Belowdot())
	KEY_DEAD_HOOK              = Key(C.Qt_Key_Dead_Hook())
	KEY_DEAD_HORN              = Key(C.Qt_Key_Dead_Horn())
	KEY_BACK                   = Key(C.Qt_Key_Back())
	KEY_FORWARD                = Key(C.Qt_Key_Forward())
	KEY_STOP                   = Key(C.Qt_Key_Stop())
	KEY_REFRESH                = Key(C.Qt_Key_Refresh())
	KEY_VOLUMEDOWN             = Key(C.Qt_Key_VolumeDown())
	KEY_VOLUMEMUTE             = Key(C.Qt_Key_VolumeMute())
	KEY_VOLUMEUP               = Key(C.Qt_Key_VolumeUp())
	KEY_BASSBOOST              = Key(C.Qt_Key_BassBoost())
	KEY_BASSUP                 = Key(C.Qt_Key_BassUp())
	KEY_BASSDOWN               = Key(C.Qt_Key_BassDown())
	KEY_TREBLEUP               = Key(C.Qt_Key_TrebleUp())
	KEY_TREBLEDOWN             = Key(C.Qt_Key_TrebleDown())
	KEY_MEDIAPLAY              = Key(C.Qt_Key_MediaPlay())
	KEY_MEDIASTOP              = Key(C.Qt_Key_MediaStop())
	KEY_MEDIAPREVIOUS          = Key(C.Qt_Key_MediaPrevious())
	KEY_MEDIANEXT              = Key(C.Qt_Key_MediaNext())
	KEY_MEDIARECORD            = Key(C.Qt_Key_MediaRecord())
	KEY_MEDIAPAUSE             = Key(C.Qt_Key_MediaPause())
	KEY_MEDIATOGGLEPLAYPAUSE   = Key(C.Qt_Key_MediaTogglePlayPause())
	KEY_HOMEPAGE               = Key(C.Qt_Key_HomePage())
	KEY_FAVORITES              = Key(C.Qt_Key_Favorites())
	KEY_SEARCH                 = Key(C.Qt_Key_Search())
	KEY_STANDBY                = Key(C.Qt_Key_Standby())
	KEY_OPENURL                = Key(C.Qt_Key_OpenUrl())
	KEY_LAUNCHMAIL             = Key(C.Qt_Key_LaunchMail())
	KEY_LAUNCHMEDIA            = Key(C.Qt_Key_LaunchMedia())
	KEY_LAUNCH0                = Key(C.Qt_Key_Launch0())
	KEY_LAUNCH1                = Key(C.Qt_Key_Launch1())
	KEY_LAUNCH2                = Key(C.Qt_Key_Launch2())
	KEY_LAUNCH3                = Key(C.Qt_Key_Launch3())
	KEY_LAUNCH4                = Key(C.Qt_Key_Launch4())
	KEY_LAUNCH5                = Key(C.Qt_Key_Launch5())
	KEY_LAUNCH6                = Key(C.Qt_Key_Launch6())
	KEY_LAUNCH7                = Key(C.Qt_Key_Launch7())
	KEY_LAUNCH8                = Key(C.Qt_Key_Launch8())
	KEY_LAUNCH9                = Key(C.Qt_Key_Launch9())
	KEY_LAUNCHA                = Key(C.Qt_Key_LaunchA())
	KEY_LAUNCHB                = Key(C.Qt_Key_LaunchB())
	KEY_LAUNCHC                = Key(C.Qt_Key_LaunchC())
	KEY_LAUNCHD                = Key(C.Qt_Key_LaunchD())
	KEY_LAUNCHE                = Key(C.Qt_Key_LaunchE())
	KEY_LAUNCHF                = Key(C.Qt_Key_LaunchF())
	KEY_LAUNCHG                = Key(C.Qt_Key_LaunchG())
	KEY_LAUNCHH                = Key(C.Qt_Key_LaunchH())
	KEY_MONBRIGHTNESSUP        = Key(C.Qt_Key_MonBrightnessUp())
	KEY_MONBRIGHTNESSDOWN      = Key(C.Qt_Key_MonBrightnessDown())
	KEY_KEYBOARDLIGHTONOFF     = Key(C.Qt_Key_KeyboardLightOnOff())
	KEY_KEYBOARDBRIGHTNESSUP   = Key(C.Qt_Key_KeyboardBrightnessUp())
	KEY_KEYBOARDBRIGHTNESSDOWN = Key(C.Qt_Key_KeyboardBrightnessDown())
	KEY_POWEROFF               = Key(C.Qt_Key_PowerOff())
	KEY_WAKEUP                 = Key(C.Qt_Key_WakeUp())
	KEY_EJECT                  = Key(C.Qt_Key_Eject())
	KEY_SCREENSAVER            = Key(C.Qt_Key_ScreenSaver())
	KEY_WWW                    = Key(C.Qt_Key_WWW())
	KEY_MEMO                   = Key(C.Qt_Key_Memo())
	KEY_LIGHTBULB              = Key(C.Qt_Key_LightBulb())
	KEY_SHOP                   = Key(C.Qt_Key_Shop())
	KEY_HISTORY                = Key(C.Qt_Key_History())
	KEY_ADDFAVORITE            = Key(C.Qt_Key_AddFavorite())
	KEY_HOTLINKS               = Key(C.Qt_Key_HotLinks())
	KEY_BRIGHTNESSADJUST       = Key(C.Qt_Key_BrightnessAdjust())
	KEY_FINANCE                = Key(C.Qt_Key_Finance())
	KEY_COMMUNITY              = Key(C.Qt_Key_Community())
	KEY_AUDIOREWIND            = Key(C.Qt_Key_AudioRewind())
	KEY_BACKFORWARD            = Key(C.Qt_Key_BackForward())
	KEY_APPLICATIONLEFT        = Key(C.Qt_Key_ApplicationLeft())
	KEY_APPLICATIONRIGHT       = Key(C.Qt_Key_ApplicationRight())
	KEY_BOOK                   = Key(C.Qt_Key_Book())
	KEY_CD                     = Key(C.Qt_Key_CD())
	KEY_CALCULATOR             = Key(C.Qt_Key_Calculator())
	KEY_TODOLIST               = Key(C.Qt_Key_ToDoList())
	KEY_CLEARGRAB              = Key(C.Qt_Key_ClearGrab())
	KEY_CLOSE                  = Key(C.Qt_Key_Close())
	KEY_COPY                   = Key(C.Qt_Key_Copy())
	KEY_CUT                    = Key(C.Qt_Key_Cut())
	KEY_DISPLAY                = Key(C.Qt_Key_Display())
	KEY_DOS                    = Key(C.Qt_Key_DOS())
	KEY_DOCUMENTS              = Key(C.Qt_Key_Documents())
	KEY_EXCEL                  = Key(C.Qt_Key_Excel())
	KEY_EXPLORER               = Key(C.Qt_Key_Explorer())
	KEY_GAME                   = Key(C.Qt_Key_Game())
	KEY_GO                     = Key(C.Qt_Key_Go())
	KEY_ITOUCH                 = Key(C.Qt_Key_iTouch())
	KEY_LOGOFF                 = Key(C.Qt_Key_LogOff())
	KEY_MARKET                 = Key(C.Qt_Key_Market())
	KEY_MEETING                = Key(C.Qt_Key_Meeting())
	KEY_MENUKB                 = Key(C.Qt_Key_MenuKB())
	KEY_MENUPB                 = Key(C.Qt_Key_MenuPB())
	KEY_MYSITES                = Key(C.Qt_Key_MySites())
	KEY_NEWS                   = Key(C.Qt_Key_News())
	KEY_OFFICEHOME             = Key(C.Qt_Key_OfficeHome())
	KEY_OPTION                 = Key(C.Qt_Key_Option())
	KEY_PASTE                  = Key(C.Qt_Key_Paste())
	KEY_PHONE                  = Key(C.Qt_Key_Phone())
	KEY_CALENDAR               = Key(C.Qt_Key_Calendar())
	KEY_REPLY                  = Key(C.Qt_Key_Reply())
	KEY_RELOAD                 = Key(C.Qt_Key_Reload())
	KEY_ROTATEWINDOWS          = Key(C.Qt_Key_RotateWindows())
	KEY_ROTATIONPB             = Key(C.Qt_Key_RotationPB())
	KEY_ROTATIONKB             = Key(C.Qt_Key_RotationKB())
	KEY_SAVE                   = Key(C.Qt_Key_Save())
	KEY_SEND                   = Key(C.Qt_Key_Send())
	KEY_SPELL                  = Key(C.Qt_Key_Spell())
	KEY_SPLITSCREEN            = Key(C.Qt_Key_SplitScreen())
	KEY_SUPPORT                = Key(C.Qt_Key_Support())
	KEY_TASKPANE               = Key(C.Qt_Key_TaskPane())
	KEY_TERMINAL               = Key(C.Qt_Key_Terminal())
	KEY_TOOLS                  = Key(C.Qt_Key_Tools())
	KEY_TRAVEL                 = Key(C.Qt_Key_Travel())
	KEY_VIDEO                  = Key(C.Qt_Key_Video())
	KEY_WORD                   = Key(C.Qt_Key_Word())
	KEY_XFER                   = Key(C.Qt_Key_Xfer())
	KEY_ZOOMIN                 = Key(C.Qt_Key_ZoomIn())
	KEY_ZOOMOUT                = Key(C.Qt_Key_ZoomOut())
	KEY_AWAY                   = Key(C.Qt_Key_Away())
	KEY_MESSENGER              = Key(C.Qt_Key_Messenger())
	KEY_WEBCAM                 = Key(C.Qt_Key_WebCam())
	KEY_MAILFORWARD            = Key(C.Qt_Key_MailForward())
	KEY_PICTURES               = Key(C.Qt_Key_Pictures())
	KEY_MUSIC                  = Key(C.Qt_Key_Music())
	KEY_BATTERY                = Key(C.Qt_Key_Battery())
	KEY_BLUETOOTH              = Key(C.Qt_Key_Bluetooth())
	KEY_WLAN                   = Key(C.Qt_Key_WLAN())
	KEY_UWB                    = Key(C.Qt_Key_UWB())
	KEY_AUDIOFORWARD           = Key(C.Qt_Key_AudioForward())
	KEY_AUDIOREPEAT            = Key(C.Qt_Key_AudioRepeat())
	KEY_AUDIORANDOMPLAY        = Key(C.Qt_Key_AudioRandomPlay())
	KEY_SUBTITLE               = Key(C.Qt_Key_Subtitle())
	KEY_AUDIOCYCLETRACK        = Key(C.Qt_Key_AudioCycleTrack())
	KEY_TIME                   = Key(C.Qt_Key_Time())
	KEY_HIBERNATE              = Key(C.Qt_Key_Hibernate())
	KEY_VIEW                   = Key(C.Qt_Key_View())
	KEY_TOPMENU                = Key(C.Qt_Key_TopMenu())
	KEY_POWERDOWN              = Key(C.Qt_Key_PowerDown())
	KEY_SUSPEND                = Key(C.Qt_Key_Suspend())
	KEY_CONTRASTADJUST         = Key(C.Qt_Key_ContrastAdjust())
	KEY_TOUCHPADTOGGLE         = Key(C.Qt_Key_TouchpadToggle())
	KEY_TOUCHPADON             = Key(C.Qt_Key_TouchpadOn())
	KEY_TOUCHPADOFF            = Key(C.Qt_Key_TouchpadOff())
	KEY_MICMUTE                = Key(C.Qt_Key_MicMute())
	KEY_RED                    = Key(C.Qt_Key_Red())
	KEY_GREEN                  = Key(C.Qt_Key_Green())
	KEY_YELLOW                 = Key(C.Qt_Key_Yellow())
	KEY_BLUE                   = Key(C.Qt_Key_Blue())
	KEY_CHANNELUP              = Key(C.Qt_Key_ChannelUp())
	KEY_CHANNELDOWN            = Key(C.Qt_Key_ChannelDown())
	KEY_MEDIALAST              = Key(C.Qt_Key_MediaLast())
	KEY_UNKNOWN                = Key(C.Qt_Key_unknown())
	KEY_CALL                   = Key(C.Qt_Key_Call())
	KEY_CAMERA                 = Key(C.Qt_Key_Camera())
	KEY_CAMERAFOCUS            = Key(C.Qt_Key_CameraFocus())
	KEY_CONTEXT1               = Key(C.Qt_Key_Context1())
	KEY_CONTEXT2               = Key(C.Qt_Key_Context2())
	KEY_CONTEXT3               = Key(C.Qt_Key_Context3())
	KEY_CONTEXT4               = Key(C.Qt_Key_Context4())
	KEY_FLIP                   = Key(C.Qt_Key_Flip())
	KEY_HANGUP                 = Key(C.Qt_Key_Hangup())
	KEY_NO                     = Key(C.Qt_Key_No())
	KEY_SELECT                 = Key(C.Qt_Key_Select())
	KEY_YES                    = Key(C.Qt_Key_Yes())
	KEY_TOGGLECALLHANGUP       = Key(C.Qt_Key_ToggleCallHangup())
	KEY_VOICEDIAL              = Key(C.Qt_Key_VoiceDial())
	KEY_LASTNUMBERREDIAL       = Key(C.Qt_Key_LastNumberRedial())
	KEY_EXECUTE                = Key(C.Qt_Key_Execute())
	KEY_PRINTER                = Key(C.Qt_Key_Printer())
	KEY_PLAY                   = Key(C.Qt_Key_Play())
	KEY_SLEEP                  = Key(C.Qt_Key_Sleep())
	KEY_ZOOM                   = Key(C.Qt_Key_Zoom())
	KEY_CANCEL                 = Key(C.Qt_Key_Cancel())
)

type KeyboardModifier int

var (
	NOMODIFIER          = KeyboardModifier(C.Qt_NoModifier())
	SHIFTMODIFIER       = KeyboardModifier(C.Qt_ShiftModifier())
	CONTROLMODIFIER     = KeyboardModifier(C.Qt_ControlModifier())
	ALTMODIFIER         = KeyboardModifier(C.Qt_AltModifier())
	METAMODIFIER        = KeyboardModifier(C.Qt_MetaModifier())
	KEYPADMODIFIER      = KeyboardModifier(C.Qt_KeypadModifier())
	GROUPSWITCHMODIFIER = KeyboardModifier(C.Qt_GroupSwitchModifier())
)

type LayoutDirection int

var (
	LEFTTORIGHT         = LayoutDirection(C.Qt_LeftToRight())
	RIGHTTOLEFT         = LayoutDirection(C.Qt_RightToLeft())
	LAYOUTDIRECTIONAUTO = LayoutDirection(C.Qt_LayoutDirectionAuto())
)

type MaskMode int

var (
	MASKINCOLOR  = MaskMode(C.Qt_MaskInColor())
	MASKOUTCOLOR = MaskMode(C.Qt_MaskOutColor())
)

type MatchFlag int

var (
	MATCHEXACTLY       = MatchFlag(C.Qt_MatchExactly())
	MATCHFIXEDSTRING   = MatchFlag(C.Qt_MatchFixedString())
	MATCHCONTAINS      = MatchFlag(C.Qt_MatchContains())
	MATCHSTARTSWITH    = MatchFlag(C.Qt_MatchStartsWith())
	MATCHENDSWITH      = MatchFlag(C.Qt_MatchEndsWith())
	MATCHCASESENSITIVE = MatchFlag(C.Qt_MatchCaseSensitive())
	MATCHREGEXP        = MatchFlag(C.Qt_MatchRegExp())
	MATCHWILDCARD      = MatchFlag(C.Qt_MatchWildcard())
	MATCHWRAP          = MatchFlag(C.Qt_MatchWrap())
	MATCHRECURSIVE     = MatchFlag(C.Qt_MatchRecursive())
)

type Modifier int

var (
	SHIFT         = Modifier(C.Qt_SHIFT())
	META          = Modifier(C.Qt_META())
	CTRL          = Modifier(C.Qt_CTRL())
	ALT           = Modifier(C.Qt_ALT())
	UNICODE_ACCEL = Modifier(C.Qt_UNICODE_ACCEL())
)

type MouseButton int

var (
	NOBUTTON      = MouseButton(C.Qt_NoButton())
	ALLBUTTONS    = MouseButton(C.Qt_AllButtons())
	LEFTBUTTON    = MouseButton(C.Qt_LeftButton())
	RIGHTBUTTON   = MouseButton(C.Qt_RightButton())
	MIDBUTTON     = MouseButton(C.Qt_MidButton())
	MIDDLEBUTTON  = MouseButton(C.Qt_MiddleButton())
	BACKBUTTON    = MouseButton(C.Qt_BackButton())
	XBUTTON1      = MouseButton(C.Qt_XButton1())
	EXTRABUTTON1  = MouseButton(C.Qt_ExtraButton1())
	FORWARDBUTTON = MouseButton(C.Qt_ForwardButton())
	XBUTTON2      = MouseButton(C.Qt_XButton2())
	EXTRABUTTON2  = MouseButton(C.Qt_ExtraButton2())
	TASKBUTTON    = MouseButton(C.Qt_TaskButton())
	EXTRABUTTON3  = MouseButton(C.Qt_ExtraButton3())
	EXTRABUTTON4  = MouseButton(C.Qt_ExtraButton4())
	EXTRABUTTON5  = MouseButton(C.Qt_ExtraButton5())
	EXTRABUTTON6  = MouseButton(C.Qt_ExtraButton6())
	EXTRABUTTON7  = MouseButton(C.Qt_ExtraButton7())
	EXTRABUTTON8  = MouseButton(C.Qt_ExtraButton8())
	EXTRABUTTON9  = MouseButton(C.Qt_ExtraButton9())
	EXTRABUTTON10 = MouseButton(C.Qt_ExtraButton10())
	EXTRABUTTON11 = MouseButton(C.Qt_ExtraButton11())
	EXTRABUTTON12 = MouseButton(C.Qt_ExtraButton12())
	EXTRABUTTON13 = MouseButton(C.Qt_ExtraButton13())
	EXTRABUTTON14 = MouseButton(C.Qt_ExtraButton14())
	EXTRABUTTON15 = MouseButton(C.Qt_ExtraButton15())
	EXTRABUTTON16 = MouseButton(C.Qt_ExtraButton16())
	EXTRABUTTON17 = MouseButton(C.Qt_ExtraButton17())
	EXTRABUTTON18 = MouseButton(C.Qt_ExtraButton18())
	EXTRABUTTON19 = MouseButton(C.Qt_ExtraButton19())
	EXTRABUTTON20 = MouseButton(C.Qt_ExtraButton20())
	EXTRABUTTON21 = MouseButton(C.Qt_ExtraButton21())
	EXTRABUTTON22 = MouseButton(C.Qt_ExtraButton22())
	EXTRABUTTON23 = MouseButton(C.Qt_ExtraButton23())
	EXTRABUTTON24 = MouseButton(C.Qt_ExtraButton24())
)

type MouseEventFlag int

var (
	MOUSEEVENTCREATEDDOUBLECLICK = MouseEventFlag(C.Qt_MouseEventCreatedDoubleClick())
)

type MouseEventSource int

var (
	MOUSEEVENTNOTSYNTHESIZED      = MouseEventSource(C.Qt_MouseEventNotSynthesized())
	MOUSEEVENTSYNTHESIZEDBYSYSTEM = MouseEventSource(C.Qt_MouseEventSynthesizedBySystem())
	MOUSEEVENTSYNTHESIZEDBYQT     = MouseEventSource(C.Qt_MouseEventSynthesizedByQt())
)

type NativeGestureType int

var ()

type NavigationMode int

var (
	NAVIGATIONMODENONE               = NavigationMode(C.Qt_NavigationModeNone())
	NAVIGATIONMODEKEYPADTABORDER     = NavigationMode(C.Qt_NavigationModeKeypadTabOrder())
	NAVIGATIONMODEKEYPADDIRECTIONAL  = NavigationMode(C.Qt_NavigationModeKeypadDirectional())
	NAVIGATIONMODECURSORAUTO         = NavigationMode(C.Qt_NavigationModeCursorAuto())
	NAVIGATIONMODECURSORFORCEVISIBLE = NavigationMode(C.Qt_NavigationModeCursorForceVisible())
)

type Orientation int

var (
	HORIZONTAL = Orientation(C.Qt_Horizontal())
	VERTICAL   = Orientation(C.Qt_Vertical())
)

type PenCapStyle int

var (
	SQUARECAP = PenCapStyle(C.Qt_SquareCap())
	FLATCAP   = PenCapStyle(C.Qt_FlatCap())
	ROUNDCAP  = PenCapStyle(C.Qt_RoundCap())
)

type PenJoinStyle int

var (
	BEVELJOIN    = PenJoinStyle(C.Qt_BevelJoin())
	MITERJOIN    = PenJoinStyle(C.Qt_MiterJoin())
	ROUNDJOIN    = PenJoinStyle(C.Qt_RoundJoin())
	SVGMITERJOIN = PenJoinStyle(C.Qt_SvgMiterJoin())
)

type PenStyle int

var (
	SOLIDLINE      = PenStyle(C.Qt_SolidLine())
	DASHDOTLINE    = PenStyle(C.Qt_DashDotLine())
	NOPEN          = PenStyle(C.Qt_NoPen())
	DASHLINE       = PenStyle(C.Qt_DashLine())
	DOTLINE        = PenStyle(C.Qt_DotLine())
	DASHDOTDOTLINE = PenStyle(C.Qt_DashDotDotLine())
	CUSTOMDASHLINE = PenStyle(C.Qt_CustomDashLine())
)

type ScreenOrientation int

var (
	PRIMARYORIENTATION           = ScreenOrientation(C.Qt_PrimaryOrientation())
	LANDSCAPEORIENTATION         = ScreenOrientation(C.Qt_LandscapeOrientation())
	PORTRAITORIENTATION          = ScreenOrientation(C.Qt_PortraitOrientation())
	INVERTEDLANDSCAPEORIENTATION = ScreenOrientation(C.Qt_InvertedLandscapeOrientation())
	INVERTEDPORTRAITORIENTATION  = ScreenOrientation(C.Qt_InvertedPortraitOrientation())
)

type ScrollBarPolicy int

var (
	SCROLLBARASNEEDED  = ScrollBarPolicy(C.Qt_ScrollBarAsNeeded())
	SCROLLBARALWAYSOFF = ScrollBarPolicy(C.Qt_ScrollBarAlwaysOff())
	SCROLLBARALWAYSON  = ScrollBarPolicy(C.Qt_ScrollBarAlwaysOn())
)

type ScrollPhase int

var (
	SCROLLBEGIN  = ScrollPhase(C.Qt_ScrollBegin())
	SCROLLUPDATE = ScrollPhase(C.Qt_ScrollUpdate())
	SCROLLEND    = ScrollPhase(C.Qt_ScrollEnd())
)

type ShortcutContext int

var (
	WIDGETSHORTCUT             = ShortcutContext(C.Qt_WidgetShortcut())
	WIDGETWITHCHILDRENSHORTCUT = ShortcutContext(C.Qt_WidgetWithChildrenShortcut())
	WINDOWSHORTCUT             = ShortcutContext(C.Qt_WindowShortcut())
	APPLICATIONSHORTCUT        = ShortcutContext(C.Qt_ApplicationShortcut())
)

type SizeHint int

var (
	MINIMUMSIZE    = SizeHint(C.Qt_MinimumSize())
	PREFERREDSIZE  = SizeHint(C.Qt_PreferredSize())
	MAXIMUMSIZE    = SizeHint(C.Qt_MaximumSize())
	MINIMUMDESCENT = SizeHint(C.Qt_MinimumDescent())
)

type SizeMode int

var (
	ABSOLUTESIZE = SizeMode(C.Qt_AbsoluteSize())
	RELATIVESIZE = SizeMode(C.Qt_RelativeSize())
)

type SortOrder int

var (
	ASCENDINGORDER  = SortOrder(C.Qt_AscendingOrder())
	DESCENDINGORDER = SortOrder(C.Qt_DescendingOrder())
)

type TextElideMode int

var (
	ELIDELEFT   = TextElideMode(C.Qt_ElideLeft())
	ELIDERIGHT  = TextElideMode(C.Qt_ElideRight())
	ELIDEMIDDLE = TextElideMode(C.Qt_ElideMiddle())
	ELIDENONE   = TextElideMode(C.Qt_ElideNone())
)

type TextFlag int

var (
	TEXTSINGLELINE            = TextFlag(C.Qt_TextSingleLine())
	TEXTDONTCLIP              = TextFlag(C.Qt_TextDontClip())
	TEXTEXPANDTABS            = TextFlag(C.Qt_TextExpandTabs())
	TEXTSHOWMNEMONIC          = TextFlag(C.Qt_TextShowMnemonic())
	TEXTWORDWRAP              = TextFlag(C.Qt_TextWordWrap())
	TEXTWRAPANYWHERE          = TextFlag(C.Qt_TextWrapAnywhere())
	TEXTHIDEMNEMONIC          = TextFlag(C.Qt_TextHideMnemonic())
	TEXTDONTPRINT             = TextFlag(C.Qt_TextDontPrint())
	TEXTINCLUDETRAILINGSPACES = TextFlag(C.Qt_TextIncludeTrailingSpaces())
	TEXTJUSTIFICATIONFORCED   = TextFlag(C.Qt_TextJustificationForced())
)

type TextFormat int

var (
	PLAINTEXT = TextFormat(C.Qt_PlainText())
	RICHTEXT  = TextFormat(C.Qt_RichText())
	AUTOTEXT  = TextFormat(C.Qt_AutoText())
)

type TextInteractionFlag int

var (
	NOTEXTINTERACTION         = TextInteractionFlag(C.Qt_NoTextInteraction())
	TEXTSELECTABLEBYMOUSE     = TextInteractionFlag(C.Qt_TextSelectableByMouse())
	TEXTSELECTABLEBYKEYBOARD  = TextInteractionFlag(C.Qt_TextSelectableByKeyboard())
	LINKSACCESSIBLEBYMOUSE    = TextInteractionFlag(C.Qt_LinksAccessibleByMouse())
	LINKSACCESSIBLEBYKEYBOARD = TextInteractionFlag(C.Qt_LinksAccessibleByKeyboard())
	TEXTEDITABLE              = TextInteractionFlag(C.Qt_TextEditable())
	TEXTEDITORINTERACTION     = TextInteractionFlag(C.Qt_TextEditorInteraction())
	TEXTBROWSERINTERACTION    = TextInteractionFlag(C.Qt_TextBrowserInteraction())
)

type TileRule int

var (
	STRETCHTILE = TileRule(C.Qt_StretchTile())
	REPEATTILE  = TileRule(C.Qt_RepeatTile())
	ROUNDTILE   = TileRule(C.Qt_RoundTile())
)

type TimeSpec int

var (
	LOCALTIME     = TimeSpec(C.Qt_LocalTime())
	UTC           = TimeSpec(C.Qt_UTC())
	OFFSETFROMUTC = TimeSpec(C.Qt_OffsetFromUTC())
	TIMEZONE      = TimeSpec(C.Qt_TimeZone())
)

type TimerType int

var (
	PRECISETIMER    = TimerType(C.Qt_PreciseTimer())
	COARSETIMER     = TimerType(C.Qt_CoarseTimer())
	VERYCOARSETIMER = TimerType(C.Qt_VeryCoarseTimer())
)

type ToolBarArea int

var (
	LEFTTOOLBARAREA   = ToolBarArea(C.Qt_LeftToolBarArea())
	RIGHTTOOLBARAREA  = ToolBarArea(C.Qt_RightToolBarArea())
	TOPTOOLBARAREA    = ToolBarArea(C.Qt_TopToolBarArea())
	BOTTOMTOOLBARAREA = ToolBarArea(C.Qt_BottomToolBarArea())
	ALLTOOLBARAREAS   = ToolBarArea(C.Qt_AllToolBarAreas())
	NOTOOLBARAREA     = ToolBarArea(C.Qt_NoToolBarArea())
)

type ToolButtonStyle int

var (
	TOOLBUTTONICONONLY       = ToolButtonStyle(C.Qt_ToolButtonIconOnly())
	TOOLBUTTONTEXTONLY       = ToolButtonStyle(C.Qt_ToolButtonTextOnly())
	TOOLBUTTONTEXTBESIDEICON = ToolButtonStyle(C.Qt_ToolButtonTextBesideIcon())
	TOOLBUTTONTEXTUNDERICON  = ToolButtonStyle(C.Qt_ToolButtonTextUnderIcon())
	TOOLBUTTONFOLLOWSTYLE    = ToolButtonStyle(C.Qt_ToolButtonFollowStyle())
)

type TouchPointState int

var (
	TOUCHPOINTPRESSED    = TouchPointState(C.Qt_TouchPointPressed())
	TOUCHPOINTMOVED      = TouchPointState(C.Qt_TouchPointMoved())
	TOUCHPOINTSTATIONARY = TouchPointState(C.Qt_TouchPointStationary())
	TOUCHPOINTRELEASED   = TouchPointState(C.Qt_TouchPointReleased())
)

type TransformationMode int

var (
	FASTTRANSFORMATION   = TransformationMode(C.Qt_FastTransformation())
	SMOOTHTRANSFORMATION = TransformationMode(C.Qt_SmoothTransformation())
)

type UIEffect int

var (
	UI_ANIMATEMENU    = UIEffect(C.Qt_UI_AnimateMenu())
	UI_FADEMENU       = UIEffect(C.Qt_UI_FadeMenu())
	UI_ANIMATECOMBO   = UIEffect(C.Qt_UI_AnimateCombo())
	UI_ANIMATETOOLTIP = UIEffect(C.Qt_UI_AnimateTooltip())
	UI_FADETOOLTIP    = UIEffect(C.Qt_UI_FadeTooltip())
	UI_ANIMATETOOLBOX = UIEffect(C.Qt_UI_AnimateToolBox())
)

type WhiteSpaceMode int

var (
	WHITESPACENORMAL = WhiteSpaceMode(C.Qt_WhiteSpaceNormal())
	WHITESPACEPRE    = WhiteSpaceMode(C.Qt_WhiteSpacePre())
	WHITESPACENOWRAP = WhiteSpaceMode(C.Qt_WhiteSpaceNoWrap())
)

type WidgetAttribute int

var (
	WA_ACCEPTDROPS                     = WidgetAttribute(C.Qt_WA_AcceptDrops())
	WA_ALWAYSSHOWTOOLTIPS              = WidgetAttribute(C.Qt_WA_AlwaysShowToolTips())
	WA_CONTENTSPROPAGATED              = WidgetAttribute(C.Qt_WA_ContentsPropagated())
	WA_CUSTOMWHATSTHIS                 = WidgetAttribute(C.Qt_WA_CustomWhatsThis())
	WA_DELETEONCLOSE                   = WidgetAttribute(C.Qt_WA_DeleteOnClose())
	WA_DISABLED                        = WidgetAttribute(C.Qt_WA_Disabled())
	WA_DONTSHOWONSCREEN                = WidgetAttribute(C.Qt_WA_DontShowOnScreen())
	WA_FORCEDISABLED                   = WidgetAttribute(C.Qt_WA_ForceDisabled())
	WA_FORCEUPDATESDISABLED            = WidgetAttribute(C.Qt_WA_ForceUpdatesDisabled())
	WA_GROUPLEADER                     = WidgetAttribute(C.Qt_WA_GroupLeader())
	WA_HOVER                           = WidgetAttribute(C.Qt_WA_Hover())
	WA_INPUTMETHODENABLED              = WidgetAttribute(C.Qt_WA_InputMethodEnabled())
	WA_KEYBOARDFOCUSCHANGE             = WidgetAttribute(C.Qt_WA_KeyboardFocusChange())
	WA_KEYCOMPRESSION                  = WidgetAttribute(C.Qt_WA_KeyCompression())
	WA_LAYOUTONENTIRERECT              = WidgetAttribute(C.Qt_WA_LayoutOnEntireRect())
	WA_LAYOUTUSESWIDGETRECT            = WidgetAttribute(C.Qt_WA_LayoutUsesWidgetRect())
	WA_MACNOCLICKTHROUGH               = WidgetAttribute(C.Qt_WA_MacNoClickThrough())
	WA_MACOPAQUESIZEGRIP               = WidgetAttribute(C.Qt_WA_MacOpaqueSizeGrip())
	WA_MACSHOWFOCUSRECT                = WidgetAttribute(C.Qt_WA_MacShowFocusRect())
	WA_MACNORMALSIZE                   = WidgetAttribute(C.Qt_WA_MacNormalSize())
	WA_MACSMALLSIZE                    = WidgetAttribute(C.Qt_WA_MacSmallSize())
	WA_MACMINISIZE                     = WidgetAttribute(C.Qt_WA_MacMiniSize())
	WA_MACVARIABLESIZE                 = WidgetAttribute(C.Qt_WA_MacVariableSize())
	WA_MACBRUSHEDMETAL                 = WidgetAttribute(C.Qt_WA_MacBrushedMetal())
	WA_MAPPED                          = WidgetAttribute(C.Qt_WA_Mapped())
	WA_MOUSENOMASK                     = WidgetAttribute(C.Qt_WA_MouseNoMask())
	WA_MOUSETRACKING                   = WidgetAttribute(C.Qt_WA_MouseTracking())
	WA_MOVED                           = WidgetAttribute(C.Qt_WA_Moved())
	WA_MSWINDOWSUSEDIRECT3D            = WidgetAttribute(C.Qt_WA_MSWindowsUseDirect3D())
	WA_NOBACKGROUND                    = WidgetAttribute(C.Qt_WA_NoBackground())
	WA_NOCHILDEVENTSFORPARENT          = WidgetAttribute(C.Qt_WA_NoChildEventsForParent())
	WA_NOCHILDEVENTSFROMCHILDREN       = WidgetAttribute(C.Qt_WA_NoChildEventsFromChildren())
	WA_NOMOUSEREPLAY                   = WidgetAttribute(C.Qt_WA_NoMouseReplay())
	WA_NOMOUSEPROPAGATION              = WidgetAttribute(C.Qt_WA_NoMousePropagation())
	WA_TRANSPARENTFORMOUSEEVENTS       = WidgetAttribute(C.Qt_WA_TransparentForMouseEvents())
	WA_NOSYSTEMBACKGROUND              = WidgetAttribute(C.Qt_WA_NoSystemBackground())
	WA_OPAQUEPAINTEVENT                = WidgetAttribute(C.Qt_WA_OpaquePaintEvent())
	WA_OUTSIDEWSRANGE                  = WidgetAttribute(C.Qt_WA_OutsideWSRange())
	WA_PAINTONSCREEN                   = WidgetAttribute(C.Qt_WA_PaintOnScreen())
	WA_PAINTUNCLIPPED                  = WidgetAttribute(C.Qt_WA_PaintUnclipped())
	WA_PENDINGMOVEEVENT                = WidgetAttribute(C.Qt_WA_PendingMoveEvent())
	WA_PENDINGRESIZEEVENT              = WidgetAttribute(C.Qt_WA_PendingResizeEvent())
	WA_QUITONCLOSE                     = WidgetAttribute(C.Qt_WA_QuitOnClose())
	WA_RESIZED                         = WidgetAttribute(C.Qt_WA_Resized())
	WA_RIGHTTOLEFT                     = WidgetAttribute(C.Qt_WA_RightToLeft())
	WA_SETCURSOR                       = WidgetAttribute(C.Qt_WA_SetCursor())
	WA_SETFONT                         = WidgetAttribute(C.Qt_WA_SetFont())
	WA_SETPALETTE                      = WidgetAttribute(C.Qt_WA_SetPalette())
	WA_SETSTYLE                        = WidgetAttribute(C.Qt_WA_SetStyle())
	WA_SHOWMODAL                       = WidgetAttribute(C.Qt_WA_ShowModal())
	WA_STATICCONTENTS                  = WidgetAttribute(C.Qt_WA_StaticContents())
	WA_STYLESHEET                      = WidgetAttribute(C.Qt_WA_StyleSheet())
	WA_TRANSLUCENTBACKGROUND           = WidgetAttribute(C.Qt_WA_TranslucentBackground())
	WA_UNDERMOUSE                      = WidgetAttribute(C.Qt_WA_UnderMouse())
	WA_UPDATESDISABLED                 = WidgetAttribute(C.Qt_WA_UpdatesDisabled())
	WA_WINDOWMODIFIED                  = WidgetAttribute(C.Qt_WA_WindowModified())
	WA_WINDOWPROPAGATION               = WidgetAttribute(C.Qt_WA_WindowPropagation())
	WA_MACALWAYSSHOWTOOLWINDOW         = WidgetAttribute(C.Qt_WA_MacAlwaysShowToolWindow())
	WA_SETLOCALE                       = WidgetAttribute(C.Qt_WA_SetLocale())
	WA_STYLEDBACKGROUND                = WidgetAttribute(C.Qt_WA_StyledBackground())
	WA_SHOWWITHOUTACTIVATING           = WidgetAttribute(C.Qt_WA_ShowWithoutActivating())
	WA_NATIVEWINDOW                    = WidgetAttribute(C.Qt_WA_NativeWindow())
	WA_DONTCREATENATIVEANCESTORS       = WidgetAttribute(C.Qt_WA_DontCreateNativeAncestors())
	WA_X11NETWMWINDOWTYPEDESKTOP       = WidgetAttribute(C.Qt_WA_X11NetWmWindowTypeDesktop())
	WA_X11NETWMWINDOWTYPEDOCK          = WidgetAttribute(C.Qt_WA_X11NetWmWindowTypeDock())
	WA_X11NETWMWINDOWTYPETOOLBAR       = WidgetAttribute(C.Qt_WA_X11NetWmWindowTypeToolBar())
	WA_X11NETWMWINDOWTYPEMENU          = WidgetAttribute(C.Qt_WA_X11NetWmWindowTypeMenu())
	WA_X11NETWMWINDOWTYPEUTILITY       = WidgetAttribute(C.Qt_WA_X11NetWmWindowTypeUtility())
	WA_X11NETWMWINDOWTYPESPLASH        = WidgetAttribute(C.Qt_WA_X11NetWmWindowTypeSplash())
	WA_X11NETWMWINDOWTYPEDIALOG        = WidgetAttribute(C.Qt_WA_X11NetWmWindowTypeDialog())
	WA_X11NETWMWINDOWTYPEDROPDOWNMENU  = WidgetAttribute(C.Qt_WA_X11NetWmWindowTypeDropDownMenu())
	WA_X11NETWMWINDOWTYPEPOPUPMENU     = WidgetAttribute(C.Qt_WA_X11NetWmWindowTypePopupMenu())
	WA_X11NETWMWINDOWTYPETOOLTIP       = WidgetAttribute(C.Qt_WA_X11NetWmWindowTypeToolTip())
	WA_X11NETWMWINDOWTYPENOTIFICATION  = WidgetAttribute(C.Qt_WA_X11NetWmWindowTypeNotification())
	WA_X11NETWMWINDOWTYPECOMBO         = WidgetAttribute(C.Qt_WA_X11NetWmWindowTypeCombo())
	WA_X11NETWMWINDOWTYPEDND           = WidgetAttribute(C.Qt_WA_X11NetWmWindowTypeDND())
	WA_MACFRAMEWORKSCALED              = WidgetAttribute(C.Qt_WA_MacFrameworkScaled())
	WA_ACCEPTTOUCHEVENTS               = WidgetAttribute(C.Qt_WA_AcceptTouchEvents())
	WA_TOUCHPADACCEPTSINGLETOUCHEVENTS = WidgetAttribute(C.Qt_WA_TouchPadAcceptSingleTouchEvents())
	WA_X11DONOTACCEPTFOCUS             = WidgetAttribute(C.Qt_WA_X11DoNotAcceptFocus())
)

type WindowFrameSection int

var (
	NOSECTION          = WindowFrameSection(C.Qt_NoSection())
	LEFTSECTION        = WindowFrameSection(C.Qt_LeftSection())
	TOPLEFTSECTION     = WindowFrameSection(C.Qt_TopLeftSection())
	TOPSECTION         = WindowFrameSection(C.Qt_TopSection())
	TOPRIGHTSECTION    = WindowFrameSection(C.Qt_TopRightSection())
	RIGHTSECTION       = WindowFrameSection(C.Qt_RightSection())
	BOTTOMRIGHTSECTION = WindowFrameSection(C.Qt_BottomRightSection())
	BOTTOMSECTION      = WindowFrameSection(C.Qt_BottomSection())
	BOTTOMLEFTSECTION  = WindowFrameSection(C.Qt_BottomLeftSection())
	TITLEBARAREA       = WindowFrameSection(C.Qt_TitleBarArea())
)

type WindowModality int

var (
	NONMODAL         = WindowModality(C.Qt_NonModal())
	WINDOWMODAL      = WindowModality(C.Qt_WindowModal())
	APPLICATIONMODAL = WindowModality(C.Qt_ApplicationModal())
)

type WindowState int

var (
	WINDOWNOSTATE    = WindowState(C.Qt_WindowNoState())
	WINDOWMINIMIZED  = WindowState(C.Qt_WindowMinimized())
	WINDOWMAXIMIZED  = WindowState(C.Qt_WindowMaximized())
	WINDOWFULLSCREEN = WindowState(C.Qt_WindowFullScreen())
	WINDOWACTIVE     = WindowState(C.Qt_WindowActive())
)

type WindowType int

var (
	WIDGET                        = WindowType(C.Qt_Widget())
	WINDOW                        = WindowType(C.Qt_Window())
	DIALOG                        = WindowType(C.Qt_Dialog())
	SHEET                         = WindowType(C.Qt_Sheet())
	DRAWER                        = WindowType(C.Qt_Drawer())
	POPUP                         = WindowType(C.Qt_Popup())
	TOOL                          = WindowType(C.Qt_Tool())
	TOOLTIP                       = WindowType(C.Qt_ToolTip())
	SPLASHSCREEN                  = WindowType(C.Qt_SplashScreen())
	DESKTOP                       = WindowType(C.Qt_Desktop())
	SUBWINDOW                     = WindowType(C.Qt_SubWindow())
	FOREIGNWINDOW                 = WindowType(C.Qt_ForeignWindow())
	COVERWINDOW                   = WindowType(C.Qt_CoverWindow())
	MSWINDOWSFIXEDSIZEDIALOGHINT  = WindowType(C.Qt_MSWindowsFixedSizeDialogHint())
	MSWINDOWSOWNDC                = WindowType(C.Qt_MSWindowsOwnDC())
	BYPASSWINDOWMANAGERHINT       = WindowType(C.Qt_BypassWindowManagerHint())
	X11BYPASSWINDOWMANAGERHINT    = WindowType(C.Qt_X11BypassWindowManagerHint())
	FRAMELESSWINDOWHINT           = WindowType(C.Qt_FramelessWindowHint())
	NODROPSHADOWWINDOWHINT        = WindowType(C.Qt_NoDropShadowWindowHint())
	CUSTOMIZEWINDOWHINT           = WindowType(C.Qt_CustomizeWindowHint())
	WINDOWTITLEHINT               = WindowType(C.Qt_WindowTitleHint())
	WINDOWSYSTEMMENUHINT          = WindowType(C.Qt_WindowSystemMenuHint())
	WINDOWMINIMIZEBUTTONHINT      = WindowType(C.Qt_WindowMinimizeButtonHint())
	WINDOWMAXIMIZEBUTTONHINT      = WindowType(C.Qt_WindowMaximizeButtonHint())
	WINDOWMINMAXBUTTONSHINT       = WindowType(C.Qt_WindowMinMaxButtonsHint())
	WINDOWCLOSEBUTTONHINT         = WindowType(C.Qt_WindowCloseButtonHint())
	WINDOWCONTEXTHELPBUTTONHINT   = WindowType(C.Qt_WindowContextHelpButtonHint())
	MACWINDOWTOOLBARBUTTONHINT    = WindowType(C.Qt_MacWindowToolBarButtonHint())
	WINDOWFULLSCREENBUTTONHINT    = WindowType(C.Qt_WindowFullscreenButtonHint())
	BYPASSGRAPHICSPROXYWIDGET     = WindowType(C.Qt_BypassGraphicsProxyWidget())
	WINDOWSHADEBUTTONHINT         = WindowType(C.Qt_WindowShadeButtonHint())
	WINDOWSTAYSONTOPHINT          = WindowType(C.Qt_WindowStaysOnTopHint())
	WINDOWSTAYSONBOTTOMHINT       = WindowType(C.Qt_WindowStaysOnBottomHint())
	WINDOWOKBUTTONHINT            = WindowType(C.Qt_WindowOkButtonHint())
	WINDOWCANCELBUTTONHINT        = WindowType(C.Qt_WindowCancelButtonHint())
	WINDOWTRANSPARENTFORINPUT     = WindowType(C.Qt_WindowTransparentForInput())
	WINDOWOVERRIDESSYSTEMGESTURES = WindowType(C.Qt_WindowOverridesSystemGestures())
	WINDOWDOESNOTACCEPTFOCUS      = WindowType(C.Qt_WindowDoesNotAcceptFocus())
	WINDOWTYPE_MASK               = WindowType(C.Qt_WindowType_Mask())
)
