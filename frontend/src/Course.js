import { useParams } from 'react-router-dom';
import React, { useEffect, useState } from 'react';
import './App.css';

function Course() {
  const { id } = useParams();
  const [ course, setCourse ] = useState({});

  useEffect(() => {
    if (id == null) {
      return
    }
    console.log(`Fetching data from http://localhost:8080/courses/${id}`)
    // Fetch data from the API
    fetch(`http://localhost:8080/courses/${id}`)
      .then(response => response.json())
      .then(data => setCourse(data))
      .catch(error => console.error('Error fetching courses:', error));
  }, [course]);

  return (
    <>
      <div className="Courses">
        {course != null ? (
          <>
            <div key={course.id} className="Course">
                <>
                  <img src={course.image_url} alt={course.title} className="Course-image" />
                  <div className="Course-details">
                    <h1 className="Course-title"><a href={"/courses/" + course.id}>{course.title}</a></h1>
                    <p className="Course-description">{course.description}</p>
                    <p className="Course-category"><strong>{course.category}</strong></p>
                  </div>
                </>
            </div>
            <form><input type="submit" value="Subscribe"/></form>
          </>
        ) : (
          <p>Loading course...</p>
        )}
      </div>
    </>
  );
}

export default Course;