
class ConfigStringItem extends React.Component {
    constructor(props) {
        super(props);
        this.handleTextChange = this.handleTextChange.bind(this);
    }

    handleTextChange(e) {
        this.props.onChange(e);
    }
    render() {
        return (
            <div className="col-md-4">
                <div className="input-group">
                    <span className="input-group-addon">{this.props.name}</span>
                    <input type="text" className="form-control" value={this.props.value} onChange={this.handleTextChange}></input>
                </div>
            </div>

        );
    }
}

class ConfigNumberItem extends React.Component {
    constructor(props) {
        super(props);
        this.handleNumberChange = this.handleNumberChange.bind(this);
    }

    handleNumberChange(e) {
        this.props.onChange(e);
    }
    render() {
        return (
            <div className="col-md-4">
                <div className="input-group">
                    <span className="input-group-addon">{this.props.name}</span>
                    <input type="number" className="form-control" value={this.props.value} onChange={this.handleNumberChange}></input>
                </div>
            </div>

        );
    }
}

class SystemConfiguration extends React.Component {
    constructor(props) {
        super(props);

        this.handleDeviceNameChange = this.handleDeviceNameChange.bind(this);
        this.handleSystemVersionChange = this.handleSystemVersionChange.bind(this);
        this.handleDeviceSKUChange = this.handleDeviceSKUChange.bind(this);
        this.handleSerialNumberChange = this.handleSerialNumberChange.bind(this);
        this.handleSoftwareBuildChange = this.handleSoftwareBuildChange.bind(this);
        this.handlePartNumberChange = this.handlePartNumberChange.bind(this);
        this.handleHardwareVersionChange = this.handleHardwareVersionChange.bind(this);
        this.handleCountryChange = this.handleCountryChange.bind(this);
        this.handleRegionChange = this.handleRegionChange.bind(this);

        this.handleUpdateButtonClick = this.handleUpdateButtonClick.bind(this);

        this.state = {
            id: "",
            deviceName: "",
            sysVersion: "",
            deviceSKU: "",
            serialNumber: "",
            softwareBuild: "",
            partNumber: "",
            hardwareVersion: "",
            country: "",
            region: "",
        };
    }
    handleDeviceNameChange(e) {
        this.setState({
            deviceName: e.target.value
        });
    }
    handleSystemVersionChange(e) {
        this.setState({
            sysVersion: e.target.value
        });
    }

    handleDeviceSKUChange(e) {
        this.setState({
            deviceSKU: e.target.value
        });
    }

    handleSerialNumberChange(e) {
        this.setState({
            serialNumber: e.target.value
        });
    }

    handleSoftwareBuildChange(e) {
        this.setState({
            softwareBuild: e.target.value
        });

    }

    handlePartNumberChange(e) {
        this.setState({
            partNumber: e.target.value
        });

    }

    handleHardwareVersionChange(e) {
        this.setState({
            hardwareVersion: e.target.value
        });

    }

    handleCountryChange(e) {
        this.setState({
            country: e.target.value
        });

    }

    handleRegionChange(e) {
        this.setState({
            region: e.target.value
        });
    }

    handleUpdateButtonClick(e) {
        var url = "/setsysconfig";
        console.log("state:", this.state);
        $.ajax({
            url: url,
            dataType: "json",
            type: "POST",
            cache: false,
            async: false,
            data: this.state,
            success: function (data) {
                console.log(data);
                //result = data;
            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });

    }

    getSysConfig() {
        var url = "/getsysconfig";
        var result = null;

        $.ajax({
            url: url,
            dataType: "json",
            type: "GET",
            cache: false,
            async: false,
            success: function (data) {
                //            console.log(data);
                result = data;

            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });
        return result;
    }

    componentDidMount() {
        var data = this.getSysConfig();
        console.log("system : ", data);
        this.setState({
            id: data.Id,
            deviceName: data.Devicename,
            sysVersion: data.Systemversion,
            deviceSKU: data.Devicesku,
            serialNumber: data.Serialnumber,
            softwareBuild: data.Softwarebuild,
            partNumber: data.Partnumber,
            hardwareVersion: data.Hardwareversion,
            country: data.Country,
            region: data.Region,
        });
    }


    render() {

        return (
            <div className="panel panel-default">
                <div className="panel-heading">System Configuration</div>
                <div className="panel-body">
                    <div className="container-fluid">
                        <div className="row">

                            <ConfigStringItem name="Device Name:" value={this.state.deviceName} onChange={this.handleDeviceNameChange} />
                            <ConfigStringItem name="System Version:" value={this.state.sysVersion} onChange={this.handleSystemVersionChange} />
                            <ConfigStringItem name="Device SKU:" value={this.state.deviceSKU} onChange={this.handleDeviceSKUChange} />
                            <ConfigStringItem name="Serial Number:" value={this.state.serialNumber} onChange={this.handleSerialNumberChange} />
                            <ConfigStringItem name="Software Build:" value={this.state.softwareBuild} onChange={this.handleSoftwareBuildChange} />
                            <ConfigStringItem name="Part Number:" value={this.state.partNumber} onChange={this.handlePartNumberChange} />
                            <ConfigStringItem name="Hardware Version:" value={this.state.hardwareVersion} onChange={this.handleHardwareVersionChange} />
                            <ConfigNumberItem name="Country:" value={this.state.country} onChange={this.handleCountryChange} />
                            <ConfigNumberItem name="Region:" value={this.state.region} onChange={this.handleRegionChange} />

                        </div>
                    </div>
                </div>
                <div className="panel-footer">
                    <input type="button" value="Update" onClick={this.handleUpdateButtonClick}></input>
                </div>
            </div>
        );
    }
}



class HardwareConfiguration extends React.Component {
    constructor(props) {
        super(props);

        this.handleNameChange = this.handleNameChange.bind(this);
        this.handlePartNumberChange = this.handlePartNumberChange.bind(this);
        this.handleRevisionChange = this.handleRevisionChange.bind(this);
        this.handleSerialNumberChange = this.handleSerialNumberChange.bind(this);

        this.handleUpdateButtonClick = this.handleUpdateButtonClick.bind(this);

        this.state = {
            id: "",
            name: "",
            partNumber: "",
            revision: "",
            serialNumber: "",
        };
    }

    handleNameChange(e) {
        this.setState({
            name: e.target.value
        });

    }


    handlePartNumberChange(e) {
        this.setState({
            partNumber: e.target.value
        });

    }

    handleRevisionChange(e) {
        this.setState({
            revision: e.target.value
        });

    }

    handleSerialNumberChange(e) {
        this.setState({
            serialNumber: e.target.value
        });

    }

    
    handleUpdateButtonClick(e) {
        var url = "/sethwconfig";
        console.log("state:", this.state);
        $.ajax({
            url: url,
            dataType: "json",
            type: "POST",
            cache: false,
            async: false,
            data: this.state,
            success: function (data) {
                console.log(data);
                //result = data;
            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });

    }





    getHwConfig() {
        var url = "/gethwconfig";
        var result = null;

        $.ajax({
            url: url,
            dataType: "json",
            type: "GET",
            cache: false,
            async: false,
            success: function (data) {
                //            console.log(data);
                result = data;

            },
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }
        });
        return result;
    }
    componentDidMount() {
        var data = this.getHwConfig();
        console.log("hardware : ", data);
        this.setState({
            id: data.Id,
            name: data.Name,
            partNumber: data.Partnumber,
            revision: data.Revision,
            serialNumber: data.Serialnumber,
        });
    }
    render() {
        return (
            <div className="panel panel-default">
                <div className="panel-heading">Hardware Configuration</div>
                <div className="panel-body">
                    <div className="container-fluid">
                        <div className="row">

                            <ConfigStringItem name="Name:" value={this.state.name} onChange={this.handleNameChange} />
                            <ConfigStringItem name="Part Number:" value={this.state.partNumber} onChange={this.handlePartNumberChange} />
                            <ConfigStringItem name="Revision:" value={this.state.revision} onChange={this.handleRevisionChange} />
                            <ConfigStringItem name="Serial Number:" value={this.state.serialNumber} onChange={this.handleSerialNumberChange} />

                        </div>
                    </div>
                </div>
                <div className="panel-footer">
                    <input type="button" value="Update" onClick={this.handleUpdateButtonClick}></input>
                </div>
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