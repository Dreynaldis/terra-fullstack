import { BrowserRouter as Router, Route, Routes } from 'react-router-dom'
import LoginPage from './pages/LoginPage'
import Navbar from './component/Navbar'
import { Navigate } from 'react-router-dom'
import ProfilePage from './pages/ProfilePage'

function App() {


  return (
    <Router>
      <Navbar />
      <Routes>
        <Route path='/login' element={<LoginPage />} />
        <Route path='/profile' element={<ProfilePage />} />
        <Route path='*' element={<Navigate to='/login' />} />
      </Routes>
    </Router>
  )
}

export default App
