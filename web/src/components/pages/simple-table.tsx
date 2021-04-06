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
    Modal, 
} from '@patternfly/react-core';
import { MinusIcon, PlusCircleIcon, PlusIcon, RedoIcon, TrashIcon } from '@patternfly/react-icons';

interface TablePageStates {
	canSelectAll : boolean
	columns_route : string[]
	rows_route : string[][]
    isModalOpen:boolean
    name:string
}

interface TablePageProps {
    // Empty
}

type Optional<T> = T | undefined

export class STable extends React.Component<TablePageProps, TablePageStates> {
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
            isModalOpen:false,
            name:"sabuj"
		  };
		}
	  
        handleModalToggle = () => {
            this.setState(({ isModalOpen }) => ({
              isModalOpen: !isModalOpen
            }));
        };
        handleChange(event: React.FormEvent<HTMLInputElement>) {
            const newValue = event.currentTarget.value;
            console.log(newValue)
            this.setState({
                name:"sabuj mondal"
            })
        }
		render() {
		  const { canSelectAll, columns_route,  rows_route,isModalOpen,name } = this.state;
	  
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
						aria-label="Selectable Table"
						cells={columns_route}
						rows={rows_route}>
						    <TableHeader />
						    <TableBody />
						</Table>
					</CardBody>
				</Card>
                <React.Fragment>
        <Button variant="primary" onClick={this.handleModalToggle}>
          Show Modal
        </Button>
        <Modal
          title="Modal Header"
          isOpen={isModalOpen}
          onClose={this.handleModalToggle}
          actions={[
            <Button key="confirm" variant="primary" onClick={this.handleModalToggle}>
              Confirm
            </Button>,
            <Button key="cancel" variant="link" onClick={this.handleModalToggle}>
              Cancel
            </Button>
          ]}
        >
          <input name="dd" value={name} onChange={this.handleChange}></input>
        </Modal>
      </React.Fragment>
			</PageSection>		
		</Stack>
		  );
		}
	  }