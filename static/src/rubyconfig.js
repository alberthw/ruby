function loadConfigFromServer(url) {
    var result = {
        id: null,
        line: "",
        speed: null
    };
    $.ajax({
        url: url,
        dataType: "json",
        cache: false,
        async: false,
        success: function (data) {
            result.id = data.Id;
            result.line = data.Serialline;
            result.speed = data.Serialspeed;
        },
        error: function (xhr, status, err) {
            console.error(this.props.url, status, err.toString());
        }
    });
    return result;
}

function postConfigToServer(url, data) {
    var result = false;
    $.ajax({
        url: url,
        dataType: "json",
        type: "POST",
        async: false,
        data: data,
        success: function (data) {
            result = true;
        },
        error: function (xhr, status, err) {
            console.error(url, status, err.toString());
        }
    });
    return result;
}

var ConfigForm = React.createClass({
    getInitialState: function () {
       return loadConfigFromServer(this.props.url);
    },
    handleSerialLineChange: function (e) {
        this.setState({
            line: e.target.value
        });
    },
    handleSerialSpeedChange: function (e) {
        this.setState({
            speed: e.target.value
        });

    },
    handleSubmit: function (e) {
        e.preventDefault();
        var d = this.state.id;
        var l = this.state.line.trim();
        var s = this.state.speed;
        postConfigToServer(this.props.url, { id: d, line: l, speed: s });
        this.getInitialState();
    },
    render: function () {
        return (
            <form className="configForm" onSubmit={this.handleSubmit}>
                <span>Serial Line: </span><input type="text" value={this.state.line} onChange={this.handleSerialLineChange}/>
                <span>Serial Speed: </span><input type="text" value={this.state.speed} onChange={this.handleSerialSpeedChange}/>
                <input type="submit" value="Update" />
            </form>
        )
    }
});

var ConfigBox = React.createClass({
    render: function () {
        return (
            <div>
                <ConfigForm url="/config" />
            </div>
        )

    }
});


ReactDOM.render(<ConfigBox />, document.getElementById("rubyconfig"));

/*

var configURL = "/config";

var SerialLine = React.createClass({
    getInitialState:function(){
        return {
            value:this.props.data
        }
    },
    dataOnChange:function(e){
        this.setState({
            value:e.target.value
        });
    },
    render:function(){
        return(
            <input type="text" id="serialline" placeholder="Serial Line" onChange={this.dataOnChange}></input>
        )
    }
});

var SerialSpeed = React.createClass({
    getInitialState:function(){
        return {
            value:this.props.data
        }
    },
    render:function(){
        return(
            <input type="text" id="serialspeed" placeholder="Serial Speed"></input>
        )
    }
});

var ConfigForm = React.createClass({
    render:function(){
        return(
            <form>
            Serial Line:<SerialLine data="test"/>
            Serial Speed:<SerialSpeed />
            </form>
        )
    }
});

ReactDOM.render(<ConfigForm />, document.getElementById("rubyconfig"));





var ConfigForm = React.createClass({
    getInitialState:function(){
        return {
            serialID : this.props.data.Id,
            serialLine:this.props.data.Serialline,
            serialSpeed:this.props.data.Serialspeed
        }
    },
    handleSerialIDChange:function(e){
        this.setState({
            value:e.target.value
        });
    },
    handleSerialLineChange:function(e){
        this.setState({
            value:e.target.value
        });
    },
    handleSerialSpeedChange:function(e){
        this.setState({
            serialSpeed:e.target.value
        });
    },
    handleSubmit:function(e){
        e.preventDefault();
  //      var id = this.state.data.Id;
        var line = this.state.data.Serialline.trim();
        var speed = this.state.data.Serialspeed;
        if (!line || !speed){
            return ;
        }
        this.props.onConfigSubmit({
            serialLine:line,
            serialSpeed:speed
        });
        this.setState({
            serialLine:"", 
            serialSpeed:""
        });
    },
    render : function(){
        return (
            <form className="configFrom" onSubmit={this.handleSubmit}>
                <a id="serialID" value={this.props.data.Id || "" } onChange={this.handleSerialIDChange} hidden>{this.props.data.Id}</a>
                <a>Serial Line:</a><input type="text" id="serialLine" placeholder="Serial Line" value={this.props.data.Serialline || ""} onChange={this.handleSerialLineChange}></input>
                <a>Serial Speed:</a><input type="text" id="serialSpeed" placeholder="Serial Speed" value={this.props.data.Serialspeed || ""} onChange={this.handleSerialSpeedChange}></input>
                <input type="submit" value="Update"></input>
            </form>
        );
    },
});

var ConfigBox = React.createClass({
    getInitialState:function(){
        this.loadConfigFromServer();
        return {data:[]};
    },
    loadConfigFromServer:function(){
    $.ajax({
            url:this.props.url,
            dataType:"json",
            cache:false,
            success:function(data){
                this.setState({data:data});
            }.bind(this),
            error:function(xhr, status, err){ 
                console.error(this.props.url, status, err.toString());
            }.bind(this)
        });
    },
    handleConfigSubmit:function(data){
        alert("line:" + data.serialLine + " speed:" + data.serialSpeed);
    },
    render : function(){
        return (
            <div className="configBox">
            <h1>Config</h1>
            <ConfigForm onConfigSubmit={this.handleConfigSubmit} data={this.state.data}></ConfigForm>
            </div>
        );

    }
});

ReactDOM.render(<ConfigBox url = {configURL} />, document.getElementById("rubyconfig"));
*/