

const Banner = () => {
    return (  <section class="hero is-large-touch has-text-centered black has-text-white round">
    <div class="hero-body">
      <p class="title is-4 white">ratemyprofessors.help</p>
      <p class="subtitle is-5 white">
        A Rate My Professors Helper! beta
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
    let data = {
        School: this.state.school,
        Course: this.state.course,
    };
    fetch("/get_rank", {
        headers: {
            'Accept': 'application/json',
            'Content-Type': 'application/json'
        },
        method: "POST",
        body: JSON.stringify(data)
    }).then((response) => {
        response.text().then(function (data) {
            let result = JSON.parse(data);
            console.log(result);
            let title = result[0].course;
            let htmlString = '';
            for (let i = 0; i<result.length;++i) {
                htmlString+= `<section class = "box">
                <div class="newResult">
                  <div class = "title is-5">${result[i].prof}</div>
                  <div>
                    <div class = "title is-2">${result[i].rating}</div>
                    <div class = "subtitle is-6">${result[i].size} Ratings</div>
                  </div>
                </div>
                </div>
                </section>`
            }
            Swal.fire({
              title: "<strong>"+result[0].course+"<\strong>",
              showClass: {
                popup: 'animate__animated animate__fadeInDown'
              },
              hideClass: {
                popup: 'animate__animated animate__fadeOutUp'
              },
              html: htmlString,
              showCloseButton: true,
              focusConfirm: false,
              confirmButtonColor: '#5186ef',
              confirmButtonText:
                'Ok!',
              confirmButtonAriaLabel: 'Thumbs up, great!',
            })
        });
    }).catch((error) => {
        console.log(error)
    });
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

