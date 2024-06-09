import { useParams } from 'react-router-dom';
import React, { useEffect, useState } from 'react';
import './App.css';
import Header from './Header';

function Course() {
  const { id } = useParams();
  const [message, setMessage] = useState("");
  const [course, setCourse] = useState({});

  useEffect(() => {
    if (!id) {
      return;
    }

    const fetchCourse = async () => {
      try {
        console.log(`Fetching data from http://localhost:8080/courses/${id}`);
        const response = await fetch(`http://localhost:8080/courses/${id}`);
        const data = await response.json();

        if (response.status !== 200) {
          setMessage(data.message);
        } else {
          setCourse(data);
        }
      } catch (error) {
        console.error('Error fetching course:', error);
      }
    };

    fetchCourse();
  }, [id]);

  return (
    <>
      <Header />
      <div>
        {message ? (
          <p>{message}</p>
        ) : ( Object.keys(course).length > 0 ? (
          <>
            <div className="course-details">
              <p><strong>{course.category}</strong></p>
              <p>{course.title}</p>
              <img className="course-image" src={course.image_url} alt={course.title} />
              <p>{course.description}</p>
              <br />
              <form><input type="submit" value="Subscribe" /></form>
            </div>
          </>
        ) : (
          <img style={{width: 36 + "px"}} src="https://raw.githubusercontent.com/Codelessly/FlutterLoadingGIFs/master/packages/circular_progress_indicator_square_large.gif" />
        ))}
      </div>
    </>
  );
}

export default Course;