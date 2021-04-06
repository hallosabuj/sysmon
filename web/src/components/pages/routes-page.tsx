import * as React from "react";
import {
    PageSection,
    PageSectionVariants
} from "@patternfly/react-core";
import { InterfaceAddresses, InterfaceDetails, Interfaces, Routes, Rules, Tables,GeneralRequest,GeneralResponse } from "../../models/route-models";

interface RoutesPageState {
    // Empty
}

interface RoutesPageProps {
    // Empty
}

export class RoutesPage extends React.Component<RoutesPageProps, RoutesPageState> {
    constructor(props: RoutesPageProps) {
        super(props);
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

    render() {
        console.log('Rendering routes page ...')
        return (
            <PageSection variant={PageSectionVariants.light}> This is routes page</PageSection>
        );
    }
}