import QtQuick 2.6
import StratifyLabs.UI 2.0

SColumn {
  property alias tableData: libraryModel;
  ListModel {
    id: libraryModel;
    ListElement {
      title: "A Masterpiece"
      author: "Gabriel"
    }
    ListElement {
      title: "Brilliance"
      author: "Jens"
    }
    ListElement {
      title: "Outstanding"
      author: "Frederik"
    }
    ListElement {
      title: "A Masterpiece"
      author: "Gabriel"
    }
  }

  SText {
    span: 4;
    style: "left top text-h3";
    text: "Striped";
  }

  SRow {
    SContainer {
      span: 8;
      height: 450;
      style: "padding-zero right";
      /* The following table will fill the container
          rather than simply taking up its
          implicitHeight. */
      STable{
        model: tableData;
        style: "fill table-striped top";

        STableColumn {
          span: 6;
          role: "title";
          title: "Title";
        }

        STableColumn {
          span: 6;
          role: "author"
          title: "Author"
        }
      }
    }
  }

  SText {
    span: 4;
    style: "left top text-h3";
    text: "Condensed";
  }

  SRow {
    STable{
      span: 8;
      model: tableData;
      style: "right table-condensed top";

      STableColumn {
        span: 6;
        role: "title";
        title: "Title";
      }

      STableColumn {
        span: 6;
        role: "author"
        title: "Author"
      }
    }
  }

  SText {
    span: 4;
    style: "left top text-h3";
    text: "Bordered";
  }

  SRow {
    STable {
      span: 8;
      model: tableData;
      style: "right table-bordered top";

      STableColumn {
        span: 6;
        role: "title";
        title: "Title";
      }

      STableColumn {
        span: 6;
        role: "author"
        title: "Author"
      }
    }
  }

}


