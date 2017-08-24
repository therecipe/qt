import QtQuick 2.6
import StratifyLabs.UI 2.0

SContainer {
  name: "Layouts";
  style: "fill";
  SPane {
    style: "block fill";
    SColumn {
      SText { style: "left text-bold"; text: "Introduction"; }
      EParagraph { text: 'StratifyLabs UI layouts are based \
on a responsive grid system inspired by the Twitter Bootstrap framework including \
containers (SContainer), rows (SRow), and columns (SColumn).'; }
      SText { style: "left text-bold"; text: "Container"; }
      EParagraph { text: 'An SContainer is provides padding \
and sizing. It can be placed directly in a standard QML object (such as \
an "ApplicationWindow") or it can be used inside an SRow or SColumn to provide padding.';
      }

      SText { style: "left text-bold"; text: "Row"; }
      SText { style: "left text-bold"; text: "Column"; }
    }

  }
}
