import * as React from 'react';
import { BrowserRouter as Router } from 'react-router-dom';
import { Route, Redirect, Switch } from 'react-router-dom';
import { NavLink } from 'react-router-dom';
import {
    Nav, NavItem, NavList,
    Page, PageHeader, PageSidebar
} from "@patternfly/react-core";
import { DashboardPage }     from "./pages/dashboard-page";
import { RoutesPage }        from "./pages/routes-page";
import { MetricsPage }       from "./pages/metrics-page";
import { SamplePage } from "./pages/sample-page";

interface AppState {
    sidebarOpen: boolean;
}

interface AppProps {
    // Empty
}

export class App extends React.Component<AppProps, AppState> {

    constructor(props: AppProps) {
        super(props);
        this.state = {
            sidebarOpen: true
        };
    }

    sidebarToggle = () => {
        this.setState({
            sidebarOpen: !this.state.sidebarOpen
        });
    }

    render() {

        const { sidebarOpen } = this.state;

        const Header = (
            <PageHeader
                logoComponent="span"
                logo="SYSMON"
                showNavToggle
                isNavOpen={sidebarOpen}
                onNavToggle={this.sidebarToggle}
            >
            </PageHeader>
        );

        const Navigation = (
            <Nav theme="dark">
                <NavList>
                    <NavItem>
                        <NavLink exact to='/web/dashboard' activeClassName="pf-m-current">Dashboard</NavLink>
                    </NavItem>
                    <NavItem>
                        <NavLink exact to='/web/routes' activeClassName="pf-m-current">Routes</NavLink>
                    </NavItem>
                    <NavItem>
                        <NavLink exact to='/web/metrics' activeClassName="pf-m-current">Metrics</NavLink>
                    </NavItem>
                    <NavItem>
                        <NavLink exact to='/web/sample' activeClassName="pf-m-current">Sample</NavLink>
                    </NavItem>
                </NavList>
            </Nav>
        );

        const Sidebar = <PageSidebar nav={ Navigation } isNavOpen={ sidebarOpen } />;

        return (
            <Router>
                <Page header={ Header } sidebar={ Sidebar }>
                    <Switch>
                        <Route exact path='/web/dashboard'>
                            <DashboardPage> </DashboardPage>
                        </Route>
                        <Route exact path='/web/routes'>
                            <RoutesPage> </RoutesPage>
                        </Route>
                        <Route exact path='/web/metrics'>
                            <MetricsPage> </MetricsPage>
                        </Route>
                        <Route exact path='/web/sample'>
                            <SamplePage> </SamplePage>
                        </Route>
                        <Route exact path='/'>
                            <Redirect to='/web/dashboard'> </Redirect>
                        </Route>
                    </Switch>
                </Page>
            </Router>
        );
    }
}