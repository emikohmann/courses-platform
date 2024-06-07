import './App.css';

function Signup() {

  const handleSignupSubmit = (e) => {
    e.preventDefault(); // Prevent the default form submission
    
    let signupRequest =  {
      email: e.target["0"].value,
      password: e.target["1"].value,
      type: e.target["2"].value
    };

    fetch('http://localhost:8080/users/signup', {
      method: 'POST',
      body: JSON.stringify( signupRequest )
    })
      .then(response => response.json())
      .catch(error => console.error('Error registering user:', error));
    
    console.log(signupRequest);
  };

  return (
    <>
      <div className="Signup">
        <form onSubmit={handleSignupSubmit}>
          <input
            type="email"
            placeholder="Email"
          />
          <input
            type="password"
            placeholder="Password"
          />
          <input
            type="text"
            placeholder="Type"
          />
          <input 
            type="submit"
            value="Register"
          />
        </form>
      </div>
    </>
  );
}

export default Signup;