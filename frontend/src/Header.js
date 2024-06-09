import './App.css';

function Header() {
  return (
    <div className="header">
      <h1><a href="/">Courses Platform</a></h1>
      <ul>
        <li><a href="/courses">Courses</a></li>
        <li><a href="/signup">Sign up</a></li>
        <li><a href="/login">Login</a></li>
      </ul>
      <hr />
    </div>
  );
}

export default Header;