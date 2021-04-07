import React, { HtmlHTMLAttributes } from 'react';
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
    Modal, 
} from '@patternfly/react-core';
import { MinusIcon, PlusCircleIcon, PlusIcon, RedoIcon, TrashIcon } from '@patternfly/react-icons';
import { GeneralRequest } from '../../models/route-models';
import { request } from 'node:http';

type Cell={
	cells:string[]
}
type Action={
	title:string
	onClick:(event:any, rowId:any, rowData:any, extra:any)=>void
}

interface TablePageStates {
	columns : string[]
	rows : Cell[]
	actions:Action[]
}

interface TablePageProps {
    // Empty
}

type Optional<T> = T | undefined

export class STable extends React.Component<TablePageProps, TablePageStates> {
    constructor(props : TablePageProps) {
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
				  cells: ['<h1>ggg</h1>', 'two', 'a', 'four', 'five']
				},
				{
				  cells: ['a', 'two', 'k', 'four', 'five']
				},
				{
				  cells: ['p', 'two', 'b', 'four', 'five'],
				}
			  ],
			  actions: [
				{
				  title: 'Remove',
				  onClick:(event, rowId, rowData, extra) => {console.log('clicked on Some action, on row: ', rowId)}
				}
			  ]
		};
	}
		render() {
		  const { columns,  rows,actions } = this.state;
	  
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
						<Table
						actions={actions}
						aria-label="Selectable Table"
						cells={columns}
						rows={rows}>
						    <TableHeader />
						    <TableBody />
						</Table>
					</CardBody>
				</Card>
                <React.Fragment>
                </React.Fragment>
			</PageSection>		
		</Stack>
		  );
		}
	  }