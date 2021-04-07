import * as React from "react";
import { InterfaceAddresses, InterfaceDetails, Interfaces,ITables, Routes,Route,IRules, Rules, Tables,GeneralRequest,GeneralResponse, Rule, IInterfaces } from "../../models/route-models";
import { 
	Table, 
	TableHeader, 
	TableBody, 
	TableVariant,
	RowSelectVariant,
    IRowData,
    IExtraData, 
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
    FormSelect,
    FormSelectOption, 
} from '@patternfly/react-core';
import { MinusIcon, PlusCircleIcon, PlusIcon, RedoIcon, TrashIcon } from '@patternfly/react-icons';
import { c_background_image_BackgroundColor } from "@patternfly/react-tokens";

type Option={
    value:string
    label:string
    disabled:boolean
}
interface RoutesPageState {
    isModalOpen : boolean
    canSelectAll : boolean
	columns_route : string[]
	rows_route : string[][]
	columns_rule : string[]
	rows_rule : string[][]
	columns : string[]
	rows : string[][]
    interfaces_columns:string[]
    interfaces_rows:string[][]
    srcIP : string
    dstIP : string
    defaultGateway : string
    interfaceName : string
    isDisabledConfirm:boolean
    options:Option[]
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
            interfaces_columns:['Index','Interface Name'],
            interfaces_rows:[],
            isDisabledConfirm:true,
            options : []
		  };

		    //this.onSelect = this.onSelect.bind(this);
		    this.toggleSelect = this.toggleSelect.bind(this);
            this.handleModalToggle = this.handleModalToggle.bind(this);
            this.handleChange = this.handleChange.bind(this);
            this.handleSubmit = this.handleSubmit.bind(this);
            this.updateRules = this.updateRules.bind(this);
            this.updateRoute = this.updateRoute.bind(this);
            this.updateTables = this.updateTables.bind(this);
            this.updateInterfaces = this.updateInterfaces.bind(this);
            this.onChange = this.onChange.bind(this);
            this.checkAllValue = this.checkAllValue.bind(this);
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
    async interfaces() :Promise<IInterfaces>{
        let response=await (await fetch("/api/v1/interfaces")).json()
        let temp:IInterfaces=response
        return temp
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
        console.log("===============")
        console.log(temp)
        console.log("===============")
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
    async updateInterfaces(){
        let response=await this.interfaces()
        if (response.interfaces){
            let interfaces:string[][]
            let temp:Option[]=[{value:"null",label:"Choose Interface",disabled:false}]
            interfaces=[]
            for(let i=0;i<response.interfaces?.length;i++){
                temp=[...temp,{value:response.interfaces[i].name,label:response.interfaces[i].name,disabled:false}]
                interfaces=[...interfaces,[response.interfaces[i].index,response.interfaces[i].name]]
            }
            this.setState({
                interfaces_rows:interfaces,
                options:temp
            })
        }
    }

    componentDidMount(){
        document.title = "Routes | SysMon"
        this.updateRoute()
        this.updateRules()
        this.updateTables()
        this.updateInterfaces()
        // For delrule
        const payload2=new GeneralRequest()
        payload2.destination="192.168.56.11";
        payload2.intermediate="192.168.122.1";
        payload2.interfaceName="virbr0";
        payload2.tableName="default";
        // this.delRoute(payload2)
        //================================
        // For delrule
        let payload4:GeneralRequest=new GeneralRequest();
        payload4.sourceIp="192.168.56.11";
        payload4.tableName="default";
        // this.delRule(payload4)
        //=================================
        // const interfaceName="wlp2s0"
        // this.interfaceByName(interfaceName)
        //==================================
    }
    
    handleModalToggle () {
        this.setState(({ isModalOpen }) => ({
          isModalOpen: !isModalOpen
        }));
    } 
      
    onSelect(event : React.FormEvent<HTMLInputElement>, isSelected : boolean, rowIndex : number,rowData: IRowData,extraData: IExtraData) {
        console.log(rowIndex)
        isSelected=true
	}

	toggleSelect(checked : boolean) {
		this.setState({
		  canSelectAll: checked
		});
	} 
    async handleChange(value: string, event: React.FormEvent<HTMLInputElement>) {
        const id = event.currentTarget.id;
        const newValue = event.currentTarget.value;
        if (id == "simple-form-src-ip-01") {
            await this.setState({
              srcIP :newValue
            });
        }
        if (id == "simple-form-dst-ip-01") {
            await this.setState({
              dstIP :newValue
            });
        }
        if (id == "simple-form-gateway-01") {
            await this.setState({
              defaultGateway :newValue
            });
        }
        this.checkAllValue()
    }
    checkAllValue(){
        if (this.state.interfaceName!="null"){
            if (this.state.srcIP.length>0 && this.state.dstIP.length>0 && this.state.defaultGateway.length>0 && this.state.interfaceName.length>0){
                this.setState({
                    isDisabledConfirm:false
                })
            }else{
                this.setState({
                    isDisabledConfirm:true
                })
            }
        }else{
            this.setState({
                isDisabledConfirm:true
            })
        }
    }         
    handleSubmit(event: React.FormEvent<HTMLButtonElement>) {
        let payload5:GeneralRequest=new GeneralRequest();
        payload5.destination=this.state.dstIP;
        payload5.intermediate=this.state.defaultGateway;
        payload5.interfaceName=this.state.interfaceName;
        payload5.sourceIp=this.state.srcIP;

        this.addTable(payload5)
        this.updateTables()
        this.updateRoute()
        this.updateRules()
        this.updateInterfaces()
        this.handleModalToggle()
        this.setState({
            srcIP :'',
            dstIP:'',
            defaultGateway:'',
            interfaceName:''
        });
        event.preventDefault();
    }

    async onChange(value: string, event: React.FormEvent<HTMLSelectElement>){
        await this.setState({
            interfaceName:value
        })
        this.checkAllValue()
    }
    render() {
        console.log('Rendering routes page ...')
        const { srcIP, dstIP, defaultGateway, interfaceName, isModalOpen, canSelectAll, columns_route, columns_rule, columns, rows_route, rows_rule, rows,interfaces_rows,interfaces_columns } = this.state;
	  
		return ( 	
            <div>    		
                <Stack>
                    <PageSection variant={PageSectionVariants.light}>
                        <Card>
                            <CardHeader>
                                <CardTitle>Routes Information</CardTitle>
                                <CardActions>
                                    <Button variant="link" icon={<PlusIcon />} onClick={this.handleModalToggle}>
                                    </Button>
                                    <Button variant="link" icon={<RedoIcon />} onClick={this.updateRoute}>
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
                                <CardTitle>Rules Information</CardTitle>
                                <CardActions>
                                    <Button variant="link" icon={<RedoIcon />} onClick={this.updateRules}>
                                    </Button>
                                </CardActions>
                            </CardHeader>
                            <CardBody>
                                <Table
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
                                <CardTitle>IP Tables</CardTitle>
                                <CardActions>
                                    <Button variant="link" icon={<RedoIcon />} onClick={this.updateTables}>
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

                    <PageSection variant={PageSectionVariants.light}>
                        <Card>
                            <CardHeader>
                                <CardTitle>Interfaces List</CardTitle>
                                <CardActions>
                                    <Button variant="link" icon={<RedoIcon />} onClick={this.updateInterfaces}>
                                    </Button>
                                </CardActions>
                            </CardHeader>
                            <CardBody>
                            <Table
                                aria-label="Selectable Table"
                                cells={interfaces_columns}
                                rows={interfaces_rows}>
                                <TableHeader />
                                <TableBody />
                                </Table>
                            </CardBody>
                        </Card>
                    </PageSection>        
                </Stack>
                <Modal
                    title="Add new routing policy"
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

                        <FormGroup label="Interface Name" fieldId="horizontal-form-title">
                            <FormSelect
                                value={this.state.interfaceName}
                                onChange={this.onChange}
                                id="horzontal-form-title"
                                name="horizontal-form-title"
                                aria-label="Your title"
                            >
                                {this.state.options.map((option, index) => (
                                <FormSelectOption isDisabled={option.disabled} key={index} value={option.value} label={option.label} />
                                ))}
                            </FormSelect>
                        </FormGroup>

                        <ActionGroup>
                            <Button key="confirm" variant="primary" onClick={this.handleSubmit} isDisabled={this.state.isDisabledConfirm}>
                                Confirm</Button>
                            <Button key="cancel" variant="link" onClick={this.handleModalToggle}>
                                Cancel</Button>
                        </ActionGroup>
                    </Form>
                </Modal>
            </div> 
        );
    }
}