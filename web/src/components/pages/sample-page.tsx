import * as React from "react";
import {
    PageSection,
    PageSectionVariants
} from "@patternfly/react-core";

interface SamplePageState {
    // Empty
}

interface SamplePageProps {
    // Empty
}

export class SamplePage extends React.Component<SamplePageProps, SamplePageState> {
    constructor(props: SamplePageProps) {
        super(props);
    }

    componentDidMount(){
        document.title = "Sample | SysMon"
    }

    render() {
        console.log('Rendering sample page ...')
        return (
            <PageSection variant={PageSectionVariants.light}> This is sample page</PageSection>
        );
    }
}