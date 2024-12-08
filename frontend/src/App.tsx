import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import LoginPage from './pages/LoginPage'
import Navbar from './component/Navbar'

function App() {


  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path='/login' element={<LoginPage />} />
      </Routes>
    </Router>
  )
}

export default App
