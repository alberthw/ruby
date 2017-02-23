import React from 'react';
import {Row, Col, Card} from 'antd';
import SystemConfiguration from './sysconfig';
import HardwareConfiguration from "./hwconfig";
import SoftwareConfiguration from "./swconfig";
import ValidateConfig from "./validateconfig";

export default class DeviceConfig extends React.Component {
    render() {
        return (
            <Card title="Device Configuration">
                <ValidateConfig />
                <Row gutter={8}>
                    <Col span={12}>
                        <SystemConfiguration/>
                    </Col>
                    <Col span={12}>
                        <HardwareConfiguration/>
                    </Col>
                </Row>
                <Row>
                    <SoftwareConfiguration/>
                </Row>
            </Card>

        );
    }
}