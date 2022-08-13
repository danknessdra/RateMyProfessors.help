

const Banner = () => {
    return (  <section class="hero is-small has-text-centered black has-text-white round">
    <div class="hero-body">
      <p class="title white">ratemyprofessors.help</p>
      <p class="subtitle white">
        A Rate My Professors Helper!
      </p>
    </div>
  </section>);
};
const Form = () => {
                    return( 
                      <section class="columns is-mobile is-centered">
                        <div class="column is-one-quarter">
                    <section class = "box ">
                       <select name="school" id="school">
                        <option class="select-school" value="" disabled selected>Select School</option>
                        <option value="De Anza">De Anza</option>
                        {/* <option value="Berkeley">Berkeley</option> */}
                    </select><br></br>
                    <input type="subject" name="subject" id="subject" placeholder="Format: Course and ID I.E CS 46B" required>
                    </input><br></br>
                    <input type="submit" value="Search!"></input>
                    </section>
                      </div>
                      </section>
                    );
};

ReactDOM.render(<Banner />, document.getElementById("home"));
ReactDOM.render(<Form />, document.getElementById("form"));
