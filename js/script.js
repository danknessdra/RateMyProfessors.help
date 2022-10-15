const Banner = () => {
  return /*#__PURE__*/React.createElement("section", {
    className: "hero is-large-touch has-text-centered black has-text-white round"
  }, /*#__PURE__*/React.createElement("div", {
    className: "hero-body"
  }, /*#__PURE__*/React.createElement("p", {
    className: "title is-4 white"
  }, "RateMyProfessors.help"), /*#__PURE__*/React.createElement("p", {
    className: "subtitle is-5 white"
  }, "A Rate My Professors Helper! (beta)")));
};

class Form extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      value: ''
    };
    this.handleSchool = this.handleSchool.bind(this);
    this.handleCourse = this.handleCourse.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleSchool(event) {
    this.setState({
      school: event.target.value
    });
  }

  handleCourse(event) {
    this.setState({
      course: event.target.value
    });
  }

  handleSubmit(event) {
    let data = {
      School: this.state.school,
      Course: this.state.course.toUpperCase()
    };
    fetch("/get_rank", {
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json'
      },
      method: "POST",
      body: JSON.stringify(data)
    }).then(response => {
      response.text().then(function (data) {
        let result = JSON.parse(data);
        // console.log(result);
        if (result===null) {
          Swal.fire({
            title: 'Invalid Input!',
            text: 'Please ensure that your course or school is in the correct format. If it is, contact @bobo#0007 on discord!',
            showClass: {
              popup: 'animate__animated animate__fadeInDown'
            },
            hideClass: {
              popup: 'animate__animated animate__fadeOutUp'
            },
            confirmButtonColor: '#5186ef',
            confirmButtonText: 'Ok!',
          })
        }
        else {
        let title = result[0].course;
        let htmlString = '';
        for (let i = 0; i < result.length; ++i) {
          htmlString += `<section class = "box">
                <div class="newResult">
                  <div class = "title is-5">${result[i].prof}</div>
                  <div>
                    <div class = "title is-2">${result[i].rating}</div>
                    <div class = "subtitle is-6">${result[i].size} Ratings</div>
                  </div>
                </div>
                </div>
                </section>`;
        }

        Swal.fire({
          title: "<strong>" + result[0].course + "<\strong>",
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
          confirmButtonText: 'Ok!',
          confirmButtonAriaLabel: 'Thumbs up, great!'
        });}
      });
    }).catch(error => {
      console.log(error);
    });
    
    event.preventDefault();
  }

  render() {
    return /*#__PURE__*/React.createElement("section", {
      className: "columns is-mobile is-centered"
    }, /*#__PURE__*/React.createElement("div", {
      className: "column is-one-quarter-desktop is-one-half-tablet"
    }, /*#__PURE__*/React.createElement("form", {
      className: "box",
      onSubmit: this.handleSubmit
    }, /*#__PURE__*/React.createElement("select", {
      name: "school",
      id: "school",
      value: this.state.school,
      onChange: this.handleSchool
    }, /*#__PURE__*/React.createElement("option", {
      className: "select-school",
      value: "",
      disabled: true,
      selected: true
    }, "Select School"), /*#__PURE__*/React.createElement("option", {
      value: "De Anza"
    }, "De Anza"), /*#__PURE__*/React.createElement("option", {
      value: "UC Berkeley"
    }, "UC Berkeley")), /*#__PURE__*/React.createElement("br", null), /*#__PURE__*/React.createElement("input", {
      type: "subject",
      name: "subject",
      id: "subject",
      value: this.state.course,
      onChange: this.handleCourse,
      placeholder: "Format: Course and ID I.E COMPSCI 61A",
      required: true
    }), /*#__PURE__*/React.createElement("br", null), /*#__PURE__*/React.createElement("input", {
      type: "submit",
      id: "search",
      value: "Search!"
    }))));
  }

}

const Footer = () => {
  return /*#__PURE__*/React.createElement("div", {
    className: "black footer1"
  }, /*#__PURE__*/React.createElement("div", {
    className: "content has-text-centered"
  }, /*#__PURE__*/React.createElement("p", null, "RateMyProfessors.help by Duy Nguyen and Raymond Lin. The source code is licensed under the MIT license.")));
}; // add container to root then render- updated for React 18


const render1 = document.getElementById("home");
const render2 = document.getElementById("form");
const render3 = document.getElementById("footer");
const root1 = ReactDOM.createRoot(render1);
const root2 = ReactDOM.createRoot(render2);
const root3 = ReactDOM.createRoot(render3);
root1.render( /*#__PURE__*/React.createElement(Banner, null));
root2.render( /*#__PURE__*/React.createElement(Form, null));
root3.render( /*#__PURE__*/React.createElement(Footer, null));