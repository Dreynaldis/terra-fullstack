import { Link,Navigate, useNavigate } from "react-router-dom"
import Cookies from "js-cookie"
import { useState, useEffect } from "react"
import axios from "axios"

const Navbar = () => {
  const navigate = useNavigate()
  const [name, setName] = useState<string | null>(null)
  const [isLoggedIn, setIsLoggedIn] = useState<boolean>(false)


  const checkLogin = async () => {
    const token = Cookies.get('token')
    const username = Cookies.get('username')
    console.log("test");
    
    if (!token || !username) {
      setIsLoggedIn(false)
      setName(null)
      Cookies.remove('token')
      Cookies.remove('username')
      navigate('/login')
    } else {
      setName(username)
      try {
        await axios.get(`${import.meta.env.VITE_API_BASE_URL}/auth/check`,{headers :{
          "Authorization": `Bearer ${token}`
        }})

        setIsLoggedIn(true)

      } catch (error) {
        setIsLoggedIn(false)
        
      }
    }
  }
  const handleLogout = () => {
    Cookies.remove('token')
    Cookies.remove('username')
    setIsLoggedIn(false)
    setName(null)
    navigate('/login')

  }
  useEffect(() => {
    if (isLoggedIn) {
      navigate('/profile')
    } else {
      checkLogin()
    }
  }, [isLoggedIn])

  return (
    <>
    <div className="mb-5 h-20">
        <nav className="fixed py-2 flex w-full z-20 top-0 justify-center">
            <div className="flex w-3/4 max-w-screen-xl mx-auto items-center justify-between">
            <a href="/" className="flex items-center">
                <img className="" src="/logo.png" alt="Terra logo" />
            </a>
            {isLoggedIn ? (
              <div>

              <Link onClick={handleLogout} to="/login" className="ml-5 py-3 px-8 text-white bg-p1 rounded-3xl hover:bg-p2 transition-colors delay-100">
              Logout
              </Link>
              </div>
            ): (
            <div>
              <Link to="/login" className="ml-5 py-3 px-8 text-white bg-p1 rounded-3xl hover:bg-p2 transition-colors delay-100">
                Login
              </Link>
            </div>

            ) }
            </div>
        </nav>
    </div>
    </>
  )
}

export default Navbar