

const Banner = () => {
    return (  <section class="hero is-small has-text-centered is-dark has-text-white">
    <div class="hero-body">
      <p class="title">Rate My Courses</p>
      <p class="subtitle">
        A Rate My Professors Algorithm
      </p>
    </div>
  </section>);
};
const Form = () => {
                    return( 
                    <section class = "container">
                    <section class = "box card">
                       <select name="school" id="school">
                        <option class="select-school" value="" disabled selected>Select School</option>
                        <option value="De Anza">De Anza</option>
                        <option value="Berkeley">Berkeley</option>
                    </select>
                    <select name="subject" id="subject">
                        <option class="select-subject" value="" disabled selected>Select Subject</option>
                        <option value="GenEd">GenEd</option>
                        <option value="nothingatm">PlaceholderMajor</option>
                    </select>
                    <input type="submit" value="Register"></input>
                    </section>
                    </section>
                    );
};

ReactDOM.render(<Banner />, document.getElementById("home"));
ReactDOM.render(<Form />, document.getElementById("form"));
