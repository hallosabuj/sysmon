import React from 'react';
import { 
	Table, 
	TableHeader, 
	TableBody, 
	TableVariant,
	RowSelectVariant, 
} from '@patternfly/react-table';
import { 
	Button,
	Card, 
	CardTitle, 
	CardActions,
	CardBody, 
	CardHeader,
	Checkbox,
	DropdownItem,
    DropdownSeparator,
   PageSection, 
	PageSectionVariants,
	Stack, 
} from '@patternfly/react-core';
import { MinusIcon, PlusCircleIcon, PlusIcon, RedoIcon, TrashIcon } from '@patternfly/react-icons';

interface TablePageStates {
	canSelectAll : boolean
	columns_route : string[]
	rows_route : string[][]
	columns_rule : string[]
	rows_rule : string[][]
	columns : string[]
	rows : string[][]
}

interface TablePageProps {
    // Empty
}

type Optional<T> = T | undefined

export class SimpleTable extends React.Component<TablePageProps, TablePageStates> {
  constructor(props : TablePageProps) {
		  super(props);
		  this.state = {

			canSelectAll: true,
			columns_route : ['Index', 'Route'],
		    rows_route : [
			['Repository one', 'Branch one'],
			['Repository two', 'Branch two'],
			['Repository three', 'Branch three']
		    ],
		    columns_rule : ['Priority', 'Rule'],
		    rows_rule : [
				['Repository one', 'Branch one'],
				['Repository two', 'Branch two'],
				['Repository three', 'Branch three']
			],
		    columns : ['Identifier', 'Name'],
		    rows : [
			['Repository one', 'Branch one'],
			['Repository two', 'Branch two'],
			['Repository three', 'Branch three']
		    ],
		  };

			// this.onSelect = this.onSelect.bind(this);
			// this.toggleSelect = this.toggleSelect.bind(this);
		}
	  
		// handleItemClick(selected: boolean, 
		// 	event: React.MouseEvent<any> | React.KeyboardEvent | MouseEvent) {
		//   this.setState({
		// 	choice: event.currentTarget.id
		//   });
		// }

		// onSelect(event : Event, isSelected : boolean, rowId : number) {
		// 	let rows = this.state.rows_rule.map((oneRow, index) => {
		// 		oneRow.selected = rowId === index;
		// 		return oneRow;
		// 	  });
		// 	  this.setState({
		// 		rows
		// 	  });
		//   }

		// toggleSelect(checked : boolean) {
		// 	this.setState({
		// 	  canSelectAll: checked
		// 	});
		//   }  
	  
		render() {
		  const { canSelectAll, columns_route, columns_rule, columns, rows_route, rows_rule, rows } = this.state;
	  
		  return ( 			
		 <Stack>
			<PageSection variant={PageSectionVariants.light}>
				<Card>
					<CardHeader>
					    <CardTitle>IP Rules Table</CardTitle>
						<CardActions>
							<Button variant="link" icon={<RedoIcon />}>
    						</Button>
							<Button variant="link" icon={<TrashIcon />}>
    						</Button>
						</CardActions>
					</CardHeader>
					<CardBody>
						{/* <Checkbox
						label="Can select all"
						className="pf-u-mb-lg"
						isChecked={canSelectAll}
						onChange={this.toggleSelect}
						aria-label="toggle select all checkbox"
						id="toggle-select-all"
						name="toggle-select-all"
						/> */}
						<Table
						//onSelect={this.onSelect}
						//selectVariant={RowSelectVariant.radio}
						aria-label="Selectable Table"
						cells={columns_rule}
						rows={rows_rule}>
						<TableHeader />
						<TableBody />
						</Table>
					</CardBody>
				</Card>
			</PageSection>	
				
			<PageSection variant={PageSectionVariants.light}>
				<Card>
				    <CardHeader>
					    <CardTitle>IP Routes Table</CardTitle>
						<CardActions>
							<Button variant="link" icon={<PlusIcon />}>
    						</Button>
						    <Button variant="link" icon={<RedoIcon />}>
    						</Button>
							<Button variant="link" icon={<TrashIcon />}>
    						</Button>
						</CardActions>
					</CardHeader>
					<CardBody>
						<Table
						aria-label="Selectable Table"
						cells={columns_route}
						rows={rows_route}>
						<TableHeader />
						<TableBody />
						</Table>
					</CardBody>
				</Card>
			</PageSection>
				
			<PageSection variant={PageSectionVariants.light}>
				<Card>
				    <CardHeader>
					    <CardTitle>IP Table</CardTitle>
						<CardActions>
						    <Button variant="link" icon={<RedoIcon />}>
    						</Button>
							<Button variant="link" icon={<TrashIcon />}>
    						</Button>
						</CardActions>
					</CardHeader>
					<CardBody>
					<Table
						aria-label="Selectable Table"
						cells={columns}
						rows={rows}>
						<TableHeader />
						<TableBody />
						</Table>
					</CardBody>
				</Card>
			</PageSection>
		</Stack>
		  );
		}
	  }