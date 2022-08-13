

const Banner = () => {
    return (  <section class="hero is-large-touch has-text-centered black has-text-white round">
    <div class="hero-body">
      <p class="title is-4 white">ratemyprofessors.help</p>
      <p class="subtitle is-5 white">
        A Rate My Professors Helper!
      </p>
    </div>
  </section>);
};
class Form extends React.Component {
  constructor(props) {
    super(props);
    this.state = {value: ''};

    this.handleSchool = this.handleSchool.bind(this);
    this.handleCourse = this.handleCourse.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }
  handleSchool(event){
    this.setState({school: event.target.value});
  }
  handleCourse(event){
    this.setState({course: event.target.value});
  }


  handleSubmit(event) {
    alert(this.state.school+" "+this.state.course);
    event.preventDefault();
  }

  render() {
    return( 
      <section class="columns is-mobile is-centered">
        <div class="column is-one-quarter-desktop is-one-half-tablet">
    <form class = "box" onSubmit={this.handleSubmit}>
       <select name="school" id="school" value={this.state.school} onChange={this.handleSchool}>
        <option class="select-school" value="" disabled selected>Select School</option>
        <option value="De Anza">De Anza</option>
        {/* <option value="Berkeley">Berkeley</option> */}
    </select><br></br>
    <input type="subject" name="subject" id="subject"value={this.state.course} onChange={this.handleCourse} placeholder="Format: Course and ID I.E CS 46B" required>
    </input><br></br>
    <input type="submit" id="search" value="Search!"></input>
    </form>
      </div>
      </section>
    );
  }
}


ReactDOM.render(<Banner />, document.getElementById("home"));
ReactDOM.render(<Form />, document.getElementById("form"));

