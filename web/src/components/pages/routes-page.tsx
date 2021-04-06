import * as React from "react";
import { InterfaceAddresses, InterfaceDetails, Interfaces, Routes, Rules, Tables,GeneralRequest,GeneralResponse } from "../../models/route-models";
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
    Modal,
    PageSection, 
	PageSectionVariants,
	Stack, 
} from '@patternfly/react-core';
import { MinusIcon, PlusCircleIcon, PlusIcon, RedoIcon, TrashIcon } from '@patternfly/react-icons';


interface RoutesPageState {
    isModalOpen : boolean
    canSelectAll : boolean
	columns_route : string[]
	rows_route : string[][]
	columns_rule : string[]
	rows_rule : string[][]
	columns : string[]
	rows : string[][]
    value : string
}

interface RoutesPageProps {
    // Empty
}

type Optional<T> = T | undefined

export class RoutesPage extends React.Component<RoutesPageProps, RoutesPageState> {
    constructor(props: RoutesPageProps) {
        super(props);
        this.state = {
            value : 'Hii',
            isModalOpen: false,
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
            this.handleModalToggle = this.handleModalToggle.bind(this);
            this.handleModalInput = this.handleModalInput.bind(this);
    }
    async routes() {
        let response=await (await fetch("/api/v1/routes")).json()
        let temp:Routes=response
        console.log(temp)
        console.log(response)
    }
    async routesByTableName(tableName:string) {
        let response=await (await fetch(`/api/v1/routes/${tableName}`)).json()
        console.log(response.routes)
    }
    async addRoute(payload:any) {
        let response=await (await fetch("/api/v1/addroute",{
            method:"post",
            headers: {'Content-Type':'application/json'},
            body: JSON.stringify(payload)
        })).json()
        let temp:GeneralResponse=response
        console.log(response)
        console.log(temp)
    }
    async delRoute(payload:any) {
        let response=await (await fetch("/api/v1/delroute",{
            method:"post",
            headers: {'Content-Type':'application/json'},
            body: JSON.stringify(payload)
        })).json()
        let temp:GeneralResponse=response
        console.log(response)
        console.log(temp)
    }
    //==========================================
    async rules() {
        let response=await (await fetch("/api/v1/rules")).json()
        console.log(response)
        let temp:Rules=response
        console.log(temp)
    }
    async addRule(payload:any) {
        let response=await (await fetch("/api/v1/addrule",{
            method:"post",
            headers: {'Content-Type':'application/json'},
            body: JSON.stringify(payload)
        })).json()
        let temp:GeneralResponse=response
        console.log(response)
        console.log(temp)
    }
    async delRule(payload:any) {
        let response=await (await fetch("/api/v1/delrule",{
            method:"post",
            headers: {'Content-Type':'application/json'},
            body: JSON.stringify(payload)
        })).json()
        let temp:GeneralResponse=response
        //console.log(response)
        console.log(temp)
    }
    //============================================
    async interfaceAddresses() {
        let response=await (await fetch("/api/v1/interfaceaddresses")).json()
        let temp:InterfaceAddresses=response
        console.log(response)
        console.log(temp)
    }
    async interfaces() {
        let response=await (await fetch("/api/v1/interfaces")).json()
        let temp:Interfaces=response
        console.log(response)
        console.log(temp)
    }
    async interfaceByName(interfaceName:string) {
        let response=await (await fetch(`/api/v1/interface/${interfaceName}`)).json()
        let temp:InterfaceDetails=response
        console.log(response)
        console.log(temp)
    }
    //================================================
    async tables() {
        let response=await (await fetch("/api/v1/tables")).json()
        let temp:Tables=response;
        console.log(response)
        console.log(temp)
    }



    componentDidMount(){
        document.title = "Routes | SysMon"
        // this.routes()
        // this.routesByTableName("main")
        // For addroute
        const payload1:GeneralRequest=new GeneralRequest()
        payload1.destination="192.168.56.11";
        payload1.intermediate="192.168.122.1";
        payload1.interfaceName="virbr0";
        payload1.tableName="default";
        // this.addRoute(payload1)
        // For delrule
        const payload2=new GeneralRequest()
        payload2.destination="192.168.56.11";
        payload2.intermediate="192.168.122.1";
        payload2.interfaceName="virbr0";
        payload2.tableName="default";
        // this.delRoute(payload2)
        //================================
        // this.rules()
        // For addrule
        let payload3:GeneralRequest=new GeneralRequest();
        payload3.sourceIp="192.168.56.11";
        payload3.tableName="default";
        // this.addRule(payload3)
        // For delrule
        let payload4:GeneralRequest=new GeneralRequest();
        payload4.sourceIp="192.168.56.11";
        payload4.tableName="default";
        // this.delRule(payload4)
        //=================================
        // this.interfaceAddresses()
        // this.interfaces()
        // const interfaceName="wlp2s0"
        // this.interfaceByName(interfaceName)
        //==================================
        // this.tables()

    }
    
    handleModalToggle () {
        this.setState(({ isModalOpen }) => ({
          isModalOpen: !isModalOpen
        }));
      }
    handleModalInput () {
        console.log(this.state.value)
        this.setState (({ value, isModalOpen }) => ({
            value: value,
            isModalOpen: !isModalOpen
        }), () => {console.log(this.state.value," ",this.state.isModalOpen," ",this.state.value)});
      }  
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
        console.log('Rendering routes page ...')
        const { value, isModalOpen, canSelectAll, columns_route, columns_rule, columns, rows_route, rows_rule, rows } = this.state;
	  
		  return ( 	
          <div>    		
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
							<Button variant="link" icon={<PlusIcon />} onClick={this.handleModalToggle}>
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
        <Modal
        title="Simple modal header"
        isOpen={isModalOpen}
        onClose={this.handleModalToggle}
        actions={[
            // <Button key="confirm" variant="primary" onClick={this.handleModalInput}>
            // Confirm
            // </Button>,
            <Button key="cancel" variant="link" onClick={this.handleModalToggle}>
            Cancel
            </Button>
        ]}
        >
        Enter your name: <input
            type="text"
            
            //value={this.state.value}
            defaultValue="Hello!"
            //onSubmit={this.handleModalInput}
         />  
         <Button key="confirm" variant="primary" onClick={this.handleModalInput}>
            Confirm
        </Button>,
        </Modal>
		</div> 
         );
    }
}