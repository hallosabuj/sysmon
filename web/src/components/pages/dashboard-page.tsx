import * as React from "react";
import {
    PageSection,
    PageSectionVariants
} from "@patternfly/react-core";

interface DashboardPageState {
    // Empty
}

interface DashboardPageProps {
    // Empty
}

export class DashboardPage extends React.Component<DashboardPageProps, DashboardPageState> {
    controller : AbortController|undefined
    timerId : any
    constructor(props: DashboardPageProps) {
        super(props);
    }

    componentDidMount(){
        document.title = "Dashboard | SysMon"
        this.controller = new AbortController()
        this.periodicPing()
        // this.doPing()
        // this.doPong()
    }
    componentWillUnmount(){
        if ( undefined !== this.controller) {
            this.controller.abort()
        }
        clearInterval(this.timerId)
    }

    render() {
        console.log('Rendering dashboard page ...')
        return (
            <PageSection variant={PageSectionVariants.light}> This is dashboard page</PageSection>
        );
    }
    async doPing(){
        let response=await fetch("/api/v1/ping")
        console.log(response.status)
        let body=await response.text()
        console.log(body)
    }
    async doPong(){
        let response=await fetch("/api/v1/pong",{
            method:"post",
            headers: {'Content-Type':'application/json'},
            body: JSON.stringify({
                name:"sabuj",
                age:111
            })
        })
        console.log(response.status)
        let body=await response.text()
        console.log(body)
    }
    async dong () {
        let request = await fetch ("/api/v1/dong/"+"hello-world",{signal:this.controller!.signal})
        if ( request.status === 200 ) {
            console.log("request.status ",request.status)
            let body = await request.text()
            console.log("Body ",body)
        } 
        else {
            console.log(request.status)
        }
    }
    async periodicPing () {
        this.timerId = setInterval(() => this.dong(), 2000)
    }
}