import React from 'react';
import {
  Table,
  TableHeader,
  TableBody,
  sortable,
  SortByDirection,
  headerCol,
  TableVariant,
  expandable,
  cellWidth,
  IRowData
} from '@patternfly/react-table';
import {
    Checkbox
} from '@patternfly/react-core';
import { type } from 'node:os';

type Row = {
    row : string[],
    selected : boolean,
}
interface MetricsPageState {
    // Empty
    canSelectAll : boolean
    columns : string[]
    rows : Row[]
}

interface MetricsPageProps {
    // Empty
}

export class DashboardPage extends React.Component<MetricsPageProps, MetricsPageState> {
  constructor(props: MetricsPageProps) {
    super(props);
    this.state = {
      columns: [
        'Repositories',
        'Branches',
        'Pull requests',
        'Workspaces',
        'Last Commit'
      ],
      rows: [
        {
          row:  ['one', 'two', 'a', 'four', 'five'] ,
          selected: false
        },
        {
          row:  ['a', 'two', 'k', 'four', 'five'] ,
          selected: false
        },
        {
          row: ['p', 'two', 'b', 'four', 'five'] ,
          selected: false
        }
      ],
      canSelectAll: true
    };
    this.onSelect = this.onSelect.bind(this);
    this.toggleSelect = this.toggleSelect.bind(this);
  }

  onSelect(event: React.FormEvent<HTMLInputElement>, isSelected: boolean, rowId: number, oneRow: IRowData ) {
    let rows;
    // if (rowId === -1) {
    //   rows = this.state.rows.map(oneRow => {
    //     oneRow.selected = isSelected;
    //     return oneRow;
    //   });
    // } else {
    //   rows = [...this.state.rows];
    //   rows[rowId].selected = isSelected;
    // }
    // this.setState({
    //   rows
    // });
  }

  toggleSelect(checked: boolean) {
    this.setState({
      canSelectAll: checked
    });
  }

  render() {
    const { columns, rows } = this.state;

    return (
      <div>
      {/* <Table aria-label="Selectable Table" onSelect={this.onSelect} cells={columns} rows={rows} canSelectAll={this.state.canSelectAll}>
        <TableHeader />
        <TableBody />
      </Table>
      <Checkbox
        label="canSelectAll"
        isChecked={this.state.canSelectAll}
        onChange={this.toggleSelect}
        aria-label="toggle select all checkbox"
        id="toggle-select-all"
        name="toggle-select-all" */}
      {/* /> */}
      Dashboard Page
      </div>
    );
  }
}
