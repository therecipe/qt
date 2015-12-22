package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"unsafe"
)

type QStyle struct {
	core.QObject
}

type QStyle_ITF interface {
	core.QObject_ITF
	QStyle_PTR() *QStyle
}

func PointerFromQStyle(ptr QStyle_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QStyle_PTR().Pointer()
	}
	return nil
}

func NewQStyleFromPointer(ptr unsafe.Pointer) *QStyle {
	var n = new(QStyle)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QStyle_") {
		n.SetObjectName("QStyle_" + qt.Identifier())
	}
	return n
}

func (ptr *QStyle) QStyle_PTR() *QStyle {
	return ptr
}

//QStyle::ComplexControl
type QStyle__ComplexControl int64

const (
	QStyle__CC_SpinBox     = QStyle__ComplexControl(0)
	QStyle__CC_ComboBox    = QStyle__ComplexControl(1)
	QStyle__CC_ScrollBar   = QStyle__ComplexControl(2)
	QStyle__CC_Slider      = QStyle__ComplexControl(3)
	QStyle__CC_ToolButton  = QStyle__ComplexControl(4)
	QStyle__CC_TitleBar    = QStyle__ComplexControl(5)
	QStyle__CC_Dial        = QStyle__ComplexControl(6)
	QStyle__CC_GroupBox    = QStyle__ComplexControl(7)
	QStyle__CC_MdiControls = QStyle__ComplexControl(8)
	QStyle__CC_CustomBase  = QStyle__ComplexControl(0xf0000000)
)

//QStyle::ContentsType
type QStyle__ContentsType int64

const (
	QStyle__CT_PushButton    = QStyle__ContentsType(0)
	QStyle__CT_CheckBox      = QStyle__ContentsType(1)
	QStyle__CT_RadioButton   = QStyle__ContentsType(2)
	QStyle__CT_ToolButton    = QStyle__ContentsType(3)
	QStyle__CT_ComboBox      = QStyle__ContentsType(4)
	QStyle__CT_Splitter      = QStyle__ContentsType(5)
	QStyle__CT_ProgressBar   = QStyle__ContentsType(6)
	QStyle__CT_MenuItem      = QStyle__ContentsType(7)
	QStyle__CT_MenuBarItem   = QStyle__ContentsType(8)
	QStyle__CT_MenuBar       = QStyle__ContentsType(9)
	QStyle__CT_Menu          = QStyle__ContentsType(10)
	QStyle__CT_TabBarTab     = QStyle__ContentsType(11)
	QStyle__CT_Slider        = QStyle__ContentsType(12)
	QStyle__CT_ScrollBar     = QStyle__ContentsType(13)
	QStyle__CT_LineEdit      = QStyle__ContentsType(14)
	QStyle__CT_SpinBox       = QStyle__ContentsType(15)
	QStyle__CT_SizeGrip      = QStyle__ContentsType(16)
	QStyle__CT_TabWidget     = QStyle__ContentsType(17)
	QStyle__CT_DialogButtons = QStyle__ContentsType(18)
	QStyle__CT_HeaderSection = QStyle__ContentsType(19)
	QStyle__CT_GroupBox      = QStyle__ContentsType(20)
	QStyle__CT_MdiControls   = QStyle__ContentsType(21)
	QStyle__CT_ItemViewItem  = QStyle__ContentsType(22)
	QStyle__CT_CustomBase    = QStyle__ContentsType(0xf0000000)
)

//QStyle::ControlElement
type QStyle__ControlElement int64

const (
	QStyle__CE_PushButton          = QStyle__ControlElement(0)
	QStyle__CE_PushButtonBevel     = QStyle__ControlElement(1)
	QStyle__CE_PushButtonLabel     = QStyle__ControlElement(2)
	QStyle__CE_CheckBox            = QStyle__ControlElement(3)
	QStyle__CE_CheckBoxLabel       = QStyle__ControlElement(4)
	QStyle__CE_RadioButton         = QStyle__ControlElement(5)
	QStyle__CE_RadioButtonLabel    = QStyle__ControlElement(6)
	QStyle__CE_TabBarTab           = QStyle__ControlElement(7)
	QStyle__CE_TabBarTabShape      = QStyle__ControlElement(8)
	QStyle__CE_TabBarTabLabel      = QStyle__ControlElement(9)
	QStyle__CE_ProgressBar         = QStyle__ControlElement(10)
	QStyle__CE_ProgressBarGroove   = QStyle__ControlElement(11)
	QStyle__CE_ProgressBarContents = QStyle__ControlElement(12)
	QStyle__CE_ProgressBarLabel    = QStyle__ControlElement(13)
	QStyle__CE_MenuItem            = QStyle__ControlElement(14)
	QStyle__CE_MenuScroller        = QStyle__ControlElement(15)
	QStyle__CE_MenuVMargin         = QStyle__ControlElement(16)
	QStyle__CE_MenuHMargin         = QStyle__ControlElement(17)
	QStyle__CE_MenuTearoff         = QStyle__ControlElement(18)
	QStyle__CE_MenuEmptyArea       = QStyle__ControlElement(19)
	QStyle__CE_MenuBarItem         = QStyle__ControlElement(20)
	QStyle__CE_MenuBarEmptyArea    = QStyle__ControlElement(21)
	QStyle__CE_ToolButtonLabel     = QStyle__ControlElement(22)
	QStyle__CE_Header              = QStyle__ControlElement(23)
	QStyle__CE_HeaderSection       = QStyle__ControlElement(24)
	QStyle__CE_HeaderLabel         = QStyle__ControlElement(25)
	QStyle__CE_ToolBoxTab          = QStyle__ControlElement(26)
	QStyle__CE_SizeGrip            = QStyle__ControlElement(27)
	QStyle__CE_Splitter            = QStyle__ControlElement(28)
	QStyle__CE_RubberBand          = QStyle__ControlElement(29)
	QStyle__CE_DockWidgetTitle     = QStyle__ControlElement(30)
	QStyle__CE_ScrollBarAddLine    = QStyle__ControlElement(31)
	QStyle__CE_ScrollBarSubLine    = QStyle__ControlElement(32)
	QStyle__CE_ScrollBarAddPage    = QStyle__ControlElement(33)
	QStyle__CE_ScrollBarSubPage    = QStyle__ControlElement(34)
	QStyle__CE_ScrollBarSlider     = QStyle__ControlElement(35)
	QStyle__CE_ScrollBarFirst      = QStyle__ControlElement(36)
	QStyle__CE_ScrollBarLast       = QStyle__ControlElement(37)
	QStyle__CE_FocusFrame          = QStyle__ControlElement(38)
	QStyle__CE_ComboBoxLabel       = QStyle__ControlElement(39)
	QStyle__CE_ToolBar             = QStyle__ControlElement(40)
	QStyle__CE_ToolBoxTabShape     = QStyle__ControlElement(41)
	QStyle__CE_ToolBoxTabLabel     = QStyle__ControlElement(42)
	QStyle__CE_HeaderEmptyArea     = QStyle__ControlElement(43)
	QStyle__CE_ColumnViewGrip      = QStyle__ControlElement(44)
	QStyle__CE_ItemViewItem        = QStyle__ControlElement(45)
	QStyle__CE_ShapedFrame         = QStyle__ControlElement(46)
	QStyle__CE_CustomBase          = QStyle__ControlElement(0xf0000000)
)

//QStyle::PixelMetric
type QStyle__PixelMetric int64

var (
	QStyle__PM_ButtonMargin                       = QStyle__PixelMetric(0)
	QStyle__PM_ButtonDefaultIndicator             = QStyle__PixelMetric(1)
	QStyle__PM_MenuButtonIndicator                = QStyle__PixelMetric(2)
	QStyle__PM_ButtonShiftHorizontal              = QStyle__PixelMetric(3)
	QStyle__PM_ButtonShiftVertical                = QStyle__PixelMetric(4)
	QStyle__PM_DefaultFrameWidth                  = QStyle__PixelMetric(5)
	QStyle__PM_SpinBoxFrameWidth                  = QStyle__PixelMetric(6)
	QStyle__PM_ComboBoxFrameWidth                 = QStyle__PixelMetric(7)
	QStyle__PM_MaximumDragDistance                = QStyle__PixelMetric(8)
	QStyle__PM_ScrollBarExtent                    = QStyle__PixelMetric(9)
	QStyle__PM_ScrollBarSliderMin                 = QStyle__PixelMetric(10)
	QStyle__PM_SliderThickness                    = QStyle__PixelMetric(11)
	QStyle__PM_SliderControlThickness             = QStyle__PixelMetric(12)
	QStyle__PM_SliderLength                       = QStyle__PixelMetric(13)
	QStyle__PM_SliderTickmarkOffset               = QStyle__PixelMetric(14)
	QStyle__PM_SliderSpaceAvailable               = QStyle__PixelMetric(15)
	QStyle__PM_DockWidgetSeparatorExtent          = QStyle__PixelMetric(16)
	QStyle__PM_DockWidgetHandleExtent             = QStyle__PixelMetric(17)
	QStyle__PM_DockWidgetFrameWidth               = QStyle__PixelMetric(18)
	QStyle__PM_TabBarTabOverlap                   = QStyle__PixelMetric(19)
	QStyle__PM_TabBarTabHSpace                    = QStyle__PixelMetric(20)
	QStyle__PM_TabBarTabVSpace                    = QStyle__PixelMetric(21)
	QStyle__PM_TabBarBaseHeight                   = QStyle__PixelMetric(22)
	QStyle__PM_TabBarBaseOverlap                  = QStyle__PixelMetric(23)
	QStyle__PM_ProgressBarChunkWidth              = QStyle__PixelMetric(24)
	QStyle__PM_SplitterWidth                      = QStyle__PixelMetric(25)
	QStyle__PM_TitleBarHeight                     = QStyle__PixelMetric(26)
	QStyle__PM_MenuScrollerHeight                 = QStyle__PixelMetric(27)
	QStyle__PM_MenuHMargin                        = QStyle__PixelMetric(28)
	QStyle__PM_MenuVMargin                        = QStyle__PixelMetric(29)
	QStyle__PM_MenuPanelWidth                     = QStyle__PixelMetric(30)
	QStyle__PM_MenuTearoffHeight                  = QStyle__PixelMetric(31)
	QStyle__PM_MenuDesktopFrameWidth              = QStyle__PixelMetric(32)
	QStyle__PM_MenuBarPanelWidth                  = QStyle__PixelMetric(33)
	QStyle__PM_MenuBarItemSpacing                 = QStyle__PixelMetric(34)
	QStyle__PM_MenuBarVMargin                     = QStyle__PixelMetric(35)
	QStyle__PM_MenuBarHMargin                     = QStyle__PixelMetric(36)
	QStyle__PM_IndicatorWidth                     = QStyle__PixelMetric(37)
	QStyle__PM_IndicatorHeight                    = QStyle__PixelMetric(38)
	QStyle__PM_ExclusiveIndicatorWidth            = QStyle__PixelMetric(39)
	QStyle__PM_ExclusiveIndicatorHeight           = QStyle__PixelMetric(40)
	QStyle__PM_DialogButtonsSeparator             = QStyle__PixelMetric(41)
	QStyle__PM_DialogButtonsButtonWidth           = QStyle__PixelMetric(42)
	QStyle__PM_DialogButtonsButtonHeight          = QStyle__PixelMetric(43)
	QStyle__PM_MdiSubWindowFrameWidth             = QStyle__PixelMetric(44)
	QStyle__PM_MDIFrameWidth                      = QStyle__PixelMetric(QStyle__PM_MdiSubWindowFrameWidth)
	QStyle__PM_MdiSubWindowMinimizedWidth         = QStyle__PixelMetric(C.QStyle_PM_MdiSubWindowMinimizedWidth_Type())
	QStyle__PM_MDIMinimizedWidth                  = QStyle__PixelMetric(QStyle__PM_MdiSubWindowMinimizedWidth)
	QStyle__PM_HeaderMargin                       = QStyle__PixelMetric(C.QStyle_PM_HeaderMargin_Type())
	QStyle__PM_HeaderMarkSize                     = QStyle__PixelMetric(C.QStyle_PM_HeaderMarkSize_Type())
	QStyle__PM_HeaderGripMargin                   = QStyle__PixelMetric(C.QStyle_PM_HeaderGripMargin_Type())
	QStyle__PM_TabBarTabShiftHorizontal           = QStyle__PixelMetric(C.QStyle_PM_TabBarTabShiftHorizontal_Type())
	QStyle__PM_TabBarTabShiftVertical             = QStyle__PixelMetric(C.QStyle_PM_TabBarTabShiftVertical_Type())
	QStyle__PM_TabBarScrollButtonWidth            = QStyle__PixelMetric(C.QStyle_PM_TabBarScrollButtonWidth_Type())
	QStyle__PM_ToolBarFrameWidth                  = QStyle__PixelMetric(C.QStyle_PM_ToolBarFrameWidth_Type())
	QStyle__PM_ToolBarHandleExtent                = QStyle__PixelMetric(C.QStyle_PM_ToolBarHandleExtent_Type())
	QStyle__PM_ToolBarItemSpacing                 = QStyle__PixelMetric(C.QStyle_PM_ToolBarItemSpacing_Type())
	QStyle__PM_ToolBarItemMargin                  = QStyle__PixelMetric(C.QStyle_PM_ToolBarItemMargin_Type())
	QStyle__PM_ToolBarSeparatorExtent             = QStyle__PixelMetric(C.QStyle_PM_ToolBarSeparatorExtent_Type())
	QStyle__PM_ToolBarExtensionExtent             = QStyle__PixelMetric(C.QStyle_PM_ToolBarExtensionExtent_Type())
	QStyle__PM_SpinBoxSliderHeight                = QStyle__PixelMetric(C.QStyle_PM_SpinBoxSliderHeight_Type())
	QStyle__PM_DefaultTopLevelMargin              = QStyle__PixelMetric(C.QStyle_PM_DefaultTopLevelMargin_Type())
	QStyle__PM_DefaultChildMargin                 = QStyle__PixelMetric(C.QStyle_PM_DefaultChildMargin_Type())
	QStyle__PM_DefaultLayoutSpacing               = QStyle__PixelMetric(C.QStyle_PM_DefaultLayoutSpacing_Type())
	QStyle__PM_ToolBarIconSize                    = QStyle__PixelMetric(C.QStyle_PM_ToolBarIconSize_Type())
	QStyle__PM_ListViewIconSize                   = QStyle__PixelMetric(C.QStyle_PM_ListViewIconSize_Type())
	QStyle__PM_IconViewIconSize                   = QStyle__PixelMetric(C.QStyle_PM_IconViewIconSize_Type())
	QStyle__PM_SmallIconSize                      = QStyle__PixelMetric(C.QStyle_PM_SmallIconSize_Type())
	QStyle__PM_LargeIconSize                      = QStyle__PixelMetric(C.QStyle_PM_LargeIconSize_Type())
	QStyle__PM_FocusFrameVMargin                  = QStyle__PixelMetric(C.QStyle_PM_FocusFrameVMargin_Type())
	QStyle__PM_FocusFrameHMargin                  = QStyle__PixelMetric(C.QStyle_PM_FocusFrameHMargin_Type())
	QStyle__PM_ToolTipLabelFrameWidth             = QStyle__PixelMetric(C.QStyle_PM_ToolTipLabelFrameWidth_Type())
	QStyle__PM_CheckBoxLabelSpacing               = QStyle__PixelMetric(C.QStyle_PM_CheckBoxLabelSpacing_Type())
	QStyle__PM_TabBarIconSize                     = QStyle__PixelMetric(C.QStyle_PM_TabBarIconSize_Type())
	QStyle__PM_SizeGripSize                       = QStyle__PixelMetric(C.QStyle_PM_SizeGripSize_Type())
	QStyle__PM_DockWidgetTitleMargin              = QStyle__PixelMetric(C.QStyle_PM_DockWidgetTitleMargin_Type())
	QStyle__PM_MessageBoxIconSize                 = QStyle__PixelMetric(C.QStyle_PM_MessageBoxIconSize_Type())
	QStyle__PM_ButtonIconSize                     = QStyle__PixelMetric(C.QStyle_PM_ButtonIconSize_Type())
	QStyle__PM_DockWidgetTitleBarButtonMargin     = QStyle__PixelMetric(C.QStyle_PM_DockWidgetTitleBarButtonMargin_Type())
	QStyle__PM_RadioButtonLabelSpacing            = QStyle__PixelMetric(C.QStyle_PM_RadioButtonLabelSpacing_Type())
	QStyle__PM_LayoutLeftMargin                   = QStyle__PixelMetric(C.QStyle_PM_LayoutLeftMargin_Type())
	QStyle__PM_LayoutTopMargin                    = QStyle__PixelMetric(C.QStyle_PM_LayoutTopMargin_Type())
	QStyle__PM_LayoutRightMargin                  = QStyle__PixelMetric(C.QStyle_PM_LayoutRightMargin_Type())
	QStyle__PM_LayoutBottomMargin                 = QStyle__PixelMetric(C.QStyle_PM_LayoutBottomMargin_Type())
	QStyle__PM_LayoutHorizontalSpacing            = QStyle__PixelMetric(C.QStyle_PM_LayoutHorizontalSpacing_Type())
	QStyle__PM_LayoutVerticalSpacing              = QStyle__PixelMetric(C.QStyle_PM_LayoutVerticalSpacing_Type())
	QStyle__PM_TabBar_ScrollButtonOverlap         = QStyle__PixelMetric(C.QStyle_PM_TabBar_ScrollButtonOverlap_Type())
	QStyle__PM_TextCursorWidth                    = QStyle__PixelMetric(C.QStyle_PM_TextCursorWidth_Type())
	QStyle__PM_TabCloseIndicatorWidth             = QStyle__PixelMetric(C.QStyle_PM_TabCloseIndicatorWidth_Type())
	QStyle__PM_TabCloseIndicatorHeight            = QStyle__PixelMetric(C.QStyle_PM_TabCloseIndicatorHeight_Type())
	QStyle__PM_ScrollView_ScrollBarSpacing        = QStyle__PixelMetric(C.QStyle_PM_ScrollView_ScrollBarSpacing_Type())
	QStyle__PM_ScrollView_ScrollBarOverlap        = QStyle__PixelMetric(C.QStyle_PM_ScrollView_ScrollBarOverlap_Type())
	QStyle__PM_SubMenuOverlap                     = QStyle__PixelMetric(C.QStyle_PM_SubMenuOverlap_Type())
	QStyle__PM_TreeViewIndentation                = QStyle__PixelMetric(C.QStyle_PM_TreeViewIndentation_Type())
	QStyle__PM_HeaderDefaultSectionSizeHorizontal = QStyle__PixelMetric(C.QStyle_PM_HeaderDefaultSectionSizeHorizontal_Type())
	QStyle__PM_HeaderDefaultSectionSizeVertical   = QStyle__PixelMetric(C.QStyle_PM_HeaderDefaultSectionSizeVertical_Type())
	QStyle__PM_CustomBase                         = QStyle__PixelMetric(0xf0000000)
)

//QStyle::PrimitiveElement
type QStyle__PrimitiveElement int64

var (
	QStyle__PE_Frame                           = QStyle__PrimitiveElement(0)
	QStyle__PE_FrameDefaultButton              = QStyle__PrimitiveElement(1)
	QStyle__PE_FrameDockWidget                 = QStyle__PrimitiveElement(2)
	QStyle__PE_FrameFocusRect                  = QStyle__PrimitiveElement(3)
	QStyle__PE_FrameGroupBox                   = QStyle__PrimitiveElement(4)
	QStyle__PE_FrameLineEdit                   = QStyle__PrimitiveElement(5)
	QStyle__PE_FrameMenu                       = QStyle__PrimitiveElement(6)
	QStyle__PE_FrameStatusBar                  = QStyle__PrimitiveElement(7)
	QStyle__PE_FrameStatusBarItem              = QStyle__PrimitiveElement(QStyle__PE_FrameStatusBar)
	QStyle__PE_FrameTabWidget                  = QStyle__PrimitiveElement(C.QStyle_PE_FrameTabWidget_Type())
	QStyle__PE_FrameWindow                     = QStyle__PrimitiveElement(C.QStyle_PE_FrameWindow_Type())
	QStyle__PE_FrameButtonBevel                = QStyle__PrimitiveElement(C.QStyle_PE_FrameButtonBevel_Type())
	QStyle__PE_FrameButtonTool                 = QStyle__PrimitiveElement(C.QStyle_PE_FrameButtonTool_Type())
	QStyle__PE_FrameTabBarBase                 = QStyle__PrimitiveElement(C.QStyle_PE_FrameTabBarBase_Type())
	QStyle__PE_PanelButtonCommand              = QStyle__PrimitiveElement(C.QStyle_PE_PanelButtonCommand_Type())
	QStyle__PE_PanelButtonBevel                = QStyle__PrimitiveElement(C.QStyle_PE_PanelButtonBevel_Type())
	QStyle__PE_PanelButtonTool                 = QStyle__PrimitiveElement(C.QStyle_PE_PanelButtonTool_Type())
	QStyle__PE_PanelMenuBar                    = QStyle__PrimitiveElement(C.QStyle_PE_PanelMenuBar_Type())
	QStyle__PE_PanelToolBar                    = QStyle__PrimitiveElement(C.QStyle_PE_PanelToolBar_Type())
	QStyle__PE_PanelLineEdit                   = QStyle__PrimitiveElement(C.QStyle_PE_PanelLineEdit_Type())
	QStyle__PE_IndicatorArrowDown              = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorArrowDown_Type())
	QStyle__PE_IndicatorArrowLeft              = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorArrowLeft_Type())
	QStyle__PE_IndicatorArrowRight             = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorArrowRight_Type())
	QStyle__PE_IndicatorArrowUp                = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorArrowUp_Type())
	QStyle__PE_IndicatorBranch                 = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorBranch_Type())
	QStyle__PE_IndicatorButtonDropDown         = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorButtonDropDown_Type())
	QStyle__PE_IndicatorViewItemCheck          = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorViewItemCheck_Type())
	QStyle__PE_IndicatorItemViewItemCheck      = QStyle__PrimitiveElement(QStyle__PE_IndicatorViewItemCheck)
	QStyle__PE_IndicatorCheckBox               = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorCheckBox_Type())
	QStyle__PE_IndicatorDockWidgetResizeHandle = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorDockWidgetResizeHandle_Type())
	QStyle__PE_IndicatorHeaderArrow            = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorHeaderArrow_Type())
	QStyle__PE_IndicatorMenuCheckMark          = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorMenuCheckMark_Type())
	QStyle__PE_IndicatorProgressChunk          = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorProgressChunk_Type())
	QStyle__PE_IndicatorRadioButton            = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorRadioButton_Type())
	QStyle__PE_IndicatorSpinDown               = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorSpinDown_Type())
	QStyle__PE_IndicatorSpinMinus              = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorSpinMinus_Type())
	QStyle__PE_IndicatorSpinPlus               = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorSpinPlus_Type())
	QStyle__PE_IndicatorSpinUp                 = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorSpinUp_Type())
	QStyle__PE_IndicatorToolBarHandle          = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorToolBarHandle_Type())
	QStyle__PE_IndicatorToolBarSeparator       = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorToolBarSeparator_Type())
	QStyle__PE_PanelTipLabel                   = QStyle__PrimitiveElement(C.QStyle_PE_PanelTipLabel_Type())
	QStyle__PE_IndicatorTabTear                = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorTabTear_Type())
	QStyle__PE_PanelScrollAreaCorner           = QStyle__PrimitiveElement(C.QStyle_PE_PanelScrollAreaCorner_Type())
	QStyle__PE_Widget                          = QStyle__PrimitiveElement(C.QStyle_PE_Widget_Type())
	QStyle__PE_IndicatorColumnViewArrow        = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorColumnViewArrow_Type())
	QStyle__PE_IndicatorItemViewItemDrop       = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorItemViewItemDrop_Type())
	QStyle__PE_PanelItemViewItem               = QStyle__PrimitiveElement(C.QStyle_PE_PanelItemViewItem_Type())
	QStyle__PE_PanelItemViewRow                = QStyle__PrimitiveElement(C.QStyle_PE_PanelItemViewRow_Type())
	QStyle__PE_PanelStatusBar                  = QStyle__PrimitiveElement(C.QStyle_PE_PanelStatusBar_Type())
	QStyle__PE_IndicatorTabClose               = QStyle__PrimitiveElement(C.QStyle_PE_IndicatorTabClose_Type())
	QStyle__PE_PanelMenu                       = QStyle__PrimitiveElement(C.QStyle_PE_PanelMenu_Type())
	QStyle__PE_CustomBase                      = QStyle__PrimitiveElement(0xf000000)
)

//QStyle::RequestSoftwareInputPanel
type QStyle__RequestSoftwareInputPanel int64

const (
	QStyle__RSIP_OnMouseClickAndAlreadyFocused = QStyle__RequestSoftwareInputPanel(0)
	QStyle__RSIP_OnMouseClick                  = QStyle__RequestSoftwareInputPanel(1)
)

//QStyle::StandardPixmap
type QStyle__StandardPixmap int64

const (
	QStyle__SP_TitleBarMenuButton               = QStyle__StandardPixmap(0)
	QStyle__SP_TitleBarMinButton                = QStyle__StandardPixmap(1)
	QStyle__SP_TitleBarMaxButton                = QStyle__StandardPixmap(2)
	QStyle__SP_TitleBarCloseButton              = QStyle__StandardPixmap(3)
	QStyle__SP_TitleBarNormalButton             = QStyle__StandardPixmap(4)
	QStyle__SP_TitleBarShadeButton              = QStyle__StandardPixmap(5)
	QStyle__SP_TitleBarUnshadeButton            = QStyle__StandardPixmap(6)
	QStyle__SP_TitleBarContextHelpButton        = QStyle__StandardPixmap(7)
	QStyle__SP_DockWidgetCloseButton            = QStyle__StandardPixmap(8)
	QStyle__SP_MessageBoxInformation            = QStyle__StandardPixmap(9)
	QStyle__SP_MessageBoxWarning                = QStyle__StandardPixmap(10)
	QStyle__SP_MessageBoxCritical               = QStyle__StandardPixmap(11)
	QStyle__SP_MessageBoxQuestion               = QStyle__StandardPixmap(12)
	QStyle__SP_DesktopIcon                      = QStyle__StandardPixmap(13)
	QStyle__SP_TrashIcon                        = QStyle__StandardPixmap(14)
	QStyle__SP_ComputerIcon                     = QStyle__StandardPixmap(15)
	QStyle__SP_DriveFDIcon                      = QStyle__StandardPixmap(16)
	QStyle__SP_DriveHDIcon                      = QStyle__StandardPixmap(17)
	QStyle__SP_DriveCDIcon                      = QStyle__StandardPixmap(18)
	QStyle__SP_DriveDVDIcon                     = QStyle__StandardPixmap(19)
	QStyle__SP_DriveNetIcon                     = QStyle__StandardPixmap(20)
	QStyle__SP_DirOpenIcon                      = QStyle__StandardPixmap(21)
	QStyle__SP_DirClosedIcon                    = QStyle__StandardPixmap(22)
	QStyle__SP_DirLinkIcon                      = QStyle__StandardPixmap(23)
	QStyle__SP_DirLinkOpenIcon                  = QStyle__StandardPixmap(24)
	QStyle__SP_FileIcon                         = QStyle__StandardPixmap(25)
	QStyle__SP_FileLinkIcon                     = QStyle__StandardPixmap(26)
	QStyle__SP_ToolBarHorizontalExtensionButton = QStyle__StandardPixmap(27)
	QStyle__SP_ToolBarVerticalExtensionButton   = QStyle__StandardPixmap(28)
	QStyle__SP_FileDialogStart                  = QStyle__StandardPixmap(29)
	QStyle__SP_FileDialogEnd                    = QStyle__StandardPixmap(30)
	QStyle__SP_FileDialogToParent               = QStyle__StandardPixmap(31)
	QStyle__SP_FileDialogNewFolder              = QStyle__StandardPixmap(32)
	QStyle__SP_FileDialogDetailedView           = QStyle__StandardPixmap(33)
	QStyle__SP_FileDialogInfoView               = QStyle__StandardPixmap(34)
	QStyle__SP_FileDialogContentsView           = QStyle__StandardPixmap(35)
	QStyle__SP_FileDialogListView               = QStyle__StandardPixmap(36)
	QStyle__SP_FileDialogBack                   = QStyle__StandardPixmap(37)
	QStyle__SP_DirIcon                          = QStyle__StandardPixmap(38)
	QStyle__SP_DialogOkButton                   = QStyle__StandardPixmap(39)
	QStyle__SP_DialogCancelButton               = QStyle__StandardPixmap(40)
	QStyle__SP_DialogHelpButton                 = QStyle__StandardPixmap(41)
	QStyle__SP_DialogOpenButton                 = QStyle__StandardPixmap(42)
	QStyle__SP_DialogSaveButton                 = QStyle__StandardPixmap(43)
	QStyle__SP_DialogCloseButton                = QStyle__StandardPixmap(44)
	QStyle__SP_DialogApplyButton                = QStyle__StandardPixmap(45)
	QStyle__SP_DialogResetButton                = QStyle__StandardPixmap(46)
	QStyle__SP_DialogDiscardButton              = QStyle__StandardPixmap(47)
	QStyle__SP_DialogYesButton                  = QStyle__StandardPixmap(48)
	QStyle__SP_DialogNoButton                   = QStyle__StandardPixmap(49)
	QStyle__SP_ArrowUp                          = QStyle__StandardPixmap(50)
	QStyle__SP_ArrowDown                        = QStyle__StandardPixmap(51)
	QStyle__SP_ArrowLeft                        = QStyle__StandardPixmap(52)
	QStyle__SP_ArrowRight                       = QStyle__StandardPixmap(53)
	QStyle__SP_ArrowBack                        = QStyle__StandardPixmap(54)
	QStyle__SP_ArrowForward                     = QStyle__StandardPixmap(55)
	QStyle__SP_DirHomeIcon                      = QStyle__StandardPixmap(56)
	QStyle__SP_CommandLink                      = QStyle__StandardPixmap(57)
	QStyle__SP_VistaShield                      = QStyle__StandardPixmap(58)
	QStyle__SP_BrowserReload                    = QStyle__StandardPixmap(59)
	QStyle__SP_BrowserStop                      = QStyle__StandardPixmap(60)
	QStyle__SP_MediaPlay                        = QStyle__StandardPixmap(61)
	QStyle__SP_MediaStop                        = QStyle__StandardPixmap(62)
	QStyle__SP_MediaPause                       = QStyle__StandardPixmap(63)
	QStyle__SP_MediaSkipForward                 = QStyle__StandardPixmap(64)
	QStyle__SP_MediaSkipBackward                = QStyle__StandardPixmap(65)
	QStyle__SP_MediaSeekForward                 = QStyle__StandardPixmap(66)
	QStyle__SP_MediaSeekBackward                = QStyle__StandardPixmap(67)
	QStyle__SP_MediaVolume                      = QStyle__StandardPixmap(68)
	QStyle__SP_MediaVolumeMuted                 = QStyle__StandardPixmap(69)
	QStyle__SP_LineEditClearButton              = QStyle__StandardPixmap(70)
	QStyle__SP_CustomBase                       = QStyle__StandardPixmap(0xf0000000)
)

//QStyle::StateFlag
type QStyle__StateFlag int64

const (
	QStyle__State_None                = QStyle__StateFlag(0x00000000)
	QStyle__State_Enabled             = QStyle__StateFlag(0x00000001)
	QStyle__State_Raised              = QStyle__StateFlag(0x00000002)
	QStyle__State_Sunken              = QStyle__StateFlag(0x00000004)
	QStyle__State_Off                 = QStyle__StateFlag(0x00000008)
	QStyle__State_NoChange            = QStyle__StateFlag(0x00000010)
	QStyle__State_On                  = QStyle__StateFlag(0x00000020)
	QStyle__State_DownArrow           = QStyle__StateFlag(0x00000040)
	QStyle__State_Horizontal          = QStyle__StateFlag(0x00000080)
	QStyle__State_HasFocus            = QStyle__StateFlag(0x00000100)
	QStyle__State_Top                 = QStyle__StateFlag(0x00000200)
	QStyle__State_Bottom              = QStyle__StateFlag(0x00000400)
	QStyle__State_FocusAtBorder       = QStyle__StateFlag(0x00000800)
	QStyle__State_AutoRaise           = QStyle__StateFlag(0x00001000)
	QStyle__State_MouseOver           = QStyle__StateFlag(0x00002000)
	QStyle__State_UpArrow             = QStyle__StateFlag(0x00004000)
	QStyle__State_Selected            = QStyle__StateFlag(0x00008000)
	QStyle__State_Active              = QStyle__StateFlag(0x00010000)
	QStyle__State_Window              = QStyle__StateFlag(0x00020000)
	QStyle__State_Open                = QStyle__StateFlag(0x00040000)
	QStyle__State_Children            = QStyle__StateFlag(0x00080000)
	QStyle__State_Item                = QStyle__StateFlag(0x00100000)
	QStyle__State_Sibling             = QStyle__StateFlag(0x00200000)
	QStyle__State_Editing             = QStyle__StateFlag(0x00400000)
	QStyle__State_KeyboardFocusChange = QStyle__StateFlag(0x00800000)
	QStyle__State_HasEditFocus        = QStyle__StateFlag(0x01000000)
	QStyle__State_ReadOnly            = QStyle__StateFlag(0x02000000)
	QStyle__State_Small               = QStyle__StateFlag(0x04000000)
	QStyle__State_Mini                = QStyle__StateFlag(0x08000000)
)

//QStyle::StyleHint
type QStyle__StyleHint int64

var (
	QStyle__SH_EtchDisabledText                               = QStyle__StyleHint(0)
	QStyle__SH_DitherDisabledText                             = QStyle__StyleHint(1)
	QStyle__SH_ScrollBar_MiddleClickAbsolutePosition          = QStyle__StyleHint(2)
	QStyle__SH_ScrollBar_ScrollWhenPointerLeavesControl       = QStyle__StyleHint(3)
	QStyle__SH_TabBar_SelectMouseType                         = QStyle__StyleHint(4)
	QStyle__SH_TabBar_Alignment                               = QStyle__StyleHint(5)
	QStyle__SH_Header_ArrowAlignment                          = QStyle__StyleHint(6)
	QStyle__SH_Slider_SnapToValue                             = QStyle__StyleHint(7)
	QStyle__SH_Slider_SloppyKeyEvents                         = QStyle__StyleHint(8)
	QStyle__SH_ProgressDialog_CenterCancelButton              = QStyle__StyleHint(9)
	QStyle__SH_ProgressDialog_TextLabelAlignment              = QStyle__StyleHint(10)
	QStyle__SH_PrintDialog_RightAlignButtons                  = QStyle__StyleHint(11)
	QStyle__SH_MainWindow_SpaceBelowMenuBar                   = QStyle__StyleHint(12)
	QStyle__SH_FontDialog_SelectAssociatedText                = QStyle__StyleHint(13)
	QStyle__SH_Menu_AllowActiveAndDisabled                    = QStyle__StyleHint(14)
	QStyle__SH_Menu_SpaceActivatesItem                        = QStyle__StyleHint(15)
	QStyle__SH_Menu_SubMenuPopupDelay                         = QStyle__StyleHint(16)
	QStyle__SH_ScrollView_FrameOnlyAroundContents             = QStyle__StyleHint(17)
	QStyle__SH_MenuBar_AltKeyNavigation                       = QStyle__StyleHint(18)
	QStyle__SH_ComboBox_ListMouseTracking                     = QStyle__StyleHint(19)
	QStyle__SH_Menu_MouseTracking                             = QStyle__StyleHint(20)
	QStyle__SH_MenuBar_MouseTracking                          = QStyle__StyleHint(21)
	QStyle__SH_ItemView_ChangeHighlightOnFocus                = QStyle__StyleHint(22)
	QStyle__SH_Widget_ShareActivation                         = QStyle__StyleHint(23)
	QStyle__SH_Workspace_FillSpaceOnMaximize                  = QStyle__StyleHint(24)
	QStyle__SH_ComboBox_Popup                                 = QStyle__StyleHint(25)
	QStyle__SH_TitleBar_NoBorder                              = QStyle__StyleHint(26)
	QStyle__SH_Slider_StopMouseOverSlider                     = QStyle__StyleHint(27)
	QStyle__SH_ScrollBar_StopMouseOverSlider                  = QStyle__StyleHint(QStyle__SH_Slider_StopMouseOverSlider)
	QStyle__SH_BlinkCursorWhenTextSelected                    = QStyle__StyleHint(C.QStyle_SH_BlinkCursorWhenTextSelected_Type())
	QStyle__SH_RichText_FullWidthSelection                    = QStyle__StyleHint(C.QStyle_SH_RichText_FullWidthSelection_Type())
	QStyle__SH_Menu_Scrollable                                = QStyle__StyleHint(C.QStyle_SH_Menu_Scrollable_Type())
	QStyle__SH_GroupBox_TextLabelVerticalAlignment            = QStyle__StyleHint(C.QStyle_SH_GroupBox_TextLabelVerticalAlignment_Type())
	QStyle__SH_GroupBox_TextLabelColor                        = QStyle__StyleHint(C.QStyle_SH_GroupBox_TextLabelColor_Type())
	QStyle__SH_Menu_SloppySubMenus                            = QStyle__StyleHint(C.QStyle_SH_Menu_SloppySubMenus_Type())
	QStyle__SH_Table_GridLineColor                            = QStyle__StyleHint(C.QStyle_SH_Table_GridLineColor_Type())
	QStyle__SH_LineEdit_PasswordCharacter                     = QStyle__StyleHint(C.QStyle_SH_LineEdit_PasswordCharacter_Type())
	QStyle__SH_DialogButtons_DefaultButton                    = QStyle__StyleHint(C.QStyle_SH_DialogButtons_DefaultButton_Type())
	QStyle__SH_ToolBox_SelectedPageTitleBold                  = QStyle__StyleHint(C.QStyle_SH_ToolBox_SelectedPageTitleBold_Type())
	QStyle__SH_TabBar_PreferNoArrows                          = QStyle__StyleHint(C.QStyle_SH_TabBar_PreferNoArrows_Type())
	QStyle__SH_ScrollBar_LeftClickAbsolutePosition            = QStyle__StyleHint(C.QStyle_SH_ScrollBar_LeftClickAbsolutePosition_Type())
	QStyle__SH_ListViewExpand_SelectMouseType                 = QStyle__StyleHint(C.QStyle_SH_ListViewExpand_SelectMouseType_Type())
	QStyle__SH_UnderlineShortcut                              = QStyle__StyleHint(C.QStyle_SH_UnderlineShortcut_Type())
	QStyle__SH_SpinBox_AnimateButton                          = QStyle__StyleHint(C.QStyle_SH_SpinBox_AnimateButton_Type())
	QStyle__SH_SpinBox_KeyPressAutoRepeatRate                 = QStyle__StyleHint(C.QStyle_SH_SpinBox_KeyPressAutoRepeatRate_Type())
	QStyle__SH_SpinBox_ClickAutoRepeatRate                    = QStyle__StyleHint(C.QStyle_SH_SpinBox_ClickAutoRepeatRate_Type())
	QStyle__SH_Menu_FillScreenWithScroll                      = QStyle__StyleHint(C.QStyle_SH_Menu_FillScreenWithScroll_Type())
	QStyle__SH_ToolTipLabel_Opacity                           = QStyle__StyleHint(C.QStyle_SH_ToolTipLabel_Opacity_Type())
	QStyle__SH_DrawMenuBarSeparator                           = QStyle__StyleHint(C.QStyle_SH_DrawMenuBarSeparator_Type())
	QStyle__SH_TitleBar_ModifyNotification                    = QStyle__StyleHint(C.QStyle_SH_TitleBar_ModifyNotification_Type())
	QStyle__SH_Button_FocusPolicy                             = QStyle__StyleHint(C.QStyle_SH_Button_FocusPolicy_Type())
	QStyle__SH_MessageBox_UseBorderForButtonSpacing           = QStyle__StyleHint(C.QStyle_SH_MessageBox_UseBorderForButtonSpacing_Type())
	QStyle__SH_TitleBar_AutoRaise                             = QStyle__StyleHint(C.QStyle_SH_TitleBar_AutoRaise_Type())
	QStyle__SH_ToolButton_PopupDelay                          = QStyle__StyleHint(C.QStyle_SH_ToolButton_PopupDelay_Type())
	QStyle__SH_FocusFrame_Mask                                = QStyle__StyleHint(C.QStyle_SH_FocusFrame_Mask_Type())
	QStyle__SH_RubberBand_Mask                                = QStyle__StyleHint(C.QStyle_SH_RubberBand_Mask_Type())
	QStyle__SH_WindowFrame_Mask                               = QStyle__StyleHint(C.QStyle_SH_WindowFrame_Mask_Type())
	QStyle__SH_SpinControls_DisableOnBounds                   = QStyle__StyleHint(C.QStyle_SH_SpinControls_DisableOnBounds_Type())
	QStyle__SH_Dial_BackgroundRole                            = QStyle__StyleHint(C.QStyle_SH_Dial_BackgroundRole_Type())
	QStyle__SH_ComboBox_LayoutDirection                       = QStyle__StyleHint(C.QStyle_SH_ComboBox_LayoutDirection_Type())
	QStyle__SH_ItemView_EllipsisLocation                      = QStyle__StyleHint(C.QStyle_SH_ItemView_EllipsisLocation_Type())
	QStyle__SH_ItemView_ShowDecorationSelected                = QStyle__StyleHint(C.QStyle_SH_ItemView_ShowDecorationSelected_Type())
	QStyle__SH_ItemView_ActivateItemOnSingleClick             = QStyle__StyleHint(C.QStyle_SH_ItemView_ActivateItemOnSingleClick_Type())
	QStyle__SH_ScrollBar_ContextMenu                          = QStyle__StyleHint(C.QStyle_SH_ScrollBar_ContextMenu_Type())
	QStyle__SH_ScrollBar_RollBetweenButtons                   = QStyle__StyleHint(C.QStyle_SH_ScrollBar_RollBetweenButtons_Type())
	QStyle__SH_Slider_AbsoluteSetButtons                      = QStyle__StyleHint(C.QStyle_SH_Slider_AbsoluteSetButtons_Type())
	QStyle__SH_Slider_PageSetButtons                          = QStyle__StyleHint(C.QStyle_SH_Slider_PageSetButtons_Type())
	QStyle__SH_Menu_KeyboardSearch                            = QStyle__StyleHint(C.QStyle_SH_Menu_KeyboardSearch_Type())
	QStyle__SH_TabBar_ElideMode                               = QStyle__StyleHint(C.QStyle_SH_TabBar_ElideMode_Type())
	QStyle__SH_DialogButtonLayout                             = QStyle__StyleHint(C.QStyle_SH_DialogButtonLayout_Type())
	QStyle__SH_ComboBox_PopupFrameStyle                       = QStyle__StyleHint(C.QStyle_SH_ComboBox_PopupFrameStyle_Type())
	QStyle__SH_MessageBox_TextInteractionFlags                = QStyle__StyleHint(C.QStyle_SH_MessageBox_TextInteractionFlags_Type())
	QStyle__SH_DialogButtonBox_ButtonsHaveIcons               = QStyle__StyleHint(C.QStyle_SH_DialogButtonBox_ButtonsHaveIcons_Type())
	QStyle__SH_SpellCheckUnderlineStyle                       = QStyle__StyleHint(C.QStyle_SH_SpellCheckUnderlineStyle_Type())
	QStyle__SH_MessageBox_CenterButtons                       = QStyle__StyleHint(C.QStyle_SH_MessageBox_CenterButtons_Type())
	QStyle__SH_Menu_SelectionWrap                             = QStyle__StyleHint(C.QStyle_SH_Menu_SelectionWrap_Type())
	QStyle__SH_ItemView_MovementWithoutUpdatingSelection      = QStyle__StyleHint(C.QStyle_SH_ItemView_MovementWithoutUpdatingSelection_Type())
	QStyle__SH_ToolTip_Mask                                   = QStyle__StyleHint(C.QStyle_SH_ToolTip_Mask_Type())
	QStyle__SH_FocusFrame_AboveWidget                         = QStyle__StyleHint(C.QStyle_SH_FocusFrame_AboveWidget_Type())
	QStyle__SH_TextControl_FocusIndicatorTextCharFormat       = QStyle__StyleHint(C.QStyle_SH_TextControl_FocusIndicatorTextCharFormat_Type())
	QStyle__SH_WizardStyle                                    = QStyle__StyleHint(C.QStyle_SH_WizardStyle_Type())
	QStyle__SH_ItemView_ArrowKeysNavigateIntoChildren         = QStyle__StyleHint(C.QStyle_SH_ItemView_ArrowKeysNavigateIntoChildren_Type())
	QStyle__SH_Menu_Mask                                      = QStyle__StyleHint(C.QStyle_SH_Menu_Mask_Type())
	QStyle__SH_Menu_FlashTriggeredItem                        = QStyle__StyleHint(C.QStyle_SH_Menu_FlashTriggeredItem_Type())
	QStyle__SH_Menu_FadeOutOnHide                             = QStyle__StyleHint(C.QStyle_SH_Menu_FadeOutOnHide_Type())
	QStyle__SH_SpinBox_ClickAutoRepeatThreshold               = QStyle__StyleHint(C.QStyle_SH_SpinBox_ClickAutoRepeatThreshold_Type())
	QStyle__SH_ItemView_PaintAlternatingRowColorsForEmptyArea = QStyle__StyleHint(C.QStyle_SH_ItemView_PaintAlternatingRowColorsForEmptyArea_Type())
	QStyle__SH_FormLayoutWrapPolicy                           = QStyle__StyleHint(C.QStyle_SH_FormLayoutWrapPolicy_Type())
	QStyle__SH_TabWidget_DefaultTabPosition                   = QStyle__StyleHint(C.QStyle_SH_TabWidget_DefaultTabPosition_Type())
	QStyle__SH_ToolBar_Movable                                = QStyle__StyleHint(C.QStyle_SH_ToolBar_Movable_Type())
	QStyle__SH_FormLayoutFieldGrowthPolicy                    = QStyle__StyleHint(C.QStyle_SH_FormLayoutFieldGrowthPolicy_Type())
	QStyle__SH_FormLayoutFormAlignment                        = QStyle__StyleHint(C.QStyle_SH_FormLayoutFormAlignment_Type())
	QStyle__SH_FormLayoutLabelAlignment                       = QStyle__StyleHint(C.QStyle_SH_FormLayoutLabelAlignment_Type())
	QStyle__SH_ItemView_DrawDelegateFrame                     = QStyle__StyleHint(C.QStyle_SH_ItemView_DrawDelegateFrame_Type())
	QStyle__SH_TabBar_CloseButtonPosition                     = QStyle__StyleHint(C.QStyle_SH_TabBar_CloseButtonPosition_Type())
	QStyle__SH_DockWidget_ButtonsHaveFrame                    = QStyle__StyleHint(C.QStyle_SH_DockWidget_ButtonsHaveFrame_Type())
	QStyle__SH_ToolButtonStyle                                = QStyle__StyleHint(C.QStyle_SH_ToolButtonStyle_Type())
	QStyle__SH_RequestSoftwareInputPanel                      = QStyle__StyleHint(C.QStyle_SH_RequestSoftwareInputPanel_Type())
	QStyle__SH_ScrollBar_Transient                            = QStyle__StyleHint(C.QStyle_SH_ScrollBar_Transient_Type())
	QStyle__SH_Menu_SupportsSections                          = QStyle__StyleHint(C.QStyle_SH_Menu_SupportsSections_Type())
	QStyle__SH_ToolTip_WakeUpDelay                            = QStyle__StyleHint(C.QStyle_SH_ToolTip_WakeUpDelay_Type())
	QStyle__SH_ToolTip_FallAsleepDelay                        = QStyle__StyleHint(C.QStyle_SH_ToolTip_FallAsleepDelay_Type())
	QStyle__SH_Widget_Animate                                 = QStyle__StyleHint(C.QStyle_SH_Widget_Animate_Type())
	QStyle__SH_Splitter_OpaqueResize                          = QStyle__StyleHint(C.QStyle_SH_Splitter_OpaqueResize_Type())
	QStyle__SH_ComboBox_UseNativePopup                        = QStyle__StyleHint(C.QStyle_SH_ComboBox_UseNativePopup_Type())
	QStyle__SH_LineEdit_PasswordMaskDelay                     = QStyle__StyleHint(C.QStyle_SH_LineEdit_PasswordMaskDelay_Type())
	QStyle__SH_TabBar_ChangeCurrentDelay                      = QStyle__StyleHint(C.QStyle_SH_TabBar_ChangeCurrentDelay_Type())
	QStyle__SH_Menu_SubMenuUniDirection                       = QStyle__StyleHint(C.QStyle_SH_Menu_SubMenuUniDirection_Type())
	QStyle__SH_Menu_SubMenuUniDirectionFailCount              = QStyle__StyleHint(C.QStyle_SH_Menu_SubMenuUniDirectionFailCount_Type())
	QStyle__SH_Menu_SubMenuSloppySelectOtherActions           = QStyle__StyleHint(C.QStyle_SH_Menu_SubMenuSloppySelectOtherActions_Type())
	QStyle__SH_Menu_SubMenuSloppyCloseTimeout                 = QStyle__StyleHint(C.QStyle_SH_Menu_SubMenuSloppyCloseTimeout_Type())
	QStyle__SH_Menu_SubMenuResetWhenReenteringParent          = QStyle__StyleHint(C.QStyle_SH_Menu_SubMenuResetWhenReenteringParent_Type())
	QStyle__SH_Menu_SubMenuDontStartSloppyOnLeave             = QStyle__StyleHint(C.QStyle_SH_Menu_SubMenuDontStartSloppyOnLeave_Type())
	QStyle__SH_CustomBase                                     = QStyle__StyleHint(0xf0000000)
)

//QStyle::SubControl
type QStyle__SubControl int64

const (
	QStyle__SC_None                      = QStyle__SubControl(0x00000000)
	QStyle__SC_ScrollBarAddLine          = QStyle__SubControl(0x00000001)
	QStyle__SC_ScrollBarSubLine          = QStyle__SubControl(0x00000002)
	QStyle__SC_ScrollBarAddPage          = QStyle__SubControl(0x00000004)
	QStyle__SC_ScrollBarSubPage          = QStyle__SubControl(0x00000008)
	QStyle__SC_ScrollBarFirst            = QStyle__SubControl(0x00000010)
	QStyle__SC_ScrollBarLast             = QStyle__SubControl(0x00000020)
	QStyle__SC_ScrollBarSlider           = QStyle__SubControl(0x00000040)
	QStyle__SC_ScrollBarGroove           = QStyle__SubControl(0x00000080)
	QStyle__SC_SpinBoxUp                 = QStyle__SubControl(0x00000001)
	QStyle__SC_SpinBoxDown               = QStyle__SubControl(0x00000002)
	QStyle__SC_SpinBoxFrame              = QStyle__SubControl(0x00000004)
	QStyle__SC_SpinBoxEditField          = QStyle__SubControl(0x00000008)
	QStyle__SC_ComboBoxFrame             = QStyle__SubControl(0x00000001)
	QStyle__SC_ComboBoxEditField         = QStyle__SubControl(0x00000002)
	QStyle__SC_ComboBoxArrow             = QStyle__SubControl(0x00000004)
	QStyle__SC_ComboBoxListBoxPopup      = QStyle__SubControl(0x00000008)
	QStyle__SC_SliderGroove              = QStyle__SubControl(0x00000001)
	QStyle__SC_SliderHandle              = QStyle__SubControl(0x00000002)
	QStyle__SC_SliderTickmarks           = QStyle__SubControl(0x00000004)
	QStyle__SC_ToolButton                = QStyle__SubControl(0x00000001)
	QStyle__SC_ToolButtonMenu            = QStyle__SubControl(0x00000002)
	QStyle__SC_TitleBarSysMenu           = QStyle__SubControl(0x00000001)
	QStyle__SC_TitleBarMinButton         = QStyle__SubControl(0x00000002)
	QStyle__SC_TitleBarMaxButton         = QStyle__SubControl(0x00000004)
	QStyle__SC_TitleBarCloseButton       = QStyle__SubControl(0x00000008)
	QStyle__SC_TitleBarNormalButton      = QStyle__SubControl(0x00000010)
	QStyle__SC_TitleBarShadeButton       = QStyle__SubControl(0x00000020)
	QStyle__SC_TitleBarUnshadeButton     = QStyle__SubControl(0x00000040)
	QStyle__SC_TitleBarContextHelpButton = QStyle__SubControl(0x00000080)
	QStyle__SC_TitleBarLabel             = QStyle__SubControl(0x00000100)
	QStyle__SC_DialGroove                = QStyle__SubControl(0x00000001)
	QStyle__SC_DialHandle                = QStyle__SubControl(0x00000002)
	QStyle__SC_DialTickmarks             = QStyle__SubControl(0x00000004)
	QStyle__SC_GroupBoxCheckBox          = QStyle__SubControl(0x00000001)
	QStyle__SC_GroupBoxLabel             = QStyle__SubControl(0x00000002)
	QStyle__SC_GroupBoxContents          = QStyle__SubControl(0x00000004)
	QStyle__SC_GroupBoxFrame             = QStyle__SubControl(0x00000008)
	QStyle__SC_MdiMinButton              = QStyle__SubControl(0x00000001)
	QStyle__SC_MdiNormalButton           = QStyle__SubControl(0x00000002)
	QStyle__SC_MdiCloseButton            = QStyle__SubControl(0x00000004)
	QStyle__SC_CustomBase                = QStyle__SubControl(0xf0000000)
	QStyle__SC_All                       = QStyle__SubControl(0xffffffff)
)

//QStyle::SubElement
type QStyle__SubElement int64

var (
	QStyle__SE_PushButtonContents         = QStyle__SubElement(0)
	QStyle__SE_PushButtonFocusRect        = QStyle__SubElement(1)
	QStyle__SE_CheckBoxIndicator          = QStyle__SubElement(2)
	QStyle__SE_CheckBoxContents           = QStyle__SubElement(3)
	QStyle__SE_CheckBoxFocusRect          = QStyle__SubElement(4)
	QStyle__SE_CheckBoxClickRect          = QStyle__SubElement(5)
	QStyle__SE_RadioButtonIndicator       = QStyle__SubElement(6)
	QStyle__SE_RadioButtonContents        = QStyle__SubElement(7)
	QStyle__SE_RadioButtonFocusRect       = QStyle__SubElement(8)
	QStyle__SE_RadioButtonClickRect       = QStyle__SubElement(9)
	QStyle__SE_ComboBoxFocusRect          = QStyle__SubElement(10)
	QStyle__SE_SliderFocusRect            = QStyle__SubElement(11)
	QStyle__SE_ProgressBarGroove          = QStyle__SubElement(12)
	QStyle__SE_ProgressBarContents        = QStyle__SubElement(13)
	QStyle__SE_ProgressBarLabel           = QStyle__SubElement(14)
	QStyle__SE_ToolBoxTabContents         = QStyle__SubElement(15)
	QStyle__SE_HeaderLabel                = QStyle__SubElement(16)
	QStyle__SE_HeaderArrow                = QStyle__SubElement(17)
	QStyle__SE_TabWidgetTabBar            = QStyle__SubElement(18)
	QStyle__SE_TabWidgetTabPane           = QStyle__SubElement(19)
	QStyle__SE_TabWidgetTabContents       = QStyle__SubElement(20)
	QStyle__SE_TabWidgetLeftCorner        = QStyle__SubElement(21)
	QStyle__SE_TabWidgetRightCorner       = QStyle__SubElement(22)
	QStyle__SE_ViewItemCheckIndicator     = QStyle__SubElement(23)
	QStyle__SE_ItemViewItemCheckIndicator = QStyle__SubElement(QStyle__SE_ViewItemCheckIndicator)
	QStyle__SE_TabBarTearIndicator        = QStyle__SubElement(C.QStyle_SE_TabBarTearIndicator_Type())
	QStyle__SE_TreeViewDisclosureItem     = QStyle__SubElement(C.QStyle_SE_TreeViewDisclosureItem_Type())
	QStyle__SE_LineEditContents           = QStyle__SubElement(C.QStyle_SE_LineEditContents_Type())
	QStyle__SE_FrameContents              = QStyle__SubElement(C.QStyle_SE_FrameContents_Type())
	QStyle__SE_DockWidgetCloseButton      = QStyle__SubElement(C.QStyle_SE_DockWidgetCloseButton_Type())
	QStyle__SE_DockWidgetFloatButton      = QStyle__SubElement(C.QStyle_SE_DockWidgetFloatButton_Type())
	QStyle__SE_DockWidgetTitleBarText     = QStyle__SubElement(C.QStyle_SE_DockWidgetTitleBarText_Type())
	QStyle__SE_DockWidgetIcon             = QStyle__SubElement(C.QStyle_SE_DockWidgetIcon_Type())
	QStyle__SE_CheckBoxLayoutItem         = QStyle__SubElement(C.QStyle_SE_CheckBoxLayoutItem_Type())
	QStyle__SE_ComboBoxLayoutItem         = QStyle__SubElement(C.QStyle_SE_ComboBoxLayoutItem_Type())
	QStyle__SE_DateTimeEditLayoutItem     = QStyle__SubElement(C.QStyle_SE_DateTimeEditLayoutItem_Type())
	QStyle__SE_DialogButtonBoxLayoutItem  = QStyle__SubElement(C.QStyle_SE_DialogButtonBoxLayoutItem_Type())
	QStyle__SE_LabelLayoutItem            = QStyle__SubElement(C.QStyle_SE_LabelLayoutItem_Type())
	QStyle__SE_ProgressBarLayoutItem      = QStyle__SubElement(C.QStyle_SE_ProgressBarLayoutItem_Type())
	QStyle__SE_PushButtonLayoutItem       = QStyle__SubElement(C.QStyle_SE_PushButtonLayoutItem_Type())
	QStyle__SE_RadioButtonLayoutItem      = QStyle__SubElement(C.QStyle_SE_RadioButtonLayoutItem_Type())
	QStyle__SE_SliderLayoutItem           = QStyle__SubElement(C.QStyle_SE_SliderLayoutItem_Type())
	QStyle__SE_SpinBoxLayoutItem          = QStyle__SubElement(C.QStyle_SE_SpinBoxLayoutItem_Type())
	QStyle__SE_ToolButtonLayoutItem       = QStyle__SubElement(C.QStyle_SE_ToolButtonLayoutItem_Type())
	QStyle__SE_FrameLayoutItem            = QStyle__SubElement(C.QStyle_SE_FrameLayoutItem_Type())
	QStyle__SE_GroupBoxLayoutItem         = QStyle__SubElement(C.QStyle_SE_GroupBoxLayoutItem_Type())
	QStyle__SE_TabWidgetLayoutItem        = QStyle__SubElement(C.QStyle_SE_TabWidgetLayoutItem_Type())
	QStyle__SE_ItemViewItemDecoration     = QStyle__SubElement(C.QStyle_SE_ItemViewItemDecoration_Type())
	QStyle__SE_ItemViewItemText           = QStyle__SubElement(C.QStyle_SE_ItemViewItemText_Type())
	QStyle__SE_ItemViewItemFocusRect      = QStyle__SubElement(C.QStyle_SE_ItemViewItemFocusRect_Type())
	QStyle__SE_TabBarTabLeftButton        = QStyle__SubElement(C.QStyle_SE_TabBarTabLeftButton_Type())
	QStyle__SE_TabBarTabRightButton       = QStyle__SubElement(C.QStyle_SE_TabBarTabRightButton_Type())
	QStyle__SE_TabBarTabText              = QStyle__SubElement(C.QStyle_SE_TabBarTabText_Type())
	QStyle__SE_ShapedFrameContents        = QStyle__SubElement(C.QStyle_SE_ShapedFrameContents_Type())
	QStyle__SE_ToolBarHandle              = QStyle__SubElement(C.QStyle_SE_ToolBarHandle_Type())
	QStyle__SE_CustomBase                 = QStyle__SubElement(0xf0000000)
)

func (ptr *QStyle) ItemPixmapRect(rectangle core.QRect_ITF, alignment int, pixmap gui.QPixmap_ITF) *core.QRect {
	defer qt.Recovering("QStyle::itemPixmapRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QStyle_ItemPixmapRect(ptr.Pointer(), core.PointerFromQRect(rectangle), C.int(alignment), gui.PointerFromQPixmap(pixmap)))
	}
	return nil
}

func (ptr *QStyle) ItemTextRect(metrics gui.QFontMetrics_ITF, rectangle core.QRect_ITF, alignment int, enabled bool, text string) *core.QRect {
	defer qt.Recovering("QStyle::itemTextRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QStyle_ItemTextRect(ptr.Pointer(), gui.PointerFromQFontMetrics(metrics), core.PointerFromQRect(rectangle), C.int(alignment), C.int(qt.GoBoolToInt(enabled)), C.CString(text)))
	}
	return nil
}

func (ptr *QStyle) Proxy() *QStyle {
	defer qt.Recovering("QStyle::proxy")

	if ptr.Pointer() != nil {
		return NewQStyleFromPointer(C.QStyle_Proxy(ptr.Pointer()))
	}
	return nil
}

func QStyle_SliderValueFromPosition(min int, max int, position int, span int, upsideDown bool) int {
	defer qt.Recovering("QStyle::sliderValueFromPosition")

	return int(C.QStyle_QStyle_SliderValueFromPosition(C.int(min), C.int(max), C.int(position), C.int(span), C.int(qt.GoBoolToInt(upsideDown))))
}

func QStyle_VisualPos(direction core.Qt__LayoutDirection, boundingRectangle core.QRect_ITF, logicalPosition core.QPoint_ITF) *core.QPoint {
	defer qt.Recovering("QStyle::visualPos")

	return core.NewQPointFromPointer(C.QStyle_QStyle_VisualPos(C.int(direction), core.PointerFromQRect(boundingRectangle), core.PointerFromQPoint(logicalPosition)))
}

func QStyle_VisualRect(direction core.Qt__LayoutDirection, boundingRectangle core.QRect_ITF, logicalRectangle core.QRect_ITF) *core.QRect {
	defer qt.Recovering("QStyle::visualRect")

	return core.NewQRectFromPointer(C.QStyle_QStyle_VisualRect(C.int(direction), core.PointerFromQRect(boundingRectangle), core.PointerFromQRect(logicalRectangle)))
}

func QStyle_AlignedRect(direction core.Qt__LayoutDirection, alignment core.Qt__AlignmentFlag, size core.QSize_ITF, rectangle core.QRect_ITF) *core.QRect {
	defer qt.Recovering("QStyle::alignedRect")

	return core.NewQRectFromPointer(C.QStyle_QStyle_AlignedRect(C.int(direction), C.int(alignment), core.PointerFromQSize(size), core.PointerFromQRect(rectangle)))
}

func (ptr *QStyle) CombinedLayoutSpacing(controls1 QSizePolicy__ControlType, controls2 QSizePolicy__ControlType, orientation core.Qt__Orientation, option QStyleOption_ITF, widget QWidget_ITF) int {
	defer qt.Recovering("QStyle::combinedLayoutSpacing")

	if ptr.Pointer() != nil {
		return int(C.QStyle_CombinedLayoutSpacing(ptr.Pointer(), C.int(controls1), C.int(controls2), C.int(orientation), PointerFromQStyleOption(option), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QStyle) DrawComplexControl(control QStyle__ComplexControl, option QStyleOptionComplex_ITF, painter gui.QPainter_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QStyle::drawComplexControl")

	if ptr.Pointer() != nil {
		C.QStyle_DrawComplexControl(ptr.Pointer(), C.int(control), PointerFromQStyleOptionComplex(option), gui.PointerFromQPainter(painter), PointerFromQWidget(widget))
	}
}

func (ptr *QStyle) DrawControl(element QStyle__ControlElement, option QStyleOption_ITF, painter gui.QPainter_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QStyle::drawControl")

	if ptr.Pointer() != nil {
		C.QStyle_DrawControl(ptr.Pointer(), C.int(element), PointerFromQStyleOption(option), gui.PointerFromQPainter(painter), PointerFromQWidget(widget))
	}
}

func (ptr *QStyle) DrawPrimitive(element QStyle__PrimitiveElement, option QStyleOption_ITF, painter gui.QPainter_ITF, widget QWidget_ITF) {
	defer qt.Recovering("QStyle::drawPrimitive")

	if ptr.Pointer() != nil {
		C.QStyle_DrawPrimitive(ptr.Pointer(), C.int(element), PointerFromQStyleOption(option), gui.PointerFromQPainter(painter), PointerFromQWidget(widget))
	}
}

func (ptr *QStyle) HitTestComplexControl(control QStyle__ComplexControl, option QStyleOptionComplex_ITF, position core.QPoint_ITF, widget QWidget_ITF) QStyle__SubControl {
	defer qt.Recovering("QStyle::hitTestComplexControl")

	if ptr.Pointer() != nil {
		return QStyle__SubControl(C.QStyle_HitTestComplexControl(ptr.Pointer(), C.int(control), PointerFromQStyleOptionComplex(option), core.PointerFromQPoint(position), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QStyle) LayoutSpacing(control1 QSizePolicy__ControlType, control2 QSizePolicy__ControlType, orientation core.Qt__Orientation, option QStyleOption_ITF, widget QWidget_ITF) int {
	defer qt.Recovering("QStyle::layoutSpacing")

	if ptr.Pointer() != nil {
		return int(C.QStyle_LayoutSpacing(ptr.Pointer(), C.int(control1), C.int(control2), C.int(orientation), PointerFromQStyleOption(option), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QStyle) PixelMetric(metric QStyle__PixelMetric, option QStyleOption_ITF, widget QWidget_ITF) int {
	defer qt.Recovering("QStyle::pixelMetric")

	if ptr.Pointer() != nil {
		return int(C.QStyle_PixelMetric(ptr.Pointer(), C.int(metric), PointerFromQStyleOption(option), PointerFromQWidget(widget)))
	}
	return 0
}

func (ptr *QStyle) ConnectPolish(f func(widget *QWidget)) {
	defer qt.Recovering("connect QStyle::polish")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "polish", f)
	}
}

func (ptr *QStyle) DisconnectPolish() {
	defer qt.Recovering("disconnect QStyle::polish")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "polish")
	}
}

//export callbackQStylePolish
func callbackQStylePolish(ptrName *C.char, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QStyle::polish")

	if signal := qt.GetSignal(C.GoString(ptrName), "polish"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func (ptr *QStyle) SizeFromContents(ty QStyle__ContentsType, option QStyleOption_ITF, contentsSize core.QSize_ITF, widget QWidget_ITF) *core.QSize {
	defer qt.Recovering("QStyle::sizeFromContents")

	if ptr.Pointer() != nil {
		return core.NewQSizeFromPointer(C.QStyle_SizeFromContents(ptr.Pointer(), C.int(ty), PointerFromQStyleOption(option), core.PointerFromQSize(contentsSize), PointerFromQWidget(widget)))
	}
	return nil
}

func QStyle_SliderPositionFromValue(min int, max int, logicalValue int, span int, upsideDown bool) int {
	defer qt.Recovering("QStyle::sliderPositionFromValue")

	return int(C.QStyle_QStyle_SliderPositionFromValue(C.int(min), C.int(max), C.int(logicalValue), C.int(span), C.int(qt.GoBoolToInt(upsideDown))))
}

func (ptr *QStyle) StandardIcon(standardIcon QStyle__StandardPixmap, option QStyleOption_ITF, widget QWidget_ITF) *gui.QIcon {
	defer qt.Recovering("QStyle::standardIcon")

	if ptr.Pointer() != nil {
		return gui.NewQIconFromPointer(C.QStyle_StandardIcon(ptr.Pointer(), C.int(standardIcon), PointerFromQStyleOption(option), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QStyle) StyleHint(hint QStyle__StyleHint, option QStyleOption_ITF, widget QWidget_ITF, returnData QStyleHintReturn_ITF) int {
	defer qt.Recovering("QStyle::styleHint")

	if ptr.Pointer() != nil {
		return int(C.QStyle_StyleHint(ptr.Pointer(), C.int(hint), PointerFromQStyleOption(option), PointerFromQWidget(widget), PointerFromQStyleHintReturn(returnData)))
	}
	return 0
}

func (ptr *QStyle) SubControlRect(control QStyle__ComplexControl, option QStyleOptionComplex_ITF, subControl QStyle__SubControl, widget QWidget_ITF) *core.QRect {
	defer qt.Recovering("QStyle::subControlRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QStyle_SubControlRect(ptr.Pointer(), C.int(control), PointerFromQStyleOptionComplex(option), C.int(subControl), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QStyle) SubElementRect(element QStyle__SubElement, option QStyleOption_ITF, widget QWidget_ITF) *core.QRect {
	defer qt.Recovering("QStyle::subElementRect")

	if ptr.Pointer() != nil {
		return core.NewQRectFromPointer(C.QStyle_SubElementRect(ptr.Pointer(), C.int(element), PointerFromQStyleOption(option), PointerFromQWidget(widget)))
	}
	return nil
}

func (ptr *QStyle) ConnectUnpolish(f func(widget *QWidget)) {
	defer qt.Recovering("connect QStyle::unpolish")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "unpolish", f)
	}
}

func (ptr *QStyle) DisconnectUnpolish() {
	defer qt.Recovering("disconnect QStyle::unpolish")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "unpolish")
	}
}

//export callbackQStyleUnpolish
func callbackQStyleUnpolish(ptrName *C.char, widget unsafe.Pointer) bool {
	defer qt.Recovering("callback QStyle::unpolish")

	if signal := qt.GetSignal(C.GoString(ptrName), "unpolish"); signal != nil {
		signal.(func(*QWidget))(NewQWidgetFromPointer(widget))
		return true
	}
	return false

}

func QStyle_VisualAlignment(direction core.Qt__LayoutDirection, alignment core.Qt__AlignmentFlag) core.Qt__AlignmentFlag {
	defer qt.Recovering("QStyle::visualAlignment")

	return core.Qt__AlignmentFlag(C.QStyle_QStyle_VisualAlignment(C.int(direction), C.int(alignment)))
}

func (ptr *QStyle) DestroyQStyle() {
	defer qt.Recovering("QStyle::~QStyle")

	if ptr.Pointer() != nil {
		C.QStyle_DestroyQStyle(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
