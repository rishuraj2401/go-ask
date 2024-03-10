// MyModal.js
import React, { useContext, useState } from 'react';
import Modal from 'react-modal';
import { DataContext } from './DataContext';

const MyModal = ({ isOpen, onRequestClose }) => {
  const {
    answer,setAnswer,
    handleAnswer,
    q, handleCloseModal
    
  }=useContext(DataContext);
//   const [inputValue, setInputValue] = useState('');
const handleA=(e)=>{
    e.preventDefault()
    console.log("this is again",q);
    handleAnswer(q)
}
  return (
    <div className="modal w-{80vw} h-{50vh}">
    <Modal
      isOpen={isOpen}
      onRequestClose={onRequestClose}
      contentLabel="My Modal"
      style={{
        content: {
          width: '55vw', // Adjust the width as needed
          height: '50vh', // Adjust the height as needed
          top: '50%', // Center the modal vertically
          left: '50%', // Center the modal horizontally
          transform: 'translate(-50%, -50%)', 
          backgroundColor:"white",
        },
      }}
    >
        <div className="inp1">
            <button type="button" onClick={handleCloseModal}>X</button>
            <textarea name="" id="" cols="30" rows="10" onChange={(e)=>setAnswer(e.target.value)} value={answer} >

            </textarea>
            <button type="submit" onClick={(e)=>handleA(e)}>Submit</button>
        </div>
    
    </Modal>
    </div>
  );
};

export default MyModal;