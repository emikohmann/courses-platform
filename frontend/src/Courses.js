import React, { useState, useEffect } from 'react';
import './App.css';
import Header from './Header';

function Courses() {

  const [query, setQuery] = useState('');
  const [search, setSearch] = useState('');
  const [courses, setCourses] = useState([]);

  const handleSearchChange = (e) => {
    setQuery(e.target.value); // Update the query state as the user types
  };

  const handleSearchSubmit = (e) => {
    e.preventDefault();
    setSearch(query);
  };

  useEffect(() => {
    setCourses([]);
    fetch(`http://localhost:8080/courses/search?query=${search}`)
      .then(response => response.json())
      .then(data => {
        setCourses(data.results)
      })
      .catch(error => console.error('Error fetching courses:', error));
  }, [search]);

  return (
    <>
      <Header />
      <form onSubmit={handleSearchSubmit}>
        <input
          type="text"
          placeholder="Search for courses..."
          value={query}
          onChange={handleSearchChange}
        />
      </form>
      <div>
        {courses.length > 0 ? (
          courses.map((course, idx) => (
            <>
              {idx > 0 ? (<hr />) : (<></>)}
              <div key={course.id} className="course-list-item">
                <img src={course.image_url} alt={course.title} />
                <div>
                  <p><strong>{course.category}</strong></p>
                  <p>{course.title}</p>
                  <p><a href={`/courses/${course.id}`}>More details</a></p>
                </div>
              </div>
            </>
          ))
        ) : (
          <img style={{width: 36 + "px"}} src="https://raw.githubusercontent.com/Codelessly/FlutterLoadingGIFs/master/packages/circular_progress_indicator_square_large.gif" />
        )}
      </div>
    </>
  );
}

export default Courses;