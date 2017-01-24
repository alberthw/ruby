class SystemConfiguration extends React.Component {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <div className="panel panel-default">
                <div className="panel-heading">System Configuration</div>
                <div className="panel-body">
                    <div className="container-fluid">
                        <div className="row">

                            <div className="col-md-3">
                                <div className="input-group">
                                    <span className="input-group-addon">Device Name:</span>
                                    <input type="text" className="form-control"></input>
                                </div>
                            </div>

                            <div className="col-md-3">
                                <div className="input-group">
                                    <span className="input-group-addon">System Version:</span>
                                    <input type="text" className="form-control"></input>
                                </div>
                            </div>

                            <div className="col-md-3">
                                <div className="input-group">
                                    <span className="input-group-addon">Device SKU:</span>
                                    <input type="text" className="form-control"></input>
                                </div>
                            </div>

                            <div className="col-md-3">
                                <div className="input-group">
                                    <span className="input-group-addon">Serial Number:</span>
                                    <input type="text" className="form-control"></input>
                                </div>
                            </div>

                            <div className="col-md-3">
                                <div className="input-group">
                                    <span className="input-group-addon">Software Build:</span>
                                    <input type="text" className="form-control"></input>
                                </div>
                            </div>

                            <div className="col-md-3">
                                <div className="input-group">
                                    <span className="input-group-addon">Part Number:</span>
                                    <input type="text" className="form-control"></input>
                                </div>
                            </div>

                            <div className="col-md-3">
                                <div className="input-group">
                                    <span className="input-group-addon">hardware Version:</span>
                                    <input type="text" className="form-control"></input>
                                </div>
                            </div>

                            <div className="col-md-3">
                                <div className="input-group">
                                    <span className="input-group-addon">Country:</span>
                                    <input type="text" className="form-control"></input>
                                </div>
                            </div>

                            <div className="col-md-3">
                                <div className="input-group">
                                    <span className="input-group-addon">Region:</span>
                                    <input type="text" className="form-control"></input>
                                </div>
                            </div>

                        </div>
                    </div>
                </div>

            </div>
        );
    }
}

class HardwareConfiguration extends React.Component {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <div className="panel panel-default">
                <div className="panel-heading">Hardware Configuration</div>
                <div className="panel-body">Hardware Configuration</div>
            </div>
        );
    }
}

class SoftwareConfiguration extends React.Component {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <div className="panel panel-default">
                <div className="panel-heading">Software Configuration</div>
                <div className="panel-body">
                    <HostBootConfiguration />
                    <HostAppConfiguration />
                    <DspAppConfiguration />
                </div>
            </div>
        );
    }
}

class HostBootConfiguration extends React.Component {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <div className="panel panel-default">
                <div className="panel-heading">Host Boot Loader Configuration</div>
                <div className="panel-body">Boot Loader Configuration</div>
            </div>
        );
    }
}

class HostAppConfiguration extends React.Component {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <div className="panel panel-default">
                <div className="panel-heading">Host Application Configuration</div>
                <div className="panel-body">Application Configuration</div>
            </div>
        );
    }
}

class DspAppConfiguration extends React.Component {
    constructor(props) {
        super(props);
    }
    render() {
        return (
            <div className="panel panel-default">
                <div className="panel-heading">DSP Application Configuration</div>
                <div className="panel-body">Application Configuration</div>
            </div>
        );
    }
}


class DeviceConfiguration extends React.Component {

    render() {
        return (
            <div className="panel panel-default panel-primary">
                <div className="panel-heading">Device Configuration</div>
                <div className="panel-body">
                    <SystemConfiguration />
                    <HardwareConfiguration />
                    <SoftwareConfiguration />
                </div>
            </div>
        );
    }

}

ReactDOM.render(<DeviceConfiguration />, document.getElementById("configuration"));