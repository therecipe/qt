import QtQuick 2.6
import StratifyLabs.UI 2.0

EPane {
  name: "Rows";
  EHeading {
    title: "Rows";
    inherits: "GridLayout";
    stratifyName: "SRow";
    defaultSize: "block";
  }

  ESectionTitle { text: "Description"; }
  SText {
    style: "block";
    text: 'An SRow arranges items in a row. \
If the items won\'t fit in the row, the items are arranged in a grid. Each row \
can have a total span of <i>STheme.grid_columns</i> (which defaults to a value of ' + STheme.grid_columns + ').';
  }

    SText {
      style: "block";
      text: "If you add the following code to your \
main application window, rows will automatically change to a grid when the screen is more \
narrow than STheme.screen_sm.";
    }

    SWell {
      text: 'onWidthChanged: {
        STheme.updateScreenSize(width);
}';
    }

    EParagraph { text: 'With that code, <i>SRow</i> will change the \
number of columns from <i>STheme.grid_columns</i> (' + STheme.grid_columns + ') to <i>STheme.grid_columns_sm</i> \
(' + STheme.grid_columns_sm + ').'; }

      SHLine{}

      ESectionTitle { text: "Responsive Example"; }
      RowExampleResponsive{}
      ECodeButton { source: "RowExampleResponsive"; }

      EParagraph { text: 'You can customize when the change is triggered \
as well as the number of columns in each mode by adding the code below to the main "ApplicationWindow" \
object.'; }

      ECodeBlockInline {
        text: '\
Component.onCompleted: {
  STheme.grid_columns_sm = 6;
  STheme.screen_sm = 600;
}';
      }

    }
