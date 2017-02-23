import React from "react";
import {Row, Col} from "antd";
import SerialConfig from "./serialconfig";
import RepositorySetting from "./reposetting";
import ReleaseFilesTable from "./releasefiles";

export default class UpgradePage extends React.Component {
    render() {
        return (
            <div>
                <Row>
                    <Col span={8}>
                        <SerialConfig url="/config"/>
                    </Col>
                    <Col span={8}>
                        <RepositorySetting/>
                    </Col>
                </Row>
                <Row>
                    <ReleaseFilesTable/>
                </Row>

            </div>
        );
    }
}
