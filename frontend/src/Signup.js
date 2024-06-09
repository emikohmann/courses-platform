import React, { useState } from 'react';
import './App.css';
import Header from './Header';

function Signup() {
  const [loading, setLoading] = useState(false);
  const [result, setResult] = useState("");

  const handleSignupSubmit = (e) => {
    e.preventDefault(); // Prevent the default form submission
    
    let signupRequest = {
      email: e.target["0"].value,
      password: e.target["1"].value,
      type: e.target["2"].value
    };

    setLoading(true);
    fetch('http://localhost:8080/users/signup', {
      method: 'POST',
      body: JSON.stringify( signupRequest )
    }).then(response => response.json())
      .then(data => setResult(data.message))
      .then(() => setLoading(false))
      .catch(error => console.error('Error registering user:', error));
  };

  return (
    <>
      <Header />
      {
        result ? (
          <p>{result}</p>
        ): (
          loading ? (
            <img style={{width: 36 + "px"}} src="https://raw.githubusercontent.com/Codelessly/FlutterLoadingGIFs/master/packages/circular_progress_indicator_square_large.gif" />
          ) : (
            <div className="Signup">
              <form onSubmit={handleSignupSubmit}>
                <p>
                  <input
                    type="email"
                    placeholder="Email"
                  />
                </p>
                <p>
                  <input
                    type="password"
                    placeholder="Password"
                  />
                </p>
                <p>
                  <select>
                    <option value="normal">Normal</option>
                    <option value="admin">Admin</option>
                  </select>
                </p>
                <p>
                  <input
                    type="submit"
                    value="Register"
                  />
                </p>
              </form>
            </div>
          )
        )
      }
    </>
  );
}

export default Signup;