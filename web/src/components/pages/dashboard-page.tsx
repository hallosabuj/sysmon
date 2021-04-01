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
    constructor(props: DashboardPageProps) {
        super(props);
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
    componentDidMount(){
        document.title = "Dashboard | SysMon"
        this.doPing()
        this.doPong()
    }

    render() {
        console.log('Rendering dashboard page ...')
        return (
            <PageSection variant={PageSectionVariants.light}> This is dashboard page</PageSection>
        );
    }
}