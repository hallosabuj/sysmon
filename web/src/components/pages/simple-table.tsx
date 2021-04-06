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
import { GeneralRequest } from '../../models/route-models';
import { request } from 'node:http';


interface TablePageStates {
	canSelectAll : boolean
	columns_route : string[]
	rows_route : string[][]
    isModalOpen:boolean
    request:GeneralRequest
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
            request:new GeneralRequest
		};
        this.handleClick = this.handleClick.bind(this);
        this.handleChange = this.handleChange.bind(this);
	}
	  
        handleModalToggle = () => {
            this.setState(({ isModalOpen }) => ({
              isModalOpen: !isModalOpen
            }));
        };
        handleChange(event: React.FormEvent<HTMLInputElement>) {
            const newValue = event.currentTarget.value;
            console.log(newValue)
            let temp=new GeneralRequest()
            temp.destination=newValue
            this.setState({
                request:temp
            })
        }
        handleClick(event: React.MouseEvent<HTMLButtonElement,MouseEvent>) {
            console.log("==========")
            console.log(this.state.request)
            event.preventDefault()
        }
		render() {
		  const { canSelectAll, columns_route,  rows_route,isModalOpen,request } = this.state;
	  
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
                        <form id="form">
                            <input name="dd" id="dd" value={request.destination} onChange={this.handleChange}></input>
                            <button onClick={this.handleClick}>Submit</button>
                        </form>
                    </Modal>
                </React.Fragment>
			</PageSection>		
		</Stack>
		  );
		}
	  }