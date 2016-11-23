class Message extends React.Component {
    constructor(props) {
        super(props);

        this.handleDeviceNameChange = this.handleDeviceNameChange.bind(this);
        this.handleSessionRequestClick = this.handleSessionRequestClick.bind(this);
        this.handleRequestChange = this.handleRequestChange.bind(this);
        this.handleResponseChange = this.handleResponseChange.bind(this);

        this.state = {
            deviceName: "",
            request: "",
            response: "",
        };
    }

    componentDidMount() {
        this.timer = setInterval(() => { this.getLatestResponse("/response") }, 1000);
    }

    componentWillUnmount() {
        clearInterval(this.timer);
    }

    getLatestResponse(url) {
        $.ajax({
            url: url,
            dataType: "json",
            cache: false,
            async: false,
            success: function (data) {
                //         console.log(data);
                this.setState({
                    response: data.Info,
                    Id: data.Id
                });
            }.bind(this),
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }.bind(this)
        });


    }

    handleDeviceNameChange(e) {
        this.setState({
            deviceName: e.target.value
        });
    }

    handleSessionRequestClick(e) {
        const data = this.state;
        var url = "/generate";
        $.ajax({
            url: url,
            dataType: "json",
            type: "POST",
            async: false,
            data: {type:2},
            success: function (data) {
                console.log(data);
                this.setState({
                    request: data
                });

            }.bind(this),
            error: function (xhr, status, err) {
                console.error(url, status, err.toString());
            }.bind(this)
        });

    }

    handleRequestChange(e) {
        this.setState({
            request: e.target.value
        });
    }

    handleResponseChange(e) {
        this.setState({
            response: e.target.value
        });
    }


    render() {

        const deviceName = this.state.deviceName;
        const request = this.state.request;
        const response = this.state.response;

        return (

            <div>
                <div>
                    <input type="button" value="SessionRequest" onClick={this.handleSessionRequestClick}></input>
                </div>
                <div>
                    <h2>Request Message:</h2>
                    <textarea cols="200" rows="4" value={request} onChange={this.handleRequestChange}></textarea>
                    <h2>Response Message:</h2>
                    <textarea cols="200" value={response} onChange={this.handleResponseChange} readOnly></textarea>
                </div>
                <div>
                    <a>Device Name:</a>
                    <input type="text" value={deviceName} onChange={this.handleDeviceNameChange} readOnly></input>

                </div>

            </div>
        );

    }

}

ReactDOM.render(
    <Message />,
    document.getElementById("message")
);

