import React from "react";
import {Row, Col} from "antd";
import SerialConfig from "./serialconfig.jsx";
import RepositorySetting from "./reposetting.jsx";
import ReleaseFilesTable from "./releasefiles.jsx";

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
                <ReleaseFilesTable/>

            </div>
        );
    }
}
