import { useParams } from 'react-router-dom';
import './App.css';

function Course() {
  const { id } = useParams();
  return (
    <>
      <div>
        {id}
      </div>
    </>
  );
}

export default Course;