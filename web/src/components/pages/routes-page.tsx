import * as React from "react";
import { InterfaceAddresses, InterfaceDetails, Interfaces,ITables, Routes,Route,IRules, Rules, Tables,GeneralRequest,GeneralResponse, Rule } from "../../models/route-models";
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
    Form,
    FormGroup,
    ActionGroup,
    TextInput, 
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
    srcIP : string
    dstIP : string
    defaultGateway : string
    interfaceName : string
}

interface RoutesPageProps {
    // Empty
}

type Optional<T> = T | undefined

export class RoutesPage extends React.Component<RoutesPageProps, RoutesPageState> {
    constructor(props: RoutesPageProps) {
        super(props);
        this.state = {
            srcIP: '',
            dstIP: '',
            defaultGateway: '',
            interfaceName: '',
            isModalOpen: false,
			canSelectAll: true,
			columns_route : ['Index', 'Route','Table Name'],
		    rows_route : [],
		    columns_rule : ['Priority', 'Rule'],
		    rows_rule : [],
		    columns : ['Identifier', 'Name'],
		    rows : [],
		  };

		    //this.onSelect = this.onSelect.bind(this);
		    this.toggleSelect = this.toggleSelect.bind(this);
            this.handleModalToggle = this.handleModalToggle.bind(this);
            this.handleChange = this.handleChange.bind(this);
            this.handleSubmit = this.handleSubmit.bind(this);
    }

    async routes() :Promise<Routes> {
        let response=await (await fetch("/api/v1/routes")).json()
        let temp:Routes=response
        return temp
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
    async rules():Promise<IRules> {
        let response=await (await fetch("/api/v1/rules")).text()
        let temp:IRules=JSON.parse(response)
        return temp
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
        // console.log(response)
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
    async tables():Promise<ITables> {
        let response=await (await fetch("/api/v1/tables")).json()
        let temp:ITables=response;
        return temp
    }
    async addTable(payload:any){
        let response=await (await fetch("/api/v1/addtable",{
            method:"post",
            headers: {'Content-Type':'application/json'},
            body: JSON.stringify(payload)
        })).json()
        let temp:GeneralResponse=response
        console.log(response)
        console.log(temp)
    }
    //==============================================

    async updateRoute(){
        let response=await this.routes()
        if (response.routes){
            let rows_route:string[][]
            rows_route=[]
            for(let i=0;i<response.routes?.length;i++){
                rows_route=[...rows_route,[response.routes[i].index,response.routes[i].route,response.routes[i].tableName]]
            }
            this.setState({
                rows_route:rows_route
            })
        }
    }
    async updateRules(){
        let response=await this.rules()
        if (response.rules){
            let rows_rule:string[][]
            rows_rule=[]
            for(let i=0;i<response.rules?.length;i++){
                rows_rule=[...rows_rule,[response.rules[i].priority,response.rules[i].rule]]
            }
            this.setState({
                rows_rule:rows_rule
            })
        }
    }
    async updateTables(){
        let response=await this.tables()
        if (response.tables){
            let tables:string[][]
            tables=[]
            for(let i=0;i<response.tables?.length;i++){
                tables=[...tables,[response.tables[i].tableNumber,response.tables[i].tableName]]
            }
            this.setState({
                rows:tables
            })
        }
    }

    componentDidMount(){
        document.title = "Routes | SysMon"
        this.updateRoute()
        this.updateRules()
        this.updateTables()
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
        let payload5:GeneralRequest=new GeneralRequest();
        payload5.destination="192.168.122.13";
        payload5.intermediate="192.168.122.1";
        payload5.interfaceName="virbr0";
        payload5.sourceIp="192.168.122.13";
        this.addTable(payload5)
    }
    
    handleModalToggle () {
        this.setState(({ isModalOpen }) => ({
          isModalOpen: !isModalOpen
        }));
      }
    // handleModalInput () {
    //     console.log(this.state.value)
    //     // let temp=document.getElementById("name").

    //     this.setState(prev => ({
    //         value: "dd"
    //       }))
    //       console.log(this.state.value)

    //     // this.setState (({ value, isModalOpen }) => ({
    //     //     value: value,
    //     //     isModalOpen: !isModalOpen
    //     // }), () => {console.log(this.state.value," ",this.state.isModalOpen," ",this.state.value)});
    //   }  
      
    // onSelect(event : Event, isSelected : boolean, rowId : number) {
		// 	let rows = this.state.rows_rule.map((oneRow, index) => {
		// 		oneRow.selected = rowId === index;
		// 		return oneRow;
		// 	  });
		// 	  this.setState({
		// 		rows
		// 	  });
		//   }

		toggleSelect(checked : boolean) {
			this.setState({
			  canSelectAll: checked
			});
		  } 
        handleChange(value: string, event: React.FormEvent<HTMLInputElement>) {
            const id = event.currentTarget.id;
            const newValue = event.currentTarget.value;
            if (id == "simple-form-src-ip-01") {
                console.log("interface ",newValue)
                this.setState({
                  srcIP :newValue
                });
            }
            if (id == "simple-form-dst-ip-01") {
                console.log("interface ",newValue)
                this.setState({
                  dstIP :newValue
                });
            }
            if (id == "simple-form-gateway-01") {
                console.log("interface ",newValue)
                this.setState({
                  defaultGateway :newValue
                });
            }
            if (id == "simple-form-interface-01") {
                console.log("interface ",newValue)
                this.setState({
                  interfaceName :newValue
                });
            }
           }         
         handleSubmit(event: React.FormEvent<HTMLButtonElement>) {
            var pkt = {
                srcIP : this.state.srcIP,
                dstIP : this.state.dstIP,
                defaultGateway : this.state.defaultGateway,
                interfaceName : this.state.interfaceName
            };
           
            
            const objJSON = JSON.stringify(pkt)
            console.log(objJSON);
            //alert('A name was submitted: ' + this.state.value1);
            event.preventDefault();
          }
    render() {
        console.log('Rendering routes page ...')
        const { srcIP, dstIP, defaultGateway, interfaceName, isModalOpen, canSelectAll, columns_route, columns_rule, columns, rows_route, rows_rule, rows } = this.state;
	  
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
        >
         <Form>
         <FormGroup label="Source IP Address"
          isRequired
          fieldId="simple-form-src-ip-01"
        > 
        <TextInput
            isRequired
            type="text"
            id="simple-form-src-ip-01"
            name="simple-form-src-ip-01"
            aria-describedby="simple-form-src-ip-01-helper"
            value={this.state.srcIP}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup label="Destination IP Address" isRequired fieldId="simple-form-dst-ip-01">
          <TextInput
            isRequired
            type="text"
            id="simple-form-dst-ip-01"
            name="simple-form-dst-ip-01"
            value={this.state.dstIP}
            onChange={this.handleChange}
          />
        </FormGroup> 
        <FormGroup label="Gateway" isRequired fieldId="simple-form-gateway-01">
          <TextInput
            isRequired
            type="tel"
            id="simple-form-gateway-01"
            name="simple-form-gateway-01"
            value={this.state.defaultGateway}
            onChange={this.handleChange}
          />
        </FormGroup>
        <FormGroup label="Interface Name" isRequired fieldId="simple-form-interface-01">
          <TextInput
            isRequired
            type="tel"
            id="simple-form-interface-01"
            name="simple-form-interface-01"
            value={this.state.interfaceName}
           onChange={this.handleChange}
          />
        </FormGroup>
        <ActionGroup>
          <Button key="confirm" variant="primary" onClick={this.handleSubmit}>
            Confirm</Button>,
          <Button key="cancel" variant="link" onClick={this.handleModalToggle}>
            Cancel</Button>
        </ActionGroup>
        </Form>
        </Modal>
		</div> 
         );
    }
}