import  { useEffect, useState } from 'react';
import axios from 'axios';

const Home = () => {
    useEffect(() => {
      // Define the URL for your Go backend
      const apiUrl = `http://localhost:8089//getQ/${page}`;
  
      // Make a GET request to the backend
      axios.get(apiUrl)
        .then(response => {
          // Update state with the fetched questions
          setQuestions(response.data);
  
          // You can handle the data in other ways as needed
          console.log(response.data);
        })
        .catch(error => {
          // Handle error
          console.error('Error fetching questions:', error);
        });
    }, [page]);
    const handleQuestionChange = (e) => {
      setQuestion(e.target.value);
    };
  
    return (
      <div>
       
       <div className="nav"></div>
       <div className="main">
        <div className="search"></div>
        
       </div>


      </div>
    );
}

export default Home;
