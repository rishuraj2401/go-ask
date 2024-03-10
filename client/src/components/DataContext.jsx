import React, { createContext, useContext, useState } from 'react';
import axios from 'axios';
import { GoogleAuthProvider, onAuthStateChanged, signInWithPopup, signOut } from 'firebase/auth';
import { auth } from '../firebase';
import { useEffect } from 'react';
export const DataContext = createContext();

const DataProvider = ({ children }) => {
  const [index, setIndex] = useState(0);
  const [searchresult, setResult] = useState([]);
  const [data, setData] = useState([]);
  const [currentPage, setPage] = useState(1);
  const [limit, setLimit] = useState(8);
  const [loader, setLoader] = useState(false);
  const [searchInput, setInput] = useState('');
  const [flag, setFlag] = useState(false);
  const ApiUrl = "http://localhost:8089/";
  const ApiSearch = "https://api.app.creatosaurus.io/creatosaurus/adminpanel/users/search";
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [rowData, setRowData] = useState({})
  const [edit, setEdit] = useState({ editE: "", editF: "", editL: "", editA: false });
  const [question, setQuestion] = useState('');
  const [questions, setQuestions] = useState([]);
  const [answer, setAnswer] = useState("")
  const [user, setUser] = useState("");
  const [q, setQ] = useState("")
  // const [user1, setUser1]= useState("rishu");

  useEffect(() => {
    // Check if the user is already authenticated 
    const unsubscribe = onAuthStateChanged(auth, (user1) => {
      if (user1) {
        // User is signed in
        // Additional logic, if needed
        setUser(user1)
        console.log('Automatic login:', user.displayName);
      } else {
        // No user is signed in
        console.log('User not signed in');
      }
    });

    // Clean up the subscription on component unmount
    return () => unsubscribe();
  }, []);


  useEffect(() => {
    if (user) postUser()
  }, [user])
  const postUser = async () => {
    try {
      const response = await axios.post('http://localhost:8089/user', { _id: user.uid, email: user.email, name: user.displayName }, {
        headers: {
          'Content-Type': 'application/json',
        },
      });
      console.log(response);
    } catch (error) {
      console.log(error);
    }
  }

  const loginHandler = async () => {
    try {
      const provider = new GoogleAuthProvider();
      const { user } = await signInWithPopup(auth, provider);
      setUser(user)
      console.log({
        name: user.displayName,
        email: user.email,
        _id: user.uid,
        photo: user.photoURL

      }, user);
    }
    catch (error) {
      console.log("error", error);
    }
  }

  const logoutHandler = () => {
    signOut(auth).then(() => setUser(""))
      .catch((err) => console.log(err))
  }
  const handleLike = (q, a) => {
    if (user) {
      console.log(`http://localhost:8089/${q}/${a}/${user}`);
      axios.get(`http://localhost:8089/like/${q}/${a}/${user}`)
        .then(() => setFlag(!flag))
        .catch(() => console.log("error in liking"))
    }
    else {
      loginHandler()
    }
  }
  const handlePostQuestion = async () => {
    try {
      const response = await axios.post('http://localhost:8089/insert', { questions: question, answer: [] }, {
        headers: {
          'Content-Type': 'application/json',
        },
      });
      console.log('Question posted successfully:', response.data);
      alert("Posted succesfully");
    } catch (error) {
      console.error('Error posting question:', error.message);
      // Handle errors, e.g., show an error message to the user
    }
  };
  const handleAnswer = async (id) => {
    if (user) {
      const endpoint = `http://localhost:8089/ans/${id}/${user.email}`;
      axios.put(endpoint, { answered: answer })
        .then(response => {
          console.log("Answer updated successfully:", response.data);
          // Handle any further actions after successful update
          handleCloseModal()
          setFlag(!flag)
        })
        .catch(error => {
          console.error('Error updating answer:', error);
          // Handle error cases
        });
    }
    else {
      loginHandler();
    }

  }
  const fetchQ = () => {
    const apiUrl = `http://localhost:8089/search/${currentPage}`;

    // Make a GET request to the backend
    axios.get(apiUrl,
      { params: { q: question } })
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
  }
  const handleOpenModal = () => {
    setIsModalOpen(true);
  };

  const handleCloseModal = () => {
    // e.preventDefault()
    setIsModalOpen(false);
  };
  const fetchSearch = async () => {
    try {
      setLoader(true);
      const response = await axios.get(`${ApiUrl}?page=${currentPage}&limit=${limit}&query=${searchInput}`);
      setData(response.data);
      console.log("fetch search is executing", currentPage);
    } catch (error) {
      console.log("Fetching error:", error);
    } finally {
      setLoader(false);
    }
  };

  const fetchData = async () => {
    try {
      setLoader(true);
      setFlag(false);
      const response = await axios.get(`${ApiUrl}?page=${currentPage}&limit=${limit}`);
      setData(response.data);
    } catch (error) {
      console.log("Fetching error:", error);
    } finally {
      setLoader(false);
    }
  };
  const handleEdit = (rowData) => {
    handleOpenModal();
    // setRowData(rowData);
    setEdit({ editE: rowData.email, editF: rowData.firstName, editL: rowData.lastName, editA: rowData.isAdmin });
    // setEditE(rowData.email);
    // setEditF(rowData.firstName);
    // setEditL(rowData.lastName);
    console.log('Editing:', rowData);
  };
  const handleOnChange = (e) => {
    e.preventDefault()
    setEdit({ ...edit, [e.target.name]: e.target.value })
  }

  const contextValue = {
    index,
    setIndex,
    searchresult,
    setResult,
    data,
    setData,
    currentPage,
    setPage,
    limit,
    setLimit,
    loader,
    setLoader,
    searchInput,
    setInput,
    flag,
    setFlag,
    fetchSearch,
    fetchData,
    ApiSearch,
    ApiUrl,
    isModalOpen,
    setIsModalOpen,
    handleOpenModal,
    handleCloseModal,
    rowData,
    setRowData,
    handleEdit,
    edit,
    setEdit,
    handleOnChange,
    handlePostQuestion,
    question,
    questions,
    setQuestion,
    setQuestions,
    fetchQ,
    handleAnswer,
    user, setUser,
    handleLike,
    loginHandler,
    logoutHandler,
    answer, setAnswer,
    q, setQ
    // editA,editF,editE,editL,
    // setEditA,setEditE,setEditF,setEditL
  };

  return (
    <DataContext.Provider value={contextValue}>
      {children}
    </DataContext.Provider>
  );
};

// const useDataContext = () => {
//   return useContext(DataContext);
// };

export { DataProvider };