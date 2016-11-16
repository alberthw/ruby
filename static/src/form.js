class NameForm extends React.Component {
    constructor(props) {
        super(props)

        this.state = {
            value: "",
            area: "Please write an essay about your favorite DOM element.",
            select: "coconut"
        };

        this.handleChange = this
            .handleChange
            .bind(this);

        this.handleAreaChange = this
            .handleAreaChange
            .bind(this);

        this.handleSelectChange = this
            .handleSelectChange
            .bind(this);

        this.handleSubmit = this
            .handleSubmit
            .bind(this);
    }

    handleChange(event) {
        this.setState({
            value: event
                .target
                .value
                .toUpperCase()
        });
    }

    handleAreaChange(event) {
        this.setState({area: event.target.value});
    }

    handleSelectChange(event) {
        this.setState({select: event.target.value});
    }

    handleSubmit(event) {
        alert('A name was submitted : ' + this.state.value + " | " + this.state.area + " | " + this.state.select);
        event.preventDefault();
    }

    render() {
        return (
            <form onSubmit={this.handleSubmit}>
                Name :
                <input type="text" value={this.state.value} onChange={this.handleChange}/>
                <br/>
                <textarea value={this.state.area} onChange={this.handleAreaChange}/>
                <br/>
                <select value={this.state.select} onChange={this.handleSelectChange}>
                    <option value="grapefruit">Grapefruit</option>
                    <option value="lime">Lime</option>
                    <option value="coconut">Coconut</option>
                    <option value="mango">Mango</option>
                </select>
                <br/>
                <input type="submit" value="Submit"/>
            </form>
        );
    }

}

ReactDOM.render(
    <NameForm/>, document.getElementById("form"));