import{ useContext, useEffect} from 'react';
// import Tabl from './Tabl';
import { DataContext } from './DataContext';
import MyModal from './MyModel';
import Navb from './Navb';
// import Navbar from './Navbar';


const Home = () => {
  const {
    data,
    currentPage,
    setPage,
    flag,
    questions, 
    fetchQ,
    user,
    handleLike,
    question, setQuestion, handlePostQuestion,
    isModalOpen, setIsModalOpen,
    q,setQ,
    loginHandler
  } = useContext(DataContext);
  useEffect(()=>{
    if(!question){
      fetchQ()
    }
  },[question])
useEffect(()=>{
  if(!question){
fetchQ();}
if(question.length>1){
  fetchQ();
}
},[currentPage,flag])
const handleOnChange=(e)=>{
e.preventDefault()
setPage(1);
setQuestion(e.target.value)
console.log(question);
}  
const handleModal=(e,qid)=>{
  e.preventDefault()
  if(user){
  setQ(qid)
setIsModalOpen(true)
console.log("this is questionId",q);}
else {
  loginHandler();
}

}
  return ( 
    <>
    <div className="nave"><Navb/></div>
   <div className="maine ">
    <div className="input">
      <input type="text" className="inp" placeholder='write your question...' onChange={(e)=>handleOnChange(e)} value={question}/>
      <button type="button" onClick={fetchQ}>Search</button>
      <button type="button" onClick={handlePostQuestion}>Post Question</button>
    </div>
    <div className="main1">
    <MyModal isOpen={isModalOpen} close={!isModalOpen}/>
      {       
       questions?.map((q)=><><div className="quest">
        <div className="quest1">
        <h3 className="qh">{q.Questions}</h3>
        <button type="button" onClick={(e)=>handleModal(e,q.ID)}>
          answer</button></div>
          <div className="quest2">
           { q.answer?(<>
           <div className="ans">
            <h5>Answers:</h5>
            <div className="ans0">{q.answer.map((a)=><><div className="ans1">
              <div className="ans2"><p1>{a.Answered}</p1></div>
              <div className="like"><p1>{a.Upvote} <button type='button'
              onClick={()=> handleLike(q.ID, a.ID)}>
                upvote</button></p1> <p1>{a.AnsBy}</p1>
                </div></div></>)}</div>
            
           </div>
           </>):(<></>)}
          </div>
        </div></>)
      }
    </div>
    <div className="next text-center">
    <button type='button' className='mx-10 px-4 py-1 bg-blue-700 rounded' style={{ opacity: `${currentPage === 1 ? '0.5' : '1'}` }} onClick={() => { setPage(currentPage - 1) }}
          disabled={currentPage === 1 }>Prev</button> {'<'}{currentPage}{'>'}
        <button type='button'  className='mx-10 bg-blue-700 px-4 py-1 rounded' style={{ opacity: `${data.length===0 ? '0.5' : '1'}` }} onClick={() => {
          setPage(currentPage + 1)
          console.log(currentPage);
        }} disabled={!questions }>Next</button>
    </div>
   </div>
    
    </>
  );
}

export default Home;